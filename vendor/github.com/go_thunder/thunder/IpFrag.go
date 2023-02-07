package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Frag struct {
	UUID FragInstance `json:"frag,omitempty"`
}

type SamplingEnable47 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type CPUThreshold struct {
	High int `json:"high,omitempty"`
	Low  int `json:"low,omitempty"`
}
type FragInstance struct {
	UUID                    string             `json:"uuid,omitempty"`
	MaxReassemblySessions   int                `json:"max-reassembly-sessions,omitempty"`
	Counters1               []SamplingEnable47 `json:"sampling-enable,omitempty"`
	High                    CPUThreshold       `json:"cpu-threshold,omitempty"`
	Timeout                 int                `json:"timeout,omitempty"`
	MaxPacketsPerReassembly int                `json:"max-packets-per-reassembly,omitempty"`
	Buff                    int                `json:"buff,omitempty"`
}

func PostIpFrag(id string, inst Frag, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpFrag")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/frag", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Frag
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpFrag REQ RES..........................", m)
			err := check_api_status("PostIpFrag", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpFrag(id string, host string) (*Frag, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpFrag")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/frag/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Frag
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpFrag REQ RES..........................", m)
			err := check_api_status("GetIpFrag", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
