package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateClientSsh() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_client_ssh`: Client Side SSH Template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateClientSshCreate,
		UpdateContext: resourceSlbTemplateClientSshUpdate,
		ReadContext:   resourceSlbTemplateClientSshRead,
		DeleteContext: resourceSlbTemplateClientSshDelete,

		Schema: map[string]*schema.Schema{
			"forward_proxy_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSH forward proxy",
			},
			"forward_proxy_hostkey": {
				Type: schema.TypeString, Optional: true, Description: "Specify private-key (Key Name)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Client SSH Template Name",
			},
			"passphrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
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
func resourceSlbTemplateClientSshCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSshCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsh(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateClientSshRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateClientSshUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSshUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsh(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateClientSshRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateClientSshDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSshDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsh(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateClientSshRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateClientSshRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateClientSsh(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateClientSsh(d *schema.ResourceData) edpt.SlbTemplateClientSsh {
	var ret edpt.SlbTemplateClientSsh
	//omit encrypted
	ret.Inst.ForwardProxyEnable = d.Get("forward_proxy_enable").(int)
	ret.Inst.ForwardProxyHostkey = d.Get("forward_proxy_hostkey").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Passphrase = d.Get("passphrase").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
