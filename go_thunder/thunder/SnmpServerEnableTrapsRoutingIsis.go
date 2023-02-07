package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerEnableTrapsRoutingIsis struct {
	IsisAuthenticationFailure SnmpServerEnableTrapsRoutingIsisInstance `json:"isis,omitempty"`
}

type SnmpServerEnableTrapsRoutingIsisInstance struct {
	IsisAdjacencyChange                  int    `json:"isisAdjacencyChange,omitempty"`
	IsisAreaMismatch                     int    `json:"isisAreaMismatch,omitempty"`
	IsisAttemptToExceedMaxSequence       int    `json:"isisAttemptToExceedMaxSequence,omitempty"`
	IsisAuthenticationFailure            int    `json:"isisAuthenticationFailure,omitempty"`
	IsisAuthenticationTypeFailure        int    `json:"isisAuthenticationTypeFailure,omitempty"`
	IsisCorruptedLSPDetected             int    `json:"isisCorruptedLSPDetected,omitempty"`
	IsisDatabaseOverload                 int    `json:"isisDatabaseOverload,omitempty"`
	IsisIDLenMismatch                    int    `json:"isisIDLenMismatch,omitempty"`
	IsisLSPTooLargeToPropagate           int    `json:"isisLSPTooLargeToPropagate,omitempty"`
	IsisManualAddressDrops               int    `json:"isisManualAddressDrops,omitempty"`
	IsisMaxAreaAddressesMismatch         int    `json:"isisMaxAreaAddressesMismatch,omitempty"`
	IsisOriginatingLSPBufferSizeMismatch int    `json:"isisOriginatingLSPBufferSizeMismatch,omitempty"`
	IsisOwnLSPPurge                      int    `json:"isisOwnLSPPurge,omitempty"`
	IsisProtocolsSupportedMismatch       int    `json:"isisProtocolsSupportedMismatch,omitempty"`
	IsisRejectedAdjacency                int    `json:"isisRejectedAdjacency,omitempty"`
	IsisSequenceNumberSkip               int    `json:"isisSequenceNumberSkip,omitempty"`
	IsisVersionSkew                      int    `json:"isisVersionSkew,omitempty"`
	UUID                                 string `json:"uuid,omitempty"`
}

func PostSnmpServerEnableTrapsRoutingIsis(id string, inst SnmpServerEnableTrapsRoutingIsis, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerEnableTrapsRoutingIsis")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/enable/traps/routing/isis", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsRoutingIsis
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSnmpServerEnableTrapsRoutingIsis", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSnmpServerEnableTrapsRoutingIsis(id string, host string) (*SnmpServerEnableTrapsRoutingIsis, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerEnableTrapsRoutingIsis")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/enable/traps/routing/isis/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerEnableTrapsRoutingIsis
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSnmpServerEnableTrapsRoutingIsis", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
