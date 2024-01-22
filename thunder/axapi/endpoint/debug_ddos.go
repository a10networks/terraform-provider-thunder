

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugDdos struct {
	Inst struct {

    ControlVar int `json:"control-var"`
    DnsCache int `json:"dns-cache"`
    Event int `json:"event"`
    EventFilter string `json:"event-filter"`
    FlowBasedDetection int `json:"flow-based-detection"`
    Level int `json:"level"`
    Uuid string `json:"uuid"`
    Zbar int `json:"zbar"`

	} `json:"ddos"`
}

func (p *DebugDdos) GetId() string{
    return "1"
}

func (p *DebugDdos) getPath() string{
    return "debug/ddos"
}

func (p *DebugDdos) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugDdos::Post")
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

func (p *DebugDdos) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugDdos::Get")
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
func (p *DebugDdos) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugDdos::Put")
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

func (p *DebugDdos) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugDdos::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
