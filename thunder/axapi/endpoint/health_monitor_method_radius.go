

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodRadius struct {
	Inst struct {

    Radius int `json:"radius"`
    RadiusEncrypted string `json:"radius-encrypted"`
    RadiusExpect int `json:"radius-expect"`
    RadiusPasswordString string `json:"radius-password-string"`
    RadiusPort int `json:"radius-port" dval:"1812"`
    RadiusResponseCode string `json:"radius-response-code"`
    RadiusSecret string `json:"radius-secret"`
    RadiusSecretEncrypted string `json:"radius-secret-encrypted"`
    RadiusUsername string `json:"radius-username"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"radius"`
}

func (p *HealthMonitorMethodRadius) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodRadius) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/radius"
}

func (p *HealthMonitorMethodRadius) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodRadius::Post")
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

func (p *HealthMonitorMethodRadius) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodRadius::Get")
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
func (p *HealthMonitorMethodRadius) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodRadius::Put")
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

func (p *HealthMonitorMethodRadius) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodRadius::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
