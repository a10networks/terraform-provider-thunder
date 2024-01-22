

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6Reroute struct {
	Inst struct {

    SuppressProtocols Ipv6RerouteSuppressProtocols1044 `json:"suppress-protocols"`
    Uuid string `json:"uuid"`

	} `json:"reroute"`
}


type Ipv6RerouteSuppressProtocols1044 struct {
    Ospf int `json:"ospf"`
    Ebgp int `json:"ebgp"`
    Ibgp int `json:"ibgp"`
    Static int `json:"static"`
    Isis int `json:"isis"`
    Rip int `json:"rip"`
    Connected int `json:"connected"`
    Uuid string `json:"uuid"`
}

func (p *Ipv6Reroute) GetId() string{
    return "1"
}

func (p *Ipv6Reroute) getPath() string{
    return "ipv6/reroute"
}

func (p *Ipv6Reroute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Reroute::Post")
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

func (p *Ipv6Reroute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Reroute::Get")
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
func (p *Ipv6Reroute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Reroute::Put")
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

func (p *Ipv6Reroute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6Reroute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
