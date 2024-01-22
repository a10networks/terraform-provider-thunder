package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpTrustedKey() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ntp_trusted_key`: trusted key\n\n__PLACEHOLDER__",
		CreateContext: resourceNtpTrustedKeyCreate,
		UpdateContext: resourceNtpTrustedKeyUpdate,
		ReadContext:   resourceNtpTrustedKeyRead,
		DeleteContext: resourceNtpTrustedKeyDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type: schema.TypeInt, Required: true, Description: "trusted key",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNtpTrustedKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpTrustedKeyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpTrustedKey(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpTrustedKeyRead(ctx, d, meta)
	}
	return diags
}

func resourceNtpTrustedKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpTrustedKeyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpTrustedKey(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpTrustedKeyRead(ctx, d, meta)
	}
	return diags
}
func resourceNtpTrustedKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpTrustedKeyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpTrustedKey(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNtpTrustedKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpTrustedKeyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpTrustedKey(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNtpTrustedKey(d *schema.ResourceData) edpt.NtpTrustedKey {
	var ret edpt.NtpTrustedKey
	ret.Inst.Key = d.Get("key").(int)
	//omit uuid
	return ret
}
