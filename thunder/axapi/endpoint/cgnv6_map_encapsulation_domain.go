

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapEncapsulationDomain struct {
	Inst struct {

    BasicMappingRule Cgnv6MapEncapsulationDomainBasicMappingRule97 `json:"basic-mapping-rule"`
    Description string `json:"description"`
    Format string `json:"format"`
    HealthCheckGateway Cgnv6MapEncapsulationDomainHealthCheckGateway99 `json:"health-check-gateway"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []Cgnv6MapEncapsulationDomainSamplingEnable `json:"sampling-enable"`
    TunnelEndpointAddress string `json:"tunnel-endpoint-address"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"domain"`
}


type Cgnv6MapEncapsulationDomainBasicMappingRule97 struct {
    RuleIpv4AddressPortSettings string `json:"rule-ipv4-address-port-settings"`
    EaLength int `json:"ea-length"`
    ShareRatio int `json:"share-ratio"`
    PortStart int `json:"port-start"`
    Uuid string `json:"uuid"`
    PrefixRuleList []Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98 `json:"prefix-rule-list"`
}


type Cgnv6MapEncapsulationDomainBasicMappingRulePrefixRuleList98 struct {
    Name string `json:"name"`
    RuleIpv6Prefix string `json:"rule-ipv6-prefix"`
    RuleIpv4Prefix string `json:"rule-ipv4-prefix"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv4AddressPortSettings string `json:"ipv4-address-port-settings"`
    EaLength int `json:"ea-length"`
    ShareRatio int `json:"share-ratio"`
    PortStart int `json:"port-start"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type Cgnv6MapEncapsulationDomainHealthCheckGateway99 struct {
    AddressList []Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100 `json:"address-list"`
    Ipv6AddressList []Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101 `json:"ipv6-address-list"`
    WithdrawRoute string `json:"withdraw-route" dval:"any-link-failure"`
    Uuid string `json:"uuid"`
}


type Cgnv6MapEncapsulationDomainHealthCheckGatewayAddressList100 struct {
    Ipv4Gateway string `json:"ipv4-gateway"`
}


type Cgnv6MapEncapsulationDomainHealthCheckGatewayIpv6AddressList101 struct {
    Ipv6Gateway string `json:"ipv6-gateway"`
}


type Cgnv6MapEncapsulationDomainSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6MapEncapsulationDomain) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6MapEncapsulationDomain) getPath() string{
    return "cgnv6/map/encapsulation/domain"
}

func (p *Cgnv6MapEncapsulationDomain) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomain::Post")
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

func (p *Cgnv6MapEncapsulationDomain) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomain::Get")
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
func (p *Cgnv6MapEncapsulationDomain) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomain::Put")
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

func (p *Cgnv6MapEncapsulationDomain) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6MapEncapsulationDomain::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
