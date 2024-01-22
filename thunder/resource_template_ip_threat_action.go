package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateIpThreatAction() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_ip_threat_action`: Create an IP Threat Action Template\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateIpThreatActionCreate,
		UpdateContext: resourceTemplateIpThreatActionUpdate,
		ReadContext:   resourceTemplateIpThreatActionRead,
		DeleteContext: resourceTemplateIpThreatActionDelete,

		Schema: map[string]*schema.Schema{
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Idle timeout in minutes(default:5)",
			},
			"log": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable logging; 'disable': Disable logging (Default);",
			},
			"template_number": {
				Type: schema.TypeInt, Required: true, Description: "Template Number",
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
func resourceTemplateIpThreatActionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateIpThreatActionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateIpThreatAction(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateIpThreatActionRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateIpThreatActionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateIpThreatActionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateIpThreatAction(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateIpThreatActionRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateIpThreatActionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateIpThreatActionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateIpThreatAction(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateIpThreatActionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateIpThreatActionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateIpThreatAction(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTemplateIpThreatAction(d *schema.ResourceData) edpt.TemplateIpThreatAction {
	var ret edpt.TemplateIpThreatAction
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.Log = d.Get("log").(string)
	ret.Inst.TemplateNumber = d.Get("template_number").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
