

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsCnameRecordStats struct {
    
    AliasName string `json:"alias-name"`
    Stats GslbZoneServiceDnsCnameRecordStatsStats `json:"stats"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsCnameRecordStats struct {
    DtGslbZoneServiceDnsCnameRecordStats GslbZoneServiceDnsCnameRecordStats `json:"dns-cname-record"`
}


type GslbZoneServiceDnsCnameRecordStatsStats struct {
    CnameHits int `json:"cname-hits"`
}

func (p *GslbZoneServiceDnsCnameRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsCnameRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-cname-record/"+p.AliasName+"/stats"
}

func (p *GslbZoneServiceDnsCnameRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsCnameRecordStats,error) {
logger.Println("GslbZoneServiceDnsCnameRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsCnameRecordStats
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
