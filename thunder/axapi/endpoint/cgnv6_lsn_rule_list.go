

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRuleList struct {
	Inst struct {

    Default Cgnv6LsnRuleListDefault88 `json:"default"`
    DomainIp Cgnv6LsnRuleListDomainIp95 `json:"domain-ip"`
    DomainListNameList []Cgnv6LsnRuleListDomainListNameList `json:"domain-list-name-list"`
    DomainNameList []Cgnv6LsnRuleListDomainNameList `json:"domain-name-list"`
    HttpMatchDomainName int `json:"http-match-domain-name"`
    IpList []Cgnv6LsnRuleListIpList `json:"ip-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"lsn-rule-list"`
}


type Cgnv6LsnRuleListDefault88 struct {
    RuleCfg []Cgnv6LsnRuleListDefaultRuleCfg89 `json:"rule-cfg"`
    Uuid string `json:"uuid"`
    SamplingEnable []Cgnv6LsnRuleListDefaultSamplingEnable94 `json:"sampling-enable"`
}


type Cgnv6LsnRuleListDefaultRuleCfg89 struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListDefaultRuleCfgTcpCfg90 `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListDefaultRuleCfgUdpCfg91 `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg92 `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListDefaultRuleCfgDscpCfg93 `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListDefaultRuleCfgTcpCfg90 struct {
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


type Cgnv6LsnRuleListDefaultRuleCfgUdpCfg91 struct {
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


type Cgnv6LsnRuleListDefaultRuleCfgIcmpOthersCfg92 struct {
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


type Cgnv6LsnRuleListDefaultRuleCfgDscpCfg93 struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDefaultSamplingEnable94 struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6LsnRuleListDomainIp95 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []Cgnv6LsnRuleListDomainIpSamplingEnable96 `json:"sampling-enable"`
}


type Cgnv6LsnRuleListDomainIpSamplingEnable96 struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6LsnRuleListDomainListNameList struct {
    NameDomainList string `json:"name-domain-list"`
    RuleCfg []Cgnv6LsnRuleListDomainListNameListRuleCfg `json:"rule-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []Cgnv6LsnRuleListDomainListNameListSamplingEnable `json:"sampling-enable"`
}


type Cgnv6LsnRuleListDomainListNameListRuleCfg struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListDomainListNameListRuleCfgTcpCfg `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListDomainListNameListRuleCfgUdpCfg `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListDomainListNameListRuleCfgIcmpOthersCfg `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListDomainListNameListRuleCfgDscpCfg `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListDomainListNameListRuleCfgTcpCfg struct {
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


type Cgnv6LsnRuleListDomainListNameListRuleCfgUdpCfg struct {
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


type Cgnv6LsnRuleListDomainListNameListRuleCfgIcmpOthersCfg struct {
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


type Cgnv6LsnRuleListDomainListNameListRuleCfgDscpCfg struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDomainListNameListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6LsnRuleListDomainNameList struct {
    NameDomain string `json:"name-domain"`
    RuleCfg []Cgnv6LsnRuleListDomainNameListRuleCfg `json:"rule-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []Cgnv6LsnRuleListDomainNameListSamplingEnable `json:"sampling-enable"`
}


type Cgnv6LsnRuleListDomainNameListRuleCfg struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListDomainNameListRuleCfgTcpCfg `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListDomainNameListRuleCfgUdpCfg `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListDomainNameListRuleCfgIcmpOthersCfg `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListDomainNameListRuleCfgDscpCfg `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListDomainNameListRuleCfgTcpCfg struct {
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


type Cgnv6LsnRuleListDomainNameListRuleCfgUdpCfg struct {
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


type Cgnv6LsnRuleListDomainNameListRuleCfgIcmpOthersCfg struct {
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


type Cgnv6LsnRuleListDomainNameListRuleCfgDscpCfg struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListDomainNameListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6LsnRuleListIpList struct {
    Ipv4Addr string `json:"ipv4-addr"`
    RuleCfg []Cgnv6LsnRuleListIpListRuleCfg `json:"rule-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []Cgnv6LsnRuleListIpListSamplingEnable `json:"sampling-enable"`
}


type Cgnv6LsnRuleListIpListRuleCfg struct {
    Proto string `json:"proto"`
    TcpCfg Cgnv6LsnRuleListIpListRuleCfgTcpCfg `json:"tcp-cfg"`
    UdpCfg Cgnv6LsnRuleListIpListRuleCfgUdpCfg `json:"udp-cfg"`
    IcmpOthersCfg Cgnv6LsnRuleListIpListRuleCfgIcmpOthersCfg `json:"icmp-others-cfg"`
    DscpCfg Cgnv6LsnRuleListIpListRuleCfgDscpCfg `json:"dscp-cfg"`
}


type Cgnv6LsnRuleListIpListRuleCfgTcpCfg struct {
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


type Cgnv6LsnRuleListIpListRuleCfgUdpCfg struct {
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


type Cgnv6LsnRuleListIpListRuleCfgIcmpOthersCfg struct {
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


type Cgnv6LsnRuleListIpListRuleCfgDscpCfg struct {
    DscpMatch string `json:"dscp-match"`
    ActionCfg string `json:"action-cfg"`
    ActionType string `json:"action-type"`
    DscpDirection string `json:"dscp-direction"`
    DscpValue string `json:"dscp-value"`
}


type Cgnv6LsnRuleListIpListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnRuleList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6LsnRuleList) getPath() string{
    return "cgnv6/lsn-rule-list"
}

func (p *Cgnv6LsnRuleList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleList::Post")
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

func (p *Cgnv6LsnRuleList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleList::Get")
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
func (p *Cgnv6LsnRuleList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleList::Put")
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

func (p *Cgnv6LsnRuleList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRuleList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
