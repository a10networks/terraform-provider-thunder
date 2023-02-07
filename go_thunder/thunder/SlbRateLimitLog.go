package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RateLimitLog struct {
	Counters1 RateLimitLogInstance `json:"rate-limit-log,omitempty"`
}

type SamplingEnable30 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type RateLimitLogInstance struct {
	Counters1 []SamplingEnable30 `json:"sampling-enable,omitempty"`
}

func PostSlbRateLimitLog(id string, inst RateLimitLog, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbRateLimitLog")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/rate-limit-log", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RateLimitLog
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbRateLimitLog REQ RES..........................", m)
			err := check_api_status("PostSlbRateLimitLog", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbRateLimitLog(id string, host string) (*RateLimitLog, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbRateLimitLog")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/rate-limit-log/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RateLimitLog
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbRateLimitLog REQ RES..........................", m)
			err := check_api_status("GetSlbRateLimitLog", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
