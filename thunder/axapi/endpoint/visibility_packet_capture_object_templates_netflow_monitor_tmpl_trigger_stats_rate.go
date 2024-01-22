

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate struct {
	Inst struct {

    CustomDenyResetEventFw4RecordsSen int `json:"custom-deny-reset-event-fw4-records-sen"`
    CustomDenyResetEventFw6RecordsSen int `json:"custom-deny-reset-event-fw6-records-sen"`
    CustomFwIddosEntryCreatedRecordsS int `json:"custom-fw-iddos-entry-created-records-s"`
    CustomFwIddosEntryDeletedRecordsS int `json:"custom-fw-iddos-entry-deleted-records-s"`
    CustomFwSesnLimitExceededRecordsS int `json:"custom-fw-sesn-limit-exceeded-records-s"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent-"`
    CustomGtpDenyEventRecordsSentFail int `json:"custom-gtp-deny-event-records-sent-fail"`
    CustomGtpInfoEventRecordsSentFail int `json:"custom-gtp-info-event-records-sent-fail"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent-"`
    CustomNatIddosL3EntryCreatedRecor int `json:"custom-nat-iddos-l3-entry-created-recor"`
    CustomNatIddosL3EntryDeletedRecor int `json:"custom-nat-iddos-l3-entry-deleted-recor"`
    CustomNatIddosL4EntryCreatedRecor int `json:"custom-nat-iddos-l4-entry-created-recor"`
    CustomNatIddosL4EntryDeletedRecor int `json:"custom-nat-iddos-l4-entry-deleted-recor"`
    CustomPortBatchingDsliteCreationRe int `json:"custom-port-batching-dslite-creation-re"`
    CustomPortBatchingDsliteDeletionRe int `json:"custom-port-batching-dslite-deletion-re"`
    CustomPortBatchingNat44CreationRec int `json:"custom-port-batching-nat44-creation-rec"`
    CustomPortBatchingNat44DeletionRec int `json:"custom-port-batching-nat44-deletion-rec"`
    CustomPortBatchingNat64CreationRec int `json:"custom-port-batching-nat64-creation-rec"`
    CustomPortBatchingNat64DeletionRec int `json:"custom-port-batching-nat64-deletion-rec"`
    CustomPortBatchingV2DsliteCreation int `json:"custom-port-batching-v2-dslite-creation"`
    CustomPortBatchingV2DsliteDeletion int `json:"custom-port-batching-v2-dslite-deletion"`
    CustomPortBatchingV2Nat44Creation int `json:"custom-port-batching-v2-nat44-creation-"`
    CustomPortBatchingV2Nat44Deletion int `json:"custom-port-batching-v2-nat44-deletion-"`
    CustomPortBatchingV2Nat64Creation int `json:"custom-port-batching-v2-nat64-creation-"`
    CustomPortBatchingV2Nat64Deletion int `json:"custom-port-batching-v2-nat64-deletion-"`
    CustomPortMappingDsliteCreationRec int `json:"custom-port-mapping-dslite-creation-rec"`
    CustomPortMappingDsliteDeletionRec int `json:"custom-port-mapping-dslite-deletion-rec"`
    CustomPortMappingNat44CreationReco int `json:"custom-port-mapping-nat44-creation-reco"`
    CustomPortMappingNat44DeletionReco int `json:"custom-port-mapping-nat44-deletion-reco"`
    CustomPortMappingNat64CreationReco int `json:"custom-port-mapping-nat64-creation-reco"`
    CustomPortMappingNat64DeletionReco int `json:"custom-port-mapping-nat64-deletion-reco"`
    CustomSessionEventDsliteCreationRe int `json:"custom-session-event-dslite-creation-re"`
    CustomSessionEventDsliteDeletionRe int `json:"custom-session-event-dslite-deletion-re"`
    CustomSessionEventFw4CreationRecor int `json:"custom-session-event-fw4-creation-recor"`
    CustomSessionEventFw4DeletionRecor int `json:"custom-session-event-fw4-deletion-recor"`
    CustomSessionEventFw6CreationRecor int `json:"custom-session-event-fw6-creation-recor"`
    CustomSessionEventFw6DeletionRecor int `json:"custom-session-event-fw6-deletion-recor"`
    CustomSessionEventNat44CreationRec int `json:"custom-session-event-nat44-creation-rec"`
    CustomSessionEventNat44DeletionRec int `json:"custom-session-event-nat44-deletion-rec"`
    CustomSessionEventNat64CreationRec int `json:"custom-session-event-nat64-creation-rec"`
    CustomSessionEventNat64DeletionRec int `json:"custom-session-event-nat64-deletion-rec"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    Duration int `json:"duration" dval:"60"`
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    PortBatchingDsliteRecordsSentFailu int `json:"port-batching-dslite-records-sent-failu"`
    PortBatchingNat44RecordsSentFailur int `json:"port-batching-nat44-records-sent-failur"`
    PortBatchingNat64RecordsSentFailur int `json:"port-batching-nat64-records-sent-failur"`
    PortBatchingV2DsliteRecordsSentFa int `json:"port-batching-v2-dslite-records-sent-fa"`
    PortBatchingV2Nat44RecordsSentFai int `json:"port-batching-v2-nat44-records-sent-fai"`
    PortBatchingV2Nat64RecordsSentFai int `json:"port-batching-v2-nat64-records-sent-fai"`
    PortMappingDsliteRecordsSentFailur int `json:"port-mapping-dslite-records-sent-failur"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    SessionEventDsliteRecordsSentFailu int `json:"session-event-dslite-records-sent-failu"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    SessionEventNat44RecordsSentFailur int `json:"session-event-nat44-records-sent-failur"`
    SessionEventNat64RecordsSentFailur int `json:"session-event-nat64-records-sent-failur"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/object-templates/netflow-monitor-tmpl/" +p.Inst.Name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesNetflowMonitorTmplTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
