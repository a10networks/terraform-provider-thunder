package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type ResourceUsage struct {
	SlbThresholdResUsagePercent ResourceUsageInstance `json:"resource-usage,omitempty"`
}

type ResourceUsageInstance struct {
	SlbThresholdResUsagePercent int `json:"slb-threshold-res-usage-percent,omitempty"`
	ProxyTemplateCount          int `json:"proxy-template-count,omitempty"`
	NatPoolAddrCount            int `json:"nat-pool-addr-count,omitempty"`
	FastTCPTemplateCount        int `json:"fast-tcp-template-count,omitempty"`
	CacheTemplateCount          int `json:"cache-template-count,omitempty"`
	HealthMonitorCount          int `json:"health-monitor-count,omitempty"`
	FastUDPTemplateCount        int `json:"fast-udp-template-count,omitempty"`
	VirtualPortCount            int `json:"virtual-port-count,omitempty"`
	ClientSslTemplateCount      int `json:"client-ssl-template-count,omitempty"`
	PersistSrcipTemplateCount   int `json:"persist-srcip-template-count,omitempty"`
	ServerSslTemplateCount      int `json:"server-ssl-template-count,omitempty"`
	HTTPTemplateCount           int `json:"http-template-count,omitempty"`
	PbslbSubnetCount            int `json:"pbslb-subnet-count,omitempty"`
	PersistCookieTemplateCount  int `json:"persist-cookie-template-count,omitempty"`
	VirtualServerCount          int `json:"virtual-server-count,omitempty"`
	StreamTemplateCount         int `json:"stream-template-count,omitempty"`
	ConnReuseTemplateCount      int `json:"conn-reuse-template-count,omitempty"`
	RealServerCount             int `json:"real-server-count,omitempty"`
	RealPortCount               int `json:"real-port-count,omitempty"`
	ServiceGroupCount           int `json:"service-group-count,omitempty"`
}

func PostSlbResourceUsage(id string, inst ResourceUsage, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbResourceUsage")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/resource-usage", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ResourceUsage
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbResourceUsage REQ RES..........................", m)
			check_api_status("PostSlbResourceUsage", data)

		}
	}

}

func GetSlbResourceUsage(id string, host string) (*ResourceUsage, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbResourceUsage")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/resource-usage/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ResourceUsage
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbResourceUsage REQ RES..........................", m)
			check_api_status("GetSlbResourceUsage", data)
			return &m, nil
		}
	}

}
