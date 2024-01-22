

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSwitchStats struct {
    
    Stats SlbSwitchStatsStats `json:"stats"`

}
type DataSlbSwitchStats struct {
    DtSlbSwitchStats SlbSwitchStats `json:"switch"`
}


type SlbSwitchStatsStats struct {
    Fwlb int `json:"fwlb"`
    Licexpire_drop int `json:"licexpire_drop"`
    Bwl_drop int `json:"bwl_drop"`
    Rx_kernel int `json:"rx_kernel"`
    Rx_arp_req int `json:"rx_arp_req"`
    Rx_arp_resp int `json:"rx_arp_resp"`
    Vlan_flood int `json:"vlan_flood"`
    L2_def_vlan_drop int `json:"l2_def_vlan_drop"`
    Ipv4_noroute_drop int `json:"ipv4_noroute_drop"`
    Ipv6_noroute_drop int `json:"ipv6_noroute_drop"`
    Prot_down_drop int `json:"prot_down_drop"`
    L2_forward int `json:"l2_forward"`
    L3_forward_ip int `json:"l3_forward_ip"`
    L3_forward_ipv6 int `json:"l3_forward_ipv6"`
    L4_process int `json:"l4_process"`
    Unknown_prot_drop int `json:"unknown_prot_drop"`
    Ttl_exceeded_drop int `json:"ttl_exceeded_drop"`
    Linkdown_drop int `json:"linkdown_drop"`
    Sport_drop int `json:"sport_drop"`
    Incorrect_len_drop int `json:"incorrect_len_drop"`
    Ip_defrag int `json:"ip_defrag"`
    Acl_deny int `json:"acl_deny"`
    Ipfrag_tcp int `json:"ipfrag_tcp"`
    Ipfrag_overlap int `json:"ipfrag_overlap"`
    Ipfrag_timeout int `json:"ipfrag_timeout"`
    Ipfrag_overload int `json:"ipfrag_overload"`
    Ipfrag_reasmoks int `json:"ipfrag_reasmoks"`
    Ipfrag_reasmfails int `json:"ipfrag_reasmfails"`
    Badpkt_drop int `json:"badpkt_drop"`
    Ipsec_drop int `json:"ipsec_drop"`
    Bpdu_rcvd int `json:"bpdu_rcvd"`
    Bpdu_sent int `json:"bpdu_sent"`
    Ctrl_syn_rate_drop int `json:"ctrl_syn_rate_drop"`
    Ip_defrag_invalid_len int `json:"ip_defrag_invalid_len"`
    Ipv4_frag_6rd_ok int `json:"ipv4_frag_6rd_ok"`
    Ipv4_frag_6rd_drop int `json:"ipv4_frag_6rd_drop"`
    No_ip_drop int `json:"no_ip_drop"`
    Ipv6frag_udp int `json:"ipv6frag_udp"`
    Ipv6frag_udp_dropped int `json:"ipv6frag_udp_dropped"`
    Ipv6frag_tcp_dropped int `json:"ipv6frag_tcp_dropped"`
    Ipv6frag_ipip_ok int `json:"ipv6frag_ipip_ok"`
    Ipv6frag_ipip_dropped int `json:"ipv6frag_ipip_dropped"`
    Ip_frag_oversize int `json:"ip_frag_oversize"`
    Ip_frag_too_many int `json:"ip_frag_too_many"`
    Ipv4_novlanfwd_drop int `json:"ipv4_novlanfwd_drop"`
    Ipv6_novlanfwd_drop int `json:"ipv6_novlanfwd_drop"`
    Fpga_error_pkt1 int `json:"fpga_error_pkt1"`
    Fpga_error_pkt2 int `json:"fpga_error_pkt2"`
    Max_arp_drop int `json:"max_arp_drop"`
    Ipv6frag_tcp int `json:"ipv6frag_tcp"`
    Ipv6frag_icmp int `json:"ipv6frag_icmp"`
    Ipv6frag_ospf int `json:"ipv6frag_ospf"`
    Ipv6frag_esp int `json:"ipv6frag_esp"`
    L4_in_ctrl_cpu int `json:"l4_in_ctrl_cpu"`
    Mgmt_svc_drop int `json:"mgmt_svc_drop"`
    Jumbo_frag_drop int `json:"jumbo_frag_drop"`
    Ipv6_jumbo_frag_drop int `json:"ipv6_jumbo_frag_drop"`
    Ipipv6_jumbo_frag_drop int `json:"ipipv6_jumbo_frag_drop"`
    Ipv6_ndisc_dad_solicits int `json:"ipv6_ndisc_dad_solicits"`
    Ipv6_ndisc_dad_adverts int `json:"ipv6_ndisc_dad_adverts"`
    Ipv6_ndisc_mac_changes int `json:"ipv6_ndisc_mac_changes"`
    Ipv6_ndisc_out_of_memory int `json:"ipv6_ndisc_out_of_memory"`
    Sp_non_ctrl_pkt_drop int `json:"sp_non_ctrl_pkt_drop"`
    Urpf_pkt_drop int `json:"urpf_pkt_drop"`
    Fw_smp_zone_mismatch int `json:"fw_smp_zone_mismatch"`
    Ipfrag_udp int `json:"ipfrag_udp"`
    Ipfrag_icmp int `json:"ipfrag_icmp"`
    Ipfrag_ospf int `json:"ipfrag_ospf"`
    Ipfrag_esp int `json:"ipfrag_esp"`
    Ipfrag_tcp_dropped int `json:"ipfrag_tcp_dropped"`
    Ipfrag_udp_dropped int `json:"ipfrag_udp_dropped"`
    Ipfrag_ipip_dropped int `json:"ipfrag_ipip_dropped"`
    Redirect_fwd_fail int `json:"redirect_fwd_fail"`
    Redirect_fwd_sent int `json:"redirect_fwd_sent"`
    Redirect_rev_fail int `json:"redirect_rev_fail"`
    Redirect_rev_sent int `json:"redirect_rev_sent"`
    Redirect_setup_fail int `json:"redirect_setup_fail"`
    Ip_frag_sent int `json:"ip_frag_sent"`
    Invalid_rx_arp_pkt int `json:"invalid_rx_arp_pkt"`
    Invalid_sender_mac_arp_drop int `json:"invalid_sender_mac_arp_drop"`
    Dev_based_arp_drop int `json:"dev_based_arp_drop"`
    Scaleout_arp_drop int `json:"scaleout_arp_drop"`
    Virtual_ip_not_found_arp_drop int `json:"virtual_ip_not_found_arp_drop"`
    Inactive_static_nat_pool_arp_drop int `json:"inactive_static_nat_pool_arp_drop"`
    Inactive_nat_pool_arp_drop int `json:"inactive_nat_pool_arp_drop"`
    Scaleout_hairpin_arp_drop int `json:"scaleout_hairpin_arp_drop"`
    Self_grat_arp_drop int `json:"self_grat_arp_drop"`
    Self_grat_nat_ip_arp_drop int `json:"self_grat_nat_ip_arp_drop"`
    Ip_not_found_arp_drop int `json:"ip_not_found_arp_drop"`
    Dev_link_down_arp_drop int `json:"dev_link_down_arp_drop"`
    Lacp_tx_intf_err_drop int `json:"lacp_tx_intf_err_drop"`
    Service_chain_sent int `json:"service_chain_sent"`
    Service_chain_rcvd int `json:"service_chain_rcvd"`
    Unnumbered_nat_error int `json:"unnumbered_nat_error"`
    Unnumbered_unsupported_drop int `json:"unnumbered_unsupported_drop"`
    Ipv6frag_gre_dropped int `json:"ipv6frag_gre_dropped"`
    Ipv6_ndisc_dad_prefix_mismatch_drop int `json:"ipv6_ndisc_dad_prefix_mismatch_drop"`
    Bw_ignore_limit int `json:"bw_ignore_limit"`
    Ppsl_drop_egr int `json:"ppsl_drop_egr"`
    Ppsl_drop_ing int `json:"ppsl_drop_ing"`
    Ppsl_ignore_limit int `json:"ppsl_ignore_limit"`
    Closed_port_syn_drop int `json:"closed_port_syn_drop"`
    Tls13_drop_req int `json:"tls13_drop_req"`
    Tls13_ignore_req int `json:"tls13_ignore_req"`
    Tls12_drop_req int `json:"tls12_drop_req"`
    Tls12_ignore_req int `json:"tls12_ignore_req"`
    Tls12_tls13_drop_req int `json:"tls12_tls13_drop_req"`
    Tls12_tls13_ignore_req int `json:"tls12_tls13_ignore_req"`
}

func (p *SlbSwitchStats) GetId() string{
    return "1"
}

func (p *SlbSwitchStats) getPath() string{
    return "slb/switch/stats"
}

func (p *SlbSwitchStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSwitchStats,error) {
logger.Println("SlbSwitchStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSwitchStats
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
