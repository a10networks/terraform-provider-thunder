

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifOper struct {
    
    Ifname string `json:"ifname"`
    Oper InterfaceLifOperOper `json:"oper"`

}
type DataInterfaceLifOper struct {
    DtInterfaceLifOper InterfaceLifOper `json:"lif"`
}


type InterfaceLifOperOper struct {
    State string `json:"state"`
    Line_protocol string `json:"line_protocol"`
    Link_type string `json:"link_type"`
    Encapsulation_type string `json:"encapsulation_type"`
    Member_id string `json:"member_id"`
    Keep_alive string `json:"keep_alive"`
    Mac string `json:"mac"`
    IgmpQuerySent int `json:"igmp-query-sent"`
    IcmpRateLimitCurrent int `json:"icmp-rate-limit-current"`
    IcmpRateOverLimitDrop int `json:"icmp-rate-over-limit-drop"`
    Icmp6RateLimitCurrent int `json:"icmp6-rate-limit-current"`
    Icmp6RateOverLimitDrop int `json:"icmp6-rate-over-limit-drop"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_list []InterfaceLifOperOperIpv4_list `json:"ipv4_list"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_list []InterfaceLifOperOperIpv6_list `json:"ipv6_list"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6LinkLocalType string `json:"ipv6-link-local-type"`
    Ipv6LinkLocalScope string `json:"ipv6-link-local-scope"`
    Ip_unnumbered_enabled int `json:"ip_unnumbered_enabled"`
    Mtu string `json:"mtu"`
}


type InterfaceLifOperOperIpv4_list struct {
    Addr string `json:"addr"`
    Mask string `json:"mask"`
}


type InterfaceLifOperOperIpv6_list struct {
    Addr string `json:"addr"`
    Prefix string `json:"prefix"`
    Is_anycast int `json:"is_anycast"`
}

func (p *InterfaceLifOper) GetId() string{
    return "1"
}

func (p *InterfaceLifOper) getPath() string{
    
    return "interface/lif/"+p.Ifname+"/oper"
}

func (p *InterfaceLifOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceLifOper,error) {
logger.Println("InterfaceLifOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceLifOper
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
