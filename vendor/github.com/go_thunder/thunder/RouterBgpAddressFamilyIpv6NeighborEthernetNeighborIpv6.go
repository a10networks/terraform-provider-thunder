package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6 struct {
	Ethernet RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Instance `json:"ethernet-neighbor-ipv6,omitempty"`
}

type RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Instance struct {
	Ethernet      int    `json:"ethernet,omitempty"`
	PeerGroupName string `json:"peer-group-name,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

func PostRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(id string, name1 string, inst RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ethernet-neighbor-ipv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(id string, name1 string, name2 string, host string) (*RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ethernet-neighbor-ipv6/"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6(id string, name1 string, name2 string, inst RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/address-family/ipv6/neighbor/ethernet-neighbor-ipv6/"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
