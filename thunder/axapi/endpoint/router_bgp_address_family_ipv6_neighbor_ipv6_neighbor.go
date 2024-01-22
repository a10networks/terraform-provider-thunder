

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor struct {
	Inst struct {

    Activate int `json:"activate"`
    AllowasIn int `json:"allowas-in"`
    AllowasInCount int `json:"allowas-in-count" dval:"3"`
    DefaultOriginate int `json:"default-originate"`
    DistributeLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists `json:"distribute-lists"`
    GracefulRestart int `json:"graceful-restart"`
    Inbound int `json:"inbound"`
    MaximumPrefix int `json:"maximum-prefix"`
    MaximumPrefixThres int `json:"maximum-prefix-thres"`
    NeighborFilterLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists `json:"neighbor-filter-lists"`
    NeighborIpv6 string `json:"neighbor-ipv6"`
    NeighborPrefixLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists `json:"neighbor-prefix-lists"`
    NeighborRouteMapLists []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    NextHopSelf int `json:"next-hop-self"`
    PeerGroupName string `json:"peer-group-name"`
    PrefixListDirection string `json:"prefix-list-direction"`
    RemovePrivateAs int `json:"remove-private-as"`
    RestartMin int `json:"restart-min"`
    RouteMap string `json:"route-map"`
    SendCommunityVal string `json:"send-community-val" dval:"both"`
    UnsuppressMap string `json:"unsuppress-map"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight"`
    AsNumber string 

	} `json:"ipv6-neighbor"`
}


type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists struct {
    DistributeList string `json:"distribute-list"`
    DistributeListDirection string `json:"distribute-list-direction"`
}


type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists struct {
    FilterList string `json:"filter-list"`
    FilterListDirection string `json:"filter-list-direction"`
}


type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists struct {
    NbrPrefixList string `json:"nbr-prefix-list"`
    NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}


type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}

func (p *RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor) GetId() string{
    return p.Inst.NeighborIpv6
}

func (p *RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv6/neighbor/ipv6-neighbor"
}

func (p *RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor::Post")
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

func (p *RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor::Get")
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
func (p *RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor::Put")
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

func (p *RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
