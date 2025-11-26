package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTemplateBindMonitor() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_template_bind_monitor`: Apply monitor template to the whole system\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTemplateBindMonitorCreate,
		UpdateContext: resourceSystemTemplateBindMonitorUpdate,
		ReadContext:   resourceSystemTemplateBindMonitorRead,
		DeleteContext: resourceSystemTemplateBindMonitorDelete,

		Schema: map[string]*schema.Schema{
			"template_monitor": {
				Type: schema.TypeInt, Required: true, Description: "Monitor template ID Number",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTemplateBindMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateBindMonitorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplateBindMonitor(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTemplateBindMonitorRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTemplateBindMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateBindMonitorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplateBindMonitor(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTemplateBindMonitorRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTemplateBindMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateBindMonitorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplateBindMonitor(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTemplateBindMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateBindMonitorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplateBindMonitor(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTemplateBindMonitor(d *schema.ResourceData) edpt.SystemTemplateBindMonitor {
	var ret edpt.SystemTemplateBindMonitor
	ret.Inst.TemplateMonitor = d.Get("template_monitor").(int)
	//omit uuid
	return ret
}
