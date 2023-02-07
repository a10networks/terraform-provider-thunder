package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SNMP struct {
	UUID SNMPInstance `json:"snmp,omitempty"`
}

type SNMPInstance struct {
	Username         string `json:"username,omitempty"`
	Oid              string `json:"oid,omitempty"`
	PrivProto        string `json:"priv-proto,omitempty"`
	UUID             string `json:"uuid,omitempty"`
	ContextName      string `json:"context-name,omitempty"`
	AuthKey          string `json:"auth-key,omitempty"`
	Interval         int    `json:"interval,omitempty"`
	ContextEngineID  string `json:"context-engine-id,omitempty"`
	SecurityLevel    string `json:"security-level,omitempty"`
	Community        string `json:"community,omitempty"`
	AuthProto        string `json:"auth-proto,omitempty"`
	Host             string `json:"host,omitempty"`
	Version          string `json:"version,omitempty"`
	UserTag          string `json:"user-tag,omitempty"`
	Interface        int    `json:"interface,omitempty"`
	PrivKey          string `json:"priv-key,omitempty"`
	SecurityEngineID string `json:"security-engine-id,omitempty"`
	Port             int    `json:"port,omitempty"`
	SnmpName         string `json:"snmp-name,omitempty"`
}

func PostSlbTemplateSNMP(id string, inst SNMP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateSNMP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/snmp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SNMP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateSNMP REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateSNMP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateSNMP(id string, snmp_name string, host string) (*SNMP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateSNMP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/snmp/"+snmp_name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SNMP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateSNMP REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateSNMP", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateSNMP(id string, snmp_name string, inst SNMP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateSNMP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/snmp/"+snmp_name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SNMP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateSNMP REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateSNMP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateSNMP(id string, snmp_name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateSNMP")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/snmp/"+snmp_name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SNMP
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
