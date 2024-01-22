

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwRateLimit struct {
	Inst struct {

    Summary FwRateLimitSummary373 `json:"summary"`
    Uuid string `json:"uuid"`

	} `json:"rate-limit"`
}


type FwRateLimitSummary373 struct {
    Uuid string `json:"uuid"`
}

func (p *FwRateLimit) GetId() string{
    return "1"
}

func (p *FwRateLimit) getPath() string{
    return "fw/rate-limit"
}

func (p *FwRateLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwRateLimit::Post")
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

func (p *FwRateLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwRateLimit::Get")
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
func (p *FwRateLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwRateLimit::Put")
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

func (p *FwRateLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwRateLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
