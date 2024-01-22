

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwGtpApnPrefixStats struct {
    
    Stats FwGtpApnPrefixStatsStats `json:"stats"`

}
type DataFwGtpApnPrefixStats struct {
    DtFwGtpApnPrefixStats FwGtpApnPrefixStats `json:"apn-prefix"`
}


type FwGtpApnPrefixStatsStats struct {
    KeyName string `json:"key-name"`
    UplinkBytes int `json:"uplink-bytes"`
    DownlinkBytes int `json:"downlink-bytes"`
    UplinkPkts int `json:"uplink-pkts"`
    DownlinkPkts int `json:"downlink-pkts"`
    GtpV0CTunnelCreated int `json:"gtp-v0-c-tunnel-created"`
    GtpV0CTunnelHalfOpen int `json:"gtp-v0-c-tunnel-half-open"`
    GtpV0CTunnelHalfClosed int `json:"gtp-v0-c-tunnel-half-closed"`
    GtpV0CTunnelClosed int `json:"gtp-v0-c-tunnel-closed"`
    GtpV0CTunnelDeleted int `json:"gtp-v0-c-tunnel-deleted"`
    GtpV0CHalfOpenTunnelClosed int `json:"gtp-v0-c-half-open-tunnel-closed"`
    GtpV1CTunnelCreated int `json:"gtp-v1-c-tunnel-created"`
    GtpV1CTunnelHalfOpen int `json:"gtp-v1-c-tunnel-half-open"`
    GtpV1CTunnelHalfClosed int `json:"gtp-v1-c-tunnel-half-closed"`
    GtpV1CTunnelClosed int `json:"gtp-v1-c-tunnel-closed"`
    GtpV1CTunnelDeleted int `json:"gtp-v1-c-tunnel-deleted"`
    GtpV1CHalfOpenTunnelClosed int `json:"gtp-v1-c-half-open-tunnel-closed"`
    GtpV2CTunnelCreated int `json:"gtp-v2-c-tunnel-created"`
    GtpV2CTunnelHalfOpen int `json:"gtp-v2-c-tunnel-half-open"`
    GtpV2CTunnelHalfClosed int `json:"gtp-v2-c-tunnel-half-closed"`
    GtpV2CTunnelClosed int `json:"gtp-v2-c-tunnel-closed"`
    GtpV2CTunnelDeleted int `json:"gtp-v2-c-tunnel-deleted"`
    GtpV2CHalfOpenTunnelClosed int `json:"gtp-v2-c-half-open-tunnel-closed"`
    GtpUTunnelCreated int `json:"gtp-u-tunnel-created"`
    GtpUTunnelDeleted int `json:"gtp-u-tunnel-deleted"`
    GtpV0CUpdatePdpRespUnsuccess int `json:"gtp-v0-c-update-pdp-resp-unsuccess"`
    GtpV1CUpdatePdpRespUnsuccess int `json:"gtp-v1-c-update-pdp-resp-unsuccess"`
    GtpV2CMod_bearerRespUnsuccess int `json:"gtp-v2-c-mod_bearer-resp-unsuccess"`
    GtpV0CCreatePdpRespUnsuccess int `json:"gtp-v0-c-create-pdp-resp-unsuccess"`
    GtpV1CCreatePdpRespUnsuccess int `json:"gtp-v1-c-create-pdp-resp-unsuccess"`
    GtpV2CCreateSessRespUnsuccess int `json:"gtp-v2-c-create-sess-resp-unsuccess"`
    GtpV2CPiggybackMessage int `json:"gtp-v2-c-piggyback-message"`
    GtpPathManagementMessage int `json:"gtp-path-management-message"`
    GtpV0CTunnelDeletedRestart int `json:"gtp-v0-c-tunnel-deleted-restart"`
    GtpV1CTunnelDeletedRestart int `json:"gtp-v1-c-tunnel-deleted-restart"`
    GtpV2CTunnelDeletedRestart int `json:"gtp-v2-c-tunnel-deleted-restart"`
    GtpV0CReservedMessageAllow int `json:"gtp-v0-c-reserved-message-allow"`
    GtpV1CReservedMessageAllow int `json:"gtp-v1-c-reserved-message-allow"`
    GtpV2CReservedMessageAllow int `json:"gtp-v2-c-reserved-message-allow"`
    GtpV1CPduNotificationRequestForward int `json:"gtp-v1-c-pdu-notification-request-forward"`
    GtpV1CPduNotificationRejectRequestForward int `json:"gtp-v1-c-pdu-notification-reject-request-forward"`
    GtpV0CPduNotificationRequestForward int `json:"gtp-v0-c-pdu-notification-request-forward"`
    GtpV0CPduNotificationRejectRequestForward int `json:"gtp-v0-c-pdu-notification-reject-request-forward"`
    GtpV0CMessageSkippedApnFilteringNoImsi int `json:"gtp-v0-c-message-skipped-apn-filtering-no-imsi"`
    GtpV1CMessageSkippedApnFilteringNoImsi int `json:"gtp-v1-c-message-skipped-apn-filtering-no-imsi"`
    GtpV2CMessageSkippedApnFilteringNoImsi int `json:"gtp-v2-c-message-skipped-apn-filtering-no-imsi"`
    GtpV0CMessageSkippedMsisdnFilteringNoMsisdn int `json:"gtp-v0-c-message-skipped-msisdn-filtering-no-msisdn"`
    GtpV1CMessageSkippedMsisdnFilteringNoMsisdn int `json:"gtp-v1-c-message-skipped-msisdn-filtering-no-msisdn"`
    GtpV2CMessageSkippedMsisdnFilteringNoMsisdn int `json:"gtp-v2-c-message-skipped-msisdn-filtering-no-msisdn"`
    GtpV0CPacketDummyMsisdn int `json:"gtp-v0-c-packet-dummy-msisdn"`
    GtpV1CPacketDummyMsisdn int `json:"gtp-v1-c-packet-dummy-msisdn"`
    GtpV2CPacketDummyMsisdn int `json:"gtp-v2-c-packet-dummy-msisdn"`
    DropVldGtpV2CMessageWithTeidZeroExpected int `json:"drop-vld-gtp-v2-c-message-with-teid-zero-expected"`
    DropVldGtpV1CMessageWithTeidZeroExpected int `json:"drop-vld-gtp-v1-c-message-with-teid-zero-expected"`
    DropVldGtpV0CMessageWithTeidZeroExpected int `json:"drop-vld-gtp-v0-c-message-with-teid-zero-expected"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldUnsupportedMessageType int `json:"drop-vld-unsupported-message-type"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    DropVldGtpBearerCountExceed int `json:"drop-vld-gtp-bearer-count-exceed"`
    DropVldGtpV2WrongLbiCreateBearerReq int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer-req"`
    GtpCHandoverInProgressWithConn int `json:"gtp-c-handover-in-progress-with-conn"`
    DropVldV0ReservedMessageDrop int `json:"drop-vld-v0-reserved-message-drop"`
    DropVldV1ReservedMessageDrop int `json:"drop-vld-v1-reserved-message-drop"`
    DropVldV2ReservedMessageDrop int `json:"drop-vld-v2-reserved-message-drop"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    DropVldGtpv0SubscriberAttrMiss int `json:"drop-vld-gtpv0-subscriber-attr-miss"`
    DropVldGtpv1SubscriberAttrMiss int `json:"drop-vld-gtpv1-subscriber-attr-miss"`
    DropVldGtpv2SubscriberAttrMiss int `json:"drop-vld-gtpv2-subscriber-attr-miss"`
    DropVldGtpV0CIeLenExceedMsgLen int `json:"drop-vld-gtp-v0-c-ie-len-exceed-msg-len"`
    DropVldGtpV1CIeLenExceedMsgLen int `json:"drop-vld-gtp-v1-c-ie-len-exceed-msg-len"`
    DropVldGtpV2CIeLenExceedMsgLen int `json:"drop-vld-gtp-v2-c-ie-len-exceed-msg-len"`
    DropVldGtpV0CMessageLengthMismatch int `json:"drop-vld-gtp-v0-c-message-length-mismatch"`
    DropVldGtpV1CMessageLengthMismatch int `json:"drop-vld-gtp-v1-c-message-length-mismatch"`
    DropVldGtpV2CMessageLengthMismatch int `json:"drop-vld-gtp-v2-c-message-length-mismatch"`
    DropVldGtpV0CMessageDroppedApnFilteringNoApn int `json:"drop-vld-gtp-v0-c-message-dropped-apn-filtering-no-apn"`
    DropVldGtpV1CMessageDroppedApnFilteringNoApn int `json:"drop-vld-gtp-v1-c-message-dropped-apn-filtering-no-apn"`
    DropVldGtpV2CMessageDroppedApnFilteringNoApn int `json:"drop-vld-gtp-v2-c-message-dropped-apn-filtering-no-apn"`
    DropFltMessageFiltering int `json:"drop-flt-message-filtering"`
    DropFltApnFiltering int `json:"drop-flt-apn-filtering"`
    DropFltMsisdnFiltering int `json:"drop-flt-msisdn-filtering"`
    DropFltRatTypeFiltering int `json:"drop-flt-rat-type-filtering"`
    DropFltGtpInGtp int `json:"drop-flt-gtp-in-gtp"`
    DropRlGtpV0CAgg int `json:"drop-rl-gtp-v0-c-agg"`
    DropRlGtpV1CAgg int `json:"drop-rl-gtp-v1-c-agg"`
    DropRlGtpV2CAgg int `json:"drop-rl-gtp-v2-c-agg"`
    DropRlGtpV1CCreatePdpRequest int `json:"drop-rl-gtp-v1-c-create-pdp-request"`
    DropRlGtpV2CCreateSessionRequest int `json:"drop-rl-gtp-v2-c-create-session-request"`
    DropRlGtpV1CUpdatePdpRequest int `json:"drop-rl-gtp-v1-c-update-pdp-request"`
    DropRlGtpV2CModifyBearerRequest int `json:"drop-rl-gtp-v2-c-modify-bearer-request"`
    DropRlGtpUTunnelCreate int `json:"drop-rl-gtp-u-tunnel-create"`
    DropRlGtpUUplinkByte int `json:"drop-rl-gtp-u-uplink-byte"`
    DropRlGtpUUplinkPacket int `json:"drop-rl-gtp-u-uplink-packet"`
    DropRlGtpUDownlinkByte int `json:"drop-rl-gtp-u-downlink-byte"`
    DropRlGtpUDownlinkPacket int `json:"drop-rl-gtp-u-downlink-packet"`
    DropRlGtpUTotalByte int `json:"drop-rl-gtp-u-total-byte"`
    DropRlGtpUTotalPacket int `json:"drop-rl-gtp-u-total-packet"`
    DropRlGtpUMaxConcurrentTunnels int `json:"drop-rl-gtp-u-max-concurrent-tunnels"`
}

func (p *FwGtpApnPrefixStats) GetId() string{
    return "1"
}

func (p *FwGtpApnPrefixStats) getPath() string{
    return "fw/gtp/apn-prefix/stats"
}

func (p *FwGtpApnPrefixStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwGtpApnPrefixStats,error) {
logger.Println("FwGtpApnPrefixStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwGtpApnPrefixStats
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
