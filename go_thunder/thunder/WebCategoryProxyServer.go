package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type WebCategoryProxyServer struct {
	ProxyHost WebCategoryProxyServerInstance `json:"proxy-server,omitempty"`
}

type WebCategoryProxyServerInstance struct {
	AuthType     string `json:"auth-type,omitempty"`
	Domain       string `json:"domain,omitempty"`
	HTTPPort     int    `json:"http-port,omitempty"`
	HTTPSPort    int    `json:"https-port,omitempty"`
	Password     int    `json:"password,omitempty"`
	ProxyHost    string `json:"proxy-host,omitempty"`
	SecretString string `json:"secret-string,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Username     string `json:"username,omitempty"`
}

func PostWebCategoryProxyServer(id string, inst WebCategoryProxyServer, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostWebCategoryProxyServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/web-category/proxy-server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryProxyServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostWebCategoryProxyServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetWebCategoryProxyServer(id string, host string) (*WebCategoryProxyServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetWebCategoryProxyServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/web-category/proxy-server", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m WebCategoryProxyServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetWebCategoryProxyServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
