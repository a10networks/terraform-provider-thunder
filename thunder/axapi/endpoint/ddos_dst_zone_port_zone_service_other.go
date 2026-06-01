package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDstZonePortZoneServiceOther struct {
	Inst struct {
		Age int `json:"age" dval:"5"`

		ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`

		DefaultActionList string `json:"default-action-list"`

		Deny int `json:"deny"`

		DynamicEntryCountWarnThreshold int `json:"dynamic-entry-count-warn-threshold"`

		DynamicEntryOverflowPolicyList []DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`

		EnableClassListOverflow int `json:"enable-class-list-overflow"`

		EnableTopK int `json:"enable-top-k"`

		EnableTopKDestination int `json:"enable-top-k-destination"`

		FasterDeEscalation int `json:"faster-de-escalation"`

		GlidCfg DdosDstZonePortZoneServiceOtherGlidCfg `json:"glid-cfg"`

		IpFilteringPolicy string `json:"ip-filtering-policy"`

		IpFilteringPolicyStatistics DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatistics246 `json:"ip-filtering-policy-statistics"`

		LevelList []DdosDstZonePortZoneServiceOtherLevelList `json:"level-list"`

		LogSrcDefaultEnable int `json:"log-src-default-enable"`

		ManualModeEnable int `json:"manual-mode-enable"`

		ManualModeList []DdosDstZonePortZoneServiceOtherManualModeList `json:"manual-mode-list"`

		MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`

		OutboundOnly int `json:"outbound-only"`

		PatternRecognition DdosDstZonePortZoneServiceOtherPatternRecognition247 `json:"pattern-recognition"`

		PatternRecognitionPuDetails DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetails248 `json:"pattern-recognition-pu-details"`

		PortInd DdosDstZonePortZoneServiceOtherPortInd249 `json:"port-ind"`

		PortOther string `json:"port-other"`

		ProgressionTracking DdosDstZonePortZoneServiceOtherProgressionTracking251 `json:"progression-tracking"`

		Protocol string `json:"protocol"`

		SameSourceDestPortDrop int `json:"same-source-dest-port-drop"`

		SetCounterBaseVal int `json:"set-counter-base-val"`

		SflowCommon int `json:"sflow-common"`

		SflowIpFilteringPolicy int `json:"sflow-ip-filtering-policy"`

		SflowPackets int `json:"sflow-packets"`

		SflowTcp DdosDstZonePortZoneServiceOtherSflowTcp `json:"sflow-tcp"`

		SrcBasedPolicyList []DdosDstZonePortZoneServiceOtherSrcBasedPolicyList `json:"src-based-policy-list"`

		Stateful int `json:"stateful"`

		TopkDestinations DdosDstZonePortZoneServiceOtherTopkDestinations252 `json:"topk-destinations"`

		TopkDstNumRecords int `json:"topk-dst-num-records" dval:"20"`

		TopkDstSortKey string `json:"topk-dst-sort-key" dval:"avg"`

		TopkNumRecords int `json:"topk-num-records" dval:"20"`

		TopkSortKey string `json:"topk-sort-key" dval:"avg"`

		TopkSources DdosDstZonePortZoneServiceOtherTopkSources253 `json:"topk-sources"`

		UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`

		Uuid string `json:"uuid"`

		ZoneName string
	} `json:"zone-service-other"`
}

type DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                                    `json:"dummy-name"`
	Glid         string                                                                    `json:"glid"`
	Action       string                                                                    `json:"action"`
	LogEnable    int                                                                       `json:"log-enable"`
	LogPeriodic  int                                                                       `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                    `json:"uuid"`
	UserTag      string                                                                    `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicyListZoneTemplate struct {
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceOtherGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZonePortZoneServiceOtherIpFilteringPolicyStatistics246 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherLevelList struct {
	LevelNum                      string                                                  `json:"level-num"`
	SrcDefaultGlid                string                                                  `json:"src-default-glid"`
	GlidAction                    string                                                  `json:"glid-action"`
	ZoneEscalationScore           int                                                     `json:"zone-escalation-score"`
	ZoneViolationActions          string                                                  `json:"zone-violation-actions"`
	SrcEscalationScore            int                                                     `json:"src-escalation-score"`
	SrcViolationActions           string                                                  `json:"src-violation-actions"`
	ZoneTemplate                  DdosDstZonePortZoneServiceOtherLevelListZoneTemplate    `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                     `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                     `json:"close-sessions-for-all-sources"`
	ClearSourcesUponDeescalation  int                                                     `json:"clear-sources-upon-deescalation"`
	StartPatternRecognition       int                                                     `json:"start-pattern-recognition"`
	ApplyExtractedFilters         int                                                     `json:"apply-extracted-filters"`
	Uuid                          string                                                  `json:"uuid"`
	UserTag                       string                                                  `json:"user-tag"`
	IndicatorList                 []DdosDstZonePortZoneServiceOtherLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZonePortZoneServiceOtherLevelListZoneTemplate struct {
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceOtherLevelListIndicatorList struct {
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

type DdosDstZonePortZoneServiceOtherManualModeList struct {
	Config                        string                                                    `json:"config"`
	SrcDefaultGlid                string                                                    `json:"src-default-glid"`
	GlidAction                    string                                                    `json:"glid-action"`
	ZoneTemplate                  DdosDstZonePortZoneServiceOtherManualModeListZoneTemplate `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                       `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                       `json:"close-sessions-for-all-sources"`
	Uuid                          string                                                    `json:"uuid"`
	UserTag                       string                                                    `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherManualModeListZoneTemplate struct {
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceOtherPatternRecognition247 struct {
	Algorithm               string `json:"algorithm"`
	Mode                    string `json:"mode"`
	Sensitivity             string `json:"sensitivity"`
	FilterThreshold         int    `json:"filter-threshold"`
	FilterInactiveThreshold int    `json:"filter-inactive-threshold"`
	TriggeredBy             string `json:"triggered-by"`
	CaptureTraffic          string `json:"capture-traffic"`
	Uuid                    string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherPatternRecognitionPuDetails248 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherPortInd249 struct {
	Uuid           string                                                    `json:"uuid"`
	SamplingEnable []DdosDstZonePortZoneServiceOtherPortIndSamplingEnable250 `json:"sampling-enable"`
}

type DdosDstZonePortZoneServiceOtherPortIndSamplingEnable250 struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceOtherProgressionTracking251 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                                 `json:"src-based-policy-name"`
	Uuid                string                                                                 `json:"uuid"`
	UserTag             string                                                                 `json:"user-tag"`
	PolicyClassListList []DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                            `json:"class-list-name"`
	ClassListGlid                  string                                                                                            `json:"class-list-glid"`
	Glid                           string                                                                                            `json:"glid"`
	GlidAction                     string                                                                                            `json:"glid-action"`
	Action                         string                                                                                            `json:"action"`
	LogEnable                      int                                                                                               `json:"log-enable"`
	MaxDynamicEntryCount           int                                                                                               `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                               `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                            `json:"uuid"`
	UserTag                        string                                                                                            `json:"user-tag"`
	SamplingEnable                 []DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                                      `json:"dummy-name"`
	Glid         string                                                                                                      `json:"glid"`
	Action       string                                                                                                      `json:"action"`
	LogEnable    int                                                                                                         `json:"log-enable"`
	LogPeriodic  int                                                                                                         `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                                      `json:"uuid"`
	UserTag      string                                                                                                      `json:"user-tag"`
}

type DdosDstZonePortZoneServiceOtherSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
	Dns     string `json:"dns"`
	Http    string `json:"http"`
	SslL4   string `json:"ssl-l4"`
	Sip     string `json:"sip"`
	Tcp     string `json:"tcp"`
	Udp     string `json:"udp"`
	Encap   string `json:"encap"`
	Logging string `json:"logging"`
}

type DdosDstZonePortZoneServiceOtherTopkDestinations252 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceOtherTopkSources253 struct {
	Uuid string `json:"uuid"`
}

func (p *DdosDstZonePortZoneServiceOther) GetId() string {
	return p.Inst.PortOther + "+" + p.Inst.Protocol
}

func (p *DdosDstZonePortZoneServiceOther) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port/zone-service-other"
}

func (p *DdosDstZonePortZoneServiceOther) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceOther::Post")
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

func (p *DdosDstZonePortZoneServiceOther) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceOther::Get")
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
func (p *DdosDstZonePortZoneServiceOther) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceOther::Put")
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

func (p *DdosDstZonePortZoneServiceOther) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneServiceOther::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
