

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSpdyProxyStats struct {
    
    Stats SlbSpdyProxyStatsStats `json:"stats"`

}
type DataSlbSpdyProxyStats struct {
    DtSlbSpdyProxyStats SlbSpdyProxyStats `json:"spdy-proxy"`
}


type SlbSpdyProxyStatsStats struct {
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Curr_http_proxy int `json:"curr_http_proxy"`
    Total_http_proxy int `json:"total_http_proxy"`
    Total_v2_proxy int `json:"total_v2_proxy"`
    Total_v3_proxy int `json:"total_v3_proxy"`
    Curr_stream int `json:"curr_stream"`
    Total_stream int `json:"total_stream"`
    Total_stream_succ int `json:"total_stream_succ"`
    Client_rst int `json:"client_rst"`
    Server_rst int `json:"server_rst"`
    Client_goaway int `json:"client_goaway"`
    Server_goaway int `json:"server_goaway"`
    Tcp_err int `json:"tcp_err"`
    Inflate_ctx int `json:"inflate_ctx"`
    Deflate_ctx int `json:"deflate_ctx"`
    Ping_sent int `json:"ping_sent"`
    Stream_not_found int `json:"stream_not_found"`
    Client_fin int `json:"client_fin"`
    Server_fin int `json:"server_fin"`
    Stream_close int `json:"stream_close"`
    Stream_err int `json:"stream_err"`
    Session_err int `json:"session_err"`
    Control_frame int `json:"control_frame"`
    Syn_frame int `json:"syn_frame"`
    Syn_reply_frame int `json:"syn_reply_frame"`
    Headers_frame int `json:"headers_frame"`
    Settings_frame int `json:"settings_frame"`
    Window_frame int `json:"window_frame"`
    Ping_frame int `json:"ping_frame"`
    Data_frame int `json:"data_frame"`
    Data_no_stream int `json:"data_no_stream"`
    Data_no_stream_no_goaway int `json:"data_no_stream_no_goaway"`
    Data_no_stream_goaway_close int `json:"data_no_stream_goaway_close"`
    Est_cb_no_tuple int `json:"est_cb_no_tuple"`
    Data_cb_no_tuple int `json:"data_cb_no_tuple"`
    Ctx_alloc_fail int `json:"ctx_alloc_fail"`
    Fin_close_session int `json:"fin_close_session"`
    Server_rst_close_stream int `json:"server_rst_close_stream"`
    Stream_found int `json:"stream_found"`
    Close_stream_session_not_found int `json:"close_stream_session_not_found"`
    Close_stream_stream_not_found int `json:"close_stream_stream_not_found"`
    Close_stream_already_closed int `json:"close_stream_already_closed"`
    Close_stream_session_close int `json:"close_stream_session_close"`
    Close_session_already_closed int `json:"close_session_already_closed"`
    Max_concurrent_stream_limit int `json:"max_concurrent_stream_limit"`
    Stream_alloc_fail int `json:"stream_alloc_fail"`
    Http_conn_alloc_fail int `json:"http_conn_alloc_fail"`
    Request_header_alloc_fail int `json:"request_header_alloc_fail"`
    Name_value_total_len_ex int `json:"name_value_total_len_ex"`
    Name_value_zero_len int `json:"name_value_zero_len"`
    Name_value_invalid_http_ver int `json:"name_value_invalid_http_ver"`
    Name_value_connection int `json:"name_value_connection"`
    Name_value_keepalive int `json:"name_value_keepalive"`
    Name_value_proxy_conn int `json:"name_value_proxy_conn"`
    Name_value_trasnfer_encod int `json:"name_value_trasnfer_encod"`
    Name_value_no_must_have int `json:"name_value_no_must_have"`
    Decompress_fail int `json:"decompress_fail"`
    Syn_after_goaway int `json:"syn_after_goaway"`
    Stream_lt_prev int `json:"stream_lt_prev"`
    Syn_stream_exist_or_even int `json:"syn_stream_exist_or_even"`
    Syn_unidir int `json:"syn_unidir"`
    Syn_reply_alr_rcvd int `json:"syn_reply_alr_rcvd"`
    Client_rst_nostream int `json:"client_rst_nostream"`
    Window_no_stream int `json:"window_no_stream"`
    Invalid_window_size int `json:"invalid_window_size"`
    Unknown_control_frame int `json:"unknown_control_frame"`
    Data_on_closed_stream int `json:"data_on_closed_stream"`
    Invalid_frame_size int `json:"invalid_frame_size"`
    Invalid_version int `json:"invalid_version"`
    Header_after_session_close int `json:"header_after_session_close"`
    Compress_ctx_alloc_fail int `json:"compress_ctx_alloc_fail"`
    Header_compress_fail int `json:"header_compress_fail"`
    Http_data_session_close int `json:"http_data_session_close"`
    Http_data_stream_not_found int `json:"http_data_stream_not_found"`
    Close_stream_not_http_proxy int `json:"close_stream_not_http_proxy"`
    Session_needs_requeue int `json:"session_needs_requeue"`
    New_stream_session_del int `json:"new_stream_session_del"`
    Fin_stream_closed int `json:"fin_stream_closed"`
    Http_close_stream_closed int `json:"http_close_stream_closed"`
    Http_err_stream_closed int `json:"http_err_stream_closed"`
    Http_hdr_stream_close int `json:"http_hdr_stream_close"`
    Http_data_stream_close int `json:"http_data_stream_close"`
    Session_close int `json:"session_close"`
}

func (p *SlbSpdyProxyStats) GetId() string{
    return "1"
}

func (p *SlbSpdyProxyStats) getPath() string{
    return "slb/spdy-proxy/stats"
}

func (p *SlbSpdyProxyStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSpdyProxyStats,error) {
logger.Println("SlbSpdyProxyStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSpdyProxyStats
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
