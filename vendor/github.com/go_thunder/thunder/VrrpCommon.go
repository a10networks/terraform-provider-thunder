package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type HostidAppendToVrid struct {
	HostidAppendToVridValue   int `json:"hostid-append-to-vrid-value,omitempty"`
	HostidAppendToVridDefault int `json:"hostid-append-to-vrid-default"`
}
type InlineModeCfg struct {
	InlineMode     int `json:"inline-mode"`
	PreferredTrunk int `json:"preferred-trunk,omitempty"`
	PreferredPort  int `json:"preferred-port,omitempty"`
}
type Common struct {
	ForwardL4PacketOnStandby int                `json:"forward-l4-packet-on-standby"`
	GetReadyTime             int                `json:"get-ready-time,omitempty"`
	HelloInterval            int                `json:"hello-interval,omitempty"`
	UUID                     string             `json:"uuid,omitempty"`
	PreemptionDelay          int                `json:"preemption-delay,omitempty"`
	SetID                    int                `json:"set-id,omitempty"`
	DeviceID                 int                `json:"device-id,omitempty"`
	ArpRetry                 int                `json:"arp-retry,omitempty"`
	DeadTimer                int                `json:"dead-timer,omitempty"`
	DisableDefaultVrid       int                `json:"disable-default-vrid"`
	TrackEventDelay          int                `json:"track-event-delay,omitempty"`
	Action                   string             `json:"action,omitempty"`
	HostidAppendToVridValue  HostidAppendToVrid `json:"hostid-append-to-vrid,omitempty"`
	RestartTime              int                `json:"restart-time,omitempty"`
	InlineMode               InlineModeCfg      `json:"inline-mode-cfg,omitempty"`
}

type CommonInstance struct {
	UUID Common `json:"common,omitempty"`
}

func GetVrrpCommon(id string, host string) (*CommonInstance, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/vrrp-a/common", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CommonInstance
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("Common instance from Read")
			logger.Println(string(data))
			logger.Println("[INFO] GetVrrpCommon REQ RES..........................", m)
			err := check_api_status("GetVrrpCommon", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostVrrpCommon(id string, vc CommonInstance, host string) error {

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
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/vrrp-a/common", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v CommonInstance
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostVrrpCommon", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
