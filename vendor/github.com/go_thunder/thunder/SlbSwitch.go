package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Switch struct {
	Counters1 SwitchInstance `json:"switch,omitempty"`
}

type SamplingEnable13 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SwitchInstance struct {
	Counters1 []SamplingEnable13 `json:"sampling-enable,omitempty"`
}

func PostSlbSwitch(id string, inst Switch, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSwitch")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/switch", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Switch
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSwitch REQ RES..........................", m)
			err := check_api_status("PostSlbSwitch", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSwitch(id string, host string) (*Switch, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSwitch")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/switch/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Switch
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSwitch REQ RES..........................", m)
			err := check_api_status("GetSlbSwitch", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
