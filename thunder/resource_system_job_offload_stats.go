package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemJobOffloadStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_job_offload_stats`: Statistics for the object job-offload\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemJobOffloadStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"jobs": {
							Type: schema.TypeInt, Optional: true, Description: "Current Jobs",
						},
						"submit": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Submitted",
						},
						"receive": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Received",
						},
						"execute": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Executed",
						},
						"snt_home": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Sent Back Home",
						},
						"rcv_home": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Received Home",
						},
						"complete": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Completed",
						},
						"fail_submit": {
							Type: schema.TypeInt, Optional: true, Description: "Jobs Failed to Submit",
						},
						"q_no_space": {
							Type: schema.TypeInt, Optional: true, Description: "No More Space in Q",
						},
						"fail_execute": {
							Type: schema.TypeInt, Optional: true, Description: "Failed to Execute Job",
						},
						"fail_complete": {
							Type: schema.TypeInt, Optional: true, Description: "Failed to Complete Job",
						},
					},
				},
			},
		},
	}
}

func resourceSystemJobOffloadStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemJobOffloadStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemJobOffloadStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemJobOffloadStatsStats := setObjectSystemJobOffloadStatsStats(res)
		d.Set("stats", SystemJobOffloadStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemJobOffloadStatsStats(ret edpt.DataSystemJobOffloadStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"jobs":          ret.DtSystemJobOffloadStats.Stats.Jobs,
			"submit":        ret.DtSystemJobOffloadStats.Stats.Submit,
			"receive":       ret.DtSystemJobOffloadStats.Stats.Receive,
			"execute":       ret.DtSystemJobOffloadStats.Stats.Execute,
			"snt_home":      ret.DtSystemJobOffloadStats.Stats.Snt_home,
			"rcv_home":      ret.DtSystemJobOffloadStats.Stats.Rcv_home,
			"complete":      ret.DtSystemJobOffloadStats.Stats.Complete,
			"fail_submit":   ret.DtSystemJobOffloadStats.Stats.Fail_submit,
			"q_no_space":    ret.DtSystemJobOffloadStats.Stats.Q_no_space,
			"fail_execute":  ret.DtSystemJobOffloadStats.Stats.Fail_execute,
			"fail_complete": ret.DtSystemJobOffloadStats.Stats.Fail_complete,
		},
	}
}

func getObjectSystemJobOffloadStatsStats(d []interface{}) edpt.SystemJobOffloadStatsStats {

	count1 := len(d)
	var ret edpt.SystemJobOffloadStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Jobs = in["jobs"].(int)
		ret.Submit = in["submit"].(int)
		ret.Receive = in["receive"].(int)
		ret.Execute = in["execute"].(int)
		ret.Snt_home = in["snt_home"].(int)
		ret.Rcv_home = in["rcv_home"].(int)
		ret.Complete = in["complete"].(int)
		ret.Fail_submit = in["fail_submit"].(int)
		ret.Q_no_space = in["q_no_space"].(int)
		ret.Fail_execute = in["fail_execute"].(int)
		ret.Fail_complete = in["fail_complete"].(int)
	}
	return ret
}

func dataToEndpointSystemJobOffloadStats(d *schema.ResourceData) edpt.SystemJobOffloadStats {
	var ret edpt.SystemJobOffloadStats

	ret.Stats = getObjectSystemJobOffloadStatsStats(d.Get("stats").([]interface{}))
	return ret
}
