package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouteMap struct {
	Tag RouteMapInstance `json:"route-map,omitempty"`
}

type RouteMapInstance struct {
	Action       string        `json:"action,omitempty"`
	Name         RouteMapMatch `json:"match,omitempty"`
	Sequence     int           `json:"sequence,omitempty"`
	AggregatorAs RouteMapSet   `json:"set,omitempty"`
	Tag          string        `json:"tag,omitempty"`
	UUID         string        `json:"uuid,omitempty"`
	UserTag      string        `json:"user-tag,omitempty"`
}

type RouteMapMatch struct {
	Name              RouteMapAsPath1         `json:"as-path,omitempty"`
	NameCfg           RouteMapCommunity       `json:"community,omitempty"`
	ExtcommunityLName RouteMapExtcommunity    `json:"extcommunity,omitempty"`
	GroupID           RouteMapGroup           `json:"group,omitempty"`
	Address           RouteMapIp1             `json:"ip,omitempty"`
	Ethernet          RouteMapInterface       `json:"interface,omitempty"`
	Address1          RouteMapIpv6            `json:"ipv6,omitempty"`
	Val               RouteMapLocalPreference `json:"local-preference,omitempty"`
	Value             RouteMapMetricMap       `json:"metric,omitempty"`
	Egp               RouteMapOrigin          `json:"origin,omitempty"`
	ValueExternal     RouteMapRouteType       `json:"route-type,omitempty"`
	ClusterID         RouteMapScaleout        `json:"scaleout,omitempty"`
	ValueTag          RouteMapTag             `json:"tag,omitempty"`
	UUID              string                  `json:"uuid,omitempty"`
}

type RouteMapSet struct {
	AggregatorAs            RouteMapAggregator      `json:"aggregator,omitempty"`
	Prepend                 RouteMapAsPath          `json:"as-path,omitempty"`
	AtomicAggregate         int                     `json:"atomic-aggregate,omitempty"`
	VStd                    RouteMapCommList        `json:"comm-list,omitempty"`
	Community               string                  `json:"community,omitempty"`
	Dampening               RouteMapDampeningCfg    `json:"dampening-cfg,omitempty"`
	ClassListName           RouteMapDdos            `json:"ddos,omitempty"`
	ValueRt                 RouteMapExtcommunity1   `json:"extcommunity,omitempty"`
	Address                 RouteMapIp2             `json:"ip,omitempty"`
	RouteMapNextHop1Address RouteMapIpv62           `json:"ipv6,omitempty"`
	ValueLevel1             RouteMapLevel           `json:"level,omitempty"`
	Val                     RouteMapLocalPreference `json:"local-preference,omitempty"`
	ValueLevel2             RouteMapMetric          `json:"metric,omitempty"`
	ValueLevel3             RouteMapMetricType      `json:"metric-type,omitempty"`
	Egp                     RouteMapOrigin          `json:"origin,omitempty"`
	OriginatorIP            RouteMapOriginatorID    `json:"originator-id,omitempty"`
	ValueTag                RouteMapTag             `json:"tag,omitempty"`
	UUID                    string                  `json:"uuid,omitempty"`
	WeightVal               RouteMapWeight          `json:"weight,omitempty"`
}

type RouteMapAsPath1 struct {
	Name string `json:"name,omitempty"`
}

type RouteMapPrefixList struct {
	Name string `json:"name,omitempty"`
}

type RouteMapTag struct {
	ValueTag int `json:"value,omitempty"`
}

type RouteMapMetricType struct {
	ValueLevel3 string `json:"value,omitempty"`
}

type RouteMapCommunity struct {
	Name RouteMapNameCfg `json:"name-cfg,omitempty"`
}

type RouteMapExtcommunityLName struct {
	ExtcommunityLName string `json:"name,omitempty"`
	ExactMatch        int    `json:"exact-match,omitempty"`
}

type RouteMapExtcommunity struct {
	ExtcommunityLName RouteMapExtcommunityLName `json:"extcommunity-l-name,omitempty"`
}

type RouteMapGroup struct {
	GroupID int    `json:"group-id,omitempty"`
	HaState string `json:"ha-state,omitempty"`
}

type RouteMapIp1 struct {
	Acl1  RouteMapAddress   `json:"address,omitempty"`
	Acl2  RouteMapNextHopIp `json:"next-hop,omitempty"`
	Name  RouteMapPeer      `json:"peer,omitempty"`
	Exact RouteMapRib       `json:"rib,omitempty"`
}

type RouteMapInterface struct {
	Ethernet int `json:"ethernet,omitempty"`
	Loopback int `json:"loopback,omitempty"`
	Trunk    int `json:"trunk,omitempty"`
	Tunnel   int `json:"tunnel,omitempty"`
	Ve       int `json:"ve,omitempty"`
}

type RouteMapIpv6 struct {
	Name           RouteMapAddress1    `json:"address-1,omitempty"`
	NextHopAclName RouteMapNextHopIpv6 `json:"next-hop-1,omitempty"`
	Acl1           RouteMapPeer        `json:"peer-1,omitempty"`
	Exact          RouteMapRib         `json:"rib,omitempty"`
}

type RouteMapLocalPreference struct {
	Val int `json:"val,omitempty"`
}

type RouteMapExternal struct {
	ValueExternal string `json:"value,omitempty"`
}

type RouteMapMetricMap struct {
	Value int `json:"value,omitempty"`
}

type RouteMapMetric struct {
	Value       string `json:"value,omitempty"`
	ValueMetric string `json:"value,omitempty"`
	ValueLevel2 string `json:"value,omitempty"`
}

type RouteMapOrigin struct {
	Egp        int `json:"egp,omitempty"`
	Igp        int `json:"igp,omitempty"`
	Incomplete int `json:"incomplete,omitempty"`
}

type RouteMapRouteType struct {
	ValueExternal RouteMapExternal `json:"external,omitempty"`
}

type RouteMapScaleout struct {
	ClusterID        int    `json:"cluster-id,omitempty"`
	OperationalState string `json:"operational-state,omitempty"`
}

type RouteMapAggregator struct {
	Asn RouteMapAggregatorAs `json:"aggregator-as,omitempty"`
}

type RouteMapAsPath struct {
	Num     int    `json:"num,omitempty"`
	Num2    int    `json:"num2,omitempty"`
	Prepend string `json:"prepend,omitempty"`
}

type RouteMapCommList struct {
	Delete     int    `json:"delete,omitempty"`
	Name       string `json:"name,omitempty"`
	NameDelete int    `json:"name-delete,omitempty"`
	VExp       int    `json:"v-exp,omitempty"`
	VExpDelete int    `json:"v-exp-delete,omitempty"`
	VStd       int    `json:"v-std,omitempty"`
}

type RouteMapDampeningCfg struct {
	Dampening           int `json:"dampening,omitempty"`
	DampeningHalfTime   int `json:"dampening-half-time,omitempty"`
	DampeningMaxSupress int `json:"dampening-max-supress,omitempty"`
	DampeningPenalty    int `json:"dampening-penalty,omitempty"`
	DampeningReuse      int `json:"dampening-reuse,omitempty"`
	DampeningSupress    int `json:"dampening-supress,omitempty"`
}

type RouteMapDdos struct {
	ClassListCid  int    `json:"class-list-cid,omitempty"`
	ClassListName string `json:"class-list-name,omitempty"`
	Zone          string `json:"zone,omitempty"`
}

type RouteMapRt struct {
	ValueRt string `json:"value,omitempty"`
}

type RouteMapSoo struct {
	ValueSoo string `json:"value,omitempty"`
}

type RouteMapExtcommunity1 struct {
	ValueRt  RouteMapRt  `json:"rt,omitempty"`
	ValueSoo RouteMapSoo `json:"soo,omitempty"`
}

type RouteMapIp2 struct {
	Address RouteMapNextHopSetIp `json:"next-hop,omitempty"`
}

type RouteMapIpv62 struct {
	RouteMapNextHop1Address RouteMapNextHop1 `json:"next-hop-1,omitempty"`
}

type RouteMapLevel struct {
	ValueLevel1 string `json:"value,omitempty"`
	ValueLevel2 string `json:"value,omitempty"`
	ValueLevel3 string `json:"value,omitempty"`
}

type RouteMapOriginatorID struct {
	OriginatorIP string `json:"originator-ip,omitempty"`
}

type RouteMapWeight struct {
	WeightVal int `json:"weight-val,omitempty"`
}

type RouteMapNameCfg struct {
	ExactMatch int    `json:"exact-match,omitempty"`
	Name       string `json:"name,omitempty"`
}

type RouteMapAddress struct {
	Acl1                int                `json:"acl1,omitempty"`
	Acl2                int                `json:"acl2,omitempty"`
	RouteMapAddressName string             `json:"name,omitempty"`
	Name                RouteMapPrefixList `json:"prefix-list,omitempty"`
}

type RouteMapPrefixList1 struct {
	Name string `json:"name,omitempty"`
}

type RouteMapNextHopIp struct {
	Acl1                  int                 `json:"acl1,omitempty"`
	Acl2                  int                 `json:"acl2,omitempty"`
	RouteMapNextHopIpName string              `json:"name,omitempty"`
	Name                  RouteMapPrefixList1 `json:"prefix-list-1,omitempty"`
}

type RouteMapPeer struct {
	Acl1 int    `json:"acl1,omitempty"`
	Acl2 int    `json:"acl2,omitempty"`
	Name string `json:"name,omitempty"`
}

type RouteMapRib struct {
	Exact       string `json:"exact,omitempty"`
	Reachable   string `json:"reachable,omitempty"`
	Unreachable string `json:"unreachable,omitempty"`
}

type RouteMapPrefixList2 struct {
	Name string `json:"name,omitempty"`
}

type RouteMapPeer1 struct {
	Acl1 int    `json:"acl1,omitempty"`
	Acl2 int    `json:"acl2,omitempty"`
	Name string `json:"name,omitempty"`
}

type RouteMapAddress1 struct {
	RouteMapAddress1Name string              `json:"name,omitempty"`
	Name                 RouteMapPrefixList2 `json:"prefix-list-2,omitempty"`
}

type RouteMapNextHopIpv6 struct {
	NextHopAclName string `json:"next-hop-acl-name,omitempty"`
	PrefixListName string `json:"prefix-list-name,omitempty"`
	V6Addr         string `json:"v6-addr,omitempty"`
}

type RouteMapAggregatorAs struct {
	Asn int    `json:"asn,omitempty"`
	IP  string `json:"ip,omitempty"`
}

type RouteMapNextHopSetIp struct {
	Address string `json:"address,omitempty"`
}

type RouteMapNextHop1 struct {
	RouteMapNextHop1Address string        `json:"address,omitempty"`
	Address                 RouteMapLocal `json:"local,omitempty"`
}

type RouteMapLocal struct {
	Address string `json:"address,omitempty"`
}

func PostRouteMap(id string, inst RouteMap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouteMap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/route-map", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouteMap", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetRouteMap(id string, name1 string, name2 string, name3 string, host string) (*RouteMap, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouteMap")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/route-map/"+name1+"+"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouteMap", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouteMap(id string, name1 string, name2 string, name3 string, inst RouteMap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouteMap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/route-map/"+name1+"+"+name2+"+"+name3, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouteMap", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteRouteMap(id string, name1 string, name2 string, name3 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouteMap")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/route-map/"+name1+"+"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouteMap", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
