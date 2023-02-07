package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Icap struct {
	Counters1 IcapInstance `json:"icap,omitempty"`
}
type SamplingEnable39 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type IcapInstance struct {
	Counters1 []SamplingEnable39 `json:"sampling-enable,omitempty"`
}

func PostSlbIcap(id string, inst Icap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbIcap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/icap", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Icap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbIcap REQ RES..........................", m)
			err := check_api_status("PostSlbIcap", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbIcap(id string, host string) (*Icap, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbIcap")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/icap", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Icap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] PostSlbIcap REQ RES..........................", m)
			err := check_api_status("GetSlbIcap", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
