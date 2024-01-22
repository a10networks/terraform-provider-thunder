package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_source_destination_adv_match_stats`: Statistics for the object adv-match\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsRead,

		Schema: map[string]*schema.Schema{
			"priority": {
				Type: schema.TypeInt, Required: true, Description: "Rule priority (1000 is highest)",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests hit this rule",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats := setObjectSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStats

	ret.Priority = d.Get("priority").(int)

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicySourceDestinationAdvMatchStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
