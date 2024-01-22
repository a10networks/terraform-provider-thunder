

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifEncapsulationDot1q struct {
	Inst struct {

    Ethernet int `json:"ethernet"`
    Tag int `json:"tag"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"dot1q"`
}

func (p *InterfaceLifEncapsulationDot1q) GetId() string{
    return "1"
}

func (p *InterfaceLifEncapsulationDot1q) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/encapsulation/dot1q"
}

func (p *InterfaceLifEncapsulationDot1q) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifEncapsulationDot1q::Post")
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

func (p *InterfaceLifEncapsulationDot1q) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifEncapsulationDot1q::Get")
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
func (p *InterfaceLifEncapsulationDot1q) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifEncapsulationDot1q::Put")
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

func (p *InterfaceLifEncapsulationDot1q) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifEncapsulationDot1q::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
