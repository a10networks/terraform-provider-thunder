package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLoggingStandardLogLogFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging_standard_log_log_filter`: setting log filter (only support request type)\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingStandardLogLogFilterCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingStandardLogLogFilterUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingStandardLogLogFilterRead,
		DeleteContext: resourceSlbTemplateDnsLoggingStandardLogLogFilterDelete,

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
func resourceSlbTemplateDnsLoggingStandardLogLogFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogLogFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLogLogFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingStandardLogLogFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingStandardLogLogFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogLogFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLogLogFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingStandardLogLogFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingStandardLogLogFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogLogFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLogLogFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingStandardLogLogFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogLogFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLogLogFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsLoggingStandardLogLogFilter(d *schema.ResourceData) edpt.SlbTemplateDnsLoggingStandardLogLogFilter {
	var ret edpt.SlbTemplateDnsLoggingStandardLogLogFilter
	ret.Inst.Feature = d.Get("feature").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.TriggerReason = d.Get("trigger_reason").(string)
	ret.Inst.Dns_logging_name = d.Get("dns_logging_name").(string)
	return ret
}
