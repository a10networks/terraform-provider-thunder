package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplatePersistSourceIp struct {
	Netmask6 SlbTemplatePersistSourceIPInstance `json:"source-ip,omitempty"`
}

type SlbTemplatePersistSourceIPInstance struct {
	DontHonorConnRules    int    `json:"dont-honor-conn-rules,omitempty"`
	EnforceHigherPriority int    `json:"enforce-higher-priority,omitempty"`
	HashPersist           int    `json:"hash-persist,omitempty"`
	InclDstIP             int    `json:"incl-dst-ip,omitempty"`
	InclSport             int    `json:"incl-sport,omitempty"`
	MatchType             int    `json:"match-type,omitempty"`
	Name                  string `json:"name,omitempty"`
	Netmask               string `json:"netmask,omitempty"`
	Netmask6              int    `json:"netmask6,omitempty"`
	PrimaryPort           int    `json:"primary-port,omitempty"`
	ScanAllMembers        int    `json:"scan-all-members,omitempty"`
	Server                int    `json:"server,omitempty"`
	ServiceGroup          int    `json:"service-group,omitempty"`
	Timeout               int    `json:"timeout,omitempty"`
	UUID                  string `json:"uuid,omitempty"`
	UserTag               string `json:"user-tag,omitempty"`
}

func PostSlbTemplatePersistSourceIp(id string, inst SlbTemplatePersistSourceIp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplatePersistSourceIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/persist/source-ip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistSourceIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplatePersistSourceIp REQ RES..........................", m)
			err := check_api_status("PostSlbTemplatePersistSourceIp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplatePersistSourceIp(id string, name string, host string) (*SlbTemplatePersistSourceIp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplatePersistSourceIp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/persist/source-ip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistSourceIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplatePersistSourceIp REQ RES..........................", m)
			err := check_api_status("GetSlbTemplatePersistSourceIp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplatePersistSourceIp(id string, name string, inst SlbTemplatePersistSourceIp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplatePersistSourceIp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/persist/source-ip/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistSourceIp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplatePersistSourceIp REQ RES..........................", m)
			err := check_api_status("PutSlbTemplatePersistSourceIp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplatePersistSourceIp(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplatePersistSourceIp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/persist/source-ip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistSourceIp
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
