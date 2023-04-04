package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwClearSessionFilter struct {
	Status FwClearSessionFilterInstance `json:"clear-session-filter,omitempty"`
}

type FwClearSessionFilterInstance struct {
	Status string `json:"status,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

func PostFwClearSessionFilter(id string, inst FwClearSessionFilter, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwClearSessionFilter")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/clear-session-filter", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwClearSessionFilter
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] POST REQ RES..........................", m)
			err := check_api_status("PostFwClearSessionFilter", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetFwClearSessionFilter(id string, host string) (*FwClearSessionFilter, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwClearSessionFilter")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/clear-session-filter/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwClearSessionFilter
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetFwClearSessionFilter", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
