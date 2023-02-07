package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsRoutingOspf struct {
	OspfLsdbOverflow SnmpServerEnableTrapsRoutingOspfInstance `json:"ospf,omitempty"`
}

type SnmpServerEnableTrapsRoutingOspfInstance struct {
	OspfIfAuthFailure           int    `json:"ospfIfAuthFailure,omitempty"`
	OspfIfConfigError           int    `json:"ospfIfConfigError,omitempty"`
	OspfIfRxBadPacket           int    `json:"ospfIfRxBadPacket,omitempty"`
	OspfIfStateChange           int    `json:"ospfIfStateChange,omitempty"`
	OspfLsdbApproachingOverflow int    `json:"ospfLsdbApproachingOverflow,omitempty"`
	OspfLsdbOverflow            int    `json:"ospfLsdbOverflow,omitempty"`
	OspfMaxAgeLsa               int    `json:"ospfMaxAgeLsa,omitempty"`
	OspfNbrStateChange          int    `json:"ospfNbrStateChange,omitempty"`
	OspfOriginateLsa            int    `json:"ospfOriginateLsa,omitempty"`
	OspfTxRetransmit            int    `json:"ospfTxRetransmit,omitempty"`
	OspfVirtIfAuthFailure       int    `json:"ospfVirtIfAuthFailure,omitempty"`
	OspfVirtIfConfigError       int    `json:"ospfVirtIfConfigError,omitempty"`
	OspfVirtIfRxBadPacket       int    `json:"ospfVirtIfRxBadPacket,omitempty"`
	OspfVirtIfStateChange       int    `json:"ospfVirtIfStateChange,omitempty"`
	OspfVirtIfTxRetransmit      int    `json:"ospfVirtIfTxRetransmit,omitempty"`
	OspfVirtNbrStateChange      int    `json:"ospfVirtNbrStateChange,omitempty"`
	UUID                        string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsRoutingOspf(id string, inst SnmpServerEnableTrapsRoutingOspf, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsRoutingOspf")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/routing/ospf", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsRoutingOspf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsRoutingOspf", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsRoutingOspf(id string, host string) (*SnmpServerEnableTrapsRoutingOspf, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsRoutingOspf")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/routing/ospf/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsRoutingOspf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsRoutingOspf", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
