package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type DnsSuffix struct {
	UUID DnsSuffixInstance `json:"suffix,omitempty"`
}

type DnsSuffixInstance struct {
	DomainName string `json:"domain-name"`
	UUID       string `json:"uuid"`
}

func PostIpDnsSuffix(id string, inst DnsSuffix, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpDnsSuffix")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/dns/suffix", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DnsSuffix
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpDnsSuffix REQ RES..........................", m)
			err := check_api_status("PostIpDnsSuffix", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpDnsSuffix(id string, host string) (*DnsSuffix, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpDnsSuffix")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/dns/suffix/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DnsSuffix
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpDnsSuffix REQ RES..........................", m)
			err := check_api_status("GetIpDnsSuffix", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
