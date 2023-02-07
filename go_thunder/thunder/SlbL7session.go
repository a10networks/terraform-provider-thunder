package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"

	"io/ioutil"
	"util"
)

type L7session struct {
	Counters1 L7sessionInstance `json:"l7session,omitempty"`
}

type SamplingEnable43 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type L7sessionInstance struct {
	Counters1 []SamplingEnable43 `json:"sampling-enable,omitempty"`
}

func PostSlbL7session(id string, inst L7session, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlL7session")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/l7session", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m L7session
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbL7session REQ RES..........................", m)
			err := check_api_status("PostSlbL7session", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbL7session(id string, host string) (*L7session, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlL7session")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/l7session/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m L7session
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbL7session REQ RES..........................", m)
			err := check_api_status("GetSlbL7session", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
