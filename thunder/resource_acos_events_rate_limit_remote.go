package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsRateLimitRemote() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_rate_limit_remote`: Configure rate limit for logs sent to remote via classic logging config\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsRateLimitRemoteCreate,
		UpdateContext: resourceAcosEventsRateLimitRemoteUpdate,
		ReadContext:   resourceAcosEventsRateLimitRemoteRead,
		DeleteContext: resourceAcosEventsRateLimitRemoteDelete,

		Schema: map[string]*schema.Schema{
			"limit": {
				Type: schema.TypeInt, Optional: true, Default: 32, Description: "Configure rate limit for logs sent to remote via classic logging config",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsRateLimitRemoteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitRemoteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitRemote(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsRateLimitRemoteRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsRateLimitRemoteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitRemoteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitRemote(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsRateLimitRemoteRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsRateLimitRemoteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitRemoteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitRemote(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsRateLimitRemoteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsRateLimitRemoteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsRateLimitRemote(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsRateLimitRemote(d *schema.ResourceData) edpt.AcosEventsRateLimitRemote {
	var ret edpt.AcosEventsRateLimitRemote
	ret.Inst.Limit = d.Get("limit").(int)
	//omit uuid
	return ret
}
