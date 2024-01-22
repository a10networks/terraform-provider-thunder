package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebCategoryReputationScopeStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_web_category_reputation_scope_stats`: Statistics for the object reputation-scope\n\n__PLACEHOLDER__",
		ReadContext: resourceWebCategoryReputationScopeStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Reputation Scope name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trustworthy": {
							Type: schema.TypeInt, Optional: true, Description: "Trustworthy level(81-100)",
						},
						"low_risk": {
							Type: schema.TypeInt, Optional: true, Description: "Low-risk level(61-80)",
						},
						"moderate_risk": {
							Type: schema.TypeInt, Optional: true, Description: "Moderate-risk level(41-60)",
						},
						"suspicious": {
							Type: schema.TypeInt, Optional: true, Description: "Suspicious level(21-40)",
						},
						"malicious": {
							Type: schema.TypeInt, Optional: true, Description: "Malicious level(1-20)",
						},
					},
				},
			},
		},
	}
}

func resourceWebCategoryReputationScopeStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebCategoryReputationScopeStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebCategoryReputationScopeStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		WebCategoryReputationScopeStatsStats := setObjectWebCategoryReputationScopeStatsStats(res)
		d.Set("stats", WebCategoryReputationScopeStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectWebCategoryReputationScopeStatsStats(ret edpt.DataWebCategoryReputationScopeStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"trustworthy":   ret.DtWebCategoryReputationScopeStats.Stats.Trustworthy,
			"low_risk":      ret.DtWebCategoryReputationScopeStats.Stats.LowRisk,
			"moderate_risk": ret.DtWebCategoryReputationScopeStats.Stats.ModerateRisk,
			"suspicious":    ret.DtWebCategoryReputationScopeStats.Stats.Suspicious,
			"malicious":     ret.DtWebCategoryReputationScopeStats.Stats.Malicious,
		},
	}
}

func getObjectWebCategoryReputationScopeStatsStats(d []interface{}) edpt.WebCategoryReputationScopeStatsStats {

	count1 := len(d)
	var ret edpt.WebCategoryReputationScopeStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Trustworthy = in["trustworthy"].(int)
		ret.LowRisk = in["low_risk"].(int)
		ret.ModerateRisk = in["moderate_risk"].(int)
		ret.Suspicious = in["suspicious"].(int)
		ret.Malicious = in["malicious"].(int)
	}
	return ret
}

func dataToEndpointWebCategoryReputationScopeStats(d *schema.ResourceData) edpt.WebCategoryReputationScopeStats {
	var ret edpt.WebCategoryReputationScopeStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectWebCategoryReputationScopeStatsStats(d.Get("stats").([]interface{}))
	return ret
}
