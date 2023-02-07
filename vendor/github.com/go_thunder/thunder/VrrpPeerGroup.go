package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type PeerGroup struct {
	UUID PeerGroupInstance `json:"peer-group,omitempty"`
}
type IPPeerAddressCfg struct {
	IPPeerAddress string `json:"ip-peer-address,omitempty"`
}
type Ipv6PeerAddressCfg struct {
	Ipv6PeerAddress string `json:"ipv6-peer-address,omitempty"`
}
type Peer struct {
	IPPeerAddress   []IPPeerAddressCfg   `json:"ip-peer-address-cfg,omitempty"`
	Ipv6PeerAddress []Ipv6PeerAddressCfg `json:"ipv6-peer-address-cfg,omitempty"`
}
type PeerGroupInstance struct {
	IPPeerAddressCfg Peer   `json:"peer,omitempty"`
	UUID             string `json:"uuid,omitempty"`
}

func GetVrrpPeerGroup(id string, host string) (*PeerGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/vrrp-a/peer-group", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m PeerGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetVrrpPeerGroup", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostVrrpPeerGroup(id string, vc PeerGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/vrrp-a/peer-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v PeerGroup
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostVrrpPeerGroup", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
