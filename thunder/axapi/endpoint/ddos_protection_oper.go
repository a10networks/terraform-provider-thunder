

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosProtectionOper struct {
    
    Ipv6SrcHashMaskBits DdosProtectionOperIpv6SrcHashMaskBits `json:"ipv6-src-hash-mask-bits"`
    Oper DdosProtectionOperOper `json:"oper"`

}
type DataDdosProtectionOper struct {
    DtDdosProtectionOper DdosProtectionOper `json:"protection"`
}


type DdosProtectionOperIpv6SrcHashMaskBits struct {
    Oper DdosProtectionOperIpv6SrcHashMaskBitsOper `json:"oper"`
}


type DdosProtectionOperIpv6SrcHashMaskBitsOper struct {
    Offsets []DdosProtectionOperIpv6SrcHashMaskBitsOperOffsets `json:"offsets"`
}


type DdosProtectionOperIpv6SrcHashMaskBitsOperOffsets struct {
    MaskBitOffset1 int `json:"mask-bit-offset-1"`
    MaskBitOffset2 int `json:"mask-bit-offset-2"`
    MaskBitOffset3 int `json:"mask-bit-offset-3"`
    MaskBitOffset4 int `json:"mask-bit-offset-4"`
    MaskBitOffset5 int `json:"mask-bit-offset-5"`
}


type DdosProtectionOperOper struct {
    DdosProtection string `json:"ddos-protection"`
    RateInterval string `json:"rate-interval"`
    Mode string `json:"mode"`
    UseRoute string `json:"use-route"`
    TapInterfaces string `json:"tap-interfaces"`
    DstAutoLearningIpv4 string `json:"dst-auto-learning-ipv4"`
    DstAutoLearningIpv6 string `json:"dst-auto-learning-ipv6"`
    SrcAutoLearningIpv4 string `json:"src-auto-learning-ipv4"`
    SrcAutoLearningIpv6 string `json:"src-auto-learning-ipv6"`
    SrcDelayLearning string `json:"src-delay-learning"`
    OneArmMode string `json:"one-arm-mode"`
    HwSynCookie string `json:"hw-syn-cookie"`
    Sync string `json:"sync"`
    SyncAutoWl string `json:"sync-auto-wl"`
    Bgp string `json:"bgp"`
    BgpAutoWl string `json:"bgp-auto-wl"`
    Vrrp string `json:"vrrp"`
    VrrpAutoWl string `json:"vrrp-auto-wl"`
    MplsPktInspect string `json:"mpls-pkt-inspect"`
    Detection string `json:"detection"`
    DdetMode string `json:"ddet-mode"`
    DdetCpus int `json:"ddet-cpus"`
    DstDynamicOverflowIpv4 string `json:"dst-dynamic-overflow-ipv4"`
    DstDynamicOverflowIpv6 string `json:"dst-dynamic-overflow-ipv6"`
    SrcDynamicOverflowIpv4 string `json:"src-dynamic-overflow-ipv4"`
    SrcDynamicOverflowIpv6 string `json:"src-dynamic-overflow-ipv6"`
    IpAnoSecL3 string `json:"ip-ano-sec-l3"`
    IpAnoSecL4Tcp string `json:"ip-ano-sec-l4-tcp"`
    IpAnoSecL4Udp string `json:"ip-ano-sec-l4-udp"`
    IpAnoDefL3 string `json:"ip-ano-def-l3"`
    IpAnoDefL4 string `json:"ip-ano-def-l4"`
    DnsCacheMode string `json:"dns-cache-mode"`
    WarmUp string `json:"warm-up"`
    DnsZoneTransferDedicatedCpus int `json:"dns-zone-transfer-dedicated-cpus"`
    SrcDstEntryLimit string `json:"src-dst-entry-limit"`
    SrcZonePortEntryLimit string `json:"src-zone-port-entry-limit"`
    InterbladeSyncAccuracy string `json:"interblade-sync-accuracy"`
    PatternRecognition string `json:"pattern-recognition"`
    PatternRecognitionCpus int `json:"pattern-recognition-cpus"`
    PatternRecognitionHardwareFilter string `json:"pattern-recognition-hardware-filter"`
    DetectionWindowSize int `json:"detection-window-size"`
    DisallowRstAckInSynAuth string `json:"disallow-rst-ack-in-syn-auth"`
    NonZeroWinSizeSyncookie string `json:"non-zero-win-size-syncookie"`
    HwBlocking string `json:"hw-blocking"`
    HwBlockingThreshold int `json:"hw-blocking-threshold"`
    InterfaceHttpHealthCheck string `json:"interface-http-health-check"`
}

func (p *DdosProtectionOper) GetId() string{
    return "1"
}

func (p *DdosProtectionOper) getPath() string{
    return "ddos/protection/oper"
}

func (p *DdosProtectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosProtectionOper,error) {
logger.Println("DdosProtectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosProtectionOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
