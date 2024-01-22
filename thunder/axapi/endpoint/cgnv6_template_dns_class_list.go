

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateDnsClassList struct {
	Inst struct {

    LidList []Cgnv6TemplateDnsClassListLidList `json:"lid-list"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`

	} `json:"class-list"`
}


type Cgnv6TemplateDnsClassListLidList struct {
    Lidnum int `json:"lidnum"`
    ConnRateLimit int `json:"conn-rate-limit"`
    Per int `json:"per"`
    OverLimitAction int `json:"over-limit-action"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    LogInterval int `json:"log-interval"`
    Dns Cgnv6TemplateDnsClassListLidListDns `json:"dns"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplateDnsClassListLidListDns struct {
    CacheAction string `json:"cache-action" dval:"cache-disable"`
    Ttl int `json:"ttl"`
    Weight int `json:"weight"`
}

func (p *Cgnv6TemplateDnsClassList) GetId() string{
    return "1"
}

func (p *Cgnv6TemplateDnsClassList) getPath() string{
    return "cgnv6/template/dns/"+p.Inst.Name+"/class-list"
}

func (p *Cgnv6TemplateDnsClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassList::Post")
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

func (p *Cgnv6TemplateDnsClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassList::Get")
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
func (p *Cgnv6TemplateDnsClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassList::Put")
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

func (p *Cgnv6TemplateDnsClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateDnsClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
