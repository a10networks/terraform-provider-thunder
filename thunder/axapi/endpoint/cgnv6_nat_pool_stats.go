

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatPoolStats struct {
    
    PoolName string `json:"pool-name"`
    Stats Cgnv6NatPoolStatsStats `json:"stats"`

}
type DataCgnv6NatPoolStats struct {
    DtCgnv6NatPoolStats Cgnv6NatPoolStats `json:"pool"`
}


type Cgnv6NatPoolStatsStats struct {
    Users int `json:"users"`
    Icmp int `json:"icmp"`
    IcmpFreed int `json:"icmp-freed"`
    IcmpTotal int `json:"icmp-total"`
    IcmpRsvd int `json:"icmp-rsvd"`
    IcmpPeak int `json:"icmp-peak"`
    IcmpHitFull int `json:"icmp-hit-full"`
    Udp int `json:"udp"`
    UdpFreed int `json:"udp-freed"`
    UdpTotal int `json:"udp-total"`
    UdpRsvd int `json:"udp-rsvd"`
    UdpPeak int `json:"udp-peak"`
    UdpHitFull int `json:"udp-hit-full"`
    UdpPortOverloaded int `json:"udp-port-overloaded"`
    UdpPortOverloadCreate int `json:"udp-port-overload-create"`
    UdpPortOverloadFree int `json:"udp-port-overload-free"`
    Tcp int `json:"tcp"`
    TcpFreed int `json:"tcp-freed"`
    TcpTotal int `json:"tcp-total"`
    TcpRsvd int `json:"tcp-rsvd"`
    TcpPeak int `json:"tcp-peak"`
    TcpHitFull int `json:"tcp-hit-full"`
    TcpPortOverloaded int `json:"tcp-port-overloaded"`
    TcpPortOverloadCreate int `json:"tcp-port-overload-create"`
    TcpPortOverloadFree int `json:"tcp-port-overload-free"`
    IpUsed int `json:"ip-used"`
    IpFree int `json:"ip-free"`
    IpTotal int `json:"ip-total"`
}

func (p *Cgnv6NatPoolStats) GetId() string{
    return "1"
}

func (p *Cgnv6NatPoolStats) getPath() string{
    
    return "cgnv6/nat/pool/"+p.PoolName+"/stats"
}

func (p *Cgnv6NatPoolStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6NatPoolStats,error) {
logger.Println("Cgnv6NatPoolStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6NatPoolStats
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
