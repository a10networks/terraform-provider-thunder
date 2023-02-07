package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type IpTcp struct {
	UUID IpTcpInstance `json:"tcp,omitempty"`
}
type SynCookie struct {
	Threshold int `json:"threshold,omitempty"`
}
type IpTcpInstance struct {
	UUID      string    `json:"uuid,omitempty"`
	Threshold SynCookie `json:"syn-cookie,omitempty"`
}

func PostIpTcp(id string, inst IpTcp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpTcp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/tcp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpTcp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpTcp REQ RES..........................", m)
			err := check_api_status("PostIpTcp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpTcp(id string, host string) (*IpTcp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpTcp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/tcp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m IpTcp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpTcp REQ RES..........................", m)
			err := check_api_status("GetIpTcp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
