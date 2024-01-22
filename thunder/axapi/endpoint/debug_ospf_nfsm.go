

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugOspfNfsm struct {
	Inst struct {

    Events int `json:"events"`
    Status int `json:"status"`
    Timers int `json:"timers"`
    Uuid string `json:"uuid"`

	} `json:"nfsm"`
}

func (p *DebugOspfNfsm) GetId() string{
    return "1"
}

func (p *DebugOspfNfsm) getPath() string{
    return "debug/ospf/nfsm"
}

func (p *DebugOspfNfsm) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfNfsm::Post")
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

func (p *DebugOspfNfsm) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfNfsm::Get")
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
func (p *DebugOspfNfsm) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfNfsm::Put")
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

func (p *DebugOspfNfsm) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfNfsm::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
