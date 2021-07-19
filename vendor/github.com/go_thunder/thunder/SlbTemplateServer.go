package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type TemplateServer struct {
	UUID TemplateServerInstance `json:"server,omitempty"`
}

type TemplateServerInstance struct {
	HealthCheckDisable     int    `json:"health-check-disable,omitempty"`
	StatsDataAction        string `json:"stats-data-action,omitempty"`
	SlowStart              int    `json:"slow-start,omitempty"`
	Weight                 int    `json:"weight,omitempty"`
	BwRateLimit            int    `json:"bw-rate-limit,omitempty"`
	SpoofingCache          int    `json:"spoofing-cache,omitempty"`
	ConnLimit              int    `json:"conn-limit,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
	Resume                 int    `json:"resume,omitempty"`
	MaxDynamicServer       int    `json:"max-dynamic-server,omitempty"`
	RateInterval           string `json:"rate-interval,omitempty"`
	Till                   int    `json:"till,omitempty"`
	Add                    int    `json:"add,omitempty"`
	MinTTLRatio            int    `json:"min-ttl-ratio,omitempty"`
	BwRateLimitNoLogging   int    `json:"bw-rate-limit-no-logging,omitempty"`
	DynamicServerPrefix    string `json:"dynamic-server-prefix,omitempty"`
	InitialSlowStart       int    `json:"initial-slow-start,omitempty"`
	Every                  int    `json:"every,omitempty"`
	ConnLimitNoLogging     int    `json:"conn-limit-no-logging,omitempty"`
	ExtendedStats          int    `json:"extended-stats,omitempty"`
	ConnRateLimitNoLogging int    `json:"conn-rate-limit-no-logging,omitempty"`
	Name                   string `json:"name,omitempty"`
	BwRateLimitDuration    int    `json:"bw-rate-limit-duration,omitempty"`
	BwRateLimitResume      int    `json:"bw-rate-limit-resume,omitempty"`
	BwRateLimitAcct        string `json:"bw-rate-limit-acct,omitempty"`
	UserTag                string `json:"user-tag,omitempty"`
	Times                  int    `json:"times,omitempty"`
	LogSelectionFailure    int    `json:"log-selection-failure,omitempty"`
	ConnRateLimit          int    `json:"conn-rate-limit,omitempty"`
	DNSQueryInterval       int    `json:"dns-query-interval,omitempty"`
	HealthCheck            string `json:"health-check,omitempty"`
}

func PostSlbTemplateServer(id string, inst TemplateServer, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TemplateServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateServer REQ RES..........................", m)
			check_api_status("PostSlbTemplateServer", data)

		}
	}

}

func GetSlbTemplateServer(id string, name string, host string) (*TemplateServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/server/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TemplateServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateServer REQ RES..........................", m)
			check_api_status("GetSlbTemplateServer", data)
			return &m, nil
		}
	}

}

func PutSlbTemplateServer(id string, name string, inst TemplateServer, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateServer")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/server/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TemplateServer
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateServer REQ RES..........................", m)
			check_api_status("PutSlbTemplateServer", data)

		}
	}

}

func DeleteSlbTemplateServer(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateServer")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/server/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TemplateServer
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
