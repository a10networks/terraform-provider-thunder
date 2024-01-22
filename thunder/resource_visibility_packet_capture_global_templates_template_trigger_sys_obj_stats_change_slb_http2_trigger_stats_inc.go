package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_http2_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"alloc_fail_total": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Alloc Fail - Total",
			},
			"bad_connection_preface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad Connection Preface",
			},
			"bad_frame_type_for_stream_state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Bad frame type for stream state",
			},
			"buff_alloc_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Buff alloc error",
			},
			"cancel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for cancel",
			},
			"cant_allocate_control_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate control frame",
			},
			"cant_allocate_goaway_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate GOAWAY frame",
			},
			"cant_allocate_ping_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate PING frame",
			},
			"cant_allocate_rst_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate RST_STREAM frame",
			},
			"cant_allocate_settings_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate SETTINGS frame",
			},
			"cant_allocate_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate stream",
			},
			"cant_allocate_window_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Cant allocate WINDOW_UPDATE frame",
			},
			"closed_state_unexpected_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in closed state",
			},
			"compression_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for compression error",
			},
			"connect_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for connect error",
			},
			"continuation_before_headers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for CONTINUATION frame with no headers frame",
			},
			"data_no_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for DATA Frame Rcvd on non-existent stream",
			},
			"data_queue_alloc_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data Queue Alloc Error",
			},
			"deflate_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for deflate alloc fail",
			},
			"enhance_your_calm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for enhance your calm error",
			},
			"err_rcvd_total": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rcvd - Total",
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
			"err_sent_flow_control": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FLOW_CONTROL_ERROR",
			},
			"err_sent_frame_size_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - FRAME_SIZE_ERROR",
			},
			"err_sent_http11_required": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - HTTP_1_1_REQUIRED",
			},
			"err_sent_inadequate_security": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INADEQUATE_SECURITY",
			},
			"err_sent_internal_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - INTERNAL_ERROR",
			},
			"err_sent_proto_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - PROTOCOL_ERROR",
			},
			"err_sent_refused_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - REFUSED_STREAM",
			},
			"err_sent_setting_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - SETTINGS_TIMEOUT",
			},
			"err_sent_stream_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - STREAM_CLOSED",
			},
			"err_sent_total": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Rent - Total",
			},
			"err_sent_your_calm": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Error Sent - ENHANCE_YOUR_CALM",
			},
			"error_max_invalid_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Max Invalid Stream Rcvd",
			},
			"exceeds_max_window_size_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with increment that results in exceeding max window",
			},
			"flow_control_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Flow Control Error",
			},
			"frame_size_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Frame Size Error",
			},
			"half_closed_remote_state_unexpected_fra": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in half closed remote state",
			},
			"header_no_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header no stream",
			},
			"header_padlen_gt_frame_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header padlen greater than frame payload size",
			},
			"headers_after_continuation": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers frame before CONTINUATION was complete",
			},
			"headers_interleaved": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for headers interleaved on streams",
			},
			"http_1_1_required": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP1.1 Required",
			},
			"idle_state_unexpected_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unxpected frame received in idle state",
			},
			"inadequate_security": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inadequate security",
			},
			"inflate_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for inflate alloc fail",
			},
			"inflate_header_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Inflate Header Fail",
			},
			"internal_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Internal Error",
			},
			"invalid_frame_during_headers": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for frame before headers were complete",
			},
			"invalid_frame_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid Frame Size Rcvd",
			},
			"invalid_push_promise": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for unexpected PUSH_PROMISE frame",
			},
			"invalid_setting_value": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for invalid setting-frame value",
			},
			"invalid_stream_id": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for received invalid stream ID",
			},
			"invalid_window_update": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for window-update value out of range",
			},
			"protocol_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Protocol Error",
			},
			"proxy_alloc_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP2 Proxy alloc Error",
			},
			"refused_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Refused Stream",
			},
			"reserved_local_state_unexpected_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved local state",
			},
			"reserved_remote_state_unexpected_frame": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Unexpected frame received in reserved remote state",
			},
			"settings_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Settings Timeout",
			},
			"split_buff_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Splitting Buffer Failed",
			},
			"streams_gt_max_concur_streams": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Streams greater than max allowed concurrent streams",
			},
			"trailers_no_end_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for trailers not marked as end-of-stream",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"wrong_stream_state": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Wrong Stream State",
			},
			"zero_window_size_on_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Window Update with zero increment rcvd",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbHttp2TriggerStatsInc
	ret.Inst.Alloc_fail_total = d.Get("alloc_fail_total").(int)
	ret.Inst.Bad_connection_preface = d.Get("bad_connection_preface").(int)
	ret.Inst.Bad_frame_type_for_stream_state = d.Get("bad_frame_type_for_stream_state").(int)
	ret.Inst.Buff_alloc_error = d.Get("buff_alloc_error").(int)
	ret.Inst.Cancel = d.Get("cancel").(int)
	ret.Inst.Cant_allocate_control_frame = d.Get("cant_allocate_control_frame").(int)
	ret.Inst.Cant_allocate_goaway_frame = d.Get("cant_allocate_goaway_frame").(int)
	ret.Inst.Cant_allocate_ping_frame = d.Get("cant_allocate_ping_frame").(int)
	ret.Inst.Cant_allocate_rst_frame = d.Get("cant_allocate_rst_frame").(int)
	ret.Inst.Cant_allocate_settings_frame = d.Get("cant_allocate_settings_frame").(int)
	ret.Inst.Cant_allocate_stream = d.Get("cant_allocate_stream").(int)
	ret.Inst.Cant_allocate_window_frame = d.Get("cant_allocate_window_frame").(int)
	ret.Inst.Closed_state_unexpected_frame = d.Get("closed_state_unexpected_frame").(int)
	ret.Inst.Compression_error = d.Get("compression_error").(int)
	ret.Inst.Connect_error = d.Get("connect_error").(int)
	ret.Inst.Continuation_before_headers = d.Get("continuation_before_headers").(int)
	ret.Inst.Data_no_stream = d.Get("data_no_stream").(int)
	ret.Inst.Data_queue_alloc_error = d.Get("data_queue_alloc_error").(int)
	ret.Inst.Deflate_alloc_fail = d.Get("deflate_alloc_fail").(int)
	ret.Inst.Enhance_your_calm = d.Get("enhance_your_calm").(int)
	ret.Inst.Err_rcvd_total = d.Get("err_rcvd_total").(int)
	ret.Inst.Err_sent_cancel = d.Get("err_sent_cancel").(int)
	ret.Inst.Err_sent_compression_err = d.Get("err_sent_compression_err").(int)
	ret.Inst.Err_sent_connect_err = d.Get("err_sent_connect_err").(int)
	ret.Inst.Err_sent_flow_control = d.Get("err_sent_flow_control").(int)
	ret.Inst.Err_sent_frame_size_err = d.Get("err_sent_frame_size_err").(int)
	ret.Inst.Err_sent_http11_required = d.Get("err_sent_http11_required").(int)
	ret.Inst.Err_sent_inadequate_security = d.Get("err_sent_inadequate_security").(int)
	ret.Inst.Err_sent_internal_err = d.Get("err_sent_internal_err").(int)
	ret.Inst.Err_sent_proto_err = d.Get("err_sent_proto_err").(int)
	ret.Inst.Err_sent_refused_stream = d.Get("err_sent_refused_stream").(int)
	ret.Inst.Err_sent_setting_timeout = d.Get("err_sent_setting_timeout").(int)
	ret.Inst.Err_sent_stream_closed = d.Get("err_sent_stream_closed").(int)
	ret.Inst.Err_sent_total = d.Get("err_sent_total").(int)
	ret.Inst.Err_sent_your_calm = d.Get("err_sent_your_calm").(int)
	ret.Inst.Error_max_invalid_stream = d.Get("error_max_invalid_stream").(int)
	ret.Inst.Exceeds_max_window_size_stream = d.Get("exceeds_max_window_size_stream").(int)
	ret.Inst.Flow_control_error = d.Get("flow_control_error").(int)
	ret.Inst.Frame_size_error = d.Get("frame_size_error").(int)
	ret.Inst.Half_closed_remote_state_unexpected_fra = d.Get("half_closed_remote_state_unexpected_fra").(int)
	ret.Inst.Header_no_stream = d.Get("header_no_stream").(int)
	ret.Inst.Header_padlen_gt_frame_payload = d.Get("header_padlen_gt_frame_payload").(int)
	ret.Inst.Headers_after_continuation = d.Get("headers_after_continuation").(int)
	ret.Inst.Headers_interleaved = d.Get("headers_interleaved").(int)
	ret.Inst.Http_1_1_required = d.Get("http_1_1_required").(int)
	ret.Inst.Idle_state_unexpected_frame = d.Get("idle_state_unexpected_frame").(int)
	ret.Inst.Inadequate_security = d.Get("inadequate_security").(int)
	ret.Inst.Inflate_alloc_fail = d.Get("inflate_alloc_fail").(int)
	ret.Inst.Inflate_header_fail = d.Get("inflate_header_fail").(int)
	ret.Inst.Internal_error = d.Get("internal_error").(int)
	ret.Inst.Invalid_frame_during_headers = d.Get("invalid_frame_during_headers").(int)
	ret.Inst.Invalid_frame_size = d.Get("invalid_frame_size").(int)
	ret.Inst.Invalid_push_promise = d.Get("invalid_push_promise").(int)
	ret.Inst.Invalid_setting_value = d.Get("invalid_setting_value").(int)
	ret.Inst.Invalid_stream_id = d.Get("invalid_stream_id").(int)
	ret.Inst.Invalid_window_update = d.Get("invalid_window_update").(int)
	ret.Inst.Protocol_error = d.Get("protocol_error").(int)
	ret.Inst.Proxy_alloc_error = d.Get("proxy_alloc_error").(int)
	ret.Inst.Refused_stream = d.Get("refused_stream").(int)
	ret.Inst.Reserved_local_state_unexpected_frame = d.Get("reserved_local_state_unexpected_frame").(int)
	ret.Inst.Reserved_remote_state_unexpected_frame = d.Get("reserved_remote_state_unexpected_frame").(int)
	ret.Inst.Settings_timeout = d.Get("settings_timeout").(int)
	ret.Inst.Split_buff_fail = d.Get("split_buff_fail").(int)
	ret.Inst.Streams_gt_max_concur_streams = d.Get("streams_gt_max_concur_streams").(int)
	ret.Inst.Trailers_no_end_stream = d.Get("trailers_no_end_stream").(int)
	//omit uuid
	ret.Inst.Wrong_stream_state = d.Get("wrong_stream_state").(int)
	ret.Inst.Zero_window_size_on_stream = d.Get("zero_window_size_on_stream").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
