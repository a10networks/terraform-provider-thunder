

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetection struct {
	Inst struct {

    AgentList []DdosDetectionAgentList `json:"agent-list"`
    DdosScript DdosDetectionDdosScript136 `json:"ddos-script"`
    Disable int `json:"disable"`
    ResourceUsage DdosDetectionResourceUsage137 `json:"resource-usage"`
    Settings DdosDetectionSettings138 `json:"settings"`
    Statistics DdosDetectionStatistics144 `json:"statistics"`
    Uuid string `json:"uuid"`

	} `json:"detection"`
}


type DdosDetectionAgentList struct {
    AgentName string `json:"agent-name"`
    AgentV4Addr string `json:"agent-v4-addr"`
    AgentV6Addr string `json:"agent-v6-addr"`
    AgentType string `json:"agent-type"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []DdosDetectionAgentListSamplingEnable `json:"sampling-enable"`
    Sflow DdosDetectionAgentListSflow `json:"sflow"`
    Netflow DdosDetectionAgentListNetflow `json:"netflow"`
}


type DdosDetectionAgentListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type DdosDetectionAgentListSflow struct {
    SflowPktSamplesCollection string `json:"sflow-pkt-samples-collection" dval:"enable"`
    Uuid string `json:"uuid"`
}


type DdosDetectionAgentListNetflow struct {
    NetflowSamplesCollection string `json:"netflow-samples-collection" dval:"enable"`
    NetflowSamplingRate int `json:"netflow-sampling-rate" dval:"1"`
    ActiveTimeout int `json:"active-timeout"`
    InactiveTimeout int `json:"inactive-timeout"`
    Uuid string `json:"uuid"`
}


type DdosDetectionDdosScript136 struct {
    File string `json:"file"`
    Action string `json:"action"`
    Uuid string `json:"uuid"`
}


type DdosDetectionResourceUsage137 struct {
    Uuid string `json:"uuid"`
}


type DdosDetectionSettings138 struct {
    DetectorMode string `json:"detector-mode"`
    DedicatedCpus int `json:"dedicated-cpus"`
    CtrlCpuUsage int `json:"ctrl-cpu-usage"`
    FullCoreEnable int `json:"full-core-enable"`
    TopKResetInterval int `json:"top-k-reset-interval"`
    PktSampling DdosDetectionSettingsPktSampling139 `json:"pkt-sampling"`
    HistogramEscalatePercentage int `json:"histogram-escalate-percentage"`
    HistogramDeEscalatePercentage int `json:"histogram-de-escalate-percentage"`
    DetectionWindowSize int `json:"detection-window-size" dval:"1"`
    InitialLearningInterval int `json:"initial-learning-interval"`
    ExportInterval int `json:"export-interval" dval:"20"`
    NotificationDebugLog string `json:"notification-debug-log"`
    NetworkObjectWindowSize string `json:"network-object-window-size" dval:"30"`
    NetworkObjectFloodingMultiple int `json:"network-object-flooding-multiple" dval:"2"`
    DeEscalationQuietTime int `json:"de-escalation-quiet-time"`
    Uuid string `json:"uuid"`
    EntrySaving DdosDetectionSettingsEntrySaving140 `json:"entry-saving"`
    StandaloneSettings DdosDetectionSettingsStandaloneSettings141 `json:"standalone-settings"`
}


type DdosDetectionSettingsPktSampling139 struct {
    OverrideRate int `json:"override-rate"`
    AssignIndex int `json:"assign-index"`
    AssignRate int `json:"assign-rate"`
}


type DdosDetectionSettingsEntrySaving140 struct {
    Interval int `json:"interval"`
    ManualSave int `json:"manual-save"`
    ManualRestore int `json:"manual-restore"`
    Uuid string `json:"uuid"`
}


type DdosDetectionSettingsStandaloneSettings141 struct {
    Action string `json:"action" dval:"disable"`
    DeEscalationQuietTime int `json:"de-escalation-quiet-time"`
    Uuid string `json:"uuid"`
    Sflow DdosDetectionSettingsStandaloneSettingsSflow142 `json:"sflow"`
    Netflow DdosDetectionSettingsStandaloneSettingsNetflow143 `json:"netflow"`
}


type DdosDetectionSettingsStandaloneSettingsSflow142 struct {
    ListeningPort int `json:"listening-port" dval:"6343"`
    Uuid string `json:"uuid"`
}


type DdosDetectionSettingsStandaloneSettingsNetflow143 struct {
    ListeningPort int `json:"listening-port" dval:"9996"`
    TemplateActiveTimeout int `json:"template-active-timeout" dval:"30"`
    Uuid string `json:"uuid"`
}


type DdosDetectionStatistics144 struct {
    Uuid string `json:"uuid"`
}

func (p *DdosDetection) GetId() string{
    return "1"
}

func (p *DdosDetection) getPath() string{
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
        if len(axResp) > 0{
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
