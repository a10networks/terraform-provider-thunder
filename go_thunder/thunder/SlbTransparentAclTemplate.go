package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type TransparentAclTemplate struct {
	Name TransparentAclTemplateInstance `json:"transparent-acl-template,omitempty"`
}

type TransparentAclTemplateInstance struct {
	Name string `json:"name,omitempty"`
}

func PostSlbTransparentAclTemplate(id string, inst TransparentAclTemplate, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTransparentAclTemplate")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/transparent-acl-template", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TransparentAclTemplate
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTransparentAclTemplate REQ RES..........................", m)
			err := check_api_status("PostSlbTransparentAclTemplate", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTransparentAclTemplate(id string, name string, host string) (*TransparentAclTemplate, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTransparentAclTemplate")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/transparent-acl-template/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TransparentAclTemplate
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTransparentAclTemplate REQ RES..........................", m)
			err := check_api_status("GetSlbTransparentAclTemplate", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
