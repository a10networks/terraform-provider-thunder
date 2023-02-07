package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouteMap struct {
	RouteMapInstanceTag RouteMapInstance `json:"route-map,omitempty"`
}

type RouteMapInstance struct {
	RouteMapInstanceAction      string                `json:"action,omitempty"`
	RouteMapInstanceMatchAsPath RouteMapInstanceMatch `json:"match,omitempty"`
	RouteMapInstanceSequence    int                   `json:"sequence,omitempty"`
	RouteMapInstanceSetIP       RouteMapInstanceSet   `json:"set,omitempty"`
	RouteMapInstanceTag         string                `json:"tag,omitempty"`
	RouteMapInstanceUUID        string                `json:"uuid,omitempty"`
	RouteMapInstanceUserTag     string                `json:"user-tag,omitempty"`
}

type RouteMapInstanceMatch struct {
	RouteMapInstanceMatchAsPathName                    RouteMapInstanceMatchAsPath          `json:"as-path,omitempty"`
	RouteMapInstanceMatchCommunityNameCfg              RouteMapInstanceMatchCommunity       `json:"community,omitempty"`
	RouteMapInstanceMatchExtcommunityExtcommunityLName RouteMapInstanceMatchExtcommunity    `json:"extcommunity,omitempty"`
	RouteMapInstanceMatchGroupGroupID                  RouteMapInstanceMatchGroup           `json:"group,omitempty"`
	RouteMapInstanceMatchIPAddress                     RouteMapInstanceMatchIP              `json:"ip,omitempty"`
	RouteMapInstanceMatchInterfaceEthernet             RouteMapInstanceMatchInterface       `json:"interface,omitempty"`
	RouteMapInstanceMatchIpv6Address1                  RouteMapInstanceMatchIpv6            `json:"ipv6,omitempty"`
	RouteMapInstanceMatchLocalPreferenceVal            RouteMapInstanceMatchLocalPreference `json:"local-preference,omitempty"`
	RouteMapInstanceMatchMetricValue                   RouteMapInstanceMatchMetric          `json:"metric,omitempty"`
	RouteMapInstanceMatchOriginEgp                     RouteMapInstanceMatchOrigin          `json:"origin,omitempty"`
	RouteMapInstanceMatchRouteTypeExternal             RouteMapInstanceMatchRouteType       `json:"route-type,omitempty"`
	RouteMapInstanceMatchScaleoutClusterID             RouteMapInstanceMatchScaleout        `json:"scaleout,omitempty"`
	RouteMapInstanceMatchTagValue                      RouteMapInstanceMatchTag             `json:"tag,omitempty"`
	RouteMapInstanceMatchUUID                          string                               `json:"uuid,omitempty"`
}

type RouteMapInstanceSet struct {
	RouteMapInstanceSetAggregatorAggregatorAs   RouteMapInstanceSetAggregator      `json:"aggregator,omitempty"`
	RouteMapInstanceSetAsPathPrepend            RouteMapInstanceSetAsPath          `json:"as-path,omitempty"`
	RouteMapInstanceSetAtomicAggregate          int                                `json:"atomic-aggregate,omitempty"`
	RouteMapInstanceSetCommListVStd             RouteMapInstanceSetCommList        `json:"comm-list,omitempty"`
	RouteMapInstanceSetCommunity                string                             `json:"community,omitempty"`
	RouteMapInstanceSetDampeningCfgDampening    RouteMapInstanceSetDampeningCfg    `json:"dampening-cfg,omitempty"`
	RouteMapInstanceSetDdosClassListName        RouteMapInstanceSetDdos            `json:"ddos,omitempty"`
	RouteMapInstanceSetExtcommunityRt           RouteMapInstanceSetExtcommunity    `json:"extcommunity,omitempty"`
	RouteMapInstanceSetIPNextHop                RouteMapInstanceSetIP              `json:"ip,omitempty"`
	RouteMapInstanceSetIpv6NextHop1             RouteMapInstanceSetIpv6            `json:"ipv6,omitempty"`
	RouteMapInstanceSetLevelValue               RouteMapInstanceSetLevel           `json:"level,omitempty"`
	RouteMapInstanceSetLocalPreferenceVal       RouteMapInstanceSetLocalPreference `json:"local-preference,omitempty"`
	RouteMapInstanceSetMetricValue              RouteMapInstanceSetMetric          `json:"metric,omitempty"`
	RouteMapInstanceSetMetricTypeValue          RouteMapInstanceSetMetricType      `json:"metric-type,omitempty"`
	RouteMapInstanceSetOriginEgp                RouteMapInstanceSetOrigin          `json:"origin,omitempty"`
	RouteMapInstanceSetOriginatorIDOriginatorIP RouteMapInstanceSetOriginatorID    `json:"originator-id,omitempty"`
	RouteMapInstanceSetTagValue                 RouteMapInstanceSetTag             `json:"tag,omitempty"`
	RouteMapInstanceSetUUID                     string                             `json:"uuid,omitempty"`
	RouteMapInstanceSetWeightWeightVal          RouteMapInstanceSetWeight          `json:"weight,omitempty"`
}

type RouteMapInstanceMatchAsPath struct {
	RouteMapInstanceMatchAsPathName string `json:"name,omitempty"`
}

type RouteMapInstanceMatchCommunity struct {
	RouteMapInstanceMatchCommunityNameCfgName RouteMapInstanceMatchCommunityNameCfg `json:"name-cfg,omitempty"`
}

type RouteMapInstanceMatchExtcommunity struct {
	RouteMapInstanceMatchExtcommunityExtcommunityLNameName RouteMapInstanceMatchExtcommunityExtcommunityLName `json:"extcommunity-l-name,omitempty"`
}

type RouteMapInstanceMatchGroup struct {
	RouteMapInstanceMatchGroupGroupID int    `json:"group-id,omitempty"`
	RouteMapInstanceMatchGroupHaState string `json:"ha-state,omitempty"`
}

type RouteMapInstanceMatchIP struct {
	RouteMapInstanceMatchIPAddressAcl1 RouteMapInstanceMatchIPAddress `json:"address,omitempty"`
	RouteMapInstanceMatchIPNextHopAcl1 RouteMapInstanceMatchIPNextHop `json:"next-hop,omitempty"`
	RouteMapInstanceMatchIPPeerAcl1    RouteMapInstanceMatchIPPeer    `json:"peer,omitempty"`
	RouteMapInstanceMatchIPRibExact    RouteMapInstanceMatchIPRib     `json:"rib,omitempty"`
}

type RouteMapInstanceMatchInterface struct {
	RouteMapInstanceMatchInterfaceEthernet int `json:"ethernet,omitempty"`
	RouteMapInstanceMatchInterfaceLoopback int `json:"loopback,omitempty"`
	RouteMapInstanceMatchInterfaceTrunk    int `json:"trunk,omitempty"`
	RouteMapInstanceMatchInterfaceTunnel   int `json:"tunnel,omitempty"`
	RouteMapInstanceMatchInterfaceVe       int `json:"ve,omitempty"`
}

type RouteMapInstanceMatchIpv6 struct {
	RouteMapInstanceMatchIpv6Address1Name           RouteMapInstanceMatchIpv6Address1 `json:"address-1,omitempty"`
	RouteMapInstanceMatchIpv6NextHop1NextHopAclName RouteMapInstanceMatchIpv6NextHop1 `json:"next-hop-1,omitempty"`
	RouteMapInstanceMatchIpv6Peer1Acl1              RouteMapInstanceMatchIpv6Peer1    `json:"peer-1,omitempty"`
	RouteMapInstanceMatchIpv6RibExact               RouteMapInstanceMatchIpv6Rib      `json:"rib,omitempty"`
}

type RouteMapInstanceMatchLocalPreference struct {
	RouteMapInstanceMatchLocalPreferenceVal int `json:"val,omitempty"`
}

type RouteMapInstanceMatchMetric struct {
	RouteMapInstanceMatchMetricValue int `json:"value,omitempty"`
}

type RouteMapInstanceMatchOrigin struct {
	RouteMapInstanceMatchOriginEgp        int `json:"egp,omitempty"`
	RouteMapInstanceMatchOriginIgp        int `json:"igp,omitempty"`
	RouteMapInstanceMatchOriginIncomplete int `json:"incomplete,omitempty"`
}

type RouteMapInstanceMatchRouteType struct {
	RouteMapInstanceMatchRouteTypeExternalValue RouteMapInstanceMatchRouteTypeExternal `json:"external,omitempty"`
}

type RouteMapInstanceMatchScaleout struct {
	RouteMapInstanceMatchScaleoutClusterID        int    `json:"cluster-id,omitempty"`
	RouteMapInstanceMatchScaleoutOperationalState string `json:"operational-state,omitempty"`
}

type RouteMapInstanceMatchTag struct {
	RouteMapInstanceMatchTagValue int `json:"value,omitempty"`
}

type RouteMapInstanceSetAggregator struct {
	RouteMapInstanceSetAggregatorAggregatorAsAsn RouteMapInstanceSetAggregatorAggregatorAs `json:"aggregator-as,omitempty"`
}

type RouteMapInstanceSetAsPath struct {
	RouteMapInstanceSetAsPathNum     string `json:"num,omitempty"`
	RouteMapInstanceSetAsPathNum2    string `json:"num2,omitempty"`
	RouteMapInstanceSetAsPathPrepend string `json:"prepend,omitempty"`
}

type RouteMapInstanceSetCommList struct {
	RouteMapInstanceSetCommListDelete     int    `json:"delete,omitempty"`
	RouteMapInstanceSetCommListName       string `json:"name,omitempty"`
	RouteMapInstanceSetCommListNameDelete int    `json:"name-delete,omitempty"`
	RouteMapInstanceSetCommListVExp       int    `json:"v-exp,omitempty"`
	RouteMapInstanceSetCommListVExpDelete int    `json:"v-exp-delete,omitempty"`
	RouteMapInstanceSetCommListVStd       int    `json:"v-std,omitempty"`
}

type RouteMapInstanceSetDampeningCfg struct {
	RouteMapInstanceSetDampeningCfgDampening           int `json:"dampening,omitempty"`
	RouteMapInstanceSetDampeningCfgDampeningHalfTime   int `json:"dampening-half-time,omitempty"`
	RouteMapInstanceSetDampeningCfgDampeningMaxSupress int `json:"dampening-max-supress,omitempty"`
	RouteMapInstanceSetDampeningCfgDampeningPenalty    int `json:"dampening-penalty,omitempty"`
	RouteMapInstanceSetDampeningCfgDampeningReuse      int `json:"dampening-reuse,omitempty"`
	RouteMapInstanceSetDampeningCfgDampeningSupress    int `json:"dampening-supress,omitempty"`
}

type RouteMapInstanceSetDdos struct {
	RouteMapInstanceSetDdosClassListCid  int    `json:"class-list-cid,omitempty"`
	RouteMapInstanceSetDdosClassListName string `json:"class-list-name,omitempty"`
	RouteMapInstanceSetDdosZone          string `json:"zone,omitempty"`
}

type RouteMapInstanceSetExtcommunity struct {
	RouteMapInstanceSetExtcommunityRtValue  RouteMapInstanceSetExtcommunityRt  `json:"rt,omitempty"`
	RouteMapInstanceSetExtcommunitySooValue RouteMapInstanceSetExtcommunitySoo `json:"soo,omitempty"`
}

type RouteMapInstanceSetIP struct {
	RouteMapInstanceSetIPNextHopAddress RouteMapInstanceSetIPNextHop `json:"next-hop,omitempty"`
}

type RouteMapInstanceSetIpv6 struct {
	RouteMapInstanceSetIpv6NextHop1Address RouteMapInstanceSetIpv6NextHop1 `json:"next-hop-1,omitempty"`
}

type RouteMapInstanceSetLevel struct {
	RouteMapInstanceSetLevelValue string `json:"value,omitempty"`
}

type RouteMapInstanceSetLocalPreference struct {
	RouteMapInstanceSetLocalPreferenceVal int `json:"val,omitempty"`
}

type RouteMapInstanceSetMetric struct {
	RouteMapInstanceSetMetricValue string `json:"value,omitempty"`
}

type RouteMapInstanceSetMetricType struct {
	RouteMapInstanceSetMetricTypeValue string `json:"value,omitempty"`
}

type RouteMapInstanceSetOrigin struct {
	RouteMapInstanceSetOriginEgp        int `json:"egp,omitempty"`
	RouteMapInstanceSetOriginIgp        int `json:"igp,omitempty"`
	RouteMapInstanceSetOriginIncomplete int `json:"incomplete,omitempty"`
}

type RouteMapInstanceSetOriginatorID struct {
	RouteMapInstanceSetOriginatorIDOriginatorIP string `json:"originator-ip,omitempty"`
}

type RouteMapInstanceSetTag struct {
	RouteMapInstanceSetTagValue int `json:"value,omitempty"`
}

type RouteMapInstanceSetWeight struct {
	RouteMapInstanceSetWeightWeightVal int `json:"weight-val,omitempty"`
}

type RouteMapInstanceMatchCommunityNameCfg struct {
	RouteMapInstanceMatchCommunityNameCfgExactMatch int    `json:"exact-match,omitempty"`
	RouteMapInstanceMatchCommunityNameCfgName       string `json:"name,omitempty"`
}

type RouteMapInstanceMatchExtcommunityExtcommunityLName struct {
	RouteMapInstanceMatchExtcommunityExtcommunityLNameExactMatch int    `json:"exact-match,omitempty"`
	RouteMapInstanceMatchExtcommunityExtcommunityLNameName       string `json:"name,omitempty"`
}

type RouteMapInstanceMatchIPAddress struct {
	RouteMapInstanceMatchIPAddressAcl1           int                                      `json:"acl1,omitempty"`
	RouteMapInstanceMatchIPAddressAcl2           int                                      `json:"acl2,omitempty"`
	RouteMapInstanceMatchIPAddressName           string                                   `json:"name,omitempty"`
	RouteMapInstanceMatchIPAddressPrefixListName RouteMapInstanceMatchIPAddressPrefixList `json:"prefix-list,omitempty"`
}

type RouteMapInstanceMatchIPNextHop struct {
	RouteMapInstanceMatchIPNextHopAcl1            int                                       `json:"acl1,omitempty"`
	RouteMapInstanceMatchIPNextHopAcl2            int                                       `json:"acl2,omitempty"`
	RouteMapInstanceMatchIPNextHopName            string                                    `json:"name,omitempty"`
	RouteMapInstanceMatchIPNextHopPrefixList1Name RouteMapInstanceMatchIPNextHopPrefixList1 `json:"prefix-list-1,omitempty"`
}

type RouteMapInstanceMatchIPPeer struct {
	RouteMapInstanceMatchIPPeerAcl1 int    `json:"acl1,omitempty"`
	RouteMapInstanceMatchIPPeerAcl2 int    `json:"acl2,omitempty"`
	RouteMapInstanceMatchIPPeerName string `json:"name,omitempty"`
}

type RouteMapInstanceMatchIPRib struct {
	RouteMapInstanceMatchIPRibExact       string `json:"exact,omitempty"`
	RouteMapInstanceMatchIPRibReachable   string `json:"reachable,omitempty"`
	RouteMapInstanceMatchIPRibUnreachable string `json:"unreachable,omitempty"`
}

type RouteMapInstanceMatchIpv6Address1 struct {
	RouteMapInstanceMatchIpv6Address1Name            string                                       `json:"name,omitempty"`
	RouteMapInstanceMatchIpv6Address1PrefixList2Name RouteMapInstanceMatchIpv6Address1PrefixList2 `json:"prefix-list-2,omitempty"`
}

type RouteMapInstanceMatchIpv6NextHop1 struct {
	RouteMapInstanceMatchIpv6NextHop1NextHopAclName string `json:"next-hop-acl-name,omitempty"`
	RouteMapInstanceMatchIpv6NextHop1PrefixListName string `json:"prefix-list-name,omitempty"`
	RouteMapInstanceMatchIpv6NextHop1V6Addr         string `json:"v6-addr,omitempty"`
}

type RouteMapInstanceMatchIpv6Peer1 struct {
	RouteMapInstanceMatchIpv6Peer1Acl1 int    `json:"acl1,omitempty"`
	RouteMapInstanceMatchIpv6Peer1Acl2 int    `json:"acl2,omitempty"`
	RouteMapInstanceMatchIpv6Peer1Name string `json:"name,omitempty"`
}

type RouteMapInstanceMatchIpv6Rib struct {
	RouteMapInstanceMatchIpv6RibExact       string `json:"exact,omitempty"`
	RouteMapInstanceMatchIpv6RibReachable   string `json:"reachable,omitempty"`
	RouteMapInstanceMatchIpv6RibUnreachable string `json:"unreachable,omitempty"`
}

type RouteMapInstanceMatchRouteTypeExternal struct {
	RouteMapInstanceMatchRouteTypeExternalValue string `json:"value,omitempty"`
}

type RouteMapInstanceSetAggregatorAggregatorAs struct {
	RouteMapInstanceSetAggregatorAggregatorAsAsn int    `json:"asn,omitempty"`
	RouteMapInstanceSetAggregatorAggregatorAsIP  string `json:"ip,omitempty"`
}

type RouteMapInstanceSetExtcommunityRt struct {
	RouteMapInstanceSetExtcommunityRtValue string `json:"value,omitempty"`
}

type RouteMapInstanceSetExtcommunitySoo struct {
	RouteMapInstanceSetExtcommunitySooValue string `json:"value,omitempty"`
}

type RouteMapInstanceSetIPNextHop struct {
	RouteMapInstanceSetIPNextHopAddress string `json:"address,omitempty"`
}

type RouteMapInstanceSetIpv6NextHop1 struct {
	RouteMapInstanceSetIpv6NextHop1Address      string                               `json:"address,omitempty"`
	RouteMapInstanceSetIpv6NextHop1LocalAddress RouteMapInstanceSetIpv6NextHop1Local `json:"local,omitempty"`
}

type RouteMapInstanceMatchIPAddressPrefixList struct {
	RouteMapInstanceMatchIPAddressPrefixListName string `json:"name,omitempty"`
}

type RouteMapInstanceMatchIPNextHopPrefixList1 struct {
	RouteMapInstanceMatchIPNextHopPrefixList1Name string `json:"name,omitempty"`
}

type RouteMapInstanceMatchIpv6Address1PrefixList2 struct {
	RouteMapInstanceMatchIpv6Address1PrefixList2Name string `json:"name,omitempty"`
}

type RouteMapInstanceSetIpv6NextHop1Local struct {
	RouteMapInstanceSetIpv6NextHop1LocalAddress string `json:"address,omitempty"`
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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/route-map", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
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
		err := json.Unmarshal(data, &m)
		if err != nil {
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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/route-map/"+name1+"+"+name2+"+"+name3, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
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
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouteMap
		err := json.Unmarshal(data, &m)
		if err != nil {
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
	return err
}
