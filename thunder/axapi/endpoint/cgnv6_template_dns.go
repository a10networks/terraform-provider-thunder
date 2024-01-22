

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateDns struct {
	Inst struct {

    ClassList Cgnv6TemplateDnsClassList111 `json:"class-list"`
    DefaultPolicy string `json:"default-policy" dval:"nocache"`
    DisableDnsTemplate int `json:"disable-dns-template"`
    Dns64 Cgnv6TemplateDnsDns64114 `json:"dns64"`
    Drop int `json:"drop"`
    Forward string `json:"forward"`
    MaxCacheSize int `json:"max-cache-size"`
    Name string `json:"name"`
    Period int `json:"period"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type Cgnv6TemplateDnsClassList111 struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    LidList []Cgnv6TemplateDnsClassListLidList112 `json:"lid-list"`
}


type Cgnv6TemplateDnsClassListLidList112 struct {
    Lidnum int `json:"lidnum"`
    ConnRateLimit int `json:"conn-rate-limit"`
    Per int `json:"per"`
    OverLimitAction int `json:"over-limit-action"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    LogInterval int `json:"log-interval"`
    Dns Cgnv6TemplateDnsClassListLidListDns113 `json:"dns"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplateDnsClassListLidListDns113 struct {
    CacheAction string `json:"cache-action" dval:"cache-disable"`
    Ttl int `json:"ttl"`
    Weight int `json:"weight"`
}


type Cgnv6TemplateDnsDns64114 struct {
    Enable int `json:"enable"`
    AnswerOnlyDisable int `json:"answer-only-disable"`
    AuthData int `json:"auth-data"`
    Cache int `json:"cache"`
    ChangeQuery int `json:"change-query"`
    CompressDisable int `json:"compress-disable"`
    DeepCheckQr int `json:"deep-check-qr"`
    DeepCheckRrDisable int `json:"deep-check-rr-disable"`
    DropCnameDisable int `json:"drop-cname-disable"`
    EdnsAppend int `json:"edns-append"`
    FastAppend int `json:"fast-append"`
    IgnoreRcode3Disable int `json:"ignore-rcode3-disable"`
    MaxQrLength int `json:"max-qr-length" dval:"128"`
    ParallelQuery int `json:"parallel-query"`
    PassiveQueryDisable int `json:"passive-query-disable"`
    Retry int `json:"retry" dval:"3"`
    SingleResponseDisable int `json:"single-response-disable"`
    Timeout int `json:"timeout" dval:"1"`
    TransPtr int `json:"trans-ptr"`
    TransPtrQuery int `json:"trans-ptr-query"`
    Ttl int `json:"ttl"`
    Uuid string `json:"uuid"`
}

func (p *Cgnv6TemplateDns) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6TemplateDns) getPath() string{
    return "cgnv6/template/dns"
}

func (p *Cgnv6TemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDns::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *Cgnv6TemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDns::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *Cgnv6TemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDns::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *Cgnv6TemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
