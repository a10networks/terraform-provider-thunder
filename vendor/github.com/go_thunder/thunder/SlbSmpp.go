package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbSmpp struct {
	Counters1 SlbSmppInstance `json:"smpp,omitempty"`
}

type SamplingEnable11 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SlbSmppInstance struct {
	Counters1 []SamplingEnable11 `json:"sampling-enable,omitempty"`
}

func PostSlbSmpp(id string, inst SlbSmpp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSmpp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/smpp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbSmpp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSmpp REQ RES..........................", m)
			err := check_api_status("PostSlbSmpp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSmpp(id string, host string) (*SlbSmpp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSmpp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/smpp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbSmpp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSmpp REQ RES..........................", m)
			err := check_api_status("GetSlbSmpp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
