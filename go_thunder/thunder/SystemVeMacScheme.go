package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SystemVeMacScheme struct {
	SystemVeMacSchemeInstanceVeMacSchemeVal SystemVeMacSchemeInstance `json:"ve-mac-scheme,omitempty"`
}

type SystemVeMacSchemeInstance struct {
	SystemVeMacSchemeInstanceUUID           string `json:"uuid,omitempty"`
	SystemVeMacSchemeInstanceVeMacSchemeVal string `json:"ve-mac-scheme-val,omitempty"`
}

func PostSystemVeMacScheme(id string, inst SystemVeMacScheme, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/system/ve-mac-scheme", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SystemVeMacScheme
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSystemVeMacScheme", data)
			if err != nil {
				return err
			}

		}
	}
	return err
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
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSystemVeMacScheme", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
