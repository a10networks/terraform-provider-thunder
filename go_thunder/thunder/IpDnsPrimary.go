package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type DnsPrimary struct {
	UUID DnsPrimaryInstance `json:"primary,omitempty"`
}
type DnsPrimaryInstance struct {
	IPV4Addr string `json:"ip-v4-addr,omitempty"`
	IPV6Addr string `json:"ip-v6-addr,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

func PostIpDnsPrimary(id string, inst DnsPrimary, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpDnsPrimary")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/dns/primary", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DnsPrimary
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpDnsPrimary REQ RES..........................", m)
			err := check_api_status("PostIpDnsPrimary", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpDnsPrimary(id string, host string) (*DnsPrimary, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpDnsPrimary")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/dns/primary/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DnsPrimary
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpDnsPrimary REQ RES..........................", m)
			err := check_api_status("GetIpDnsPrimary", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
