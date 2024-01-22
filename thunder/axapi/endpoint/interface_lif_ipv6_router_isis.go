

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifIpv6RouterIsis struct {
	Inst struct {

    Tag string `json:"tag"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"isis"`
}

func (p *InterfaceLifIpv6RouterIsis) GetId() string{
    return "1"
}

func (p *InterfaceLifIpv6RouterIsis) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/ipv6/router/isis"
}

func (p *InterfaceLifIpv6RouterIsis) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6RouterIsis::Post")
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

func (p *InterfaceLifIpv6RouterIsis) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6RouterIsis::Get")
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
func (p *InterfaceLifIpv6RouterIsis) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6RouterIsis::Put")
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

func (p *InterfaceLifIpv6RouterIsis) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpv6RouterIsis::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
