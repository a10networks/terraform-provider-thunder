package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetRuleActionGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_traffic_control_rule_set_rule_action_group`: Configure action-group\n\n__PLACEHOLDER__",
		CreateContext: resourceTrafficControlRuleSetRuleActionGroupCreate,
		UpdateContext: resourceTrafficControlRuleSetRuleActionGroupUpdate,
		ReadContext:   resourceTrafficControlRuleSetRuleActionGroupRead,
		DeleteContext: resourceTrafficControlRuleSetRuleActionGroupDelete,

		Schema: map[string]*schema.Schema{
			"limit_policy": {
				Type: schema.TypeInt, Optional: true, Description: "Limit policy Template",
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
func resourceTrafficControlRuleSetRuleActionGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleActionGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleActionGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlRuleSetRuleActionGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceTrafficControlRuleSetRuleActionGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleActionGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleActionGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlRuleSetRuleActionGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceTrafficControlRuleSetRuleActionGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleActionGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleActionGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTrafficControlRuleSetRuleActionGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleActionGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleActionGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTrafficControlRuleSetRuleActionGroup(d *schema.ResourceData) edpt.TrafficControlRuleSetRuleActionGroup {
	var ret edpt.TrafficControlRuleSetRuleActionGroup
	ret.Inst.LimitPolicy = d.Get("limit_policy").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
