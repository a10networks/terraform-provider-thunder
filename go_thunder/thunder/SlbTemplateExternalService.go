package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type ExternalService struct {
	UUID ExternalServiceInstance `json:"external-service,omitempty"`
}

type RequestHeaderForwardList struct {
	RequestHeaderForward string `json:"request-header-forward,omitempty"`
}

type BypassIPCfg1 struct {
	Mask     string `json:"mask,omitempty"`
	BypassIP string `json:"bypass-ip,omitempty"`
}
type ExternalServiceInstance struct {
	Name                                   string                     `json:"name,omitempty"`
	SharedPartitionPersistSourceIPTemplate int                        `json:"shared-partition-persist-source-ip-template,omitempty"`
	Type                                   string                     `json:"type,omitempty"`
	SourceIP                               string                     `json:"source-ip,omitempty"`
	RequestHeaderForward                   []RequestHeaderForwardList `json:"request-header-forward-list,omitempty"`
	TemplateTCPProxyShared                 string                     `json:"template-tcp-proxy-shared,omitempty"`
	UserTag                                string                     `json:"user-tag,omitempty"`
	SharedPartitionTCPProxyTemplate        int                        `json:"shared-partition-tcp-proxy-template,omitempty"`
	Action                                 string                     `json:"action,omitempty"`
	ServiceGroup                           string                     `json:"service-group,omitempty"`
	FailureAction                          string                     `json:"failure-action,omitempty"`
	Timeout                                int                        `json:"timeout,omitempty"`
	TCPProxy                               string                     `json:"tcp-proxy,omitempty"`
	TemplatePersistSourceIPShared          string                     `json:"template-persist-source-ip-shared,omitempty"`
	BypassIP                               []BypassIPCfg1             `json:"bypass-ip-cfg,omitempty"`
	UUID                                   string                     `json:"uuid,omitempty"`
}

func PostSlbTemplateExternalService(id string, inst ExternalService, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateExternalService")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/external-service", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ExternalService
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateExternalService REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateExternalService", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateExternalService(id string, name string, host string) (*ExternalService, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateExternalService")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/external-service/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ExternalService
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateExternalService REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateExternalService", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateExternalService(id string, name string, inst ExternalService, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateExternalService")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/external-service/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ExternalService
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateExternalService REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateExternalService", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateExternalService(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateExternalService")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/external-service/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ExternalService
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
