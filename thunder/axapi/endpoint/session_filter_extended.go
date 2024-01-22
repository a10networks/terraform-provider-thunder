

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SessionFilterExtended struct {
	Inst struct {

    FilterRel string `json:"filter-rel" dval:"AND"`
    FwdFilter string `json:"fwd-filter"`
    Name string `json:"name"`
    RevFilter string `json:"rev-filter"`
    SessionType string `json:"session-type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"session-filter-extended"`
}

func (p *SessionFilterExtended) GetId() string{
    return p.Inst.Name
}

func (p *SessionFilterExtended) getPath() string{
    return "session-filter-extended"
}

func (p *SessionFilterExtended) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilterExtended::Post")
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

func (p *SessionFilterExtended) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilterExtended::Get")
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
func (p *SessionFilterExtended) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilterExtended::Put")
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

func (p *SessionFilterExtended) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SessionFilterExtended::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
