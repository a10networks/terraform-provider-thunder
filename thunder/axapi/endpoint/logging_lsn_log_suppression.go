

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLsnLogSuppression struct {
	Inst struct {

    Count1 int `json:"count1" dval:"100"`
    Time int `json:"time" dval:"30"`
    Uuid string `json:"uuid"`

	} `json:"log-suppression"`
}

func (p *LoggingLsnLogSuppression) GetId() string{
    return "1"
}

func (p *LoggingLsnLogSuppression) getPath() string{
    return "logging/lsn/log-suppression"
}

func (p *LoggingLsnLogSuppression) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnLogSuppression::Post")
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

func (p *LoggingLsnLogSuppression) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnLogSuppression::Get")
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
func (p *LoggingLsnLogSuppression) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnLogSuppression::Put")
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

func (p *LoggingLsnLogSuppression) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLsnLogSuppression::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
