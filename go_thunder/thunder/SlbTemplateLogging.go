package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Logging struct {
	Name LoggingInstance `json:"logging,omitempty"`
}
type LoggingInstance struct {
	PoolShared                      string `json:"pool-shared,omitempty"`
	Name                            string `json:"name,omitempty"`
	Format                          string `json:"format,omitempty"`
	Auto                            string `json:"auto,omitempty"`
	KeepEnd                         int    `json:"keep-end,omitempty"`
	LocalLogging                    int    `json:"local-logging,omitempty"`
	Mask                            string `json:"mask,omitempty"`
	TemplateTCPProxyShared          string `json:"template-tcp-proxy-shared,omitempty"`
	SharedPartitionTCPProxyTemplate int    `json:"shared-partition-tcp-proxy-template,omitempty"`
	KeepStart                       int    `json:"keep-start,omitempty"`
	ServiceGroup                    string `json:"service-group,omitempty"`
	PcreMask                        string `json:"pcre-mask,omitempty"`
	UserTag                         string `json:"user-tag,omitempty"`
	TCPProxy                        string `json:"tcp-proxy,omitempty"`
	SharedPartitionPool             int    `json:"shared-partition-pool,omitempty"`
	Pool                            string `json:"pool,omitempty"`
	UUID                            string `json:"uuid,omitempty"`
}

func GetLogging(id string, name string, host string) (*Logging, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/logging/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Logging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetLogging REQ RES..........................", m)
			err := check_api_status("GetLogging", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostLogging(id string, sg Logging, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/logging", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Logging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostLogging", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func PutLogging(id string, name string, sg Logging, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/logging/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Logging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PutLogging", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func DeleteLogging(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/logging/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Logging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
