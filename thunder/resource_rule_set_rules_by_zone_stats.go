package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRulesByZoneStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_rules_by_zone_stats`: Statistics for the object rules-by-zone\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetRulesByZoneStatsRead,

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

func resourceRuleSetRulesByZoneStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRulesByZoneStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRulesByZoneStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetRulesByZoneStatsStats := setObjectRuleSetRulesByZoneStatsStats(res)
		d.Set("stats", RuleSetRulesByZoneStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetRulesByZoneStatsStats(ret edpt.DataRuleSetRulesByZoneStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dummy": ret.DtRuleSetRulesByZoneStats.Stats.Dummy,
		},
	}
}

func getObjectRuleSetRulesByZoneStatsStats(d []interface{}) edpt.RuleSetRulesByZoneStatsStats {

	count1 := len(d)
	var ret edpt.RuleSetRulesByZoneStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dummy = in["dummy"].(int)
	}
	return ret
}

func dataToEndpointRuleSetRulesByZoneStats(d *schema.ResourceData) edpt.RuleSetRulesByZoneStats {
	var ret edpt.RuleSetRulesByZoneStats

	ret.Stats = getObjectRuleSetRulesByZoneStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
