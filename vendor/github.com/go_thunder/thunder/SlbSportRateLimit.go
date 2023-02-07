package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SportRateLimit struct {
	Counters1 SportRateLimitInstance `json:"sport-rate-limit,omitempty"`
}

type SamplingEnable15 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SportRateLimitInstance struct {
	Counters1 []SamplingEnable15 `json:"sampling-enable,omitempty"`
}

func PostSlbSportRateLimit(id string, inst SportRateLimit, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSportRateLimit")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/sport-rate-limit", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SportRateLimit
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSportRateLimit REQ RES..........................", m)
			err := check_api_status("PostSlbSportRateLimit", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSportRateLimit(id string, host string) (*SportRateLimit, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSportRateLimit")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/sport-rate-limit/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SportRateLimit
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSportRateLimit REQ RES..........................", m)
			err := check_api_status("GetSlbSportRateLimit", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
