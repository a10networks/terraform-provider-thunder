package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type L4 struct {
	Counters1 L4Instance `json:"l4,omitempty"`
}

type SamplingEnable42 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type L4Instance struct {
	Counters1 []SamplingEnable42 `json:"sampling-enable,omitempty"`
}

func PostSlbL4(id string, inst L4, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlL4")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/l4", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m L4
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbL4 REQ RES..........................", m)
			err := check_api_status("PostSlbL4", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbL4(id string, host string) (*L4, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlL4")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/l4/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m L4
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbL4 REQ RES..........................", m)
			err := check_api_status("GetSlbL4", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
