

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosLongStats struct {
    
    Stats DdosLongStatsStats `json:"stats"`

}
type DataDdosLongStats struct {
    DtDdosLongStats DdosLongStats `json:"long"`
}


type DdosLongStatsStats struct {
    Tcp_syncookie_sent int `json:"tcp_syncookie_sent"`
    Tcp_syncookie_pass int `json:"tcp_syncookie_pass"`
    Tcp_syncookie_sent_fail int `json:"tcp_syncookie_sent_fail"`
    Tcp_syncookie_check_fail int `json:"tcp_syncookie_check_fail"`
    Tcp_syncookie_fail_bl int `json:"tcp_syncookie_fail_bl"`
    Tcp_outrst int `json:"tcp_outrst"`
    Tcp_syn_received int `json:"tcp_syn_received"`
    Tcp_syn_rate int `json:"tcp_syn_rate"`
    Udp_exceed_drop_conn_prate int `json:"udp_exceed_drop_conn_prate"`
    Dns_malform_drop int `json:"dns_malform_drop"`
    Dns_qry_any_drop int `json:"dns_qry_any_drop"`
    Tcp_reset_client int `json:"tcp_reset_client"`
    Tcp_reset_server int `json:"tcp_reset_server"`
    Dst_entry_learn int `json:"dst_entry_learn"`
    Dst_entry_hit int `json:"dst_entry_hit"`
    Src_entry_learn int `json:"src_entry_learn"`
    Src_entry_hit int `json:"src_entry_hit"`
    Sync_src_wl_sent int `json:"sync_src_wl_sent"`
    Sync_src_dst_wl_sent int `json:"sync_src_dst_wl_sent"`
    Sync_dst_wl_sent int `json:"sync_dst_wl_sent"`
    Sync_src_wl_rcv int `json:"sync_src_wl_rcv"`
    Sync_src_dst_wl_rcv int `json:"sync_src_dst_wl_rcv"`
    Sync_dst_wl_rcv int `json:"sync_dst_wl_rcv"`
    Dst_port_pkt_rate_exceed int `json:"dst_port_pkt_rate_exceed"`
    Dst_port_conn_limit_exceed int `json:"dst_port_conn_limit_exceed"`
    Dst_port_conn_rate_exceed int `json:"dst_port_conn_rate_exceed"`
    Dst_sport_pkt_rate_exceed int `json:"dst_sport_pkt_rate_exceed"`
    Dst_sport_conn_limit_exceed int `json:"dst_sport_conn_limit_exceed"`
    Dst_sport_conn_rate_exceed int `json:"dst_sport_conn_rate_exceed"`
    Dst_ipproto_pkt_rate_exceed int `json:"dst_ipproto_pkt_rate_exceed"`
    Tcp_ack_no_syn int `json:"tcp_ack_no_syn"`
    Tcp_out_of_order int `json:"tcp_out_of_order"`
    Tcp_zero_window int `json:"tcp_zero_window"`
    Tcp_retransmit int `json:"tcp_retransmit"`
    Tcp_action_on_ack_start int `json:"tcp_action_on_ack_start"`
    Tcp_action_on_ack_matched int `json:"tcp_action_on_ack_matched"`
    Tcp_action_on_ack_passed int `json:"tcp_action_on_ack_passed"`
    Tcp_action_on_ack_failed int `json:"tcp_action_on_ack_failed"`
    Tcp_action_on_ack_timeout int `json:"tcp_action_on_ack_timeout"`
    Tcp_action_on_ack_reset int `json:"tcp_action_on_ack_reset"`
    Src_entry_aged int `json:"src_entry_aged"`
    Dst_entry_aged int `json:"dst_entry_aged"`
    Tcp_zero_wind_bl int `json:"tcp_zero_wind_bl"`
    Tcp_out_of_seq_bl int `json:"tcp_out_of_seq_bl"`
    Tcp_retransmit_bl int `json:"tcp_retransmit_bl"`
    Syn_auth_skip int `json:"syn_auth_skip"`
    Udp_retry_pass int `json:"udp_retry_pass"`
    Dns_auth_udp_pass int `json:"dns_auth_udp_pass"`
    Udp_dst_wellknown_port_drop int `json:"udp_dst_wellknown_port_drop"`
    Udp_ntp_monlist_req_drop int `json:"udp_ntp_monlist_req_drop"`
    Udp_ntp_monlist_resp_drop int `json:"udp_ntp_monlist_resp_drop"`
    Udp_payload_too_big_drop int `json:"udp_payload_too_big_drop"`
    Udp_payload_too_small_drop int `json:"udp_payload_too_small_drop"`
    Tcp_rexmit_syn_limit_drop int `json:"tcp_rexmit_syn_limit_drop"`
    Tcp_rexmit_syn_limit_bl int `json:"tcp_rexmit_syn_limit_bl"`
    Tcp_exceed_drop_conn_prate int `json:"tcp_exceed_drop_conn_prate"`
    Ip_tunnel_rcvd int `json:"ip_tunnel_rcvd"`
    Ipv6_tunnel_rcvd int `json:"ipv6_tunnel_rcvd"`
    Gre_tunnel_rcvd int `json:"gre_tunnel_rcvd"`
    Gre_v6_tunnel_rcvd int `json:"gre_v6_tunnel_rcvd"`
    Dns_tcp_auth_pass int `json:"dns_tcp_auth_pass"`
    Jumbo_frag_drop int `json:"jumbo_frag_drop"`
    Entry_create_fail_drop int `json:"entry_create_fail_drop"`
    Dst_port_kbit_rate_exceed int `json:"dst_port_kbit_rate_exceed"`
    Dst_sport_kbit_rate_exceed int `json:"dst_sport_kbit_rate_exceed"`
    Ip_tunnel_encap int `json:"ip_tunnel_encap"`
    Ip_tunnel_encap_fail int `json:"ip_tunnel_encap_fail"`
    Ip_tunnel_decap int `json:"ip_tunnel_decap"`
    Ip_tunnel_decap_fail int `json:"ip_tunnel_decap_fail"`
    Ip_tunnel_rate_limit_inner int `json:"ip_tunnel_rate_limit_inner"`
    Ipv6_tunnel_encap int `json:"ipv6_tunnel_encap"`
    Ipv6_tunnel_encap_fail int `json:"ipv6_tunnel_encap_fail"`
    Ipv6_tunnel_decap int `json:"ipv6_tunnel_decap"`
    Ipv6_tunnel_decap_fail int `json:"ipv6_tunnel_decap_fail"`
    Ipv6_tunnel_rate_limit_inner int `json:"ipv6_tunnel_rate_limit_inner"`
    Ip_gre_tunnel_encap int `json:"ip_gre_tunnel_encap"`
    Ip_gre_tunnel_encap_fail int `json:"ip_gre_tunnel_encap_fail"`
    Ip_gre_tunnel_decap int `json:"ip_gre_tunnel_decap"`
    Ip_gre_tunnel_decap_fail int `json:"ip_gre_tunnel_decap_fail"`
    Ip_gre_tunnel_rate_limit_inner int `json:"ip_gre_tunnel_rate_limit_inner"`
    Ip_gre_tunnel_encap_key int `json:"ip_gre_tunnel_encap_key"`
    Ip_gre_tunnel_decap_key int `json:"ip_gre_tunnel_decap_key"`
    Ip_gre_tunnel_decap_key_drop int `json:"ip_gre_tunnel_decap_key_drop"`
    Ipv6_gre_tunnel_encap int `json:"ipv6_gre_tunnel_encap"`
    Ipv6_gre_tunnel_encap_fail int `json:"ipv6_gre_tunnel_encap_fail"`
    Ipv6_gre_tunnel_decap int `json:"ipv6_gre_tunnel_decap"`
    Ipv6_gre_tunnel_decap_fail int `json:"ipv6_gre_tunnel_decap_fail"`
    Ipv6_gre_tunnel_rate_limit_inner int `json:"ipv6_gre_tunnel_rate_limit_inner"`
    Ipv6_gre_tunnel_encap_key int `json:"ipv6_gre_tunnel_encap_key"`
    Ipv6_gre_tunnel_decap_key int `json:"ipv6_gre_tunnel_decap_key"`
    Ipv6_gre_tunnel_decap_key_drop int `json:"ipv6_gre_tunnel_decap_key_drop"`
    Ip_vxlan_tunnel_rcvd int `json:"ip_vxlan_tunnel_rcvd"`
    Ip_vxlan_tunnel_invalid_vni int `json:"ip_vxlan_tunnel_invalid_vni"`
    Ip_vxlan_tunnel_decap int `json:"ip_vxlan_tunnel_decap"`
    Ip_vxlan_tunnel_decap_err int `json:"ip_vxlan_tunnel_decap_err"`
    Jumbo_frag_drop_by_filter int `json:"jumbo_frag_drop_by_filter"`
    Jumbo_frag_drop_before_slb int `json:"jumbo_frag_drop_before_slb"`
    Jumbo_outgoing_mtu_exceed_drop int `json:"jumbo_outgoing_mtu_exceed_drop"`
    Jumbo_in_tunnel_drop int `json:"jumbo_in_tunnel_drop"`
    Tcp_progression_violation_exceed int `json:"tcp_progression_violation_exceed"`
    Tcp_progression_violation_exceed_bl int `json:"tcp_progression_violation_exceed_bl"`
    Tcp_progression_violation_exceed_drop int `json:"tcp_progression_violation_exceed_drop"`
    Tcp_progression_violation_exceed_reset int `json:"tcp_progression_violation_exceed_reset"`
}

func (p *DdosLongStats) GetId() string{
    return "1"
}

func (p *DdosLongStats) getPath() string{
    return "ddos/long/stats"
}

func (p *DdosLongStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosLongStats,error) {
logger.Println("DdosLongStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosLongStats
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
