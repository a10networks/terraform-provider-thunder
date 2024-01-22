package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbRateLimitLogStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_rate_limit_log_stats`: Statistics for the object rate-limit-log\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbRateLimitLogStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_log_times": {
							Type: schema.TypeInt, Optional: true, Description: "Total log times",
						},
						"total_log_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Total log messages",
						},
						"local_log_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Local log messages",
						},
						"remote_log_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Remote log messages",
						},
						"local_log_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Local rate (per sec)",
						},
						"remote_log_rate": {
							Type: schema.TypeInt, Optional: true, Description: "Remote rate (per sec)",
						},
						"msg_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Log message too big",
						},
						"buff_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Buffer alloc fail",
						},
						"no_route": {
							Type: schema.TypeInt, Optional: true, Description: "No route",
						},
						"buff_send_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Buffer send fail",
						},
						"alloc_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Log-session alloc",
						},
						"free_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Log-session free",
						},
						"conn_alloc_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Log-session alloc fail",
						},
						"no_repeat_msg": {
							Type: schema.TypeInt, Optional: true, Description: "No repeat message",
						},
						"local_log_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Local log dropped due to rate-limit",
						},
					},
				},
			},
		},
	}
}

func resourceSlbRateLimitLogStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbRateLimitLogStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbRateLimitLogStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbRateLimitLogStatsStats := setObjectSlbRateLimitLogStatsStats(res)
		d.Set("stats", SlbRateLimitLogStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbRateLimitLogStatsStats(ret edpt.DataSlbRateLimitLogStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_log_times":   ret.DtSlbRateLimitLogStats.Stats.Total_log_times,
			"total_log_msg":     ret.DtSlbRateLimitLogStats.Stats.Total_log_msg,
			"local_log_msg":     ret.DtSlbRateLimitLogStats.Stats.Local_log_msg,
			"remote_log_msg":    ret.DtSlbRateLimitLogStats.Stats.Remote_log_msg,
			"local_log_rate":    ret.DtSlbRateLimitLogStats.Stats.Local_log_rate,
			"remote_log_rate":   ret.DtSlbRateLimitLogStats.Stats.Remote_log_rate,
			"msg_too_big":       ret.DtSlbRateLimitLogStats.Stats.Msg_too_big,
			"buff_alloc_fail":   ret.DtSlbRateLimitLogStats.Stats.Buff_alloc_fail,
			"no_route":          ret.DtSlbRateLimitLogStats.Stats.No_route,
			"buff_send_fail":    ret.DtSlbRateLimitLogStats.Stats.Buff_send_fail,
			"alloc_conn":        ret.DtSlbRateLimitLogStats.Stats.Alloc_conn,
			"free_conn":         ret.DtSlbRateLimitLogStats.Stats.Free_conn,
			"conn_alloc_fail":   ret.DtSlbRateLimitLogStats.Stats.Conn_alloc_fail,
			"no_repeat_msg":     ret.DtSlbRateLimitLogStats.Stats.No_repeat_msg,
			"local_log_dropped": ret.DtSlbRateLimitLogStats.Stats.Local_log_dropped,
		},
	}
}

func getObjectSlbRateLimitLogStatsStats(d []interface{}) edpt.SlbRateLimitLogStatsStats {

	count1 := len(d)
	var ret edpt.SlbRateLimitLogStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_log_times = in["total_log_times"].(int)
		ret.Total_log_msg = in["total_log_msg"].(int)
		ret.Local_log_msg = in["local_log_msg"].(int)
		ret.Remote_log_msg = in["remote_log_msg"].(int)
		ret.Local_log_rate = in["local_log_rate"].(int)
		ret.Remote_log_rate = in["remote_log_rate"].(int)
		ret.Msg_too_big = in["msg_too_big"].(int)
		ret.Buff_alloc_fail = in["buff_alloc_fail"].(int)
		ret.No_route = in["no_route"].(int)
		ret.Buff_send_fail = in["buff_send_fail"].(int)
		ret.Alloc_conn = in["alloc_conn"].(int)
		ret.Free_conn = in["free_conn"].(int)
		ret.Conn_alloc_fail = in["conn_alloc_fail"].(int)
		ret.No_repeat_msg = in["no_repeat_msg"].(int)
		ret.Local_log_dropped = in["local_log_dropped"].(int)
	}
	return ret
}

func dataToEndpointSlbRateLimitLogStats(d *schema.ResourceData) edpt.SlbRateLimitLogStats {
	var ret edpt.SlbRateLimitLogStats

	ret.Stats = getObjectSlbRateLimitLogStatsStats(d.Get("stats").([]interface{}))
	return ret
}
