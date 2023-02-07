package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsSlb struct {
	All SnmpServerEnableTrapsSlbInstance `json:"slb,omitempty"`
}

type SnmpServerEnableTrapsSlbInstance struct {
	All                    int    `json:"all,omitempty"`
	ApplicationBufferLimit int    `json:"application-buffer-limit,omitempty"`
	BwRateLimitExceed      int    `json:"bw-rate-limit-exceed,omitempty"`
	BwRateLimitResume      int    `json:"bw-rate-limit-resume,omitempty"`
	GatewayDown            int    `json:"gateway-down,omitempty"`
	GatewayUp              int    `json:"gateway-up,omitempty"`
	ServerConnLimit        int    `json:"server-conn-limit,omitempty"`
	ServerConnResume       int    `json:"server-conn-resume,omitempty"`
	ServerDisabled         int    `json:"server-disabled,omitempty"`
	ServerDown             int    `json:"server-down,omitempty"`
	ServerSelectionFailure int    `json:"server-selection-failure,omitempty"`
	ServerUp               int    `json:"server-up,omitempty"`
	ServiceConnLimit       int    `json:"service-conn-limit,omitempty"`
	ServiceConnResume      int    `json:"service-conn-resume,omitempty"`
	ServiceDown            int    `json:"service-down,omitempty"`
	ServiceGroupDown       int    `json:"service-group-down,omitempty"`
	ServiceGroupMemberDown int    `json:"service-group-member-down,omitempty"`
	ServiceGroupMemberUp   int    `json:"service-group-member-up,omitempty"`
	ServiceGroupUp         int    `json:"service-group-up,omitempty"`
	ServiceUp              int    `json:"service-up,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
	VipConnlimit           int    `json:"vip-connlimit,omitempty"`
	VipConnratelimit       int    `json:"vip-connratelimit,omitempty"`
	VipDown                int    `json:"vip-down,omitempty"`
	VipPortConnlimit       int    `json:"vip-port-connlimit,omitempty"`
	VipPortConnratelimit   int    `json:"vip-port-connratelimit,omitempty"`
	VipPortDown            int    `json:"vip-port-down,omitempty"`
	VipPortUp              int    `json:"vip-port-up,omitempty"`
	VipUp                  int    `json:"vip-up,omitempty"`
}

func PostSnmpServerEnableTrapsSlb(id string, inst SnmpServerEnableTrapsSlb, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsSlb")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/slb", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSlb
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsSlb", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsSlb(id string, host string) (*SnmpServerEnableTrapsSlb, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsSlb")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/slb/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsSlb
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsSlb", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
