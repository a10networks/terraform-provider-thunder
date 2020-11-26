package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerHostIpv4Host struct {
	UUID SnmpServerHostIpv4HostInstance `json:"ipv4-host,omitempty"`
}

type SnmpServerHostIpv4HostInstance struct {
	Ipv4Addr  string `json:"ipv4-addr,omitempty"`
	UDPPort   int    `json:"udp-port,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	User      string `json:"user,omitempty"`
	V1V2CComm string `json:"v1-v2c-comm,omitempty"`
	Version   string `json:"version,omitempty"`
}

func PostSnmpServerHostIpv4Host(id string, inst SnmpServerHostIpv4Host, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerHostIpv4Host")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/host/ipv4-host", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv4Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSnmpServerHostIpv4Host", data)

		}
	}

}

func GetSnmpServerHostIpv4Host(id string, name1 string, name2 string, host string) (*SnmpServerHostIpv4Host, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerHostIpv4Host")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/host/ipv4-host/"+name1+"+"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv4Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSnmpServerHostIpv4Host", data)
			return &m, nil
		}
	}

}

func PutSnmpServerHostIpv4Host(id string, name1 string, name2 string, inst SnmpServerHostIpv4Host, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerHostIpv4Host")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/host/ipv4-host/"+name1+"+"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv4Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			check_api_status("PutSnmpServerHostIpv4Host", data)

		}
	}

}

func DeleteSnmpServerHostIpv4Host(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerHostIpv4Host")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/host/ipv4-host/"+name1+"+"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv4Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)

		}
	}
	return nil
}
