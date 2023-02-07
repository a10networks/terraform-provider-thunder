package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor struct {
	NeighborIpv4 RouterBgpAddressFamilyIpv6NeighborIpv4NeighborInstance `json:"ipv4-neighbor,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborInstance struct {
	Activate            int                                                                   `json:"activate,omitempty"`
	AllowasIn           int                                                                   `json:"allowas-in,omitempty"`
	AllowasInCount      int                                                                   `json:"allowas-in-count,omitempty"`
	DefaultOriginate    int                                                                   `json:"default-originate,omitempty"`
	DistributeList      []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborDistributeLists       `json:"distribute-lists,omitempty"`
	GracefulRestart     int                                                                   `json:"graceful-restart,omitempty"`
	Inbound             int                                                                   `json:"inbound,omitempty"`
	MaximumPrefix       int                                                                   `json:"maximum-prefix,omitempty"`
	MaximumPrefixThres  int                                                                   `json:"maximum-prefix-thres,omitempty"`
	FilterList          []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborNeighborFilterLists   `json:"neighbor-filter-lists,omitempty"`
	NeighborIpv4        string                                                                `json:"neighbor-ipv4,omitempty"`
	NbrPrefixList       []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborNeighborPrefixLists   `json:"neighbor-prefix-lists,omitempty"`
	NbrRouteMap         []RouterBgpAddressFamilyIpv6NeighborIpv4NeighborNeighborRouteMapLists `json:"neighbor-route-map-lists,omitempty"`
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

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborDistributeLists struct {
	DistributeList          string `json:"distribute-list,omitempty"`
	DistributeListDirection string `json:"distribute-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborNeighborFilterLists struct {
	FilterList          string `json:"filter-list,omitempty"`
	FilterListDirection string `json:"filter-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborNeighborPrefixLists struct {
	NbrPrefixList          string `json:"nbr-prefix-list,omitempty"`
	NbrPrefixListDirection string `json:"nbr-prefix-list-direction,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborIpv4NeighborNeighborRouteMapLists struct {
	NbrRmapDirection string `json:"nbr-rmap-direction,omitempty"`
	NbrRouteMap      string `json:"nbr-route-map,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor(id string, name1 string, inst RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ipv4-neighbor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor(id string, name1 string, name2 string, host string) (*RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ipv4-neighbor/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor(id string, name1 string, name2 string, inst RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ipv4-neighbor/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborIpv4Neighbor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpAddressFamilyIpv6NeighborIpv4Neighbor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
