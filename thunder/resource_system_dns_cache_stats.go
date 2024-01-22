package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDnsCacheStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_dns_cache_stats`: Statistics for the object dns-cache\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemDnsCacheStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_q": {
							Type: schema.TypeInt, Optional: true, Description: "Total query",
						},
						"total_r": {
							Type: schema.TypeInt, Optional: true, Description: "Total server response",
						},
						"hit": {
							Type: schema.TypeInt, Optional: true, Description: "Total cache hit",
						},
						"bad_q": {
							Type: schema.TypeInt, Optional: true, Description: "Query not passed",
						},
						"encode_q": {
							Type: schema.TypeInt, Optional: true, Description: "Query encoded",
						},
						"multiple_q": {
							Type: schema.TypeInt, Optional: true, Description: "Query with multiple questions",
						},
						"oversize_q": {
							Type: schema.TypeInt, Optional: true, Description: "Query exceed cache size",
						},
						"bad_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response not passed",
						},
						"oversize_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response exceed cache size",
						},
						"encode_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response encoded",
						},
						"multiple_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response with multiple questions",
						},
						"answer_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response with multiple answers",
						},
						"ttl_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response with short TTL",
						},
						"ageout": {
							Type: schema.TypeInt, Optional: true, Description: "Total aged out",
						},
						"bad_answer": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Answer",
						},
						"ageout_weight": {
							Type: schema.TypeInt, Optional: true, Description: "Total aged for lower weight",
						},
						"total_log": {
							Type: schema.TypeInt, Optional: true, Description: "Total stats log sent",
						},
						"total_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Total allocated",
						},
						"total_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total freed",
						},
						"current_allocate": {
							Type: schema.TypeInt, Optional: true, Description: "Current allocate",
						},
						"current_data_allocate": {
							Type: schema.TypeInt, Optional: true, Description: "Current data allocate",
						},
						"resolver_queue_full": {
							Type: schema.TypeInt, Optional: true, Description: "Resolver task queue full",
						},
						"truncated_r": {
							Type: schema.TypeInt, Optional: true, Description: "Response with Truncation bit set",
						},
					},
				},
			},
		},
	}
}

func resourceSystemDnsCacheStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsCacheStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsCacheStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemDnsCacheStatsStats := setObjectSystemDnsCacheStatsStats(res)
		d.Set("stats", SystemDnsCacheStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemDnsCacheStatsStats(ret edpt.DataSystemDnsCacheStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_q":               ret.DtSystemDnsCacheStats.Stats.Total_q,
			"total_r":               ret.DtSystemDnsCacheStats.Stats.Total_r,
			"hit":                   ret.DtSystemDnsCacheStats.Stats.Hit,
			"bad_q":                 ret.DtSystemDnsCacheStats.Stats.Bad_q,
			"encode_q":              ret.DtSystemDnsCacheStats.Stats.Encode_q,
			"multiple_q":            ret.DtSystemDnsCacheStats.Stats.Multiple_q,
			"oversize_q":            ret.DtSystemDnsCacheStats.Stats.Oversize_q,
			"bad_r":                 ret.DtSystemDnsCacheStats.Stats.Bad_r,
			"oversize_r":            ret.DtSystemDnsCacheStats.Stats.Oversize_r,
			"encode_r":              ret.DtSystemDnsCacheStats.Stats.Encode_r,
			"multiple_r":            ret.DtSystemDnsCacheStats.Stats.Multiple_r,
			"answer_r":              ret.DtSystemDnsCacheStats.Stats.Answer_r,
			"ttl_r":                 ret.DtSystemDnsCacheStats.Stats.Ttl_r,
			"ageout":                ret.DtSystemDnsCacheStats.Stats.Ageout,
			"bad_answer":            ret.DtSystemDnsCacheStats.Stats.Bad_answer,
			"ageout_weight":         ret.DtSystemDnsCacheStats.Stats.Ageout_weight,
			"total_log":             ret.DtSystemDnsCacheStats.Stats.Total_log,
			"total_alloc":           ret.DtSystemDnsCacheStats.Stats.Total_alloc,
			"total_freed":           ret.DtSystemDnsCacheStats.Stats.Total_freed,
			"current_allocate":      ret.DtSystemDnsCacheStats.Stats.Current_allocate,
			"current_data_allocate": ret.DtSystemDnsCacheStats.Stats.Current_data_allocate,
			"resolver_queue_full":   ret.DtSystemDnsCacheStats.Stats.Resolver_queue_full,
			"truncated_r":           ret.DtSystemDnsCacheStats.Stats.Truncated_r,
		},
	}
}

func getObjectSystemDnsCacheStatsStats(d []interface{}) edpt.SystemDnsCacheStatsStats {

	count1 := len(d)
	var ret edpt.SystemDnsCacheStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_q = in["total_q"].(int)
		ret.Total_r = in["total_r"].(int)
		ret.Hit = in["hit"].(int)
		ret.Bad_q = in["bad_q"].(int)
		ret.Encode_q = in["encode_q"].(int)
		ret.Multiple_q = in["multiple_q"].(int)
		ret.Oversize_q = in["oversize_q"].(int)
		ret.Bad_r = in["bad_r"].(int)
		ret.Oversize_r = in["oversize_r"].(int)
		ret.Encode_r = in["encode_r"].(int)
		ret.Multiple_r = in["multiple_r"].(int)
		ret.Answer_r = in["answer_r"].(int)
		ret.Ttl_r = in["ttl_r"].(int)
		ret.Ageout = in["ageout"].(int)
		ret.Bad_answer = in["bad_answer"].(int)
		ret.Ageout_weight = in["ageout_weight"].(int)
		ret.Total_log = in["total_log"].(int)
		ret.Total_alloc = in["total_alloc"].(int)
		ret.Total_freed = in["total_freed"].(int)
		ret.Current_allocate = in["current_allocate"].(int)
		ret.Current_data_allocate = in["current_data_allocate"].(int)
		ret.Resolver_queue_full = in["resolver_queue_full"].(int)
		ret.Truncated_r = in["truncated_r"].(int)
	}
	return ret
}

func dataToEndpointSystemDnsCacheStats(d *schema.ResourceData) edpt.SystemDnsCacheStats {
	var ret edpt.SystemDnsCacheStats

	ret.Stats = getObjectSystemDnsCacheStatsStats(d.Get("stats").([]interface{}))
	return ret
}
