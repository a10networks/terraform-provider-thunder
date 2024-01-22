

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor struct {
	Inst struct {

    Activate int `json:"activate"`
    AllowasIn int `json:"allowas-in"`
    AllowasInCount int `json:"allowas-in-count" dval:"3"`
    Inbound int `json:"inbound"`
    MaximumPrefix int `json:"maximum-prefix"`
    MaximumPrefixThres int `json:"maximum-prefix-thres"`
    NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    NextHopSelf int `json:"next-hop-self"`
    PeerGroup string `json:"peer-group"`
    RemovePrivateAs int `json:"remove-private-as"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight"`
    AsNumber string 

	} `json:"peer-group-neighbor"`
}


type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}

func (p *RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor) GetId() string{
    return p.Inst.PeerGroup
}

func (p *RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv6/neighbor/peer-group-neighbor"
}

func (p *RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor::Post")
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

func (p *RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor::Get")
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
func (p *RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor::Put")
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

func (p *RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
