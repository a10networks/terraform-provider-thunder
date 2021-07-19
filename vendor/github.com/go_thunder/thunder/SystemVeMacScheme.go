package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SystemVeMacScheme struct {
	VeMacSchemeVal SystemVeMacSchemeInstance `json:"ve-mac-scheme,omitempty"`
}

type SystemVeMacSchemeInstance struct {
	UUID           string `json:"uuid,omitempty"`
	VeMacSchemeVal string `json:"ve-mac-scheme-val,omitempty"`
}

func PostSystemVeMacScheme(id string, inst SystemVeMacScheme, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSystemVeMacScheme")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/system/ve-mac-scheme", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SystemVeMacScheme
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSystemVeMacScheme", data)

		}
	}

}

func GetSystemVeMacScheme(id string, host string) (*SystemVeMacScheme, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSystemVeMacScheme")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/system/ve-mac-scheme", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SystemVeMacScheme
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSystemVeMacScheme", data)
			return &m, nil
		}
	}

}
