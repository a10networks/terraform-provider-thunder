package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FileCsr struct {
	FileCsrInstanceFile FileCsrInstance `json:"csr,omitempty"`
}

type FileCsrInstance struct {
	FileCsrInstanceAction     string `json:"action,omitempty"`
	FileCsrInstanceDstFile    string `json:"dst-file,omitempty"`
	FileCsrInstanceFile       string `json:"file,omitempty"`
	FileCsrInstanceFileHandle string `json:"file-handle,omitempty"`
	FileCsrInstanceSize       int    `json:"size,omitempty"`
	FileCsrInstanceUUID       string `json:"uuid,omitempty"`
}

func PostFileCsr(id string, inst FileCsr, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFileCsr")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/file/csr", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FileCsr
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostFileCsr", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFileCsr(id string, host string) (*FileCsr, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFileCsr")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/file/csr", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FileCsr
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetFileCsr", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
