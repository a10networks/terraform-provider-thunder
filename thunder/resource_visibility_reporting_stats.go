package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityReportingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_reporting_stats`: Statistics for the object reporting\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityReportingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_transmit_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total log transmit failures",
						},
						"buffer_alloc_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Total reporting buffer allocation failures",
						},
						"notif_jobs_in_queue": {
							Type: schema.TypeInt, Optional: true, Description: "Total notification jobs in queue",
						},
						"enqueue_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Total enqueue jobs failed",
						},
						"enqueue_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Total enqueue jobs passed",
						},
						"dequeued": {
							Type: schema.TypeInt, Optional: true, Description: "Total jobs dequeued",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityReportingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityReportingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityReportingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityReportingStatsStats := setObjectVisibilityReportingStatsStats(res)
		d.Set("stats", VisibilityReportingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityReportingStatsStats(ret edpt.DataVisibilityReportingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"log_transmit_failure": ret.DtVisibilityReportingStats.Stats.LogTransmitFailure,
			"buffer_alloc_failure": ret.DtVisibilityReportingStats.Stats.BufferAllocFailure,
			"notif_jobs_in_queue":  ret.DtVisibilityReportingStats.Stats.NotifJobsInQueue,
			"enqueue_fail":         ret.DtVisibilityReportingStats.Stats.EnqueueFail,
			"enqueue_pass":         ret.DtVisibilityReportingStats.Stats.EnqueuePass,
			"dequeued":             ret.DtVisibilityReportingStats.Stats.Dequeued,
		},
	}
}

func getObjectVisibilityReportingStatsStats(d []interface{}) edpt.VisibilityReportingStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityReportingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LogTransmitFailure = in["log_transmit_failure"].(int)
		ret.BufferAllocFailure = in["buffer_alloc_failure"].(int)
		ret.NotifJobsInQueue = in["notif_jobs_in_queue"].(int)
		ret.EnqueueFail = in["enqueue_fail"].(int)
		ret.EnqueuePass = in["enqueue_pass"].(int)
		ret.Dequeued = in["dequeued"].(int)
	}
	return ret
}

func dataToEndpointVisibilityReportingStats(d *schema.ResourceData) edpt.VisibilityReportingStats {
	var ret edpt.VisibilityReportingStats

	ret.Stats = getObjectVisibilityReportingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
