package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateUdp struct {
	SlbTemplateUDPInstanceName SlbTemplateUDPInstance `json:"udp,omitempty"`
}

type SlbTemplateUDPInstance struct {
	SlbTemplateUDPInstanceAge                    int    `json:"age,omitempty"`
	SlbTemplateUDPInstanceAvp                    string `json:"avp,omitempty"`
	SlbTemplateUDPInstanceDisableClearSession    int    `json:"disable-clear-session,omitempty"`
	SlbTemplateUDPInstanceIdleTimeout            int    `json:"idle-timeout,omitempty"`
	SlbTemplateUDPInstanceImmediate              int    `json:"immediate,omitempty"`
	SlbTemplateUDPInstanceName                   string `json:"name,omitempty"`
	SlbTemplateUDPInstanceQos                    int    `json:"qos,omitempty"`
	SlbTemplateUDPInstanceRadiusLbMethodHashType string `json:"radius-lb-method-hash-type,omitempty"`
	SlbTemplateUDPInstanceReSelectIfServerDown   int    `json:"re-select-if-server-down,omitempty"`
	SlbTemplateUDPInstanceShort                  int    `json:"short,omitempty"`
	SlbTemplateUDPInstanceStatelessConnTimeout   int    `json:"stateless-conn-timeout,omitempty"`
	SlbTemplateUDPInstanceUUID                   string `json:"uuid,omitempty"`
	SlbTemplateUDPInstanceUserTag                string `json:"user-tag,omitempty"`
}

func PostSlbTemplateUdp(id string, inst SlbTemplateUdp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateUdp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/udp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateUdp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateUdp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateUdp(id string, name1 string, host string) (*SlbTemplateUdp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateUdp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/udp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateUdp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateUdp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateUdp(id string, name1 string, inst SlbTemplateUdp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateUdp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/udp/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateUdp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateUdp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateUdp(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateUdp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/udp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateUdp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateUdp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
