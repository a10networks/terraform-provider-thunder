

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsTxtRecordStats struct {
    
    RecordName string `json:"record-name"`
    Stats GslbZoneServiceDnsTxtRecordStatsStats `json:"stats"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsTxtRecordStats struct {
    DtGslbZoneServiceDnsTxtRecordStats GslbZoneServiceDnsTxtRecordStats `json:"dns-txt-record"`
}


type GslbZoneServiceDnsTxtRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsTxtRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsTxtRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-txt-record/"+p.RecordName+"/stats"
}

func (p *GslbZoneServiceDnsTxtRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsTxtRecordStats,error) {
logger.Println("GslbZoneServiceDnsTxtRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsTxtRecordStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
