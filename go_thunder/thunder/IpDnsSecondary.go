package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type DnsSecondary struct {
	UUID DnsSecondaryInstance `json:"Secondary,omitempty"`
}
type DnsSecondaryInstance struct {
	IPV4Addr string `json:"ip-v4-addr,omitempty"`
	IPV6Addr string `json:"ip-v6-addr,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

func PostIpDnsSecondary(id string, inst DnsSecondary, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpDnsSecondary")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/dns/Secondary", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DnsSecondary
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpDnsSecondary REQ RES..........................", m)
			err := check_api_status("PostIpDnsSecondary", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpDnsSecondary(id string, host string) (*DnsSecondary, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpDnsSecondary")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/dns/Secondary/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DnsSecondary
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpDnsSecondary REQ RES..........................", m)
			err := check_api_status("GetIpDnsSecondary", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
