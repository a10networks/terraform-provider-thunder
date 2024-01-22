

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceOper struct {
    
    DnsMxRecordList []GslbZoneServiceOperDnsMxRecordList `json:"dns-mx-record-list"`
    DnsNsRecordList []GslbZoneServiceOperDnsNsRecordList `json:"dns-ns-record-list"`
    Oper GslbZoneServiceOperOper `json:"oper"`
    ServiceName string `json:"service-name"`
    ServicePort int `json:"service-port"`
    Name string 

}
type DataGslbZoneServiceOper struct {
    DtGslbZoneServiceOper GslbZoneServiceOper `json:"service"`
}


type GslbZoneServiceOperDnsMxRecordList struct {
    MxName string `json:"mx-name"`
    Oper GslbZoneServiceOperDnsMxRecordListOper `json:"oper"`
}


type GslbZoneServiceOperDnsMxRecordListOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
    Priority int `json:"priority"`
}


type GslbZoneServiceOperDnsNsRecordList struct {
    NsName string `json:"ns-name"`
    Oper GslbZoneServiceOperDnsNsRecordListOper `json:"oper"`
}


type GslbZoneServiceOperDnsNsRecordListOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
}


type GslbZoneServiceOperOper struct {
    State string `json:"state"`
    CacheList []GslbZoneServiceOperOperCacheList `json:"cache-list"`
    SessionList []GslbZoneServiceOperOperSessionList `json:"session-list"`
    Matched int `json:"matched"`
    TotalSessions int `json:"total-sessions"`
    DnsARecordList []GslbZoneServiceOperOperDnsARecordList `json:"dns-a-record-list"`
}


type GslbZoneServiceOperOperCacheList struct {
    Alias string `json:"alias"`
    CacheLength int `json:"cache-length"`
    CacheTtl int `json:"cache-ttl"`
    CacheDnsFlag string `json:"cache-dns-flag"`
    QuestionRecords int `json:"question-records"`
    AnswerRecords int `json:"answer-records"`
    AuthorityRecords int `json:"authority-records"`
    AdditionalRecords int `json:"additional-records"`
}


type GslbZoneServiceOperOperSessionList struct {
    Client string `json:"client"`
    Best string `json:"best"`
    Mode string `json:"mode"`
    Hits int `json:"hits"`
    LastSecondHits int `json:"last-second-hits"`
    Ttl string `json:"ttl"`
    Update int `json:"update"`
    Aging int `json:"aging"`
}


type GslbZoneServiceOperOperDnsARecordList struct {
    Ip string `json:"ip"`
    RecTtl int `json:"rec-ttl"`
}

func (p *GslbZoneServiceOper) GetId() string{
    return "1"
}

func (p *GslbZoneServiceOper) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +strconv.Itoa(p.ServicePort)+"+"+p.ServiceName+"/oper"
}

func (p *GslbZoneServiceOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceOper,error) {
logger.Println("GslbZoneServiceOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceOper
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
