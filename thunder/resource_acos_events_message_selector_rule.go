package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsMessageSelectorRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_message_selector_rule`: Configure rules to select messages for which logging is enabled/blocked\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsMessageSelectorRuleCreate,
		UpdateContext: resourceAcosEventsMessageSelectorRuleUpdate,
		ReadContext:   resourceAcosEventsMessageSelectorRuleRead,
		DeleteContext: resourceAcosEventsMessageSelectorRuleDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "send", Description: "'send': log messages selected by this rule will be sent; 'drop': log messages selected by this rule will be dropped;",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceAcosEventsMessageSelectorRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelectorRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageSelectorRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsMessageSelectorRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelectorRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsMessageSelectorRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsMessageSelectorRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelectorRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsMessageSelectorRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsMessageSelectorRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsMessageSelectorRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsMessageSelectorRule(d *schema.ResourceData) edpt.AcosEventsMessageSelectorRule {
	var ret edpt.AcosEventsMessageSelectorRule
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.Index = d.Get("index").(int)
	ret.Inst.MessageId = d.Get("message_id").(string)
	ret.Inst.MessageIdScope = d.Get("message_id_scope").(string)
	ret.Inst.SeverityOper = d.Get("severity_oper").(string)
	ret.Inst.SeverityVal = d.Get("severity_val").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
