package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type AdminSession struct {
	AdminSessionInstanceClear AdminSessionInstance `json:"admin-session,omitempty"`
}

type AdminSessionInstance struct {
	AdminSessionInstanceAll   int    `json:"all,omitempty"`
	AdminSessionInstanceClear int    `json:"clear,omitempty"`
	AdminSessionInstanceSid   int    `json:"sid,omitempty"`
	AdminSessionInstanceUUID  string `json:"uuid,omitempty"`
}

func PostAdminSession(id string, inst AdminSession, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostAdminSession")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/admin-session", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AdminSession
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostAdminSession", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetAdminSession(id string, host string) (*AdminSession, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetAdminSession")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/admin-session", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m AdminSession
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetAdminSession", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
