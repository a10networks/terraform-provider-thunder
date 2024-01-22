

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateSip struct {
	Inst struct {

    Dst DdosZoneTemplateSipDst `json:"dst"`
    FilterHeaderList []DdosZoneTemplateSipFilterHeaderList `json:"filter-header-list"`
    IdleTimeout DdosZoneTemplateSipIdleTimeout `json:"idle-timeout"`
    MalformedSip DdosZoneTemplateSipMalformedSip316 `json:"malformed-sip"`
    MultiPuThresholdDistribution DdosZoneTemplateSipMultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    SipTmplName string `json:"sip-tmpl-name"`
    Src DdosZoneTemplateSipSrc `json:"src"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"sip"`
}


type DdosZoneTemplateSipDst struct {
    SipRequestRateLimit DdosZoneTemplateSipDstSipRequestRateLimit `json:"sip-request-rate-limit"`
}


type DdosZoneTemplateSipDstSipRequestRateLimit struct {
    DstSipRateActionListName string `json:"dst-sip-rate-action-list-name"`
    DstSipRateAction string `json:"dst-sip-rate-action" dval:"drop"`
    Method DdosZoneTemplateSipDstSipRequestRateLimitMethod `json:"method"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethod struct {
    InviteCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodInviteCfg `json:"invite-cfg"`
    RegisterCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodRegisterCfg `json:"register-cfg"`
    OptionsCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodOptionsCfg `json:"options-cfg"`
    ByeCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodByeCfg `json:"bye-cfg"`
    SubscribeCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodSubscribeCfg `json:"subscribe-cfg"`
    NotifyCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodNotifyCfg `json:"notify-cfg"`
    ReferCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodReferCfg `json:"refer-cfg"`
    MessageCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodMessageCfg `json:"message-cfg"`
    UpdateCfg DdosZoneTemplateSipDstSipRequestRateLimitMethodUpdateCfg `json:"update-cfg"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodInviteCfg struct {
    Invite int `json:"INVITE"`
    DstSipInviteRate int `json:"dst-sip-invite-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodRegisterCfg struct {
    Register int `json:"REGISTER"`
    DstSipRegisterRate int `json:"dst-sip-register-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodOptionsCfg struct {
    Options int `json:"OPTIONS"`
    DstSipOptionsRate int `json:"dst-sip-options-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodByeCfg struct {
    Bye int `json:"BYE"`
    DstSipByeRate int `json:"dst-sip-bye-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodSubscribeCfg struct {
    Subscribe int `json:"SUBSCRIBE"`
    DstSipSubscribeRate int `json:"dst-sip-subscribe-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodNotifyCfg struct {
    Notify int `json:"NOTIFY"`
    DstSipNotifyRate int `json:"dst-sip-notify-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodReferCfg struct {
    Refer int `json:"REFER"`
    DstSipReferRate int `json:"dst-sip-refer-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodMessageCfg struct {
    Message int `json:"MESSAGE"`
    DstSipMessageRate int `json:"dst-sip-message-rate"`
}


type DdosZoneTemplateSipDstSipRequestRateLimitMethodUpdateCfg struct {
    Update int `json:"UPDATE"`
    DstSipUpdateRate int `json:"dst-sip-update-rate"`
}


type DdosZoneTemplateSipFilterHeaderList struct {
    SipFilterName string `json:"sip-filter-name"`
    SipFilterHeaderSeq int `json:"sip-filter-header-seq"`
    SipHeaderCfg DdosZoneTemplateSipFilterHeaderListSipHeaderCfg `json:"sip-header-cfg"`
    SipFilterActionListName string `json:"sip-filter-action-list-name"`
    SipFilterAction string `json:"sip-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateSipFilterHeaderListSipHeaderCfg struct {
    SipFilterHeaderRegex string `json:"sip-filter-header-regex"`
    SipFilterHeaderInverseMatch int `json:"sip-filter-header-inverse-match"`
}


type DdosZoneTemplateSipIdleTimeout struct {
    IdleTimeoutValue int `json:"idle-timeout-value"`
    IgnoreZeroPayload int `json:"ignore-zero-payload"`
    IdleTimeoutActionListName string `json:"idle-timeout-action-list-name"`
    IdleTimeoutAction string `json:"idle-timeout-action"`
}


type DdosZoneTemplateSipMalformedSip316 struct {
    MalformedSipCheck string `json:"malformed-sip-check"`
    MalformedSipMaxLineSize int `json:"malformed-sip-max-line-size" dval:"32511"`
    MalformedSipMaxUriLength int `json:"malformed-sip-max-uri-length" dval:"32511"`
    MalformedSipMaxHeaderNameLength int `json:"malformed-sip-max-header-name-length" dval:"63"`
    MalformedSipMaxHeaderValueLength int `json:"malformed-sip-max-header-value-length" dval:"32511"`
    MalformedSipCallIdMaxLength int `json:"malformed-sip-call-id-max-length" dval:"32511"`
    MalformedSipSdpMaxLength int `json:"malformed-sip-sdp-max-length" dval:"32511"`
    MalformedSipActionListName string `json:"malformed-sip-action-list-name"`
    MalformedSipAction string `json:"malformed-sip-action" dval:"drop"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateSipMultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosZoneTemplateSipSrc struct {
    SipRequestRateLimit DdosZoneTemplateSipSrcSipRequestRateLimit `json:"sip-request-rate-limit"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimit struct {
    SrcSipRateActionListName string `json:"src-sip-rate-action-list-name"`
    SrcSipRateAction string `json:"src-sip-rate-action" dval:"drop"`
    Method DdosZoneTemplateSipSrcSipRequestRateLimitMethod `json:"method"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethod struct {
    InviteCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodInviteCfg `json:"invite-cfg"`
    RegisterCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodRegisterCfg `json:"register-cfg"`
    OptionsCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodOptionsCfg `json:"options-cfg"`
    ByeCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodByeCfg `json:"bye-cfg"`
    SubscribeCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg `json:"subscribe-cfg"`
    NotifyCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodNotifyCfg `json:"notify-cfg"`
    ReferCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodReferCfg `json:"refer-cfg"`
    MessageCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodMessageCfg `json:"message-cfg"`
    UpdateCfg DdosZoneTemplateSipSrcSipRequestRateLimitMethodUpdateCfg `json:"update-cfg"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodInviteCfg struct {
    Invite int `json:"INVITE"`
    SrcSipInviteRate int `json:"src-sip-invite-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodRegisterCfg struct {
    Register int `json:"REGISTER"`
    SrcSipRegisterRate int `json:"src-sip-register-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodOptionsCfg struct {
    Options int `json:"OPTIONS"`
    SrcSipOptionsRate int `json:"src-sip-options-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodByeCfg struct {
    Bye int `json:"BYE"`
    SrcSipByeRate int `json:"src-sip-bye-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg struct {
    Subscribe int `json:"SUBSCRIBE"`
    SrcSipSubscribeRate int `json:"src-sip-subscribe-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodNotifyCfg struct {
    Notify int `json:"NOTIFY"`
    SrcSipNotifyRate int `json:"src-sip-notify-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodReferCfg struct {
    Refer int `json:"REFER"`
    SrcSipReferRate int `json:"src-sip-refer-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodMessageCfg struct {
    Message int `json:"MESSAGE"`
    SrcSipMessageRate int `json:"src-sip-message-rate"`
}


type DdosZoneTemplateSipSrcSipRequestRateLimitMethodUpdateCfg struct {
    Update int `json:"UPDATE"`
    SrcSipUpdateRate int `json:"src-sip-update-rate"`
}

func (p *DdosZoneTemplateSip) GetId() string{
    return url.QueryEscape(p.Inst.SipTmplName)
}

func (p *DdosZoneTemplateSip) getPath() string{
    return "ddos/zone-template/sip"
}

func (p *DdosZoneTemplateSip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSip::Post")
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

func (p *DdosZoneTemplateSip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSip::Get")
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
func (p *DdosZoneTemplateSip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSip::Put")
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

func (p *DdosZoneTemplateSip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateSip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
