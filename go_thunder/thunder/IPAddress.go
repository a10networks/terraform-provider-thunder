package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type IPAddress struct {
	UUID IPAddressInstance `json:"address,omitempty"`
}

type IPAddressInstance struct {
	IPAddr string `json:"ip-addr,omitempty"`
	IPMask string `json:"ip-mask,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

func PostIPAddress(id string, inst IPAddress, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIPAddress")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/address", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IPAddress
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIPAddress REQ RES..........................", m)
			err := check_api_status("PostIPAddress", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetIPAddress(id string, host string) (*IPAddress, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIPAddress")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/address/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IPAddress
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIPAddress REQ RES..........................", m)
			err := check_api_status("GetIPAddress", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
