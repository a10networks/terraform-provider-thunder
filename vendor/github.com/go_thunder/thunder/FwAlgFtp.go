package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwAlgFtp struct {
	DefaultPortDisable FwAlgFtpInstance `json:"ftp,omitempty"`
}

type FwAlgFtpInstance struct {
	DefaultPortDisable string                   `json:"default-port-disable,omitempty"`
	Counters1          []FwAlgFtpSamplingEnable `json:"sampling-enable,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
}

type FwAlgFtpSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwAlgFtp(id string, inst FwAlgFtp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwAlgFtp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/alg/ftp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgFtp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwAlgFtp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwAlgFtp(id string, host string) (*FwAlgFtp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwAlgFtp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/alg/ftp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwAlgFtp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwAlgFtp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
