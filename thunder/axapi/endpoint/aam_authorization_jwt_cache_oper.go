

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthorizationJwtCacheOper struct {
    
    Oper AamAuthorizationJwtCacheOperOper `json:"oper"`

}
type DataAamAuthorizationJwtCacheOper struct {
    DtAamAuthorizationJwtCacheOper AamAuthorizationJwtCacheOper `json:"cache"`
}


type AamAuthorizationJwtCacheOperOper struct {
    Token_cached int `json:"token_cached"`
    Token_cache_hit int `json:"token_cache_hit"`
    Max_token_cache int `json:"max_token_cache"`
    CacheList []AamAuthorizationJwtCacheOperOperCacheList `json:"cache-list"`
    Audience string `json:"audience"`
}


type AamAuthorizationJwtCacheOperOperCacheList struct {
    CacheId int `json:"cache-id"`
    Issuer string `json:"issuer"`
    Subject string `json:"subject"`
    Audience string `json:"audience"`
    ClientIp string `json:"client-ip"`
    Ttl int `json:"ttl"`
}

func (p *AamAuthorizationJwtCacheOper) GetId() string{
    return "1"
}

func (p *AamAuthorizationJwtCacheOper) getPath() string{
    return "aam/authorization/jwt/cache/oper"
}

func (p *AamAuthorizationJwtCacheOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAamAuthorizationJwtCacheOper,error) {
logger.Println("AamAuthorizationJwtCacheOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAamAuthorizationJwtCacheOper
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
