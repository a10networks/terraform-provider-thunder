

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AutomaticUpdateCheckNow struct {
	Inst struct {

    FeatureName string `json:"feature-name"`
    FromStagingServer int `json:"from-staging-server"`
    ProdVer string `json:"prod-ver"`
    StageVer string `json:"stage-ver"`

	} `json:"check-now"`
}

func (p *AutomaticUpdateCheckNow) GetId() string{
    return "1"
}

func (p *AutomaticUpdateCheckNow) getPath() string{
    return "automatic-update/check-now"
}

func (p *AutomaticUpdateCheckNow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateCheckNow::Post")
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

func (p *AutomaticUpdateCheckNow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateCheckNow::Get")
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
func (p *AutomaticUpdateCheckNow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateCheckNow::Put")
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

func (p *AutomaticUpdateCheckNow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AutomaticUpdateCheckNow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
