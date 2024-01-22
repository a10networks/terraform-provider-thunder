package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SctpRateLimitDestination() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sctp_rate_limit_destination`: Configure SCTP destination rate-limit\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SctpRateLimitDestinationCreate,
		UpdateContext: resourceCgnv6SctpRateLimitDestinationUpdate,
		ReadContext:   resourceCgnv6SctpRateLimitDestinationRead,
		DeleteContext: resourceCgnv6SctpRateLimitDestinationDelete,

		Schema: map[string]*schema.Schema{
			"ip": {
				Type: schema.TypeString, Required: true, Description: "IP address",
			},
			"rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Rate limit in packets per second",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SctpRateLimitDestinationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitDestinationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitDestination(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpRateLimitDestinationRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SctpRateLimitDestinationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitDestinationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitDestination(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpRateLimitDestinationRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SctpRateLimitDestinationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitDestinationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitDestination(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SctpRateLimitDestinationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitDestinationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitDestination(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SctpRateLimitDestination(d *schema.ResourceData) edpt.Cgnv6SctpRateLimitDestination {
	var ret edpt.Cgnv6SctpRateLimitDestination
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.RateLimit = d.Get("rate_limit").(int)
	//omit uuid
	return ret
}
