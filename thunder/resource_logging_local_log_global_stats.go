package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLoggingLocalLogGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_logging_local_log_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceLoggingLocalLogGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enqueue": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log enqueue",
						},
						"enqueue_full": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log queue full",
						},
						"enqueue_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log enqueue error",
						},
						"dequeue": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log dequeue",
						},
						"dequeue_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log dequeue processing error",
						},
						"raw_log": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log raw logs",
						},
						"raw_log_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total raw log logging error",
						},
						"log_summarized": {
							Type: schema.TypeInt, Optional: true, Description: "Total raw log summarized",
						},
						"l1_log_summarized": {
							Type: schema.TypeInt, Optional: true, Description: "Total layer 1 log summarized",
						},
						"l2_log_summarized": {
							Type: schema.TypeInt, Optional: true, Description: "Total layer 2 log summarized",
						},
						"log_summarized_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log summarization error",
						},
						"aam_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log AAM raw database",
						},
						"ep_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP raw database",
						},
						"fw_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall raw database",
						},
						"aam_top_user_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log AAM top user summary database",
						},
						"ep_top_user_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP top user summary database",
						},
						"ep_top_src_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP top client summary database",
						},
						"ep_top_dst_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP top destination summary database",
						},
						"ep_top_domain_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP top domain summary database",
						},
						"ep_top_web_category_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP top web-category summary database",
						},
						"ep_top_host_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log EP top host summary database",
						},
						"fw_top_app_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall top application summary database",
						},
						"fw_top_src_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall top source summary database",
						},
						"fw_top_app_src_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall top application and source summary database",
						},
						"fw_top_category_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall top category summary database",
						},
						"db_erro": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log database create error",
						},
						"query": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log axapi query",
						},
						"response": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log axapi response",
						},
						"query_error": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log axapi query error",
						},
						"fw_top_thr_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall top threat summary database",
						},
						"fw_top_thr_src_db": {
							Type: schema.TypeInt, Optional: true, Description: "Total local-log Firewall top threat and source summary database",
						},
					},
				},
			},
		},
	}
}

func resourceLoggingLocalLogGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLoggingLocalLogGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLoggingLocalLogGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		LoggingLocalLogGlobalStatsStats := setObjectLoggingLocalLogGlobalStatsStats(res)
		d.Set("stats", LoggingLocalLogGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectLoggingLocalLogGlobalStatsStats(ret edpt.DataLoggingLocalLogGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"enqueue":                ret.DtLoggingLocalLogGlobalStats.Stats.Enqueue,
			"enqueue_full":           ret.DtLoggingLocalLogGlobalStats.Stats.EnqueueFull,
			"enqueue_error":          ret.DtLoggingLocalLogGlobalStats.Stats.EnqueueError,
			"dequeue":                ret.DtLoggingLocalLogGlobalStats.Stats.Dequeue,
			"dequeue_error":          ret.DtLoggingLocalLogGlobalStats.Stats.DequeueError,
			"raw_log":                ret.DtLoggingLocalLogGlobalStats.Stats.RawLog,
			"raw_log_error":          ret.DtLoggingLocalLogGlobalStats.Stats.RawLogError,
			"log_summarized":         ret.DtLoggingLocalLogGlobalStats.Stats.LogSummarized,
			"l1_log_summarized":      ret.DtLoggingLocalLogGlobalStats.Stats.L1LogSummarized,
			"l2_log_summarized":      ret.DtLoggingLocalLogGlobalStats.Stats.L2LogSummarized,
			"log_summarized_error":   ret.DtLoggingLocalLogGlobalStats.Stats.LogSummarizedError,
			"aam_db":                 ret.DtLoggingLocalLogGlobalStats.Stats.AamDb,
			"ep_db":                  ret.DtLoggingLocalLogGlobalStats.Stats.EpDb,
			"fw_db":                  ret.DtLoggingLocalLogGlobalStats.Stats.FwDb,
			"aam_top_user_db":        ret.DtLoggingLocalLogGlobalStats.Stats.AamTopUserDb,
			"ep_top_user_db":         ret.DtLoggingLocalLogGlobalStats.Stats.EpTopUserDb,
			"ep_top_src_db":          ret.DtLoggingLocalLogGlobalStats.Stats.EpTopSrcDb,
			"ep_top_dst_db":          ret.DtLoggingLocalLogGlobalStats.Stats.EpTopDstDb,
			"ep_top_domain_db":       ret.DtLoggingLocalLogGlobalStats.Stats.EpTopDomainDb,
			"ep_top_web_category_db": ret.DtLoggingLocalLogGlobalStats.Stats.EpTopWebCategoryDb,
			"ep_top_host_db":         ret.DtLoggingLocalLogGlobalStats.Stats.EpTopHostDb,
			"fw_top_app_db":          ret.DtLoggingLocalLogGlobalStats.Stats.FwTopAppDb,
			"fw_top_src_db":          ret.DtLoggingLocalLogGlobalStats.Stats.FwTopSrcDb,
			"fw_top_app_src_db":      ret.DtLoggingLocalLogGlobalStats.Stats.FwTopAppSrcDb,
			"fw_top_category_db":     ret.DtLoggingLocalLogGlobalStats.Stats.FwTopCategoryDb,
			"db_erro":                ret.DtLoggingLocalLogGlobalStats.Stats.DbErro,
			"query":                  ret.DtLoggingLocalLogGlobalStats.Stats.Query,
			"response":               ret.DtLoggingLocalLogGlobalStats.Stats.Response,
			"query_error":            ret.DtLoggingLocalLogGlobalStats.Stats.QueryError,
			"fw_top_thr_db":          ret.DtLoggingLocalLogGlobalStats.Stats.FwTopThrDb,
			"fw_top_thr_src_db":      ret.DtLoggingLocalLogGlobalStats.Stats.FwTopThrSrcDb,
		},
	}
}

func getObjectLoggingLocalLogGlobalStatsStats(d []interface{}) edpt.LoggingLocalLogGlobalStatsStats {

	count1 := len(d)
	var ret edpt.LoggingLocalLogGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enqueue = in["enqueue"].(int)
		ret.EnqueueFull = in["enqueue_full"].(int)
		ret.EnqueueError = in["enqueue_error"].(int)
		ret.Dequeue = in["dequeue"].(int)
		ret.DequeueError = in["dequeue_error"].(int)
		ret.RawLog = in["raw_log"].(int)
		ret.RawLogError = in["raw_log_error"].(int)
		ret.LogSummarized = in["log_summarized"].(int)
		ret.L1LogSummarized = in["l1_log_summarized"].(int)
		ret.L2LogSummarized = in["l2_log_summarized"].(int)
		ret.LogSummarizedError = in["log_summarized_error"].(int)
		ret.AamDb = in["aam_db"].(int)
		ret.EpDb = in["ep_db"].(int)
		ret.FwDb = in["fw_db"].(int)
		ret.AamTopUserDb = in["aam_top_user_db"].(int)
		ret.EpTopUserDb = in["ep_top_user_db"].(int)
		ret.EpTopSrcDb = in["ep_top_src_db"].(int)
		ret.EpTopDstDb = in["ep_top_dst_db"].(int)
		ret.EpTopDomainDb = in["ep_top_domain_db"].(int)
		ret.EpTopWebCategoryDb = in["ep_top_web_category_db"].(int)
		ret.EpTopHostDb = in["ep_top_host_db"].(int)
		ret.FwTopAppDb = in["fw_top_app_db"].(int)
		ret.FwTopSrcDb = in["fw_top_src_db"].(int)
		ret.FwTopAppSrcDb = in["fw_top_app_src_db"].(int)
		ret.FwTopCategoryDb = in["fw_top_category_db"].(int)
		ret.DbErro = in["db_erro"].(int)
		ret.Query = in["query"].(int)
		ret.Response = in["response"].(int)
		ret.QueryError = in["query_error"].(int)
		ret.FwTopThrDb = in["fw_top_thr_db"].(int)
		ret.FwTopThrSrcDb = in["fw_top_thr_src_db"].(int)
	}
	return ret
}

func dataToEndpointLoggingLocalLogGlobalStats(d *schema.ResourceData) edpt.LoggingLocalLogGlobalStats {
	var ret edpt.LoggingLocalLogGlobalStats

	ret.Stats = getObjectLoggingLocalLogGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
