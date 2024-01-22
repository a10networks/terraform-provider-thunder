package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6NatIcmpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_nat_icmpv6`: ICMPv6 configuration for IPv6 NAT\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6NatIcmpv6Create,
		UpdateContext: resourceCgnv6NatIcmpv6Update,
		ReadContext:   resourceCgnv6NatIcmpv6Read,
		DeleteContext: resourceCgnv6NatIcmpv6Delete,

		Schema: map[string]*schema.Schema{
			"respond_to_ping": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Respond to ICMPv6 echo requests to NAT pool IPs (default: disabled)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6NatIcmpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatIcmpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6NatIcmpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6NatIcmpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6NatIcmpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6NatIcmpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6NatIcmpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6NatIcmpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6NatIcmpv6(d *schema.ResourceData) edpt.Cgnv6NatIcmpv6 {
	var ret edpt.Cgnv6NatIcmpv6
	ret.Inst.RespondToPing = d.Get("respond_to_ping").(int)
	//omit uuid
	return ret
}
