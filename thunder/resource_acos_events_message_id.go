package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessageId() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_message_id`: Configure log message\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsMessageIdCreate,
		UpdateContext: resourceAcosEventsMessageIdUpdate,
		ReadContext:   resourceAcosEventsMessageIdRead,
		DeleteContext: resourceAcosEventsMessageIdDelete,

		Schema: map[string]*schema.Schema{
			"log_msg": {
				Type: schema.TypeString, Required: true, Description: "Specify log message-id lineage",
			},
			"message_id_scope_route": {
				Type: schema.TypeString, Required: true, Description: "'all': Log messages at this level and all sub-trees; 'node-only': Log messages at this node only; 'children-only': Log messages at all sub-trees only; 'log-field-only': This is a log-field (Default);",
			},
			"property": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"severity_val": {
										Type: schema.TypeString, Optional: true, Description: "'emergency': System unusable log messages (Most Important); 'alert': Action must be taken immediately; 'critical': Critical conditions; 'error': Error conditions; 'warning': Warning conditions; 'notification': Normal but significant conditions; 'information': Informational messages; 'debugging': Debug level messages (Least Important);",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"log_route": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_route_val": {
										Type: schema.TypeString, Optional: true, Description: "'local-only': send logs to local-only; 'remote-only': send logs to remote-only; 'local-and-remote': send logs to both local and remote;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
						"rate_limit": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rate_limit_val": {
										Type: schema.TypeString, Optional: true, Description: "'enable': enable rate-limiting of logs; 'disable': disable rate-limiting of logs;",
									},
									"uuid": {
										Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
									},
								},
							},
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsMessageIdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageId(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsMessageIdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageId(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageIdRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsMessageIdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageId(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsMessageIdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageIdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageId(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAcosEventsMessageIdProperty56(d []interface{}) edpt.AcosEventsMessageIdProperty56 {

	count1 := len(d)
	var ret edpt.AcosEventsMessageIdProperty56
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Severity = getObjectAcosEventsMessageIdPropertySeverity57(in["severity"].([]interface{}))
		ret.LogRoute = getObjectAcosEventsMessageIdPropertyLogRoute58(in["log_route"].([]interface{}))
		ret.RateLimit = getObjectAcosEventsMessageIdPropertyRateLimit59(in["rate_limit"].([]interface{}))
	}
	return ret
}

func getObjectAcosEventsMessageIdPropertySeverity57(d []interface{}) edpt.AcosEventsMessageIdPropertySeverity57 {

	count1 := len(d)
	var ret edpt.AcosEventsMessageIdPropertySeverity57
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SeverityVal = in["severity_val"].(string)
		//omit uuid
	}
	return ret
}

func getObjectAcosEventsMessageIdPropertyLogRoute58(d []interface{}) edpt.AcosEventsMessageIdPropertyLogRoute58 {

	count1 := len(d)
	var ret edpt.AcosEventsMessageIdPropertyLogRoute58
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogRouteVal = in["log_route_val"].(string)
		//omit uuid
	}
	return ret
}

func getObjectAcosEventsMessageIdPropertyRateLimit59(d []interface{}) edpt.AcosEventsMessageIdPropertyRateLimit59 {

	count1 := len(d)
	var ret edpt.AcosEventsMessageIdPropertyRateLimit59
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RateLimitVal = in["rate_limit_val"].(string)
		//omit uuid
	}
	return ret
}

func dataToEndpointAcosEventsMessageId(d *schema.ResourceData) edpt.AcosEventsMessageId {
	var ret edpt.AcosEventsMessageId
	ret.Inst.LogMsg = d.Get("log_msg").(string)
	ret.Inst.MessageIdScopeRoute = d.Get("message_id_scope_route").(string)
	ret.Inst.Property = getObjectAcosEventsMessageIdProperty56(d.Get("property").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
