package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FastHttpProxy struct {
	UUID FastHttpProxyInstance `json:"fast-http-proxy,omitempty"`
}

type SamplingEnable18 struct {
	Counters1 string `json:"counters1,omitempty"`
	Counters2 string `json:"counters2,omitempty"`
}
type FastHttpProxyInstance struct {
	Counters1 []SamplingEnable18 `json:"sampling-enable,omitempty"`
	UUID      string             `json:"uuid,omitempty"`
}

func GetSlbFastHttpProxy(id string, host string) (*FastHttpProxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/fast-http-proxy", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FastHttpProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetSlbFastHttpProxy REQ RES..........................", m)
			err := check_api_status("GetSlbFastHttpProxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostSlbFastHttpProxy(id string, vc FastHttpProxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/fast-http-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v FastHttpProxy
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostSlbFastHttpProxy", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
