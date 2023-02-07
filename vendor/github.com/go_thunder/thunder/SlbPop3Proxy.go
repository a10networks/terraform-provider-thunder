package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Pop3Proxy struct {
	Counters1 Pop3ProxyInstance `json:"pop3-proxy,omitempty"`
}

type SamplingEnable28 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type Pop3ProxyInstance struct {
	Counters1 []SamplingEnable28 `json:"sampling-enable,omitempty"`
}

func PostSlbPop3Proxy(id string, inst Pop3Proxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbPop3Proxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/pop3-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Pop3Proxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbPop3Proxy REQ RES..........................", m)
			err := check_api_status("PostSlbPop3Proxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbPop3Proxy(id string, host string) (*Pop3Proxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbPop3Proxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/pop3-proxy/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Pop3Proxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbPop3Proxy REQ RES..........................", m)
			err := check_api_status("GetSlbPop3Proxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
