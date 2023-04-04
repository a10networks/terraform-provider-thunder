package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwAlgSip struct {
	DefaultPortDisable FwAlgSipInstance `json:"sip,omitempty"`
}

type FwAlgSipInstance struct {
	DefaultPortDisable string                   `json:"default-port-disable,omitempty"`
	Counters1          []FwAlgSipSamplingEnable `json:"sampling-enable,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
}

type FwAlgSipSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwAlgSip(id string, inst FwAlgSip, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwAlgSip")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/alg/sip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgSip
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("POSTFwAlgSip", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwAlgSip(id string, host string) (*FwAlgSip, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwAlgSip")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/alg/sip/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgSip
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwAlgSip", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
