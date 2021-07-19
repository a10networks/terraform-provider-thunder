package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SnmpServerHostIpv6Host struct {
	UUID SnmpServerHostIpv6HostInstance `json:"ipv6-host,omitempty"`
}

type SnmpServerHostIpv6HostInstance struct {
	Ipv6Addr  string `json:"ipv6-addr,omitempty"`
	UDPPort   int    `json:"udp-port,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	User      string `json:"user,omitempty"`
	V1V2CComm string `json:"v1-v2c-comm,omitempty"`
	Version   string `json:"version,omitempty"`
}

func PostSnmpServerHostIpv6Host(id string, inst SnmpServerHostIpv6Host, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSnmpServerHostIpv6Host")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/snmp-server/host/ipv6-host", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv6Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			check_api_status("PostSnmpServerHostIpv6Host", data)

		}
	}

}

func GetSnmpServerHostIpv6Host(id string, name1 string, name2 string, host string) (*SnmpServerHostIpv6Host, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSnmpServerHostIpv6Host")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/snmp-server/host/ipv6-host/"+name1+"+"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv6Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			check_api_status("GetSnmpServerHostIpv6Host", data)
			return &m, nil
		}
	}

}

func PutSnmpServerHostIpv6Host(id string, name1 string, name2 string, inst SnmpServerHostIpv6Host, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSnmpServerHostIpv6Host")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/snmp-server/host/ipv6-host/"+name1+"+"+name2, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv6Host
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			check_api_status("PutSnmpServerHostIpv6Host", data)

		}
	}

}

func DeleteSnmpServerHostIpv6Host(id string, name1 string, name2 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSnmpServerHostIpv6Host")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/snmp-server/host/ipv6-host/"+name1+"+"+name2, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SnmpServerHostIpv6Host
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
