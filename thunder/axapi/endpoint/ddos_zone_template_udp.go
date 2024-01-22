

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateUdp struct {
	Inst struct {

    Age int `json:"age" dval:"2"`
    FilterList []DdosZoneTemplateUdpFilterList `json:"filter-list"`
    FilterMatchType string `json:"filter-match-type" dval:"default"`
    KnownRespSrcPortCfg DdosZoneTemplateUdpKnownRespSrcPortCfg `json:"known-resp-src-port-cfg"`
    MaxPayloadSizeCfg DdosZoneTemplateUdpMaxPayloadSizeCfg `json:"max-payload-size-cfg"`
    MinPayloadSizeCfg DdosZoneTemplateUdpMinPayloadSizeCfg `json:"min-payload-size-cfg"`
    Name string `json:"name"`
    NtpMonlistCfg DdosZoneTemplateUdpNtpMonlistCfg `json:"ntp-monlist-cfg"`
    PerConnPktRateCfg DdosZoneTemplateUdpPerConnPktRateCfg `json:"per-conn-pkt-rate-cfg"`
    PerConnRateInterval string `json:"per-conn-rate-interval" dval:"1sec"`
    PreviousSaltTimeout int `json:"previous-salt-timeout" dval:"1"`
    PublicIpv4Addr string `json:"public-ipv4-addr"`
    PublicIpv6Addr string `json:"public-ipv6-addr"`
    SpoofDetectFailAction string `json:"spoof-detect-fail-action"`
    SpoofDetectFailActionListName string `json:"spoof-detect-fail-action-list-name"`
    SpoofDetectMinDelay int `json:"spoof-detect-min-delay"`
    SpoofDetectMinDelayInterval string `json:"spoof-detect-min-delay-interval" dval:"1sec"`
    SpoofDetectPassAction string `json:"spoof-detect-pass-action"`
    SpoofDetectPassActionListName string `json:"spoof-detect-pass-action-list-name"`
    SpoofDetectRetryTimeout int `json:"spoof-detect-retry-timeout"`
    TokenAuthentication int `json:"token-authentication"`
    TokenAuthenticationFormula string `json:"token-authentication-formula"`
    TokenAuthenticationHwAssistDisable int `json:"token-authentication-hw-assist-disable"`
    TokenAuthenticationPublicAddress int `json:"token-authentication-public-address"`
    TokenAuthenticationSaltPrefix int `json:"token-authentication-salt-prefix"`
    TokenAuthenticationSaltPrefixCurr int `json:"token-authentication-salt-prefix-curr"`
    TokenAuthenticationSaltPrefixPrev int `json:"token-authentication-salt-prefix-prev"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"udp"`
}


type DdosZoneTemplateUdpFilterList struct {
    UdpFilterName string `json:"udp-filter-name"`
    UdpFilterSeq int `json:"udp-filter-seq"`
    UdpFilterRegex string `json:"udp-filter-regex"`
    UdpFilterInverseMatch int `json:"udp-filter-inverse-match"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    UdpFilterActionListName string `json:"udp-filter-action-list-name"`
    UdpFilterAction string `json:"udp-filter-action" dval:"drop"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateUdpKnownRespSrcPortCfg struct {
    KnownRespSrcPort int `json:"known-resp-src-port"`
    KnownRespSrcPortActionListName string `json:"known-resp-src-port-action-list-name"`
    KnownRespSrcPortAction string `json:"known-resp-src-port-action"`
    ExcludeSrcRespPort int `json:"exclude-src-resp-port"`
}


type DdosZoneTemplateUdpMaxPayloadSizeCfg struct {
    MaxPayloadSize int `json:"max-payload-size"`
    MaxPayloadSizeActionListName string `json:"max-payload-size-action-list-name"`
    MaxPayloadSizeAction string `json:"max-payload-size-action" dval:"drop"`
}


type DdosZoneTemplateUdpMinPayloadSizeCfg struct {
    MinPayloadSize int `json:"min-payload-size"`
    MinPayloadSizeActionListName string `json:"min-payload-size-action-list-name"`
    MinPayloadSizeAction string `json:"min-payload-size-action" dval:"drop"`
}


type DdosZoneTemplateUdpNtpMonlistCfg struct {
    NtpMonlist int `json:"ntp-monlist"`
    NtpMonlistActionListName string `json:"ntp-monlist-action-list-name"`
    NtpMonlistAction string `json:"ntp-monlist-action" dval:"drop"`
}


type DdosZoneTemplateUdpPerConnPktRateCfg struct {
    PerConnPktRateLimit int `json:"per-conn-pkt-rate-limit"`
    PerConnPktRateActionListName string `json:"per-conn-pkt-rate-action-list-name"`
    PerConnPktRateAction string `json:"per-conn-pkt-rate-action" dval:"drop"`
}

func (p *DdosZoneTemplateUdp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosZoneTemplateUdp) getPath() string{
    return "ddos/zone-template/udp"
}

func (p *DdosZoneTemplateUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateUdp::Post")
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

func (p *DdosZoneTemplateUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateUdp::Get")
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
func (p *DdosZoneTemplateUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateUdp::Put")
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

func (p *DdosZoneTemplateUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
