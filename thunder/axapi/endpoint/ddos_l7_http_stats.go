

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosL7HttpStats struct {
    
    Stats DdosL7HttpStatsStats `json:"stats"`

}
type DataDdosL7HttpStats struct {
    DtDdosL7HttpStats DdosL7HttpStats `json:"l7-http"`
}


type DdosL7HttpStatsStats struct {
    Req_processed int `json:"req_processed"`
    Req_ofo int `json:"req_ofo"`
    Ofo_timer_expired int `json:"ofo_timer_expired"`
    Ofo_queue_exceed int `json:"ofo_queue_exceed"`
    Ofo int `json:"ofo"`
    Partial_hdr int `json:"partial_hdr"`
    Http_idle_timeout int `json:"http_idle_timeout"`
    New_syn int `json:"new_syn"`
    Retrans int `json:"retrans"`
    Retrans_fin int `json:"retrans_fin"`
    Retrans_push int `json:"retrans_push"`
    Retrans_rst int `json:"retrans_rst"`
    Req_retrans int `json:"req_retrans"`
    Request int `json:"request"`
    Req_content_len int `json:"req_content_len"`
    Src_req_rate_exceed int `json:"src_req_rate_exceed"`
    Dst_req_rate_exceed int `json:"dst_req_rate_exceed"`
    Lower_than_mss_exceed int `json:"lower_than_mss_exceed"`
    Parsereq_fail int `json:"parsereq_fail"`
    Neg_req_remain int `json:"neg_req_remain"`
    Neg_rsp_remain int `json:"neg_rsp_remain"`
    Invalid_header int `json:"invalid_header"`
    Too_many_headers int `json:"too_many_headers"`
    Header_name_too_long int `json:"header_name_too_long"`
    Invalid_hdr_name int `json:"invalid_hdr_name"`
    Invalid_hdr_val int `json:"invalid_hdr_val"`
    Line_too_long int `json:"line_too_long"`
    Client_rst int `json:"client_rst"`
    Hps_server_rst int `json:"hps_server_rst"`
    Ddos_policy_violation int `json:"ddos_policy_violation"`
    Policy_drop int `json:"policy_drop"`
    Error_condition int `json:"error_condition"`
    Http11 int `json:"http11"`
    Http10 int `json:"http10"`
    Rsp_chunk int `json:"rsp_chunk"`
    Http_get int `json:"http_get"`
    Http_head int `json:"http_head"`
    Http_put int `json:"http_put"`
    Http_post int `json:"http_post"`
    Http_trace int `json:"http_trace"`
    Http_options int `json:"http_options"`
    Http_connect int `json:"http_connect"`
    Http_del int `json:"http_del"`
    Http_unknown int `json:"http_unknown"`
    Hps_req_sz_1k int `json:"hps_req_sz_1k"`
    Hps_req_sz_2k int `json:"hps_req_sz_2k"`
    Hps_req_sz_4k int `json:"hps_req_sz_4k"`
    Hps_req_sz_8k int `json:"hps_req_sz_8k"`
    Hps_req_sz_16k int `json:"hps_req_sz_16k"`
    Hps_req_sz_32k int `json:"hps_req_sz_32k"`
    Hps_req_sz_64k int `json:"hps_req_sz_64k"`
    Hps_req_sz_256k int `json:"hps_req_sz_256k"`
    Hps_req_sz_256k_plus int `json:"hps_req_sz_256k_plus"`
    Hps_rsp_11 int `json:"hps_rsp_11"`
    Hps_rsp_10 int `json:"hps_rsp_10"`
    Hps_rsp_sz_1k int `json:"hps_rsp_sz_1k"`
    Hps_rsp_sz_2k int `json:"hps_rsp_sz_2k"`
    Hps_rsp_sz_4k int `json:"hps_rsp_sz_4k"`
    Hps_rsp_sz_8k int `json:"hps_rsp_sz_8k"`
    Hps_rsp_sz_16k int `json:"hps_rsp_sz_16k"`
    Hps_rsp_sz_32k int `json:"hps_rsp_sz_32k"`
    Hps_rsp_sz_64k int `json:"hps_rsp_sz_64k"`
    Hps_rsp_sz_256k int `json:"hps_rsp_sz_256k"`
    Hps_rsp_sz_256k_plus int `json:"hps_rsp_sz_256k_plus"`
    Hps_rsp_status_1xx int `json:"hps_rsp_status_1xx"`
    Hps_rsp_status_2xx int `json:"hps_rsp_status_2xx"`
    Hps_rsp_status_3xx int `json:"hps_rsp_status_3xx"`
    Hps_rsp_status_4xx int `json:"hps_rsp_status_4xx"`
    Hps_rsp_status_5xx int `json:"hps_rsp_status_5xx"`
    Hps_rsp_status_504_ax int `json:"hps_rsp_status_504_AX"`
    Hps_rsp_status_6xx int `json:"hps_rsp_status_6xx"`
    Hps_rsp_status_unknown int `json:"hps_rsp_status_unknown"`
    Chunk_sz_512 int `json:"chunk_sz_512"`
    Chunk_sz_1k int `json:"chunk_sz_1k"`
    Chunk_sz_2k int `json:"chunk_sz_2k"`
    Chunk_sz_4k int `json:"chunk_sz_4k"`
    Chunk_sz_gt_4k int `json:"chunk_sz_gt_4k"`
    Chunk_bad int `json:"chunk_bad"`
    Challenge_fail int `json:"challenge_fail"`
    Challenge_ud_sent int `json:"challenge_ud_sent"`
    Challenge_ud_fail int `json:"challenge_ud_fail"`
    Challenge_js_sent int `json:"challenge_js_sent"`
    Challenge_js_fail int `json:"challenge_js_fail"`
    Malform_bad_chunk int `json:"malform_bad_chunk"`
    Malform_content_len_too_long int `json:"malform_content_len_too_long"`
    Malform_header_name_too_long int `json:"malform_header_name_too_long"`
    Malform_line_too_long int `json:"malform_line_too_long"`
    Malform_req_line_too_long int `json:"malform_req_line_too_long"`
    Malform_too_many_headers int `json:"malform_too_many_headers"`
    Window_small int `json:"window_small"`
    Window_small_drop int `json:"window_small_drop"`
    Alloc_fail int `json:"alloc_fail"`
    Use_hdr_ip_as_source int `json:"use_hdr_ip_as_source"`
    Agent_filter_match int `json:"agent_filter_match"`
    Agent_filter_blacklist int `json:"agent_filter_blacklist"`
    Referer_filter_match int `json:"referer_filter_match"`
    Referer_filter_blacklist int `json:"referer_filter_blacklist"`
    Dst_filter_match int `json:"dst_filter_match"`
    Dst_filter_not_match int `json:"dst_filter_not_match"`
    Dst_filter_action_blacklist int `json:"dst_filter_action_blacklist"`
    Dst_filter_action_drop int `json:"dst_filter_action_drop"`
    Dst_filter_action_default_pass int `json:"dst_filter_action_default_pass"`
    Dst_post_rate_exceed int `json:"dst_post_rate_exceed"`
    Src_post_rate_exceed int `json:"src_post_rate_exceed"`
    Dst_resp_rate_exceed int `json:"dst_resp_rate_exceed"`
    Dst_filter_action_whitelist int `json:"dst_filter_action_whitelist"`
    Src_filter_match int `json:"src_filter_match"`
    Src_filter_not_match int `json:"src_filter_not_match"`
    Src_filter_action_blacklist int `json:"src_filter_action_blacklist"`
    Src_filter_action_drop int `json:"src_filter_action_drop"`
    Src_filter_action_default_pass int `json:"src_filter_action_default_pass"`
    Src_filter_action_whitelist int `json:"src_filter_action_whitelist"`
    Src_dst_filter_match int `json:"src_dst_filter_match"`
    Src_dst_filter_not_match int `json:"src_dst_filter_not_match"`
    Src_dst_filter_action_blacklist int `json:"src_dst_filter_action_blacklist"`
    Src_dst_filter_action_drop int `json:"src_dst_filter_action_drop"`
    Src_dst_filter_action_default_pass int `json:"src_dst_filter_action_default_pass"`
    Src_dst_filter_action_whitelist int `json:"src_dst_filter_action_whitelist"`
    Dst_filter_rate_exceed int `json:"dst_filter_rate_exceed"`
    Dst_filter_action_ignore int `json:"dst_filter_action_ignore"`
    Dst_filter_action_reset int `json:"dst_filter_action_reset"`
    Uri_filter_match int `json:"uri_filter_match"`
    Http_auth_drop int `json:"http_auth_drop"`
    Http_auth_resp int `json:"http_auth_resp"`
    Header_processing_time_0 int `json:"header_processing_time_0"`
    Header_processing_time_1 int `json:"header_processing_time_1"`
    Header_processing_time_2 int `json:"header_processing_time_2"`
    Header_processing_time_3 int `json:"header_processing_time_3"`
    Header_processing_incomplete int `json:"header_processing_incomplete"`
    Malform_req_line_too_small int `json:"malform_req_line_too_small"`
    Malform_req_line_invalid_method int `json:"malform_req_line_invalid_method"`
}

func (p *DdosL7HttpStats) GetId() string{
    return "1"
}

func (p *DdosL7HttpStats) getPath() string{
    return "ddos/l7-http/stats"
}

func (p *DdosL7HttpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosL7HttpStats,error) {
logger.Println("DdosL7HttpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosL7HttpStats
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
