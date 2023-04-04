package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type ReqmodIcap struct {
	UUID ReqmodIcapInstance `json:"reqmod-icap,omitempty"`
}
type BypassIPCfg struct {
	Mask     string `json:"mask,omitempty"`
	BypassIP string `json:"bypass-ip,omitempty"`
}
type ReqmodIcapInstance struct {
	MinPayloadSize                         int           `json:"min-payload-size,omitempty"`
	SharedPartitionPersistSourceIPTemplate int           `json:"shared-partition-persist-source-ip-template,omitempty"`
	TemplateTCPProxyShared                 string        `json:"template-tcp-proxy-shared,omitempty"`
	AllowedHTTPMethods                     string        `json:"allowed-http-methods,omitempty"`
	UUID                                   string        `json:"uuid,omitempty"`
	ServiceURL                             string        `json:"service-url,omitempty"`
	SharedPartitionTCPProxyTemplate        int           `json:"shared-partition-tcp-proxy-template,omitempty"`
	ServiceGroup                           string        `json:"service-group,omitempty"`
	TCPProxy                               string        `json:"tcp-proxy,omitempty"`
	Preview                                int           `json:"preview,omitempty"`
	DisableHTTPServerReset                 int           `json:"disable-http-server-reset,omitempty"`
	ServerSsl                              string        `json:"server-ssl,omitempty"`
	FailClose                              int           `json:"fail-close,omitempty"`
	Mask                                   []BypassIPCfg `json:"bypass-ip-cfg,omitempty"`
	TemplatePersistSourceIPShared          string        `json:"template-persist-source-ip-shared,omitempty"`
	IncludeProtocolInURI                   int           `json:"include-protocol-in-uri,omitempty"`
	Logging                                string        `json:"logging,omitempty"`
	Name                                   string        `json:"name,omitempty"`
	UserTag                                string        `json:"user-tag,omitempty"`
	LogOnlyAllowedMethod                   int           `json:"log-only-allowed-method,omitempty"`
	Action                                 string        `json:"action,omitempty"`
	Cylance                                int           `json:"cylance,omitempty"`
	SourceIP                               string        `json:"source-ip,omitempty"`
}

func PostSlbTemplateReqmodIcap(id string, inst ReqmodIcap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateReqmodIcap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/reqmod-icap", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ReqmodIcap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateReqmodIcap REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateReqmodIcap", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateReqmodIcap(id string, name string, host string) (*ReqmodIcap, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateReqmodIcap")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/reqmod-icap/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ReqmodIcap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateReqmodIcap REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateReqmodIcap", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateReqmodIcap(id string, name string, inst ReqmodIcap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateReqmodIcap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/reqmod-icap/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ReqmodIcap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateReqmodIcap", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateReqmodIcap(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateReqmodIcap")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/reqmod-icap/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ReqmodIcap
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
