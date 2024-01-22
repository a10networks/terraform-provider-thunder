

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitorSecondaryMonitorDeleteDebugFile struct {
	Inst struct {

    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`

	} `json:"delete-debug-file"`
}

func (p *VisibilityMonitorSecondaryMonitorDeleteDebugFile) GetId() string{
    return p.Inst.DebugIpAddr
}

func (p *VisibilityMonitorSecondaryMonitorDeleteDebugFile) getPath() string{
    return "visibility/monitor/secondary-monitor/delete-debug-file"
}

func (p *VisibilityMonitorSecondaryMonitorDeleteDebugFile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDeleteDebugFile::Post")
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

func (p *VisibilityMonitorSecondaryMonitorDeleteDebugFile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDeleteDebugFile::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *VisibilityMonitorSecondaryMonitorDeleteDebugFile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDeleteDebugFile::Put")
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

func (p *VisibilityMonitorSecondaryMonitorDeleteDebugFile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDeleteDebugFile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
