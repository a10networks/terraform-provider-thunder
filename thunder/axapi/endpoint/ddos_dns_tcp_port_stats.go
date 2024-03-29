

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsTcpPortStats struct {
    
    Stats DdosDnsTcpPortStatsStats `json:"stats"`

}
type DataDdosDnsTcpPortStats struct {
    DtDdosDnsTcpPortStats DdosDnsTcpPortStats `json:"dns-tcp-port"`
}


type DdosDnsTcpPortStatsStats struct {
    Rate_limit_type0 int `json:"rate_limit_type0"`
    Rate_limit_type1 int `json:"rate_limit_type1"`
    Rate_limit_type2 int `json:"rate_limit_type2"`
    Rate_limit_type3 int `json:"rate_limit_type3"`
    Rate_limit_type4 int `json:"rate_limit_type4"`
    Nxdomain_rate_exceed int `json:"nxdomain_rate_exceed"`
    Is_nxdomain int `json:"is_nxdomain"`
    Nxdomain_bl_drop int `json:"nxdomain_bl_drop"`
    Dns_fqdn_label_len_exceed int `json:"dns_fqdn_label_len_exceed"`
    Dns_malform_drop int `json:"dns_malform_drop"`
    Dns_fqdn_stage_2_rate_exceed int `json:"dns_fqdn_stage_2_rate_exceed"`
    Req_sent int `json:"req_sent"`
    Req_incomplete int `json:"req_incomplete"`
    Req_size_exceed int `json:"req_size_exceed"`
    Req_retrans int `json:"req_retrans"`
    Query_type_a int `json:"query_type_a"`
    Query_type_aaaa int `json:"query_type_aaaa"`
    Query_type_ns int `json:"query_type_ns"`
    Query_type_cname int `json:"query_type_cname"`
    Query_type_any int `json:"query_type_any"`
    Query_type_srv int `json:"query_type_srv"`
    Query_type_mx int `json:"query_type_mx"`
    Query_type_soa int `json:"query_type_soa"`
    Query_type_opt int `json:"query_type_opt"`
    Port_rcvd int `json:"port_rcvd"`
    Port_drop int `json:"port_drop"`
    Port_pkt_sent int `json:"port_pkt_sent"`
    Port_pkt_rate_exceed int `json:"port_pkt_rate_exceed"`
    Port_kbit_rate_exceed int `json:"port_kbit_rate_exceed"`
    Port_conn_rate_exceed int `json:"port_conn_rate_exceed"`
    Port_conn_limm_exceed int `json:"port_conn_limm_exceed"`
    Dg_action_permit int `json:"dg_action_permit"`
    Dg_action_deny int `json:"dg_action_deny"`
    Dg_action_default int `json:"dg_action_default"`
    Port_bytes int `json:"port_bytes"`
    Outbound_port_bytes int `json:"outbound_port_bytes"`
    Outbound_port_rcvd int `json:"outbound_port_rcvd"`
    Outbound_port_pkt_sent int `json:"outbound_port_pkt_sent"`
    Port_bytes_sent int `json:"port_bytes_sent"`
    Port_bytes_drop int `json:"port_bytes_drop"`
    Port_src_bl int `json:"port_src_bl"`
    Filter_auth_fail int `json:"filter_auth_fail"`
    Syn_auth_fail int `json:"syn_auth_fail"`
    Ack_auth_fail int `json:"ack_auth_fail"`
    Syn_cookie_fail int `json:"syn_cookie_fail"`
    Sess_create int `json:"sess_create"`
    Filter_action_blacklist int `json:"filter_action_blacklist"`
    Filter_action_drop int `json:"filter_action_drop"`
    Filter_action_default_pass int `json:"filter_action_default_pass"`
    Filter_action_whitelist int `json:"filter_action_whitelist"`
    Exceed_drop_prate_src int `json:"exceed_drop_prate_src"`
    Exceed_drop_crate_src int `json:"exceed_drop_crate_src"`
    Exceed_drop_climit_src int `json:"exceed_drop_climit_src"`
    Exceed_drop_brate_src int `json:"exceed_drop_brate_src"`
    Outbound_port_bytes_sent int `json:"outbound_port_bytes_sent"`
    Outbound_port_drop int `json:"outbound_port_drop"`
    Outbound_port_bytes_drop int `json:"outbound_port_bytes_drop"`
    Src_rate_limit_type0 int `json:"src_rate_limit_type0"`
    Src_rate_limit_type1 int `json:"src_rate_limit_type1"`
    Src_rate_limit_type2 int `json:"src_rate_limit_type2"`
    Src_rate_limit_type3 int `json:"src_rate_limit_type3"`
    Src_rate_limit_type4 int `json:"src_rate_limit_type4"`
    Exceed_drop_brate_src_pkt int `json:"exceed_drop_brate_src_pkt"`
    Port_kbit_rate_exceed_pkt int `json:"port_kbit_rate_exceed_pkt"`
    Syn_cookie_sent int `json:"syn_cookie_sent"`
    Ack_retry_init int `json:"ack_retry_init"`
    Ack_retry_gap_drop int `json:"ack_retry_gap_drop"`
    Conn_prate_excd int `json:"conn_prate_excd"`
    Out_of_seq_excd int `json:"out_of_seq_excd"`
    Retransmit_excd int `json:"retransmit_excd"`
    Zero_window_excd int `json:"zero_window_excd"`
    Syn_retry_init int `json:"syn_retry_init"`
    Syn_retry_gap_drop int `json:"syn_retry_gap_drop"`
    Ack_retry_pass int `json:"ack_retry_pass"`
    Syn_retry_pass int `json:"syn_retry_pass"`
    Fqdn_rate_by_label_count_check int `json:"fqdn_rate_by_label_count_check"`
    Fqdn_rate_by_label_count_exceed int `json:"fqdn_rate_by_label_count_exceed"`
    Policy_drop int `json:"policy_drop"`
    Fqdn_label_count_exceed int `json:"fqdn_label_count_exceed"`
    Rrtype_drop int `json:"rrtype_drop"`
    Src_policy_drop int `json:"src_policy_drop"`
    Src_dns_fqdn_label_len_exceed int `json:"src_dns_fqdn_label_len_exceed"`
    Src_fqdn_label_count_exceed int `json:"src_fqdn_label_count_exceed"`
    Src_dns_malform_drop int `json:"src_dns_malform_drop"`
    Tcp_auth_drop int `json:"tcp_auth_drop"`
    Tcp_auth_resp int `json:"tcp_auth_resp"`
    Dns_query_class_in int `json:"dns_query_class_in"`
    Dns_query_class_csnet int `json:"dns_query_class_csnet"`
    Dns_query_class_chaos int `json:"dns_query_class_chaos"`
    Dns_query_class_hs int `json:"dns_query_class_hs"`
    Dns_query_class_none int `json:"dns_query_class_none"`
    Dns_query_class_any int `json:"dns_query_class_any"`
    Dns_query_class_whitelist_miss int `json:"dns_query_class_whitelist_miss"`
    Bl int `json:"bl"`
    Src_drop int `json:"src_drop"`
    Frag_rcvd int `json:"frag_rcvd"`
    Frag_drop int `json:"frag_drop"`
    Sess_create_inbound int `json:"sess_create_inbound"`
    Sess_create_outbound int `json:"sess_create_outbound"`
    Conn_create_from_syn int `json:"conn_create_from_syn"`
    Conn_create_from_ack int `json:"conn_create_from_ack"`
    Conn_close int `json:"conn_close"`
    Conn_close_w_rst int `json:"conn_close_w_rst"`
    Conn_close_w_fin int `json:"conn_close_w_fin"`
    Conn_close_w_idle int `json:"conn_close_w_idle"`
    Conn_close_half_open int `json:"conn_close_half_open"`
    Sess_aged int `json:"sess_aged"`
    Syn_drop int `json:"syn_drop"`
    Syn_auth_pass int `json:"syn_auth_pass"`
    Unauth_drop int `json:"unauth_drop"`
    Rst_cookie_fail int `json:"rst_cookie_fail"`
    Syn_retry_failed int `json:"syn_retry_failed"`
    Filter1_match int `json:"filter1_match"`
    Filter2_match int `json:"filter2_match"`
    Filter3_match int `json:"filter3_match"`
    Filter4_match int `json:"filter4_match"`
    Filter5_match int `json:"filter5_match"`
    Filter_none_match int `json:"filter_none_match"`
    Filter_total_not_match int `json:"filter_total_not_match"`
    Src_syn_auth_fail int `json:"src_syn_auth_fail"`
    Src_syn_cookie_sent int `json:"src_syn_cookie_sent"`
    Src_syn_cookie_fail int `json:"src_syn_cookie_fail"`
    Src_unauth_drop int `json:"src_unauth_drop"`
    Src_rst_cookie_fail int `json:"src_rst_cookie_fail"`
    Src_syn_retry_init int `json:"src_syn_retry_init"`
    Src_syn_retry_gap_drop int `json:"src_syn_retry_gap_drop"`
    Src_syn_retry_failed int `json:"src_syn_retry_failed"`
    Src_ack_retry_init int `json:"src_ack_retry_init"`
    Src_ack_retry_gap_drop int `json:"src_ack_retry_gap_drop"`
    Src_ack_auth_fail int `json:"src_ack_auth_fail"`
    Src_out_of_seq_excd int `json:"src_out_of_seq_excd"`
    Src_retransmit_excd int `json:"src_retransmit_excd"`
    Src_zero_window_excd int `json:"src_zero_window_excd"`
    Src_conn_pkt_rate_excd int `json:"src_conn_pkt_rate_excd"`
    Src_filter_action_blacklist int `json:"src_filter_action_blacklist"`
    Src_filter_action_drop int `json:"src_filter_action_drop"`
    Src_filter_action_default_pass int `json:"src_filter_action_default_pass"`
    Src_filter_action_whitelist int `json:"src_filter_action_whitelist"`
    Force_tcp_auth_pass int `json:"force_tcp_auth_pass"`
    Nxdomain_drop int `json:"nxdomain_drop"`
    Dns_malform int `json:"dns_malform"`
    Src_dns_malform int `json:"src_dns_malform"`
    Tcp_rexmit_syn_limit_drop int `json:"tcp_rexmit_syn_limit_drop"`
    Tcp_rexmit_syn_limit_bl int `json:"tcp_rexmit_syn_limit_bl"`
    Conn_ofo_rate_excd int `json:"conn_ofo_rate_excd"`
    Conn_rexmit_rate_excd int `json:"conn_rexmit_rate_excd"`
    Conn_zwindow_rate_excd int `json:"conn_zwindow_rate_excd"`
    Src_conn_ofo_rate_excd int `json:"src_conn_ofo_rate_excd"`
    Src_conn_rexmit_rate_excd int `json:"src_conn_rexmit_rate_excd"`
    Src_conn_zwindow_rate_excd int `json:"src_conn_zwindow_rate_excd"`
    Ack_retry_rto_pass int `json:"ack_retry_rto_pass"`
    Ack_retry_rto_fail int `json:"ack_retry_rto_fail"`
    Ack_retry_rto_progress int `json:"ack_retry_rto_progress"`
    Syn_retry_rto_pass int `json:"syn_retry_rto_pass"`
    Syn_retry_rto_fail int `json:"syn_retry_rto_fail"`
    Syn_retry_rto_progress int `json:"syn_retry_rto_progress"`
    Src_syn_retry_rto_pass int `json:"src_syn_retry_rto_pass"`
    Src_syn_retry_rto_fail int `json:"src_syn_retry_rto_fail"`
    Src_syn_retry_rto_progress int `json:"src_syn_retry_rto_progress"`
    Src_ack_retry_rto_pass int `json:"src_ack_retry_rto_pass"`
    Src_ack_retry_rto_fail int `json:"src_ack_retry_rto_fail"`
    Src_ack_retry_rto_progress int `json:"src_ack_retry_rto_progress"`
    Wellknown_sport_drop int `json:"wellknown_sport_drop"`
    Src_well_known_port int `json:"src_well_known_port"`
    Dg_request_check_fail int `json:"dg_request_check_fail"`
    Query_type_any_drop int `json:"query_type_any_drop"`
    Src_query_type_any_drop int `json:"src_query_type_any_drop"`
    Src_auth_drop int `json:"src_auth_drop"`
    Src_frag_drop int `json:"src_frag_drop"`
    Frag_timeout int `json:"frag_timeout"`
    Create_conn_non_syn_dropped int `json:"create_conn_non_syn_dropped"`
    Src_create_conn_non_syn_dropped int `json:"src_create_conn_non_syn_dropped"`
    Dg_rate_exceed int `json:"dg_rate_exceed"`
    Port_syn_rate_exceed int `json:"port_syn_rate_exceed"`
    Pattern_recognition_proceeded int `json:"pattern_recognition_proceeded"`
    Pattern_not_found int `json:"pattern_not_found"`
    Pattern_recognition_generic_error int `json:"pattern_recognition_generic_error"`
    Pattern_filter1_match int `json:"pattern_filter1_match"`
    Pattern_filter2_match int `json:"pattern_filter2_match"`
    Pattern_filter3_match int `json:"pattern_filter3_match"`
    Pattern_filter4_match int `json:"pattern_filter4_match"`
    Pattern_filter5_match int `json:"pattern_filter5_match"`
    Pattern_filter_drop int `json:"pattern_filter_drop"`
    Src_syn_rate_exceed int `json:"src_syn_rate_exceed"`
    Non_query_opcode_pass_through int `json:"non_query_opcode_pass_through"`
    Src_filter1_match int `json:"src_filter1_match"`
    Src_filter2_match int `json:"src_filter2_match"`
    Src_filter3_match int `json:"src_filter3_match"`
    Src_filter4_match int `json:"src_filter4_match"`
    Src_filter5_match int `json:"src_filter5_match"`
    Src_filter_none_match int `json:"src_filter_none_match"`
    Src_filter_total_not_match int `json:"src_filter_total_not_match"`
    Src_filter_auth_fail int `json:"src_filter_auth_fail"`
    Syn_tfo_rcv int `json:"syn_tfo_rcv"`
    Ack_retry_timeout int `json:"ack_retry_timeout"`
    Ack_retry_reset int `json:"ack_retry_reset"`
    Ack_retry_blacklist int `json:"ack_retry_blacklist"`
    Src_ack_retry_timeout int `json:"src_ack_retry_timeout"`
    Src_ack_retry_reset int `json:"src_ack_retry_reset"`
    Src_ack_retry_blacklist int `json:"src_ack_retry_blacklist"`
    Syn_retry_timeout int `json:"syn_retry_timeout"`
    Syn_retry_reset int `json:"syn_retry_reset"`
    Syn_retry_blacklist int `json:"syn_retry_blacklist"`
    Src_syn_retry_timeout int `json:"src_syn_retry_timeout"`
    Src_syn_retry_reset int `json:"src_syn_retry_reset"`
    Src_syn_retry_blacklist int `json:"src_syn_retry_blacklist"`
    Sflow_internal_samples_packed int `json:"sflow_internal_samples_packed"`
    Sflow_external_samples_packed int `json:"sflow_external_samples_packed"`
    Sflow_internal_packets_sent int `json:"sflow_internal_packets_sent"`
    Sflow_external_packets_sent int `json:"sflow_external_packets_sent"`
    Pattern_recognition_sampling_started int `json:"pattern_recognition_sampling_started"`
    Pattern_recognition_pattern_changed int `json:"pattern_recognition_pattern_changed"`
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Synack_reset_sent int `json:"synack_reset_sent"`
    Synack_multiple_attempts_per_ip_detected int `json:"synack_multiple_attempts_per_ip_detected"`
    Prog_first_req_time_exceed int `json:"prog_first_req_time_exceed"`
    Prog_req_resp_time_exceed int `json:"prog_req_resp_time_exceed"`
    Prog_request_len_exceed int `json:"prog_request_len_exceed"`
    Prog_response_len_exceed int `json:"prog_response_len_exceed"`
    Prog_resp_req_ratio_exceed int `json:"prog_resp_req_ratio_exceed"`
    Prog_resp_req_time_exceed int `json:"prog_resp_req_time_exceed"`
    Prog_conn_sent_exceed int `json:"prog_conn_sent_exceed"`
    Prog_conn_rcvd_exceed int `json:"prog_conn_rcvd_exceed"`
    Prog_conn_time_exceed int `json:"prog_conn_time_exceed"`
    Prog_conn_rcvd_sent_ratio_exceed int `json:"prog_conn_rcvd_sent_ratio_exceed"`
    Prog_win_sent_exceed int `json:"prog_win_sent_exceed"`
    Prog_win_rcvd_exceed int `json:"prog_win_rcvd_exceed"`
    Prog_win_rcvd_sent_ratio_exceed int `json:"prog_win_rcvd_sent_ratio_exceed"`
    Snat_fail int `json:"snat_fail"`
    Prog_exceed_drop int `json:"prog_exceed_drop"`
    Prog_exceed_bl int `json:"prog_exceed_bl"`
    Prog_conn_exceed_drop int `json:"prog_conn_exceed_drop"`
    Prog_conn_exceed_bl int `json:"prog_conn_exceed_bl"`
    Prog_win_exceed_drop int `json:"prog_win_exceed_drop"`
    Prog_win_exceed_bl int `json:"prog_win_exceed_bl"`
    Exceed_action_drop int `json:"exceed_action_drop"`
    Syn_auth_rst_ack_drop int `json:"syn_auth_rst_ack_drop"`
    Prog_exceed_reset int `json:"prog_exceed_reset"`
    Prog_conn_exceed_reset int `json:"prog_conn_exceed_reset"`
    Prog_win_exceed_reset int `json:"prog_win_exceed_reset"`
    Conn_create_from_synack int `json:"conn_create_from_synack"`
    Port_synack_rate_exceed int `json:"port_synack_rate_exceed"`
    Prog_conn_samples int `json:"prog_conn_samples"`
    Prog_req_samples int `json:"prog_req_samples"`
    Prog_win_samples int `json:"prog_win_samples"`
    Prog_conn_samples_processed int `json:"prog_conn_samples_processed"`
    Prog_req_samples_processed int `json:"prog_req_samples_processed"`
    Prog_win_samples_processed int `json:"prog_win_samples_processed"`
    Tcp_auth_rst int `json:"tcp_auth_rst"`
    Src_tcp_auth_rst int `json:"src_tcp_auth_rst"`
}

func (p *DdosDnsTcpPortStats) GetId() string{
    return "1"
}

func (p *DdosDnsTcpPortStats) getPath() string{
    return "ddos/dns-tcp-port/stats"
}

func (p *DdosDnsTcpPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDnsTcpPortStats,error) {
logger.Println("DdosDnsTcpPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDnsTcpPortStats
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
