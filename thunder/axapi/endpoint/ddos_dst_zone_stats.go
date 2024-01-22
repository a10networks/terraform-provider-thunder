

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneStats struct {
    
    Stats DdosDstZoneStatsStats `json:"stats"`
    ZoneName string `json:"zone-name"`

}
type DataDdosDstZoneStats struct {
    DtDdosDstZoneStats DdosDstZoneStats `json:"zone"`
}


type DdosDstZoneStatsStats struct {
    Zone_tcp_any_exceed int `json:"zone_tcp_any_exceed"`
    Zone_tcp_pkt_rate_exceed int `json:"zone_tcp_pkt_rate_exceed"`
    Zone_tcp_conn_rate_exceed int `json:"zone_tcp_conn_rate_exceed"`
    Zone_udp_any_exceed int `json:"zone_udp_any_exceed"`
    Zone_udp_pkt_rate_exceed int `json:"zone_udp_pkt_rate_exceed"`
    Zone_udp_conn_limit_exceed int `json:"zone_udp_conn_limit_exceed"`
    Zone_udp_conn_rate_exceed int `json:"zone_udp_conn_rate_exceed"`
    Zone_icmp_pkt_rate_exceed int `json:"zone_icmp_pkt_rate_exceed"`
    Zone_other_pkt_rate_exceed int `json:"zone_other_pkt_rate_exceed"`
    Zone_other_frag_pkt_rate_exceed int `json:"zone_other_frag_pkt_rate_exceed"`
    Zone_port_pkt_rate_exceed int `json:"zone_port_pkt_rate_exceed"`
    Zone_port_conn_limit_exceed int `json:"zone_port_conn_limit_exceed"`
    Zone_port_conn_rate_exceed int `json:"zone_port_conn_rate_exceed"`
    Zone_pkt_sent int `json:"zone_pkt_sent"`
    Zone_udp_pkt_sent int `json:"zone_udp_pkt_sent"`
    Zone_tcp_pkt_sent int `json:"zone_tcp_pkt_sent"`
    Zone_icmp_pkt_sent int `json:"zone_icmp_pkt_sent"`
    Zone_other_pkt_sent int `json:"zone_other_pkt_sent"`
    Zone_tcp_conn_limit_exceed int `json:"zone_tcp_conn_limit_exceed"`
    Zone_tcp_pkt_rcvd int `json:"zone_tcp_pkt_rcvd"`
    Zone_udp_pkt_rcvd int `json:"zone_udp_pkt_rcvd"`
    Zone_icmp_pkt_rcvd int `json:"zone_icmp_pkt_rcvd"`
    Zone_other_pkt_rcvd int `json:"zone_other_pkt_rcvd"`
    Zone_udp_filter_match int `json:"zone_udp_filter_match"`
    Zone_udp_filter_not_match int `json:"zone_udp_filter_not_match"`
    Zone_udp_filter_action_blacklist int `json:"zone_udp_filter_action_blacklist"`
    Zone_udp_filter_action_drop int `json:"zone_udp_filter_action_drop"`
    Zone_tcp_syn int `json:"zone_tcp_syn"`
    Zone_tcp_syn_drop int `json:"zone_tcp_syn_drop"`
    Zone_tcp_src_rate_drop int `json:"zone_tcp_src_rate_drop"`
    Zone_udp_src_rate_drop int `json:"zone_udp_src_rate_drop"`
    Zone_icmp_src_rate_drop int `json:"zone_icmp_src_rate_drop"`
    Zone_other_frag_src_rate_drop int `json:"zone_other_frag_src_rate_drop"`
    Zone_other_src_rate_drop int `json:"zone_other_src_rate_drop"`
    Zone_tcp_drop int `json:"zone_tcp_drop"`
    Zone_udp_drop int `json:"zone_udp_drop"`
    Zone_icmp_drop int `json:"zone_icmp_drop"`
    Zone_frag_drop int `json:"zone_frag_drop"`
    Zone_other_drop int `json:"zone_other_drop"`
    Zone_tcp_auth int `json:"zone_tcp_auth"`
    Zone_udp_filter_action_default_pass int `json:"zone_udp_filter_action_default_pass"`
    Zone_tcp_filter_match int `json:"zone_tcp_filter_match"`
    Zone_tcp_filter_not_match int `json:"zone_tcp_filter_not_match"`
    Zone_tcp_filter_action_blacklist int `json:"zone_tcp_filter_action_blacklist"`
    Zone_tcp_filter_action_drop int `json:"zone_tcp_filter_action_drop"`
    Zone_tcp_filter_action_default_pass int `json:"zone_tcp_filter_action_default_pass"`
    Zone_udp_filter_action_whitelist int `json:"zone_udp_filter_action_whitelist"`
    Zone_udp_kibit_rate_drop int `json:"zone_udp_kibit_rate_drop"`
    Zone_tcp_kibit_rate_drop int `json:"zone_tcp_kibit_rate_drop"`
    Zone_icmp_kibit_rate_drop int `json:"zone_icmp_kibit_rate_drop"`
    Zone_other_kibit_rate_drop int `json:"zone_other_kibit_rate_drop"`
    Zone_port_undef_drop int `json:"zone_port_undef_drop"`
    Zone_port_bl int `json:"zone_port_bl"`
    Zone_src_port_bl int `json:"zone_src_port_bl"`
    Zone_port_kbit_rate_exceed int `json:"zone_port_kbit_rate_exceed"`
    Zone_tcp_src_drop int `json:"zone_tcp_src_drop"`
    Zone_udp_src_drop int `json:"zone_udp_src_drop"`
    Zone_icmp_src_drop int `json:"zone_icmp_src_drop"`
    Zone_other_src_drop int `json:"zone_other_src_drop"`
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
    Zone_tcp_session_created int `json:"zone_tcp_session_created"`
    Zone_udp_session_created int `json:"zone_udp_session_created"`
    Zone_tcp_filter_action_whitelist int `json:"zone_tcp_filter_action_whitelist"`
    Zone_other_filter_match int `json:"zone_other_filter_match"`
    Zone_other_filter_not_match int `json:"zone_other_filter_not_match"`
    Zone_other_filter_action_blacklist int `json:"zone_other_filter_action_blacklist"`
    Zone_other_filter_action_drop int `json:"zone_other_filter_action_drop"`
    Zone_other_filter_action_whitelist int `json:"zone_other_filter_action_whitelist"`
    Zone_other_filter_action_default_pass int `json:"zone_other_filter_action_default_pass"`
    Zone_blackhole_inject int `json:"zone_blackhole_inject"`
    Zone_blackhole_withdraw int `json:"zone_blackhole_withdraw"`
    Zone_tcp_out_of_seq_excd int `json:"zone_tcp_out_of_seq_excd"`
    Zone_tcp_retransmit_excd int `json:"zone_tcp_retransmit_excd"`
    Zone_tcp_zero_window_excd int `json:"zone_tcp_zero_window_excd"`
    Zone_tcp_conn_prate_excd int `json:"zone_tcp_conn_prate_excd"`
    Zone_tcp_action_on_ack_init int `json:"zone_tcp_action_on_ack_init"`
    Zone_tcp_action_on_ack_gap_drop int `json:"zone_tcp_action_on_ack_gap_drop"`
    Zone_tcp_action_on_ack_fail int `json:"zone_tcp_action_on_ack_fail"`
    Zone_tcp_action_on_ack_pass int `json:"zone_tcp_action_on_ack_pass"`
    Zone_tcp_action_on_syn_init int `json:"zone_tcp_action_on_syn_init"`
    Zone_tcp_action_on_syn_gap_drop int `json:"zone_tcp_action_on_syn_gap_drop"`
    Zone_tcp_action_on_syn_fail int `json:"zone_tcp_action_on_syn_fail"`
    Zone_tcp_action_on_syn_pass int `json:"zone_tcp_action_on_syn_pass"`
    Zone_payload_too_small int `json:"zone_payload_too_small"`
    Zone_payload_too_big int `json:"zone_payload_too_big"`
    Zone_udp_conn_prate_excd int `json:"zone_udp_conn_prate_excd"`
    Zone_udp_ntp_monlist_req int `json:"zone_udp_ntp_monlist_req"`
    Zone_udp_ntp_monlist_resp int `json:"zone_udp_ntp_monlist_resp"`
    Zone_udp_wellknown_sport_drop int `json:"zone_udp_wellknown_sport_drop"`
    Zone_udp_retry_init int `json:"zone_udp_retry_init"`
    Zone_udp_retry_pass int `json:"zone_udp_retry_pass"`
    Zone_tcp_bytes_drop int `json:"zone_tcp_bytes_drop"`
    Zone_udp_bytes_drop int `json:"zone_udp_bytes_drop"`
    Zone_icmp_bytes_drop int `json:"zone_icmp_bytes_drop"`
    Zone_other_bytes_drop int `json:"zone_other_bytes_drop"`
    Zone_out_no_route int `json:"zone_out_no_route"`
    Outbound_bytes_sent int `json:"outbound_bytes_sent"`
    Outbound_drop int `json:"outbound_drop"`
    Outbound_bytes_drop int `json:"outbound_bytes_drop"`
    Outbound_pkt_sent int `json:"outbound_pkt_sent"`
    Inbound_bytes_sent int `json:"inbound_bytes_sent"`
    Inbound_bytes_drop int `json:"inbound_bytes_drop"`
    Zone_src_port_pkt_rate_exceed int `json:"zone_src_port_pkt_rate_exceed"`
    Zone_src_port_kbit_rate_exceed int `json:"zone_src_port_kbit_rate_exceed"`
    Zone_src_port_conn_limit_exceed int `json:"zone_src_port_conn_limit_exceed"`
    Zone_src_port_conn_rate_exceed int `json:"zone_src_port_conn_rate_exceed"`
    Zone_ip_proto_pkt_rate_exceed int `json:"zone_ip_proto_pkt_rate_exceed"`
    Zone_ip_proto_kbit_rate_exceed int `json:"zone_ip_proto_kbit_rate_exceed"`
    Zone_tcp_port_any_exceed int `json:"zone_tcp_port_any_exceed"`
    Zone_udp_port_any_exceed int `json:"zone_udp_port_any_exceed"`
    Zone_tcp_auth_pass int `json:"zone_tcp_auth_pass"`
    Zone_tcp_rst_cookie_fail int `json:"zone_tcp_rst_cookie_fail"`
    Zone_tcp_unauth_drop int `json:"zone_tcp_unauth_drop"`
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
    Zone_port_kbit_rate_exceed_pkt int `json:"zone_port_kbit_rate_exceed_pkt"`
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
    Dst_drop int `json:"dst_drop"`
    Dst_entry_pkt_rate_exceed int `json:"dst_entry_pkt_rate_exceed"`
    Dst_entry_kbit_rate_exceed int `json:"dst_entry_kbit_rate_exceed"`
    Dst_entry_conn_limit_exceed int `json:"dst_entry_conn_limit_exceed"`
    Dst_entry_conn_rate_exceed int `json:"dst_entry_conn_rate_exceed"`
    Dst_entry_frag_pkt_rate_exceed int `json:"dst_entry_frag_pkt_rate_exceed"`
    Dst_l4_tcp_blacklist_drop int `json:"dst_l4_tcp_blacklist_drop"`
    Dst_l4_udp_blacklist_drop int `json:"dst_l4_udp_blacklist_drop"`
    Dst_l4_icmp_blacklist_drop int `json:"dst_l4_icmp_blacklist_drop"`
    Dst_l4_other_blacklist_drop int `json:"dst_l4_other_blacklist_drop"`
    Dst_icmp_any_exceed int `json:"dst_icmp_any_exceed"`
    Dst_other_any_exceed int `json:"dst_other_any_exceed"`
    Tcp_rexmit_syn_limit_drop int `json:"tcp_rexmit_syn_limit_drop"`
    Tcp_rexmit_syn_limit_bl int `json:"tcp_rexmit_syn_limit_bl"`
    Dst_clist_overflow_policy_at_learning int `json:"dst_clist_overflow_policy_at_learning"`
    Zone_frag_rcvd int `json:"zone_frag_rcvd"`
    Zone_tcp_wellknown_sport_drop int `json:"zone_tcp_wellknown_sport_drop"`
    Src_tcp_wellknown_sport_drop int `json:"src_tcp_wellknown_sport_drop"`
    Secondary_dst_entry_pkt_rate_exceed int `json:"secondary_dst_entry_pkt_rate_exceed"`
    Secondary_dst_entry_kbit_rate_exceed int `json:"secondary_dst_entry_kbit_rate_exceed"`
    Secondary_dst_entry_conn_limit_exceed int `json:"secondary_dst_entry_conn_limit_exceed"`
    Secondary_dst_entry_conn_rate_exceed int `json:"secondary_dst_entry_conn_rate_exceed"`
    Secondary_dst_entry_frag_pkt_rate_exceed int `json:"secondary_dst_entry_frag_pkt_rate_exceed"`
    Src_udp_retry_gap_drop int `json:"src_udp_retry_gap_drop"`
    Dst_entry_kbit_rate_exceed_count int `json:"dst_entry_kbit_rate_exceed_count"`
    Secondary_entry_learn int `json:"secondary_entry_learn"`
    Secondary_entry_hit int `json:"secondary_entry_hit"`
    Secondary_entry_miss int `json:"secondary_entry_miss"`
    Secondary_entry_aged int `json:"secondary_entry_aged"`
    Secondary_entry_learning_thre_exceed int `json:"secondary_entry_learning_thre_exceed"`
    Zone_port_undef_hit int `json:"zone_port_undef_hit"`
    Zone_tcp_action_on_ack_timeout int `json:"zone_tcp_action_on_ack_timeout"`
    Zone_tcp_action_on_ack_reset int `json:"zone_tcp_action_on_ack_reset"`
    Zone_tcp_action_on_ack_blacklist int `json:"zone_tcp_action_on_ack_blacklist"`
    Src_tcp_action_on_ack_timeout int `json:"src_tcp_action_on_ack_timeout"`
    Src_tcp_action_on_ack_reset int `json:"src_tcp_action_on_ack_reset"`
    Src_tcp_action_on_ack_blacklist int `json:"src_tcp_action_on_ack_blacklist"`
    Zone_tcp_action_on_syn_timeout int `json:"zone_tcp_action_on_syn_timeout"`
    Zone_tcp_action_on_syn_reset int `json:"zone_tcp_action_on_syn_reset"`
    Zone_tcp_action_on_syn_blacklist int `json:"zone_tcp_action_on_syn_blacklist"`
    Src_tcp_action_on_syn_timeout int `json:"src_tcp_action_on_syn_timeout"`
    Src_tcp_action_on_syn_reset int `json:"src_tcp_action_on_syn_reset"`
    Src_tcp_action_on_syn_blacklist int `json:"src_tcp_action_on_syn_blacklist"`
    Zone_udp_frag_pkt_rate_exceed int `json:"zone_udp_frag_pkt_rate_exceed"`
    Zone_udp_frag_src_rate_drop int `json:"zone_udp_frag_src_rate_drop"`
    Zone_tcp_frag_pkt_rate_exceed int `json:"zone_tcp_frag_pkt_rate_exceed"`
    Zone_tcp_frag_src_rate_drop int `json:"zone_tcp_frag_src_rate_drop"`
    Zone_icmp_frag_pkt_rate_exceed int `json:"zone_icmp_frag_pkt_rate_exceed"`
    Zone_icmp_frag_src_rate_drop int `json:"zone_icmp_frag_src_rate_drop"`
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
    Source_entry_total int `json:"source_entry_total"`
    Source_entry_udp int `json:"source_entry_udp"`
    Source_entry_tcp int `json:"source_entry_tcp"`
    Source_entry_icmp int `json:"source_entry_icmp"`
    Source_entry_other int `json:"source_entry_other"`
    Dst_exceed_action_tunnel int `json:"dst_exceed_action_tunnel"`
    Dst_udp_retry_timeout_blacklist int `json:"dst_udp_retry_timeout_blacklist"`
    Src_udp_auth_timeout int `json:"src_udp_auth_timeout"`
    Zone_src_udp_retry_timeout_blacklist int `json:"zone_src_udp_retry_timeout_blacklist"`
    Src_udp_retry_pass int `json:"src_udp_retry_pass"`
    Secondary_port_learn int `json:"secondary_port_learn"`
    Secondary_port_aged int `json:"secondary_port_aged"`
    Dst_entry_outbound_udp_session_created int `json:"dst_entry_outbound_udp_session_created"`
    Dst_entry_outbound_udp_session_aged int `json:"dst_entry_outbound_udp_session_aged"`
    Dst_entry_outbound_tcp_session_created int `json:"dst_entry_outbound_tcp_session_created"`
    Dst_entry_outbound_tcp_session_aged int `json:"dst_entry_outbound_tcp_session_aged"`
    Dst_entry_outbound_pkt_rate_exceed int `json:"dst_entry_outbound_pkt_rate_exceed"`
    Dst_entry_outbound_kbit_rate_exceed int `json:"dst_entry_outbound_kbit_rate_exceed"`
    Dst_entry_outbound_kbit_rate_exceed_count int `json:"dst_entry_outbound_kbit_rate_exceed_count"`
    Dst_entry_outbound_conn_limit_exceed int `json:"dst_entry_outbound_conn_limit_exceed"`
    Dst_entry_outbound_conn_rate_exceed int `json:"dst_entry_outbound_conn_rate_exceed"`
    Dst_entry_outbound_frag_pkt_rate_exceed int `json:"dst_entry_outbound_frag_pkt_rate_exceed"`
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
    East_west_inbound_rcv_pkt int `json:"east_west_inbound_rcv_pkt"`
    East_west_inbound_drop_pkt int `json:"east_west_inbound_drop_pkt"`
    East_west_inbound_fwd_pkt int `json:"east_west_inbound_fwd_pkt"`
    East_west_inbound_rcv_byte int `json:"east_west_inbound_rcv_byte"`
    East_west_inbound_drop_byte int `json:"east_west_inbound_drop_byte"`
    East_west_inbound_fwd_byte int `json:"east_west_inbound_fwd_byte"`
    East_west_outbound_rcv_pkt int `json:"east_west_outbound_rcv_pkt"`
    East_west_outbound_drop_pkt int `json:"east_west_outbound_drop_pkt"`
    East_west_outbound_fwd_pkt int `json:"east_west_outbound_fwd_pkt"`
    East_west_outbound_rcv_byte int `json:"east_west_outbound_rcv_byte"`
    East_west_outbound_drop_byte int `json:"east_west_outbound_drop_byte"`
    East_west_outbound_fwd_byte int `json:"east_west_outbound_fwd_byte"`
    Dst_exceed_action_drop int `json:"dst_exceed_action_drop"`
    Prog_conn_samples int `json:"prog_conn_samples"`
    Prog_req_samples int `json:"prog_req_samples"`
    Prog_win_samples int `json:"prog_win_samples"`
    Victim_ip_learned int `json:"victim_ip_learned"`
    Victim_ip_aged int `json:"victim_ip_aged"`
    Prog_conn_samples_processed int `json:"prog_conn_samples_processed"`
    Prog_req_samples_processed int `json:"prog_req_samples_processed"`
    Prog_win_samples_processed int `json:"prog_win_samples_processed"`
    Dst_src_learn_overflow int `json:"dst_src_learn_overflow"`
    Dst_tcp_auth_rst int `json:"dst_tcp_auth_rst"`
}

func (p *DdosDstZoneStats) GetId() string{
    return "1"
}

func (p *DdosDstZoneStats) getPath() string{
    
    return "ddos/dst/zone/"+p.ZoneName+"/stats"
}

func (p *DdosDstZoneStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneStats,error) {
logger.Println("DdosDstZoneStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneStats
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
