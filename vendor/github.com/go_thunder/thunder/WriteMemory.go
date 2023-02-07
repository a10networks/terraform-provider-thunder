package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type WriteMemory struct {
	Profile WriteMemoryInstance `json:"memory,omitempty"`
}

type WriteMemoryInstance struct {
	Destination        string `json:"destination,omitempty"`
	Partition          string `json:"partition,omitempty"`
	Profile            string `json:"profile,omitempty"`
	SpecifiedPartition string `json:"specified-partition,omitempty"`
}

func PostWriteMemory(id string, inst WriteMemory, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostWriteMemory")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/write/memory", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WriteMemory
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostWriteMemory REQ RES..........................", m)
			err := check_api_status("PostWriteMemory", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetWriteMemory(id string, host string) (*WriteMemory, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetWriteMemory")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/write/memory/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WriteMemory
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GEGetWriteMemoryT REQ RES..........................", m)
			err := check_api_status("GetWriteMemory", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
