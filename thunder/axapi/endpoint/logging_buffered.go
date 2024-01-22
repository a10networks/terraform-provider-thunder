

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingBuffered struct {
	Inst struct {

    Buffersize int `json:"buffersize" dval:"30000"`
    Levelname string `json:"levelname" dval:"debugging"`
    PartitionBuffersize int `json:"partition-buffersize" dval:"3000"`
    Uuid string `json:"uuid"`

	} `json:"buffered"`
}

func (p *LoggingBuffered) GetId() string{
    return "1"
}

func (p *LoggingBuffered) getPath() string{
    return "logging/buffered"
}

func (p *LoggingBuffered) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingBuffered::Post")
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

func (p *LoggingBuffered) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingBuffered::Get")
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
func (p *LoggingBuffered) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingBuffered::Put")
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

func (p *LoggingBuffered) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingBuffered::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
