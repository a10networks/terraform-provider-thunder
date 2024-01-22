package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_slb_templ_cache_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"content_toobig": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheContentToobig, help content_toobig",
			},
			"content_toosmall": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheContentToosmall, help content_toosmall",
			},
			"entry_create_failures": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheEntryCreateFailures, help entry_create_failures",
			},
			"header_save_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for header_save_error",
			},
			"nc_req_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcReqHeader, help nc_req_header",
			},
			"nc_res_header": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcResHeader, help nc_res_header",
			},
			"rv_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheRvFailure, help rv_failure",
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
func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc
	ret.Inst.Content_toobig = d.Get("content_toobig").(int)
	ret.Inst.Content_toosmall = d.Get("content_toosmall").(int)
	ret.Inst.Entry_create_failures = d.Get("entry_create_failures").(int)
	ret.Inst.Header_save_error = d.Get("header_save_error").(int)
	ret.Inst.Nc_req_header = d.Get("nc_req_header").(int)
	ret.Inst.Nc_res_header = d.Get("nc_res_header").(int)
	ret.Inst.Rv_failure = d.Get("rv_failure").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
