package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwLogging struct {
	SamplingEnable FwLoggingInstance `json:"logging-instance,omitempty"`
}

type FwLoggingInstance struct {
	Name      string                    `json:"name,omitempty"`
	Counters1 []FwLoggingSamplingEnable `json:"sampling-enable,omitempty"`
	UUID      string                    `json:"uuid,omitempty"`
}

type FwLoggingSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwLogging(id string, inst FwLogging, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwLogging")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/logging", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwLogging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			check_api_status("PostFwLogging", data)

		}
	}

}

func GetFwLogging(id string, host string) (*FwLogging, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwLogging")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/logging/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwLogging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			check_api_status("GetFwLogging", data)
			return &m, nil
		}
	}

}
