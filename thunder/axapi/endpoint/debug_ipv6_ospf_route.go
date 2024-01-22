

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugIpv6OspfRoute struct {
	Inst struct {

    Ase int `json:"ase"`
    Ia int `json:"ia"`
    Install int `json:"install"`
    Spf int `json:"spf"`
    Uuid string `json:"uuid"`

	} `json:"route"`
}

func (p *DebugIpv6OspfRoute) GetId() string{
    return "1"
}

func (p *DebugIpv6OspfRoute) getPath() string{
    return "debug/ipv6/ospf/route"
}

func (p *DebugIpv6OspfRoute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6OspfRoute::Post")
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

func (p *DebugIpv6OspfRoute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6OspfRoute::Get")
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
func (p *DebugIpv6OspfRoute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6OspfRoute::Put")
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

func (p *DebugIpv6OspfRoute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugIpv6OspfRoute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
