

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Flowspec struct {
	Inst struct {

    DestAddrType string `json:"dest-addr-type"`
    DestIpHost string `json:"dest-ip-host"`
    DestIpSubnet string `json:"dest-ip-subnet"`
    DestIpv6Host string `json:"dest-ipv6-host"`
    DestIpv6Subnet string `json:"dest-ipv6-subnet"`
    DestinationPortList []FlowspecDestinationPortList `json:"destination-port-list"`
    DscpList []FlowspecDscpList `json:"dscp-list"`
    FilteringAction FlowspecFilteringAction353 `json:"filtering-action"`
    FragmentationOptionList []FlowspecFragmentationOptionList `json:"fragmentation-option-list"`
    IcmpCodeList []FlowspecIcmpCodeList `json:"icmp-code-list"`
    IcmpTypeList []FlowspecIcmpTypeList `json:"icmp-type-list"`
    Name string `json:"name"`
    OperationalMode FlowspecOperationalMode354 `json:"operational-mode"`
    PacketLengthList []FlowspecPacketLengthList `json:"packet-length-list"`
    PortList []FlowspecPortList `json:"port-list"`
    ProtocolList []FlowspecProtocolList `json:"protocol-list"`
    SourcePortList []FlowspecSourcePortList `json:"source-port-list"`
    SrcAddrType string `json:"src-addr-type"`
    SrcIpHost string `json:"src-ip-host"`
    SrcIpSubnet string `json:"src-ip-subnet"`
    SrcIpv6Host string `json:"src-ipv6-host"`
    SrcIpv6Subnet string `json:"src-ipv6-subnet"`
    TcpFlags string `json:"tcp-flags"`
    TcpFlagsBitmask string `json:"tcp-flags-bitmask"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"flowspec"`
}


type FlowspecDestinationPortList struct {
    PortAttribute string `json:"port-attribute"`
    PortNum int `json:"port-num"`
    PortNumEnd int `json:"port-num-end"`
    Uuid string `json:"uuid"`
}


type FlowspecDscpList struct {
    DscpAttribute string `json:"dscp-attribute"`
    DscpVal int `json:"dscp-val"`
    DscpValEnd int `json:"dscp-val-end"`
    Uuid string `json:"uuid"`
}


type FlowspecFilteringAction353 struct {
    TerminalAction int `json:"terminal-action"`
    SampleLog int `json:"sample-log"`
    TrafficRate int `json:"traffic-rate"`
    TrafficMarking string `json:"traffic-marking"`
    DscpVal int `json:"dscp-val"`
    TrafficClass int `json:"traffic-class"`
    Redirect string `json:"redirect"`
    NextHopNlriType string `json:"next-hop-nlri-type"`
    IpHostNlri string `json:"ip-host-nlri"`
    CopyIpHostNlri int `json:"copy-ip-host-nlri"`
    Ipv6HostNlri string `json:"ipv6-host-nlri"`
    CopyIpv6HostNlri int `json:"copy-ipv6-host-nlri"`
    NextHopType string `json:"next-hop-type"`
    IpHost string `json:"ip-host"`
    CopyIpHost int `json:"copy-ip-host"`
    Ipv6Host string `json:"ipv6-host"`
    CopyIpv6Host int `json:"copy-ipv6-host"`
    VrfTargetString string `json:"vrf-target-string"`
    VrfTargetIp string `json:"vrf-target-ip"`
    IpHostRt string `json:"ip-host-rt"`
    ValueIpHost int `json:"value-ip-host"`
    EcommCustomHex string `json:"ecomm-custom-hex"`
    Uuid string `json:"uuid"`
}


type FlowspecFragmentationOptionList struct {
    FragAttribute string `json:"frag-attribute"`
    Uuid string `json:"uuid"`
}


type FlowspecIcmpCodeList struct {
    IcmpCodeAttribute string `json:"icmp-code-attribute"`
    Code int `json:"code"`
    CodeEnd int `json:"code-end"`
    Uuid string `json:"uuid"`
}


type FlowspecIcmpTypeList struct {
    IcmpTypeAttribute string `json:"icmp-type-attribute"`
    Type int `json:"type"`
    TypeEnd int `json:"type-end"`
    Uuid string `json:"uuid"`
}


type FlowspecOperationalMode354 struct {
    Mode string `json:"mode" dval:"disabled"`
    Uuid string `json:"uuid"`
}


type FlowspecPacketLengthList struct {
    PacketLengthAttribute string `json:"packet-length-attribute"`
    Length int `json:"length"`
    LengthEnd int `json:"length-end"`
    Uuid string `json:"uuid"`
}


type FlowspecPortList struct {
    PortAttribute string `json:"port-attribute"`
    PortNum int `json:"port-num"`
    PortNumEnd int `json:"port-num-end"`
    Uuid string `json:"uuid"`
}


type FlowspecProtocolList struct {
    ProtoAttribute string `json:"proto-attribute"`
    ProtoNum int `json:"proto-num"`
    ProtoNumEnd int `json:"proto-num-end"`
    Uuid string `json:"uuid"`
}


type FlowspecSourcePortList struct {
    PortAttribute string `json:"port-attribute"`
    PortNum int `json:"port-num"`
    PortNumEnd int `json:"port-num-end"`
    Uuid string `json:"uuid"`
}

func (p *Flowspec) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Flowspec) getPath() string{
    return "flowspec"
}

func (p *Flowspec) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Flowspec::Post")
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

func (p *Flowspec) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Flowspec::Get")
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
func (p *Flowspec) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Flowspec::Put")
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

func (p *Flowspec) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Flowspec::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
