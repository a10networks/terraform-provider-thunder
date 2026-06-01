package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 6_0_8-219
type RouteMap struct {
	Inst struct {
		Action string `json:"action"`

		Match RouteMapMatch1180 `json:"match"`

		Sequence int `json:"sequence"`

		Set RouteMapSet1210 `json:"set"`

		Tag string `json:"tag"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`
	} `json:"route-map"`
}

type RouteMapMatch1180 struct {
	AsPath          RouteMapMatchAsPath1181          `json:"as-path"`
	Community       RouteMapMatchCommunity1182       `json:"community"`
	Extcommunity    RouteMapMatchExtcommunity1184    `json:"extcommunity"`
	LargeCommunity  RouteMapMatchLargeCommunity1186  `json:"large-community"`
	Group           RouteMapMatchGroup1188           `json:"group"`
	Scaleout        RouteMapMatchScaleout1189        `json:"scaleout"`
	Interface       RouteMapMatchInterface1190       `json:"interface"`
	LocalPreference RouteMapMatchLocalPreference1191 `json:"local-preference"`
	Origin          RouteMapMatchOrigin1192          `json:"origin"`
	Ip              RouteMapMatchIp1193              `json:"ip"`
	Ipv6            RouteMapMatchIpv61200            `json:"ipv6"`
	Metric          RouteMapMatchMetric1206          `json:"metric"`
	RouteType       RouteMapMatchRouteType1207       `json:"route-type"`
	Tag             RouteMapMatchTag1209             `json:"tag"`
	Uuid            string                           `json:"uuid"`
}

type RouteMapMatchAsPath1181 struct {
	Name string `json:"name"`
}

type RouteMapMatchCommunity1182 struct {
	NameCfg RouteMapMatchCommunityNameCfg1183 `json:"name-cfg"`
}

type RouteMapMatchCommunityNameCfg1183 struct {
	Name       string `json:"name"`
	ExactMatch int    `json:"exact-match"`
}

type RouteMapMatchExtcommunity1184 struct {
	ExtcommunityLName RouteMapMatchExtcommunityExtcommunityLName1185 `json:"extcommunity-l-name"`
}

type RouteMapMatchExtcommunityExtcommunityLName1185 struct {
	Name       string `json:"name"`
	ExactMatch int    `json:"exact-match"`
}

type RouteMapMatchLargeCommunity1186 struct {
	LNameCfg RouteMapMatchLargeCommunityLNameCfg1187 `json:"l-name-cfg"`
}

type RouteMapMatchLargeCommunityLNameCfg1187 struct {
	Name       string `json:"name"`
	ExactMatch int    `json:"exact-match"`
}

type RouteMapMatchGroup1188 struct {
	GroupId int    `json:"group-id"`
	HaState string `json:"ha-state"`
}

type RouteMapMatchScaleout1189 struct {
	ClusterId        int    `json:"cluster-id"`
	OperationalState string `json:"operational-state"`
	Ipv4             string `json:"ipv4"`
	Ipv6             string `json:"ipv6"`
}

type RouteMapMatchInterface1190 struct {
	Ethernet int `json:"ethernet"`
	Loopback int `json:"loopback"`
	Trunk    int `json:"trunk"`
	Ve       int `json:"ve"`
	Tunnel   int `json:"tunnel"`
}

type RouteMapMatchLocalPreference1191 struct {
	Val int `json:"val"`
}

type RouteMapMatchOrigin1192 struct {
	Egp        int `json:"egp"`
	Igp        int `json:"igp"`
	Incomplete int `json:"incomplete"`
}

type RouteMapMatchIp1193 struct {
	Address RouteMapMatchIpAddress1194 `json:"address"`
	NextHop RouteMapMatchIpNextHop1196 `json:"next-hop"`
	Peer    RouteMapMatchIpPeer1198    `json:"peer"`
	Rib     RouteMapMatchIpRib1199     `json:"rib"`
}

type RouteMapMatchIpAddress1194 struct {
	Acl1       int                                  `json:"acl1"`
	Acl2       int                                  `json:"acl2"`
	Name       string                               `json:"name"`
	PrefixList RouteMapMatchIpAddressPrefixList1195 `json:"prefix-list"`
}

type RouteMapMatchIpAddressPrefixList1195 struct {
	Name string `json:"name"`
}

type RouteMapMatchIpNextHop1196 struct {
	Acl1        int                                   `json:"acl1"`
	Acl2        int                                   `json:"acl2"`
	Name        string                                `json:"name"`
	PrefixList1 RouteMapMatchIpNextHopPrefixList11197 `json:"prefix-list-1"`
}

type RouteMapMatchIpNextHopPrefixList11197 struct {
	Name string `json:"name"`
}

type RouteMapMatchIpPeer1198 struct {
	Acl1 int    `json:"acl1"`
	Acl2 int    `json:"acl2"`
	Name string `json:"name"`
}

type RouteMapMatchIpRib1199 struct {
	Exact       string `json:"exact"`
	Reachable   string `json:"reachable"`
	Unreachable string `json:"unreachable"`
}

type RouteMapMatchIpv61200 struct {
	Address1 RouteMapMatchIpv6Address11201 `json:"address-1"`
	NextHop1 RouteMapMatchIpv6NextHop11203 `json:"next-hop-1"`
	Peer1    RouteMapMatchIpv6Peer11204    `json:"peer-1"`
	Rib      RouteMapMatchIpv6Rib1205      `json:"rib"`
}

type RouteMapMatchIpv6Address11201 struct {
	Name        string                                   `json:"name"`
	PrefixList2 RouteMapMatchIpv6Address1PrefixList21202 `json:"prefix-list-2"`
}

type RouteMapMatchIpv6Address1PrefixList21202 struct {
	Name string `json:"name"`
}

type RouteMapMatchIpv6NextHop11203 struct {
	NextHopAclName string `json:"next-hop-acl-name"`
	V6Addr         string `json:"v6-addr"`
	PrefixListName string `json:"prefix-list-name"`
}

type RouteMapMatchIpv6Peer11204 struct {
	Acl1 int    `json:"acl1"`
	Acl2 int    `json:"acl2"`
	Name string `json:"name"`
}

type RouteMapMatchIpv6Rib1205 struct {
	Exact       string `json:"exact"`
	Reachable   string `json:"reachable"`
	Unreachable string `json:"unreachable"`
}

type RouteMapMatchMetric1206 struct {
	Value int `json:"value"`
}

type RouteMapMatchRouteType1207 struct {
	External RouteMapMatchRouteTypeExternal1208 `json:"external"`
}

type RouteMapMatchRouteTypeExternal1208 struct {
	Value string `json:"value"`
}

type RouteMapMatchTag1209 struct {
	Value int `json:"value"`
}

type RouteMapSet1210 struct {
	Ip              RouteMapSetIp1211              `json:"ip"`
	Ddos            RouteMapSetDdos1213            `json:"ddos"`
	Ipv6            RouteMapSetIpv61214            `json:"ipv6"`
	Level           RouteMapSetLevel1217           `json:"level"`
	Metric          RouteMapSetMetric1218          `json:"metric"`
	MetricType      RouteMapSetMetricType1219      `json:"metric-type"`
	Tag             RouteMapSetTag1220             `json:"tag"`
	Aggregator      RouteMapSetAggregator1221      `json:"aggregator"`
	AsPath          RouteMapSetAsPath1223          `json:"as-path"`
	AtomicAggregate int                            `json:"atomic-aggregate"`
	CommList        RouteMapSetCommList1224        `json:"comm-list"`
	Community       string                         `json:"community"`
	DampeningCfg    RouteMapSetDampeningCfg1225    `json:"dampening-cfg"`
	Extcommunity    RouteMapSetExtcommunity1226    `json:"extcommunity"`
	LocalPreference RouteMapSetLocalPreference1229 `json:"local-preference"`
	OriginatorId    RouteMapSetOriginatorId1230    `json:"originator-id"`
	Weight          RouteMapSetWeight1231          `json:"weight"`
	Origin          RouteMapSetOrigin1232          `json:"origin"`
	LargeCommList   RouteMapSetLargeCommList1233   `json:"large-comm-list"`
	LargeCommunity  string                         `json:"large-community"`
	Uuid            string                         `json:"uuid"`
}

type RouteMapSetIp1211 struct {
	NextHop RouteMapSetIpNextHop1212 `json:"next-hop"`
}

type RouteMapSetIpNextHop1212 struct {
	Address string `json:"address"`
}

type RouteMapSetDdos1213 struct {
	ClassListName string `json:"class-list-name"`
	ClassListCid  int    `json:"class-list-cid"`
	Zone          string `json:"zone"`
}

type RouteMapSetIpv61214 struct {
	NextHop1 RouteMapSetIpv6NextHop11215 `json:"next-hop-1"`
}

type RouteMapSetIpv6NextHop11215 struct {
	Address string                           `json:"address"`
	Local   RouteMapSetIpv6NextHop1Local1216 `json:"local"`
}

type RouteMapSetIpv6NextHop1Local1216 struct {
	Address string `json:"address"`
}

type RouteMapSetLevel1217 struct {
	Value string `json:"value"`
}

type RouteMapSetMetric1218 struct {
	Value string `json:"value"`
}

type RouteMapSetMetricType1219 struct {
	Value string `json:"value"`
}

type RouteMapSetTag1220 struct {
	Value int `json:"value"`
}

type RouteMapSetAggregator1221 struct {
	AggregatorAs RouteMapSetAggregatorAggregatorAs1222 `json:"aggregator-as"`
}

type RouteMapSetAggregatorAggregatorAs1222 struct {
	Asn int    `json:"asn"`
	Ip  string `json:"ip"`
}

type RouteMapSetAsPath1223 struct {
	Prepend string `json:"prepend"`
	Num     string `json:"num"`
	Num2    string `json:"num2"`
}

type RouteMapSetCommList1224 struct {
	VStd       int    `json:"v-std"`
	Delete     int    `json:"delete"`
	VExp       int    `json:"v-exp"`
	VExpDelete int    `json:"v-exp-delete"`
	Name       string `json:"name"`
	NameDelete int    `json:"name-delete"`
}

type RouteMapSetDampeningCfg1225 struct {
	Dampening           int `json:"dampening"`
	DampeningHalfTime   int `json:"dampening-half-time"`
	DampeningReuse      int `json:"dampening-reuse"`
	DampeningSupress    int `json:"dampening-supress"`
	DampeningMaxSupress int `json:"dampening-max-supress"`
	DampeningPenalty    int `json:"dampening-penalty"`
}

type RouteMapSetExtcommunity1226 struct {
	Rt  RouteMapSetExtcommunityRt1227  `json:"rt"`
	Soo RouteMapSetExtcommunitySoo1228 `json:"soo"`
}

type RouteMapSetExtcommunityRt1227 struct {
	Value string `json:"value"`
}

type RouteMapSetExtcommunitySoo1228 struct {
	Value string `json:"value"`
}

type RouteMapSetLocalPreference1229 struct {
	Val int `json:"val"`
}

type RouteMapSetOriginatorId1230 struct {
	OriginatorIp string `json:"originator-ip"`
}

type RouteMapSetWeight1231 struct {
	WeightVal int `json:"weight-val"`
}

type RouteMapSetOrigin1232 struct {
	Egp        int `json:"egp"`
	Igp        int `json:"igp"`
	Incomplete int `json:"incomplete"`
}

type RouteMapSetLargeCommList1233 struct {
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
