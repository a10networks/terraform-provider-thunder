package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwHelperSessions struct {
	IdleTimeout FwHelperSessionsInstance `json:"helper-sessions,omitempty"`
}

type FwHelperSessionsInstance struct {
	IdleTimeout int    `json:"idle-timeout,omitempty"`
	Limit       int    `json:"limit,omitempty"`
	Mode        string `json:"mode,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

func PostFwHelperSessions(id string, inst FwHelperSessions, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwHelperSessions")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/helper-sessions", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwHelperSessions
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwHelperSessions", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwHelperSessions(id string, host string) (*FwHelperSessions, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwHelperSessions")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/helper-sessions/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwHelperSessions
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwHelperSessions", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
