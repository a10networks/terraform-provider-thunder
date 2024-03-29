package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSpdyProxy() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_spdy_proxy`: Configure SPDY Proxy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbSpdyProxyCreate,
		UpdateContext: resourceSlbSpdyProxyUpdate,
		ReadContext:   resourceSlbSpdyProxyRead,
		DeleteContext: resourceSlbSpdyProxyDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'curr_proxy': Curr Proxy Conns; 'total_proxy': Total Proxy Conns; 'curr_http_proxy': Curr HTTP Proxy Conns; 'total_http_proxy': Total HTTP Proxy Conns; 'total_v2_proxy': Version 2 Streams; 'total_v3_proxy': Version 3 Streams; 'curr_stream': Curr Streams; 'total_stream': Total Streams; 'total_stream_succ': Streams(succ); 'client_rst': client_rst; 'server_rst': Server RST sent; 'client_goaway': client_goaway; 'server_goaway': Server GOAWAY sent; 'tcp_err': TCP sock error; 'inflate_ctx': Inflate context; 'deflate_ctx': Deflate context; 'ping_sent': PING sent; 'stream_not_found': STREAM not found; 'client_fin': Client FIN; 'server_fin': Server FIN; 'stream_close': Stream close; 'stream_err': Stream err; 'session_err': Session err; 'control_frame': Control frame received; 'syn_frame': SYN stream frame received; 'syn_reply_frame': SYN reply frame received; 'headers_frame': Headers frame received; 'settings_frame': Setting frame received; 'window_frame': Window update frame received; 'ping_frame': Ping frame received; 'data_frame': Data frame received; 'data_no_stream': Data no stream found; 'data_no_stream_no_goaway': Data no stream and no goaway; 'data_no_stream_goaway_close': Data no stream and no goaway and close session; 'est_cb_no_tuple': Est callback no tuple; 'data_cb_no_tuple': Data callback no tuple; 'ctx_alloc_fail': Context alloc fail; 'fin_close_session': FIN close session; 'server_rst_close_stream': Server RST close stream; 'stream_found': Stream found; 'close_stream_session_not_found': Close stream session not found; 'close_stream_stream_not_found': Close stream stream not found; 'close_stream_already_closed': Closing closed stream; 'close_stream_session_close': Stream close session close; 'close_session_already_closed': Closing closed session; 'max_concurrent_stream_limit': Max concurrent stream limit; 'stream_alloc_fail': Stream alloc fail; 'http_conn_alloc_fail': HTTP connection allocation fail; 'request_header_alloc_fail': Request/Header allocation fail; 'name_value_total_len_ex': Name value total length exceeded; 'name_value_zero_len': Name value zero name length; 'name_value_invalid_http_ver': Name value invalid http version; 'name_value_connection': Name value connection; 'name_value_keepalive': Name value keep alive; 'name_value_proxy_conn': Name value proxy-connection; 'name_value_trasnfer_encod': Name value transfer encoding; 'name_value_no_must_have': Name value no must have; 'decompress_fail': Decompress fail; 'syn_after_goaway': SYN after goaway; 'stream_lt_prev': Stream id less than previous; 'syn_stream_exist_or_even': Stream already exists; 'syn_unidir': Unidirectional SYN; 'syn_reply_alr_rcvd': SYN reply already received; 'client_rst_nostream': Close RST stream not found; 'window_no_stream': Window update no stream found; 'invalid_window_size': Invalid window size; 'unknown_control_frame': Unknown control frame; 'data_on_closed_stream': Data on closed stream; 'invalid_frame_size': Invalid frame size; 'invalid_version': Invalid version; 'header_after_session_close': Header after session close; 'compress_ctx_alloc_fail': Compression context allocation fail; 'header_compress_fail': Header compress fail; 'http_data_session_close': HTTP data session close; 'http_data_stream_not_found': HTTP data stream not found; 'close_stream_not_http_proxy': Close Stream not http-proxy; 'session_needs_requeue': Session needs requeue; 'new_stream_session_del': New Stream after Session delete; 'fin_stream_closed': HTTP FIN stream already closed; 'http_close_stream_closed': HTTP close stream already closed; 'http_err_stream_closed': HTTP error stream already closed; 'http_hdr_stream_close': HTTP header stream already closed; 'http_data_stream_close': HTTP data stream already closed; 'session_close': Session close;",
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
func resourceSlbSpdyProxyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSpdyProxyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSpdyProxy(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSpdyProxyRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbSpdyProxyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSpdyProxyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSpdyProxy(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbSpdyProxyRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbSpdyProxyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSpdyProxyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSpdyProxy(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbSpdyProxyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSpdyProxyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSpdyProxy(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbSpdyProxySamplingEnable(d []interface{}) []edpt.SlbSpdyProxySamplingEnable {

	count1 := len(d)
	ret := make([]edpt.SlbSpdyProxySamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbSpdyProxySamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbSpdyProxy(d *schema.ResourceData) edpt.SlbSpdyProxy {
	var ret edpt.SlbSpdyProxy
	ret.Inst.SamplingEnable = getSliceSlbSpdyProxySamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
