

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkOper struct {
    
    Ifnum int `json:"ifnum"`
    Oper InterfaceTrunkOperOper `json:"oper"`

}
type DataInterfaceTrunkOper struct {
    DtInterfaceTrunkOper InterfaceTrunkOper `json:"trunk"`
}


type InterfaceTrunkOperOper struct {
    State string `json:"state"`
    LineProtocol string `json:"line-protocol"`
    TrunkMember []InterfaceTrunkOperOperTrunkMember `json:"trunk-member"`
    Hardware string `json:"Hardware"`
    Address string `json:"Address"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6LinkLocalType string `json:"ipv6-link-local-type"`
    Ipv6LinkLocalScope string `json:"ipv6-link-local-scope"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_list []InterfaceTrunkOperOperIpv4_list `json:"ipv4_list"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_list []InterfaceTrunkOperOperIpv6_list `json:"ipv6_list"`
    IgmpQuerySent int `json:"igmp-query-sent"`
    IcmpRateLimitCurrent int `json:"icmp-rate-limit-current"`
    IcmpRateOverLimitDrop int `json:"icmp-rate-over-limit-drop"`
    Icmp6RateLimitCurrent int `json:"icmp6-rate-limit-current"`
    Icmp6RateOverLimitDrop int `json:"icmp6-rate-over-limit-drop"`
    VlanId int `json:"vlan-id"`
    Ip_unnumbered_oper int `json:"ip_unnumbered_oper"`
    Ip_unnumbered_enabled int `json:"ip_unnumbered_enabled"`
    Ip_unnumbered_mac_learned int `json:"ip_unnumbered_mac_learned"`
    Ip_unnumbered_peer_lla string `json:"ip_unnumbered_peer_lla"`
    Mtu int `json:"mtu"`
}


type InterfaceTrunkOperOperTrunkMember struct {
    Members int `json:"members"`
}


type InterfaceTrunkOperOperIpv4_list struct {
    Addr string `json:"addr"`
    Mask string `json:"mask"`
}


type InterfaceTrunkOperOperIpv6_list struct {
    Addr string `json:"addr"`
    Prefix string `json:"prefix"`
    Is_anycast int `json:"is_anycast"`
}

func (p *InterfaceTrunkOper) GetId() string{
    return "1"
}

func (p *InterfaceTrunkOper) getPath() string{
    
    return "interface/trunk/" +strconv.Itoa(p.Ifnum)+"/oper"
}

func (p *InterfaceTrunkOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceTrunkOper,error) {
logger.Println("InterfaceTrunkOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceTrunkOper
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
