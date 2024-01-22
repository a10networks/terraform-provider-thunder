

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL7SipStats struct {
    
    Stats DdosL7SipStatsStats `json:"stats"`

}
type DataDdosL7SipStats struct {
    DtDdosL7SipStats DdosL7SipStats `json:"l7-sip"`
}


type DdosL7SipStatsStats struct {
    Policy_drop int `json:"policy_drop"`
    Policy_violation int `json:"policy_violation"`
    Idle_timeout int `json:"idle_timeout"`
    Ofo_timeout int `json:"ofo_timeout"`
    Seq_check_ofo int `json:"seq_check_ofo"`
    Pkts_ofo_total int `json:"pkts_ofo_total"`
    Ofo_queue_size_exceed int `json:"ofo_queue_size_exceed"`
    Seq_check_retrans_fin int `json:"seq_check_retrans_fin"`
    Seq_check_retrans_rst int `json:"seq_check_retrans_rst"`
    Seq_check_retrans_push int `json:"seq_check_retrans_push"`
    Seq_check_retrans_other int `json:"seq_check_retrans_other"`
    Pkts_retrans_total int `json:"pkts_retrans_total"`
    Client_rst int `json:"client_rst"`
    Error_condition int `json:"error_condition"`
    Request_method_ack int `json:"request_method_ack"`
    Request_method_bye int `json:"request_method_bye"`
    Request_method_cancel int `json:"request_method_cancel"`
    Request_method_invite int `json:"request_method_invite"`
    Request_method_info int `json:"request_method_info"`
    Request_method_message int `json:"request_method_message"`
    Request_method_notify int `json:"request_method_notify"`
    Request_method_options int `json:"request_method_options"`
    Request_method_prack int `json:"request_method_prack"`
    Request_method_publish int `json:"request_method_publish"`
    Request_method_register int `json:"request_method_register"`
    Request_method_refer int `json:"request_method_refer"`
    Request_method_subscribe int `json:"request_method_subscribe"`
    Request_method_update int `json:"request_method_update"`
    Request_method_unknown int `json:"request_method_unknown"`
    Request_unknown_version int `json:"request_unknown_version"`
    Keep_alive_msg int `json:"keep_alive_msg"`
    Rate1_limit_exceed int `json:"rate1_limit_exceed"`
    Rate2_limit_exceed int `json:"rate2_limit_exceed"`
    Src_rate1_limit_exceed int `json:"src_rate1_limit_exceed"`
    Src_rate2_limit_exceed int `json:"src_rate2_limit_exceed"`
    Response_1xx int `json:"response_1xx"`
    Response_2xx int `json:"response_2xx"`
    Response_3xx int `json:"response_3xx"`
    Response_4xx int `json:"response_4xx"`
    Response_5xx int `json:"response_5xx"`
    Response_6xx int `json:"response_6xx"`
    Response_unknown int `json:"response_unknown"`
    Response_unknown_version int `json:"response_unknown_version"`
    Read_start_line_error int `json:"read_start_line_error"`
    Invalid_start_line_error int `json:"invalid_start_line_error"`
    Parse_start_line_error int `json:"parse_start_line_error"`
    Line_too_long int `json:"line_too_long"`
    Line_mem_allocated int `json:"line_mem_allocated"`
    Line_mem_freed int `json:"line_mem_freed"`
    Max_uri_len_exceed int `json:"max_uri_len_exceed"`
    Too_many_header int `json:"too_many_header"`
    Invalid_header int `json:"invalid_header"`
    Header_name_too_long int `json:"header_name_too_long"`
    Parse_header_fail_error int `json:"parse_header_fail_error"`
    Max_header_value_len_exceed int `json:"max_header_value_len_exceed"`
    Max_call_id_len_exceed int `json:"max_call_id_len_exceed"`
    Header_filter_match int `json:"header_filter_match"`
    Header_filter_not_match int `json:"header_filter_not_match"`
    Header_filter_none_match int `json:"header_filter_none_match"`
    Header_filter_action_drop int `json:"header_filter_action_drop"`
    Header_filter_action_blacklist int `json:"header_filter_action_blacklist"`
    Header_filter_action_whitelist int `json:"header_filter_action_whitelist"`
    Header_filter_action_default_pass int `json:"header_filter_action_default_pass"`
    Max_sdp_len_exceed int `json:"max_sdp_len_exceed"`
    Body_too_big int `json:"body_too_big"`
    Get_content_fail_error int `json:"get_content_fail_error"`
    Concatenate_msg int `json:"concatenate_msg"`
    Mem_alloc_fail_error int `json:"mem_alloc_fail_error"`
    Malform_request int `json:"malform_request"`
    Src_header_filter_match int `json:"src_header_filter_match"`
    Src_header_filter_not_match int `json:"src_header_filter_not_match"`
    Src_header_filter_action_drop int `json:"src_header_filter_action_drop"`
    Src_header_filter_action_blacklist int `json:"src_header_filter_action_blacklist"`
    Src_header_filter_action_whitelist int `json:"src_header_filter_action_whitelist"`
    Src_header_filter_action_default_pass int `json:"src_header_filter_action_default_pass"`
    Src_dst_header_filter_match int `json:"src_dst_header_filter_match"`
    Src_dst_header_filter_not_match int `json:"src_dst_header_filter_not_match"`
    Src_dst_header_filter_action_drop int `json:"src_dst_header_filter_action_drop"`
    Src_dst_header_filter_action_blacklist int `json:"src_dst_header_filter_action_blacklist"`
    Src_dst_header_filter_action_whitelist int `json:"src_dst_header_filter_action_whitelist"`
    Src_dst_header_filter_action_default_pass int `json:"src_dst_header_filter_action_default_pass"`
}

func (p *DdosL7SipStats) GetId() string{
    return "1"
}

func (p *DdosL7SipStats) getPath() string{
    return "ddos/l7-sip/stats"
}

func (p *DdosL7SipStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL7SipStats,error) {
logger.Println("DdosL7SipStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL7SipStats
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
