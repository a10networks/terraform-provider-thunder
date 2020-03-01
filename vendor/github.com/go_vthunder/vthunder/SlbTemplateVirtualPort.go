package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type VirtualPort struct {
	UUID VirtualPortInstance `json:"virtual-port,omitempty"`
}

type VirtualPortInstance struct {
	ResetUnknownConn       int    `json:"reset-unknown-conn,omitempty"`
	IgnoreTCPMsl           int    `json:"ignore-tcp-msl,omitempty"`
	Rate                   int    `json:"rate,omitempty"`
	SnatMsl                int    `json:"snat-msl,omitempty"`
	AllowSynOtherflags     int    `json:"allow-syn-otherflags,omitempty"`
	Aflow                  int    `json:"aflow,omitempty"`
	ConnLimit              int    `json:"conn-limit,omitempty"`
	DropUnknownConn        int    `json:"drop-unknown-conn,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
	ResetL7OnFailover      int    `json:"reset-l7-on-failover,omitempty"`
	PktRateType            string `json:"pkt-rate-type,omitempty"`
	RateInterval           string `json:"rate-interval,omitempty"`
	SnatPortPreserve       int    `json:"snat-port-preserve,omitempty"`
	ConnRateLimitReset     int    `json:"conn-rate-limit-reset,omitempty"`
	WhenRrEnable           int    `json:"when-rr-enable,omitempty"`
	NonSynInitiation       int    `json:"non-syn-initiation,omitempty"`
	ConnLimitReset         int    `json:"conn-limit-reset,omitempty"`
	Dscp                   int    `json:"dscp,omitempty"`
	PktRateLimitReset      int    `json:"pkt-rate-limit-reset,omitempty"`
	ConnLimitNoLogging     int    `json:"conn-limit-no-logging,omitempty"`
	ConnRateLimitNoLogging int    `json:"conn-rate-limit-no-logging,omitempty"`
	LogOptions             string `json:"log-options,omitempty"`
	Name                   string `json:"name,omitempty"`
	AllowVipToRportMapping int    `json:"allow-vip-to-rport-mapping,omitempty"`
	PktRateInterval        string `json:"pkt-rate-interval,omitempty"`
	UserTag                string `json:"user-tag,omitempty"`
	ConnRateLimit          int    `json:"conn-rate-limit,omitempty"`
}

func PostSlbTemplateVirtualPort(id string, inst VirtualPort, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateVirtualPort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/virtual-port", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetSlbTemplateVirtualPort(id string, name string, host string) (*VirtualPort, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateVirtualPort")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/virtual-port/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}

}

func PutSlbTemplateVirtualPort(id string, name string, inst VirtualPort, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateVirtualPort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/virtual-port/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualPort
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteSlbTemplateVirtualPort(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateVirtualPort")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/virtual-port/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualPort
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
