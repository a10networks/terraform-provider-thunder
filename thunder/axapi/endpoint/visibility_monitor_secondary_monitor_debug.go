

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VisibilityMonitorSecondaryMonitorDebug struct {
	Inst struct {

    DebugIpAddr string `json:"debug-ip-addr"`
    DebugPort int `json:"debug-port"`
    DebugProtocol string `json:"debug-protocol"`
    Uuid string `json:"uuid"`

	} `json:"debug"`
}

func (p *VisibilityMonitorSecondaryMonitorDebug) GetId() string{
    return p.Inst.DebugIpAddr+"+"+strconv.Itoa(p.Inst.DebugPort)+"+"+p.Inst.DebugProtocol
}

func (p *VisibilityMonitorSecondaryMonitorDebug) getPath() string{
    return "visibility/monitor/secondary-monitor/debug"
}

func (p *VisibilityMonitorSecondaryMonitorDebug) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDebug::Post")
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

func (p *VisibilityMonitorSecondaryMonitorDebug) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDebug::Get")
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
func (p *VisibilityMonitorSecondaryMonitorDebug) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDebug::Put")
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

func (p *VisibilityMonitorSecondaryMonitorDebug) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityMonitorSecondaryMonitorDebug::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
