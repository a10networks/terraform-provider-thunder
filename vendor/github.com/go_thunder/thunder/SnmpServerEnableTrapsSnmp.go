package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsSnmp struct {
	Linkup SnmpServerEnableTrapsSnmpInstance `json:"snmp,omitempty"`
}

type SnmpServerEnableTrapsSnmpInstance struct {
	All      int    `json:"all,omitempty"`
	Linkdown int    `json:"linkdown,omitempty"`
	Linkup   int    `json:"linkup,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsSnmp(id string, inst SnmpServerEnableTrapsSnmp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsSnmp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/snmp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSnmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsSnmp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsSnmp(id string, host string) (*SnmpServerEnableTrapsSnmp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsSnmp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/snmp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSnmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsSnmp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
