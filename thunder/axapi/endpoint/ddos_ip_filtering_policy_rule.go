

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosIpFilteringPolicyRule struct {
	Inst struct {

    Action string `json:"action" dval:"drop"`
    DstIp string `json:"dst-ip"`
    DstIpv6 string `json:"dst-ipv6"`
    DstPort int `json:"dst-port"`
    DstPortEnd int `json:"dst-port-end"`
    DstPortStart int `json:"dst-port-start"`
    Glid string `json:"glid"`
    IcmpCode int `json:"icmp-code"`
    IcmpType int `json:"icmp-type"`
    ProtoNum int `json:"proto-num"`
    Protocol string `json:"protocol"`
    Seq int `json:"seq"`
    SrcIp string `json:"src-ip"`
    SrcIpv6 string `json:"src-ipv6"`
    SrcPort int `json:"src-port"`
    SrcPortEnd int `json:"src-port-end"`
    SrcPortStart int `json:"src-port-start"`
    TcpFlag string `json:"tcp-flag"`
    TcpFlagsBitmask string `json:"tcp-flags-bitmask"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"rule"`
}

func (p *DdosIpFilteringPolicyRule) GetId() string{
    return strconv.Itoa(p.Inst.Seq)
}

func (p *DdosIpFilteringPolicyRule) getPath() string{
    return "ddos/ip-filtering-policy/" +p.Inst.Name + "/rule"
}

func (p *DdosIpFilteringPolicyRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicyRule::Post")
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

func (p *DdosIpFilteringPolicyRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicyRule::Get")
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
func (p *DdosIpFilteringPolicyRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicyRule::Put")
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

func (p *DdosIpFilteringPolicyRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosIpFilteringPolicyRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
