package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDetection struct {
	Inst struct {
		AgentGroupList []DdosDetectionAgentGroupList `json:"agent-group-list"`

		AgentList []DdosDetectionAgentList `json:"agent-list"`

		DdosScript DdosDetectionDdosScript154 `json:"ddos-script"`

		Disable int `json:"disable"`

		EntrySaving DdosDetectionEntrySaving155 `json:"entry-saving"`

		ResourceUsage DdosDetectionResourceUsage156 `json:"resource-usage"`

		Settings DdosDetectionSettings157 `json:"settings"`

		Statistics DdosDetectionStatistics164 `json:"statistics"`

		Trustlist DdosDetectionTrustlist165 `json:"trustlist"`

		Uuid string `json:"uuid"`
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

type DdosDetectionDdosScript154 struct {
	File   string `json:"file"`
	Action string `json:"action"`
	Uuid   string `json:"uuid"`
}

type DdosDetectionEntrySaving155 struct {
	ClearSavedData int    `json:"clear-saved-data"`
	ManualSave     int    `json:"manual-save"`
	ManualRestore  int    `json:"manual-restore"`
	Uuid           string `json:"uuid"`
}

type DdosDetectionResourceUsage156 struct {
	Uuid string `json:"uuid"`
}

type DdosDetectionSettings157 struct {
	DetectorMode                     string                                     `json:"detector-mode"`
	DedicatedCpus                    int                                        `json:"dedicated-cpus"`
	CtrlCpuUsage                     int                                        `json:"ctrl-cpu-usage"`
	FullCoreEnable                   int                                        `json:"full-core-enable"`
	TopKResetInterval                int                                        `json:"top-k-reset-interval"`
	PktSampling                      []DdosDetectionSettingsPktSampling158      `json:"pkt-sampling"`
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
	EntrySaving                      DdosDetectionSettingsEntrySaving159        `json:"entry-saving"`
	StandaloneSettings               DdosDetectionSettingsStandaloneSettings160 `json:"standalone-settings"`
	ZoneNotifications                DdosDetectionSettingsZoneNotifications163  `json:"zone-notifications"`
}

type DdosDetectionSettingsPktSampling158 struct {
	OverrideRate int `json:"override-rate"`
	StartLevel   int `json:"start-level" dval:"1"`
}

type DdosDetectionSettingsEntrySaving159 struct {
	DisableBootupRestore int    `json:"disable-bootup-restore"`
	Interval             int    `json:"interval"`
	Uuid                 string `json:"uuid"`
}

type DdosDetectionSettingsStandaloneSettings160 struct {
	Action  string                                            `json:"action" dval:"disable"`
	Uuid    string                                            `json:"uuid"`
	Sflow   DdosDetectionSettingsStandaloneSettingsSflow161   `json:"sflow"`
	Netflow DdosDetectionSettingsStandaloneSettingsNetflow162 `json:"netflow"`
}

type DdosDetectionSettingsStandaloneSettingsSflow161 struct {
	ListeningPort int    `json:"listening-port" dval:"6343"`
	Uuid          string `json:"uuid"`
}

type DdosDetectionSettingsStandaloneSettingsNetflow162 struct {
	ListeningPort         int    `json:"listening-port" dval:"9996"`
	TemplateActiveTimeout int    `json:"template-active-timeout" dval:"30"`
	DistributeByDuration  string `json:"distribute-by-duration" dval:"enable"`
	Uuid                  string `json:"uuid"`
}

type DdosDetectionSettingsZoneNotifications163 struct {
	SourceEntry string `json:"source-entry" dval:"disable"`
	Uuid        string `json:"uuid"`
}

type DdosDetectionStatistics164 struct {
	Uuid string `json:"uuid"`
}

type DdosDetectionTrustlist165 struct {
	V4ClassList string `json:"v4-class-list"`
	V6ClassList string `json:"v6-class-list"`
	Uuid        string `json:"uuid"`
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
