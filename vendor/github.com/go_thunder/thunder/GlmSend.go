package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type GlmSend struct {
	LicenseRequest GlmSendInstance `json:"send,omitempty"`
}

type GlmSendInstance struct {
	LicenseRequest int `json:"license-request,omitempty"`
}

func PostGlmSend(id string, inst GlmSend, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostGlmSend")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/glm/send", bytes.NewReader(payloadBytes), headers)
	logger.Println("post resp --> ", resp)
	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("post - data -> ", data)
		return err
		var m GlmSend
		erro := json.Unmarshal(data, &m)

		if erro != nil {
			logger.Println("Unmarshal error ", erro)

		} else {
			logger.Println("[INFO] PostGlmSend REQ RES..........................", m)
			err := check_api_status("PostGlmSend", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetGlmSend(id string, host string) (*GlmSend, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetGlmSend")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/glm/send/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)

		var m GlmSend
		erro := json.Unmarshal(data, &m)

		if erro != nil {
			logger.Println("Unmarshal error ", erro)
			return nil, err
		} else {
			logger.Println("[INFO] GetGlmSend REQ RES..........................", m)
			err := check_api_status("GetGlmSend", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
