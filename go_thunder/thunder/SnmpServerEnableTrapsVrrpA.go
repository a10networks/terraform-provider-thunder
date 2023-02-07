package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsVrrpA struct {
	Active SnmpServerEnableTrapsVrrpAInstance `json:"vrrp-a,omitempty"`
}

type SnmpServerEnableTrapsVrrpAInstance struct {
	Active  int    `json:"active,omitempty"`
	All     int    `json:"all,omitempty"`
	Standby int    `json:"standby,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsVrrpA(id string, inst SnmpServerEnableTrapsVrrpA, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsVrrpA")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/vrrp-a", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsVrrpA
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsVrrpA", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsVrrpA(id string, host string) (*SnmpServerEnableTrapsVrrpA, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsVrrpA")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/vrrp-a/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsVrrpA
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsVrrpA", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
