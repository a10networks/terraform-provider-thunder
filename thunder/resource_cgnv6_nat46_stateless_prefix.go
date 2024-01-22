package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat46StatelessPrefix() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat46_stateless_prefix`: Enable Stateless NAT46 IPv6 source address\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6Nat46StatelessPrefixCreate,
		UpdateContext: resourceCgnv6Nat46StatelessPrefixUpdate,
		ReadContext:   resourceCgnv6Nat46StatelessPrefixRead,
		DeleteContext: resourceCgnv6Nat46StatelessPrefixDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_prefix": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 prefix",
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
func resourceCgnv6Nat46StatelessPrefixCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPrefixCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPrefix(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessPrefixRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6Nat46StatelessPrefixUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPrefixUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPrefix(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6Nat46StatelessPrefixRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6Nat46StatelessPrefixDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPrefixDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPrefix(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6Nat46StatelessPrefixRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat46StatelessPrefixRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat46StatelessPrefix(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6Nat46StatelessPrefix(d *schema.ResourceData) edpt.Cgnv6Nat46StatelessPrefix {
	var ret edpt.Cgnv6Nat46StatelessPrefix
	ret.Inst.Ipv6Prefix = d.Get("ipv6_prefix").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
