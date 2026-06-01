package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_slb_service_tmpl`: Configure template for slb.server.service\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplDelete,

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
						"es_resp_300": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 300",
						},
						"es_resp_400": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 400",
						},
						"es_resp_500": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 500",
						},
						"resp_3xx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 3xx",
						},
						"resp_4xx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 4xx",
						},
						"resp_5xx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 5xx",
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
						"es_resp_300": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 300",
						},
						"es_resp_400": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 400",
						},
						"es_resp_500": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 500",
						},
						"resp_3xx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 3xx",
						},
						"resp_4xx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 4xx",
						},
						"resp_5xx": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Response status 5xx",
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
func resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbServiceTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbServiceTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbServiceTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesSlbServiceTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbServiceTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsInc2830(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsInc2830 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsInc2830
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Es_resp_300 = in["es_resp_300"].(int)
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Resp3xx = in["resp_3xx"].(int)
		ret.Resp4xx = in["resp_4xx"].(int)
		ret.Resp5xx = in["resp_5xx"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate2831(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate2831 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate2831
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Es_resp_300 = in["es_resp_300"].(int)
		ret.Es_resp_400 = in["es_resp_400"].(int)
		ret.Es_resp_500 = in["es_resp_500"].(int)
		ret.Resp3xx = in["resp_3xx"].(int)
		ret.Resp4xx = in["resp_4xx"].(int)
		ret.Resp5xx = in["resp_5xx"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsSeverity2832(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsSeverity2832 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsSeverity2832
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesSlbServiceTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesSlbServiceTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsInc2830(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsRate2831(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesSlbServiceTmplTriggerStatsSeverity2832(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
