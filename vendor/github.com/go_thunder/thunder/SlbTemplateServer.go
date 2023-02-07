package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateServer struct {
	SlbTemplateServerInstanceName SlbTemplateServerInstance `json:"server,omitempty"`
}

type SlbTemplateServerInstance struct {
	SlbTemplateServerInstanceAdd                    int    `json:"add,omitempty"`
	SlbTemplateServerInstanceBwRateLimit            int    `json:"bw-rate-limit,omitempty"`
	SlbTemplateServerInstanceBwRateLimitAcct        string `json:"bw-rate-limit-acct,omitempty"`
	SlbTemplateServerInstanceBwRateLimitDuration    int    `json:"bw-rate-limit-duration,omitempty"`
	SlbTemplateServerInstanceBwRateLimitNoLogging   int    `json:"bw-rate-limit-no-logging,omitempty"`
	SlbTemplateServerInstanceBwRateLimitResume      int    `json:"bw-rate-limit-resume,omitempty"`
	SlbTemplateServerInstanceConnLimit              int    `json:"conn-limit,omitempty"`
	SlbTemplateServerInstanceConnLimitNoLogging     int    `json:"conn-limit-no-logging,omitempty"`
	SlbTemplateServerInstanceConnRateLimit          int    `json:"conn-rate-limit,omitempty"`
	SlbTemplateServerInstanceConnRateLimitNoLogging int    `json:"conn-rate-limit-no-logging,omitempty"`
	SlbTemplateServerInstanceDNSFailInterval        int    `json:"dns-fail-interval,omitempty"`
	SlbTemplateServerInstanceDNSQueryInterval       int    `json:"dns-query-interval,omitempty"`
	SlbTemplateServerInstanceDynamicServerPrefix    string `json:"dynamic-server-prefix,omitempty"`
	SlbTemplateServerInstanceEvery                  int    `json:"every,omitempty"`
	SlbTemplateServerInstanceExtendedStats          int    `json:"extended-stats,omitempty"`
	SlbTemplateServerInstanceHealthCheck            string `json:"health-check,omitempty"`
	SlbTemplateServerInstanceHealthCheckDisable     int    `json:"health-check-disable,omitempty"`
	SlbTemplateServerInstanceInitialSlowStart       int    `json:"initial-slow-start,omitempty"`
	SlbTemplateServerInstanceLogSelectionFailure    int    `json:"log-selection-failure,omitempty"`
	SlbTemplateServerInstanceMaxDynamicServer       int    `json:"max-dynamic-server,omitempty"`
	SlbTemplateServerInstanceMinTTLRatio            int    `json:"min-ttl-ratio,omitempty"`
	SlbTemplateServerInstanceName                   string `json:"name,omitempty"`
	SlbTemplateServerInstanceRateInterval           string `json:"rate-interval,omitempty"`
	SlbTemplateServerInstanceResume                 int    `json:"resume,omitempty"`
	SlbTemplateServerInstanceSlowStart              int    `json:"slow-start,omitempty"`
	SlbTemplateServerInstanceSpoofingCache          int    `json:"spoofing-cache,omitempty"`
	SlbTemplateServerInstanceStatsDataAction        string `json:"stats-data-action,omitempty"`
	SlbTemplateServerInstanceTill                   int    `json:"till,omitempty"`
	SlbTemplateServerInstanceTimes                  int    `json:"times,omitempty"`
	SlbTemplateServerInstanceUUID                   string `json:"uuid,omitempty"`
	SlbTemplateServerInstanceUserTag                string `json:"user-tag,omitempty"`
	SlbTemplateServerInstanceWeight                 int    `json:"weight,omitempty"`
}

func PostSlbTemplateServer(id string, inst SlbTemplateServer, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/server", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateServer(id string, name1 string, host string) (*SlbTemplateServer, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateServer")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/server/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateServer", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateServer(id string, name1 string, inst SlbTemplateServer, host string) error {

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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/server/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateServer(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateServer")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/server/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServer
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateServer", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
