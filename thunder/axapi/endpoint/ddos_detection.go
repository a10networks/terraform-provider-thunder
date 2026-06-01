package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type DdosDetection struct {
	Inst struct {
		AgentGroupList []DdosDetectionAgentGroupList `json:"agent-group-list"`

		AgentList []DdosDetectionAgentList `json:"agent-list"`

		DdosScript DdosDetectionDdosScript157 `json:"ddos-script"`

		Disable int `json:"disable"`

		EntrySaving DdosDetectionEntrySaving158 `json:"entry-saving"`

		ResourceUsage DdosDetectionResourceUsage159 `json:"resource-usage"`

		Settings DdosDetectionSettings160 `json:"settings"`

		Statistics DdosDetectionStatistics167 `json:"statistics"`

		Trustlist DdosDetectionTrustlist168 `json:"trustlist"`

		Uuid string `json:"uuid"`

		XflowInterfaceSelectionList []DdosDetectionXflowInterfaceSelectionList `json:"xflow-interface-selection-list"`
	} `json:"detection"`
}

type DdosDetectionAgentGroupList struct {
	AgentGroupName string                             `json:"agent-group-name"`
	Agent          []DdosDetectionAgentGroupListAgent `json:"agent"`
	Uuid           string                             `json:"uuid"`
	UserTag        string                             `json:"user-tag"`
}

type DdosDetectionAgentGroupListAgent struct {
	AgentName string `json:"agent-name"`
}

type DdosDetectionAgentList struct {
	AgentName      string                                 `json:"agent-name"`
	AgentV4Addr    string                                 `json:"agent-v4-addr"`
	AgentV6Addr    string                                 `json:"agent-v6-addr"`
	AgentType      string                                 `json:"agent-type"`
	Uuid           string                                 `json:"uuid"`
	UserTag        string                                 `json:"user-tag"`
	SamplingEnable []DdosDetectionAgentListSamplingEnable `json:"sampling-enable"`
	Sflow          DdosDetectionAgentListSflow            `json:"sflow"`
	Netflow        DdosDetectionAgentListNetflow          `json:"netflow"`
	Snmp           DdosDetectionAgentListSnmp             `json:"snmp"`
}

type DdosDetectionAgentListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

type DdosDetectionAgentListSflow struct {
	SflowPktSamplesCollection string `json:"sflow-pkt-samples-collection" dval:"enable"`
	Uuid                      string `json:"uuid"`
}

type DdosDetectionAgentListNetflow struct {
	NetflowSamplesCollection string `json:"netflow-samples-collection" dval:"enable"`
	NetflowSamplingRate      int    `json:"netflow-sampling-rate" dval:"1"`
	ActiveTimeout            int    `json:"active-timeout"`
	InactiveTimeout          int    `json:"inactive-timeout"`
	Uuid                     string `json:"uuid"`
}

type DdosDetectionAgentListSnmp struct {
	Ipv4Addr        string `json:"ipv4-addr"`
	CommunityString string `json:"community-string"`
	Refresh         int    `json:"refresh"`
	Uuid            string `json:"uuid"`
}

type DdosDetectionDdosScript157 struct {
	File   string `json:"file"`
	Action string `json:"action"`
	Uuid   string `json:"uuid"`
}

type DdosDetectionEntrySaving158 struct {
	ClearSavedData int    `json:"clear-saved-data"`
	ManualSave     int    `json:"manual-save"`
	ManualRestore  int    `json:"manual-restore"`
	Uuid           string `json:"uuid"`
}

type DdosDetectionResourceUsage159 struct {
	Uuid string `json:"uuid"`
}

type DdosDetectionSettings160 struct {
	DetectorMode                     string                                     `json:"detector-mode"`
	DedicatedCpus                    int                                        `json:"dedicated-cpus"`
	CtrlCpuUsage                     int                                        `json:"ctrl-cpu-usage"`
	FullCoreEnable                   int                                        `json:"full-core-enable"`
	TopKResetInterval                int                                        `json:"top-k-reset-interval"`
	PktSampling                      []DdosDetectionSettingsPktSampling161      `json:"pkt-sampling"`
	HistogramEscalatePercentage      int                                        `json:"histogram-escalate-percentage"`
	HistogramDeEscalatePercentage    int                                        `json:"histogram-de-escalate-percentage"`
	DetectionWindowSize              int                                        `json:"detection-window-size" dval:"1"`
	InitialLearningInterval          int                                        `json:"initial-learning-interval"`
	ExportInterval                   int                                        `json:"export-interval" dval:"20"`
	NotificationDebugLog             string                                     `json:"notification-debug-log"`
	NetworkObjectWindowSize          string                                     `json:"network-object-window-size" dval:"30"`
	NetworkObjectFloodingMultiple    int                                        `json:"network-object-flooding-multiple" dval:"2"`
	DeEscalationQuietTime            int                                        `json:"de-escalation-quiet-time"`
	NetworkObjectSubnetNotifyPercent int                                        `json:"network-object-subnet-notify-percent"`
	Uuid                             string                                     `json:"uuid"`
	EntrySaving                      DdosDetectionSettingsEntrySaving162        `json:"entry-saving"`
	StandaloneSettings               DdosDetectionSettingsStandaloneSettings163 `json:"standalone-settings"`
	ZoneNotifications                DdosDetectionSettingsZoneNotifications166  `json:"zone-notifications"`
}

type DdosDetectionSettingsPktSampling161 struct {
	OverrideRate int `json:"override-rate"`
	StartLevel   int `json:"start-level" dval:"1"`
}

type DdosDetectionSettingsEntrySaving162 struct {
	DisableBootupRestore int    `json:"disable-bootup-restore"`
	Interval             int    `json:"interval"`
	Uuid                 string `json:"uuid"`
}

type DdosDetectionSettingsStandaloneSettings163 struct {
	Action  string                                            `json:"action" dval:"disable"`
	Uuid    string                                            `json:"uuid"`
	Sflow   DdosDetectionSettingsStandaloneSettingsSflow164   `json:"sflow"`
	Netflow DdosDetectionSettingsStandaloneSettingsNetflow165 `json:"netflow"`
}

type DdosDetectionSettingsStandaloneSettingsSflow164 struct {
	ListeningPort int    `json:"listening-port" dval:"6343"`
	Uuid          string `json:"uuid"`
}

type DdosDetectionSettingsStandaloneSettingsNetflow165 struct {
	ListeningPort         int    `json:"listening-port" dval:"9996"`
	TemplateActiveTimeout int    `json:"template-active-timeout" dval:"30"`
	DistributeByDuration  string `json:"distribute-by-duration" dval:"enable"`
	Uuid                  string `json:"uuid"`
}

type DdosDetectionSettingsZoneNotifications166 struct {
	SourceEntry string `json:"source-entry" dval:"disable"`
	Uuid        string `json:"uuid"`
}

type DdosDetectionStatistics167 struct {
	Uuid string `json:"uuid"`
}

type DdosDetectionTrustlist168 struct {
	V4ClassList string `json:"v4-class-list"`
	V6ClassList string `json:"v6-class-list"`
	Uuid        string `json:"uuid"`
}

type DdosDetectionXflowInterfaceSelectionList struct {
	Type    string                                        `json:"type"`
	Uuid    string                                        `json:"uuid"`
	UserTag string                                        `json:"user-tag"`
	Regex   DdosDetectionXflowInterfaceSelectionListRegex `json:"regex"`
}

type DdosDetectionXflowInterfaceSelectionListRegex struct {
	RuleList []DdosDetectionXflowInterfaceSelectionListRegexRuleList `json:"rule-list"`
	Uuid     string                                                  `json:"uuid"`
}

type DdosDetectionXflowInterfaceSelectionListRegexRuleList struct {
	SingleRegex string `json:"single-regex"`
}

func (p *DdosDetection) GetId() string {
	return "1"
}

func (p *DdosDetection) getPath() string {
	return "ddos/detection"
}

func (p *DdosDetection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetection::Post")
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

func (p *DdosDetection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetection::Get")
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
func (p *DdosDetection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetection::Put")
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

func (p *DdosDetection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetection::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
