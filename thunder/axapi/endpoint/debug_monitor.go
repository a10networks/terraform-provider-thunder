

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugMonitor struct {
	Inst struct {

    AllSlots int `json:"all-slots"`
    Cpuid int `json:"cpuid"`
    Filename string `json:"filename"`
    Filesize int `json:"filesize" dval:"1"`
    NoStop int `json:"no-stop"`
    Timeout int `json:"timeout" dval:"5"`
    Uuid string `json:"uuid"`

	} `json:"monitor"`
}

func (p *DebugMonitor) GetId() string{
    return "1"
}

func (p *DebugMonitor) getPath() string{
    return "debug/monitor"
}

func (p *DebugMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugMonitor::Post")
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

func (p *DebugMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugMonitor::Get")
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
func (p *DebugMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugMonitor::Put")
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

func (p *DebugMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
