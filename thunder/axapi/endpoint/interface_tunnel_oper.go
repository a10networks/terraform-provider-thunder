

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnelOper struct {
    
    Ifnum int `json:"ifnum"`
    Oper InterfaceTunnelOperOper `json:"oper"`

}
type DataInterfaceTunnelOper struct {
    DtInterfaceTunnelOper InterfaceTunnelOper `json:"tunnel"`
}


type InterfaceTunnelOperOper struct {
    State string `json:"state"`
    LineProtocol string `json:"line-protocol"`
    LinkType string `json:"link-type"`
    Mac string `json:"mac"`
    ConfigSpeed int `json:"config-speed"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6LinkLocalType string `json:"ipv6-link-local-type"`
    Ipv6LinkLocalScope string `json:"ipv6-link-local-scope"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_list []InterfaceTunnelOperOperIpv4_list `json:"ipv4_list"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_list []InterfaceTunnelOperOperIpv6_list `json:"ipv6_list"`
}


type InterfaceTunnelOperOperIpv4_list struct {
    Addr string `json:"addr"`
    Mask string `json:"mask"`
}


type InterfaceTunnelOperOperIpv6_list struct {
    Addr string `json:"addr"`
    Prefix string `json:"prefix"`
    Is_anycast int `json:"is_anycast"`
}

func (p *InterfaceTunnelOper) GetId() string{
    return "1"
}

func (p *InterfaceTunnelOper) getPath() string{
    
    return "interface/tunnel/" +strconv.Itoa(p.Ifnum)+"/oper"
}

func (p *InterfaceTunnelOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceTunnelOper,error) {
logger.Println("InterfaceTunnelOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceTunnelOper
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
