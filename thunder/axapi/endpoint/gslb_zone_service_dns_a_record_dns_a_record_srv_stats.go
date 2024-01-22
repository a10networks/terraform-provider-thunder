

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsARecordDnsARecordSrvStats struct {
    
    Stats GslbZoneServiceDnsARecordDnsARecordSrvStatsStats `json:"stats"`
    Svrname string `json:"svrname"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsARecordDnsARecordSrvStats struct {
    DtGslbZoneServiceDnsARecordDnsARecordSrvStats GslbZoneServiceDnsARecordDnsARecordSrvStats `json:"dns-a-record-srv"`
}


type GslbZoneServiceDnsARecordDnsARecordSrvStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsARecordDnsARecordSrvStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsARecordDnsARecordSrvStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-a-record/dns-a-record-srv/"+p.Svrname+"/stats"
}

func (p *GslbZoneServiceDnsARecordDnsARecordSrvStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsARecordDnsARecordSrvStats,error) {
logger.Println("GslbZoneServiceDnsARecordDnsARecordSrvStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsARecordDnsARecordSrvStats
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
