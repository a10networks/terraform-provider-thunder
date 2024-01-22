package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessageIdPropertyRateLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_message_id_property_rate_limit`: Enable/disable rate limiting of log message(s)\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsMessageIdPropertyRateLimitCreate,
		UpdateContext: resourceAcosEventsMessageIdPropertyRateLimitUpdate,
		ReadContext:   resourceAcosEventsMessageIdPropertyRateLimitRead,
		DeleteContext: resourceAcosEventsMessageIdPropertyRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"rate_limit_val": {
				Type: schema.TypeString, Optional: true, Description: "'enable': enable rate-limiting of logs; 'disable': disable rate-limiting of logs;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"log_msg": {
				Type: schema.TypeString, Required: true, Description: "LogMsg",
			},
			"message_id_scope_route": {
				Type: schema.TypeString, Required: true, Description: "MessageIdScopeRoute",
			},
		},
	}
}
func resourceAcosEventsMessageIdPropertyRateLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyRateLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyRateLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdPropertyRateLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsMessageIdPropertyRateLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyRateLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyRateLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdPropertyRateLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsMessageIdPropertyRateLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyRateLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyRateLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsMessageIdPropertyRateLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyRateLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyRateLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsMessageIdPropertyRateLimit(d *schema.ResourceData) edpt.AcosEventsMessageIdPropertyRateLimit {
	var ret edpt.AcosEventsMessageIdPropertyRateLimit
	ret.Inst.RateLimitVal = d.Get("rate_limit_val").(string)
	//omit uuid
	ret.Inst.LogMsg = d.Get("log_msg").(string)
	ret.Inst.MessageIdScopeRoute = d.Get("message_id_scope_route").(string)
	return ret
}
