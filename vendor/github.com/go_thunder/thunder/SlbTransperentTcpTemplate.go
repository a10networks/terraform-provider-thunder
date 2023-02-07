package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type TransperentTcpTemplate struct {
	Name TransperentTcpTemplateInstance `json:"transparent-tcp-template,omitempty"`
}

type TransperentTcpTemplateInstance struct {
	Name string `json:"name,omitempty"`
}

func PostSlbTransperentTcpTemplate(id string, inst TransperentTcpTemplate, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTransperentTcpTemplate")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/transparent-tcp-template", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TransperentTcpTemplate
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTransperentTcpTemplate REQ RES..........................", m)
			err := check_api_status("PostSlbTransperentTcpTemplate", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTransperentTcpTemplate(id string, name string, host string) (*TransperentTcpTemplate, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTransperentTcpTemplate")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/transparent-tcp-template/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TransperentTcpTemplate
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTransperentTcpTemplate REQ RES..........................", m)
			err := check_api_status("GetSlbTransperentTcpTemplate", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
