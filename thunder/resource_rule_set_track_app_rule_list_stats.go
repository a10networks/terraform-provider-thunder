package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetTrackAppRuleListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_track_app_rule_list_stats`: Statistics for the object track-app-rule-list\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetTrackAppRuleListStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dummy": {
							Type: schema.TypeInt, Optional: true, Description: "Entry for a10countergen",
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

func resourceRuleSetTrackAppRuleListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetTrackAppRuleListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetTrackAppRuleListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetTrackAppRuleListStatsStats := setObjectRuleSetTrackAppRuleListStatsStats(res)
		d.Set("stats", RuleSetTrackAppRuleListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetTrackAppRuleListStatsStats(ret edpt.DataRuleSetTrackAppRuleListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dummy": ret.DtRuleSetTrackAppRuleListStats.Stats.Dummy,
		},
	}
}

func getObjectRuleSetTrackAppRuleListStatsStats(d []interface{}) edpt.RuleSetTrackAppRuleListStatsStats {

	count1 := len(d)
	var ret edpt.RuleSetTrackAppRuleListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func dataToEndpointRuleSetTrackAppRuleListStats(d *schema.ResourceData) edpt.RuleSetTrackAppRuleListStats {
	var ret edpt.RuleSetTrackAppRuleListStats

	ret.Stats = getObjectRuleSetTrackAppRuleListStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
