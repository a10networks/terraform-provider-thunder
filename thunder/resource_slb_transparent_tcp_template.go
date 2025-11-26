package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTransparentTcpTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_transparent_tcp_template`: TCP setting for transparent TCP sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTransparentTcpTemplateCreate,
		UpdateContext: resourceSlbTransparentTcpTemplateUpdate,
		ReadContext:   resourceSlbTransparentTcpTemplateRead,
		DeleteContext: resourceSlbTransparentTcpTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "Specify template name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbTransparentTcpTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTransparentTcpTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTransparentTcpTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTransparentTcpTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTransparentTcpTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTransparentTcpTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTransparentTcpTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTransparentTcpTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTransparentTcpTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTransparentTcpTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTransparentTcpTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTransparentTcpTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTransparentTcpTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTransparentTcpTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTransparentTcpTemplate(d *schema.ResourceData) edpt.SlbTransparentTcpTemplate {
	var ret edpt.SlbTransparentTcpTemplate
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
