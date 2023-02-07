package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type ClientSSH struct {
	UUID ClientSSHInstance `json:"client-ssh,omitempty"`
}

type ClientSSHInstance struct {
	UUID                string `json:"uuid,omitempty"`
	Encrypted           string `json:"encrypted,omitempty"`
	UserTag             string `json:"user-tag,omitempty"`
	ForwardProxyEnable  int    `json:"forward-proxy-enable,omitempty"`
	Passphrase          string `json:"passphrase,omitempty"`
	ForwardProxyHostkey string `json:"forward-proxy-hostkey,omitempty"`
	Name                string `json:"name,omitempty"`
}

func PostTemplateClientSsh(id string, inst ClientSSH, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplateClientSsh")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/client-ssh", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSH
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostTemplateClientSsh REQ RES..........................", m)
			err := check_api_status("PostTemplateClientSsh", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetTemplateClientSsh(id string, name string, host string) (*ClientSSH, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplateClientSsh")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/client-ssh/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSH
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetTemplateClientSsh REQ RES..........................", m)
			err := check_api_status("GetTemplateClientSsh", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutTemplateClientSsh(id string, name string, inst ClientSSH, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplateClientSsh")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/client-ssh/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSH
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutTemplateClientSsh REQ RES..........................", m)
			err := check_api_status("PutTemplateClientSsh", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteTemplateClientSsh(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplateClientSsh")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/client-ssh/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClientSSH
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
