package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Fix struct {
	Name FixInstance `json:"fix,omitempty"`
}
type TagSwitching struct {
	ServiceGroup  string `json:"service-group,omitempty"`
	Equals        string `json:"equals,omitempty"`
	SwitchingType string `json:"switching-type,omitempty"`
}
type FixInstance struct {
	Logging        string         `json:"logging,omitempty"`
	Name           string         `json:"name,omitempty"`
	ServiceGroup   []TagSwitching `json:"tag-switching,omitempty"`
	UserTag        string         `json:"user-tag,omitempty"`
	InsertClientIP int            `json:"insert-client-ip,omitempty"`
	UUID           string         `json:"uuid,omitempty"`
}

func PostTemplateFix(id string, inst Fix, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplateFix")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/fix", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Fix
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostTemplateFix REQ RES..........................", m)
			err := check_api_status("PostTemplateFix", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetTemplateFix(id string, name string, host string) (*Fix, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplateFix")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/fix/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Fix
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetTemplateFix REQ RES..........................", m)
			err := check_api_status("GetTemplateFix", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutTemplateFix(id string, name string, inst Fix, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplateFix")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/fix/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Fix
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutTemplateFix REQ RES..........................", m)
			err := check_api_status("PutTemplateFix", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteTemplateFix(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplateFix")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/fix/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Fix
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
