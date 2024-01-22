package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowEventNotificationDdos() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_event_notification_ddos`: Send notification of DDoS events\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowEventNotificationDdosCreate,
		UpdateContext: resourceSflowEventNotificationDdosUpdate,
		ReadContext:   resourceSflowEventNotificationDdosRead,
		DeleteContext: resourceSflowEventNotificationDdosDelete,

		Schema: map[string]*schema.Schema{
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable sflow notification for DDOS events; 'disable': Disable sflow notification for DDOS events;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowEventNotificationDdosCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowEventNotificationDdosCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowEventNotificationDdos(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowEventNotificationDdosRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowEventNotificationDdosUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowEventNotificationDdosUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowEventNotificationDdos(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowEventNotificationDdosRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowEventNotificationDdosDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowEventNotificationDdosDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowEventNotificationDdos(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowEventNotificationDdosRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowEventNotificationDdosRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowEventNotificationDdos(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowEventNotificationDdos(d *schema.ResourceData) edpt.SflowEventNotificationDdos {
	var ret edpt.SflowEventNotificationDdos
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
