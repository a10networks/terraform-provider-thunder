package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogParameterizationMessageSelectorRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_log_parameterization_message_selector_rule`: Configure rules to select messages for which logging is enabled/blocked\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLogParameterizationMessageSelectorRuleCreate,
		UpdateContext: resourceAcosEventsLogParameterizationMessageSelectorRuleUpdate,
		ReadContext:   resourceAcosEventsLogParameterizationMessageSelectorRuleRead,
		DeleteContext: resourceAcosEventsLogParameterizationMessageSelectorRuleDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "send", Description: "'send': log messages selected by this rule will be sent (Default); 'drop': log messages selected by this rule will be dropped;",
			},
			"index": {
				Type: schema.TypeInt, Required: true, Description: "Specify rule index - rules are applied in numeric order",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsLogParameterizationMessageSelectorRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelectorRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogParameterizationMessageSelectorRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLogParameterizationMessageSelectorRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelectorRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogParameterizationMessageSelectorRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLogParameterizationMessageSelectorRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelectorRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLogParameterizationMessageSelectorRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogParameterizationMessageSelectorRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogParameterizationMessageSelectorRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsLogParameterizationMessageSelectorRule(d *schema.ResourceData) edpt.AcosEventsLogParameterizationMessageSelectorRule {
	var ret edpt.AcosEventsLogParameterizationMessageSelectorRule
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Index = d.Get("index").(int)
	ret.Inst.MessageId = d.Get("message_id").(string)
	ret.Inst.MessageIdScope = d.Get("message_id_scope").(string)
	ret.Inst.SeverityOper = d.Get("severity_oper").(string)
	ret.Inst.SeverityVal = d.Get("severity_val").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
