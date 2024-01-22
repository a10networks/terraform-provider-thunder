

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SixrdDomainStats struct {
    
    Name string `json:"name"`
    Stats Cgnv6SixrdDomainStatsStats `json:"stats"`

}
type DataCgnv6SixrdDomainStats struct {
    DtCgnv6SixrdDomainStats Cgnv6SixrdDomainStats `json:"domain"`
}


type Cgnv6SixrdDomainStatsStats struct {
    OutboundTcpPacketsReceived int `json:"outbound-tcp-packets-received"`
    OutboundUdpPacketsReceived int `json:"outbound-udp-packets-received"`
    OutboundIcmpPacketsReceived int `json:"outbound-icmp-packets-received"`
    OutboundOtherPacketsReceived int `json:"outbound-other-packets-received"`
    OutboundPacketsDrop int `json:"outbound-packets-drop"`
    OutboundIpv6DestUnreachable int `json:"outbound-ipv6-dest-unreachable"`
    OutboundFragmentIpv6 int `json:"outbound-fragment-ipv6"`
    InboundTcpPacketsReceived int `json:"inbound-tcp-packets-received"`
    InboundUdpPacketsReceived int `json:"inbound-udp-packets-received"`
    InboundIcmpPacketsReceived int `json:"inbound-icmp-packets-received"`
    InboundOtherPacketsReceived int `json:"inbound-other-packets-received"`
    InboundPacketsDrop int `json:"inbound-packets-drop"`
    InboundIpv4DestUnreachable int `json:"inbound-ipv4-dest-unreachable"`
    InboundFragmentIpv4 int `json:"inbound-fragment-ipv4"`
    InboundTunnelFragmentIpv6 int `json:"inbound-tunnel-fragment-ipv6"`
    VportMatched int `json:"vport-matched"`
    UnknownDelegatedPrefix int `json:"unknown-delegated-prefix"`
    PacketTooBig int `json:"packet-too-big"`
    NotLocalIp int `json:"not-local-ip"`
    FragmentError int `json:"fragment-error"`
    OtherError int `json:"other-error"`
}

func (p *Cgnv6SixrdDomainStats) GetId() string{
    return "1"
}

func (p *Cgnv6SixrdDomainStats) getPath() string{
    
    return "cgnv6/sixrd/domain/"+p.Name+"/stats"
}

func (p *Cgnv6SixrdDomainStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6SixrdDomainStats,error) {
logger.Println("Cgnv6SixrdDomainStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6SixrdDomainStats
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
