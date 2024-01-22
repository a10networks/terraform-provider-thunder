package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rate_limit`: Rate limit configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceRateLimitCreate,
		UpdateContext: resourceRateLimitUpdate,
		ReadContext:   resourceRateLimitRead,
		DeleteContext: resourceRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"maxpktnum": {
				Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Max number of packets",
			},
			"rl_type": {
				Type: schema.TypeString, Optional: true, Description: "'ctrl': The max number of packets that can be sent to kernel in 100ms;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRateLimit(d *schema.ResourceData) edpt.RateLimit {
	var ret edpt.RateLimit
	ret.Inst.Maxpktnum = d.Get("maxpktnum").(int)
	ret.Inst.RlType = d.Get("rl_type").(string)
	//omit uuid
	return ret
}
