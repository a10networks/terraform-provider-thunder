

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplatePolicy struct {
	Inst struct {

    ClassList Cgnv6TemplatePolicyClassList125 `json:"class-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"policy"`
}


type Cgnv6TemplatePolicyClassList125 struct {
    Name string `json:"name"`
    ClientIpL3Dest int `json:"client-ip-l3-dest"`
    ClientIpL7Header int `json:"client-ip-l7-header"`
    HeaderName string `json:"header-name"`
    Uuid string `json:"uuid"`
    LidList []Cgnv6TemplatePolicyClassListLidList126 `json:"lid-list"`
}


type Cgnv6TemplatePolicyClassListLidList126 struct {
    Lidnum int `json:"lidnum"`
    ConnLimit int `json:"conn-limit"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnPer int `json:"conn-per"`
    RequestLimit int `json:"request-limit"`
    RequestRateLimit int `json:"request-rate-limit"`
    RequestPer int `json:"request-per"`
    OverLimitAction int `json:"over-limit-action"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    Interval int `json:"interval"`
    Dns64 Cgnv6TemplatePolicyClassListLidListDns64127 `json:"dns64"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6TemplatePolicyClassListLidListDns64127 struct {
    Disable int `json:"disable"`
    ExclusiveAnswer int `json:"exclusive-answer"`
    Prefix string `json:"prefix"`
}

func (p *Cgnv6TemplatePolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6TemplatePolicy) getPath() string{
    return "cgnv6/template/policy"
}

func (p *Cgnv6TemplatePolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicy::Post")
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

func (p *Cgnv6TemplatePolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicy::Get")
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
func (p *Cgnv6TemplatePolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicy::Put")
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

func (p *Cgnv6TemplatePolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplatePolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
