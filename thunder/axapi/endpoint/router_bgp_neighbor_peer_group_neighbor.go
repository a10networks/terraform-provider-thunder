

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpNeighborPeerGroupNeighbor struct {
	Inst struct {

    Activate int `json:"activate" dval:"1"`
    AdvertisementInterval int `json:"advertisement-interval"`
    AllowasIn int `json:"allowas-in"`
    AllowasInCount int `json:"allowas-in-count" dval:"3"`
    AsOriginationInterval int `json:"as-origination-interval"`
    Bfd int `json:"bfd"`
    CollideEstablished int `json:"collide-established"`
    Connect int `json:"connect"`
    DefaultOriginate int `json:"default-originate"`
    Description string `json:"description"`
    DontCapabilityNegotiate int `json:"dont-capability-negotiate"`
    Dynamic int `json:"dynamic"`
    EbgpMultihop int `json:"ebgp-multihop"`
    EbgpMultihopHopCount int `json:"ebgp-multihop-hop-count"`
    EnforceMultihop int `json:"enforce-multihop"`
    Ethernet int `json:"ethernet"`
    ExtendedNexthop int `json:"extended-nexthop"`
    Inbound int `json:"inbound"`
    Lif string `json:"lif"`
    Loopback int `json:"loopback"`
    MaximumPrefix int `json:"maximum-prefix"`
    MaximumPrefixThres int `json:"maximum-prefix-thres"`
    Multihop int `json:"multihop"`
    NeighborRouteMapLists []RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    OverrideCapability int `json:"override-capability"`
    PassEncrypted string `json:"pass-encrypted"`
    PassValue string `json:"pass-value"`
    Passive int `json:"passive"`
    PeerGroup string `json:"peer-group"`
    PeerGroupKey int `json:"peer-group-key"`
    PeerGroupRemoteAs string `json:"peer-group-remote-as"`
    RemovePrivateAs int `json:"remove-private-as"`
    RouteMap string `json:"route-map"`
    RouteRefresh int `json:"route-refresh" dval:"1"`
    Shutdown int `json:"shutdown"`
    StrictCapabilityMatch int `json:"strict-capability-match"`
    TimersHoldtime int `json:"timers-holdtime" dval:"90"`
    TimersKeepalive int `json:"timers-keepalive" dval:"30"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    UpdateSourceIp string `json:"update-source-ip"`
    UpdateSourceIpv6 string `json:"update-source-ipv6"`
    Uuid string `json:"uuid"`
    Ve int `json:"ve"`
    Weight int `json:"weight"`
    AsNumber string 

	} `json:"peer-group-neighbor"`
}


type RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}

func (p *RouterBgpNeighborPeerGroupNeighbor) GetId() string{
    return p.Inst.PeerGroup
}

func (p *RouterBgpNeighborPeerGroupNeighbor) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/neighbor/peer-group-neighbor"
}

func (p *RouterBgpNeighborPeerGroupNeighbor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborPeerGroupNeighbor::Post")
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

func (p *RouterBgpNeighborPeerGroupNeighbor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborPeerGroupNeighbor::Get")
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
func (p *RouterBgpNeighborPeerGroupNeighbor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborPeerGroupNeighbor::Put")
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

func (p *RouterBgpNeighborPeerGroupNeighbor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborPeerGroupNeighbor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
