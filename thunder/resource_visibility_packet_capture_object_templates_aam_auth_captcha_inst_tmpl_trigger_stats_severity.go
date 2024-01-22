package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_auth_captcha_inst_tmpl_trigger_stats_severity`: Configure triggers severity for system object stats\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityDelete,

		Schema: map[string]*schema.Schema{
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all drop counters (Default disabled)",
			},
			"drop_alert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert drop counters (Default disabled)",
			},
			"drop_critical": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical drop counters (Default disabled)",
			},
			"drop_warning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning drop counters (Default disabled)",
			},
			"error": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all error counters (Default disabled)",
			},
			"error_alert": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all alert error counters (Default disabled)",
			},
			"error_critical": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all critical error counters (Default disabled)",
			},
			"error_warning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable packet capture on all warning error counters (Default disabled)",
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
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthCaptchaInstTmplTriggerStatsSeverity
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.DropAlert = d.Get("drop_alert").(int)
	ret.Inst.DropCritical = d.Get("drop_critical").(int)
	ret.Inst.DropWarning = d.Get("drop_warning").(int)
	ret.Inst.Error = d.Get("error").(int)
	ret.Inst.ErrorAlert = d.Get("error_alert").(int)
	ret.Inst.ErrorCritical = d.Get("error_critical").(int)
	ret.Inst.ErrorWarning = d.Get("error_warning").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
