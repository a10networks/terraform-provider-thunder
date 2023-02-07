package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RespmodIcap struct {
	UUID RespmodIcapInstance `json:"respmod-icap,omitempty"`
}

type BypassIPCfg2 struct {
	Mask     string `json:"mask,omitempty"`
	BypassIP string `json:"bypass-ip,omitempty"`
}
type RespmodIcapInstance struct {
	MinPayloadSize                         int            `json:"min-payload-size,omitempty"`
	SharedPartitionPersistSourceIPTemplate int            `json:"shared-partition-persist-source-ip-template,omitempty"`
	TemplateTCPProxyShared                 string         `json:"template-tcp-proxy-shared,omitempty"`
	UUID                                   string         `json:"uuid,omitempty"`
	ServiceURL                             string         `json:"service-url,omitempty"`
	SharedPartitionTCPProxyTemplate        int            `json:"shared-partition-tcp-proxy-template,omitempty"`
	ServiceGroup                           string         `json:"service-group,omitempty"`
	TCPProxy                               string         `json:"tcp-proxy,omitempty"`
	Preview                                int            `json:"preview,omitempty"`
	DisableHTTPServerReset                 int            `json:"disable-http-server-reset,omitempty"`
	ServerSsl                              string         `json:"server-ssl,omitempty"`
	FailClose                              int            `json:"fail-close,omitempty"`
	Mask                                   []BypassIPCfg2 `json:"bypass-ip-cfg,omitempty"`
	TemplatePersistSourceIPShared          string         `json:"template-persist-source-ip-shared,omitempty"`
	IncludeProtocolInURI                   int            `json:"include-protocol-in-uri,omitempty"`
	Logging                                string         `json:"logging,omitempty"`
	Name                                   string         `json:"name,omitempty"`
	UserTag                                string         `json:"user-tag,omitempty"`
	LogOnlyAllowedMethod                   int            `json:"log-only-allowed-method,omitempty"`
	Action                                 string         `json:"action,omitempty"`
	Cylance                                int            `json:"cylance,omitempty"`
	SourceIP                               string         `json:"source-ip,omitempty"`
}

func PostSlbTemplateRespmodIcap(id string, inst RespmodIcap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateRespmodIcap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/respmod-icap", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RespmodIcap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateRespmodIcap REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateRespmodIcap", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateRespmodIcap(id string, name string, host string) (*RespmodIcap, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateRespmodIcap")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/respmod-icap/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RespmodIcap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateRespmodIcap REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateRespmodIcap", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateRespmodIcap(id string, name string, inst RespmodIcap, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateRespmodIcap")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/respmod-icap/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RespmodIcap
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateRespmodIcap REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateRespmodIcap", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateRespmodIcap(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateRespmodIcap")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/respmod-icap/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RespmodIcap
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
