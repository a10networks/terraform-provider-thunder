

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosInterfaceHttpHealthCheck struct {
	Inst struct {

    ChallengeMethod string `json:"challenge-method"`
    ChallengeRedirectCode string `json:"challenge-redirect-code" dval:"302"`
    ChallengeUriEncode int `json:"challenge-uri-encode"`
    Enable string `json:"enable"`
    Uuid string `json:"uuid"`

	} `json:"interface-http-health-check"`
}

func (p *DdosInterfaceHttpHealthCheck) GetId() string{
    return "1"
}

func (p *DdosInterfaceHttpHealthCheck) getPath() string{
    return "ddos/interface-http-health-check"
}

func (p *DdosInterfaceHttpHealthCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosInterfaceHttpHealthCheck::Post")
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

func (p *DdosInterfaceHttpHealthCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosInterfaceHttpHealthCheck::Get")
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
func (p *DdosInterfaceHttpHealthCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosInterfaceHttpHealthCheck::Put")
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

func (p *DdosInterfaceHttpHealthCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosInterfaceHttpHealthCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
