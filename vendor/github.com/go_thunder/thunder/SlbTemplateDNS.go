package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type DNS struct {
	Name DNSInstance `json:"dns,omitempty"`
}

type LidList struct {
	ActionValue     string `json:"action-value,omitempty"`
	Log             int    `json:"log,omitempty"`
	Lidnum          int    `json:"lidnum,omitempty"`
	OverLimitAction int    `json:"over-limit-action,omitempty"`
	Per             int    `json:"per,omitempty"`
	Lockout         int    `json:"lockout,omitempty"`
	UserTag         string `json:"user-tag,omitempty"`
	ConnRateLimit   int    `json:"conn-rate-limit,omitempty"`
	LogInterval     int    `json:"log-interval,omitempty"`
	UUID            string `json:"uuid,omitempty"`
}
type ClassList struct {
	ActionValue []LidList `json:"lid-list,omitempty"`
	Name        string    `json:"name,omitempty"`
	UUID        string    `json:"uuid,omitempty"`
}

type ResponseRateLimiting struct {
	FilterResponseRate int    `json:"filter-response-rate,omitempty"`
	SlipRate           int    `json:"slip-rate,omitempty"`
	ResponseRate       int    `json:"response-rate,omitempty"`
	Window             int    `json:"window,omitempty"`
	Action             string `json:"action,omitempty"`
	EnableLog          int    `json:"enable-log,omitempty"`
	UUID               string `json:"uuid,omitempty"`
}
type DNSInstance struct {
	DnssecServiceGroup string               `json:"dnssec-service-group,omitempty"`
	Name               string               `json:"name,omitempty"`
	LidList            ClassList            `json:"class-list,omitempty"`
	FilterResponseRate ResponseRateLimiting `json:"response-rate-limiting,omitempty"`
	Drop               int                  `json:"drop,omitempty"`
	Period             int                  `json:"period,omitempty"`
	UserTag            string               `json:"user-tag,omitempty"`
	QueryIDSwitch      int                  `json:"query-id-switch,omitempty"`
	EnableCacheSharing int                  `json:"enable-cache-sharing,omitempty"`
	RedirectToTCPPort  int                  `json:"redirect-to-tcp-port,omitempty"`
	MaxQueryLength     int                  `json:"max-query-length,omitempty"`
	DisableDNSTemplate int                  `json:"disable-dns-template,omitempty"`
	Forward            string               `json:"forward,omitempty"`
	MaxCacheSize       int                  `json:"max-cache-size,omitempty"`
	DefaultPolicy      string               `json:"default-policy,omitempty"`
	MaxCacheEntrySize  int                  `json:"max-cache-entry-size,omitempty"`
	UUID               string               `json:"uuid,omitempty"`
}

func PostTemplateDNS(id string, inst DNS, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplateDNS")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}
	logger.Println("LINK = https://" + host + "/axapi/v3/slb/template/dns" + "\n id = " + id)
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/dns", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DNS
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetTemplateDNS(id string, name string, host string) (*DNS, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplateDNS")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/dns/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DNS
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

func PutTemplateDNS(id string, name string, inst DNS, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplateDNS")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/dns/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DNS
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteTemplateDNS(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplateDNS")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/dns/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DNS
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
