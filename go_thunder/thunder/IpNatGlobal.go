package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type NatGlobal struct {
	UUID NatGlobalInstance `json:"nat-global,omitempty"`
}

type NatGlobalInstance struct {
	UUID             string `json:"uuid,omitempty"`
	ResetIdleTcpConn int    `json:"reset-idle-tcp-conn,omitempty"`
}

func PostIpNatGlobal(id string, inst NatGlobal, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpNatGlobal")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/nat-global", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpNatGlobal REQ RES..........................", m)
			err := check_api_status("PostIpNatGlobal", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpNatGlobal(id string, host string) (*NatGlobal, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpNatGlobal")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/nat-global/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NatGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpNatGlobal REQ RES..........................", m)
			err := check_api_status("GetIpNatGlobal", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
