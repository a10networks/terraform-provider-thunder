

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetRuleActionGroup struct {
	Inst struct {

    Cgnv6 int `json:"cgnv6"`
    Cgnv6DsLite string `json:"cgnv6-ds-lite"`
    Cgnv6DsLiteLsnLid int `json:"cgnv6-ds-lite-lsn-lid"`
    Cgnv6LsnLid int `json:"cgnv6-lsn-lid"`
    Cgnv6Policy string `json:"cgnv6-policy"`
    DenyFwLog string `json:"deny-fw-log"`
    DenyLog int `json:"deny-log"`
    DenyLogTemplateType string `json:"deny-log-template-type"`
    DscpNumber int `json:"dscp-number"`
    DscpValue string `json:"dscp-value"`
    Forward int `json:"forward"`
    InspectPayload int `json:"inspect-payload"`
    Ipsec int `json:"ipsec"`
    IpsecGroup int `json:"ipsec-group"`
    ListenOnPort int `json:"listen-on-port"`
    LoggingTemplateList []RuleSetRuleActionGroupLoggingTemplateList `json:"logging-template-list"`
    PermitLimitPolicy int `json:"permit-limit-policy"`
    PermitLog int `json:"permit-log"`
    PermitRespondToUserMac int `json:"permit-respond-to-user-mac"`
    ResetFwLog string `json:"reset-fw-log"`
    ResetLog int `json:"reset-log"`
    ResetLogTemplateType string `json:"reset-log-template-type"`
    ResetRespondToUserMac int `json:"reset-respond-to-user-mac"`
    SetDscp int `json:"set-dscp"`
    Type string `json:"type"`
    Uuid string `json:"uuid"`
    VpnIpsecGroupName string `json:"vpn-ipsec-group-name"`
    VpnIpsecName string `json:"vpn-ipsec-name"`
    Name string 

	} `json:"action-group"`
}


type RuleSetRuleActionGroupLoggingTemplateList struct {
    PermitLogTemplateType string `json:"permit-log-template-type"`
    PermitFwLog string `json:"permit-fw-log"`
    PermitCgnv6Log string `json:"permit-cgnv6-log"`
    PermitNetflowLog string `json:"permit-netflow-log"`
}

func (p *RuleSetRuleActionGroup) GetId() string{
    return "1"
}

func (p *RuleSetRuleActionGroup) getPath() string{
    return "rule-set/" +p.Inst.Name + "/rule/" +p.Inst.Name + "/action-group"
}

func (p *RuleSetRuleActionGroup) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleActionGroup::Post")
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

func (p *RuleSetRuleActionGroup) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleActionGroup::Get")
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
func (p *RuleSetRuleActionGroup) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleActionGroup::Put")
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

func (p *RuleSetRuleActionGroup) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RuleSetRuleActionGroup::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
