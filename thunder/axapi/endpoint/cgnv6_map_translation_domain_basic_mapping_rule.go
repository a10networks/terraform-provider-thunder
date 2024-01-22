

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapTranslationDomainBasicMappingRule struct {
	Inst struct {

    EaLength int `json:"ea-length"`
    PortStart int `json:"port-start"`
    PrefixRuleList []Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList `json:"prefix-rule-list"`
    RuleIpv4AddressPortSettings string `json:"rule-ipv4-address-port-settings"`
    ShareRatio int `json:"share-ratio"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"basic-mapping-rule"`
}


type Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList struct {
    Name string `json:"name"`
    RuleIpv6Prefix string `json:"rule-ipv6-prefix"`
    RuleIpv4Prefix string `json:"rule-ipv4-prefix"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *Cgnv6MapTranslationDomainBasicMappingRule) GetId() string{
    return "1"
}

func (p *Cgnv6MapTranslationDomainBasicMappingRule) getPath() string{
    return "cgnv6/map/translation/domain/" +p.Inst.Name + "/basic-mapping-rule"
}

func (p *Cgnv6MapTranslationDomainBasicMappingRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRule::Post")
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

func (p *Cgnv6MapTranslationDomainBasicMappingRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRule::Get")
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
func (p *Cgnv6MapTranslationDomainBasicMappingRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRule::Put")
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

func (p *Cgnv6MapTranslationDomainBasicMappingRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomainBasicMappingRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
