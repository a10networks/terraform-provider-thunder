

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate struct {
	Inst struct {

    DropVldCountryCodeMismatch int `json:"drop-vld-country-code-mismatch"`
    DropVldCrossLayerCorrelation int `json:"drop-vld-cross-layer-correlation"`
    DropVldGtpBearerCountExceed int `json:"drop-vld-gtp-bearer-count-exceed"`
    DropVldGtpIeRepeatCountExceed int `json:"drop-vld-gtp-ie-repeat-count-exceed"`
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
    DropVldMandatoryIeInGroupedIe int `json:"drop-vld-mandatory-ie-in-grouped-ie"`
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
    DropVldV0ReservedMessageDrop int `json:"drop-vld-v0-reserved-message-drop"`
    DropVldV1ReservedMessageDrop int `json:"drop-vld-v1-reserved-message-drop"`
    DropVldV2ReservedMessageDrop int `json:"drop-vld-v2-reserved-message-drop"`
    DropVldVersionNotSupported int `json:"drop-vld-version-not-supported"`
    Duration int `json:"duration" dval:"60"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/object-templates/templ-gtp-plcy-tmpl/" +p.Inst.Name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesTemplGtpPlcyTmplTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
