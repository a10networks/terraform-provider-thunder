

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwMatchOper struct {
    
    Oper FwMatchOperOper `json:"oper"`

}
type DataFwMatchOper struct {
    DtFwMatchOper FwMatchOper `json:"match"`
}


type FwMatchOperOper struct {
    ErrorMsg string `json:"error-msg"`
    ActiveAccessControlRuleSet string `json:"active-access-control-rule-set"`
    ActiveTrafficControlRuleSet string `json:"active-traffic-control-rule-set"`
    RuleIds []FwMatchOperOperRuleIds `json:"rule-ids"`
    MatchingRulesFetched int `json:"matching-rules-fetched"`
    MatchingRulesTotal int `json:"matching-rules-total"`
    ShowAll int `json:"show-all"`
    Vlan int `json:"vlan"`
    SrcIpv4Addr string `json:"src-ipv4-addr"`
    DstIpv4Addr string `json:"dst-ipv4-addr"`
    SrcIpv6Addr string `json:"src-ipv6-addr"`
    DstIpv6Addr string `json:"dst-ipv6-addr"`
    Tcp int `json:"tcp"`
    Udp int `json:"udp"`
    Icmp int `json:"icmp"`
    Icmpv6 int `json:"icmpv6"`
    SrcPort int `json:"src-port"`
    DstPort int `json:"dst-port"`
    IcmpType int `json:"icmp-type"`
}


type FwMatchOperOperRuleIds struct {
    MatchingRuleType int `json:"matching-rule-type"`
    MatchingRule string `json:"matching-rule"`
}

func (p *FwMatchOper) GetId() string{
    return "1"
}

func (p *FwMatchOper) getPath() string{
    return "fw/match/oper"
}

func (p *FwMatchOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwMatchOper,error) {
logger.Println("FwMatchOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwMatchOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
