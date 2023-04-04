package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Sync struct {
	AllPartitions SyncInstance `json:"sync,omitempty"`
}
type SyncInstance struct {
	AllPartitions      int    `json:"all-partitions,omitempty"`
	PrivateKey         string `json:"private-key,omitempty"`
	PartitionName      string `json:"partition-name,omitempty"`
	Pwd                string `json:"pwd,omitempty"`
	AutoAuthentication int    `json:"auto-authentication,omitempty"`
	Address            string `json:"address,omitempty"`
	Shared             int    `json:"shared,omitempty"`
	Type               string `json:"type,omitempty"`
	PwdEnc             string `json:"pwd-enc,omitempty"`
	Usr                string `json:"usr,omitempty"`
}

func GetConfigureSync(id string, host string) (*Sync, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/configure/sync", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Sync
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetConfigureSync", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostConfigureSync(id string, vc Sync, host string) error {

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
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/configure/sync", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v Sync
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostConfigureSync", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
