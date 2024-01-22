package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableSiteLicense() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_site_license`: Enable Site License for Azure\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableSiteLicenseCreate,
		UpdateContext: resourceEnableSiteLicenseUpdate,
		ReadContext:   resourceEnableSiteLicenseRead,
		DeleteContext: resourceEnableSiteLicenseDelete,

		Schema: map[string]*schema.Schema{
			"hash_key": {
				Type: schema.TypeString, Required: true, Description: "Input Hash key to enable licensing",
			},
		},
	}
}
func resourceEnableSiteLicenseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableSiteLicenseCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableSiteLicense(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableSiteLicenseRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableSiteLicenseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableSiteLicenseUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableSiteLicense(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableSiteLicenseRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableSiteLicenseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableSiteLicenseDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableSiteLicense(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableSiteLicenseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableSiteLicenseRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableSiteLicense(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointEnableSiteLicense(d *schema.ResourceData) edpt.EnableSiteLicense {
	var ret edpt.EnableSiteLicense
	ret.Inst.HashKey = d.Get("hash_key").(string)
	return ret
}
