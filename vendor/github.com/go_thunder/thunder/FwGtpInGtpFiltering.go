package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwGtpInGtpFiltering struct {
	GtpInGtpValue FwGtpInGtpFilteringInstance `json:"gtp-in-gtp-filtering,omitempty"`
}

type FwGtpInGtpFilteringInstance struct {
	GtpInGtpValue string `json:"gtp-in-gtp-value,omitempty"`
	UUID          string `json:"uuid,omitempty"`
}

func PostFwGtpInGtpFiltering(id string, inst FwGtpInGtpFiltering, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwGtpInGtpFiltering")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/gtp-in-gtp-filtering", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGtpInGtpFiltering
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwGtpInGtpFiltering", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwGtpInGtpFiltering(id string, host string) (*FwGtpInGtpFiltering, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwGtpInGtpFiltering")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/gtp-in-gtp-filtering/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwGtpInGtpFiltering
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwGtpInGtpFiltering", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
