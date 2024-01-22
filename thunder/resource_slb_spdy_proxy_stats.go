package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSpdyProxyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_spdy_proxy_stats`: Statistics for the object spdy-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSpdyProxyStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Curr Proxy Conns",
						},
						"total_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total Proxy Conns",
						},
						"curr_http_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Curr HTTP Proxy Conns",
						},
						"total_http_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP Proxy Conns",
						},
						"total_v2_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Version 2 Streams",
						},
						"total_v3_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Version 3 Streams",
						},
						"curr_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Curr Streams",
						},
						"total_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Total Streams",
						},
						"total_stream_succ": {
							Type: schema.TypeInt, Optional: true, Description: "Streams(succ)",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Server RST sent",
						},
						"client_goaway": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"server_goaway": {
							Type: schema.TypeInt, Optional: true, Description: "Server GOAWAY sent",
						},
						"tcp_err": {
							Type: schema.TypeInt, Optional: true, Description: "TCP sock error",
						},
						"inflate_ctx": {
							Type: schema.TypeInt, Optional: true, Description: "Inflate context",
						},
						"deflate_ctx": {
							Type: schema.TypeInt, Optional: true, Description: "Deflate context",
						},
						"ping_sent": {
							Type: schema.TypeInt, Optional: true, Description: "PING sent",
						},
						"stream_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "STREAM not found",
						},
						"client_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Client FIN",
						},
						"server_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Server FIN",
						},
						"stream_close": {
							Type: schema.TypeInt, Optional: true, Description: "Stream close",
						},
						"stream_err": {
							Type: schema.TypeInt, Optional: true, Description: "Stream err",
						},
						"session_err": {
							Type: schema.TypeInt, Optional: true, Description: "Session err",
						},
						"control_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Control frame received",
						},
						"syn_frame": {
							Type: schema.TypeInt, Optional: true, Description: "SYN stream frame received",
						},
						"syn_reply_frame": {
							Type: schema.TypeInt, Optional: true, Description: "SYN reply frame received",
						},
						"headers_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Headers frame received",
						},
						"settings_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Setting frame received",
						},
						"window_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Window update frame received",
						},
						"ping_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Ping frame received",
						},
						"data_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Data frame received",
						},
						"data_no_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Data no stream found",
						},
						"data_no_stream_no_goaway": {
							Type: schema.TypeInt, Optional: true, Description: "Data no stream and no goaway",
						},
						"data_no_stream_goaway_close": {
							Type: schema.TypeInt, Optional: true, Description: "Data no stream and no goaway and close session",
						},
						"est_cb_no_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Est callback no tuple",
						},
						"data_cb_no_tuple": {
							Type: schema.TypeInt, Optional: true, Description: "Data callback no tuple",
						},
						"ctx_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Context alloc fail",
						},
						"fin_close_session": {
							Type: schema.TypeInt, Optional: true, Description: "FIN close session",
						},
						"server_rst_close_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Server RST close stream",
						},
						"stream_found": {
							Type: schema.TypeInt, Optional: true, Description: "Stream found",
						},
						"close_stream_session_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Close stream session not found",
						},
						"close_stream_stream_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Close stream stream not found",
						},
						"close_stream_already_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Closing closed stream",
						},
						"close_stream_session_close": {
							Type: schema.TypeInt, Optional: true, Description: "Stream close session close",
						},
						"close_session_already_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Closing closed session",
						},
						"max_concurrent_stream_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Max concurrent stream limit",
						},
						"stream_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Stream alloc fail",
						},
						"http_conn_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP connection allocation fail",
						},
						"request_header_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Request/Header allocation fail",
						},
						"name_value_total_len_ex": {
							Type: schema.TypeInt, Optional: true, Description: "Name value total length exceeded",
						},
						"name_value_zero_len": {
							Type: schema.TypeInt, Optional: true, Description: "Name value zero name length",
						},
						"name_value_invalid_http_ver": {
							Type: schema.TypeInt, Optional: true, Description: "Name value invalid http version",
						},
						"name_value_connection": {
							Type: schema.TypeInt, Optional: true, Description: "Name value connection",
						},
						"name_value_keepalive": {
							Type: schema.TypeInt, Optional: true, Description: "Name value keep alive",
						},
						"name_value_proxy_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Name value proxy-connection",
						},
						"name_value_trasnfer_encod": {
							Type: schema.TypeInt, Optional: true, Description: "Name value transfer encoding",
						},
						"name_value_no_must_have": {
							Type: schema.TypeInt, Optional: true, Description: "Name value no must have",
						},
						"decompress_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Decompress fail",
						},
						"syn_after_goaway": {
							Type: schema.TypeInt, Optional: true, Description: "SYN after goaway",
						},
						"stream_lt_prev": {
							Type: schema.TypeInt, Optional: true, Description: "Stream id less than previous",
						},
						"syn_stream_exist_or_even": {
							Type: schema.TypeInt, Optional: true, Description: "Stream already exists",
						},
						"syn_unidir": {
							Type: schema.TypeInt, Optional: true, Description: "Unidirectional SYN",
						},
						"syn_reply_alr_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "SYN reply already received",
						},
						"client_rst_nostream": {
							Type: schema.TypeInt, Optional: true, Description: "Close RST stream not found",
						},
						"window_no_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Window update no stream found",
						},
						"invalid_window_size": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid window size",
						},
						"unknown_control_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown control frame",
						},
						"data_on_closed_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Data on closed stream",
						},
						"invalid_frame_size": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid frame size",
						},
						"invalid_version": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid version",
						},
						"header_after_session_close": {
							Type: schema.TypeInt, Optional: true, Description: "Header after session close",
						},
						"compress_ctx_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Compression context allocation fail",
						},
						"header_compress_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Header compress fail",
						},
						"http_data_session_close": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP data session close",
						},
						"http_data_stream_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP data stream not found",
						},
						"close_stream_not_http_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Close Stream not http-proxy",
						},
						"session_needs_requeue": {
							Type: schema.TypeInt, Optional: true, Description: "Session needs requeue",
						},
						"new_stream_session_del": {
							Type: schema.TypeInt, Optional: true, Description: "New Stream after Session delete",
						},
						"fin_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP FIN stream already closed",
						},
						"http_close_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP close stream already closed",
						},
						"http_err_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP error stream already closed",
						},
						"http_hdr_stream_close": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP header stream already closed",
						},
						"http_data_stream_close": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP data stream already closed",
						},
						"session_close": {
							Type: schema.TypeInt, Optional: true, Description: "Session close",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSpdyProxyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSpdyProxyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSpdyProxyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSpdyProxyStatsStats := setObjectSlbSpdyProxyStatsStats(res)
		d.Set("stats", SlbSpdyProxyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSpdyProxyStatsStats(ret edpt.DataSlbSpdyProxyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":                     ret.DtSlbSpdyProxyStats.Stats.Curr_proxy,
			"total_proxy":                    ret.DtSlbSpdyProxyStats.Stats.Total_proxy,
			"curr_http_proxy":                ret.DtSlbSpdyProxyStats.Stats.Curr_http_proxy,
			"total_http_proxy":               ret.DtSlbSpdyProxyStats.Stats.Total_http_proxy,
			"total_v2_proxy":                 ret.DtSlbSpdyProxyStats.Stats.Total_v2_proxy,
			"total_v3_proxy":                 ret.DtSlbSpdyProxyStats.Stats.Total_v3_proxy,
			"curr_stream":                    ret.DtSlbSpdyProxyStats.Stats.Curr_stream,
			"total_stream":                   ret.DtSlbSpdyProxyStats.Stats.Total_stream,
			"total_stream_succ":              ret.DtSlbSpdyProxyStats.Stats.Total_stream_succ,
			"client_rst":                     ret.DtSlbSpdyProxyStats.Stats.Client_rst,
			"server_rst":                     ret.DtSlbSpdyProxyStats.Stats.Server_rst,
			"client_goaway":                  ret.DtSlbSpdyProxyStats.Stats.Client_goaway,
			"server_goaway":                  ret.DtSlbSpdyProxyStats.Stats.Server_goaway,
			"tcp_err":                        ret.DtSlbSpdyProxyStats.Stats.Tcp_err,
			"inflate_ctx":                    ret.DtSlbSpdyProxyStats.Stats.Inflate_ctx,
			"deflate_ctx":                    ret.DtSlbSpdyProxyStats.Stats.Deflate_ctx,
			"ping_sent":                      ret.DtSlbSpdyProxyStats.Stats.Ping_sent,
			"stream_not_found":               ret.DtSlbSpdyProxyStats.Stats.Stream_not_found,
			"client_fin":                     ret.DtSlbSpdyProxyStats.Stats.Client_fin,
			"server_fin":                     ret.DtSlbSpdyProxyStats.Stats.Server_fin,
			"stream_close":                   ret.DtSlbSpdyProxyStats.Stats.Stream_close,
			"stream_err":                     ret.DtSlbSpdyProxyStats.Stats.Stream_err,
			"session_err":                    ret.DtSlbSpdyProxyStats.Stats.Session_err,
			"control_frame":                  ret.DtSlbSpdyProxyStats.Stats.Control_frame,
			"syn_frame":                      ret.DtSlbSpdyProxyStats.Stats.Syn_frame,
			"syn_reply_frame":                ret.DtSlbSpdyProxyStats.Stats.Syn_reply_frame,
			"headers_frame":                  ret.DtSlbSpdyProxyStats.Stats.Headers_frame,
			"settings_frame":                 ret.DtSlbSpdyProxyStats.Stats.Settings_frame,
			"window_frame":                   ret.DtSlbSpdyProxyStats.Stats.Window_frame,
			"ping_frame":                     ret.DtSlbSpdyProxyStats.Stats.Ping_frame,
			"data_frame":                     ret.DtSlbSpdyProxyStats.Stats.Data_frame,
			"data_no_stream":                 ret.DtSlbSpdyProxyStats.Stats.Data_no_stream,
			"data_no_stream_no_goaway":       ret.DtSlbSpdyProxyStats.Stats.Data_no_stream_no_goaway,
			"data_no_stream_goaway_close":    ret.DtSlbSpdyProxyStats.Stats.Data_no_stream_goaway_close,
			"est_cb_no_tuple":                ret.DtSlbSpdyProxyStats.Stats.Est_cb_no_tuple,
			"data_cb_no_tuple":               ret.DtSlbSpdyProxyStats.Stats.Data_cb_no_tuple,
			"ctx_alloc_fail":                 ret.DtSlbSpdyProxyStats.Stats.Ctx_alloc_fail,
			"fin_close_session":              ret.DtSlbSpdyProxyStats.Stats.Fin_close_session,
			"server_rst_close_stream":        ret.DtSlbSpdyProxyStats.Stats.Server_rst_close_stream,
			"stream_found":                   ret.DtSlbSpdyProxyStats.Stats.Stream_found,
			"close_stream_session_not_found": ret.DtSlbSpdyProxyStats.Stats.Close_stream_session_not_found,
			"close_stream_stream_not_found":  ret.DtSlbSpdyProxyStats.Stats.Close_stream_stream_not_found,
			"close_stream_already_closed":    ret.DtSlbSpdyProxyStats.Stats.Close_stream_already_closed,
			"close_stream_session_close":     ret.DtSlbSpdyProxyStats.Stats.Close_stream_session_close,
			"close_session_already_closed":   ret.DtSlbSpdyProxyStats.Stats.Close_session_already_closed,
			"max_concurrent_stream_limit":    ret.DtSlbSpdyProxyStats.Stats.Max_concurrent_stream_limit,
			"stream_alloc_fail":              ret.DtSlbSpdyProxyStats.Stats.Stream_alloc_fail,
			"http_conn_alloc_fail":           ret.DtSlbSpdyProxyStats.Stats.Http_conn_alloc_fail,
			"request_header_alloc_fail":      ret.DtSlbSpdyProxyStats.Stats.Request_header_alloc_fail,
			"name_value_total_len_ex":        ret.DtSlbSpdyProxyStats.Stats.Name_value_total_len_ex,
			"name_value_zero_len":            ret.DtSlbSpdyProxyStats.Stats.Name_value_zero_len,
			"name_value_invalid_http_ver":    ret.DtSlbSpdyProxyStats.Stats.Name_value_invalid_http_ver,
			"name_value_connection":          ret.DtSlbSpdyProxyStats.Stats.Name_value_connection,
			"name_value_keepalive":           ret.DtSlbSpdyProxyStats.Stats.Name_value_keepalive,
			"name_value_proxy_conn":          ret.DtSlbSpdyProxyStats.Stats.Name_value_proxy_conn,
			"name_value_trasnfer_encod":      ret.DtSlbSpdyProxyStats.Stats.Name_value_trasnfer_encod,
			"name_value_no_must_have":        ret.DtSlbSpdyProxyStats.Stats.Name_value_no_must_have,
			"decompress_fail":                ret.DtSlbSpdyProxyStats.Stats.Decompress_fail,
			"syn_after_goaway":               ret.DtSlbSpdyProxyStats.Stats.Syn_after_goaway,
			"stream_lt_prev":                 ret.DtSlbSpdyProxyStats.Stats.Stream_lt_prev,
			"syn_stream_exist_or_even":       ret.DtSlbSpdyProxyStats.Stats.Syn_stream_exist_or_even,
			"syn_unidir":                     ret.DtSlbSpdyProxyStats.Stats.Syn_unidir,
			"syn_reply_alr_rcvd":             ret.DtSlbSpdyProxyStats.Stats.Syn_reply_alr_rcvd,
			"client_rst_nostream":            ret.DtSlbSpdyProxyStats.Stats.Client_rst_nostream,
			"window_no_stream":               ret.DtSlbSpdyProxyStats.Stats.Window_no_stream,
			"invalid_window_size":            ret.DtSlbSpdyProxyStats.Stats.Invalid_window_size,
			"unknown_control_frame":          ret.DtSlbSpdyProxyStats.Stats.Unknown_control_frame,
			"data_on_closed_stream":          ret.DtSlbSpdyProxyStats.Stats.Data_on_closed_stream,
			"invalid_frame_size":             ret.DtSlbSpdyProxyStats.Stats.Invalid_frame_size,
			"invalid_version":                ret.DtSlbSpdyProxyStats.Stats.Invalid_version,
			"header_after_session_close":     ret.DtSlbSpdyProxyStats.Stats.Header_after_session_close,
			"compress_ctx_alloc_fail":        ret.DtSlbSpdyProxyStats.Stats.Compress_ctx_alloc_fail,
			"header_compress_fail":           ret.DtSlbSpdyProxyStats.Stats.Header_compress_fail,
			"http_data_session_close":        ret.DtSlbSpdyProxyStats.Stats.Http_data_session_close,
			"http_data_stream_not_found":     ret.DtSlbSpdyProxyStats.Stats.Http_data_stream_not_found,
			"close_stream_not_http_proxy":    ret.DtSlbSpdyProxyStats.Stats.Close_stream_not_http_proxy,
			"session_needs_requeue":          ret.DtSlbSpdyProxyStats.Stats.Session_needs_requeue,
			"new_stream_session_del":         ret.DtSlbSpdyProxyStats.Stats.New_stream_session_del,
			"fin_stream_closed":              ret.DtSlbSpdyProxyStats.Stats.Fin_stream_closed,
			"http_close_stream_closed":       ret.DtSlbSpdyProxyStats.Stats.Http_close_stream_closed,
			"http_err_stream_closed":         ret.DtSlbSpdyProxyStats.Stats.Http_err_stream_closed,
			"http_hdr_stream_close":          ret.DtSlbSpdyProxyStats.Stats.Http_hdr_stream_close,
			"http_data_stream_close":         ret.DtSlbSpdyProxyStats.Stats.Http_data_stream_close,
			"session_close":                  ret.DtSlbSpdyProxyStats.Stats.Session_close,
		},
	}
}

func getObjectSlbSpdyProxyStatsStats(d []interface{}) edpt.SlbSpdyProxyStatsStats {

	count1 := len(d)
	var ret edpt.SlbSpdyProxyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Curr_http_proxy = in["curr_http_proxy"].(int)
		ret.Total_http_proxy = in["total_http_proxy"].(int)
		ret.Total_v2_proxy = in["total_v2_proxy"].(int)
		ret.Total_v3_proxy = in["total_v3_proxy"].(int)
		ret.Curr_stream = in["curr_stream"].(int)
		ret.Total_stream = in["total_stream"].(int)
		ret.Total_stream_succ = in["total_stream_succ"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Server_rst = in["server_rst"].(int)
		ret.Client_goaway = in["client_goaway"].(int)
		ret.Server_goaway = in["server_goaway"].(int)
		ret.Tcp_err = in["tcp_err"].(int)
		ret.Inflate_ctx = in["inflate_ctx"].(int)
		ret.Deflate_ctx = in["deflate_ctx"].(int)
		ret.Ping_sent = in["ping_sent"].(int)
		ret.Stream_not_found = in["stream_not_found"].(int)
		ret.Client_fin = in["client_fin"].(int)
		ret.Server_fin = in["server_fin"].(int)
		ret.Stream_close = in["stream_close"].(int)
		ret.Stream_err = in["stream_err"].(int)
		ret.Session_err = in["session_err"].(int)
		ret.Control_frame = in["control_frame"].(int)
		ret.Syn_frame = in["syn_frame"].(int)
		ret.Syn_reply_frame = in["syn_reply_frame"].(int)
		ret.Headers_frame = in["headers_frame"].(int)
		ret.Settings_frame = in["settings_frame"].(int)
		ret.Window_frame = in["window_frame"].(int)
		ret.Ping_frame = in["ping_frame"].(int)
		ret.Data_frame = in["data_frame"].(int)
		ret.Data_no_stream = in["data_no_stream"].(int)
		ret.Data_no_stream_no_goaway = in["data_no_stream_no_goaway"].(int)
		ret.Data_no_stream_goaway_close = in["data_no_stream_goaway_close"].(int)
		ret.Est_cb_no_tuple = in["est_cb_no_tuple"].(int)
		ret.Data_cb_no_tuple = in["data_cb_no_tuple"].(int)
		ret.Ctx_alloc_fail = in["ctx_alloc_fail"].(int)
		ret.Fin_close_session = in["fin_close_session"].(int)
		ret.Server_rst_close_stream = in["server_rst_close_stream"].(int)
		ret.Stream_found = in["stream_found"].(int)
		ret.Close_stream_session_not_found = in["close_stream_session_not_found"].(int)
		ret.Close_stream_stream_not_found = in["close_stream_stream_not_found"].(int)
		ret.Close_stream_already_closed = in["close_stream_already_closed"].(int)
		ret.Close_stream_session_close = in["close_stream_session_close"].(int)
		ret.Close_session_already_closed = in["close_session_already_closed"].(int)
		ret.Max_concurrent_stream_limit = in["max_concurrent_stream_limit"].(int)
		ret.Stream_alloc_fail = in["stream_alloc_fail"].(int)
		ret.Http_conn_alloc_fail = in["http_conn_alloc_fail"].(int)
		ret.Request_header_alloc_fail = in["request_header_alloc_fail"].(int)
		ret.Name_value_total_len_ex = in["name_value_total_len_ex"].(int)
		ret.Name_value_zero_len = in["name_value_zero_len"].(int)
		ret.Name_value_invalid_http_ver = in["name_value_invalid_http_ver"].(int)
		ret.Name_value_connection = in["name_value_connection"].(int)
		ret.Name_value_keepalive = in["name_value_keepalive"].(int)
		ret.Name_value_proxy_conn = in["name_value_proxy_conn"].(int)
		ret.Name_value_trasnfer_encod = in["name_value_trasnfer_encod"].(int)
		ret.Name_value_no_must_have = in["name_value_no_must_have"].(int)
		ret.Decompress_fail = in["decompress_fail"].(int)
		ret.Syn_after_goaway = in["syn_after_goaway"].(int)
		ret.Stream_lt_prev = in["stream_lt_prev"].(int)
		ret.Syn_stream_exist_or_even = in["syn_stream_exist_or_even"].(int)
		ret.Syn_unidir = in["syn_unidir"].(int)
		ret.Syn_reply_alr_rcvd = in["syn_reply_alr_rcvd"].(int)
		ret.Client_rst_nostream = in["client_rst_nostream"].(int)
		ret.Window_no_stream = in["window_no_stream"].(int)
		ret.Invalid_window_size = in["invalid_window_size"].(int)
		ret.Unknown_control_frame = in["unknown_control_frame"].(int)
		ret.Data_on_closed_stream = in["data_on_closed_stream"].(int)
		ret.Invalid_frame_size = in["invalid_frame_size"].(int)
		ret.Invalid_version = in["invalid_version"].(int)
		ret.Header_after_session_close = in["header_after_session_close"].(int)
		ret.Compress_ctx_alloc_fail = in["compress_ctx_alloc_fail"].(int)
		ret.Header_compress_fail = in["header_compress_fail"].(int)
		ret.Http_data_session_close = in["http_data_session_close"].(int)
		ret.Http_data_stream_not_found = in["http_data_stream_not_found"].(int)
		ret.Close_stream_not_http_proxy = in["close_stream_not_http_proxy"].(int)
		ret.Session_needs_requeue = in["session_needs_requeue"].(int)
		ret.New_stream_session_del = in["new_stream_session_del"].(int)
		ret.Fin_stream_closed = in["fin_stream_closed"].(int)
		ret.Http_close_stream_closed = in["http_close_stream_closed"].(int)
		ret.Http_err_stream_closed = in["http_err_stream_closed"].(int)
		ret.Http_hdr_stream_close = in["http_hdr_stream_close"].(int)
		ret.Http_data_stream_close = in["http_data_stream_close"].(int)
		ret.Session_close = in["session_close"].(int)
	}
	return ret
}

func dataToEndpointSlbSpdyProxyStats(d *schema.ResourceData) edpt.SlbSpdyProxyStats {
	var ret edpt.SlbSpdyProxyStats

	ret.Stats = getObjectSlbSpdyProxyStatsStats(d.Get("stats").([]interface{}))
	return ret
}
