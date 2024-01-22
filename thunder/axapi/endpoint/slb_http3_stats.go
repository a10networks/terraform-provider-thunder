

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHttp3Stats struct {
    
    Stats SlbHttp3StatsStats `json:"stats"`

}
type DataSlbHttp3Stats struct {
    DtSlbHttp3Stats SlbHttp3Stats `json:"http3"`
}


type SlbHttp3StatsStats struct {
    Client_conn_curr int `json:"client_conn_curr"`
    Server_conn_curr int `json:"server_conn_curr"`
    Client_conn_total int `json:"client_conn_total"`
    Server_conn_total int `json:"server_conn_total"`
    Client_conn_peak int `json:"client_conn_peak"`
    Server_conn_peak int `json:"server_conn_peak"`
    Client_request_streams_curr int `json:"client_request_streams_curr"`
    Server_request_streams_curr int `json:"server_request_streams_curr"`
    Client_request_streams_total int `json:"client_request_streams_total"`
    Server_request_streams_total int `json:"server_request_streams_total"`
    Client_request_push_curr int `json:"client_request_push_curr"`
    Server_request_push_curr int `json:"server_request_push_curr"`
    Client_request_push_total int `json:"client_request_push_total"`
    Server_request_push_total int `json:"server_request_push_total"`
    Client_request_other_curr int `json:"client_request_other_curr"`
    Server_request_other_curr int `json:"server_request_other_curr"`
    Client_request_other_total int `json:"client_request_other_total"`
    Server_request_other_total int `json:"server_request_other_total"`
    Client_frame_type_headers_rcvd int `json:"client_frame_type_headers_rcvd"`
    Client_frame_type_headers_sent int `json:"client_frame_type_headers_sent"`
    Client_frame_type_data_rcvd int `json:"client_frame_type_data_rcvd"`
    Client_frame_type_data_sent int `json:"client_frame_type_data_sent"`
    Client_frame_type_cancel_push_rcvd int `json:"client_frame_type_cancel_push_rcvd"`
    Client_frame_type_cancel_push_sent int `json:"client_frame_type_cancel_push_sent"`
    Client_frame_type_settings_rcvd int `json:"client_frame_type_settings_rcvd"`
    Client_frame_type_settings_sent int `json:"client_frame_type_settings_sent"`
    Client_frame_type_push_promise_rcvd int `json:"client_frame_type_push_promise_rcvd"`
    Client_frame_type_push_promise_sent int `json:"client_frame_type_push_promise_sent"`
    Client_frame_type_goaway_rcvd int `json:"client_frame_type_goaway_rcvd"`
    Client_frame_type_goaway_sent int `json:"client_frame_type_goaway_sent"`
    Client_frame_type_max_push_id_rcvd int `json:"client_frame_type_max_push_id_rcvd"`
    Client_frame_type_max_push_id_sent int `json:"client_frame_type_max_push_id_sent"`
    Client_frame_type_unknown_rcvd int `json:"client_frame_type_unknown_rcvd"`
    Client_header_frames_to_app int `json:"client_header_frames_to_app"`
    Client_data_frames_to_app int `json:"client_data_frames_to_app"`
    Client_header_bytes_rcvd int `json:"client_header_bytes_rcvd"`
    Client_header_bytes_sent int `json:"client_header_bytes_sent"`
    Client_data_bytes_rcvd int `json:"client_data_bytes_rcvd"`
    Client_data_bytes_sent int `json:"client_data_bytes_sent"`
    Client_other_frame_bytes_rcvd int `json:"client_other_frame_bytes_rcvd"`
    Client_other_frame_bytes_sent int `json:"client_other_frame_bytes_sent"`
    Client_heading_bytes_rcvd int `json:"client_heading_bytes_rcvd"`
    Client_heading_bytes_sent int `json:"client_heading_bytes_sent"`
    Client_total_bytes_rcvd int `json:"client_total_bytes_rcvd"`
    Client_total_bytes_sent int `json:"client_total_bytes_sent"`
    Server_frame_type_headers_rcvd int `json:"server_frame_type_headers_rcvd"`
    Server_frame_type_headers_sent int `json:"server_frame_type_headers_sent"`
    Server_frame_type_data_rcvd int `json:"server_frame_type_data_rcvd"`
    Server_frame_type_data_sent int `json:"server_frame_type_data_sent"`
    Server_frame_type_cancel_push_rcvd int `json:"server_frame_type_cancel_push_rcvd"`
    Server_frame_type_cancel_push_sent int `json:"server_frame_type_cancel_push_sent"`
    Server_frame_type_settings_rcvd int `json:"server_frame_type_settings_rcvd"`
    Server_frame_type_settings_sent int `json:"server_frame_type_settings_sent"`
    Server_frame_type_push_promise_rcvd int `json:"server_frame_type_push_promise_rcvd"`
    Server_frame_type_push_promise_sent int `json:"server_frame_type_push_promise_sent"`
    Server_frame_type_goaway_rcvd int `json:"server_frame_type_goaway_rcvd"`
    Server_frame_type_goaway_sent int `json:"server_frame_type_goaway_sent"`
    Server_frame_type_max_push_id_rcvd int `json:"server_frame_type_max_push_id_rcvd"`
    Server_frame_type_max_push_id_sent int `json:"server_frame_type_max_push_id_sent"`
    Server_frame_type_unknown_rcvd int `json:"server_frame_type_unknown_rcvd"`
    Server_header_frames_to_app int `json:"server_header_frames_to_app"`
    Server_data_frames_to_app int `json:"server_data_frames_to_app"`
    Server_header_bytes_rcvd int `json:"server_header_bytes_rcvd"`
    Server_header_bytes_sent int `json:"server_header_bytes_sent"`
    Server_data_bytes_rcvd int `json:"server_data_bytes_rcvd"`
    Server_data_bytes_sent int `json:"server_data_bytes_sent"`
    Server_other_frame_bytes_rcvd int `json:"server_other_frame_bytes_rcvd"`
    Server_other_frame_bytes_sent int `json:"server_other_frame_bytes_sent"`
    Server_heading_bytes_rcvd int `json:"server_heading_bytes_rcvd"`
    Server_heading_bytes_sent int `json:"server_heading_bytes_sent"`
    Server_total_bytes_rcvd int `json:"server_total_bytes_rcvd"`
    Server_total_bytes_sent int `json:"server_total_bytes_sent"`
    Invalid_argument int `json:"invalid_argument"`
    Invalid_state int `json:"invalid_state"`
    Wouldblock int `json:"wouldblock"`
    Stream_in_use int `json:"stream_in_use"`
    Push_id_blocked int `json:"push_id_blocked"`
    Malformed_http_header int `json:"malformed_http_header"`
    Remove_http_header int `json:"remove_http_header"`
    Malformed_http_messaging int `json:"malformed_http_messaging"`
    Too_late int `json:"too_late"`
    Qpack_fatal int `json:"qpack_fatal"`
    Qpack_header_too_large int `json:"qpack_header_too_large"`
    Ignore_stream int `json:"ignore_stream"`
    Stream_not_found int `json:"stream_not_found"`
    Ignore_push_promise int `json:"ignore_push_promise"`
    Qpack_decompression_failed int `json:"qpack_decompression_failed"`
    Qpack_encoder_stream_error int `json:"qpack_encoder_stream_error"`
    Qpack_decoder_stream_error int `json:"qpack_decoder_stream_error"`
    H3_frame_unexpected int `json:"h3_frame_unexpected"`
    H3_frame_error int `json:"h3_frame_error"`
    H3_missing_settings int `json:"h3_missing_settings"`
    H3_internal_error int `json:"h3_internal_error"`
    H3_closed_critical_stream int `json:"h3_closed_critical_stream"`
    H3_general_protocol_error int `json:"h3_general_protocol_error"`
    H3_id_error int `json:"h3_id_error"`
    H3_settings_error int `json:"h3_settings_error"`
    H3_stream_creation_error int `json:"h3_stream_creation_error"`
    Fatal int `json:"fatal"`
    Conn_alloc_error int `json:"conn_alloc_error"`
    Alloc_fail_total int `json:"alloc_fail_total"`
    Http3_rejected int `json:"http3_rejected"`
}

func (p *SlbHttp3Stats) GetId() string{
    return "1"
}

func (p *SlbHttp3Stats) getPath() string{
    return "slb/http3/stats"
}

func (p *SlbHttp3Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHttp3Stats,error) {
logger.Println("SlbHttp3Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHttp3Stats
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
