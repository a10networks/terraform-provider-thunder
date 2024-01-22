

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSportRateLimit struct {
	Inst struct {

    SamplingEnable []SlbSportRateLimitSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"sport-rate-limit"`
}


type SlbSportRateLimitSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbSportRateLimit) GetId() string{
    return "1"
}

func (p *SlbSportRateLimit) getPath() string{
    return "slb/sport-rate-limit"
}

func (p *SlbSportRateLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSportRateLimit::Post")
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

func (p *SlbSportRateLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSportRateLimit::Get")
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
func (p *SlbSportRateLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSportRateLimit::Put")
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

func (p *SlbSportRateLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSportRateLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
