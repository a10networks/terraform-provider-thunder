

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetOper struct {
    
    Ifnum int `json:"ifnum"`
    Oper InterfaceEthernetOperOper `json:"oper"`

}
type DataInterfaceEthernetOper struct {
    DtInterfaceEthernetOper InterfaceEthernetOper `json:"ethernet"`
}


type InterfaceEthernetOperOper struct {
    State string `json:"state"`
    Line_protocol string `json:"line_protocol"`
    Link_type string `json:"link_type"`
    Mac string `json:"mac"`
    Config_speed string `json:"config_speed"`
    Actual_speed string `json:"actual_speed"`
    Config_duplexity string `json:"config_duplexity"`
    Actual_duplexity string `json:"actual_duplexity"`
    Media_type string `json:"media_type"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv4_addr_count int `json:"ipv4_addr_count"`
    Ipv4_list []InterfaceEthernetOperOperIpv4_list `json:"ipv4_list"`
    Ipv6_addr_count int `json:"ipv6_addr_count"`
    Ipv6_list []InterfaceEthernetOperOperIpv6_list `json:"ipv6_list"`
    Ipv6LinkLocal string `json:"ipv6-link-local"`
    Ipv6LinkLocalPrefix string `json:"ipv6-link-local-prefix"`
    Ipv6LinkLocalType string `json:"ipv6-link-local-type"`
    Ipv6LinkLocalScope string `json:"ipv6-link-local-scope"`
    IgmpQuerySent int `json:"igmp-query-sent"`
    IcmpRateLimitCurrent int `json:"icmp-rate-limit-current"`
    IcmpRateOverLimitDrop int `json:"icmp-rate-over-limit-drop"`
    Icmp6RateLimitCurrent int `json:"icmp6-rate-limit-current"`
    Icmp6RateOverLimitDrop int `json:"icmp6-rate-over-limit-drop"`
    IsTagged int `json:"is-tagged"`
    VlanId int `json:"vlan-id"`
    TaggedVlanList string `json:"tagged-vlan-list"`
    Span_mode string `json:"span_mode"`
    Span_port_state string `json:"span_port_state"`
    Is_lead_member int `json:"is_lead_member"`
    Is_blocked int `json:"is_blocked"`
    Current_vnp_id int `json:"current_vnp_id"`
    Port_vnp_id int `json:"port_vnp_id"`
    Is_pristine int `json:"is_pristine"`
    Rate_byte_rcvd int `json:"rate_byte_rcvd"`
    Rate_byte_sent int `json:"rate_byte_sent"`
    Rate_pkt_rcvd int `json:"rate_pkt_rcvd"`
    Rate_pkt_sent int `json:"rate_pkt_sent"`
    Input_utilization int `json:"input_utilization"`
    Output_utilization int `json:"output_utilization"`
    Is_device_transparent int `json:"is_device_transparent"`
    Incoming_pkts_mirrored int `json:"incoming_pkts_mirrored"`
    Outgoing_pkts_mirrored int `json:"outgoing_pkts_mirrored"`
    Incoming_pkts_monitored int `json:"incoming_pkts_monitored"`
    Outgoing_pkts_monitored int `json:"outgoing_pkts_monitored"`
    Ip_unnumbered_oper int `json:"ip_unnumbered_oper"`
    Ip_unnumbered_enabled int `json:"ip_unnumbered_enabled"`
    Ip_unnumbered_mac_learned int `json:"ip_unnumbered_mac_learned"`
    Ip_unnumbered_peer_lla string `json:"ip_unnumbered_peer_lla"`
    Last_count_clear_at string `json:"last_count_clear_at"`
    Last_up_event_at string `json:"last_up_event_at"`
    Last_down_event_at string `json:"last_down_event_at"`
}


type InterfaceEthernetOperOperIpv4_list struct {
    Addr string `json:"addr"`
    Mask string `json:"mask"`
}


type InterfaceEthernetOperOperIpv6_list struct {
    Addr string `json:"addr"`
    Prefix string `json:"prefix"`
    Is_anycast int `json:"is_anycast"`
}

func (p *InterfaceEthernetOper) GetId() string{
    return "1"
}

func (p *InterfaceEthernetOper) getPath() string{
    
    return "interface/ethernet/" +strconv.Itoa(p.Ifnum)+"/oper"
}

func (p *InterfaceEthernetOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataInterfaceEthernetOper,error) {
logger.Println("InterfaceEthernetOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataInterfaceEthernetOper
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
