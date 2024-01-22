package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_slb_templ_cache_tmpl`: Configure template for slb.template.cache\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplDelete,

		Schema: map[string]*schema.Schema{
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Specify name of the capture-config to use with this template",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Packet Capture Template Name",
			},
			"trigger_stats_inc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nc_req_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcReqHeader, help nc_req_header",
						},
						"nc_res_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcResHeader, help nc_res_header",
						},
						"rv_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheRvFailure, help rv_failure",
						},
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
						"nc_req_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcReqHeader, help nc_req_header",
						},
						"nc_res_header": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheNcResHeader, help nc_res_header",
						},
						"rv_failure": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for slbTemplateCacheRvFailure, help rv_failure",
						},
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"trigger_stats_severity": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
						},
						"error_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
						},
						"error_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
						},
						"error_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
						},
						"drop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
						},
						"drop_alert": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
						},
						"drop_warning": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
						},
						"drop_critical": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc2711(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc2711 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc2711
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Nc_req_header = in["nc_req_header"].(int)
		ret.Nc_res_header = in["nc_res_header"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		ret.Header_save_error = in["header_save_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate2712(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate2712 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate2712
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Nc_req_header = in["nc_req_header"].(int)
		ret.Nc_res_header = in["nc_res_header"].(int)
		ret.Rv_failure = in["rv_failure"].(int)
		ret.Content_toobig = in["content_toobig"].(int)
		ret.Content_toosmall = in["content_toosmall"].(int)
		ret.Entry_create_failures = in["entry_create_failures"].(int)
		ret.Header_save_error = in["header_save_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsSeverity2713(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsSeverity2713 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsSeverity2713
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Error = in["error"].(int)
		ret.ErrorAlert = in["error_alert"].(int)
		ret.ErrorWarning = in["error_warning"].(int)
		ret.ErrorCritical = in["error_critical"].(int)
		ret.Drop = in["drop"].(int)
		ret.DropAlert = in["drop_alert"].(int)
		ret.DropWarning = in["drop_warning"].(int)
		ret.DropCritical = in["drop_critical"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsInc2711(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsRate2712(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesSlbTemplCacheTmplTriggerStatsSeverity2713(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
