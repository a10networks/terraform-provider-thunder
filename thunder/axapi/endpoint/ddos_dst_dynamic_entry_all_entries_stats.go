

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntryAllEntriesStats struct {
    
    Stats DdosDstDynamicEntryAllEntriesStatsStats `json:"stats"`

}
type DataDdosDstDynamicEntryAllEntriesStats struct {
    DtDdosDstDynamicEntryAllEntriesStats DdosDstDynamicEntryAllEntriesStats `json:"all-entries"`
}


type DdosDstDynamicEntryAllEntriesStatsStats struct {
    Dst_tcp_any_exceed int `json:"dst_tcp_any_exceed"`
    Dst_tcp_pkt_rate_exceed int `json:"dst_tcp_pkt_rate_exceed"`
    Dst_tcp_conn_rate_exceed int `json:"dst_tcp_conn_rate_exceed"`
    Dst_udp_any_exceed int `json:"dst_udp_any_exceed"`
    Dst_udp_pkt_rate_exceed int `json:"dst_udp_pkt_rate_exceed"`
    Dst_udp_conn_limit_exceed int `json:"dst_udp_conn_limit_exceed"`
    Dst_udp_conn_rate_exceed int `json:"dst_udp_conn_rate_exceed"`
    Dst_icmp_pkt_rate_exceed int `json:"dst_icmp_pkt_rate_exceed"`
    Dst_other_pkt_rate_exceed int `json:"dst_other_pkt_rate_exceed"`
    Dst_other_frag_pkt_rate_exceed int `json:"dst_other_frag_pkt_rate_exceed"`
    Dst_port_pkt_rate_exceed int `json:"dst_port_pkt_rate_exceed"`
    Dst_port_conn_limit_exceed int `json:"dst_port_conn_limit_exceed"`
    Dst_port_conn_rate_exceed int `json:"dst_port_conn_rate_exceed"`
    Dst_pkt_sent int `json:"dst_pkt_sent"`
    Dst_udp_pkt_sent int `json:"dst_udp_pkt_sent"`
    Dst_tcp_pkt_sent int `json:"dst_tcp_pkt_sent"`
    Dst_icmp_pkt_sent int `json:"dst_icmp_pkt_sent"`
    Dst_other_pkt_sent int `json:"dst_other_pkt_sent"`
    Dst_tcp_conn_limit_exceed int `json:"dst_tcp_conn_limit_exceed"`
    Dst_tcp_pkt_rcvd int `json:"dst_tcp_pkt_rcvd"`
    Dst_udp_pkt_rcvd int `json:"dst_udp_pkt_rcvd"`
    Dst_icmp_pkt_rcvd int `json:"dst_icmp_pkt_rcvd"`
    Dst_other_pkt_rcvd int `json:"dst_other_pkt_rcvd"`
    Dst_udp_filter_match int `json:"dst_udp_filter_match"`
    Dst_udp_filter_not_match int `json:"dst_udp_filter_not_match"`
    Dst_udp_filter_action_blacklist int `json:"dst_udp_filter_action_blacklist"`
    Dst_udp_filter_action_drop int `json:"dst_udp_filter_action_drop"`
    Dst_tcp_syn int `json:"dst_tcp_syn"`
    Dst_tcp_syn_drop int `json:"dst_tcp_syn_drop"`
    Dst_tcp_src_rate_drop int `json:"dst_tcp_src_rate_drop"`
    Dst_udp_src_rate_drop int `json:"dst_udp_src_rate_drop"`
    Dst_icmp_src_rate_drop int `json:"dst_icmp_src_rate_drop"`
    Dst_other_frag_src_rate_drop int `json:"dst_other_frag_src_rate_drop"`
    Dst_other_src_rate_drop int `json:"dst_other_src_rate_drop"`
    Dst_tcp_drop int `json:"dst_tcp_drop"`
    Dst_udp_drop int `json:"dst_udp_drop"`
    Dst_icmp_drop int `json:"dst_icmp_drop"`
    Dst_frag_drop int `json:"dst_frag_drop"`
    Dst_other_drop int `json:"dst_other_drop"`
    Dst_tcp_auth int `json:"dst_tcp_auth"`
    Dst_udp_filter_action_default_pass int `json:"dst_udp_filter_action_default_pass"`
    Dst_tcp_filter_match int `json:"dst_tcp_filter_match"`
    Dst_tcp_filter_not_match int `json:"dst_tcp_filter_not_match"`
    Dst_tcp_filter_action_blacklist int `json:"dst_tcp_filter_action_blacklist"`
    Dst_tcp_filter_action_drop int `json:"dst_tcp_filter_action_drop"`
    Dst_tcp_filter_action_default_pass int `json:"dst_tcp_filter_action_default_pass"`
    Dst_udp_filter_action_whitelist int `json:"dst_udp_filter_action_whitelist"`
    Dst_udp_kibit_rate_drop int `json:"dst_udp_kibit_rate_drop"`
    Dst_tcp_kibit_rate_drop int `json:"dst_tcp_kibit_rate_drop"`
    Dst_icmp_kibit_rate_drop int `json:"dst_icmp_kibit_rate_drop"`
    Dst_other_kibit_rate_drop int `json:"dst_other_kibit_rate_drop"`
    Dst_port_undef_drop int `json:"dst_port_undef_drop"`
    Dst_port_bl int `json:"dst_port_bl"`
    Dst_src_port_bl int `json:"dst_src_port_bl"`
    Dst_port_kbit_rate_exceed int `json:"dst_port_kbit_rate_exceed"`
    Dst_tcp_src_drop int `json:"dst_tcp_src_drop"`
    Dst_udp_src_drop int `json:"dst_udp_src_drop"`
    Dst_icmp_src_drop int `json:"dst_icmp_src_drop"`
    Dst_other_src_drop int `json:"dst_other_src_drop"`
    Tcp_syn_rcvd int `json:"tcp_syn_rcvd"`
    Tcp_syn_ack_rcvd int `json:"tcp_syn_ack_rcvd"`
    Tcp_ack_rcvd int `json:"tcp_ack_rcvd"`
    Tcp_fin_rcvd int `json:"tcp_fin_rcvd"`
    Tcp_rst_rcvd int `json:"tcp_rst_rcvd"`
    Ingress_bytes int `json:"ingress_bytes"`
    Egress_bytes int `json:"egress_bytes"`
    Ingress_packets int `json:"ingress_packets"`
    Egress_packets int `json:"egress_packets"`
    Tcp_fwd_recv int `json:"tcp_fwd_recv"`
    Udp_fwd_recv int `json:"udp_fwd_recv"`
    Icmp_fwd_recv int `json:"icmp_fwd_recv"`
    Tcp_syn_cookie_fail int `json:"tcp_syn_cookie_fail"`
    Dst_tcp_session_created int `json:"dst_tcp_session_created"`
    Dst_udp_session_created int `json:"dst_udp_session_created"`
    Dst_tcp_filter_action_whitelist int `json:"dst_tcp_filter_action_whitelist"`
    Dst_other_filter_match int `json:"dst_other_filter_match"`
    Dst_other_filter_not_match int `json:"dst_other_filter_not_match"`
    Dst_other_filter_action_blacklist int `json:"dst_other_filter_action_blacklist"`
    Dst_other_filter_action_drop int `json:"dst_other_filter_action_drop"`
    Dst_other_filter_action_whitelist int `json:"dst_other_filter_action_whitelist"`
    Dst_other_filter_action_default_pass int `json:"dst_other_filter_action_default_pass"`
    Dst_blackhole_inject int `json:"dst_blackhole_inject"`
    Dst_blackhole_withdraw int `json:"dst_blackhole_withdraw"`
    Dst_tcp_out_of_seq_excd int `json:"dst_tcp_out_of_seq_excd"`
    Dst_tcp_retransmit_excd int `json:"dst_tcp_retransmit_excd"`
    Dst_tcp_zero_window_excd int `json:"dst_tcp_zero_window_excd"`
    Dst_tcp_conn_prate_excd int `json:"dst_tcp_conn_prate_excd"`
    Dst_tcp_action_on_ack_init int `json:"dst_tcp_action_on_ack_init"`
    Dst_tcp_action_on_ack_gap_drop int `json:"dst_tcp_action_on_ack_gap_drop"`
    Dst_tcp_action_on_ack_fail int `json:"dst_tcp_action_on_ack_fail"`
    Dst_tcp_action_on_ack_pass int `json:"dst_tcp_action_on_ack_pass"`
    Dst_tcp_action_on_syn_init int `json:"dst_tcp_action_on_syn_init"`
    Dst_tcp_action_on_syn_gap_drop int `json:"dst_tcp_action_on_syn_gap_drop"`
    Dst_tcp_action_on_syn_fail int `json:"dst_tcp_action_on_syn_fail"`
    Dst_tcp_action_on_syn_pass int `json:"dst_tcp_action_on_syn_pass"`
    Udp_payload_too_small int `json:"udp_payload_too_small"`
    Udp_payload_too_big int `json:"udp_payload_too_big"`
    Dst_udp_conn_prate_excd int `json:"dst_udp_conn_prate_excd"`
    Dst_udp_ntp_monlist_req int `json:"dst_udp_ntp_monlist_req"`
    Dst_udp_ntp_monlist_resp int `json:"dst_udp_ntp_monlist_resp"`
    Dst_udp_wellknown_sport_drop int `json:"dst_udp_wellknown_sport_drop"`
    Dst_udp_retry_init int `json:"dst_udp_retry_init"`
    Dst_udp_retry_pass int `json:"dst_udp_retry_pass"`
    Dst_tcp_bytes_drop int `json:"dst_tcp_bytes_drop"`
    Dst_udp_bytes_drop int `json:"dst_udp_bytes_drop"`
    Dst_icmp_bytes_drop int `json:"dst_icmp_bytes_drop"`
    Dst_other_bytes_drop int `json:"dst_other_bytes_drop"`
    Dst_out_no_route int `json:"dst_out_no_route"`
    Outbound_bytes_sent int `json:"outbound_bytes_sent"`
    Outbound_pkt_drop int `json:"outbound_pkt_drop"`
    Outbound_bytes_drop int `json:"outbound_bytes_drop"`
    Outbound_pkt_sent int `json:"outbound_pkt_sent"`
    Inbound_bytes_sent int `json:"inbound_bytes_sent"`
    Inbound_bytes_drop int `json:"inbound_bytes_drop"`
    Dst_src_port_pkt_rate_exceed int `json:"dst_src_port_pkt_rate_exceed"`
    Dst_src_port_kbit_rate_exceed int `json:"dst_src_port_kbit_rate_exceed"`
    Dst_src_port_conn_limit_exceed int `json:"dst_src_port_conn_limit_exceed"`
    Dst_src_port_conn_rate_exceed int `json:"dst_src_port_conn_rate_exceed"`
    Dst_ip_proto_pkt_rate_exceed int `json:"dst_ip_proto_pkt_rate_exceed"`
    Dst_ip_proto_kbit_rate_exceed int `json:"dst_ip_proto_kbit_rate_exceed"`
    Dst_tcp_port_any_exceed int `json:"dst_tcp_port_any_exceed"`
    Dst_udp_port_any_exceed int `json:"dst_udp_port_any_exceed"`
    Dst_tcp_auth_pass int `json:"dst_tcp_auth_pass"`
    Dst_tcp_rst_cookie_fail int `json:"dst_tcp_rst_cookie_fail"`
    Dst_tcp_unauth_drop int `json:"dst_tcp_unauth_drop"`
    Src_tcp_syn_auth_fail int `json:"src_tcp_syn_auth_fail"`
    Src_tcp_syn_cookie_sent int `json:"src_tcp_syn_cookie_sent"`
    Src_tcp_syn_cookie_fail int `json:"src_tcp_syn_cookie_fail"`
    Src_tcp_rst_cookie_fail int `json:"src_tcp_rst_cookie_fail"`
    Src_tcp_unauth_drop int `json:"src_tcp_unauth_drop"`
    Src_tcp_action_on_syn_init int `json:"src_tcp_action_on_syn_init"`
    Src_tcp_action_on_syn_gap_drop int `json:"src_tcp_action_on_syn_gap_drop"`
    Src_tcp_action_on_syn_fail int `json:"src_tcp_action_on_syn_fail"`
    Src_tcp_action_on_ack_init int `json:"src_tcp_action_on_ack_init"`
    Src_tcp_action_on_ack_gap_drop int `json:"src_tcp_action_on_ack_gap_drop"`
    Src_tcp_action_on_ack_fail int `json:"src_tcp_action_on_ack_fail"`
    Src_tcp_out_of_seq_excd int `json:"src_tcp_out_of_seq_excd"`
    Src_tcp_retransmit_excd int `json:"src_tcp_retransmit_excd"`
    Src_tcp_zero_window_excd int `json:"src_tcp_zero_window_excd"`
    Src_tcp_conn_prate_excd int `json:"src_tcp_conn_prate_excd"`
    Src_udp_min_payload int `json:"src_udp_min_payload"`
    Src_udp_max_payload int `json:"src_udp_max_payload"`
    Src_udp_conn_prate_excd int `json:"src_udp_conn_prate_excd"`
    Src_udp_ntp_monlist_req int `json:"src_udp_ntp_monlist_req"`
    Src_udp_ntp_monlist_resp int `json:"src_udp_ntp_monlist_resp"`
    Src_udp_wellknown_sport_drop int `json:"src_udp_wellknown_sport_drop"`
    Src_udp_retry_init int `json:"src_udp_retry_init"`
    Dst_udp_retry_gap_drop int `json:"dst_udp_retry_gap_drop"`
    Dst_udp_retry_fail int `json:"dst_udp_retry_fail"`
    Dst_tcp_session_aged int `json:"dst_tcp_session_aged"`
    Dst_udp_session_aged int `json:"dst_udp_session_aged"`
    Dst_tcp_conn_close int `json:"dst_tcp_conn_close"`
    Dst_tcp_conn_close_half_open int `json:"dst_tcp_conn_close_half_open"`
    Dst_l4_tcp_auth int `json:"dst_l4_tcp_auth"`
    Tcp_l4_syn_cookie_fail int `json:"tcp_l4_syn_cookie_fail"`
    Tcp_l4_rst_cookie_fail int `json:"tcp_l4_rst_cookie_fail"`
    Tcp_l4_unauth_drop int `json:"tcp_l4_unauth_drop"`
    Src_tcp_filter_action_blacklist int `json:"src_tcp_filter_action_blacklist"`
    Src_tcp_filter_action_whitelist int `json:"src_tcp_filter_action_whitelist"`
    Src_tcp_filter_action_drop int `json:"src_tcp_filter_action_drop"`
    Src_tcp_filter_action_default_pass int `json:"src_tcp_filter_action_default_pass"`
    Src_udp_filter_action_blacklist int `json:"src_udp_filter_action_blacklist"`
    Src_udp_filter_action_whitelist int `json:"src_udp_filter_action_whitelist"`
    Src_udp_filter_action_drop int `json:"src_udp_filter_action_drop"`
    Src_udp_filter_action_default_pass int `json:"src_udp_filter_action_default_pass"`
    Src_other_filter_action_blacklist int `json:"src_other_filter_action_blacklist"`
    Src_other_filter_action_whitelist int `json:"src_other_filter_action_whitelist"`
    Src_other_filter_action_drop int `json:"src_other_filter_action_drop"`
    Src_other_filter_action_default_pass int `json:"src_other_filter_action_default_pass"`
    Tcp_invalid_syn int `json:"tcp_invalid_syn"`
    Dst_tcp_conn_close_w_rst int `json:"dst_tcp_conn_close_w_rst"`
    Dst_tcp_conn_close_w_fin int `json:"dst_tcp_conn_close_w_fin"`
    Dst_tcp_conn_close_w_idle int `json:"dst_tcp_conn_close_w_idle"`
    Dst_tcp_conn_create_from_syn int `json:"dst_tcp_conn_create_from_syn"`
    Dst_tcp_conn_create_from_ack int `json:"dst_tcp_conn_create_from_ack"`
    Src_frag_drop int `json:"src_frag_drop"`
    Dst_l4_tcp_blacklist_drop int `json:"dst_l4_tcp_blacklist_drop"`
    Dst_l4_udp_blacklist_drop int `json:"dst_l4_udp_blacklist_drop"`
    Dst_l4_icmp_blacklist_drop int `json:"dst_l4_icmp_blacklist_drop"`
    Dst_l4_other_blacklist_drop int `json:"dst_l4_other_blacklist_drop"`
    Src_l4_tcp_blacklist_drop int `json:"src_l4_tcp_blacklist_drop"`
    Src_l4_udp_blacklist_drop int `json:"src_l4_udp_blacklist_drop"`
    Src_l4_icmp_blacklist_drop int `json:"src_l4_icmp_blacklist_drop"`
    Src_l4_other_blacklist_drop int `json:"src_l4_other_blacklist_drop"`
    Dst_port_kbit_rate_exceed_pkt int `json:"dst_port_kbit_rate_exceed_pkt"`
    Dst_tcp_bytes_rcv int `json:"dst_tcp_bytes_rcv"`
    Dst_udp_bytes_rcv int `json:"dst_udp_bytes_rcv"`
    Dst_icmp_bytes_rcv int `json:"dst_icmp_bytes_rcv"`
    Dst_other_bytes_rcv int `json:"dst_other_bytes_rcv"`
    Dst_tcp_bytes_sent int `json:"dst_tcp_bytes_sent"`
    Dst_udp_bytes_sent int `json:"dst_udp_bytes_sent"`
    Dst_icmp_bytes_sent int `json:"dst_icmp_bytes_sent"`
    Dst_other_bytes_sent int `json:"dst_other_bytes_sent"`
    Dst_udp_auth_drop int `json:"dst_udp_auth_drop"`
    Dst_tcp_auth_drop int `json:"dst_tcp_auth_drop"`
    Dst_tcp_auth_resp int `json:"dst_tcp_auth_resp"`
    Inbound_pkt_drop int `json:"inbound_pkt_drop"`
    Dst_entry_pkt_rate_exceed int `json:"dst_entry_pkt_rate_exceed"`
    Dst_entry_kbit_rate_exceed int `json:"dst_entry_kbit_rate_exceed"`
    Dst_entry_conn_limit_exceed int `json:"dst_entry_conn_limit_exceed"`
    Dst_entry_conn_rate_exceed int `json:"dst_entry_conn_rate_exceed"`
    Dst_entry_frag_pkt_rate_exceed int `json:"dst_entry_frag_pkt_rate_exceed"`
    Dst_icmp_any_exceed int `json:"dst_icmp_any_exceed"`
    Dst_other_any_exceed int `json:"dst_other_any_exceed"`
    Src_dst_pair_entry_total int `json:"src_dst_pair_entry_total"`
    Src_dst_pair_entry_udp int `json:"src_dst_pair_entry_udp"`
    Src_dst_pair_entry_tcp int `json:"src_dst_pair_entry_tcp"`
    Src_dst_pair_entry_icmp int `json:"src_dst_pair_entry_icmp"`
    Src_dst_pair_entry_other int `json:"src_dst_pair_entry_other"`
    Dst_clist_overflow_policy_at_learning int `json:"dst_clist_overflow_policy_at_learning"`
    Tcp_rexmit_syn_limit_drop int `json:"tcp_rexmit_syn_limit_drop"`
    Tcp_rexmit_syn_limit_bl int `json:"tcp_rexmit_syn_limit_bl"`
    Dst_tcp_wellknown_sport_drop int `json:"dst_tcp_wellknown_sport_drop"`
    Src_tcp_wellknown_sport_drop int `json:"src_tcp_wellknown_sport_drop"`
    Dst_frag_rcvd int `json:"dst_frag_rcvd"`
    No_policy_class_list_match int `json:"no_policy_class_list_match"`
    Src_udp_retry_gap_drop int `json:"src_udp_retry_gap_drop"`
    Dst_entry_kbit_rate_exceed_count int `json:"dst_entry_kbit_rate_exceed_count"`
    Dst_port_undef_hit int `json:"dst_port_undef_hit"`
    Dst_tcp_action_on_ack_timeout int `json:"dst_tcp_action_on_ack_timeout"`
    Dst_tcp_action_on_ack_reset int `json:"dst_tcp_action_on_ack_reset"`
    Dst_tcp_action_on_ack_blacklist int `json:"dst_tcp_action_on_ack_blacklist"`
    Src_tcp_action_on_ack_timeout int `json:"src_tcp_action_on_ack_timeout"`
    Src_tcp_action_on_ack_reset int `json:"src_tcp_action_on_ack_reset"`
    Src_tcp_action_on_ack_blacklist int `json:"src_tcp_action_on_ack_blacklist"`
    Dst_tcp_action_on_syn_timeout int `json:"dst_tcp_action_on_syn_timeout"`
    Dst_tcp_action_on_syn_reset int `json:"dst_tcp_action_on_syn_reset"`
    Dst_tcp_action_on_syn_blacklist int `json:"dst_tcp_action_on_syn_blacklist"`
    Src_tcp_action_on_syn_timeout int `json:"src_tcp_action_on_syn_timeout"`
    Src_tcp_action_on_syn_reset int `json:"src_tcp_action_on_syn_reset"`
    Src_tcp_action_on_syn_blacklist int `json:"src_tcp_action_on_syn_blacklist"`
    Dst_udp_frag_pkt_rate_exceed int `json:"dst_udp_frag_pkt_rate_exceed"`
    Dst_udp_frag_src_rate_drop int `json:"dst_udp_frag_src_rate_drop"`
    Dst_tcp_frag_pkt_rate_exceed int `json:"dst_tcp_frag_pkt_rate_exceed"`
    Dst_tcp_frag_src_rate_drop int `json:"dst_tcp_frag_src_rate_drop"`
    Dst_icmp_frag_pkt_rate_exceed int `json:"dst_icmp_frag_pkt_rate_exceed"`
    Dst_icmp_frag_src_rate_drop int `json:"dst_icmp_frag_src_rate_drop"`
    Sflow_internal_samples_packed int `json:"sflow_internal_samples_packed"`
    Sflow_external_samples_packed int `json:"sflow_external_samples_packed"`
    Sflow_internal_packets_sent int `json:"sflow_internal_packets_sent"`
    Sflow_external_packets_sent int `json:"sflow_external_packets_sent"`
    Dns_outbound_total_query int `json:"dns_outbound_total_query"`
    Dns_outbound_query_malformed int `json:"dns_outbound_query_malformed"`
    Dns_outbound_query_resp_chk_failed int `json:"dns_outbound_query_resp_chk_failed"`
    Dns_outbound_query_resp_chk_blacklisted int `json:"dns_outbound_query_resp_chk_blacklisted"`
    Dns_outbound_query_resp_chk_refused_sent int `json:"dns_outbound_query_resp_chk_refused_sent"`
    Dns_outbound_query_resp_chk_reset_sent int `json:"dns_outbound_query_resp_chk_reset_sent"`
    Dns_outbound_query_resp_chk_no_resp_sent int `json:"dns_outbound_query_resp_chk_no_resp_sent"`
    Dns_outbound_query_resp_size_exceed int `json:"dns_outbound_query_resp_size_exceed"`
    Dns_outbound_query_sess_timed_out int `json:"dns_outbound_query_sess_timed_out"`
    Dst_exceed_action_tunnel int `json:"dst_exceed_action_tunnel"`
    Src_udp_auth_timeout int `json:"src_udp_auth_timeout"`
    Src_udp_retry_pass int `json:"src_udp_retry_pass"`
    Dst_hw_drop_rule_insert int `json:"dst_hw_drop_rule_insert"`
    Dst_hw_drop_rule_remove int `json:"dst_hw_drop_rule_remove"`
    Src_hw_drop_rule_insert int `json:"src_hw_drop_rule_insert"`
    Src_hw_drop_rule_remove int `json:"src_hw_drop_rule_remove"`
    Prog_first_req_time_exceed int `json:"prog_first_req_time_exceed"`
    Prog_req_resp_time_exceed int `json:"prog_req_resp_time_exceed"`
    Prog_request_len_exceed int `json:"prog_request_len_exceed"`
    Prog_response_len_exceed int `json:"prog_response_len_exceed"`
    Prog_resp_req_ratio_exceed int `json:"prog_resp_req_ratio_exceed"`
    Prog_resp_req_time_exceed int `json:"prog_resp_req_time_exceed"`
    Entry_sync_message_received int `json:"entry_sync_message_received"`
    Entry_sync_message_sent int `json:"entry_sync_message_sent"`
    Prog_conn_sent_exceed int `json:"prog_conn_sent_exceed"`
    Prog_conn_rcvd_exceed int `json:"prog_conn_rcvd_exceed"`
    Prog_conn_time_exceed int `json:"prog_conn_time_exceed"`
    Prog_conn_rcvd_sent_ratio_exceed int `json:"prog_conn_rcvd_sent_ratio_exceed"`
    Prog_win_sent_exceed int `json:"prog_win_sent_exceed"`
    Prog_win_rcvd_exceed int `json:"prog_win_rcvd_exceed"`
    Prog_win_rcvd_sent_ratio_exceed int `json:"prog_win_rcvd_sent_ratio_exceed"`
    Prog_exceed_drop int `json:"prog_exceed_drop"`
    Prog_exceed_bl int `json:"prog_exceed_bl"`
    Prog_conn_exceed_drop int `json:"prog_conn_exceed_drop"`
    Prog_conn_exceed_bl int `json:"prog_conn_exceed_bl"`
    Prog_win_exceed_drop int `json:"prog_win_exceed_drop"`
    Prog_win_exceed_bl int `json:"prog_win_exceed_bl"`
    Dst_exceed_action_drop int `json:"dst_exceed_action_drop"`
    Src_hw_drop int `json:"src_hw_drop"`
    Dst_tcp_auth_rst int `json:"dst_tcp_auth_rst"`
    Dst_src_learn_overflow int `json:"dst_src_learn_overflow"`
    Tcp_fwd_sent int `json:"tcp_fwd_sent"`
    Udp_fwd_sent int `json:"udp_fwd_sent"`
}

func (p *DdosDstDynamicEntryAllEntriesStats) GetId() string{
    return "1"
}

func (p *DdosDstDynamicEntryAllEntriesStats) getPath() string{
    return "ddos/dst/dynamic-entry/all-entries/stats"
}

func (p *DdosDstDynamicEntryAllEntriesStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstDynamicEntryAllEntriesStats,error) {
logger.Println("DdosDstDynamicEntryAllEntriesStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstDynamicEntryAllEntriesStats
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
