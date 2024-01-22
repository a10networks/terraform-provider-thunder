package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogParameterizationMessageSelector() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_log_parameterization_message_selector`: Configure message selector to select messages to be logged/blocked\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLogParameterizationMessageSelectorCreate,
		UpdateContext: resourceAcosEventsLogParameterizationMessageSelectorUpdate,
		ReadContext:   resourceAcosEventsLogParameterizationMessageSelectorRead,
		DeleteContext: resourceAcosEventsLogParameterizationMessageSelectorDelete,

		Schema: map[string]*schema.Schema{
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": {
							Type: schema.TypeInt, Required: true, Description: "Specify rule index - rules are applied in numeric order",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Default: "send", Description: "'send': log messages selected by this rule will be sent (Default); 'drop': log messages selected by this rule will be dropped;",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsLogParameterizationMessageSelectorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelector(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogParameterizationMessageSelectorRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLogParameterizationMessageSelectorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelector(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogParameterizationMessageSelectorRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLogParameterizationMessageSelectorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelector(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLogParameterizationMessageSelectorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelector(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceAcosEventsLogParameterizationMessageSelectorRuleList(d []interface{}) []edpt.AcosEventsLogParameterizationMessageSelectorRuleList {

	count1 := len(d)
	ret := make([]edpt.AcosEventsLogParameterizationMessageSelectorRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AcosEventsLogParameterizationMessageSelectorRuleList
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

func dataToEndpointAcosEventsLogParameterizationMessageSelector(d *schema.ResourceData) edpt.AcosEventsLogParameterizationMessageSelector {
	var ret edpt.AcosEventsLogParameterizationMessageSelector
	ret.Inst.RuleList = getSliceAcosEventsLogParameterizationMessageSelectorRuleList(d.Get("rule_list").([]interface{}))
	//omit uuid
	return ret
}
