

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwLoggingStats struct {
    
    Gtp FwLoggingStatsGtp `json:"gtp"`
    Stats FwLoggingStatsStats `json:"stats"`

}
type DataFwLoggingStats struct {
    DtFwLoggingStats FwLoggingStats `json:"logging"`
}


type FwLoggingStatsGtp struct {
    Stats FwLoggingStatsGtpStats `json:"stats"`
}


type FwLoggingStatsGtpStats struct {
    Log_type_gtp_invalid_teid int `json:"log_type_gtp_invalid_teid"`
    Log_gtp_type_reserved_ie_present int `json:"log_gtp_type_reserved_ie_present"`
    Log_type_gtp_mandatory_ie_missing int `json:"log_type_gtp_mandatory_ie_missing"`
    Log_type_gtp_mandatory_ie_inside_grouped_ie_missing int `json:"log_type_gtp_mandatory_ie_inside_grouped_ie_missing"`
    Log_type_gtp_msisdn_filtering int `json:"log_type_gtp_msisdn_filtering"`
    Log_type_gtp_out_of_order_ie int `json:"log_type_gtp_out_of_order_ie"`
    Log_type_gtp_out_of_state_ie int `json:"log_type_gtp_out_of_state_ie"`
    Log_type_enduser_ip_spoofed int `json:"log_type_enduser_ip_spoofed"`
    Log_type_crosslayer_correlation int `json:"log_type_crosslayer_correlation"`
    Log_type_message_not_supported int `json:"log_type_message_not_supported"`
    Log_type_out_of_state int `json:"log_type_out_of_state"`
    Log_type_max_msg_length int `json:"log_type_max_msg_length"`
    Log_type_gtp_message_filtering int `json:"log_type_gtp_message_filtering"`
    Log_type_gtp_apn_filtering int `json:"log_type_gtp_apn_filtering"`
    Log_type_gtp_rat_type_filtering int `json:"log_type_gtp_rat_type_filtering"`
    Log_type_country_code_mismatch int `json:"log_type_country_code_mismatch"`
    Log_type_gtp_in_gtp_filtering int `json:"log_type_gtp_in_gtp_filtering"`
    Log_type_gtp_node_restart int `json:"log_type_gtp_node_restart"`
    Log_type_gtp_seq_num_mismatch int `json:"log_type_gtp_seq_num_mismatch"`
    Log_type_gtp_rate_limit_periodic int `json:"log_type_gtp_rate_limit_periodic"`
    Log_type_gtp_invalid_message_length int `json:"log_type_gtp_invalid_message_length"`
    Log_type_gtp_hdr_invalid_protocol_flag int `json:"log_type_gtp_hdr_invalid_protocol_flag"`
    Log_type_gtp_hdr_invalid_spare_bits int `json:"log_type_gtp_hdr_invalid_spare_bits"`
    Log_type_gtp_hdr_invalid_piggy_flag int `json:"log_type_gtp_hdr_invalid_piggy_flag"`
    Log_type_gtp_invalid_version int `json:"log_type_gtp_invalid_version"`
    Log_type_gtp_invalid_ports int `json:"log_type_gtp_invalid_ports"`
}


type FwLoggingStatsStats struct {
    Log_message_sent int `json:"log_message_sent"`
    Log_type_reset int `json:"log_type_reset"`
    Log_type_deny int `json:"log_type_deny"`
    Log_type_session_closed int `json:"log_type_session_closed"`
    Log_type_session_opened int `json:"log_type_session_opened"`
    Rule_not_logged int `json:"rule_not_logged"`
    LogDropped int `json:"log-dropped"`
    TcpSessionCreated int `json:"tcp-session-created"`
    TcpSessionDeleted int `json:"tcp-session-deleted"`
    UdpSessionCreated int `json:"udp-session-created"`
    UdpSessionDeleted int `json:"udp-session-deleted"`
    IcmpSessionDeleted int `json:"icmp-session-deleted"`
    IcmpSessionCreated int `json:"icmp-session-created"`
    Icmpv6SessionDeleted int `json:"icmpv6-session-deleted"`
    Icmpv6SessionCreated int `json:"icmpv6-session-created"`
    OtherSessionDeleted int `json:"other-session-deleted"`
    OtherSessionCreated int `json:"other-session-created"`
    HttpRequestLogged int `json:"http-request-logged"`
    HttpLoggingInvalidFormat int `json:"http-logging-invalid-format"`
    SctpSessionCreated int `json:"sctp-session-created"`
    SctpSessionDeleted int `json:"sctp-session-deleted"`
    Log_type_sctp_inner_proto_filter int `json:"log_type_sctp_inner_proto_filter"`
    IddosBlackholeEntryCreate int `json:"iddos-blackhole-entry-create"`
    IddosBlackholeEntryDelete int `json:"iddos-blackhole-entry-delete"`
    SessionLimitExceeded int `json:"session-limit-exceeded"`
}

func (p *FwLoggingStats) GetId() string{
    return "1"
}

func (p *FwLoggingStats) getPath() string{
    return "fw/logging/stats"
}

func (p *FwLoggingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwLoggingStats,error) {
logger.Println("FwLoggingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwLoggingStats
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
