package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSslSetCompatibleCipher() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ssl_set_compatible_cipher`: To Enable/disable setting common cipher suite for SSL connections in management plane\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemSslSetCompatibleCipherCreate,
		UpdateContext: resourceSystemSslSetCompatibleCipherUpdate,
		ReadContext:   resourceSystemSslSetCompatibleCipherRead,
		DeleteContext: resourceSystemSslSetCompatibleCipherDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable setting common cipher suite in management plane",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemSslSetCompatibleCipherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslSetCompatibleCipherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslSetCompatibleCipher(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslSetCompatibleCipherRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemSslSetCompatibleCipherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslSetCompatibleCipherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslSetCompatibleCipher(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemSslSetCompatibleCipherRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemSslSetCompatibleCipherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslSetCompatibleCipherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslSetCompatibleCipher(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemSslSetCompatibleCipherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemSslSetCompatibleCipherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemSslSetCompatibleCipher(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemSslSetCompatibleCipher(d *schema.ResourceData) edpt.SystemSslSetCompatibleCipher {
	var ret edpt.SystemSslSetCompatibleCipher
	ret.Inst.Disable = d.Get("disable").(int)
	//omit uuid
	return ret
}
