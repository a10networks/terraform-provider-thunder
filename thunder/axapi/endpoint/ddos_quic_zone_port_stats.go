

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosQuicZonePortStats struct {
    
    Stats DdosQuicZonePortStatsStats `json:"stats"`

}
type DataDdosQuicZonePortStats struct {
    DtDdosQuicZonePortStats DdosQuicZonePortStats `json:"quic-zone-port"`
}


type DdosQuicZonePortStatsStats struct {
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
    Filter_auth_fail int `json:"filter_auth_fail"`
    Spoof_detect_fail int `json:"spoof_detect_fail"`
    Port_bytes int `json:"port_bytes"`
    Outbound_port_bytes int `json:"outbound_port_bytes"`
    Outbound_port_rcvd int `json:"outbound_port_rcvd"`
    Outbound_port_pkt_sent int `json:"outbound_port_pkt_sent"`
    Port_bytes_sent int `json:"port_bytes_sent"`
    Port_bytes_drop int `json:"port_bytes_drop"`
    Port_src_bl int `json:"port_src_bl"`
    Port_src_escalation int `json:"port_src_escalation"`
    Current_es_level int `json:"current_es_level"`
    Sess_create int `json:"sess_create"`
    Payload_too_small int `json:"payload_too_small"`
    Payload_too_big int `json:"payload_too_big"`
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
    Exceed_drop_brate_src_pkt int `json:"exceed_drop_brate_src_pkt"`
    Port_kbit_rate_exceed_pkt int `json:"port_kbit_rate_exceed_pkt"`
    Udp_retry_init int `json:"udp_retry_init"`
    Conn_prate_excd int `json:"conn_prate_excd"`
    Ntp_monlist_req int `json:"ntp_monlist_req"`
    Ntp_monlist_resp int `json:"ntp_monlist_resp"`
    Wellknown_sport_drop int `json:"wellknown_sport_drop"`
    Udp_auth_drop int `json:"udp_auth_drop"`
    Bl int `json:"bl"`
    Src_drop int `json:"src_drop"`
    Frag_rcvd int `json:"frag_rcvd"`
    Frag_drop int `json:"frag_drop"`
    Sess_create_inbound int `json:"sess_create_inbound"`
    Sess_create_outbound int `json:"sess_create_outbound"`
    Sess_aged int `json:"sess_aged"`
    Udp_retry_pass int `json:"udp_retry_pass"`
    Udp_retry_gap_drop int `json:"udp_retry_gap_drop"`
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
    Secondary_port_pkt_rate_exceed int `json:"secondary_port_pkt_rate_exceed"`
    Secondary_port_kbit_rate_exceed int `json:"secondary_port_kbit_rate_exceed"`
    Secondary_port_kbit_rate_exceed_pkt int `json:"secondary_port_kbit_rate_exceed_pkt"`
    Secondary_port_conn_rate_exceed int `json:"secondary_port_conn_rate_exceed"`
    Secondary_port_conn_limm_exceed int `json:"secondary_port_conn_limm_exceed"`
    Src_udp_retry_gap_drop int `json:"src_udp_retry_gap_drop"`
    Src_udp_auth_drop int `json:"src_udp_auth_drop"`
    Src_frag_drop int `json:"src_frag_drop"`
    No_policy_class_list_match int `json:"no_policy_class_list_match"`
    Frag_timeout int `json:"frag_timeout"`
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
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Dst_udp_retry_timeout_blacklist int `json:"dst_udp_retry_timeout_blacklist"`
    Src_udp_auth_timeout int `json:"src_udp_auth_timeout"`
    Zone_src_udp_retry_timeout_blacklist int `json:"zone_src_udp_retry_timeout_blacklist"`
    Src_udp_retry_pass int `json:"src_udp_retry_pass"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Secondary_port_hit int `json:"secondary_port_hit"`
    Processed int `json:"processed"`
    Initial int `json:"initial"`
    Version_negotiation int `json:"version_negotiation"`
    Rtt0 int `json:"rtt0"`
    Handshake int `json:"handshake"`
    Retry int `json:"retry"`
    Invalid_packet_type int `json:"invalid_packet_type"`
    Packet_parse_failed int `json:"packet_parse_failed"`
    Version_match_action_taken int `json:"version_match_action_taken"`
    Version_match_action_drop int `json:"version_match_action_drop"`
    Version_match_action_blacklist int `json:"version_match_action_blacklist"`
    Malformed_action_taken int `json:"malformed_action_taken"`
    Malformed_action_drop int `json:"malformed_action_drop"`
    Malformed_action_blacklist int `json:"malformed_action_blacklist"`
    Malformed_dcid_len_max_exceed int `json:"malformed_dcid_len_max_exceed"`
    Malformed_scid_len_max_exceed int `json:"malformed_scid_len_max_exceed"`
    Initial_force_tcp int `json:"initial_force_tcp"`
    Fixed_bit_not_set int `json:"fixed_bit_not_set"`
    Src_zone_service_entry_learned int `json:"src_zone_service_entry_learned"`
    Src_zone_service_entry_aged int `json:"src_zone_service_entry_aged"`
    Dst_hw_drop_inserted int `json:"dst_hw_drop_inserted"`
    Dst_hw_drop_removed int `json:"dst_hw_drop_removed"`
    Src_hw_drop_inserted int `json:"src_hw_drop_inserted"`
    Src_hw_drop_removed int `json:"src_hw_drop_removed"`
    Token_authentication_mismatched int `json:"token_authentication_mismatched"`
    Token_authentication_invalid int `json:"token_authentication_invalid"`
    Token_authentication_curr_salt_matched int `json:"token_authentication_curr_salt_matched"`
    Token_authentication_prev_salt_matched int `json:"token_authentication_prev_salt_matched"`
    Token_authentication_session_created int `json:"token_authentication_session_created"`
    Token_authentication_session_created_fail int `json:"token_authentication_session_created_fail"`
    Snat_fail int `json:"snat_fail"`
    Exceed_action_drop int `json:"exceed_action_drop"`
    Ew_inbound_port_rcv int `json:"ew_inbound_port_rcv"`
    Ew_inbound_port_drop int `json:"ew_inbound_port_drop"`
    Ew_inbound_port_sent int `json:"ew_inbound_port_sent"`
    Ew_inbound_port_byte_rcv int `json:"ew_inbound_port_byte_rcv"`
    Ew_inbound_port_byte_drop int `json:"ew_inbound_port_byte_drop"`
    Ew_inbound_port_byte_sent int `json:"ew_inbound_port_byte_sent"`
    Ew_outbound_port_rcv int `json:"ew_outbound_port_rcv"`
    Ew_outbound_port_drop int `json:"ew_outbound_port_drop"`
    Ew_outbound_port_sent int `json:"ew_outbound_port_sent"`
    Ew_outbound_port_byte_rcv int `json:"ew_outbound_port_byte_rcv"`
    Ew_outbound_port_byte_sent int `json:"ew_outbound_port_byte_sent"`
    Ew_outbound_port_byte_drop int `json:"ew_outbound_port_byte_drop"`
    No_route_drop int `json:"no_route_drop"`
    Unauth_src_session_reset int `json:"unauth_src_session_reset"`
    Src_hw_drop int `json:"src_hw_drop"`
    Addr_filter_drop int `json:"addr_filter_drop"`
    Addr_filter_bl int `json:"addr_filter_bl"`
    Src_learn_overflow int `json:"src_learn_overflow"`
}

func (p *DdosQuicZonePortStats) GetId() string{
    return "1"
}

func (p *DdosQuicZonePortStats) getPath() string{
    return "ddos/quic-zone-port/stats"
}

func (p *DdosQuicZonePortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosQuicZonePortStats,error) {
logger.Println("DdosQuicZonePortStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosQuicZonePortStats
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
