package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type ServerSSH struct {
	UUID ServerSSHInstance `json:"server-ssh,omitempty"`
}
type SamplingEnable2 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type ServerSSHInstance struct {
	Counters1          []SamplingEnable2 `json:"sampling-enable,omitempty"`
	ForwardProxyEnable int               `json:"forward-proxy-enable,omitempty"`
	Name               string            `json:"name,omitempty"`
	UserTag            string            `json:"user-tag,omitempty"`
	UUID               string            `json:"uuid,omitempty"`
}

func PostSlbTemplateServerSSH(id string, inst ServerSSH, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateServerSSH")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/server-ssh", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSH
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateServerSSH REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateServerSSH", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateServerSSH(id string, name string, host string) (*ServerSSH, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateServerSSH")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/server-ssh/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSH
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateServerSSH REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateServerSSH", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateServerSSH(id string, name string, inst ServerSSH, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateServerSSH")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/server-ssh/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSH
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateServerSSH REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateServerSSH", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateServerSSH(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateServerSSH")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/server-ssh/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSH
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
