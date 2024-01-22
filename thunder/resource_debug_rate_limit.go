package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_rate_limit`: Debug Rate Limiting\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugRateLimitCreate,
		UpdateContext: resourceDebugRateLimitUpdate,
		ReadContext:   resourceDebugRateLimitRead,
		DeleteContext: resourceDebugRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugRateLimit(d *schema.ResourceData) edpt.DebugRateLimit {
	var ret edpt.DebugRateLimit
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
