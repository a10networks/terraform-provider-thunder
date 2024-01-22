

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodTacplus struct {
	Inst struct {

    SecretEncrypted string `json:"secret-encrypted"`
    Tacplus int `json:"tacplus"`
    TacplusEncrypted string `json:"tacplus-encrypted"`
    TacplusPassword int `json:"tacplus-password"`
    TacplusPasswordString string `json:"tacplus-password-string"`
    TacplusPort int `json:"tacplus-port" dval:"49"`
    TacplusSecret int `json:"tacplus-secret"`
    TacplusSecretString string `json:"tacplus-secret-string"`
    TacplusType string `json:"tacplus-type" dval:"inbound-ascii-login"`
    TacplusUsername string `json:"tacplus-username"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"tacplus"`
}

func (p *HealthMonitorMethodTacplus) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodTacplus) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/tacplus"
}

func (p *HealthMonitorMethodTacplus) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTacplus::Post")
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

func (p *HealthMonitorMethodTacplus) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTacplus::Get")
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
func (p *HealthMonitorMethodTacplus) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTacplus::Put")
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

func (p *HealthMonitorMethodTacplus) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTacplus::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
