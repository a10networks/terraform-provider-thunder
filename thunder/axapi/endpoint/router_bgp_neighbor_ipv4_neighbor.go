

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpNeighborIpv4Neighbor struct {
	Inst struct {

    AcosApplicationOnly int `json:"acos-application-only"`
    Activate int `json:"activate" dval:"1"`
    AdvertisementInterval int `json:"advertisement-interval"`
    AllowasIn int `json:"allowas-in"`
    AllowasInCount int `json:"allowas-in-count" dval:"3"`
    AsOriginationInterval int `json:"as-origination-interval"`
    Bfd int `json:"bfd"`
    BfdEncrypted string `json:"bfd-encrypted"`
    BfdValue string `json:"bfd-value"`
    CollideEstablished int `json:"collide-established"`
    Connect int `json:"connect"`
    DefaultOriginate int `json:"default-originate"`
    Description string `json:"description"`
    DisallowInfiniteHoldtime int `json:"disallow-infinite-holdtime"`
    DistributeLists []RouterBgpNeighborIpv4NeighborDistributeLists `json:"distribute-lists"`
    DontCapabilityNegotiate int `json:"dont-capability-negotiate"`
    Dynamic int `json:"dynamic"`
    EbgpMultihop int `json:"ebgp-multihop"`
    EbgpMultihopHopCount int `json:"ebgp-multihop-hop-count"`
    EnforceMultihop int `json:"enforce-multihop"`
    Ethernet int `json:"ethernet"`
    GracefulRestart int `json:"graceful-restart"`
    Inbound int `json:"inbound"`
    KeyId int `json:"key-id"`
    KeyType string `json:"key-type"`
    Lif string `json:"lif"`
    Loopback int `json:"loopback"`
    MaximumPrefix int `json:"maximum-prefix"`
    MaximumPrefixThres int `json:"maximum-prefix-thres"`
    Multihop int `json:"multihop"`
    NbrRemoteAs string `json:"nbr-remote-as"`
    NeighborFilterLists []RouterBgpNeighborIpv4NeighborNeighborFilterLists `json:"neighbor-filter-lists"`
    NeighborIpv4 string `json:"neighbor-ipv4"`
    NeighborPrefixLists []RouterBgpNeighborIpv4NeighborNeighborPrefixLists `json:"neighbor-prefix-lists"`
    NeighborRouteMapLists []RouterBgpNeighborIpv4NeighborNeighborRouteMapLists `json:"neighbor-route-map-lists"`
    NextHopSelf int `json:"next-hop-self"`
    OverrideCapability int `json:"override-capability"`
    PassEncrypted string `json:"pass-encrypted"`
    PassValue string `json:"pass-value"`
    Passive int `json:"passive"`
    PeerGroupName string `json:"peer-group-name"`
    PrefixListDirection string `json:"prefix-list-direction"`
    RemovePrivateAs int `json:"remove-private-as"`
    RestartMin int `json:"restart-min"`
    RouteMap string `json:"route-map"`
    RouteRefresh int `json:"route-refresh" dval:"1"`
    SendCommunityVal string `json:"send-community-val" dval:"both"`
    Shutdown int `json:"shutdown"`
    StrictCapabilityMatch int `json:"strict-capability-match"`
    Telemetry int `json:"telemetry"`
    TimersHoldtime int `json:"timers-holdtime" dval:"90"`
    TimersKeepalive int `json:"timers-keepalive" dval:"30"`
    Trunk int `json:"trunk"`
    Tunnel int `json:"tunnel"`
    UnsuppressMap string `json:"unsuppress-map"`
    UpdateSourceIp string `json:"update-source-ip"`
    UpdateSourceIpv6 string `json:"update-source-ipv6"`
    Uuid string `json:"uuid"`
    Ve int `json:"ve"`
    Weight int `json:"weight"`
    AsNumber string 

	} `json:"ipv4-neighbor"`
}


type RouterBgpNeighborIpv4NeighborDistributeLists struct {
    DistributeList string `json:"distribute-list"`
    DistributeListDirection string `json:"distribute-list-direction"`
}


type RouterBgpNeighborIpv4NeighborNeighborFilterLists struct {
    FilterList string `json:"filter-list"`
    FilterListDirection string `json:"filter-list-direction"`
}


type RouterBgpNeighborIpv4NeighborNeighborPrefixLists struct {
    NbrPrefixList string `json:"nbr-prefix-list"`
    NbrPrefixListDirection string `json:"nbr-prefix-list-direction"`
}


type RouterBgpNeighborIpv4NeighborNeighborRouteMapLists struct {
    NbrRouteMap string `json:"nbr-route-map"`
    NbrRmapDirection string `json:"nbr-rmap-direction"`
}

func (p *RouterBgpNeighborIpv4Neighbor) GetId() string{
    return p.Inst.NeighborIpv4
}

func (p *RouterBgpNeighborIpv4Neighbor) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/neighbor/ipv4-neighbor"
}

func (p *RouterBgpNeighborIpv4Neighbor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborIpv4Neighbor::Post")
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

func (p *RouterBgpNeighborIpv4Neighbor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborIpv4Neighbor::Get")
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
func (p *RouterBgpNeighborIpv4Neighbor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborIpv4Neighbor::Put")
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

func (p *RouterBgpNeighborIpv4Neighbor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpNeighborIpv4Neighbor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
