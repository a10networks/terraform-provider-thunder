package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTcpWindowCheck struct {
	Status FwTCPWindowCheckInstance `json:"tcp-window-check,omitempty"`
}

type FwTCPWindowCheckInstance struct {
	Counters1 []FwTCPWindowCheckSamplingEnable `json:"sampling-enable,omitempty"`
	Status    string                           `json:"status,omitempty"`
	UUID      string                           `json:"uuid,omitempty"`
}

type FwTCPWindowCheckSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwTcpWindowCheck(id string, inst FwTcpWindowCheck, host string) error {

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
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpWindowCheck
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTcpWindowCheck REQ RES..........................", m)
			err := check_api_status("PostFwTcpWindowCheck", data)
			if err != nil {
				return err
			}

		}
	}
	return err
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
			err := check_api_status("GetFwTcpWindowCheck", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
