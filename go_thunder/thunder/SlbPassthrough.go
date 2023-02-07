package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Passthrough struct {
	Counters1 PassthroughInstance `json:"passthrough,omitempty"`
}

type SamplingEnable24 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type PassthroughInstance struct {
	Counters1 []SamplingEnable24 `json:"sampling-enable,omitempty"`
}

func PostSlbPassthrough(id string, inst Passthrough, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbPassthrough")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/passthrough", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Passthrough
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbPassthrough REQ RES..........................", m)
			err := check_api_status("PostSlbPassthrough", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbPassthrough(id string, host string) (*Passthrough, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbPassthrough")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/passthrough/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Passthrough
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbPassthrough REQ RES..........................", m)
			err := check_api_status("GetSlbPassthrough", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
