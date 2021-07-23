package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplatePersistCookie struct {
	PassPhrase SlbTemplatePersistCookieInstance `json:"cookie,omitempty"`
}

type SlbTemplatePersistCookieInstance struct {
	CookieName         string `json:"cookie-name,omitempty"`
	Domain             string `json:"domain,omitempty"`
	DontHonorConnRules int    `json:"dont-honor-conn-rules,omitempty"`
	EncryptLevel       int    `json:"encrypt-level,omitempty"`
	Encrypted          string `json:"encrypted,omitempty"`
	Expire             int    `json:"expire,omitempty"`
	Httponly           int    `json:"httponly,omitempty"`
	InsertAlways       int    `json:"insert-always,omitempty"`
	MatchType          int    `json:"match-type,omitempty"`
	Name               string `json:"name,omitempty"`
	PassPhrase         string `json:"pass-phrase,omitempty"`
	PassThru           int    `json:"pass-thru,omitempty"`
	Path               string `json:"path,omitempty"`
	ScanAllMembers     int    `json:"scan-all-members,omitempty"`
	Secure             int    `json:"secure,omitempty"`
	Server             int    `json:"server,omitempty"`
	ServerServiceGroup int    `json:"server-service-group,omitempty"`
	ServiceGroup       int    `json:"service-group,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	UserTag            string `json:"user-tag,omitempty"`
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
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/persist/cookie", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplatePersistCookie REQ RES..........................", m)
			err := check_api_status("PostSlbTemplatePersistCookie", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetSlbTemplatePersistCookie(id string, name string, host string) (*SlbTemplatePersistCookie, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplatePersistCookie")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/persist/cookie/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplatePersistCookie REQ RES..........................", m)
			err := check_api_status("GetSlbTemplatePersistCookie", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplatePersistCookie(id string, name string, inst SlbTemplatePersistCookie, host string) error {

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
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/persist/cookie/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplatePersistCookie REQ RES..........................", m)
			err := check_api_status("PutSlbTemplatePersistCookie", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteSlbTemplatePersistCookie(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplatePersistCookie")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/persist/cookie/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplatePersistCookie
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
