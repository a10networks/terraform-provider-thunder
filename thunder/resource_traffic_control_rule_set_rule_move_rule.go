package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetRuleMoveRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_traffic_control_rule_set_rule_move_rule`: Move Rule\n\n__PLACEHOLDER__",
		CreateContext: resourceTrafficControlRuleSetRuleMoveRuleCreate,
		UpdateContext: resourceTrafficControlRuleSetRuleMoveRuleUpdate,
		ReadContext:   resourceTrafficControlRuleSetRuleMoveRuleRead,
		DeleteContext: resourceTrafficControlRuleSetRuleMoveRuleDelete,

		Schema: map[string]*schema.Schema{
			"location": {
				Type: schema.TypeString, Optional: true, Default: "bottom", Description: "'top': top; 'before': before; 'after': after; 'bottom': bottom;",
			},
			"target_rule": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"rule_name": {
				Type: schema.TypeString, Required: true, Description: "Rule_name",
			},
			"rule_set_name": {
				Type: schema.TypeString, Required: true, Description: "Rule_set_name",
			},
		},
	}
}
func resourceTrafficControlRuleSetRuleMoveRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleMoveRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleMoveRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlRuleSetRuleMoveRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceTrafficControlRuleSetRuleMoveRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleMoveRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleMoveRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlRuleSetRuleMoveRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceTrafficControlRuleSetRuleMoveRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleMoveRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleMoveRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTrafficControlRuleSetRuleMoveRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleMoveRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleMoveRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTrafficControlRuleSetRuleMoveRule(d *schema.ResourceData) edpt.TrafficControlRuleSetRuleMoveRule {
	var ret edpt.TrafficControlRuleSetRuleMoveRule
	ret.Inst.Location = d.Get("location").(string)
	ret.Inst.TargetRule = d.Get("target_rule").(string)
	ret.Inst.Rule_name = d.Get("rule_name").(string)
	ret.Inst.Rule_set_name = d.Get("rule_set_name").(string)
	return ret
}
