package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttp3Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_http3_stats`: Statistics for the object http3\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbHttp3StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_conn_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current HTTP/3 Client Connections",
						},
						"server_conn_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current HTTP/3 Server Connections",
						},
						"client_conn_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP/3 Client Connections",
						},
						"server_conn_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total HTTP/3 Server Connections",
						},
						"client_conn_peak": {
							Type: schema.TypeInt, Optional: true, Description: "Peak HTTP/3 Client Connections",
						},
						"server_conn_peak": {
							Type: schema.TypeInt, Optional: true, Description: "Peak HTTP/3 Server Connections",
						},
						"client_request_streams_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current Request Streams on client side",
						},
						"server_request_streams_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current Request Streams on server side",
						},
						"client_request_streams_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request Streams on client side",
						},
						"server_request_streams_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Request Streams on server side",
						},
						"client_request_push_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current Push Streams on client side",
						},
						"server_request_push_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current Push Streams on server side",
						},
						"client_request_push_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Push Streams on client side",
						},
						"server_request_push_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Push Streams on server side",
						},
						"client_request_other_curr": {
							Type: schema.TypeInt, Optional: true, Description: "Current Other Streams on client side (control, decoder, encoder)",
						},
						"server_request_other_curr": {
							Type: schema.TypeInt, Optional: true, Description: "urrent Other Streams on server side (control, decoder, encoder)",
						},
						"client_request_other_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Streams on client side (control, decoder, encoder)",
						},
						"server_request_other_total": {
							Type: schema.TypeInt, Optional: true, Description: "Total Other Streams on server side (control, decoder, encoder)",
						},
						"client_frame_type_headers_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "HEADERS Frame received on client side",
						},
						"client_frame_type_headers_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HEADERS Frame sent on client side",
						},
						"client_frame_type_data_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame received on client side",
						},
						"client_frame_type_data_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame sent on client side",
						},
						"client_frame_type_cancel_push_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "CANCEL PUSH Frame received on client side",
						},
						"client_frame_type_cancel_push_sent": {
							Type: schema.TypeInt, Optional: true, Description: "CANCEL PUSH Frame sent on client side",
						},
						"client_frame_type_settings_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS Frame received on client side",
						},
						"client_frame_type_settings_sent": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS Frame sent on client side",
						},
						"client_frame_type_push_promise_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "PUSH PROMISE Frame received on client side",
						},
						"client_frame_type_push_promise_sent": {
							Type: schema.TypeInt, Optional: true, Description: "PUSH PROMISE Frame sent on client side",
						},
						"client_frame_type_goaway_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GOAWAY Frame received on client side",
						},
						"client_frame_type_goaway_sent": {
							Type: schema.TypeInt, Optional: true, Description: "GOAWAY Frame sent on client side",
						},
						"client_frame_type_max_push_id_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "MAX PUSH ID Frame received on client side",
						},
						"client_frame_type_max_push_id_sent": {
							Type: schema.TypeInt, Optional: true, Description: "MAX PUSH ID Frame sent on client side",
						},
						"client_frame_type_unknown_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Frame received on client side",
						},
						"client_header_frames_to_app": {
							Type: schema.TypeInt, Optional: true, Description: "HEADER Frames passed to HTTP layer on client side",
						},
						"client_data_frames_to_app": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frames passed to HTTP layer on client side",
						},
						"client_header_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in HEADER frames on client side",
						},
						"client_header_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in HEADER frames on client side",
						},
						"client_data_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in DATA frames on client side",
						},
						"client_data_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in DATA frames on client side",
						},
						"client_other_frame_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on client side",
						},
						"client_other_frame_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on client side",
						},
						"client_heading_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in HEADERS/DATA frame/stream heading on client side",
						},
						"client_heading_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in HEADERS/DATA frame/stream heading on client side",
						},
						"client_total_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bytes received on client side",
						},
						"client_total_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bytes sent on client side",
						},
						"server_frame_type_headers_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "HEADERS Frame received on server side",
						},
						"server_frame_type_headers_sent": {
							Type: schema.TypeInt, Optional: true, Description: "HEADERS Frame sent on server side",
						},
						"server_frame_type_data_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame received on server side",
						},
						"server_frame_type_data_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frame sent on server side",
						},
						"server_frame_type_cancel_push_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "CANCEL PUSH Frame received on server side",
						},
						"server_frame_type_cancel_push_sent": {
							Type: schema.TypeInt, Optional: true, Description: "CANCEL PUSH Frame sent on server side",
						},
						"server_frame_type_settings_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS Frame received on server side",
						},
						"server_frame_type_settings_sent": {
							Type: schema.TypeInt, Optional: true, Description: "SETTINGS Frame sent on server side",
						},
						"server_frame_type_push_promise_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "PUSH PROMISE Frame received on server side",
						},
						"server_frame_type_push_promise_sent": {
							Type: schema.TypeInt, Optional: true, Description: "PUSH PROMISE Frame sent on server side",
						},
						"server_frame_type_goaway_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GOAWAY Frame received on server side",
						},
						"server_frame_type_goaway_sent": {
							Type: schema.TypeInt, Optional: true, Description: "GOAWAY Frame sent on server side",
						},
						"server_frame_type_max_push_id_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "MAX PUSH ID Frame received on server side",
						},
						"server_frame_type_max_push_id_sent": {
							Type: schema.TypeInt, Optional: true, Description: "MAX PUSH ID Frame sent on server side",
						},
						"server_frame_type_unknown_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Frame received on server side",
						},
						"server_header_frames_to_app": {
							Type: schema.TypeInt, Optional: true, Description: "HEADER Frames passed to HTTP layer on server side",
						},
						"server_data_frames_to_app": {
							Type: schema.TypeInt, Optional: true, Description: "DATA Frames passed to HTTP layer on server side",
						},
						"server_header_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in HEADER frames on server side",
						},
						"server_header_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in HEADER frames on server side",
						},
						"server_data_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in DATA frames on server side",
						},
						"server_data_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in DATA frames on server side",
						},
						"server_other_frame_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on server side",
						},
						"server_other_frame_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on server side",
						},
						"server_heading_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes received in HEADERS/DATA frame/stream heading on server side",
						},
						"server_heading_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Bytes sent in HEADERS/DATA frame/stream heading on server side",
						},
						"server_total_bytes_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bytes received on server side",
						},
						"server_total_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total Bytes sent on server side",
						},
						"invalid_argument": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Argument",
						},
						"invalid_state": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid State",
						},
						"wouldblock": {
							Type: schema.TypeInt, Optional: true, Description: "Wouldblock",
						},
						"stream_in_use": {
							Type: schema.TypeInt, Optional: true, Description: "Stream In Use",
						},
						"push_id_blocked": {
							Type: schema.TypeInt, Optional: true, Description: "Push Id Blocked",
						},
						"malformed_http_header": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Http Header",
						},
						"remove_http_header": {
							Type: schema.TypeInt, Optional: true, Description: "Remove Http Header",
						},
						"malformed_http_messaging": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Http Messaging",
						},
						"too_late": {
							Type: schema.TypeInt, Optional: true, Description: "Too Late",
						},
						"qpack_fatal": {
							Type: schema.TypeInt, Optional: true, Description: "Qpack Fatal",
						},
						"qpack_header_too_large": {
							Type: schema.TypeInt, Optional: true, Description: "Qpack Header Too Large",
						},
						"ignore_stream": {
							Type: schema.TypeInt, Optional: true, Description: "Ignore Stream",
						},
						"stream_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Stream Not Found",
						},
						"ignore_push_promise": {
							Type: schema.TypeInt, Optional: true, Description: "Ignore Push Promise",
						},
						"qpack_decompression_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Qpack Decompression Failed",
						},
						"qpack_encoder_stream_error": {
							Type: schema.TypeInt, Optional: true, Description: "Qpack Encoder Stream Error",
						},
						"qpack_decoder_stream_error": {
							Type: schema.TypeInt, Optional: true, Description: "Qpack Decoder Stream Error",
						},
						"h3_frame_unexpected": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Frame Unexpected",
						},
						"h3_frame_error": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Frame Error",
						},
						"h3_missing_settings": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Missing Settings",
						},
						"h3_internal_error": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Internal Error",
						},
						"h3_closed_critical_stream": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Closed Critical Stream",
						},
						"h3_general_protocol_error": {
							Type: schema.TypeInt, Optional: true, Description: "H3 General Protocol Error",
						},
						"h3_id_error": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Id Error",
						},
						"h3_settings_error": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Settings Error",
						},
						"h3_stream_creation_error": {
							Type: schema.TypeInt, Optional: true, Description: "H3 Stream Creation Error",
						},
						"fatal": {
							Type: schema.TypeInt, Optional: true, Description: "Fatal Error",
						},
						"conn_alloc_error": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP/3 Connection Allocation Error",
						},
						"alloc_fail_total": {
							Type: schema.TypeInt, Optional: true, Description: "Memory Allocation Failures",
						},
						"http3_rejected": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP3 Rejected",
						},
					},
				},
			},
		},
	}
}

func resourceSlbHttp3StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp3StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp3Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbHttp3StatsStats := setObjectSlbHttp3StatsStats(res)
		d.Set("stats", SlbHttp3StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbHttp3StatsStats(ret edpt.DataSlbHttp3Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"client_conn_curr":                    ret.DtSlbHttp3Stats.Stats.Client_conn_curr,
			"server_conn_curr":                    ret.DtSlbHttp3Stats.Stats.Server_conn_curr,
			"client_conn_total":                   ret.DtSlbHttp3Stats.Stats.Client_conn_total,
			"server_conn_total":                   ret.DtSlbHttp3Stats.Stats.Server_conn_total,
			"client_conn_peak":                    ret.DtSlbHttp3Stats.Stats.Client_conn_peak,
			"server_conn_peak":                    ret.DtSlbHttp3Stats.Stats.Server_conn_peak,
			"client_request_streams_curr":         ret.DtSlbHttp3Stats.Stats.Client_request_streams_curr,
			"server_request_streams_curr":         ret.DtSlbHttp3Stats.Stats.Server_request_streams_curr,
			"client_request_streams_total":        ret.DtSlbHttp3Stats.Stats.Client_request_streams_total,
			"server_request_streams_total":        ret.DtSlbHttp3Stats.Stats.Server_request_streams_total,
			"client_request_push_curr":            ret.DtSlbHttp3Stats.Stats.Client_request_push_curr,
			"server_request_push_curr":            ret.DtSlbHttp3Stats.Stats.Server_request_push_curr,
			"client_request_push_total":           ret.DtSlbHttp3Stats.Stats.Client_request_push_total,
			"server_request_push_total":           ret.DtSlbHttp3Stats.Stats.Server_request_push_total,
			"client_request_other_curr":           ret.DtSlbHttp3Stats.Stats.Client_request_other_curr,
			"server_request_other_curr":           ret.DtSlbHttp3Stats.Stats.Server_request_other_curr,
			"client_request_other_total":          ret.DtSlbHttp3Stats.Stats.Client_request_other_total,
			"server_request_other_total":          ret.DtSlbHttp3Stats.Stats.Server_request_other_total,
			"client_frame_type_headers_rcvd":      ret.DtSlbHttp3Stats.Stats.Client_frame_type_headers_rcvd,
			"client_frame_type_headers_sent":      ret.DtSlbHttp3Stats.Stats.Client_frame_type_headers_sent,
			"client_frame_type_data_rcvd":         ret.DtSlbHttp3Stats.Stats.Client_frame_type_data_rcvd,
			"client_frame_type_data_sent":         ret.DtSlbHttp3Stats.Stats.Client_frame_type_data_sent,
			"client_frame_type_cancel_push_rcvd":  ret.DtSlbHttp3Stats.Stats.Client_frame_type_cancel_push_rcvd,
			"client_frame_type_cancel_push_sent":  ret.DtSlbHttp3Stats.Stats.Client_frame_type_cancel_push_sent,
			"client_frame_type_settings_rcvd":     ret.DtSlbHttp3Stats.Stats.Client_frame_type_settings_rcvd,
			"client_frame_type_settings_sent":     ret.DtSlbHttp3Stats.Stats.Client_frame_type_settings_sent,
			"client_frame_type_push_promise_rcvd": ret.DtSlbHttp3Stats.Stats.Client_frame_type_push_promise_rcvd,
			"client_frame_type_push_promise_sent": ret.DtSlbHttp3Stats.Stats.Client_frame_type_push_promise_sent,
			"client_frame_type_goaway_rcvd":       ret.DtSlbHttp3Stats.Stats.Client_frame_type_goaway_rcvd,
			"client_frame_type_goaway_sent":       ret.DtSlbHttp3Stats.Stats.Client_frame_type_goaway_sent,
			"client_frame_type_max_push_id_rcvd":  ret.DtSlbHttp3Stats.Stats.Client_frame_type_max_push_id_rcvd,
			"client_frame_type_max_push_id_sent":  ret.DtSlbHttp3Stats.Stats.Client_frame_type_max_push_id_sent,
			"client_frame_type_unknown_rcvd":      ret.DtSlbHttp3Stats.Stats.Client_frame_type_unknown_rcvd,
			"client_header_frames_to_app":         ret.DtSlbHttp3Stats.Stats.Client_header_frames_to_app,
			"client_data_frames_to_app":           ret.DtSlbHttp3Stats.Stats.Client_data_frames_to_app,
			"client_header_bytes_rcvd":            ret.DtSlbHttp3Stats.Stats.Client_header_bytes_rcvd,
			"client_header_bytes_sent":            ret.DtSlbHttp3Stats.Stats.Client_header_bytes_sent,
			"client_data_bytes_rcvd":              ret.DtSlbHttp3Stats.Stats.Client_data_bytes_rcvd,
			"client_data_bytes_sent":              ret.DtSlbHttp3Stats.Stats.Client_data_bytes_sent,
			"client_other_frame_bytes_rcvd":       ret.DtSlbHttp3Stats.Stats.Client_other_frame_bytes_rcvd,
			"client_other_frame_bytes_sent":       ret.DtSlbHttp3Stats.Stats.Client_other_frame_bytes_sent,
			"client_heading_bytes_rcvd":           ret.DtSlbHttp3Stats.Stats.Client_heading_bytes_rcvd,
			"client_heading_bytes_sent":           ret.DtSlbHttp3Stats.Stats.Client_heading_bytes_sent,
			"client_total_bytes_rcvd":             ret.DtSlbHttp3Stats.Stats.Client_total_bytes_rcvd,
			"client_total_bytes_sent":             ret.DtSlbHttp3Stats.Stats.Client_total_bytes_sent,
			"server_frame_type_headers_rcvd":      ret.DtSlbHttp3Stats.Stats.Server_frame_type_headers_rcvd,
			"server_frame_type_headers_sent":      ret.DtSlbHttp3Stats.Stats.Server_frame_type_headers_sent,
			"server_frame_type_data_rcvd":         ret.DtSlbHttp3Stats.Stats.Server_frame_type_data_rcvd,
			"server_frame_type_data_sent":         ret.DtSlbHttp3Stats.Stats.Server_frame_type_data_sent,
			"server_frame_type_cancel_push_rcvd":  ret.DtSlbHttp3Stats.Stats.Server_frame_type_cancel_push_rcvd,
			"server_frame_type_cancel_push_sent":  ret.DtSlbHttp3Stats.Stats.Server_frame_type_cancel_push_sent,
			"server_frame_type_settings_rcvd":     ret.DtSlbHttp3Stats.Stats.Server_frame_type_settings_rcvd,
			"server_frame_type_settings_sent":     ret.DtSlbHttp3Stats.Stats.Server_frame_type_settings_sent,
			"server_frame_type_push_promise_rcvd": ret.DtSlbHttp3Stats.Stats.Server_frame_type_push_promise_rcvd,
			"server_frame_type_push_promise_sent": ret.DtSlbHttp3Stats.Stats.Server_frame_type_push_promise_sent,
			"server_frame_type_goaway_rcvd":       ret.DtSlbHttp3Stats.Stats.Server_frame_type_goaway_rcvd,
			"server_frame_type_goaway_sent":       ret.DtSlbHttp3Stats.Stats.Server_frame_type_goaway_sent,
			"server_frame_type_max_push_id_rcvd":  ret.DtSlbHttp3Stats.Stats.Server_frame_type_max_push_id_rcvd,
			"server_frame_type_max_push_id_sent":  ret.DtSlbHttp3Stats.Stats.Server_frame_type_max_push_id_sent,
			"server_frame_type_unknown_rcvd":      ret.DtSlbHttp3Stats.Stats.Server_frame_type_unknown_rcvd,
			"server_header_frames_to_app":         ret.DtSlbHttp3Stats.Stats.Server_header_frames_to_app,
			"server_data_frames_to_app":           ret.DtSlbHttp3Stats.Stats.Server_data_frames_to_app,
			"server_header_bytes_rcvd":            ret.DtSlbHttp3Stats.Stats.Server_header_bytes_rcvd,
			"server_header_bytes_sent":            ret.DtSlbHttp3Stats.Stats.Server_header_bytes_sent,
			"server_data_bytes_rcvd":              ret.DtSlbHttp3Stats.Stats.Server_data_bytes_rcvd,
			"server_data_bytes_sent":              ret.DtSlbHttp3Stats.Stats.Server_data_bytes_sent,
			"server_other_frame_bytes_rcvd":       ret.DtSlbHttp3Stats.Stats.Server_other_frame_bytes_rcvd,
			"server_other_frame_bytes_sent":       ret.DtSlbHttp3Stats.Stats.Server_other_frame_bytes_sent,
			"server_heading_bytes_rcvd":           ret.DtSlbHttp3Stats.Stats.Server_heading_bytes_rcvd,
			"server_heading_bytes_sent":           ret.DtSlbHttp3Stats.Stats.Server_heading_bytes_sent,
			"server_total_bytes_rcvd":             ret.DtSlbHttp3Stats.Stats.Server_total_bytes_rcvd,
			"server_total_bytes_sent":             ret.DtSlbHttp3Stats.Stats.Server_total_bytes_sent,
			"invalid_argument":                    ret.DtSlbHttp3Stats.Stats.Invalid_argument,
			"invalid_state":                       ret.DtSlbHttp3Stats.Stats.Invalid_state,
			"wouldblock":                          ret.DtSlbHttp3Stats.Stats.Wouldblock,
			"stream_in_use":                       ret.DtSlbHttp3Stats.Stats.Stream_in_use,
			"push_id_blocked":                     ret.DtSlbHttp3Stats.Stats.Push_id_blocked,
			"malformed_http_header":               ret.DtSlbHttp3Stats.Stats.Malformed_http_header,
			"remove_http_header":                  ret.DtSlbHttp3Stats.Stats.Remove_http_header,
			"malformed_http_messaging":            ret.DtSlbHttp3Stats.Stats.Malformed_http_messaging,
			"too_late":                            ret.DtSlbHttp3Stats.Stats.Too_late,
			"qpack_fatal":                         ret.DtSlbHttp3Stats.Stats.Qpack_fatal,
			"qpack_header_too_large":              ret.DtSlbHttp3Stats.Stats.Qpack_header_too_large,
			"ignore_stream":                       ret.DtSlbHttp3Stats.Stats.Ignore_stream,
			"stream_not_found":                    ret.DtSlbHttp3Stats.Stats.Stream_not_found,
			"ignore_push_promise":                 ret.DtSlbHttp3Stats.Stats.Ignore_push_promise,
			"qpack_decompression_failed":          ret.DtSlbHttp3Stats.Stats.Qpack_decompression_failed,
			"qpack_encoder_stream_error":          ret.DtSlbHttp3Stats.Stats.Qpack_encoder_stream_error,
			"qpack_decoder_stream_error":          ret.DtSlbHttp3Stats.Stats.Qpack_decoder_stream_error,
			"h3_frame_unexpected":                 ret.DtSlbHttp3Stats.Stats.H3_frame_unexpected,
			"h3_frame_error":                      ret.DtSlbHttp3Stats.Stats.H3_frame_error,
			"h3_missing_settings":                 ret.DtSlbHttp3Stats.Stats.H3_missing_settings,
			"h3_internal_error":                   ret.DtSlbHttp3Stats.Stats.H3_internal_error,
			"h3_closed_critical_stream":           ret.DtSlbHttp3Stats.Stats.H3_closed_critical_stream,
			"h3_general_protocol_error":           ret.DtSlbHttp3Stats.Stats.H3_general_protocol_error,
			"h3_id_error":                         ret.DtSlbHttp3Stats.Stats.H3_id_error,
			"h3_settings_error":                   ret.DtSlbHttp3Stats.Stats.H3_settings_error,
			"h3_stream_creation_error":            ret.DtSlbHttp3Stats.Stats.H3_stream_creation_error,
			"fatal":                               ret.DtSlbHttp3Stats.Stats.Fatal,
			"conn_alloc_error":                    ret.DtSlbHttp3Stats.Stats.Conn_alloc_error,
			"alloc_fail_total":                    ret.DtSlbHttp3Stats.Stats.Alloc_fail_total,
			"http3_rejected":                      ret.DtSlbHttp3Stats.Stats.Http3_rejected,
		},
	}
}

func getObjectSlbHttp3StatsStats(d []interface{}) edpt.SlbHttp3StatsStats {

	count1 := len(d)
	var ret edpt.SlbHttp3StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Client_conn_curr = in["client_conn_curr"].(int)
		ret.Server_conn_curr = in["server_conn_curr"].(int)
		ret.Client_conn_total = in["client_conn_total"].(int)
		ret.Server_conn_total = in["server_conn_total"].(int)
		ret.Client_conn_peak = in["client_conn_peak"].(int)
		ret.Server_conn_peak = in["server_conn_peak"].(int)
		ret.Client_request_streams_curr = in["client_request_streams_curr"].(int)
		ret.Server_request_streams_curr = in["server_request_streams_curr"].(int)
		ret.Client_request_streams_total = in["client_request_streams_total"].(int)
		ret.Server_request_streams_total = in["server_request_streams_total"].(int)
		ret.Client_request_push_curr = in["client_request_push_curr"].(int)
		ret.Server_request_push_curr = in["server_request_push_curr"].(int)
		ret.Client_request_push_total = in["client_request_push_total"].(int)
		ret.Server_request_push_total = in["server_request_push_total"].(int)
		ret.Client_request_other_curr = in["client_request_other_curr"].(int)
		ret.Server_request_other_curr = in["server_request_other_curr"].(int)
		ret.Client_request_other_total = in["client_request_other_total"].(int)
		ret.Server_request_other_total = in["server_request_other_total"].(int)
		ret.Client_frame_type_headers_rcvd = in["client_frame_type_headers_rcvd"].(int)
		ret.Client_frame_type_headers_sent = in["client_frame_type_headers_sent"].(int)
		ret.Client_frame_type_data_rcvd = in["client_frame_type_data_rcvd"].(int)
		ret.Client_frame_type_data_sent = in["client_frame_type_data_sent"].(int)
		ret.Client_frame_type_cancel_push_rcvd = in["client_frame_type_cancel_push_rcvd"].(int)
		ret.Client_frame_type_cancel_push_sent = in["client_frame_type_cancel_push_sent"].(int)
		ret.Client_frame_type_settings_rcvd = in["client_frame_type_settings_rcvd"].(int)
		ret.Client_frame_type_settings_sent = in["client_frame_type_settings_sent"].(int)
		ret.Client_frame_type_push_promise_rcvd = in["client_frame_type_push_promise_rcvd"].(int)
		ret.Client_frame_type_push_promise_sent = in["client_frame_type_push_promise_sent"].(int)
		ret.Client_frame_type_goaway_rcvd = in["client_frame_type_goaway_rcvd"].(int)
		ret.Client_frame_type_goaway_sent = in["client_frame_type_goaway_sent"].(int)
		ret.Client_frame_type_max_push_id_rcvd = in["client_frame_type_max_push_id_rcvd"].(int)
		ret.Client_frame_type_max_push_id_sent = in["client_frame_type_max_push_id_sent"].(int)
		ret.Client_frame_type_unknown_rcvd = in["client_frame_type_unknown_rcvd"].(int)
		ret.Client_header_frames_to_app = in["client_header_frames_to_app"].(int)
		ret.Client_data_frames_to_app = in["client_data_frames_to_app"].(int)
		ret.Client_header_bytes_rcvd = in["client_header_bytes_rcvd"].(int)
		ret.Client_header_bytes_sent = in["client_header_bytes_sent"].(int)
		ret.Client_data_bytes_rcvd = in["client_data_bytes_rcvd"].(int)
		ret.Client_data_bytes_sent = in["client_data_bytes_sent"].(int)
		ret.Client_other_frame_bytes_rcvd = in["client_other_frame_bytes_rcvd"].(int)
		ret.Client_other_frame_bytes_sent = in["client_other_frame_bytes_sent"].(int)
		ret.Client_heading_bytes_rcvd = in["client_heading_bytes_rcvd"].(int)
		ret.Client_heading_bytes_sent = in["client_heading_bytes_sent"].(int)
		ret.Client_total_bytes_rcvd = in["client_total_bytes_rcvd"].(int)
		ret.Client_total_bytes_sent = in["client_total_bytes_sent"].(int)
		ret.Server_frame_type_headers_rcvd = in["server_frame_type_headers_rcvd"].(int)
		ret.Server_frame_type_headers_sent = in["server_frame_type_headers_sent"].(int)
		ret.Server_frame_type_data_rcvd = in["server_frame_type_data_rcvd"].(int)
		ret.Server_frame_type_data_sent = in["server_frame_type_data_sent"].(int)
		ret.Server_frame_type_cancel_push_rcvd = in["server_frame_type_cancel_push_rcvd"].(int)
		ret.Server_frame_type_cancel_push_sent = in["server_frame_type_cancel_push_sent"].(int)
		ret.Server_frame_type_settings_rcvd = in["server_frame_type_settings_rcvd"].(int)
		ret.Server_frame_type_settings_sent = in["server_frame_type_settings_sent"].(int)
		ret.Server_frame_type_push_promise_rcvd = in["server_frame_type_push_promise_rcvd"].(int)
		ret.Server_frame_type_push_promise_sent = in["server_frame_type_push_promise_sent"].(int)
		ret.Server_frame_type_goaway_rcvd = in["server_frame_type_goaway_rcvd"].(int)
		ret.Server_frame_type_goaway_sent = in["server_frame_type_goaway_sent"].(int)
		ret.Server_frame_type_max_push_id_rcvd = in["server_frame_type_max_push_id_rcvd"].(int)
		ret.Server_frame_type_max_push_id_sent = in["server_frame_type_max_push_id_sent"].(int)
		ret.Server_frame_type_unknown_rcvd = in["server_frame_type_unknown_rcvd"].(int)
		ret.Server_header_frames_to_app = in["server_header_frames_to_app"].(int)
		ret.Server_data_frames_to_app = in["server_data_frames_to_app"].(int)
		ret.Server_header_bytes_rcvd = in["server_header_bytes_rcvd"].(int)
		ret.Server_header_bytes_sent = in["server_header_bytes_sent"].(int)
		ret.Server_data_bytes_rcvd = in["server_data_bytes_rcvd"].(int)
		ret.Server_data_bytes_sent = in["server_data_bytes_sent"].(int)
		ret.Server_other_frame_bytes_rcvd = in["server_other_frame_bytes_rcvd"].(int)
		ret.Server_other_frame_bytes_sent = in["server_other_frame_bytes_sent"].(int)
		ret.Server_heading_bytes_rcvd = in["server_heading_bytes_rcvd"].(int)
		ret.Server_heading_bytes_sent = in["server_heading_bytes_sent"].(int)
		ret.Server_total_bytes_rcvd = in["server_total_bytes_rcvd"].(int)
		ret.Server_total_bytes_sent = in["server_total_bytes_sent"].(int)
		ret.Invalid_argument = in["invalid_argument"].(int)
		ret.Invalid_state = in["invalid_state"].(int)
		ret.Wouldblock = in["wouldblock"].(int)
		ret.Stream_in_use = in["stream_in_use"].(int)
		ret.Push_id_blocked = in["push_id_blocked"].(int)
		ret.Malformed_http_header = in["malformed_http_header"].(int)
		ret.Remove_http_header = in["remove_http_header"].(int)
		ret.Malformed_http_messaging = in["malformed_http_messaging"].(int)
		ret.Too_late = in["too_late"].(int)
		ret.Qpack_fatal = in["qpack_fatal"].(int)
		ret.Qpack_header_too_large = in["qpack_header_too_large"].(int)
		ret.Ignore_stream = in["ignore_stream"].(int)
		ret.Stream_not_found = in["stream_not_found"].(int)
		ret.Ignore_push_promise = in["ignore_push_promise"].(int)
		ret.Qpack_decompression_failed = in["qpack_decompression_failed"].(int)
		ret.Qpack_encoder_stream_error = in["qpack_encoder_stream_error"].(int)
		ret.Qpack_decoder_stream_error = in["qpack_decoder_stream_error"].(int)
		ret.H3_frame_unexpected = in["h3_frame_unexpected"].(int)
		ret.H3_frame_error = in["h3_frame_error"].(int)
		ret.H3_missing_settings = in["h3_missing_settings"].(int)
		ret.H3_internal_error = in["h3_internal_error"].(int)
		ret.H3_closed_critical_stream = in["h3_closed_critical_stream"].(int)
		ret.H3_general_protocol_error = in["h3_general_protocol_error"].(int)
		ret.H3_id_error = in["h3_id_error"].(int)
		ret.H3_settings_error = in["h3_settings_error"].(int)
		ret.H3_stream_creation_error = in["h3_stream_creation_error"].(int)
		ret.Fatal = in["fatal"].(int)
		ret.Conn_alloc_error = in["conn_alloc_error"].(int)
		ret.Alloc_fail_total = in["alloc_fail_total"].(int)
		ret.Http3_rejected = in["http3_rejected"].(int)
	}
	return ret
}

func dataToEndpointSlbHttp3Stats(d *schema.ResourceData) edpt.SlbHttp3Stats {
	var ret edpt.SlbHttp3Stats

	ret.Stats = getObjectSlbHttp3StatsStats(d.Get("stats").([]interface{}))
	return ret
}
