package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLoggingCustomLog() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging_custom_log`: setting customizing log\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingCustomLogCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingCustomLogUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingCustomLogRead,
		DeleteContext: resourceSlbTemplateDnsLoggingCustomLogDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable this log",
			},
			"format": {
				Type: schema.TypeString, Optional: true, Description: "Request Message (Custom message string)",
			},
			"trigger_reason": {
				Type: schema.TypeString, Required: true, Description: "'request': log when request comes from client; 'response': log when response to client;",
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
func resourceSlbTemplateDnsLoggingCustomLogCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLog(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingCustomLogRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingCustomLogUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLog(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingCustomLogRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingCustomLogDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLog(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingCustomLogRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCustomLogRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingCustomLog(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsLoggingCustomLog(d *schema.ResourceData) edpt.SlbTemplateDnsLoggingCustomLog {
	var ret edpt.SlbTemplateDnsLoggingCustomLog
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.Format = d.Get("format").(string)
	ret.Inst.TriggerReason = d.Get("trigger_reason").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Dns_logging_name = d.Get("dns_logging_name").(string)
	return ret
}
