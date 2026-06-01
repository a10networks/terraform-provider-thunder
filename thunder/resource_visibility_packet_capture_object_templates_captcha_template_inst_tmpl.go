package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_captcha_template_inst_tmpl`: Configure template for slb.template.captcha.instance\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplDelete,

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
						"parse_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total JSON Response Parse Failure",
						},
						"json_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
						},
						"timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
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
						"parse_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total JSON Response Parse Failure",
						},
						"json_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
						},
						"timeout_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
						},
						"other_error": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
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
func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc2779(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc2779 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc2779
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate2780(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate2780 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate2780
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ThresholdExceededBy = in["threshold_exceeded_by"].(int)
		ret.Duration = in["duration"].(int)
		ret.ParseFail = in["parse_fail"].(int)
		ret.JsonFail = in["json_fail"].(int)
		ret.TimeoutError = in["timeout_error"].(int)
		ret.OtherError = in["other_error"].(int)
		//omit uuid
	}
	return ret
}

func getObjectVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsSeverity2781(d []interface{}) edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsSeverity2781 {

	count1 := len(d)
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsSeverity2781
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

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmpl
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TriggerStatsInc = getObjectVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc2779(d.Get("trigger_stats_inc").([]interface{}))
	ret.Inst.TriggerStatsRate = getObjectVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsRate2780(d.Get("trigger_stats_rate").([]interface{}))
	ret.Inst.TriggerStatsSeverity = getObjectVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsSeverity2781(d.Get("trigger_stats_severity").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
