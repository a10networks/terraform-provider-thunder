package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpsGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ips_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceIpsGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ips_matched_total": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Total",
						},
						"ips_matched_http": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched HTTP",
						},
						"ips_matched_dns": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched DNS",
						},
						"ips_matched_other": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Other",
						},
						"ips_matched_action_pass": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Action Pass",
						},
						"ips_matched_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Action Drop",
						},
						"ips_matched_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Action Blacklist",
						},
						"ips_matched_severity_high": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Severity High",
						},
						"ips_matched_severity_medium": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Severity Medium",
						},
						"ips_matched_severity_low": {
							Type: schema.TypeInt, Optional: true, Description: "IPS Matched Severity Low",
						},
					},
				},
			},
		},
	}
}

func resourceIpsGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpsGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpsGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		IpsGlobalStatsStats := setObjectIpsGlobalStatsStats(res)
		d.Set("stats", IpsGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpsGlobalStatsStats(ret edpt.DataIpsGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ips_matched_total":            ret.DtIpsGlobalStats.Stats.Ips_matched_total,
			"ips_matched_http":             ret.DtIpsGlobalStats.Stats.Ips_matched_http,
			"ips_matched_dns":              ret.DtIpsGlobalStats.Stats.Ips_matched_dns,
			"ips_matched_other":            ret.DtIpsGlobalStats.Stats.Ips_matched_other,
			"ips_matched_action_pass":      ret.DtIpsGlobalStats.Stats.Ips_matched_action_pass,
			"ips_matched_action_drop":      ret.DtIpsGlobalStats.Stats.Ips_matched_action_drop,
			"ips_matched_action_blacklist": ret.DtIpsGlobalStats.Stats.Ips_matched_action_blacklist,
			"ips_matched_severity_high":    ret.DtIpsGlobalStats.Stats.Ips_matched_severity_high,
			"ips_matched_severity_medium":  ret.DtIpsGlobalStats.Stats.Ips_matched_severity_medium,
			"ips_matched_severity_low":     ret.DtIpsGlobalStats.Stats.Ips_matched_severity_low,
		},
	}
}

func getObjectIpsGlobalStatsStats(d []interface{}) edpt.IpsGlobalStatsStats {

	count1 := len(d)
	var ret edpt.IpsGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ips_matched_total = in["ips_matched_total"].(int)
		ret.Ips_matched_http = in["ips_matched_http"].(int)
		ret.Ips_matched_dns = in["ips_matched_dns"].(int)
		ret.Ips_matched_other = in["ips_matched_other"].(int)
		ret.Ips_matched_action_pass = in["ips_matched_action_pass"].(int)
		ret.Ips_matched_action_drop = in["ips_matched_action_drop"].(int)
		ret.Ips_matched_action_blacklist = in["ips_matched_action_blacklist"].(int)
		ret.Ips_matched_severity_high = in["ips_matched_severity_high"].(int)
		ret.Ips_matched_severity_medium = in["ips_matched_severity_medium"].(int)
		ret.Ips_matched_severity_low = in["ips_matched_severity_low"].(int)
	}
	return ret
}

func dataToEndpointIpsGlobalStats(d *schema.ResourceData) edpt.IpsGlobalStats {
	var ret edpt.IpsGlobalStats

	ret.Stats = getObjectIpsGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
