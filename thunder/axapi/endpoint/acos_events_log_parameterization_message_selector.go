

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsLogParameterizationMessageSelector struct {
	Inst struct {

    RuleList []AcosEventsLogParameterizationMessageSelectorRuleList `json:"rule-list"`
    Uuid string `json:"uuid"`

	} `json:"message-selector"`
}


type AcosEventsLogParameterizationMessageSelectorRuleList struct {
    Index int `json:"index"`
    Action string `json:"action" dval:"send"`
    MessageId string `json:"message-id"`
    MessageIdScope string `json:"message-id-scope"`
    SeverityOper string `json:"severity-oper"`
    SeverityVal string `json:"severity-val"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *AcosEventsLogParameterizationMessageSelector) GetId() string{
    return "1"
}

func (p *AcosEventsLogParameterizationMessageSelector) getPath() string{
    return "acos-events/log-parameterization/message-selector"
}

func (p *AcosEventsLogParameterizationMessageSelector) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogParameterizationMessageSelector::Post")
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

func (p *AcosEventsLogParameterizationMessageSelector) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogParameterizationMessageSelector::Get")
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
func (p *AcosEventsLogParameterizationMessageSelector) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogParameterizationMessageSelector::Put")
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

func (p *AcosEventsLogParameterizationMessageSelector) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsLogParameterizationMessageSelector::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
