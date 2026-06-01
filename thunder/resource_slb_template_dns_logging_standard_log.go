package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLoggingStandardLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging_standard_log`: setting standard log\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingStandardLogCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingStandardLogUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingStandardLogRead,
		DeleteContext: resourceSlbTemplateDnsLoggingStandardLogDelete,

		Schema: map[string]*schema.Schema{
			"log_filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature": {
							Type: schema.TypeString, Required: true, Description: "'RPZ': log when rpz feature hit;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"trigger_reason": {
				Type: schema.TypeString, Required: true, Description: "'request': log when request comes from client;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dns_logging_name": {
				Type: schema.TypeString, Required: true, Description: "Dns_logging_name",
			},
		},
	}
}
func resourceSlbTemplateDnsLoggingStandardLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingStandardLogRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingStandardLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingStandardLogRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingStandardLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingStandardLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingStandardLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingStandardLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsLoggingStandardLogLogFilterList(d []interface{}) []edpt.SlbTemplateDnsLoggingStandardLogLogFilterList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLoggingStandardLogLogFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLoggingStandardLogLogFilterList
		oi.Feature = in["feature"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsLoggingStandardLog(d *schema.ResourceData) edpt.SlbTemplateDnsLoggingStandardLog {
	var ret edpt.SlbTemplateDnsLoggingStandardLog
	ret.Inst.LogFilterList = getSliceSlbTemplateDnsLoggingStandardLogLogFilterList(d.Get("log_filter_list").([]interface{}))
	ret.Inst.TriggerReason = d.Get("trigger_reason").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Dns_logging_name = d.Get("dns_logging_name").(string)
	return ret
}
