

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosProtection struct {
	Inst struct {

    BlacklistReasonTracking int `json:"blacklist-reason-tracking"`
    CloseSessForUnauthSrcWithoutRst int `json:"close-sess-for-unauth-src-without-rst"`
    DisableAdvancedCoreAnalysis int `json:"disable-advanced-core-analysis"`
    DisableDelayDynamicSrcLearning int `json:"disable-delay-dynamic-src-learning"`
    DisableOnReboot int `json:"disable-on-reboot"`
    DisallowRstAckInSynAuth int `json:"disallow-rst-ack-in-syn-auth"`
    EnableNow int `json:"enable-now"`
    FastAging DdosProtectionFastAging `json:"fast-aging"`
    FastPathDisable int `json:"fast-path-disable"`
    ForceRoutingOnTransp int `json:"force-routing-on-transp"`
    ForceTrafficToSameBladeDisable int `json:"force-traffic-to-same-blade-disable"`
    HwBlockingEnable int `json:"hw-blocking-enable"`
    HwBlockingThresholdLimit int `json:"hw-blocking-threshold-limit" dval:"10000"`
    Ipv6SrcHashMaskBits DdosProtectionIpv6SrcHashMaskBits291 `json:"ipv6-src-hash-mask-bits"`
    Mpls int `json:"mpls"`
    MultiPuZoneDistribution DdosProtectionMultiPuZoneDistribution292 `json:"multi-pu-zone-distribution"`
    NonZeroWinSizeSyncookie int `json:"non-zero-win-size-syncookie"`
    ProgressionTracking string `json:"progression-tracking" dval:"enable"`
    RateInterval string `json:"rate-interval" dval:"100ms"`
    RexmitSynLog int `json:"rexmit-syn-log"`
    SrcDstEntryLimit string `json:"src-dst-entry-limit" dval:"16M"`
    SrcIpHashBit int `json:"src-ip-hash-bit" dval:"2"`
    SrcIpv6HashBit int `json:"src-ipv6-hash-bit" dval:"2"`
    SrcZonePortEntryLimit string `json:"src-zone-port-entry-limit" dval:"16M"`
    Toggle string `json:"toggle" dval:"disable"`
    UseRoute int `json:"use-route"`
    Uuid string `json:"uuid"`

	} `json:"protection"`
}


type DdosProtectionFastAging struct {
    HalfOpenConnRatio int `json:"half-open-conn-ratio" dval:"25"`
    HalfOpenConnThreshold int `json:"half-open-conn-threshold" dval:"1"`
}


type DdosProtectionIpv6SrcHashMaskBits291 struct {
    MaskBitOffset1 int `json:"mask-bit-offset-1"`
    MaskBitOffset2 int `json:"mask-bit-offset-2"`
    MaskBitOffset3 int `json:"mask-bit-offset-3"`
    MaskBitOffset4 int `json:"mask-bit-offset-4"`
    MaskBitOffset5 int `json:"mask-bit-offset-5"`
    Uuid string `json:"uuid"`
}


type DdosProtectionMultiPuZoneDistribution292 struct {
    DistributionMethod string `json:"distribution-method" dval:"traffic-rate"`
    CpuThresholdPerEntry int `json:"cpu-threshold-per-entry" dval:"60"`
    CpuThresholdPerPu int `json:"cpu-threshold-per-pu" dval:"80"`
    RatePktThreshold int `json:"rate-pkt-threshold" dval:"55000000"`
    RateKbitThreshold int `json:"rate-kbit-threshold" dval:"150000000"`
    Uuid string `json:"uuid"`
}

func (p *DdosProtection) GetId() string{
    return "1"
}

func (p *DdosProtection) getPath() string{
    return "ddos/protection"
}

func (p *DdosProtection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtection::Post")
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

func (p *DdosProtection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtection::Get")
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
func (p *DdosProtection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtection::Put")
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

func (p *DdosProtection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
