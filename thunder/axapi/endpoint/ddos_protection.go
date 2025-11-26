package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
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

		Ipv6SrcHashMaskBits DdosProtectionIpv6SrcHashMaskBits330 `json:"ipv6-src-hash-mask-bits"`

		Mpls int `json:"mpls"`

		MultiPuZoneDistribution DdosProtectionMultiPuZoneDistribution331 `json:"multi-pu-zone-distribution"`

		NonZeroWinSizeSyncookie int `json:"non-zero-win-size-syncookie"`

		PerServiceSzpEntryLimit DdosProtectionPerServiceSzpEntryLimit332 `json:"per-service-szp-entry-limit"`

		PktRateLimitOnReassemble string `json:"pkt-rate-limit-on-reassemble" dval:"disable"`

		ProgressionTracking string `json:"progression-tracking" dval:"enable"`

		RateInterval string `json:"rate-interval" dval:"100ms"`

		RateLimitSyncInterval int `json:"rate-limit-sync-interval" dval:"3"`

		RexmitSynLog int `json:"rexmit-syn-log"`

		SrcDstEntryLimit string `json:"src-dst-entry-limit" dval:"16M"`

		SrcHashFunction string `json:"src-hash-function" dval:"v1"`

		SrcIpHashBit int `json:"src-ip-hash-bit" dval:"2"`

		SrcIpv6HashBit int `json:"src-ipv6-hash-bit" dval:"2"`

		SrcZonePortEntryLimit string `json:"src-zone-port-entry-limit" dval:"16M"`

		SzpClistWarnThreshold int `json:"szp-clist-warn-threshold"`

		SzpWarnExceedEnable int `json:"szp-warn-exceed-enable"`

		SzpWarnThreshold int `json:"szp-warn-threshold"`

		Toggle string `json:"toggle" dval:"disable"`

		UseRoute int `json:"use-route"`

		Uuid string `json:"uuid"`

		VxlanOutboundCheck string `json:"vxlan-outbound-check" dval:"disable"`
	} `json:"protection"`
}

type DdosProtectionFastAging struct {
	HalfOpenConnRatio     int `json:"half-open-conn-ratio" dval:"25"`
	HalfOpenConnThreshold int `json:"half-open-conn-threshold" dval:"1"`
}

type DdosProtectionIpv6SrcHashMaskBits330 struct {
	MaskBitOffset1 int    `json:"mask-bit-offset-1"`
	MaskBitOffset2 int    `json:"mask-bit-offset-2"`
	MaskBitOffset3 int    `json:"mask-bit-offset-3"`
	MaskBitOffset4 int    `json:"mask-bit-offset-4"`
	MaskBitOffset5 int    `json:"mask-bit-offset-5"`
	Uuid           string `json:"uuid"`
}

type DdosProtectionMultiPuZoneDistribution331 struct {
	RegularRebalance string `json:"regular-rebalance" dval:"disable"`
	Uuid             string `json:"uuid"`
}

type DdosProtectionPerServiceSzpEntryLimit332 struct {
	DnsTcpLimit           int    `json:"dns-tcp-limit"`
	DnsUdpLimit           int    `json:"dns-udp-limit"`
	HttpLimit             int    `json:"http-limit"`
	TcpLimit              int    `json:"tcp-limit"`
	UdpLimit              int    `json:"udp-limit"`
	SslL4Limit            int    `json:"ssl-l4-limit"`
	SipUdpLimit           int    `json:"sip-udp-limit"`
	SipTcpLimit           int    `json:"sip-tcp-limit"`
	QuicLimit             int    `json:"quic-limit"`
	IpProtoIcmpV4Limit    int    `json:"ip-proto-icmp-v4-limit"`
	IpProtoIcmpV6Limit    int    `json:"ip-proto-icmp-v6-limit"`
	IpProtoOtherLimit     int    `json:"ip-proto-other-limit"`
	IpProtoGreLimit       int    `json:"ip-proto-gre-limit"`
	IpProtoIpv4EncapLimit int    `json:"ip-proto-ipv4-encap-limit"`
	IpProtoIpv6EncapLimit int    `json:"ip-proto-ipv6-encap-limit"`
	IpProtoCustomLimit    int    `json:"ip-proto-custom-limit"`
	Uuid                  string `json:"uuid"`
}

func (p *DdosProtection) GetId() string {
	return "1"
}

func (p *DdosProtection) getPath() string {
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
		if len(axResp) > 0 {
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
