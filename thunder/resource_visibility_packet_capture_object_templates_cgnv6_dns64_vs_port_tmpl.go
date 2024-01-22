package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_cgnv6_dns64_vs_port_tmpl`: Configure template for cgnv6.dns64-virtualserver.port\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplDelete,

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
						"es_total_failure_actions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
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
						"es_total_failure_actions": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total failure actions",
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
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc2666(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc2666 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc2666
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsRate2667(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsRate2667 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsRate2667
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.Es_total_failure_actions = in["es_total_failure_actions"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity2668(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity2668 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity2668
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsInc2666(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsRate2667(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCgnv6Dns64VsPortTmplTriggerStatsSeverity2668(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
