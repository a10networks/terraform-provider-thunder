package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKeyKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_key_key`: Configure a key\n\n__PLACEHOLDER__",
		CreateContext: resourceKeyKeyCreate,
		UpdateContext: resourceKeyKeyUpdate,
		ReadContext:   resourceKeyKeyRead,
		DeleteContext: resourceKeyKeyDelete,

		Schema: map[string]*schema.Schema{
			"key_number": {
				Type: schema.TypeInt, Required: true, Description: "Key identifier number",
			},
			"key_string": {
				Type: schema.TypeString, Optional: true, Description: "Set key string (The key)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"key_chain_name": {
				Type: schema.TypeString, Required: true, Description: "KeyChainName",
			},
			"key_chain_flag": {
				Type: schema.TypeString, Required: true, Description: "KeyChainFlag",
			},
		},
	}
}
func resourceKeyKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKeyKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceKeyKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceKeyKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKeyKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceKeyKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceKeyKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKeyKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceKeyKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKeyKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointKeyKey(d *schema.ResourceData) edpt.KeyKey {
	var ret edpt.KeyKey
	ret.Inst.KeyNumber = d.Get("key_number").(int)
	ret.Inst.KeyString = d.Get("key_string").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.KeyChainName = d.Get("key_chain_name").(string)
	ret.Inst.KeyChainFlag = d.Get("key_chain_flag").(string)
	return ret
}
