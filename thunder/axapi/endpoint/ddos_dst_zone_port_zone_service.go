package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type DdosDstZonePortZoneService struct {
	Inst struct {
		Age int `json:"age" dval:"5"`

		ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`

		CaptureConfig DdosDstZonePortZoneServiceCaptureConfig `json:"capture-config"`

		DefaultActionList string `json:"default-action-list"`

		Deny int `json:"deny"`

		DynamicEntryCountWarnThreshold int `json:"dynamic-entry-count-warn-threshold"`

		DynamicEntryOverflowPolicyList []DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList `json:"dynamic-entry-overflow-policy-list"`

		EnableClassListOverflow int `json:"enable-class-list-overflow"`

		EnableTopK int `json:"enable-top-k"`

		EnableTopKDestination int `json:"enable-top-k-destination"`

		FasterDeEscalation int `json:"faster-de-escalation"`

		GlidCfg DdosDstZonePortZoneServiceGlidCfg `json:"glid-cfg"`

		IpFilteringPolicy string `json:"ip-filtering-policy"`

		IpFilteringPolicyStatistics DdosDstZonePortZoneServiceIpFilteringPolicyStatistics254 `json:"ip-filtering-policy-statistics"`

		LevelList []DdosDstZonePortZoneServiceLevelList `json:"level-list"`

		LogSrcDefaultEnable int `json:"log-src-default-enable"`

		ManualModeEnable int `json:"manual-mode-enable"`

		ManualModeList []DdosDstZonePortZoneServiceManualModeList `json:"manual-mode-list"`

		MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`

		OutboundOnly int `json:"outbound-only"`

		PatternRecognition DdosDstZonePortZoneServicePatternRecognition255 `json:"pattern-recognition"`

		PatternRecognitionPuDetails DdosDstZonePortZoneServicePatternRecognitionPuDetails256 `json:"pattern-recognition-pu-details"`

		PortInd DdosDstZonePortZoneServicePortInd257 `json:"port-ind"`

		PortNum int `json:"port-num"`

		ProgressionTracking DdosDstZonePortZoneServiceProgressionTracking259 `json:"progression-tracking"`

		Protocol string `json:"protocol"`

		SameSourceDestPortDrop int `json:"same-source-dest-port-drop"`

		SetCounterBaseVal int `json:"set-counter-base-val"`

		SflowCommon int `json:"sflow-common"`

		SflowHttp int `json:"sflow-http"`

		SflowIpFilteringPolicy int `json:"sflow-ip-filtering-policy"`

		SflowPackets int `json:"sflow-packets"`

		SflowTcp DdosDstZonePortZoneServiceSflowTcp `json:"sflow-tcp"`

		SrcBasedPolicyList []DdosDstZonePortZoneServiceSrcBasedPolicyList `json:"src-based-policy-list"`

		Stateful int `json:"stateful"`

		TopkDestinations DdosDstZonePortZoneServiceTopkDestinations260 `json:"topk-destinations"`

		TopkDstNumRecords int `json:"topk-dst-num-records" dval:"20"`

		TopkDstSortKey string `json:"topk-dst-sort-key" dval:"avg"`

		TopkNumRecords int `json:"topk-num-records" dval:"20"`

		TopkSortKey string `json:"topk-sort-key" dval:"avg"`

		TopkSources DdosDstZonePortZoneServiceTopkSources261 `json:"topk-sources"`

		UnlimitedDynamicEntryCount int `json:"unlimited-dynamic-entry-count"`

		Uuid string `json:"uuid"`

		Virtualhosts DdosDstZonePortZoneServiceVirtualhosts262 `json:"virtualhosts"`

		ZoneName string
	} `json:"zone-service"`
}

type DdosDstZonePortZoneServiceCaptureConfig struct {
	CaptureConfigName string `json:"capture-config-name"`
	CaptureConfigMode string `json:"capture-config-mode"`
}

type DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyList struct {
	DummyName    string                                                               `json:"dummy-name"`
	Glid         string                                                               `json:"glid"`
	Action       string                                                               `json:"action"`
	LogEnable    int                                                                  `json:"log-enable"`
	LogPeriodic  int                                                                  `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                               `json:"uuid"`
	UserTag      string                                                               `json:"user-tag"`
}

type DdosDstZonePortZoneServiceDynamicEntryOverflowPolicyListZoneTemplate struct {
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

type DdosDstZonePortZoneServiceGlidCfg struct {
	Glid        string `json:"glid"`
	GlidAction  string `json:"glid-action"`
	ActionList  string `json:"action-list"`
	PerAddrGlid string `json:"per-addr-glid"`
}

type DdosDstZonePortZoneServiceIpFilteringPolicyStatistics254 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceLevelList struct {
	LevelNum                      string                                             `json:"level-num"`
	SrcDefaultGlid                string                                             `json:"src-default-glid"`
	GlidAction                    string                                             `json:"glid-action"`
	ZoneEscalationScore           int                                                `json:"zone-escalation-score"`
	ZoneViolationActions          string                                             `json:"zone-violation-actions"`
	SrcEscalationScore            int                                                `json:"src-escalation-score"`
	SrcViolationActions           string                                             `json:"src-violation-actions"`
	ZoneTemplate                  DdosDstZonePortZoneServiceLevelListZoneTemplate    `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                `json:"close-sessions-for-all-sources"`
	ClearSourcesUponDeescalation  int                                                `json:"clear-sources-upon-deescalation"`
	StartSignatureExtraction      int                                                `json:"start-signature-extraction"`
	StartPatternRecognition       int                                                `json:"start-pattern-recognition"`
	ApplyExtractedFilters         int                                                `json:"apply-extracted-filters"`
	Uuid                          string                                             `json:"uuid"`
	UserTag                       string                                             `json:"user-tag"`
	IndicatorList                 []DdosDstZonePortZoneServiceLevelListIndicatorList `json:"indicator-list"`
}

type DdosDstZonePortZoneServiceLevelListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServiceLevelListIndicatorList struct {
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

type DdosDstZonePortZoneServiceManualModeList struct {
	Config                        string                                               `json:"config"`
	SrcDefaultGlid                string                                               `json:"src-default-glid"`
	GlidAction                    string                                               `json:"glid-action"`
	ZoneTemplate                  DdosDstZonePortZoneServiceManualModeListZoneTemplate `json:"zone-template"`
	CloseSessionsForUnauthSources int                                                  `json:"close-sessions-for-unauth-sources"`
	CloseSessionsForAllSources    int                                                  `json:"close-sessions-for-all-sources"`
	Uuid                          string                                               `json:"uuid"`
	UserTag                       string                                               `json:"user-tag"`
}

type DdosDstZonePortZoneServiceManualModeListZoneTemplate struct {
	Quic  string `json:"quic"`
	Dns   string `json:"dns"`
	Http  string `json:"http"`
	SslL4 string `json:"ssl-l4"`
	Sip   string `json:"sip"`
	Tcp   string `json:"tcp"`
	Udp   string `json:"udp"`
	Encap string `json:"encap"`
}

type DdosDstZonePortZoneServicePatternRecognition255 struct {
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

type DdosDstZonePortZoneServicePatternRecognitionPuDetails256 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServicePortInd257 struct {
	Uuid           string                                               `json:"uuid"`
	SamplingEnable []DdosDstZonePortZoneServicePortIndSamplingEnable258 `json:"sampling-enable"`
}

type DdosDstZonePortZoneServicePortIndSamplingEnable258 struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceProgressionTracking259 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceSflowTcp struct {
	SflowTcpBasic    int `json:"sflow-tcp-basic"`
	SflowTcpStateful int `json:"sflow-tcp-stateful"`
	SflowSampSni     int `json:"sflow-samp-sni"`
}

type DdosDstZonePortZoneServiceSrcBasedPolicyList struct {
	SrcBasedPolicyName  string                                                            `json:"src-based-policy-name"`
	Uuid                string                                                            `json:"uuid"`
	UserTag             string                                                            `json:"user-tag"`
	PolicyClassListList []DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList `json:"policy-class-list-list"`
}

type DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListList struct {
	ClassListName                  string                                                                                       `json:"class-list-name"`
	ClassListGlid                  string                                                                                       `json:"class-list-glid"`
	Glid                           string                                                                                       `json:"glid"`
	GlidAction                     string                                                                                       `json:"glid-action"`
	Action                         string                                                                                       `json:"action"`
	LogEnable                      int                                                                                          `json:"log-enable"`
	LogPeriodic                    int                                                                                          `json:"log-periodic"`
	MaxDynamicEntryCount           int                                                                                          `json:"max-dynamic-entry-count"`
	DynamicEntryCountWarnThreshold int                                                                                          `json:"dynamic-entry-count-warn-threshold"`
	ZoneTemplate                   DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListZoneTemplate                  `json:"zone-template"`
	Uuid                           string                                                                                       `json:"uuid"`
	UserTag                        string                                                                                       `json:"user-tag"`
	SamplingEnable                 []DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable              `json:"sampling-enable"`
	ClassListOverflowPolicyList    []DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList `json:"class-list-overflow-policy-list"`
}

type DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListZoneTemplate struct {
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

type DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyList struct {
	DummyName    string                                                                                                 `json:"dummy-name"`
	Glid         string                                                                                                 `json:"glid"`
	Action       string                                                                                                 `json:"action"`
	LogEnable    int                                                                                                    `json:"log-enable"`
	LogPeriodic  int                                                                                                    `json:"log-periodic"`
	ZoneTemplate DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate `json:"zone-template"`
	Uuid         string                                                                                                 `json:"uuid"`
	UserTag      string                                                                                                 `json:"user-tag"`
}

type DdosDstZonePortZoneServiceSrcBasedPolicyListPolicyClassListListClassListOverflowPolicyListZoneTemplate struct {
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

type DdosDstZonePortZoneServiceTopkDestinations260 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceTopkSources261 struct {
	Uuid string `json:"uuid"`
}

type DdosDstZonePortZoneServiceVirtualhosts262 struct {
	VhostsConfig      string                                                     `json:"vhosts-config"`
	SourceTrackingAll int                                                        `json:"source-tracking-all"`
	Uuid              string                                                     `json:"uuid"`
	VirtualhostList   []DdosDstZonePortZoneServiceVirtualhostsVirtualhostList263 `json:"virtualhost-list"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostList263 struct {
	Vhost              string                                                               `json:"vhost"`
	Servername         []DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264 `json:"servername"`
	ServernameList     string                                                               `json:"servername-list"`
	ServernameMatchAny int                                                                  `json:"servername-match-any"`
	ServernameNoSni    int                                                                  `json:"servername-no-sni"`
	SourceTracking     string                                                               `json:"source-tracking" dval:"follow"`
	GlidCfg            DdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg265      `json:"glid-cfg"`
	Deny               int                                                                  `json:"deny"`
	Uuid               string                                                               `json:"uuid"`
	UserTag            string                                                               `json:"user-tag"`
	LevelList          []DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266  `json:"level-list"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListServername264 struct {
	MatchType       string `json:"match-type"`
	HostMatchString string `json:"host-match-string"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListGlidCfg265 struct {
	Glid       string `json:"glid"`
	GlidAction string `json:"glid-action"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelList266 struct {
	LevelNum       string                                                                        `json:"level-num"`
	SrcDefaultGlid string                                                                        `json:"src-default-glid"`
	GlidAction     string                                                                        `json:"glid-action"`
	ZoneTemplate   DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate267 `json:"zone-template"`
	Uuid           string                                                                        `json:"uuid"`
	UserTag        string                                                                        `json:"user-tag"`
}

type DdosDstZonePortZoneServiceVirtualhostsVirtualhostListLevelListZoneTemplate267 struct {
	SslL4 string `json:"ssl-l4"`
	Tcp   string `json:"tcp"`
}

func (p *DdosDstZonePortZoneService) GetId() string {
	return strconv.Itoa(p.Inst.PortNum) + "+" + p.Inst.Protocol
}

func (p *DdosDstZonePortZoneService) getPath() string {
	return "ddos/dst/zone/" + p.Inst.ZoneName + "/port/zone-service"
}

func (p *DdosDstZonePortZoneService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneService::Post")
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

func (p *DdosDstZonePortZoneService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneService::Get")
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
func (p *DdosDstZonePortZoneService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneService::Put")
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

func (p *DdosDstZonePortZoneService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDstZonePortZoneService::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
