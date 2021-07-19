package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type IpNatPool struct {
	UseIfIP IpNatPoolInstance `json:"pool,omitempty"`
}

type IpNatPoolInstance struct {
	ChunkNetmask     string `json:"chunk-netmask,omitempty"`
	EndAddress       string `json:"end-address,omitempty"`
	Ethernet         int    `json:"ethernet,omitempty"`
	Gateway          string `json:"gateway,omitempty"`
	IPRr             int    `json:"ip-rr,omitempty"`
	Netmask          string `json:"netmask,omitempty"`
	PoolName         string `json:"pool-name,omitempty"`
	PortOverload     int    `json:"port-overload,omitempty"`
	ScaleoutDeviceID int    `json:"scaleout-device-id,omitempty"`
	StartAddress     string `json:"start-address,omitempty"`
	UUID             string `json:"uuid,omitempty"`
	UseIfIP          int    `json:"use-if-ip,omitempty"`
	Vrid             int    `json:"vrid,omitempty"`
}

func PostIpNatPool(id string, inst IpNatPool, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpNatPool")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/nat/pool", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpNatPool
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpNatPool REQ RES..........................", m)
			check_api_status("PostIpNatPool", data)

		}
	}

}

func GetIpNatPool(id string, name string, host string) (*IpNatPool, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpNatPool")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/nat/pool/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpNatPool
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpNatPool REQ RES..........................", m)
			check_api_status("GetIpNatPool", data)
			return &m, nil
		}
	}

}

func PutIpNatPool(id string, name string, inst IpNatPool, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutIpNatPool")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/ip/nat/pool/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpNatPool
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutIpNatPool REQ RES..........................", m)
			check_api_status("PutIpNatPool", data)

		}
	}

}

func DeleteIpNatPool(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteIpNatPool")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/ip/nat/pool/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpNatPool
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
