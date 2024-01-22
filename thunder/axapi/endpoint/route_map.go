

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type RouteMap struct {
	Inst struct {

    Action string `json:"action"`
    Match RouteMapMatch1097 `json:"match"`
    Sequence int `json:"sequence"`
    Set RouteMapSet1127 `json:"set"`
    Tag string `json:"tag"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"route-map"`
}


type RouteMapMatch1097 struct {
    AsPath RouteMapMatchAsPath1098 `json:"as-path"`
    Community RouteMapMatchCommunity1099 `json:"community"`
    Extcommunity RouteMapMatchExtcommunity1101 `json:"extcommunity"`
    LargeCommunity RouteMapMatchLargeCommunity1103 `json:"large-community"`
    Group RouteMapMatchGroup1105 `json:"group"`
    Scaleout RouteMapMatchScaleout1106 `json:"scaleout"`
    Interface RouteMapMatchInterface1107 `json:"interface"`
    LocalPreference RouteMapMatchLocalPreference1108 `json:"local-preference"`
    Origin RouteMapMatchOrigin1109 `json:"origin"`
    Ip RouteMapMatchIp1110 `json:"ip"`
    Ipv6 RouteMapMatchIpv61117 `json:"ipv6"`
    Metric RouteMapMatchMetric1123 `json:"metric"`
    RouteType RouteMapMatchRouteType1124 `json:"route-type"`
    Tag RouteMapMatchTag1126 `json:"tag"`
    Uuid string `json:"uuid"`
}


type RouteMapMatchAsPath1098 struct {
    Name string `json:"name"`
}


type RouteMapMatchCommunity1099 struct {
    NameCfg RouteMapMatchCommunityNameCfg1100 `json:"name-cfg"`
}


type RouteMapMatchCommunityNameCfg1100 struct {
    Name string `json:"name"`
    ExactMatch int `json:"exact-match"`
}


type RouteMapMatchExtcommunity1101 struct {
    ExtcommunityLName RouteMapMatchExtcommunityExtcommunityLName1102 `json:"extcommunity-l-name"`
}


type RouteMapMatchExtcommunityExtcommunityLName1102 struct {
    Name string `json:"name"`
    ExactMatch int `json:"exact-match"`
}


type RouteMapMatchLargeCommunity1103 struct {
    LNameCfg RouteMapMatchLargeCommunityLNameCfg1104 `json:"l-name-cfg"`
}


type RouteMapMatchLargeCommunityLNameCfg1104 struct {
    Name string `json:"name"`
    ExactMatch int `json:"exact-match"`
}


type RouteMapMatchGroup1105 struct {
    GroupId int `json:"group-id"`
    HaState string `json:"ha-state"`
}


type RouteMapMatchScaleout1106 struct {
    ClusterId int `json:"cluster-id"`
    OperationalState string `json:"operational-state"`
}


type RouteMapMatchInterface1107 struct {
    Ethernet int `json:"ethernet"`
    Loopback int `json:"loopback"`
    Trunk int `json:"trunk"`
    Ve int `json:"ve"`
    Tunnel int `json:"tunnel"`
}


type RouteMapMatchLocalPreference1108 struct {
    Val int `json:"val"`
}


type RouteMapMatchOrigin1109 struct {
    Egp int `json:"egp"`
    Igp int `json:"igp"`
    Incomplete int `json:"incomplete"`
}


type RouteMapMatchIp1110 struct {
    Address RouteMapMatchIpAddress1111 `json:"address"`
    NextHop RouteMapMatchIpNextHop1113 `json:"next-hop"`
    Peer RouteMapMatchIpPeer1115 `json:"peer"`
    Rib RouteMapMatchIpRib1116 `json:"rib"`
}


type RouteMapMatchIpAddress1111 struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
    PrefixList RouteMapMatchIpAddressPrefixList1112 `json:"prefix-list"`
}


type RouteMapMatchIpAddressPrefixList1112 struct {
    Name string `json:"name"`
}


type RouteMapMatchIpNextHop1113 struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
    PrefixList1 RouteMapMatchIpNextHopPrefixList11114 `json:"prefix-list-1"`
}


type RouteMapMatchIpNextHopPrefixList11114 struct {
    Name string `json:"name"`
}


type RouteMapMatchIpPeer1115 struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
}


type RouteMapMatchIpRib1116 struct {
    Exact string `json:"exact"`
    Reachable string `json:"reachable"`
    Unreachable string `json:"unreachable"`
}


type RouteMapMatchIpv61117 struct {
    Address1 RouteMapMatchIpv6Address11118 `json:"address-1"`
    NextHop1 RouteMapMatchIpv6NextHop11120 `json:"next-hop-1"`
    Peer1 RouteMapMatchIpv6Peer11121 `json:"peer-1"`
    Rib RouteMapMatchIpv6Rib1122 `json:"rib"`
}


type RouteMapMatchIpv6Address11118 struct {
    Name string `json:"name"`
    PrefixList2 RouteMapMatchIpv6Address1PrefixList21119 `json:"prefix-list-2"`
}


type RouteMapMatchIpv6Address1PrefixList21119 struct {
    Name string `json:"name"`
}


type RouteMapMatchIpv6NextHop11120 struct {
    NextHopAclName string `json:"next-hop-acl-name"`
    V6Addr string `json:"v6-addr"`
    PrefixListName string `json:"prefix-list-name"`
}


type RouteMapMatchIpv6Peer11121 struct {
    Acl1 int `json:"acl1"`
    Acl2 int `json:"acl2"`
    Name string `json:"name"`
}


type RouteMapMatchIpv6Rib1122 struct {
    Exact string `json:"exact"`
    Reachable string `json:"reachable"`
    Unreachable string `json:"unreachable"`
}


type RouteMapMatchMetric1123 struct {
    Value int `json:"value"`
}


type RouteMapMatchRouteType1124 struct {
    External RouteMapMatchRouteTypeExternal1125 `json:"external"`
}


type RouteMapMatchRouteTypeExternal1125 struct {
    Value string `json:"value"`
}


type RouteMapMatchTag1126 struct {
    Value int `json:"value"`
}


type RouteMapSet1127 struct {
    Ip RouteMapSetIp1128 `json:"ip"`
    Ddos RouteMapSetDdos1130 `json:"ddos"`
    Ipv6 RouteMapSetIpv61131 `json:"ipv6"`
    Level RouteMapSetLevel1134 `json:"level"`
    Metric RouteMapSetMetric1135 `json:"metric"`
    MetricType RouteMapSetMetricType1136 `json:"metric-type"`
    Tag RouteMapSetTag1137 `json:"tag"`
    Aggregator RouteMapSetAggregator1138 `json:"aggregator"`
    AsPath RouteMapSetAsPath1140 `json:"as-path"`
    AtomicAggregate int `json:"atomic-aggregate"`
    CommList RouteMapSetCommList1141 `json:"comm-list"`
    Community string `json:"community"`
    DampeningCfg RouteMapSetDampeningCfg1142 `json:"dampening-cfg"`
    Extcommunity RouteMapSetExtcommunity1143 `json:"extcommunity"`
    LocalPreference RouteMapSetLocalPreference1146 `json:"local-preference"`
    OriginatorId RouteMapSetOriginatorId1147 `json:"originator-id"`
    Weight RouteMapSetWeight1148 `json:"weight"`
    Origin RouteMapSetOrigin1149 `json:"origin"`
    LargeCommList RouteMapSetLargeCommList1150 `json:"large-comm-list"`
    LargeCommunity string `json:"large-community"`
    Uuid string `json:"uuid"`
}


type RouteMapSetIp1128 struct {
    NextHop RouteMapSetIpNextHop1129 `json:"next-hop"`
}


type RouteMapSetIpNextHop1129 struct {
    Address string `json:"address"`
}


type RouteMapSetDdos1130 struct {
    ClassListName string `json:"class-list-name"`
    ClassListCid int `json:"class-list-cid"`
    Zone string `json:"zone"`
}


type RouteMapSetIpv61131 struct {
    NextHop1 RouteMapSetIpv6NextHop11132 `json:"next-hop-1"`
}


type RouteMapSetIpv6NextHop11132 struct {
    Address string `json:"address"`
    Local RouteMapSetIpv6NextHop1Local1133 `json:"local"`
}


type RouteMapSetIpv6NextHop1Local1133 struct {
    Address string `json:"address"`
}


type RouteMapSetLevel1134 struct {
    Value string `json:"value"`
}


type RouteMapSetMetric1135 struct {
    Value string `json:"value"`
}


type RouteMapSetMetricType1136 struct {
    Value string `json:"value"`
}


type RouteMapSetTag1137 struct {
    Value int `json:"value"`
}


type RouteMapSetAggregator1138 struct {
    AggregatorAs RouteMapSetAggregatorAggregatorAs1139 `json:"aggregator-as"`
}


type RouteMapSetAggregatorAggregatorAs1139 struct {
    Asn int `json:"asn"`
    Ip string `json:"ip"`
}


type RouteMapSetAsPath1140 struct {
    Prepend string `json:"prepend"`
    Num string `json:"num"`
    Num2 string `json:"num2"`
}


type RouteMapSetCommList1141 struct {
    VStd int `json:"v-std"`
    Delete int `json:"delete"`
    VExp int `json:"v-exp"`
    VExpDelete int `json:"v-exp-delete"`
    Name string `json:"name"`
    NameDelete int `json:"name-delete"`
}


type RouteMapSetDampeningCfg1142 struct {
    Dampening int `json:"dampening"`
    DampeningHalfTime int `json:"dampening-half-time"`
    DampeningReuse int `json:"dampening-reuse"`
    DampeningSupress int `json:"dampening-supress"`
    DampeningMaxSupress int `json:"dampening-max-supress"`
    DampeningPenalty int `json:"dampening-penalty"`
}


type RouteMapSetExtcommunity1143 struct {
    Rt RouteMapSetExtcommunityRt1144 `json:"rt"`
    Soo RouteMapSetExtcommunitySoo1145 `json:"soo"`
}


type RouteMapSetExtcommunityRt1144 struct {
    Value string `json:"value"`
}


type RouteMapSetExtcommunitySoo1145 struct {
    Value string `json:"value"`
}


type RouteMapSetLocalPreference1146 struct {
    Val int `json:"val"`
}


type RouteMapSetOriginatorId1147 struct {
    OriginatorIp string `json:"originator-ip"`
}


type RouteMapSetWeight1148 struct {
    WeightVal int `json:"weight-val"`
}


type RouteMapSetOrigin1149 struct {
    Egp int `json:"egp"`
    Igp int `json:"igp"`
    Incomplete int `json:"incomplete"`
}


type RouteMapSetLargeCommList1150 struct {
    LVStd int `json:"l-v-std"`
    LVStdDelete int `json:"l-v-std-delete"`
    LVExp int `json:"l-v-exp"`
    LVExpDelete int `json:"l-v-exp-delete"`
    LName string `json:"l-name"`
    LargeNameDelete int `json:"large-name-delete"`
}

func (p *RouteMap) GetId() string{
    return p.Inst.Tag+"+"+p.Inst.Action+"+"+strconv.Itoa(p.Inst.Sequence)
}

func (p *RouteMap) getPath() string{
    return "route-map"
}

func (p *RouteMap) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMap::Post")
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

func (p *RouteMap) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMap::Get")
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
func (p *RouteMap) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMap::Put")
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

func (p *RouteMap) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouteMap::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
