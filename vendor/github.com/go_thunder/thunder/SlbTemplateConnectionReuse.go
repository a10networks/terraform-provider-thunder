package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Connection_Reuse struct {
	Name ConnectionReuseInstance `json:"connection-reuse,omitempty"`
}

type ConnectionReuseInstance struct {
	Preopen        int    `json:"preopen,omitempty"`
	UUID           string `json:"uuid,omitempty"`
	KeepAliveConn  int    `json:"keep-alive-conn,omitempty"`
	UserTag        string `json:"user-tag,omitempty"`
	LimitPerServer int    `json:"limit-per-server,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	NumConnPerPort int    `json:"num-conn-per-port,omitempty"`
	Name           string `json:"name,omitempty"`
}

func PostTemplateConnectionReuse(id string, inst Connection_Reuse, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplateConnectionReuse")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/connection-reuse", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Connection_Reuse
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostTemplateConnectionReuse REQ RES..........................", m)
			err := check_api_status("PostTemplateConnectionReuse", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetTemplateConnectionReuse(id string, name string, host string) (*Connection_Reuse, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplateConnectionReuse")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/connection-reuse/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Connection_Reuse
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetTemplateConnectionReuse REQ RES..........................", m)
			err := check_api_status("GetTemplateConnectionReuse", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutTemplateConnectionReuse(id string, name string, inst Connection_Reuse, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplateConnectionReuse")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/connection-reuse/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Connection_Reuse
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutTemplateConnectionReuse REQ RES..........................", m)
			err := check_api_status("PutTemplateConnectionReuse", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteTemplateConnectionReuse(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplateConnectionReuse")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/connection-reuse/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Connection_Reuse
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
