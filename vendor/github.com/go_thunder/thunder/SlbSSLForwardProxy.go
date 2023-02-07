package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SSLForwardProxy struct {
	Counters1 SSLForwardProxyInstance `json:"ssl-forward-proxy,omitempty"`
}

type SamplingEnable38 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SSLForwardProxyInstance struct {
	Counters1 []SamplingEnable38 `json:"sampling-enable,omitempty"`
}

func PostSlbSSLForwardProxy(id string, inst SSLForwardProxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSSLForwardProxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/ssl-forward-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SSLForwardProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSSLForwardProxy REQ RES..........................", m)
			err := check_api_status("PostSlbSSLForwardProxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSSLForwardProxy(id string, host string) (*SSLForwardProxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSSLForwardProxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/ssl-forward-proxy/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SSLForwardProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSSLForwardProxy REQ RES..........................", m)
			err := check_api_status("GetSlbSSLForwardProxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
