package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSpdyProxyOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_spdy_proxy_oper`: Operational Status for the object spdy-proxy\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSpdyProxyOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"curr_http_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_http_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_v2_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_v3_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"curr_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_stream_succ": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_goaway": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_goaway": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inflate_ctx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"deflate_ctx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ping_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_reply_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"headers_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"settings_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"window_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ping_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_no_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_no_stream_no_goaway": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_no_stream_goaway_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"est_cb_no_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_cb_no_tuple": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ctx_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fin_close_session": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_rst_close_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_stream_session_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_stream_stream_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_stream_already_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_stream_session_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_session_already_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"max_concurrent_stream_limit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_conn_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"request_header_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_total_len_ex": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_zero_len": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_invalid_http_ver": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_connection": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_keepalive": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_proxy_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_trasnfer_encod": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"name_value_no_must_have": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"decompress_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_after_goaway": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_lt_prev": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_stream_exist_or_even": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_unidir": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syn_reply_alr_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"client_rst_nostream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"window_no_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_window_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unknown_control_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_on_closed_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_frame_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_version": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_after_session_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compress_ctx_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_compress_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_data_session_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_data_stream_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"close_stream_not_http_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_needs_requeue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"new_stream_session_del": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fin_stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_close_stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_err_stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_hdr_stream_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_data_stream_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"session_close": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSpdyProxyOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSpdyProxyOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSpdyProxyOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSpdyProxyOperOper := setObjectSlbSpdyProxyOperOper(res)
		d.Set("oper", SlbSpdyProxyOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSpdyProxyOperOper(ret edpt.DataSlbSpdyProxyOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l4_cpu_list": setSliceSlbSpdyProxyOperOperL4CpuList(ret.DtSlbSpdyProxyOper.Oper.L4CpuList),
			"cpu_count":   ret.DtSlbSpdyProxyOper.Oper.CpuCount,
		},
	}
}

func setSliceSlbSpdyProxyOperOperL4CpuList(d []edpt.SlbSpdyProxyOperOperL4CpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr_proxy"] = item.Curr_proxy
		in["total_proxy"] = item.Total_proxy
		in["curr_http_proxy"] = item.Curr_http_proxy
		in["total_http_proxy"] = item.Total_http_proxy
		in["total_v2_proxy"] = item.Total_v2_proxy
		in["total_v3_proxy"] = item.Total_v3_proxy
		in["curr_stream"] = item.Curr_stream
		in["total_stream"] = item.Total_stream
		in["total_stream_succ"] = item.Total_stream_succ
		in["client_rst"] = item.Client_rst
		in["server_rst"] = item.Server_rst
		in["client_goaway"] = item.Client_goaway
		in["server_goaway"] = item.Server_goaway
		in["tcp_err"] = item.Tcp_err
		in["inflate_ctx"] = item.Inflate_ctx
		in["deflate_ctx"] = item.Deflate_ctx
		in["ping_sent"] = item.Ping_sent
		in["stream_not_found"] = item.Stream_not_found
		in["client_fin"] = item.Client_fin
		in["server_fin"] = item.Server_fin
		in["stream_close"] = item.Stream_close
		in["stream_err"] = item.Stream_err
		in["session_err"] = item.Session_err
		in["control_frame"] = item.Control_frame
		in["syn_frame"] = item.Syn_frame
		in["syn_reply_frame"] = item.Syn_reply_frame
		in["headers_frame"] = item.Headers_frame
		in["settings_frame"] = item.Settings_frame
		in["window_frame"] = item.Window_frame
		in["ping_frame"] = item.Ping_frame
		in["data_frame"] = item.Data_frame
		in["data_no_stream"] = item.Data_no_stream
		in["data_no_stream_no_goaway"] = item.Data_no_stream_no_goaway
		in["data_no_stream_goaway_close"] = item.Data_no_stream_goaway_close
		in["est_cb_no_tuple"] = item.Est_cb_no_tuple
		in["data_cb_no_tuple"] = item.Data_cb_no_tuple
		in["ctx_alloc_fail"] = item.Ctx_alloc_fail
		in["fin_close_session"] = item.Fin_close_session
		in["server_rst_close_stream"] = item.Server_rst_close_stream
		in["stream_found"] = item.Stream_found
		in["close_stream_session_not_found"] = item.Close_stream_session_not_found
		in["close_stream_stream_not_found"] = item.Close_stream_stream_not_found
		in["close_stream_already_closed"] = item.Close_stream_already_closed
		in["close_stream_session_close"] = item.Close_stream_session_close
		in["close_session_already_closed"] = item.Close_session_already_closed
		in["max_concurrent_stream_limit"] = item.Max_concurrent_stream_limit
		in["stream_alloc_fail"] = item.Stream_alloc_fail
		in["http_conn_alloc_fail"] = item.Http_conn_alloc_fail
		in["request_header_alloc_fail"] = item.Request_header_alloc_fail
		in["name_value_total_len_ex"] = item.Name_value_total_len_ex
		in["name_value_zero_len"] = item.Name_value_zero_len
		in["name_value_invalid_http_ver"] = item.Name_value_invalid_http_ver
		in["name_value_connection"] = item.Name_value_connection
		in["name_value_keepalive"] = item.Name_value_keepalive
		in["name_value_proxy_conn"] = item.Name_value_proxy_conn
		in["name_value_trasnfer_encod"] = item.Name_value_trasnfer_encod
		in["name_value_no_must_have"] = item.Name_value_no_must_have
		in["decompress_fail"] = item.Decompress_fail
		in["syn_after_goaway"] = item.Syn_after_goaway
		in["stream_lt_prev"] = item.Stream_lt_prev
		in["syn_stream_exist_or_even"] = item.Syn_stream_exist_or_even
		in["syn_unidir"] = item.Syn_unidir
		in["syn_reply_alr_rcvd"] = item.Syn_reply_alr_rcvd
		in["client_rst_nostream"] = item.Client_rst_nostream
		in["window_no_stream"] = item.Window_no_stream
		in["invalid_window_size"] = item.Invalid_window_size
		in["unknown_control_frame"] = item.Unknown_control_frame
		in["data_on_closed_stream"] = item.Data_on_closed_stream
		in["invalid_frame_size"] = item.Invalid_frame_size
		in["invalid_version"] = item.Invalid_version
		in["header_after_session_close"] = item.Header_after_session_close
		in["compress_ctx_alloc_fail"] = item.Compress_ctx_alloc_fail
		in["header_compress_fail"] = item.Header_compress_fail
		in["http_data_session_close"] = item.Http_data_session_close
		in["http_data_stream_not_found"] = item.Http_data_stream_not_found
		in["close_stream_not_http_proxy"] = item.Close_stream_not_http_proxy
		in["session_needs_requeue"] = item.Session_needs_requeue
		in["new_stream_session_del"] = item.New_stream_session_del
		in["fin_stream_closed"] = item.Fin_stream_closed
		in["http_close_stream_closed"] = item.Http_close_stream_closed
		in["http_err_stream_closed"] = item.Http_err_stream_closed
		in["http_hdr_stream_close"] = item.Http_hdr_stream_close
		in["http_data_stream_close"] = item.Http_data_stream_close
		in["session_close"] = item.Session_close
		result = append(result, in)
	}
	return result
}

func getObjectSlbSpdyProxyOperOper(d []interface{}) edpt.SlbSpdyProxyOperOper {

	count1 := len(d)
	var ret edpt.SlbSpdyProxyOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4CpuList = getSliceSlbSpdyProxyOperOperL4CpuList(in["l4_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbSpdyProxyOperOperL4CpuList(d []interface{}) []edpt.SlbSpdyProxyOperOperL4CpuList {

	count1 := len(d)
	ret := make([]edpt.SlbSpdyProxyOperOperL4CpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSpdyProxyOperOperL4CpuList
		oi.Curr_proxy = in["curr_proxy"].(int)
		oi.Total_proxy = in["total_proxy"].(int)
		oi.Curr_http_proxy = in["curr_http_proxy"].(int)
		oi.Total_http_proxy = in["total_http_proxy"].(int)
		oi.Total_v2_proxy = in["total_v2_proxy"].(int)
		oi.Total_v3_proxy = in["total_v3_proxy"].(int)
		oi.Curr_stream = in["curr_stream"].(int)
		oi.Total_stream = in["total_stream"].(int)
		oi.Total_stream_succ = in["total_stream_succ"].(int)
		oi.Client_rst = in["client_rst"].(int)
		oi.Server_rst = in["server_rst"].(int)
		oi.Client_goaway = in["client_goaway"].(int)
		oi.Server_goaway = in["server_goaway"].(int)
		oi.Tcp_err = in["tcp_err"].(int)
		oi.Inflate_ctx = in["inflate_ctx"].(int)
		oi.Deflate_ctx = in["deflate_ctx"].(int)
		oi.Ping_sent = in["ping_sent"].(int)
		oi.Stream_not_found = in["stream_not_found"].(int)
		oi.Client_fin = in["client_fin"].(int)
		oi.Server_fin = in["server_fin"].(int)
		oi.Stream_close = in["stream_close"].(int)
		oi.Stream_err = in["stream_err"].(int)
		oi.Session_err = in["session_err"].(int)
		oi.Control_frame = in["control_frame"].(int)
		oi.Syn_frame = in["syn_frame"].(int)
		oi.Syn_reply_frame = in["syn_reply_frame"].(int)
		oi.Headers_frame = in["headers_frame"].(int)
		oi.Settings_frame = in["settings_frame"].(int)
		oi.Window_frame = in["window_frame"].(int)
		oi.Ping_frame = in["ping_frame"].(int)
		oi.Data_frame = in["data_frame"].(int)
		oi.Data_no_stream = in["data_no_stream"].(int)
		oi.Data_no_stream_no_goaway = in["data_no_stream_no_goaway"].(int)
		oi.Data_no_stream_goaway_close = in["data_no_stream_goaway_close"].(int)
		oi.Est_cb_no_tuple = in["est_cb_no_tuple"].(int)
		oi.Data_cb_no_tuple = in["data_cb_no_tuple"].(int)
		oi.Ctx_alloc_fail = in["ctx_alloc_fail"].(int)
		oi.Fin_close_session = in["fin_close_session"].(int)
		oi.Server_rst_close_stream = in["server_rst_close_stream"].(int)
		oi.Stream_found = in["stream_found"].(int)
		oi.Close_stream_session_not_found = in["close_stream_session_not_found"].(int)
		oi.Close_stream_stream_not_found = in["close_stream_stream_not_found"].(int)
		oi.Close_stream_already_closed = in["close_stream_already_closed"].(int)
		oi.Close_stream_session_close = in["close_stream_session_close"].(int)
		oi.Close_session_already_closed = in["close_session_already_closed"].(int)
		oi.Max_concurrent_stream_limit = in["max_concurrent_stream_limit"].(int)
		oi.Stream_alloc_fail = in["stream_alloc_fail"].(int)
		oi.Http_conn_alloc_fail = in["http_conn_alloc_fail"].(int)
		oi.Request_header_alloc_fail = in["request_header_alloc_fail"].(int)
		oi.Name_value_total_len_ex = in["name_value_total_len_ex"].(int)
		oi.Name_value_zero_len = in["name_value_zero_len"].(int)
		oi.Name_value_invalid_http_ver = in["name_value_invalid_http_ver"].(int)
		oi.Name_value_connection = in["name_value_connection"].(int)
		oi.Name_value_keepalive = in["name_value_keepalive"].(int)
		oi.Name_value_proxy_conn = in["name_value_proxy_conn"].(int)
		oi.Name_value_trasnfer_encod = in["name_value_trasnfer_encod"].(int)
		oi.Name_value_no_must_have = in["name_value_no_must_have"].(int)
		oi.Decompress_fail = in["decompress_fail"].(int)
		oi.Syn_after_goaway = in["syn_after_goaway"].(int)
		oi.Stream_lt_prev = in["stream_lt_prev"].(int)
		oi.Syn_stream_exist_or_even = in["syn_stream_exist_or_even"].(int)
		oi.Syn_unidir = in["syn_unidir"].(int)
		oi.Syn_reply_alr_rcvd = in["syn_reply_alr_rcvd"].(int)
		oi.Client_rst_nostream = in["client_rst_nostream"].(int)
		oi.Window_no_stream = in["window_no_stream"].(int)
		oi.Invalid_window_size = in["invalid_window_size"].(int)
		oi.Unknown_control_frame = in["unknown_control_frame"].(int)
		oi.Data_on_closed_stream = in["data_on_closed_stream"].(int)
		oi.Invalid_frame_size = in["invalid_frame_size"].(int)
		oi.Invalid_version = in["invalid_version"].(int)
		oi.Header_after_session_close = in["header_after_session_close"].(int)
		oi.Compress_ctx_alloc_fail = in["compress_ctx_alloc_fail"].(int)
		oi.Header_compress_fail = in["header_compress_fail"].(int)
		oi.Http_data_session_close = in["http_data_session_close"].(int)
		oi.Http_data_stream_not_found = in["http_data_stream_not_found"].(int)
		oi.Close_stream_not_http_proxy = in["close_stream_not_http_proxy"].(int)
		oi.Session_needs_requeue = in["session_needs_requeue"].(int)
		oi.New_stream_session_del = in["new_stream_session_del"].(int)
		oi.Fin_stream_closed = in["fin_stream_closed"].(int)
		oi.Http_close_stream_closed = in["http_close_stream_closed"].(int)
		oi.Http_err_stream_closed = in["http_err_stream_closed"].(int)
		oi.Http_hdr_stream_close = in["http_hdr_stream_close"].(int)
		oi.Http_data_stream_close = in["http_data_stream_close"].(int)
		oi.Session_close = in["session_close"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSpdyProxyOper(d *schema.ResourceData) edpt.SlbSpdyProxyOper {
	var ret edpt.SlbSpdyProxyOper

	ret.Oper = getObjectSlbSpdyProxyOperOper(d.Get("oper").([]interface{}))
	return ret
}
