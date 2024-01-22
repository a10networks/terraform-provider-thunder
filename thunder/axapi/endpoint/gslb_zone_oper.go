

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneOper struct {
    
    DnsCaaRecordList []GslbZoneOperDnsCaaRecordList `json:"dns-caa-record-list"`
    DnsMxRecordList []GslbZoneOperDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneOperDnsNsRecordList `json:"dns-ns-record-list"`
    Name string `json:"name"`
    Oper GslbZoneOperOper `json:"oper"`
    ServiceList []GslbZoneOperServiceList `json:"service-list"`

}
type DataGslbZoneOper struct {
    DtGslbZoneOper GslbZoneOper `json:"zone"`
}


type GslbZoneOperDnsCaaRecordList struct {
    CriticalFlag int `json:"critical-flag"`
    PropertyTag string `json:"property-tag"`
    Rdata string `json:"rdata"`
    Oper GslbZoneOperDnsCaaRecordListOper `json:"oper"`
}


type GslbZoneOperDnsCaaRecordListOper struct {
    LastServer string `json:"last-server"`
}


type GslbZoneOperDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Oper GslbZoneOperDnsMxRecordListOper `json:"oper"`
}


type GslbZoneOperDnsMxRecordListOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
    Priority int `json:"priority"`
}


type GslbZoneOperDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Oper GslbZoneOperDnsNsRecordListOper `json:"oper"`
}


type GslbZoneOperDnsNsRecordListOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
}


type GslbZoneOperOper struct {
    State string `json:"state"`
    DnsSoaRecordList []GslbZoneOperOperDnsSoaRecordList `json:"dns-soa-record-list"`
}


type GslbZoneOperOperDnsSoaRecordList struct {
    Name string `json:"name"`
    Type string `json:"type"`
    Expire int `json:"expire"`
    Refresh int `json:"refresh"`
    Serial int `json:"serial"`
    Retry int `json:"retry"`
    Ttl int `json:"ttl"`
    MxName string `json:"mx-name"`
}


type GslbZoneOperServiceList struct {
    ServicePort int `json:"service-port"`
    ServiceName string `json:"service-name"`
    Oper GslbZoneOperServiceListOper `json:"oper"`
    DnsMxRecordList []GslbZoneOperServiceListDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneOperServiceListDnsNsRecordList `json:"dns-ns-record-list"`
}


type GslbZoneOperServiceListOper struct {
    State string `json:"state"`
    CacheList []GslbZoneOperServiceListOperCacheList `json:"cache-list"`
    SessionList []GslbZoneOperServiceListOperSessionList `json:"session-list"`
    Matched int `json:"matched"`
    TotalSessions int `json:"total-sessions"`
    DnsARecordList []GslbZoneOperServiceListOperDnsARecordList `json:"dns-a-record-list"`
}


type GslbZoneOperServiceListOperCacheList struct {
    Alias string `json:"alias"`
    CacheLength int `json:"cache-length"`
    CacheTtl int `json:"cache-ttl"`
    CacheDnsFlag string `json:"cache-dns-flag"`
    QuestionRecords int `json:"question-records"`
    AnswerRecords int `json:"answer-records"`
    AuthorityRecords int `json:"authority-records"`
    AdditionalRecords int `json:"additional-records"`
}


type GslbZoneOperServiceListOperSessionList struct {
    Client string `json:"client"`
    Best string `json:"best"`
    Mode string `json:"mode"`
    Hits int `json:"hits"`
    LastSecondHits int `json:"last-second-hits"`
    Ttl string `json:"ttl"`
    Update int `json:"update"`
    Aging int `json:"aging"`
}


type GslbZoneOperServiceListOperDnsARecordList struct {
    Ip string `json:"ip"`
    RecTtl int `json:"rec-ttl"`
}


type GslbZoneOperServiceListDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Oper GslbZoneOperServiceListDnsMxRecordListOper `json:"oper"`
}


type GslbZoneOperServiceListDnsMxRecordListOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
    Priority int `json:"priority"`
}


type GslbZoneOperServiceListDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Oper GslbZoneOperServiceListDnsNsRecordListOper `json:"oper"`
}


type GslbZoneOperServiceListDnsNsRecordListOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
}

func (p *GslbZoneOper) GetId() string{
    return "1"
}

func (p *GslbZoneOper) getPath() string{
    
    return "gslb/zone/"+p.Name+"/oper"
}

func (p *GslbZoneOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneOper,error) {
logger.Println("GslbZoneOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneOper
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
