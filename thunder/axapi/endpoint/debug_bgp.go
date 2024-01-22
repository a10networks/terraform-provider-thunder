

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugBgp struct {
	Inst struct {

    All int `json:"all"`
    Bfd int `json:"bfd"`
    Dampening int `json:"dampening"`
    Events int `json:"events"`
    Filters int `json:"filters"`
    Fsm int `json:"fsm"`
    In int `json:"in"`
    Keepalives int `json:"keepalives"`
    Nht int `json:"nht"`
    Nsm int `json:"nsm"`
    Out int `json:"out"`
    Updates int `json:"updates"`
    Uuid string `json:"uuid"`

	} `json:"bgp"`
}

func (p *DebugBgp) GetId() string{
    return "1"
}

func (p *DebugBgp) getPath() string{
    return "debug/bgp"
}

func (p *DebugBgp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugBgp::Post")
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

func (p *DebugBgp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugBgp::Get")
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
func (p *DebugBgp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugBgp::Put")
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

func (p *DebugBgp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugBgp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
