

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsMessageSelector struct {
	Inst struct {

    Name string `json:"name"`
    RuleList []AcosEventsMessageSelectorRuleList `json:"rule-list"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"message-selector"`
}


type AcosEventsMessageSelectorRuleList struct {
    Index int `json:"index"`
    Action string `json:"action" dval:"send"`
    MessageId string `json:"message-id"`
    MessageIdScope string `json:"message-id-scope"`
    SeverityOper string `json:"severity-oper"`
    SeverityVal string `json:"severity-val"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *AcosEventsMessageSelector) GetId() string{
    return p.Inst.Name
}

func (p *AcosEventsMessageSelector) getPath() string{
    return "acos-events/message-selector"
}

func (p *AcosEventsMessageSelector) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelector::Post")
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

func (p *AcosEventsMessageSelector) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelector::Get")
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
func (p *AcosEventsMessageSelector) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelector::Put")
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

func (p *AcosEventsMessageSelector) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelector::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
