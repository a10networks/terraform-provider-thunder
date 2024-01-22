package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateServerSsh() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_server_ssh`: Server Side SSH Template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateServerSshCreate,
		UpdateContext: resourceSlbTemplateServerSshUpdate,
		ReadContext:   resourceSlbTemplateServerSshRead,
		DeleteContext: resourceSlbTemplateServerSshDelete,

		Schema: map[string]*schema.Schema{
			"forward_proxy_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSH forward proxy",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Server SSH Template Name",
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
func resourceSlbTemplateServerSshCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSshCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsh(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerSshRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateServerSshUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSshUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsh(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateServerSshRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateServerSshDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSshDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsh(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateServerSshRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateServerSshRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateServerSsh(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateServerSsh(d *schema.ResourceData) edpt.SlbTemplateServerSsh {
	var ret edpt.SlbTemplateServerSsh
	ret.Inst.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
