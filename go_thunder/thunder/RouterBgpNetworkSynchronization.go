package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RouterBgpNetworkSynchronization struct {
	NetworkSynchronization RouterBgpNetworkSynchronizationInstance `json:"synchronization,omitempty"`
}

type RouterBgpNetworkSynchronizationInstance struct {
	NetworkSynchronization int    `json:"network-synchronization,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
}

func PostRouterBgpNetworkSynchronization(id string, name1 string, inst RouterBgpNetworkSynchronization, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterBgpNetworkSynchronization")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/network/synchronization", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNetworkSynchronization
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterBgpNetworkSynchronization", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterBgpNetworkSynchronization(id string, name1 string, host string) (*RouterBgpNetworkSynchronization, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterBgpNetworkSynchronization")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/bgp/"+name1+"/network/synchronization", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterBgpNetworkSynchronization
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterBgpNetworkSynchronization", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
