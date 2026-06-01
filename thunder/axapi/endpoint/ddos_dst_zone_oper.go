package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZoneOper struct {
	Detection DdosDstZoneOperDetection `json:"detection"`

	IpProto DdosDstZoneOperIpProto `json:"ip-proto"`

	Oper DdosDstZoneOperOper `json:"oper"`

	OutboundPolicy DdosDstZoneOperOutboundPolicy `json:"outbound-policy"`

	PacketAnomalyDetection DdosDstZoneOperPacketAnomalyDetection `json:"packet-anomaly-detection"`

	Port DdosDstZoneOperPort `json:"port"`

	PortRangeList []DdosDstZoneOperPortRangeList `json:"port-range-list"`

	SrcIpFiltering DdosDstZoneOperSrcIpFiltering `json:"src-ip-filtering"`

	SrcPort DdosDstZoneOperSrcPort `json:"src-port"`

	SrcPortRangeList []DdosDstZoneOperSrcPortRangeList `json:"src-port-range-list"`

	TopkDestinations DdosDstZoneOperTopkDestinations `json:"topk-destinations"`

	ZoneName string `json:"zone-name"`
}
type DataDdosDstZoneOper struct {
	DtDdosDstZoneOper DdosDstZoneOper `json:"zone"`
}

type DdosDstZoneOperDetection struct {
	Oper              DdosDstZoneOperDetectionOper              `json:"oper"`
	OutboundDetection DdosDstZoneOperDetectionOutboundDetection `json:"outbound-detection"`
	ServiceDiscovery  DdosDstZoneOperDetectionServiceDiscovery  `json:"service-discovery"`
	VictimIpDetection DdosDstZoneOperDetectionVictimIpDetection `json:"victim-ip-detection"`
}

type DdosDstZoneOperDetectionOper struct {
}

type DdosDstZoneOperDetectionOutboundDetection struct {
	Oper             DdosDstZoneOperDetectionOutboundDetectionOper             `json:"oper"`
	TopkSourceSubnet DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet `json:"topk-source-subnet"`
}

type DdosDstZoneOperDetectionOutboundDetectionOper struct {
	DiscoveryTimestamp string                                                   `json:"discovery-timestamp"`
	EntryList          []DdosDstZoneOperDetectionOutboundDetectionOperEntryList `json:"entry-list"`
}

type DdosDstZoneOperDetectionOutboundDetectionOperEntryList struct {
	LocationType     string                                                             `json:"location-type"`
	LocationName     string                                                             `json:"location-name"`
	Indicators       []DdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators `json:"indicators"`
	DataSource       string                                                             `json:"data-source"`
	Anomaly          string                                                             `json:"anomaly"`
	AnomalyTimestamp string                                                             `json:"anomaly-timestamp"`
	InitialLearning  string                                                             `json:"initial-learning"`
	ActiveTime       int                                                                `json:"active-time"`
}

type DdosDstZoneOperDetectionOutboundDetectionOperEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	NonZeroMinimum    string `json:"non-zero-minimum"`
	Average           string `json:"average"`
	AdaptiveThreshold string `json:"adaptive-threshold"`
}

type DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnet struct {
	Oper DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper `json:"oper"`
}

type DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOper struct {
	EntryList []DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList `json:"entry-list"`
}

type DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryList struct {
	LocationType string                                                                             `json:"location-type"`
	LocationName string                                                                             `json:"location-name"`
	Indicators   []DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicators struct {
	IndicatorName  string                                                                                          `json:"indicator-name"`
	IndicatorIndex int                                                                                             `json:"indicator-index"`
	SourceSubnets  []DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets `json:"source-subnets"`
}

type DdosDstZoneOperDetectionOutboundDetectionTopkSourceSubnetOperEntryListIndicatorsSourceSubnets struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperDetectionServiceDiscovery struct {
	Oper DdosDstZoneOperDetectionServiceDiscoveryOper `json:"oper"`
}

type DdosDstZoneOperDetectionServiceDiscoveryOper struct {
	DiscoveredServiceList []DdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList `json:"discovered-service-list"`
}

type DdosDstZoneOperDetectionServiceDiscoveryOperDiscoveredServiceList struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Rate     int    `json:"rate"`
}

type DdosDstZoneOperDetectionVictimIpDetection struct {
	Oper DdosDstZoneOperDetectionVictimIpDetectionOper `json:"oper"`
}

type DdosDstZoneOperDetectionVictimIpDetectionOper struct {
	IpEntryList         []DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList `json:"ip-entry-list"`
	IpEntryCount        int                                                        `json:"ip-entry-count"`
	TotalIpEntryCount   int                                                        `json:"total-ip-entry-count"`
	WithSelectedDetails int                                                        `json:"with-selected-details"`
	ActiveList          int                                                        `json:"active-list"`
	VictimList          int                                                        `json:"victim-list"`
	Ipv4Ip              string                                                     `json:"ipv4-ip"`
	Ipv6Ip              string                                                     `json:"ipv6-ip"`
}

type DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryList struct {
	IpAddressStr            string                                                               `json:"ip-address-str"`
	Indicators              []DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators `json:"indicators"`
	IsLearningDone          int                                                                  `json:"is-learning-done"`
	IsHistogramLearningDone int                                                                  `json:"is-histogram-learning-done"`
	IsIpAnomaly             int                                                                  `json:"is-ip-anomaly"`
	Is_static_threshold     int                                                                  `json:"is_static_threshold"`
	EscalationTimestamp     string                                                               `json:"escalation-timestamp"`
	DeEscalationTimestamp   string                                                               `json:"de-escalation-timestamp"`
}

type DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicators struct {
	IndicatorName  string                                                                    `json:"indicator-name"`
	IndicatorIndex int                                                                       `json:"indicator-index"`
	Value          []DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue `json:"value"`
	IsAnomaly      int                                                                       `json:"is-anomaly"`
}

type DdosDstZoneOperDetectionVictimIpDetectionOperIpEntryListIndicatorsValue struct {
	Current   string `json:"current"`
	Threshold string `json:"threshold"`
}

type DdosDstZoneOperIpProto struct {
	Oper            DdosDstZoneOperIpProtoOper              `json:"oper"`
	ProtoNumberList []DdosDstZoneOperIpProtoProtoNumberList `json:"proto-number-list"`
	ProtoTcpUdpList []DdosDstZoneOperIpProtoProtoTcpUdpList `json:"proto-tcp-udp-list"`
	ProtoNameList   []DdosDstZoneOperIpProtoProtoNameList   `json:"proto-name-list"`
}

type DdosDstZoneOperIpProtoOper struct {
}

type DdosDstZoneOperIpProtoProtoNumberList struct {
	ProtocolNum                 int                                                              `json:"protocol-num"`
	Oper                        DdosDstZoneOperIpProtoProtoNumberListOper                        `json:"oper"`
	IpFilteringPolicyStatistics DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
	SrcBasedPolicyList          []DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyList        `json:"src-based-policy-list"`
	PortInd                     DdosDstZoneOperIpProtoProtoNumberListPortInd                     `json:"port-ind"`
	TopkSources                 DdosDstZoneOperIpProtoProtoNumberListTopkSources                 `json:"topk-sources"`
	TopkDestinations            DdosDstZoneOperIpProtoProtoNumberListTopkDestinations            `json:"topk-destinations"`
	ProgressionTracking         DdosDstZoneOperIpProtoProtoNumberListProgressionTracking         `json:"progression-tracking"`
}

type DdosDstZoneOperIpProtoProtoNumberListOper struct {
	Ddos_entry_list       []DdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                        `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                        `json:"service-displayed-count"`
	ReportingStatus       int                                                        `json:"reporting-status"`
	Sources               int                                                        `json:"sources"`
	OverflowPolicy        int                                                        `json:"overflow-policy"`
	SourcesAllEntries     int                                                        `json:"sources-all-entries"`
	ClassList             string                                                     `json:"class-list"`
	SubnetIpAddr          string                                                     `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                     `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                     `json:"ipv6"`
	Exceeded              int                                                        `json:"exceeded"`
	BlackListed           int                                                        `json:"black-listed"`
	WhiteListed           int                                                        `json:"white-listed"`
	Authenticated         int                                                        `json:"authenticated"`
	Level                 int                                                        `json:"level"`
	AppStat               int                                                        `json:"app-stat"`
	Indicators            int                                                        `json:"indicators"`
	IndicatorDetail       int                                                        `json:"indicator-detail"`
	HwBlacklisted         int                                                        `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                        `json:"suffix-request-rate"`
	DomainName            string                                                     `json:"domain-name"`
}

type DdosDstZoneOperIpProtoProtoNumberListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	BlReasoningRcode       string `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string `json:"bl-reasoning-timestamp"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyStatistics struct {
	Oper DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZoneOperIpProtoProtoNumberListIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                       `json:"src-based-policy-name"`
	Oper                DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyListOper                  `json:"oper"`
	PolicyClassListList []DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyListOper struct {
}

type DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName string                                                                         `json:"class-list-name"`
	Oper          DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNumberListSrcBasedPolicyListPolicyClassListListOper struct {
	CurrentConnections     int    `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        int    `json:"connection-limit"`
	CurrentConnectionRate  int    `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    int    `json:"connection-rate-limit"`
	CurrentPacketRate      int    `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        int    `json:"packet-rate-limit"`
	CurrentKbitRate        int    `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          int    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  int    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    int    `json:"frag-packet-rate-limit"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperIpProtoProtoNumberListPortInd struct {
	Oper DdosDstZoneOperIpProtoProtoNumberListPortIndOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNumberListPortIndOper struct {
	SrcEntryList        []DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList `json:"src-entry-list"`
	Indicators          []DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators   `json:"indicators"`
	DetectionDataSource string                                                         `json:"detection-data-source"`
	TotalScore          string                                                         `json:"total-score"`
	CurrentLevel        string                                                         `json:"current-level"`
	EscalationTimestamp string                                                         `json:"escalation-timestamp"`
	InitialLearning     string                                                         `json:"initial-learning"`
	ActiveTime          int                                                            `json:"active-time"`
	BaselineWindowSize  int                                                            `json:"baseline-window-size"`
	SourcesAllEntries   int                                                            `json:"sources-all-entries"`
	SubnetIpAddr        string                                                         `json:"subnet-ip-addr"`
	SubnetIpv6Addr      string                                                         `json:"subnet-ipv6-addr"`
	Ipv6                string                                                         `json:"ipv6"`
	Details             int                                                            `json:"details"`
	Sources             int                                                            `json:"sources"`
}

type DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryList struct {
	SrcAddressStr       string                                                                   `json:"src-address-str"`
	Indicators          []DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators `json:"indicators"`
	DetectionDataSource string                                                                   `json:"detection-data-source"`
	TotalScore          string                                                                   `json:"total-score"`
	CurrentLevel        string                                                                   `json:"current-level"`
	SrcLevel            string                                                                   `json:"src-level"`
	EscalationTimestamp string                                                                   `json:"escalation-timestamp"`
	InitialLearning     string                                                                   `json:"initial-learning"`
	ActiveTime          int                                                                      `json:"active-time"`
}

type DdosDstZoneOperIpProtoProtoNumberListPortIndOperSrcEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	SrcMaximum        string `json:"src-maximum"`
	SrcMinimum        string `json:"src-minimum"`
	SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
	SrcAverage        string `json:"src-average"`
	Score             string `json:"score"`
}

type DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicators struct {
	IndicatorName         string                                                                   `json:"indicator-name"`
	IndicatorIndex        int                                                                      `json:"indicator-index"`
	Rate                  string                                                                   `json:"rate"`
	ZoneMaximum           string                                                                   `json:"zone-maximum"`
	ZoneMinimum           string                                                                   `json:"zone-minimum"`
	ZoneNonZeroMinimum    string                                                                   `json:"zone-non-zero-minimum"`
	ZoneAverage           string                                                                   `json:"zone-average"`
	ZoneAdaptiveThreshold string                                                                   `json:"zone-adaptive-threshold"`
	SrcMaximum            string                                                                   `json:"src-maximum"`
	IndicatorCfg          []DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
	Score                 string                                                                   `json:"score"`
}

type DdosDstZoneOperIpProtoProtoNumberListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkSources struct {
	Oper DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOper struct {
	Indicators    []DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList  `json:"entry-list"`
	NextIndicator int                                                              `json:"next-indicator"`
	Finished      int                                                              `json:"finished"`
	Details       int                                                              `json:"details"`
	TopKKey       string                                                           `json:"top-k-key"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicators struct {
	IndicatorName  string                                                                  `json:"indicator-name"`
	IndicatorIndex int                                                                     `json:"indicator-index"`
	Sources        []DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources `json:"sources"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperIndicatorsSources struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryList struct {
	AddressStr string                                                                    `json:"address-str"`
	Indicators []DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkSourcesOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkDestinations struct {
	Oper DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOper struct {
	Indicators    []DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                                                   `json:"next-indicator"`
	Finished      int                                                                   `json:"finished"`
	Details       int                                                                   `json:"details"`
	TopKKey       string                                                                `json:"top-k-key"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                                            `json:"indicator-name"`
	IndicatorIndex int                                                                               `json:"indicator-index"`
	Destinations   []DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryList struct {
	AddressStr string                                                                         `json:"address-str"`
	Indicators []DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperIpProtoProtoNumberListTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperIpProtoProtoNumberListProgressionTracking struct {
	Oper DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOper struct {
	Indicators          []DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators `json:"indicators"`
	LearningDetails     int                                                                      `json:"learning-details"`
	LearningBrief       int                                                                      `json:"learning-brief"`
	RecommendedTemplate int                                                                      `json:"recommended-template"`
	TemplateDebugTable  int                                                                      `json:"template-debug-table"`
}

type DdosDstZoneOperIpProtoProtoNumberListProgressionTrackingOperIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	NumSample         int    `json:"num-sample"`
	Average           string `json:"average"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	StandardDeviation string `json:"standard-deviation"`
}

type DdosDstZoneOperIpProtoProtoTcpUdpList struct {
	Protocol                    string                                                           `json:"protocol"`
	Oper                        DdosDstZoneOperIpProtoProtoTcpUdpListOper                        `json:"oper"`
	IpFilteringPolicyStatistics DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
}

type DdosDstZoneOperIpProtoProtoTcpUdpListOper struct {
	Ddos_entry_list       []DdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                        `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                        `json:"service-displayed-count"`
	ReportingStatus       int                                                        `json:"reporting-status"`
	Sources               int                                                        `json:"sources"`
	OverflowPolicy        int                                                        `json:"overflow-policy"`
	SourcesAllEntries     int                                                        `json:"sources-all-entries"`
	ClassList             string                                                     `json:"class-list"`
	SubnetIpAddr          string                                                     `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                     `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                     `json:"ipv6"`
	Exceeded              int                                                        `json:"exceeded"`
	BlackListed           int                                                        `json:"black-listed"`
	WhiteListed           int                                                        `json:"white-listed"`
	Authenticated         int                                                        `json:"authenticated"`
	Level                 int                                                        `json:"level"`
	AppStat               int                                                        `json:"app-stat"`
	Indicators            int                                                        `json:"indicators"`
	IndicatorDetail       int                                                        `json:"indicator-detail"`
	HwBlacklisted         int                                                        `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                        `json:"suffix-request-rate"`
	DomainName            string                                                     `json:"domain-name"`
}

type DdosDstZoneOperIpProtoProtoTcpUdpListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	BlReasoningRcode       string `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string `json:"bl-reasoning-timestamp"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyStatistics struct {
	Oper DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZoneOperIpProtoProtoTcpUdpListIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZoneOperIpProtoProtoNameList struct {
	Protocol                    string                                                         `json:"protocol"`
	Oper                        DdosDstZoneOperIpProtoProtoNameListOper                        `json:"oper"`
	IpFilteringPolicyStatistics DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
	SrcBasedPolicyList          []DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyList        `json:"src-based-policy-list"`
	PortInd                     DdosDstZoneOperIpProtoProtoNameListPortInd                     `json:"port-ind"`
	TopkSources                 DdosDstZoneOperIpProtoProtoNameListTopkSources                 `json:"topk-sources"`
	ProgressionTracking         DdosDstZoneOperIpProtoProtoNameListProgressionTracking         `json:"progression-tracking"`
	TopkDestinations            DdosDstZoneOperIpProtoProtoNameListTopkDestinations            `json:"topk-destinations"`
}

type DdosDstZoneOperIpProtoProtoNameListOper struct {
	Ddos_entry_list       []DdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                      `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                      `json:"service-displayed-count"`
	ReportingStatus       int                                                      `json:"reporting-status"`
	Sources               int                                                      `json:"sources"`
	OverflowPolicy        int                                                      `json:"overflow-policy"`
	SourcesAllEntries     int                                                      `json:"sources-all-entries"`
	ClassList             string                                                   `json:"class-list"`
	SubnetIpAddr          string                                                   `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                   `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                   `json:"ipv6"`
	Exceeded              int                                                      `json:"exceeded"`
	BlackListed           int                                                      `json:"black-listed"`
	WhiteListed           int                                                      `json:"white-listed"`
	Authenticated         int                                                      `json:"authenticated"`
	Level                 int                                                      `json:"level"`
	AppStat               int                                                      `json:"app-stat"`
	Indicators            int                                                      `json:"indicators"`
	IndicatorDetail       int                                                      `json:"indicator-detail"`
	HwBlacklisted         int                                                      `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                      `json:"suffix-request-rate"`
	DomainName            string                                                   `json:"domain-name"`
}

type DdosDstZoneOperIpProtoProtoNameListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	BlReasoningRcode       string `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string `json:"bl-reasoning-timestamp"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyStatistics struct {
	Oper DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZoneOperIpProtoProtoNameListIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                     `json:"src-based-policy-name"`
	Oper                DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyListOper                  `json:"oper"`
	PolicyClassListList []DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyListOper struct {
}

type DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName string                                                                       `json:"class-list-name"`
	Oper          DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNameListSrcBasedPolicyListPolicyClassListListOper struct {
	CurrentConnections     int    `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        int    `json:"connection-limit"`
	CurrentConnectionRate  int    `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    int    `json:"connection-rate-limit"`
	CurrentPacketRate      int    `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        int    `json:"packet-rate-limit"`
	CurrentKbitRate        int    `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          int    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  int    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    int    `json:"frag-packet-rate-limit"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperIpProtoProtoNameListPortInd struct {
	Oper DdosDstZoneOperIpProtoProtoNameListPortIndOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNameListPortIndOper struct {
	SrcEntryList        []DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList `json:"src-entry-list"`
	Indicators          []DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators   `json:"indicators"`
	DetectionDataSource string                                                       `json:"detection-data-source"`
	TotalScore          string                                                       `json:"total-score"`
	CurrentLevel        string                                                       `json:"current-level"`
	EscalationTimestamp string                                                       `json:"escalation-timestamp"`
	InitialLearning     string                                                       `json:"initial-learning"`
	ActiveTime          int                                                          `json:"active-time"`
	BaselineWindowSize  int                                                          `json:"baseline-window-size"`
	SourcesAllEntries   int                                                          `json:"sources-all-entries"`
	SubnetIpAddr        string                                                       `json:"subnet-ip-addr"`
	SubnetIpv6Addr      string                                                       `json:"subnet-ipv6-addr"`
	Ipv6                string                                                       `json:"ipv6"`
	Details             int                                                          `json:"details"`
	Sources             int                                                          `json:"sources"`
}

type DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryList struct {
	SrcAddressStr       string                                                                 `json:"src-address-str"`
	Indicators          []DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators `json:"indicators"`
	DetectionDataSource string                                                                 `json:"detection-data-source"`
	TotalScore          string                                                                 `json:"total-score"`
	CurrentLevel        string                                                                 `json:"current-level"`
	SrcLevel            string                                                                 `json:"src-level"`
	EscalationTimestamp string                                                                 `json:"escalation-timestamp"`
	InitialLearning     string                                                                 `json:"initial-learning"`
	ActiveTime          int                                                                    `json:"active-time"`
}

type DdosDstZoneOperIpProtoProtoNameListPortIndOperSrcEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	SrcMaximum        string `json:"src-maximum"`
	SrcMinimum        string `json:"src-minimum"`
	SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
	SrcAverage        string `json:"src-average"`
	Score             string `json:"score"`
}

type DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicators struct {
	IndicatorName         string                                                                 `json:"indicator-name"`
	IndicatorIndex        int                                                                    `json:"indicator-index"`
	Rate                  string                                                                 `json:"rate"`
	ZoneMaximum           string                                                                 `json:"zone-maximum"`
	ZoneMinimum           string                                                                 `json:"zone-minimum"`
	ZoneNonZeroMinimum    string                                                                 `json:"zone-non-zero-minimum"`
	ZoneAverage           string                                                                 `json:"zone-average"`
	ZoneAdaptiveThreshold string                                                                 `json:"zone-adaptive-threshold"`
	SrcMaximum            string                                                                 `json:"src-maximum"`
	IndicatorCfg          []DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
	Score                 string                                                                 `json:"score"`
}

type DdosDstZoneOperIpProtoProtoNameListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkSources struct {
	Oper DdosDstZoneOperIpProtoProtoNameListTopkSourcesOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkSourcesOper struct {
	Indicators    []DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList  `json:"entry-list"`
	NextIndicator int                                                            `json:"next-indicator"`
	Finished      int                                                            `json:"finished"`
	Details       int                                                            `json:"details"`
	TopKKey       string                                                         `json:"top-k-key"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicators struct {
	IndicatorName  string                                                                `json:"indicator-name"`
	IndicatorIndex int                                                                   `json:"indicator-index"`
	Sources        []DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources `json:"sources"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperIndicatorsSources struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryList struct {
	AddressStr string                                                                  `json:"address-str"`
	Indicators []DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkSourcesOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperIpProtoProtoNameListProgressionTracking struct {
	Oper DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOper struct {
	Indicators          []DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators `json:"indicators"`
	LearningDetails     int                                                                    `json:"learning-details"`
	LearningBrief       int                                                                    `json:"learning-brief"`
	RecommendedTemplate int                                                                    `json:"recommended-template"`
	TemplateDebugTable  int                                                                    `json:"template-debug-table"`
}

type DdosDstZoneOperIpProtoProtoNameListProgressionTrackingOperIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	NumSample         int    `json:"num-sample"`
	Average           string `json:"average"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	StandardDeviation string `json:"standard-deviation"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkDestinations struct {
	Oper DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper `json:"oper"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOper struct {
	Indicators    []DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                                                 `json:"next-indicator"`
	Finished      int                                                                 `json:"finished"`
	Details       int                                                                 `json:"details"`
	TopKKey       string                                                              `json:"top-k-key"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                                          `json:"indicator-name"`
	IndicatorIndex int                                                                             `json:"indicator-index"`
	Destinations   []DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryList struct {
	AddressStr string                                                                       `json:"address-str"`
	Indicators []DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperIpProtoProtoNameListTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperOper struct {
	Ddos_entry_list           []DdosDstZoneOperOperDdos_entry_list           `json:"ddos_entry_list"`
	TotalDynamicEntryCount    string                                         `json:"total-dynamic-entry-count"`
	UdpDynamicEntryCount      string                                         `json:"udp-dynamic-entry-count"`
	TcpDynamicEntryCount      string                                         `json:"tcp-dynamic-entry-count"`
	IcmpDynamicEntryCount     string                                         `json:"icmp-dynamic-entry-count"`
	OtherDynamicEntryCount    string                                         `json:"other-dynamic-entry-count"`
	TrafficDistributionStatus []DdosDstZoneOperOperTrafficDistributionStatus `json:"traffic-distribution-status"`
	EntryDisplayedCount       int                                            `json:"entry-displayed-count"`
	ServiceDisplayedCount     int                                            `json:"service-displayed-count"`
	NoT2IdxPortCount          int                                            `json:"no-t2-idx-port-count"`
	Addresses                 int                                            `json:"addresses"`
	SubnetIpAddr              string                                         `json:"subnet-ip-addr"`
	SubnetIpv6Addr            string                                         `json:"subnet-ipv6-addr"`
	AllAddresses              int                                            `json:"all-addresses"`
	IpProtoNum                int                                            `json:"ip-proto-num"`
	AllIpProtos               int                                            `json:"all-ip-protos"`
	PortNum                   int                                            `json:"port-num"`
	PortRangeStart            int                                            `json:"port-range-start"`
	PortRangeEnd              int                                            `json:"port-range-end"`
	Protocol                  string                                         `json:"protocol"`
	AllPorts                  int                                            `json:"all-ports"`
	DynamicExpandSubnet       int                                            `json:"dynamic-expand-subnet"`
	Blackhole                 int                                            `json:"blackhole"`
}

type DdosDstZoneOperOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	PortStr                string `json:"port-str"`
	OperationalMode        string `json:"operational-mode"`
	BwState                string `json:"bw-state"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	AgeStr                 string `json:"age-str"`
	LockupTime             int    `json:"lockup-time"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperOperTrafficDistributionStatus struct {
	MasterPu string                                                 `json:"master-pu"`
	ActivePu []DdosDstZoneOperOperTrafficDistributionStatusActivePu `json:"active-pu"`
}

type DdosDstZoneOperOperTrafficDistributionStatusActivePu struct {
	PuId string `json:"pu-id"`
}

type DdosDstZoneOperOutboundPolicy struct {
	Oper DdosDstZoneOperOutboundPolicyOper `json:"oper"`
}

type DdosDstZoneOperOutboundPolicyOper struct {
	PolicyName            string                                                 `json:"policy-name"`
	NoClassListMatch      int                                                    `json:"no-class-list-match"`
	PolicyClassList       []DdosDstZoneOperOutboundPolicyOperPolicyClassList     `json:"policy-class-list"`
	GeoTrackingStatistics DdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics `json:"geo-tracking-statistics"`
	TrackingEntryList     []DdosDstZoneOperOutboundPolicyOperTrackingEntryList   `json:"tracking-entry-list"`
	PolicyRate            int                                                    `json:"policy-rate"`
	PolicyStatistics      int                                                    `json:"policy-statistics"`
	TrackingEntryFilter   int                                                    `json:"tracking-entry-filter"`
}

type DdosDstZoneOperOutboundPolicyOperPolicyClassList struct {
	ClassListName          string `json:"class-list-name"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	AgeStr                 string `json:"age-str"`
	LockupTime             int    `json:"lockup-time"`
	PacketReceived         int    `json:"packet-received"`
	PacketDropped          int    `json:"packet-dropped"`
	PacketRateExceed       int    `json:"packet-rate-exceed"`
	KbitRateExceed         int    `json:"kBit-rate-exceed"`
	KbitRateExceedCount    int    `json:"kBit-rate-exceed-count"`
	ConnectionsExceed      int    `json:"connections-exceed"`
	ConnectionRateExceed   int    `json:"connection-rate-exceed"`
	FragPacketRate         int    `json:"frag-packet-rate"`
}

type DdosDstZoneOperOutboundPolicyOperGeoTrackingStatistics struct {
	PacketReceived                  int `json:"packet-received"`
	PacketDropped                   int `json:"packet-dropped"`
	PacketRateExceed                int `json:"packet-rate-exceed"`
	KbitRateExceed                  int `json:"kBit-rate-exceed"`
	KbitRateExceedCount             int `json:"kBit-rate-exceed-count"`
	ConnectionsExceed               int `json:"connections-exceed"`
	ConnectionRateExceed            int `json:"connection-rate-exceed"`
	FragPacketRate                  int `json:"frag-packet-rate"`
	TrackingEntryLearn              int `json:"tracking-entry-learn"`
	TrackingEntryAged               int `json:"tracking-entry-aged"`
	TrackingEntryLearningThreExceed int `json:"tracking-entry-learning-thre-exceed"`
}

type DdosDstZoneOperOutboundPolicyOperTrackingEntryList struct {
	GeoLocationName        string `json:"geo-location-name"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	Age                    int    `json:"age"`
}

type DdosDstZoneOperPacketAnomalyDetection struct {
	Oper DdosDstZoneOperPacketAnomalyDetectionOper `json:"oper"`
}

type DdosDstZoneOperPacketAnomalyDetectionOper struct {
	Indicators []DdosDstZoneOperPacketAnomalyDetectionOperIndicators `json:"indicators"`
	DataSource string                                                `json:"data-source"`
}

type DdosDstZoneOperPacketAnomalyDetectionOperIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	Maximum        string `json:"maximum"`
	Minimum        string `json:"minimum"`
	Threshold      string `json:"threshold"`
	IsAnomaly      int    `json:"is-anomaly"`
}

type DdosDstZoneOperPort struct {
	Oper                 DdosDstZoneOperPortOper                   `json:"oper"`
	ZoneServiceList      []DdosDstZoneOperPortZoneServiceList      `json:"zone-service-list"`
	ZoneServiceOtherList []DdosDstZoneOperPortZoneServiceOtherList `json:"zone-service-other-list"`
}

type DdosDstZoneOperPortOper struct {
}

type DdosDstZoneOperPortZoneServiceList struct {
	PortNum                     int                                                           `json:"port-num"`
	Protocol                    string                                                        `json:"protocol"`
	Oper                        DdosDstZoneOperPortZoneServiceListOper                        `json:"oper"`
	PatternRecognition          DdosDstZoneOperPortZoneServiceListPatternRecognition          `json:"pattern-recognition"`
	PatternRecognitionPuDetails DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
	IpFilteringPolicyStatistics DdosDstZoneOperPortZoneServiceListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
	PortInd                     DdosDstZoneOperPortZoneServiceListPortInd                     `json:"port-ind"`
	TopkSources                 DdosDstZoneOperPortZoneServiceListTopkSources                 `json:"topk-sources"`
	ProgressionTracking         DdosDstZoneOperPortZoneServiceListProgressionTracking         `json:"progression-tracking"`
	TopkDestinations            DdosDstZoneOperPortZoneServiceListTopkDestinations            `json:"topk-destinations"`
	SrcBasedPolicyList          []DdosDstZoneOperPortZoneServiceListSrcBasedPolicyList        `json:"src-based-policy-list"`
	Virtualhosts                DdosDstZoneOperPortZoneServiceListVirtualhosts                `json:"virtualhosts"`
}

type DdosDstZoneOperPortZoneServiceListOper struct {
	Ddos_entry_list       []DdosDstZoneOperPortZoneServiceListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                     `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                     `json:"service-displayed-count"`
	ReportingStatus       int                                                     `json:"reporting-status"`
	Sources               int                                                     `json:"sources"`
	OverflowPolicy        int                                                     `json:"overflow-policy"`
	SourcesAllEntries     int                                                     `json:"sources-all-entries"`
	ClassList             string                                                  `json:"class-list"`
	SubnetIpAddr          string                                                  `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                  `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                  `json:"ipv6"`
	Exceeded              int                                                     `json:"exceeded"`
	BlackListed           int                                                     `json:"black-listed"`
	WhiteListed           int                                                     `json:"white-listed"`
	Authenticated         int                                                     `json:"authenticated"`
	Level                 int                                                     `json:"level"`
	AppStat               int                                                     `json:"app-stat"`
	Indicators            int                                                     `json:"indicators"`
	IndicatorDetail       int                                                     `json:"indicator-detail"`
	L4ExtRate             int                                                     `json:"l4-ext-rate"`
	HwBlacklisted         int                                                     `json:"hw-blacklisted"`
	HwBlacklistedStats    int                                                     `json:"hw-blacklisted-stats"`
	SuffixRequestRate     int                                                     `json:"suffix-request-rate"`
	DomainName            string                                                  `json:"domain-name"`
}

type DdosDstZoneOperPortZoneServiceListOperDdos_entry_list struct {
	DstAddressStr          string                                                                   `json:"dst-address-str"`
	BwState                string                                                                   `json:"bw-state"`
	Is_auth_passed         string                                                                   `json:"is_auth_passed"`
	Level                  int                                                                      `json:"level"`
	BlReasoningRcode       string                                                                   `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string                                                                   `json:"bl-reasoning-timestamp"`
	CurrentConnections     string                                                                   `json:"current-connections"`
	IsConnectionsExceed    int                                                                      `json:"is-connections-exceed"`
	ConnectionLimit        string                                                                   `json:"connection-limit"`
	CurrentConnectionRate  string                                                                   `json:"current-connection-rate"`
	IsConnectionRateExceed int                                                                      `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string                                                                   `json:"connection-rate-limit"`
	CurrentPacketRate      string                                                                   `json:"current-packet-rate"`
	IsPacketRateExceed     int                                                                      `json:"is-packet-rate-exceed"`
	PacketRateLimit        string                                                                   `json:"packet-rate-limit"`
	CurrentKbitRate        string                                                                   `json:"current-kBit-rate"`
	IsKbitRateExceed       int                                                                      `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string                                                                   `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string                                                                   `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int                                                                      `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string                                                                   `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string                                                                   `json:"current-app-stat1"`
	IsAppStat1Exceed       int                                                                      `json:"is-app-stat1-exceed"`
	AppStat1Limit          string                                                                   `json:"app-stat1-limit"`
	CurrentAppStat2        string                                                                   `json:"current-app-stat2"`
	IsAppStat2Exceed       int                                                                      `json:"is-app-stat2-exceed"`
	AppStat2Limit          string                                                                   `json:"app-stat2-limit"`
	CurrentAppStat3        string                                                                   `json:"current-app-stat3"`
	IsAppStat3Exceed       int                                                                      `json:"is-app-stat3-exceed"`
	AppStat3Limit          string                                                                   `json:"app-stat3-limit"`
	CurrentAppStat4        string                                                                   `json:"current-app-stat4"`
	IsAppStat4Exceed       int                                                                      `json:"is-app-stat4-exceed"`
	AppStat4Limit          string                                                                   `json:"app-stat4-limit"`
	CurrentAppStat5        string                                                                   `json:"current-app-stat5"`
	IsAppStat5Exceed       int                                                                      `json:"is-app-stat5-exceed"`
	AppStat5Limit          string                                                                   `json:"app-stat5-limit"`
	CurrentAppStat6        string                                                                   `json:"current-app-stat6"`
	IsAppStat6Exceed       int                                                                      `json:"is-app-stat6-exceed"`
	AppStat6Limit          string                                                                   `json:"app-stat6-limit"`
	CurrentAppStat7        string                                                                   `json:"current-app-stat7"`
	IsAppStat7Exceed       int                                                                      `json:"is-app-stat7-exceed"`
	AppStat7Limit          string                                                                   `json:"app-stat7-limit"`
	CurrentAppStat8        string                                                                   `json:"current-app-stat8"`
	IsAppStat8Exceed       int                                                                      `json:"is-app-stat8-exceed"`
	AppStat8Limit          string                                                                   `json:"app-stat8-limit"`
	Age                    int                                                                      `json:"age"`
	LockupTime             int                                                                      `json:"lockup-time"`
	DynamicEntryCount      string                                                                   `json:"dynamic-entry-count"`
	DynamicEntryLimit      string                                                                   `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string                                                                   `json:"dynamic-entry-warn-state"`
	SflowSourceId          int                                                                      `json:"sflow-source-id"`
	HttpFilterRates        []DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates   `json:"http-filter-rates"`
	ResponseSizeRates      []DdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates `json:"response-size-rates"`
	HwBlockedRules         []DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHwBlockedRules    `json:"hw-blocked-rules"`
	DebugStr               string                                                                   `json:"debug-str"`
}

type DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHttpFilterRates struct {
	HttpFilterRateName          string `json:"http-filter-rate-name"`
	IsHttpFilterRateLimitExceed int    `json:"is-http-filter-rate-limit-exceed"`
	CurrentHttpFilterRate       string `json:"current-http-filter-rate"`
	HttpFilterRateLimit         string `json:"http-filter-rate-limit"`
}

type DdosDstZoneOperPortZoneServiceListOperDdos_entry_listResponseSizeRates struct {
	ResponseSizeRateName          string `json:"response-size-rate-name"`
	IsResponseSizeRateLimitExceed int    `json:"is-response-size-rate-limit-exceed"`
	CurrentResponseSizeRate       string `json:"current-response-size-rate"`
	ResponseSizeRateLimit         string `json:"response-size-rate-limit"`
}

type DdosDstZoneOperPortZoneServiceListOperDdos_entry_listHwBlockedRules struct {
	RuleDstIp       string `json:"rule-dst-ip"`
	HwBlockingState string `json:"hw-blocking-state"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognition struct {
	Oper DdosDstZoneOperPortZoneServiceListPatternRecognitionOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognitionOper struct {
	State            string                                                               `json:"state"`
	Timestamp        string                                                               `json:"timestamp"`
	PeacePktCount    int                                                                  `json:"peace-pkt-count"`
	WarPktCount      int                                                                  `json:"war-pkt-count"`
	WarPktPercentage int                                                                  `json:"war-pkt-percentage"`
	FilterThreshold  int                                                                  `json:"filter-threshold"`
	FilterCount      int                                                                  `json:"filter-count"`
	FilterList       []DdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList `json:"filter-list"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognitionOperFilterList struct {
	ProcessingUnit string `json:"processing-unit"`
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetails struct {
	Oper DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOper struct {
	AllFilters []DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFilters struct {
	ProcessingUnit   string                                                                                  `json:"processing-unit"`
	State            string                                                                                  `json:"state"`
	Timestamp        string                                                                                  `json:"timestamp"`
	PeacePktCount    int                                                                                     `json:"peace-pkt-count"`
	WarPktCount      int                                                                                     `json:"war-pkt-count"`
	WarPktPercentage int                                                                                     `json:"war-pkt-percentage"`
	FilterThreshold  int                                                                                     `json:"filter-threshold"`
	FilterCount      int                                                                                     `json:"filter-count"`
	FilterList       []DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}

type DdosDstZoneOperPortZoneServiceListPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZoneOperPortZoneServiceListIpFilteringPolicyStatistics struct {
	Oper DdosDstZoneOperPortZoneServiceListIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZoneOperPortZoneServiceListIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZoneOperPortZoneServiceListIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZoneOperPortZoneServiceListPortInd struct {
	Oper DdosDstZoneOperPortZoneServiceListPortIndOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListPortIndOper struct {
	SrcEntryList        []DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList `json:"src-entry-list"`
	Indicators          []DdosDstZoneOperPortZoneServiceListPortIndOperIndicators   `json:"indicators"`
	DetectionDataSource string                                                      `json:"detection-data-source"`
	TotalScore          string                                                      `json:"total-score"`
	CurrentLevel        string                                                      `json:"current-level"`
	EscalationTimestamp string                                                      `json:"escalation-timestamp"`
	InitialLearning     string                                                      `json:"initial-learning"`
	ActiveTime          int                                                         `json:"active-time"`
	BaselineWindowSize  int                                                         `json:"baseline-window-size"`
	SourcesAllEntries   int                                                         `json:"sources-all-entries"`
	SubnetIpAddr        string                                                      `json:"subnet-ip-addr"`
	SubnetIpv6Addr      string                                                      `json:"subnet-ipv6-addr"`
	Ipv6                string                                                      `json:"ipv6"`
	Details             int                                                         `json:"details"`
	Sources             int                                                         `json:"sources"`
}

type DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryList struct {
	SrcAddressStr       string                                                                `json:"src-address-str"`
	Indicators          []DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators `json:"indicators"`
	DetectionDataSource string                                                                `json:"detection-data-source"`
	TotalScore          string                                                                `json:"total-score"`
	CurrentLevel        string                                                                `json:"current-level"`
	SrcLevel            string                                                                `json:"src-level"`
	EscalationTimestamp string                                                                `json:"escalation-timestamp"`
	InitialLearning     string                                                                `json:"initial-learning"`
	ActiveTime          int                                                                   `json:"active-time"`
}

type DdosDstZoneOperPortZoneServiceListPortIndOperSrcEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	SrcMaximum        string `json:"src-maximum"`
	SrcMinimum        string `json:"src-minimum"`
	SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
	SrcAverage        string `json:"src-average"`
	Score             string `json:"score"`
}

type DdosDstZoneOperPortZoneServiceListPortIndOperIndicators struct {
	IndicatorName         string                                                                `json:"indicator-name"`
	IndicatorIndex        int                                                                   `json:"indicator-index"`
	Rate                  string                                                                `json:"rate"`
	ZoneMaximum           string                                                                `json:"zone-maximum"`
	ZoneMinimum           string                                                                `json:"zone-minimum"`
	ZoneNonZeroMinimum    string                                                                `json:"zone-non-zero-minimum"`
	ZoneAverage           string                                                                `json:"zone-average"`
	ZoneAdaptiveThreshold string                                                                `json:"zone-adaptive-threshold"`
	SrcMaximum            string                                                                `json:"src-maximum"`
	IndicatorCfg          []DdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
	Score                 string                                                                `json:"score"`
}

type DdosDstZoneOperPortZoneServiceListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperPortZoneServiceListTopkSources struct {
	Oper DdosDstZoneOperPortZoneServiceListTopkSourcesOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListTopkSourcesOper struct {
	Indicators    []DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList  `json:"entry-list"`
	NextIndicator int                                                           `json:"next-indicator"`
	Finished      int                                                           `json:"finished"`
	Details       int                                                           `json:"details"`
	TopKKey       string                                                        `json:"top-k-key"`
}

type DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicators struct {
	IndicatorName  string                                                               `json:"indicator-name"`
	IndicatorIndex int                                                                  `json:"indicator-index"`
	Sources        []DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources `json:"sources"`
}

type DdosDstZoneOperPortZoneServiceListTopkSourcesOperIndicatorsSources struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryList struct {
	AddressStr string                                                                 `json:"address-str"`
	Indicators []DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperPortZoneServiceListTopkSourcesOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperPortZoneServiceListProgressionTracking struct {
	Oper DdosDstZoneOperPortZoneServiceListProgressionTrackingOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListProgressionTrackingOper struct {
	Indicators          []DdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators `json:"indicators"`
	LearningDetails     int                                                                   `json:"learning-details"`
	LearningBrief       int                                                                   `json:"learning-brief"`
	RecommendedTemplate int                                                                   `json:"recommended-template"`
	TemplateDebugTable  int                                                                   `json:"template-debug-table"`
}

type DdosDstZoneOperPortZoneServiceListProgressionTrackingOperIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	NumSample         int    `json:"num-sample"`
	Average           string `json:"average"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	StandardDeviation string `json:"standard-deviation"`
}

type DdosDstZoneOperPortZoneServiceListTopkDestinations struct {
	Oper DdosDstZoneOperPortZoneServiceListTopkDestinationsOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListTopkDestinationsOper struct {
	Indicators    []DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                                                `json:"next-indicator"`
	Finished      int                                                                `json:"finished"`
	Details       int                                                                `json:"details"`
	TopKKey       string                                                             `json:"top-k-key"`
}

type DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                                         `json:"indicator-name"`
	IndicatorIndex int                                                                            `json:"indicator-index"`
	Destinations   []DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZoneOperPortZoneServiceListTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryList struct {
	AddressStr string                                                                      `json:"address-str"`
	Indicators []DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperPortZoneServiceListTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperPortZoneServiceListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                    `json:"src-based-policy-name"`
	Oper                DdosDstZoneOperPortZoneServiceListSrcBasedPolicyListOper                  `json:"oper"`
	PolicyClassListList []DdosDstZoneOperPortZoneServiceListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneOperPortZoneServiceListSrcBasedPolicyListOper struct {
}

type DdosDstZoneOperPortZoneServiceListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName string                                                                      `json:"class-list-name"`
	Oper          DdosDstZoneOperPortZoneServiceListSrcBasedPolicyListPolicyClassListListOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListSrcBasedPolicyListPolicyClassListListOper struct {
	CurrentConnections     int    `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        int    `json:"connection-limit"`
	CurrentConnectionRate  int    `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    int    `json:"connection-rate-limit"`
	CurrentPacketRate      int    `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        int    `json:"packet-rate-limit"`
	CurrentKbitRate        int    `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          int    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  int    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    int    `json:"frag-packet-rate-limit"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhosts struct {
	Oper            DdosDstZoneOperPortZoneServiceListVirtualhostsOper              `json:"oper"`
	VirtualhostList []DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostList `json:"virtualhost-list"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsOper struct {
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostList struct {
	Vhost string                                                            `json:"vhost"`
	Oper  DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOper struct {
	Ddos_entry_list       []DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                                                `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                                                `json:"service-displayed-count"`
	ReportingStatus       int                                                                                `json:"reporting-status"`
	Sources               int                                                                                `json:"sources"`
	OverflowPolicy        int                                                                                `json:"overflow-policy"`
	SourcesAllEntries     int                                                                                `json:"sources-all-entries"`
	ClassList             string                                                                             `json:"class-list"`
	SubnetIpAddr          string                                                                             `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                                             `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                                             `json:"ipv6"`
	Exceeded              int                                                                                `json:"exceeded"`
	BlackListed           int                                                                                `json:"black-listed"`
	WhiteListed           int                                                                                `json:"white-listed"`
	Authenticated         int                                                                                `json:"authenticated"`
	Level                 int                                                                                `json:"level"`
	AppStat               int                                                                                `json:"app-stat"`
	Indicators            int                                                                                `json:"indicators"`
	IndicatorDetail       int                                                                                `json:"indicator-detail"`
	L4ExtRate             int                                                                                `json:"l4-ext-rate"`
	HwBlacklisted         int                                                                                `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                                                `json:"suffix-request-rate"`
	DomainName            string                                                                             `json:"domain-name"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_list struct {
	DstAddressStr          string                                                                                              `json:"dst-address-str"`
	BwState                string                                                                                              `json:"bw-state"`
	Is_auth_passed         string                                                                                              `json:"is_auth_passed"`
	Level                  int                                                                                                 `json:"level"`
	BlReasoningRcode       string                                                                                              `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string                                                                                              `json:"bl-reasoning-timestamp"`
	CurrentConnections     string                                                                                              `json:"current-connections"`
	IsConnectionsExceed    int                                                                                                 `json:"is-connections-exceed"`
	ConnectionLimit        string                                                                                              `json:"connection-limit"`
	CurrentConnectionRate  string                                                                                              `json:"current-connection-rate"`
	IsConnectionRateExceed int                                                                                                 `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string                                                                                              `json:"connection-rate-limit"`
	CurrentPacketRate      string                                                                                              `json:"current-packet-rate"`
	IsPacketRateExceed     int                                                                                                 `json:"is-packet-rate-exceed"`
	PacketRateLimit        string                                                                                              `json:"packet-rate-limit"`
	CurrentKbitRate        string                                                                                              `json:"current-kBit-rate"`
	IsKbitRateExceed       int                                                                                                 `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string                                                                                              `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string                                                                                              `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int                                                                                                 `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string                                                                                              `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string                                                                                              `json:"current-app-stat1"`
	IsAppStat1Exceed       int                                                                                                 `json:"is-app-stat1-exceed"`
	AppStat1Limit          string                                                                                              `json:"app-stat1-limit"`
	CurrentAppStat2        string                                                                                              `json:"current-app-stat2"`
	IsAppStat2Exceed       int                                                                                                 `json:"is-app-stat2-exceed"`
	AppStat2Limit          string                                                                                              `json:"app-stat2-limit"`
	CurrentAppStat3        string                                                                                              `json:"current-app-stat3"`
	IsAppStat3Exceed       int                                                                                                 `json:"is-app-stat3-exceed"`
	AppStat3Limit          string                                                                                              `json:"app-stat3-limit"`
	CurrentAppStat4        string                                                                                              `json:"current-app-stat4"`
	IsAppStat4Exceed       int                                                                                                 `json:"is-app-stat4-exceed"`
	AppStat4Limit          string                                                                                              `json:"app-stat4-limit"`
	CurrentAppStat5        string                                                                                              `json:"current-app-stat5"`
	IsAppStat5Exceed       int                                                                                                 `json:"is-app-stat5-exceed"`
	AppStat5Limit          string                                                                                              `json:"app-stat5-limit"`
	CurrentAppStat6        string                                                                                              `json:"current-app-stat6"`
	IsAppStat6Exceed       int                                                                                                 `json:"is-app-stat6-exceed"`
	AppStat6Limit          string                                                                                              `json:"app-stat6-limit"`
	CurrentAppStat7        string                                                                                              `json:"current-app-stat7"`
	IsAppStat7Exceed       int                                                                                                 `json:"is-app-stat7-exceed"`
	AppStat7Limit          string                                                                                              `json:"app-stat7-limit"`
	CurrentAppStat8        string                                                                                              `json:"current-app-stat8"`
	IsAppStat8Exceed       int                                                                                                 `json:"is-app-stat8-exceed"`
	AppStat8Limit          string                                                                                              `json:"app-stat8-limit"`
	Age                    int                                                                                                 `json:"age"`
	LockupTime             int                                                                                                 `json:"lockup-time"`
	DynamicEntryCount      string                                                                                              `json:"dynamic-entry-count"`
	DynamicEntryLimit      string                                                                                              `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string                                                                                              `json:"dynamic-entry-warn-state"`
	SflowSourceId          int                                                                                                 `json:"sflow-source-id"`
	HttpFilterRates        []DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_listHttpFilterRates   `json:"http-filter-rates"`
	ResponseSizeRates      []DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_listResponseSizeRates `json:"response-size-rates"`
	HwBlockedRules         []DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_listHwBlockedRules    `json:"hw-blocked-rules"`
	DebugStr               string                                                                                              `json:"debug-str"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_listHttpFilterRates struct {
	HttpFilterRateName          string `json:"http-filter-rate-name"`
	IsHttpFilterRateLimitExceed int    `json:"is-http-filter-rate-limit-exceed"`
	CurrentHttpFilterRate       string `json:"current-http-filter-rate"`
	HttpFilterRateLimit         string `json:"http-filter-rate-limit"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_listResponseSizeRates struct {
	ResponseSizeRateName          string `json:"response-size-rate-name"`
	IsResponseSizeRateLimitExceed int    `json:"is-response-size-rate-limit-exceed"`
	CurrentResponseSizeRate       string `json:"current-response-size-rate"`
	ResponseSizeRateLimit         string `json:"response-size-rate-limit"`
}

type DdosDstZoneOperPortZoneServiceListVirtualhostsVirtualhostListOperDdos_entry_listHwBlockedRules struct {
	RuleDstIp       string `json:"rule-dst-ip"`
	HwBlockingState string `json:"hw-blocking-state"`
}

type DdosDstZoneOperPortZoneServiceOtherList struct {
	PortOther                   string                                                             `json:"port-other"`
	Protocol                    string                                                             `json:"protocol"`
	Oper                        DdosDstZoneOperPortZoneServiceOtherListOper                        `json:"oper"`
	IpFilteringPolicyStatistics DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
	PatternRecognition          DdosDstZoneOperPortZoneServiceOtherListPatternRecognition          `json:"pattern-recognition"`
	PatternRecognitionPuDetails DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
	PortInd                     DdosDstZoneOperPortZoneServiceOtherListPortInd                     `json:"port-ind"`
	TopkSources                 DdosDstZoneOperPortZoneServiceOtherListTopkSources                 `json:"topk-sources"`
	ProgressionTracking         DdosDstZoneOperPortZoneServiceOtherListProgressionTracking         `json:"progression-tracking"`
	TopkDestinations            DdosDstZoneOperPortZoneServiceOtherListTopkDestinations            `json:"topk-destinations"`
	SrcBasedPolicyList          []DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyList        `json:"src-based-policy-list"`
}

type DdosDstZoneOperPortZoneServiceOtherListOper struct {
	Ddos_entry_list       []DdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                          `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                          `json:"service-displayed-count"`
	ReportingStatus       int                                                          `json:"reporting-status"`
	Sources               int                                                          `json:"sources"`
	OverflowPolicy        int                                                          `json:"overflow-policy"`
	SourcesAllEntries     int                                                          `json:"sources-all-entries"`
	ClassList             string                                                       `json:"class-list"`
	SubnetIpAddr          string                                                       `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                       `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                       `json:"ipv6"`
	Exceeded              int                                                          `json:"exceeded"`
	BlackListed           int                                                          `json:"black-listed"`
	WhiteListed           int                                                          `json:"white-listed"`
	Authenticated         int                                                          `json:"authenticated"`
	Level                 int                                                          `json:"level"`
	AppStat               int                                                          `json:"app-stat"`
	Indicators            int                                                          `json:"indicators"`
	IndicatorDetail       int                                                          `json:"indicator-detail"`
	L4ExtRate             int                                                          `json:"l4-ext-rate"`
	HwBlacklisted         int                                                          `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                          `json:"suffix-request-rate"`
	DomainName            string                                                       `json:"domain-name"`
}

type DdosDstZoneOperPortZoneServiceOtherListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	BlReasoningRcode       string `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string `json:"bl-reasoning-timestamp"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string `json:"current-app-stat1"`
	IsAppStat1Exceed       int    `json:"is-app-stat1-exceed"`
	AppStat1Limit          string `json:"app-stat1-limit"`
	CurrentAppStat2        string `json:"current-app-stat2"`
	IsAppStat2Exceed       int    `json:"is-app-stat2-exceed"`
	AppStat2Limit          string `json:"app-stat2-limit"`
	CurrentAppStat3        string `json:"current-app-stat3"`
	IsAppStat3Exceed       int    `json:"is-app-stat3-exceed"`
	AppStat3Limit          string `json:"app-stat3-limit"`
	CurrentAppStat4        string `json:"current-app-stat4"`
	IsAppStat4Exceed       int    `json:"is-app-stat4-exceed"`
	AppStat4Limit          string `json:"app-stat4-limit"`
	CurrentAppStat5        string `json:"current-app-stat5"`
	IsAppStat5Exceed       int    `json:"is-app-stat5-exceed"`
	AppStat5Limit          string `json:"app-stat5-limit"`
	CurrentAppStat6        string `json:"current-app-stat6"`
	IsAppStat6Exceed       int    `json:"is-app-stat6-exceed"`
	AppStat6Limit          string `json:"app-stat6-limit"`
	CurrentAppStat7        string `json:"current-app-stat7"`
	IsAppStat7Exceed       int    `json:"is-app-stat7-exceed"`
	AppStat7Limit          string `json:"app-stat7-limit"`
	CurrentAppStat8        string `json:"current-app-stat8"`
	IsAppStat8Exceed       int    `json:"is-app-stat8-exceed"`
	AppStat8Limit          string `json:"app-stat8-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyStatistics struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZoneOperPortZoneServiceOtherListIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognition struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOper struct {
	State            string                                                                    `json:"state"`
	Timestamp        string                                                                    `json:"timestamp"`
	PeacePktCount    int                                                                       `json:"peace-pkt-count"`
	WarPktCount      int                                                                       `json:"war-pkt-count"`
	WarPktPercentage int                                                                       `json:"war-pkt-percentage"`
	FilterThreshold  int                                                                       `json:"filter-threshold"`
	FilterCount      int                                                                       `json:"filter-count"`
	FilterList       []DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList `json:"filter-list"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionOperFilterList struct {
	ProcessingUnit string `json:"processing-unit"`
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetails struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOper struct {
	AllFilters []DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFilters struct {
	ProcessingUnit   string                                                                                       `json:"processing-unit"`
	State            string                                                                                       `json:"state"`
	Timestamp        string                                                                                       `json:"timestamp"`
	PeacePktCount    int                                                                                          `json:"peace-pkt-count"`
	WarPktCount      int                                                                                          `json:"war-pkt-count"`
	WarPktPercentage int                                                                                          `json:"war-pkt-percentage"`
	FilterThreshold  int                                                                                          `json:"filter-threshold"`
	FilterCount      int                                                                                          `json:"filter-count"`
	FilterList       []DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}

type DdosDstZoneOperPortZoneServiceOtherListPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZoneOperPortZoneServiceOtherListPortInd struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListPortIndOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListPortIndOper struct {
	SrcEntryList        []DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList `json:"src-entry-list"`
	Indicators          []DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators   `json:"indicators"`
	DetectionDataSource string                                                           `json:"detection-data-source"`
	TotalScore          string                                                           `json:"total-score"`
	CurrentLevel        string                                                           `json:"current-level"`
	EscalationTimestamp string                                                           `json:"escalation-timestamp"`
	InitialLearning     string                                                           `json:"initial-learning"`
	ActiveTime          int                                                              `json:"active-time"`
	BaselineWindowSize  int                                                              `json:"baseline-window-size"`
	SourcesAllEntries   int                                                              `json:"sources-all-entries"`
	SubnetIpAddr        string                                                           `json:"subnet-ip-addr"`
	SubnetIpv6Addr      string                                                           `json:"subnet-ipv6-addr"`
	Ipv6                string                                                           `json:"ipv6"`
	Details             int                                                              `json:"details"`
	Sources             int                                                              `json:"sources"`
}

type DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryList struct {
	SrcAddressStr       string                                                                     `json:"src-address-str"`
	Indicators          []DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators `json:"indicators"`
	DetectionDataSource string                                                                     `json:"detection-data-source"`
	TotalScore          string                                                                     `json:"total-score"`
	CurrentLevel        string                                                                     `json:"current-level"`
	SrcLevel            string                                                                     `json:"src-level"`
	EscalationTimestamp string                                                                     `json:"escalation-timestamp"`
	InitialLearning     string                                                                     `json:"initial-learning"`
	ActiveTime          int                                                                        `json:"active-time"`
}

type DdosDstZoneOperPortZoneServiceOtherListPortIndOperSrcEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	SrcMaximum        string `json:"src-maximum"`
	SrcMinimum        string `json:"src-minimum"`
	SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
	SrcAverage        string `json:"src-average"`
	Score             string `json:"score"`
}

type DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicators struct {
	IndicatorName         string                                                                     `json:"indicator-name"`
	IndicatorIndex        int                                                                        `json:"indicator-index"`
	Rate                  string                                                                     `json:"rate"`
	ZoneMaximum           string                                                                     `json:"zone-maximum"`
	ZoneMinimum           string                                                                     `json:"zone-minimum"`
	ZoneNonZeroMinimum    string                                                                     `json:"zone-non-zero-minimum"`
	ZoneAverage           string                                                                     `json:"zone-average"`
	ZoneAdaptiveThreshold string                                                                     `json:"zone-adaptive-threshold"`
	SrcMaximum            string                                                                     `json:"src-maximum"`
	IndicatorCfg          []DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
	Score                 string                                                                     `json:"score"`
}

type DdosDstZoneOperPortZoneServiceOtherListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkSources struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOper struct {
	Indicators    []DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList  `json:"entry-list"`
	NextIndicator int                                                                `json:"next-indicator"`
	Finished      int                                                                `json:"finished"`
	Details       int                                                                `json:"details"`
	TopKKey       string                                                             `json:"top-k-key"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicators struct {
	IndicatorName  string                                                                    `json:"indicator-name"`
	IndicatorIndex int                                                                       `json:"indicator-index"`
	Sources        []DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources `json:"sources"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperIndicatorsSources struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryList struct {
	AddressStr string                                                                      `json:"address-str"`
	Indicators []DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkSourcesOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperPortZoneServiceOtherListProgressionTracking struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOper struct {
	Indicators          []DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators `json:"indicators"`
	LearningDetails     int                                                                        `json:"learning-details"`
	LearningBrief       int                                                                        `json:"learning-brief"`
	RecommendedTemplate int                                                                        `json:"recommended-template"`
	TemplateDebugTable  int                                                                        `json:"template-debug-table"`
}

type DdosDstZoneOperPortZoneServiceOtherListProgressionTrackingOperIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	NumSample         int    `json:"num-sample"`
	Average           string `json:"average"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	StandardDeviation string `json:"standard-deviation"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkDestinations struct {
	Oper DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOper struct {
	Indicators    []DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                                                     `json:"next-indicator"`
	Finished      int                                                                     `json:"finished"`
	Details       int                                                                     `json:"details"`
	TopKKey       string                                                                  `json:"top-k-key"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                                              `json:"indicator-name"`
	IndicatorIndex int                                                                                 `json:"indicator-index"`
	Destinations   []DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryList struct {
	AddressStr string                                                                           `json:"address-str"`
	Indicators []DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperPortZoneServiceOtherListTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                         `json:"src-based-policy-name"`
	Oper                DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyListOper                  `json:"oper"`
	PolicyClassListList []DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyListOper struct {
}

type DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName string                                                                           `json:"class-list-name"`
	Oper          DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListOper `json:"oper"`
}

type DdosDstZoneOperPortZoneServiceOtherListSrcBasedPolicyListPolicyClassListListOper struct {
	CurrentConnections     int    `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        int    `json:"connection-limit"`
	CurrentConnectionRate  int    `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    int    `json:"connection-rate-limit"`
	CurrentPacketRate      int    `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        int    `json:"packet-rate-limit"`
	CurrentKbitRate        int    `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          int    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  int    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    int    `json:"frag-packet-rate-limit"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperPortRangeList struct {
	PortRangeStart              int                                                     `json:"port-range-start"`
	PortRangeEnd                int                                                     `json:"port-range-end"`
	Protocol                    string                                                  `json:"protocol"`
	Oper                        DdosDstZoneOperPortRangeListOper                        `json:"oper"`
	IpFilteringPolicyStatistics DdosDstZoneOperPortRangeListIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`
	PatternRecognition          DdosDstZoneOperPortRangeListPatternRecognition          `json:"pattern-recognition"`
	PatternRecognitionPuDetails DdosDstZoneOperPortRangeListPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`
	PortInd                     DdosDstZoneOperPortRangeListPortInd                     `json:"port-ind"`
	TopkSources                 DdosDstZoneOperPortRangeListTopkSources                 `json:"topk-sources"`
	TopkDestinations            DdosDstZoneOperPortRangeListTopkDestinations            `json:"topk-destinations"`
	ProgressionTracking         DdosDstZoneOperPortRangeListProgressionTracking         `json:"progression-tracking"`
	SrcBasedPolicyList          []DdosDstZoneOperPortRangeListSrcBasedPolicyList        `json:"src-based-policy-list"`
	Virtualhosts                DdosDstZoneOperPortRangeListVirtualhosts                `json:"virtualhosts"`
}

type DdosDstZoneOperPortRangeListOper struct {
	Ddos_entry_list       []DdosDstZoneOperPortRangeListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                               `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                               `json:"service-displayed-count"`
	ReportingStatus       int                                               `json:"reporting-status"`
	Sources               int                                               `json:"sources"`
	OverflowPolicy        int                                               `json:"overflow-policy"`
	SourcesAllEntries     int                                               `json:"sources-all-entries"`
	ClassList             string                                            `json:"class-list"`
	SubnetIpAddr          string                                            `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                            `json:"subnet-ipv6-addr"`
	Ipv6                  string                                            `json:"ipv6"`
	Exceeded              int                                               `json:"exceeded"`
	BlackListed           int                                               `json:"black-listed"`
	WhiteListed           int                                               `json:"white-listed"`
	Authenticated         int                                               `json:"authenticated"`
	Level                 int                                               `json:"level"`
	AppStat               int                                               `json:"app-stat"`
	Indicators            int                                               `json:"indicators"`
	IndicatorDetail       int                                               `json:"indicator-detail"`
	L4ExtRate             int                                               `json:"l4-ext-rate"`
	HwBlacklisted         int                                               `json:"hw-blacklisted"`
	SuffixRequestRate     int                                               `json:"suffix-request-rate"`
	DomainName            string                                            `json:"domain-name"`
}

type DdosDstZoneOperPortRangeListOperDdos_entry_list struct {
	DstAddressStr          string                                                             `json:"dst-address-str"`
	BwState                string                                                             `json:"bw-state"`
	Is_auth_passed         string                                                             `json:"is_auth_passed"`
	Level                  int                                                                `json:"level"`
	BlReasoningRcode       string                                                             `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string                                                             `json:"bl-reasoning-timestamp"`
	CurrentConnections     string                                                             `json:"current-connections"`
	IsConnectionsExceed    int                                                                `json:"is-connections-exceed"`
	ConnectionLimit        string                                                             `json:"connection-limit"`
	CurrentConnectionRate  string                                                             `json:"current-connection-rate"`
	IsConnectionRateExceed int                                                                `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string                                                             `json:"connection-rate-limit"`
	CurrentPacketRate      string                                                             `json:"current-packet-rate"`
	IsPacketRateExceed     int                                                                `json:"is-packet-rate-exceed"`
	PacketRateLimit        string                                                             `json:"packet-rate-limit"`
	CurrentKbitRate        string                                                             `json:"current-kBit-rate"`
	IsKbitRateExceed       int                                                                `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string                                                             `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string                                                             `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int                                                                `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string                                                             `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string                                                             `json:"current-app-stat1"`
	IsAppStat1Exceed       int                                                                `json:"is-app-stat1-exceed"`
	AppStat1Limit          string                                                             `json:"app-stat1-limit"`
	CurrentAppStat2        string                                                             `json:"current-app-stat2"`
	IsAppStat2Exceed       int                                                                `json:"is-app-stat2-exceed"`
	AppStat2Limit          string                                                             `json:"app-stat2-limit"`
	CurrentAppStat3        string                                                             `json:"current-app-stat3"`
	IsAppStat3Exceed       int                                                                `json:"is-app-stat3-exceed"`
	AppStat3Limit          string                                                             `json:"app-stat3-limit"`
	CurrentAppStat4        string                                                             `json:"current-app-stat4"`
	IsAppStat4Exceed       int                                                                `json:"is-app-stat4-exceed"`
	AppStat4Limit          string                                                             `json:"app-stat4-limit"`
	CurrentAppStat5        string                                                             `json:"current-app-stat5"`
	IsAppStat5Exceed       int                                                                `json:"is-app-stat5-exceed"`
	AppStat5Limit          string                                                             `json:"app-stat5-limit"`
	CurrentAppStat6        string                                                             `json:"current-app-stat6"`
	IsAppStat6Exceed       int                                                                `json:"is-app-stat6-exceed"`
	AppStat6Limit          string                                                             `json:"app-stat6-limit"`
	CurrentAppStat7        string                                                             `json:"current-app-stat7"`
	IsAppStat7Exceed       int                                                                `json:"is-app-stat7-exceed"`
	AppStat7Limit          string                                                             `json:"app-stat7-limit"`
	CurrentAppStat8        string                                                             `json:"current-app-stat8"`
	IsAppStat8Exceed       int                                                                `json:"is-app-stat8-exceed"`
	AppStat8Limit          string                                                             `json:"app-stat8-limit"`
	Age                    int                                                                `json:"age"`
	LockupTime             int                                                                `json:"lockup-time"`
	DynamicEntryCount      string                                                             `json:"dynamic-entry-count"`
	DynamicEntryLimit      string                                                             `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string                                                             `json:"dynamic-entry-warn-state"`
	SflowSourceId          int                                                                `json:"sflow-source-id"`
	HttpFilterRates        []DdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates   `json:"http-filter-rates"`
	ResponseSizeRates      []DdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates `json:"response-size-rates"`
	DebugStr               string                                                             `json:"debug-str"`
}

type DdosDstZoneOperPortRangeListOperDdos_entry_listHttpFilterRates struct {
	HttpFilterRateName          string `json:"http-filter-rate-name"`
	IsHttpFilterRateLimitExceed int    `json:"is-http-filter-rate-limit-exceed"`
	CurrentHttpFilterRate       string `json:"current-http-filter-rate"`
	HttpFilterRateLimit         string `json:"http-filter-rate-limit"`
}

type DdosDstZoneOperPortRangeListOperDdos_entry_listResponseSizeRates struct {
	ResponseSizeRateName          string `json:"response-size-rate-name"`
	IsResponseSizeRateLimitExceed int    `json:"is-response-size-rate-limit-exceed"`
	CurrentResponseSizeRate       string `json:"current-response-size-rate"`
	ResponseSizeRateLimit         string `json:"response-size-rate-limit"`
}

type DdosDstZoneOperPortRangeListIpFilteringPolicyStatistics struct {
	Oper DdosDstZoneOperPortRangeListIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZoneOperPortRangeListIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZoneOperPortRangeListIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZoneOperPortRangeListPatternRecognition struct {
	Oper DdosDstZoneOperPortRangeListPatternRecognitionOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListPatternRecognitionOper struct {
	State            string                                                         `json:"state"`
	Timestamp        string                                                         `json:"timestamp"`
	PeacePktCount    int                                                            `json:"peace-pkt-count"`
	WarPktCount      int                                                            `json:"war-pkt-count"`
	WarPktPercentage int                                                            `json:"war-pkt-percentage"`
	FilterThreshold  int                                                            `json:"filter-threshold"`
	FilterCount      int                                                            `json:"filter-count"`
	FilterList       []DdosDstZoneOperPortRangeListPatternRecognitionOperFilterList `json:"filter-list"`
}

type DdosDstZoneOperPortRangeListPatternRecognitionOperFilterList struct {
	ProcessingUnit string `json:"processing-unit"`
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZoneOperPortRangeListPatternRecognitionPuDetails struct {
	Oper DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOper struct {
	AllFilters []DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}

type DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFilters struct {
	ProcessingUnit   string                                                                            `json:"processing-unit"`
	State            string                                                                            `json:"state"`
	Timestamp        string                                                                            `json:"timestamp"`
	PeacePktCount    int                                                                               `json:"peace-pkt-count"`
	WarPktCount      int                                                                               `json:"war-pkt-count"`
	WarPktPercentage int                                                                               `json:"war-pkt-percentage"`
	FilterThreshold  int                                                                               `json:"filter-threshold"`
	FilterCount      int                                                                               `json:"filter-count"`
	FilterList       []DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}

type DdosDstZoneOperPortRangeListPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZoneOperPortRangeListPortInd struct {
	Oper DdosDstZoneOperPortRangeListPortIndOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListPortIndOper struct {
	SrcEntryList        []DdosDstZoneOperPortRangeListPortIndOperSrcEntryList `json:"src-entry-list"`
	Indicators          []DdosDstZoneOperPortRangeListPortIndOperIndicators   `json:"indicators"`
	DetectionDataSource string                                                `json:"detection-data-source"`
	TotalScore          string                                                `json:"total-score"`
	CurrentLevel        string                                                `json:"current-level"`
	EscalationTimestamp string                                                `json:"escalation-timestamp"`
	InitialLearning     string                                                `json:"initial-learning"`
	ActiveTime          int                                                   `json:"active-time"`
	BaselineWindowSize  int                                                   `json:"baseline-window-size"`
	SourcesAllEntries   int                                                   `json:"sources-all-entries"`
	SubnetIpAddr        string                                                `json:"subnet-ip-addr"`
	SubnetIpv6Addr      string                                                `json:"subnet-ipv6-addr"`
	Ipv6                string                                                `json:"ipv6"`
	Details             int                                                   `json:"details"`
	Sources             int                                                   `json:"sources"`
}

type DdosDstZoneOperPortRangeListPortIndOperSrcEntryList struct {
	SrcAddressStr       string                                                          `json:"src-address-str"`
	Indicators          []DdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators `json:"indicators"`
	DetectionDataSource string                                                          `json:"detection-data-source"`
	TotalScore          string                                                          `json:"total-score"`
	CurrentLevel        string                                                          `json:"current-level"`
	SrcLevel            string                                                          `json:"src-level"`
	EscalationTimestamp string                                                          `json:"escalation-timestamp"`
	InitialLearning     string                                                          `json:"initial-learning"`
	ActiveTime          int                                                             `json:"active-time"`
}

type DdosDstZoneOperPortRangeListPortIndOperSrcEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	SrcMaximum        string `json:"src-maximum"`
	SrcMinimum        string `json:"src-minimum"`
	SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
	SrcAverage        string `json:"src-average"`
	Score             string `json:"score"`
}

type DdosDstZoneOperPortRangeListPortIndOperIndicators struct {
	IndicatorName         string                                                          `json:"indicator-name"`
	IndicatorIndex        int                                                             `json:"indicator-index"`
	Rate                  string                                                          `json:"rate"`
	ZoneMaximum           string                                                          `json:"zone-maximum"`
	ZoneMinimum           string                                                          `json:"zone-minimum"`
	ZoneNonZeroMinimum    string                                                          `json:"zone-non-zero-minimum"`
	ZoneAverage           string                                                          `json:"zone-average"`
	ZoneAdaptiveThreshold string                                                          `json:"zone-adaptive-threshold"`
	SrcMaximum            string                                                          `json:"src-maximum"`
	IndicatorCfg          []DdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
	Score                 string                                                          `json:"score"`
}

type DdosDstZoneOperPortRangeListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperPortRangeListTopkSources struct {
	Oper DdosDstZoneOperPortRangeListTopkSourcesOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListTopkSourcesOper struct {
	Indicators    []DdosDstZoneOperPortRangeListTopkSourcesOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperPortRangeListTopkSourcesOperEntryList  `json:"entry-list"`
	NextIndicator int                                                     `json:"next-indicator"`
	Finished      int                                                     `json:"finished"`
	Details       int                                                     `json:"details"`
	TopKKey       string                                                  `json:"top-k-key"`
}

type DdosDstZoneOperPortRangeListTopkSourcesOperIndicators struct {
	IndicatorName  string                                                         `json:"indicator-name"`
	IndicatorIndex int                                                            `json:"indicator-index"`
	Sources        []DdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources `json:"sources"`
}

type DdosDstZoneOperPortRangeListTopkSourcesOperIndicatorsSources struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperPortRangeListTopkSourcesOperEntryList struct {
	AddressStr string                                                           `json:"address-str"`
	Indicators []DdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperPortRangeListTopkSourcesOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperPortRangeListTopkDestinations struct {
	Oper DdosDstZoneOperPortRangeListTopkDestinationsOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListTopkDestinationsOper struct {
	Indicators    []DdosDstZoneOperPortRangeListTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperPortRangeListTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                                          `json:"next-indicator"`
	Finished      int                                                          `json:"finished"`
	Details       int                                                          `json:"details"`
	TopKKey       string                                                       `json:"top-k-key"`
}

type DdosDstZoneOperPortRangeListTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                                   `json:"indicator-name"`
	IndicatorIndex int                                                                      `json:"indicator-index"`
	Destinations   []DdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZoneOperPortRangeListTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperPortRangeListTopkDestinationsOperEntryList struct {
	AddressStr string                                                                `json:"address-str"`
	Indicators []DdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperPortRangeListTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZoneOperPortRangeListProgressionTracking struct {
	Oper DdosDstZoneOperPortRangeListProgressionTrackingOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListProgressionTrackingOper struct {
	Indicators          []DdosDstZoneOperPortRangeListProgressionTrackingOperIndicators `json:"indicators"`
	LearningDetails     int                                                             `json:"learning-details"`
	LearningBrief       int                                                             `json:"learning-brief"`
	RecommendedTemplate int                                                             `json:"recommended-template"`
	TemplateDebugTable  int                                                             `json:"template-debug-table"`
}

type DdosDstZoneOperPortRangeListProgressionTrackingOperIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	NumSample         int    `json:"num-sample"`
	Average           string `json:"average"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	StandardDeviation string `json:"standard-deviation"`
}

type DdosDstZoneOperPortRangeListSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                              `json:"src-based-policy-name"`
	Oper                DdosDstZoneOperPortRangeListSrcBasedPolicyListOper                  `json:"oper"`
	PolicyClassListList []DdosDstZoneOperPortRangeListSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZoneOperPortRangeListSrcBasedPolicyListOper struct {
}

type DdosDstZoneOperPortRangeListSrcBasedPolicyListPolicyClassListList struct {
	ClassListName string                                                                `json:"class-list-name"`
	Oper          DdosDstZoneOperPortRangeListSrcBasedPolicyListPolicyClassListListOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListSrcBasedPolicyListPolicyClassListListOper struct {
	CurrentConnections     int    `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        int    `json:"connection-limit"`
	CurrentConnectionRate  int    `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    int    `json:"connection-rate-limit"`
	CurrentPacketRate      int    `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        int    `json:"packet-rate-limit"`
	CurrentKbitRate        int    `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          int    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  int    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    int    `json:"frag-packet-rate-limit"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperPortRangeListVirtualhosts struct {
	Oper            DdosDstZoneOperPortRangeListVirtualhostsOper              `json:"oper"`
	VirtualhostList []DdosDstZoneOperPortRangeListVirtualhostsVirtualhostList `json:"virtualhost-list"`
}

type DdosDstZoneOperPortRangeListVirtualhostsOper struct {
}

type DdosDstZoneOperPortRangeListVirtualhostsVirtualhostList struct {
	Vhost string                                                      `json:"vhost"`
	Oper  DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOper `json:"oper"`
}

type DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOper struct {
	Ddos_entry_list       []DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                                          `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                                          `json:"service-displayed-count"`
	ReportingStatus       int                                                                          `json:"reporting-status"`
	Sources               int                                                                          `json:"sources"`
	OverflowPolicy        int                                                                          `json:"overflow-policy"`
	SourcesAllEntries     int                                                                          `json:"sources-all-entries"`
	ClassList             string                                                                       `json:"class-list"`
	SubnetIpAddr          string                                                                       `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                                       `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                                       `json:"ipv6"`
	Exceeded              int                                                                          `json:"exceeded"`
	BlackListed           int                                                                          `json:"black-listed"`
	WhiteListed           int                                                                          `json:"white-listed"`
	Authenticated         int                                                                          `json:"authenticated"`
	Level                 int                                                                          `json:"level"`
	AppStat               int                                                                          `json:"app-stat"`
	Indicators            int                                                                          `json:"indicators"`
	IndicatorDetail       int                                                                          `json:"indicator-detail"`
	L4ExtRate             int                                                                          `json:"l4-ext-rate"`
	HwBlacklisted         int                                                                          `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                                          `json:"suffix-request-rate"`
	DomainName            string                                                                       `json:"domain-name"`
}

type DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOperDdos_entry_list struct {
	DstAddressStr          string                                                                                        `json:"dst-address-str"`
	BwState                string                                                                                        `json:"bw-state"`
	Is_auth_passed         string                                                                                        `json:"is_auth_passed"`
	Level                  int                                                                                           `json:"level"`
	BlReasoningRcode       string                                                                                        `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string                                                                                        `json:"bl-reasoning-timestamp"`
	CurrentConnections     string                                                                                        `json:"current-connections"`
	IsConnectionsExceed    int                                                                                           `json:"is-connections-exceed"`
	ConnectionLimit        string                                                                                        `json:"connection-limit"`
	CurrentConnectionRate  string                                                                                        `json:"current-connection-rate"`
	IsConnectionRateExceed int                                                                                           `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string                                                                                        `json:"connection-rate-limit"`
	CurrentPacketRate      string                                                                                        `json:"current-packet-rate"`
	IsPacketRateExceed     int                                                                                           `json:"is-packet-rate-exceed"`
	PacketRateLimit        string                                                                                        `json:"packet-rate-limit"`
	CurrentKbitRate        string                                                                                        `json:"current-kBit-rate"`
	IsKbitRateExceed       int                                                                                           `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string                                                                                        `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string                                                                                        `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int                                                                                           `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string                                                                                        `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string                                                                                        `json:"current-app-stat1"`
	IsAppStat1Exceed       int                                                                                           `json:"is-app-stat1-exceed"`
	AppStat1Limit          string                                                                                        `json:"app-stat1-limit"`
	CurrentAppStat2        string                                                                                        `json:"current-app-stat2"`
	IsAppStat2Exceed       int                                                                                           `json:"is-app-stat2-exceed"`
	AppStat2Limit          string                                                                                        `json:"app-stat2-limit"`
	CurrentAppStat3        string                                                                                        `json:"current-app-stat3"`
	IsAppStat3Exceed       int                                                                                           `json:"is-app-stat3-exceed"`
	AppStat3Limit          string                                                                                        `json:"app-stat3-limit"`
	CurrentAppStat4        string                                                                                        `json:"current-app-stat4"`
	IsAppStat4Exceed       int                                                                                           `json:"is-app-stat4-exceed"`
	AppStat4Limit          string                                                                                        `json:"app-stat4-limit"`
	CurrentAppStat5        string                                                                                        `json:"current-app-stat5"`
	IsAppStat5Exceed       int                                                                                           `json:"is-app-stat5-exceed"`
	AppStat5Limit          string                                                                                        `json:"app-stat5-limit"`
	CurrentAppStat6        string                                                                                        `json:"current-app-stat6"`
	IsAppStat6Exceed       int                                                                                           `json:"is-app-stat6-exceed"`
	AppStat6Limit          string                                                                                        `json:"app-stat6-limit"`
	CurrentAppStat7        string                                                                                        `json:"current-app-stat7"`
	IsAppStat7Exceed       int                                                                                           `json:"is-app-stat7-exceed"`
	AppStat7Limit          string                                                                                        `json:"app-stat7-limit"`
	CurrentAppStat8        string                                                                                        `json:"current-app-stat8"`
	IsAppStat8Exceed       int                                                                                           `json:"is-app-stat8-exceed"`
	AppStat8Limit          string                                                                                        `json:"app-stat8-limit"`
	Age                    int                                                                                           `json:"age"`
	LockupTime             int                                                                                           `json:"lockup-time"`
	DynamicEntryCount      string                                                                                        `json:"dynamic-entry-count"`
	DynamicEntryLimit      string                                                                                        `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string                                                                                        `json:"dynamic-entry-warn-state"`
	SflowSourceId          int                                                                                           `json:"sflow-source-id"`
	HttpFilterRates        []DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOperDdos_entry_listHttpFilterRates   `json:"http-filter-rates"`
	ResponseSizeRates      []DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOperDdos_entry_listResponseSizeRates `json:"response-size-rates"`
	DebugStr               string                                                                                        `json:"debug-str"`
}

type DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOperDdos_entry_listHttpFilterRates struct {
	HttpFilterRateName          string `json:"http-filter-rate-name"`
	IsHttpFilterRateLimitExceed int    `json:"is-http-filter-rate-limit-exceed"`
	CurrentHttpFilterRate       string `json:"current-http-filter-rate"`
	HttpFilterRateLimit         string `json:"http-filter-rate-limit"`
}

type DdosDstZoneOperPortRangeListVirtualhostsVirtualhostListOperDdos_entry_listResponseSizeRates struct {
	ResponseSizeRateName          string `json:"response-size-rate-name"`
	IsResponseSizeRateLimitExceed int    `json:"is-response-size-rate-limit-exceed"`
	CurrentResponseSizeRate       string `json:"current-response-size-rate"`
	ResponseSizeRateLimit         string `json:"response-size-rate-limit"`
}

type DdosDstZoneOperSrcIpFiltering struct {
	Oper DdosDstZoneOperSrcIpFilteringOper `json:"oper"`
}

type DdosDstZoneOperSrcIpFilteringOper struct {
	ClassList []DdosDstZoneOperSrcIpFilteringOperClassList `json:"class-list"`
}

type DdosDstZoneOperSrcIpFilteringOperClassList struct {
	Name string `json:"name"`
	Hit  int    `json:"hit"`
}

type DdosDstZoneOperSrcPort struct {
	Oper                 DdosDstZoneOperSrcPortOper                   `json:"oper"`
	ZoneSrcPortList      []DdosDstZoneOperSrcPortZoneSrcPortList      `json:"zone-src-port-list"`
	ZoneSrcPortOtherList []DdosDstZoneOperSrcPortZoneSrcPortOtherList `json:"zone-src-port-other-list"`
}

type DdosDstZoneOperSrcPortOper struct {
}

type DdosDstZoneOperSrcPortZoneSrcPortList struct {
	PortNum  int                                          `json:"port-num"`
	Protocol string                                       `json:"protocol"`
	Oper     DdosDstZoneOperSrcPortZoneSrcPortListOper    `json:"oper"`
	PortInd  DdosDstZoneOperSrcPortZoneSrcPortListPortInd `json:"port-ind"`
}

type DdosDstZoneOperSrcPortZoneSrcPortListOper struct {
	Ddos_entry_list       []DdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                        `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                        `json:"service-displayed-count"`
	ReportingStatus       int                                                        `json:"reporting-status"`
	Sources               int                                                        `json:"sources"`
	SourcesAllEntries     int                                                        `json:"sources-all-entries"`
	SubnetIpAddr          string                                                     `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                     `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                     `json:"ipv6"`
	HwBlacklisted         int                                                        `json:"hw-blacklisted"`
}

type DdosDstZoneOperSrcPortZoneSrcPortListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string `json:"current-app-stat1"`
	IsAppStat1Exceed       int    `json:"is-app-stat1-exceed"`
	AppStat1Limit          string `json:"app-stat1-limit"`
	CurrentAppStat2        string `json:"current-app-stat2"`
	IsAppStat2Exceed       int    `json:"is-app-stat2-exceed"`
	AppStat2Limit          string `json:"app-stat2-limit"`
	CurrentAppStat3        string `json:"current-app-stat3"`
	IsAppStat3Exceed       int    `json:"is-app-stat3-exceed"`
	AppStat3Limit          string `json:"app-stat3-limit"`
	CurrentAppStat4        string `json:"current-app-stat4"`
	IsAppStat4Exceed       int    `json:"is-app-stat4-exceed"`
	AppStat4Limit          string `json:"app-stat4-limit"`
	CurrentAppStat5        string `json:"current-app-stat5"`
	IsAppStat5Exceed       int    `json:"is-app-stat5-exceed"`
	AppStat5Limit          string `json:"app-stat5-limit"`
	CurrentAppStat6        string `json:"current-app-stat6"`
	IsAppStat6Exceed       int    `json:"is-app-stat6-exceed"`
	AppStat6Limit          string `json:"app-stat6-limit"`
	CurrentAppStat7        string `json:"current-app-stat7"`
	IsAppStat7Exceed       int    `json:"is-app-stat7-exceed"`
	AppStat7Limit          string `json:"app-stat7-limit"`
	CurrentAppStat8        string `json:"current-app-stat8"`
	IsAppStat8Exceed       int    `json:"is-app-stat8-exceed"`
	AppStat8Limit          string `json:"app-stat8-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperSrcPortZoneSrcPortListPortInd struct {
	Oper DdosDstZoneOperSrcPortZoneSrcPortListPortIndOper `json:"oper"`
}

type DdosDstZoneOperSrcPortZoneSrcPortListPortIndOper struct {
	Indicators          []DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators `json:"indicators"`
	DetectionDataSource string                                                       `json:"detection-data-source"`
	CurrentLevel        string                                                       `json:"current-level"`
	Details             int                                                          `json:"details"`
}

type DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicators struct {
	IndicatorName  string                                                                   `json:"indicator-name"`
	IndicatorIndex int                                                                      `json:"indicator-index"`
	Rate           string                                                                   `json:"rate"`
	IndicatorCfg   []DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}

type DdosDstZoneOperSrcPortZoneSrcPortListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherList struct {
	PortOther string                                            `json:"port-other"`
	Protocol  string                                            `json:"protocol"`
	Oper      DdosDstZoneOperSrcPortZoneSrcPortOtherListOper    `json:"oper"`
	PortInd   DdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd `json:"port-ind"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherListOper struct {
	Ddos_entry_list       []DdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                             `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                             `json:"service-displayed-count"`
	ReportingStatus       int                                                             `json:"reporting-status"`
	Sources               int                                                             `json:"sources"`
	SourcesAllEntries     int                                                             `json:"sources-all-entries"`
	SubnetIpAddr          string                                                          `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                          `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                          `json:"ipv6"`
	HwBlacklisted         int                                                             `json:"hw-blacklisted"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string `json:"current-app-stat1"`
	IsAppStat1Exceed       int    `json:"is-app-stat1-exceed"`
	AppStat1Limit          string `json:"app-stat1-limit"`
	CurrentAppStat2        string `json:"current-app-stat2"`
	IsAppStat2Exceed       int    `json:"is-app-stat2-exceed"`
	AppStat2Limit          string `json:"app-stat2-limit"`
	CurrentAppStat3        string `json:"current-app-stat3"`
	IsAppStat3Exceed       int    `json:"is-app-stat3-exceed"`
	AppStat3Limit          string `json:"app-stat3-limit"`
	CurrentAppStat4        string `json:"current-app-stat4"`
	IsAppStat4Exceed       int    `json:"is-app-stat4-exceed"`
	AppStat4Limit          string `json:"app-stat4-limit"`
	CurrentAppStat5        string `json:"current-app-stat5"`
	IsAppStat5Exceed       int    `json:"is-app-stat5-exceed"`
	AppStat5Limit          string `json:"app-stat5-limit"`
	CurrentAppStat6        string `json:"current-app-stat6"`
	IsAppStat6Exceed       int    `json:"is-app-stat6-exceed"`
	AppStat6Limit          string `json:"app-stat6-limit"`
	CurrentAppStat7        string `json:"current-app-stat7"`
	IsAppStat7Exceed       int    `json:"is-app-stat7-exceed"`
	AppStat7Limit          string `json:"app-stat7-limit"`
	CurrentAppStat8        string `json:"current-app-stat8"`
	IsAppStat8Exceed       int    `json:"is-app-stat8-exceed"`
	AppStat8Limit          string `json:"app-stat8-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherListPortInd struct {
	Oper DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper `json:"oper"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOper struct {
	Indicators          []DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators `json:"indicators"`
	DetectionDataSource string                                                            `json:"detection-data-source"`
	CurrentLevel        string                                                            `json:"current-level"`
	Details             int                                                               `json:"details"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicators struct {
	IndicatorName  string                                                                        `json:"indicator-name"`
	IndicatorIndex int                                                                           `json:"indicator-index"`
	Rate           string                                                                        `json:"rate"`
	IndicatorCfg   []DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}

type DdosDstZoneOperSrcPortZoneSrcPortOtherListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperSrcPortRangeList struct {
	SrcPortRangeStart int                                    `json:"src-port-range-start"`
	SrcPortRangeEnd   int                                    `json:"src-port-range-end"`
	Protocol          string                                 `json:"protocol"`
	Oper              DdosDstZoneOperSrcPortRangeListOper    `json:"oper"`
	PortInd           DdosDstZoneOperSrcPortRangeListPortInd `json:"port-ind"`
}

type DdosDstZoneOperSrcPortRangeListOper struct {
	Ddos_entry_list       []DdosDstZoneOperSrcPortRangeListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                  `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                  `json:"service-displayed-count"`
	ReportingStatus       int                                                  `json:"reporting-status"`
	Sources               int                                                  `json:"sources"`
	SourcesAllEntries     int                                                  `json:"sources-all-entries"`
	SubnetIpAddr          string                                               `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                               `json:"subnet-ipv6-addr"`
	Ipv6                  string                                               `json:"ipv6"`
	HwBlacklisted         int                                                  `json:"hw-blacklisted"`
}

type DdosDstZoneOperSrcPortRangeListOperDdos_entry_list struct {
	DstAddressStr          string `json:"dst-address-str"`
	BwState                string `json:"bw-state"`
	Is_auth_passed         string `json:"is_auth_passed"`
	Level                  int    `json:"level"`
	CurrentConnections     string `json:"current-connections"`
	IsConnectionsExceed    int    `json:"is-connections-exceed"`
	ConnectionLimit        string `json:"connection-limit"`
	CurrentConnectionRate  string `json:"current-connection-rate"`
	IsConnectionRateExceed int    `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string `json:"connection-rate-limit"`
	CurrentPacketRate      string `json:"current-packet-rate"`
	IsPacketRateExceed     int    `json:"is-packet-rate-exceed"`
	PacketRateLimit        string `json:"packet-rate-limit"`
	CurrentKbitRate        string `json:"current-kBit-rate"`
	IsKbitRateExceed       int    `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int    `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string `json:"current-app-stat1"`
	IsAppStat1Exceed       int    `json:"is-app-stat1-exceed"`
	AppStat1Limit          string `json:"app-stat1-limit"`
	CurrentAppStat2        string `json:"current-app-stat2"`
	IsAppStat2Exceed       int    `json:"is-app-stat2-exceed"`
	AppStat2Limit          string `json:"app-stat2-limit"`
	CurrentAppStat3        string `json:"current-app-stat3"`
	IsAppStat3Exceed       int    `json:"is-app-stat3-exceed"`
	AppStat3Limit          string `json:"app-stat3-limit"`
	CurrentAppStat4        string `json:"current-app-stat4"`
	IsAppStat4Exceed       int    `json:"is-app-stat4-exceed"`
	AppStat4Limit          string `json:"app-stat4-limit"`
	CurrentAppStat5        string `json:"current-app-stat5"`
	IsAppStat5Exceed       int    `json:"is-app-stat5-exceed"`
	AppStat5Limit          string `json:"app-stat5-limit"`
	CurrentAppStat6        string `json:"current-app-stat6"`
	IsAppStat6Exceed       int    `json:"is-app-stat6-exceed"`
	AppStat6Limit          string `json:"app-stat6-limit"`
	CurrentAppStat7        string `json:"current-app-stat7"`
	IsAppStat7Exceed       int    `json:"is-app-stat7-exceed"`
	AppStat7Limit          string `json:"app-stat7-limit"`
	CurrentAppStat8        string `json:"current-app-stat8"`
	IsAppStat8Exceed       int    `json:"is-app-stat8-exceed"`
	AppStat8Limit          string `json:"app-stat8-limit"`
	Age                    int    `json:"age"`
	LockupTime             int    `json:"lockup-time"`
	DynamicEntryCount      string `json:"dynamic-entry-count"`
	DynamicEntryLimit      string `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string `json:"dynamic-entry-warn-state"`
	SflowSourceId          int    `json:"sflow-source-id"`
	DebugStr               string `json:"debug-str"`
}

type DdosDstZoneOperSrcPortRangeListPortInd struct {
	Oper DdosDstZoneOperSrcPortRangeListPortIndOper `json:"oper"`
}

type DdosDstZoneOperSrcPortRangeListPortIndOper struct {
	Indicators          []DdosDstZoneOperSrcPortRangeListPortIndOperIndicators `json:"indicators"`
	DetectionDataSource string                                                 `json:"detection-data-source"`
	CurrentLevel        string                                                 `json:"current-level"`
	Details             int                                                    `json:"details"`
}

type DdosDstZoneOperSrcPortRangeListPortIndOperIndicators struct {
	IndicatorName  string                                                             `json:"indicator-name"`
	IndicatorIndex int                                                                `json:"indicator-index"`
	Rate           string                                                             `json:"rate"`
	IndicatorCfg   []DdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
}

type DdosDstZoneOperSrcPortRangeListPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZoneOperTopkDestinations struct {
	Oper DdosDstZoneOperTopkDestinationsOper `json:"oper"`
}

type DdosDstZoneOperTopkDestinationsOper struct {
	Indicators    []DdosDstZoneOperTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZoneOperTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                             `json:"next-indicator"`
	Finished      int                                             `json:"finished"`
	Details       int                                             `json:"details"`
	TopKKey       string                                          `json:"top-k-key"`
}

type DdosDstZoneOperTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                      `json:"indicator-name"`
	IndicatorIndex int                                                         `json:"indicator-index"`
	Destinations   []DdosDstZoneOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZoneOperTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZoneOperTopkDestinationsOperEntryList struct {
	AddressStr string                                                   `json:"address-str"`
	Indicators []DdosDstZoneOperTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZoneOperTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

func (p *DdosDstZoneOper) GetId() string {
	return "1"
}

func (p *DdosDstZoneOper) getPath() string {
	return "ddos/dst/zone/" + p.ZoneName + "/oper"
}

func (p *DdosDstZoneOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneOper, error) {
	logger.Println("DdosDstZoneOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZoneOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
