package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwActiveRuleSet struct {
	SessionAging FwActiveRuleSetInstance `json:"active-rule-set-instance,omitempty"`
}

type FwActiveRuleSetInstance struct {
	Name             string `json:"name,omitempty"`
	OverrideNatAging int    `json:"override-nat-aging,omitempty"`
	SessionAging     string `json:"session-aging,omitempty"`
	UUID             string `json:"uuid,omitempty"`
}

func PostFwActiveRuleSet(id string, inst FwActiveRuleSet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwActiveRuleSet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/active-rule-set", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwActiveRuleSet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("PostFwActiveRuleSet", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetFwActiveRuleSet(id string, host string) (*FwActiveRuleSet, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwActiveRuleSet")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/active-rule-set/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwActiveRuleSet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwActiveRuleSet", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
