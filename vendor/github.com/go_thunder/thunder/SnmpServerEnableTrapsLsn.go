package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsLsn struct {
	All SnmpServerEnableTrapsLsnInstance `json:"lsn,omitempty"`
}

type SnmpServerEnableTrapsLsnInstance struct {
	All                           int    `json:"all,omitempty"`
	FixedNatPortMappingFileChange int    `json:"fixed-nat-port-mapping-file-change,omitempty"`
	MaxIpportThreshold            int    `json:"max-ipport-threshold,omitempty"`
	MaxPortThreshold              int    `json:"max-port-threshold,omitempty"`
	PerIPPortUsageThreshold       int    `json:"per-ip-port-usage-threshold,omitempty"`
	TotalPortUsageThreshold       int    `json:"total-port-usage-threshold,omitempty"`
	TrafficExceeded               int    `json:"traffic-exceeded,omitempty"`
	UUID                          string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsLsn(id string, inst SnmpServerEnableTrapsLsn, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsLsn")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/lsn", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsLsn
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsLsn", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsLsn(id string, host string) (*SnmpServerEnableTrapsLsn, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsLsn")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/lsn/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsLsn
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsLsn", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
