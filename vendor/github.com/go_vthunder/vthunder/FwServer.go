package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type FwServer struct {
	UUID FwServerInstance `json:"server-instance,omitempty"`
}

type FwServerInstance struct {
	Action             string                   `json:"action,omitempty"`
	FqdnName           string                   `json:"fqdn-name,omitempty"`
	HealthCheck        string                   `json:"health-check,omitempty"`
	HealthCheckDisable int                      `json:"health-check-disable,omitempty"`
	Host               string                   `json:"host,omitempty"`
	Name               string                   `json:"name,omitempty"`
	Protocol           []FwServerPortList       `json:"port-list,omitempty"`
	ResolveAs          string                   `json:"resolve-as,omitempty"`
	Counters1          []FwServerSamplingEnable `json:"sampling-enable,omitempty"`
	ServerIpv6Addr     string                   `json:"server-ipv6-addr,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	UserTag            string                   `json:"user-tag,omitempty"`
}

type FwServerPortList struct {
	Action             string                   `json:"action,omitempty"`
	HealthCheck        string                   `json:"health-check,omitempty"`
	HealthCheckDisable int                      `json:"health-check-disable,omitempty"`
	PortNumber         int                      `json:"port-number,omitempty"`
	Protocol           string                   `json:"protocol,omitempty"`
	Counters1          []FwServerSamplingEnable `json:"sampling-enable,omitempty"`
	UUID               string                   `json:"uuid,omitempty"`
	UserTag            string                   `json:"user-tag,omitempty"`
}

type FwServerSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

func PostFwServer(id string, inst FwServer, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetFwServer(id string, name string, host string) (*FwServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/server/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}

func PutFwServer(id string, name string, inst FwServer, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutFwServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/fw/server/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteFwServer(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteFwServer")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/fw/server/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
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
