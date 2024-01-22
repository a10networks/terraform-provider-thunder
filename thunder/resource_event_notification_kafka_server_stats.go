package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventNotificationKafkaServerStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_event_notification_kafka_server_stats`: Statistics for the object server\n\n__PLACEHOLDER__",
		ReadContext: resourceEventNotificationKafkaServerStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pr_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR logs sent",
						},
						"avro_device_status_topic": {
							Type: schema.TypeInt, Optional: true, Description: "Device Status Metrics sent",
						},
						"avro_partition_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "Partition Metrics sent",
						},
						"avro_generic_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Generic Metrics sent",
						},
						"pr_acos_harmony_topic_enqueue_err": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR dropped,enq error on acos queues",
						},
						"pr_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR dropped,enq error analytics queues",
						},
						"avro_generic_failed_encoding": {
							Type: schema.TypeInt, Optional: true, Description: "Generic Metrics dropped,encoding error",
						},
						"avro_generic_failed_sending": {
							Type: schema.TypeInt, Optional: true, Description: "Generic Metrics dropped,sending failure",
						},
						"avro_device_status_topic_enqueue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Device Status dropped,enq error on acos queues",
						},
						"avro_device_status_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Device Status dropped,enq error analytics queues",
						},
						"avro_partition_metrics_topic_enqueue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Part metrics dropped,enq error on acos queues",
						},
						"avro_partition_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Part metrics dropped,enq error analytics queues",
						},
						"kafka_unknown_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown type dropped,enq error analytics queues",
						},
						"kafka_broker_down": {
							Type: schema.TypeInt, Optional: true, Description: "Messages dropped,analytics down",
						},
						"kafka_queue_full_err": {
							Type: schema.TypeInt, Optional: true, Description: "Messages dropped,acos analytics queue full",
						},
						"pr_throttle_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR dropped,log throttling",
						},
						"pr_not_allowed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR dropped, not allowed to be sent",
						},
						"pr_be_ttfb_anomaly": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR back-end ttfb is negative",
						},
						"pr_be_ttlb_anomaly": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR back-end ttlb is negative",
						},
						"pr_in_latency_threshold_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR on latency threshold exceeded",
						},
						"pr_out_latency_threshold_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR out latency threshold exceeded",
						},
						"pr_out_latency_anomaly": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR out latency negative",
						},
						"pr_in_latency_anomaly": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PR on latency negative",
						},
						"kafka_topic_error": {
							Type: schema.TypeInt, Optional: true, Description: "Module not supported by analytics",
						},
						"pc_encoding_failed": {
							Type: schema.TypeInt, Optional: true, Description: "L4 PC logs dropped,encoding error",
						},
						"pc_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "L4 PC logs sent",
						},
						"pc_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "L4 PC logs dropped,enq error analytics queues",
						},
						"cgn_pc_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "CGN PC logs sent",
						},
						"cgn_pc_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "CGN PC logs dropped,enq error analytics queues",
						},
						"cgn_pe_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "CGN PE logs sent",
						},
						"cgn_pe_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "CGN PE logs dropped,enq error analytics queues",
						},
						"fw_pc_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "FW PC logs sent",
						},
						"fw_pc_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "FW PC logs dropped,enq error analytics queues",
						},
						"fw_deny_pc_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "FW DENY PC logs sent",
						},
						"fw_deny_pc_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "FW DENY PC logs dropped,enq error analytics queues",
						},
						"fw_rst_pc_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "FW RST PC logs sent",
						},
						"fw_rst_pc_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "FW RST PC logs dropped,enq error analytics queues",
						},
						"cgn_summary_error_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "CGN PE logs sent",
						},
						"cgn_summary_error_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "CGN PE logs dropped,enq error analytics queues",
						},
						"rule_set_application_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "AppFW metrics sent",
						},
						"rule_set_application_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "AppFW metrics dropped,enq error analytics queues",
						},
						"slb_ssl_stats_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "SSL metrics sent",
						},
						"slb_ssl_stats_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "SSL metrics dropped,enq error analytics queues",
						},
						"slb_client_ssl_counters_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "Client SSL metrics sent",
						},
						"slb_client_ssl_counters_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Cilent SSL metrics dropped,enq error analytics qs",
						},
						"slb_server_ssl_counters_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL metrics sent",
						},
						"slb_server_ssl_counters_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Server SSL metrics dropped,enq error analytics qs",
						},
						"pc_throttle_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L4 PC logs dropped,throttling",
						},
						"metrics_dropped_pt_missing": {
							Type: schema.TypeInt, Optional: true, Description: "Metrics dropped,missing partition tenant mapping",
						},
						"ssli_pc_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "SSLi PC topic counter from acos to harmony",
						},
						"ssli_pc_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "SSLi PC topic to harmony dequeue error",
						},
						"ssli_pe_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "SSLi PE topic counter from acos to harmony",
						},
						"ssli_pe_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "SSLi PE topic to harmony dequeue error",
						},
						"analytics_bus_restart": {
							Type: schema.TypeInt, Optional: true, Description: "Analytics bus restart count",
						},
						"waf_learn_pr_topic": {
							Type: schema.TypeInt, Optional: true, Description: "WAF learn topic counter",
						},
						"waf_learn_pr_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "WAF learn metrics dropped,enq error analytics qs",
						},
						"waf_events_topic": {
							Type: schema.TypeInt, Optional: true, Description: "WAF events topic counter",
						},
						"waf_events_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "WAF events metrics dropped,enq error analytics qs",
						},
						"visibility_topn_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "Visibility TopN sent",
						},
						"visibility_topn_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Visibility TopN metrics dropped,enq error analytics qs",
						},
						"hc_logs_sent_to_master": {
							Type: schema.TypeInt, Optional: true, Description: "HC logs sent to master",
						},
						"hc_logs_received_from_blade": {
							Type: schema.TypeInt, Optional: true, Description: "HC logs received from blade",
						},
						"hc_oper_sent_to_master": {
							Type: schema.TypeInt, Optional: true, Description: "HC oper to master",
						},
						"hc_oper_received_from_blade": {
							Type: schema.TypeInt, Optional: true, Description: "HC oper received from blade",
						},
						"hc_counters_sent_to_master": {
							Type: schema.TypeInt, Optional: true, Description: "HC counters sent to master",
						},
						"hc_counters_received_from_blade": {
							Type: schema.TypeInt, Optional: true, Description: "HC counters received from blade",
						},
						"hc_counters_dropped_from_blade": {
							Type: schema.TypeInt, Optional: true, Description: "HC counters dropped from blade (uuid or size mismatch)",
						},
						"pe_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PE logs sent",
						},
						"pe_acos_harmony_topic_enqueue_err": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PE dropped,enq error on acos queues",
						},
						"pe_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "L7 PE dropped,enq error analytics queues",
						},
						"vpn_ipsec_sa_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "IPSec SA metrics sent",
						},
						"vpn_ipsec_sa_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "IPSec SA metrics dropped,enq error analytics qs",
						},
						"vpn_ike_gateway_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "IKE gateway metrics sent",
						},
						"vpn_ike_gateway_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "IKE gateway metrics dropped,enq error analytics qs",
						},
						"vpn_stats_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "VPN STATS metrics sent",
						},
						"vpn_stats_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "VPN STATS metrics dropped,enq error analytics qs",
						},
						"cgn_port_usage_hstgrm_acos_harmony_topic": {
							Type: schema.TypeInt, Optional: true, Description: "CGN Port Usage Histogram HC Export",
						},
						"cgn_port_usage_hstgrm_acos_harmony_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "CGN Port Usage Histogram HC Export Failed",
						},
						"avro_system_env_topic": {
							Type: schema.TypeInt, Optional: true, Description: "System environment sent",
						},
						"avro_system_env_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "System Environmet dropped,enq error analytics queues",
						},
						"cert_pinning_list_topic": {
							Type: schema.TypeInt, Optional: true, Description: "Cert-pinning candidate list sent",
						},
						"cert_pinning_list_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "Cert-pinning candidate list dropped,enq error analytics queues",
						},
						"ngwaf_hc_ep_topic": {
							Type: schema.TypeInt, Optional: true, Description: "NGWAF HC PE export",
						},
						"ngwaf_hc_ep_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "NGWAF HC PE export failed",
						},
						"ngwaf_hc_metrics_topic": {
							Type: schema.TypeInt, Optional: true, Description: "NGWAF HC metrics export",
						},
						"ngwaf_hc_metrics_topic_dequeue_err": {
							Type: schema.TypeInt, Optional: true, Description: "NGWAF HC metrics export failed",
						},
					},
				},
			},
		},
	}
}

func resourceEventNotificationKafkaServerStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceEventNotificationKafkaServerStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointEventNotificationKafkaServerStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		EventNotificationKafkaServerStatsStats := setObjectEventNotificationKafkaServerStatsStats(res)
		d.Set("stats", EventNotificationKafkaServerStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectEventNotificationKafkaServerStatsStats(ret edpt.DataEventNotificationKafkaServerStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"pr_acos_harmony_topic":                                ret.DtEventNotificationKafkaServerStats.Stats.PrAcosHarmonyTopic,
			"avro_device_status_topic":                             ret.DtEventNotificationKafkaServerStats.Stats.AvroDeviceStatusTopic,
			"avro_partition_metrics_topic":                         ret.DtEventNotificationKafkaServerStats.Stats.AvroPartitionMetricsTopic,
			"avro_generic_sent":                                    ret.DtEventNotificationKafkaServerStats.Stats.AvroGenericSent,
			"pr_acos_harmony_topic_enqueue_err":                    ret.DtEventNotificationKafkaServerStats.Stats.PrAcosHarmonyTopicEnqueueErr,
			"pr_acos_harmony_topic_dequeue_err":                    ret.DtEventNotificationKafkaServerStats.Stats.PrAcosHarmonyTopicDequeueErr,
			"avro_generic_failed_encoding":                         ret.DtEventNotificationKafkaServerStats.Stats.AvroGenericFailedEncoding,
			"avro_generic_failed_sending":                          ret.DtEventNotificationKafkaServerStats.Stats.AvroGenericFailedSending,
			"avro_device_status_topic_enqueue_err":                 ret.DtEventNotificationKafkaServerStats.Stats.AvroDeviceStatusTopicEnqueueErr,
			"avro_device_status_topic_dequeue_err":                 ret.DtEventNotificationKafkaServerStats.Stats.AvroDeviceStatusTopicDequeueErr,
			"avro_partition_metrics_topic_enqueue_err":             ret.DtEventNotificationKafkaServerStats.Stats.AvroPartitionMetricsTopicEnqueueErr,
			"avro_partition_metrics_topic_dequeue_err":             ret.DtEventNotificationKafkaServerStats.Stats.AvroPartitionMetricsTopicDequeueErr,
			"kafka_unknown_topic_dequeue_err":                      ret.DtEventNotificationKafkaServerStats.Stats.KafkaUnknownTopicDequeueErr,
			"kafka_broker_down":                                    ret.DtEventNotificationKafkaServerStats.Stats.KafkaBrokerDown,
			"kafka_queue_full_err":                                 ret.DtEventNotificationKafkaServerStats.Stats.KafkaQueueFullErr,
			"pr_throttle_drop":                                     ret.DtEventNotificationKafkaServerStats.Stats.PrThrottleDrop,
			"pr_not_allowed_drop":                                  ret.DtEventNotificationKafkaServerStats.Stats.PrNotAllowedDrop,
			"pr_be_ttfb_anomaly":                                   ret.DtEventNotificationKafkaServerStats.Stats.PrBeTtfbAnomaly,
			"pr_be_ttlb_anomaly":                                   ret.DtEventNotificationKafkaServerStats.Stats.PrBeTtlbAnomaly,
			"pr_in_latency_threshold_exceed":                       ret.DtEventNotificationKafkaServerStats.Stats.PrInLatencyThresholdExceed,
			"pr_out_latency_threshold_exceed":                      ret.DtEventNotificationKafkaServerStats.Stats.PrOutLatencyThresholdExceed,
			"pr_out_latency_anomaly":                               ret.DtEventNotificationKafkaServerStats.Stats.PrOutLatencyAnomaly,
			"pr_in_latency_anomaly":                                ret.DtEventNotificationKafkaServerStats.Stats.PrInLatencyAnomaly,
			"kafka_topic_error":                                    ret.DtEventNotificationKafkaServerStats.Stats.KafkaTopicError,
			"pc_encoding_failed":                                   ret.DtEventNotificationKafkaServerStats.Stats.PcEncodingFailed,
			"pc_acos_harmony_topic":                                ret.DtEventNotificationKafkaServerStats.Stats.PcAcosHarmonyTopic,
			"pc_acos_harmony_topic_dequeue_err":                    ret.DtEventNotificationKafkaServerStats.Stats.PcAcosHarmonyTopicDequeueErr,
			"cgn_pc_acos_harmony_topic":                            ret.DtEventNotificationKafkaServerStats.Stats.CgnPcAcosHarmonyTopic,
			"cgn_pc_acos_harmony_topic_dequeue_err":                ret.DtEventNotificationKafkaServerStats.Stats.CgnPcAcosHarmonyTopicDequeueErr,
			"cgn_pe_acos_harmony_topic":                            ret.DtEventNotificationKafkaServerStats.Stats.CgnPeAcosHarmonyTopic,
			"cgn_pe_acos_harmony_topic_dequeue_err":                ret.DtEventNotificationKafkaServerStats.Stats.CgnPeAcosHarmonyTopicDequeueErr,
			"fw_pc_acos_harmony_topic":                             ret.DtEventNotificationKafkaServerStats.Stats.FwPcAcosHarmonyTopic,
			"fw_pc_acos_harmony_topic_dequeue_err":                 ret.DtEventNotificationKafkaServerStats.Stats.FwPcAcosHarmonyTopicDequeueErr,
			"fw_deny_pc_acos_harmony_topic":                        ret.DtEventNotificationKafkaServerStats.Stats.FwDenyPcAcosHarmonyTopic,
			"fw_deny_pc_acos_harmony_topic_dequeue_err":            ret.DtEventNotificationKafkaServerStats.Stats.FwDenyPcAcosHarmonyTopicDequeueErr,
			"fw_rst_pc_acos_harmony_topic":                         ret.DtEventNotificationKafkaServerStats.Stats.FwRstPcAcosHarmonyTopic,
			"fw_rst_pc_acos_harmony_topic_dequeue_err":             ret.DtEventNotificationKafkaServerStats.Stats.FwRstPcAcosHarmonyTopicDequeueErr,
			"cgn_summary_error_acos_harmony_topic":                 ret.DtEventNotificationKafkaServerStats.Stats.CgnSummaryErrorAcosHarmonyTopic,
			"cgn_summary_error_acos_harmony_topic_dequeue_err":     ret.DtEventNotificationKafkaServerStats.Stats.CgnSummaryErrorAcosHarmonyTopicDequeueErr,
			"rule_set_application_metrics_topic":                   ret.DtEventNotificationKafkaServerStats.Stats.RuleSetApplicationMetricsTopic,
			"rule_set_application_metrics_topic_dequeue_err":       ret.DtEventNotificationKafkaServerStats.Stats.RuleSetApplicationMetricsTopicDequeueErr,
			"slb_ssl_stats_metrics_topic":                          ret.DtEventNotificationKafkaServerStats.Stats.SlbSslStatsMetricsTopic,
			"slb_ssl_stats_metrics_topic_dequeue_err":              ret.DtEventNotificationKafkaServerStats.Stats.SlbSslStatsMetricsTopicDequeueErr,
			"slb_client_ssl_counters_metrics_topic":                ret.DtEventNotificationKafkaServerStats.Stats.SlbClientSslCountersMetricsTopic,
			"slb_client_ssl_counters_metrics_topic_dequeue_err":    ret.DtEventNotificationKafkaServerStats.Stats.SlbClientSslCountersMetricsTopicDequeueErr,
			"slb_server_ssl_counters_metrics_topic":                ret.DtEventNotificationKafkaServerStats.Stats.SlbServerSslCountersMetricsTopic,
			"slb_server_ssl_counters_metrics_topic_dequeue_err":    ret.DtEventNotificationKafkaServerStats.Stats.SlbServerSslCountersMetricsTopicDequeueErr,
			"pc_throttle_drop":                                     ret.DtEventNotificationKafkaServerStats.Stats.PcThrottleDrop,
			"metrics_dropped_pt_missing":                           ret.DtEventNotificationKafkaServerStats.Stats.MetricsDroppedPtMissing,
			"ssli_pc_acos_harmony_topic":                           ret.DtEventNotificationKafkaServerStats.Stats.SsliPcAcosHarmonyTopic,
			"ssli_pc_acos_harmony_topic_dequeue_err":               ret.DtEventNotificationKafkaServerStats.Stats.SsliPcAcosHarmonyTopicDequeueErr,
			"ssli_pe_acos_harmony_topic":                           ret.DtEventNotificationKafkaServerStats.Stats.SsliPeAcosHarmonyTopic,
			"ssli_pe_acos_harmony_topic_dequeue_err":               ret.DtEventNotificationKafkaServerStats.Stats.SsliPeAcosHarmonyTopicDequeueErr,
			"analytics_bus_restart":                                ret.DtEventNotificationKafkaServerStats.Stats.AnalyticsBusRestart,
			"waf_learn_pr_topic":                                   ret.DtEventNotificationKafkaServerStats.Stats.WafLearnPrTopic,
			"waf_learn_pr_topic_dequeue_err":                       ret.DtEventNotificationKafkaServerStats.Stats.WafLearnPrTopicDequeueErr,
			"waf_events_topic":                                     ret.DtEventNotificationKafkaServerStats.Stats.WafEventsTopic,
			"waf_events_topic_dequeue_err":                         ret.DtEventNotificationKafkaServerStats.Stats.WafEventsTopicDequeueErr,
			"visibility_topn_harmony_topic":                        ret.DtEventNotificationKafkaServerStats.Stats.VisibilityTopnHarmonyTopic,
			"visibility_topn_harmony_topic_dequeue_err":            ret.DtEventNotificationKafkaServerStats.Stats.VisibilityTopnHarmonyTopicDequeueErr,
			"hc_logs_sent_to_master":                               ret.DtEventNotificationKafkaServerStats.Stats.HcLogsSentToMaster,
			"hc_logs_received_from_blade":                          ret.DtEventNotificationKafkaServerStats.Stats.HcLogsReceivedFromBlade,
			"hc_oper_sent_to_master":                               ret.DtEventNotificationKafkaServerStats.Stats.HcOperSentToMaster,
			"hc_oper_received_from_blade":                          ret.DtEventNotificationKafkaServerStats.Stats.HcOperReceivedFromBlade,
			"hc_counters_sent_to_master":                           ret.DtEventNotificationKafkaServerStats.Stats.HcCountersSentToMaster,
			"hc_counters_received_from_blade":                      ret.DtEventNotificationKafkaServerStats.Stats.HcCountersReceivedFromBlade,
			"hc_counters_dropped_from_blade":                       ret.DtEventNotificationKafkaServerStats.Stats.HcCountersDroppedFromBlade,
			"pe_acos_harmony_topic":                                ret.DtEventNotificationKafkaServerStats.Stats.PeAcosHarmonyTopic,
			"pe_acos_harmony_topic_enqueue_err":                    ret.DtEventNotificationKafkaServerStats.Stats.PeAcosHarmonyTopicEnqueueErr,
			"pe_acos_harmony_topic_dequeue_err":                    ret.DtEventNotificationKafkaServerStats.Stats.PeAcosHarmonyTopicDequeueErr,
			"vpn_ipsec_sa_metrics_topic":                           ret.DtEventNotificationKafkaServerStats.Stats.VpnIpsecSaMetricsTopic,
			"vpn_ipsec_sa_metrics_topic_dequeue_err":               ret.DtEventNotificationKafkaServerStats.Stats.VpnIpsecSaMetricsTopicDequeueErr,
			"vpn_ike_gateway_metrics_topic":                        ret.DtEventNotificationKafkaServerStats.Stats.VpnIkeGatewayMetricsTopic,
			"vpn_ike_gateway_metrics_topic_dequeue_err":            ret.DtEventNotificationKafkaServerStats.Stats.VpnIkeGatewayMetricsTopicDequeueErr,
			"vpn_stats_metrics_topic":                              ret.DtEventNotificationKafkaServerStats.Stats.VpnStatsMetricsTopic,
			"vpn_stats_metrics_topic_dequeue_err":                  ret.DtEventNotificationKafkaServerStats.Stats.VpnStatsMetricsTopicDequeueErr,
			"cgn_port_usage_hstgrm_acos_harmony_topic":             ret.DtEventNotificationKafkaServerStats.Stats.CgnPortUsageHstgrmAcosHarmonyTopic,
			"cgn_port_usage_hstgrm_acos_harmony_topic_dequeue_err": ret.DtEventNotificationKafkaServerStats.Stats.CgnPortUsageHstgrmAcosHarmonyTopicDequeueErr,
			"avro_system_env_topic":                                ret.DtEventNotificationKafkaServerStats.Stats.AvroSystemEnvTopic,
			"avro_system_env_dequeue_err":                          ret.DtEventNotificationKafkaServerStats.Stats.AvroSystemEnvDequeueErr,
			"cert_pinning_list_topic":                              ret.DtEventNotificationKafkaServerStats.Stats.CertPinningListTopic,
			"cert_pinning_list_topic_dequeue_err":                  ret.DtEventNotificationKafkaServerStats.Stats.CertPinningListTopicDequeueErr,
			"ngwaf_hc_ep_topic":                                    ret.DtEventNotificationKafkaServerStats.Stats.NgwafHcEpTopic,
			"ngwaf_hc_ep_topic_dequeue_err":                        ret.DtEventNotificationKafkaServerStats.Stats.NgwafHcEpTopicDequeueErr,
			"ngwaf_hc_metrics_topic":                               ret.DtEventNotificationKafkaServerStats.Stats.NgwafHcMetricsTopic,
			"ngwaf_hc_metrics_topic_dequeue_err":                   ret.DtEventNotificationKafkaServerStats.Stats.NgwafHcMetricsTopicDequeueErr,
		},
	}
}

func getObjectEventNotificationKafkaServerStatsStats(d []interface{}) edpt.EventNotificationKafkaServerStatsStats {

	count1 := len(d)
	var ret edpt.EventNotificationKafkaServerStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PrAcosHarmonyTopic = in["pr_acos_harmony_topic"].(int)
		ret.AvroDeviceStatusTopic = in["avro_device_status_topic"].(int)
		ret.AvroPartitionMetricsTopic = in["avro_partition_metrics_topic"].(int)
		ret.AvroGenericSent = in["avro_generic_sent"].(int)
		ret.PrAcosHarmonyTopicEnqueueErr = in["pr_acos_harmony_topic_enqueue_err"].(int)
		ret.PrAcosHarmonyTopicDequeueErr = in["pr_acos_harmony_topic_dequeue_err"].(int)
		ret.AvroGenericFailedEncoding = in["avro_generic_failed_encoding"].(int)
		ret.AvroGenericFailedSending = in["avro_generic_failed_sending"].(int)
		ret.AvroDeviceStatusTopicEnqueueErr = in["avro_device_status_topic_enqueue_err"].(int)
		ret.AvroDeviceStatusTopicDequeueErr = in["avro_device_status_topic_dequeue_err"].(int)
		ret.AvroPartitionMetricsTopicEnqueueErr = in["avro_partition_metrics_topic_enqueue_err"].(int)
		ret.AvroPartitionMetricsTopicDequeueErr = in["avro_partition_metrics_topic_dequeue_err"].(int)
		ret.KafkaUnknownTopicDequeueErr = in["kafka_unknown_topic_dequeue_err"].(int)
		ret.KafkaBrokerDown = in["kafka_broker_down"].(int)
		ret.KafkaQueueFullErr = in["kafka_queue_full_err"].(int)
		ret.PrThrottleDrop = in["pr_throttle_drop"].(int)
		ret.PrNotAllowedDrop = in["pr_not_allowed_drop"].(int)
		ret.PrBeTtfbAnomaly = in["pr_be_ttfb_anomaly"].(int)
		ret.PrBeTtlbAnomaly = in["pr_be_ttlb_anomaly"].(int)
		ret.PrInLatencyThresholdExceed = in["pr_in_latency_threshold_exceed"].(int)
		ret.PrOutLatencyThresholdExceed = in["pr_out_latency_threshold_exceed"].(int)
		ret.PrOutLatencyAnomaly = in["pr_out_latency_anomaly"].(int)
		ret.PrInLatencyAnomaly = in["pr_in_latency_anomaly"].(int)
		ret.KafkaTopicError = in["kafka_topic_error"].(int)
		ret.PcEncodingFailed = in["pc_encoding_failed"].(int)
		ret.PcAcosHarmonyTopic = in["pc_acos_harmony_topic"].(int)
		ret.PcAcosHarmonyTopicDequeueErr = in["pc_acos_harmony_topic_dequeue_err"].(int)
		ret.CgnPcAcosHarmonyTopic = in["cgn_pc_acos_harmony_topic"].(int)
		ret.CgnPcAcosHarmonyTopicDequeueErr = in["cgn_pc_acos_harmony_topic_dequeue_err"].(int)
		ret.CgnPeAcosHarmonyTopic = in["cgn_pe_acos_harmony_topic"].(int)
		ret.CgnPeAcosHarmonyTopicDequeueErr = in["cgn_pe_acos_harmony_topic_dequeue_err"].(int)
		ret.FwPcAcosHarmonyTopic = in["fw_pc_acos_harmony_topic"].(int)
		ret.FwPcAcosHarmonyTopicDequeueErr = in["fw_pc_acos_harmony_topic_dequeue_err"].(int)
		ret.FwDenyPcAcosHarmonyTopic = in["fw_deny_pc_acos_harmony_topic"].(int)
		ret.FwDenyPcAcosHarmonyTopicDequeueErr = in["fw_deny_pc_acos_harmony_topic_dequeue_err"].(int)
		ret.FwRstPcAcosHarmonyTopic = in["fw_rst_pc_acos_harmony_topic"].(int)
		ret.FwRstPcAcosHarmonyTopicDequeueErr = in["fw_rst_pc_acos_harmony_topic_dequeue_err"].(int)
		ret.CgnSummaryErrorAcosHarmonyTopic = in["cgn_summary_error_acos_harmony_topic"].(int)
		ret.CgnSummaryErrorAcosHarmonyTopicDequeueErr = in["cgn_summary_error_acos_harmony_topic_dequeue_err"].(int)
		ret.RuleSetApplicationMetricsTopic = in["rule_set_application_metrics_topic"].(int)
		ret.RuleSetApplicationMetricsTopicDequeueErr = in["rule_set_application_metrics_topic_dequeue_err"].(int)
		ret.SlbSslStatsMetricsTopic = in["slb_ssl_stats_metrics_topic"].(int)
		ret.SlbSslStatsMetricsTopicDequeueErr = in["slb_ssl_stats_metrics_topic_dequeue_err"].(int)
		ret.SlbClientSslCountersMetricsTopic = in["slb_client_ssl_counters_metrics_topic"].(int)
		ret.SlbClientSslCountersMetricsTopicDequeueErr = in["slb_client_ssl_counters_metrics_topic_dequeue_err"].(int)
		ret.SlbServerSslCountersMetricsTopic = in["slb_server_ssl_counters_metrics_topic"].(int)
		ret.SlbServerSslCountersMetricsTopicDequeueErr = in["slb_server_ssl_counters_metrics_topic_dequeue_err"].(int)
		ret.PcThrottleDrop = in["pc_throttle_drop"].(int)
		ret.MetricsDroppedPtMissing = in["metrics_dropped_pt_missing"].(int)
		ret.SsliPcAcosHarmonyTopic = in["ssli_pc_acos_harmony_topic"].(int)
		ret.SsliPcAcosHarmonyTopicDequeueErr = in["ssli_pc_acos_harmony_topic_dequeue_err"].(int)
		ret.SsliPeAcosHarmonyTopic = in["ssli_pe_acos_harmony_topic"].(int)
		ret.SsliPeAcosHarmonyTopicDequeueErr = in["ssli_pe_acos_harmony_topic_dequeue_err"].(int)
		ret.AnalyticsBusRestart = in["analytics_bus_restart"].(int)
		ret.WafLearnPrTopic = in["waf_learn_pr_topic"].(int)
		ret.WafLearnPrTopicDequeueErr = in["waf_learn_pr_topic_dequeue_err"].(int)
		ret.WafEventsTopic = in["waf_events_topic"].(int)
		ret.WafEventsTopicDequeueErr = in["waf_events_topic_dequeue_err"].(int)
		ret.VisibilityTopnHarmonyTopic = in["visibility_topn_harmony_topic"].(int)
		ret.VisibilityTopnHarmonyTopicDequeueErr = in["visibility_topn_harmony_topic_dequeue_err"].(int)
		ret.HcLogsSentToMaster = in["hc_logs_sent_to_master"].(int)
		ret.HcLogsReceivedFromBlade = in["hc_logs_received_from_blade"].(int)
		ret.HcOperSentToMaster = in["hc_oper_sent_to_master"].(int)
		ret.HcOperReceivedFromBlade = in["hc_oper_received_from_blade"].(int)
		ret.HcCountersSentToMaster = in["hc_counters_sent_to_master"].(int)
		ret.HcCountersReceivedFromBlade = in["hc_counters_received_from_blade"].(int)
		ret.HcCountersDroppedFromBlade = in["hc_counters_dropped_from_blade"].(int)
		ret.PeAcosHarmonyTopic = in["pe_acos_harmony_topic"].(int)
		ret.PeAcosHarmonyTopicEnqueueErr = in["pe_acos_harmony_topic_enqueue_err"].(int)
		ret.PeAcosHarmonyTopicDequeueErr = in["pe_acos_harmony_topic_dequeue_err"].(int)
		ret.VpnIpsecSaMetricsTopic = in["vpn_ipsec_sa_metrics_topic"].(int)
		ret.VpnIpsecSaMetricsTopicDequeueErr = in["vpn_ipsec_sa_metrics_topic_dequeue_err"].(int)
		ret.VpnIkeGatewayMetricsTopic = in["vpn_ike_gateway_metrics_topic"].(int)
		ret.VpnIkeGatewayMetricsTopicDequeueErr = in["vpn_ike_gateway_metrics_topic_dequeue_err"].(int)
		ret.VpnStatsMetricsTopic = in["vpn_stats_metrics_topic"].(int)
		ret.VpnStatsMetricsTopicDequeueErr = in["vpn_stats_metrics_topic_dequeue_err"].(int)
		ret.CgnPortUsageHstgrmAcosHarmonyTopic = in["cgn_port_usage_hstgrm_acos_harmony_topic"].(int)
		ret.CgnPortUsageHstgrmAcosHarmonyTopicDequeueErr = in["cgn_port_usage_hstgrm_acos_harmony_topic_dequeue_err"].(int)
		ret.AvroSystemEnvTopic = in["avro_system_env_topic"].(int)
		ret.AvroSystemEnvDequeueErr = in["avro_system_env_dequeue_err"].(int)
		ret.CertPinningListTopic = in["cert_pinning_list_topic"].(int)
		ret.CertPinningListTopicDequeueErr = in["cert_pinning_list_topic_dequeue_err"].(int)
		ret.NgwafHcEpTopic = in["ngwaf_hc_ep_topic"].(int)
		ret.NgwafHcEpTopicDequeueErr = in["ngwaf_hc_ep_topic_dequeue_err"].(int)
		ret.NgwafHcMetricsTopic = in["ngwaf_hc_metrics_topic"].(int)
		ret.NgwafHcMetricsTopicDequeueErr = in["ngwaf_hc_metrics_topic_dequeue_err"].(int)
	}
	return ret
}

func dataToEndpointEventNotificationKafkaServerStats(d *schema.ResourceData) edpt.EventNotificationKafkaServerStats {
	var ret edpt.EventNotificationKafkaServerStats

	ret.Stats = getObjectEventNotificationKafkaServerStatsStats(d.Get("stats").([]interface{}))
	return ret
}
