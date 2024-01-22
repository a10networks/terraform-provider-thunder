

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwLoggingGtpStats struct {
    
    Stats FwLoggingGtpStatsStats `json:"stats"`

}
type DataFwLoggingGtpStats struct {
    DtFwLoggingGtpStats FwLoggingGtpStats `json:"gtp"`
}


type FwLoggingGtpStatsStats struct {
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

func (p *FwLoggingGtpStats) GetId() string{
    return "1"
}

func (p *FwLoggingGtpStats) getPath() string{
    return "fw/logging/gtp/stats"
}

func (p *FwLoggingGtpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwLoggingGtpStats,error) {
logger.Println("FwLoggingGtpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwLoggingGtpStats
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
