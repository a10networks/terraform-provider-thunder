

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthGlobal struct {
	Inst struct {

    CheckRate int `json:"check-rate" dval:"1000"`
    DisableAutoAdjust int `json:"disable-auto-adjust"`
    ExternalRate int `json:"external-rate" dval:"2"`
    Interval int `json:"interval" dval:"5"`
    MultiProcess int `json:"multi-process" dval:"1"`
    Per int `json:"per" dval:"2"`
    Retry int `json:"retry" dval:"3"`
    Timeout int `json:"timeout" dval:"5"`
    UpRetry int `json:"up-retry" dval:"1"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}

func (p *HealthGlobal) GetId() string{
    return "1"
}

func (p *HealthGlobal) getPath() string{
    return "health/global"
}

func (p *HealthGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthGlobal::Post")
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

func (p *HealthGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthGlobal::Get")
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
func (p *HealthGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthGlobal::Put")
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

func (p *HealthGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
