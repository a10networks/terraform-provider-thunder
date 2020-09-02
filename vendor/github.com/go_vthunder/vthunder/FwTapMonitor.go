package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwTapMonitor struct {
	Status FwTapMonitorInstance `json:"tap-monitor-instance,omitempty"`
}

type FwTapMonitorInstance struct {
	Status string                `json:"status,omitempty"`
	TapEth []FwTapMonitorPortCfg `json:"tap-port-cfg,omitempty"`
	UUID   string                `json:"uuid,omitempty"`
}

type FwTapMonitorPortCfg struct {
	TapEth  int `json:"tap-eth,omitempty"`
	TapVlan int `json:"tap-vlan,omitempty"`
}

func PostFwTapMonitor(id string, inst FwTapMonitor, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTapMonitor")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/tap-monitor", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTapMonitor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetFwTapMonitor(id string, host string) (*FwTapMonitor, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTapMonitor")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/tap-monitor/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTapMonitor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}
