

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsMessageId struct {
	Inst struct {

    LogMsg string `json:"log-msg"`
    MessageIdScopeRoute string `json:"message-id-scope-route"`
    Property AcosEventsMessageIdProperty56 `json:"property"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"message-id"`
}


type AcosEventsMessageIdProperty56 struct {
    Severity AcosEventsMessageIdPropertySeverity57 `json:"severity"`
    LogRoute AcosEventsMessageIdPropertyLogRoute58 `json:"log-route"`
    RateLimit AcosEventsMessageIdPropertyRateLimit59 `json:"rate-limit"`
}


type AcosEventsMessageIdPropertySeverity57 struct {
    SeverityVal string `json:"severity-val"`
    Uuid string `json:"uuid"`
}


type AcosEventsMessageIdPropertyLogRoute58 struct {
    LogRouteVal string `json:"log-route-val"`
    Uuid string `json:"uuid"`
}


type AcosEventsMessageIdPropertyRateLimit59 struct {
    RateLimitVal string `json:"rate-limit-val"`
    Uuid string `json:"uuid"`
}

func (p *AcosEventsMessageId) GetId() string{
    return p.Inst.LogMsg+"+"+p.Inst.MessageIdScopeRoute
}

func (p *AcosEventsMessageId) getPath() string{
    return "acos-events/message-id"
}

func (p *AcosEventsMessageId) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageId::Post")
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

func (p *AcosEventsMessageId) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageId::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *AcosEventsMessageId) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageId::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *AcosEventsMessageId) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageId::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
