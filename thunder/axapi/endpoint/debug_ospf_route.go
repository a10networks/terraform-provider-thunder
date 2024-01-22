

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugOspfRoute struct {
	Inst struct {

    Ase int `json:"ase"`
    Ia int `json:"ia"`
    Install int `json:"install"`
    Spf int `json:"spf"`
    Uuid string `json:"uuid"`

	} `json:"route"`
}

func (p *DebugOspfRoute) GetId() string{
    return "1"
}

func (p *DebugOspfRoute) getPath() string{
    return "debug/ospf/route"
}

func (p *DebugOspfRoute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfRoute::Post")
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

func (p *DebugOspfRoute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfRoute::Get")
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
func (p *DebugOspfRoute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfRoute::Put")
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

func (p *DebugOspfRoute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfRoute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
