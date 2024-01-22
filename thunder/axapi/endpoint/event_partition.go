

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EventPartition struct {
	Inst struct {

    Email string `json:"email"`
    Logging string `json:"logging"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VnpEvents string `json:"vnp-events"`

	} `json:"partition"`
}

func (p *EventPartition) GetId() string{
    return p.Inst.VnpEvents
}

func (p *EventPartition) getPath() string{
    return "event/partition"
}

func (p *EventPartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EventPartition::Post")
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

func (p *EventPartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EventPartition::Get")
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
func (p *EventPartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EventPartition::Put")
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

func (p *EventPartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EventPartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
