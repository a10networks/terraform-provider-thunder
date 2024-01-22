

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateDnsClassListLid struct {
	Inst struct {

    ActionValue string `json:"action-value"`
    ConnRateLimit int `json:"conn-rate-limit"`
    Dns Cgnv6TemplateDnsClassListLidDns `json:"dns"`
    Lidnum int `json:"lidnum"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    LogInterval int `json:"log-interval"`
    OverLimitAction int `json:"over-limit-action"`
    Per int `json:"per"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"lid"`
}


type Cgnv6TemplateDnsClassListLidDns struct {
    CacheAction string `json:"cache-action" dval:"cache-disable"`
    Ttl int `json:"ttl"`
    Weight int `json:"weight"`
}

func (p *Cgnv6TemplateDnsClassListLid) GetId() string{
    return strconv.Itoa(p.Inst.Lidnum)
}

func (p *Cgnv6TemplateDnsClassListLid) getPath() string{
    return "cgnv6/template/dns/" +p.Inst.Name + "/class-list/lid"
}

func (p *Cgnv6TemplateDnsClassListLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassListLid::Post")
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

func (p *Cgnv6TemplateDnsClassListLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassListLid::Get")
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
func (p *Cgnv6TemplateDnsClassListLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassListLid::Put")
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

func (p *Cgnv6TemplateDnsClassListLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassListLid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
