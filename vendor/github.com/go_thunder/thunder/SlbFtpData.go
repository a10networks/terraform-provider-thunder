package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type FTPData struct {
	UUID FTPDataInstance `json:"ftp-data,omitempty"`
}

type SamplingEnable21 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type FTPDataInstance struct {
	Counters1 []SamplingEnable21 `json:"sampling-enable,omitempty"`
	UUID      string             `json:"uuid,omitempty"`
}

func GetSlbFTPData(id string, host string) (*FTPData, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/ftp-data", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FTPData
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			check_api_status("GetSlbFTPData", data)
			return &m, nil
		}
	}
}

func PostSlbFTPData(id string, vc FTPData, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/ftp-data", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v FTPData
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			check_api_status("PostSlbFTPData", data)
		}
	}

}
