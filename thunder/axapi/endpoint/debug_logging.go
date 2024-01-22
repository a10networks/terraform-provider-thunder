

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugLogging struct {
	Inst struct {

    All int `json:"all"`
    Command int `json:"command"`
    Error int `json:"error"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}

func (p *DebugLogging) GetId() string{
    return "1"
}

func (p *DebugLogging) getPath() string{
    return "debug/logging"
}

func (p *DebugLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLogging::Post")
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

func (p *DebugLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLogging::Get")
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
func (p *DebugLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLogging::Put")
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

func (p *DebugLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
