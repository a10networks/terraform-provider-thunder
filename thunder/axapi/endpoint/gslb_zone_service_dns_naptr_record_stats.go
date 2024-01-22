

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsNaptrRecordStats struct {
    
    Flag string `json:"flag"`
    NaptrTarget string `json:"naptr-target"`
    ServiceProto string `json:"service-proto"`
    Stats GslbZoneServiceDnsNaptrRecordStatsStats `json:"stats"`
    ServicePort string 
    ServiceName string 
    Name string 

}
type DataGslbZoneServiceDnsNaptrRecordStats struct {
    DtGslbZoneServiceDnsNaptrRecordStats GslbZoneServiceDnsNaptrRecordStats `json:"dns-naptr-record"`
}


type GslbZoneServiceDnsNaptrRecordStatsStats struct {
    NaptrHits int `json:"naptr-hits"`
}

func (p *GslbZoneServiceDnsNaptrRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsNaptrRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-naptr-record/"+p.NaptrTarget+"+"+p.ServiceProto+"+"+p.Flag+"/stats"
}

func (p *GslbZoneServiceDnsNaptrRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsNaptrRecordStats,error) {
logger.Println("GslbZoneServiceDnsNaptrRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsNaptrRecordStats
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
