

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsPtrRecordStats struct {
    
    PtrName string `json:"ptr-name"`
    Stats GslbZoneServiceDnsPtrRecordStatsStats `json:"stats"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsPtrRecordStats struct {
    DtGslbZoneServiceDnsPtrRecordStats GslbZoneServiceDnsPtrRecordStats `json:"dns-ptr-record"`
}


type GslbZoneServiceDnsPtrRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsPtrRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsPtrRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-ptr-record/"+p.PtrName+"/stats"
}

func (p *GslbZoneServiceDnsPtrRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsPtrRecordStats,error) {
logger.Println("GslbZoneServiceDnsPtrRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsPtrRecordStats
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
