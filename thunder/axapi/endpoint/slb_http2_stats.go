

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHttp2Stats struct {
    
    Stats SlbHttp2StatsStats `json:"stats"`

}
type DataSlbHttp2Stats struct {
    DtSlbHttp2Stats SlbHttp2Stats `json:"http2"`
}


type SlbHttp2StatsStats struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Connection_preface_rcvd int `json:"connection_preface_rcvd"`
    Control_frame int `json:"control_frame"`
    Headers_frame int `json:"headers_frame"`
    Continuation_frame int `json:"continuation_frame"`
    Rst_frame_rcvd int `json:"rst_frame_rcvd"`
    Settings_frame int `json:"settings_frame"`
    Window_update_frame int `json:"window_update_frame"`
    Ping_frame int `json:"ping_frame"`
    Goaway_frame int `json:"goaway_frame"`
    Priority_frame int `json:"priority_frame"`
    Data_frame int `json:"data_frame"`
    Unknown_frame int `json:"unknown_frame"`
    Connection_preface_sent int `json:"connection_preface_sent"`
    Settings_frame_sent int `json:"settings_frame_sent"`
    Settings_ack_sent int `json:"settings_ack_sent"`
    Empty_settings_sent int `json:"empty_settings_sent"`
    Ping_frame_sent int `json:"ping_frame_sent"`
    Window_update_frame_sent int `json:"window_update_frame_sent"`
    Rst_frame_sent int `json:"rst_frame_sent"`
    Goaway_frame_sent int `json:"goaway_frame_sent"`
    Header_to_app int `json:"header_to_app"`
    Data_to_app int `json:"data_to_app"`
    Protocol_error int `json:"protocol_error"`
    Internal_error int `json:"internal_error"`
    Proxy_alloc_error int `json:"proxy_alloc_error"`
    Split_buff_fail int `json:"split_buff_fail"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Error_max_invalid_stream int `json:"error_max_invalid_stream"`
    Data_no_stream int `json:"data_no_stream"`
    Flow_control_error int `json:"flow_control_error"`
    Settings_timeout int `json:"settings_timeout"`
    Frame_size_error int `json:"frame_size_error"`
    Refused_stream int `json:"refused_stream"`
    Cancel int `json:"cancel"`
    Compression_error int `json:"compression_error"`
    Connect_error int `json:"connect_error"`
    Enhance_your_calm int `json:"enhance_your_calm"`
    Inadequate_security int `json:"inadequate_security"`
    Http_1_1_required int `json:"http_1_1_required"`
    Deflate_alloc_fail int `json:"deflate_alloc_fail"`
    Inflate_alloc_fail int `json:"inflate_alloc_fail"`
    Inflate_header_fail int `json:"inflate_header_fail"`
    Bad_connection_preface int `json:"bad_connection_preface"`
    Cant_allocate_control_frame int `json:"cant_allocate_control_frame"`
    Cant_allocate_settings_frame int `json:"cant_allocate_settings_frame"`
    Bad_frame_type_for_stream_state int `json:"bad_frame_type_for_stream_state"`
    Wrong_stream_state int `json:"wrong_stream_state"`
    Data_queue_alloc_error int `json:"data_queue_alloc_error"`
    Buff_alloc_error int `json:"buff_alloc_error"`
    Cant_allocate_rst_frame int `json:"cant_allocate_rst_frame"`
    Cant_allocate_goaway_frame int `json:"cant_allocate_goaway_frame"`
    Cant_allocate_ping_frame int `json:"cant_allocate_ping_frame"`
    Cant_allocate_stream int `json:"cant_allocate_stream"`
    Cant_allocate_window_frame int `json:"cant_allocate_window_frame"`
    Header_no_stream int `json:"header_no_stream"`
    Header_padlen_gt_frame_payload int `json:"header_padlen_gt_frame_payload"`
    Streams_gt_max_concur_streams int `json:"streams_gt_max_concur_streams"`
    Idle_state_unexpected_frame int `json:"idle_state_unexpected_frame"`
    Reserved_local_state_unexpected_frame int `json:"reserved_local_state_unexpected_frame"`
    Reserved_remote_state_unexpected_frame int `json:"reserved_remote_state_unexpected_frame"`
    Half_closed_remote_state_unexpected_frame int `json:"half_closed_remote_state_unexpected_frame"`
    Closed_state_unexpected_frame int `json:"closed_state_unexpected_frame"`
    Zero_window_size_on_stream int `json:"zero_window_size_on_stream"`
    Exceeds_max_window_size_stream int `json:"exceeds_max_window_size_stream"`
    Stream_closed int `json:"stream_closed"`
    Continuation_before_headers int `json:"continuation_before_headers"`
    Invalid_frame_during_headers int `json:"invalid_frame_during_headers"`
    Headers_after_continuation int `json:"headers_after_continuation"`
    Push_promise_frame_sent int `json:"push_promise_frame_sent"`
    Invalid_push_promise int `json:"invalid_push_promise"`
    Invalid_stream_id int `json:"invalid_stream_id"`
    Headers_interleaved int `json:"headers_interleaved"`
    Trailers_no_end_stream int `json:"trailers_no_end_stream"`
    Invalid_setting_value int `json:"invalid_setting_value"`
    Invalid_window_update int `json:"invalid_window_update"`
    Frame_header_bytes_received int `json:"frame_header_bytes_received"`
    Frame_header_bytes_sent int `json:"frame_header_bytes_sent"`
    Control_bytes_received int `json:"control_bytes_received"`
    Control_bytes_sent int `json:"control_bytes_sent"`
    Header_bytes_received int `json:"header_bytes_received"`
    Header_bytes_sent int `json:"header_bytes_sent"`
    Data_bytes_received int `json:"data_bytes_received"`
    Data_bytes_sent int `json:"data_bytes_sent"`
    Total_bytes_received int `json:"total_bytes_received"`
    Total_bytes_sent int `json:"total_bytes_sent"`
    Peak_proxy int `json:"peak_proxy"`
    Control_frame_sent int `json:"control_frame_sent"`
    Continuation_frame_sent int `json:"continuation_frame_sent"`
    Data_frame_sent int `json:"data_frame_sent"`
    Headers_frame_sent int `json:"headers_frame_sent"`
    Priority_frame_sent int `json:"priority_frame_sent"`
    Settings_ack_rcvd int `json:"settings_ack_rcvd"`
    Empty_settings_rcvd int `json:"empty_settings_rcvd"`
    Alloc_fail_total int `json:"alloc_fail_total"`
    Err_rcvd_total int `json:"err_rcvd_total"`
    Err_sent_total int `json:"err_sent_total"`
    Err_sent_proto_err int `json:"err_sent_proto_err"`
    Err_sent_internal_err int `json:"err_sent_internal_err"`
    Err_sent_flow_control int `json:"err_sent_flow_control"`
    Err_sent_setting_timeout int `json:"err_sent_setting_timeout"`
    Err_sent_stream_closed int `json:"err_sent_stream_closed"`
    Err_sent_frame_size_err int `json:"err_sent_frame_size_err"`
    Err_sent_refused_stream int `json:"err_sent_refused_stream"`
    Err_sent_cancel int `json:"err_sent_cancel"`
    Err_sent_compression_err int `json:"err_sent_compression_err"`
    Err_sent_connect_err int `json:"err_sent_connect_err"`
    Err_sent_your_calm int `json:"err_sent_your_calm"`
    Err_sent_inadequate_security int `json:"err_sent_inadequate_security"`
    Err_sent_http11_required int `json:"err_sent_http11_required"`
    Http2_rejected int `json:"http2_rejected"`
    Current_stream int `json:"current_stream"`
    Stream_create int `json:"stream_create"`
    Stream_free int `json:"stream_free"`
    End_stream_rcvd int `json:"end_stream_rcvd"`
    End_stream_sent int `json:"end_stream_sent"`
}

func (p *SlbHttp2Stats) GetId() string{
    return "1"
}

func (p *SlbHttp2Stats) getPath() string{
    return "slb/http2/stats"
}

func (p *SlbHttp2Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHttp2Stats,error) {
logger.Println("SlbHttp2Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHttp2Stats
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
