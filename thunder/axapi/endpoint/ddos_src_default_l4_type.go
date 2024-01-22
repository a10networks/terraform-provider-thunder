

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcDefaultL4Type struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Protocol string `json:"protocol"`
    Template DdosSrcDefaultL4TypeTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"l4-type"`
}


type DdosSrcDefaultL4TypeTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}

func (p *DdosSrcDefaultL4Type) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosSrcDefaultL4Type) getPath() string{
    return "ddos/src/default/" +p.Inst.DefaultAddressType + "/l4-type"
}

func (p *DdosSrcDefaultL4Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefaultL4Type::Post")
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

func (p *DdosSrcDefaultL4Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefaultL4Type::Get")
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
func (p *DdosSrcDefaultL4Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefaultL4Type::Put")
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

func (p *DdosSrcDefaultL4Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefaultL4Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
