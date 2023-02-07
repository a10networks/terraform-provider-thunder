package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbPersist struct {
	Counters1 SlbPersistInstance `json:"persist,omitempty"`
}

type SamplingEnable26 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SlbPersistInstance struct {
	Counters1 []SamplingEnable26 `json:"sampling-enable,omitempty"`
}

func PostSlbpersist(id string, inst SlbPersist, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbpersist")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/persist", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbPersist
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbpersist REQ RES..........................", m)
			err := check_api_status("PostSlbpersist", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbpersist(id string, host string) (*SlbPersist, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbpersist")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/persist/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbPersist
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbpersist REQ RES..........................", m)
			err := check_api_status("GetSlbpersist", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
