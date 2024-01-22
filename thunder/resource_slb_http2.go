package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttp2() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_http2`: Configure http2\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHttp2Create,
		UpdateContext: resourceSlbHttp2Update,
		ReadContext:   resourceSlbHttp2Read,
		DeleteContext: resourceSlbHttp2Delete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_proxy': Curr Proxy Conns; 'total_proxy': Total Proxy Conns; 'connection_preface_rcvd': Connection preface rcvd; 'control_frame': Control Frame Rcvd; 'headers_frame': HEADERS Frame Rcvd; 'continuation_frame': CONTINUATION Frame Rcvd; 'rst_frame_rcvd': RST_STREAM Frame Rcvd; 'settings_frame': SETTINGS Frame Rcvd; 'window_update_frame': WINDOW_UPDATE Frame Rcvd; 'ping_frame': PING Frame Rcvd; 'goaway_frame': GOAWAY Frame Rcvd; 'priority_frame': PRIORITY Frame Rcvd; 'data_frame': DATA Frame Recvd; 'unknown_frame': Unknown Frame Recvd; 'connection_preface_sent': Connection preface sent; 'settings_frame_sent': SETTINGS Frame Sent; 'settings_ack_sent': SETTINGS ACK Frame Sent; 'empty_settings_sent': Empty SETTINGS Frame Sent; 'ping_frame_sent': PING Frame Sent; 'window_update_frame_sent': WINDOW_UPDATE Frame Sent; 'rst_frame_sent': RST_STREAM Frame Sent; 'goaway_frame_sent': GOAWAY Frame Sent; 'header_to_app': HEADER Frame to HTTP; 'data_to_app': DATA Frame to HTTP; 'protocol_error': Protocol Error; 'internal_error': Internal Error; 'proxy_alloc_error': HTTP2 Proxy alloc Error; 'split_buff_fail': Splitting Buffer Failed; 'invalid_frame_size': Invalid Frame Size Rcvd; 'error_max_invalid_stream': Max Invalid Stream Rcvd; 'data_no_stream': DATA Frame Rcvd on non-existent stream; 'flow_control_error': Flow Control Error; 'settings_timeout': Settings Timeout; 'frame_size_error': Frame Size Error; 'refused_stream': Refused Stream; 'cancel': cancel; 'compression_error': compression error; 'connect_error': connect error; 'enhance_your_calm': enhance your calm error; 'inadequate_security': inadequate security; 'http_1_1_required': HTTP1.1 Required; 'deflate_alloc_fail': deflate alloc fail; 'inflate_alloc_fail': inflate alloc fail; 'inflate_header_fail': Inflate Header Fail; 'bad_connection_preface': Bad Connection Preface; 'cant_allocate_control_frame': Cant allocate control frame; 'cant_allocate_settings_frame': Cant allocate SETTINGS frame; 'bad_frame_type_for_stream_state': Bad frame type for stream state; 'wrong_stream_state': Wrong Stream State; 'data_queue_alloc_error': Data Queue Alloc Error; 'buff_alloc_error': Buff alloc error; 'cant_allocate_rst_frame': Cant allocate RST_STREAM frame; 'cant_allocate_goaway_frame': Cant allocate GOAWAY frame; 'cant_allocate_ping_frame': Cant allocate PING frame; 'cant_allocate_stream': Cant allocate stream; 'cant_allocate_window_frame': Cant allocate WINDOW_UPDATE frame; 'header_no_stream': header no stream; 'header_padlen_gt_frame_payload': Header padlen greater than frame payload size; 'streams_gt_max_concur_streams': Streams greater than max allowed concurrent streams; 'idle_state_unexpected_frame': Unxpected frame received in idle state; 'reserved_local_state_unexpected_frame': Unexpected frame received in reserved local state; 'reserved_remote_state_unexpected_frame': Unexpected frame received in reserved remote state; 'half_closed_remote_state_unexpected_frame': Unexpected frame received in half closed remote state; 'closed_state_unexpected_frame': Unexpected frame received in closed state; 'zero_window_size_on_stream': Window Update with zero increment rcvd; 'exceeds_max_window_size_stream': Window Update with increment that results in exceeding max window; 'stream_closed': stream closed; 'continuation_before_headers': CONTINUATION frame with no headers frame; 'invalid_frame_during_headers': frame before headers were complete; 'headers_after_continuation': headers frame before CONTINUATION was complete; 'push_promise_frame_sent': Push Promise Frame Sent; 'invalid_push_promise': unexpected PUSH_PROMISE frame; 'invalid_stream_id': received invalid stream ID; 'headers_interleaved': headers interleaved on streams; 'trailers_no_end_stream': trailers not marked as end-of-stream; 'invalid_setting_value': invalid setting-frame value; 'invalid_window_update': window-update value out of range; 'frame_header_bytes_received': frame header bytes received; 'frame_header_bytes_sent': frame header bytes sent; 'control_bytes_received': HTTP/2 control frame bytes received; 'control_bytes_sent': HTTP/2 control frame bytes sent; 'header_bytes_received': HTTP/2 header bytes received; 'header_bytes_sent': HTTP/2 header bytes sent; 'data_bytes_received': HTTP/2 data bytes received; 'data_bytes_sent': HTTP/2 data bytes sent; 'total_bytes_received': HTTP/2 total bytes received; 'total_bytes_sent': HTTP/2 total bytes sent; 'peak_proxy': Peak Proxy Conns; 'control_frame_sent': Control Frame Sent; 'continuation_frame_sent': CONTINUATION Frame Sent; 'data_frame_sent': DATA Frame Sent; 'headers_frame_sent': HEADERS Frame Sent; 'priority_frame_sent': PRIORITY Frame Sent; 'settings_ack_rcvd': SETTINGS ACK Frame Rcvd; 'empty_settings_rcvd': Empty SETTINGS Frame Rcvd; 'alloc_fail_total': Alloc Fail - Total; 'err_rcvd_total': Error Rcvd - Total; 'err_sent_total': Error Rent - Total; 'err_sent_proto_err': Error Sent - PROTOCOL_ERROR; 'err_sent_internal_err': Error Sent - INTERNAL_ERROR; 'err_sent_flow_control': Error Sent - FLOW_CONTROL_ERROR; 'err_sent_setting_timeout': Error Sent - SETTINGS_TIMEOUT; 'err_sent_stream_closed': Error Sent - STREAM_CLOSED; 'err_sent_frame_size_err': Error Sent - FRAME_SIZE_ERROR; 'err_sent_refused_stream': Error Sent - REFUSED_STREAM; 'err_sent_cancel': Error Sent - CANCEL; 'err_sent_compression_err': Error Sent - COMPRESSION_ERROR; 'err_sent_connect_err': Error Sent - CONNECT_ERROR; 'err_sent_your_calm': Error Sent - ENHANCE_YOUR_CALM; 'err_sent_inadequate_security': Error Sent - INADEQUATE_SECURITY; 'err_sent_http11_required': Error Sent - HTTP_1_1_REQUIRED; 'http2_rejected': HTTP2 Rejected; 'current_stream': Current Streams; 'stream_create': Stream Create; 'stream_free': Stream Free; 'end_stream_rcvd': End Stream Recieved; 'end_stream_sent': End Stream Sent;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbHttp2Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp2Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp2(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttp2Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbHttp2Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp2Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp2(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttp2Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbHttp2Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp2Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp2(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHttp2Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp2Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp2(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHttp2SamplingEnable(d []interface{}) []edpt.SlbHttp2SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHttp2SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHttp2SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHttp2(d *schema.ResourceData) edpt.SlbHttp2 {
	var ret edpt.SlbHttp2
	ret.Inst.SamplingEnable = getSliceSlbHttp2SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
