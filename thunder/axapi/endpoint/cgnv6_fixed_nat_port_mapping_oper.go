

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatPortMappingOper struct {
    
    Oper Cgnv6FixedNatPortMappingOperOper `json:"oper"`

}
type DataCgnv6FixedNatPortMappingOper struct {
    DtCgnv6FixedNatPortMappingOper Cgnv6FixedNatPortMappingOper `json:"port-mapping"`
}


type Cgnv6FixedNatPortMappingOperOper struct {
    MappingList []Cgnv6FixedNatPortMappingOperOperMappingList `json:"mapping-list"`
    Partition string `json:"partition"`
    InsideUserV4 string `json:"inside-user-v4"`
    InsideUserV6 string `json:"inside-user-v6"`
    NatIp string `json:"nat-ip"`
    NatPort int `json:"nat-port"`
}


type Cgnv6FixedNatPortMappingOperOperMappingList struct {
    NatAddress string `json:"nat-address"`
    TcpPortStart int `json:"tcp-port-start"`
    TcpPortEnd int `json:"tcp-port-end"`
    UdpPortStart int `json:"udp-port-start"`
    UdpPortEnd int `json:"udp-port-end"`
    IcmpPortStart int `json:"icmp-port-start"`
    IcmpPortEnd int `json:"icmp-port-end"`
    AssignedTo string `json:"assigned-to"`
}

func (p *Cgnv6FixedNatPortMappingOper) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatPortMappingOper) getPath() string{
    return "cgnv6/fixed-nat/port-mapping/oper"
}

func (p *Cgnv6FixedNatPortMappingOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatPortMappingOper,error) {
logger.Println("Cgnv6FixedNatPortMappingOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatPortMappingOper
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
