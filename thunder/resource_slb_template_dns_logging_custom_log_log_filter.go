package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLoggingCustomLogLogFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging_custom_log_log_filter`: setting log filter (only support request type)\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingCustomLogLogFilterCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingCustomLogLogFilterUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingCustomLogLogFilterRead,
		DeleteContext: resourceSlbTemplateDnsLoggingCustomLogLogFilterDelete,

		Schema: map[string]*schema.Schema{
			"feature": {
				Type: schema.TypeString, Required: true, Description: "'RPZ': log when rpz feature hit;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"trigger_reason": {
				Type: schema.TypeString, Required: true, Description: "TriggerReason",
			},
			"dns_logging_name": {
				Type: schema.TypeString, Required: true, Description: "Dns_logging_name",
			},
		},
	}
}
func resourceSlbTemplateDnsLoggingCustomLogLogFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogLogFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLogLogFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingCustomLogLogFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingCustomLogLogFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogLogFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLogLogFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingCustomLogLogFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingCustomLogLogFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogLogFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLogLogFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingCustomLogLogFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogLogFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLogLogFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsLoggingCustomLogLogFilter(d *schema.ResourceData) edpt.SlbTemplateDnsLoggingCustomLogLogFilter {
	var ret edpt.SlbTemplateDnsLoggingCustomLogLogFilter
	ret.Inst.Feature = d.Get("feature").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.TriggerReason = d.Get("trigger_reason").(string)
	ret.Inst.Dns_logging_name = d.Get("dns_logging_name").(string)
	return ret
}
