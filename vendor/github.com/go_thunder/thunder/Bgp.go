package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Bgp1 struct {
	ExtendedAsnCap BgpInstance `json:"bgp,omitempty"`
}

type BgpInstance struct {
	ExtendedAsnCap int            `json:"extended-asn-cap,omitempty"`
	Enable         NexthopTrigger `json:"nexthop-trigger,omitempty"`
	UUID           string         `json:"uuid,omitempty"`
}

type NexthopTrigger struct {
	Delay  int `json:"delay,omitempty"`
	Enable int `json:"enable,omitempty"`
}

func PostBgp(id string, inst Bgp1, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostBgp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/bgp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Bgp1
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostBgp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetBgp(id string, host string) (*Bgp1, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetBgp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/bgp", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Bgp1
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetBgp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
