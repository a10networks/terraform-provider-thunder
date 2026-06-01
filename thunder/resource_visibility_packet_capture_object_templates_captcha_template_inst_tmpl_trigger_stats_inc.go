package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_captcha_template_inst_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncDelete,

		Schema: map[string]*schema.Schema{
			"json_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Failure JSON Response",
			},
			"other_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Other Error",
			},
			"parse_fail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total JSON Response Parse Failure",
			},
			"timeout_error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for Total Timeout",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"captcha_template_inst_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "Captcha_template_inst_tmpl_name",
			},
		},
	}
}
func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsIncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesCaptchaTemplateInstTmplTriggerStatsInc
	ret.Inst.JsonFail = d.Get("json_fail").(int)
	ret.Inst.OtherError = d.Get("other_error").(int)
	ret.Inst.ParseFail = d.Get("parse_fail").(int)
	ret.Inst.TimeoutError = d.Get("timeout_error").(int)
	//omit uuid
	ret.Inst.Captcha_template_inst_tmpl_name = d.Get("captcha_template_inst_tmpl_name").(string)
	return ret
}
