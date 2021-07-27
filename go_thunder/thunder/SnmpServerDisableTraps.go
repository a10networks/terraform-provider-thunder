package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerDisableTraps struct {
	All SnmpServerDisableTrapsInstance `json:"traps,omitempty"`
}

type SnmpServerDisableTrapsInstance struct {
	All       int    `json:"all,omitempty"`
	Gslb      int    `json:"gslb,omitempty"`
	Slb       int    `json:"slb,omitempty"`
	SlbChange int    `json:"slb-change,omitempty"`
	Snmp      int    `json:"snmp,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	VrrpA     int    `json:"vrrp-a,omitempty"`
}

func PostSnmpServerDisableTraps(id string, inst SnmpServerDisableTraps, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerDisableTraps")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/disable/traps", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerDisableTraps
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerDisableTraps", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerDisableTraps(id string, host string) (*SnmpServerDisableTraps, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerDisableTraps")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/disable/traps/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerDisableTraps
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerDisableTraps", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
