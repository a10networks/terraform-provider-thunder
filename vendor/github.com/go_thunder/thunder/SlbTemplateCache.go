package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type Cache struct {
	UUID CacheInstance `json:"cache,omitempty"`
}

type LocalURIPolicy struct {
	LocalURI string `json:"local-uri,omitempty"`
}
type SamplingEnable4 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type URIPolicy struct {
	CacheAction string `json:"cache-action,omitempty"`
	CacheValue  int    `json:"cache-value,omitempty"`
	URI         string `json:"uri,omitempty"`
	Invalidate  string `json:"invalidate,omitempty"`
}
type CacheInstance struct {
	AcceptReloadReq      int               `json:"accept-reload-req,omitempty"`
	Name                 string            `json:"name,omitempty"`
	DefaultPolicyNocache int               `json:"default-policy-nocache,omitempty"`
	Age                  int               `json:"age,omitempty"`
	DisableInsertVia     int               `json:"disable-insert-via,omitempty"`
	UserTag              string            `json:"user-tag,omitempty"`
	LocalURI             []LocalURIPolicy  `json:"local-uri-policy,omitempty"`
	Counters1            []SamplingEnable4 `json:"sampling-enable,omitempty"`
	ReplacementPolicy    string            `json:"replacement-policy,omitempty"`
	DisableInsertAge     int               `json:"disable-insert-age,omitempty"`
	MaxContentSize       int               `json:"max-content-size,omitempty"`
	MaxCacheSize         int               `json:"max-cache-size,omitempty"`
	Logging              string            `json:"logging,omitempty"`
	URI                  []URIPolicy       `json:"uri-policy,omitempty"`
	RemoveCookies        int               `json:"remove-cookies,omitempty"`
	VerifyHost           int               `json:"verify-host,omitempty"`
	MinContentSize       int               `json:"min-content-size,omitempty"`
	UUID                 string            `json:"uuid,omitempty"`
}

func PostSlbTemplateCache(id string, inst Cache, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateCache")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/cache", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Cache
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetSlbTemplateCache(id string, name string, host string) (*Cache, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateCache")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/cache/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Cache
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

func PutSlbTemplateCache(id string, name string, inst Cache, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateCache")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/cache/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Cache
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteSlbTemplateCache(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateCache")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/cache/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Cache
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
