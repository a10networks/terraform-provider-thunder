package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type DdosDstZone struct {
	Inst struct {
		ActionList string `json:"action-list"`

		AdvertisedEnable int `json:"advertised-enable"`

		CaptureConfigList []DdosDstZoneCaptureConfigList `json:"capture-config-list"`

		Collector []DdosDstZoneCollector `json:"collector"`

		ContinuousLearning int `json:"continuous-learning"`

		Description string `json:"description"`

		DestNatIp string `json:"dest-nat-ip"`

		DestNatIpv6 string `json:"dest-nat-ipv6"`

		Detection DdosDstZoneDetection278 `json:"detection"`

		DropFragPkt int `json:"drop-frag-pkt"`

		EnableTopK []DdosDstZoneEnableTopK `json:"enable-top-k"`

		ForceOperationalMode int `json:"force-operational-mode"`

		Glid string `json:"glid"`

		HwBlacklistBlocking DdosDstZoneHwBlacklistBlocking290 `json:"hw-blacklist-blocking"`

		InboundForwardDscp int `json:"inbound-forward-dscp"`

		Ip []DdosDstZoneIp `json:"ip"`

		IpProto DdosDstZoneIpProto291 `json:"ip-proto"`

		Ipv6 []DdosDstZoneIpv6 `json:"ipv6"`

		IsFromWizard int `json:"is-from-wizard"`

		LogEnable int `json:"log-enable"`

		LogHighFrequency int `json:"log-high-frequency"`

		LogPeriodic int `json:"log-periodic"`

		NonRestrictive int `json:"non-restrictive"`

		OperationalMode string `json:"operational-mode" dval:"idle"`

		OutboundForwardDscp int `json:"outbound-forward-dscp"`

		OutboundPolicy DdosDstZoneOutboundPolicy292 `json:"outbound-policy"`

		PacketAnomalyDetection DdosDstZonePacketAnomalyDetection293 `json:"packet-anomaly-detection"`

		PatternRecognitionHwFilterEnable int `json:"pattern-recognition-hw-filter-enable"`

		PatternRecognitionSensitivity string `json:"pattern-recognition-sensitivity"`

		PerAddrGlid string `json:"per-addr-glid"`

		Port DdosDstZonePort294 `json:"port"`

		PortRangeList []DdosDstZonePortRangeList `json:"port-range-list"`

		RateLimit int `json:"rate-limit" dval:"1"`

		ReportingDisabled int `json:"reporting-disabled"`

		SamplingEnable []DdosDstZoneSamplingEnable `json:"sampling-enable"`

		SetCounterBaseVal int `json:"set-counter-base-val"`

		SflowCommon int `json:"sflow-common"`

		SflowHttp int `json:"sflow-http"`

		SflowLayer4 int `json:"sflow-layer-4"`

		SflowPackets int `json:"sflow-packets"`

		SflowTcp DdosDstZoneSflowTcp `json:"sflow-tcp"`

		SourceNatPool string `json:"source-nat-pool"`

		SrcPort DdosDstZoneSrcPort295 `json:"src-port"`

		SrcPortRangeList []DdosDstZoneSrcPortRangeList `json:"src-port-range-list"`

		SrcPrefixLen int `json:"src-prefix-len"`

		TelemetryEnable int `json:"telemetry-enable"`

		ThresholdSensitivity string `json:"threshold-sensitivity" dval:"OFF"`

		TopkDestinations DdosDstZoneTopkDestinations296 `json:"topk-destinations"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		WebGui DdosDstZoneWebGui297 `json:"web-gui"`

		ZoneMigrateToOpposite int `json:"zone-migrate-to-opposite"`

		ZoneName string `json:"zone-name"`

		ZoneProfile string `json:"zone-profile"`

		ZoneTemplate DdosDstZoneZoneTemplate `json:"zone-template"`
	} `json:"zone"`
}

type DdosDstZoneCaptureConfigList struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
	Uuid string `json:"uuid"`
}

type DdosDstZoneCollector struct {
	SflowName string `json:"sflow-name"`
}

type DdosDstZoneDetection278 struct {
	Settings               string                                        `json:"settings"`
	Toggle                 string                                        `json:"toggle" dval:"enable"`
	Uuid                   string                                        `json:"uuid"`
	Notification           DdosDstZoneDetectionNotification279           `json:"notification"`
	OutboundDetection      DdosDstZoneDetectionOutboundDetection281      `json:"outbound-detection"`
	ServiceDiscovery       DdosDstZoneDetectionServiceDiscovery285       `json:"service-discovery"`
	PacketAnomalyDetection DdosDstZoneDetectionPacketAnomalyDetection286 `json:"packet-anomaly-detection"`
	VictimIpDetection      DdosDstZoneDetectionVictimIpDetection288      `json:"victim-ip-detection"`
}

type DdosDstZoneDetectionNotification279 struct {
	Configuration string                                            `json:"configuration"`
	Notification  []DdosDstZoneDetectionNotificationNotification280 `json:"notification"`
	Uuid          string                                            `json:"uuid"`
}

type DdosDstZoneDetectionNotificationNotification280 struct {
	NotificationTemplateName string `json:"notification-template-name"`
}

type DdosDstZoneDetectionOutboundDetection281 struct {
	Configuration    string                                                   `json:"configuration"`
	Toggle           string                                                   `json:"toggle" dval:"disable"`
	DiscoveryMethod  string                                                   `json:"discovery-method"`
	DiscoveryRecord  int                                                      `json:"discovery-record" dval:"10"`
	EnableTopK       []DdosDstZoneDetectionOutboundDetectionEnableTopK282     `json:"enable-top-k"`
	Uuid             string                                                   `json:"uuid"`
	IndicatorList    []DdosDstZoneDetectionOutboundDetectionIndicatorList283  `json:"indicator-list"`
	TopkSourceSubnet DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet284 `json:"topk-source-subnet"`
}

type DdosDstZoneDetectionOutboundDetectionEnableTopK282 struct {
	TopkType       string `json:"topk-type"`
	TopkNetmask    int    `json:"topk-netmask" dval:"128"`
	TopkNumRecords int    `json:"topk-num-records" dval:"20"`
}

type DdosDstZoneDetectionOutboundDetectionIndicatorList283 struct {
	Type              string `json:"type"`
	TcpWindowSize     int    `json:"tcp-window-size"`
	DataPacketSize    int    `json:"data-packet-size"`
	ThresholdNum      int    `json:"threshold-num"`
	ThresholdLargeNum int    `json:"threshold-large-num"`
	ThresholdStr      string `json:"threshold-str"`
	Uuid              string `json:"uuid"`
	UserTag           string `json:"user-tag"`
}

type DdosDstZoneDetectionOutboundDetectionTopkSourceSubnet284 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneDetectionServiceDiscovery285 struct {
	Configuration    string `json:"configuration"`
	Toggle           string `json:"toggle" dval:"disable"`
	PktRateThreshold int    `json:"pkt-rate-threshold" dval:"10"`
	Uuid             string `json:"uuid"`
}

type DdosDstZoneDetectionPacketAnomalyDetection286 struct {
	Configuration string                                                       `json:"configuration"`
	Toggle        string                                                       `json:"toggle" dval:"enable"`
	Uuid          string                                                       `json:"uuid"`
	IndicatorList []DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList287 `json:"indicator-list"`
}

type DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList287 struct {
	Type         string `json:"type"`
	ThresholdNum int    `json:"threshold-num" dval:"100"`
	Uuid         string `json:"uuid"`
	UserTag      string `json:"user-tag"`
}

type DdosDstZoneDetectionVictimIpDetection288 struct {
	Configuration   string                                                  `json:"configuration"`
	Toggle          string                                                  `json:"toggle" dval:"disable"`
	HistogramToggle string                                                  `json:"histogram-toggle" dval:"histogram-disable"`
	Uuid            string                                                  `json:"uuid"`
	IndicatorList   []DdosDstZoneDetectionVictimIpDetectionIndicatorList289 `json:"indicator-list"`
}

type DdosDstZoneDetectionVictimIpDetectionIndicatorList289 struct {
	Type           string `json:"type"`
	IpThresholdNum int    `json:"ip-threshold-num"`
	Uuid           string `json:"uuid"`
	UserTag        string `json:"user-tag"`
}

type DdosDstZoneEnableTopK struct {
	TopkType       string `json:"topk-type"`
	TopkNumRecords int    `json:"topk-num-records" dval:"20"`
	TopkSortKey    string `json:"topk-sort-key" dval:"avg"`
}

type DdosDstZoneHwBlacklistBlocking290 struct {
	DstEnable int    `json:"dst-enable"`
	SrcEnable int    `json:"src-enable"`
	Uuid      string `json:"uuid"`
}

type DdosDstZoneIp struct {
	IpAddr             string `json:"ip-addr"`
	SubnetIpAddr       string `json:"subnet-ip-addr"`
	ExpandIpSubnet     int    `json:"expand-ip-subnet"`
	ExpandIpSubnetMode string `json:"expand-ip-subnet-mode" dval:"default"`
}

type DdosDstZoneIpProto291 struct {
	ProtoNumberList []DdosDstZoneIpProtoProtoNumberList `json:"proto-number-list"`
	ProtoTcpUdpList []DdosDstZoneIpProtoProtoTcpUdpList `json:"proto-tcp-udp-list"`
	ProtoNameList   []DdosDstZoneIpProtoProtoNameList   `json:"proto-name-list"`
}

type DdosDstZoneIpProtoProtoNumberList struct {
	ProtocolNum                    int                                                               `json:"protocol-num"`
	ManualModeEnable               int                                                               `json:"manual-mode-enable"`
	Deny                           int                                                               `json:"deny"`
	EspInspect                     DdosDstZoneIpProtoProtoNumberListEspInspect                       `json:"esp-inspect"`
	GlidCfg                        DdosDstZoneIpProtoProtoNumberListGlidCfg                          `json:"glid-cfg"`
	DropFragPkt                    int                                                               `json:"drop-frag-pkt"`
	UnlimitedDynamicEntryCount     int                                                               `json:"unlimited-dynamic-entry-count"`
	MaxDynamicEntryCount           int                                                               `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                               `json:"dynamic-entry-count-warn-threshold"`
	ApplyPolicyOnOverflow          int                                                               `json:"apply-policy-on-overflow"`
	EnableTopK                     int                                                               `json:"enable-top-k"`
	TopkNumRecords                 int                                                               `json:"topk-num-records" dval:"20"`
	TopkSortKey                    string                                                            `json:"topk-sort-key" dval:"avg"`
	EnableTopKDestination          int                                                               `json:"enable-top-k-destination"`
	TopkDstNumRecords              int                                                               `json:"topk-dst-num-records" dval:"20"`
	TopkDstSortKey                 string                                                            `json:"topk-dst-sort-key" dval:"avg"`
	SetCounterBaseVal              int                                                               `json:"set-counter-base-val"`
	Age                            int                                                               `json:"age" dval:"5"`
	EnableClassListOverflow        int                                                               `json:"enable-class-list-overflow"`
	FasterDeEscalation             int                                                               `json:"faster-de-escalation"`
	SflowIpFilteringPolicy         int                                                               `json:"sflow-ip-filtering-policy"`
	IpFilteringPolicy              string                                                            `json:"ip-filtering-policy"`
	Uuid                           string                                                            `json:"uuid"`
	IpFilteringPolicyStatistics    DdosDstZoneIpProtoProtoNumberListIpFilteringPolicyStatistics      `json:"ip-filtering-policy-statistics"`
	SrcBasedPolicyList             []DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList             `json:"src-based-policy-list"`
	DynamicEntryOverflowPolicyList []DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
	LevelList                      []DdosDstZoneIpProtoProtoNumberListLevelList                      `json:"level-list"`
	ManualModeList                 []DdosDstZoneIpProtoProtoNumberListManualModeList                 `json:"manual-mode-list"`
	PortInd                        DdosDstZoneIpProtoProtoNumberListPortInd                          `json:"port-ind"`
	TopkSources                    DdosDstZoneIpProtoProtoNumberListTopkSources                      `json:"topk-sources"`
	TopkDestinations               DdosDstZoneIpProtoProtoNumberListTopkDestinations                 `json:"topk-destinations"`
	ProgressionTracking            DdosDstZoneIpProtoProtoNumberListProgressionTracking              `json:"progression-tracking"`
}

type DdosDstZoneIpProtoProtoNumberListEspInspect struct {
	AuthAlgorithm    string `json:"auth-algorithm"`
	EncryptAlgorithm string `json:"encrypt-algorithm"`
	Mode             string `json:"mode"`
}

type DdosDstZoneIpProtoProtoNumberListGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZoneIpProtoProtoNumberListIpFilteringPolicyStatistics struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                   `json:"src-based-policy-name"`
	Uuid                string                                                                   `json:"uuid"`
	UserTag             string                                                                   `json:"user-tag"`
	PolicyClassListList []DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                              `json:"class-list-name"`
	ClassListGlid                  string                                                                                              `json:"class-list-glid"`
	Glid                           string                                                                                              `json:"glid"`
	GlidAction                     string                                                                                              `json:"glid-action"`
	Action                         string                                                                                              `json:"action"`
	LogEnable                      int                                                                                                 `json:"log-enable"`
	LogPeriodic                    int                                                                                                 `json:"log-periodic"`
	MaxDynamicEntryCount           int                                                                                                 `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                                 `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                              `json:"uuid"`
	UserTag                        string                                                                                              `json:"user-tag"`
	SamplingEnable                 []DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
	Logging string `json:"logging"`
	IpProto string `json:"ip-proto"`
}

type DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                                        `json:"dummy-name"`
	Glid         string                                                                                                        `json:"glid"`
	Action       string                                                                                                        `json:"action"`
	LogEnable    int                                                                                                           `json:"log-enable"`
	LogPeriodic  int                                                                                                           `json:"log-periodic"`
	ZoneTemplate DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                                        `json:"uuid"`
	UserTag      string                                                                                                        `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
	IpProto string `json:"ip-proto"`
}

type DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                                      `json:"dummy-name"`
	Glid         string                                                                      `json:"glid"`
	Action       string                                                                      `json:"action"`
	ZoneTemplate DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                      `json:"uuid"`
	UserTag      string                                                                      `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNumberListDynamicEntryOverflowPolicyListZoneTemplate struct {
	IpProto string `json:"ip-proto"`
}

type DdosDstZoneIpProtoProtoNumberListLevelList struct {
	LevelNum             string                                                    `json:"level-num"`
	SrcDefaultGlid       string                                                    `json:"src-default-glid"`
	GlidAction           string                                                    `json:"glid-action"`
	ZoneEscalationScore  int                                                       `json:"zone-escalation-score"`
	ZoneViolationActions string                                                    `json:"zone-violation-actions"`
	SrcEscalationScore   int                                                       `json:"src-escalation-score"`
	SrcViolationActions  string                                                    `json:"src-violation-actions"`
	ZoneTemplate         DdosDstZoneIpProtoProtoNumberListLevelListZoneTemplate    `json:"zone-template"`
	Uuid                 string                                                    `json:"uuid"`
	UserTag              string                                                    `json:"user-tag"`
	IndicatorList        []DdosDstZoneIpProtoProtoNumberListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZoneIpProtoProtoNumberListLevelListZoneTemplate struct {
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNumberListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	DataPacketSize        int    `json:"data-packet-size"`
	Score                 int    `json:"score"`
	SrcThresholdNum       int    `json:"src-threshold-num"`
	SrcThresholdLargeNum  int    `json:"src-threshold-large-num"`
	SrcThresholdStr       string `json:"src-threshold-str"`
	SrcViolationActions   string `json:"src-violation-actions"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	ZoneThresholdStr      string `json:"zone-threshold-str"`
	ZoneViolationActions  string `json:"zone-violation-actions"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNumberListManualModeList struct {
	Config         string                                                      `json:"config"`
	SrcDefaultGlid string                                                      `json:"src-default-glid"`
	GlidAction     string                                                      `json:"glid-action"`
	ZoneTemplate   DdosDstZoneIpProtoProtoNumberListManualModeListZoneTemplate `json:"zone-template"`
	Uuid           string                                                      `json:"uuid"`
	UserTag        string                                                      `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNumberListManualModeListZoneTemplate struct {
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNumberListPortInd struct {
	Uuid           string                                                   `json:"uuid"`
	SamplingEnable []DdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable `json:"sampling-enable"`
}

type DdosDstZoneIpProtoProtoNumberListPortIndSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZoneIpProtoProtoNumberListTopkSources struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNumberListTopkDestinations struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNumberListProgressionTracking struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoTcpUdpList struct {
	Protocol                    string                                                       `json:"protocol"`
	Deny                        int                                                          `json:"deny"`
	GlidCfg                     DdosDstZoneIpProtoProtoTcpUdpListGlidCfg                     `json:"glid-cfg"`
	DropFragPkt                 int                                                          `json:"drop-frag-pkt"`
	SetCounterBaseVal           int                                                          `json:"set-counter-base-val"`
	SflowIpFilteringPolicy      int                                                          `json:"sflow-ip-filtering-policy"`
	IpFilteringPolicy           string                                                       `json:"ip-filtering-policy"`
	SameSourceDestPortDrop      int                                                          `json:"same-source-dest-port-drop"`
	Uuid                        string                                                       `json:"uuid"`
	IpFilteringPolicyStatistics DdosDstZoneIpProtoProtoTcpUdpListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
}

type DdosDstZoneIpProtoProtoTcpUdpListGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	PerAddrGlid string `json:"per-addr-glid"`
	ActionList  string `json:"action-list"`
}

type DdosDstZoneIpProtoProtoTcpUdpListIpFilteringPolicyStatistics struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNameList struct {
	Protocol                       string                                                          `json:"protocol"`
	ManualModeEnable               int                                                             `json:"manual-mode-enable"`
	Deny                           int                                                             `json:"deny"`
	GlidCfg                        DdosDstZoneIpProtoProtoNameListGlidCfg                          `json:"glid-cfg"`
	TunnelDecap                    int                                                             `json:"tunnel-decap"`
	KeyCfg                         []DdosDstZoneIpProtoProtoNameListKeyCfg                         `json:"key-cfg"`
	TunnelRateLimit                int                                                             `json:"tunnel-rate-limit"`
	DropFragPkt                    int                                                             `json:"drop-frag-pkt"`
	UnlimitedDynamicEntryCount     int                                                             `json:"unlimited-dynamic-entry-count"`
	MaxDynamicEntryCount           int                                                             `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                             `json:"dynamic-entry-count-warn-threshold"`
	ApplyPolicyOnOverflow          int                                                             `json:"apply-policy-on-overflow"`
	EnableTopK                     int                                                             `json:"enable-top-k"`
	TopkNumRecords                 int                                                             `json:"topk-num-records" dval:"20"`
	TopkSortKey                    string                                                          `json:"topk-sort-key" dval:"avg"`
	EnableTopKDestination          int                                                             `json:"enable-top-k-destination"`
	TopkDstNumRecords              int                                                             `json:"topk-dst-num-records" dval:"20"`
	TopkDstSortKey                 string                                                          `json:"topk-dst-sort-key" dval:"avg"`
	SetCounterBaseVal              int                                                             `json:"set-counter-base-val"`
	Age                            int                                                             `json:"age" dval:"5"`
	EnableClassListOverflow        int                                                             `json:"enable-class-list-overflow"`
	FasterDeEscalation             int                                                             `json:"faster-de-escalation"`
	SflowIpFilteringPolicy         int                                                             `json:"sflow-ip-filtering-policy"`
	IpFilteringPolicy              string                                                          `json:"ip-filtering-policy"`
	Uuid                           string                                                          `json:"uuid"`
	IpFilteringPolicyStatistics    DdosDstZoneIpProtoProtoNameListIpFilteringPolicyStatistics      `json:"ip-filtering-policy-statistics"`
	LevelList                      []DdosDstZoneIpProtoProtoNameListLevelList                      `json:"level-list"`
	ManualModeList                 []DdosDstZoneIpProtoProtoNameListManualModeList                 `json:"manual-mode-list"`
	SrcBasedPolicyList             []DdosDstZoneIpProtoProtoNameListSrcBasedPolicyList             `json:"src-based-policy-list"`
	DynamicEntryOverflowPolicyList []DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
	PortInd                        DdosDstZoneIpProtoProtoNameListPortInd                          `json:"port-ind"`
	TopkSources                    DdosDstZoneIpProtoProtoNameListTopkSources                      `json:"topk-sources"`
	ProgressionTracking            DdosDstZoneIpProtoProtoNameListProgressionTracking              `json:"progression-tracking"`
	TopkDestinations               DdosDstZoneIpProtoProtoNameListTopkDestinations                 `json:"topk-destinations"`
}

type DdosDstZoneIpProtoProtoNameListGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZoneIpProtoProtoNameListKeyCfg struct {
	Key string `json:"key"`
}

type DdosDstZoneIpProtoProtoNameListIpFilteringPolicyStatistics struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNameListLevelList struct {
	LevelNum             string                                                  `json:"level-num"`
	SrcDefaultGlid       string                                                  `json:"src-default-glid"`
	GlidAction           string                                                  `json:"glid-action"`
	ZoneEscalationScore  int                                                     `json:"zone-escalation-score"`
	ZoneViolationActions string                                                  `json:"zone-violation-actions"`
	SrcEscalationScore   int                                                     `json:"src-escalation-score"`
	SrcViolationActions  string                                                  `json:"src-violation-actions"`
	ZoneTemplate         DdosDstZoneIpProtoProtoNameListLevelListZoneTemplate    `json:"zone-template"`
	Uuid                 string                                                  `json:"uuid"`
	UserTag              string                                                  `json:"user-tag"`
	IndicatorList        []DdosDstZoneIpProtoProtoNameListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZoneIpProtoProtoNameListLevelListZoneTemplate struct {
	IcmpV4  string `json:"icmp-v4"`
	IcmpV6  string `json:"icmp-v6"`
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNameListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	DataPacketSize        int    `json:"data-packet-size"`
	Score                 int    `json:"score"`
	SrcThresholdNum       int    `json:"src-threshold-num"`
	SrcThresholdLargeNum  int    `json:"src-threshold-large-num"`
	SrcThresholdStr       string `json:"src-threshold-str"`
	SrcViolationActions   string `json:"src-violation-actions"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	ZoneThresholdStr      string `json:"zone-threshold-str"`
	ZoneViolationActions  string `json:"zone-violation-actions"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNameListManualModeList struct {
	Config         string                                                    `json:"config"`
	SrcDefaultGlid string                                                    `json:"src-default-glid"`
	GlidAction     string                                                    `json:"glid-action"`
	ZoneTemplate   DdosDstZoneIpProtoProtoNameListManualModeListZoneTemplate `json:"zone-template"`
	Uuid           string                                                    `json:"uuid"`
	UserTag        string                                                    `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNameListManualModeListZoneTemplate struct {
	IcmpV4  string `json:"icmp-v4"`
	IcmpV6  string `json:"icmp-v6"`
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNameListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                 `json:"src-based-policy-name"`
	Uuid                string                                                                 `json:"uuid"`
	UserTag             string                                                                 `json:"user-tag"`
	PolicyClassListList []DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                            `json:"class-list-name"`
	ClassListGlid                  string                                                                                            `json:"class-list-glid"`
	Glid                           string                                                                                            `json:"glid"`
	GlidAction                     string                                                                                            `json:"glid-action"`
	Action                         string                                                                                            `json:"action"`
	LogEnable                      int                                                                                               `json:"log-enable"`
	LogPeriodic                    int                                                                                               `json:"log-periodic"`
	MaxDynamicEntryCount           int                                                                                               `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                               `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                            `json:"uuid"`
	UserTag                        string                                                                                            `json:"user-tag"`
	SamplingEnable                 []DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
	Logging string `json:"logging"`
	IcmpV4  string `json:"icmp-v4"`
	IcmpV6  string `json:"icmp-v6"`
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                                      `json:"dummy-name"`
	Glid         string                                                                                                      `json:"glid"`
	Action       string                                                                                                      `json:"action"`
	LogEnable    int                                                                                                         `json:"log-enable"`
	LogPeriodic  int                                                                                                         `json:"log-periodic"`
	ZoneTemplate DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                                      `json:"uuid"`
	UserTag      string                                                                                                      `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
	IcmpV4  string `json:"icmp-v4"`
	IcmpV6  string `json:"icmp-v6"`
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                                    `json:"dummy-name"`
	Glid         string                                                                    `json:"glid"`
	Action       string                                                                    `json:"action"`
	ZoneTemplate DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                    `json:"uuid"`
	UserTag      string                                                                    `json:"user-tag"`
}

type DdosDstZoneIpProtoProtoNameListDynamicEntryOverflowPolicyListZoneTemplate struct {
	IcmpV4  string `json:"icmp-v4"`
	IcmpV6  string `json:"icmp-v6"`
	IpProto string `json:"ip-proto"`
	Encap   string `json:"encap"`
}

type DdosDstZoneIpProtoProtoNameListPortInd struct {
	Uuid           string                                                 `json:"uuid"`
	SamplingEnable []DdosDstZoneIpProtoProtoNameListPortIndSamplingEnable `json:"sampling-enable"`
}

type DdosDstZoneIpProtoProtoNameListPortIndSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZoneIpProtoProtoNameListTopkSources struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNameListProgressionTracking struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpProtoProtoNameListTopkDestinations struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneIpv6 struct {
	Ip6Addr              string `json:"ip6-addr"`
	SubnetIpv6Addr       string `json:"subnet-ipv6-addr"`
	ExpandIpv6Subnet     int    `json:"expand-ipv6-subnet"`
	ExpandIpv6SubnetMode string `json:"expand-ipv6-subnet-mode" dval:"default"`
}

type DdosDstZoneOutboundPolicy292 struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type DdosDstZonePacketAnomalyDetection293 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePort294 struct {
	ZoneServiceList      []DdosDstZonePortZoneServiceList      `json:"zone-service-list"`
	ZoneServiceOtherList []DdosDstZonePortZoneServiceOtherList `json:"zone-service-other-list"`
}

type DdosDstZonePortZoneServiceList struct {
	PortNum                        int                                                            `json:"port-num"`
	Protocol                       string                                                         `json:"protocol"`
	ManualModeEnable               int                                                            `json:"manual-mode-enable"`
	Deny                           int                                                            `json:"deny"`
	GlidCfg                        DdosDstZonePortZoneServiceListGlidCfg                          `json:"glid-cfg"`
	Stateful                       int                                                            `json:"stateful"`
	DefaultActionList              string                                                         `json:"default-action-list"`
	SflowCommon                    int                                                            `json:"sflow-common"`
	SflowPackets                   int                                                            `json:"sflow-packets"`
	SflowIpFilteringPolicy         int                                                            `json:"sflow-ip-filtering-policy"`
	SflowTcp                       DdosDstZonePortZoneServiceListSflowTcp                         `json:"sflow-tcp"`
	SflowHttp                      int                                                            `json:"sflow-http"`
	UnlimitedDynamicEntryCount     int                                                            `json:"unlimited-dynamic-entry-count"`
	MaxDynamicEntryCount           int                                                            `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                            `json:"dynamic-entry-count-warn-threshold"`
	ApplyPolicyOnOverflow          int                                                            `json:"apply-policy-on-overflow"`
	EnableClassListOverflow        int                                                            `json:"enable-class-list-overflow"`
	Age                            int                                                            `json:"age" dval:"5"`
	EnableTopK                     int                                                            `json:"enable-top-k"`
	TopkNumRecords                 int                                                            `json:"topk-num-records" dval:"20"`
	TopkSortKey                    string                                                         `json:"topk-sort-key" dval:"avg"`
	EnableTopKDestination          int                                                            `json:"enable-top-k-destination"`
	TopkDstNumRecords              int                                                            `json:"topk-dst-num-records" dval:"20"`
	TopkDstSortKey                 string                                                         `json:"topk-dst-sort-key" dval:"avg"`
	CaptureConfig                  DdosDstZonePortZoneServiceListCaptureConfig                    `json:"capture-config"`
	SetCounterBaseVal              int                                                            `json:"set-counter-base-val"`
	OutboundOnly                   int                                                            `json:"outbound-only"`
	FasterDeEscalation             int                                                            `json:"faster-de-escalation"`
	IpFilteringPolicy              string                                                         `json:"ip-filtering-policy"`
	SameSourceDestPortDrop         int                                                            `json:"same-source-dest-port-drop"`
	Uuid                           string                                                         `json:"uuid"`
	PatternRecognition             DdosDstZonePortZoneServiceListPatternRecognition               `json:"pattern-recognition"`
	PatternRecognitionPuDetails    DdosDstZonePortZoneServiceListPatternRecognitionPuDetails      `json:"pattern-recognition-pu-details"`
	IpFilteringPolicyStatistics    DdosDstZonePortZoneServiceListIpFilteringPolicyStatistics      `json:"ip-filtering-policy-statistics"`
	LevelList                      []DdosDstZonePortZoneServiceListLevelList                      `json:"level-list"`
	ManualModeList                 []DdosDstZonePortZoneServiceListManualModeList                 `json:"manual-mode-list"`
	PortInd                        DdosDstZonePortZoneServiceListPortInd                          `json:"port-ind"`
	TopkSources                    DdosDstZonePortZoneServiceListTopkSources                      `json:"topk-sources"`
	ProgressionTracking            DdosDstZonePortZoneServiceListProgressionTracking              `json:"progression-tracking"`
	TopkDestinations               DdosDstZonePortZoneServiceListTopkDestinations                 `json:"topk-destinations"`
	SrcBasedPolicyList             []DdosDstZonePortZoneServiceListSrcBasedPolicyList             `json:"src-based-policy-list"`
	DynamicEntryOverflowPolicyList []DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
	Virtualhosts                   DdosDstZonePortZoneServiceListVirtualhosts                     `json:"virtualhosts"`
}

type DdosDstZonePortZoneServiceListGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZonePortZoneServiceListSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
	SflowSampSni     int `json:"sflow-samp-sni"`
}

type DdosDstZonePortZoneServiceListCaptureConfig struct {
	CaptureConfigName string `json:"capture-config-name"`
	CaptureConfigMode string `json:"capture-config-mode"`
}

type DdosDstZonePortZoneServiceListPatternRecognition struct {
	Algorithm               string `json:"algorithm"`
	Mode                    string `json:"mode"`
	Sensitivity             string `json:"sensitivity"`
	FilterThreshold         int    `json:"filter-threshold"`
	FilterInactiveThreshold int    `json:"filter-inactive-threshold"`
	TriggeredBy             string `json:"triggered-by"`
	CaptureTraffic          string `json:"capture-traffic"`
	AppPayloadOffset        int    `json:"app-payload-offset"`
	Uuid                    string `json:"uuid"`
}

type DdosDstZonePortZoneServiceListPatternRecognitionPuDetails struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceListIpFilteringPolicyStatistics struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceListLevelList struct {
	LevelNum                      string                                                 `json:"level-num"`
	SrcDefaultGlid                string                                                 `json:"src-default-glid"`
	GlidAction                    string                                                 `json:"glid-action"`
	ZoneEscalationScore           int                                                    `json:"zone-escalation-score"`
	ZoneViolationActions          string                                                 `json:"zone-violation-actions"`
	SrcEscalationScore            int                                                    `json:"src-escalation-score"`
	SrcViolationActions           string                                                 `json:"src-violation-actions"`
	ZoneTemplate                  DdosDstZonePortZoneServiceListLevelListZoneTemplate    `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                    `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                    `json:"close-sessions-for-all-sources"`
	ClearSourcesUponDeescalation  int                                                    `json:"clear-sources-upon-deescalation"`
	StartSignatureExtraction      int                                                    `json:"start-signature-extraction"`
	StartPatternRecognition       int                                                    `json:"start-pattern-recognition"`
	ApplyExtractedFilters         int                                                    `json:"apply-extracted-filters"`
	Uuid                          string                                                 `json:"uuid"`
	UserTag                       string                                                 `json:"user-tag"`
	IndicatorList                 []DdosDstZonePortZoneServiceListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZonePortZoneServiceListLevelListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	TcpWindowSize         int    `json:"tcp-window-size"`
	DataPacketSize        int    `json:"data-packet-size"`
	Score                 int    `json:"score"`
	SrcThresholdNum       int    `json:"src-threshold-num"`
	SrcThresholdLargeNum  int    `json:"src-threshold-large-num"`
	SrcThresholdStr       string `json:"src-threshold-str"`
	SrcViolationActions   string `json:"src-violation-actions"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdStr      string `json:"zone-threshold-str"`
	ZoneViolationActions  string `json:"zone-violation-actions"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZonePortZoneServiceListManualModeList struct {
	Config                        string                                                   `json:"config"`
	SrcDefaultGlid                string                                                   `json:"src-default-glid"`
	GlidAction                    string                                                   `json:"glid-action"`
	ZoneTemplate                  DdosDstZonePortZoneServiceListManualModeListZoneTemplate `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                      `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                      `json:"close-sessions-for-all-sources"`
	Uuid                          string                                                   `json:"uuid"`
	UserTag                       string                                                   `json:"user-tag"`
}

type DdosDstZonePortZoneServiceListManualModeListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceListPortInd struct {
	Uuid           string                                                `json:"uuid"`
	SamplingEnable []DdosDstZonePortZoneServiceListPortIndSamplingEnable `json:"sampling-enable"`
}

type DdosDstZonePortZoneServiceListPortIndSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceListTopkSources struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceListProgressionTracking struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceListTopkDestinations struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                `json:"src-based-policy-name"`
	Uuid                string                                                                `json:"uuid"`
	UserTag             string                                                                `json:"user-tag"`
	PolicyClassListList []DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                           `json:"class-list-name"`
	ClassListGlid                  string                                                                                           `json:"class-list-glid"`
	Glid                           string                                                                                           `json:"glid"`
	GlidAction                     string                                                                                           `json:"glid-action"`
	Action                         string                                                                                           `json:"action"`
	LogEnable                      int                                                                                              `json:"log-enable"`
	LogPeriodic                    int                                                                                              `json:"log-periodic"`
	MaxDynamicEntryCount           int                                                                                              `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                              `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                           `json:"uuid"`
	UserTag                        string                                                                                           `json:"user-tag"`
	SamplingEnable                 []DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
	Quic    string `json:"quic"`
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                                     `json:"dummy-name"`
	Glid         string                                                                                                     `json:"glid"`
	Action       string                                                                                                     `json:"action"`
	LogEnable    int                                                                                                        `json:"log-enable"`
	LogPeriodic  int                                                                                                        `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                                     `json:"uuid"`
	UserTag      string                                                                                                     `json:"user-tag"`
}

type DdosDstZonePortZoneServiceListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
	Quic    string `json:"quic"`
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                                   `json:"dummy-name"`
	Glid         string                                                                   `json:"glid"`
	Action       string                                                                   `json:"action"`
	LogEnable    int                                                                      `json:"log-enable"`
	LogPeriodic  int                                                                      `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                   `json:"uuid"`
	UserTag      string                                                                   `json:"user-tag"`
}

type DdosDstZonePortZoneServiceListDynamicEntryOverflowPolicyListZoneTemplate struct {
	Quic    string `json:"quic"`
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceListVirtualhosts struct {
	VhostsConfig      string                                                      `json:"vhosts-config"`
	SourceTrackingAll int                                                         `json:"source-tracking-all"`
	Uuid              string                                                      `json:"uuid"`
	VirtualhostList   []DdosDstZonePortZoneServiceListVirtualhostsVirtualhostList `json:"virtualhost-list"`
}

type DdosDstZonePortZoneServiceListVirtualhostsVirtualhostList struct {
	Vhost              string                                                                `json:"vhost"`
	Servername         []DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListServername `json:"servername"`
	ServernameList     string                                                                `json:"servername-list"`
	ServernameMatchAny int                                                                   `json:"servername-match-any"`
	ServernameNoSni    int                                                                   `json:"servername-no-sni"`
	SourceTracking     string                                                                `json:"source-tracking" dval:"follow"`
	GlidCfg            DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListGlidCfg      `json:"glid-cfg"`
	Deny               int                                                                   `json:"deny"`
	Uuid               string                                                                `json:"uuid"`
	UserTag            string                                                                `json:"user-tag"`
	LevelList          []DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListLevelList  `json:"level-list"`
}

type DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListServername struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

type DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListLevelList struct {
	LevelNum       string                                                                         `json:"level-num"`
	SrcDefaultGlid string                                                                         `json:"src-default-glid"`
	GlidAction     string                                                                         `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListLevelListZoneTemplate `json:"zone-template"`
	Uuid           string                                                                         `json:"uuid"`
	UserTag        string                                                                         `json:"user-tag"`
}

type DdosDstZonePortZoneServiceListVirtualhostsVirtualhostListLevelListZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

type DdosDstZonePortZoneServiceOtherList struct {
	PortOther                      string                                                              `json:"port-other"`
	Protocol                       string                                                              `json:"protocol"`
	ManualModeEnable               int                                                                 `json:"manual-mode-enable"`
	EnableTopK                     int                                                                 `json:"enable-top-k"`
	TopkNumRecords                 int                                                                 `json:"topk-num-records" dval:"20"`
	TopkSortKey                    string                                                              `json:"topk-sort-key" dval:"avg"`
	EnableTopKDestination          int                                                                 `json:"enable-top-k-destination"`
	TopkDstNumRecords              int                                                                 `json:"topk-dst-num-records" dval:"20"`
	TopkDstSortKey                 string                                                              `json:"topk-dst-sort-key" dval:"avg"`
	Deny                           int                                                                 `json:"deny"`
	GlidCfg                        DdosDstZonePortZoneServiceOtherListGlidCfg                          `json:"glid-cfg"`
	Stateful                       int                                                                 `json:"stateful"`
	DefaultActionList              string                                                              `json:"default-action-list"`
	SflowCommon                    int                                                                 `json:"sflow-common"`
	SflowPackets                   int                                                                 `json:"sflow-packets"`
	SflowTcp                       DdosDstZonePortZoneServiceOtherListSflowTcp                         `json:"sflow-tcp"`
	SflowIpFilteringPolicy         int                                                                 `json:"sflow-ip-filtering-policy"`
	UnlimitedDynamicEntryCount     int                                                                 `json:"unlimited-dynamic-entry-count"`
	MaxDynamicEntryCount           int                                                                 `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                 `json:"dynamic-entry-count-warn-threshold"`
	ApplyPolicyOnOverflow          int                                                                 `json:"apply-policy-on-overflow"`
	SetCounterBaseVal              int                                                                 `json:"set-counter-base-val"`
	EnableClassListOverflow        int                                                                 `json:"enable-class-list-overflow"`
	Age                            int                                                                 `json:"age" dval:"5"`
	OutboundOnly                   int                                                                 `json:"outbound-only"`
	FasterDeEscalation             int                                                                 `json:"faster-de-escalation"`
	IpFilteringPolicy              string                                                              `json:"ip-filtering-policy"`
	SameSourceDestPortDrop         int                                                                 `json:"same-source-dest-port-drop"`
	Uuid                           string                                                              `json:"uuid"`
	IpFilteringPolicyStatistics    DdosDstZonePortZoneServiceOtherListIpFilteringPolicyStatistics      `json:"ip-filtering-policy-statistics"`
	PatternRecognition             DdosDstZonePortZoneServiceOtherListPatternRecognition               `json:"pattern-recognition"`
	PatternRecognitionPuDetails    DdosDstZonePortZoneServiceOtherListPatternRecognitionPuDetails      `json:"pattern-recognition-pu-details"`
	LevelList                      []DdosDstZonePortZoneServiceOtherListLevelList                      `json:"level-list"`
	ManualModeList                 []DdosDstZonePortZoneServiceOtherListManualModeList                 `json:"manual-mode-list"`
	PortInd                        DdosDstZonePortZoneServiceOtherListPortInd                          `json:"port-ind"`
	TopkSources                    DdosDstZonePortZoneServiceOtherListTopkSources                      `json:"topk-sources"`
	ProgressionTracking            DdosDstZonePortZoneServiceOtherListProgressionTracking              `json:"progression-tracking"`
	TopkDestinations               DdosDstZonePortZoneServiceOtherListTopkDestinations                 `json:"topk-destinations"`
	SrcBasedPolicyList             []DdosDstZonePortZoneServiceOtherListSrcBasedPolicyList             `json:"src-based-policy-list"`
	DynamicEntryOverflowPolicyList []DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
}

type DdosDstZonePortZoneServiceOtherListGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZonePortZoneServiceOtherListSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
}

type DdosDstZonePortZoneServiceOtherListIpFilteringPolicyStatistics struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherListPatternRecognition struct {
	Algorithm               string `json:"algorithm"`
	Mode                    string `json:"mode"`
	Sensitivity             string `json:"sensitivity"`
	FilterThreshold         int    `json:"filter-threshold"`
	FilterInactiveThreshold int    `json:"filter-inactive-threshold"`
	TriggeredBy             string `json:"triggered-by"`
	CaptureTraffic          string `json:"capture-traffic"`
	Uuid                    string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherListPatternRecognitionPuDetails struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherListLevelList struct {
	LevelNum                      string                                                      `json:"level-num"`
	SrcDefaultGlid                string                                                      `json:"src-default-glid"`
	GlidAction                    string                                                      `json:"glid-action"`
	ZoneEscalationScore           int                                                         `json:"zone-escalation-score"`
	ZoneViolationActions          string                                                      `json:"zone-violation-actions"`
	SrcEscalationScore            int                                                         `json:"src-escalation-score"`
	SrcViolationActions           string                                                      `json:"src-violation-actions"`
	ZoneTemplate                  DdosDstZonePortZoneServiceOtherListLevelListZoneTemplate    `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                         `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                         `json:"close-sessions-for-all-sources"`
	ClearSourcesUponDeescalation  int                                                         `json:"clear-sources-upon-deescalation"`
	StartPatternRecognition       int                                                         `json:"start-pattern-recognition"`
	ApplyExtractedFilters         int                                                         `json:"apply-extracted-filters"`
	Uuid                          string                                                      `json:"uuid"`
	UserTag                       string                                                      `json:"user-tag"`
	IndicatorList                 []DdosDstZonePortZoneServiceOtherListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZonePortZoneServiceOtherListLevelListZoneTemplate struct {
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceOtherListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	TcpWindowSize         int    `json:"tcp-window-size"`
	DataPacketSize        int    `json:"data-packet-size"`
	Score                 int    `json:"score"`
	SrcThresholdNum       int    `json:"src-threshold-num"`
	SrcThresholdLargeNum  int    `json:"src-threshold-large-num"`
	SrcThresholdStr       string `json:"src-threshold-str"`
	SrcViolationActions   string `json:"src-violation-actions"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	ZoneThresholdStr      string `json:"zone-threshold-str"`
	ZoneViolationActions  string `json:"zone-violation-actions"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherListManualModeList struct {
	Config                        string                                                        `json:"config"`
	SrcDefaultGlid                string                                                        `json:"src-default-glid"`
	GlidAction                    string                                                        `json:"glid-action"`
	ZoneTemplate                  DdosDstZonePortZoneServiceOtherListManualModeListZoneTemplate `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                           `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                           `json:"close-sessions-for-all-sources"`
	Uuid                          string                                                        `json:"uuid"`
	UserTag                       string                                                        `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherListManualModeListZoneTemplate struct {
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceOtherListPortInd struct {
	Uuid           string                                                     `json:"uuid"`
	SamplingEnable []DdosDstZonePortZoneServiceOtherListPortIndSamplingEnable `json:"sampling-enable"`
}

type DdosDstZonePortZoneServiceOtherListPortIndSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceOtherListTopkSources struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherListProgressionTracking struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherListTopkDestinations struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                     `json:"src-based-policy-name"`
	Uuid                string                                                                     `json:"uuid"`
	UserTag             string                                                                     `json:"user-tag"`
	PolicyClassListList []DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                                `json:"class-list-name"`
	ClassListGlid                  string                                                                                                `json:"class-list-glid"`
	Glid                           string                                                                                                `json:"glid"`
	GlidAction                     string                                                                                                `json:"glid-action"`
	Action                         string                                                                                                `json:"action"`
	MaxDynamicEntryCount           int                                                                                                   `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                                   `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                                `json:"uuid"`
	UserTag                        string                                                                                                `json:"user-tag"`
	SamplingEnable                 []DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                                          `json:"dummy-name"`
	Glid         string                                                                                                          `json:"glid"`
	Action       string                                                                                                          `json:"action"`
	LogEnable    int                                                                                                             `json:"log-enable"`
	LogPeriodic  int                                                                                                             `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                                          `json:"uuid"`
	UserTag      string                                                                                                          `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                                        `json:"dummy-name"`
	Glid         string                                                                        `json:"glid"`
	Action       string                                                                        `json:"action"`
	LogEnable    int                                                                           `json:"log-enable"`
	LogPeriodic  int                                                                           `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                        `json:"uuid"`
	UserTag      string                                                                        `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherListDynamicEntryOverflowPolicyListZoneTemplate struct {
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortRangeList struct {
	PortRangeStart                 int                                                      `json:"port-range-start"`
	PortRangeEnd                   int                                                      `json:"port-range-end"`
	Protocol                       string                                                   `json:"protocol"`
	ManualModeEnable               int                                                      `json:"manual-mode-enable"`
	Deny                           int                                                      `json:"deny"`
	GlidCfg                        DdosDstZonePortRangeListGlidCfg                          `json:"glid-cfg"`
	Stateful                       int                                                      `json:"stateful"`
	DefaultActionList              string                                                   `json:"default-action-list"`
	SflowCommon                    int                                                      `json:"sflow-common"`
	SflowPackets                   int                                                      `json:"sflow-packets"`
	SflowIpFilteringPolicy         int                                                      `json:"sflow-ip-filtering-policy"`
	SflowTcp                       DdosDstZonePortRangeListSflowTcp                         `json:"sflow-tcp"`
	SflowHttp                      int                                                      `json:"sflow-http"`
	UnlimitedDynamicEntryCount     int                                                      `json:"unlimited-dynamic-entry-count"`
	MaxDynamicEntryCount           int                                                      `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                      `json:"dynamic-entry-count-warn-threshold"`
	ApplyPolicyOnOverflow          int                                                      `json:"apply-policy-on-overflow"`
	EnableClassListOverflow        int                                                      `json:"enable-class-list-overflow"`
	EnableTopK                     int                                                      `json:"enable-top-k"`
	TopkNumRecords                 int                                                      `json:"topk-num-records" dval:"20"`
	TopkSortKey                    string                                                   `json:"topk-sort-key" dval:"avg"`
	EnableTopKDestination          int                                                      `json:"enable-top-k-destination"`
	TopkDstNumRecords              int                                                      `json:"topk-dst-num-records" dval:"20"`
	TopkDstSortKey                 string                                                   `json:"topk-dst-sort-key" dval:"avg"`
	SetCounterBaseVal              int                                                      `json:"set-counter-base-val"`
	Age                            int                                                      `json:"age" dval:"5"`
	OutboundOnly                   int                                                      `json:"outbound-only"`
	FasterDeEscalation             int                                                      `json:"faster-de-escalation"`
	CaptureConfig                  DdosDstZonePortRangeListCaptureConfig                    `json:"capture-config"`
	IpFilteringPolicy              string                                                   `json:"ip-filtering-policy"`
	SameSourceDestPortDrop         int                                                      `json:"same-source-dest-port-drop"`
	Uuid                           string                                                   `json:"uuid"`
	UserTag                        string                                                   `json:"user-tag"`
	IpFilteringPolicyStatistics    DdosDstZonePortRangeListIpFilteringPolicyStatistics      `json:"ip-filtering-policy-statistics"`
	PatternRecognition             DdosDstZonePortRangeListPatternRecognition               `json:"pattern-recognition"`
	PatternRecognitionPuDetails    DdosDstZonePortRangeListPatternRecognitionPuDetails      `json:"pattern-recognition-pu-details"`
	LevelList                      []DdosDstZonePortRangeListLevelList                      `json:"level-list"`
	ManualModeList                 []DdosDstZonePortRangeListManualModeList                 `json:"manual-mode-list"`
	PortInd                        DdosDstZonePortRangeListPortInd                          `json:"port-ind"`
	TopkSources                    DdosDstZonePortRangeListTopkSources                      `json:"topk-sources"`
	TopkDestinations               DdosDstZonePortRangeListTopkDestinations                 `json:"topk-destinations"`
	ProgressionTracking            DdosDstZonePortRangeListProgressionTracking              `json:"progression-tracking"`
	SrcBasedPolicyList             []DdosDstZonePortRangeListSrcBasedPolicyList             `json:"src-based-policy-list"`
	DynamicEntryOverflowPolicyList []DdosDstZonePortRangeListDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`
	Virtualhosts                   DdosDstZonePortRangeListVirtualhosts                     `json:"virtualhosts"`
}

type DdosDstZonePortRangeListGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZonePortRangeListSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
}

type DdosDstZonePortRangeListCaptureConfig struct {
	CaptureConfigName string `json:"capture-config-name"`
	CaptureConfigMode string `json:"capture-config-mode"`
}

type DdosDstZonePortRangeListIpFilteringPolicyStatistics struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeListPatternRecognition struct {
	Algorithm               string `json:"algorithm"`
	Mode                    string `json:"mode"`
	Sensitivity             string `json:"sensitivity"`
	FilterThreshold         int    `json:"filter-threshold"`
	FilterInactiveThreshold int    `json:"filter-inactive-threshold"`
	TriggeredBy             string `json:"triggered-by"`
	CaptureTraffic          string `json:"capture-traffic"`
	AppPayloadOffset        int    `json:"app-payload-offset"`
	Uuid                    string `json:"uuid"`
}

type DdosDstZonePortRangeListPatternRecognitionPuDetails struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeListLevelList struct {
	LevelNum                      string                                           `json:"level-num"`
	SrcDefaultGlid                string                                           `json:"src-default-glid"`
	GlidAction                    string                                           `json:"glid-action"`
	ZoneEscalationScore           int                                              `json:"zone-escalation-score"`
	ZoneViolationActions          string                                           `json:"zone-violation-actions"`
	SrcEscalationScore            int                                              `json:"src-escalation-score"`
	SrcViolationActions           string                                           `json:"src-violation-actions"`
	ZoneTemplate                  DdosDstZonePortRangeListLevelListZoneTemplate    `json:"zone-template"`
	CloseSessionsForUnauthSources int                                              `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                              `json:"close-sessions-for-all-sources"`
	ClearSourcesUponDeescalation  int                                              `json:"clear-sources-upon-deescalation"`
	StartPatternRecognition       int                                              `json:"start-pattern-recognition"`
	ApplyExtractedFilters         int                                              `json:"apply-extracted-filters"`
	Uuid                          string                                           `json:"uuid"`
	UserTag                       string                                           `json:"user-tag"`
	IndicatorList                 []DdosDstZonePortRangeListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZonePortRangeListLevelListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortRangeListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	TcpWindowSize         int    `json:"tcp-window-size"`
	DataPacketSize        int    `json:"data-packet-size"`
	Score                 int    `json:"score"`
	SrcThresholdNum       int    `json:"src-threshold-num"`
	SrcThresholdLargeNum  int    `json:"src-threshold-large-num"`
	SrcThresholdStr       string `json:"src-threshold-str"`
	SrcViolationActions   string `json:"src-violation-actions"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	ZoneThresholdStr      string `json:"zone-threshold-str"`
	ZoneViolationActions  string `json:"zone-violation-actions"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZonePortRangeListManualModeList struct {
	Config                        string                                             `json:"config"`
	SrcDefaultGlid                string                                             `json:"src-default-glid"`
	GlidAction                    string                                             `json:"glid-action"`
	ZoneTemplate                  DdosDstZonePortRangeListManualModeListZoneTemplate `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                `json:"close-sessions-for-all-sources"`
	Uuid                          string                                             `json:"uuid"`
	UserTag                       string                                             `json:"user-tag"`
}

type DdosDstZonePortRangeListManualModeListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortRangeListPortInd struct {
	Uuid           string                                          `json:"uuid"`
	SamplingEnable []DdosDstZonePortRangeListPortIndSamplingEnable `json:"sampling-enable"`
}

type DdosDstZonePortRangeListPortIndSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortRangeListTopkSources struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeListTopkDestinations struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeListProgressionTracking struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                          `json:"src-based-policy-name"`
	Uuid                string                                                          `json:"uuid"`
	UserTag             string                                                          `json:"user-tag"`
	PolicyClassListList []DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                     `json:"class-list-name"`
	ClassListGlid                  string                                                                                     `json:"class-list-glid"`
	Glid                           string                                                                                     `json:"glid"`
	GlidAction                     string                                                                                     `json:"glid-action"`
	Action                         string                                                                                     `json:"action"`
	MaxDynamicEntryCount           int                                                                                        `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                        `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                     `json:"uuid"`
	UserTag                        string                                                                                     `json:"user-tag"`
	SamplingEnable                 []DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
	Quic    string `json:"quic"`
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                               `json:"dummy-name"`
	Glid         string                                                                                               `json:"glid"`
	Action       string                                                                                               `json:"action"`
	LogEnable    int                                                                                                  `json:"log-enable"`
	LogPeriodic  int                                                                                                  `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                               `json:"uuid"`
	UserTag      string                                                                                               `json:"user-tag"`
}

type DdosDstZonePortRangeListSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
	Quic    string `json:"quic"`
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortRangeListDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                             `json:"dummy-name"`
	Glid         string                                                             `json:"glid"`
	Action       string                                                             `json:"action"`
	LogEnable    int                                                                `json:"log-enable"`
	LogPeriodic  int                                                                `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortRangeListDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                             `json:"uuid"`
	UserTag      string                                                             `json:"user-tag"`
}

type DdosDstZonePortRangeListDynamicEntryOverflowPolicyListZoneTemplate struct {
	Quic    string `json:"quic"`
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortRangeListVirtualhosts struct {
	VhostsConfig      string                                                `json:"vhosts-config"`
	SourceTrackingAll int                                                   `json:"source-tracking-all"`
	Uuid              string                                                `json:"uuid"`
	VirtualhostList   []DdosDstZonePortRangeListVirtualhostsVirtualhostList `json:"virtualhost-list"`
}

type DdosDstZonePortRangeListVirtualhostsVirtualhostList struct {
	Vhost              string                                                          `json:"vhost"`
	Servername         []DdosDstZonePortRangeListVirtualhostsVirtualhostListServername `json:"servername"`
	ServernameList     string                                                          `json:"servername-list"`
	ServernameMatchAny int                                                             `json:"servername-match-any"`
	SourceTracking     string                                                          `json:"source-tracking" dval:"follow"`
	GlidCfg            DdosDstZonePortRangeListVirtualhostsVirtualhostListGlidCfg      `json:"glid-cfg"`
	Deny               int                                                             `json:"deny"`
	Uuid               string                                                          `json:"uuid"`
	UserTag            string                                                          `json:"user-tag"`
	LevelList          []DdosDstZonePortRangeListVirtualhostsVirtualhostListLevelList  `json:"level-list"`
}

type DdosDstZonePortRangeListVirtualhostsVirtualhostListServername struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

type DdosDstZonePortRangeListVirtualhostsVirtualhostListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortRangeListVirtualhostsVirtualhostListLevelList struct {
	LevelNum       string                                                                   `json:"level-num"`
	SrcDefaultGlid string                                                                   `json:"src-default-glid"`
	GlidAction     string                                                                   `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortRangeListVirtualhostsVirtualhostListLevelListZoneTemplate `json:"zone-template"`
	Uuid           string                                                                   `json:"uuid"`
	UserTag        string                                                                   `json:"user-tag"`
}

type DdosDstZonePortRangeListVirtualhostsVirtualhostListLevelListZoneTemplate struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

type DdosDstZoneSamplingEnable struct {
	Counters1 string `json:"counters1"`
	Counters2 string `json:"counters2"`
	Counters3 string `json:"counters3"`
}

type DdosDstZoneSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
	SflowSampSession int `json:"sflow-samp-session"`
}

type DdosDstZoneSrcPort295 struct {
	ZoneSrcPortList      []DdosDstZoneSrcPortZoneSrcPortList      `json:"zone-src-port-list"`
	ZoneSrcPortOtherList []DdosDstZoneSrcPortZoneSrcPortOtherList `json:"zone-src-port-other-list"`
}

type DdosDstZoneSrcPortZoneSrcPortList struct {
	PortNum             int                                           `json:"port-num"`
	Protocol            string                                        `json:"protocol"`
	Deny                int                                           `json:"deny"`
	GlidCfg             DdosDstZoneSrcPortZoneSrcPortListGlidCfg      `json:"glid-cfg"`
	OutboundSrcTracking string                                        `json:"outbound-src-tracking" dval:"disable"`
	ZoneTemplate        DdosDstZoneSrcPortZoneSrcPortListZoneTemplate `json:"zone-template"`
	DefaultActionList   string                                        `json:"default-action-list"`
	SetCounterBaseVal   int                                           `json:"set-counter-base-val"`
	Uuid                string                                        `json:"uuid"`
	PortInd             DdosDstZoneSrcPortZoneSrcPortListPortInd      `json:"port-ind"`
	LevelList           []DdosDstZoneSrcPortZoneSrcPortListLevelList  `json:"level-list"`
}

type DdosDstZoneSrcPortZoneSrcPortListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZoneSrcPortZoneSrcPortListZoneTemplate struct {
	SrcUdp string `json:"src-udp"`
	SrcTcp string `json:"src-tcp"`
	SrcDns string `json:"src-dns"`
}

type DdosDstZoneSrcPortZoneSrcPortListPortInd struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneSrcPortZoneSrcPortListLevelList struct {
	LevelNum      string                                                    `json:"level-num"`
	Uuid          string                                                    `json:"uuid"`
	UserTag       string                                                    `json:"user-tag"`
	IndicatorList []DdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZoneSrcPortZoneSrcPortListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZoneSrcPortZoneSrcPortOtherList struct {
	PortOther         string                                             `json:"port-other"`
	Protocol          string                                             `json:"protocol"`
	Deny              int                                                `json:"deny"`
	GlidCfg           DdosDstZoneSrcPortZoneSrcPortOtherListGlidCfg      `json:"glid-cfg"`
	ZoneTemplate      DdosDstZoneSrcPortZoneSrcPortOtherListZoneTemplate `json:"zone-template"`
	DefaultActionList string                                             `json:"default-action-list"`
	SetCounterBaseVal int                                                `json:"set-counter-base-val"`
	Uuid              string                                             `json:"uuid"`
	PortInd           DdosDstZoneSrcPortZoneSrcPortOtherListPortInd      `json:"port-ind"`
	LevelList         []DdosDstZoneSrcPortZoneSrcPortOtherListLevelList  `json:"level-list"`
}

type DdosDstZoneSrcPortZoneSrcPortOtherListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZoneSrcPortZoneSrcPortOtherListZoneTemplate struct {
	SrcUdp string `json:"src-udp"`
	SrcTcp string `json:"src-tcp"`
}

type DdosDstZoneSrcPortZoneSrcPortOtherListPortInd struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneSrcPortZoneSrcPortOtherListLevelList struct {
	LevelNum      string                                                         `json:"level-num"`
	Uuid          string                                                         `json:"uuid"`
	UserTag       string                                                         `json:"user-tag"`
	IndicatorList []DdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZoneSrcPortZoneSrcPortOtherListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZoneSrcPortRangeList struct {
	SrcPortRangeStart int                                      `json:"src-port-range-start"`
	SrcPortRangeEnd   int                                      `json:"src-port-range-end"`
	Protocol          string                                   `json:"protocol"`
	Deny              int                                      `json:"deny"`
	GlidCfg           DdosDstZoneSrcPortRangeListGlidCfg       `json:"glid-cfg"`
	ZoneTemplate      DdosDstZoneSrcPortRangeListZoneTemplate  `json:"zone-template"`
	DefaultActionList string                                   `json:"default-action-list"`
	CaptureConfig     DdosDstZoneSrcPortRangeListCaptureConfig `json:"capture-config"`
	SetCounterBaseVal int                                      `json:"set-counter-base-val"`
	Uuid              string                                   `json:"uuid"`
	UserTag           string                                   `json:"user-tag"`
	PortInd           DdosDstZoneSrcPortRangeListPortInd       `json:"port-ind"`
	LevelList         []DdosDstZoneSrcPortRangeListLevelList   `json:"level-list"`
}

type DdosDstZoneSrcPortRangeListGlidCfg struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZoneSrcPortRangeListZoneTemplate struct {
	SrcUdp string `json:"src-udp"`
	SrcTcp string `json:"src-tcp"`
}

type DdosDstZoneSrcPortRangeListCaptureConfig struct {
	CaptureConfigName string `json:"capture-config-name"`
	CaptureConfigMode string `json:"capture-config-mode"`
}

type DdosDstZoneSrcPortRangeListPortInd struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneSrcPortRangeListLevelList struct {
	LevelNum      string                                              `json:"level-num"`
	Uuid          string                                              `json:"uuid"`
	UserTag       string                                              `json:"user-tag"`
	IndicatorList []DdosDstZoneSrcPortRangeListLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZoneSrcPortRangeListLevelListIndicatorList struct {
	Type                  string `json:"type"`
	ZoneThresholdNum      int    `json:"zone-threshold-num"`
	ZoneThresholdLargeNum int    `json:"zone-threshold-large-num"`
	Uuid                  string `json:"uuid"`
	UserTag               string `json:"user-tag"`
}

type DdosDstZoneTopkDestinations296 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZoneWebGui297 struct {
	Status                 string                         `json:"status" dval:"newly"`
	ActivatedAfterLearning int                            `json:"activated-after-learning"`
	CreateTime             string                         `json:"create-time"`
	ModifyTime             string                         `json:"modify-time"`
	Sensitivity            string                         `json:"sensitivity" dval:"3"`
	Uuid                   string                         `json:"uuid"`
	Learning               DdosDstZoneWebGuiLearning298   `json:"learning"`
	Protection             DdosDstZoneWebGuiProtection299 `json:"protection"`
}

type DdosDstZoneWebGuiLearning298 struct {
	Duration     string `json:"duration" dval:"6hour"`
	StartingTime string `json:"starting-time"`
	Uuid         string `json:"uuid"`
}

type DdosDstZoneWebGuiProtection299 struct {
	Port          DdosDstZoneWebGuiProtectionPort300            `json:"port"`
	IpProto       DdosDstZoneWebGuiProtectionIpProto303         `json:"ip-proto"`
	PortRangeList []DdosDstZoneWebGuiProtectionPortRangeList305 `json:"port-range-list"`
}

type DdosDstZoneWebGuiProtectionPort300 struct {
	ZoneServiceList      []DdosDstZoneWebGuiProtectionPortZoneServiceList301      `json:"zone-service-list"`
	ZoneServiceOtherList []DdosDstZoneWebGuiProtectionPortZoneServiceOtherList302 `json:"zone-service-other-list"`
}

type DdosDstZoneWebGuiProtectionPortZoneServiceList301 struct {
	PortNum  int    `json:"port-num"`
	Protocol string `json:"protocol"`
	Pbe      string `json:"pbe"`
	Uuid     string `json:"uuid"`
}

type DdosDstZoneWebGuiProtectionPortZoneServiceOtherList302 struct {
	PortOther string `json:"port-other"`
	Protocol  string `json:"protocol"`
	Pbe       string `json:"pbe"`
	Uuid      string `json:"uuid"`
}

type DdosDstZoneWebGuiProtectionIpProto303 struct {
	ProtoNameList []DdosDstZoneWebGuiProtectionIpProtoProtoNameList304 `json:"proto-name-list"`
}

type DdosDstZoneWebGuiProtectionIpProtoProtoNameList304 struct {
	Protocol string `json:"protocol"`
	Pbe      string `json:"pbe"`
	Uuid     string `json:"uuid"`
	UserTag  string `json:"user-tag"`
}

type DdosDstZoneWebGuiProtectionPortRangeList305 struct {
	PortRangeStart int    `json:"port-range-start"`
	PortRangeEnd   int    `json:"port-range-end"`
	Protocol       string `json:"protocol"`
	Pbe            string `json:"pbe"`
	Uuid           string `json:"uuid"`
	UserTag        string `json:"user-tag"`
}

type DdosDstZoneZoneTemplate struct {
	Logging string `json:"logging"`
}

func (p *DdosDstZone) GetId() string {
	return url.QueryEscape(p.Inst.ZoneName)
}

func (p *DdosDstZone) getPath() string {
	return "ddos/dst/zone"
}

func (p *DdosDstZone) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZone::Post")
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

func (p *DdosDstZone) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZone::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosDstZone) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZone::Put")
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

func (p *DdosDstZone) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZone::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
