package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6HttpAlgStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_http_alg_stats`: Statistics for the object http-alg\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6HttpAlgStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request_processed": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Request Processed",
						},
						"request_insert_msisdn_performed": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP MSISDN Insertion Performed",
						},
						"request_insert_client_ip_performed": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Client IP Insertion Performed",
						},
						"request_insert_msisdn_unavailable": {
							Type: schema.TypeInt, Optional: true, Description: "Inserted MSISDN is 0000 (MSISDN Unavailable)",
						},
						"queued_session_too_many": {
							Type: schema.TypeInt, Optional: true, Description: "Queued Session Exceed Drop",
						},
						"radius_query_succeed": {
							Type: schema.TypeInt, Optional: true, Description: "MSISDN Query Succeed",
						},
						"radius_query_failed": {
							Type: schema.TypeInt, Optional: true, Description: "MSISDN Query Failed",
						},
						"radius_requst_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Query Request Sent",
						},
						"radius_requst_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Query Request Dropped",
						},
						"radius_response_received": {
							Type: schema.TypeInt, Optional: true, Description: "Query Response Received",
						},
						"radius_response_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Query Response Dropped",
						},
						"out_of_memory_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Out-of-Memory Dropped",
						},
						"queue_len_exceed_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Queue Length Exceed Dropped",
						},
						"out_of_order_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Out-of-Order Dropped",
						},
						"header_insertion_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Buff Insertion Failed",
						},
						"header_removal_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Buff Removal Failed",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6HttpAlgStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6HttpAlgStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6HttpAlgStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6HttpAlgStatsStats := setObjectCgnv6HttpAlgStatsStats(res)
		d.Set("stats", Cgnv6HttpAlgStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6HttpAlgStatsStats(ret edpt.DataCgnv6HttpAlgStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_processed":                  ret.DtCgnv6HttpAlgStats.Stats.RequestProcessed,
			"request_insert_msisdn_performed":    ret.DtCgnv6HttpAlgStats.Stats.RequestInsertMsisdnPerformed,
			"request_insert_client_ip_performed": ret.DtCgnv6HttpAlgStats.Stats.RequestInsertClientIpPerformed,
			"request_insert_msisdn_unavailable":  ret.DtCgnv6HttpAlgStats.Stats.RequestInsertMsisdnUnavailable,
			"queued_session_too_many":            ret.DtCgnv6HttpAlgStats.Stats.QueuedSessionTooMany,
			"radius_query_succeed":               ret.DtCgnv6HttpAlgStats.Stats.RadiusQuerySucceed,
			"radius_query_failed":                ret.DtCgnv6HttpAlgStats.Stats.RadiusQueryFailed,
			"radius_requst_sent":                 ret.DtCgnv6HttpAlgStats.Stats.RadiusRequstSent,
			"radius_requst_dropped":              ret.DtCgnv6HttpAlgStats.Stats.RadiusRequstDropped,
			"radius_response_received":           ret.DtCgnv6HttpAlgStats.Stats.RadiusResponseReceived,
			"radius_response_dropped":            ret.DtCgnv6HttpAlgStats.Stats.RadiusResponseDropped,
			"out_of_memory_dropped":              ret.DtCgnv6HttpAlgStats.Stats.OutOfMemoryDropped,
			"queue_len_exceed_dropped":           ret.DtCgnv6HttpAlgStats.Stats.QueueLenExceedDropped,
			"out_of_order_dropped":               ret.DtCgnv6HttpAlgStats.Stats.OutOfOrderDropped,
			"header_insertion_failed":            ret.DtCgnv6HttpAlgStats.Stats.HeaderInsertionFailed,
			"header_removal_failed":              ret.DtCgnv6HttpAlgStats.Stats.HeaderRemovalFailed,
		},
	}
}

func getObjectCgnv6HttpAlgStatsStats(d []interface{}) edpt.Cgnv6HttpAlgStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6HttpAlgStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RequestProcessed = in["request_processed"].(int)
		ret.RequestInsertMsisdnPerformed = in["request_insert_msisdn_performed"].(int)
		ret.RequestInsertClientIpPerformed = in["request_insert_client_ip_performed"].(int)
		ret.RequestInsertMsisdnUnavailable = in["request_insert_msisdn_unavailable"].(int)
		ret.QueuedSessionTooMany = in["queued_session_too_many"].(int)
		ret.RadiusQuerySucceed = in["radius_query_succeed"].(int)
		ret.RadiusQueryFailed = in["radius_query_failed"].(int)
		ret.RadiusRequstSent = in["radius_requst_sent"].(int)
		ret.RadiusRequstDropped = in["radius_requst_dropped"].(int)
		ret.RadiusResponseReceived = in["radius_response_received"].(int)
		ret.RadiusResponseDropped = in["radius_response_dropped"].(int)
		ret.OutOfMemoryDropped = in["out_of_memory_dropped"].(int)
		ret.QueueLenExceedDropped = in["queue_len_exceed_dropped"].(int)
		ret.OutOfOrderDropped = in["out_of_order_dropped"].(int)
		ret.HeaderInsertionFailed = in["header_insertion_failed"].(int)
		ret.HeaderRemovalFailed = in["header_removal_failed"].(int)
	}
	return ret
}

func dataToEndpointCgnv6HttpAlgStats(d *schema.ResourceData) edpt.Cgnv6HttpAlgStats {
	var ret edpt.Cgnv6HttpAlgStats

	ret.Stats = getObjectCgnv6HttpAlgStatsStats(d.Get("stats").([]interface{}))
	return ret
}
