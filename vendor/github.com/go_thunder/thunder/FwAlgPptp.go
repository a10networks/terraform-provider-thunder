package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwAlgPptp struct {
	DefaultPortDisable FwAlgPptpInstance `json:"pptp,omitempty"`
}

type FwAlgPptpInstance struct {
	DefaultPortDisable string                    `json:"default-port-disable,omitempty"`
	Counters1          []FwAlgPptpSamplingEnable `json:"sampling-enable,omitempty"`
	UUID               string                    `json:"uuid,omitempty"`
}

type FwAlgPptpSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwAlgPptp(id string, inst FwAlgPptp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwAlgPptp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/alg/pptp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgPptp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwAlgPptp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwAlgPptp(id string, host string) (*FwAlgPptp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwAlgPptp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/alg/pptp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgPptp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwAlgPptp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
