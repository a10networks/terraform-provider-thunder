package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwTcpWindowCheck struct {
	Status FwTCPWindowCheckInstance `json:"tcp-window-check-instance,omitempty"`
}

type FwTCPWindowCheckInstance struct {
	Counters1 []FwTCPWindowCheckSamplingEnable `json:"sampling-enable,omitempty"`
	Status    string                           `json:"status,omitempty"`
	UUID      string                           `json:"uuid,omitempty"`
}

type FwTCPWindowCheckSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwTcpWindowCheck(id string, inst FwTcpWindowCheck, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTcpWindowCheck")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/tcp-window-check", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpWindowCheck
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTcpWindowCheck REQ RES..........................", m)
			check_api_status("PostFwTcpWindowCheck", data)

		}
	}

}

func GetFwTcpWindowCheck(id string, host string) (*FwTcpWindowCheck, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTcpWindowCheck")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/tcp-window-check/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpWindowCheck
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwTcpWindowCheck REQ RES..........................", m)
			check_api_status("GetFwTcpWindowCheck", data)
			return &m, nil
		}
	}

}
