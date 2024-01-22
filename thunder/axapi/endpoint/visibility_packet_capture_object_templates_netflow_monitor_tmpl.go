

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc2699 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate2700 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsSeverity2701 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"netflow-monitor-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsInc2699 struct {
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate2700 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsSeverity2701 struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/netflow-monitor-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
