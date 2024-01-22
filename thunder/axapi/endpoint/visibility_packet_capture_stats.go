

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureStats struct {
    
    AutomatedCaptures VisibilityPacketCaptureStatsAutomatedCaptures `json:"automated-captures"`
    CaptureConfigList []VisibilityPacketCaptureStatsCaptureConfigList `json:"capture-config-list"`
    Stats VisibilityPacketCaptureStatsStats `json:"stats"`

}
type DataVisibilityPacketCaptureStats struct {
    DtVisibilityPacketCaptureStats VisibilityPacketCaptureStats `json:"packet-capture"`
}


type VisibilityPacketCaptureStatsAutomatedCaptures struct {
    Stats VisibilityPacketCaptureStatsAutomatedCapturesStats `json:"stats"`
}


type VisibilityPacketCaptureStatsAutomatedCapturesStats struct {
    TotalFailure int `json:"total-failure"`
}


type VisibilityPacketCaptureStatsCaptureConfigList struct {
    Name string `json:"name"`
    Stats VisibilityPacketCaptureStatsCaptureConfigListStats `json:"stats"`
}


type VisibilityPacketCaptureStatsCaptureConfigListStats struct {
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


type VisibilityPacketCaptureStatsStats struct {
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

func (p *VisibilityPacketCaptureStats) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureStats) getPath() string{
    return "visibility/packet-capture/stats"
}

func (p *VisibilityPacketCaptureStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityPacketCaptureStats,error) {
logger.Println("VisibilityPacketCaptureStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityPacketCaptureStats
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
