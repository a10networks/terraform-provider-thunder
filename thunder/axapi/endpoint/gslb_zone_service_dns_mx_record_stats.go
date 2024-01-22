

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsMxRecordStats struct {
    
    MxName string `json:"mx-name"`
    Stats GslbZoneServiceDnsMxRecordStatsStats `json:"stats"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsMxRecordStats struct {
    DtGslbZoneServiceDnsMxRecordStats GslbZoneServiceDnsMxRecordStats `json:"dns-mx-record"`
}


type GslbZoneServiceDnsMxRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsMxRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsMxRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-mx-record/"+p.MxName+"/stats"
}

func (p *GslbZoneServiceDnsMxRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsMxRecordStats,error) {
logger.Println("GslbZoneServiceDnsMxRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsMxRecordStats
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
