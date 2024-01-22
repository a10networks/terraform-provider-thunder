package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttp2Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_http2_stats`: Statistics for the object http2\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHttp2StatsRead,

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
						"connection_preface_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Connection preface rcvd",
						},
						"control_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Control Frame Rcvd",
						},
						"headers_frame": {
							Type: schema.TypeInt, Optional: true, Description: "HEADERS Frame Rcvd",
						},
						"continuation_frame": {
							Type: schema.TypeInt, Optional: true, Description: "CONTINUATION Frame Rcvd",
						},
						"rst_frame_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "RST_STREAM Frame Rcvd",
						},
						"settings_frame": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS Frame Rcvd",
						},
						"window_update_frame": {
							Type: schema.TypeInt, Optional: true, Description: "WINDOW_UPDATE Frame Rcvd",
						},
						"ping_frame": {
							Type: schema.TypeInt, Optional: true, Description: "PING Frame Rcvd",
						},
						"goaway_frame": {
							Type: schema.TypeInt, Optional: true, Description: "GOAWAY Frame Rcvd",
						},
						"priority_frame": {
							Type: schema.TypeInt, Optional: true, Description: "PRIORITY Frame Rcvd",
						},
						"data_frame": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame Recvd",
						},
						"unknown_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Frame Recvd",
						},
						"connection_preface_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Connection preface sent",
						},
						"settings_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS Frame Sent",
						},
						"settings_ack_sent": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS ACK Frame Sent",
						},
						"empty_settings_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Empty SETTINGS Frame Sent",
						},
						"ping_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "PING Frame Sent",
						},
						"window_update_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "WINDOW_UPDATE Frame Sent",
						},
						"rst_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "RST_STREAM Frame Sent",
						},
						"goaway_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "GOAWAY Frame Sent",
						},
						"header_to_app": {
							Type: schema.TypeInt, Optional: true, Description: "HEADER Frame to HTTP",
						},
						"data_to_app": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame to HTTP",
						},
						"protocol_error": {
							Type: schema.TypeInt, Optional: true, Description: "Protocol Error",
						},
						"internal_error": {
							Type: schema.TypeInt, Optional: true, Description: "Internal Error",
						},
						"proxy_alloc_error": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP2 Proxy alloc Error",
						},
						"split_buff_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Splitting Buffer Failed",
						},
						"invalid_frame_size": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Frame Size Rcvd",
						},
						"error_max_invalid_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Max Invalid Stream Rcvd",
						},
						"data_no_stream": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame Rcvd on non-existent stream",
						},
						"flow_control_error": {
							Type: schema.TypeInt, Optional: true, Description: "Flow Control Error",
						},
						"settings_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Settings Timeout",
						},
						"frame_size_error": {
							Type: schema.TypeInt, Optional: true, Description: "Frame Size Error",
						},
						"refused_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Refused Stream",
						},
						"cancel": {
							Type: schema.TypeInt, Optional: true, Description: "cancel",
						},
						"compression_error": {
							Type: schema.TypeInt, Optional: true, Description: "compression error",
						},
						"connect_error": {
							Type: schema.TypeInt, Optional: true, Description: "connect error",
						},
						"enhance_your_calm": {
							Type: schema.TypeInt, Optional: true, Description: "enhance your calm error",
						},
						"inadequate_security": {
							Type: schema.TypeInt, Optional: true, Description: "inadequate security",
						},
						"http_1_1_required": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP1.1 Required",
						},
						"deflate_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "deflate alloc fail",
						},
						"inflate_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "inflate alloc fail",
						},
						"inflate_header_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Inflate Header Fail",
						},
						"bad_connection_preface": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Connection Preface",
						},
						"cant_allocate_control_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate control frame",
						},
						"cant_allocate_settings_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate SETTINGS frame",
						},
						"bad_frame_type_for_stream_state": {
							Type: schema.TypeInt, Optional: true, Description: "Bad frame type for stream state",
						},
						"wrong_stream_state": {
							Type: schema.TypeInt, Optional: true, Description: "Wrong Stream State",
						},
						"data_queue_alloc_error": {
							Type: schema.TypeInt, Optional: true, Description: "Data Queue Alloc Error",
						},
						"buff_alloc_error": {
							Type: schema.TypeInt, Optional: true, Description: "Buff alloc error",
						},
						"cant_allocate_rst_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate RST_STREAM frame",
						},
						"cant_allocate_goaway_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate GOAWAY frame",
						},
						"cant_allocate_ping_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate PING frame",
						},
						"cant_allocate_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate stream",
						},
						"cant_allocate_window_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Cant allocate WINDOW_UPDATE frame",
						},
						"header_no_stream": {
							Type: schema.TypeInt, Optional: true, Description: "header no stream",
						},
						"header_padlen_gt_frame_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Header padlen greater than frame payload size",
						},
						"streams_gt_max_concur_streams": {
							Type: schema.TypeInt, Optional: true, Description: "Streams greater than max allowed concurrent streams",
						},
						"idle_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unxpected frame received in idle state",
						},
						"reserved_local_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unexpected frame received in reserved local state",
						},
						"reserved_remote_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unexpected frame received in reserved remote state",
						},
						"half_closed_remote_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unexpected frame received in half closed remote state",
						},
						"closed_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Description: "Unexpected frame received in closed state",
						},
						"zero_window_size_on_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Window Update with zero increment rcvd",
						},
						"exceeds_max_window_size_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Window Update with increment that results in exceeding max window",
						},
						"stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "stream closed",
						},
						"continuation_before_headers": {
							Type: schema.TypeInt, Optional: true, Description: "CONTINUATION frame with no headers frame",
						},
						"invalid_frame_during_headers": {
							Type: schema.TypeInt, Optional: true, Description: "frame before headers were complete",
						},
						"headers_after_continuation": {
							Type: schema.TypeInt, Optional: true, Description: "headers frame before CONTINUATION was complete",
						},
						"push_promise_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Push Promise Frame Sent",
						},
						"invalid_push_promise": {
							Type: schema.TypeInt, Optional: true, Description: "unexpected PUSH_PROMISE frame",
						},
						"invalid_stream_id": {
							Type: schema.TypeInt, Optional: true, Description: "received invalid stream ID",
						},
						"headers_interleaved": {
							Type: schema.TypeInt, Optional: true, Description: "headers interleaved on streams",
						},
						"trailers_no_end_stream": {
							Type: schema.TypeInt, Optional: true, Description: "trailers not marked as end-of-stream",
						},
						"invalid_setting_value": {
							Type: schema.TypeInt, Optional: true, Description: "invalid setting-frame value",
						},
						"invalid_window_update": {
							Type: schema.TypeInt, Optional: true, Description: "window-update value out of range",
						},
						"frame_header_bytes_received": {
							Type: schema.TypeInt, Optional: true, Description: "frame header bytes received",
						},
						"frame_header_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "frame header bytes sent",
						},
						"control_bytes_received": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 control frame bytes received",
						},
						"control_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 control frame bytes sent",
						},
						"header_bytes_received": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 header bytes received",
						},
						"header_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 header bytes sent",
						},
						"data_bytes_received": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 data bytes received",
						},
						"data_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 data bytes sent",
						},
						"total_bytes_received": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 total bytes received",
						},
						"total_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/2 total bytes sent",
						},
						"peak_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Peak Proxy Conns",
						},
						"control_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Control Frame Sent",
						},
						"continuation_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "CONTINUATION Frame Sent",
						},
						"data_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame Sent",
						},
						"headers_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HEADERS Frame Sent",
						},
						"priority_frame_sent": {
							Type: schema.TypeInt, Optional: true, Description: "PRIORITY Frame Sent",
						},
						"settings_ack_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS ACK Frame Rcvd",
						},
						"empty_settings_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Empty SETTINGS Frame Rcvd",
						},
						"alloc_fail_total": {
							Type: schema.TypeInt, Optional: true, Description: "Alloc Fail - Total",
						},
						"err_rcvd_total": {
							Type: schema.TypeInt, Optional: true, Description: "Error Rcvd - Total",
						},
						"err_sent_total": {
							Type: schema.TypeInt, Optional: true, Description: "Error Rent - Total",
						},
						"err_sent_proto_err": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - PROTOCOL_ERROR",
						},
						"err_sent_internal_err": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - INTERNAL_ERROR",
						},
						"err_sent_flow_control": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - FLOW_CONTROL_ERROR",
						},
						"err_sent_setting_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - SETTINGS_TIMEOUT",
						},
						"err_sent_stream_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - STREAM_CLOSED",
						},
						"err_sent_frame_size_err": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - FRAME_SIZE_ERROR",
						},
						"err_sent_refused_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - REFUSED_STREAM",
						},
						"err_sent_cancel": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - CANCEL",
						},
						"err_sent_compression_err": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - COMPRESSION_ERROR",
						},
						"err_sent_connect_err": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - CONNECT_ERROR",
						},
						"err_sent_your_calm": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - ENHANCE_YOUR_CALM",
						},
						"err_sent_inadequate_security": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - INADEQUATE_SECURITY",
						},
						"err_sent_http11_required": {
							Type: schema.TypeInt, Optional: true, Description: "Error Sent - HTTP_1_1_REQUIRED",
						},
						"http2_rejected": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP2 Rejected",
						},
						"current_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Current Streams",
						},
						"stream_create": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Create",
						},
						"stream_free": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Free",
						},
						"end_stream_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "End Stream Recieved",
						},
						"end_stream_sent": {
							Type: schema.TypeInt, Optional: true, Description: "End Stream Sent",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHttp2StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp2StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp2Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHttp2StatsStats := setObjectSlbHttp2StatsStats(res)
		d.Set("stats", SlbHttp2StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHttp2StatsStats(ret edpt.DataSlbHttp2Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_proxy":                                ret.DtSlbHttp2Stats.Stats.Curr_proxy,
			"total_proxy":                               ret.DtSlbHttp2Stats.Stats.Total_proxy,
			"connection_preface_rcvd":                   ret.DtSlbHttp2Stats.Stats.Connection_preface_rcvd,
			"control_frame":                             ret.DtSlbHttp2Stats.Stats.Control_frame,
			"headers_frame":                             ret.DtSlbHttp2Stats.Stats.Headers_frame,
			"continuation_frame":                        ret.DtSlbHttp2Stats.Stats.Continuation_frame,
			"rst_frame_rcvd":                            ret.DtSlbHttp2Stats.Stats.Rst_frame_rcvd,
			"settings_frame":                            ret.DtSlbHttp2Stats.Stats.Settings_frame,
			"window_update_frame":                       ret.DtSlbHttp2Stats.Stats.Window_update_frame,
			"ping_frame":                                ret.DtSlbHttp2Stats.Stats.Ping_frame,
			"goaway_frame":                              ret.DtSlbHttp2Stats.Stats.Goaway_frame,
			"priority_frame":                            ret.DtSlbHttp2Stats.Stats.Priority_frame,
			"data_frame":                                ret.DtSlbHttp2Stats.Stats.Data_frame,
			"unknown_frame":                             ret.DtSlbHttp2Stats.Stats.Unknown_frame,
			"connection_preface_sent":                   ret.DtSlbHttp2Stats.Stats.Connection_preface_sent,
			"settings_frame_sent":                       ret.DtSlbHttp2Stats.Stats.Settings_frame_sent,
			"settings_ack_sent":                         ret.DtSlbHttp2Stats.Stats.Settings_ack_sent,
			"empty_settings_sent":                       ret.DtSlbHttp2Stats.Stats.Empty_settings_sent,
			"ping_frame_sent":                           ret.DtSlbHttp2Stats.Stats.Ping_frame_sent,
			"window_update_frame_sent":                  ret.DtSlbHttp2Stats.Stats.Window_update_frame_sent,
			"rst_frame_sent":                            ret.DtSlbHttp2Stats.Stats.Rst_frame_sent,
			"goaway_frame_sent":                         ret.DtSlbHttp2Stats.Stats.Goaway_frame_sent,
			"header_to_app":                             ret.DtSlbHttp2Stats.Stats.Header_to_app,
			"data_to_app":                               ret.DtSlbHttp2Stats.Stats.Data_to_app,
			"protocol_error":                            ret.DtSlbHttp2Stats.Stats.Protocol_error,
			"internal_error":                            ret.DtSlbHttp2Stats.Stats.Internal_error,
			"proxy_alloc_error":                         ret.DtSlbHttp2Stats.Stats.Proxy_alloc_error,
			"split_buff_fail":                           ret.DtSlbHttp2Stats.Stats.Split_buff_fail,
			"invalid_frame_size":                        ret.DtSlbHttp2Stats.Stats.Invalid_frame_size,
			"error_max_invalid_stream":                  ret.DtSlbHttp2Stats.Stats.Error_max_invalid_stream,
			"data_no_stream":                            ret.DtSlbHttp2Stats.Stats.Data_no_stream,
			"flow_control_error":                        ret.DtSlbHttp2Stats.Stats.Flow_control_error,
			"settings_timeout":                          ret.DtSlbHttp2Stats.Stats.Settings_timeout,
			"frame_size_error":                          ret.DtSlbHttp2Stats.Stats.Frame_size_error,
			"refused_stream":                            ret.DtSlbHttp2Stats.Stats.Refused_stream,
			"cancel":                                    ret.DtSlbHttp2Stats.Stats.Cancel,
			"compression_error":                         ret.DtSlbHttp2Stats.Stats.Compression_error,
			"connect_error":                             ret.DtSlbHttp2Stats.Stats.Connect_error,
			"enhance_your_calm":                         ret.DtSlbHttp2Stats.Stats.Enhance_your_calm,
			"inadequate_security":                       ret.DtSlbHttp2Stats.Stats.Inadequate_security,
			"http_1_1_required":                         ret.DtSlbHttp2Stats.Stats.Http_1_1_required,
			"deflate_alloc_fail":                        ret.DtSlbHttp2Stats.Stats.Deflate_alloc_fail,
			"inflate_alloc_fail":                        ret.DtSlbHttp2Stats.Stats.Inflate_alloc_fail,
			"inflate_header_fail":                       ret.DtSlbHttp2Stats.Stats.Inflate_header_fail,
			"bad_connection_preface":                    ret.DtSlbHttp2Stats.Stats.Bad_connection_preface,
			"cant_allocate_control_frame":               ret.DtSlbHttp2Stats.Stats.Cant_allocate_control_frame,
			"cant_allocate_settings_frame":              ret.DtSlbHttp2Stats.Stats.Cant_allocate_settings_frame,
			"bad_frame_type_for_stream_state":           ret.DtSlbHttp2Stats.Stats.Bad_frame_type_for_stream_state,
			"wrong_stream_state":                        ret.DtSlbHttp2Stats.Stats.Wrong_stream_state,
			"data_queue_alloc_error":                    ret.DtSlbHttp2Stats.Stats.Data_queue_alloc_error,
			"buff_alloc_error":                          ret.DtSlbHttp2Stats.Stats.Buff_alloc_error,
			"cant_allocate_rst_frame":                   ret.DtSlbHttp2Stats.Stats.Cant_allocate_rst_frame,
			"cant_allocate_goaway_frame":                ret.DtSlbHttp2Stats.Stats.Cant_allocate_goaway_frame,
			"cant_allocate_ping_frame":                  ret.DtSlbHttp2Stats.Stats.Cant_allocate_ping_frame,
			"cant_allocate_stream":                      ret.DtSlbHttp2Stats.Stats.Cant_allocate_stream,
			"cant_allocate_window_frame":                ret.DtSlbHttp2Stats.Stats.Cant_allocate_window_frame,
			"header_no_stream":                          ret.DtSlbHttp2Stats.Stats.Header_no_stream,
			"header_padlen_gt_frame_payload":            ret.DtSlbHttp2Stats.Stats.Header_padlen_gt_frame_payload,
			"streams_gt_max_concur_streams":             ret.DtSlbHttp2Stats.Stats.Streams_gt_max_concur_streams,
			"idle_state_unexpected_frame":               ret.DtSlbHttp2Stats.Stats.Idle_state_unexpected_frame,
			"reserved_local_state_unexpected_frame":     ret.DtSlbHttp2Stats.Stats.Reserved_local_state_unexpected_frame,
			"reserved_remote_state_unexpected_frame":    ret.DtSlbHttp2Stats.Stats.Reserved_remote_state_unexpected_frame,
			"half_closed_remote_state_unexpected_frame": ret.DtSlbHttp2Stats.Stats.Half_closed_remote_state_unexpected_frame,
			"closed_state_unexpected_frame":             ret.DtSlbHttp2Stats.Stats.Closed_state_unexpected_frame,
			"zero_window_size_on_stream":                ret.DtSlbHttp2Stats.Stats.Zero_window_size_on_stream,
			"exceeds_max_window_size_stream":            ret.DtSlbHttp2Stats.Stats.Exceeds_max_window_size_stream,
			"stream_closed":                             ret.DtSlbHttp2Stats.Stats.Stream_closed,
			"continuation_before_headers":               ret.DtSlbHttp2Stats.Stats.Continuation_before_headers,
			"invalid_frame_during_headers":              ret.DtSlbHttp2Stats.Stats.Invalid_frame_during_headers,
			"headers_after_continuation":                ret.DtSlbHttp2Stats.Stats.Headers_after_continuation,
			"push_promise_frame_sent":                   ret.DtSlbHttp2Stats.Stats.Push_promise_frame_sent,
			"invalid_push_promise":                      ret.DtSlbHttp2Stats.Stats.Invalid_push_promise,
			"invalid_stream_id":                         ret.DtSlbHttp2Stats.Stats.Invalid_stream_id,
			"headers_interleaved":                       ret.DtSlbHttp2Stats.Stats.Headers_interleaved,
			"trailers_no_end_stream":                    ret.DtSlbHttp2Stats.Stats.Trailers_no_end_stream,
			"invalid_setting_value":                     ret.DtSlbHttp2Stats.Stats.Invalid_setting_value,
			"invalid_window_update":                     ret.DtSlbHttp2Stats.Stats.Invalid_window_update,
			"frame_header_bytes_received":               ret.DtSlbHttp2Stats.Stats.Frame_header_bytes_received,
			"frame_header_bytes_sent":                   ret.DtSlbHttp2Stats.Stats.Frame_header_bytes_sent,
			"control_bytes_received":                    ret.DtSlbHttp2Stats.Stats.Control_bytes_received,
			"control_bytes_sent":                        ret.DtSlbHttp2Stats.Stats.Control_bytes_sent,
			"header_bytes_received":                     ret.DtSlbHttp2Stats.Stats.Header_bytes_received,
			"header_bytes_sent":                         ret.DtSlbHttp2Stats.Stats.Header_bytes_sent,
			"data_bytes_received":                       ret.DtSlbHttp2Stats.Stats.Data_bytes_received,
			"data_bytes_sent":                           ret.DtSlbHttp2Stats.Stats.Data_bytes_sent,
			"total_bytes_received":                      ret.DtSlbHttp2Stats.Stats.Total_bytes_received,
			"total_bytes_sent":                          ret.DtSlbHttp2Stats.Stats.Total_bytes_sent,
			"peak_proxy":                                ret.DtSlbHttp2Stats.Stats.Peak_proxy,
			"control_frame_sent":                        ret.DtSlbHttp2Stats.Stats.Control_frame_sent,
			"continuation_frame_sent":                   ret.DtSlbHttp2Stats.Stats.Continuation_frame_sent,
			"data_frame_sent":                           ret.DtSlbHttp2Stats.Stats.Data_frame_sent,
			"headers_frame_sent":                        ret.DtSlbHttp2Stats.Stats.Headers_frame_sent,
			"priority_frame_sent":                       ret.DtSlbHttp2Stats.Stats.Priority_frame_sent,
			"settings_ack_rcvd":                         ret.DtSlbHttp2Stats.Stats.Settings_ack_rcvd,
			"empty_settings_rcvd":                       ret.DtSlbHttp2Stats.Stats.Empty_settings_rcvd,
			"alloc_fail_total":                          ret.DtSlbHttp2Stats.Stats.Alloc_fail_total,
			"err_rcvd_total":                            ret.DtSlbHttp2Stats.Stats.Err_rcvd_total,
			"err_sent_total":                            ret.DtSlbHttp2Stats.Stats.Err_sent_total,
			"err_sent_proto_err":                        ret.DtSlbHttp2Stats.Stats.Err_sent_proto_err,
			"err_sent_internal_err":                     ret.DtSlbHttp2Stats.Stats.Err_sent_internal_err,
			"err_sent_flow_control":                     ret.DtSlbHttp2Stats.Stats.Err_sent_flow_control,
			"err_sent_setting_timeout":                  ret.DtSlbHttp2Stats.Stats.Err_sent_setting_timeout,
			"err_sent_stream_closed":                    ret.DtSlbHttp2Stats.Stats.Err_sent_stream_closed,
			"err_sent_frame_size_err":                   ret.DtSlbHttp2Stats.Stats.Err_sent_frame_size_err,
			"err_sent_refused_stream":                   ret.DtSlbHttp2Stats.Stats.Err_sent_refused_stream,
			"err_sent_cancel":                           ret.DtSlbHttp2Stats.Stats.Err_sent_cancel,
			"err_sent_compression_err":                  ret.DtSlbHttp2Stats.Stats.Err_sent_compression_err,
			"err_sent_connect_err":                      ret.DtSlbHttp2Stats.Stats.Err_sent_connect_err,
			"err_sent_your_calm":                        ret.DtSlbHttp2Stats.Stats.Err_sent_your_calm,
			"err_sent_inadequate_security":              ret.DtSlbHttp2Stats.Stats.Err_sent_inadequate_security,
			"err_sent_http11_required":                  ret.DtSlbHttp2Stats.Stats.Err_sent_http11_required,
			"http2_rejected":                            ret.DtSlbHttp2Stats.Stats.Http2_rejected,
			"current_stream":                            ret.DtSlbHttp2Stats.Stats.Current_stream,
			"stream_create":                             ret.DtSlbHttp2Stats.Stats.Stream_create,
			"stream_free":                               ret.DtSlbHttp2Stats.Stats.Stream_free,
			"end_stream_rcvd":                           ret.DtSlbHttp2Stats.Stats.End_stream_rcvd,
			"end_stream_sent":                           ret.DtSlbHttp2Stats.Stats.End_stream_sent,
		},
	}
}

func getObjectSlbHttp2StatsStats(d []interface{}) edpt.SlbHttp2StatsStats {

	count1 := len(d)
	var ret edpt.SlbHttp2StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_proxy = in["curr_proxy"].(int)
		ret.Total_proxy = in["total_proxy"].(int)
		ret.Connection_preface_rcvd = in["connection_preface_rcvd"].(int)
		ret.Control_frame = in["control_frame"].(int)
		ret.Headers_frame = in["headers_frame"].(int)
		ret.Continuation_frame = in["continuation_frame"].(int)
		ret.Rst_frame_rcvd = in["rst_frame_rcvd"].(int)
		ret.Settings_frame = in["settings_frame"].(int)
		ret.Window_update_frame = in["window_update_frame"].(int)
		ret.Ping_frame = in["ping_frame"].(int)
		ret.Goaway_frame = in["goaway_frame"].(int)
		ret.Priority_frame = in["priority_frame"].(int)
		ret.Data_frame = in["data_frame"].(int)
		ret.Unknown_frame = in["unknown_frame"].(int)
		ret.Connection_preface_sent = in["connection_preface_sent"].(int)
		ret.Settings_frame_sent = in["settings_frame_sent"].(int)
		ret.Settings_ack_sent = in["settings_ack_sent"].(int)
		ret.Empty_settings_sent = in["empty_settings_sent"].(int)
		ret.Ping_frame_sent = in["ping_frame_sent"].(int)
		ret.Window_update_frame_sent = in["window_update_frame_sent"].(int)
		ret.Rst_frame_sent = in["rst_frame_sent"].(int)
		ret.Goaway_frame_sent = in["goaway_frame_sent"].(int)
		ret.Header_to_app = in["header_to_app"].(int)
		ret.Data_to_app = in["data_to_app"].(int)
		ret.Protocol_error = in["protocol_error"].(int)
		ret.Internal_error = in["internal_error"].(int)
		ret.Proxy_alloc_error = in["proxy_alloc_error"].(int)
		ret.Split_buff_fail = in["split_buff_fail"].(int)
		ret.Invalid_frame_size = in["invalid_frame_size"].(int)
		ret.Error_max_invalid_stream = in["error_max_invalid_stream"].(int)
		ret.Data_no_stream = in["data_no_stream"].(int)
		ret.Flow_control_error = in["flow_control_error"].(int)
		ret.Settings_timeout = in["settings_timeout"].(int)
		ret.Frame_size_error = in["frame_size_error"].(int)
		ret.Refused_stream = in["refused_stream"].(int)
		ret.Cancel = in["cancel"].(int)
		ret.Compression_error = in["compression_error"].(int)
		ret.Connect_error = in["connect_error"].(int)
		ret.Enhance_your_calm = in["enhance_your_calm"].(int)
		ret.Inadequate_security = in["inadequate_security"].(int)
		ret.Http_1_1_required = in["http_1_1_required"].(int)
		ret.Deflate_alloc_fail = in["deflate_alloc_fail"].(int)
		ret.Inflate_alloc_fail = in["inflate_alloc_fail"].(int)
		ret.Inflate_header_fail = in["inflate_header_fail"].(int)
		ret.Bad_connection_preface = in["bad_connection_preface"].(int)
		ret.Cant_allocate_control_frame = in["cant_allocate_control_frame"].(int)
		ret.Cant_allocate_settings_frame = in["cant_allocate_settings_frame"].(int)
		ret.Bad_frame_type_for_stream_state = in["bad_frame_type_for_stream_state"].(int)
		ret.Wrong_stream_state = in["wrong_stream_state"].(int)
		ret.Data_queue_alloc_error = in["data_queue_alloc_error"].(int)
		ret.Buff_alloc_error = in["buff_alloc_error"].(int)
		ret.Cant_allocate_rst_frame = in["cant_allocate_rst_frame"].(int)
		ret.Cant_allocate_goaway_frame = in["cant_allocate_goaway_frame"].(int)
		ret.Cant_allocate_ping_frame = in["cant_allocate_ping_frame"].(int)
		ret.Cant_allocate_stream = in["cant_allocate_stream"].(int)
		ret.Cant_allocate_window_frame = in["cant_allocate_window_frame"].(int)
		ret.Header_no_stream = in["header_no_stream"].(int)
		ret.Header_padlen_gt_frame_payload = in["header_padlen_gt_frame_payload"].(int)
		ret.Streams_gt_max_concur_streams = in["streams_gt_max_concur_streams"].(int)
		ret.Idle_state_unexpected_frame = in["idle_state_unexpected_frame"].(int)
		ret.Reserved_local_state_unexpected_frame = in["reserved_local_state_unexpected_frame"].(int)
		ret.Reserved_remote_state_unexpected_frame = in["reserved_remote_state_unexpected_frame"].(int)
		ret.Half_closed_remote_state_unexpected_frame = in["half_closed_remote_state_unexpected_frame"].(int)
		ret.Closed_state_unexpected_frame = in["closed_state_unexpected_frame"].(int)
		ret.Zero_window_size_on_stream = in["zero_window_size_on_stream"].(int)
		ret.Exceeds_max_window_size_stream = in["exceeds_max_window_size_stream"].(int)
		ret.Stream_closed = in["stream_closed"].(int)
		ret.Continuation_before_headers = in["continuation_before_headers"].(int)
		ret.Invalid_frame_during_headers = in["invalid_frame_during_headers"].(int)
		ret.Headers_after_continuation = in["headers_after_continuation"].(int)
		ret.Push_promise_frame_sent = in["push_promise_frame_sent"].(int)
		ret.Invalid_push_promise = in["invalid_push_promise"].(int)
		ret.Invalid_stream_id = in["invalid_stream_id"].(int)
		ret.Headers_interleaved = in["headers_interleaved"].(int)
		ret.Trailers_no_end_stream = in["trailers_no_end_stream"].(int)
		ret.Invalid_setting_value = in["invalid_setting_value"].(int)
		ret.Invalid_window_update = in["invalid_window_update"].(int)
		ret.Frame_header_bytes_received = in["frame_header_bytes_received"].(int)
		ret.Frame_header_bytes_sent = in["frame_header_bytes_sent"].(int)
		ret.Control_bytes_received = in["control_bytes_received"].(int)
		ret.Control_bytes_sent = in["control_bytes_sent"].(int)
		ret.Header_bytes_received = in["header_bytes_received"].(int)
		ret.Header_bytes_sent = in["header_bytes_sent"].(int)
		ret.Data_bytes_received = in["data_bytes_received"].(int)
		ret.Data_bytes_sent = in["data_bytes_sent"].(int)
		ret.Total_bytes_received = in["total_bytes_received"].(int)
		ret.Total_bytes_sent = in["total_bytes_sent"].(int)
		ret.Peak_proxy = in["peak_proxy"].(int)
		ret.Control_frame_sent = in["control_frame_sent"].(int)
		ret.Continuation_frame_sent = in["continuation_frame_sent"].(int)
		ret.Data_frame_sent = in["data_frame_sent"].(int)
		ret.Headers_frame_sent = in["headers_frame_sent"].(int)
		ret.Priority_frame_sent = in["priority_frame_sent"].(int)
		ret.Settings_ack_rcvd = in["settings_ack_rcvd"].(int)
		ret.Empty_settings_rcvd = in["empty_settings_rcvd"].(int)
		ret.Alloc_fail_total = in["alloc_fail_total"].(int)
		ret.Err_rcvd_total = in["err_rcvd_total"].(int)
		ret.Err_sent_total = in["err_sent_total"].(int)
		ret.Err_sent_proto_err = in["err_sent_proto_err"].(int)
		ret.Err_sent_internal_err = in["err_sent_internal_err"].(int)
		ret.Err_sent_flow_control = in["err_sent_flow_control"].(int)
		ret.Err_sent_setting_timeout = in["err_sent_setting_timeout"].(int)
		ret.Err_sent_stream_closed = in["err_sent_stream_closed"].(int)
		ret.Err_sent_frame_size_err = in["err_sent_frame_size_err"].(int)
		ret.Err_sent_refused_stream = in["err_sent_refused_stream"].(int)
		ret.Err_sent_cancel = in["err_sent_cancel"].(int)
		ret.Err_sent_compression_err = in["err_sent_compression_err"].(int)
		ret.Err_sent_connect_err = in["err_sent_connect_err"].(int)
		ret.Err_sent_your_calm = in["err_sent_your_calm"].(int)
		ret.Err_sent_inadequate_security = in["err_sent_inadequate_security"].(int)
		ret.Err_sent_http11_required = in["err_sent_http11_required"].(int)
		ret.Http2_rejected = in["http2_rejected"].(int)
		ret.Current_stream = in["current_stream"].(int)
		ret.Stream_create = in["stream_create"].(int)
		ret.Stream_free = in["stream_free"].(int)
		ret.End_stream_rcvd = in["end_stream_rcvd"].(int)
		ret.End_stream_sent = in["end_stream_sent"].(int)
	}
	return ret
}

func dataToEndpointSlbHttp2Stats(d *schema.ResourceData) edpt.SlbHttp2Stats {
	var ret edpt.SlbHttp2Stats

	ret.Stats = getObjectSlbHttp2StatsStats(d.Get("stats").([]interface{}))
	return ret
}
