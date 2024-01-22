

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateTcp struct {
	Inst struct {

    AckAuthentication DdosZoneTemplateTcpAckAuthentication `json:"ack-authentication"`
    AckAuthenticationSynackReset int `json:"ack-authentication-synack-reset"`
    ActionOnAckRtoRetryCount int `json:"action-on-ack-rto-retry-count"`
    ActionOnSynRtoRetryCount int `json:"action-on-syn-rto-retry-count"`
    Age int `json:"age" dval:"2"`
    AllowSynOtherflags int `json:"allow-syn-otherflags"`
    AllowSynackSkipAuthentications int `json:"allow-synack-skip-authentications"`
    AllowTcpTfo int `json:"allow-tcp-tfo"`
    Concurrent int `json:"concurrent"`
    ConnRateLimitOnSynOnly int `json:"conn-rate-limit-on-syn-only"`
    CreateConnOnSynOnly int `json:"create-conn-on-syn-only"`
    Dst DdosZoneTemplateTcpDst `json:"dst"`
    FilterList []DdosZoneTemplateTcpFilterList `json:"filter-list"`
    FilterMatchType string `json:"filter-match-type" dval:"default"`
    KnownRespSrcPortCfg DdosZoneTemplateTcpKnownRespSrcPortCfg `json:"known-resp-src-port-cfg"`
    MaxRexmitSynPerFlowCfg DdosZoneTemplateTcpMaxRexmitSynPerFlowCfg `json:"max-rexmit-syn-per-flow-cfg"`
    Name string `json:"name"`
    OutOfSeqCfg DdosZoneTemplateTcpOutOfSeqCfg `json:"out-of-seq-cfg"`
    PerConnOutOfSeqRateCfg DdosZoneTemplateTcpPerConnOutOfSeqRateCfg `json:"per-conn-out-of-seq-rate-cfg"`
    PerConnPktRateCfg DdosZoneTemplateTcpPerConnPktRateCfg `json:"per-conn-pkt-rate-cfg"`
    PerConnRateInterval string `json:"per-conn-rate-interval" dval:"1sec"`
    PerConnRetransmitRateCfg DdosZoneTemplateTcpPerConnRetransmitRateCfg `json:"per-conn-retransmit-rate-cfg"`
    PerConnZeroWinRateCfg DdosZoneTemplateTcpPerConnZeroWinRateCfg `json:"per-conn-zero-win-rate-cfg"`
    ProgressionTracking DdosZoneTemplateTcpProgressionTracking320 `json:"progression-tracking"`
    RetransmitCfg DdosZoneTemplateTcpRetransmitCfg `json:"retransmit-cfg"`
    Src DdosZoneTemplateTcpSrc `json:"src"`
    SynAuthentication DdosZoneTemplateTcpSynAuthentication `json:"syn-authentication"`
    SynCookie int `json:"syn-cookie"`
    SynackRateLimit int `json:"synack-rate-limit"`
    TrackTogetherWithSyn int `json:"track-together-with-syn"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZeroWinCfg DdosZoneTemplateTcpZeroWinCfg `json:"zero-win-cfg"`

	} `json:"tcp"`
}


type DdosZoneTemplateTcpAckAuthentication struct {
    AckAuthTimeout int `json:"ack-auth-timeout"`
    AckAuthMinDelay int `json:"ack-auth-min-delay"`
    AckAuthOnly int `json:"ack-auth-only"`
    AckAuthRto int `json:"ack-auth-rto"`
    AckAuthPassActionListName string `json:"ack-auth-pass-action-list-name"`
    AckAuthPassAction string `json:"ack-auth-pass-action"`
    AckAuthFailActionListName string `json:"ack-auth-fail-action-list-name"`
    AckAuthFailAction string `json:"ack-auth-fail-action"`
}


type DdosZoneTemplateTcpDst struct {
    RateLimit DdosZoneTemplateTcpDstRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateTcpDstRateLimit struct {
    SynRateLimit DdosZoneTemplateTcpDstRateLimitSynRateLimit `json:"syn-rate-limit"`
}


type DdosZoneTemplateTcpDstRateLimitSynRateLimit struct {
    DstSynRateLimit int `json:"dst-syn-rate-limit"`
    DstSynRateAction string `json:"dst-syn-rate-action" dval:"drop"`
}


type DdosZoneTemplateTcpFilterList struct {
    TcpFilterName string `json:"tcp-filter-name"`
    TcpFilterSeq int `json:"tcp-filter-seq"`
    TcpFilterRegex string `json:"tcp-filter-regex"`
    TcpFilterInverseMatch int `json:"tcp-filter-inverse-match"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    TcpFilterActionListName string `json:"tcp-filter-action-list-name"`
    TcpFilterAction string `json:"tcp-filter-action" dval:"drop"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateTcpKnownRespSrcPortCfg struct {
    KnownRespSrcPort int `json:"known-resp-src-port"`
    KnownRespSrcPortActionListName string `json:"known-resp-src-port-action-list-name"`
    KnownRespSrcPortAction string `json:"known-resp-src-port-action"`
    ExcludeSrcRespPort int `json:"exclude-src-resp-port"`
}


type DdosZoneTemplateTcpMaxRexmitSynPerFlowCfg struct {
    MaxRexmitSynPerFlow int `json:"max-rexmit-syn-per-flow"`
    MaxRexmitSynPerFlowActionListName string `json:"max-rexmit-syn-per-flow-action-list-name"`
    MaxRexmitSynPerFlowAction string `json:"max-rexmit-syn-per-flow-action" dval:"drop"`
}


type DdosZoneTemplateTcpOutOfSeqCfg struct {
    OutOfSeq int `json:"out-of-seq"`
    OutOfSeqActionListName string `json:"out-of-seq-action-list-name"`
    OutOfSeqAction string `json:"out-of-seq-action" dval:"drop"`
}


type DdosZoneTemplateTcpPerConnOutOfSeqRateCfg struct {
    PerConnOutOfSeqRateLimit int `json:"per-conn-out-of-seq-rate-limit"`
    PerConnOutOfSeqRateActionListName string `json:"per-conn-out-of-seq-rate-action-list-name"`
    PerConnOutOfSeqRateAction string `json:"per-conn-out-of-seq-rate-action" dval:"drop"`
}


type DdosZoneTemplateTcpPerConnPktRateCfg struct {
    PerConnPktRateLimit int `json:"per-conn-pkt-rate-limit"`
    PerConnPktRateActionListName string `json:"per-conn-pkt-rate-action-list-name"`
    PerConnPktRateAction string `json:"per-conn-pkt-rate-action" dval:"drop"`
}


type DdosZoneTemplateTcpPerConnRetransmitRateCfg struct {
    PerConnRetransmitRateLimit int `json:"per-conn-retransmit-rate-limit"`
    PerConnRetransmitRateActionListName string `json:"per-conn-retransmit-rate-action-list-name"`
    PerConnRetransmitRateAction string `json:"per-conn-retransmit-rate-action" dval:"drop"`
}


type DdosZoneTemplateTcpPerConnZeroWinRateCfg struct {
    PerConnZeroWinRateLimit int `json:"per-conn-zero-win-rate-limit"`
    PerConnZeroWinRateActionListName string `json:"per-conn-zero-win-rate-action-list-name"`
    PerConnZeroWinRateAction string `json:"per-conn-zero-win-rate-action" dval:"drop"`
}


type DdosZoneTemplateTcpProgressionTracking320 struct {
    ProgressionTrackingEnabled string `json:"progression-tracking-enabled"`
    RequestResponseModel string `json:"request-response-model" dval:"enable"`
    Violation int `json:"violation"`
    IgnoreTlsHandshake int `json:"ignore-TLS-handshake"`
    ResponseLengthMax int `json:"response-length-max"`
    ResponseLengthMin int `json:"response-length-min"`
    RequestLengthMin int `json:"request-length-min"`
    RequestLengthMax int `json:"request-length-max"`
    ResponseRequestMinRatio int `json:"response-request-min-ratio"`
    ResponseRequestMaxRatio int `json:"response-request-max-ratio"`
    FirstRequestMaxTime int `json:"first-request-max-time"`
    RequestToResponseMaxTime int `json:"request-to-response-max-time"`
    ResponseToRequestMaxTime int `json:"response-to-request-max-time"`
    ProfilingRequestResponseModel int `json:"profiling-request-response-model"`
    ProfilingConnectionLifeModel int `json:"profiling-connection-life-model"`
    ProfilingTimeWindowModel int `json:"profiling-time-window-model"`
    ProgressionTrackingActionListName string `json:"progression-tracking-action-list-name"`
    ProgressionTrackingAction string `json:"progression-tracking-action" dval:"drop"`
    Uuid string `json:"uuid"`
    ConnectionTracking DdosZoneTemplateTcpProgressionTrackingConnectionTracking321 `json:"connection-tracking"`
    TimeWindowTracking DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking322 `json:"time-window-tracking"`
}


type DdosZoneTemplateTcpProgressionTrackingConnectionTracking321 struct {
    ProgressionTrackingConnEnabled string `json:"progression-tracking-conn-enabled"`
    ConnSentMax int `json:"conn-sent-max"`
    ConnSentMin int `json:"conn-sent-min"`
    ConnRcvdMax int `json:"conn-rcvd-max"`
    ConnRcvdMin int `json:"conn-rcvd-min"`
    ConnRcvdSentRatioMin int `json:"conn-rcvd-sent-ratio-min"`
    ConnRcvdSentRatioMax int `json:"conn-rcvd-sent-ratio-max"`
    ConnDurationMax int `json:"conn-duration-max"`
    ConnDurationMin int `json:"conn-duration-min"`
    ConnViolation int `json:"conn-violation"`
    ProgressionTrackingConnActionListName string `json:"progression-tracking-conn-action-list-name"`
    ProgressionTrackingConnAction string `json:"progression-tracking-conn-action" dval:"drop"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking322 struct {
    ProgressionTrackingWinEnabled string `json:"progression-tracking-win-enabled"`
    WindowSentMax int `json:"window-sent-max"`
    WindowSentMin int `json:"window-sent-min"`
    WindowRcvdMax int `json:"window-rcvd-max"`
    WindowRcvdMin int `json:"window-rcvd-min"`
    WindowRcvdSentRatioMin int `json:"window-rcvd-sent-ratio-min"`
    WindowRcvdSentRatioMax int `json:"window-rcvd-sent-ratio-max"`
    WindowViolation int `json:"window-violation"`
    ProgressionTrackingWindowsActionListName string `json:"progression-tracking-windows-action-list-name"`
    ProgressionTrackingWindowsAction string `json:"progression-tracking-windows-action" dval:"drop"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateTcpRetransmitCfg struct {
    Retransmit int `json:"retransmit"`
    RetransmitActionListName string `json:"retransmit-action-list-name"`
    RetransmitAction string `json:"retransmit-action" dval:"drop"`
}


type DdosZoneTemplateTcpSrc struct {
    RateLimit DdosZoneTemplateTcpSrcRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateTcpSrcRateLimit struct {
    SynRateLimit DdosZoneTemplateTcpSrcRateLimitSynRateLimit `json:"syn-rate-limit"`
}


type DdosZoneTemplateTcpSrcRateLimitSynRateLimit struct {
    SrcSynRateLimit int `json:"src-syn-rate-limit"`
    SrcSynRateActionListName string `json:"src-syn-rate-action-list-name"`
    SrcSynRateAction string `json:"src-syn-rate-action" dval:"drop"`
}


type DdosZoneTemplateTcpSynAuthentication struct {
    SynAuthType string `json:"syn-auth-type"`
    SynAuthTimeout int `json:"syn-auth-timeout"`
    SynAuthMinDelay int `json:"syn-auth-min-delay"`
    SynAuthRto int `json:"syn-auth-rto"`
    SynAuthPassActionListName string `json:"syn-auth-pass-action-list-name"`
    SynAuthPassAction string `json:"syn-auth-pass-action"`
    SynAuthFailActionListName string `json:"syn-auth-fail-action-list-name"`
    SynAuthFailAction string `json:"syn-auth-fail-action"`
    AllowRa int `json:"allow-ra"`
}


type DdosZoneTemplateTcpZeroWinCfg struct {
    ZeroWin int `json:"zero-win"`
    ZeroWinActionListName string `json:"zero-win-action-list-name"`
    ZeroWinAction string `json:"zero-win-action" dval:"drop"`
}

func (p *DdosZoneTemplateTcp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosZoneTemplateTcp) getPath() string{
    return "ddos/zone-template/tcp"
}

func (p *DdosZoneTemplateTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcp::Post")
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

func (p *DdosZoneTemplateTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcp::Get")
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
func (p *DdosZoneTemplateTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcp::Put")
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

func (p *DdosZoneTemplateTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
