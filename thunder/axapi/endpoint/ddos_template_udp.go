

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateUdp struct {
	Inst struct {

    Age int `json:"age"`
    DropKnownRespSrcPortCfg DdosTemplateUdpDropKnownRespSrcPortCfg `json:"drop-known-resp-src-port-cfg"`
    DropNtpMonlist int `json:"drop-ntp-monlist"`
    FilterList []DdosTemplateUdpFilterList `json:"filter-list"`
    MaxPayloadSize int `json:"max-payload-size"`
    MinPayloadSize int `json:"min-payload-size"`
    Name string `json:"name"`
    PerConnPktRateLimit int `json:"per-conn-pkt-rate-limit"`
    PerConnRateInterval string `json:"per-conn-rate-interval" dval:"1sec"`
    PreviousSaltTimeout int `json:"previous-salt-timeout" dval:"1"`
    PublicIpv4Addr string `json:"public-ipv4-addr"`
    PublicIpv6Addr string `json:"public-ipv6-addr"`
    SpoofDetectCfg DdosTemplateUdpSpoofDetectCfg `json:"spoof-detect-cfg"`
    TokenAuthentication int `json:"token-authentication"`
    TokenAuthenticationFormula string `json:"token-authentication-formula"`
    TokenAuthenticationHwAssistDisable int `json:"token-authentication-hw-assist-disable"`
    TokenAuthenticationPublicAddress int `json:"token-authentication-public-address"`
    TokenAuthenticationSaltPrefix int `json:"token-authentication-salt-prefix"`
    TokenAuthenticationSaltPrefixCurr int `json:"token-authentication-salt-prefix-curr"`
    TokenAuthenticationSaltPrefixPrev int `json:"token-authentication-salt-prefix-prev"`
    TunnelEncap DdosTemplateUdpTunnelEncap `json:"tunnel-encap"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"udp"`
}


type DdosTemplateUdpDropKnownRespSrcPortCfg struct {
    DropKnownRespSrcPort int `json:"drop-known-resp-src-port"`
    ExcludeSrcRespPort int `json:"exclude-src-resp-port"`
}


type DdosTemplateUdpFilterList struct {
    UdpFilterSeq int `json:"udp-filter-seq"`
    UdpFilterRegex string `json:"udp-filter-regex"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    UdpFilterUnmatched int `json:"udp-filter-unmatched"`
    UdpFilterAction string `json:"udp-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosTemplateUdpSpoofDetectCfg struct {
    SpoofDetect int `json:"spoof-detect"`
    MinRetryGapInterval string `json:"min-retry-gap-interval" dval:"1sec"`
    SpoofDetectRetryTimeoutValOnly int `json:"spoof-detect-retry-timeout-val-only" dval:"5"`
    MinRetryGap int `json:"min-retry-gap"`
    SpoofDetectRetryTimeout int `json:"spoof-detect-retry-timeout" dval:"5"`
}


type DdosTemplateUdpTunnelEncap struct {
    IpEncap int `json:"ip-encap"`
    Always DdosTemplateUdpTunnelEncapAlways `json:"always"`
    GreEncap int `json:"gre-encap"`
    GreAlways DdosTemplateUdpTunnelEncapGreAlways `json:"gre-always"`
}


type DdosTemplateUdpTunnelEncapAlways struct {
    Ipv4Addr string `json:"ipv4-addr"`
    PreserveSrcIpv4 int `json:"preserve-src-ipv4"`
    Ipv6Addr string `json:"ipv6-addr"`
    PreserveSrcIpv6 int `json:"preserve-src-ipv6"`
}


type DdosTemplateUdpTunnelEncapGreAlways struct {
    GreIpv4 string `json:"gre-ipv4"`
    KeyIpv4 string `json:"key-ipv4"`
    PreserveSrcIpv4Gre int `json:"preserve-src-ipv4-gre"`
    GreIpv6 string `json:"gre-ipv6"`
    KeyIpv6 string `json:"key-ipv6"`
    PreserveSrcIpv6Gre int `json:"preserve-src-ipv6-gre"`
}

func (p *DdosTemplateUdp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosTemplateUdp) getPath() string{
    return "ddos/template/udp"
}

func (p *DdosTemplateUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateUdp::Post")
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

func (p *DdosTemplateUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateUdp::Get")
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
func (p *DdosTemplateUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateUdp::Put")
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

func (p *DdosTemplateUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
