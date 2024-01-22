package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAamAaaPolicyAaaRuleStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_aam_aaa_policy_aaa_rule_stats`: Statistics for the object aaa-rule\n\n__PLACEHOLDER__",
		ReadContext: resourceAamAaaPolicyAaaRuleStatsRead,

		Schema: map[string]*schema.Schema{
			"index": {
				Type: schema.TypeInt, Required: true, Description: "Specify AAA rule index",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hit_deny": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hit_auth": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"hit_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"failure_bypass": {
							Type: schema.TypeInt, Optional: true, Description: "",
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

func resourceAamAaaPolicyAaaRuleStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAamAaaPolicyAaaRuleStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAamAaaPolicyAaaRuleStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AamAaaPolicyAaaRuleStatsStats := setObjectAamAaaPolicyAaaRuleStatsStats(res)
		d.Set("stats", AamAaaPolicyAaaRuleStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAamAaaPolicyAaaRuleStatsStats(ret edpt.DataAamAaaPolicyAaaRuleStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_count":    ret.DtAamAaaPolicyAaaRuleStats.Stats.Total_count,
			"hit_deny":       ret.DtAamAaaPolicyAaaRuleStats.Stats.Hit_deny,
			"hit_auth":       ret.DtAamAaaPolicyAaaRuleStats.Stats.Hit_auth,
			"hit_bypass":     ret.DtAamAaaPolicyAaaRuleStats.Stats.Hit_bypass,
			"failure_bypass": ret.DtAamAaaPolicyAaaRuleStats.Stats.Failure_bypass,
		},
	}
}

func getObjectAamAaaPolicyAaaRuleStatsStats(d []interface{}) edpt.AamAaaPolicyAaaRuleStatsStats {

	count1 := len(d)
	var ret edpt.AamAaaPolicyAaaRuleStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_count = in["total_count"].(int)
		ret.Hit_deny = in["hit_deny"].(int)
		ret.Hit_auth = in["hit_auth"].(int)
		ret.Hit_bypass = in["hit_bypass"].(int)
		ret.Failure_bypass = in["failure_bypass"].(int)
	}
	return ret
}

func dataToEndpointAamAaaPolicyAaaRuleStats(d *schema.ResourceData) edpt.AamAaaPolicyAaaRuleStats {
	var ret edpt.AamAaaPolicyAaaRuleStats

	ret.Index = d.Get("index").(int)

	ret.Stats = getObjectAamAaaPolicyAaaRuleStatsStats(d.Get("stats").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
