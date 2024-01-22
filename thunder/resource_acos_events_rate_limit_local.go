package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsRateLimitLocal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_rate_limit_local`: Configure Rate Limit for Local logs\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsRateLimitLocalCreate,
		UpdateContext: resourceAcosEventsRateLimitLocalUpdate,
		ReadContext:   resourceAcosEventsRateLimitLocalRead,
		DeleteContext: resourceAcosEventsRateLimitLocalDelete,

		Schema: map[string]*schema.Schema{
			"limit": {
				Type: schema.TypeInt, Optional: true, Default: 32, Description: "Configure Rate Limit for Local logs",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsRateLimitLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitLocalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitLocal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsRateLimitLocalRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsRateLimitLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitLocalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitLocal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsRateLimitLocalRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsRateLimitLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitLocalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitLocal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsRateLimitLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitLocalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitLocal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsRateLimitLocal(d *schema.ResourceData) edpt.AcosEventsRateLimitLocal {
	var ret edpt.AcosEventsRateLimitLocal
	ret.Inst.Limit = d.Get("limit").(int)
	//omit uuid
	return ret
}
