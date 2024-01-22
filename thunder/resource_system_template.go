package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_template`: Apply template to the whole system\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTemplateCreate,
		UpdateContext: resourceSystemTemplateUpdate,
		ReadContext:   resourceSystemTemplateRead,
		DeleteContext: resourceSystemTemplateDelete,

		Schema: map[string]*schema.Schema{
			"template_policy": {
				Type: schema.TypeString, Optional: true, Description: "Apply policy template to the whole system (Policy template name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTemplate(d *schema.ResourceData) edpt.SystemTemplate {
	var ret edpt.SystemTemplate
	ret.Inst.TemplatePolicy = d.Get("template_policy").(string)
	//omit uuid
	return ret
}
