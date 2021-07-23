package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type VirtualServer struct {
	UUID VirtualServerInstance `json:"virtual-server,omitempty"`
}

type VirtualServerInstance struct {
	ConnLimit                  int    `json:"conn-limit,omitempty"`
	ConnRateLimitNoLogging     int    `json:"conn-rate-limit-no-logging,omitempty"`
	Name                       string `json:"name,omitempty"`
	IcmpLockupPeriod           int    `json:"icmp-lockup-period,omitempty"`
	ConnLimitReset             int    `json:"conn-limit-reset,omitempty"`
	RateInterval               string `json:"rate-interval,omitempty"`
	UserTag                    string `json:"user-tag,omitempty"`
	Icmpv6RateLimit            int    `json:"icmpv6-rate-limit,omitempty"`
	SubnetGratuitousArp        int    `json:"subnet-gratuitous-arp,omitempty"`
	Icmpv6Lockup               int    `json:"icmpv6-lockup,omitempty"`
	ConnRateLimitReset         int    `json:"conn-rate-limit-reset,omitempty"`
	TCPStackTfoBackoffTime     int    `json:"tcp-stack-tfo-backoff-time,omitempty"`
	TCPStackTfoCookieTimeLimit int    `json:"tcp-stack-tfo-cookie-time-limit,omitempty"`
	ConnLimitNoLogging         int    `json:"conn-limit-no-logging,omitempty"`
	Icmpv6LockupPeriod         int    `json:"icmpv6-lockup-period,omitempty"`
	ConnRateLimit              int    `json:"conn-rate-limit,omitempty"`
	TCPStackTfoActiveConnLimit int    `json:"tcp-stack-tfo-active-conn-limit,omitempty"`
	IcmpLockup                 int    `json:"icmp-lockup,omitempty"`
	IcmpRateLimit              int    `json:"icmp-rate-limit,omitempty"`
	UUID                       string `json:"uuid,omitempty"`
}

func PostSlbTemplateVirtualServer(id string, inst VirtualServer, host string) error {

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
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/virtual-server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateVirtualServer REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetSlbTemplateVirtualServer(id string, name string, host string) (*VirtualServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateVirtualServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/virtual-server/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateVirtualServer REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateVirtualServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateVirtualServer(id string, name string, inst VirtualServer, host string) error {

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
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/virtual-server/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateVirtualServer REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateVirtualServer", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteSlbTemplateVirtualServer(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateVirtualServer")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/virtual-server/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m VirtualServer
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
