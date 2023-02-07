package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwGtp struct {
	SamplingEnable FwGtpInstance `json:"gtp,omitempty"`
}

type FwGtpInstance struct {
	GtpValue  string                `json:"gtp-value,omitempty"`
	Counters1 []FwGtpSamplingEnable `json:"sampling-enable,omitempty"`
	UUID      string                `json:"uuid,omitempty"`
}

type FwGtpSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwGtp(id string, inst FwGtp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwGtp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/gtp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGtp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwGtp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwGtp(id string, host string) (*FwGtp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwGtp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/gtp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGtp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwGtp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
