

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSharedPollMode struct {
	Inst struct {

    Disable int `json:"disable"`
    Enable int `json:"enable"`

	} `json:"shared-poll-mode"`
}

func (p *SystemSharedPollMode) GetId() string{
    return "1"
}

func (p *SystemSharedPollMode) getPath() string{
    return "system/shared-poll-mode"
}

func (p *SystemSharedPollMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSharedPollMode::Post")
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

func (p *SystemSharedPollMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSharedPollMode::Get")
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
func (p *SystemSharedPollMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSharedPollMode::Put")
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

func (p *SystemSharedPollMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSharedPollMode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
