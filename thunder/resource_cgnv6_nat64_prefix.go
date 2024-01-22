package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64Prefix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat64_prefix`: Configure NAT64 Prefix\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat64PrefixCreate,
		UpdateContext: resourceCgnv6Nat64PrefixUpdate,
		ReadContext:   resourceCgnv6Nat64PrefixRead,
		DeleteContext: resourceCgnv6Nat64PrefixDelete,

		Schema: map[string]*schema.Schema{
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Class-list to match for NAT64",
			},
			"prefix_val": {
				Type: schema.TypeString, Required: true, Description: "NAT64 Prefix",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6Nat64PrefixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64PrefixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64Prefix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64PrefixRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat64PrefixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64PrefixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64Prefix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat64PrefixRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat64PrefixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64PrefixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64Prefix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat64PrefixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64PrefixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64Prefix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat64Prefix(d *schema.ResourceData) edpt.Cgnv6Nat64Prefix {
	var ret edpt.Cgnv6Nat64Prefix
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.PrefixVal = d.Get("prefix_val").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
