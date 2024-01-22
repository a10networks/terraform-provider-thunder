package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionAgentStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_detection_agent_stats`: Statistics for the object agent\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDetectionAgentStatsRead,

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
						"netflow_v9_sample_flowset_bad_padding": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Sample Flowset Bad Padding",
						},
						"netflow_v9_max_records_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Sample Max Records Error",
						},
						"netflow_v9_template_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow v9 Template Not Found",
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
							Type: schema.TypeInt, Optional: true, Description: "Netflow Data Record Copy Fail, Local MEM size error",
						},
						"netflow_record_rse_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Data Record Reduced Size Invalid",
						},
						"netflow_sample_flow_dur_error": {
							Type: schema.TypeInt, Optional: true, Description: "Netflow Sample Flow Duration Error",
						},
						"flow_dst_entry_miss": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Destination Entry Lookup Failures",
						},
						"flow_ip_proto_or_port_miss": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Destination Service Lookup Failures",
						},
						"flow_detection_msgq_full": {
							Type: schema.TypeInt, Optional: true, Description: "Detection Message Enqueue Failures",
						},
						"flow_network_entry_miss": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Destination Network-object Entry Lookup Failures",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDetectionAgentStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionAgentStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionAgentStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDetectionAgentStatsStats := setObjectDdosDetectionAgentStatsStats(res)
		d.Set("stats", DdosDetectionAgentStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDetectionAgentStatsStats(ret edpt.DataDdosDetectionAgentStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"sflow_packets_received":                 ret.DtDdosDetectionAgentStats.Stats.SflowPacketsReceived,
			"sflow_samples_received":                 ret.DtDdosDetectionAgentStats.Stats.SflowSamplesReceived,
			"sflow_samples_bad_len":                  ret.DtDdosDetectionAgentStats.Stats.SflowSamplesBadLen,
			"sflow_samples_non_std":                  ret.DtDdosDetectionAgentStats.Stats.SflowSamplesNonStd,
			"sflow_samples_skipped":                  ret.DtDdosDetectionAgentStats.Stats.SflowSamplesSkipped,
			"sflow_sample_record_bad_len":            ret.DtDdosDetectionAgentStats.Stats.SflowSampleRecordBadLen,
			"sflow_samples_sent_for_detection":       ret.DtDdosDetectionAgentStats.Stats.SflowSamplesSentForDetection,
			"sflow_sample_record_invalid_layer2":     ret.DtDdosDetectionAgentStats.Stats.SflowSampleRecordInvalidLayer2,
			"sflow_sample_ipv6_hdr_parse_fail":       ret.DtDdosDetectionAgentStats.Stats.SflowSampleIpv6HdrParseFail,
			"sflow_disabled":                         ret.DtDdosDetectionAgentStats.Stats.SflowDisabled,
			"netflow_disabled":                       ret.DtDdosDetectionAgentStats.Stats.NetflowDisabled,
			"netflow_v5_packets_received":            ret.DtDdosDetectionAgentStats.Stats.NetflowV5PacketsReceived,
			"netflow_v5_samples_received":            ret.DtDdosDetectionAgentStats.Stats.NetflowV5SamplesReceived,
			"netflow_v5_samples_sent_for_detection":  ret.DtDdosDetectionAgentStats.Stats.NetflowV5SamplesSentForDetection,
			"netflow_v5_sample_records_bad_len":      ret.DtDdosDetectionAgentStats.Stats.NetflowV5SampleRecordsBadLen,
			"netflow_v5_max_records_exceed":          ret.DtDdosDetectionAgentStats.Stats.NetflowV5MaxRecordsExceed,
			"netflow_v9_packets_received":            ret.DtDdosDetectionAgentStats.Stats.NetflowV9PacketsReceived,
			"netflow_v9_samples_received":            ret.DtDdosDetectionAgentStats.Stats.NetflowV9SamplesReceived,
			"netflow_v9_samples_sent_for_detection":  ret.DtDdosDetectionAgentStats.Stats.NetflowV9SamplesSentForDetection,
			"netflow_v9_sample_records_bad_len":      ret.DtDdosDetectionAgentStats.Stats.NetflowV9SampleRecordsBadLen,
			"netflow_v9_sample_flowset_bad_padding":  ret.DtDdosDetectionAgentStats.Stats.NetflowV9SampleFlowsetBadPadding,
			"netflow_v9_max_records_exceed":          ret.DtDdosDetectionAgentStats.Stats.NetflowV9MaxRecordsExceed,
			"netflow_v9_template_not_found":          ret.DtDdosDetectionAgentStats.Stats.NetflowV9TemplateNotFound,
			"netflow_v10_packets_received":           ret.DtDdosDetectionAgentStats.Stats.NetflowV10PacketsReceived,
			"netflow_v10_samples_received":           ret.DtDdosDetectionAgentStats.Stats.NetflowV10SamplesReceived,
			"netflow_v10_samples_sent_for_detection": ret.DtDdosDetectionAgentStats.Stats.NetflowV10SamplesSentForDetection,
			"netflow_v10_sample_records_bad_len":     ret.DtDdosDetectionAgentStats.Stats.NetflowV10SampleRecordsBadLen,
			"netflow_v10_max_records_exceed":         ret.DtDdosDetectionAgentStats.Stats.NetflowV10MaxRecordsExceed,
			"netflow_tcp_sample_received":            ret.DtDdosDetectionAgentStats.Stats.NetflowTcpSampleReceived,
			"netflow_udp_sample_received":            ret.DtDdosDetectionAgentStats.Stats.NetflowUdpSampleReceived,
			"netflow_icmp_sample_received":           ret.DtDdosDetectionAgentStats.Stats.NetflowIcmpSampleReceived,
			"netflow_other_sample_received":          ret.DtDdosDetectionAgentStats.Stats.NetflowOtherSampleReceived,
			"netflow_record_copy_oom_error":          ret.DtDdosDetectionAgentStats.Stats.NetflowRecordCopyOomError,
			"netflow_record_rse_invalid":             ret.DtDdosDetectionAgentStats.Stats.NetflowRecordRseInvalid,
			"netflow_sample_flow_dur_error":          ret.DtDdosDetectionAgentStats.Stats.NetflowSampleFlowDurError,
			"flow_dst_entry_miss":                    ret.DtDdosDetectionAgentStats.Stats.FlowDstEntryMiss,
			"flow_ip_proto_or_port_miss":             ret.DtDdosDetectionAgentStats.Stats.FlowIpProtoOrPortMiss,
			"flow_detection_msgq_full":               ret.DtDdosDetectionAgentStats.Stats.FlowDetectionMsgqFull,
			"flow_network_entry_miss":                ret.DtDdosDetectionAgentStats.Stats.FlowNetworkEntryMiss,
		},
	}
}

func getObjectDdosDetectionAgentStatsStats(d []interface{}) edpt.DdosDetectionAgentStatsStats {

	count1 := len(d)
	var ret edpt.DdosDetectionAgentStatsStats
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
		ret.NetflowV9SampleFlowsetBadPadding = in["netflow_v9_sample_flowset_bad_padding"].(int)
		ret.NetflowV9MaxRecordsExceed = in["netflow_v9_max_records_exceed"].(int)
		ret.NetflowV9TemplateNotFound = in["netflow_v9_template_not_found"].(int)
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
		ret.FlowDstEntryMiss = in["flow_dst_entry_miss"].(int)
		ret.FlowIpProtoOrPortMiss = in["flow_ip_proto_or_port_miss"].(int)
		ret.FlowDetectionMsgqFull = in["flow_detection_msgq_full"].(int)
		ret.FlowNetworkEntryMiss = in["flow_network_entry_miss"].(int)
	}
	return ret
}

func dataToEndpointDdosDetectionAgentStats(d *schema.ResourceData) edpt.DdosDetectionAgentStats {
	var ret edpt.DdosDetectionAgentStats

	ret.AgentName = d.Get("agent_name").(string)

	ret.Stats = getObjectDdosDetectionAgentStatsStats(d.Get("stats").([]interface{}))
	return ret
}
