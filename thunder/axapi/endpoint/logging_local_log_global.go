

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogGlobal struct {
	Inst struct {

    SamplingEnable []LoggingLocalLogGlobalSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type LoggingLocalLogGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *LoggingLocalLogGlobal) GetId() string{
    return "1"
}

func (p *LoggingLocalLogGlobal) getPath() string{
    return "logging/local-log/global"
}

func (p *LoggingLocalLogGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogGlobal::Post")
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

func (p *LoggingLocalLogGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogGlobal::Get")
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
func (p *LoggingLocalLogGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogGlobal::Put")
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

func (p *LoggingLocalLogGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
