

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceVeOper struct {
    
    Ifnum int `json:"ifnum"`
    Oper InterfaceVeOperOper `json:"oper"`

}
type DataInterfaceVeOper struct {
    DtInterfaceVeOper InterfaceVeOper `json:"ve"`
}


type InterfaceVeOperOper struct {
    State string `json:"state"`
    Line_protocol string `json:"line_protocol"`
    Link_type string `json:"link_type"`
    Mac string `json:"mac"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6LinkLocalType string `json:"ipv6-link-local-type"`
    Ipv6LinkLocalScope string `json:"ipv6-link-local-scope"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_list []InterfaceVeOperOperIpv4_list `json:"ipv4_list"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_list []InterfaceVeOperOperIpv6_list `json:"ipv6_list"`
    IgmpQuerySent int `json:"igmp-query-sent"`
    IcmpRateLimitCurrent int `json:"icmp-rate-limit-current"`
    IcmpRateOverLimitDrop int `json:"icmp-rate-over-limit-drop"`
    Icmp6RateLimitCurrent int `json:"icmp6-rate-limit-current"`
    Icmp6RateOverLimitDrop int `json:"icmp6-rate-over-limit-drop"`
    UserTrunkId int `json:"user-trunk-id"`
    Ip_unnumbered_oper int `json:"ip_unnumbered_oper"`
    Ip_unnumbered_enabled int `json:"ip_unnumbered_enabled"`
    Ip_unnumbered_mac_learned int `json:"ip_unnumbered_mac_learned"`
    Ip_unnumbered_peer_lla string `json:"ip_unnumbered_peer_lla"`
    Mtu int `json:"mtu"`
}


type InterfaceVeOperOperIpv4_list struct {
    Addr string `json:"addr"`
    Mask string `json:"mask"`
}


type InterfaceVeOperOperIpv6_list struct {
    Addr string `json:"addr"`
    Prefix string `json:"prefix"`
    Is_anycast int `json:"is_anycast"`
}

func (p *InterfaceVeOper) GetId() string{
    return "1"
}

func (p *InterfaceVeOper) getPath() string{
    
    return "interface/ve/" +strconv.Itoa(p.Ifnum)+"/oper"
}

func (p *InterfaceVeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceVeOper,error) {
logger.Println("InterfaceVeOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceVeOper
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
