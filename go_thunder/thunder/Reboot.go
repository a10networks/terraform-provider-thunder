package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Reboot struct {
	All RebootInstance `json:"reboot,omitempty"`
}
type RebootInstance struct {
	All         int    `json:"all,omitempty"`
	DayOfMonth  int    `json:"day-of-month,omitempty"`
	Reason3     string `json:"reason-3,omitempty"`
	Reason2     string `json:"reason-2,omitempty"`
	In          string `json:"in,omitempty"`
	Month2      string `json:"month-2,omitempty"`
	Month       string `json:"month,omitempty"`
	Device      int    `json:"device,omitempty"`
	Reason      string `json:"reason,omitempty"`
	At          int    `json:"at,omitempty"`
	Time        string `json:"time,omitempty"`
	Cancel      int    `json:"cancel,omitempty"`
	DayOfMonth2 int    `json:"day-of-month-2,omitempty"`
}

func GetReboot(id string, host string) (*Reboot, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/reboot", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Reboot
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetReboot", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostReboot(id string, vc Reboot, host string) error {

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
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/reboot", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v Reboot
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostReboot", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
