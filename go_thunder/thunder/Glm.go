package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Glm struct {
	UUID GlmInstance `json:"glm,omitempty"`
}

type GlmInstance struct {
	AllocateBandwidth int         `json:"allocate-bandwidth,omitempty"`
	ApplianceName     string      `json:"appliance-name,omitempty"`
	Burst             int         `json:"burst,omitempty"`
	EnableRequests    int         `json:"enable-requests,omitempty"`
	Enterprise        string      `json:"enterprise,omitempty"`
	Interval          int         `json:"interval,omitempty"`
	Port              int         `json:"port,omitempty"`
	Username          ProxyServer `json:"proxy-server,omitempty"`
	Token             string      `json:"token,omitempty"`
	UUID              string      `json:"uuid,omitempty"`
	UseMgmtPort       int         `json:"use-mgmt-port,omitempty"`
}

type ProxyServer struct {
	Encrypted    string `json:"encrypted,omitempty"`
	Host         string `json:"host,omitempty"`
	Password     int    `json:"password,omitempty"`
	Port         int    `json:"port,omitempty"`
	SecretString string `json:"secret-string,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Username     string `json:"username,omitempty"`
}

func GetGlm(id string, host string) (*Glm, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/glm", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Glm
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GetGlm REQ RES..........................", m)
			err := check_api_status("GetGlm", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostGlm(id string, vc Glm, host string) error {

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
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/glm", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v Glm
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostGlm", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
