package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkIcmpv6RateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_icmpv6_rate_limit`: Limit the rate of ICMPv6 packet\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkIcmpv6RateLimitCreate,
		UpdateContext: resourceNetworkIcmpv6RateLimitUpdate,
		ReadContext:   resourceNetworkIcmpv6RateLimitRead,
		DeleteContext: resourceNetworkIcmpv6RateLimitDelete,

		Schema: map[string]*schema.Schema{
			"icmpv6_lockup": {
				Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
			},
			"icmpv6_lockup_period": {
				Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
			},
			"icmpv6_normal_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkIcmpv6RateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpv6RateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpv6RateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkIcmpv6RateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkIcmpv6RateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpv6RateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpv6RateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkIcmpv6RateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkIcmpv6RateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpv6RateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpv6RateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkIcmpv6RateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpv6RateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpv6RateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkIcmpv6RateLimit(d *schema.ResourceData) edpt.NetworkIcmpv6RateLimit {
	var ret edpt.NetworkIcmpv6RateLimit
	ret.Inst.Icmpv6Lockup = d.Get("icmpv6_lockup").(int)
	ret.Inst.Icmpv6LockupPeriod = d.Get("icmpv6_lockup_period").(int)
	ret.Inst.Icmpv6NormalRateLimit = d.Get("icmpv6_normal_rate_limit").(int)
	//omit uuid
	return ret
}
