

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NatInsideSourceList struct {
	Inst struct {

    ListName string `json:"list-name"`
    Pool string `json:"pool"`
    Uuid string `json:"uuid"`

	} `json:"list"`
}

func (p *Ipv6NatInsideSourceList) GetId() string{
    return p.Inst.ListName
}

func (p *Ipv6NatInsideSourceList) getPath() string{
    return "ipv6/nat/inside/source/list"
}

func (p *Ipv6NatInsideSourceList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatInsideSourceList::Post")
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

func (p *Ipv6NatInsideSourceList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatInsideSourceList::Get")
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
func (p *Ipv6NatInsideSourceList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatInsideSourceList::Put")
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

func (p *Ipv6NatInsideSourceList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatInsideSourceList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
