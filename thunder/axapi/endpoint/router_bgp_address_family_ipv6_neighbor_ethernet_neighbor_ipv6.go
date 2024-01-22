

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 struct {
	Inst struct {

    Ethernet int `json:"ethernet"`
    PeerGroupName string `json:"peer-group-name"`
    Uuid string `json:"uuid"`
    AsNumber string 

	} `json:"ethernet-neighbor-ipv6"`
}

func (p *RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6) GetId() string{
    return strconv.Itoa(p.Inst.Ethernet)
}

func (p *RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv6/neighbor/ethernet-neighbor-ipv6"
}

func (p *RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6::Post")
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

func (p *RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6::Get")
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
func (p *RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6::Put")
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

func (p *RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
