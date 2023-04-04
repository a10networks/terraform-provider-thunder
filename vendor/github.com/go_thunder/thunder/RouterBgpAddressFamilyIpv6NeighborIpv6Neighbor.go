package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor struct {
	NeighborIpv6 RouterBgpAddressFamilyIpv6NeighborIpv6NeighborInstance `json:"ipv6-neighbor,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborInstance struct {
	Activate            int                                                                   `json:"activate,omitempty"`
	AllowasIn           int                                                                   `json:"allowas-in,omitempty"`
	AllowasInCount      int                                                                   `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                                                                   `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists       `json:"distribute-lists,omitempty"`
	GracefulRestart     int                                                                   `json:"graceful-restart,omitempty"`
	Inbound             int                                                                   `json:"inbound,omitempty"`
	MaximumPrefix       int                                                                   `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                                                                   `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv6        string                                                                `json:"neighbor-ipv6,omitempty"`
	NbrPrefixList       []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
	NextHopSelf         int                                                                   `json:"next-hop-self,omitempty"`
	PeerGroupName       string                                                                `json:"peer-group-name,omitempty"`
	PrefixListDirection string                                                                `json:"prefix-list-direction,omitempty"`
	RemovePrivateAs     int                                                                   `json:"remove-private-as,omitempty"`
	RouteMap            string                                                                `json:"route-map,omitempty"`
	SendCommunityVal    string                                                                `json:"send-community-val,omitempty"`
	UUID                string                                                                `json:"uuid,omitempty"`
	UnsuppressMap       string                                                                `json:"unsuppress-map,omitempty"`
	Weight              int                                                                   `json:"weight,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborDistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv6NeighborNeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(id string, name1 string, inst RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ipv6-neighbor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(id string, name1 string, name2 string, host string) (*RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ipv6-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor(id string, name1 string, name2 string, inst RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ipv6-neighbor/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborIpv6Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpAddressFamilyIpv6NeighborIpv6Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
