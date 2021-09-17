package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwLocalLog struct {
	LocalLogging FwLocalLogInstance `json:"local-log-instance,omitempty"`
}

type FwLocalLogInstance struct {
	LocalLogging int    `json:"local-logging,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostFwLocalLog(id string, inst FwLocalLog, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwLocalLog")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/local-log", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwLocalLog
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwLocalLog", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwLocalLog(id string, host string) (*FwLocalLog, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwLocalLog")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/local-log/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwLocalLog
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwLocalLog", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
