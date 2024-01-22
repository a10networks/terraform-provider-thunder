

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AutomaticUpdateConfig struct {
	Inst struct {

    Daily int `json:"daily"`
    DayTime string `json:"day-time"`
    Debug int `json:"debug"`
    DisableSslVerify int `json:"disable-ssl-verify"`
    FeatureName string `json:"feature-name"`
    Schedule int `json:"schedule"`
    Uuid string `json:"uuid"`
    WeekDay string `json:"week-day"`
    WeekTime string `json:"week-time"`
    Weekly int `json:"weekly"`

	} `json:"config"`
}

func (p *AutomaticUpdateConfig) GetId() string{
    return p.Inst.FeatureName
}

func (p *AutomaticUpdateConfig) getPath() string{
    return "automatic-update/config"
}

func (p *AutomaticUpdateConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateConfig::Post")
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

func (p *AutomaticUpdateConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateConfig::Get")
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
func (p *AutomaticUpdateConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateConfig::Put")
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

func (p *AutomaticUpdateConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
