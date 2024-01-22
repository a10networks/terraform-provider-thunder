package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsStatisticsStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_acos_events_statistics_stats`: Statistics for the object statistics\n\n__PLACEHOLDER__",
		ReadContext: resourceAcosEventsStatisticsStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Messages sent, to Remote",
						},
						"msg_sent_logdb": {
							Type: schema.TypeInt, Optional: true, Description: "Messages sent, to LogDB",
						},
						"msg_dropped_format_not_defined": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, format not defined",
						},
						"msg_dropped_malloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, malloc failure",
						},
						"msg_dropped_no_template": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, no active template",
						},
						"msg_dropped_selector": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, selector does not enable msg",
						},
						"msg_dropped_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, invalid length",
						},
						"msg_dropped_craft_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, msg crafting failed",
						},
						"msg_dropped_local_log_ratelimit": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, local log ratelimited",
						},
						"msg_dropped_send_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, send failed",
						},
						"msg_dropped_no_active_member": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, no active member in collector grp",
						},
						"msg_dropped_route_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, Route lookup failed",
						},
						"msg_dropped_other": {
							Type: schema.TypeInt, Optional: true, Description: "Messages Dropped, unexpected error",
						},
						"param_msg_sent_to_hc": {
							Type: schema.TypeInt, Optional: true, Description: "Parameterized log sent to HC",
						},
						"param_msg_sent_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parameterized log send to HC failed",
						},
						"param_msg_encode_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Parameterized log AVRO encoding failed",
						},
					},
				},
			},
		},
	}
}

func resourceAcosEventsStatisticsStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsStatisticsStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsStatisticsStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		AcosEventsStatisticsStatsStats := setObjectAcosEventsStatisticsStatsStats(res)
		d.Set("stats", AcosEventsStatisticsStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectAcosEventsStatisticsStatsStats(ret edpt.DataAcosEventsStatisticsStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"msg_sent":                        ret.DtAcosEventsStatisticsStats.Stats.Msg_sent,
			"msg_sent_logdb":                  ret.DtAcosEventsStatisticsStats.Stats.Msg_sent_logdb,
			"msg_dropped_format_not_defined":  ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_format_not_defined,
			"msg_dropped_malloc_failure":      ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_malloc_failure,
			"msg_dropped_no_template":         ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_no_template,
			"msg_dropped_selector":            ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_selector,
			"msg_dropped_too_long":            ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_too_long,
			"msg_dropped_craft_fail":          ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_craft_fail,
			"msg_dropped_local_log_ratelimit": ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_local_log_ratelimit,
			"msg_dropped_send_failed":         ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_send_failed,
			"msg_dropped_no_active_member":    ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_no_active_member,
			"msg_dropped_route_fail":          ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_route_fail,
			"msg_dropped_other":               ret.DtAcosEventsStatisticsStats.Stats.Msg_dropped_other,
			"param_msg_sent_to_hc":            ret.DtAcosEventsStatisticsStats.Stats.Param_msg_sent_to_hc,
			"param_msg_sent_fail":             ret.DtAcosEventsStatisticsStats.Stats.Param_msg_sent_fail,
			"param_msg_encode_fail":           ret.DtAcosEventsStatisticsStats.Stats.Param_msg_encode_fail,
		},
	}
}

func getObjectAcosEventsStatisticsStatsStats(d []interface{}) edpt.AcosEventsStatisticsStatsStats {

	count1 := len(d)
	var ret edpt.AcosEventsStatisticsStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Msg_sent = in["msg_sent"].(int)
		ret.Msg_sent_logdb = in["msg_sent_logdb"].(int)
		ret.Msg_dropped_format_not_defined = in["msg_dropped_format_not_defined"].(int)
		ret.Msg_dropped_malloc_failure = in["msg_dropped_malloc_failure"].(int)
		ret.Msg_dropped_no_template = in["msg_dropped_no_template"].(int)
		ret.Msg_dropped_selector = in["msg_dropped_selector"].(int)
		ret.Msg_dropped_too_long = in["msg_dropped_too_long"].(int)
		ret.Msg_dropped_craft_fail = in["msg_dropped_craft_fail"].(int)
		ret.Msg_dropped_local_log_ratelimit = in["msg_dropped_local_log_ratelimit"].(int)
		ret.Msg_dropped_send_failed = in["msg_dropped_send_failed"].(int)
		ret.Msg_dropped_no_active_member = in["msg_dropped_no_active_member"].(int)
		ret.Msg_dropped_route_fail = in["msg_dropped_route_fail"].(int)
		ret.Msg_dropped_other = in["msg_dropped_other"].(int)
		ret.Param_msg_sent_to_hc = in["param_msg_sent_to_hc"].(int)
		ret.Param_msg_sent_fail = in["param_msg_sent_fail"].(int)
		ret.Param_msg_encode_fail = in["param_msg_encode_fail"].(int)
	}
	return ret
}

func dataToEndpointAcosEventsStatisticsStats(d *schema.ResourceData) edpt.AcosEventsStatisticsStats {
	var ret edpt.AcosEventsStatisticsStats

	ret.Stats = getObjectAcosEventsStatisticsStatsStats(d.Get("stats").([]interface{}))
	return ret
}
