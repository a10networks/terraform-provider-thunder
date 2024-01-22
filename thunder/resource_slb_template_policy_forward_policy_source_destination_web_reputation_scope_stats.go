package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_source_destination_web_reputation_scope_stats`: Statistics for the object web-reputation-scope\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests matching this destination rule",
						},
					},
				},
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

func resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats := setObjectSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStats

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeStatsStats(d.Get("stats").([]interface{}))

	ret.WebReputationScope = d.Get("web_reputation_scope").(string)

	ret.Name = d.Get("name").(string)
	return ret
}
