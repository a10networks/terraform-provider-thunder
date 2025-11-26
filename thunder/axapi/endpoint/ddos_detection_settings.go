package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDetectionSettings struct {
	Inst struct {
		CtrlCpuUsage int `json:"ctrl-cpu-usage"`

		DeEscalationQuietTime int `json:"de-escalation-quiet-time"`

		DedicatedCpus int `json:"dedicated-cpus"`

		DetectionWindowSize int `json:"detection-window-size" dval:"1"`

		DetectorMode string `json:"detector-mode"`

		EntrySaving DdosDetectionSettingsEntrySaving149 `json:"entry-saving"`

		ExportInterval int `json:"export-interval" dval:"20"`

		FullCoreEnable int `json:"full-core-enable"`

		HistogramDeEscalatePercentage int `json:"histogram-de-escalate-percentage"`

		HistogramEscalatePercentage int `json:"histogram-escalate-percentage"`

		InitialLearningInterval int `json:"initial-learning-interval"`

		NetworkObjectFloodingMultiple int `json:"network-object-flooding-multiple" dval:"2"`

		NetworkObjectSubnetNotifyPercent int `json:"network-object-subnet-notify-percent"`

		NetworkObjectWindowSize string `json:"network-object-window-size" dval:"30"`

		NotificationDebugLog string `json:"notification-debug-log"`

		PktSampling []DdosDetectionSettingsPktSampling `json:"pkt-sampling"`

		StandaloneSettings DdosDetectionSettingsStandaloneSettings150 `json:"standalone-settings"`

		TopKResetInterval int `json:"top-k-reset-interval"`

		Uuid string `json:"uuid"`

		ZoneNotifications DdosDetectionSettingsZoneNotifications153 `json:"zone-notifications"`
	} `json:"settings"`
}

type DdosDetectionSettingsEntrySaving149 struct {
	DisableBootupRestore int    `json:"disable-bootup-restore"`
	Interval             int    `json:"interval"`
	Uuid                 string `json:"uuid"`
}

type DdosDetectionSettingsPktSampling struct {
	OverrideRate int `json:"override-rate"`
	StartLevel   int `json:"start-level" dval:"1"`
}

type DdosDetectionSettingsStandaloneSettings150 struct {
	Action  string                                            `json:"action" dval:"disable"`
	Uuid    string                                            `json:"uuid"`
	Sflow   DdosDetectionSettingsStandaloneSettingsSflow151   `json:"sflow"`
	Netflow DdosDetectionSettingsStandaloneSettingsNetflow152 `json:"netflow"`
}

type DdosDetectionSettingsStandaloneSettingsSflow151 struct {
	ListeningPort int    `json:"listening-port" dval:"6343"`
	Uuid          string `json:"uuid"`
}

type DdosDetectionSettingsStandaloneSettingsNetflow152 struct {
	ListeningPort         int    `json:"listening-port" dval:"9996"`
	TemplateActiveTimeout int    `json:"template-active-timeout" dval:"30"`
	DistributeByDuration  string `json:"distribute-by-duration" dval:"enable"`
	Uuid                  string `json:"uuid"`
}

type DdosDetectionSettingsZoneNotifications153 struct {
	SourceEntry string `json:"source-entry" dval:"disable"`
	Uuid        string `json:"uuid"`
}

func (p *DdosDetectionSettings) GetId() string {
	return "1"
}

func (p *DdosDetectionSettings) getPath() string {
	return "ddos/detection/settings"
}

func (p *DdosDetectionSettings) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettings::Post")
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

func (p *DdosDetectionSettings) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettings::Get")
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
func (p *DdosDetectionSettings) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettings::Put")
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

func (p *DdosDetectionSettings) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDetectionSettings::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
