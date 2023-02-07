package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTcpMssClamp struct {
	MssSubtract FwTcpMssClampInstance `json:"mss-clamp,omitempty"`
}

type FwTcpMssClampInstance struct {
	Min          int    `json:"min,omitempty"`
	MssClampType string `json:"mss-clamp-type,omitempty"`
	MssSubtract  int    `json:"mss-subtract,omitempty"`
	MssValue     int    `json:"mss-value,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostFwTcpMssClamp(id string, inst FwTcpMssClamp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTcpMssClamp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/tcp/mss-clamp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpMssClamp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTcpMssClamp REQ RES..........................", m)
			err := check_api_status("PostFwTcpMssClamp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwTcpMssClamp(id string, host string) (*FwTcpMssClamp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTcpMssClamp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/tcp/mss-clamp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpMssClamp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwTcpMssClamp REQ RES..........................", m)
			err := check_api_status("GetFwTcpMssClamp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
