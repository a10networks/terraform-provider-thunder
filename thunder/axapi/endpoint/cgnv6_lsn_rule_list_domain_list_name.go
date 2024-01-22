

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDomainListName struct {
	Inst struct {

    NameDomainList string `json:"name-domain-list"`
    RuleCfg []Cgnv6LsnRuleListDomainListNameRuleCfg `json:"rule-cfg"`
    SamplingEnable []Cgnv6LsnRuleListDomainListNameSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"domain-list-name"`
}


type Cgnv6LsnRuleListDomainListNameRuleCfg struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListDomainListNameRuleCfgTcpCfg `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListDomainListNameRuleCfgUdpCfg `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListDomainListNameRuleCfgIcmpOthersCfg `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListDomainListNameRuleCfgDscpCfg `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListDomainListNameRuleCfgTcpCfg struct {
    StartPort int `json:"start-port"`
    EndPort int `json:"end-port"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    Ipv4List string `json:"ipv4-list"`
    PortList string `json:"port-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Pool string `json:"pool"`
    Shared int `json:"shared"`
    HttpAlg string `json:"http-alg"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDomainListNameRuleCfgUdpCfg struct {
    StartPort int `json:"start-port"`
    EndPort int `json:"end-port"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    Ipv4List string `json:"ipv4-list"`
    PortList string `json:"port-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Pool string `json:"pool"`
    Shared int `json:"shared"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDomainListNameRuleCfgIcmpOthersCfg struct {
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    Ipv4List string `json:"ipv4-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Pool string `json:"pool"`
    Shared int `json:"shared"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDomainListNameRuleCfgDscpCfg struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDomainListNameSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnRuleListDomainListName) GetId() string{
    return p.Inst.NameDomainList
}

func (p *Cgnv6LsnRuleListDomainListName) getPath() string{
    return "cgnv6/lsn-rule-list/" +p.Inst.Name + "/domain-list-name"
}

func (p *Cgnv6LsnRuleListDomainListName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDomainListName::Post")
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

func (p *Cgnv6LsnRuleListDomainListName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDomainListName::Get")
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
func (p *Cgnv6LsnRuleListDomainListName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDomainListName::Put")
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

func (p *Cgnv6LsnRuleListDomainListName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDomainListName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
