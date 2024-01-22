package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbDnsResponseRateLimitingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_dns_response_rate_limiting_stats`: Statistics for the object dns-response-rate-limiting\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbDnsResponseRateLimitingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"curr_entries": {
							Type: schema.TypeInt, Optional: true, Description: "Current Entry Count",
						},
						"total_created": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Created",
						},
						"total_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Inserted",
						},
						"total_withdrew": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Withdrew",
						},
						"total_ready_to_free": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Ready To Free",
						},
						"total_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total Entry Freed",
						},
						"total_logs": {
							Type: schema.TypeInt, Optional: true, Description: "Total Logs",
						},
						"total_overflow_entry_hits": {
							Type: schema.TypeInt, Optional: true, Description: "Total Overflow Entry Hits",
						},
						"total_refill": {
							Type: schema.TypeInt, Optional: true, Description: "Total Refills",
						},
						"total_credit_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Total Credit Exceeded",
						},
						"other_thread_refill": {
							Type: schema.TypeInt, Optional: true, Description: "Other Thread Refilling",
						},
						"err_entry_create_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Creation Failure",
						},
						"err_entry_create_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Creation Out of Memory",
						},
						"err_entry_ext_create_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Extension Creation Out of Memory",
						},
						"err_entry_insert_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Insert Failed",
						},
						"err_vport_fail_match": {
							Type: schema.TypeInt, Optional: true, Description: "Failed to Match Vport",
						},
					},
				},
			},
		},
	}
}

func resourceSlbDnsResponseRateLimitingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbDnsResponseRateLimitingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbDnsResponseRateLimitingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbDnsResponseRateLimitingStatsStats := setObjectSlbDnsResponseRateLimitingStatsStats(res)
		d.Set("stats", SlbDnsResponseRateLimitingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbDnsResponseRateLimitingStatsStats(ret edpt.DataSlbDnsResponseRateLimitingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"curr_entries":              ret.DtSlbDnsResponseRateLimitingStats.Stats.Curr_entries,
			"total_created":             ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_created,
			"total_inserted":            ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_inserted,
			"total_withdrew":            ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_withdrew,
			"total_ready_to_free":       ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_ready_to_free,
			"total_freed":               ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_freed,
			"total_logs":                ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_logs,
			"total_overflow_entry_hits": ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_overflow_entry_hits,
			"total_refill":              ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_refill,
			"total_credit_exceeded":     ret.DtSlbDnsResponseRateLimitingStats.Stats.Total_credit_exceeded,
			"other_thread_refill":       ret.DtSlbDnsResponseRateLimitingStats.Stats.Other_thread_refill,
			"err_entry_create_failed":   ret.DtSlbDnsResponseRateLimitingStats.Stats.Err_entry_create_failed,
			"err_entry_create_oom":      ret.DtSlbDnsResponseRateLimitingStats.Stats.Err_entry_create_oom,
			"err_entry_ext_create_oom":  ret.DtSlbDnsResponseRateLimitingStats.Stats.Err_entry_ext_create_oom,
			"err_entry_insert_failed":   ret.DtSlbDnsResponseRateLimitingStats.Stats.Err_entry_insert_failed,
			"err_vport_fail_match":      ret.DtSlbDnsResponseRateLimitingStats.Stats.Err_vport_fail_match,
		},
	}
}

func getObjectSlbDnsResponseRateLimitingStatsStats(d []interface{}) edpt.SlbDnsResponseRateLimitingStatsStats {

	count1 := len(d)
	var ret edpt.SlbDnsResponseRateLimitingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Curr_entries = in["curr_entries"].(int)
		ret.Total_created = in["total_created"].(int)
		ret.Total_inserted = in["total_inserted"].(int)
		ret.Total_withdrew = in["total_withdrew"].(int)
		ret.Total_ready_to_free = in["total_ready_to_free"].(int)
		ret.Total_freed = in["total_freed"].(int)
		ret.Total_logs = in["total_logs"].(int)
		ret.Total_overflow_entry_hits = in["total_overflow_entry_hits"].(int)
		ret.Total_refill = in["total_refill"].(int)
		ret.Total_credit_exceeded = in["total_credit_exceeded"].(int)
		ret.Other_thread_refill = in["other_thread_refill"].(int)
		ret.Err_entry_create_failed = in["err_entry_create_failed"].(int)
		ret.Err_entry_create_oom = in["err_entry_create_oom"].(int)
		ret.Err_entry_ext_create_oom = in["err_entry_ext_create_oom"].(int)
		ret.Err_entry_insert_failed = in["err_entry_insert_failed"].(int)
		ret.Err_vport_fail_match = in["err_vport_fail_match"].(int)
	}
	return ret
}

func dataToEndpointSlbDnsResponseRateLimitingStats(d *schema.ResourceData) edpt.SlbDnsResponseRateLimitingStats {
	var ret edpt.SlbDnsResponseRateLimitingStats

	ret.Stats = getObjectSlbDnsResponseRateLimitingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
