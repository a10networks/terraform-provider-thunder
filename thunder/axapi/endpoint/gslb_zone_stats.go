

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneStats struct {
    
    DnsCaaRecordList []GslbZoneStatsDnsCaaRecordList `json:"dns-caa-record-list"`
    DnsMxRecordList []GslbZoneStatsDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneStatsDnsNsRecordList `json:"dns-ns-record-list"`
    Name string `json:"name"`
    ServiceList []GslbZoneStatsServiceList `json:"service-list"`
    Stats GslbZoneStatsStats `json:"stats"`

}
type DataGslbZoneStats struct {
    DtGslbZoneStats GslbZoneStats `json:"zone"`
}


type GslbZoneStatsDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Stats GslbZoneStatsDnsCaaRecordListStats `json:"stats"`
}


type GslbZoneStatsDnsCaaRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Stats GslbZoneStatsDnsMxRecordListStats `json:"stats"`
}


type GslbZoneStatsDnsMxRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Stats GslbZoneStatsDnsNsRecordListStats `json:"stats"`
}


type GslbZoneStatsDnsNsRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsServiceList struct {
    ServicePort int `json:"service-port"`
    ServiceName string `json:"service-name"`
    Stats GslbZoneStatsServiceListStats `json:"stats"`
    DnsCnameRecordList []GslbZoneStatsServiceListDnsCnameRecordList `json:"dns-cname-record-list"`
    DnsMxRecordList []GslbZoneStatsServiceListDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneStatsServiceListDnsNsRecordList `json:"dns-ns-record-list"`
    DnsPtrRecordList []GslbZoneStatsServiceListDnsPtrRecordList `json:"dns-ptr-record-list"`
    DnsSrvRecordList []GslbZoneStatsServiceListDnsSrvRecordList `json:"dns-srv-record-list"`
    DnsNaptrRecordList []GslbZoneStatsServiceListDnsNaptrRecordList `json:"dns-naptr-record-list"`
    DnsTxtRecordList []GslbZoneStatsServiceListDnsTxtRecordList `json:"dns-txt-record-list"`
    DnsCaaRecordList []GslbZoneStatsServiceListDnsCaaRecordList `json:"dns-caa-record-list"`
}


type GslbZoneStatsServiceListStats struct {
    ReceivedQuery int `json:"received-query"`
    SentResponse int `json:"sent-response"`
    ProxyModeResponse int `json:"proxy-mode-response"`
    CacheModeResponse int `json:"cache-mode-response"`
    ServerModeResponse int `json:"server-mode-response"`
    StickyModeResponse int `json:"sticky-mode-response"`
    BackupModeResponse int `json:"backup-mode-response"`
}


type GslbZoneStatsServiceListDnsCnameRecordList struct {
    AliasName string `json:"alias-name"`
    Stats GslbZoneStatsServiceListDnsCnameRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsCnameRecordListStats struct {
    CnameHits int `json:"cname-hits"`
}


type GslbZoneStatsServiceListDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Stats GslbZoneStatsServiceListDnsMxRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsMxRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsServiceListDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Stats GslbZoneStatsServiceListDnsNsRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsNsRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsServiceListDnsPtrRecordList struct {
    PtrName string `json:"ptr-name"`
    Stats GslbZoneStatsServiceListDnsPtrRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsPtrRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsServiceListDnsSrvRecordList struct {
    SrvName string `json:"srv-name"`
    Port int `json:"port"`
    Stats GslbZoneStatsServiceListDnsSrvRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsSrvRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsServiceListDnsNaptrRecordList struct {
    NaptrTarget string `json:"naptr-target"`
    ServiceProto string `json:"service-proto"`
    Flag string `json:"flag"`
    Stats GslbZoneStatsServiceListDnsNaptrRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsNaptrRecordListStats struct {
    NaptrHits int `json:"naptr-hits"`
}


type GslbZoneStatsServiceListDnsTxtRecordList struct {
    RecordName string `json:"record-name"`
    Stats GslbZoneStatsServiceListDnsTxtRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsTxtRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsServiceListDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Stats GslbZoneStatsServiceListDnsCaaRecordListStats `json:"stats"`
}


type GslbZoneStatsServiceListDnsCaaRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneStatsStats struct {
    ReceivedQuery int `json:"received-query"`
    SentResponse int `json:"sent-response"`
    ProxyModeResponse int `json:"proxy-mode-response"`
    CacheModeResponse int `json:"cache-mode-response"`
    ServerModeResponse int `json:"server-mode-response"`
    StickyModeResponse int `json:"sticky-mode-response"`
    BackupModeResponse int `json:"backup-mode-response"`
}

func (p *GslbZoneStats) GetId() string{
    return "1"
}

func (p *GslbZoneStats) getPath() string{
    
    return "gslb/zone/"+p.Name+"/stats"
}

func (p *GslbZoneStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneStats,error) {
logger.Println("GslbZoneStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneStats
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
