package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_sflow_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceSflowGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_packet_sample_records": {
							Type: schema.TypeInt, Optional: true, Description: "Total packet sample records",
						},
						"total_counter_sample_records": {
							Type: schema.TypeInt, Optional: true, Description: "Total counter sample records",
						},
						"total_sflow_packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total sflow packets sent",
						},
						"total_sflow_local_packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Total sflow packets sent  desc {}",
						},
						"total_sflow_packets_sent_mgmt": {
							Type: schema.TypeInt, Optional: true, Description: "Total sflow packets sent via Mgmt Interface",
						},
						"total_sflow_packets_drop_mgmt": {
							Type: schema.TypeInt, Optional: true, Description: "sflow packets dropped because of rate limit via Mgmt Interface",
						},
					},
				},
			},
		},
	}
}

func resourceSflowGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SflowGlobalStatsStats := setObjectSflowGlobalStatsStats(res)
		d.Set("stats", SflowGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSflowGlobalStatsStats(ret edpt.DataSflowGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_packet_sample_records":    ret.DtSflowGlobalStats.Stats.TotalPacketSampleRecords,
			"total_counter_sample_records":   ret.DtSflowGlobalStats.Stats.TotalCounterSampleRecords,
			"total_sflow_packets_sent":       ret.DtSflowGlobalStats.Stats.TotalSflowPacketsSent,
			"total_sflow_local_packets_sent": ret.DtSflowGlobalStats.Stats.TotalSflowLocalPacketsSent,
			"total_sflow_packets_sent_mgmt":  ret.DtSflowGlobalStats.Stats.TotalSflowPacketsSentMgmt,
			"total_sflow_packets_drop_mgmt":  ret.DtSflowGlobalStats.Stats.TotalSflowPacketsDropMgmt,
		},
	}
}

func getObjectSflowGlobalStatsStats(d []interface{}) edpt.SflowGlobalStatsStats {

	count1 := len(d)
	var ret edpt.SflowGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalPacketSampleRecords = in["total_packet_sample_records"].(int)
		ret.TotalCounterSampleRecords = in["total_counter_sample_records"].(int)
		ret.TotalSflowPacketsSent = in["total_sflow_packets_sent"].(int)
		ret.TotalSflowLocalPacketsSent = in["total_sflow_local_packets_sent"].(int)
		ret.TotalSflowPacketsSentMgmt = in["total_sflow_packets_sent_mgmt"].(int)
		ret.TotalSflowPacketsDropMgmt = in["total_sflow_packets_drop_mgmt"].(int)
	}
	return ret
}

func dataToEndpointSflowGlobalStats(d *schema.ResourceData) edpt.SflowGlobalStats {
	var ret edpt.SflowGlobalStats

	ret.Stats = getObjectSflowGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
