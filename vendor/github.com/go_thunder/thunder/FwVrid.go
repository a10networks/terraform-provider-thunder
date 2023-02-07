package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwVrid struct {
	Vrid FwVridInstance `json:"vrid,omitempty"`
}

type FwVridInstance struct {
	UUID string `json:"uuid,omitempty"`
	Vrid int    `json:"vrid,omitempty"`
}

func PostFwVrid(id string, inst FwVrid, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwVrid")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/vrid", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwVrid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwVrid REQ RES..........................", m)
			err := check_api_status("PostFwVrid", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwVrid(id string, host string) (*FwVrid, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwVrid")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/vrid/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwVrid
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwVrid REQ RES..........................", m)
			err := check_api_status("GetFwVrid", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
