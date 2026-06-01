package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpAuthKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ntp_auth_key`: authentication key\n\n__PLACEHOLDER__",
		CreateContext: resourceNtpAuthKeyCreate,
		UpdateContext: resourceNtpAuthKeyUpdate,
		ReadContext:   resourceNtpAuthKeyRead,
		DeleteContext: resourceNtpAuthKeyDelete,

		Schema: map[string]*schema.Schema{
			"alg_type": {
				Type: schema.TypeString, Optional: true, Description: "'M': encryption using MD5; 'SHA': encryption using SHA; 'SHA1': encryption using SHA1;",
			},
			"asc_key": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"hex_key": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"key": {
				Type: schema.TypeInt, Required: true, Description: "authentication key",
			},
			"key_type": {
				Type: schema.TypeString, Optional: true, Description: "'ascii': key string in ASCII form; 'hex': key string in hex form;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNtpAuthKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpAuthKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpAuthKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpAuthKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceNtpAuthKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpAuthKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpAuthKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpAuthKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceNtpAuthKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpAuthKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpAuthKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNtpAuthKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpAuthKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpAuthKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNtpAuthKey(d *schema.ResourceData) edpt.NtpAuthKey {
	var ret edpt.NtpAuthKey
	ret.Inst.AlgType = d.Get("alg_type").(string)
	ret.Inst.AscKey = d.Get("asc_key").(string)
	//omit encrypted
	//omit hex_encrypted
	ret.Inst.HexKey = d.Get("hex_key").(string)
	ret.Inst.Key = d.Get("key").(int)
	ret.Inst.KeyType = d.Get("key_type").(string)
	//omit uuid
	return ret
}
