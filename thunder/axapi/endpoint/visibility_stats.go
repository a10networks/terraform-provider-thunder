

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityStats struct {
    
    MonEntityTelemetryData VisibilityStatsMonEntityTelemetryData `json:"mon-entity-telemetry-data"`
    PacketCapture VisibilityStatsPacketCapture `json:"packet-capture"`
    Reporting VisibilityStatsReporting `json:"reporting"`
    Stats VisibilityStatsStats `json:"stats"`
    Topn VisibilityStatsTopn `json:"topn"`

}
type DataVisibilityStats struct {
    DtVisibilityStats VisibilityStats `json:"visibility"`
}


type VisibilityStatsMonEntityTelemetryData struct {
    Stats VisibilityStatsMonEntityTelemetryDataStats `json:"stats"`
}


type VisibilityStatsMonEntityTelemetryDataStats struct {
    In_pkts int `json:"in_pkts"`
    Out_pkts int `json:"out_pkts"`
    In_bytes int `json:"in_bytes"`
    Out_bytes int `json:"out_bytes"`
    Errors int `json:"errors"`
    In_small_pkt int `json:"in_small_pkt"`
    In_frag int `json:"in_frag"`
    Out_small_pkt int `json:"out_small_pkt"`
    Out_frag int `json:"out_frag"`
    NewConn int `json:"new-conn"`
    Avg_data_cpu_util int `json:"avg_data_cpu_util"`
    Outside_intf_util int `json:"outside_intf_util"`
    ConcurrentConn int `json:"concurrent-conn"`
    In_bytes_per_out_bytes int `json:"in_bytes_per_out_bytes"`
    Drop_pkts_per_pkts int `json:"drop_pkts_per_pkts"`
    Tcp_in_syn int `json:"tcp_in_syn"`
    Tcp_out_syn int `json:"tcp_out_syn"`
    Tcp_in_fin int `json:"tcp_in_fin"`
    Tcp_out_fin int `json:"tcp_out_fin"`
    Tcp_in_payload int `json:"tcp_in_payload"`
    Tcp_out_payload int `json:"tcp_out_payload"`
    Tcp_in_rexmit int `json:"tcp_in_rexmit"`
    Tcp_out_rexmit int `json:"tcp_out_rexmit"`
    Tcp_in_rst int `json:"tcp_in_rst"`
    Tcp_out_rst int `json:"tcp_out_rst"`
    Tcp_in_empty_ack int `json:"tcp_in_empty_ack"`
    Tcp_out_empty_ack int `json:"tcp_out_empty_ack"`
    Tcp_in_zero_wnd int `json:"tcp_in_zero_wnd"`
    Tcp_out_zero_wnd int `json:"tcp_out_zero_wnd"`
    Tcp_conn_miss int `json:"tcp_conn_miss"`
    Tcp_fwd_syn_per_fin int `json:"tcp_fwd_syn_per_fin"`
}


type VisibilityStatsPacketCapture struct {
    Stats VisibilityStatsPacketCaptureStats `json:"stats"`
    CaptureConfigList []VisibilityStatsPacketCaptureCaptureConfigList `json:"capture-config-list"`
    AutomatedCaptures VisibilityStatsPacketCaptureAutomatedCaptures `json:"automated-captures"`
}


type VisibilityStatsPacketCaptureStats struct {
    ConcurrentCaptureCreatedByCtrIncrement int `json:"Concurrent-capture-created-by-ctr-increment"`
    ConcurrentCaptureCreatedByCtrAnomaly int `json:"Concurrent-capture-created-by-ctr-anomaly"`
    ConcurrentCaptureCreatedByOtherFeature int `json:"Concurrent-capture-created-by-other-feature"`
    ConcurrentCaptureCreateFailedByCtrIncrement int `json:"Concurrent-capture-create-failed-by-ctr-increment"`
    ConcurrentCaptureCreateFailedByCtrAnomaly int `json:"Concurrent-capture-create-failed-by-ctr-anomaly"`
    ConcurrentCaptureCreateFailedByOtherFeature int `json:"Concurrent-capture-create-failed-by-other-feature"`
    ConcurrentCaptureCreateFailedOom int `json:"Concurrent-capture-create-failed-oom"`
    ConcurrentCaptureLimitReached int `json:"Concurrent-capture-limit-reached"`
    ConcurrentCaptureByCtrIncrementFreed int `json:"Concurrent-capture-by-ctr-increment-freed"`
    ConcurrentCaptureByCtrAnomalyFreed int `json:"Concurrent-capture-by-ctr-anomaly-freed"`
    ConcurrentCaptureByCtrOtherFeatureFreed int `json:"Concurrent-capture-by-ctr-other-feature-freed"`
    GlobalCaptureFinished int `json:"Global-capture-finished"`
    ConcurrentCaptureFinished int `json:"Concurrent-capture-finished"`
    PktcaptureWithNoConnSuccess int `json:"pktcapture-with-no-conn-success"`
    PktcaptureWithNoConnFailure int `json:"pktcapture-with-no-conn-failure"`
    PktcaptureWithConnButNotTaggedSuccess int `json:"pktcapture-with-conn-but-not-tagged-success"`
    PktcaptureWithConnButNotTaggedFailure int `json:"pktcapture-with-conn-but-not-tagged-failure"`
    PktcaptureWithConnSuccessGlobal int `json:"pktcapture-with-conn-success-global"`
    PktcaptureWithConnSuccess int `json:"pktcapture-with-conn-success"`
    PktcaptureWithConnFailureGlobal int `json:"pktcapture-with-conn-failure-global"`
    PktcaptureWithConnFailure int `json:"pktcapture-with-conn-failure"`
    PktcaptureFailureWaitForBlock int `json:"pktcapture-failure-wait-for-block"`
    PktcaptureFailureFileSizeRchd int `json:"pktcapture-failure-file-size-rchd"`
    NumConnsTaggedGlobalIncrement int `json:"num-conns-tagged-global-increment"`
    NumConnsTaggedGlobalAnomaly int `json:"num-conns-tagged-global-anomaly"`
    NumConnsTaggedGlobalOtherFeature int `json:"num-conns-tagged-global-other-feature"`
    NumConnsTaggedGlobalIncrementFail int `json:"num-conns-tagged-global-increment-fail"`
    NumConnsTaggedGlobalAnomalyFail int `json:"num-conns-tagged-global-anomaly-fail"`
    NumConnsTaggedGlobalOtherFeatureFail int `json:"num-conns-tagged-global-other-feature-fail"`
    NumConnsTaggedGlobalIncrementMaxed int `json:"num-conns-tagged-global-increment-maxed"`
    NumConnsTaggedGlobalAnomalyMaxed int `json:"num-conns-tagged-global-anomaly-maxed"`
    NumConnsTaggedGlobalOtherFeatureMaxed int `json:"num-conns-tagged-global-other-feature-maxed"`
    NumConnsTaggedIncrement int `json:"num-conns-tagged-increment"`
    NumConnsTaggedAnomaly int `json:"num-conns-tagged-anomaly"`
    NumConnsTaggedOtherFeature int `json:"num-conns-tagged-other-feature"`
    NumConnsTaggedIncrementFail int `json:"num-conns-tagged-increment-fail"`
    NumConnsTaggedAnomalyFail int `json:"num-conns-tagged-anomaly-fail"`
    NumConnsTaggedOtherFeatureFail int `json:"num-conns-tagged-other-feature-fail"`
    NumConnsTaggedIncrementMaxed int `json:"num-conns-tagged-increment-maxed"`
    NumConnsTaggedAnomalyMaxed int `json:"num-conns-tagged-anomaly-maxed"`
    NumConnsTaggedOtherFeatureMaxed int `json:"num-conns-tagged-other-feature-maxed"`
    NumConnsUntagged int `json:"num-conns-untagged"`
    PktcaptureTriggeredByIncrement int `json:"pktcapture-triggered-by-increment"`
    PktcaptureTriggeredByAnomaly int `json:"pktcapture-triggered-by-anomaly"`
    PktcaptureTriggeredByOtherFeature int `json:"pktcapture-triggered-by-other-feature"`
    NumOfAnomaliesDetected int `json:"num-of-anomalies-detected"`
    NumOfAnomaliesCleared int `json:"num-of-anomalies-cleared"`
    NumPcapsCreated int `json:"num-pcaps-created"`
    NumTmpPcapsCreated int `json:"num-tmp-pcaps-created"`
    NumPcapsCreateFailed int `json:"num-pcaps-create-failed"`
    PktcapOom int `json:"pktcap-oom"`
    FailedDiskFull int `json:"failed-disk-full"`
    ConnExtFailed int `json:"conn-ext-failed"`
    SkipAsConnAlreadyRecapture int `json:"skip-as-conn-already-recapture"`
    SkipCaptureAsConnCreatedBeforeSmp int `json:"skip-capture-as-conn-created-before-smp"`
    FailedAsReturnCompletedSet int `json:"failed-as-return-completed-set"`
    NonPktPath int `json:"non-pkt-path"`
    PktAlreadyCaptured int `json:"pkt-already-captured"`
    WrongCtrIncremented int `json:"wrong-ctr-incremented"`
    AutoPcapFileMerged int `json:"auto-pcap-file-merged"`
    AutoPcapFileMergedFailed int `json:"auto-pcap-file-merged-failed"`
    NumGlobalTmplCreated int `json:"num-global-tmpl-created"`
    NumObjectTmplCreated int `json:"num-object-tmpl-created"`
    NumGlobalTmplDeleted int `json:"num-global-tmpl-deleted"`
    NumObjectTmplDeleted int `json:"num-object-tmpl-deleted"`
    NumCaptureConfigCreated int `json:"num-capture-config-created"`
    NumDynamicCaptureConfigCreated int `json:"num-dynamic-capture-config-created"`
    NumCaptureConfigDeleted int `json:"num-capture-config-deleted"`
    NumDynamicCaptureConfigDeleted int `json:"num-dynamic-capture-config-deleted"`
    NumCaptureConfigDeleteQ int `json:"num-capture-config-delete-q"`
    NumDynamicCaptureConfigDeleteQ int `json:"num-dynamic-capture-config-delete-q"`
    NumCaptureConfigLinked int `json:"num-capture-config-linked"`
    NumDynamicCaptureConfigLinked int `json:"num-dynamic-capture-config-linked"`
    NumCaptureConfigUnlinked int `json:"num-capture-config-unlinked"`
    NumDynamicCaptureConfigUnlinked int `json:"num-dynamic-capture-config-unlinked"`
    NumGlobalCountersRegistered int `json:"num-global-counters-registered"`
    NumGlobalCountersDeregistered int `json:"num-global-counters-deregistered"`
    NumPerObjectCountersRegistered int `json:"num-per-object-counters-registered"`
    NumPerObjectCountersDeregistered int `json:"num-per-object-counters-deregistered"`
}


type VisibilityStatsPacketCaptureCaptureConfigList struct {
    Name string `json:"name"`
    Stats VisibilityStatsPacketCaptureCaptureConfigListStats `json:"stats"`
}


type VisibilityStatsPacketCaptureCaptureConfigListStats struct {
    ConcurrentCaptureCreatedByCtrIncrement int `json:"Concurrent-capture-created-by-ctr-increment"`
    ConcurrentCaptureCreatedByCtrAnomaly int `json:"Concurrent-capture-created-by-ctr-anomaly"`
    ConcurrentCaptureCreateFailedByCtrIncrement int `json:"Concurrent-capture-create-failed-by-ctr-increment"`
    ConcurrentCaptureCreateFailedByCtrAnomaly int `json:"Concurrent-capture-create-failed-by-ctr-anomaly"`
    ConcurrentCaptureCreateFailedByOtherFeature int `json:"Concurrent-capture-create-failed-by-other-feature"`
    ConcurrentCaptureCreateFailedOom int `json:"Concurrent-capture-create-failed-oom"`
    ConcurrentCaptureLimitReached int `json:"Concurrent-capture-limit-reached"`
    ConcurrentCaptureByCtrIncrementFreed int `json:"Concurrent-capture-by-ctr-increment-freed"`
    ConcurrentCaptureByCtrAnomalyFreed int `json:"Concurrent-capture-by-ctr-anomaly-freed"`
    ConcurrentCaptureByCtrOtherFeatureFreed int `json:"Concurrent-capture-by-ctr-other-feature-freed"`
    GlobalCaptureFinished int `json:"Global-capture-finished"`
    ConcurrentCaptureFinished int `json:"Concurrent-capture-finished"`
    PktcaptureWithNoConnSuccess int `json:"pktcapture-with-no-conn-success"`
    PktcaptureWithNoConnFailure int `json:"pktcapture-with-no-conn-failure"`
    PktcaptureWithConnButNotTaggedSuccess int `json:"pktcapture-with-conn-but-not-tagged-success"`
    PktcaptureWithConnButNotTaggedFailure int `json:"pktcapture-with-conn-but-not-tagged-failure"`
    PktcaptureWithConnSuccessGlobal int `json:"pktcapture-with-conn-success-global"`
    PktcaptureWithConnSuccess int `json:"pktcapture-with-conn-success"`
    PktcaptureWithConnFailureGlobal int `json:"pktcapture-with-conn-failure-global"`
    PktcaptureWithConnFailure int `json:"pktcapture-with-conn-failure"`
    PktcaptureFailureWaitForBlock int `json:"pktcapture-failure-wait-for-block"`
    PktcaptureFailureFileSizeRchd int `json:"pktcapture-failure-file-size-rchd"`
    NumConnsTaggedGlobalIncrement int `json:"num-conns-tagged-global-increment"`
    NumConnsTaggedGlobalAnomaly int `json:"num-conns-tagged-global-anomaly"`
    NumConnsTaggedGlobalOtherFeature int `json:"num-conns-tagged-global-other-feature"`
    NumConnsTaggedGlobalIncrementFail int `json:"num-conns-tagged-global-increment-fail"`
    NumConnsTaggedGlobalAnomalyFail int `json:"num-conns-tagged-global-anomaly-fail"`
    NumConnsTaggedGlobalOtherFeatureFail int `json:"num-conns-tagged-global-other-feature-fail"`
    NumConnsTaggedGlobalIncrementMaxed int `json:"num-conns-tagged-global-increment-maxed"`
    NumConnsTaggedGlobalAnomalyMaxed int `json:"num-conns-tagged-global-anomaly-maxed"`
    NumConnsTaggedGlobalOtherFeatureMaxed int `json:"num-conns-tagged-global-other-feature-maxed"`
    NumConnsTaggedIncrement int `json:"num-conns-tagged-increment"`
    NumConnsTaggedAnomaly int `json:"num-conns-tagged-anomaly"`
    NumConnsTaggedOtherFeature int `json:"num-conns-tagged-other-feature"`
    NumConnsTaggedIncrementFail int `json:"num-conns-tagged-increment-fail"`
    NumConnsTaggedAnomalyFail int `json:"num-conns-tagged-anomaly-fail"`
    NumConnsTaggedOtherFeatureFail int `json:"num-conns-tagged-other-feature-fail"`
    NumConnsTaggedIncrementMaxed int `json:"num-conns-tagged-increment-maxed"`
    NumConnsTaggedAnomalyMaxed int `json:"num-conns-tagged-anomaly-maxed"`
    NumConnsTaggedOtherFeatureMaxed int `json:"num-conns-tagged-other-feature-maxed"`
    NumConnsUntagged int `json:"num-conns-untagged"`
    PktcaptureTriggeredByIncrement int `json:"pktcapture-triggered-by-increment"`
    PktcaptureTriggeredByAnomaly int `json:"pktcapture-triggered-by-anomaly"`
    PktcaptureTriggeredByOtherFeature int `json:"pktcapture-triggered-by-other-feature"`
    NumOfAnomaliesDetected int `json:"num-of-anomalies-detected"`
    NumOfAnomaliesCleared int `json:"num-of-anomalies-cleared"`
    NumPcapsCreated int `json:"num-pcaps-created"`
    NumTmpPcapsCreated int `json:"num-tmp-pcaps-created"`
    NumPcapsCreateFailed int `json:"num-pcaps-create-failed"`
    PktcapOom int `json:"pktcap-oom"`
    FailedDiskFull int `json:"failed-disk-full"`
    ConnExtFailed int `json:"conn-ext-failed"`
    SkipAsConnAlreadyRecapture int `json:"skip-as-conn-already-recapture"`
    SkipCaptureAsConnCreatedBeforeSmp int `json:"skip-capture-as-conn-created-before-smp"`
    FailedAsReturnCompletedSet int `json:"failed-as-return-completed-set"`
    NonPktPath int `json:"non-pkt-path"`
    PktAlreadyCaptured int `json:"pkt-already-captured"`
    WrongCtrIncremented int `json:"wrong-ctr-incremented"`
    AutoPcapFileMerged int `json:"auto-pcap-file-merged"`
    AutoPcapFileMergedFailed int `json:"auto-pcap-file-merged-failed"`
    NumDynamicCaptureConfigCreated int `json:"num-dynamic-capture-config-created"`
    NumDynamicCaptureConfigDeleteQ int `json:"num-dynamic-capture-config-delete-q"`
    NumDynamicCaptureConfigDeleted int `json:"num-dynamic-capture-config-deleted"`
    NumGlobalCountersRegistered int `json:"num-global-counters-registered"`
    NumGlobalCountersDeregistered int `json:"num-global-counters-deregistered"`
    NumPerObjectCountersRegistered int `json:"num-per-object-counters-registered"`
    NumPerObjectCountersDeregistered int `json:"num-per-object-counters-deregistered"`
}


type VisibilityStatsPacketCaptureAutomatedCaptures struct {
    Stats VisibilityStatsPacketCaptureAutomatedCapturesStats `json:"stats"`
}


type VisibilityStatsPacketCaptureAutomatedCapturesStats struct {
    TotalFailure int `json:"total-failure"`
}


type VisibilityStatsReporting struct {
    Stats VisibilityStatsReportingStats `json:"stats"`
}


type VisibilityStatsReportingStats struct {
    LogTransmitFailure int `json:"log-transmit-failure"`
    BufferAllocFailure int `json:"buffer-alloc-failure"`
    NotifJobsInQueue int `json:"notif-jobs-in-queue"`
    EnqueueFail int `json:"enqueue-fail"`
    EnqueuePass int `json:"enqueue-pass"`
    Dequeued int `json:"dequeued"`
}


type VisibilityStatsStats struct {
    MonEntityLimitExceed int `json:"mon-entity-limit-exceed"`
    HaEntityCreateSent int `json:"ha-entity-create-sent"`
    HaEntityDeleteSent int `json:"ha-entity-delete-sent"`
    HaEntityAnomalyOnSent int `json:"ha-entity-anomaly-on-sent"`
    HaEntityAnomalyOffSent int `json:"ha-entity-anomaly-off-sent"`
    HaEntityPeriodicSyncSent int `json:"ha-entity-periodic-sync-sent"`
    OutOfMemoryAllocFailures int `json:"out-of-memory-alloc-failures"`
    LwMonEntityCreated int `json:"lw-mon-entity-created"`
    LwMonEntityDeleted int `json:"lw-mon-entity-deleted"`
    LwMonEntityLimitExceed int `json:"lw-mon-entity-limit-exceed"`
    LwOutOfMemoryAllocFailures int `json:"lw-out-of-memory-alloc-failures"`
    MonEntityRrdFileTimestampErr int `json:"mon-entity-rrd-file-timestamp-err"`
    MonEntityRrdUpdateErr int `json:"mon-entity-rrd-update-err"`
    MonEntityRrdLastUpdateFetchFailedErr int `json:"mon-entity-rrd-last-update-fetch-failed-err"`
    MonEntityRrdTuneErr int `json:"mon-entity-rrd-tune-err"`
    MonEntityRrdOutOfMemoryErr int `json:"mon-entity-rrd-out-of-memory-err"`
    MonEntityRrdFileCreateErr int `json:"mon-entity-rrd-file-create-err"`
}


type VisibilityStatsTopn struct {
    Stats VisibilityStatsTopnStats `json:"stats"`
}


type VisibilityStatsTopnStats struct {
    HeapAllocSuccess int `json:"heap-alloc-success"`
    HeapAllocFailed int `json:"heap-alloc-failed"`
    HeapAllocOom int `json:"heap-alloc-oom"`
    ObjRegSuccess int `json:"obj-reg-success"`
    ObjRegFailed int `json:"obj-reg-failed"`
    ObjRegOom int `json:"obj-reg-oom"`
    HeapDeleted int `json:"heap-deleted"`
    ObjDeleted int `json:"obj-deleted"`
}

func (p *VisibilityStats) GetId() string{
    return "1"
}

func (p *VisibilityStats) getPath() string{
    return "visibility/stats"
}

func (p *VisibilityStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityStats,error) {
logger.Println("VisibilityStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityStats
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
