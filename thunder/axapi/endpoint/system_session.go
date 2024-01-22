

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSession struct {
	Inst struct {

    SamplingEnable []SystemSessionSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"session"`
}


type SystemSessionSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemSession) GetId() string{
    return "1"
}

func (p *SystemSession) getPath() string{
    return "system/session"
}

func (p *SystemSession) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSession::Post")
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

func (p *SystemSession) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSession::Get")
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
func (p *SystemSession) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSession::Put")
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

func (p *SystemSession) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSession::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
