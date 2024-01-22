

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleListIp struct {
	Inst struct {

    Ipv4Addr string `json:"ipv4-addr"`
    RuleCfg []Cgnv6LsnRuleListIpRuleCfg `json:"rule-cfg"`
    SamplingEnable []Cgnv6LsnRuleListIpSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"ip"`
}


type Cgnv6LsnRuleListIpRuleCfg struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListIpRuleCfgTcpCfg `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListIpRuleCfgUdpCfg `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListIpRuleCfgIcmpOthersCfg `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListIpRuleCfgDscpCfg `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListIpRuleCfgTcpCfg struct {
    StartPort int `json:"start-port"`
    EndPort int `json:"end-port"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    Ipv4List string `json:"ipv4-list"`
    PortList string `json:"port-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Domain string `json:"domain"`
    Pool string `json:"pool"`
    Shared int `json:"shared"`
    HttpAlg string `json:"http-alg"`
    TimeoutVal int `json:"timeout-val"`
    Fast string `json:"fast"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListIpRuleCfgUdpCfg struct {
    StartPort int `json:"start-port"`
    EndPort int `json:"end-port"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    Ipv4List string `json:"ipv4-list"`
    PortList string `json:"port-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Domain string `json:"domain"`
    Pool string `json:"pool"`
    Shared int `json:"shared"`
    TimeoutVal int `json:"timeout-val"`
    Fast string `json:"fast"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListIpRuleCfgIcmpOthersCfg struct {
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    Ipv4List string `json:"ipv4-list"`
    NoSnat int `json:"no-snat"`
    Vrid int `json:"vrid"`
    Domain string `json:"domain"`
    Pool string `json:"pool"`
    Shared int `json:"shared"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListIpRuleCfgDscpCfg struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListIpSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnRuleListIp) GetId() string{
    return url.QueryEscape(p.Inst.Ipv4Addr)
}

func (p *Cgnv6LsnRuleListIp) getPath() string{
    return "cgnv6/lsn-rule-list/" +p.Inst.Name + "/ip"
}

func (p *Cgnv6LsnRuleListIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListIp::Post")
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

func (p *Cgnv6LsnRuleListIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListIp::Get")
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
func (p *Cgnv6LsnRuleListIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListIp::Put")
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

func (p *Cgnv6LsnRuleListIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleListIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
