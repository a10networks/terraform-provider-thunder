package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessageIdPropertySeverity() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_message_id_property_severity`: Configure Serverity of log message(s)\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsMessageIdPropertySeverityCreate,
		UpdateContext: resourceAcosEventsMessageIdPropertySeverityUpdate,
		ReadContext:   resourceAcosEventsMessageIdPropertySeverityRead,
		DeleteContext: resourceAcosEventsMessageIdPropertySeverityDelete,

		Schema: map[string]*schema.Schema{
			"severity_val": {
				Type: schema.TypeString, Optional: true, Description: "'emergency': System unusable log messages (Most Important); 'alert': Action must be taken immediately; 'critical': Critical conditions; 'error': Error conditions; 'warning': Warning conditions; 'notification': Normal but significant conditions; 'information': Informational messages; 'debugging': Debug level messages (Least Important);",
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
func resourceAcosEventsMessageIdPropertySeverityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertySeverityCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertySeverity(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdPropertySeverityRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsMessageIdPropertySeverityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertySeverityUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertySeverity(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdPropertySeverityRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsMessageIdPropertySeverityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertySeverityDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertySeverity(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsMessageIdPropertySeverityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdPropertySeverityRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageIdPropertySeverity(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsMessageIdPropertySeverity(d *schema.ResourceData) edpt.AcosEventsMessageIdPropertySeverity {
	var ret edpt.AcosEventsMessageIdPropertySeverity
	ret.Inst.SeverityVal = d.Get("severity_val").(string)
	//omit uuid
	ret.Inst.LogMsg = d.Get("log_msg").(string)
	ret.Inst.MessageIdScopeRoute = d.Get("message_id_scope_route").(string)
	return ret
}
