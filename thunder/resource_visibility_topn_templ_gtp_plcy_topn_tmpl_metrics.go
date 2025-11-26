package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnTemplGtpPlcyTopnTmplMetrics() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_templ_gtp_plcy_topn_tmpl_metrics`: Configure topn metrics for template.gtp-policy\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsCreate,
		UpdateContext: resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsUpdate,
		ReadContext:   resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsRead,
		DeleteContext: resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsDelete,

		Schema: map[string]*schema.Schema{
			"rl_message_monitor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP Message forwarded via monitor mode at rate-limit policy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"templ_gtp_plcy_topn_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "Templ_gtp_plcy_topn_tmpl_name",
			},
		},
	}
}
func resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmplMetrics(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmplMetrics(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmplMetrics(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplMetricsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmplMetrics(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmplMetrics(d *schema.ResourceData) edpt.VisibilityTopnTemplGtpPlcyTopnTmplMetrics {
	var ret edpt.VisibilityTopnTemplGtpPlcyTopnTmplMetrics
	ret.Inst.RlMessageMonitor = d.Get("rl_message_monitor").(int)
	//omit uuid
	ret.Inst.Templ_gtp_plcy_topn_tmpl_name = d.Get("templ_gtp_plcy_topn_tmpl_name").(string)
	return ret
}
