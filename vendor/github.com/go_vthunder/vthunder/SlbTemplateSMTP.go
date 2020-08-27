package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SMTP struct {
	UUID SMTPInstance `json:"smtp,omitempty"`
}

type ClientDomainSwitching struct {
	ServiceGroup  string `json:"service-group,omitempty"`
	MatchString   string `json:"match-string,omitempty"`
	SwitchingType string `json:"switching-type,omitempty"`
}
type Template1 struct {
	Logging string `json:"logging,omitempty"`
}
type CommandDisable struct {
	DisableType string `json:"disable-type,omitempty"`
}
type SMTPInstance struct {
	Name               string                  `json:"name,omitempty"`
	UserTag            string                  `json:"user-tag,omitempty"`
	ServerDomain       string                  `json:"server-domain,omitempty"`
	ServiceGroup       []ClientDomainSwitching `json:"client-domain-switching,omitempty"`
	ClientStarttlsType string                  `json:"client-starttls-type,omitempty"`
	Logging            Template1               `json:"template,omitempty"`
	DisableType        []CommandDisable        `json:"command-disable,omitempty"`
	ServerStarttlsType string                  `json:"server-starttls-type,omitempty"`
	ServiceReadyMsg    string                  `json:"service-ready-msg,omitempty"`
	UUID               string                  `json:"uuid,omitempty"`
}

func PostSlbTemplateSMTP(id string, inst SMTP, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateSMTP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/smtp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SMTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetSlbTemplateSMTP(id string, name string, host string) (*SMTP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateSMTP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/smtp/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SMTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}

func PutSlbTemplateSMTP(id string, name string, inst SMTP, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateSMTP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/smtp/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SMTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteSlbTemplateSMTP(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateSMTP")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/smtp/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SMTP
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
