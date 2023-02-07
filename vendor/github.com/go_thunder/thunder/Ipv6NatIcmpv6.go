package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type NatIcmpv6 struct {
	UUID NatIcmpv6Instance `json:"icmpv6,omitempty"`
}
type NatIcmpv6Instance struct {
	RespondToPing int    `json:"respond-to-ping,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

func PostIpv6NatIcmpv6(id string, inst NatIcmpv6, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpv6NatIcmpv6")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ipv6/nat/icmpv6", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatIcmpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpv6NatIcmpv6 REQ RES..........................", m)
			err := check_api_status("PostIpv6NatIcmpv6", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpv6NatIcmpv6(id string, host string) (*NatIcmpv6, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpv6NatIcmpv6")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ipv6/nat/icmpv6/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatIcmpv6
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpv6NatIcmpv6 REQ RES..........................", m)
			err := check_api_status("GetIpv6NatIcmpv6", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
