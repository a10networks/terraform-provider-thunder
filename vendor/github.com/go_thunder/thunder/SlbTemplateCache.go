package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateCache struct {
	SlbTemplateCacheInstanceName SlbTemplateCacheInstance `json:"cache,omitempty"`
}

type SlbTemplateCacheInstance struct {
	SlbTemplateCacheInstanceAcceptReloadReq         int                                      `json:"accept-reload-req,omitempty"`
	SlbTemplateCacheInstanceAge                     int                                      `json:"age,omitempty"`
	SlbTemplateCacheInstanceDefaultPolicyNocache    int                                      `json:"default-policy-nocache,omitempty"`
	SlbTemplateCacheInstanceDisableInsertAge        int                                      `json:"disable-insert-age,omitempty"`
	SlbTemplateCacheInstanceDisableInsertVia        int                                      `json:"disable-insert-via,omitempty"`
	SlbTemplateCacheInstanceLocalURIPolicyLocalURI  []SlbTemplateCacheInstanceLocalURIPolicy `json:"local-uri-policy,omitempty"`
	SlbTemplateCacheInstanceLogging                 string                                   `json:"logging,omitempty"`
	SlbTemplateCacheInstanceMaxCacheSize            int                                      `json:"max-cache-size,omitempty"`
	SlbTemplateCacheInstanceMaxContentSize          int                                      `json:"max-content-size,omitempty"`
	SlbTemplateCacheInstanceMinContentSize          int                                      `json:"min-content-size,omitempty"`
	SlbTemplateCacheInstanceName                    string                                   `json:"name,omitempty"`
	SlbTemplateCacheInstancePacketCaptureTemplate   string                                   `json:"packet-capture-template,omitempty"`
	SlbTemplateCacheInstanceRemoveCookies           int                                      `json:"remove-cookies,omitempty"`
	SlbTemplateCacheInstanceReplacementPolicy       string                                   `json:"replacement-policy,omitempty"`
	SlbTemplateCacheInstanceSamplingEnableCounters1 []SlbTemplateCacheInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	SlbTemplateCacheInstanceURIPolicyURI            []SlbTemplateCacheInstanceURIPolicy      `json:"uri-policy,omitempty"`
	SlbTemplateCacheInstanceUUID                    string                                   `json:"uuid,omitempty"`
	SlbTemplateCacheInstanceUserTag                 string                                   `json:"user-tag,omitempty"`
	SlbTemplateCacheInstanceVerifyHost              int                                      `json:"verify-host,omitempty"`
}

type SlbTemplateCacheInstanceLocalURIPolicy struct {
	SlbTemplateCacheInstanceLocalURIPolicyLocalURI string `json:"local-uri,omitempty"`
}

type SlbTemplateCacheInstanceSamplingEnable struct {
	SlbTemplateCacheInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type SlbTemplateCacheInstanceURIPolicy struct {
	SlbTemplateCacheInstanceURIPolicyCacheAction string `json:"cache-action,omitempty"`
	SlbTemplateCacheInstanceURIPolicyCacheValue  int    `json:"cache-value,omitempty"`
	SlbTemplateCacheInstanceURIPolicyInvalidate  string `json:"invalidate,omitempty"`
	SlbTemplateCacheInstanceURIPolicyURI         string `json:"uri,omitempty"`
}

func PostSlbTemplateCache(id string, inst SlbTemplateCache, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/cache", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateCache
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateCache", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateCache(id string, name1 string, host string) (*SlbTemplateCache, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateCache")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/cache/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateCache
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateCache", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateCache(id string, name1 string, inst SlbTemplateCache, host string) error {

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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/cache/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateCache
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateCache", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateCache(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateCache")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/cache/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateCache
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateCache", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
