package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_logging`: Logging template Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateLoggingCreate,
		UpdateContext: resourceDdosZoneTemplateLoggingUpdate,
		ReadContext:   resourceDdosZoneTemplateLoggingRead,
		DeleteContext: resourceDdosZoneTemplateLoggingDelete,

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
func resourceDdosZoneTemplateLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosZoneTemplateLogging(d *schema.ResourceData) edpt.DdosZoneTemplateLogging {
	var ret edpt.DdosZoneTemplateLogging
	ret.Inst.EnableActionLogging = d.Get("enable_action_logging").(int)
	ret.Inst.LogFormatCef = d.Get("log_format_cef").(int)
	ret.Inst.LogFormatCustom = d.Get("log_format_custom").(string)
	ret.Inst.LoggingTmplName = d.Get("logging_tmpl_name").(string)
	ret.Inst.UseObjName = d.Get("use_obj_name").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
