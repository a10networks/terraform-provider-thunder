package thunder

import (
	"bytes"
	"github.com/clarketm/json"
	"io/ioutil"
	"util"
)

// based on ACOS 5_2_1-P3_70
type EndpointLoggingHostIpv6addr struct {
	Inst LoggingHostIpv6addr `json:"ipv6addr"`
}
type LoggingHostIpv6addr struct {
	HostIpv6    string `json:"host-ipv6,omitempty"`
	OverTls     int    `json:"over-tls,omitempty"`
	Port        int    `json:"port,omitempty"`
	Tcp         int    `json:"tcp,omitempty"`
	UseMgmtPort int    `json:"use-mgmt-port,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
}

func (p *EndpointLoggingHostIpv6addr) GetId() string {
	return p.Inst.HostIpv6
}

func PostEndpointLoggingHostIpv6addr(auth_token string, inst EndpointLoggingHostIpv6addr, host string) error {
	logger := util.GetLoggerInstance()
	logger.Println("PostLoggingHostIpv6addr")
	headers := GetReqHeader(auth_token)
	payloadBytes, err := json.Marshal(inst)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	resp, err := doHttp("POST", "https://"+host+"/axapi/v3/logging/host/ipv6addr", bytes.NewReader(payloadBytes), headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointLoggingHostIpv6addr
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("PostLoggingHostIpv6addr", data)
		}
	}
	return err
}

func GetEndpointLoggingHostIpv6addr(auth_token string, host string, instId string) (*EndpointLoggingHostIpv6addr, error) {
	logger := util.GetLoggerInstance()
	logger.Println("GetLoggingHostIpv6addr")
	headers := GetReqHeader(auth_token)
	resp, err := doHttp("GET", "https://"+host+"/axapi/v3/logging/host/ipv6addr/"+instId, nil, headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointLoggingHostIpv6addr
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("GetLoggingHostIpv6addr", data)
			if err == nil {
				return &m, nil
			}
		}
	}
	return nil, err
}

func PutEndpointLoggingHostIpv6addr(auth_token string, inst EndpointLoggingHostIpv6addr, host string) error {
	logger := util.GetLoggerInstance()
	logger.Println("PutLoggingHostIpv6addr")
	headers := GetReqHeader(auth_token)
	payloadBytes, err := json.Marshal(inst)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))

	resp, err := doHttp("PUT", "https://"+host+"/axapi/v3/logging/host/ipv6addr/"+inst.GetId(), bytes.NewReader(payloadBytes), headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", string(data))
		var m EndpointLoggingHostIpv6addr
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("PutLoggingHostIpv6addr", data)
		}
	}
	return err
}

func DeleteEndpointLoggingHostIpv6addr(auth_token string, host string, instId string) error {
	logger := util.GetLoggerInstance()
	logger.Println("DeleteLoggingHostIpv6addr")
	headers := GetReqHeader(auth_token)
	resp, err := doHttp("DELETE", "https://"+host+"/axapi/v3/logging/host/ipv6addr/"+instId, nil, headers)
	if err != nil {
		logger.Println("HTTP request failed with error", err)
	} else {
		logger.Println("HTTP status", resp.Status)
		data, _ := ioutil.ReadAll(resp.Body)
		logger.Println("response:", data)
		var m EndpointLoggingHostIpv6addr
		err = json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		} else {
			err = checkAxApiResponseStatus("DeleteLoggingHostIpv6addr", data)
		}
	}
	return err
}
