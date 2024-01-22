

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NeighborStatic struct {
	Inst struct {

    Ethernet int `json:"ethernet"`
    Ipv6Addr string `json:"ipv6-addr"`
    Mac string `json:"mac"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    Uuid string `json:"uuid"`
    Vlan int `json:"vlan"`

	} `json:"static"`
}

func (p *Ipv6NeighborStatic) GetId() string{
    return p.Inst.Ipv6Addr+"+"+strconv.Itoa(p.Inst.Vlan)
}

func (p *Ipv6NeighborStatic) getPath() string{
    return "ipv6/neighbor/static"
}

func (p *Ipv6NeighborStatic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NeighborStatic::Post")
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

func (p *Ipv6NeighborStatic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NeighborStatic::Get")
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
func (p *Ipv6NeighborStatic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NeighborStatic::Put")
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

func (p *Ipv6NeighborStatic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NeighborStatic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
