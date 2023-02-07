package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor struct {
	PeerGroup RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborInstance `json:"peer-group-neighbor,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborInstance struct {
	Activate            int                                                                  `json:"activate,omitempty"`
	AllowasIn           int                                                                  `json:"allowas-in,omitempty"`
	AllowasInCount      int                                                                  `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                                                                  `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDistributeLists `json:"distribute-lists,omitempty"`
	Inbound             int                                                                  `json:"inbound,omitempty"`
	MaximumPrefix       int                                                                  `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                                                                  `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborFilterLists     `json:"neighbor-filter-lists,omitempty"`
	NbrPrefixList       []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborPrefixLists     `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRouteMapLists   `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                                                                  `json:"next-hop-self,omitempty"`
	PeerGroup           string                                                               `json:"peer-group,omitempty"`
	PrefixListDirection string                                                               `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                                                                  `json:"remove-private-as,omitempty"`
	RouteMap            string                                                               `json:"route-map,omitempty"`
	SendCommunityVal    string                                                               `json:"send-community-val,omitempty"`
	UUID                string                                                               `json:"uuid,omitempty"`
	UnsuppressMap       string                                                               `json:"unsuppress-map,omitempty"`
	Weight              int                                                                  `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborDistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(id string, name1 string, inst RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/peer-group-neighbor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(id string, name1 string, name2 string, host string) (*RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/peer-group-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor(id string, name1 string, name2 string, inst RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/peer-group-neighbor/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpAddressFamilyIpv6NeighborPeerGroupNeighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
