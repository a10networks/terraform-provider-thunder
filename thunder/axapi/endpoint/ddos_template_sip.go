

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateSip struct {
	Inst struct {

    Action string `json:"action" dval:"drop"`
    Dst DdosTemplateSipDst `json:"dst"`
    FilterHeaderList []DdosTemplateSipFilterHeaderList `json:"filter-header-list"`
    IdleTimeout int `json:"idle-timeout"`
    IgnoreZeroPayload int `json:"ignore-zero-payload"`
    MalformedSip DdosTemplateSipMalformedSip298 `json:"malformed-sip"`
    MultiPuThresholdDistribution DdosTemplateSipMultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    SipTmplName string `json:"sip-tmpl-name"`
    Src DdosTemplateSipSrc `json:"src"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"sip"`
}


type DdosTemplateSipDst struct {
    SipRequestRateLimit DdosTemplateSipDstSipRequestRateLimit `json:"sip-request-rate-limit"`
}


type DdosTemplateSipDstSipRequestRateLimit struct {
    Method DdosTemplateSipDstSipRequestRateLimitMethod `json:"method"`
}


type DdosTemplateSipDstSipRequestRateLimitMethod struct {
    InviteCfg DdosTemplateSipDstSipRequestRateLimitMethodInviteCfg `json:"invite-cfg"`
    RegisterCfg DdosTemplateSipDstSipRequestRateLimitMethodRegisterCfg `json:"register-cfg"`
    OptionsCfg DdosTemplateSipDstSipRequestRateLimitMethodOptionsCfg `json:"options-cfg"`
    ByeCfg DdosTemplateSipDstSipRequestRateLimitMethodByeCfg `json:"bye-cfg"`
    SubscribeCfg DdosTemplateSipDstSipRequestRateLimitMethodSubscribeCfg `json:"subscribe-cfg"`
    NotifyCfg DdosTemplateSipDstSipRequestRateLimitMethodNotifyCfg `json:"notify-cfg"`
    ReferCfg DdosTemplateSipDstSipRequestRateLimitMethodReferCfg `json:"refer-cfg"`
    MessageCfg DdosTemplateSipDstSipRequestRateLimitMethodMessageCfg `json:"message-cfg"`
    UpdateCfg DdosTemplateSipDstSipRequestRateLimitMethodUpdateCfg `json:"update-cfg"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodInviteCfg struct {
    DstSipInviteCfgFlag int `json:"dst-sip-invite-cfg-flag"`
    DstSipInviteRate int `json:"dst-sip-invite-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodRegisterCfg struct {
    DstSipRegisterCfgFlag int `json:"dst-sip-register-cfg-flag"`
    DstSipRegisterRate int `json:"dst-sip-register-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodOptionsCfg struct {
    DstSipOptionsCfgFlag int `json:"dst-sip-options-cfg-flag"`
    DstSipOptionsRate int `json:"dst-sip-options-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodByeCfg struct {
    DstSipByeCfgFlag int `json:"dst-sip-bye-cfg-flag"`
    DstSipByeRate int `json:"dst-sip-bye-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodSubscribeCfg struct {
    DstSipSubscribeCfgFlag int `json:"dst-sip-subscribe-cfg-flag"`
    DstSipSubscribeRate int `json:"dst-sip-subscribe-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodNotifyCfg struct {
    DstSipNotifyCfgFlag int `json:"dst-sip-notify-cfg-flag"`
    DstSipNotifyRate int `json:"dst-sip-notify-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodReferCfg struct {
    DstSipReferCfgFlag int `json:"dst-sip-refer-cfg-flag"`
    DstSipReferRate int `json:"dst-sip-refer-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodMessageCfg struct {
    DstSipMessageCfgFlag int `json:"dst-sip-message-cfg-flag"`
    DstSipMessageRate int `json:"dst-sip-message-rate"`
}


type DdosTemplateSipDstSipRequestRateLimitMethodUpdateCfg struct {
    DstSipUpdateCfgFlag int `json:"dst-sip-update-cfg-flag"`
    DstSipUpdateRate int `json:"dst-sip-update-rate"`
}


type DdosTemplateSipFilterHeaderList struct {
    SipFilterHeaderSeq int `json:"sip-filter-header-seq"`
    SipFilterHeaderRegex string `json:"sip-filter-header-regex"`
    SipFilterHeaderUnmatched int `json:"sip-filter-header-unmatched"`
    SipFilterHeaderBlacklist int `json:"sip-filter-header-blacklist"`
    SipFilterHeaderWhitelist int `json:"sip-filter-header-whitelist"`
    SipFilterHeaderCountOnly int `json:"sip-filter-header-count-only"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosTemplateSipMalformedSip298 struct {
    MalformedSipCheck string `json:"malformed-sip-check"`
    MalformedSipMaxLineSize int `json:"malformed-sip-max-line-size" dval:"32511"`
    MalformedSipMaxUriLength int `json:"malformed-sip-max-uri-length" dval:"32511"`
    MalformedSipMaxHeaderNameLength int `json:"malformed-sip-max-header-name-length" dval:"63"`
    MalformedSipMaxHeaderValueLength int `json:"malformed-sip-max-header-value-length" dval:"32511"`
    MalformedSipCallIdMaxLength int `json:"malformed-sip-call-id-max-length" dval:"32511"`
    MalformedSipSdpMaxLength int `json:"malformed-sip-sdp-max-length" dval:"32511"`
    Uuid string `json:"uuid"`
}


type DdosTemplateSipMultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosTemplateSipSrc struct {
    SipRequestRateLimit DdosTemplateSipSrcSipRequestRateLimit `json:"sip-request-rate-limit"`
}


type DdosTemplateSipSrcSipRequestRateLimit struct {
    Method DdosTemplateSipSrcSipRequestRateLimitMethod `json:"method"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethod struct {
    InviteCfg DdosTemplateSipSrcSipRequestRateLimitMethodInviteCfg `json:"invite-cfg"`
    RegisterCfg DdosTemplateSipSrcSipRequestRateLimitMethodRegisterCfg `json:"register-cfg"`
    OptionsCfg DdosTemplateSipSrcSipRequestRateLimitMethodOptionsCfg `json:"options-cfg"`
    ByeCfg DdosTemplateSipSrcSipRequestRateLimitMethodByeCfg `json:"bye-cfg"`
    SubscribeCfg DdosTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg `json:"subscribe-cfg"`
    NotifyCfg DdosTemplateSipSrcSipRequestRateLimitMethodNotifyCfg `json:"notify-cfg"`
    ReferCfg DdosTemplateSipSrcSipRequestRateLimitMethodReferCfg `json:"refer-cfg"`
    MessageCfg DdosTemplateSipSrcSipRequestRateLimitMethodMessageCfg `json:"message-cfg"`
    UpdateCfg DdosTemplateSipSrcSipRequestRateLimitMethodUpdateCfg `json:"update-cfg"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodInviteCfg struct {
    SrcSipInviteCfgFlag int `json:"src-sip-invite-cfg-flag"`
    SrcSipInviteRate int `json:"src-sip-invite-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodRegisterCfg struct {
    SrcSipRegisterCfgFlag int `json:"src-sip-register-cfg-flag"`
    SrcSipRegisterRate int `json:"src-sip-register-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodOptionsCfg struct {
    SrcSipOptionsCfgFlag int `json:"src-sip-options-cfg-flag"`
    SrcSipOptionsRate int `json:"src-sip-options-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodByeCfg struct {
    SrcSipByeCfgFlag int `json:"src-sip-bye-cfg-flag"`
    SrcSipByeRate int `json:"src-sip-bye-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodSubscribeCfg struct {
    SrcSipSubscribeCfgFlag int `json:"src-sip-subscribe-cfg-flag"`
    SrcSipSubscribeRate int `json:"src-sip-subscribe-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodNotifyCfg struct {
    SrcSipNotifyCfgFlag int `json:"src-sip-notify-cfg-flag"`
    SrcSipNotifyRate int `json:"src-sip-notify-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodReferCfg struct {
    SrcSipReferCfgFlag int `json:"src-sip-refer-cfg-flag"`
    SrcSipReferRate int `json:"src-sip-refer-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodMessageCfg struct {
    SrcSipMessageCfgFlag int `json:"src-sip-message-cfg-flag"`
    SrcSipMessageRate int `json:"src-sip-message-rate"`
}


type DdosTemplateSipSrcSipRequestRateLimitMethodUpdateCfg struct {
    SrcSipUpdateCfgFlag int `json:"src-sip-update-cfg-flag"`
    SrcSipUpdateRate int `json:"src-sip-update-rate"`
}

func (p *DdosTemplateSip) GetId() string{
    return url.QueryEscape(p.Inst.SipTmplName)
}

func (p *DdosTemplateSip) getPath() string{
    return "ddos/template/sip"
}

func (p *DdosTemplateSip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSip::Post")
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

func (p *DdosTemplateSip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSip::Get")
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
func (p *DdosTemplateSip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSip::Put")
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

func (p *DdosTemplateSip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
