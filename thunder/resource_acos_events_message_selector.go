package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessageSelector() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_message_selector`: Configure message selector to select messages to be logged/blocked\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsMessageSelectorCreate,
		UpdateContext: resourceAcosEventsMessageSelectorUpdate,
		ReadContext:   resourceAcosEventsMessageSelectorRead,
		DeleteContext: resourceAcosEventsMessageSelectorDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify message selector name",
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": {
							Type: schema.TypeInt, Required: true, Description: "Specify rule index - rules are applied in numeric order",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "send", Description: "'send': log messages selected by this rule will be sent; 'drop': log messages selected by this rule will be dropped;",
						},
						"message_id": {
							Type: schema.TypeString, Optional: true, Description: "Select a specific message by message-id and optionally severity",
						},
						"message_id_scope": {
							Type: schema.TypeString, Optional: true, Description: "'all': Log messages at this level and all sub-trees; 'node-only': Log messages at this node only; 'children-only': Log messages at all sub-trees only; 'log-field-only': Log message for this Log Field only;",
						},
						"severity_oper": {
							Type: schema.TypeString, Optional: true, Description: "'equal-and-higher': emergency is highest, debugging lowest; 'equal': single severity;",
						},
						"severity_val": {
							Type: schema.TypeString, Optional: true, Description: "'emergency': System unusable log messages (Most Important); 'alert': Action must be taken immediately; 'critical': Critical conditions; 'error': Error conditions; 'warning': Warning conditions; 'notification': Normal but significant conditions; 'information': Informational messages; 'debugging': Debug level messages (Least Important);",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
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
func resourceAcosEventsMessageSelectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelector(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageSelectorRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsMessageSelectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelector(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageSelectorRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsMessageSelectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelector(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsMessageSelectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelector(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsMessageSelectorRuleList(d []interface{}) []edpt.AcosEventsMessageSelectorRuleList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsMessageSelectorRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsMessageSelectorRuleList
		oi.Index = in["index"].(int)
		oi.Action = in["action"].(string)
		oi.MessageId = in["message_id"].(string)
		oi.MessageIdScope = in["message_id_scope"].(string)
		oi.SeverityOper = in["severity_oper"].(string)
		oi.SeverityVal = in["severity_val"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointAcosEventsMessageSelector(d *schema.ResourceData) edpt.AcosEventsMessageSelector {
	var ret edpt.AcosEventsMessageSelector
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RuleList = getSliceAcosEventsMessageSelectorRuleList(d.Get("rule_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
