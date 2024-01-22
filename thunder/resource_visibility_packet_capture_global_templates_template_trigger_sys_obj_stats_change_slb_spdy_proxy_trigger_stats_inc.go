package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_slb_spdy_proxy_trigger_stats_inc`: Configure stats to trigger packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"compress_ctx_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Compression context allocation fail",
			},
			"ctx_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Context alloc fail",
			},
			"data_cb_no_tuple": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data callback no tuple",
			},
			"data_no_stream": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream found",
			},
			"data_no_stream_goaway_close": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream and no goaway and close session",
			},
			"data_no_stream_no_goaway": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Data no stream and no goaway",
			},
			"decompress_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Decompress fail",
			},
			"est_cb_no_tuple": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Est callback no tuple",
			},
			"header_compress_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Header compress fail",
			},
			"http_conn_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP connection allocation fail",
			},
			"http_err_stream_closed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for HTTP error stream already closed",
			},
			"invalid_frame_size": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid frame size",
			},
			"invalid_version": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Invalid version",
			},
			"request_header_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Request/Header allocation fail",
			},
			"session_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Session err",
			},
			"stream_alloc_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream alloc fail",
			},
			"stream_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Stream err",
			},
			"stream_not_found": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for STREAM not found",
			},
			"tcp_err": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for TCP sock error",
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
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbSpdyProxyTriggerStatsInc
	ret.Inst.Compress_ctx_alloc_fail = d.Get("compress_ctx_alloc_fail").(int)
	ret.Inst.Ctx_alloc_fail = d.Get("ctx_alloc_fail").(int)
	ret.Inst.Data_cb_no_tuple = d.Get("data_cb_no_tuple").(int)
	ret.Inst.Data_no_stream = d.Get("data_no_stream").(int)
	ret.Inst.Data_no_stream_goaway_close = d.Get("data_no_stream_goaway_close").(int)
	ret.Inst.Data_no_stream_no_goaway = d.Get("data_no_stream_no_goaway").(int)
	ret.Inst.Decompress_fail = d.Get("decompress_fail").(int)
	ret.Inst.Est_cb_no_tuple = d.Get("est_cb_no_tuple").(int)
	ret.Inst.Header_compress_fail = d.Get("header_compress_fail").(int)
	ret.Inst.Http_conn_alloc_fail = d.Get("http_conn_alloc_fail").(int)
	ret.Inst.Http_err_stream_closed = d.Get("http_err_stream_closed").(int)
	ret.Inst.Invalid_frame_size = d.Get("invalid_frame_size").(int)
	ret.Inst.Invalid_version = d.Get("invalid_version").(int)
	ret.Inst.Request_header_alloc_fail = d.Get("request_header_alloc_fail").(int)
	ret.Inst.Session_err = d.Get("session_err").(int)
	ret.Inst.Stream_alloc_fail = d.Get("stream_alloc_fail").(int)
	ret.Inst.Stream_err = d.Get("stream_err").(int)
	ret.Inst.Stream_not_found = d.Get("stream_not_found").(int)
	ret.Inst.Tcp_err = d.Get("tcp_err").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
