package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbHttp3() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_http3`: Configure http3\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbHttp3Create,
		UpdateContext: resourceSlbHttp3Update,
		ReadContext:   resourceSlbHttp3Read,
		DeleteContext: resourceSlbHttp3Delete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'client_conn_curr': Current HTTP/3 Client Connections; 'server_conn_curr': Current HTTP/3 Server Connections; 'client_conn_total': Total HTTP/3 Client Connections; 'server_conn_total': Total HTTP/3 Server Connections; 'client_conn_peak': Peak HTTP/3 Client Connections; 'server_conn_peak': Peak HTTP/3 Server Connections; 'client_request_streams_curr': Current Request Streams on client side; 'server_request_streams_curr': Current Request Streams on server side; 'client_request_streams_total': Total Request Streams on client side; 'server_request_streams_total': Total Request Streams on server side; 'client_request_push_curr': Current Push Streams on client side; 'server_request_push_curr': Current Push Streams on server side; 'client_request_push_total': Total Push Streams on client side; 'server_request_push_total': Total Push Streams on server side; 'client_request_other_curr': Current Other Streams on client side (control, decoder, encoder); 'server_request_other_curr': urrent Other Streams on server side (control, decoder, encoder); 'client_request_other_total': Total Other Streams on client side (control, decoder, encoder); 'server_request_other_total': Total Other Streams on server side (control, decoder, encoder); 'client_frame_type_headers_rcvd': HEADERS Frame received on client side; 'client_frame_type_headers_sent': HEADERS Frame sent on client side; 'client_frame_type_data_rcvd': DATA Frame received on client side; 'client_frame_type_data_sent': DATA Frame sent on client side; 'client_frame_type_cancel_push_rcvd': CANCEL PUSH Frame received on client side; 'client_frame_type_cancel_push_sent': CANCEL PUSH Frame sent on client side; 'client_frame_type_settings_rcvd': SETTINGS Frame received on client side; 'client_frame_type_settings_sent': SETTINGS Frame sent on client side; 'client_frame_type_push_promise_rcvd': PUSH PROMISE Frame received on client side; 'client_frame_type_push_promise_sent': PUSH PROMISE Frame sent on client side; 'client_frame_type_goaway_rcvd': GOAWAY Frame received on client side; 'client_frame_type_goaway_sent': GOAWAY Frame sent on client side; 'client_frame_type_max_push_id_rcvd': MAX PUSH ID Frame received on client side; 'client_frame_type_max_push_id_sent': MAX PUSH ID Frame sent on client side; 'client_frame_type_unknown_rcvd': Unknown Frame received on client side; 'client_header_frames_to_app': HEADER Frames passed to HTTP layer on client side; 'client_data_frames_to_app': DATA Frames passed to HTTP layer on client side; 'client_header_bytes_rcvd': Bytes received in HEADER frames on client side; 'client_header_bytes_sent': Bytes sent in HEADER frames on client side; 'client_data_bytes_rcvd': Bytes received in DATA frames on client side; 'client_data_bytes_sent': Bytes sent in DATA frames on client side; 'client_other_frame_bytes_rcvd': Bytes received in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on client side; 'client_other_frame_bytes_sent': Bytes sent in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on client side; 'client_heading_bytes_rcvd': Bytes received in HEADERS/DATA frame/stream heading on client side; 'client_heading_bytes_sent': Bytes sent in HEADERS/DATA frame/stream heading on client side; 'client_total_bytes_rcvd': Total Bytes received on client side; 'client_total_bytes_sent': Total Bytes sent on client side; 'server_frame_type_headers_rcvd': HEADERS Frame received on server side; 'server_frame_type_headers_sent': HEADERS Frame sent on server side; 'server_frame_type_data_rcvd': DATA Frame received on server side; 'server_frame_type_data_sent': DATA Frame sent on server side; 'server_frame_type_cancel_push_rcvd': CANCEL PUSH Frame received on server side; 'server_frame_type_cancel_push_sent': CANCEL PUSH Frame sent on server side; 'server_frame_type_settings_rcvd': SETTINGS Frame received on server side; 'server_frame_type_settings_sent': SETTINGS Frame sent on server side; 'server_frame_type_push_promise_rcvd': PUSH PROMISE Frame received on server side; 'server_frame_type_push_promise_sent': PUSH PROMISE Frame sent on server side; 'server_frame_type_goaway_rcvd': GOAWAY Frame received on server side; 'server_frame_type_goaway_sent': GOAWAY Frame sent on server side; 'server_frame_type_max_push_id_rcvd': MAX PUSH ID Frame received on server side; 'server_frame_type_max_push_id_sent': MAX PUSH ID Frame sent on server side; 'server_frame_type_unknown_rcvd': Unknown Frame received on server side; 'server_header_frames_to_app': HEADER Frames passed to HTTP layer on server side; 'server_data_frames_to_app': DATA Frames passed to HTTP layer on server side; 'server_header_bytes_rcvd': Bytes received in HEADER frames on server side; 'server_header_bytes_sent': Bytes sent in HEADER frames on server side; 'server_data_bytes_rcvd': Bytes received in DATA frames on server side; 'server_data_bytes_sent': Bytes sent in DATA frames on server side; 'server_other_frame_bytes_rcvd': Bytes received in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on server side; 'server_other_frame_bytes_sent': Bytes sent in other frames (SETTINGS, GOAWAY, PUSH_PROMISE etc) on server side; 'server_heading_bytes_rcvd': Bytes received in HEADERS/DATA frame/stream heading on server side; 'server_heading_bytes_sent': Bytes sent in HEADERS/DATA frame/stream heading on server side; 'server_total_bytes_rcvd': Total Bytes received on server side; 'server_total_bytes_sent': Total Bytes sent on server side; 'invalid_argument': Invalid Argument; 'invalid_state': Invalid State; 'wouldblock': Wouldblock; 'stream_in_use': Stream In Use; 'push_id_blocked': Push Id Blocked; 'malformed_http_header': Malformed Http Header; 'remove_http_header': Remove Http Header; 'malformed_http_messaging': Malformed Http Messaging; 'too_late': Too Late; 'qpack_fatal': Qpack Fatal; 'qpack_header_too_large': Qpack Header Too Large; 'ignore_stream': Ignore Stream; 'stream_not_found': Stream Not Found; 'ignore_push_promise': Ignore Push Promise; 'qpack_decompression_failed': Qpack Decompression Failed; 'qpack_encoder_stream_error': Qpack Encoder Stream Error; 'qpack_decoder_stream_error': Qpack Decoder Stream Error; 'h3_frame_unexpected': H3 Frame Unexpected; 'h3_frame_error': H3 Frame Error; 'h3_missing_settings': H3 Missing Settings; 'h3_internal_error': H3 Internal Error; 'h3_closed_critical_stream': H3 Closed Critical Stream; 'h3_general_protocol_error': H3 General Protocol Error; 'h3_id_error': H3 Id Error; 'h3_settings_error': H3 Settings Error; 'h3_stream_creation_error': H3 Stream Creation Error; 'fatal': Fatal Error; 'conn_alloc_error': HTTP/3 Connection Allocation Error; 'alloc_fail_total': Memory Allocation Failures; 'http3_rejected': HTTP3 Rejected;",
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
func resourceSlbHttp3Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp3Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp3(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttp3Read(ctx, d, meta)
	}
	return diags
}

func resourceSlbHttp3Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp3Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp3(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbHttp3Read(ctx, d, meta)
	}
	return diags
}
func resourceSlbHttp3Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp3Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp3(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbHttp3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbHttp3Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbHttp3(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbHttp3SamplingEnable(d []interface{}) []edpt.SlbHttp3SamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbHttp3SamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbHttp3SamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbHttp3(d *schema.ResourceData) edpt.SlbHttp3 {
	var ret edpt.SlbHttp3
	ret.Inst.SamplingEnable = getSliceSlbHttp3SamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
