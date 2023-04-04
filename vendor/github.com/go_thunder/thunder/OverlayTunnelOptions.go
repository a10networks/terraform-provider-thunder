package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Options struct {
	UUID OptionsInstance `json:"options,omitempty"`
}
type OptionsInstance struct {
	NvgreKeyModeLower24 int    `json:"nvgre-key-mode-lower24,omitempty"`
	UUID                string `json:"uuid,omitempty"`
	TCPMssAdjustDisable int    `json:"tcp-mss-adjust-disable,omitempty"`
	GatewayMac          string `json:"gateway-mac,omitempty"`
	IPDscpPreserve      int    `json:"ip-dscp-preserve,omitempty"`
	NvgreDisableFlowID  int    `json:"nvgre-disable-flow-id,omitempty"`
	VxlanDestPort       int    `json:"vxlan-dest-port,omitempty"`
}

func GetOptions(id string, host string) (*Options, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/overlay-tunnel/options", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Options
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("Options instance from Read")
			logger.Println(string(data))
			logger.Println("[INFO] GetOptions REQ RES..........................", m)
			err := check_api_status("GetOptions", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostOptions(id string, vc Options, host string) error {

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
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/overlay-tunnel/options", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v Options
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostOptions", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
