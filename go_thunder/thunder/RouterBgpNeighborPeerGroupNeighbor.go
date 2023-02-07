package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpNeighborPeerGroupNeighbor struct {
	PeerGroup RouterBgpNeighborPeerGroupNeighborInstance `json:"peer-group-neighbor,omitempty"`
}

type RouterBgpNeighborPeerGroupNeighborInstance struct {
	Activate                 int                                                       `json:"activate,omitempty"`
	AdvertisementInterval    int                                                       `json:"advertisement-interval,omitempty"`
	AllowasIn                int                                                       `json:"allowas-in,omitempty"`
	AllowasInCount           int                                                       `json:"allowas-in-count,omitempty"`
	AsOriginationInterval    int                                                       `json:"as-origination-interval,omitempty"`
	Bfd                      int                                                       `json:"bfd,omitempty"`
	CollideEstablished       int                                                       `json:"collide-established,omitempty"`
	Connect                  int                                                       `json:"connect,omitempty"`
	DefaultOriginate         int                                                       `json:"default-originate,omitempty"`
	Description              string                                                    `json:"description,omitempty"`
	DisallowInfiniteHoldtime int                                                       `json:"disallow-infinite-holdtime,omitempty"`
	DistributeList           []RouterBgpNeighborPeerGroupNeighborDistributeLists       `json:"distribute-lists,omitempty"`
	DontCapabilityNegotiate  int                                                       `json:"dont-capability-negotiate,omitempty"`
	Dynamic                  int                                                       `json:"dynamic,omitempty"`
	EbgpMultihop             int                                                       `json:"ebgp-multihop,omitempty"`
	EbgpMultihopHopCount     int                                                       `json:"ebgp-multihop-hop-count,omitempty"`
	EnforceMultihop          int                                                       `json:"enforce-multihop,omitempty"`
	Ethernet                 int                                                       `json:"ethernet,omitempty"`
	ExtendedNexthop          int                                                       `json:"extended-nexthop,omitempty"`
	Inbound                  int                                                       `json:"inbound,omitempty"`
	Lif                      int                                                       `json:"lif,omitempty"`
	Loopback                 int                                                       `json:"loopback,omitempty"`
	MaximumPrefix            int                                                       `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres       int                                                       `json:"maximum-prefix-thres,omitempty"`
	Multihop                 int                                                       `json:"multihop,omitempty"`
	FilterList               []RouterBgpNeighborPeerGroupNeighborNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NbrPrefixList            []RouterBgpNeighborPeerGroupNeighborNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap              []RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf              int                                                       `json:"next-hop-self,omitempty"`
	OverrideCapability       int                                                       `json:"override-capability,omitempty"`
	PassValue                string                                                    `json:"pass-value,omitempty"`
	Passive                  int                                                       `json:"passive,omitempty"`
	PeerGroup                string                                                    `json:"peer-group,omitempty"`
	PeerGroupKey             int                                                       `json:"peer-group-key,omitempty"`
	PeerGroupRemoteAs        int                                                       `json:"peer-group-remote-as,omitempty"`
	PrefixListDirection      string                                                    `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs          int                                                       `json:"remove-private-as,omitempty"`
	RouteMap                 string                                                    `json:"route-map,omitempty"`
	RouteRefresh             int                                                       `json:"route-refresh,omitempty"`
	SendCommunityVal         string                                                    `json:"send-community-val,omitempty"`
	Shutdown                 int                                                       `json:"shutdown,omitempty"`
	StrictCapabilityMatch    int                                                       `json:"strict-capability-match,omitempty"`
	TimersHoldtime           int                                                       `json:"timers-holdtime,omitempty"`
	TimersKeepalive          int                                                       `json:"timers-keepalive,omitempty"`
	Trunk                    int                                                       `json:"trunk,omitempty"`
	Tunnel                   int                                                       `json:"tunnel,omitempty"`
	UUID                     string                                                    `json:"uuid,omitempty"`
	UnsuppressMap            string                                                    `json:"unsuppress-map,omitempty"`
	UpdateSourceIP           string                                                    `json:"update-source-ip,omitempty"`
	UpdateSourceIpv6         string                                                    `json:"update-source-ipv6,omitempty"`
	Ve                       int                                                       `json:"ve,omitempty"`
	Weight                   int                                                       `json:"weight,omitempty"`
}

type RouterBgpNeighborPeerGroupNeighborDistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpNeighborPeerGroupNeighborNeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpNeighborPeerGroupNeighborNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpNeighborPeerGroupNeighborNeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

func PostRouterBgpNeighborPeerGroupNeighbor(id string, name1 string, inst RouterBgpNeighborPeerGroupNeighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpNeighborPeerGroupNeighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/peer-group-neighbor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpNeighborPeerGroupNeighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpNeighborPeerGroupNeighbor(id string, name1 string, name2 string, host string) (*RouterBgpNeighborPeerGroupNeighbor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpNeighborPeerGroupNeighbor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/peer-group-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpNeighborPeerGroupNeighbor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpNeighborPeerGroupNeighbor(id string, name1 string, name2 string, inst RouterBgpNeighborPeerGroupNeighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpNeighborPeerGroupNeighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/peer-group-neighbor/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpNeighborPeerGroupNeighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRouterBgpNeighborPeerGroupNeighbor(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterBgpNeighborPeerGroupNeighbor")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/peer-group-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouterBgpNeighborPeerGroupNeighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
