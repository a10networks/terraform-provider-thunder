

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosIpFilteringPolicy struct {
	Inst struct {

    DefaultAction string `json:"default-action" dval:"permit"`
    Name string `json:"name"`
    RuleList []DdosIpFilteringPolicyRuleList `json:"rule-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ip-filtering-policy"`
}


type DdosIpFilteringPolicyRuleList struct {
    Seq int `json:"seq"`
    Action string `json:"action" dval:"drop"`
    Glid string `json:"glid"`
    SrcIp string `json:"src-ip"`
    SrcIpv6 string `json:"src-ipv6"`
    DstIp string `json:"dst-ip"`
    DstIpv6 string `json:"dst-ipv6"`
    Protocol string `json:"protocol"`
    ProtoNum int `json:"proto-num"`
    SrcPort int `json:"src-port"`
    SrcPortStart int `json:"src-port-start"`
    SrcPortEnd int `json:"src-port-end"`
    DstPort int `json:"dst-port"`
    DstPortStart int `json:"dst-port-start"`
    DstPortEnd int `json:"dst-port-end"`
    TcpFlag string `json:"tcp-flag"`
    TcpFlagsBitmask string `json:"tcp-flags-bitmask"`
    IcmpType int `json:"icmp-type"`
    IcmpCode int `json:"icmp-code"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosIpFilteringPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosIpFilteringPolicy) getPath() string{
    return "ddos/ip-filtering-policy"
}

func (p *DdosIpFilteringPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicy::Post")
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

func (p *DdosIpFilteringPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicy::Get")
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
func (p *DdosIpFilteringPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicy::Put")
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

func (p *DdosIpFilteringPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
