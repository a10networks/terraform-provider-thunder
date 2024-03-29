

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosHttpPortStats struct {
    
    Stats DdosHttpPortStatsStats `json:"stats"`

}
type DataDdosHttpPortStats struct {
    DtDdosHttpPortStats DdosHttpPortStats `json:"http-port"`
}


type DdosHttpPortStatsStats struct {
    Ofo_timer_expired int `json:"ofo_timer_expired"`
    Partial_header int `json:"partial_header"`
    Slow_post int `json:"slow_post"`
    Idle_timeout int `json:"idle_timeout"`
    Pkts_ofo int `json:"pkts_ofo"`
    Pkts_retrans_fin int `json:"pkts_retrans_fin"`
    Pkts_retrans_rst int `json:"pkts_retrans_rst"`
    Pkts_retrans_push int `json:"pkts_retrans_push"`
    Pkts_retrans int `json:"pkts_retrans"`
    Chunk_bad int `json:"chunk_bad"`
    Chunk_sz_512 int `json:"chunk_sz_512"`
    Chunk_sz_1k int `json:"chunk_sz_1k"`
    Chunk_sz_2k int `json:"chunk_sz_2k"`
    Chunk_sz_4k int `json:"chunk_sz_4k"`
    Chunk_sz_gt_4k int `json:"chunk_sz_gt_4k"`
    Negative_resp_remain int `json:"negative_resp_remain"`
    Too_many_headers int `json:"too_many_headers"`
    Header_name_too_long int `json:"header_name_too_long"`
    Req_http11 int `json:"req_http11"`
    Req_http10 int `json:"req_http10"`
    Req_get int `json:"req_get"`
    Req_head int `json:"req_head"`
    Req_put int `json:"req_put"`
    Req_post int `json:"req_post"`
    Req_trace int `json:"req_trace"`
    Req_options int `json:"req_options"`
    Req_connect int `json:"req_connect"`
    Req_delete int `json:"req_delete"`
    Req_unknown int `json:"req_unknown"`
    Line_too_long int `json:"line_too_long"`
    Req_content_len int `json:"req_content_len"`
    Rsp_chunk int `json:"rsp_chunk"`
    Req int `json:"req"`
    Negative_req_remain int `json:"negative_req_remain"`
    Client_rst int `json:"client_rst"`
    Parsereq_fail int `json:"parsereq_fail"`
    Req_retran int `json:"req_retran"`
    Req_ofo int `json:"req_ofo"`
    Invalid_header int `json:"invalid_header"`
    Policy_violation int `json:"policy_violation"`
    Less_than_mss_exceed int `json:"less_than_mss_exceed"`
    Dst_req_rate_exceed int `json:"dst_req_rate_exceed"`
    Src_req_rate_exceed int `json:"src_req_rate_exceed"`
    Processed int `json:"processed"`
    New_syn int `json:"new_syn"`
    Policy_drop int `json:"policy_drop"`
    Er_condition int `json:"er_condition"`
    Ofo_queue_sz_exceed int `json:"ofo_queue_sz_exceed"`
    Fail_alloc int `json:"fail_alloc"`
    Fail_get_line int `json:"fail_get_line"`
    Fail_alloc_hdr int `json:"fail_alloc_hdr"`
    Invalid_hdr_name int `json:"invalid_hdr_name"`
    Invalid_hdr_val int `json:"invalid_hdr_val"`
    Ud_challenge_failed int `json:"ud_challenge_failed"`
    Js_challenge_failed int `json:"js_challenge_failed"`
    Challenge_failed int `json:"challenge_failed"`
    Js_challenge_sent int `json:"js_challenge_sent"`
    Ud_challenge_sent int `json:"ud_challenge_sent"`
    Parse_bad_chunk int `json:"parse_bad_chunk"`
    Content_length_too_long int `json:"content_length_too_long"`
    Parse_too_many_headers int `json:"parse_too_many_headers"`
    Parse_header_name_too_long int `json:"parse_header_name_too_long"`
    Parse_line_too_long int `json:"parse_line_too_long"`
    Parse_req_line_too_long int `json:"parse_req_line_too_long"`
    Window_small int `json:"window_small"`
    Window_small_drop int `json:"window_small_drop"`
    Use_hdr_src_ip int `json:"use_hdr_src_ip"`
    User_agent_filter_match int `json:"user_agent_filter_match"`
    User_agent_filter_blacklist int `json:"user_agent_filter_blacklist"`
    Referer_filter_match int `json:"referer_filter_match"`
    Referer_filter_blacklist int `json:"referer_filter_blacklist"`
    Req_sz_1k int `json:"req_sz_1k"`
    Req_sz_2k int `json:"req_sz_2k"`
    Req_sz_4k int `json:"req_sz_4k"`
    Req_sz_8k int `json:"req_sz_8k"`
    Req_sz_16k int `json:"req_sz_16k"`
    Req_sz_32k int `json:"req_sz_32k"`
    Req_sz_64k int `json:"req_sz_64k"`
    Req_sz_256k int `json:"req_sz_256k"`
    Req_sz_gt_256k int `json:"req_sz_gt_256k"`
    Filter_match int `json:"filter_match"`
    Filter_not_match int `json:"filter_not_match"`
    Filter_action_blacklist int `json:"filter_action_blacklist"`
    Filter_action_drop int `json:"filter_action_drop"`
    Filter_action_default_pass int `json:"filter_action_default_pass"`
    Dst_post_rate_exceed int `json:"dst_post_rate_exceed"`
    Src_post_rate_exceed int `json:"src_post_rate_exceed"`
    Dst_resp_rate_exceed int `json:"dst_resp_rate_exceed"`
    Header_processing_time_0 int `json:"header_processing_time_0"`
    Header_processing_time_1 int `json:"header_processing_time_1"`
    Header_processing_time_2 int `json:"header_processing_time_2"`
    Header_processing_time_3 int `json:"header_processing_time_3"`
    Header_processing_incomplete int `json:"header_processing_incomplete"`
    Resp_code_1xx int `json:"resp_code_1xx"`
    Resp_code_2xx int `json:"resp_code_2xx"`
    Resp_code_3xx int `json:"resp_code_3xx"`
    Resp_code_4xx int `json:"resp_code_4xx"`
    Resp_code_5xx int `json:"resp_code_5xx"`
    Resp_code_other int `json:"resp_code_other"`
    Filter_action_whitelist int `json:"filter_action_whitelist"`
    Filter1_match int `json:"filter1_match"`
    Filter2_match int `json:"filter2_match"`
    Filter3_match int `json:"filter3_match"`
    Filter4_match int `json:"filter4_match"`
    Filter5_match int `json:"filter5_match"`
    Filter_none_match int `json:"filter_none_match"`
    Port_rcvd int `json:"port_rcvd"`
    Port_drop int `json:"port_drop"`
    Port_pkt_sent int `json:"port_pkt_sent"`
    Port_pkt_rate_exceed int `json:"port_pkt_rate_exceed"`
    Port_kbit_rate_exceed int `json:"port_kbit_rate_exceed"`
    Port_conn_rate_exceed int `json:"port_conn_rate_exceed"`
    Port_conn_limm_exceed int `json:"port_conn_limm_exceed"`
    Filter1_drop int `json:"filter1_drop"`
    Filter2_drop int `json:"filter2_drop"`
    Filter3_drop int `json:"filter3_drop"`
    Filter4_drop int `json:"filter4_drop"`
    Filter5_drop int `json:"filter5_drop"`
    Filter6_match int `json:"filter6_match"`
    Filter6_drop int `json:"filter6_drop"`
    Filter7_match int `json:"filter7_match"`
    Filter7_drop int `json:"filter7_drop"`
    Filter8_match int `json:"filter8_match"`
    Filter8_drop int `json:"filter8_drop"`
    Filter9_match int `json:"filter9_match"`
    Filter9_drop int `json:"filter9_drop"`
    Filter10_match int `json:"filter10_match"`
    Filter10_drop int `json:"filter10_drop"`
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
    Exceed_drop_prate_src int `json:"exceed_drop_prate_src"`
    Exceed_drop_crate_src int `json:"exceed_drop_crate_src"`
    Exceed_drop_climit_src int `json:"exceed_drop_climit_src"`
    Exceed_drop_brate_src int `json:"exceed_drop_brate_src"`
    Outbound_port_bytes_sent int `json:"outbound_port_bytes_sent"`
    Outbound_port_drop int `json:"outbound_port_drop"`
    Outbound_port_bytes_drop int `json:"outbound_port_bytes_drop"`
    Uri_filter_match int `json:"uri_filter_match"`
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
    Tcp_filter_action_blacklist int `json:"tcp_filter_action_blacklist"`
    Tcp_filter_action_drop int `json:"tcp_filter_action_drop"`
    Src_policy_drop int `json:"src_policy_drop"`
    Src_dos_policy_violation int `json:"src_dos_policy_violation"`
    Src_challenge_ud_fail int `json:"src_challenge_ud_fail"`
    Src_challenge_js_fail int `json:"src_challenge_js_fail"`
    Src_js_challenge_sent int `json:"src_js_challenge_sent"`
    Src_ud_challenge_sent int `json:"src_ud_challenge_sent"`
    Tcp_auth_drop int `json:"tcp_auth_drop"`
    Tcp_auth_resp int `json:"tcp_auth_resp"`
    Http_auth_drop int `json:"http_auth_drop"`
    Http_auth_resp int `json:"http_auth_resp"`
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
    Tcp_filter1_match int `json:"tcp_filter1_match"`
    Tcp_filter2_match int `json:"tcp_filter2_match"`
    Tcp_filter3_match int `json:"tcp_filter3_match"`
    Tcp_filter4_match int `json:"tcp_filter4_match"`
    Tcp_filter5_match int `json:"tcp_filter5_match"`
    Tcp_filter_none_match int `json:"tcp_filter_none_match"`
    Tcp_filter_total_not_match int `json:"tcp_filter_total_not_match"`
    Tcp_filter_action_default_pass int `json:"tcp_filter_action_default_pass"`
    Tcp_filter_action_whitelist int `json:"tcp_filter_action_whitelist"`
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
    Resp_mib_chunk_bad int `json:"resp_mib_chunk_bad"`
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
    Src_http_auth_drop int `json:"src_http_auth_drop"`
    Src_auth_drop int `json:"src_auth_drop"`
    Src_frag_drop int `json:"src_frag_drop"`
    Frag_timeout int `json:"frag_timeout"`
    Create_conn_non_syn_dropped int `json:"create_conn_non_syn_dropped"`
    Src_create_conn_non_syn_dropped int `json:"src_create_conn_non_syn_dropped"`
    Pkts_connect_passthru int `json:"pkts_connect_passthru"`
    Parse_req_line_too_small int `json:"parse_req_line_too_small"`
    Parse_req_line_invalid_method int `json:"parse_req_line_invalid_method"`
    Port_syn_rate_exceed int `json:"port_syn_rate_exceed"`
    Src_syn_rate_exceed int `json:"src_syn_rate_exceed"`
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
    Src_header_filter_match int `json:"src_header_filter_match"`
    Src_header_filter_not_match int `json:"src_header_filter_not_match"`
    Src_header_filter_action_blacklist int `json:"src_header_filter_action_blacklist"`
    Src_header_filter_action_drop int `json:"src_header_filter_action_drop"`
    Src_header_filter_action_default_pass int `json:"src_header_filter_action_default_pass"`
    Src_header_filter_action_whitelist int `json:"src_header_filter_action_whitelist"`
    Src_header_filter_none_match int `json:"src_header_filter_none_match"`
    Src_header_filter1_match int `json:"src_header_filter1_match"`
    Src_header_filter2_match int `json:"src_header_filter2_match"`
    Src_header_filter3_match int `json:"src_header_filter3_match"`
    Src_header_filter4_match int `json:"src_header_filter4_match"`
    Src_header_filter5_match int `json:"src_header_filter5_match"`
    Src_header_filter6_match int `json:"src_header_filter6_match"`
    Src_header_filter7_match int `json:"src_header_filter7_match"`
    Src_header_filter8_match int `json:"src_header_filter8_match"`
    Src_header_filter9_match int `json:"src_header_filter9_match"`
    Src_header_filter10_match int `json:"src_header_filter10_match"`
    Src_header_filter1_drop int `json:"src_header_filter1_drop"`
    Src_header_filter2_drop int `json:"src_header_filter2_drop"`
    Src_header_filter3_drop int `json:"src_header_filter3_drop"`
    Src_header_filter4_drop int `json:"src_header_filter4_drop"`
    Src_header_filter5_drop int `json:"src_header_filter5_drop"`
    Src_header_filter6_drop int `json:"src_header_filter6_drop"`
    Src_header_filter7_drop int `json:"src_header_filter7_drop"`
    Src_header_filter8_drop int `json:"src_header_filter8_drop"`
    Src_header_filter9_drop int `json:"src_header_filter9_drop"`
    Src_header_filter10_drop int `json:"src_header_filter10_drop"`
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Pattern_recognition_proceeded int `json:"pattern_recognition_proceeded"`
    Pattern_not_found int `json:"pattern_not_found"`
    Pattern_recognition_generic_error int `json:"pattern_recognition_generic_error"`
    Pattern_filter1_match int `json:"pattern_filter1_match"`
    Pattern_filter2_match int `json:"pattern_filter2_match"`
    Pattern_filter3_match int `json:"pattern_filter3_match"`
    Pattern_filter4_match int `json:"pattern_filter4_match"`
    Pattern_filter5_match int `json:"pattern_filter5_match"`
    Pattern_filter_drop int `json:"pattern_filter_drop"`
    Pattern_recognition_sampling_started int `json:"pattern_recognition_sampling_started"`
    Pattern_recognition_pattern_changed int `json:"pattern_recognition_pattern_changed"`
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

func (p *DdosHttpPortStats) GetId() string{
    return "1"
}

func (p *DdosHttpPortStats) getPath() string{
    return "ddos/http-port/stats"
}

func (p *DdosHttpPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosHttpPortStats,error) {
logger.Println("DdosHttpPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosHttpPortStats
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
