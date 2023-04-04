package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTopKRules struct {
	UUID FwTopKRulesInstance `json:"top-k-rules,omitempty"`
}

type FwTopKRulesInstance struct {
	UUID string `json:"uuid,omitempty"`
}

func PostFwTopKRules(id string, inst FwTopKRules, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTopKRules")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/top-k-rules", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTopKRules
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTopKRules REQ RES..........................", m)
			err := check_api_status("PostFwTopKRules", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwTopKRules(id string, host string) (*FwTopKRules, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTopKRules")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/top-k-rules/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTopKRules
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwTopKRules REQ RES..........................", m)
			err := check_api_status("GetFwTopKRules", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
