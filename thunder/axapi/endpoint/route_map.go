package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type RouteMap struct {
	Inst struct {
		Action string `json:"action"`

		Match RouteMapMatch1181 `json:"match"`

		Sequence int `json:"sequence"`

		Set RouteMapSet1211 `json:"set"`

		Tag string `json:"tag"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"route-map"`
}

type RouteMapMatch1181 struct {
	AsPath          RouteMapMatchAsPath1182          `json:"as-path"`
	Community       RouteMapMatchCommunity1183       `json:"community"`
	Extcommunity    RouteMapMatchExtcommunity1185    `json:"extcommunity"`
	LargeCommunity  RouteMapMatchLargeCommunity1187  `json:"large-community"`
	Group           RouteMapMatchGroup1189           `json:"group"`
	Scaleout        RouteMapMatchScaleout1190        `json:"scaleout"`
	Interface       RouteMapMatchInterface1191       `json:"interface"`
	LocalPreference RouteMapMatchLocalPreference1192 `json:"local-preference"`
	Origin          RouteMapMatchOrigin1193          `json:"origin"`
	Ip              RouteMapMatchIp1194              `json:"ip"`
	Ipv6            RouteMapMatchIpv61201            `json:"ipv6"`
	Metric          RouteMapMatchMetric1207          `json:"metric"`
	RouteType       RouteMapMatchRouteType1208       `json:"route-type"`
	Tag             RouteMapMatchTag1210             `json:"tag"`
	Uuid            string                           `json:"uuid"`
}

type RouteMapMatchAsPath1182 struct {
	Name string `json:"name"`
}

type RouteMapMatchCommunity1183 struct {
	NameCfg RouteMapMatchCommunityNameCfg1184 `json:"name-cfg"`
}

type RouteMapMatchCommunityNameCfg1184 struct {
	Name       string `json:"name"`
	ExactMatch int    `json:"exact-match"`
}

type RouteMapMatchExtcommunity1185 struct {
	ExtcommunityLName RouteMapMatchExtcommunityExtcommunityLName1186 `json:"extcommunity-l-name"`
}

type RouteMapMatchExtcommunityExtcommunityLName1186 struct {
	Name       string `json:"name"`
	ExactMatch int    `json:"exact-match"`
}

type RouteMapMatchLargeCommunity1187 struct {
	LNameCfg RouteMapMatchLargeCommunityLNameCfg1188 `json:"l-name-cfg"`
}

type RouteMapMatchLargeCommunityLNameCfg1188 struct {
	Name       string `json:"name"`
	ExactMatch int    `json:"exact-match"`
}

type RouteMapMatchGroup1189 struct {
	GroupId int    `json:"group-id"`
	HaState string `json:"ha-state"`
}

type RouteMapMatchScaleout1190 struct {
	ClusterId        int    `json:"cluster-id"`
	OperationalState string `json:"operational-state"`
	Ipv4             string `json:"ipv4"`
	Ipv6             string `json:"ipv6"`
}

type RouteMapMatchInterface1191 struct {
	Ethernet int `json:"ethernet"`
	Loopback int `json:"loopback"`
	Trunk    int `json:"trunk"`
	Ve       int `json:"ve"`
	Tunnel   int `json:"tunnel"`
}

type RouteMapMatchLocalPreference1192 struct {
	Val int `json:"val"`
}

type RouteMapMatchOrigin1193 struct {
	Egp        int `json:"egp"`
	Igp        int `json:"igp"`
	Incomplete int `json:"incomplete"`
}

type RouteMapMatchIp1194 struct {
	Address RouteMapMatchIpAddress1195 `json:"address"`
	NextHop RouteMapMatchIpNextHop1197 `json:"next-hop"`
	Peer    RouteMapMatchIpPeer1199    `json:"peer"`
	Rib     RouteMapMatchIpRib1200     `json:"rib"`
}

type RouteMapMatchIpAddress1195 struct {
	Acl1       int                                  `json:"acl1"`
	Acl2       int                                  `json:"acl2"`
	Name       string                               `json:"name"`
	PrefixList RouteMapMatchIpAddressPrefixList1196 `json:"prefix-list"`
}

type RouteMapMatchIpAddressPrefixList1196 struct {
	Name string `json:"name"`
}

type RouteMapMatchIpNextHop1197 struct {
	Acl1        int                                   `json:"acl1"`
	Acl2        int                                   `json:"acl2"`
	Name        string                                `json:"name"`
	PrefixList1 RouteMapMatchIpNextHopPrefixList11198 `json:"prefix-list-1"`
}

type RouteMapMatchIpNextHopPrefixList11198 struct {
	Name string `json:"name"`
}

type RouteMapMatchIpPeer1199 struct {
	Acl1 int    `json:"acl1"`
	Acl2 int    `json:"acl2"`
	Name string `json:"name"`
}

type RouteMapMatchIpRib1200 struct {
	Exact       string `json:"exact"`
	Reachable   string `json:"reachable"`
	Unreachable string `json:"unreachable"`
}

type RouteMapMatchIpv61201 struct {
	Address1 RouteMapMatchIpv6Address11202 `json:"address-1"`
	NextHop1 RouteMapMatchIpv6NextHop11204 `json:"next-hop-1"`
	Peer1    RouteMapMatchIpv6Peer11205    `json:"peer-1"`
	Rib      RouteMapMatchIpv6Rib1206      `json:"rib"`
}

type RouteMapMatchIpv6Address11202 struct {
	Name        string                                   `json:"name"`
	PrefixList2 RouteMapMatchIpv6Address1PrefixList21203 `json:"prefix-list-2"`
}

type RouteMapMatchIpv6Address1PrefixList21203 struct {
	Name string `json:"name"`
}

type RouteMapMatchIpv6NextHop11204 struct {
	NextHopAclName string `json:"next-hop-acl-name"`
	V6Addr         string `json:"v6-addr"`
	PrefixListName string `json:"prefix-list-name"`
}

type RouteMapMatchIpv6Peer11205 struct {
	Acl1 int    `json:"acl1"`
	Acl2 int    `json:"acl2"`
	Name string `json:"name"`
}

type RouteMapMatchIpv6Rib1206 struct {
	Exact       string `json:"exact"`
	Reachable   string `json:"reachable"`
	Unreachable string `json:"unreachable"`
}

type RouteMapMatchMetric1207 struct {
	Value int `json:"value"`
}

type RouteMapMatchRouteType1208 struct {
	External RouteMapMatchRouteTypeExternal1209 `json:"external"`
}

type RouteMapMatchRouteTypeExternal1209 struct {
	Value string `json:"value"`
}

type RouteMapMatchTag1210 struct {
	Value int `json:"value"`
}

type RouteMapSet1211 struct {
	Ip              RouteMapSetIp1212              `json:"ip"`
	Ddos            RouteMapSetDdos1214            `json:"ddos"`
	Ipv6            RouteMapSetIpv61215            `json:"ipv6"`
	Level           RouteMapSetLevel1218           `json:"level"`
	Metric          RouteMapSetMetric1219          `json:"metric"`
	MetricType      RouteMapSetMetricType1220      `json:"metric-type"`
	Tag             RouteMapSetTag1221             `json:"tag"`
	Aggregator      RouteMapSetAggregator1222      `json:"aggregator"`
	AsPath          RouteMapSetAsPath1224          `json:"as-path"`
	AtomicAggregate int                            `json:"atomic-aggregate"`
	CommList        RouteMapSetCommList1225        `json:"comm-list"`
	Community       string                         `json:"community"`
	DampeningCfg    RouteMapSetDampeningCfg1226    `json:"dampening-cfg"`
	Extcommunity    RouteMapSetExtcommunity1227    `json:"extcommunity"`
	LocalPreference RouteMapSetLocalPreference1230 `json:"local-preference"`
	OriginatorId    RouteMapSetOriginatorId1231    `json:"originator-id"`
	Weight          RouteMapSetWeight1232          `json:"weight"`
	Origin          RouteMapSetOrigin1233          `json:"origin"`
	LargeCommList   RouteMapSetLargeCommList1234   `json:"large-comm-list"`
	LargeCommunity  string                         `json:"large-community"`
	Uuid            string                         `json:"uuid"`
}

type RouteMapSetIp1212 struct {
	NextHop RouteMapSetIpNextHop1213 `json:"next-hop"`
}

type RouteMapSetIpNextHop1213 struct {
	Address string `json:"address"`
}

type RouteMapSetDdos1214 struct {
	ClassListName string `json:"class-list-name"`
	ClassListCid  int    `json:"class-list-cid"`
	Zone          string `json:"zone"`
}

type RouteMapSetIpv61215 struct {
	NextHop1 RouteMapSetIpv6NextHop11216 `json:"next-hop-1"`
}

type RouteMapSetIpv6NextHop11216 struct {
	Address string                           `json:"address"`
	Local   RouteMapSetIpv6NextHop1Local1217 `json:"local"`
}

type RouteMapSetIpv6NextHop1Local1217 struct {
	Address string `json:"address"`
}

type RouteMapSetLevel1218 struct {
	Value string `json:"value"`
}

type RouteMapSetMetric1219 struct {
	Value string `json:"value"`
}

type RouteMapSetMetricType1220 struct {
	Value string `json:"value"`
}

type RouteMapSetTag1221 struct {
	Value int `json:"value"`
}

type RouteMapSetAggregator1222 struct {
	AggregatorAs RouteMapSetAggregatorAggregatorAs1223 `json:"aggregator-as"`
}

type RouteMapSetAggregatorAggregatorAs1223 struct {
	Asn int    `json:"asn"`
	Ip  string `json:"ip"`
}

type RouteMapSetAsPath1224 struct {
	Prepend string `json:"prepend"`
	Num     string `json:"num"`
	Num2    string `json:"num2"`
}

type RouteMapSetCommList1225 struct {
	VStd       int    `json:"v-std"`
	Delete     int    `json:"delete"`
	VExp       int    `json:"v-exp"`
	VExpDelete int    `json:"v-exp-delete"`
	Name       string `json:"name"`
	NameDelete int    `json:"name-delete"`
}

type RouteMapSetDampeningCfg1226 struct {
	Dampening           int `json:"dampening"`
	DampeningHalfTime   int `json:"dampening-half-time"`
	DampeningReuse      int `json:"dampening-reuse"`
	DampeningSupress    int `json:"dampening-supress"`
	DampeningMaxSupress int `json:"dampening-max-supress"`
	DampeningPenalty    int `json:"dampening-penalty"`
}

type RouteMapSetExtcommunity1227 struct {
	Rt  RouteMapSetExtcommunityRt1228  `json:"rt"`
	Soo RouteMapSetExtcommunitySoo1229 `json:"soo"`
}

type RouteMapSetExtcommunityRt1228 struct {
	Value string `json:"value"`
}

type RouteMapSetExtcommunitySoo1229 struct {
	Value string `json:"value"`
}

type RouteMapSetLocalPreference1230 struct {
	Val int `json:"val"`
}

type RouteMapSetOriginatorId1231 struct {
	OriginatorIp string `json:"originator-ip"`
}

type RouteMapSetWeight1232 struct {
	WeightVal int `json:"weight-val"`
}

type RouteMapSetOrigin1233 struct {
	Egp        int `json:"egp"`
	Igp        int `json:"igp"`
	Incomplete int `json:"incomplete"`
}

type RouteMapSetLargeCommList1234 struct {
	LVStd           int    `json:"l-v-std"`
	LVStdDelete     int    `json:"l-v-std-delete"`
	LVExp           int    `json:"l-v-exp"`
	LVExpDelete     int    `json:"l-v-exp-delete"`
	LName           string `json:"l-name"`
	LargeNameDelete int    `json:"large-name-delete"`
}

func (p *RouteMap) GetId() string {
	return p.Inst.Tag + "+" + p.Inst.Action + "+" + strconv.Itoa(p.Inst.Sequence)
}

func (p *RouteMap) getPath() string {
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
		if len(axResp) > 0 {
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
