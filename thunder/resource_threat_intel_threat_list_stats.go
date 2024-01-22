package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelThreatListStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_threat_intel_threat_list_stats`: Statistics for the object threat-list\n\n__PLACEHOLDER__",
		ReadContext: resourceThreatIntelThreatListStatsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Threat category List name",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spam_sources": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for spam sources",
						},
						"windows_exploits": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for windows exploits",
						},
						"web_attacks": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for web attacks",
						},
						"botnets": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for botnets",
						},
						"scanners": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for scanners",
						},
						"dos_attacks": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for dos attacks",
						},
						"reputation": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for reputation",
						},
						"phishing": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for phishing",
						},
						"proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for proxy",
						},
						"mobile_threats": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for mobile threats",
						},
						"tor_proxy": {
							Type: schema.TypeInt, Optional: true, Description: "Hits for tor-proxy",
						},
						"total_hits": {
							Type: schema.TypeInt, Optional: true, Description: "Total hits for threat-list",
						},
					},
				},
			},
		},
	}
}

func resourceThreatIntelThreatListStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelThreatListStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelThreatListStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ThreatIntelThreatListStatsStats := setObjectThreatIntelThreatListStatsStats(res)
		d.Set("stats", ThreatIntelThreatListStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectThreatIntelThreatListStatsStats(ret edpt.DataThreatIntelThreatListStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"spam_sources":     ret.DtThreatIntelThreatListStats.Stats.SpamSources,
			"windows_exploits": ret.DtThreatIntelThreatListStats.Stats.WindowsExploits,
			"web_attacks":      ret.DtThreatIntelThreatListStats.Stats.WebAttacks,
			"botnets":          ret.DtThreatIntelThreatListStats.Stats.Botnets,
			"scanners":         ret.DtThreatIntelThreatListStats.Stats.Scanners,
			"dos_attacks":      ret.DtThreatIntelThreatListStats.Stats.DosAttacks,
			"reputation":       ret.DtThreatIntelThreatListStats.Stats.Reputation,
			"phishing":         ret.DtThreatIntelThreatListStats.Stats.Phishing,
			"proxy":            ret.DtThreatIntelThreatListStats.Stats.Proxy,
			"mobile_threats":   ret.DtThreatIntelThreatListStats.Stats.MobileThreats,
			"tor_proxy":        ret.DtThreatIntelThreatListStats.Stats.TorProxy,
			"total_hits":       ret.DtThreatIntelThreatListStats.Stats.TotalHits,
		},
	}
}

func getObjectThreatIntelThreatListStatsStats(d []interface{}) edpt.ThreatIntelThreatListStatsStats {

	count1 := len(d)
	var ret edpt.ThreatIntelThreatListStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SpamSources = in["spam_sources"].(int)
		ret.WindowsExploits = in["windows_exploits"].(int)
		ret.WebAttacks = in["web_attacks"].(int)
		ret.Botnets = in["botnets"].(int)
		ret.Scanners = in["scanners"].(int)
		ret.DosAttacks = in["dos_attacks"].(int)
		ret.Reputation = in["reputation"].(int)
		ret.Phishing = in["phishing"].(int)
		ret.Proxy = in["proxy"].(int)
		ret.MobileThreats = in["mobile_threats"].(int)
		ret.TorProxy = in["tor_proxy"].(int)
		ret.TotalHits = in["total_hits"].(int)
	}
	return ret
}

func dataToEndpointThreatIntelThreatListStats(d *schema.ResourceData) edpt.ThreatIntelThreatListStats {
	var ret edpt.ThreatIntelThreatListStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectThreatIntelThreatListStatsStats(d.Get("stats").([]interface{}))
	return ret
}
