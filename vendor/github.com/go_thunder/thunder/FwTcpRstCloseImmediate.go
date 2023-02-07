package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTcpRstCloseImmediate struct {
	Status FwTCPRstCloseImmediateInstance `json:"tcp-rst-close-immediate,omitempty"`
}

type FwTCPRstCloseImmediateInstance struct {
	Status string `json:"status,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

func PostFwTcpRstCloseImmediate(id string, inst FwTcpRstCloseImmediate, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTcpRstCloseImmediate")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/tcp-rst-close-immediate", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpRstCloseImmediate
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTcpRstCloseImmediate REQ RES..........................", m)
			err := check_api_status("PostFwTcpRstCloseImmediate", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwTcpRstCloseImmediate(id string, host string) (*FwTcpRstCloseImmediate, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTcpRstCloseImmediate")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/tcp-rst-close-immediate/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpRstCloseImmediate
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwTcpRstCloseImmediate REQ RES..........................", m)
			err := check_api_status("GetFwTcpRstCloseImmediate", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
