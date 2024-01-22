

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackOper struct {
    
    Ifnum int `json:"ifnum"`
    Oper InterfaceLoopbackOperOper `json:"oper"`

}
type DataInterfaceLoopbackOper struct {
    DtInterfaceLoopbackOper InterfaceLoopbackOper `json:"loopback"`
}


type InterfaceLoopbackOperOper struct {
    State string `json:"state"`
    Line_protocol string `json:"line_protocol"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6LinkLocalType string `json:"ipv6-link-local-type"`
    Ipv6LinkLocalScope string `json:"ipv6-link-local-scope"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_list []InterfaceLoopbackOperOperIpv4_list `json:"ipv4_list"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_list []InterfaceLoopbackOperOperIpv6_list `json:"ipv6_list"`
}


type InterfaceLoopbackOperOperIpv4_list struct {
    Addr string `json:"addr"`
    Mask string `json:"mask"`
}


type InterfaceLoopbackOperOperIpv6_list struct {
    Addr string `json:"addr"`
    Prefix string `json:"prefix"`
    Is_anycast int `json:"is_anycast"`
}

func (p *InterfaceLoopbackOper) GetId() string{
    return "1"
}

func (p *InterfaceLoopbackOper) getPath() string{
    
    return "interface/loopback/" +strconv.Itoa(p.Ifnum)+"/oper"
}

func (p *InterfaceLoopbackOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceLoopbackOper,error) {
logger.Println("InterfaceLoopbackOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceLoopbackOper
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
