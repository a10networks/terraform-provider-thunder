package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Hostname struct {
	UUID HostnameInstance `json:"hostname,omitempty"`
}

type HostnameInstance struct {
	UUID  string `json:"uuid,omitempty"`
	Value string `json:"value,omitempty"`
}

func PostHostname(id string, inst Hostname, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostHostname")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/hostname", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Hostname
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostHostname REQ RES..........................", m)
			err := check_api_status("PostHostname", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetHostname(id string, host string) (*Hostname, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetHostname")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/hostname/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Hostname
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetHostname REQ RES..........................", m)
			err := check_api_status("GetHostname", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
