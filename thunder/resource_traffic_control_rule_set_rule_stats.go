package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetRuleStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_traffic_control_rule_set_rule_stats`: Statistics for the object rule\n\n__PLACEHOLDER__",
		ReadContext: resourceTrafficControlRuleSetRuleStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hit_count": {
							Type: schema.TypeInt, Optional: true, Description: "Hit counts",
						},
					},
				},
			},
		},
	}
}

func resourceTrafficControlRuleSetRuleStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRuleStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TrafficControlRuleSetRuleStatsStats := setObjectTrafficControlRuleSetRuleStatsStats(res)
		d.Set("stats", TrafficControlRuleSetRuleStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectTrafficControlRuleSetRuleStatsStats(ret edpt.DataTrafficControlRuleSetRuleStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hit_count": ret.DtTrafficControlRuleSetRuleStats.Stats.HitCount,
		},
	}
}

func getObjectTrafficControlRuleSetRuleStatsStats(d []interface{}) edpt.TrafficControlRuleSetRuleStatsStats {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetRuleStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCount = in["hit_count"].(int)
	}
	return ret
}

func dataToEndpointTrafficControlRuleSetRuleStats(d *schema.ResourceData) edpt.TrafficControlRuleSetRuleStats {
	var ret edpt.TrafficControlRuleSetRuleStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectTrafficControlRuleSetRuleStatsStats(d.Get("stats").([]interface{}))
	return ret
}
