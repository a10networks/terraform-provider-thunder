package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type WebCategoryReputationScope struct {
	Name WebCategoryReputationScopeInstance `json:"reputation-scope,omitempty"`
}

type WebCategoryReputationScopeInstance struct {
	GreaterTrustworthy WebCategoryGreaterThan           `json:"greater-than,omitempty"`
	LessTrustworthy    WebCategoryLessThan              `json:"less-than,omitempty"`
	Name               string                           `json:"name,omitempty"`
	Counters1          []WebCategoryScopeSamplingEnable `json:"sampling-enable,omitempty"`
	UUID               string                           `json:"uuid,omitempty"`
	UserTag            string                           `json:"user-tag,omitempty"`
}

type WebCategoryGreaterThan struct {
	GreaterLowRisk      int `json:"greater-low-risk,omitempty"`
	GreaterMalicious    int `json:"greater-malicious,omitempty"`
	GreaterModerateRisk int `json:"greater-moderate-risk,omitempty"`
	GreaterSuspicious   int `json:"greater-suspicious,omitempty"`
	GreaterThreshold    int `json:"greater-threshold,omitempty"`
	GreaterTrustworthy  int `json:"greater-trustworthy,omitempty"`
}

type WebCategoryLessThan struct {
	LessLowRisk      int `json:"less-low-risk,omitempty"`
	LessMalicious    int `json:"less-malicious,omitempty"`
	LessModerateRisk int `json:"less-moderate-risk,omitempty"`
	LessSuspicious   int `json:"less-suspicious,omitempty"`
	LessThreshold    int `json:"less-threshold,omitempty"`
	LessTrustworthy  int `json:"less-trustworthy,omitempty"`
}

type WebCategoryScopeSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostWebCategoryReputationScope(id string, inst WebCategoryReputationScope, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostWebCategoryReputationScope")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/web-category/reputation-scope", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryReputationScope
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostWebCategoryReputationScope", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetWebCategoryReputationScope(id string, name1 string, host string) (*WebCategoryReputationScope, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetWebCategoryReputationScope")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/web-category/reputation-scope/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryReputationScope
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetWebCategoryReputationScope", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutWebCategoryReputationScope(id string, name1 string, inst WebCategoryReputationScope, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutWebCategoryReputationScope")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/web-category/reputation-scope/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryReputationScope
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutWebCategoryReputationScope", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteWebCategoryReputationScope(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteWebCategoryReputationScope")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/web-category/reputation-scope/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryReputationScope
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteWebCategoryReputationScope", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
