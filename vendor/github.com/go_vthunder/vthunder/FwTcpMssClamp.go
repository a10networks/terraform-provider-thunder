package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwTcpMssClamp struct {
	MssSubtract FwTcpMssClampInstance `json:"mss-clamp-instance,omitempty"`
}

type FwTcpMssClampInstance struct {
	Min          int    `json:"min,omitempty"`
	MssClampType string `json:"mss-clamp-type,omitempty"`
	MssSubtract  int    `json:"mss-subtract,omitempty"`
	MssValue     int    `json:"mss-value,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostFwTcpMssClamp(id string, inst FwTcpMssClamp, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTcpMssClamp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/tcp/mss-clamp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpMssClamp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetFwTcpMssClamp(id string, host string) (*FwTcpMssClamp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTcpMssClamp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/tcp/mss-clamp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpMssClamp
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
