package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type DynamicService struct {
	UUID DynamicServiceInstance `json:"dynamic-service,omitempty"`
}

type DNSServer struct {
	Ipv6DNSServer string `json:"ipv6-dns-server,omitempty"`
	Ipv4DNSServer string `json:"ipv4-dns-server,omitempty"`
}
type DynamicServiceInstance struct {
	Ipv6DNSServer []DNSServer `json:"dns-server,omitempty"`
	Name          string      `json:"name,omitempty"`
	UserTag       string      `json:"user-tag,omitempty"`
	UUID          string      `json:"uuid,omitempty"`
}

func PostSlbTemplateDynamicService(id string, inst DynamicService, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateDynamicService")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/dynamic-service", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DynamicService
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateDynamicService REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateDynamicService", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateDynamicService(id string, name string, host string) (*DynamicService, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateDynamicService")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/dynamic-service/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DynamicService
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateDynamicService REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateDynamicService", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateDynamicService(id string, name string, inst DynamicService, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateDynamicService")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/dynamic-service/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DynamicService
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateDynamicService REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateDynamicService", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateDynamicService(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateDynamicService")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/dynamic-service/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DynamicService
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
