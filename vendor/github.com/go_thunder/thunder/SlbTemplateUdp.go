package go_thunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type UDP struct {
	Name UDPInstance `json:"udp"`
}
type UDPInstance struct {
	Qos                  int    `json:"qos"`
	Name                 string `json:"name"`
	StatelessConnTimeout int    `json:"stateless-conn-timeout"`
	IdleTimeout          int    `json:"idle-timeout"`
	ReSelectIfServerDown int    `json:"re-select-if-server-down"`
	Immediate            int    `json:"immediate"`
	UserTag              string `json:"user-tag"`
}

func GetUdp(id string, name string, host string) (*UDP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/udp/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m UDP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetUdp REQ RES..........................", m)
			err := check_api_status("GetUdp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostUdp(id string, sg UDP, host string) error {

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
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/udp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m UDP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostUdp", data)
			if err != nil {
				return err
			}
		}
	}
return err
}

func PutUdp(id string, name string, sg UDP, host string) error {

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
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/udp/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m UDP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PutUdp", data)
			if err != nil {
				return err
			}
		}
	}
return err
}

func DeleteUdp(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/udp/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m UDP
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
