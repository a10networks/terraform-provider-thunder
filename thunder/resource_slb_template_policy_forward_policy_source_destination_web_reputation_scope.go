package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_policy_forward_policy_source_destination_web_reputation_scope`: Configure web-category reputation-scope for destination matching\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeCreate,
		UpdateContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeUpdate,
		ReadContext:   resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeRead,
		DeleteContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "Action to be performed",
			},
			"dual_stack_action": {
				Type: schema.TypeString, Optional: true, Description: "Dual-stack action to be performed",
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Description: "Priority value of the action(higher the number higher the priority)",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'host': Match hostname; 'url': match URL;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"web_reputation_scope": {
				Type: schema.TypeString, Required: true, Description: "Destination Web Reputation Scope Name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScope
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.DualStackAction = d.Get("dual_stack_action").(string)
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.Type = d.Get("type").(string)
	//omit uuid
	ret.Inst.WebReputationScope = d.Get("web_reputation_scope").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
