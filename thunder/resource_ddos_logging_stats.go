package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosLoggingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_logging_stats`: Statistics for the object logging\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosLoggingStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_msg_quota_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Log Message Per CPU Quota Exceed",
						},
						"log_msg_oom": {
							Type: schema.TypeInt, Optional: true, Description: "Log Message Out of Memory Error",
						},
						"log_queue_full": {
							Type: schema.TypeInt, Optional: true, Description: "Log Message Queue Full",
						},
						"log_msg_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Log Message Sent",
						},
						"log_msg_send_err": {
							Type: schema.TypeInt, Optional: true, Description: "Log Message Send Error",
						},
					},
				},
			},
		},
	}
}

func resourceDdosLoggingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLoggingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLoggingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosLoggingStatsStats := setObjectDdosLoggingStatsStats(res)
		d.Set("stats", DdosLoggingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosLoggingStatsStats(ret edpt.DataDdosLoggingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"log_msg_quota_exceed": ret.DtDdosLoggingStats.Stats.Log_msg_quota_exceed,
			"log_msg_oom":          ret.DtDdosLoggingStats.Stats.Log_msg_oom,
			"log_queue_full":       ret.DtDdosLoggingStats.Stats.Log_queue_full,
			"log_msg_sent":         ret.DtDdosLoggingStats.Stats.Log_msg_sent,
			"log_msg_send_err":     ret.DtDdosLoggingStats.Stats.Log_msg_send_err,
		},
	}
}

func getObjectDdosLoggingStatsStats(d []interface{}) edpt.DdosLoggingStatsStats {

	count1 := len(d)
	var ret edpt.DdosLoggingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Log_msg_quota_exceed = in["log_msg_quota_exceed"].(int)
		ret.Log_msg_oom = in["log_msg_oom"].(int)
		ret.Log_queue_full = in["log_queue_full"].(int)
		ret.Log_msg_sent = in["log_msg_sent"].(int)
		ret.Log_msg_send_err = in["log_msg_send_err"].(int)
	}
	return ret
}

func dataToEndpointDdosLoggingStats(d *schema.ResourceData) edpt.DdosLoggingStats {
	var ret edpt.DdosLoggingStats

	ret.Stats = getObjectDdosLoggingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
