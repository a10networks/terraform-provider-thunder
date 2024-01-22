

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListDefault struct {
	Inst struct {

    RuleCfg []Cgnv6LsnRuleListDefaultRuleCfg `json:"rule-cfg"`
    SamplingEnable []Cgnv6LsnRuleListDefaultSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"default"`
}


type Cgnv6LsnRuleListDefaultRuleCfg struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListDefaultRuleCfgTcpCfg `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListDefaultRuleCfgUdpCfg `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListDefaultRuleCfgDscpCfg `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListDefaultRuleCfgTcpCfg struct {
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


type Cgnv6LsnRuleListDefaultRuleCfgUdpCfg struct {
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


type Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg struct {
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


type Cgnv6LsnRuleListDefaultRuleCfgDscpCfg struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDefaultSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnRuleListDefault) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRuleListDefault) getPath() string{
    return "cgnv6/lsn-rule-list/" +p.Inst.Name + "/default"
}

func (p *Cgnv6LsnRuleListDefault) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDefault::Post")
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

func (p *Cgnv6LsnRuleListDefault) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDefault::Get")
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
func (p *Cgnv6LsnRuleListDefault) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDefault::Put")
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

func (p *Cgnv6LsnRuleListDefault) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListDefault::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
