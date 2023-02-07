package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Port struct {
	UUID PortInstance `json:"port,omitempty"`
}

type PortInstance struct {
	HealthCheckDisable     int    `json:"health-check-disable,omitempty"`
	StatsDataAction        string `json:"stats-data-action,omitempty"`
	ReselOnReset           int    `json:"resel-on-reset,omitempty"`
	DestNat                int    `json:"dest-nat,omitempty"`
	RestoreSvcTime         int    `json:"restore-svc-time,omitempty"`
	RequestRateLimit       int    `json:"request-rate-limit,omitempty"`
	DynamicMemberPriority  int    `json:"dynamic-member-priority,omitempty"`
	BwRateLimit            int    `json:"bw-rate-limit,omitempty"`
	SlowStart              int    `json:"slow-start,omitempty"`
	Decrement              int    `json:"decrement,omitempty"`
	ConnLimit              int    `json:"conn-limit,omitempty"`
	Retry                  int    `json:"retry,omitempty"`
	Weight                 int    `json:"weight,omitempty"`
	InbandHealthCheck      int    `json:"inband-health-check,omitempty"`
	Resume                 int    `json:"resume,omitempty"`
	RateInterval           string `json:"rate-interval,omitempty"`
	NoSsl                  int    `json:"no-ssl,omitempty"`
	Till                   int    `json:"till,omitempty"`
	FlapPeriod             int    `json:"flap-period,omitempty"`
	SubGroup               int    `json:"sub-group,omitempty"`
	DampeningFlaps         int    `json:"dampening-flaps,omitempty"`
	BwRateLimitNoLogging   int    `json:"bw-rate-limit-no-logging,omitempty"`
	DownGracePeriod        int    `json:"down-grace-period,omitempty"`
	InitialSlowStart       int    `json:"initial-slow-start,omitempty"`
	Dscp                   int    `json:"dscp,omitempty"`
	RequestRateInterval    string `json:"request-rate-interval,omitempty"`
	Add                    int    `json:"add,omitempty"`
	Every                  int    `json:"every,omitempty"`
	SharedPartitionPool    int    `json:"shared-partition-pool,omitempty"`
	ConnLimitNoLogging     int    `json:"conn-limit-no-logging,omitempty"`
	ExtendedStats          int    `json:"extended-stats,omitempty"`
	UUID                   string `json:"uuid,omitempty"`
	Reset                  int    `json:"reset,omitempty"`
	DelSessionOnServerDown int    `json:"del-session-on-server-down,omitempty"`
	ConnRateLimitNoLogging int    `json:"conn-rate-limit-no-logging,omitempty"`
	Name                   string `json:"name,omitempty"`
	TemplatePortPoolShared string `json:"template-port-pool-shared,omitempty"`
	BwRateLimitDuration    int    `json:"bw-rate-limit-duration,omitempty"`
	BwRateLimitResume      int    `json:"bw-rate-limit-resume,omitempty"`
	UserTag                string `json:"user-tag,omitempty"`
	Times                  int    `json:"times,omitempty"`
	RequestRateNoLogging   int    `json:"request-rate-no-logging,omitempty"`
	DownTimer              int    `json:"down-timer,omitempty"`
	ConnRateLimit          int    `json:"conn-rate-limit,omitempty"`
	SourceNat              string `json:"source-nat,omitempty"`
	Reassign               int    `json:"reassign,omitempty"`
	HealthCheck            string `json:"health-check,omitempty"`
}

func PostTemplatePort(id string, inst Port, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplatePort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/port", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Port
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostTemplatePort REQ RES..........................", m)
			err := check_api_status("PostTemplatePort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetTemplatePort(id string, name string, host string) (*Port, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplatePort")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/port/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Port
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetTemplatePort REQ RES..........................", m)
			err := check_api_status("GetTemplatePort", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutTemplatePort(id string, name string, inst Port, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplatePort")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/port/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Port
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutTemplatePort REQ RES..........................", m)
			err := check_api_status("PutTemplatePort", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteTemplatePort(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplatePort")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/port/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Port
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
