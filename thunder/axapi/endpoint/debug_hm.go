

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugHm struct {
	Inst struct {

    Level int `json:"level"`
    MethodType string `json:"method-type"`
    PinUid int `json:"pin-uid"`
    Uuid string `json:"uuid"`

	} `json:"hm"`
}

func (p *DebugHm) GetId() string{
    return "1"
}

func (p *DebugHm) getPath() string{
    return "debug/hm"
}

func (p *DebugHm) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugHm::Post")
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

func (p *DebugHm) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugHm::Get")
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
func (p *DebugHm) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugHm::Put")
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

func (p *DebugHm) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugHm::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
