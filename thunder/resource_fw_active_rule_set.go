package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwActiveRuleSet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_active_rule_set`: Active firewall policy\n\n__PLACEHOLDER__",
		CreateContext: resourceFwActiveRuleSetCreate,
		UpdateContext: resourceFwActiveRuleSetUpdate,
		ReadContext:   resourceFwActiveRuleSetRead,
		DeleteContext: resourceFwActiveRuleSetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Rule set name",
			},
			"override_nat_aging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override NAT idle-timeout",
			},
			"session_aging": {
				Type: schema.TypeString, Optional: true, Description: "Session Aging Template",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwActiveRuleSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwActiveRuleSetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwActiveRuleSet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwActiveRuleSetRead(ctx, d, meta)
	}
	return diags
}

func resourceFwActiveRuleSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwActiveRuleSetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwActiveRuleSet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwActiveRuleSetRead(ctx, d, meta)
	}
	return diags
}
func resourceFwActiveRuleSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwActiveRuleSetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwActiveRuleSet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwActiveRuleSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwActiveRuleSetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwActiveRuleSet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwActiveRuleSet(d *schema.ResourceData) edpt.FwActiveRuleSet {
	var ret edpt.FwActiveRuleSet
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.OverrideNatAging = d.Get("override_nat_aging").(int)
	ret.Inst.SessionAging = d.Get("session_aging").(string)
	//omit uuid
	return ret
}
