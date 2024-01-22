

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitorSecondaryMonitor struct {
	Inst struct {

    DebugList []VisibilityMonitorSecondaryMonitorDebugList `json:"debug-list"`
    DeleteDebugFile VisibilityMonitorSecondaryMonitorDeleteDebugFile1914 `json:"delete-debug-file"`
    MonEntityTopk int `json:"mon-entity-topk"`
    ReplayDebugFile VisibilityMonitorSecondaryMonitorReplayDebugFile1915 `json:"replay-debug-file"`
    SecondaryMonitoringKey string `json:"secondary-monitoring-key"`
    SourceEntityTopk int `json:"source-entity-topk"`
    Uuid string `json:"uuid"`

	} `json:"secondary-monitor"`
}


type VisibilityMonitorSecondaryMonitorDebugList struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
    Uuid string `json:"uuid"`
}


type VisibilityMonitorSecondaryMonitorDeleteDebugFile1914 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}


type VisibilityMonitorSecondaryMonitorReplayDebugFile1915 struct {
    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
}

func (p *VisibilityMonitorSecondaryMonitor) GetId() string{
    return "1"
}

func (p *VisibilityMonitorSecondaryMonitor) getPath() string{
    return "visibility/monitor/secondary-monitor"
}

func (p *VisibilityMonitorSecondaryMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitor::Post")
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

func (p *VisibilityMonitorSecondaryMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitor::Get")
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
func (p *VisibilityMonitorSecondaryMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitor::Put")
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

func (p *VisibilityMonitorSecondaryMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
