

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDnsCacheOper struct {
    
    Oper SlbDnsCacheOperOper `json:"oper"`

}
type DataSlbDnsCacheOper struct {
    DtSlbDnsCacheOper SlbDnsCacheOper `json:"dns-cache"`
}


type SlbDnsCacheOperOper struct {
    CacheClient []SlbDnsCacheOperOperCacheClient `json:"cache-client"`
    CacheEntry []SlbDnsCacheOperOperCacheEntry `json:"cache-entry"`
    Total int `json:"total"`
    Client int `json:"client"`
    Entry int `json:"entry"`
    Global int `json:"global"`
    CacheContent int `json:"cache-content"`
    Vport int `json:"vport"`
    VsName string `json:"vs-name"`
    PortType string `json:"port-type"`
    PortNum int `json:"port-num"`
    TypeValue int `json:"type-value"`
    FqdnDomain string `json:"fqdn-domain"`
    ClassString string `json:"class-string"`
    ClassValue int `json:"class-value"`
    TypeString string `json:"type-string"`
    DomainName string `json:"domain-name"`
    ContentMode string `json:"content-mode"`
    RdataSizeValue int `json:"rdata-size-value"`
    RdataAll int `json:"rdata-all"`
    RecordNumValue int `json:"record-num-value"`
    RecordAll int `json:"record-all"`
}


type SlbDnsCacheOperOperCacheClient struct {
    Domain string `json:"domain"`
    Address string `json:"address"`
    Unit_type string `json:"unit_type"`
    Curr_rate int `json:"curr_rate"`
    Over_rate_limit_times int `json:"over_rate_limit_times"`
    Lockup int `json:"lockup"`
    Lockup_time int `json:"lockup_time"`
}


type SlbDnsCacheOperOperCacheEntry struct {
    Name string `json:"name"`
    Domain string `json:"domain"`
    Dnssec int `json:"dnssec"`
    Cache_negative int `json:"cache_negative"`
    Cache_type int `json:"cache_type"`
    Cache_class int `json:"cache_class"`
    Q_length int `json:"q_length"`
    R_length int `json:"r_length"`
    Ttl int `json:"ttl"`
    Age int `json:"age"`
    Weight int `json:"weight"`
    Hits int `json:"hits"`
    An_count int `json:"an_count"`
    Ns_count int `json:"ns_count"`
    Ar_count int `json:"ar_count"`
    EntryRecord []SlbDnsCacheOperOperCacheEntryEntryRecord `json:"entry-record"`
}


type SlbDnsCacheOperOperCacheEntryEntryRecord struct {
    Record_type int `json:"record_type"`
    Record_class int `json:"record_class"`
    Record_ttl int `json:"record_ttl"`
    Record_rdlen int `json:"record_rdlen"`
    Record_rdata string `json:"record_rdata"`
    Record_rdata_tc int `json:"record_rdata_tc"`
}

func (p *SlbDnsCacheOper) GetId() string{
    return "1"
}

func (p *SlbDnsCacheOper) getPath() string{
    return "slb/dns-cache/oper"
}

func (p *SlbDnsCacheOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDnsCacheOper,error) {
logger.Println("SlbDnsCacheOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDnsCacheOper
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
