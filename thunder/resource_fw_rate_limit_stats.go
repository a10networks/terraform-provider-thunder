package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwRateLimitStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_rate_limit_stats`: Statistics for the object rate-limit\n\n__PLACEHOLDER__",
		ReadContext: resourceFwRateLimitStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ratelimit_used_total_mem": {
							Type: schema.TypeInt, Optional: true, Description: "Total Memory Used For Rate-limiting (bytes)",
						},
						"ratelimit_entry_count_t2_key": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Total Rate-limit Entries",
						},
						"ratelimit_entry_count_fw_rule_uid": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope Aggregate",
						},
						"ratelimit_entry_count_ip_addr": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope IPv4 Address",
						},
						"ratelimit_entry_count_ip6_addr": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope IPv6 Address",
						},
						"ratelimit_entry_count_session_id": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope Session ID",
						},
						"ratelimit_entry_count_rule_ipv4_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope IPv4 Prefix",
						},
						"ratelimit_entry_count_rule_ipv6_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope IPv6 Prefix",
						},
						"ratelimit_entry_count_parent_uid": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Parent Rate-limit Entries with Scope Aggregate",
						},
						"ratelimit_entry_count_parent_ipv4_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Parent Rate-limit Entries with Scope IPv4 Prefix",
						},
						"ratelimit_entry_count_parent_ipv6_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Parent Rate-limit Entries with Scope IPv6 Prefix",
						},
						"ratelimit_entry_count_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries Allocated Totally",
						},
						"ratelimit_entry_count_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries Freed Totally",
						},
						"ratelimit_entry_count_rule_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Rate-limit Entries with Scope IP",
						},
						"ratelimit_entry_count_parent_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Parent Rate-limit Entries with Scope IP",
						},
						"ratelimit_entry_count_radius_usergroup": {
							Type: schema.TypeInt, Optional: true, Description: "The total number of rate-limiting entries with the scope RADIUS user group.",
						},
						"ratelimit_entry_count_parent_radius_usergroup": {
							Type: schema.TypeInt, Optional: true, Description: "The total number of parent rate-limiting entries with the scope RADIUS user group.",
						},
						"ratelimit_entry_count_radius_userid": {
							Type: schema.TypeInt, Optional: true, Description: "The total number of rate-limiting entries with the scope RADIUS user ID.",
						},
						"ratelimit_entry_count_parent_radius_userid": {
							Type: schema.TypeInt, Optional: true, Description: "The total number of parent rate-limiting entries with the scope RADIUS user ID.",
						},
					},
				},
			},
		},
	}
}

func resourceFwRateLimitStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwRateLimitStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwRateLimitStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwRateLimitStatsStats := setObjectFwRateLimitStatsStats(res)
		d.Set("stats", FwRateLimitStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwRateLimitStatsStats(ret edpt.DataFwRateLimitStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"ratelimit_used_total_mem":                      ret.DtFwRateLimitStats.Stats.Ratelimit_used_total_mem,
			"ratelimit_entry_count_t2_key":                  ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_t2_key,
			"ratelimit_entry_count_fw_rule_uid":             ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_fw_rule_uid,
			"ratelimit_entry_count_ip_addr":                 ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_ip_addr,
			"ratelimit_entry_count_ip6_addr":                ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_ip6_addr,
			"ratelimit_entry_count_session_id":              ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_session_id,
			"ratelimit_entry_count_rule_ipv4_prefix":        ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_rule_ipv4_prefix,
			"ratelimit_entry_count_rule_ipv6_prefix":        ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_rule_ipv6_prefix,
			"ratelimit_entry_count_parent_uid":              ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_parent_uid,
			"ratelimit_entry_count_parent_ipv4_prefix":      ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_parent_ipv4_prefix,
			"ratelimit_entry_count_parent_ipv6_prefix":      ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_parent_ipv6_prefix,
			"ratelimit_entry_count_allocated":               ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_allocated,
			"ratelimit_entry_count_freed":                   ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_freed,
			"ratelimit_entry_count_rule_ip":                 ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_rule_ip,
			"ratelimit_entry_count_parent_ip":               ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_parent_ip,
			"ratelimit_entry_count_radius_usergroup":        ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_radius_usergroup,
			"ratelimit_entry_count_parent_radius_usergroup": ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_parent_radius_usergroup,
			"ratelimit_entry_count_radius_userid":           ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_radius_userid,
			"ratelimit_entry_count_parent_radius_userid":    ret.DtFwRateLimitStats.Stats.Ratelimit_entry_count_parent_radius_userid,
		},
	}
}

func getObjectFwRateLimitStatsStats(d []interface{}) edpt.FwRateLimitStatsStats {

	count1 := len(d)
	var ret edpt.FwRateLimitStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ratelimit_used_total_mem = in["ratelimit_used_total_mem"].(int)
		ret.Ratelimit_entry_count_t2_key = in["ratelimit_entry_count_t2_key"].(int)
		ret.Ratelimit_entry_count_fw_rule_uid = in["ratelimit_entry_count_fw_rule_uid"].(int)
		ret.Ratelimit_entry_count_ip_addr = in["ratelimit_entry_count_ip_addr"].(int)
		ret.Ratelimit_entry_count_ip6_addr = in["ratelimit_entry_count_ip6_addr"].(int)
		ret.Ratelimit_entry_count_session_id = in["ratelimit_entry_count_session_id"].(int)
		ret.Ratelimit_entry_count_rule_ipv4_prefix = in["ratelimit_entry_count_rule_ipv4_prefix"].(int)
		ret.Ratelimit_entry_count_rule_ipv6_prefix = in["ratelimit_entry_count_rule_ipv6_prefix"].(int)
		ret.Ratelimit_entry_count_parent_uid = in["ratelimit_entry_count_parent_uid"].(int)
		ret.Ratelimit_entry_count_parent_ipv4_prefix = in["ratelimit_entry_count_parent_ipv4_prefix"].(int)
		ret.Ratelimit_entry_count_parent_ipv6_prefix = in["ratelimit_entry_count_parent_ipv6_prefix"].(int)
		ret.Ratelimit_entry_count_allocated = in["ratelimit_entry_count_allocated"].(int)
		ret.Ratelimit_entry_count_freed = in["ratelimit_entry_count_freed"].(int)
		ret.Ratelimit_entry_count_rule_ip = in["ratelimit_entry_count_rule_ip"].(int)
		ret.Ratelimit_entry_count_parent_ip = in["ratelimit_entry_count_parent_ip"].(int)
		ret.Ratelimit_entry_count_radius_usergroup = in["ratelimit_entry_count_radius_usergroup"].(int)
		ret.Ratelimit_entry_count_parent_radius_usergroup = in["ratelimit_entry_count_parent_radius_usergroup"].(int)
		ret.Ratelimit_entry_count_radius_userid = in["ratelimit_entry_count_radius_userid"].(int)
		ret.Ratelimit_entry_count_parent_radius_userid = in["ratelimit_entry_count_parent_radius_userid"].(int)
	}
	return ret
}

func dataToEndpointFwRateLimitStats(d *schema.ResourceData) edpt.FwRateLimitStats {
	var ret edpt.FwRateLimitStats

	ret.Stats = getObjectFwRateLimitStatsStats(d.Get("stats").([]interface{}))
	return ret
}
