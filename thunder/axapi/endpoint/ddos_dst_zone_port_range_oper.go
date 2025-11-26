package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type DdosDstZonePortRangeOper struct {
	IpFilteringPolicyStatistics DdosDstZonePortRangeOperIpFilteringPolicyStatistics `json:"ip-filtering-policy-statistics"`

	Oper DdosDstZonePortRangeOperOper `json:"oper"`

	PatternRecognition DdosDstZonePortRangeOperPatternRecognition `json:"pattern-recognition"`

	PatternRecognitionPuDetails DdosDstZonePortRangeOperPatternRecognitionPuDetails `json:"pattern-recognition-pu-details"`

	PortInd DdosDstZonePortRangeOperPortInd `json:"port-ind"`

	PortRangeEnd int `json:"port-range-end"`

	PortRangeStart int `json:"port-range-start"`

	ProgressionTracking DdosDstZonePortRangeOperProgressionTracking `json:"progression-tracking"`

	Protocol string `json:"protocol"`

	SrcBasedPolicyList []DdosDstZonePortRangeOperSrcBasedPolicyList `json:"src-based-policy-list"`

	TopkDestinations DdosDstZonePortRangeOperTopkDestinations `json:"topk-destinations"`

	TopkSources DdosDstZonePortRangeOperTopkSources `json:"topk-sources"`

	Virtualhosts DdosDstZonePortRangeOperVirtualhosts `json:"virtualhosts"`

	ZoneName string
}
type DataDdosDstZonePortRangeOper struct {
	DtDdosDstZonePortRangeOper DdosDstZonePortRangeOper `json:"port-range"`
}

type DdosDstZonePortRangeOperIpFilteringPolicyStatistics struct {
	Oper DdosDstZonePortRangeOperIpFilteringPolicyStatisticsOper `json:"oper"`
}

type DdosDstZonePortRangeOperIpFilteringPolicyStatisticsOper struct {
	RuleList []DdosDstZonePortRangeOperIpFilteringPolicyStatisticsOperRuleList `json:"rule-list"`
}

type DdosDstZonePortRangeOperIpFilteringPolicyStatisticsOperRuleList struct {
	Seq                   int `json:"seq"`
	Hits                  int `json:"hits"`
	Blacklisted_src_count int `json:"blacklisted_src_count"`
}

type DdosDstZonePortRangeOperOper struct {
	Ddos_entry_list       []DdosDstZonePortRangeOperOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                           `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                           `json:"service-displayed-count"`
	ReportingStatus       int                                           `json:"reporting-status"`
	Sources               int                                           `json:"sources"`
	OverflowPolicy        int                                           `json:"overflow-policy"`
	SourcesAllEntries     int                                           `json:"sources-all-entries"`
	ClassList             string                                        `json:"class-list"`
	SubnetIpAddr          string                                        `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                        `json:"subnet-ipv6-addr"`
	Ipv6                  string                                        `json:"ipv6"`
	Exceeded              int                                           `json:"exceeded"`
	BlackListed           int                                           `json:"black-listed"`
	WhiteListed           int                                           `json:"white-listed"`
	Authenticated         int                                           `json:"authenticated"`
	Level                 int                                           `json:"level"`
	AppStat               int                                           `json:"app-stat"`
	Indicators            int                                           `json:"indicators"`
	IndicatorDetail       int                                           `json:"indicator-detail"`
	L4ExtRate             int                                           `json:"l4-ext-rate"`
	HwBlacklisted         int                                           `json:"hw-blacklisted"`
	SuffixRequestRate     int                                           `json:"suffix-request-rate"`
	DomainName            string                                        `json:"domain-name"`
}

type DdosDstZonePortRangeOperOperDdos_entry_list struct {
	DstAddressStr          string                                                         `json:"dst-address-str"`
	BwState                string                                                         `json:"bw-state"`
	Is_auth_passed         string                                                         `json:"is_auth_passed"`
	Level                  int                                                            `json:"level"`
	BlReasoningRcode       string                                                         `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string                                                         `json:"bl-reasoning-timestamp"`
	CurrentConnections     string                                                         `json:"current-connections"`
	IsConnectionsExceed    int                                                            `json:"is-connections-exceed"`
	ConnectionLimit        string                                                         `json:"connection-limit"`
	CurrentConnectionRate  string                                                         `json:"current-connection-rate"`
	IsConnectionRateExceed int                                                            `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string                                                         `json:"connection-rate-limit"`
	CurrentPacketRate      string                                                         `json:"current-packet-rate"`
	IsPacketRateExceed     int                                                            `json:"is-packet-rate-exceed"`
	PacketRateLimit        string                                                         `json:"packet-rate-limit"`
	CurrentKbitRate        string                                                         `json:"current-kBit-rate"`
	IsKbitRateExceed       int                                                            `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string                                                         `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string                                                         `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int                                                            `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string                                                         `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string                                                         `json:"current-app-stat1"`
	IsAppStat1Exceed       int                                                            `json:"is-app-stat1-exceed"`
	AppStat1Limit          string                                                         `json:"app-stat1-limit"`
	CurrentAppStat2        string                                                         `json:"current-app-stat2"`
	IsAppStat2Exceed       int                                                            `json:"is-app-stat2-exceed"`
	AppStat2Limit          string                                                         `json:"app-stat2-limit"`
	CurrentAppStat3        string                                                         `json:"current-app-stat3"`
	IsAppStat3Exceed       int                                                            `json:"is-app-stat3-exceed"`
	AppStat3Limit          string                                                         `json:"app-stat3-limit"`
	CurrentAppStat4        string                                                         `json:"current-app-stat4"`
	IsAppStat4Exceed       int                                                            `json:"is-app-stat4-exceed"`
	AppStat4Limit          string                                                         `json:"app-stat4-limit"`
	CurrentAppStat5        string                                                         `json:"current-app-stat5"`
	IsAppStat5Exceed       int                                                            `json:"is-app-stat5-exceed"`
	AppStat5Limit          string                                                         `json:"app-stat5-limit"`
	CurrentAppStat6        string                                                         `json:"current-app-stat6"`
	IsAppStat6Exceed       int                                                            `json:"is-app-stat6-exceed"`
	AppStat6Limit          string                                                         `json:"app-stat6-limit"`
	CurrentAppStat7        string                                                         `json:"current-app-stat7"`
	IsAppStat7Exceed       int                                                            `json:"is-app-stat7-exceed"`
	AppStat7Limit          string                                                         `json:"app-stat7-limit"`
	CurrentAppStat8        string                                                         `json:"current-app-stat8"`
	IsAppStat8Exceed       int                                                            `json:"is-app-stat8-exceed"`
	AppStat8Limit          string                                                         `json:"app-stat8-limit"`
	Age                    int                                                            `json:"age"`
	LockupTime             int                                                            `json:"lockup-time"`
	DynamicEntryCount      string                                                         `json:"dynamic-entry-count"`
	DynamicEntryLimit      string                                                         `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string                                                         `json:"dynamic-entry-warn-state"`
	SflowSourceId          int                                                            `json:"sflow-source-id"`
	HttpFilterRates        []DdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates   `json:"http-filter-rates"`
	ResponseSizeRates      []DdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates `json:"response-size-rates"`
	DebugStr               string                                                         `json:"debug-str"`
}

type DdosDstZonePortRangeOperOperDdos_entry_listHttpFilterRates struct {
	HttpFilterRateName          string `json:"http-filter-rate-name"`
	IsHttpFilterRateLimitExceed int    `json:"is-http-filter-rate-limit-exceed"`
	CurrentHttpFilterRate       string `json:"current-http-filter-rate"`
	HttpFilterRateLimit         string `json:"http-filter-rate-limit"`
}

type DdosDstZonePortRangeOperOperDdos_entry_listResponseSizeRates struct {
	ResponseSizeRateName          string `json:"response-size-rate-name"`
	IsResponseSizeRateLimitExceed int    `json:"is-response-size-rate-limit-exceed"`
	CurrentResponseSizeRate       string `json:"current-response-size-rate"`
	ResponseSizeRateLimit         string `json:"response-size-rate-limit"`
}

type DdosDstZonePortRangeOperPatternRecognition struct {
	Oper DdosDstZonePortRangeOperPatternRecognitionOper `json:"oper"`
}

type DdosDstZonePortRangeOperPatternRecognitionOper struct {
	State            string                                                     `json:"state"`
	Timestamp        string                                                     `json:"timestamp"`
	PeacePktCount    int                                                        `json:"peace-pkt-count"`
	WarPktCount      int                                                        `json:"war-pkt-count"`
	WarPktPercentage int                                                        `json:"war-pkt-percentage"`
	FilterThreshold  int                                                        `json:"filter-threshold"`
	FilterCount      int                                                        `json:"filter-count"`
	FilterList       []DdosDstZonePortRangeOperPatternRecognitionOperFilterList `json:"filter-list"`
}

type DdosDstZonePortRangeOperPatternRecognitionOperFilterList struct {
	ProcessingUnit string `json:"processing-unit"`
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZonePortRangeOperPatternRecognitionPuDetails struct {
	Oper DdosDstZonePortRangeOperPatternRecognitionPuDetailsOper `json:"oper"`
}

type DdosDstZonePortRangeOperPatternRecognitionPuDetailsOper struct {
	AllFilters []DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters `json:"all-filters"`
}

type DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFilters struct {
	ProcessingUnit   string                                                                        `json:"processing-unit"`
	State            string                                                                        `json:"state"`
	Timestamp        string                                                                        `json:"timestamp"`
	PeacePktCount    int                                                                           `json:"peace-pkt-count"`
	WarPktCount      int                                                                           `json:"war-pkt-count"`
	WarPktPercentage int                                                                           `json:"war-pkt-percentage"`
	FilterThreshold  int                                                                           `json:"filter-threshold"`
	FilterCount      int                                                                           `json:"filter-count"`
	FilterList       []DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList `json:"filter-list"`
}

type DdosDstZonePortRangeOperPatternRecognitionPuDetailsOperAllFiltersFilterList struct {
	FilterEnabled  int    `json:"filter-enabled"`
	HardwareFilter int    `json:"hardware-filter"`
	FilterExpr     string `json:"filter-expr"`
	FilterDesc     string `json:"filter-desc"`
	SampleRatio    int    `json:"sample-ratio"`
}

type DdosDstZonePortRangeOperPortInd struct {
	Oper DdosDstZonePortRangeOperPortIndOper `json:"oper"`
}

type DdosDstZonePortRangeOperPortIndOper struct {
	SrcEntryList        []DdosDstZonePortRangeOperPortIndOperSrcEntryList `json:"src-entry-list"`
	Indicators          []DdosDstZonePortRangeOperPortIndOperIndicators   `json:"indicators"`
	DetectionDataSource string                                            `json:"detection-data-source"`
	TotalScore          string                                            `json:"total-score"`
	CurrentLevel        string                                            `json:"current-level"`
	EscalationTimestamp string                                            `json:"escalation-timestamp"`
	InitialLearning     string                                            `json:"initial-learning"`
	ActiveTime          int                                               `json:"active-time"`
	SourcesAllEntries   int                                               `json:"sources-all-entries"`
	SubnetIpAddr        string                                            `json:"subnet-ip-addr"`
	SubnetIpv6Addr      string                                            `json:"subnet-ipv6-addr"`
	Ipv6                string                                            `json:"ipv6"`
	Details             int                                               `json:"details"`
	Sources             int                                               `json:"sources"`
}

type DdosDstZonePortRangeOperPortIndOperSrcEntryList struct {
	SrcAddressStr       string                                                      `json:"src-address-str"`
	Indicators          []DdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators `json:"indicators"`
	DetectionDataSource string                                                      `json:"detection-data-source"`
	TotalScore          string                                                      `json:"total-score"`
	CurrentLevel        string                                                      `json:"current-level"`
	SrcLevel            string                                                      `json:"src-level"`
	EscalationTimestamp string                                                      `json:"escalation-timestamp"`
	InitialLearning     string                                                      `json:"initial-learning"`
	ActiveTime          int                                                         `json:"active-time"`
}

type DdosDstZonePortRangeOperPortIndOperSrcEntryListIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	Rate              string `json:"rate"`
	SrcMaximum        string `json:"src-maximum"`
	SrcMinimum        string `json:"src-minimum"`
	SrcNonZeroMinimum string `json:"src-non-zero-minimum"`
	SrcAverage        string `json:"src-average"`
	Score             string `json:"score"`
}

type DdosDstZonePortRangeOperPortIndOperIndicators struct {
	IndicatorName         string                                                      `json:"indicator-name"`
	IndicatorIndex        int                                                         `json:"indicator-index"`
	Rate                  string                                                      `json:"rate"`
	ZoneMaximum           string                                                      `json:"zone-maximum"`
	ZoneMinimum           string                                                      `json:"zone-minimum"`
	ZoneNonZeroMinimum    string                                                      `json:"zone-non-zero-minimum"`
	ZoneAverage           string                                                      `json:"zone-average"`
	ZoneAdaptiveThreshold string                                                      `json:"zone-adaptive-threshold"`
	SrcMaximum            string                                                      `json:"src-maximum"`
	IndicatorCfg          []DdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg `json:"indicator-cfg"`
	Score                 string                                                      `json:"score"`
}

type DdosDstZonePortRangeOperPortIndOperIndicatorsIndicatorCfg struct {
	Level           int    `json:"level"`
	ZoneThreshold   string `json:"zone-threshold"`
	SourceThreshold string `json:"source-threshold"`
}

type DdosDstZonePortRangeOperProgressionTracking struct {
	Oper DdosDstZonePortRangeOperProgressionTrackingOper `json:"oper"`
}

type DdosDstZonePortRangeOperProgressionTrackingOper struct {
	Indicators          []DdosDstZonePortRangeOperProgressionTrackingOperIndicators `json:"indicators"`
	LearningDetails     int                                                         `json:"learning-details"`
	LearningBrief       int                                                         `json:"learning-brief"`
	RecommendedTemplate int                                                         `json:"recommended-template"`
	TemplateDebugTable  int                                                         `json:"template-debug-table"`
}

type DdosDstZonePortRangeOperProgressionTrackingOperIndicators struct {
	IndicatorName     string `json:"indicator-name"`
	IndicatorIndex    int    `json:"indicator-index"`
	NumSample         int    `json:"num-sample"`
	Average           string `json:"average"`
	Maximum           string `json:"maximum"`
	Minimum           string `json:"minimum"`
	StandardDeviation string `json:"standard-deviation"`
}

type DdosDstZonePortRangeOperSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                          `json:"src-based-policy-name"`
	Oper                DdosDstZonePortRangeOperSrcBasedPolicyListOper                  `json:"oper"`
	PolicyClassListList []DdosDstZonePortRangeOperSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortRangeOperSrcBasedPolicyListOper struct {
}

type DdosDstZonePortRangeOperSrcBasedPolicyListPolicyClassListList struct {
	ClassListName string                                                            `json:"class-list-name"`
	Oper          DdosDstZonePortRangeOperSrcBasedPolicyListPolicyClassListListOper `json:"oper"`
}

type DdosDstZonePortRangeOperSrcBasedPolicyListPolicyClassListListOper struct {
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

type DdosDstZonePortRangeOperTopkDestinations struct {
	Oper DdosDstZonePortRangeOperTopkDestinationsOper `json:"oper"`
}

type DdosDstZonePortRangeOperTopkDestinationsOper struct {
	Indicators    []DdosDstZonePortRangeOperTopkDestinationsOperIndicators `json:"indicators"`
	EntryList     []DdosDstZonePortRangeOperTopkDestinationsOperEntryList  `json:"entry-list"`
	NextIndicator int                                                      `json:"next-indicator"`
	Finished      int                                                      `json:"finished"`
	Details       int                                                      `json:"details"`
	TopKKey       string                                                   `json:"top-k-key"`
}

type DdosDstZonePortRangeOperTopkDestinationsOperIndicators struct {
	IndicatorName  string                                                               `json:"indicator-name"`
	IndicatorIndex int                                                                  `json:"indicator-index"`
	Destinations   []DdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations `json:"destinations"`
}

type DdosDstZonePortRangeOperTopkDestinationsOperIndicatorsDestinations struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZonePortRangeOperTopkDestinationsOperEntryList struct {
	AddressStr string                                                            `json:"address-str"`
	Indicators []DdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators `json:"indicators"`
}

type DdosDstZonePortRangeOperTopkDestinationsOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZonePortRangeOperTopkSources struct {
	Oper DdosDstZonePortRangeOperTopkSourcesOper `json:"oper"`
}

type DdosDstZonePortRangeOperTopkSourcesOper struct {
	Indicators    []DdosDstZonePortRangeOperTopkSourcesOperIndicators `json:"indicators"`
	EntryList     []DdosDstZonePortRangeOperTopkSourcesOperEntryList  `json:"entry-list"`
	NextIndicator int                                                 `json:"next-indicator"`
	Finished      int                                                 `json:"finished"`
	Details       int                                                 `json:"details"`
	TopKKey       string                                              `json:"top-k-key"`
}

type DdosDstZonePortRangeOperTopkSourcesOperIndicators struct {
	IndicatorName  string                                                     `json:"indicator-name"`
	IndicatorIndex int                                                        `json:"indicator-index"`
	Sources        []DdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources `json:"sources"`
}

type DdosDstZonePortRangeOperTopkSourcesOperIndicatorsSources struct {
	Address string `json:"address"`
	Rate    string `json:"rate"`
}

type DdosDstZonePortRangeOperTopkSourcesOperEntryList struct {
	AddressStr string                                                       `json:"address-str"`
	Indicators []DdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators `json:"indicators"`
}

type DdosDstZonePortRangeOperTopkSourcesOperEntryListIndicators struct {
	IndicatorName  string `json:"indicator-name"`
	IndicatorIndex int    `json:"indicator-index"`
	Rate           string `json:"rate"`
	MaxPeak        string `json:"max-peak"`
	PsdWdwCnt      int    `json:"psd-wdw-cnt"`
}

type DdosDstZonePortRangeOperVirtualhosts struct {
	Oper            DdosDstZonePortRangeOperVirtualhostsOper              `json:"oper"`
	VirtualhostList []DdosDstZonePortRangeOperVirtualhostsVirtualhostList `json:"virtualhost-list"`
}

type DdosDstZonePortRangeOperVirtualhostsOper struct {
}

type DdosDstZonePortRangeOperVirtualhostsVirtualhostList struct {
	Vhost string                                                  `json:"vhost"`
	Oper  DdosDstZonePortRangeOperVirtualhostsVirtualhostListOper `json:"oper"`
}

type DdosDstZonePortRangeOperVirtualhostsVirtualhostListOper struct {
	Ddos_entry_list       []DdosDstZonePortRangeOperVirtualhostsVirtualhostListOperDdos_entry_list `json:"ddos_entry_list"`
	EntryDisplayedCount   int                                                                      `json:"entry-displayed-count"`
	ServiceDisplayedCount int                                                                      `json:"service-displayed-count"`
	ReportingStatus       int                                                                      `json:"reporting-status"`
	Sources               int                                                                      `json:"sources"`
	OverflowPolicy        int                                                                      `json:"overflow-policy"`
	SourcesAllEntries     int                                                                      `json:"sources-all-entries"`
	ClassList             string                                                                   `json:"class-list"`
	SubnetIpAddr          string                                                                   `json:"subnet-ip-addr"`
	SubnetIpv6Addr        string                                                                   `json:"subnet-ipv6-addr"`
	Ipv6                  string                                                                   `json:"ipv6"`
	Exceeded              int                                                                      `json:"exceeded"`
	BlackListed           int                                                                      `json:"black-listed"`
	WhiteListed           int                                                                      `json:"white-listed"`
	Authenticated         int                                                                      `json:"authenticated"`
	Level                 int                                                                      `json:"level"`
	AppStat               int                                                                      `json:"app-stat"`
	Indicators            int                                                                      `json:"indicators"`
	IndicatorDetail       int                                                                      `json:"indicator-detail"`
	L4ExtRate             int                                                                      `json:"l4-ext-rate"`
	HwBlacklisted         int                                                                      `json:"hw-blacklisted"`
	SuffixRequestRate     int                                                                      `json:"suffix-request-rate"`
	DomainName            string                                                                   `json:"domain-name"`
}

type DdosDstZonePortRangeOperVirtualhostsVirtualhostListOperDdos_entry_list struct {
	DstAddressStr          string                                                                                    `json:"dst-address-str"`
	BwState                string                                                                                    `json:"bw-state"`
	Is_auth_passed         string                                                                                    `json:"is_auth_passed"`
	Level                  int                                                                                       `json:"level"`
	BlReasoningRcode       string                                                                                    `json:"bl-reasoning-rcode"`
	BlReasoningTimestamp   string                                                                                    `json:"bl-reasoning-timestamp"`
	CurrentConnections     string                                                                                    `json:"current-connections"`
	IsConnectionsExceed    int                                                                                       `json:"is-connections-exceed"`
	ConnectionLimit        string                                                                                    `json:"connection-limit"`
	CurrentConnectionRate  string                                                                                    `json:"current-connection-rate"`
	IsConnectionRateExceed int                                                                                       `json:"is-connection-rate-exceed"`
	ConnectionRateLimit    string                                                                                    `json:"connection-rate-limit"`
	CurrentPacketRate      string                                                                                    `json:"current-packet-rate"`
	IsPacketRateExceed     int                                                                                       `json:"is-packet-rate-exceed"`
	PacketRateLimit        string                                                                                    `json:"packet-rate-limit"`
	CurrentKbitRate        string                                                                                    `json:"current-kBit-rate"`
	IsKbitRateExceed       int                                                                                       `json:"is-kBit-rate-exceed"`
	KbitRateLimit          string                                                                                    `json:"kBit-rate-limit"`
	CurrentFragPacketRate  string                                                                                    `json:"current-frag-packet-rate"`
	IsFragPacketRateExceed int                                                                                       `json:"is-frag-packet-rate-exceed"`
	FragPacketRateLimit    string                                                                                    `json:"frag-packet-rate-limit"`
	CurrentAppStat1        string                                                                                    `json:"current-app-stat1"`
	IsAppStat1Exceed       int                                                                                       `json:"is-app-stat1-exceed"`
	AppStat1Limit          string                                                                                    `json:"app-stat1-limit"`
	CurrentAppStat2        string                                                                                    `json:"current-app-stat2"`
	IsAppStat2Exceed       int                                                                                       `json:"is-app-stat2-exceed"`
	AppStat2Limit          string                                                                                    `json:"app-stat2-limit"`
	CurrentAppStat3        string                                                                                    `json:"current-app-stat3"`
	IsAppStat3Exceed       int                                                                                       `json:"is-app-stat3-exceed"`
	AppStat3Limit          string                                                                                    `json:"app-stat3-limit"`
	CurrentAppStat4        string                                                                                    `json:"current-app-stat4"`
	IsAppStat4Exceed       int                                                                                       `json:"is-app-stat4-exceed"`
	AppStat4Limit          string                                                                                    `json:"app-stat4-limit"`
	CurrentAppStat5        string                                                                                    `json:"current-app-stat5"`
	IsAppStat5Exceed       int                                                                                       `json:"is-app-stat5-exceed"`
	AppStat5Limit          string                                                                                    `json:"app-stat5-limit"`
	CurrentAppStat6        string                                                                                    `json:"current-app-stat6"`
	IsAppStat6Exceed       int                                                                                       `json:"is-app-stat6-exceed"`
	AppStat6Limit          string                                                                                    `json:"app-stat6-limit"`
	CurrentAppStat7        string                                                                                    `json:"current-app-stat7"`
	IsAppStat7Exceed       int                                                                                       `json:"is-app-stat7-exceed"`
	AppStat7Limit          string                                                                                    `json:"app-stat7-limit"`
	CurrentAppStat8        string                                                                                    `json:"current-app-stat8"`
	IsAppStat8Exceed       int                                                                                       `json:"is-app-stat8-exceed"`
	AppStat8Limit          string                                                                                    `json:"app-stat8-limit"`
	Age                    int                                                                                       `json:"age"`
	LockupTime             int                                                                                       `json:"lockup-time"`
	DynamicEntryCount      string                                                                                    `json:"dynamic-entry-count"`
	DynamicEntryLimit      string                                                                                    `json:"dynamic-entry-limit"`
	DynamicEntryWarnState  string                                                                                    `json:"dynamic-entry-warn-state"`
	SflowSourceId          int                                                                                       `json:"sflow-source-id"`
	HttpFilterRates        []DdosDstZonePortRangeOperVirtualhostsVirtualhostListOperDdos_entry_listHttpFilterRates   `json:"http-filter-rates"`
	ResponseSizeRates      []DdosDstZonePortRangeOperVirtualhostsVirtualhostListOperDdos_entry_listResponseSizeRates `json:"response-size-rates"`
	DebugStr               string                                                                                    `json:"debug-str"`
}

type DdosDstZonePortRangeOperVirtualhostsVirtualhostListOperDdos_entry_listHttpFilterRates struct {
	HttpFilterRateName          string `json:"http-filter-rate-name"`
	IsHttpFilterRateLimitExceed int    `json:"is-http-filter-rate-limit-exceed"`
	CurrentHttpFilterRate       string `json:"current-http-filter-rate"`
	HttpFilterRateLimit         string `json:"http-filter-rate-limit"`
}

type DdosDstZonePortRangeOperVirtualhostsVirtualhostListOperDdos_entry_listResponseSizeRates struct {
	ResponseSizeRateName          string `json:"response-size-rate-name"`
	IsResponseSizeRateLimitExceed int    `json:"is-response-size-rate-limit-exceed"`
	CurrentResponseSizeRate       string `json:"current-response-size-rate"`
	ResponseSizeRateLimit         string `json:"response-size-rate-limit"`
}

func (p *DdosDstZonePortRangeOper) GetId() string {
	return "1"
}

func (p *DdosDstZonePortRangeOper) getPath() string {

	return "ddos/dst/zone/" + p.ZoneName + "/port-range/" + strconv.Itoa(p.PortRangeStart) + "+" + strconv.Itoa(p.PortRangeEnd) + "+" + p.Protocol + "/oper"
}

func (p *DdosDstZonePortRangeOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeOper, error) {
	logger.Println("DdosDstZonePortRangeOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataDdosDstZonePortRangeOper
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
