package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_traffic_control_rule_set_stats`: Statistics for the object rule-set\n\n__PLACEHOLDER__",
		ReadContext: resourceTrafficControlRuleSetStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule set name",
			},
			"rule_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
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
				},
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

func resourceTrafficControlRuleSetStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		TrafficControlRuleSetStatsRuleList := setSliceTrafficControlRuleSetStatsRuleList(res)
		d.Set("rule_list", TrafficControlRuleSetStatsRuleList)
		TrafficControlRuleSetStatsStats := setObjectTrafficControlRuleSetStatsStats(res)
		d.Set("stats", TrafficControlRuleSetStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setSliceTrafficControlRuleSetStatsRuleList(d edpt.DataTrafficControlRuleSetStats) []map[string]interface{} {
	result := []map[string]interface{}{}

	for _, item := range d.DtTrafficControlRuleSetStats.RuleList {
		in := make(map[string]interface{})
		in["name"] = item.Name
		in["stats"] = setObjectTrafficControlRuleSetStatsRuleListStats(item.Stats)
		result = append(result, in)
	}
	return result
}

func setObjectTrafficControlRuleSetStatsRuleListStats(d edpt.TrafficControlRuleSetStatsRuleListStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["hit_count"] = d.HitCount
	result = append(result, in)
	return result
}

func setObjectTrafficControlRuleSetStatsStats(ret edpt.DataTrafficControlRuleSetStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hit_count": ret.DtTrafficControlRuleSetStats.Stats.HitCount,
		},
	}
}

func getSliceTrafficControlRuleSetStatsRuleList(d []interface{}) []edpt.TrafficControlRuleSetStatsRuleList {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetStatsRuleList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetStatsRuleList
		oi.Name = in["name"].(string)
		oi.Stats = getObjectTrafficControlRuleSetStatsRuleListStats(in["stats"].([]interface{}))
		ret = append(ret, oi)
	}
	return ret
}

func getObjectTrafficControlRuleSetStatsRuleListStats(d []interface{}) edpt.TrafficControlRuleSetStatsRuleListStats {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetStatsRuleListStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCount = in["hit_count"].(int)
	}
	return ret
}

func getObjectTrafficControlRuleSetStatsStats(d []interface{}) edpt.TrafficControlRuleSetStatsStats {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCount = in["hit_count"].(int)
	}
	return ret
}

func dataToEndpointTrafficControlRuleSetStats(d *schema.ResourceData) edpt.TrafficControlRuleSetStats {
	var ret edpt.TrafficControlRuleSetStats

	ret.Name = d.Get("name").(string)

	ret.RuleList = getSliceTrafficControlRuleSetStatsRuleList(d.Get("rule_list").([]interface{}))

	ret.Stats = getObjectTrafficControlRuleSetStatsStats(d.Get("stats").([]interface{}))
	return ret
}
