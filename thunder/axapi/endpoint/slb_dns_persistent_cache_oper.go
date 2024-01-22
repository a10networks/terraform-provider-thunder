

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDnsPersistentCacheOper struct {
    
    Oper SlbDnsPersistentCacheOperOper `json:"oper"`

}
type DataSlbDnsPersistentCacheOper struct {
    DtSlbDnsPersistentCacheOper SlbDnsPersistentCacheOper `json:"dns-persistent-cache"`
}


type SlbDnsPersistentCacheOperOper struct {
    CacheEntry []SlbDnsPersistentCacheOperOperCacheEntry `json:"cache-entry"`
    Total_cache int `json:"total_cache"`
}


type SlbDnsPersistentCacheOperOperCacheEntry struct {
    Domain string `json:"domain"`
    Dnssec int `json:"dnssec"`
    Cache_type int `json:"cache_type"`
    Cache_class int `json:"cache_class"`
    Q_length int `json:"q_length"`
    R_length int `json:"r_length"`
    Ttl int `json:"ttl"`
}

func (p *SlbDnsPersistentCacheOper) GetId() string{
    return "1"
}

func (p *SlbDnsPersistentCacheOper) getPath() string{
    return "slb/dns-persistent-cache/oper"
}

func (p *SlbDnsPersistentCacheOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDnsPersistentCacheOper,error) {
logger.Println("SlbDnsPersistentCacheOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDnsPersistentCacheOper
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
