package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRuleStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_rule_stats`: Statistics for the object rule\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetRuleStatsRead,

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
						"permit_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Permitted bytes counter",
						},
						"deny_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Denied bytes counter",
						},
						"reset_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Reset bytes counter",
						},
						"permit_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Permitted packets counter",
						},
						"deny_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Denied packets counter",
						},
						"reset_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Reset packets counter",
						},
						"active_session_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Active TCP session counter",
						},
						"active_session_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Active UDP session counter",
						},
						"active_session_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Active ICMP session counter",
						},
						"active_session_other": {
							Type: schema.TypeInt, Optional: true, Description: "Active other protocol session counter",
						},
						"session_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP session counter",
						},
						"session_udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP session counter",
						},
						"session_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP session counter",
						},
						"session_other": {
							Type: schema.TypeInt, Optional: true, Description: "Other protocol session counter",
						},
						"active_session_sctp": {
							Type: schema.TypeInt, Optional: true, Description: "Active SCTP session counter",
						},
						"session_sctp": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP session counter",
						},
						"hitcount_timestamp": {
							Type: schema.TypeInt, Optional: true, Description: "Last hit counts timestamp",
						},
						"rate_limit_drops": {
							Type: schema.TypeInt, Optional: true, Description: "Rate Limit Drops",
						},
					},
				},
			},
		},
	}
}

func resourceRuleSetRuleStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetRuleStatsStats := setObjectRuleSetRuleStatsStats(res)
		d.Set("stats", RuleSetRuleStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetRuleStatsStats(ret edpt.DataRuleSetRuleStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hit_count":            ret.DtRuleSetRuleStats.Stats.HitCount,
			"permit_bytes":         ret.DtRuleSetRuleStats.Stats.PermitBytes,
			"deny_bytes":           ret.DtRuleSetRuleStats.Stats.DenyBytes,
			"reset_bytes":          ret.DtRuleSetRuleStats.Stats.ResetBytes,
			"permit_packets":       ret.DtRuleSetRuleStats.Stats.PermitPackets,
			"deny_packets":         ret.DtRuleSetRuleStats.Stats.DenyPackets,
			"reset_packets":        ret.DtRuleSetRuleStats.Stats.ResetPackets,
			"active_session_tcp":   ret.DtRuleSetRuleStats.Stats.ActiveSessionTcp,
			"active_session_udp":   ret.DtRuleSetRuleStats.Stats.ActiveSessionUdp,
			"active_session_icmp":  ret.DtRuleSetRuleStats.Stats.ActiveSessionIcmp,
			"active_session_other": ret.DtRuleSetRuleStats.Stats.ActiveSessionOther,
			"session_tcp":          ret.DtRuleSetRuleStats.Stats.SessionTcp,
			"session_udp":          ret.DtRuleSetRuleStats.Stats.SessionUdp,
			"session_icmp":         ret.DtRuleSetRuleStats.Stats.SessionIcmp,
			"session_other":        ret.DtRuleSetRuleStats.Stats.SessionOther,
			"active_session_sctp":  ret.DtRuleSetRuleStats.Stats.ActiveSessionSctp,
			"session_sctp":         ret.DtRuleSetRuleStats.Stats.SessionSctp,
			"hitcount_timestamp":   ret.DtRuleSetRuleStats.Stats.HitcountTimestamp,
			"rate_limit_drops":     ret.DtRuleSetRuleStats.Stats.RateLimitDrops,
		},
	}
}

func getObjectRuleSetRuleStatsStats(d []interface{}) edpt.RuleSetRuleStatsStats {

	count1 := len(d)
	var ret edpt.RuleSetRuleStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HitCount = in["hit_count"].(int)
		ret.PermitBytes = in["permit_bytes"].(int)
		ret.DenyBytes = in["deny_bytes"].(int)
		ret.ResetBytes = in["reset_bytes"].(int)
		ret.PermitPackets = in["permit_packets"].(int)
		ret.DenyPackets = in["deny_packets"].(int)
		ret.ResetPackets = in["reset_packets"].(int)
		ret.ActiveSessionTcp = in["active_session_tcp"].(int)
		ret.ActiveSessionUdp = in["active_session_udp"].(int)
		ret.ActiveSessionIcmp = in["active_session_icmp"].(int)
		ret.ActiveSessionOther = in["active_session_other"].(int)
		ret.SessionTcp = in["session_tcp"].(int)
		ret.SessionUdp = in["session_udp"].(int)
		ret.SessionIcmp = in["session_icmp"].(int)
		ret.SessionOther = in["session_other"].(int)
		ret.ActiveSessionSctp = in["active_session_sctp"].(int)
		ret.SessionSctp = in["session_sctp"].(int)
		ret.HitcountTimestamp = in["hitcount_timestamp"].(int)
		ret.RateLimitDrops = in["rate_limit_drops"].(int)
	}
	return ret
}

func dataToEndpointRuleSetRuleStats(d *schema.ResourceData) edpt.RuleSetRuleStats {
	var ret edpt.RuleSetRuleStats

	ret.Name = d.Get("name").(string)

	ret.Stats = getObjectRuleSetRuleStatsStats(d.Get("stats").([]interface{}))
	return ret
}
