package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttp2Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_http2_oper`: Operational Status for the object http2\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHttp2OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http2_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"curr_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"peak_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_proxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_preface_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connection_preface_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"continuation_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"continuation_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_to_app": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"goaway_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"goaway_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"headers_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"headers_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_to_app": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ping_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ping_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"priority_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"priority_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rst_frame_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rst_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"settings_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"settings_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"settings_ack_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"settings_ack_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"empty_settings_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"empty_settings_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"window_update_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"window_update_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unknown_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"split_buff_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_frame_size": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"error_max_invalid_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_no_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_connection_preface": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bad_frame_type_for_stream_state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"wrong_stream_state": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"alloc_fail_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxy_alloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"deflate_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inflate_alloc_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_queue_alloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"buff_alloc_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_control_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_settings_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_rst_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_goaway_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_ping_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cant_allocate_window_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inflate_header_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_no_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_padlen_gt_frame_payload": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"streams_gt_max_concur_streams": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"idle_state_unexpected_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reserved_local_state_unexpected_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reserved_remote_state_unexpected_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"half_closed_remote_state_unexpected_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"closed_state_unexpected_frame": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"zero_window_size_on_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"exceeds_max_window_size_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"continuation_before_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_frame_during_headers": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"headers_after_continuation": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_push_promise": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_stream_id": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"headers_interleaved": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"trailers_no_end_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_setting_value": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"invalid_window_update": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_rcvd_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"protocol_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"internal_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"flow_control_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"settings_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frame_size_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"refused_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"cancel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"compression_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connect_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"enhance_your_calm": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inadequate_security": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http_1_1_required": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_total": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_proto_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_internal_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_flow_control": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_setting_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_stream_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_frame_size_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_refused_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_cancel": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_compression_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_connect_err": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_your_calm": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_inadequate_security": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"err_sent_http11_required": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frame_header_bytes_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"frame_header_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_bytes_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"control_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_bytes_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"header_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_bytes_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"data_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_bytes_received": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"total_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"push_promise_frame_sent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"http2_rejected": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"current_stream": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_create": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stream_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_stream_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"end_stream_sent": {
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

func resourceSlbHttp2OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp2OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp2Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHttp2OperOper := setObjectSlbHttp2OperOper(res)
		d.Set("oper", SlbHttp2OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHttp2OperOper(ret edpt.DataSlbHttp2Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"http2_cpu_list": setSliceSlbHttp2OperOperHttp2CpuList(ret.DtSlbHttp2Oper.Oper.Http2CpuList),
			"cpu_count":      ret.DtSlbHttp2Oper.Oper.CpuCount,
		},
	}
}

func setSliceSlbHttp2OperOperHttp2CpuList(d []edpt.SlbHttp2OperOperHttp2CpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["curr_proxy"] = item.Curr_proxy
		in["peak_proxy"] = item.Peak_proxy
		in["total_proxy"] = item.Total_proxy
		in["connection_preface_rcvd"] = item.Connection_preface_rcvd
		in["connection_preface_sent"] = item.Connection_preface_sent
		in["control_frame"] = item.Control_frame
		in["control_frame_sent"] = item.Control_frame_sent
		in["continuation_frame"] = item.Continuation_frame
		in["continuation_frame_sent"] = item.Continuation_frame_sent
		in["data_frame"] = item.Data_frame
		in["data_frame_sent"] = item.Data_frame_sent
		in["data_to_app"] = item.Data_to_app
		in["goaway_frame"] = item.Goaway_frame
		in["goaway_frame_sent"] = item.Goaway_frame_sent
		in["headers_frame"] = item.Headers_frame
		in["headers_frame_sent"] = item.Headers_frame_sent
		in["header_to_app"] = item.Header_to_app
		in["ping_frame"] = item.Ping_frame
		in["ping_frame_sent"] = item.Ping_frame_sent
		in["priority_frame"] = item.Priority_frame
		in["priority_frame_sent"] = item.Priority_frame_sent
		in["rst_frame_rcvd"] = item.Rst_frame_rcvd
		in["rst_frame_sent"] = item.Rst_frame_sent
		in["settings_frame"] = item.Settings_frame
		in["settings_frame_sent"] = item.Settings_frame_sent
		in["settings_ack_rcvd"] = item.Settings_ack_rcvd
		in["settings_ack_sent"] = item.Settings_ack_sent
		in["empty_settings_rcvd"] = item.Empty_settings_rcvd
		in["empty_settings_sent"] = item.Empty_settings_sent
		in["window_update_frame"] = item.Window_update_frame
		in["window_update_frame_sent"] = item.Window_update_frame_sent
		in["unknown_frame"] = item.Unknown_frame
		in["split_buff_fail"] = item.Split_buff_fail
		in["invalid_frame_size"] = item.Invalid_frame_size
		in["error_max_invalid_stream"] = item.Error_max_invalid_stream
		in["data_no_stream"] = item.Data_no_stream
		in["bad_connection_preface"] = item.Bad_connection_preface
		in["bad_frame_type_for_stream_state"] = item.Bad_frame_type_for_stream_state
		in["wrong_stream_state"] = item.Wrong_stream_state
		in["alloc_fail_total"] = item.Alloc_fail_total
		in["proxy_alloc_error"] = item.Proxy_alloc_error
		in["deflate_alloc_fail"] = item.Deflate_alloc_fail
		in["inflate_alloc_fail"] = item.Inflate_alloc_fail
		in["data_queue_alloc_error"] = item.Data_queue_alloc_error
		in["buff_alloc_error"] = item.Buff_alloc_error
		in["cant_allocate_control_frame"] = item.Cant_allocate_control_frame
		in["cant_allocate_settings_frame"] = item.Cant_allocate_settings_frame
		in["cant_allocate_rst_frame"] = item.Cant_allocate_rst_frame
		in["cant_allocate_goaway_frame"] = item.Cant_allocate_goaway_frame
		in["cant_allocate_ping_frame"] = item.Cant_allocate_ping_frame
		in["cant_allocate_stream"] = item.Cant_allocate_stream
		in["cant_allocate_window_frame"] = item.Cant_allocate_window_frame
		in["inflate_header_fail"] = item.Inflate_header_fail
		in["header_no_stream"] = item.Header_no_stream
		in["header_padlen_gt_frame_payload"] = item.Header_padlen_gt_frame_payload
		in["streams_gt_max_concur_streams"] = item.Streams_gt_max_concur_streams
		in["idle_state_unexpected_frame"] = item.Idle_state_unexpected_frame
		in["reserved_local_state_unexpected_frame"] = item.Reserved_local_state_unexpected_frame
		in["reserved_remote_state_unexpected_frame"] = item.Reserved_remote_state_unexpected_frame
		in["half_closed_remote_state_unexpected_frame"] = item.Half_closed_remote_state_unexpected_frame
		in["closed_state_unexpected_frame"] = item.Closed_state_unexpected_frame
		in["zero_window_size_on_stream"] = item.Zero_window_size_on_stream
		in["exceeds_max_window_size_stream"] = item.Exceeds_max_window_size_stream
		in["continuation_before_headers"] = item.Continuation_before_headers
		in["invalid_frame_during_headers"] = item.Invalid_frame_during_headers
		in["headers_after_continuation"] = item.Headers_after_continuation
		in["invalid_push_promise"] = item.Invalid_push_promise
		in["invalid_stream_id"] = item.Invalid_stream_id
		in["headers_interleaved"] = item.Headers_interleaved
		in["trailers_no_end_stream"] = item.Trailers_no_end_stream
		in["invalid_setting_value"] = item.Invalid_setting_value
		in["invalid_window_update"] = item.Invalid_window_update
		in["err_rcvd_total"] = item.Err_rcvd_total
		in["protocol_error"] = item.Protocol_error
		in["internal_error"] = item.Internal_error
		in["flow_control_error"] = item.Flow_control_error
		in["settings_timeout"] = item.Settings_timeout
		in["stream_closed"] = item.Stream_closed
		in["frame_size_error"] = item.Frame_size_error
		in["refused_stream"] = item.Refused_stream
		in["cancel"] = item.Cancel
		in["compression_error"] = item.Compression_error
		in["connect_error"] = item.Connect_error
		in["enhance_your_calm"] = item.Enhance_your_calm
		in["inadequate_security"] = item.Inadequate_security
		in["http_1_1_required"] = item.Http_1_1_required
		in["err_sent_total"] = item.Err_sent_total
		in["err_sent_proto_err"] = item.Err_sent_proto_err
		in["err_sent_internal_err"] = item.Err_sent_internal_err
		in["err_sent_flow_control"] = item.Err_sent_flow_control
		in["err_sent_setting_timeout"] = item.Err_sent_setting_timeout
		in["err_sent_stream_closed"] = item.Err_sent_stream_closed
		in["err_sent_frame_size_err"] = item.Err_sent_frame_size_err
		in["err_sent_refused_stream"] = item.Err_sent_refused_stream
		in["err_sent_cancel"] = item.Err_sent_cancel
		in["err_sent_compression_err"] = item.Err_sent_compression_err
		in["err_sent_connect_err"] = item.Err_sent_connect_err
		in["err_sent_your_calm"] = item.Err_sent_your_calm
		in["err_sent_inadequate_security"] = item.Err_sent_inadequate_security
		in["err_sent_http11_required"] = item.Err_sent_http11_required
		in["frame_header_bytes_received"] = item.Frame_header_bytes_received
		in["frame_header_bytes_sent"] = item.Frame_header_bytes_sent
		in["control_bytes_received"] = item.Control_bytes_received
		in["control_bytes_sent"] = item.Control_bytes_sent
		in["header_bytes_received"] = item.Header_bytes_received
		in["header_bytes_sent"] = item.Header_bytes_sent
		in["data_bytes_received"] = item.Data_bytes_received
		in["data_bytes_sent"] = item.Data_bytes_sent
		in["total_bytes_received"] = item.Total_bytes_received
		in["total_bytes_sent"] = item.Total_bytes_sent
		in["push_promise_frame_sent"] = item.Push_promise_frame_sent
		in["http2_rejected"] = item.Http2_rejected
		in["current_stream"] = item.Current_stream
		in["stream_create"] = item.Stream_create
		in["stream_free"] = item.Stream_free
		in["end_stream_rcvd"] = item.End_stream_rcvd
		in["end_stream_sent"] = item.End_stream_sent
		result = append(result, in)
	}
	return result
}

func getObjectSlbHttp2OperOper(d []interface{}) edpt.SlbHttp2OperOper {

	count1 := len(d)
	var ret edpt.SlbHttp2OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Http2CpuList = getSliceSlbHttp2OperOperHttp2CpuList(in["http2_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbHttp2OperOperHttp2CpuList(d []interface{}) []edpt.SlbHttp2OperOperHttp2CpuList {

	count1 := len(d)
	ret := make([]edpt.SlbHttp2OperOperHttp2CpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHttp2OperOperHttp2CpuList
		oi.Curr_proxy = in["curr_proxy"].(int)
		oi.Peak_proxy = in["peak_proxy"].(int)
		oi.Total_proxy = in["total_proxy"].(int)
		oi.Connection_preface_rcvd = in["connection_preface_rcvd"].(int)
		oi.Connection_preface_sent = in["connection_preface_sent"].(int)
		oi.Control_frame = in["control_frame"].(int)
		oi.Control_frame_sent = in["control_frame_sent"].(int)
		oi.Continuation_frame = in["continuation_frame"].(int)
		oi.Continuation_frame_sent = in["continuation_frame_sent"].(int)
		oi.Data_frame = in["data_frame"].(int)
		oi.Data_frame_sent = in["data_frame_sent"].(int)
		oi.Data_to_app = in["data_to_app"].(int)
		oi.Goaway_frame = in["goaway_frame"].(int)
		oi.Goaway_frame_sent = in["goaway_frame_sent"].(int)
		oi.Headers_frame = in["headers_frame"].(int)
		oi.Headers_frame_sent = in["headers_frame_sent"].(int)
		oi.Header_to_app = in["header_to_app"].(int)
		oi.Ping_frame = in["ping_frame"].(int)
		oi.Ping_frame_sent = in["ping_frame_sent"].(int)
		oi.Priority_frame = in["priority_frame"].(int)
		oi.Priority_frame_sent = in["priority_frame_sent"].(int)
		oi.Rst_frame_rcvd = in["rst_frame_rcvd"].(int)
		oi.Rst_frame_sent = in["rst_frame_sent"].(int)
		oi.Settings_frame = in["settings_frame"].(int)
		oi.Settings_frame_sent = in["settings_frame_sent"].(int)
		oi.Settings_ack_rcvd = in["settings_ack_rcvd"].(int)
		oi.Settings_ack_sent = in["settings_ack_sent"].(int)
		oi.Empty_settings_rcvd = in["empty_settings_rcvd"].(int)
		oi.Empty_settings_sent = in["empty_settings_sent"].(int)
		oi.Window_update_frame = in["window_update_frame"].(int)
		oi.Window_update_frame_sent = in["window_update_frame_sent"].(int)
		oi.Unknown_frame = in["unknown_frame"].(int)
		oi.Split_buff_fail = in["split_buff_fail"].(int)
		oi.Invalid_frame_size = in["invalid_frame_size"].(int)
		oi.Error_max_invalid_stream = in["error_max_invalid_stream"].(int)
		oi.Data_no_stream = in["data_no_stream"].(int)
		oi.Bad_connection_preface = in["bad_connection_preface"].(int)
		oi.Bad_frame_type_for_stream_state = in["bad_frame_type_for_stream_state"].(int)
		oi.Wrong_stream_state = in["wrong_stream_state"].(int)
		oi.Alloc_fail_total = in["alloc_fail_total"].(int)
		oi.Proxy_alloc_error = in["proxy_alloc_error"].(int)
		oi.Deflate_alloc_fail = in["deflate_alloc_fail"].(int)
		oi.Inflate_alloc_fail = in["inflate_alloc_fail"].(int)
		oi.Data_queue_alloc_error = in["data_queue_alloc_error"].(int)
		oi.Buff_alloc_error = in["buff_alloc_error"].(int)
		oi.Cant_allocate_control_frame = in["cant_allocate_control_frame"].(int)
		oi.Cant_allocate_settings_frame = in["cant_allocate_settings_frame"].(int)
		oi.Cant_allocate_rst_frame = in["cant_allocate_rst_frame"].(int)
		oi.Cant_allocate_goaway_frame = in["cant_allocate_goaway_frame"].(int)
		oi.Cant_allocate_ping_frame = in["cant_allocate_ping_frame"].(int)
		oi.Cant_allocate_stream = in["cant_allocate_stream"].(int)
		oi.Cant_allocate_window_frame = in["cant_allocate_window_frame"].(int)
		oi.Inflate_header_fail = in["inflate_header_fail"].(int)
		oi.Header_no_stream = in["header_no_stream"].(int)
		oi.Header_padlen_gt_frame_payload = in["header_padlen_gt_frame_payload"].(int)
		oi.Streams_gt_max_concur_streams = in["streams_gt_max_concur_streams"].(int)
		oi.Idle_state_unexpected_frame = in["idle_state_unexpected_frame"].(int)
		oi.Reserved_local_state_unexpected_frame = in["reserved_local_state_unexpected_frame"].(int)
		oi.Reserved_remote_state_unexpected_frame = in["reserved_remote_state_unexpected_frame"].(int)
		oi.Half_closed_remote_state_unexpected_frame = in["half_closed_remote_state_unexpected_frame"].(int)
		oi.Closed_state_unexpected_frame = in["closed_state_unexpected_frame"].(int)
		oi.Zero_window_size_on_stream = in["zero_window_size_on_stream"].(int)
		oi.Exceeds_max_window_size_stream = in["exceeds_max_window_size_stream"].(int)
		oi.Continuation_before_headers = in["continuation_before_headers"].(int)
		oi.Invalid_frame_during_headers = in["invalid_frame_during_headers"].(int)
		oi.Headers_after_continuation = in["headers_after_continuation"].(int)
		oi.Invalid_push_promise = in["invalid_push_promise"].(int)
		oi.Invalid_stream_id = in["invalid_stream_id"].(int)
		oi.Headers_interleaved = in["headers_interleaved"].(int)
		oi.Trailers_no_end_stream = in["trailers_no_end_stream"].(int)
		oi.Invalid_setting_value = in["invalid_setting_value"].(int)
		oi.Invalid_window_update = in["invalid_window_update"].(int)
		oi.Err_rcvd_total = in["err_rcvd_total"].(int)
		oi.Protocol_error = in["protocol_error"].(int)
		oi.Internal_error = in["internal_error"].(int)
		oi.Flow_control_error = in["flow_control_error"].(int)
		oi.Settings_timeout = in["settings_timeout"].(int)
		oi.Stream_closed = in["stream_closed"].(int)
		oi.Frame_size_error = in["frame_size_error"].(int)
		oi.Refused_stream = in["refused_stream"].(int)
		oi.Cancel = in["cancel"].(int)
		oi.Compression_error = in["compression_error"].(int)
		oi.Connect_error = in["connect_error"].(int)
		oi.Enhance_your_calm = in["enhance_your_calm"].(int)
		oi.Inadequate_security = in["inadequate_security"].(int)
		oi.Http_1_1_required = in["http_1_1_required"].(int)
		oi.Err_sent_total = in["err_sent_total"].(int)
		oi.Err_sent_proto_err = in["err_sent_proto_err"].(int)
		oi.Err_sent_internal_err = in["err_sent_internal_err"].(int)
		oi.Err_sent_flow_control = in["err_sent_flow_control"].(int)
		oi.Err_sent_setting_timeout = in["err_sent_setting_timeout"].(int)
		oi.Err_sent_stream_closed = in["err_sent_stream_closed"].(int)
		oi.Err_sent_frame_size_err = in["err_sent_frame_size_err"].(int)
		oi.Err_sent_refused_stream = in["err_sent_refused_stream"].(int)
		oi.Err_sent_cancel = in["err_sent_cancel"].(int)
		oi.Err_sent_compression_err = in["err_sent_compression_err"].(int)
		oi.Err_sent_connect_err = in["err_sent_connect_err"].(int)
		oi.Err_sent_your_calm = in["err_sent_your_calm"].(int)
		oi.Err_sent_inadequate_security = in["err_sent_inadequate_security"].(int)
		oi.Err_sent_http11_required = in["err_sent_http11_required"].(int)
		oi.Frame_header_bytes_received = in["frame_header_bytes_received"].(int)
		oi.Frame_header_bytes_sent = in["frame_header_bytes_sent"].(int)
		oi.Control_bytes_received = in["control_bytes_received"].(int)
		oi.Control_bytes_sent = in["control_bytes_sent"].(int)
		oi.Header_bytes_received = in["header_bytes_received"].(int)
		oi.Header_bytes_sent = in["header_bytes_sent"].(int)
		oi.Data_bytes_received = in["data_bytes_received"].(int)
		oi.Data_bytes_sent = in["data_bytes_sent"].(int)
		oi.Total_bytes_received = in["total_bytes_received"].(int)
		oi.Total_bytes_sent = in["total_bytes_sent"].(int)
		oi.Push_promise_frame_sent = in["push_promise_frame_sent"].(int)
		oi.Http2_rejected = in["http2_rejected"].(int)
		oi.Current_stream = in["current_stream"].(int)
		oi.Stream_create = in["stream_create"].(int)
		oi.Stream_free = in["stream_free"].(int)
		oi.End_stream_rcvd = in["end_stream_rcvd"].(int)
		oi.End_stream_sent = in["end_stream_sent"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHttp2Oper(d *schema.ResourceData) edpt.SlbHttp2Oper {
	var ret edpt.SlbHttp2Oper

	ret.Oper = getObjectSlbHttp2OperOper(d.Get("oper").([]interface{}))
	return ret
}
