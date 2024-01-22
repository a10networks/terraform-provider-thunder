

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingLocalLogAppFw struct {
	Inst struct {

    DotPlot LoggingLocalLogAppFwDotPlot1050 `json:"dot-plot"`
    TopN LoggingLocalLogAppFwTopN1051 `json:"top-n"`
    Uuid string `json:"uuid"`

	} `json:"app-fw"`
}


type LoggingLocalLogAppFwDotPlot1050 struct {
    Uuid string `json:"uuid"`
}


type LoggingLocalLogAppFwTopN1051 struct {
    Uuid string `json:"uuid"`
}

func (p *LoggingLocalLogAppFw) GetId() string{
    return "1"
}

func (p *LoggingLocalLogAppFw) getPath() string{
    return "logging/local-log/app-fw"
}

func (p *LoggingLocalLogAppFw) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogAppFw::Post")
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

func (p *LoggingLocalLogAppFw) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogAppFw::Get")
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
func (p *LoggingLocalLogAppFw) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogAppFw::Put")
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

func (p *LoggingLocalLogAppFw) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingLocalLogAppFw::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
