

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsTemplateMessageSelector struct {
	Inst struct {

    CollectorGroupList []AcosEventsTemplateMessageSelectorCollectorGroupList `json:"collector-group-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"message-selector"`
}


type AcosEventsTemplateMessageSelectorCollectorGroupList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *AcosEventsTemplateMessageSelector) GetId() string{
    return p.Inst.Name
}

func (p *AcosEventsTemplateMessageSelector) getPath() string{
    return "acos-events/template/message-selector"
}

func (p *AcosEventsTemplateMessageSelector) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsTemplateMessageSelector::Post")
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

func (p *AcosEventsTemplateMessageSelector) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsTemplateMessageSelector::Get")
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
func (p *AcosEventsTemplateMessageSelector) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsTemplateMessageSelector::Put")
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

func (p *AcosEventsTemplateMessageSelector) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AcosEventsTemplateMessageSelector::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
