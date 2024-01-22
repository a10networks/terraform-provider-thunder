

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EventNotificationKafkaServerStats struct {
    
    Stats EventNotificationKafkaServerStatsStats `json:"stats"`

}
type DataEventNotificationKafkaServerStats struct {
    DtEventNotificationKafkaServerStats EventNotificationKafkaServerStats `json:"server"`
}


type EventNotificationKafkaServerStatsStats struct {
    PrAcosHarmonyTopic int `json:"pr-acos-harmony-topic"`
    AvroDeviceStatusTopic int `json:"avro-device-status-topic"`
    AvroPartitionMetricsTopic int `json:"avro-partition-metrics-topic"`
    AvroGenericSent int `json:"avro-generic-sent"`
    PrAcosHarmonyTopicEnqueueErr int `json:"pr-acos-harmony-topic-enqueue-err"`
    PrAcosHarmonyTopicDequeueErr int `json:"pr-acos-harmony-topic-dequeue-err"`
    AvroGenericFailedEncoding int `json:"avro-generic-failed-encoding"`
    AvroGenericFailedSending int `json:"avro-generic-failed-sending"`
    AvroDeviceStatusTopicEnqueueErr int `json:"avro-device-status-topic-enqueue-err"`
    AvroDeviceStatusTopicDequeueErr int `json:"avro-device-status-topic-dequeue-err"`
    AvroPartitionMetricsTopicEnqueueErr int `json:"avro-partition-metrics-topic-enqueue-err"`
    AvroPartitionMetricsTopicDequeueErr int `json:"avro-partition-metrics-topic-dequeue-err"`
    KafkaUnknownTopicDequeueErr int `json:"kafka-unknown-topic-dequeue-err"`
    KafkaBrokerDown int `json:"kafka-broker-down"`
    KafkaQueueFullErr int `json:"kafka-queue-full-err"`
    PrThrottleDrop int `json:"pr-throttle-drop"`
    PrNotAllowedDrop int `json:"pr-not-allowed-drop"`
    PrBeTtfbAnomaly int `json:"pr-be-ttfb-anomaly"`
    PrBeTtlbAnomaly int `json:"pr-be-ttlb-anomaly"`
    PrInLatencyThresholdExceed int `json:"pr-in-latency-threshold-exceed"`
    PrOutLatencyThresholdExceed int `json:"pr-out-latency-threshold-exceed"`
    PrOutLatencyAnomaly int `json:"pr-out-latency-anomaly"`
    PrInLatencyAnomaly int `json:"pr-in-latency-anomaly"`
    KafkaTopicError int `json:"kafka-topic-error"`
    PcEncodingFailed int `json:"pc-encoding-failed"`
    PcAcosHarmonyTopic int `json:"pc-acos-harmony-topic"`
    PcAcosHarmonyTopicDequeueErr int `json:"pc-acos-harmony-topic-dequeue-err"`
    CgnPcAcosHarmonyTopic int `json:"cgn-pc-acos-harmony-topic"`
    CgnPcAcosHarmonyTopicDequeueErr int `json:"cgn-pc-acos-harmony-topic-dequeue-err"`
    CgnPeAcosHarmonyTopic int `json:"cgn-pe-acos-harmony-topic"`
    CgnPeAcosHarmonyTopicDequeueErr int `json:"cgn-pe-acos-harmony-topic-dequeue-err"`
    FwPcAcosHarmonyTopic int `json:"fw-pc-acos-harmony-topic"`
    FwPcAcosHarmonyTopicDequeueErr int `json:"fw-pc-acos-harmony-topic-dequeue-err"`
    FwDenyPcAcosHarmonyTopic int `json:"fw-deny-pc-acos-harmony-topic"`
    FwDenyPcAcosHarmonyTopicDequeueErr int `json:"fw-deny-pc-acos-harmony-topic-dequeue-err"`
    FwRstPcAcosHarmonyTopic int `json:"fw-rst-pc-acos-harmony-topic"`
    FwRstPcAcosHarmonyTopicDequeueErr int `json:"fw-rst-pc-acos-harmony-topic-dequeue-err"`
    CgnSummaryErrorAcosHarmonyTopic int `json:"cgn-summary-error-acos-harmony-topic"`
    CgnSummaryErrorAcosHarmonyTopicDequeueErr int `json:"cgn-summary-error-acos-harmony-topic-dequeue-err"`
    RuleSetApplicationMetricsTopic int `json:"rule-set-application-metrics-topic"`
    RuleSetApplicationMetricsTopicDequeueErr int `json:"rule-set-application-metrics-topic-dequeue-err"`
    SlbSslStatsMetricsTopic int `json:"slb-ssl-stats-metrics-topic"`
    SlbSslStatsMetricsTopicDequeueErr int `json:"slb-ssl-stats-metrics-topic-dequeue-err"`
    SlbClientSslCountersMetricsTopic int `json:"slb-client-ssl-counters-metrics-topic"`
    SlbClientSslCountersMetricsTopicDequeueErr int `json:"slb-client-ssl-counters-metrics-topic-dequeue-err"`
    SlbServerSslCountersMetricsTopic int `json:"slb-server-ssl-counters-metrics-topic"`
    SlbServerSslCountersMetricsTopicDequeueErr int `json:"slb-server-ssl-counters-metrics-topic-dequeue-err"`
    PcThrottleDrop int `json:"pc-throttle-drop"`
    MetricsDroppedPtMissing int `json:"metrics-dropped-pt-missing"`
    SsliPcAcosHarmonyTopic int `json:"ssli-pc-acos-harmony-topic"`
    SsliPcAcosHarmonyTopicDequeueErr int `json:"ssli-pc-acos-harmony-topic-dequeue-err"`
    SsliPeAcosHarmonyTopic int `json:"ssli-pe-acos-harmony-topic"`
    SsliPeAcosHarmonyTopicDequeueErr int `json:"ssli-pe-acos-harmony-topic-dequeue-err"`
    AnalyticsBusRestart int `json:"analytics-bus-restart"`
    WafLearnPrTopic int `json:"waf-learn-pr-topic"`
    WafLearnPrTopicDequeueErr int `json:"waf-learn-pr-topic-dequeue-err"`
    WafEventsTopic int `json:"waf-events-topic"`
    WafEventsTopicDequeueErr int `json:"waf-events-topic-dequeue-err"`
    VisibilityTopnHarmonyTopic int `json:"visibility-topn-harmony-topic"`
    VisibilityTopnHarmonyTopicDequeueErr int `json:"visibility-topn-harmony-topic-dequeue-err"`
    HcLogsSentToMaster int `json:"hc-logs-sent-to-master"`
    HcLogsReceivedFromBlade int `json:"hc-logs-received-from-blade"`
    HcOperSentToMaster int `json:"hc-oper-sent-to-master"`
    HcOperReceivedFromBlade int `json:"hc-oper-received-from-blade"`
    HcCountersSentToMaster int `json:"hc-counters-sent-to-master"`
    HcCountersReceivedFromBlade int `json:"hc-counters-received-from-blade"`
    HcCountersDroppedFromBlade int `json:"hc-counters-dropped-from-blade"`
    PeAcosHarmonyTopic int `json:"pe-acos-harmony-topic"`
    PeAcosHarmonyTopicEnqueueErr int `json:"pe-acos-harmony-topic-enqueue-err"`
    PeAcosHarmonyTopicDequeueErr int `json:"pe-acos-harmony-topic-dequeue-err"`
    VpnIpsecSaMetricsTopic int `json:"vpn-ipsec-sa-metrics-topic"`
    VpnIpsecSaMetricsTopicDequeueErr int `json:"vpn-ipsec-sa-metrics-topic-dequeue-err"`
    VpnIkeGatewayMetricsTopic int `json:"vpn-ike-gateway-metrics-topic"`
    VpnIkeGatewayMetricsTopicDequeueErr int `json:"vpn-ike-gateway-metrics-topic-dequeue-err"`
    VpnStatsMetricsTopic int `json:"vpn-stats-metrics-topic"`
    VpnStatsMetricsTopicDequeueErr int `json:"vpn-stats-metrics-topic-dequeue-err"`
    CgnPortUsageHstgrmAcosHarmonyTopic int `json:"cgn-port-usage-hstgrm-acos-harmony-topic"`
    CgnPortUsageHstgrmAcosHarmonyTopicDequeueErr int `json:"cgn-port-usage-hstgrm-acos-harmony-topic-dequeue-err"`
    AvroSystemEnvTopic int `json:"avro-system-env-topic"`
    AvroSystemEnvDequeueErr int `json:"avro-system-env-dequeue-err"`
    CertPinningListTopic int `json:"cert-pinning-list-topic"`
    CertPinningListTopicDequeueErr int `json:"cert-pinning-list-topic-dequeue-err"`
    NgwafHcEpTopic int `json:"ngwaf-hc-ep-topic"`
    NgwafHcEpTopicDequeueErr int `json:"ngwaf-hc-ep-topic-dequeue-err"`
    NgwafHcMetricsTopic int `json:"ngwaf-hc-metrics-topic"`
    NgwafHcMetricsTopicDequeueErr int `json:"ngwaf-hc-metrics-topic-dequeue-err"`
}

func (p *EventNotificationKafkaServerStats) GetId() string{
    return "1"
}

func (p *EventNotificationKafkaServerStats) getPath() string{
    return "event-notification/kafka/server/stats"
}

func (p *EventNotificationKafkaServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataEventNotificationKafkaServerStats,error) {
logger.Println("EventNotificationKafkaServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataEventNotificationKafkaServerStats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
