

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingSyslog struct {
	Inst struct {

    SyslogLevelname string `json:"syslog-levelname" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"syslog"`
}

func (p *LoggingSyslog) GetId() string{
    return "1"
}

func (p *LoggingSyslog) getPath() string{
    return "logging/syslog"
}

func (p *LoggingSyslog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingSyslog::Post")
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

func (p *LoggingSyslog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingSyslog::Get")
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
func (p *LoggingSyslog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingSyslog::Put")
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

func (p *LoggingSyslog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingSyslog::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
