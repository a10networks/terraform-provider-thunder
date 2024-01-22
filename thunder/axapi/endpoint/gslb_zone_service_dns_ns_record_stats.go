

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsNsRecordStats struct {
    
    NsName string `json:"ns-name"`
    Stats GslbZoneServiceDnsNsRecordStatsStats `json:"stats"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsNsRecordStats struct {
    DtGslbZoneServiceDnsNsRecordStats GslbZoneServiceDnsNsRecordStats `json:"dns-ns-record"`
}


type GslbZoneServiceDnsNsRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsNsRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsNsRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-ns-record/"+p.NsName+"/stats"
}

func (p *GslbZoneServiceDnsNsRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsNsRecordStats,error) {
logger.Println("GslbZoneServiceDnsNsRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsNsRecordStats
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
