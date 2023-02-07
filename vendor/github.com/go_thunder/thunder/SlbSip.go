package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbSip struct {
	Counters1 SlbSipInstance `json:"sip,omitempty"`
}
type SamplingEnable10 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type SlbSipInstance struct {
	Counters1 []SamplingEnable10 `json:"sampling-enable,omitempty"`
}

func PostSlbSip(id string, inst SlbSip, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbSip")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/sip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbSip
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbSip REQ RES..........................", m)
			err := check_api_status("PostSlbSip", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbSip(id string, host string) (*SlbSip, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbSip")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/sip/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbSip
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbSip REQ RES..........................", m)
			err := check_api_status("GetSlbSip", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
