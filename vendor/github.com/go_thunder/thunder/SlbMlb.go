package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Mlb struct {
	Counters1 MlbInstance `json:"mlb,omitempty"`
}

type SamplingEnable44 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type MlbInstance struct {
	Counters1 []SamplingEnable44 `json:"sampling-enable,omitempty"`
}

func PostSlbMlb(id string, inst Mlb, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbMlb")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/mlb", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Mlb
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbMlb REQ RES..........................", m)
			err := check_api_status("PostSlbMlb", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbMlb(id string, host string) (*Mlb, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbMlb")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/mlb/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Mlb
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbMlb REQ RES..........................", m)
			err := check_api_status("GetSlbMlb", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
