

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateDnsDns64 struct {
	Inst struct {

    AnswerOnlyDisable int `json:"answer-only-disable"`
    AuthData int `json:"auth-data"`
    Cache int `json:"cache"`
    ChangeQuery int `json:"change-query"`
    CompressDisable int `json:"compress-disable"`
    DeepCheckQr int `json:"deep-check-qr"`
    DeepCheckRrDisable int `json:"deep-check-rr-disable"`
    DropCnameDisable int `json:"drop-cname-disable"`
    EdnsAppend int `json:"edns-append"`
    Enable int `json:"enable"`
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
    Name string 

	} `json:"dns64"`
}

func (p *Cgnv6TemplateDnsDns64) GetId() string{
    return "1"
}

func (p *Cgnv6TemplateDnsDns64) getPath() string{
    return "cgnv6/template/dns/" +p.Inst.Name + "/dns64"
}

func (p *Cgnv6TemplateDnsDns64) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsDns64::Post")
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

func (p *Cgnv6TemplateDnsDns64) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsDns64::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *Cgnv6TemplateDnsDns64) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsDns64::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *Cgnv6TemplateDnsDns64) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsDns64::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
