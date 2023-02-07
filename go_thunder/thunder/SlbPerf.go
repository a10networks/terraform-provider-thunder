package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Perf struct {
	Counters1 PerfInstance `json:"perf,omitempty"`
}

type SamplingEnable25 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type PerfInstance struct {
	Counters1 []SamplingEnable25 `json:"sampling-enable,omitempty"`
}

func PostSlbPerf(id string, inst Perf, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbPerf")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/perf", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Perf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbPerf REQ RES..........................", m)
			err := check_api_status("PostSlbPerf", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetSlbPerf(id string, host string) (*Perf, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbPerf")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/perf/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Perf
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbPerf REQ RES..........................", m)
			err := check_api_status("GetSlbPerf", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
