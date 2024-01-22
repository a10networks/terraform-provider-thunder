

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsARecordDnsARecordIpv4Stats struct {
    
    DnsARecordIp string `json:"dns-a-record-ip"`
    Stats GslbZoneServiceDnsARecordDnsARecordIpv4StatsStats `json:"stats"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsARecordDnsARecordIpv4Stats struct {
    DtGslbZoneServiceDnsARecordDnsARecordIpv4Stats GslbZoneServiceDnsARecordDnsARecordIpv4Stats `json:"dns-a-record-ipv4"`
}


type GslbZoneServiceDnsARecordDnsARecordIpv4StatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv4Stats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv4Stats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-a-record/dns-a-record-ipv4/"+p.DnsARecordIp+"/stats"
}

func (p *GslbZoneServiceDnsARecordDnsARecordIpv4Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsARecordDnsARecordIpv4Stats,error) {
logger.Println("GslbZoneServiceDnsARecordDnsARecordIpv4Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsARecordDnsARecordIpv4Stats
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
