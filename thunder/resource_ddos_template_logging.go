package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_logging`: Logging template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateLoggingCreate,
		UpdateContext: resourceDdosTemplateLoggingUpdate,
		ReadContext:   resourceDdosTemplateLoggingRead,
		DeleteContext: resourceDdosTemplateLoggingDelete,

		Schema: map[string]*schema.Schema{
			"enable_action_logging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log action taken",
			},
			"log_format_cef": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log in CEF format",
			},
			"log_format_custom": {
				Type: schema.TypeString, Optional: true, Description: "Customize log format",
			},
			"logging_tmpl_name": {
				Type: schema.TypeString, Required: true, Description: "DDOS Logging Template Name",
			},
			"use_obj_name": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show obj name instead of ip in the log",
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
func resourceDdosTemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosTemplateLogging(d *schema.ResourceData) edpt.DdosTemplateLogging {
	var ret edpt.DdosTemplateLogging
	ret.Inst.EnableActionLogging = d.Get("enable_action_logging").(int)
	ret.Inst.LogFormatCef = d.Get("log_format_cef").(int)
	ret.Inst.LogFormatCustom = d.Get("log_format_custom").(string)
	ret.Inst.LoggingTmplName = d.Get("logging_tmpl_name").(string)
	ret.Inst.UseObjName = d.Get("use_obj_name").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
