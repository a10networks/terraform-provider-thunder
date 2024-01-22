

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapTranslationDomainBasicMappingRulePrefixRule struct {
	Inst struct {

    Ipv4Netmask string `json:"ipv4-netmask"`
    Name string `json:"name"`
    RuleIpv4Prefix string `json:"rule-ipv4-prefix"`
    RuleIpv6Prefix string `json:"rule-ipv6-prefix"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"prefix-rule"`
}

func (p *Cgnv6MapTranslationDomainBasicMappingRulePrefixRule) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6MapTranslationDomainBasicMappingRulePrefixRule) getPath() string{
    return "cgnv6/map/translation/domain/basic-mapping-rule/prefix-rule"
}

func (p *Cgnv6MapTranslationDomainBasicMappingRulePrefixRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRulePrefixRule::Post")
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

func (p *Cgnv6MapTranslationDomainBasicMappingRulePrefixRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRulePrefixRule::Get")
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
func (p *Cgnv6MapTranslationDomainBasicMappingRulePrefixRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRulePrefixRule::Put")
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

func (p *Cgnv6MapTranslationDomainBasicMappingRulePrefixRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRulePrefixRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
