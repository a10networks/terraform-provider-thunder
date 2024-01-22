package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64FragmentationDfBitTransparency() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_fragmentation_df_bit_transparency`: Add an empty IPv6 fragmentation header if IPv4 DF bit is zero (default:disabled)\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64FragmentationDfBitTransparencyCreate,
		UpdateContext: resourceCgnv6Nat64FragmentationDfBitTransparencyUpdate,
		ReadContext:   resourceCgnv6Nat64FragmentationDfBitTransparencyRead,
		DeleteContext: resourceCgnv6Nat64FragmentationDfBitTransparencyDelete,

		Schema: map[string]*schema.Schema{
			"df_bit_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Add an empty IPv6 fragmentation header if IPv4 DF bit is zero;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6Nat64FragmentationDfBitTransparencyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationDfBitTransparencyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationDfBitTransparency(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64FragmentationDfBitTransparencyRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64FragmentationDfBitTransparencyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationDfBitTransparencyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationDfBitTransparency(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64FragmentationDfBitTransparencyRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64FragmentationDfBitTransparencyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationDfBitTransparencyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationDfBitTransparency(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64FragmentationDfBitTransparencyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64FragmentationDfBitTransparencyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64FragmentationDfBitTransparency(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64FragmentationDfBitTransparency(d *schema.ResourceData) edpt.Cgnv6Nat64FragmentationDfBitTransparency {
	var ret edpt.Cgnv6Nat64FragmentationDfBitTransparency
	ret.Inst.DfBitValue = d.Get("df_bit_value").(string)
	//omit uuid
	return ret
}
