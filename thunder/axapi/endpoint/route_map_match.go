

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type RouteMapMatch struct {
	Inst struct {

    AsPath RouteMapMatchAsPath `json:"as-path"`
    Community RouteMapMatchCommunity `json:"community"`
    Extcommunity RouteMapMatchExtcommunity `json:"extcommunity"`
    Group RouteMapMatchGroup `json:"group"`
    Interface RouteMapMatchInterface `json:"interface"`
    Ip RouteMapMatchIp `json:"ip"`
    Ipv6 RouteMapMatchIpv6 `json:"ipv6"`
    LargeCommunity RouteMapMatchLargeCommunity `json:"large-community"`
    LocalPreference RouteMapMatchLocalPreference `json:"local-preference"`
    Metric RouteMapMatchMetric `json:"metric"`
    Origin RouteMapMatchOrigin `json:"origin"`
    RouteType RouteMapMatchRouteType `json:"route-type"`
    Scaleout RouteMapMatchScaleout `json:"scaleout"`
    Tag RouteMapMatchTag `json:"tag"`
    Uuid string `json:"uuid"`
    Sequence string 
    Action string 

	} `json:"match"`
}


type RouteMapMatchAsPath struct {
    Name string `json:"name"`
}


type RouteMapMatchCommunity struct {
    NameCfg RouteMapMatchCommunityNameCfg `json:"name-cfg"`
}


type RouteMapMatchCommunityNameCfg struct {
    Name string `json:"name"`
    ExactMatch int `json:"exact-match"`
}


type RouteMapMatchExtcommunity struct {
    ExtcommunityLName RouteMapMatchExtcommunityExtcommunityLName `json:"extcommunity-l-name"`
}


type RouteMapMatchExtcommunityExtcommunityLName struct {
    Name string `json:"name"`
    ExactMatch int `json:"exact-match"`
}


type RouteMapMatchGroup struct {
    GroupId int `json:"group-id"`
    HaState string `json:"ha-state"`
}


type RouteMapMatchInterface struct {
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
}


type RouteMapMatchIp struct {
    Address RouteMapMatchIpAddress `json:"address"`
    NextHop RouteMapMatchIpNextHop `json:"next-hop"`
    Peer RouteMapMatchIpPeer `json:"peer"`
    Rib RouteMapMatchIpRib `json:"rib"`
}


type RouteMapMatchIpAddress struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
    PrefixList RouteMapMatchIpAddressPrefixList `json:"prefix-list"`
}


type RouteMapMatchIpAddressPrefixList struct {
    Name string `json:"name"`
}


type RouteMapMatchIpNextHop struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
    PrefixList1 RouteMapMatchIpNextHopPrefixList1 `json:"prefix-list-1"`
}


type RouteMapMatchIpNextHopPrefixList1 struct {
    Name string `json:"name"`
}


type RouteMapMatchIpPeer struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
}


type RouteMapMatchIpRib struct {
    Exact string `json:"exact"`
    Reachable string `json:"reachable"`
    Unreachable string `json:"unreachable"`
}


type RouteMapMatchIpv6 struct {
    Address1 RouteMapMatchIpv6Address1 `json:"address-1"`
    NextHop1 RouteMapMatchIpv6NextHop1 `json:"next-hop-1"`
    Peer1 RouteMapMatchIpv6Peer1 `json:"peer-1"`
    Rib RouteMapMatchIpv6Rib `json:"rib"`
}


type RouteMapMatchIpv6Address1 struct {
    Name string `json:"name"`
    PrefixList2 RouteMapMatchIpv6Address1PrefixList2 `json:"prefix-list-2"`
}


type RouteMapMatchIpv6Address1PrefixList2 struct {
    Name string `json:"name"`
}


type RouteMapMatchIpv6NextHop1 struct {
    NextHopAclName string `json:"next-hop-acl-name"`
    V6Addr string `json:"v6-addr"`
    PrefixListName string `json:"prefix-list-name"`
}


type RouteMapMatchIpv6Peer1 struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
}


type RouteMapMatchIpv6Rib struct {
    Exact string `json:"exact"`
    Reachable string `json:"reachable"`
    Unreachable string `json:"unreachable"`
}


type RouteMapMatchLargeCommunity struct {
    LNameCfg RouteMapMatchLargeCommunityLNameCfg `json:"l-name-cfg"`
}


type RouteMapMatchLargeCommunityLNameCfg struct {
    Name string `json:"name"`
    ExactMatch int `json:"exact-match"`
}


type RouteMapMatchLocalPreference struct {
    Val int `json:"val"`
}


type RouteMapMatchMetric struct {
    Value int `json:"value"`
}


type RouteMapMatchOrigin struct {
    Egp int `json:"egp"`
    Igp int `json:"igp"`
    Incomplete int `json:"incomplete"`
}


type RouteMapMatchRouteType struct {
    External RouteMapMatchRouteTypeExternal `json:"external"`
}


type RouteMapMatchRouteTypeExternal struct {
    Value string `json:"value"`
}


type RouteMapMatchScaleout struct {
    ClusterId int `json:"cluster-id"`
    OperationalState string `json:"operational-state"`
}


type RouteMapMatchTag struct {
    Value int `json:"value"`
}

func (p *RouteMapMatch) GetId() string{
    return "1"
}

func (p *RouteMapMatch) getPath() string{
    return "route-map/"+strconv.Itoa(p.Inst.Tag.Value)+"+" +p.Inst.Action + "+" +p.Inst.Sequence + "/match"
}

func (p *RouteMapMatch) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapMatch::Post")
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

func (p *RouteMapMatch) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapMatch::Get")
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
func (p *RouteMapMatch) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapMatch::Put")
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

func (p *RouteMapMatch) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMapMatch::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
