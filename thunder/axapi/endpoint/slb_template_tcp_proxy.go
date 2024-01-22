

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateTcpProxy struct {
	Inst struct {

    AckAggressiveness string `json:"ack-aggressiveness" dval:"low"`
    AliveIfActive int `json:"alive-if-active"`
    BackendWscale int `json:"backend-wscale"`
    DelSessionOnServerDown int `json:"del-session-on-server-down"`
    Disable int `json:"disable"`
    DisableAbc int `json:"disable-abc"`
    DisableSack int `json:"disable-sack"`
    DisableTcpTimestamps int `json:"disable-tcp-timestamps"`
    DisableWindowScale int `json:"disable-window-scale"`
    Down int `json:"down"`
    DynamicBufferAllocation int `json:"dynamic-buffer-allocation"`
    EarlyRetransmit int `json:"early-retransmit"`
    FinTimeout int `json:"fin-timeout"`
    ForceDeleteTimeout int `json:"force-delete-timeout"`
    ForceDeleteTimeout100ms int `json:"force-delete-timeout-100ms"`
    HalfCloseIdleTimeout int `json:"half-close-idle-timeout"`
    HalfOpenIdleTimeout int `json:"half-open-idle-timeout"`
    IdleTimeout int `json:"idle-timeout" dval:"600"`
    InitCwnd int `json:"init-cwnd" dval:"10"`
    InitialWindowSize int `json:"initial-window-size"`
    InsertClientIp int `json:"insert-client-ip"`
    InvalidRateLimit int `json:"invalid-rate-limit" dval:"500"`
    KeepaliveInterval int `json:"keepalive-interval"`
    KeepaliveProbes int `json:"keepalive-probes"`
    LimitedSlowstart int `json:"limited-slowstart"`
    Maxburst int `json:"maxburst" dval:"25"`
    MinRto int `json:"min-rto" dval:"200"`
    Mss int `json:"mss" dval:"1460"`
    Nagle int `json:"nagle"`
    NakedAckOnHandshake int `json:"naked-ack-on-handshake"`
    Name string `json:"name"`
    ProxyHeader SlbTemplateTcpProxyProxyHeader `json:"proxy-header"`
    PshFlagOptimization int `json:"psh-flag-optimization"`
    Qos int `json:"qos"`
    ReassemblyLimit int `json:"reassembly-limit" dval:"25"`
    ReassemblyTimeout int `json:"reassembly-timeout" dval:"30"`
    ReceiveBuffer int `json:"receive-buffer" dval:"200000"`
    Reno int `json:"reno"`
    ResetFwd int `json:"reset-fwd"`
    ResetRev int `json:"reset-rev"`
    RetransmitRetries int `json:"retransmit-retries" dval:"5"`
    ServerDownAction string `json:"server-down-action"`
    SynRetries int `json:"syn-retries" dval:"5"`
    Timewait int `json:"timewait" dval:"5"`
    TransmitBuffer int `json:"transmit-buffer" dval:"200000"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"tcp-proxy"`
}


type SlbTemplateTcpProxyProxyHeader struct {
    ProxyHeaderAction string `json:"proxy-header-action"`
    Version string `json:"version"`
}

func (p *SlbTemplateTcpProxy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateTcpProxy) getPath() string{
    return "slb/template/tcp-proxy"
}

func (p *SlbTemplateTcpProxy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcpProxy::Post")
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

func (p *SlbTemplateTcpProxy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcpProxy::Get")
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
func (p *SlbTemplateTcpProxy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcpProxy::Put")
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

func (p *SlbTemplateTcpProxy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateTcpProxy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
