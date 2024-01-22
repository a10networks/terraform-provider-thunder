package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEnableBgpAdvertisement() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_enable_bgp_advertisement`: Enable BGP Advertisement\n\n__PLACEHOLDER__",
		CreateContext: resourceEnableBgpAdvertisementCreate,
		UpdateContext: resourceEnableBgpAdvertisementUpdate,
		ReadContext:   resourceEnableBgpAdvertisementRead,
		DeleteContext: resourceEnableBgpAdvertisementDelete,

		Schema: map[string]*schema.Schema{
			"enable_bgp_advertisement": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable BGP Advertisement",
			},
		},
	}
}
func resourceEnableBgpAdvertisementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableBgpAdvertisementCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableBgpAdvertisement(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableBgpAdvertisementRead(ctx, d, meta)
	}
	return diags
}

func resourceEnableBgpAdvertisementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableBgpAdvertisementUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableBgpAdvertisement(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceEnableBgpAdvertisementRead(ctx, d, meta)
	}
	return diags
}
func resourceEnableBgpAdvertisementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableBgpAdvertisementDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableBgpAdvertisement(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceEnableBgpAdvertisementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEnableBgpAdvertisementRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEnableBgpAdvertisement(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointEnableBgpAdvertisement(d *schema.ResourceData) edpt.EnableBgpAdvertisement {
	var ret edpt.EnableBgpAdvertisement
	ret.Inst.EnableBgpAdvertisement = d.Get("enable_bgp_advertisement").(int)
	return ret
}
