

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsMessageSelectorRule struct {
	Inst struct {

    Action string `json:"action" dval:"send"`
    Index int `json:"index"`
    MessageId string `json:"message-id"`
    MessageIdScope string `json:"message-id-scope"`
    SeverityOper string `json:"severity-oper"`
    SeverityVal string `json:"severity-val"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"rule"`
}

func (p *AcosEventsMessageSelectorRule) GetId() string{
    return strconv.Itoa(p.Inst.Index)
}

func (p *AcosEventsMessageSelectorRule) getPath() string{
    return "acos-events/message-selector/" +p.Inst.Name + "/rule"
}

func (p *AcosEventsMessageSelectorRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelectorRule::Post")
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

func (p *AcosEventsMessageSelectorRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelectorRule::Get")
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
func (p *AcosEventsMessageSelectorRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelectorRule::Put")
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

func (p *AcosEventsMessageSelectorRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsMessageSelectorRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
