

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityTopnGtpNetworkElementTopnTmplMetrics struct {
	Inst struct {

    DownlinkBytes int `json:"downlink-bytes"`
    DownlinkPkts int `json:"downlink-pkts"`
    DropFltApnFiltering int `json:"drop-flt-apn-filtering"`
    DropFltGtpInGtp int `json:"drop-flt-gtp-in-gtp"`
    DropFltMessageFiltering int `json:"drop-flt-message-filtering"`
    DropFltMsisdnFiltering int `json:"drop-flt-msisdn-filtering"`
    DropFltRatTypeFiltering int `json:"drop-flt-rat-type-filtering"`
    DropRlGtpUDownlinkByte int `json:"drop-rl-gtp-u-downlink-byte"`
    DropRlGtpUDownlinkPacket int `json:"drop-rl-gtp-u-downlink-packet"`
    DropRlGtpUMaxConcurrentTunnels int `json:"drop-rl-gtp-u-max-concurrent-tunnels"`
    DropRlGtpUTotalByte int `json:"drop-rl-gtp-u-total-byte"`
    DropRlGtpUTotalPacket int `json:"drop-rl-gtp-u-total-packet"`
    DropRlGtpUTunnelCreate int `json:"drop-rl-gtp-u-tunnel-create"`
    DropRlGtpUUplinkByte int `json:"drop-rl-gtp-u-uplink-byte"`
    DropRlGtpUUplinkPacket int `json:"drop-rl-gtp-u-uplink-packet"`
    DropRlGtpV0CAgg int `json:"drop-rl-gtp-v0-c-agg"`
    DropRlGtpV1CAgg int `json:"drop-rl-gtp-v1-c-agg"`
    DropRlGtpV1CCreatePdpRequest int `json:"drop-rl-gtp-v1-c-create-pdp-request"`
    DropRlGtpV1CUpdatePdpRequest int `json:"drop-rl-gtp-v1-c-update-pdp-request"`
    DropRlGtpV2CAgg int `json:"drop-rl-gtp-v2-c-agg"`
    DropRlGtpV2CCreateSessionRequest int `json:"drop-rl-gtp-v2-c-create-session-request"`
    DropRlGtpV2CModifyBearerRequest int `json:"drop-rl-gtp-v2-c-modify-bearer-request"`
    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpUSpoofedSourceAddress int `json:"drop-vld-gtp-u-spoofed-source-address"`
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldInvalidFlowLabelV0 int `json:"drop-vld-invalid-flow-label-v0"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldInvalidTeid int `json:"drop-vld-invalid-teid"`
    DropVldMandatoryInformationElement int `json:"drop-vld-mandatory-information-element"`
    DropVldMessageLength int `json:"drop-vld-message-length"`
    DropVldOutOfOrderIe int `json:"drop-vld-out-of-order-ie"`
    DropVldOutOfState int `json:"drop-vld-out-of-state"`
    DropVldOutOfStateIe int `json:"drop-vld-out-of-state-ie"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
    DropVldReservedFieldSet int `json:"drop-vld-reserved-field-set"`
    DropVldReservedInformationElement int `json:"drop-vld-reserved-information-element"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldTunnelIdFlag int `json:"drop-vld-tunnel-id-flag"`
    DropVldUnsupportedMessageType int `json:"drop-vld-unsupported-message-type"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    Drop_vldGtpBearerCountExceed int `json:"drop_vld-gtp-bearer-count-exceed"`
    GtpCHandoverInProgressWithConn int `json:"gtp-c-handover-in-progress-with-conn"`
    GtpPathManagementMessage int `json:"gtp-path-management-message"`
    GtpUTunnelCreated int `json:"gtp-u-tunnel-created"`
    GtpUTunnelDeleted int `json:"gtp-u-tunnel-deleted"`
    GtpV0CCreatePdpRespUnsuccess int `json:"gtp-v0-c-create-pdp-resp-unsuccess"`
    GtpV0CHalfOpenTunnelClosed int `json:"gtp-v0-c-half-open-tunnel-closed"`
    GtpV0CReservedMessageAllow int `json:"gtp-v0-c-reserved-message-allow"`
    GtpV0CTunnelClosed int `json:"gtp-v0-c-tunnel-closed"`
    GtpV0CTunnelCreated int `json:"gtp-v0-c-tunnel-created"`
    GtpV0CTunnelDeleted int `json:"gtp-v0-c-tunnel-deleted"`
    GtpV0CTunnelDeletedRestart int `json:"gtp-v0-c-tunnel-deleted-restart"`
    GtpV0CTunnelHalfClosed int `json:"gtp-v0-c-tunnel-half-closed"`
    GtpV0CTunnelHalfOpen int `json:"gtp-v0-c-tunnel-half-open"`
    GtpV0CUpdatePdpRespUnsuccess int `json:"gtp-v0-c-update-pdp-resp-unsuccess"`
    GtpV1CCreatePdpRespUnsuccess int `json:"gtp-v1-c-create-pdp-resp-unsuccess"`
    GtpV1CHalfOpenTunnelClosed int `json:"gtp-v1-c-half-open-tunnel-closed"`
    GtpV1CReservedMessageAllow int `json:"gtp-v1-c-reserved-message-allow"`
    GtpV1CTunnelClosed int `json:"gtp-v1-c-tunnel-closed"`
    GtpV1CTunnelCreated int `json:"gtp-v1-c-tunnel-created"`
    GtpV1CTunnelDeleted int `json:"gtp-v1-c-tunnel-deleted"`
    GtpV1CTunnelDeletedRestart int `json:"gtp-v1-c-tunnel-deleted-restart"`
    GtpV1CTunnelHalfClosed int `json:"gtp-v1-c-tunnel-half-closed"`
    GtpV1CTunnelHalfOpen int `json:"gtp-v1-c-tunnel-half-open"`
    GtpV1CUpdatePdpRespUnsuccess int `json:"gtp-v1-c-update-pdp-resp-unsuccess"`
    GtpV2CCreateSessRespUnsuccess int `json:"gtp-v2-c-create-sess-resp-unsuccess"`
    GtpV2CHalfOpenTunnelClosed int `json:"gtp-v2-c-half-open-tunnel-closed"`
    GtpV2CMod_bearerRespUnsuccess int `json:"gtp-v2-c-mod_bearer-resp-unsuccess"`
    GtpV2CPiggybackMessage int `json:"gtp-v2-c-piggyback-message"`
    GtpV2CReservedMessageAllow int `json:"gtp-v2-c-reserved-message-allow"`
    GtpV2CTunnelClosed int `json:"gtp-v2-c-tunnel-closed"`
    GtpV2CTunnelCreated int `json:"gtp-v2-c-tunnel-created"`
    GtpV2CTunnelDeleted int `json:"gtp-v2-c-tunnel-deleted"`
    GtpV2CTunnelDeletedRestart int `json:"gtp-v2-c-tunnel-deleted-restart"`
    GtpV2CTunnelHalfClosed int `json:"gtp-v2-c-tunnel-half-closed"`
    GtpV2CTunnelHalfOpen int `json:"gtp-v2-c-tunnel-half-open"`
    UplinkBytes int `json:"uplink-bytes"`
    UplinkPkts int `json:"uplink-pkts"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"metrics"`
}

func (p *VisibilityTopnGtpNetworkElementTopnTmplMetrics) GetId() string{
    return "1"
}

func (p *VisibilityTopnGtpNetworkElementTopnTmplMetrics) getPath() string{
    return "visibility/topn/gtp-network-element-topn-tmpl/" +p.Inst.Name + "/metrics"
}

func (p *VisibilityTopnGtpNetworkElementTopnTmplMetrics) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpNetworkElementTopnTmplMetrics::Post")
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

func (p *VisibilityTopnGtpNetworkElementTopnTmplMetrics) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpNetworkElementTopnTmplMetrics::Get")
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
func (p *VisibilityTopnGtpNetworkElementTopnTmplMetrics) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpNetworkElementTopnTmplMetrics::Put")
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

func (p *VisibilityTopnGtpNetworkElementTopnTmplMetrics) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpNetworkElementTopnTmplMetrics::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
