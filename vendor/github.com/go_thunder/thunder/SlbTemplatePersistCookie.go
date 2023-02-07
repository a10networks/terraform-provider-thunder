package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplatePersistCookie struct {
	SlbTemplatePersistCookieInstanceName SlbTemplatePersistCookieInstance `json:"cookie,omitempty"`
}

type SlbTemplatePersistCookieInstance struct {
	SlbTemplatePersistCookieInstanceCookieName         string `json:"cookie-name,omitempty"`
	SlbTemplatePersistCookieInstanceDomain             string `json:"domain,omitempty"`
	SlbTemplatePersistCookieInstanceDontHonorConnRules int    `json:"dont-honor-conn-rules,omitempty"`
	SlbTemplatePersistCookieInstanceEncryptLevel       int    `json:"encrypt-level"`
	SlbTemplatePersistCookieInstanceExpire             int    `json:"expire,omitempty"`
	SlbTemplatePersistCookieInstanceHttponly           int    `json:"httponly,omitempty"`
	SlbTemplatePersistCookieInstanceInsertAlways       int    `json:"insert-always,omitempty"`
	SlbTemplatePersistCookieInstanceMatchType          int    `json:"match-type,omitempty"`
	SlbTemplatePersistCookieInstanceName               string `json:"name,omitempty"`
	SlbTemplatePersistCookieInstancePassPhrase         string `json:"pass-phrase,omitempty"`
	SlbTemplatePersistCookieInstancePassThru           int    `json:"pass-thru,omitempty"`
	SlbTemplatePersistCookieInstancePath               string `json:"path,omitempty"`
	SlbTemplatePersistCookieInstancePrefix             string `json:"prefix,omitempty"`
	SlbTemplatePersistCookieInstanceSamesite           string `json:"samesite,omitempty"`
	SlbTemplatePersistCookieInstanceScanAllMembers     int    `json:"scan-all-members,omitempty"`
	SlbTemplatePersistCookieInstanceSecure             int    `json:"secure,omitempty"`
	SlbTemplatePersistCookieInstanceServer             int    `json:"server,omitempty"`
	SlbTemplatePersistCookieInstanceServerServiceGroup int    `json:"server-service-group,omitempty"`
	SlbTemplatePersistCookieInstanceServiceGroup       int    `json:"service-group,omitempty"`
	SlbTemplatePersistCookieInstanceUUID               string `json:"uuid,omitempty"`
	SlbTemplatePersistCookieInstanceUserTag            string `json:"user-tag,omitempty"`
}

func PostSlbTemplatePersistCookie(id string, inst SlbTemplatePersistCookie, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplatePersistCookie")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/persist/cookie", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplatePersistCookie", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplatePersistCookie(id string, name1 string, host string) (*SlbTemplatePersistCookie, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplatePersistCookie")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/persist/cookie/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplatePersistCookie", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplatePersistCookie(id string, name1 string, inst SlbTemplatePersistCookie, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplatePersistCookie")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/persist/cookie/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplatePersistCookie", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplatePersistCookie(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplatePersistCookie")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/persist/cookie/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplatePersistCookie", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
