

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTokenAuthenticationPlayerMode struct {
	Inst struct {

    Mode string `json:"mode" dval:"many-to-one"`
    Uuid string `json:"uuid"`

	} `json:"player-mode"`
}

func (p *DdosTokenAuthenticationPlayerMode) GetId() string{
    return "1"
}

func (p *DdosTokenAuthenticationPlayerMode) getPath() string{
    return "ddos/token-authentication/player-mode"
}

func (p *DdosTokenAuthenticationPlayerMode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTokenAuthenticationPlayerMode::Post")
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

func (p *DdosTokenAuthenticationPlayerMode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTokenAuthenticationPlayerMode::Get")
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
func (p *DdosTokenAuthenticationPlayerMode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTokenAuthenticationPlayerMode::Put")
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

func (p *DdosTokenAuthenticationPlayerMode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTokenAuthenticationPlayerMode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
