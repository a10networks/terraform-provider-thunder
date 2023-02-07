package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwGlobal struct {
	AlgProcessing FwGlobalInstance `json:"global,omitempty"`
}

type FwGlobalInstance struct {
	AlgProcessing              string                   `json:"alg-processing,omitempty"`
	DisableApplicationProtocol []FwGlobalDisableAppList `json:"disable-app-list,omitempty"`
	DisableIPFwSessions        int                      `json:"disable-ip-fw-sessions,omitempty"`
	ExtendedMatching           string                   `json:"extended-matching,omitempty"`
	ListenOnPortTimeout        int                      `json:"listen-on-port-timeout,omitempty"`
	NatipDdosProtection        string                   `json:"natip-ddos-protection,omitempty"`
	PermitDefaultAction        string                   `json:"permit-default-action,omitempty"`
	RespondToUserMac           int                      `json:"respond-to-user-mac,omitempty"`
	Counters1                  []FwGlobalSamplingEnable `json:"sampling-enable,omitempty"`
	UUID                       string                   `json:"uuid,omitempty"`
}

type FwGlobalDisableAppList struct {
	DisableApplicationCategory string `json:"disable-application-category,omitempty"`
	DisableApplicationProtocol string `json:"disable-application-protocol,omitempty"`
}

type FwGlobalSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwGlobal(id string, inst FwGlobal, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwGlobal")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes 0 - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/global", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwGlobal", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwGlobal(id string, host string) (*FwGlobal, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwGlobal")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/global/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwGlobal", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
