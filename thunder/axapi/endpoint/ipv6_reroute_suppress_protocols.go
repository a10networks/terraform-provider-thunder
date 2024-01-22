

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RerouteSuppressProtocols struct {
	Inst struct {

    Connected int `json:"connected"`
    Ebgp int `json:"ebgp"`
    Ibgp int `json:"ibgp"`
    Isis int `json:"isis"`
    Ospf int `json:"ospf"`
    Rip int `json:"rip"`
    Static int `json:"static"`
    Uuid string `json:"uuid"`

	} `json:"suppress-protocols"`
}

func (p *Ipv6RerouteSuppressProtocols) GetId() string{
    return "1"
}

func (p *Ipv6RerouteSuppressProtocols) getPath() string{
    return "ipv6/reroute/suppress-protocols"
}

func (p *Ipv6RerouteSuppressProtocols) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RerouteSuppressProtocols::Post")
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

func (p *Ipv6RerouteSuppressProtocols) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RerouteSuppressProtocols::Get")
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
func (p *Ipv6RerouteSuppressProtocols) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RerouteSuppressProtocols::Put")
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

func (p *Ipv6RerouteSuppressProtocols) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RerouteSuppressProtocols::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
