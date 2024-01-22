package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_aaa_policy_tmpl_trigger_stats_severity`: Configure triggers severity for system object stats\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityDelete,

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
func resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAaaPolicyTmplTriggerStatsSeverity
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
