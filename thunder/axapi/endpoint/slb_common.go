package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

//based on ACOS 5_2_1-P4_70
type SlbCommon struct {
	Inst struct {
		AflexTableEntryAgingInterval   int                              `json:"aflex-table-entry-aging-interval" dval:"1"`
		AflexTableEntrySync            SlbCommonAflexTableEntrySync     `json:"aflex-table-entry-sync"`
		AfterDisable                   int                              `json:"after-disable"`
		AllowInGatewayMode             int                              `json:"allow-in-gateway-mode"`
		AutoNatNoIpRefresh             string                           `json:"auto-nat-no-ip-refresh" dval:"enable"`
		AutoTranslatePort              int                              `json:"auto-translate-port"`
		CompressBlockSize              int                              `json:"compress-block-size"`
		ConnRateLimit                  SlbCommonConnRateLimit           `json:"conn-rate-limit"`
		DdosPktCountThresh             int                              `json:"ddos-pkt-count-thresh" dval:"100"`
		DdosPktSizeThresh              int                              `json:"ddos-pkt-size-thresh" dval:"64"`
		DdosProtection                 SlbCommonDdosProtection          `json:"ddos-protection"`
		DisableAdaptiveResourceCheck   int                              `json:"disable-adaptive-resource-check"`
		DisablePersistScoring          int                              `json:"disable-persist-scoring"`
		DisablePortMasking             int                              `json:"disable-port-masking"`
		DisableServerAutoReselect      int                              `json:"disable-server-auto-reselect"`
		DnsCacheAge                    int                              `json:"dns-cache-age" dval:"300"`
		DnsCacheEnable                 int                              `json:"dns-cache-enable"`
		DnsCacheEntrySize              int                              `json:"dns-cache-entry-size" dval:"256"`
		DnsResponseRateLimiting        SlbCommonDnsResponseRateLimiting `json:"dns-response-rate-limiting"`
		DnsVipStateless                int                              `json:"dns-vip-stateless"`
		DropIcmpToVipWhenVipDown       int                              `json:"drop-icmp-to-vip-when-vip-down"`
		DsrHealthCheckEnable           int                              `json:"dsr-health-check-enable"`
		EcmpHash                       string                           `json:"ecmp-hash" dval:"system-default"`
		EnableL7ReqAcct                int                              `json:"enable-l7-req-acct"`
		Entity                         string                           `json:"entity"`
		ExcludeDestination             string                           `json:"exclude-destination"`
		ExtendedStats                  int                              `json:"extended-stats"`
		FastPathDisable                int                              `json:"fast-path-disable"`
		GatewayHealthCheck             int                              `json:"gateway-health-check"`
		GracefulShutdown               int                              `json:"graceful-shutdown"`
		GracefulShutdownEnable         int                              `json:"graceful-shutdown-enable"`
		HealthCheckToAllVip            int                              `json:"health-check-to-all-vip"`
		HonorServerResponseTtl         int                              `json:"honor-server-response-ttl"`
		HttpFastEnable                 int                              `json:"http-fast-enable"`
		HwCompression                  int                              `json:"hw-compression"`
		HwSynRr                        int                              `json:"hw-syn-rr"`
		Interval                       int                              `json:"interval" dval:"5"`
		Ipv4Offset                     int                              `json:"ipv4-offset"`
		L2l3TrunkLbDisable             int                              `json:"l2l3-trunk-lb-disable"`
		LogForResetUnknownConn         int                              `json:"log-for-reset-unknown-conn"`
		LowLatency                     int                              `json:"low-latency"`
		MaxBuffQueuedPerConn           int                              `json:"max-buff-queued-per-conn" dval:"1000"`
		MaxHttpHeaderCount             int                              `json:"max-http-header-count" dval:"90"`
		MaxLocalRate                   int                              `json:"max-local-rate" dval:"32"`
		MaxRemoteRate                  int                              `json:"max-remote-rate" dval:"15000"`
		MslTime                        int                              `json:"msl-time" dval:"2"`
		MssTable                       int                              `json:"mss-table" dval:"536"`
		N5New                          int                              `json:"N5-new"`
		N5Old                          int                              `json:"N5-old"`
		NoAutoUpOnAflex                int                              `json:"no-auto-up-on-aflex"`
		OneServerConnHmRate            int                              `json:"one-server-conn-hm-rate"`
		OverridePort                   int                              `json:"override-port"`
		PerThrPercent                  int                              `json:"per-thr-percent"`
		PingSweepDetection             string                           `json:"ping-sweep-detection" dval:"disable"`
		PktRateForResetUnknownConn     int                              `json:"pkt-rate-for-reset-unknown-conn"`
		PlayerIdCheckEnable            int                              `json:"player-id-check-enable"`
		PortScanDetection              string                           `json:"port-scan-detection" dval:"disable"`
		Qat                            int                              `json:"QAT"`
		Range                          int                              `json:"range"`
		RangeEnd                       int                              `json:"range-end"`
		RangeStart                     int                              `json:"range-start"`
		RateLimitLogging               int                              `json:"rate-limit-logging"`
		ResetStaleSession              int                              `json:"reset-stale-session"`
		ResolvePortConflict            int                              `json:"resolve-port-conflict"`
		ResponseType                   string                           `json:"response-type"`
		ScaleOut                       int                              `json:"scale-out"`
		ScaleOutTrafficMap             int                              `json:"scale-out-traffic-map"`
		ServiceGroupOnNoDestNatVports  string                           `json:"service-group-on-no-dest-nat-vports" dval:"enforce-different"`
		ShowSlbServerLegacyCmd         int                              `json:"show-slb-server-legacy-cmd"`
		ShowSlbServiceGroupLegacyCmd   int                              `json:"show-slb-service-group-legacy-cmd"`
		ShowSlbVirtualServerLegacyCmd  int                              `json:"show-slb-virtual-server-legacy-cmd"`
		SnatGwyForL3                   int                              `json:"snat-gwy-for-l3"`
		SnatOnVip                      int                              `json:"snat-on-vip"`
		SnatPreserve                   SlbCommonSnatPreserve            `json:"snat-preserve"`
		Software                       int                              `json:"software"`
		SoftwareTls13                  int                              `json:"software-tls13"`
		SortRes                        int                              `json:"sort-res"`
		SsliCertNotReadyInspectLimit   int                              `json:"ssli-cert-not-ready-inspect-limit" dval:"2000"`
		SsliCertNotReadyInspectTimeout int                              `json:"ssli-cert-not-ready-inspect-timeout" dval:"10"`
		SsliSniHashEnable              int                              `json:"ssli-sni-hash-enable"`
		StatelessSgMultiBinding        int                              `json:"stateless-sg-multi-binding"`
		StatsDataDisable               int                              `json:"stats-data-disable"`
		SubstituteSourceMac            int                              `json:"substitute-source-mac"`
		Timeout                        int                              `json:"timeout" dval:"15"`
		TrafficMapType                 string                           `json:"traffic-map-type" dval:"vport"`
		TtlThreshold                   int                              `json:"ttl-threshold"`
		UseDefaultSessCount            int                              `json:"use-default-sess-count"`
		UseMssTab                      int                              `json:"use-mss-tab"`
		Uuid                           string                           `json:"uuid"`
	} `json:"common"`
}

type SlbCommonAflexTableEntrySync struct {
	AflexTableEntrySyncEnable      int    `json:"aflex-table-entry-sync-enable"`
	AflexTableEntrySyncMaxKeyLen   int    `json:"aflex-table-entry-sync-max-key-len" dval:"1000"`
	AflexTableEntrySyncMaxValueLen int    `json:"aflex-table-entry-sync-max-value-len" dval:"1000"`
	AflexTableEntrySyncMinLifetime int    `json:"aflex-table-entry-sync-min-lifetime"`
	Uuid                           string `json:"uuid"`
}

type SlbCommonConnRateLimit struct {
	SrcIpList []SlbCommonConnRateLimitSrcIpList `json:"src-ip-list"`
}

type SlbCommonConnRateLimitSrcIpList struct {
	Protocol     string `json:"protocol"`
	Limit        int    `json:"limit"`
	LimitPeriod  string `json:"limit-period"`
	Shared       int    `json:"shared"`
	ExceedAction int    `json:"exceed-action"`
	Log          int    `json:"log"`
	LockOut      int    `json:"lock-out"`
	Uuid         string `json:"uuid"`
}

type SlbCommonDdosProtection struct {
	IpdEnableToggle  string                                  `json:"ipd-enable-toggle" dval:"disable"`
	Logging          SlbCommonDdosProtectionLogging          `json:"logging"`
	PacketsPerSecond SlbCommonDdosProtectionPacketsPerSecond `json:"packets-per-second"`
}

type SlbCommonDdosProtectionLogging struct {
	IpdLoggingToggle string `json:"ipd-logging-toggle" dval:"enable"`
}

type SlbCommonDdosProtectionPacketsPerSecond struct {
	IpdTcp int `json:"ipd-tcp" dval:"200"`
	IpdUdp int `json:"ipd-udp" dval:"200"`
}

type SlbCommonDnsResponseRateLimiting struct {
	MaxTableEntries int    `json:"max-table-entries"`
	Uuid            string `json:"uuid"`
}

type SlbCommonSnatPreserve struct {
	Range []SlbCommonSnatPreserveRange `json:"range"`
}

type SlbCommonSnatPreserveRange struct {
	Port1 int `json:"port1" dval:"1025"`
	Port2 int `json:"port2" dval:"1025"`
}

func (p *SlbCommon) GetId() string {
	return "1"
}

func (p *SlbCommon) getPath() string {
	return "slb/common"
}

func (p *SlbCommon) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommon::Post")
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

func (p *SlbCommon) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommon::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SlbCommon) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommon::Put")
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

func (p *SlbCommon) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbCommon::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
