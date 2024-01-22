

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheStats struct {
    
    Name string `json:"name"`
    Stats DdosDnsCacheStatsStats `json:"stats"`

}
type DataDdosDnsCacheStats struct {
    DtDdosDnsCacheStats DdosDnsCacheStats `json:"dns-cache"`
}


type DdosDnsCacheStatsStats struct {
    TotalCachedFqdn int `json:"total-cached-fqdn"`
    TotalCachedRecords int `json:"total-cached-records"`
    FqdnA int `json:"fqdn-a"`
    FqdnAaaa int `json:"fqdn-aaaa"`
    FqdnCname int `json:"fqdn-cname"`
    FqdnNs int `json:"fqdn-ns"`
    FqdnMx int `json:"fqdn-mx"`
    FqdnSoa int `json:"fqdn-soa"`
    FqdnSrv int `json:"fqdn-srv"`
    FqdnTxt int `json:"fqdn-txt"`
    FqdnPtr int `json:"fqdn-ptr"`
    FqdnOther int `json:"fqdn-other"`
    FqdnWildcard int `json:"fqdn-wildcard"`
    FqdnDelegation int `json:"fqdn-delegation"`
    ShardSize int `json:"shard-size"`
    RespExtSize int `json:"resp-ext-size"`
    ARecord int `json:"a-record"`
    AaaaRecord int `json:"aaaa-record"`
    CnameRecord int `json:"cname-record"`
    NsRecord int `json:"ns-record"`
    MxRecord int `json:"mx-record"`
    SoaRecord int `json:"soa-record"`
    SrvRecord int `json:"srv-record"`
    TxtRecord int `json:"txt-record"`
    PtrRecord int `json:"ptr-record"`
    OtherRecord int `json:"other-record"`
    FqdnInShardFilter int `json:"fqdn-in-shard-filter"`
}

func (p *DdosDnsCacheStats) GetId() string{
    return "1"
}

func (p *DdosDnsCacheStats) getPath() string{
    
    return "ddos/dns-cache/"+p.Name+"/stats"
}

func (p *DdosDnsCacheStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDnsCacheStats,error) {
logger.Println("DdosDnsCacheStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDnsCacheStats
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
