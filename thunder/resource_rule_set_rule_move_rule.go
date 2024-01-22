package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRuleMoveRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rule_set_rule_move_rule`: Move Rule\n\n__PLACEHOLDER__",
		CreateContext: resourceRuleSetRuleMoveRuleCreate,
		UpdateContext: resourceRuleSetRuleMoveRuleUpdate,
		ReadContext:   resourceRuleSetRuleMoveRuleRead,
		DeleteContext: resourceRuleSetRuleMoveRuleDelete,

		Schema: map[string]*schema.Schema{
			"location": {
				Type: schema.TypeString, Optional: true, Default: "bottom", Description: "'top': top; 'before': before; 'after': after; 'bottom': bottom;",
			},
			"target_rule": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceRuleSetRuleMoveRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleMoveRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleMoveRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRuleMoveRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceRuleSetRuleMoveRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleMoveRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleMoveRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRuleMoveRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceRuleSetRuleMoveRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleMoveRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleMoveRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRuleSetRuleMoveRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleMoveRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleMoveRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRuleSetRuleMoveRule(d *schema.ResourceData) edpt.RuleSetRuleMoveRule {
	var ret edpt.RuleSetRuleMoveRule
	ret.Inst.Location = d.Get("location").(string)
	ret.Inst.TargetRule = d.Get("target_rule").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
