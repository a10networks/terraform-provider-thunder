

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatPoolOper struct {
    
    Oper Cgnv6NatPoolOperOper `json:"oper"`
    PoolName string `json:"pool-name"`

}
type DataCgnv6NatPoolOper struct {
    DtCgnv6NatPoolOper Cgnv6NatPoolOper `json:"pool"`
}


type Cgnv6NatPoolOperOper struct {
    NatIpList []Cgnv6NatPoolOperOperNatIpList `json:"nat-ip-list"`
}


type Cgnv6NatPoolOperOperNatIpList struct {
    IpAddress string `json:"ip-address"`
    Users int `json:"users"`
    IcmpUsed int `json:"icmp-used"`
    IcmpFreed int `json:"icmp-freed"`
    IcmpTotal int `json:"icmp-total"`
    IcmpReserved int `json:"icmp-reserved"`
    IcmpPeak int `json:"icmp-peak"`
    IcmpHitFull int `json:"icmp-hit-full"`
    UdpUsed int `json:"udp-used"`
    UdpFreed int `json:"udp-freed"`
    UdpTotal int `json:"udp-total"`
    UdpReserved int `json:"udp-reserved"`
    UdpPeak int `json:"udp-peak"`
    UdpHitFull int `json:"udp-hit-full"`
    UdpPortOverloaded int `json:"udp-port-overloaded"`
    TcpUsed int `json:"tcp-used"`
    TcpFreed int `json:"tcp-freed"`
    TcpTotal int `json:"tcp-total"`
    TcpReserved int `json:"tcp-reserved"`
    TcpPeak int `json:"tcp-peak"`
    TcpHitFull int `json:"tcp-hit-full"`
    TcpPortOverloaded int `json:"tcp-port-overloaded"`
    RtspUsed int `json:"rtsp-used"`
    Obsoleted int `json:"obsoleted"`
}

func (p *Cgnv6NatPoolOper) GetId() string{
    return "1"
}

func (p *Cgnv6NatPoolOper) getPath() string{
    
    return "cgnv6/nat/pool/"+p.PoolName+"/oper"
}

func (p *Cgnv6NatPoolOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6NatPoolOper,error) {
logger.Println("Cgnv6NatPoolOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6NatPoolOper
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
