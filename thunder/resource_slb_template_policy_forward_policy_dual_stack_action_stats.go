package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicyDualStackActionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_dual_stack_action_stats`: Statistics for the object dual-stack-action\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicyDualStackActionStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Action name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests forward by this action",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicyDualStackActionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyDualStackActionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyDualStackActionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicyDualStackActionStatsStats := setObjectSlbTemplatePolicyForwardPolicyDualStackActionStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicyDualStackActionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicyDualStackActionStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicyDualStackActionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicyDualStackActionStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicyDualStackActionStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicyDualStackActionStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicyDualStackActionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicyDualStackActionStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicyDualStackActionStats {
	var ret edpt.SlbTemplatePolicyForwardPolicyDualStackActionStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicyDualStackActionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
