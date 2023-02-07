package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateTcp struct {
	SlbTemplateTCPInstanceName SlbTemplateTCPInstance `json:"tcp,omitempty"`
}

type SlbTemplateTCPInstance struct {
	SlbTemplateTCPInstanceAliveIfActive           int    `json:"alive-if-active,omitempty"`
	SlbTemplateTCPInstanceDelSessionOnServerDown  int    `json:"del-session-on-server-down,omitempty"`
	SlbTemplateTCPInstanceDisable                 int    `json:"disable,omitempty"`
	SlbTemplateTCPInstanceDown                    int    `json:"down,omitempty"`
	SlbTemplateTCPInstanceForceDeleteTimeout      int    `json:"force-delete-timeout,omitempty"`
	SlbTemplateTCPInstanceForceDeleteTimeout100Ms int    `json:"force-delete-timeout-100ms,omitempty"`
	SlbTemplateTCPInstanceHalfCloseIdleTimeout    int    `json:"half-close-idle-timeout,omitempty"`
	SlbTemplateTCPInstanceHalfOpenIdleTimeout     int    `json:"half-open-idle-timeout,omitempty"`
	SlbTemplateTCPInstanceIdleTimeout             int    `json:"idle-timeout,omitempty"`
	SlbTemplateTCPInstanceInitialWindowSize       int    `json:"initial-window-size,omitempty"`
	SlbTemplateTCPInstanceInsertClientIP          int    `json:"insert-client-ip,omitempty"`
	SlbTemplateTCPInstanceLanFastAck              int    `json:"lan-fast-ack,omitempty"`
	SlbTemplateTCPInstanceLogging                 string `json:"logging,omitempty"`
	SlbTemplateTCPInstanceName                    string `json:"name,omitempty"`
	SlbTemplateTCPInstanceQos                     int    `json:"qos,omitempty"`
	SlbTemplateTCPInstanceReSelectIfServerDown    int    `json:"re-select-if-server-down,omitempty"`
	SlbTemplateTCPInstanceResetFollowFin          int    `json:"reset-follow-fin,omitempty"`
	SlbTemplateTCPInstanceResetFwd                int    `json:"reset-fwd,omitempty"`
	SlbTemplateTCPInstanceResetRev                int    `json:"reset-rev,omitempty"`
	SlbTemplateTCPInstanceUUID                    string `json:"uuid,omitempty"`
	SlbTemplateTCPInstanceUserTag                 string `json:"user-tag,omitempty"`
}

func PostSlbTemplateTcp(id string, inst SlbTemplateTcp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateTcp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/tcp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateTcp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateTcp(id string, name1 string, host string) (*SlbTemplateTcp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateTcp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/tcp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateTcp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateTcp(id string, name1 string, inst SlbTemplateTcp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateTcp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/tcp/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateTcp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateTcp(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateTcp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/tcp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateTcp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
