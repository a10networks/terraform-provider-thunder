package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceThreatIntelWebrootGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_threat_intel_webroot_global_stats`: Statistics for the object webroot-global\n\n__PLACEHOLDER__",
		ReadContext: resourceThreatIntelWebrootGlobalStatsRead,

		Schema: map[string]*schema.Schema{
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
						"rtu_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "Number of lookups in RTU cache",
						},
						"database_lookup": {
							Type: schema.TypeInt, Optional: true, Description: "Number of lookups in database",
						},
						"non_malicious_ips": {
							Type: schema.TypeInt, Optional: true, Description: "IP's not found in database or RTU cache",
						},
					},
				},
			},
		},
	}
}

func resourceThreatIntelWebrootGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceThreatIntelWebrootGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointThreatIntelWebrootGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		ThreatIntelWebrootGlobalStatsStats := setObjectThreatIntelWebrootGlobalStatsStats(res)
		d.Set("stats", ThreatIntelWebrootGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectThreatIntelWebrootGlobalStatsStats(ret edpt.DataThreatIntelWebrootGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"spam_sources":      ret.DtThreatIntelWebrootGlobalStats.Stats.SpamSources,
			"windows_exploits":  ret.DtThreatIntelWebrootGlobalStats.Stats.WindowsExploits,
			"web_attacks":       ret.DtThreatIntelWebrootGlobalStats.Stats.WebAttacks,
			"botnets":           ret.DtThreatIntelWebrootGlobalStats.Stats.Botnets,
			"scanners":          ret.DtThreatIntelWebrootGlobalStats.Stats.Scanners,
			"dos_attacks":       ret.DtThreatIntelWebrootGlobalStats.Stats.DosAttacks,
			"reputation":        ret.DtThreatIntelWebrootGlobalStats.Stats.Reputation,
			"phishing":          ret.DtThreatIntelWebrootGlobalStats.Stats.Phishing,
			"proxy":             ret.DtThreatIntelWebrootGlobalStats.Stats.Proxy,
			"mobile_threats":    ret.DtThreatIntelWebrootGlobalStats.Stats.MobileThreats,
			"tor_proxy":         ret.DtThreatIntelWebrootGlobalStats.Stats.TorProxy,
			"rtu_lookup":        ret.DtThreatIntelWebrootGlobalStats.Stats.RtuLookup,
			"database_lookup":   ret.DtThreatIntelWebrootGlobalStats.Stats.DatabaseLookup,
			"non_malicious_ips": ret.DtThreatIntelWebrootGlobalStats.Stats.NonMaliciousIps,
		},
	}
}

func getObjectThreatIntelWebrootGlobalStatsStats(d []interface{}) edpt.ThreatIntelWebrootGlobalStatsStats {

	count1 := len(d)
	var ret edpt.ThreatIntelWebrootGlobalStatsStats
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
		ret.RtuLookup = in["rtu_lookup"].(int)
		ret.DatabaseLookup = in["database_lookup"].(int)
		ret.NonMaliciousIps = in["non_malicious_ips"].(int)
	}
	return ret
}

func dataToEndpointThreatIntelWebrootGlobalStats(d *schema.ResourceData) edpt.ThreatIntelWebrootGlobalStats {
	var ret edpt.ThreatIntelWebrootGlobalStats

	ret.Stats = getObjectThreatIntelWebrootGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
