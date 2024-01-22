package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkIcmpRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_icmp_rate_limit`: Limit the rate of ICMP packet\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkIcmpRateLimitCreate,
		UpdateContext: resourceNetworkIcmpRateLimitUpdate,
		ReadContext:   resourceNetworkIcmpRateLimitRead,
		DeleteContext: resourceNetworkIcmpRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"icmp_lockup": {
				Type: schema.TypeInt, Optional: true, Description: "Enter lockup state when ICMP rate exceeds lockup rate limit (Maximum rate limit. If exceeds this limit, drop all ICMP packet for a time period)",
			},
			"icmp_lockup_period": {
				Type: schema.TypeInt, Optional: true, Description: "Lockup period (second)",
			},
			"icmp_normal_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Normal rate limit. If exceeds this limit, drop the ICMP packet that goes over the limit",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkIcmpRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkIcmpRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkIcmpRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkIcmpRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkIcmpRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkIcmpRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkIcmpRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkIcmpRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNetworkIcmpRateLimit(d *schema.ResourceData) edpt.NetworkIcmpRateLimit {
	var ret edpt.NetworkIcmpRateLimit
	ret.Inst.IcmpLockup = d.Get("icmp_lockup").(int)
	ret.Inst.IcmpLockupPeriod = d.Get("icmp_lockup_period").(int)
	ret.Inst.IcmpNormalRateLimit = d.Get("icmp_normal_rate_limit").(int)
	//omit uuid
	return ret
}
