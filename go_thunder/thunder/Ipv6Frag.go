package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type IPv6Frag struct {
	UUID IPv6FragInstance `json:"frag,omitempty"`
}
type SamplingEnable49 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type IPv6FragInstance struct {
	Counters1   []SamplingEnable49 `json:"sampling-enable,omitempty"`
	UUID        string             `json:"uuid,omitempty"`
	FragTimeout int                `json:"frag-timeout,omitempty"`
}

func PostIpv6Frag(id string, inst IPv6Frag, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpv6Frag")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ipv6/frag", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IPv6Frag
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpv6Frag REQ RES..........................", m)
			err := check_api_status("PostIpv6Frag", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpv6Frag(id string, host string) (*IPv6Frag, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpv6Frag")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ipv6/frag/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IPv6Frag
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetIpv6Frag", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
