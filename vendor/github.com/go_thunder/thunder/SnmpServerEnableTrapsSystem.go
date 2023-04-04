package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsSystem struct {
	All SnmpServerEnableTrapsSystemInstance `json:"system,omitempty"`
}

type SnmpServerEnableTrapsSystemInstance struct {
	All                int    `json:"all,omitempty"`
	ControlCPUHigh     int    `json:"control-cpu-high,omitempty"`
	DataCPUHigh        int    `json:"data-cpu-high,omitempty"`
	Fan                int    `json:"fan,omitempty"`
	FileSysReadOnly    int    `json:"file-sys-read-only,omitempty"`
	HighDiskUse        int    `json:"high-disk-use,omitempty"`
	HighMemoryUse      int    `json:"high-memory-use,omitempty"`
	HighTemp           int    `json:"high-temp,omitempty"`
	LicenseManagement  int    `json:"license-management,omitempty"`
	LowTemp            int    `json:"low-temp,omitempty"`
	PacketDrop         int    `json:"packet-drop,omitempty"`
	Power              int    `json:"power,omitempty"`
	PriDisk            int    `json:"pri-disk,omitempty"`
	Restart            int    `json:"restart,omitempty"`
	SecDisk            int    `json:"sec-disk,omitempty"`
	Shutdown           int    `json:"shutdown,omitempty"`
	SmpResourceEvent   int    `json:"smp-resource-event,omitempty"`
	Start              int    `json:"start,omitempty"`
	SyslogSeverityOne  int    `json:"syslog-severity-one,omitempty"`
	TacacsServerUpDown int    `json:"tacacs-server-up-down,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsSystem(id string, inst SnmpServerEnableTrapsSystem, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsSystem")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/system", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSystem
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsSystem", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsSystem(id string, host string) (*SnmpServerEnableTrapsSystem, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsSystem")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/system/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSystem
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsSystem", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
