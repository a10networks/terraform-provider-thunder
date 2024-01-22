

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsMessageIdPropertyRateLimit struct {
	Inst struct {

    RateLimitVal string `json:"rate-limit-val"`
    Uuid string `json:"uuid"`
    LogMsg string 
    MessageIdScopeRoute string 

	} `json:"rate-limit"`
}

func (p *AcosEventsMessageIdPropertyRateLimit) GetId() string{
    return "1"
}

func (p *AcosEventsMessageIdPropertyRateLimit) getPath() string{
    return "acos-events/message-id/" +p.Inst.LogMsg + "+" +p.Inst.MessageIdScopeRoute + "/property/rate-limit"
}

func (p *AcosEventsMessageIdPropertyRateLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageIdPropertyRateLimit::Post")
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

func (p *AcosEventsMessageIdPropertyRateLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageIdPropertyRateLimit::Get")
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
func (p *AcosEventsMessageIdPropertyRateLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageIdPropertyRateLimit::Put")
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

func (p *AcosEventsMessageIdPropertyRateLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageIdPropertyRateLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
