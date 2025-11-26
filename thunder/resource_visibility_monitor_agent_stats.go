package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityMonitorAgentStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_visibility_monitor_agent_stats`: Statistics for the object agent\n\n__PLACEHOLDER__",
		ReadContext: resourceVisibilityMonitorAgentStatsRead,

		Schema: map[string]*schema.Schema{
			"agent_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name for the agent",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sflow_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Packets Received",
						},
						"sflow_samples_received": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Samples Received",
						},
						"sflow_samples_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Samples Bad Length",
						},
						"sflow_samples_non_std": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Samples Non-standard",
						},
						"sflow_samples_skipped": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Samples Skipped",
						},
						"sflow_sample_record_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Sample Records Bad Length",
						},
						"sflow_samples_sent_for_detection": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Samples Processed For Detection",
						},
						"sflow_sample_record_invalid_layer2": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Sample Records Unknown Layer-2",
						},
						"sflow_sample_ipv6_hdr_parse_fail": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Sample IPv6 Record Header Parse Failures",
						},
						"sflow_disabled": {
							Type: schema.TypeInt, Optional: true, Description: "sFlow Packet Samples Processing Disabled",
						},
						"netflow_disabled": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Flow Samples Processing Disabled",
						},
						"netflow_v5_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Packets Received",
						},
						"netflow_v5_samples_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Samples Received",
						},
						"netflow_v5_samples_sent_for_detection": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Samples Processed For Detection",
						},
						"netflow_v5_sample_records_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Sample Records Bad Length",
						},
						"netflow_v5_max_records_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v5 Sample Max Records Error",
						},
						"netflow_v9_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Packets Received",
						},
						"netflow_v9_samples_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Samples Received",
						},
						"netflow_v9_samples_sent_for_detection": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Samples Processed For Detection",
						},
						"netflow_v9_sample_records_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Sample Records Bad Length",
						},
						"netflow_v9_max_records_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Sample Max Records Error",
						},
						"netflow_v10_packets_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10 Packets Received",
						},
						"netflow_v10_samples_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10 Samples Received",
						},
						"netflow_v10_samples_sent_for_detection": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10 Samples Procssed For Detection",
						},
						"netflow_v10_sample_records_bad_len": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10 Sample Records Bad Length",
						},
						"netflow_v10_max_records_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v10 Sample Max records Error",
						},
						"netflow_tcp_sample_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow TCP Samples Received",
						},
						"netflow_udp_sample_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow UDP Samples received",
						},
						"netflow_icmp_sample_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow ICMP Samples Received",
						},
						"netflow_other_sample_received": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow OTHER Samples Received",
						},
						"netflow_record_copy_oom_error": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Data Record Copy Fail OOM",
						},
						"netflow_record_rse_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Data Record Reduced Size Invalid",
						},
						"netflow_sample_flow_dur_error": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Sample Flow Duration Error",
						},
					},
				},
			},
		},
	}
}

func resourceVisibilityMonitorAgentStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityMonitorAgentStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityMonitorAgentStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VisibilityMonitorAgentStatsStats := setObjectVisibilityMonitorAgentStatsStats(res)
		d.Set("stats", VisibilityMonitorAgentStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVisibilityMonitorAgentStatsStats(ret edpt.DataVisibilityMonitorAgentStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sflow_packets_received":                 ret.DtVisibilityMonitorAgentStats.Stats.SflowPacketsReceived,
			"sflow_samples_received":                 ret.DtVisibilityMonitorAgentStats.Stats.SflowSamplesReceived,
			"sflow_samples_bad_len":                  ret.DtVisibilityMonitorAgentStats.Stats.SflowSamplesBadLen,
			"sflow_samples_non_std":                  ret.DtVisibilityMonitorAgentStats.Stats.SflowSamplesNonStd,
			"sflow_samples_skipped":                  ret.DtVisibilityMonitorAgentStats.Stats.SflowSamplesSkipped,
			"sflow_sample_record_bad_len":            ret.DtVisibilityMonitorAgentStats.Stats.SflowSampleRecordBadLen,
			"sflow_samples_sent_for_detection":       ret.DtVisibilityMonitorAgentStats.Stats.SflowSamplesSentForDetection,
			"sflow_sample_record_invalid_layer2":     ret.DtVisibilityMonitorAgentStats.Stats.SflowSampleRecordInvalidLayer2,
			"sflow_sample_ipv6_hdr_parse_fail":       ret.DtVisibilityMonitorAgentStats.Stats.SflowSampleIpv6HdrParseFail,
			"sflow_disabled":                         ret.DtVisibilityMonitorAgentStats.Stats.SflowDisabled,
			"netflow_disabled":                       ret.DtVisibilityMonitorAgentStats.Stats.NetflowDisabled,
			"netflow_v5_packets_received":            ret.DtVisibilityMonitorAgentStats.Stats.NetflowV5PacketsReceived,
			"netflow_v5_samples_received":            ret.DtVisibilityMonitorAgentStats.Stats.NetflowV5SamplesReceived,
			"netflow_v5_samples_sent_for_detection":  ret.DtVisibilityMonitorAgentStats.Stats.NetflowV5SamplesSentForDetection,
			"netflow_v5_sample_records_bad_len":      ret.DtVisibilityMonitorAgentStats.Stats.NetflowV5SampleRecordsBadLen,
			"netflow_v5_max_records_exceed":          ret.DtVisibilityMonitorAgentStats.Stats.NetflowV5MaxRecordsExceed,
			"netflow_v9_packets_received":            ret.DtVisibilityMonitorAgentStats.Stats.NetflowV9PacketsReceived,
			"netflow_v9_samples_received":            ret.DtVisibilityMonitorAgentStats.Stats.NetflowV9SamplesReceived,
			"netflow_v9_samples_sent_for_detection":  ret.DtVisibilityMonitorAgentStats.Stats.NetflowV9SamplesSentForDetection,
			"netflow_v9_sample_records_bad_len":      ret.DtVisibilityMonitorAgentStats.Stats.NetflowV9SampleRecordsBadLen,
			"netflow_v9_max_records_exceed":          ret.DtVisibilityMonitorAgentStats.Stats.NetflowV9MaxRecordsExceed,
			"netflow_v10_packets_received":           ret.DtVisibilityMonitorAgentStats.Stats.NetflowV10PacketsReceived,
			"netflow_v10_samples_received":           ret.DtVisibilityMonitorAgentStats.Stats.NetflowV10SamplesReceived,
			"netflow_v10_samples_sent_for_detection": ret.DtVisibilityMonitorAgentStats.Stats.NetflowV10SamplesSentForDetection,
			"netflow_v10_sample_records_bad_len":     ret.DtVisibilityMonitorAgentStats.Stats.NetflowV10SampleRecordsBadLen,
			"netflow_v10_max_records_exceed":         ret.DtVisibilityMonitorAgentStats.Stats.NetflowV10MaxRecordsExceed,
			"netflow_tcp_sample_received":            ret.DtVisibilityMonitorAgentStats.Stats.NetflowTcpSampleReceived,
			"netflow_udp_sample_received":            ret.DtVisibilityMonitorAgentStats.Stats.NetflowUdpSampleReceived,
			"netflow_icmp_sample_received":           ret.DtVisibilityMonitorAgentStats.Stats.NetflowIcmpSampleReceived,
			"netflow_other_sample_received":          ret.DtVisibilityMonitorAgentStats.Stats.NetflowOtherSampleReceived,
			"netflow_record_copy_oom_error":          ret.DtVisibilityMonitorAgentStats.Stats.NetflowRecordCopyOomError,
			"netflow_record_rse_invalid":             ret.DtVisibilityMonitorAgentStats.Stats.NetflowRecordRseInvalid,
			"netflow_sample_flow_dur_error":          ret.DtVisibilityMonitorAgentStats.Stats.NetflowSampleFlowDurError,
		},
	}
}

func getObjectVisibilityMonitorAgentStatsStats(d []interface{}) edpt.VisibilityMonitorAgentStatsStats {

	count1 := len(d)
	var ret edpt.VisibilityMonitorAgentStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SflowPacketsReceived = in["sflow_packets_received"].(int)
		ret.SflowSamplesReceived = in["sflow_samples_received"].(int)
		ret.SflowSamplesBadLen = in["sflow_samples_bad_len"].(int)
		ret.SflowSamplesNonStd = in["sflow_samples_non_std"].(int)
		ret.SflowSamplesSkipped = in["sflow_samples_skipped"].(int)
		ret.SflowSampleRecordBadLen = in["sflow_sample_record_bad_len"].(int)
		ret.SflowSamplesSentForDetection = in["sflow_samples_sent_for_detection"].(int)
		ret.SflowSampleRecordInvalidLayer2 = in["sflow_sample_record_invalid_layer2"].(int)
		ret.SflowSampleIpv6HdrParseFail = in["sflow_sample_ipv6_hdr_parse_fail"].(int)
		ret.SflowDisabled = in["sflow_disabled"].(int)
		ret.NetflowDisabled = in["netflow_disabled"].(int)
		ret.NetflowV5PacketsReceived = in["netflow_v5_packets_received"].(int)
		ret.NetflowV5SamplesReceived = in["netflow_v5_samples_received"].(int)
		ret.NetflowV5SamplesSentForDetection = in["netflow_v5_samples_sent_for_detection"].(int)
		ret.NetflowV5SampleRecordsBadLen = in["netflow_v5_sample_records_bad_len"].(int)
		ret.NetflowV5MaxRecordsExceed = in["netflow_v5_max_records_exceed"].(int)
		ret.NetflowV9PacketsReceived = in["netflow_v9_packets_received"].(int)
		ret.NetflowV9SamplesReceived = in["netflow_v9_samples_received"].(int)
		ret.NetflowV9SamplesSentForDetection = in["netflow_v9_samples_sent_for_detection"].(int)
		ret.NetflowV9SampleRecordsBadLen = in["netflow_v9_sample_records_bad_len"].(int)
		ret.NetflowV9MaxRecordsExceed = in["netflow_v9_max_records_exceed"].(int)
		ret.NetflowV10PacketsReceived = in["netflow_v10_packets_received"].(int)
		ret.NetflowV10SamplesReceived = in["netflow_v10_samples_received"].(int)
		ret.NetflowV10SamplesSentForDetection = in["netflow_v10_samples_sent_for_detection"].(int)
		ret.NetflowV10SampleRecordsBadLen = in["netflow_v10_sample_records_bad_len"].(int)
		ret.NetflowV10MaxRecordsExceed = in["netflow_v10_max_records_exceed"].(int)
		ret.NetflowTcpSampleReceived = in["netflow_tcp_sample_received"].(int)
		ret.NetflowUdpSampleReceived = in["netflow_udp_sample_received"].(int)
		ret.NetflowIcmpSampleReceived = in["netflow_icmp_sample_received"].(int)
		ret.NetflowOtherSampleReceived = in["netflow_other_sample_received"].(int)
		ret.NetflowRecordCopyOomError = in["netflow_record_copy_oom_error"].(int)
		ret.NetflowRecordRseInvalid = in["netflow_record_rse_invalid"].(int)
		ret.NetflowSampleFlowDurError = in["netflow_sample_flow_dur_error"].(int)
	}
	return ret
}

func dataToEndpointVisibilityMonitorAgentStats(d *schema.ResourceData) edpt.VisibilityMonitorAgentStats {
	var ret edpt.VisibilityMonitorAgentStats

	ret.AgentName = d.Get("agent_name").(string)

	ret.Stats = getObjectVisibilityMonitorAgentStatsStats(d.Get("stats").([]interface{}))
	return ret
}
