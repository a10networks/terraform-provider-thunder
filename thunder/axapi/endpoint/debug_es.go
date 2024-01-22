

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugEs struct {
	Inst struct {

    Level int `json:"level"`
    Uuid string `json:"uuid"`

	} `json:"es"`
}

func (p *DebugEs) GetId() string{
    return "1"
}

func (p *DebugEs) getPath() string{
    return "debug/es"
}

func (p *DebugEs) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugEs::Post")
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

func (p *DebugEs) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugEs::Get")
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
func (p *DebugEs) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugEs::Put")
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

func (p *DebugEs) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugEs::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
