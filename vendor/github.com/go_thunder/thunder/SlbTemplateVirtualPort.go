package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateVirtualPort struct {
	SlbTemplateVirtualPortInstanceName SlbTemplateVirtualPortInstance `json:"virtual-port,omitempty"`
}

type SlbTemplateVirtualPortInstance struct {
	SlbTemplateVirtualPortInstanceAflow                  int    `json:"aflow,omitempty"`
	SlbTemplateVirtualPortInstanceAllowSynOtherflags     int    `json:"allow-syn-otherflags,omitempty"`
	SlbTemplateVirtualPortInstanceAllowVipToRportMapping int    `json:"allow-vip-to-rport-mapping,omitempty"`
	SlbTemplateVirtualPortInstanceConnLimit              int    `json:"conn-limit,omitempty"`
	SlbTemplateVirtualPortInstanceConnLimitNoLogging     int    `json:"conn-limit-no-logging,omitempty"`
	SlbTemplateVirtualPortInstanceConnLimitReset         int    `json:"conn-limit-reset,omitempty"`
	SlbTemplateVirtualPortInstanceConnRateLimit          int    `json:"conn-rate-limit,omitempty"`
	SlbTemplateVirtualPortInstanceConnRateLimitNoLogging int    `json:"conn-rate-limit-no-logging,omitempty"`
	SlbTemplateVirtualPortInstanceConnRateLimitReset     int    `json:"conn-rate-limit-reset,omitempty"`
	SlbTemplateVirtualPortInstanceDropUnknownConn        int    `json:"drop-unknown-conn,omitempty"`
	SlbTemplateVirtualPortInstanceDscp                   int    `json:"dscp,omitempty"`
	SlbTemplateVirtualPortInstanceIgnoreTCPMsl           int    `json:"ignore-tcp-msl,omitempty"`
	SlbTemplateVirtualPortInstanceLogOptions             string `json:"log-options,omitempty"`
	SlbTemplateVirtualPortInstanceName                   string `json:"name,omitempty"`
	SlbTemplateVirtualPortInstanceNonSynInitiation       int    `json:"non-syn-initiation,omitempty"`
	SlbTemplateVirtualPortInstancePktRateInterval        string `json:"pkt-rate-interval,omitempty"`
	SlbTemplateVirtualPortInstancePktRateLimitReset      int    `json:"pkt-rate-limit-reset,omitempty"`
	SlbTemplateVirtualPortInstancePktRateType            string `json:"pkt-rate-type,omitempty"`
	SlbTemplateVirtualPortInstanceRate                   int    `json:"rate,omitempty"`
	SlbTemplateVirtualPortInstanceRateInterval           string `json:"rate-interval,omitempty"`
	SlbTemplateVirtualPortInstanceResetL7OnFailover      int    `json:"reset-l7-on-failover,omitempty"`
	SlbTemplateVirtualPortInstanceResetUnknownConn       int    `json:"reset-unknown-conn,omitempty"`
	SlbTemplateVirtualPortInstanceSnatMsl                int    `json:"snat-msl,omitempty"`
	SlbTemplateVirtualPortInstanceSnatPortPreserve       int    `json:"snat-port-preserve,omitempty"`
	SlbTemplateVirtualPortInstanceUUID                   string `json:"uuid,omitempty"`
	SlbTemplateVirtualPortInstanceUserTag                string `json:"user-tag,omitempty"`
	SlbTemplateVirtualPortInstanceWhenRrEnable           int    `json:"when-rr-enable,omitempty"`
}

func PostSlbTemplateVirtualPort(id string, inst SlbTemplateVirtualPort, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/virtual-port", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualPort
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateVirtualPort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateVirtualPort(id string, name1 string, host string) (*SlbTemplateVirtualPort, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateVirtualPort")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/virtual-port/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualPort
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateVirtualPort", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateVirtualPort(id string, name1 string, inst SlbTemplateVirtualPort, host string) error {

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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/virtual-port/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualPort
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateVirtualPort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateVirtualPort(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateVirtualPort")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/virtual-port/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualPort
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateVirtualPort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
