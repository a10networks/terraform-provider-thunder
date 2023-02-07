package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type NatIcmp struct {
	UUID NatIcmpInstance `json:"icmp,omitempty"`
}
type NatIcmpInstance struct {
	AlwaysSourceNatErrors int    `json:"always-source-nat-errors,omitempty"`
	RespondToPing         int    `json:"respond-to-ping,omitempty"`
	UUID                  string `json:"uuid,omitempty"`
}

func PostIpNatIcmp(id string, inst NatIcmp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpNatIcmp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/nat/icmp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatIcmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpNatIcmp REQ RES..........................", m)
			err := check_api_status("PostIpNatIcmp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpNatIcmp(id string, host string) (*NatIcmp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpNatIcmp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/nat/icmp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatIcmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpNatIcmp REQ RES..........................", m)
			err := check_api_status("GetIpNatIcmp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
