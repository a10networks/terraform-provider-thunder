

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugOspfLsa struct {
	Inst struct {

    Flooding int `json:"flooding"`
    Gererate int `json:"gererate"`
    Install int `json:"install"`
    Maxage int `json:"maxage"`
    Refresh int `json:"refresh"`
    Uuid string `json:"uuid"`

	} `json:"lsa"`
}

func (p *DebugOspfLsa) GetId() string{
    return "1"
}

func (p *DebugOspfLsa) getPath() string{
    return "debug/ospf/lsa"
}

func (p *DebugOspfLsa) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfLsa::Post")
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

func (p *DebugOspfLsa) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfLsa::Get")
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
func (p *DebugOspfLsa) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfLsa::Put")
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

func (p *DebugOspfLsa) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugOspfLsa::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
