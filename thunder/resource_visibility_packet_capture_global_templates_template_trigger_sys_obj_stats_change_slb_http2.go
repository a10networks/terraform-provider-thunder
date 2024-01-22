package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http2`: Configure triggers for slb.http2 object\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Create,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Update,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Read,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Delete,

		Schema: map[string]*schema.Schema{
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Protocol Error",
						},
						"internal_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
						},
						"proxy_alloc_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP2 Proxy alloc Error",
						},
						"split_buff_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Splitting Buffer Failed",
						},
						"invalid_frame_size": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Frame Size Rcvd",
						},
						"error_max_invalid_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max Invalid Stream Rcvd",
						},
						"data_no_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DATA Frame Rcvd on non-existent stream",
						},
						"flow_control_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Flow Control Error",
						},
						"settings_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Settings Timeout",
						},
						"frame_size_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Frame Size Error",
						},
						"refused_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Refused Stream",
						},
						"cancel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cancel",
						},
						"compression_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression error",
						},
						"connect_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connect error",
						},
						"enhance_your_calm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for enhance your calm error",
						},
						"inadequate_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inadequate security",
						},
						"http_1_1_required": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP1.1 Required",
						},
						"deflate_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for deflate alloc fail",
						},
						"inflate_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inflate alloc fail",
						},
						"inflate_header_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inflate Header Fail",
						},
						"bad_connection_preface": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Connection Preface",
						},
						"cant_allocate_control_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate control frame",
						},
						"cant_allocate_settings_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate SETTINGS frame",
						},
						"bad_frame_type_for_stream_state": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad frame type for stream state",
						},
						"wrong_stream_state": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wrong Stream State",
						},
						"data_queue_alloc_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Queue Alloc Error",
						},
						"buff_alloc_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff alloc error",
						},
						"cant_allocate_rst_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate RST_STREAM frame",
						},
						"cant_allocate_goaway_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate GOAWAY frame",
						},
						"cant_allocate_ping_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate PING frame",
						},
						"cant_allocate_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate stream",
						},
						"cant_allocate_window_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate WINDOW_UPDATE frame",
						},
						"header_no_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header no stream",
						},
						"header_padlen_gt_frame_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header padlen greater than frame payload size",
						},
						"streams_gt_max_concur_streams": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Streams greater than max allowed concurrent streams",
						},
						"idle_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unxpected frame received in idle state",
						},
						"reserved_local_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved local state",
						},
						"reserved_remote_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved remote state",
						},
						"half_closed_remote_state_unexpected_fra": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in half closed remote state",
						},
						"closed_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in closed state",
						},
						"zero_window_size_on_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with zero increment rcvd",
						},
						"exceeds_max_window_size_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with increment that results in exceeding max window",
						},
						"continuation_before_headers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CONTINUATION frame with no headers frame",
						},
						"invalid_frame_during_headers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for frame before headers were complete",
						},
						"headers_after_continuation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers frame before CONTINUATION was complete",
						},
						"invalid_push_promise": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected PUSH_PROMISE frame",
						},
						"invalid_stream_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received invalid stream ID",
						},
						"headers_interleaved": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers interleaved on streams",
						},
						"trailers_no_end_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for trailers not marked as end-of-stream",
						},
						"invalid_setting_value": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid setting-frame value",
						},
						"invalid_window_update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for window-update value out of range",
						},
						"alloc_fail_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Alloc Fail - Total",
						},
						"err_rcvd_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rcvd - Total",
						},
						"err_sent_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rent - Total",
						},
						"err_sent_proto_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - PROTOCOL_ERROR",
						},
						"err_sent_internal_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INTERNAL_ERROR",
						},
						"err_sent_flow_control": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FLOW_CONTROL_ERROR",
						},
						"err_sent_setting_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - SETTINGS_TIMEOUT",
						},
						"err_sent_stream_closed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - STREAM_CLOSED",
						},
						"err_sent_frame_size_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FRAME_SIZE_ERROR",
						},
						"err_sent_refused_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - REFUSED_STREAM",
						},
						"err_sent_cancel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CANCEL",
						},
						"err_sent_compression_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - COMPRESSION_ERROR",
						},
						"err_sent_connect_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CONNECT_ERROR",
						},
						"err_sent_your_calm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - ENHANCE_YOUR_CALM",
						},
						"err_sent_inadequate_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INADEQUATE_SECURITY",
						},
						"err_sent_http11_required": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - HTTP_1_1_REQUIRED",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_rate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"threshold_exceeded_by": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
						},
						"duration": {
							Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
						},
						"protocol_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Protocol Error",
						},
						"internal_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
						},
						"proxy_alloc_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP2 Proxy alloc Error",
						},
						"split_buff_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Splitting Buffer Failed",
						},
						"invalid_frame_size": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Frame Size Rcvd",
						},
						"error_max_invalid_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max Invalid Stream Rcvd",
						},
						"data_no_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DATA Frame Rcvd on non-existent stream",
						},
						"flow_control_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Flow Control Error",
						},
						"settings_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Settings Timeout",
						},
						"frame_size_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Frame Size Error",
						},
						"refused_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Refused Stream",
						},
						"cancel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cancel",
						},
						"compression_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression error",
						},
						"connect_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connect error",
						},
						"enhance_your_calm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for enhance your calm error",
						},
						"inadequate_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inadequate security",
						},
						"http_1_1_required": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP1.1 Required",
						},
						"deflate_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for deflate alloc fail",
						},
						"inflate_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inflate alloc fail",
						},
						"inflate_header_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inflate Header Fail",
						},
						"bad_connection_preface": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Connection Preface",
						},
						"cant_allocate_control_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate control frame",
						},
						"cant_allocate_settings_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate SETTINGS frame",
						},
						"bad_frame_type_for_stream_state": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad frame type for stream state",
						},
						"wrong_stream_state": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wrong Stream State",
						},
						"data_queue_alloc_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Queue Alloc Error",
						},
						"buff_alloc_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff alloc error",
						},
						"cant_allocate_rst_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate RST_STREAM frame",
						},
						"cant_allocate_goaway_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate GOAWAY frame",
						},
						"cant_allocate_ping_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate PING frame",
						},
						"cant_allocate_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate stream",
						},
						"cant_allocate_window_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate WINDOW_UPDATE frame",
						},
						"header_no_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header no stream",
						},
						"header_padlen_gt_frame_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header padlen greater than frame payload size",
						},
						"streams_gt_max_concur_streams": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Streams greater than max allowed concurrent streams",
						},
						"idle_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unxpected frame received in idle state",
						},
						"reserved_local_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved local state",
						},
						"reserved_remote_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved remote state",
						},
						"half_closed_remote_state_unexpected_fra": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in half closed remote state",
						},
						"closed_state_unexpected_frame": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in closed state",
						},
						"zero_window_size_on_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with zero increment rcvd",
						},
						"exceeds_max_window_size_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with increment that results in exceeding max window",
						},
						"continuation_before_headers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CONTINUATION frame with no headers frame",
						},
						"invalid_frame_during_headers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for frame before headers were complete",
						},
						"headers_after_continuation": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers frame before CONTINUATION was complete",
						},
						"invalid_push_promise": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected PUSH_PROMISE frame",
						},
						"invalid_stream_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received invalid stream ID",
						},
						"headers_interleaved": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers interleaved on streams",
						},
						"trailers_no_end_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for trailers not marked as end-of-stream",
						},
						"invalid_setting_value": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid setting-frame value",
						},
						"invalid_window_update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for window-update value out of range",
						},
						"alloc_fail_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Alloc Fail - Total",
						},
						"err_rcvd_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rcvd - Total",
						},
						"err_sent_total": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rent - Total",
						},
						"err_sent_proto_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - PROTOCOL_ERROR",
						},
						"err_sent_internal_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INTERNAL_ERROR",
						},
						"err_sent_flow_control": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FLOW_CONTROL_ERROR",
						},
						"err_sent_setting_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - SETTINGS_TIMEOUT",
						},
						"err_sent_stream_closed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - STREAM_CLOSED",
						},
						"err_sent_frame_size_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FRAME_SIZE_ERROR",
						},
						"err_sent_refused_stream": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - REFUSED_STREAM",
						},
						"err_sent_cancel": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CANCEL",
						},
						"err_sent_compression_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - COMPRESSION_ERROR",
						},
						"err_sent_connect_err": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - CONNECT_ERROR",
						},
						"err_sent_your_calm": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - ENHANCE_YOUR_CALM",
						},
						"err_sent_inadequate_security": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INADEQUATE_SECURITY",
						},
						"err_sent_http11_required": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - HTTP_1_1_REQUIRED",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Read(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Read(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2037(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2037 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2037
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.Half_closed_remote_state_unexpected_fra = in["half_closed_remote_state_unexpected_fra"].(int)
		ret.Closed_state_unexpected_frame = in["closed_state_unexpected_frame"].(int)
		ret.Zero_window_size_on_stream = in["zero_window_size_on_stream"].(int)
		ret.Exceeds_max_window_size_stream = in["exceeds_max_window_size_stream"].(int)
		ret.Continuation_before_headers = in["continuation_before_headers"].(int)
		ret.Invalid_frame_during_headers = in["invalid_frame_during_headers"].(int)
		ret.Headers_after_continuation = in["headers_after_continuation"].(int)
		ret.Invalid_push_promise = in["invalid_push_promise"].(int)
		ret.Invalid_stream_id = in["invalid_stream_id"].(int)
		ret.Headers_interleaved = in["headers_interleaved"].(int)
		ret.Trailers_no_end_stream = in["trailers_no_end_stream"].(int)
		ret.Invalid_setting_value = in["invalid_setting_value"].(int)
		ret.Invalid_window_update = in["invalid_window_update"].(int)
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
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2038(d []interface{}) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2038 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2038
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
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
		ret.Half_closed_remote_state_unexpected_fra = in["half_closed_remote_state_unexpected_fra"].(int)
		ret.Closed_state_unexpected_frame = in["closed_state_unexpected_frame"].(int)
		ret.Zero_window_size_on_stream = in["zero_window_size_on_stream"].(int)
		ret.Exceeds_max_window_size_stream = in["exceeds_max_window_size_stream"].(int)
		ret.Continuation_before_headers = in["continuation_before_headers"].(int)
		ret.Invalid_frame_during_headers = in["invalid_frame_during_headers"].(int)
		ret.Headers_after_continuation = in["headers_after_continuation"].(int)
		ret.Invalid_push_promise = in["invalid_push_promise"].(int)
		ret.Invalid_stream_id = in["invalid_stream_id"].(int)
		ret.Headers_interleaved = in["headers_interleaved"].(int)
		ret.Trailers_no_end_stream = in["trailers_no_end_stream"].(int)
		ret.Invalid_setting_value = in["invalid_setting_value"].(int)
		ret.Invalid_window_update = in["invalid_window_update"].(int)
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
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2 {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc2037(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsRate2038(d.Get("trigger_stats_rate").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
