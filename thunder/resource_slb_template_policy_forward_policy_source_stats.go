package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_source_stats`: Statistics for the object source\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicySourceStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "source destination match rule name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hits": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests matching this source rule",
						},
						"destination_match_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Number of requests without matching destination rule",
						},
						"no_host_info": {
							Type: schema.TypeInt, Optional: true, Description: "Failed to parse ip or host information from request",
						},
					},
				},
			},
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicySourceStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicySourceStatsStats := setObjectSlbTemplatePolicyForwardPolicySourceStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicySourceStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicySourceStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicySourceStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits":                        ret.DtSlbTemplatePolicyForwardPolicySourceStats.Stats.Hits,
			"destination_match_not_found": ret.DtSlbTemplatePolicyForwardPolicySourceStats.Stats.DestinationMatchNotFound,
			"no_host_info":                ret.DtSlbTemplatePolicyForwardPolicySourceStats.Stats.NoHostInfo,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicySourceStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
		ret.DestinationMatchNotFound = in["destination_match_not_found"].(int)
		ret.NoHostInfo = in["no_host_info"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceStats {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicySourceStatsStats(d.Get("stats").([]interface{}))
	return ret
}
