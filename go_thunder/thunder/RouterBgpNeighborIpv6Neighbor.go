package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpNeighborIpv6Neighbor struct {
	NeighborIpv6 RouterBgpNeighborIpv6NeighborInstance `json:"ipv6-neighbor,omitempty"`
}

type RouterBgpNeighborIpv6NeighborInstance struct {
	AcosApplicationOnly      int                                                  `json:"acos-application-only,omitempty"`
	Activate                 int                                                  `json:"activate,omitempty"`
	AdvertisementInterval    int                                                  `json:"advertisement-interval,omitempty"`
	AllowasIn                int                                                  `json:"allowas-in,omitempty"`
	AllowasInCount           int                                                  `json:"allowas-in-count,omitempty"`
	AsOriginationInterval    int                                                  `json:"as-origination-interval,omitempty"`
	Bfd                      int                                                  `json:"bfd,omitempty"`
	BfdValue                 string                                               `json:"bfd-value,omitempty"`
	CollideEstablished       int                                                  `json:"collide-established,omitempty"`
	Connect                  int                                                  `json:"connect,omitempty"`
	DefaultOriginate         int                                                  `json:"default-originate,omitempty"`
	Description              string                                               `json:"description,omitempty"`
	DisallowInfiniteHoldtime int                                                  `json:"disallow-infinite-holdtime,omitempty"`
	DistributeList           []RouterBgpNeighborIpv6NeighborDistributeLists       `json:"distribute-lists,omitempty"`
	DontCapabilityNegotiate  int                                                  `json:"dont-capability-negotiate,omitempty"`
	Dynamic                  int                                                  `json:"dynamic,omitempty"`
	EbgpMultihop             int                                                  `json:"ebgp-multihop,omitempty"`
	EbgpMultihopHopCount     int                                                  `json:"ebgp-multihop-hop-count,omitempty"`
	EnforceMultihop          int                                                  `json:"enforce-multihop,omitempty"`
	Ethernet                 int                                                  `json:"ethernet,omitempty"`
	ExtendedNexthop          int                                                  `json:"extended-nexthop,omitempty"`
	GracefulRestart          int                                                  `json:"graceful-restart,omitempty"`
	Inbound                  int                                                  `json:"inbound,omitempty"`
	KeyID                    int                                                  `json:"key-id,omitempty"`
	KeyType                  string                                               `json:"key-type,omitempty"`
	Lif                      int                                                  `json:"lif,omitempty"`
	Loopback                 int                                                  `json:"loopback,omitempty"`
	MaximumPrefix            int                                                  `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres       int                                                  `json:"maximum-prefix-thres,omitempty"`
	Multihop                 int                                                  `json:"multihop,omitempty"`
	NbrRemoteAs              int                                                  `json:"nbr-remote-as,omitempty"`
	FilterList               []RouterBgpNeighborIpv6NeighborNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv6             string                                               `json:"neighbor-ipv6,omitempty"`
	NbrPrefixList            []RouterBgpNeighborIpv6NeighborNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap              []RouterBgpNeighborIpv6NeighborNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf              int                                                  `json:"next-hop-self,omitempty"`
	OverrideCapability       int                                                  `json:"override-capability,omitempty"`
	PassValue                string                                               `json:"pass-value,omitempty"`
	Passive                  int                                                  `json:"passive,omitempty"`
	PeerGroupName            string                                               `json:"peer-group-name,omitempty"`
	PrefixListDirection      string                                               `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs          int                                                  `json:"remove-private-as,omitempty"`
	RouteMap                 string                                               `json:"route-map,omitempty"`
	RouteRefresh             int                                                  `json:"route-refresh,omitempty"`
	SendCommunityVal         string                                               `json:"send-community-val,omitempty"`
	Shutdown                 int                                                  `json:"shutdown,omitempty"`
	StrictCapabilityMatch    int                                                  `json:"strict-capability-match,omitempty"`
	TimersHoldtime           int                                                  `json:"timers-holdtime,omitempty"`
	TimersKeepalive          int                                                  `json:"timers-keepalive,omitempty"`
	Trunk                    int                                                  `json:"trunk,omitempty"`
	Tunnel                   int                                                  `json:"tunnel,omitempty"`
	UUID                     string                                               `json:"uuid,omitempty"`
	UnsuppressMap            string                                               `json:"unsuppress-map,omitempty"`
	UpdateSourceIP           string                                               `json:"update-source-ip,omitempty"`
	UpdateSourceIpv6         string                                               `json:"update-source-ipv6,omitempty"`
	Ve                       int                                                  `json:"ve,omitempty"`
	Weight                   int                                                  `json:"weight,omitempty"`
}

type RouterBgpNeighborIpv6NeighborDistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpNeighborIpv6NeighborNeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpNeighborIpv6NeighborNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpNeighborIpv6NeighborNeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

func PostRouterBgpNeighborIpv6Neighbor(id string, name1 string, inst RouterBgpNeighborIpv6Neighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpNeighborIpv6Neighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/ipv6-neighbor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpNeighborIpv6Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpNeighborIpv6Neighbor(id string, name1 string, name2 string, host string) (*RouterBgpNeighborIpv6Neighbor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpNeighborIpv6Neighbor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/ipv6-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpNeighborIpv6Neighbor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpNeighborIpv6Neighbor(id string, name1 string, name2 string, inst RouterBgpNeighborIpv6Neighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpNeighborIpv6Neighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/ipv6-neighbor/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpNeighborIpv6Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRouterBgpNeighborIpv6Neighbor(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterBgpNeighborIpv6Neighbor")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/neighbor/ipv6-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouterBgpNeighborIpv6Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
