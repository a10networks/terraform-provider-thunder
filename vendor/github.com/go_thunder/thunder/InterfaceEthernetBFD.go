package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type EthernetBFD struct {
	UUID EthernetBFDInstance `json:"bfd,omitempty"`
}

type IntervalCfg struct {
	Interval   int `json:"interval,omitempty"`
	MinRx      int `json:"min-rx,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}
type Authentication struct {
	Encrypted string `json:"encrypted,omitempty"`
	Password  string `json:"password,omitempty"`
	Method    string `json:"method,omitempty"`
	KeyID     int    `json:"key-id,omitempty"`
}
type EthernetBFDInstance struct {
	Interval  IntervalCfg    `json:"interval-cfg,omitempty"`
	Encrypted Authentication `json:"authentication,omitempty"`
	Echo      int            `json:"echo,omitempty"`
	UUID      string         `json:"uuid,omitempty"`
	Demand    int            `json:"demand,omitempty"`
}

func PostInterfaceEthernetBFD(id string, name int, inst EthernetBFD, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernetBFD")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(name)+"/bfd", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetBFD
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceEthernetBFD REQ RES..........................", m)
			check_api_status("PostInterfaceEthernetBFD", data)

		}
	}

}

func GetInterfaceEthernetBFD(id string, name string, host string) (*EthernetBFD, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernetBFD")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name+"/bfd", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetBFD
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceEthernetBFD REQ RES..........................", m)
			check_api_status("GetInterfaceEthernetBFD", data)
			return &m, nil
		}
	}

}
