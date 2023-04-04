package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"

	"io/ioutil"
	"util"
)

type SpdyProxy struct {
	Counters1 SpdyProxyInstance `json:"spdy-proxy,omitempty"`
}
type SamplingEnable14 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SpdyProxyInstance struct {
	Counters1 []SamplingEnable14 `json:"sampling-enable,omitempty"`
}

func PostSlbSpdyProxy(id string, inst SpdyProxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSpdyProxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/spdy-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SpdyProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSpdyProxy REQ RES..........................", m)
			err := check_api_status("PostSlbSpdyProxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSpdyProxy(id string, host string) (*SpdyProxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSpdyProxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/spdy-proxy/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SpdyProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSpdyProxy REQ RES..........................", m)
			err := check_api_status("GetSlbSpdyProxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
