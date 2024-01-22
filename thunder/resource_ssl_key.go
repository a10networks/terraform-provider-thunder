package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSslKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ssl_key`: ssl certificate and key file information and management commands\n\n__PLACEHOLDER__",
		CreateContext: resourceSslKeyCreate,
		UpdateContext: resourceSslKeyUpdate,
		ReadContext:   resourceSslKeyRead,
		DeleteContext: resourceSslKeyDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "ssl certificate local file name",
			},
			"private_key": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}
func resourceSslKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSslKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceSslKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSslKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceSslKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSslKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSslKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSslKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSslKey(d *schema.ResourceData) edpt.SslKey {
	var ret edpt.SslKey
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PrivateKey = d.Get("private_key").(string)
	return ret
}
