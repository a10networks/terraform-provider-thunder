

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingEmailBuffer struct {
	Inst struct {

    Number int `json:"number" dval:"50"`
    Time int `json:"time" dval:"10"`
    Uuid string `json:"uuid"`

	} `json:"buffer"`
}

func (p *LoggingEmailBuffer) GetId() string{
    return "1"
}

func (p *LoggingEmailBuffer) getPath() string{
    return "logging/email/buffer"
}

func (p *LoggingEmailBuffer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailBuffer::Post")
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

func (p *LoggingEmailBuffer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailBuffer::Get")
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
func (p *LoggingEmailBuffer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailBuffer::Put")
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

func (p *LoggingEmailBuffer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailBuffer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
