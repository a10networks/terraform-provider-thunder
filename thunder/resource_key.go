package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_key`: Authentication key management\n\n__PLACEHOLDER__",
		CreateContext: resourceKeyCreate,
		UpdateContext: resourceKeyUpdate,
		ReadContext:   resourceKeyRead,
		DeleteContext: resourceKeyDelete,

		Schema: map[string]*schema.Schema{
			"key_chain_flag": {
				Type: schema.TypeInt, Required: true, Description: "Key-chain management",
			},
			"key_chain_name": {
				Type: schema.TypeString, Required: true, Description: "Key-chain name",
			},
			"key_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_number": {
							Type: schema.TypeInt, Required: true, Description: "Key identifier number",
						},
						"key_string": {
							Type: schema.TypeString, Optional: true, Description: "Set key string (The key)",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
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
func resourceKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceKeyKeyList(d []interface{}) []edpt.KeyKeyList {

	count1 := len(d)
	ret := make([]edpt.KeyKeyList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.KeyKeyList
		oi.KeyNumber = in["key_number"].(int)
		oi.KeyString = in["key_string"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointKey(d *schema.ResourceData) edpt.Key {
	var ret edpt.Key
	ret.Inst.KeyChainFlag = d.Get("key_chain_flag").(int)
	ret.Inst.KeyChainName = d.Get("key_chain_name").(string)
	ret.Inst.KeyList = getSliceKeyKeyList(d.Get("key_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
