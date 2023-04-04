package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RcCacheGlobal struct {
	Counters1 RcCacheGlobalInstance `json:"rc-cache-global,omitempty"`
}

type SamplingEnable31 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type RcCacheGlobalInstance struct {
	Counters1 []SamplingEnable31 `json:"sampling-enable,omitempty"`
}

func PostSlbRcCacheGlobal(id string, inst RcCacheGlobal, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbRcCacheGlobal")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/rc-cache-global", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RcCacheGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbRcCacheGlobal REQ RES..........................", m)
			err := check_api_status("PostSlbRcCacheGlobal", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbRcCacheGlobal(id string, host string) (*RcCacheGlobal, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbRcCacheGlobal")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/rc-cache-global/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RcCacheGlobal
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbRcCacheGlobal REQ RES..........................", m)
			err := check_api_status("GetSlbRcCacheGlobal", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
