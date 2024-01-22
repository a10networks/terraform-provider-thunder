

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodCompound struct {
	Inst struct {

    Compound int `json:"compound"`
    RpnString string `json:"rpn-string"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"compound"`
}

func (p *HealthMonitorMethodCompound) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodCompound) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/compound"
}

func (p *HealthMonitorMethodCompound) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodCompound::Post")
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

func (p *HealthMonitorMethodCompound) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodCompound::Get")
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
func (p *HealthMonitorMethodCompound) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodCompound::Put")
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

func (p *HealthMonitorMethodCompound) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodCompound::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
