

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplatePolicyClassListLid struct {
	Inst struct {

    ActionValue string `json:"action-value"`
    ConnLimit int `json:"conn-limit"`
    ConnPer int `json:"conn-per"`
    ConnRateLimit int `json:"conn-rate-limit"`
    Dns64 Cgnv6TemplatePolicyClassListLidDns64 `json:"dns64"`
    Interval int `json:"interval"`
    Lidnum int `json:"lidnum"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    OverLimitAction int `json:"over-limit-action"`
    RequestLimit int `json:"request-limit"`
    RequestPer int `json:"request-per"`
    RequestRateLimit int `json:"request-rate-limit"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"lid"`
}


type Cgnv6TemplatePolicyClassListLidDns64 struct {
    Disable int `json:"disable"`
    ExclusiveAnswer int `json:"exclusive-answer"`
    Prefix string `json:"prefix"`
}

func (p *Cgnv6TemplatePolicyClassListLid) GetId() string{
    return strconv.Itoa(p.Inst.Lidnum)
}

func (p *Cgnv6TemplatePolicyClassListLid) getPath() string{
    return "cgnv6/template/policy/" +p.Inst.Name + "/class-list/lid"
}

func (p *Cgnv6TemplatePolicyClassListLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicyClassListLid::Post")
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

func (p *Cgnv6TemplatePolicyClassListLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicyClassListLid::Get")
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
func (p *Cgnv6TemplatePolicyClassListLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicyClassListLid::Put")
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

func (p *Cgnv6TemplatePolicyClassListLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicyClassListLid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
