

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwHelperSessions struct {
	Inst struct {

    IdleTimeout int `json:"idle-timeout" dval:"1"`
    Limit int `json:"limit"`
    Mode string `json:"mode"`
    Uuid string `json:"uuid"`

	} `json:"helper-sessions"`
}

func (p *FwHelperSessions) GetId() string{
    return "1"
}

func (p *FwHelperSessions) getPath() string{
    return "fw/helper-sessions"
}

func (p *FwHelperSessions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwHelperSessions::Post")
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

func (p *FwHelperSessions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwHelperSessions::Get")
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
func (p *FwHelperSessions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwHelperSessions::Put")
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

func (p *FwHelperSessions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwHelperSessions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
