package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_policy_forward_policy_source_destination_any_stats`: Statistics for the object any\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsRead,

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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAnyStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats := setObjectSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats(res)
		d.Set("stats", SlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats(ret edpt.DataSlbTemplatePolicyForwardPolicySourceDestinationAnyStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hits": ret.DtSlbTemplatePolicyForwardPolicySourceDestinationAnyStats.Stats.Hits,
		},
	}
}

func getObjectSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats(d []interface{}) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats {

	count1 := len(d)
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hits = in["hits"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplatePolicyForwardPolicySourceDestinationAnyStats(d *schema.ResourceData) edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnyStats {
	var ret edpt.SlbTemplatePolicyForwardPolicySourceDestinationAnyStats

	ret.Stats = getObjectSlbTemplatePolicyForwardPolicySourceDestinationAnyStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
