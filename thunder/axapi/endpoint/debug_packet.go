

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugPacket struct {
	Inst struct {

    AllIpv4 int `json:"all-ipv4"`
    AllIpv6 int `json:"all-ipv6"`
    AllSctpPorts int `json:"all-sctp-ports"`
    AllTcpPorts int `json:"all-tcp-ports"`
    AllUdpPorts int `json:"all-udp-ports"`
    Arp int `json:"arp"`
    Count1 int `json:"count1"`
    Detail int `json:"detail"`
    Ethernet int `json:"ethernet"`
    Icmp int `json:"icmp"`
    Icmpv6 int `json:"icmpv6"`
    Interface int `json:"interface"`
    Ip int `json:"ip"`
    Ipv4ad string `json:"ipv4ad"`
    Ipv6 int `json:"ipv6"`
    Ipv6ad string `json:"ipv6ad"`
    L3Protocol int `json:"l3-protocol"`
    L4Protocol int `json:"l4-protocol"`
    Neighbor int `json:"neighbor"`
    PortRange int `json:"port-range"`
    Sctp int `json:"sctp"`
    Tcp int `json:"tcp"`
    Timestamp int `json:"timestamp"`
    Udp int `json:"udp"`
    Uuid string `json:"uuid"`
    Ve int `json:"ve"`

	} `json:"packet"`
}

func (p *DebugPacket) GetId() string{
    return "1"
}

func (p *DebugPacket) getPath() string{
    return "debug/packet"
}

func (p *DebugPacket) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPacket::Post")
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

func (p *DebugPacket) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPacket::Get")
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
func (p *DebugPacket) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPacket::Put")
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

func (p *DebugPacket) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugPacket::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
