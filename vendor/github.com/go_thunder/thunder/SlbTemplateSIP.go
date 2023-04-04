package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SIP struct {
	UUID SIPInstance `json:"sip,omitempty"`
}

type ServerRequestHeader struct {
	ServerRequestHeaderInsert    string `json:"server-request-header-insert,omitempty"`
	ServerRequestEraseAll        int    `json:"server-request-erase-all,omitempty"`
	InsertConditionServerRequest string `json:"insert-condition-server-request,omitempty"`
	ServerRequestHeaderErase     string `json:"server-request-header-erase,omitempty"`
}
type ServerResponseHeader struct {
	ServerResponseHeaderInsert    string `json:"server-response-header-insert,omitempty"`
	InsertConditionServerResponse string `json:"insert-condition-server-response,omitempty"`
	ServerResponseHeaderErase     string `json:"server-response-header-erase,omitempty"`
	ServerResponseEraseAll        int    `json:"server-response-erase-all,omitempty"`
}
type ClientRequestHeader struct {
	ClientRequestHeaderErase     string `json:"client-request-header-erase,omitempty"`
	ClientRequestHeaderInsert    string `json:"client-request-header-insert,omitempty"`
	ClientRequestEraseAll        int    `json:"client-request-erase-all,omitempty"`
	InsertConditionClientRequest string `json:"insert-condition-client-request,omitempty"`
}
type ClientResponseHeader struct {
	ClientResponseEraseAll        int    `json:"client-response-erase-all,omitempty"`
	InsertConditionClientResponse string `json:"insert-condition-client-response,omitempty"`
	ClientResponseHeaderErase     string `json:"client-response-header-erase,omitempty"`
	ClientResponseHeaderInsert    string `json:"client-response-header-insert,omitempty"`
}
type ExcludeTranslation struct {
	TranslationValue string `json:"translation-value,omitempty"`
	HeaderString     string `json:"header-string,omitempty"`
}
type SIPInstance struct {
	ServerRequestHeaderInsert    []ServerRequestHeader  `json:"server-request-header,omitempty"`
	SmpCallIDRtpSession          int                    `json:"smp-call-id-rtp-session,omitempty"`
	KeepServerIPIfMatchACL       int                    `json:"keep-server-ip-if-match-acl,omitempty"`
	ClientKeepAlive              int                    `json:"client-keep-alive,omitempty"`
	AlgSourceNat                 int                    `json:"alg-source-nat,omitempty"`
	UUID                         string                 `json:"uuid,omitempty"`
	ServerResponseHeaderInsert   []ServerResponseHeader `json:"server-response-header,omitempty"`
	ServerSelectionPerRequest    int                    `json:"server-selection-per-request,omitempty"`
	ClientRequestHeaderErase     []ClientRequestHeader  `json:"client-request-header,omitempty"`
	PstnGw                       string                 `json:"pstn-gw,omitempty"`
	ServiceGroup                 string                 `json:"service-group,omitempty"`
	InsertClientIP               int                    `json:"insert-client-ip,omitempty"`
	FailedClientSelection        int                    `json:"failed-client-selection,omitempty"`
	FailedClientSelectionMessage string                 `json:"failed-client-selection-message,omitempty"`
	CallIDPersistDisable         int                    `json:"call-id-persist-disable,omitempty"`
	ACLID                        int                    `json:"acl-id,omitempty"`
	AlgDestNat                   int                    `json:"alg-dest-nat,omitempty"`
	ServerKeepAlive              int                    `json:"server-keep-alive,omitempty"`
	ClientResponseEraseAll       []ClientResponseHeader `json:"client-response-header,omitempty"`
	FailedServerSelectionMessage string                 `json:"failed-server-selection-message,omitempty"`
	Name                         string                 `json:"name,omitempty"`
	TranslationValue             []ExcludeTranslation   `json:"exclude-translation,omitempty"`
	Interval                     int                    `json:"interval,omitempty"`
	UserTag                      string                 `json:"user-tag,omitempty"`
	DialogAware                  int                    `json:"dialog-aware,omitempty"`
	FailedServerSelection        int                    `json:"failed-server-selection,omitempty"`
	DropWhenClientFail           int                    `json:"drop-when-client-fail,omitempty"`
	Timeout                      int                    `json:"timeout,omitempty"`
	DropWhenServerFail           int                    `json:"drop-when-server-fail,omitempty"`
	ACLNameValue                 string                 `json:"acl-name-value,omitempty"`
}

func PostSlbTemplateSIP(id string, inst SIP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateSIP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/sip", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateSIP REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateSIP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateSIP(id string, name string, host string) (*SIP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateSIP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/sip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateSIP REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateSIP", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateSIP(id string, name string, inst SIP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateSIP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/sip/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SIP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateSIP REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateSIP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateSIP(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateSIP")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/sip/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SIP
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
