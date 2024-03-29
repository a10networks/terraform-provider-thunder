

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsUdpPortStats struct {
    
    Stats DdosDnsUdpPortStatsStats `json:"stats"`

}
type DataDdosDnsUdpPortStats struct {
    DtDdosDnsUdpPortStats DdosDnsUdpPortStats `json:"dns-udp-port"`
}


type DdosDnsUdpPortStatsStats struct {
    Force_tcp_auth int `json:"force_tcp_auth"`
    Udp_auth int `json:"udp_auth"`
    Rate_limit_type0 int `json:"rate_limit_type0"`
    Rate_limit_type1 int `json:"rate_limit_type1"`
    Rate_limit_type2 int `json:"rate_limit_type2"`
    Rate_limit_type3 int `json:"rate_limit_type3"`
    Rate_limit_type4 int `json:"rate_limit_type4"`
    Nxdomain_bl_drop int `json:"nxdomain_bl_drop"`
    Is_nxdomain int `json:"is_nxdomain"`
    Nxdomain_drop int `json:"nxdomain_drop"`
    Udp_auth_fail int `json:"udp_auth_fail"`
    Dns_malform_drop int `json:"dns_malform_drop"`
    Fqdn_stage_2_rate_exceed int `json:"fqdn_stage_2_rate_exceed"`
    Req_sent int `json:"req_sent"`
    Req_size_exceed int `json:"req_size_exceed"`
    Req_retrans int `json:"req_retrans"`
    Fqdn_label_len_exceed int `json:"fqdn_label_len_exceed"`
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
    Spoof_detect_fail int `json:"spoof_detect_fail"`
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
    Udp_auth_retry_fail int `json:"udp_auth_retry_fail"`
    Src_rate_limit_type0 int `json:"src_rate_limit_type0"`
    Src_rate_limit_type1 int `json:"src_rate_limit_type1"`
    Src_rate_limit_type2 int `json:"src_rate_limit_type2"`
    Src_rate_limit_type3 int `json:"src_rate_limit_type3"`
    Src_rate_limit_type4 int `json:"src_rate_limit_type4"`
    Exceed_drop_brate_src_pkt int `json:"exceed_drop_brate_src_pkt"`
    Port_kbit_rate_exceed_pkt int `json:"port_kbit_rate_exceed_pkt"`
    Udp_retry_init int `json:"udp_retry_init"`
    Conn_prate_excd int `json:"conn_prate_excd"`
    Ntp_monlist_req int `json:"ntp_monlist_req"`
    Ntp_monlist_resp int `json:"ntp_monlist_resp"`
    Wellknown_sport_drop int `json:"wellknown_sport_drop"`
    Payload_too_small int `json:"payload_too_small"`
    Payload_too_big int `json:"payload_too_big"`
    Fqdn_rate_by_label_count_check int `json:"fqdn_rate_by_label_count_check"`
    Fqdn_rate_by_label_count_exceed int `json:"fqdn_rate_by_label_count_exceed"`
    Udp_auth_retry_gap_drop int `json:"udp_auth_retry_gap_drop"`
    Fqdn_label_count_exceed int `json:"fqdn_label_count_exceed"`
    Rrtype_drop int `json:"rrtype_drop"`
    Src_dns_fqdn_label_len_exceed int `json:"src_dns_fqdn_label_len_exceed"`
    Src_fqdn_label_count_exceed int `json:"src_fqdn_label_count_exceed"`
    Src_udp_auth_fail int `json:"src_udp_auth_fail"`
    Src_force_tcp_auth int `json:"src_force_tcp_auth"`
    Src_dns_malform_drop int `json:"src_dns_malform_drop"`
    Src_udp_auth int `json:"src_udp_auth"`
    Udp_auth_pass int `json:"udp_auth_pass"`
    Force_tcp_auth_timeout int `json:"force_tcp_auth_timeout"`
    Udp_auth_drop int `json:"udp_auth_drop"`
    Dns_auth_drop int `json:"dns_auth_drop"`
    Dns_auth_resp int `json:"dns_auth_resp"`
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
    Sess_aged int `json:"sess_aged"`
    Udp_retry_pass int `json:"udp_retry_pass"`
    Udp_retry_gap_drop int `json:"udp_retry_gap_drop"`
    Filter1_match int `json:"filter1_match"`
    Filter2_match int `json:"filter2_match"`
    Filter3_match int `json:"filter3_match"`
    Filter4_match int `json:"filter4_match"`
    Filter5_match int `json:"filter5_match"`
    Filter_none_match int `json:"filter_none_match"`
    Src_payload_too_small int `json:"src_payload_too_small"`
    Src_payload_too_big int `json:"src_payload_too_big"`
    Src_ntp_monlist_req int `json:"src_ntp_monlist_req"`
    Src_ntp_monlist_resp int `json:"src_ntp_monlist_resp"`
    Src_well_known_port int `json:"src_well_known_port"`
    Src_conn_pkt_rate_excd int `json:"src_conn_pkt_rate_excd"`
    Src_udp_retry_init int `json:"src_udp_retry_init"`
    Src_filter_action_blacklist int `json:"src_filter_action_blacklist"`
    Src_filter_action_drop int `json:"src_filter_action_drop"`
    Src_filter_action_default_pass int `json:"src_filter_action_default_pass"`
    Src_filter_action_whitelist int `json:"src_filter_action_whitelist"`
    Nxdomain_rate_exceed int `json:"nxdomain_rate_exceed"`
    Dns_malform int `json:"dns_malform"`
    Src_dns_malform int `json:"src_dns_malform"`
    Dg_request_check_fail int `json:"dg_request_check_fail"`
    Query_type_any_drop int `json:"query_type_any_drop"`
    Src_query_type_any_drop int `json:"src_query_type_any_drop"`
    Src_udp_auth_drop int `json:"src_udp_auth_drop"`
    Src_dns_auth_drop int `json:"src_dns_auth_drop"`
    Src_frag_drop int `json:"src_frag_drop"`
    Frag_timeout int `json:"frag_timeout"`
    Dg_rate_exceed int `json:"dg_rate_exceed"`
    Pattern_recognition_proceeded int `json:"pattern_recognition_proceeded"`
    Pattern_not_found int `json:"pattern_not_found"`
    Pattern_recognition_generic_error int `json:"pattern_recognition_generic_error"`
    Pattern_filter1_match int `json:"pattern_filter1_match"`
    Pattern_filter2_match int `json:"pattern_filter2_match"`
    Pattern_filter3_match int `json:"pattern_filter3_match"`
    Pattern_filter4_match int `json:"pattern_filter4_match"`
    Pattern_filter5_match int `json:"pattern_filter5_match"`
    Pattern_filter_drop int `json:"pattern_filter_drop"`
    Non_query_opcode_pass_through int `json:"non_query_opcode_pass_through"`
    Src_udp_retry_gap_drop int `json:"src_udp_retry_gap_drop"`
    Src_filter1_match int `json:"src_filter1_match"`
    Src_filter2_match int `json:"src_filter2_match"`
    Src_filter3_match int `json:"src_filter3_match"`
    Src_filter4_match int `json:"src_filter4_match"`
    Src_filter5_match int `json:"src_filter5_match"`
    Src_filter_none_match int `json:"src_filter_none_match"`
    Src_filter_total_not_match int `json:"src_filter_total_not_match"`
    Src_filter_auth_fail int `json:"src_filter_auth_fail"`
    Filter_total_not_match int `json:"filter_total_not_match"`
    Sflow_internal_samples_packed int `json:"sflow_internal_samples_packed"`
    Sflow_external_samples_packed int `json:"sflow_external_samples_packed"`
    Sflow_internal_packets_sent int `json:"sflow_internal_packets_sent"`
    Sflow_external_packets_sent int `json:"sflow_external_packets_sent"`
    Pattern_recognition_sampling_started int `json:"pattern_recognition_sampling_started"`
    Pattern_recognition_pattern_changed int `json:"pattern_recognition_pattern_changed"`
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Src_udp_auth_timeout int `json:"src_udp_auth_timeout"`
    Src_udp_retry_pass int `json:"src_udp_retry_pass"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Token_authentication_mismatched int `json:"token_authentication_mismatched"`
    Token_authentication_invalid int `json:"token_authentication_invalid"`
    Token_authentication_curr_salt_matched int `json:"token_authentication_curr_salt_matched"`
    Token_authentication_prev_salt_matched int `json:"token_authentication_prev_salt_matched"`
    Token_authentication_session_created int `json:"token_authentication_session_created"`
    Token_authentication_session_created_fail int `json:"token_authentication_session_created_fail"`
    Snat_fail int `json:"snat_fail"`
    Exceed_action_drop int `json:"exceed_action_drop"`
}

func (p *DdosDnsUdpPortStats) GetId() string{
    return "1"
}

func (p *DdosDnsUdpPortStats) getPath() string{
    return "ddos/dns-udp-port/stats"
}

func (p *DdosDnsUdpPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDnsUdpPortStats,error) {
logger.Println("DdosDnsUdpPortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDnsUdpPortStats
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
