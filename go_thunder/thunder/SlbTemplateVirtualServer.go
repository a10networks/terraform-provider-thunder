package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateVirtualServer struct {
	SlbTemplateVirtualServerInstanceName SlbTemplateVirtualServerInstance `json:"virtual-server,omitempty"`
}

type SlbTemplateVirtualServerInstance struct {
	SlbTemplateVirtualServerInstanceConnLimit                  int    `json:"conn-limit,omitempty"`
	SlbTemplateVirtualServerInstanceConnLimitNoLogging         int    `json:"conn-limit-no-logging,omitempty"`
	SlbTemplateVirtualServerInstanceConnLimitReset             int    `json:"conn-limit-reset,omitempty"`
	SlbTemplateVirtualServerInstanceConnRateLimit              int    `json:"conn-rate-limit,omitempty"`
	SlbTemplateVirtualServerInstanceConnRateLimitNoLogging     int    `json:"conn-rate-limit-no-logging,omitempty"`
	SlbTemplateVirtualServerInstanceConnRateLimitReset         int    `json:"conn-rate-limit-reset,omitempty"`
	SlbTemplateVirtualServerInstanceDisableWhenAllPortsDown    int    `json:"disable-when-all-ports-down,omitempty"`
	SlbTemplateVirtualServerInstanceDisableWhenAnyPortDown     int    `json:"disable-when-any-port-down,omitempty"`
	SlbTemplateVirtualServerInstanceIcmpLockup                 int    `json:"icmp-lockup,omitempty"`
	SlbTemplateVirtualServerInstanceIcmpLockupPeriod           int    `json:"icmp-lockup-period,omitempty"`
	SlbTemplateVirtualServerInstanceIcmpRateLimit              int    `json:"icmp-rate-limit,omitempty"`
	SlbTemplateVirtualServerInstanceIcmpv6Lockup               int    `json:"icmpv6-lockup,omitempty"`
	SlbTemplateVirtualServerInstanceIcmpv6LockupPeriod         int    `json:"icmpv6-lockup-period,omitempty"`
	SlbTemplateVirtualServerInstanceIcmpv6RateLimit            int    `json:"icmpv6-rate-limit,omitempty"`
	SlbTemplateVirtualServerInstanceName                       string `json:"name,omitempty"`
	SlbTemplateVirtualServerInstanceRateInterval               string `json:"rate-interval,omitempty"`
	SlbTemplateVirtualServerInstanceSubnetGratuitousArp        int    `json:"subnet-gratuitous-arp,omitempty"`
	SlbTemplateVirtualServerInstanceTCPStackTfoActiveConnLimit int    `json:"tcp-stack-tfo-active-conn-limit,omitempty"`
	SlbTemplateVirtualServerInstanceTCPStackTfoBackoffTime     int    `json:"tcp-stack-tfo-backoff-time,omitempty"`
	SlbTemplateVirtualServerInstanceTCPStackTfoCookieTimeLimit int    `json:"tcp-stack-tfo-cookie-time-limit,omitempty"`
	SlbTemplateVirtualServerInstanceUUID                       string `json:"uuid,omitempty"`
	SlbTemplateVirtualServerInstanceUserTag                    string `json:"user-tag,omitempty"`
}

func PostSlbTemplateVirtualServer(id string, inst SlbTemplateVirtualServer, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateVirtualServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/virtual-server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateVirtualServer(id string, name1 string, host string) (*SlbTemplateVirtualServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateVirtualServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/virtual-server/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateVirtualServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateVirtualServer(id string, name1 string, inst SlbTemplateVirtualServer, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateVirtualServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/virtual-server/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateVirtualServer(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateVirtualServer")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/virtual-server/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateVirtualServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
