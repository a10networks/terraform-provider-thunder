package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Icmpv6 struct {
	UUID Icmpv6Instance `json:"icmpv6,omitempty"`
}

type Icmpv6Instance struct {
	Redirect    int    `json:"redirect"`
	Unreachable int    `json:"unreachable"`
	UUID        string `json:"uuid"`
}

func PostIpv6Icmpv6(id string, inst Icmpv6, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpv6Icmpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ipv6/icmpv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Icmpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpv6Icmpv6 REQ RES..........................", m)
			err := check_api_status("PostIpv6Icmpv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpv6Icmpv6(id string, host string) (*Icmpv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpv6Icmpv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ipv6/icmpv6/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Icmpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GEGetIpv6Icmpv6 REQ RES..........................", m)
			err := check_api_status("GetIpv6Icmpv6", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
