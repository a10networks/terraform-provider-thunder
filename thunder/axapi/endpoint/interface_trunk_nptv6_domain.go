

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkNptv6Domain struct {
	Inst struct {

    BindType string `json:"bind-type"`
    DomainName string `json:"domain-name"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"domain"`
}

func (p *InterfaceTrunkNptv6Domain) GetId() string{
    return p.Inst.DomainName+"+"+p.Inst.BindType
}

func (p *InterfaceTrunkNptv6Domain) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/nptv6/domain"
}

func (p *InterfaceTrunkNptv6Domain) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkNptv6Domain::Post")
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

func (p *InterfaceTrunkNptv6Domain) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkNptv6Domain::Get")
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
func (p *InterfaceTrunkNptv6Domain) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkNptv6Domain::Put")
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

func (p *InterfaceTrunkNptv6Domain) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkNptv6Domain::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
