

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingEmailLevel struct {
	Inst struct {

    EmailLevelname string `json:"email-levelname" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"level"`
}

func (p *LoggingEmailLevel) GetId() string{
    return "1"
}

func (p *LoggingEmailLevel) getPath() string{
    return "logging/email/level"
}

func (p *LoggingEmailLevel) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailLevel::Post")
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

func (p *LoggingEmailLevel) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailLevel::Get")
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
func (p *LoggingEmailLevel) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailLevel::Put")
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

func (p *LoggingEmailLevel) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailLevel::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
