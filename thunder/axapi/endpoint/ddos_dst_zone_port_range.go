package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type DdosDstZonePortRange struct {
	Inst struct {
		Age int `json:"age" dval:"5"`

		ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`

		CaptureConfig DdosDstZonePortRangeCaptureConfig `json:"capture-config"`

		DefaultActionList string `json:"default-action-list"`

		Deny int `json:"deny"`

		DynamicEntryCountWarnThreshold int `json:"dynamic-entry-count-warn-threshold"`

		DynamicEntryOverflowPolicyList []DdosDstZonePortRangeDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`

		EnableClassListOverflow int `json:"enable-class-list-overflow"`

		EnableTopK int `json:"enable-top-k"`

		EnableTopKDestination int `json:"enable-top-k-destination"`

		FasterDeEscalation int `json:"faster-de-escalation"`

		GlidCfg DdosDstZonePortRangeGlidCfg `json:"glid-cfg"`

		IpFilteringPolicy string `json:"ip-filtering-policy"`

		IpFilteringPolicyStatistics DdosDstZonePortRangeIpFilteringPolicyStatistics232 `json:"ip-filtering-policy-statistics"`

		LevelList []DdosDstZonePortRangeLevelList `json:"level-list"`

		ManualModeEnable int `json:"manual-mode-enable"`

		ManualModeList []DdosDstZonePortRangeManualModeList `json:"manual-mode-list"`

		MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`

		OutboundOnly int `json:"outbound-only"`

		PatternRecognition DdosDstZonePortRangePatternRecognition233 `json:"pattern-recognition"`

		PatternRecognitionPuDetails DdosDstZonePortRangePatternRecognitionPuDetails234 `json:"pattern-recognition-pu-details"`

		PortInd DdosDstZonePortRangePortInd235 `json:"port-ind"`

		PortRangeEnd int `json:"port-range-end"`

		PortRangeStart int `json:"port-range-start"`

		ProgressionTracking DdosDstZonePortRangeProgressionTracking237 `json:"progression-tracking"`

		Protocol string `json:"protocol"`

		SameSourceDestPortDrop int `json:"same-source-dest-port-drop"`

		SetCounterBaseVal int `json:"set-counter-base-val"`

		SflowCommon int `json:"sflow-common"`

		SflowHttp int `json:"sflow-http"`

		SflowIpFilteringPolicy int `json:"sflow-ip-filtering-policy"`

		SflowPackets int `json:"sflow-packets"`

		SflowTcp DdosDstZonePortRangeSflowTcp `json:"sflow-tcp"`

		SrcBasedPolicyList []DdosDstZonePortRangeSrcBasedPolicyList `json:"src-based-policy-list"`

		Stateful int `json:"stateful"`

		TopkDestinations DdosDstZonePortRangeTopkDestinations238 `json:"topk-destinations"`

		TopkDstNumRecords int `json:"topk-dst-num-records" dval:"20"`

		TopkDstSortKey string `json:"topk-dst-sort-key" dval:"avg"`

		TopkNumRecords int `json:"topk-num-records" dval:"20"`

		TopkSortKey string `json:"topk-sort-key" dval:"avg"`

		TopkSources DdosDstZonePortRangeTopkSources239 `json:"topk-sources"`

		UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		Virtualhosts DdosDstZonePortRangeVirtualhosts240 `json:"virtualhosts"`

		ZoneName string
	} `json:"port-range"`
}

type DdosDstZonePortRangeCaptureConfig struct {
	CaptureConfigName string `json:"capture-config-name"`
	CaptureConfigMode string `json:"capture-config-mode"`
}

type DdosDstZonePortRangeDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                         `json:"dummy-name"`
	Glid         string                                                         `json:"glid"`
	Action       string                                                         `json:"action"`
	LogEnable    int                                                            `json:"log-enable"`
	LogPeriodic  int                                                            `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortRangeDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                         `json:"uuid"`
	UserTag      string                                                         `json:"user-tag"`
}

type DdosDstZonePortRangeDynamicEntryOverflowPolicyListZoneTemplate struct {
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

type DdosDstZonePortRangeGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZonePortRangeIpFilteringPolicyStatistics232 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeLevelList struct {
	LevelNum                      string                                       `json:"level-num"`
	SrcDefaultGlid                string                                       `json:"src-default-glid"`
	GlidAction                    string                                       `json:"glid-action"`
	ZoneEscalationScore           int                                          `json:"zone-escalation-score"`
	ZoneViolationActions          string                                       `json:"zone-violation-actions"`
	SrcEscalationScore            int                                          `json:"src-escalation-score"`
	SrcViolationActions           string                                       `json:"src-violation-actions"`
	ZoneTemplate                  DdosDstZonePortRangeLevelListZoneTemplate    `json:"zone-template"`
	CloseSessionsForUnauthSources int                                          `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                          `json:"close-sessions-for-all-sources"`
	ClearSourcesUponDeescalation  int                                          `json:"clear-sources-upon-deescalation"`
	StartPatternRecognition       int                                          `json:"start-pattern-recognition"`
	ApplyExtractedFilters         int                                          `json:"apply-extracted-filters"`
	Uuid                          string                                       `json:"uuid"`
	UserTag                       string                                       `json:"user-tag"`
	IndicatorList                 []DdosDstZonePortRangeLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZonePortRangeLevelListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortRangeLevelListIndicatorList struct {
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

type DdosDstZonePortRangeManualModeList struct {
	Config                        string                                         `json:"config"`
	SrcDefaultGlid                string                                         `json:"src-default-glid"`
	GlidAction                    string                                         `json:"glid-action"`
	ZoneTemplate                  DdosDstZonePortRangeManualModeListZoneTemplate `json:"zone-template"`
	CloseSessionsForUnauthSources int                                            `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                            `json:"close-sessions-for-all-sources"`
	Uuid                          string                                         `json:"uuid"`
	UserTag                       string                                         `json:"user-tag"`
}

type DdosDstZonePortRangeManualModeListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortRangePatternRecognition233 struct {
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

type DdosDstZonePortRangePatternRecognitionPuDetails234 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangePortInd235 struct {
	Uuid           string                                         `json:"uuid"`
	SamplingEnable []DdosDstZonePortRangePortIndSamplingEnable236 `json:"sampling-enable"`
}

type DdosDstZonePortRangePortIndSamplingEnable236 struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortRangeProgressionTracking237 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
}

type DdosDstZonePortRangeSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                      `json:"src-based-policy-name"`
	Uuid                string                                                      `json:"uuid"`
	UserTag             string                                                      `json:"user-tag"`
	PolicyClassListList []DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                 `json:"class-list-name"`
	ClassListGlid                  string                                                                                 `json:"class-list-glid"`
	Glid                           string                                                                                 `json:"glid"`
	GlidAction                     string                                                                                 `json:"glid-action"`
	Action                         string                                                                                 `json:"action"`
	MaxDynamicEntryCount           int                                                                                    `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                    `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                 `json:"uuid"`
	UserTag                        string                                                                                 `json:"user-tag"`
	SamplingEnable                 []DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
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

type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                           `json:"dummy-name"`
	Glid         string                                                                                           `json:"glid"`
	Action       string                                                                                           `json:"action"`
	LogEnable    int                                                                                              `json:"log-enable"`
	LogPeriodic  int                                                                                              `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                           `json:"uuid"`
	UserTag      string                                                                                           `json:"user-tag"`
}

type DdosDstZonePortRangeSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
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

type DdosDstZonePortRangeTopkDestinations238 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeTopkSources239 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortRangeVirtualhosts240 struct {
	VhostsConfig      string                                               `json:"vhosts-config"`
	SourceTrackingAll int                                                  `json:"source-tracking-all"`
	Uuid              string                                               `json:"uuid"`
	VirtualhostList   []DdosDstZonePortRangeVirtualhostsVirtualhostList241 `json:"virtualhost-list"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostList241 struct {
	Vhost              string                                                         `json:"vhost"`
	Servername         []DdosDstZonePortRangeVirtualhostsVirtualhostListServername242 `json:"servername"`
	ServernameList     string                                                         `json:"servername-list"`
	ServernameMatchAny int                                                            `json:"servername-match-any"`
	SourceTracking     string                                                         `json:"source-tracking" dval:"follow"`
	GlidCfg            DdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg243      `json:"glid-cfg"`
	Deny               int                                                            `json:"deny"`
	Uuid               string                                                         `json:"uuid"`
	UserTag            string                                                         `json:"user-tag"`
	LevelList          []DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList244  `json:"level-list"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListServername242 struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListGlidCfg243 struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListLevelList244 struct {
	LevelNum       string                                                                  `json:"level-num"`
	SrcDefaultGlid string                                                                  `json:"src-default-glid"`
	GlidAction     string                                                                  `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate245 `json:"zone-template"`
	Uuid           string                                                                  `json:"uuid"`
	UserTag        string                                                                  `json:"user-tag"`
}

type DdosDstZonePortRangeVirtualhostsVirtualhostListLevelListZoneTemplate245 struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

func (p *DdosDstZonePortRange) GetId() string {
	return strconv.Itoa(p.Inst.PortRangeStart) + "+" + strconv.Itoa(p.Inst.PortRangeEnd) + "+" + p.Inst.Protocol
}

func (p *DdosDstZonePortRange) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port-range"
}

func (p *DdosDstZonePortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRange::Post")
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

func (p *DdosDstZonePortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRange::Get")
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
func (p *DdosDstZonePortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRange::Put")
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

func (p *DdosDstZonePortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortRange::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
