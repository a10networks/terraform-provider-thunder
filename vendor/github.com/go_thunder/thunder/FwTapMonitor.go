package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTapMonitor struct {
	Status FwTapMonitorInstance `json:"tap-monitor,omitempty"`
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

func PostFwTapMonitor(id string, inst FwTapMonitor, host string) error {

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
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTapMonitor
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTapMonitor REQ RES..........................", m)
			err := check_api_status("PostFwTapMonitor", data)
			if err != nil {
				return err
			}

		}
	}
	return err
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
			err := check_api_status("GetFwTapMonitor", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
