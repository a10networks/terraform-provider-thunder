

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLoopbackIpRouterIsis struct {
	Inst struct {

    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"isis"`
}

func (p *InterfaceLoopbackIpRouterIsis) GetId() string{
    return "1"
}

func (p *InterfaceLoopbackIpRouterIsis) getPath() string{
    return "interface/loopback/" +p.Inst.Ifnum + "/ip/router/isis"
}

func (p *InterfaceLoopbackIpRouterIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRouterIsis::Post")
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

func (p *InterfaceLoopbackIpRouterIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRouterIsis::Get")
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
func (p *InterfaceLoopbackIpRouterIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRouterIsis::Put")
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

func (p *InterfaceLoopbackIpRouterIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLoopbackIpRouterIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
