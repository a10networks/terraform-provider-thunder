package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTemplateLogging struct {
	FwTemplateLoggingInstanceName FwTemplateLoggingInstance `json:"logging,omitempty"`
}

type FwTemplateLoggingInstance struct {
	FwTemplateLoggingInstanceFacility                      string                                          `json:"facility,omitempty"`
	FwTemplateLoggingInstanceFormat                        string                                          `json:"format,omitempty"`
	FwTemplateLoggingInstanceIncludeDestFqdn               int                                             `json:"include-dest-fqdn,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPHeaderCfg          FwTemplateLoggingInstanceIncludeHTTP            `json:"include-http,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfg FwTemplateLoggingInstanceIncludeRadiusAttribute `json:"include-radius-attribute,omitempty"`
	FwTemplateLoggingInstanceLogHTTPRequests               FwTemplateLoggingInstanceLog                    `json:"log,omitempty"`
	FwTemplateLoggingInstanceMergedStyle                   int                                             `json:"merged-style,omitempty"`
	FwTemplateLoggingInstanceName                          string                                          `json:"name,omitempty"`
	FwTemplateLoggingInstanceResolution                    string                                          `json:"resolution,omitempty"`
	FwTemplateLoggingInstanceRuleRuleHTTPRequests          FwTemplateLoggingInstanceRule                   `json:"rule,omitempty"`
	FwTemplateLoggingInstanceServiceGroup                  string                                          `json:"service-group,omitempty"`
	FwTemplateLoggingInstanceSeverity                      string                                          `json:"severity,omitempty"`
	FwTemplateLoggingInstanceSourceAddressIP               FwTemplateLoggingInstanceSourceAddress          `json:"source-address,omitempty"`
	FwTemplateLoggingInstanceUUID                          string                                          `json:"uuid,omitempty"`
	FwTemplateLoggingInstanceUserTag                       string                                          `json:"user-tag,omitempty"`
}

type FwTemplateLoggingInstanceIncludeHTTP struct {
	FwTemplateLoggingInstanceIncludeHTTPFileExtension       int                                             `json:"file-extension,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPHeaderCfgHTTPHeader []FwTemplateLoggingInstanceIncludeHTTPHeaderCfg `json:"header-cfg,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPL4SessionInfo       int                                             `json:"l4-session-info,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPMethod              int                                             `json:"method,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPRequestNumber       int                                             `json:"request-number,omitempty"`
}

type FwTemplateLoggingInstanceIncludeRadiusAttribute struct {
	FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttr         []FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfg `json:"attr-cfg,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributeFramedIpv6Prefix    int                                                      `json:"framed-ipv6-prefix,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributeInsertIfNotExisting int                                                      `json:"insert-if-not-existing,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributeNoQuote             int                                                      `json:"no-quote,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributePrefixLength        string                                                   `json:"prefix-length,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributeZeroInCustomAttr    int                                                      `json:"zero-in-custom-attr,omitempty"`
}

type FwTemplateLoggingInstanceLog struct {
	FwTemplateLoggingInstanceLogHTTPRequests string `json:"http-requests,omitempty"`
}

type FwTemplateLoggingInstanceRule struct {
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPort FwTemplateLoggingInstanceRuleRuleHTTPRequests `json:"rule-http-requests,omitempty"`
}

type FwTemplateLoggingInstanceSourceAddress struct {
	FwTemplateLoggingInstanceSourceAddressIP   string `json:"ip,omitempty"`
	FwTemplateLoggingInstanceSourceAddressIpv6 string `json:"ipv6,omitempty"`
	FwTemplateLoggingInstanceSourceAddressUUID string `json:"uuid,omitempty"`
}

type FwTemplateLoggingInstanceIncludeHTTPHeaderCfg struct {
	FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCustomHeaderName string `json:"custom-header-name,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPHeaderCfgCustomMaxLength  int    `json:"custom-max-length,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPHeaderCfgHTTPHeader       string `json:"http-header,omitempty"`
	FwTemplateLoggingInstanceIncludeHTTPHeaderCfgMaxLength        int    `json:"max-length,omitempty"`
}

type FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfg struct {
	FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttr      string `json:"attr,omitempty"`
	FwTemplateLoggingInstanceIncludeRadiusAttributeAttrCfgAttrEvent string `json:"attr-event,omitempty"`
}

type FwTemplateLoggingInstanceRuleRuleHTTPRequests struct {
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortDestPortNumber []FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPort `json:"dest-port,omitempty"`
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsDisableSequenceCheck   int                                                     `json:"disable-sequence-check,omitempty"`
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsIncludeAllHeaders      int                                                     `json:"include-all-headers,omitempty"`
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsLogEveryHTTPRequest    int                                                     `json:"log-every-http-request,omitempty"`
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsMaxURLLen              int                                                     `json:"max-url-len,omitempty"`
}

type FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPort struct {
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortDestPortNumber   int `json:"dest-port-number,omitempty"`
	FwTemplateLoggingInstanceRuleRuleHTTPRequestsDestPortIncludeByteCount int `json:"include-byte-count,omitempty"`
}

func PostFwTemplateLogging(id string, inst FwTemplateLogging, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTemplateLogging")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/template/logging", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTemplateLogging
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostFwTemplateLogging", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwTemplateLogging(id string, name1 string, host string) (*FwTemplateLogging, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTemplateLogging")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/template/logging/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTemplateLogging
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetFwTemplateLogging", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutFwTemplateLogging(id string, name1 string, inst FwTemplateLogging, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutFwTemplateLogging")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/fw/template/logging/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTemplateLogging
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutFwTemplateLogging", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteFwTemplateLogging(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteFwTemplateLogging")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/fw/template/logging/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTemplateLogging
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteFwTemplateLogging", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
