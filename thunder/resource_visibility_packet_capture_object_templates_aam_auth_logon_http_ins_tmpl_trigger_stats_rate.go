package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_packet_capture_object_templates_aam_auth_logon_http_ins_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateCreate,
		UpdateContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateUpdate,
		ReadContext:   resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateRead,
		DeleteContext: resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type: schema.TypeInt, Optional: true, Default: 60, Description: "Time in seconds to look for the anomaly, default is 60",
			},
			"spn_krb_faiure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable automatic packet-capture for SPN Kerberos Failure",
			},
			"threshold_exceeded_by": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Set the threshold to the number of times greater than the previous duration to start the capture, default is 5",
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
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate(d *schema.ResourceData) edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate {
	var ret edpt.VisibilityPacketCaptureObjectTemplatesAamAuthLogonHttpInsTmplTriggerStatsRate
	ret.Inst.Duration = d.Get("duration").(int)
	ret.Inst.Spn_krb_faiure = d.Get("spn_krb_faiure").(int)
	ret.Inst.ThresholdExceededBy = d.Get("threshold_exceeded_by").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
