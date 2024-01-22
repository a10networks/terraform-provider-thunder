package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessageIdPropertyLogRoute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_message_id_property_log_route`: Configure Log route of log message(s)\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsMessageIdPropertyLogRouteCreate,
		UpdateContext: resourceAcosEventsMessageIdPropertyLogRouteUpdate,
		ReadContext:   resourceAcosEventsMessageIdPropertyLogRouteRead,
		DeleteContext: resourceAcosEventsMessageIdPropertyLogRouteDelete,

		Schema: map[string]*schema.Schema{
			"log_route_val": {
				Type: schema.TypeString, Optional: true, Description: "'local-only': send logs to local-only; 'remote-only': send logs to remote-only; 'local-and-remote': send logs to both local and remote;",
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
func resourceAcosEventsMessageIdPropertyLogRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyLogRouteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyLogRoute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdPropertyLogRouteRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsMessageIdPropertyLogRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyLogRouteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyLogRoute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdPropertyLogRouteRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsMessageIdPropertyLogRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyLogRouteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyLogRoute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsMessageIdPropertyLogRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertyLogRouteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertyLogRoute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsMessageIdPropertyLogRoute(d *schema.ResourceData) edpt.AcosEventsMessageIdPropertyLogRoute {
	var ret edpt.AcosEventsMessageIdPropertyLogRoute
	ret.Inst.LogRouteVal = d.Get("log_route_val").(string)
	//omit uuid
	ret.Inst.LogMsg = d.Get("log_msg").(string)
	ret.Inst.MessageIdScopeRoute = d.Get("message_id_scope_route").(string)
	return ret
}
