package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicyActionStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_action_stats`: Statistics for the object action\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicyActionStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Action policy name",
			},
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
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicyActionStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicyActionStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicyActionStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicyActionStatsStats := setObjectSlbTemplatePolicyForwardPolicyActionStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicyActionStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicyActionStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicyActionStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicyActionStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicyActionStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicyActionStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicyActionStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicyActionStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicyActionStats {
	var ret edpt.SlbTemplatePolicyForwardPolicyActionStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicyActionStatsStats(d.Get("stats").([]interface{}))
	return ret
}
