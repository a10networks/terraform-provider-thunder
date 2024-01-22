package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbAflowStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_aflow_stats`: Statistics for the object aflow\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbAflowStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pause_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Pause connection",
						},
						"pause_conn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Pause connection fail",
						},
						"resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume connection",
						},
						"event_resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume conn by event",
						},
						"timer_resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume conn by timer",
						},
						"try_to_resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume conn by trying",
						},
						"retry_resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume conn by retry",
						},
						"error_resume_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Resume conn by error",
						},
						"open_new_server_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Open new server conn",
						},
						"reuse_server_idle_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Reuse idle server conn",
						},
						"inc_aflow_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Inc aFlow limit",
						},
					},
				},
			},
		},
	}
}

func resourceSlbAflowStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAflowStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAflowStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbAflowStatsStats := setObjectSlbAflowStatsStats(res)
		d.Set("stats", SlbAflowStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbAflowStatsStats(ret edpt.DataSlbAflowStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pause_conn":             ret.DtSlbAflowStats.Stats.Pause_conn,
			"pause_conn_fail":        ret.DtSlbAflowStats.Stats.Pause_conn_fail,
			"resume_conn":            ret.DtSlbAflowStats.Stats.Resume_conn,
			"event_resume_conn":      ret.DtSlbAflowStats.Stats.Event_resume_conn,
			"timer_resume_conn":      ret.DtSlbAflowStats.Stats.Timer_resume_conn,
			"try_to_resume_conn":     ret.DtSlbAflowStats.Stats.Try_to_resume_conn,
			"retry_resume_conn":      ret.DtSlbAflowStats.Stats.Retry_resume_conn,
			"error_resume_conn":      ret.DtSlbAflowStats.Stats.Error_resume_conn,
			"open_new_server_conn":   ret.DtSlbAflowStats.Stats.Open_new_server_conn,
			"reuse_server_idle_conn": ret.DtSlbAflowStats.Stats.Reuse_server_idle_conn,
			"inc_aflow_limit":        ret.DtSlbAflowStats.Stats.Inc_aflow_limit,
		},
	}
}

func getObjectSlbAflowStatsStats(d []interface{}) edpt.SlbAflowStatsStats {

	count1 := len(d)
	var ret edpt.SlbAflowStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Pause_conn = in["pause_conn"].(int)
		ret.Pause_conn_fail = in["pause_conn_fail"].(int)
		ret.Resume_conn = in["resume_conn"].(int)
		ret.Event_resume_conn = in["event_resume_conn"].(int)
		ret.Timer_resume_conn = in["timer_resume_conn"].(int)
		ret.Try_to_resume_conn = in["try_to_resume_conn"].(int)
		ret.Retry_resume_conn = in["retry_resume_conn"].(int)
		ret.Error_resume_conn = in["error_resume_conn"].(int)
		ret.Open_new_server_conn = in["open_new_server_conn"].(int)
		ret.Reuse_server_idle_conn = in["reuse_server_idle_conn"].(int)
		ret.Inc_aflow_limit = in["inc_aflow_limit"].(int)
	}
	return ret
}

func dataToEndpointSlbAflowStats(d *schema.ResourceData) edpt.SlbAflowStats {
	var ret edpt.SlbAflowStats

	ret.Stats = getObjectSlbAflowStatsStats(d.Get("stats").([]interface{}))
	return ret
}
