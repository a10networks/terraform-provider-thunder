

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapTranslationDomain struct {
	Inst struct {

    BasicMappingRule Cgnv6MapTranslationDomainBasicMappingRule102 `json:"basic-mapping-rule"`
    DefaultMappingRule Cgnv6MapTranslationDomainDefaultMappingRule104 `json:"default-mapping-rule"`
    Description string `json:"description"`
    HealthCheckGateway Cgnv6MapTranslationDomainHealthCheckGateway105 `json:"health-check-gateway"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []Cgnv6MapTranslationDomainSamplingEnable `json:"sampling-enable"`
    Tcp Cgnv6MapTranslationDomainTcp `json:"tcp"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"domain"`
}


type Cgnv6MapTranslationDomainBasicMappingRule102 struct {
    RuleIpv4AddressPortSettings string `json:"rule-ipv4-address-port-settings"`
    EaLength int `json:"ea-length"`
    ShareRatio int `json:"share-ratio"`
    PortStart int `json:"port-start"`
    Uuid string `json:"uuid"`
    PrefixRuleList []Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103 `json:"prefix-rule-list"`
}


type Cgnv6MapTranslationDomainBasicMappingRulePrefixRuleList103 struct {
    Name string `json:"name"`
    RuleIpv6Prefix string `json:"rule-ipv6-prefix"`
    RuleIpv4Prefix string `json:"rule-ipv4-prefix"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6MapTranslationDomainDefaultMappingRule104 struct {
    RuleIpv6Prefix string `json:"rule-ipv6-prefix"`
    Uuid string `json:"uuid"`
}


type Cgnv6MapTranslationDomainHealthCheckGateway105 struct {
    AddressList []Cgnv6MapTranslationDomainHealthCheckGatewayAddressList106 `json:"address-list"`
    Ipv6AddressList []Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107 `json:"ipv6-address-list"`
    WithdrawRoute string `json:"withdraw-route" dval:"any-link-failure"`
    Uuid string `json:"uuid"`
}


type Cgnv6MapTranslationDomainHealthCheckGatewayAddressList106 struct {
    Ipv4Gateway string `json:"ipv4-gateway"`
}


type Cgnv6MapTranslationDomainHealthCheckGatewayIpv6AddressList107 struct {
    Ipv6Gateway string `json:"ipv6-gateway"`
}


type Cgnv6MapTranslationDomainSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6MapTranslationDomainTcp struct {
    MssClamp Cgnv6MapTranslationDomainTcpMssClamp `json:"mss-clamp"`
}


type Cgnv6MapTranslationDomainTcpMssClamp struct {
    MssClampType string `json:"mss-clamp-type" dval:"none"`
    MssValue int `json:"mss-value"`
    MssSubtract int `json:"mss-subtract"`
    Min int `json:"min" dval:"516"`
}

func (p *Cgnv6MapTranslationDomain) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6MapTranslationDomain) getPath() string{
    return "cgnv6/map/translation/domain"
}

func (p *Cgnv6MapTranslationDomain) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomain::Post")
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

func (p *Cgnv6MapTranslationDomain) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomain::Get")
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
func (p *Cgnv6MapTranslationDomain) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomain::Put")
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

func (p *Cgnv6MapTranslationDomain) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapTranslationDomain::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
