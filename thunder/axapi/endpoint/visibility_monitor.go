

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitor struct {
	Inst struct {

    AgentList []VisibilityMonitorAgentList `json:"agent-list"`
    DebugList []VisibilityMonitorDebugList `json:"debug-list"`
    DeleteDebugFile VisibilityMonitorDeleteDebugFile1916 `json:"delete-debug-file"`
    IndexSessions int `json:"index-sessions"`
    IndexSessionsType string `json:"index-sessions-type"`
    MonEntityTopk int `json:"mon-entity-topk"`
    MonitorKey string `json:"monitor-key"`
    Netflow VisibilityMonitorNetflow1917 `json:"netflow"`
    PrimaryMonitor string `json:"primary-monitor"`
    ReplayDebugFile VisibilityMonitorReplayDebugFile1918 `json:"replay-debug-file"`
    SecondaryMonitor VisibilityMonitorSecondaryMonitor1919 `json:"secondary-monitor"`
    Sflow VisibilityMonitorSflow1923 `json:"sflow"`
    SourceEntityTopk int `json:"source-entity-topk"`
    Template VisibilityMonitorTemplate `json:"template"`
    Uuid string `json:"uuid"`

	} `json:"monitor"`
}


type VisibilityMonitorAgentList struct {
    AgentName string `json:"agent-name"`
    AgentV4Addr string `json:"agent-v4-addr"`
    AgentV6Addr string `json:"agent-v6-addr"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []VisibilityMonitorAgentListSamplingEnable `json:"sampling-enable"`
}


type VisibilityMonitorAgentListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VisibilityMonitorDebugList struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorDeleteDebugFile1916 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorNetflow1917 struct {
    ListeningPort int `json:"listening-port" dval:"9996"`
    TemplateActiveTimeout int `json:"template-active-timeout" dval:"30"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorReplayDebugFile1918 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorSecondaryMonitor1919 struct {
    SecondaryMonitoringKey string `json:"secondary-monitoring-key"`
    MonEntityTopk int `json:"mon-entity-topk"`
    SourceEntityTopk int `json:"source-entity-topk"`
    Uuid string `json:"uuid"`
    DebugList []VisibilityMonitorSecondaryMonitorDebugList1920 `json:"debug-list"`
    DeleteDebugFile VisibilityMonitorSecondaryMonitorDeleteDebugFile1921 `json:"delete-debug-file"`
    ReplayDebugFile VisibilityMonitorSecondaryMonitorReplayDebugFile1922 `json:"replay-debug-file"`
}


type VisibilityMonitorSecondaryMonitorDebugList1920 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorSecondaryMonitorDeleteDebugFile1921 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorSecondaryMonitorReplayDebugFile1922 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorSflow1923 struct {
    ListeningPort int `json:"listening-port" dval:"6343"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorTemplate struct {
    Notification []VisibilityMonitorTemplateNotification `json:"notification"`
}


type VisibilityMonitorTemplateNotification struct {
    NotifTemplateName string `json:"notif-template-name"`
}

func (p *VisibilityMonitor) GetId() string{
    return "1"
}

func (p *VisibilityMonitor) getPath() string{
    return "visibility/monitor"
}

func (p *VisibilityMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitor::Post")
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

func (p *VisibilityMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitor::Get")
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
func (p *VisibilityMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitor::Put")
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

func (p *VisibilityMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
