package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SctpRateLimitSource() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sctp_rate_limit_source`: Configure SCTP source rate-limit\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SctpRateLimitSourceCreate,
		UpdateContext: resourceCgnv6SctpRateLimitSourceUpdate,
		ReadContext:   resourceCgnv6SctpRateLimitSourceRead,
		DeleteContext: resourceCgnv6SctpRateLimitSourceDelete,

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
func resourceCgnv6SctpRateLimitSourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitSourceCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitSource(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpRateLimitSourceRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SctpRateLimitSourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitSourceUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitSource(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpRateLimitSourceRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SctpRateLimitSourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitSourceDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitSource(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SctpRateLimitSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpRateLimitSourceRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpRateLimitSource(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SctpRateLimitSource(d *schema.ResourceData) edpt.Cgnv6SctpRateLimitSource {
	var ret edpt.Cgnv6SctpRateLimitSource
	ret.Inst.Ip = d.Get("ip").(string)
	ret.Inst.RateLimit = d.Get("rate_limit").(int)
	//omit uuid
	return ret
}
