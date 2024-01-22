

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorStats struct {
    
    Name string `json:"name"`
    Stats NetflowMonitorStatsStats `json:"stats"`

}
type DataNetflowMonitorStats struct {
    DtNetflowMonitorStats NetflowMonitorStats `json:"monitor"`
}


type NetflowMonitorStatsStats struct {
    PacketsSent int `json:"packets-sent"`
    BytesSent int `json:"bytes-sent"`
    Nat44RecordsSent int `json:"nat44-records-sent"`
    Nat44RecordsSentFailure int `json:"nat44-records-sent-failure"`
    Nat64RecordsSent int `json:"nat64-records-sent"`
    Nat64RecordsSentFailure int `json:"nat64-records-sent-failure"`
    DsliteRecordsSent int `json:"dslite-records-sent"`
    DsliteRecordsSentFailure int `json:"dslite-records-sent-failure"`
    SessionEventNat44RecordsSent int `json:"session-event-nat44-records-sent"`
    SessionEventNat44RecordsSentFailure int `json:"session-event-nat44-records-sent-failure"`
    SessionEventNat64RecordsSent int `json:"session-event-nat64-records-sent"`
    SessionEventNat64RecordsSentFailure int `json:"session-event-nat64-records-sent-failure"`
    SessionEventDsliteRecordsSent int `json:"session-event-dslite-records-sent"`
    SessionEventDsliteRecordsSentFailure int `json:"session-event-dslite-records-sent-failure"`
    SessionEventFw4RecordsSent int `json:"session-event-fw4-records-sent"`
    SessionEventFw4RecordsSentFailure int `json:"session-event-fw4-records-sent-failure"`
    SessionEventFw6RecordsSent int `json:"session-event-fw6-records-sent"`
    SessionEventFw6RecordsSentFailure int `json:"session-event-fw6-records-sent-failure"`
    PortMappingNat44RecordsSent int `json:"port-mapping-nat44-records-sent"`
    PortMappingNat44RecordsSentFailure int `json:"port-mapping-nat44-records-sent-failure"`
    PortMappingNat64RecordsSent int `json:"port-mapping-nat64-records-sent"`
    PortMappingNat64RecordsSentFailure int `json:"port-mapping-nat64-records-sent-failure"`
    PortMappingDsliteRecordsSent int `json:"port-mapping-dslite-records-sent"`
    PortMappingDsliteRecordsSentFailure int `json:"port-mapping-dslite-records-sent-failure"`
    NetflowV5RecordsSent int `json:"netflow-v5-records-sent"`
    NetflowV5RecordsSentFailure int `json:"netflow-v5-records-sent-failure"`
    NetflowV5ExtRecordsSent int `json:"netflow-v5-ext-records-sent"`
    NetflowV5ExtRecordsSentFailure int `json:"netflow-v5-ext-records-sent-failure"`
    PortBatchingNat44RecordsSent int `json:"port-batching-nat44-records-sent"`
    PortBatchingNat44RecordsSentFailure int `json:"port-batching-nat44-records-sent-failure"`
    PortBatchingNat64RecordsSent int `json:"port-batching-nat64-records-sent"`
    PortBatchingNat64RecordsSentFailure int `json:"port-batching-nat64-records-sent-failure"`
    PortBatchingDsliteRecordsSent int `json:"port-batching-dslite-records-sent"`
    PortBatchingDsliteRecordsSentFailure int `json:"port-batching-dslite-records-sent-failure"`
    PortBatchingV2Nat44RecordsSent int `json:"port-batching-v2-nat44-records-sent"`
    PortBatchingV2Nat44RecordsSentFailure int `json:"port-batching-v2-nat44-records-sent-failure"`
    PortBatchingV2Nat64RecordsSent int `json:"port-batching-v2-nat64-records-sent"`
    PortBatchingV2Nat64RecordsSentFailure int `json:"port-batching-v2-nat64-records-sent-failure"`
    PortBatchingV2DsliteRecordsSent int `json:"port-batching-v2-dslite-records-sent"`
    PortBatchingV2DsliteRecordsSentFailure int `json:"port-batching-v2-dslite-records-sent-failure"`
    CustomSessionEventNat44CreationRecordsSent int `json:"custom-session-event-nat44-creation-records-sent"`
    CustomSessionEventNat44CreationRecordsSentFailure int `json:"custom-session-event-nat44-creation-records-sent-failure"`
    CustomSessionEventNat64CreationRecordsSent int `json:"custom-session-event-nat64-creation-records-sent"`
    CustomSessionEventNat64CreationRecordsSentFailure int `json:"custom-session-event-nat64-creation-records-sent-failure"`
    CustomSessionEventDsliteCreationRecordsSent int `json:"custom-session-event-dslite-creation-records-sent"`
    CustomSessionEventDsliteCreationRecordsSentFailure int `json:"custom-session-event-dslite-creation-records-sent-failure"`
    CustomSessionEventNat44DeletionRecordsSent int `json:"custom-session-event-nat44-deletion-records-sent"`
    CustomSessionEventNat44DeletionRecordsSentFailure int `json:"custom-session-event-nat44-deletion-records-sent-failure"`
    CustomSessionEventNat64DeletionRecordsSent int `json:"custom-session-event-nat64-deletion-records-sent"`
    CustomSessionEventNat64DeletionRecordsSentFailure int `json:"custom-session-event-nat64-deletion-records-sent-failure"`
    CustomSessionEventDsliteDeletionRecordsSent int `json:"custom-session-event-dslite-deletion-records-sent"`
    CustomSessionEventDsliteDeletionRecordsSentFailure int `json:"custom-session-event-dslite-deletion-records-sent-failure"`
    CustomSessionEventFw4CreationRecordsSent int `json:"custom-session-event-fw4-creation-records-sent"`
    CustomSessionEventFw4CreationRecordsSentFailure int `json:"custom-session-event-fw4-creation-records-sent-failure"`
    CustomSessionEventFw6CreationRecordsSent int `json:"custom-session-event-fw6-creation-records-sent"`
    CustomSessionEventFw6CreationRecordsSentFailure int `json:"custom-session-event-fw6-creation-records-sent-failure"`
    CustomSessionEventFw4DeletionRecordsSent int `json:"custom-session-event-fw4-deletion-records-sent"`
    CustomSessionEventFw4DeletionRecordsSentFailure int `json:"custom-session-event-fw4-deletion-records-sent-failure"`
    CustomSessionEventFw6DeletionRecordsSent int `json:"custom-session-event-fw6-deletion-records-sent"`
    CustomSessionEventFw6DeletionRecordsSentFailure int `json:"custom-session-event-fw6-deletion-records-sent-failure"`
    CustomDenyResetEventFw4RecordsSent int `json:"custom-deny-reset-event-fw4-records-sent"`
    CustomDenyResetEventFw4RecordsSentFailure int `json:"custom-deny-reset-event-fw4-records-sent-failure"`
    CustomDenyResetEventFw6RecordsSent int `json:"custom-deny-reset-event-fw6-records-sent"`
    CustomDenyResetEventFw6RecordsSentFailure int `json:"custom-deny-reset-event-fw6-records-sent-failure"`
    CustomPortMappingNat44CreationRecordsSent int `json:"custom-port-mapping-nat44-creation-records-sent"`
    CustomPortMappingNat44CreationRecordsSentFailure int `json:"custom-port-mapping-nat44-creation-records-sent-failure"`
    CustomPortMappingNat64CreationRecordsSent int `json:"custom-port-mapping-nat64-creation-records-sent"`
    CustomPortMappingNat64CreationRecordsSentFailure int `json:"custom-port-mapping-nat64-creation-records-sent-failure"`
    CustomPortMappingDsliteCreationRecordsSent int `json:"custom-port-mapping-dslite-creation-records-sent"`
    CustomPortMappingDsliteCreationRecordsSentFailure int `json:"custom-port-mapping-dslite-creation-records-sent-failure"`
    CustomPortMappingNat44DeletionRecordsSent int `json:"custom-port-mapping-nat44-deletion-records-sent"`
    CustomPortMappingNat44DeletionRecordsSentFailure int `json:"custom-port-mapping-nat44-deletion-records-sent-failure"`
    CustomPortMappingNat64DeletionRecordsSent int `json:"custom-port-mapping-nat64-deletion-records-sent"`
    CustomPortMappingNat64DeletionRecordsSentFailure int `json:"custom-port-mapping-nat64-deletion-records-sent-failure"`
    CustomPortMappingDsliteDeletionRecordsSent int `json:"custom-port-mapping-dslite-deletion-records-sent"`
    CustomPortMappingDsliteDeletionRecordsSentFailure int `json:"custom-port-mapping-dslite-deletion-records-sent-failure"`
    CustomPortBatchingNat44CreationRecordsSent int `json:"custom-port-batching-nat44-creation-records-sent"`
    CustomPortBatchingNat44CreationRecordsSentFailure int `json:"custom-port-batching-nat44-creation-records-sent-failure"`
    CustomPortBatchingNat64CreationRecordsSent int `json:"custom-port-batching-nat64-creation-records-sent"`
    CustomPortBatchingNat64CreationRecordsSentFailure int `json:"custom-port-batching-nat64-creation-records-sent-failure"`
    CustomPortBatchingDsliteCreationRecordsSent int `json:"custom-port-batching-dslite-creation-records-sent"`
    CustomPortBatchingDsliteCreationRecordsSentFailure int `json:"custom-port-batching-dslite-creation-records-sent-failure"`
    CustomPortBatchingNat44DeletionRecordsSent int `json:"custom-port-batching-nat44-deletion-records-sent"`
    CustomPortBatchingNat44DeletionRecordsSentFailure int `json:"custom-port-batching-nat44-deletion-records-sent-failure"`
    CustomPortBatchingNat64DeletionRecordsSent int `json:"custom-port-batching-nat64-deletion-records-sent"`
    CustomPortBatchingNat64DeletionRecordsSentFailure int `json:"custom-port-batching-nat64-deletion-records-sent-failure"`
    CustomPortBatchingDsliteDeletionRecordsSent int `json:"custom-port-batching-dslite-deletion-records-sent"`
    CustomPortBatchingDsliteDeletionRecordsSentFailure int `json:"custom-port-batching-dslite-deletion-records-sent-failure"`
    CustomPortBatchingV2Nat44CreationRecordsSent int `json:"custom-port-batching-v2-nat44-creation-records-sent"`
    CustomPortBatchingV2Nat44CreationRecordsSentFailure int `json:"custom-port-batching-v2-nat44-creation-records-sent-failure"`
    CustomPortBatchingV2Nat64CreationRecordsSent int `json:"custom-port-batching-v2-nat64-creation-records-sent"`
    CustomPortBatchingV2Nat64CreationRecordsSentFailure int `json:"custom-port-batching-v2-nat64-creation-records-sent-failure"`
    CustomPortBatchingV2DsliteCreationRecordsSent int `json:"custom-port-batching-v2-dslite-creation-records-sent"`
    CustomPortBatchingV2DsliteCreationRecordsSentFailure int `json:"custom-port-batching-v2-dslite-creation-records-sent-failure"`
    CustomPortBatchingV2Nat44DeletionRecordsSent int `json:"custom-port-batching-v2-nat44-deletion-records-sent"`
    CustomPortBatchingV2Nat44DeletionRecordsSentFailure int `json:"custom-port-batching-v2-nat44-deletion-records-sent-failure"`
    CustomPortBatchingV2Nat64DeletionRecordsSent int `json:"custom-port-batching-v2-nat64-deletion-records-sent"`
    CustomPortBatchingV2Nat64DeletionRecordsSentFailure int `json:"custom-port-batching-v2-nat64-deletion-records-sent-failure"`
    CustomPortBatchingV2DsliteDeletionRecordsSent int `json:"custom-port-batching-v2-dslite-deletion-records-sent"`
    CustomPortBatchingV2DsliteDeletionRecordsSentFailure int `json:"custom-port-batching-v2-dslite-deletion-records-sent-failure"`
    ReducedLogsByDestination int `json:"reduced-logs-by-destination"`
    CustomGtpCTunnelEventRecordsSent int `json:"custom-gtp-c-tunnel-event-records-sent"`
    CustomGtpCTunnelEventRecordsSentFailure int `json:"custom-gtp-c-tunnel-event-records-sent-failure"`
    CustomGtpUTunnelEventRecordsSent int `json:"custom-gtp-u-tunnel-event-records-sent"`
    CustomGtpUTunnelEventRecordsSentFailure int `json:"custom-gtp-u-tunnel-event-records-sent-failure"`
    CustomGtpDenyEventRecordsSent int `json:"custom-gtp-deny-event-records-sent"`
    CustomGtpDenyEventRecordsSentFailure int `json:"custom-gtp-deny-event-records-sent-failure"`
    CustomGtpInfoEventRecordsSent int `json:"custom-gtp-info-event-records-sent"`
    CustomGtpInfoEventRecordsSentFailure int `json:"custom-gtp-info-event-records-sent-failure"`
    CustomFwIddosEntryCreatedRecordsSent int `json:"custom-fw-iddos-entry-created-records-sent"`
    CustomFwIddosEntryCreatedRecordsSentFailure int `json:"custom-fw-iddos-entry-created-records-sent-failure"`
    CustomFwIddosEntryDeletedRecordsSent int `json:"custom-fw-iddos-entry-deleted-records-sent"`
    CustomFwIddosEntryDeletedRecordsSentFailure int `json:"custom-fw-iddos-entry-deleted-records-sent-failure"`
    CustomFwSesnLimitExceededRecordsSent int `json:"custom-fw-sesn-limit-exceeded-records-sent"`
    CustomFwSesnLimitExceededRecordsSentFailure int `json:"custom-fw-sesn-limit-exceeded-records-sent-failure"`
    CustomNatIddosL3EntryCreatedRecordsSent int `json:"custom-nat-iddos-l3-entry-created-records-sent"`
    CustomNatIddosL3EntryCreatedRecordsSentFailure int `json:"custom-nat-iddos-l3-entry-created-records-sent-failure"`
    CustomNatIddosL3EntryDeletedRecordsSent int `json:"custom-nat-iddos-l3-entry-deleted-records-sent"`
    CustomNatIddosL3EntryDeletedRecordsSentFailure int `json:"custom-nat-iddos-l3-entry-deleted-records-sent-failure"`
    CustomNatIddosL4EntryCreatedRecordsSent int `json:"custom-nat-iddos-l4-entry-created-records-sent"`
    CustomNatIddosL4EntryCreatedRecordsSentFailure int `json:"custom-nat-iddos-l4-entry-created-records-sent-failure"`
    CustomNatIddosL4EntryDeletedRecordsSent int `json:"custom-nat-iddos-l4-entry-deleted-records-sent"`
    CustomNatIddosL4EntryDeletedRecordsSentFailure int `json:"custom-nat-iddos-l4-entry-deleted-records-sent-failure"`
    CustomGtpRateLimitPeriodicRecordsSent int `json:"custom-gtp-rate-limit-periodic-records-sent"`
    CustomGtpRateLimitPeriodicRecordsSentFailure int `json:"custom-gtp-rate-limit-periodic-records-sent-failure"`
    DdosGeneralStatRecordsSent int `json:"ddos-general-stat-records-sent"`
    DdosGeneralStatRecordsSentFailure int `json:"ddos-general-stat-records-sent-failure"`
    DdosHttpStatRecordsSent int `json:"ddos-http-stat-records-sent"`
    DdosHttpStatRecordsSentFailure int `json:"ddos-http-stat-records-sent-failure"`
}

func (p *NetflowMonitorStats) GetId() string{
    return "1"
}

func (p *NetflowMonitorStats) getPath() string{
    
    return "netflow/monitor/"+p.Name+"/stats"
}

func (p *NetflowMonitorStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetflowMonitorStats,error) {
logger.Println("NetflowMonitorStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetflowMonitorStats
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
