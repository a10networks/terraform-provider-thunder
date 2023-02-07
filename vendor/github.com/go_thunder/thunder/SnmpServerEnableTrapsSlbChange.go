package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsSlbChange struct {
	All SnmpServerEnableTrapsSlbChangeInstance `json:"slb-change,omitempty"`
}

type SnmpServerEnableTrapsSlbChangeInstance struct {
	All                     int    `json:"all,omitempty"`
	ConnectionResourceEvent int    `json:"connection-resource-event,omitempty"`
	ResourceUsageWarning    int    `json:"resource-usage-warning,omitempty"`
	Server                  int    `json:"server,omitempty"`
	ServerPort              int    `json:"server-port,omitempty"`
	SslCertChange           int    `json:"ssl-cert-change,omitempty"`
	SslCertExpire           int    `json:"ssl-cert-expire,omitempty"`
	SystemThreshold         int    `json:"system-threshold,omitempty"`
	UUID                    string `json:"uuid,omitempty"`
	Vip                     int    `json:"vip,omitempty"`
	VipPort                 int    `json:"vip-port,omitempty"`
}

func PostSnmpServerEnableTrapsSlbChange(id string, inst SnmpServerEnableTrapsSlbChange, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsSlbChange")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/slb-change", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSlbChange
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsSlbChange", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsSlbChange(id string, host string) (*SnmpServerEnableTrapsSlbChange, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsSlbChange")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/slb-change/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSlbChange
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsSlbChange", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
