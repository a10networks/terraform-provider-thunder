

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceStats struct {
    
    DnsCaaRecordList []GslbZoneServiceStatsDnsCaaRecordList `json:"dns-caa-record-list"`
    DnsCnameRecordList []GslbZoneServiceStatsDnsCnameRecordList `json:"dns-cname-record-list"`
    DnsMxRecordList []GslbZoneServiceStatsDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNaptrRecordList []GslbZoneServiceStatsDnsNaptrRecordList `json:"dns-naptr-record-list"`
    DnsNsRecordList []GslbZoneServiceStatsDnsNsRecordList `json:"dns-ns-record-list"`
    DnsPtrRecordList []GslbZoneServiceStatsDnsPtrRecordList `json:"dns-ptr-record-list"`
    DnsSrvRecordList []GslbZoneServiceStatsDnsSrvRecordList `json:"dns-srv-record-list"`
    DnsTxtRecordList []GslbZoneServiceStatsDnsTxtRecordList `json:"dns-txt-record-list"`
    ServiceName string `json:"service-name"`
    ServicePort int `json:"service-port"`
    Stats GslbZoneServiceStatsStats `json:"stats"`
    Name string 

}
type DataGslbZoneServiceStats struct {
    DtGslbZoneServiceStats GslbZoneServiceStats `json:"service"`
}


type GslbZoneServiceStatsDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Stats GslbZoneServiceStatsDnsCaaRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsCaaRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneServiceStatsDnsCnameRecordList struct {
    AliasName string `json:"alias-name"`
    Stats GslbZoneServiceStatsDnsCnameRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsCnameRecordListStats struct {
    CnameHits int `json:"cname-hits"`
}


type GslbZoneServiceStatsDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Stats GslbZoneServiceStatsDnsMxRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsMxRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneServiceStatsDnsNaptrRecordList struct {
    NaptrTarget string `json:"naptr-target"`
    ServiceProto string `json:"service-proto"`
    Flag string `json:"flag"`
    Stats GslbZoneServiceStatsDnsNaptrRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsNaptrRecordListStats struct {
    NaptrHits int `json:"naptr-hits"`
}


type GslbZoneServiceStatsDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Stats GslbZoneServiceStatsDnsNsRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsNsRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneServiceStatsDnsPtrRecordList struct {
    PtrName string `json:"ptr-name"`
    Stats GslbZoneServiceStatsDnsPtrRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsPtrRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneServiceStatsDnsSrvRecordList struct {
    SrvName string `json:"srv-name"`
    Port int `json:"port"`
    Stats GslbZoneServiceStatsDnsSrvRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsSrvRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneServiceStatsDnsTxtRecordList struct {
    RecordName string `json:"record-name"`
    Stats GslbZoneServiceStatsDnsTxtRecordListStats `json:"stats"`
}


type GslbZoneServiceStatsDnsTxtRecordListStats struct {
    Hits int `json:"hits"`
}


type GslbZoneServiceStatsStats struct {
    ReceivedQuery int `json:"received-query"`
    SentResponse int `json:"sent-response"`
    ProxyModeResponse int `json:"proxy-mode-response"`
    CacheModeResponse int `json:"cache-mode-response"`
    ServerModeResponse int `json:"server-mode-response"`
    StickyModeResponse int `json:"sticky-mode-response"`
    BackupModeResponse int `json:"backup-mode-response"`
}

func (p *GslbZoneServiceStats) GetId() string{
    return "1"
}

func (p *GslbZoneServiceStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +strconv.Itoa(p.ServicePort)+"+"+p.ServiceName+"/stats"
}

func (p *GslbZoneServiceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceStats,error) {
logger.Println("GslbZoneServiceStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceStats
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
