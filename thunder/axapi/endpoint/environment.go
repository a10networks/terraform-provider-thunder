

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Environment struct {
	Inst struct {

    ThresholdCfg EnvironmentThresholdCfg `json:"threshold-cfg"`
    UpdateInterval EnvironmentUpdateInterval `json:"update-interval"`
    Uuid string `json:"uuid"`

	} `json:"environment"`
}


type EnvironmentThresholdCfg struct {
    Low int `json:"low"`
    Medium int `json:"medium"`
    High int `json:"high"`
}


type EnvironmentUpdateInterval struct {
    Interval int `json:"interval" dval:"30"`
}

func (p *Environment) GetId() string{
    return "1"
}

func (p *Environment) getPath() string{
    return "environment"
}

func (p *Environment) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Environment::Post")
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

func (p *Environment) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Environment::Get")
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
func (p *Environment) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Environment::Put")
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

func (p *Environment) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Environment::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
