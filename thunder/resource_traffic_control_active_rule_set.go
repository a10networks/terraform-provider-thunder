package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlActiveRuleSet() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_traffic_control_active_rule_set`: Active traffic-control policy\n\n__PLACEHOLDER__",
		CreateContext: resourceTrafficControlActiveRuleSetCreate,
		UpdateContext: resourceTrafficControlActiveRuleSetUpdate,
		ReadContext:   resourceTrafficControlActiveRuleSetRead,
		DeleteContext: resourceTrafficControlActiveRuleSetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Rule set name",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTrafficControlActiveRuleSetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlActiveRuleSetCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlActiveRuleSet(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlActiveRuleSetRead(ctx, d, meta)
	}
	return diags
}

func resourceTrafficControlActiveRuleSetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlActiveRuleSetUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlActiveRuleSet(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlActiveRuleSetRead(ctx, d, meta)
	}
	return diags
}
func resourceTrafficControlActiveRuleSetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlActiveRuleSetDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlActiveRuleSet(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTrafficControlActiveRuleSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlActiveRuleSetRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlActiveRuleSet(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointTrafficControlActiveRuleSet(d *schema.ResourceData) edpt.TrafficControlActiveRuleSet {
	var ret edpt.TrafficControlActiveRuleSet
	ret.Inst.Name = d.Get("name").(string)
	//omit uuid
	return ret
}
