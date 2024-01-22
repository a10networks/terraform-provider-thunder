package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLocalLoggingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_acos_events_local_logging_stats`: Statistics for the object local-logging\n\n__PLACEHOLDER__",
		ReadContext: resourceAcosEventsLocalLoggingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"init_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Local logging Init Successful",
						},
						"init_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Local logging Init Fail",
						},
						"freed": {
							Type: schema.TypeInt, Optional: true, Description: "Local logging Stopped",
						},
						"disk_over_thres": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs Dropped, Disk reached threshold",
						},
						"rate_limited": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs Dropped, Rate limited",
						},
						"not_inited": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs Dropped, Local logging not inited",
						},
						"sent_to_store": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs sent to be stored",
						},
						"sent_to_store_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of Logs sent to be stored Failed",
						},
						"store_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs failed to be stored",
						},
						"in_logs": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs successfully stored",
						},
						"in_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Number of bytes successfully stored",
						},
						"in_logs_backlog": {
							Type: schema.TypeInt, Optional: true, Description: "Number of backlogs loaded from disk",
						},
						"in_bytes_backlog": {
							Type: schema.TypeInt, Optional: true, Description: "Number of backlog bytes loaded from disk",
						},
						"in_store_fail_no_space": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs Dropped, failed without disk space",
						},
						"in_discard_logs": {
							Type: schema.TypeInt, Optional: true, Description: "Number of old logs discarded to fit in new logs",
						},
						"in_discard_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Number of old bytes discarded to fit in new logs",
						},
						"out_logs": {
							Type: schema.TypeInt, Optional: true, Description: "Number of logs sent to log servers",
						},
						"out_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Number of bytes sent to log-servers",
						},
						"out_error": {
							Type: schema.TypeInt, Optional: true, Description: "Number of errors during send",
						},
						"remaining_logs": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of remaining logs yet to be sent",
						},
						"remaining_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of remaining bytes yet to be sent",
						},
					},
				},
			},
		},
	}
}

func resourceAcosEventsLocalLoggingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLocalLoggingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLocalLoggingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AcosEventsLocalLoggingStatsStats := setObjectAcosEventsLocalLoggingStatsStats(res)
		d.Set("stats", AcosEventsLocalLoggingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAcosEventsLocalLoggingStatsStats(ret edpt.DataAcosEventsLocalLoggingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"init_pass":              ret.DtAcosEventsLocalLoggingStats.Stats.InitPass,
			"init_fail":              ret.DtAcosEventsLocalLoggingStats.Stats.InitFail,
			"freed":                  ret.DtAcosEventsLocalLoggingStats.Stats.Freed,
			"disk_over_thres":        ret.DtAcosEventsLocalLoggingStats.Stats.DiskOverThres,
			"rate_limited":           ret.DtAcosEventsLocalLoggingStats.Stats.RateLimited,
			"not_inited":             ret.DtAcosEventsLocalLoggingStats.Stats.NotInited,
			"sent_to_store":          ret.DtAcosEventsLocalLoggingStats.Stats.SentToStore,
			"sent_to_store_fail":     ret.DtAcosEventsLocalLoggingStats.Stats.SentToStoreFail,
			"store_fail":             ret.DtAcosEventsLocalLoggingStats.Stats.StoreFail,
			"in_logs":                ret.DtAcosEventsLocalLoggingStats.Stats.InLogs,
			"in_bytes":               ret.DtAcosEventsLocalLoggingStats.Stats.InBytes,
			"in_logs_backlog":        ret.DtAcosEventsLocalLoggingStats.Stats.InLogsBacklog,
			"in_bytes_backlog":       ret.DtAcosEventsLocalLoggingStats.Stats.InBytesBacklog,
			"in_store_fail_no_space": ret.DtAcosEventsLocalLoggingStats.Stats.InStoreFailNoSpace,
			"in_discard_logs":        ret.DtAcosEventsLocalLoggingStats.Stats.InDiscardLogs,
			"in_discard_bytes":       ret.DtAcosEventsLocalLoggingStats.Stats.InDiscardBytes,
			"out_logs":               ret.DtAcosEventsLocalLoggingStats.Stats.OutLogs,
			"out_bytes":              ret.DtAcosEventsLocalLoggingStats.Stats.OutBytes,
			"out_error":              ret.DtAcosEventsLocalLoggingStats.Stats.OutError,
			"remaining_logs":         ret.DtAcosEventsLocalLoggingStats.Stats.RemainingLogs,
			"remaining_bytes":        ret.DtAcosEventsLocalLoggingStats.Stats.RemainingBytes,
		},
	}
}

func getObjectAcosEventsLocalLoggingStatsStats(d []interface{}) edpt.AcosEventsLocalLoggingStatsStats {

	count1 := len(d)
	var ret edpt.AcosEventsLocalLoggingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.InitPass = in["init_pass"].(int)
		ret.InitFail = in["init_fail"].(int)
		ret.Freed = in["freed"].(int)
		ret.DiskOverThres = in["disk_over_thres"].(int)
		ret.RateLimited = in["rate_limited"].(int)
		ret.NotInited = in["not_inited"].(int)
		ret.SentToStore = in["sent_to_store"].(int)
		ret.SentToStoreFail = in["sent_to_store_fail"].(int)
		ret.StoreFail = in["store_fail"].(int)
		ret.InLogs = in["in_logs"].(int)
		ret.InBytes = in["in_bytes"].(int)
		ret.InLogsBacklog = in["in_logs_backlog"].(int)
		ret.InBytesBacklog = in["in_bytes_backlog"].(int)
		ret.InStoreFailNoSpace = in["in_store_fail_no_space"].(int)
		ret.InDiscardLogs = in["in_discard_logs"].(int)
		ret.InDiscardBytes = in["in_discard_bytes"].(int)
		ret.OutLogs = in["out_logs"].(int)
		ret.OutBytes = in["out_bytes"].(int)
		ret.OutError = in["out_error"].(int)
		ret.RemainingLogs = in["remaining_logs"].(int)
		ret.RemainingBytes = in["remaining_bytes"].(int)
	}
	return ret
}

func dataToEndpointAcosEventsLocalLoggingStats(d *schema.ResourceData) edpt.AcosEventsLocalLoggingStats {
	var ret edpt.AcosEventsLocalLoggingStats

	ret.Stats = getObjectAcosEventsLocalLoggingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
