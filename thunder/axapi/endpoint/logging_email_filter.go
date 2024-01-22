

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type LoggingEmailFilter struct {
	Inst struct {

    Expression string `json:"expression"`
    FilterId int `json:"filter-id"`
    Trigger int `json:"trigger"`
    Uuid string `json:"uuid"`

	} `json:"filter"`
}

func (p *LoggingEmailFilter) GetId() string{
    return strconv.Itoa(p.Inst.FilterId)
}

func (p *LoggingEmailFilter) getPath() string{
    return "logging/email/filter"
}

func (p *LoggingEmailFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailFilter::Post")
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

func (p *LoggingEmailFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailFilter::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *LoggingEmailFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailFilter::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *LoggingEmailFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingEmailFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
