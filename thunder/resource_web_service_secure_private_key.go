package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebServiceSecurePrivateKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_service_secure_private_key`: Web-service secure private-key operation\n\n__PLACEHOLDER__",
		CreateContext: resourceWebServiceSecurePrivateKeyCreate,
		UpdateContext: resourceWebServiceSecurePrivateKeyUpdate,
		ReadContext:   resourceWebServiceSecurePrivateKeyRead,
		DeleteContext: resourceWebServiceSecurePrivateKeyDelete,

		Schema: map[string]*schema.Schema{
			"file_url": {
				Type: schema.TypeString, Optional: true, Description: "File URL",
			},
			"load": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load WEB private-key",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
		},
	}
}
func resourceWebServiceSecurePrivateKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecurePrivateKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecurePrivateKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecurePrivateKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceWebServiceSecurePrivateKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecurePrivateKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecurePrivateKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecurePrivateKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceWebServiceSecurePrivateKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecurePrivateKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecurePrivateKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebServiceSecurePrivateKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecurePrivateKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecurePrivateKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointWebServiceSecurePrivateKey(d *schema.ResourceData) edpt.WebServiceSecurePrivateKey {
	var ret edpt.WebServiceSecurePrivateKey
	ret.Inst.FileUrl = d.Get("file_url").(string)
	ret.Inst.Load = d.Get("load").(int)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	return ret
}
