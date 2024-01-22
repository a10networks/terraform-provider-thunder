

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor struct {
	Inst struct {

    Activate int `json:"activate"`
    NeighborIpv4 string `json:"neighbor-ipv4"`
    NeighborRouteMapLists []RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    SendCommunityVal string `json:"send-community-val" dval:"both"`
    Uuid string `json:"uuid"`
    AsNumber string 

	} `json:"ipv4-neighbor"`
}


type RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4NeighborNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}

func (p *RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor) GetId() string{
    return p.Inst.NeighborIpv4
}

func (p *RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv6-flowspec/neighbor/ipv4-neighbor"
}

func (p *RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor::Post")
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

func (p *RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor::Get")
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
func (p *RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor::Put")
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

func (p *RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6FlowspecNeighborIpv4Neighbor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
