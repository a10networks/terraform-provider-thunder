

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityTopnGtpApnPrefixTopnTmpl struct {
	Inst struct {

    Interval string `json:"interval"`
    Metrics VisibilityTopnGtpApnPrefixTopnTmplMetrics3132 `json:"metrics"`
    Name string `json:"name"`
    TopnSize int `json:"topn-size"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"gtp-apn-prefix-topn-tmpl"`
}


type VisibilityTopnGtpApnPrefixTopnTmplMetrics3132 struct {
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
    DropVldGtpV2WrongLbiCreateBearer int `json:"drop-vld-gtp-v2-wrong-lbi-create-bearer"`
    GtpCHandoverInProgressWithConn int `json:"gtp-c-handover-in-progress-with-conn"`
    DropVldInvalidPktLenPiggyback int `json:"drop-vld-invalid-pkt-len-piggyback"`
    DropVldSanityFailedPiggyback int `json:"drop-vld-sanity-failed-piggyback"`
    DropVldSequenceNumCorrelation int `json:"drop-vld-sequence-num-correlation"`
    DropVldGtpv0SeqnumBufferFull int `json:"drop-vld-gtpv0-seqnum-buffer-full"`
    DropVldGtpv1SeqnumBufferFull int `json:"drop-vld-gtpv1-seqnum-buffer-full"`
    DropVldGtpv2SeqnumBufferFull int `json:"drop-vld-gtpv2-seqnum-buffer-full"`
    DropVldGtpInvalidImsiLenDrop int `json:"drop-vld-gtp-invalid-imsi-len-drop"`
    DropVldGtpInvalidApnLenDrop int `json:"drop-vld-gtp-invalid-apn-len-drop"`
    DropVldProtocolFlagUnset int `json:"drop-vld-protocol-flag-unset"`
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
    Uuid string `json:"uuid"`
}

func (p *VisibilityTopnGtpApnPrefixTopnTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityTopnGtpApnPrefixTopnTmpl) getPath() string{
    return "visibility/topn/gtp-apn-prefix-topn-tmpl"
}

func (p *VisibilityTopnGtpApnPrefixTopnTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpApnPrefixTopnTmpl::Post")
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

func (p *VisibilityTopnGtpApnPrefixTopnTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpApnPrefixTopnTmpl::Get")
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
func (p *VisibilityTopnGtpApnPrefixTopnTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpApnPrefixTopnTmpl::Put")
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

func (p *VisibilityTopnGtpApnPrefixTopnTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityTopnGtpApnPrefixTopnTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
