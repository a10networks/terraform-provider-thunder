

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingEmailAddress struct {
	Inst struct {

    EmailList []LoggingEmailAddressEmailList `json:"email-list"`
    Uuid string `json:"uuid"`

	} `json:"email-address"`
}


type LoggingEmailAddressEmailList struct {
    EmailAddress string `json:"email-address"`
}

func (p *LoggingEmailAddress) GetId() string{
    return "1"
}

func (p *LoggingEmailAddress) getPath() string{
    return "logging/email-address"
}

func (p *LoggingEmailAddress) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailAddress::Post")
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

func (p *LoggingEmailAddress) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailAddress::Get")
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
func (p *LoggingEmailAddress) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailAddress::Put")
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

func (p *LoggingEmailAddress) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailAddress::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
