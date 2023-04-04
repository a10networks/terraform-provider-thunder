package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwServer struct {
	FwServerInstanceName FwServerInstance `json:"server,omitempty"`
}

type FwServerInstance struct {
	FwServerInstanceAction                  string                           `json:"action,omitempty"`
	FwServerInstanceFqdnName                string                           `json:"fqdn-name,omitempty"`
	FwServerInstanceHealthCheck             string                           `json:"health-check,omitempty"`
	FwServerInstanceHealthCheckDisable      int                              `json:"health-check-disable,omitempty"`
	FwServerInstanceHost                    string                           `json:"host,omitempty"`
	FwServerInstanceName                    string                           `json:"name,omitempty"`
	FwServerInstancePortListPortNumber      []FwServerInstancePortList       `json:"port-list,omitempty"`
	FwServerInstanceResolveAs               string                           `json:"resolve-as,omitempty"`
	FwServerInstanceSamplingEnableCounters1 []FwServerInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	FwServerInstanceServerIpv6Addr          string                           `json:"server-ipv6-addr,omitempty"`
	FwServerInstanceUUID                    string                           `json:"uuid,omitempty"`
	FwServerInstanceUserTag                 string                           `json:"user-tag,omitempty"`
}

type FwServerInstancePortList struct {
	FwServerInstancePortListAction                  string                                   `json:"action,omitempty"`
	FwServerInstancePortListHealthCheck             string                                   `json:"health-check,omitempty"`
	FwServerInstancePortListHealthCheckDisable      int                                      `json:"health-check-disable,omitempty"`
	FwServerInstancePortListPacketCaptureTemplate   string                                   `json:"packet-capture-template,omitempty"`
	FwServerInstancePortListPortNumber              int                                      `json:"port-number,omitempty"`
	FwServerInstancePortListProtocol                string                                   `json:"protocol,omitempty"`
	FwServerInstancePortListSamplingEnableCounters1 []FwServerInstancePortListSamplingEnable `json:"sampling-enable,omitempty"`
	FwServerInstancePortListUUID                    string                                   `json:"uuid,omitempty"`
	FwServerInstancePortListUserTag                 string                                   `json:"user-tag,omitempty"`
}

type FwServerInstanceSamplingEnable struct {
	FwServerInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type FwServerInstancePortListSamplingEnable struct {
	FwServerInstancePortListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

func PostFwServer(id string, inst FwServer, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostFwServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwServer(id string, name1 string, host string) (*FwServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/server/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetFwServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutFwServer(id string, name1 string, inst FwServer, host string) error {

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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/fw/server/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutFwServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteFwServer(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteFwServer")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/fw/server/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteFwServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
