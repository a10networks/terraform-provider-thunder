

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv4Flowspec struct {
	Inst struct {

    Neighbor RouterBgpAddressFamilyIpv4FlowspecNeighbor1151 `json:"neighbor"`
    Uuid string `json:"uuid"`
    AsNumber string 

	} `json:"ipv4-flowspec"`
}


type RouterBgpAddressFamilyIpv4FlowspecNeighbor1151 struct {
    Ipv4NeighborList []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList `json:"ipv4-neighbor-list"`
    Ipv6NeighborList []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList `json:"ipv6-neighbor-list"`
}


type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborList struct {
    NeighborIpv4 string `json:"neighbor-ipv4"`
    Activate int `json:"activate"`
    NeighborRouteMapLists []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    SendCommunityVal string `json:"send-community-val" dval:"both"`
    Uuid string `json:"uuid"`
}


type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv4NeighborListNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}


type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborList struct {
    NeighborIpv6 string `json:"neighbor-ipv6"`
    Activate int `json:"activate"`
    NeighborRouteMapLists []RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    SendCommunityVal string `json:"send-community-val" dval:"both"`
    Uuid string `json:"uuid"`
}


type RouterBgpAddressFamilyIpv4FlowspecNeighborIpv6NeighborListNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}

func (p *RouterBgpAddressFamilyIpv4Flowspec) GetId() string{
    return "1"
}

func (p *RouterBgpAddressFamilyIpv4Flowspec) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv4-flowspec"
}

func (p *RouterBgpAddressFamilyIpv4Flowspec) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv4Flowspec::Post")
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

func (p *RouterBgpAddressFamilyIpv4Flowspec) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv4Flowspec::Get")
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
func (p *RouterBgpAddressFamilyIpv4Flowspec) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv4Flowspec::Put")
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

func (p *RouterBgpAddressFamilyIpv4Flowspec) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv4Flowspec::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
