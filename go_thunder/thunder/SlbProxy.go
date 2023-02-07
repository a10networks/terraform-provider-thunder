package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Proxy struct {
	Counters1 ProxyInstance `json:"proxy,omitempty"`
}

type SamplingEnable29 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type ProxyInstance struct {
	Counters1 []SamplingEnable29 `json:"sampling-enable,omitempty"`
}

func PostSlbProxy(id string, inst Proxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbProxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Proxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbProxy REQ RES..........................", m)
			err := check_api_status("PostSlbProxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbProxy(id string, host string) (*Proxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbProxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/proxy/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Proxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbProxy REQ RES..........................", m)
			err := check_api_status("GetSlbProxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
