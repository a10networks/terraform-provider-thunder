

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateTcp struct {
	Inst struct {

    AckAuthenticationSynackReset int `json:"ack-authentication-synack-reset"`
    ActionCfg DdosTemplateTcpActionCfg `json:"action-cfg"`
    ActionOnAckRtoRetryCount int `json:"action-on-ack-rto-retry-count"`
    ActionOnSynRtoRetryCount int `json:"action-on-syn-rto-retry-count"`
    ActionSynCfg DdosTemplateTcpActionSynCfg `json:"action-syn-cfg"`
    Age int `json:"age"`
    AllowRa int `json:"allow-ra"`
    AllowSynOtherflags int `json:"allow-syn-otherflags"`
    AllowSynackSkipAuthentications int `json:"allow-synack-skip-authentications"`
    AllowTcpTfo int `json:"allow-tcp-tfo"`
    BlackListOutOfSeq int `json:"black-list-out-of-seq"`
    BlackListRetransmit int `json:"black-list-retransmit"`
    BlackListZeroWin int `json:"black-list-zero-win"`
    ConnRateLimitOnSynOnly int `json:"conn-rate-limit-on-syn-only"`
    CreateConnOnSynOnly int `json:"create-conn-on-syn-only"`
    DropKnownRespSrcPortCfg DdosTemplateTcpDropKnownRespSrcPortCfg `json:"drop-known-resp-src-port-cfg"`
    Dst DdosTemplateTcpDst `json:"dst"`
    FilterList []DdosTemplateTcpFilterList `json:"filter-list"`
    Name string `json:"name"`
    PerConnOutOfSeqRateAction string `json:"per-conn-out-of-seq-rate-action" dval:"drop"`
    PerConnOutOfSeqRateLimit int `json:"per-conn-out-of-seq-rate-limit"`
    PerConnPktRateAction string `json:"per-conn-pkt-rate-action" dval:"drop"`
    PerConnPktRateLimit int `json:"per-conn-pkt-rate-limit"`
    PerConnRateInterval string `json:"per-conn-rate-interval" dval:"1sec"`
    PerConnRetransmitRateAction string `json:"per-conn-retransmit-rate-action" dval:"drop"`
    PerConnRetransmitRateLimit int `json:"per-conn-retransmit-rate-limit"`
    PerConnZeroWinRateAction string `json:"per-conn-zero-win-rate-action" dval:"drop"`
    PerConnZeroWinRateLimit int `json:"per-conn-zero-win-rate-limit"`
    ProgressionTracking DdosTemplateTcpProgressionTracking302 `json:"progression-tracking"`
    Src DdosTemplateTcpSrc `json:"src"`
    SynAuth string `json:"syn-auth" dval:"send-rst"`
    SynCookie int `json:"syn-cookie"`
    SynackRateLimit int `json:"synack-rate-limit"`
    TrackTogetherWithSyn int `json:"track-together-with-syn"`
    TunnelEncap DdosTemplateTcpTunnelEncap `json:"tunnel-encap"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type DdosTemplateTcpActionCfg struct {
    ActionOnAck int `json:"action-on-ack"`
    Reset int `json:"reset"`
    Timeout int `json:"timeout"`
    MinRetryGap int `json:"min-retry-gap"`
    AuthenticateOnly int `json:"authenticate-only"`
    RtoAuthentication int `json:"rto-authentication"`
}


type DdosTemplateTcpActionSynCfg struct {
    ActionOnSyn int `json:"action-on-syn"`
    ActionOnSynReset int `json:"action-on-syn-reset"`
    ActionOnSynTimeout int `json:"action-on-syn-timeout"`
    ActionOnSynGap int `json:"action-on-syn-gap"`
    ActionOnSynRto int `json:"action-on-syn-rto"`
}


type DdosTemplateTcpDropKnownRespSrcPortCfg struct {
    DropKnownRespSrcPort int `json:"drop-known-resp-src-port"`
    ExcludeSrcRespPort int `json:"exclude-src-resp-port"`
}


type DdosTemplateTcpDst struct {
    RateLimit DdosTemplateTcpDstRateLimit `json:"rate-limit"`
}


type DdosTemplateTcpDstRateLimit struct {
    SynRateLimit DdosTemplateTcpDstRateLimitSynRateLimit `json:"syn-rate-limit"`
}


type DdosTemplateTcpDstRateLimitSynRateLimit struct {
    DstSynRateLimit int `json:"dst-syn-rate-limit"`
    DstSynRateAction string `json:"dst-syn-rate-action" dval:"drop"`
}


type DdosTemplateTcpFilterList struct {
    TcpFilterSeq int `json:"tcp-filter-seq"`
    TcpFilterRegex string `json:"tcp-filter-regex"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    TcpFilterUnmatched int `json:"tcp-filter-unmatched"`
    TcpFilterAction string `json:"tcp-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosTemplateTcpProgressionTracking302 struct {
    ProgressionTrackingEnabled string `json:"progression-tracking-enabled"`
    RequestResponseModel string `json:"request-response-model" dval:"enable"`
    Violation int `json:"violation"`
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
    ConnectionTracking DdosTemplateTcpProgressionTrackingConnectionTracking303 `json:"connection-tracking"`
    TimeWindowTracking DdosTemplateTcpProgressionTrackingTimeWindowTracking304 `json:"time-window-tracking"`
}


type DdosTemplateTcpProgressionTrackingConnectionTracking303 struct {
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


type DdosTemplateTcpProgressionTrackingTimeWindowTracking304 struct {
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


type DdosTemplateTcpSrc struct {
    RateLimit DdosTemplateTcpSrcRateLimit `json:"rate-limit"`
}


type DdosTemplateTcpSrcRateLimit struct {
    SynRateLimit DdosTemplateTcpSrcRateLimitSynRateLimit `json:"syn-rate-limit"`
}


type DdosTemplateTcpSrcRateLimitSynRateLimit struct {
    SrcSynRateLimit int `json:"src-syn-rate-limit"`
    SrcSynRateAction string `json:"src-syn-rate-action" dval:"drop"`
}


type DdosTemplateTcpTunnelEncap struct {
    IpCfg DdosTemplateTcpTunnelEncapIpCfg `json:"ip-cfg"`
    GreCfg DdosTemplateTcpTunnelEncapGreCfg `json:"gre-cfg"`
}


type DdosTemplateTcpTunnelEncapIpCfg struct {
    IpEncap int `json:"ip-encap"`
    Always DdosTemplateTcpTunnelEncapIpCfgAlways `json:"always"`
}


type DdosTemplateTcpTunnelEncapIpCfgAlways struct {
    Ipv4Addr string `json:"ipv4-addr"`
    PreserveSrcIpv4 int `json:"preserve-src-ipv4"`
    Ipv6Addr string `json:"ipv6-addr"`
    PreserveSrcIpv6 int `json:"preserve-src-ipv6"`
}


type DdosTemplateTcpTunnelEncapGreCfg struct {
    GreEncap int `json:"gre-encap"`
    GreAlways DdosTemplateTcpTunnelEncapGreCfgGreAlways `json:"gre-always"`
}


type DdosTemplateTcpTunnelEncapGreCfgGreAlways struct {
    GreIpv4 string `json:"gre-ipv4"`
    KeyIpv4 string `json:"key-ipv4"`
    PreserveSrcIpv4Gre int `json:"preserve-src-ipv4-gre"`
    GreIpv6 string `json:"gre-ipv6"`
    KeyIpv6 string `json:"key-ipv6"`
    PreserveSrcIpv6Gre int `json:"preserve-src-ipv6-gre"`
}

func (p *DdosTemplateTcp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosTemplateTcp) getPath() string{
    return "ddos/template/tcp"
}

func (p *DdosTemplateTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcp::Post")
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

func (p *DdosTemplateTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcp::Get")
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
func (p *DdosTemplateTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcp::Put")
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

func (p *DdosTemplateTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
