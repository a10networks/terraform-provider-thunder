package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateHttpPolicy struct {
	SlbTemplateHTTPPolicyInstanceName SlbTemplateHTTPPolicyInstance `json:"http-policy,omitempty"`
}

type SlbTemplateHTTPPolicyInstance struct {
	SlbTemplateHTTPPolicyInstanceCookieName                   string                                            `json:"cookie-name,omitempty"`
	SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocation  []SlbTemplateHTTPPolicyInstanceGeoLocationMatch   `json:"geo-location-match,omitempty"`
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchType          []SlbTemplateHTTPPolicyInstanceHTTPPolicyMatch    `json:"http-policy-match,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListMultiMatch []SlbTemplateHTTPPolicyInstanceMultiMatchRuleList `json:"multi-match-rule-list,omitempty"`
	SlbTemplateHTTPPolicyInstanceName                         string                                            `json:"name,omitempty"`
	SlbTemplateHTTPPolicyInstanceUUID                         string                                            `json:"uuid,omitempty"`
	SlbTemplateHTTPPolicyInstanceUserTag                      string                                            `json:"user-tag,omitempty"`
}

type SlbTemplateHTTPPolicyInstanceGeoLocationMatch struct {
	SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocation             string `json:"geo-location,omitempty"`
	SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocationServiceGroup string `json:"geo-location-service-group,omitempty"`
	SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocationTemplate     string `json:"geo-location-template,omitempty"`
	SlbTemplateHTTPPolicyInstanceGeoLocationMatchGeoLocationTemplateName string `json:"geo-location-template-name,omitempty"`
}

type SlbTemplateHTTPPolicyInstanceHTTPPolicyMatch struct {
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchMatchString  string `json:"match-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchMatchType    string `json:"match-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchServiceGroup string `json:"service-group,omitempty"`
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchTemplate     string `json:"template,omitempty"`
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchTemplateName string `json:"template-name,omitempty"`
	SlbTemplateHTTPPolicyInstanceHTTPPolicyMatchType         string `json:"type,omitempty"`
}

type SlbTemplateHTTPPolicyInstanceMultiMatchRuleList struct {
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameContainsString        string `json:"cookie-name-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameContainsType          string `json:"cookie-name-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEndsWithString        string `json:"cookie-name-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEndsWithType          string `json:"cookie-name-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEqualsString          string `json:"cookie-name-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameEqualsType            string `json:"cookie-name-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameStartsWithString      string `json:"cookie-name-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieNameStartsWithType        string `json:"cookie-name-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueContainsString       string `json:"cookie-value-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueContainsType         string `json:"cookie-value-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEndsWithString       string `json:"cookie-value-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEndsWithType         string `json:"cookie-value-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEqualsString         string `json:"cookie-value-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueEqualsType           string `json:"cookie-value-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueStartsWithString     string `json:"cookie-value-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListCookieValueStartsWithType       string `json:"cookie-value-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameContainsString        string `json:"header-name-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameContainsType          string `json:"header-name-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEndsWithString        string `json:"header-name-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEndsWithType          string `json:"header-name-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEqualsString          string `json:"header-name-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameEqualsType            string `json:"header-name-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameStartsWithString      string `json:"header-name-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderNameStartsWithType        string `json:"header-name-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueContainsString       string `json:"header-value-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueContainsType         string `json:"header-value-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEndsWithString       string `json:"header-value-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEndsWithType         string `json:"header-value-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEqualsString         string `json:"header-value-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueEqualsType           string `json:"header-value-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueStartsWithString     string `json:"header-value-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHeaderValueStartsWithType       string `json:"header-value-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostContainsString              string `json:"host-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostContainsType                string `json:"host-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEndsWithString              string `json:"host-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEndsWithType                string `json:"host-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEqualsString                string `json:"host-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostEqualsType                  string `json:"host-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostStartsWithString            string `json:"host-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListHostStartsWithType              string `json:"host-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListMultiMatch                      string `json:"multi-match,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameContainsString    string `json:"query-param-name-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameContainsType      string `json:"query-param-name-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEndsWithString    string `json:"query-param-name-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEndsWithType      string `json:"query-param-name-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEqualsString      string `json:"query-param-name-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameEqualsType        string `json:"query-param-name-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameStartsWithString  string `json:"query-param-name-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamNameStartsWithType    string `json:"query-param-name-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueContainsString   string `json:"query-param-value-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueContainsType     string `json:"query-param-value-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEndsWithString   string `json:"query-param-value-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEndsWithType     string `json:"query-param-value-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEqualsString     string `json:"query-param-value-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueEqualsType       string `json:"query-param-value-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueStartsWithString string `json:"query-param-value-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListQueryParamValueStartsWithType   string `json:"query-param-value-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListSeqNum                          int    `json:"seq-num,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListServiceGroup                    string `json:"service-group,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListTemplateWaf                     string `json:"template-waf,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLContainsString               string `json:"url-contains-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLContainsType                 string `json:"url-contains-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEndsWithString               string `json:"url-ends-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEndsWithType                 string `json:"url-ends-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEqualsString                 string `json:"url-equals-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLEqualsType                   string `json:"url-equals-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLStartsWithString             string `json:"url-starts-with-string,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListURLStartsWithType               string `json:"url-starts-with-type,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListUUID                            string `json:"uuid,omitempty"`
	SlbTemplateHTTPPolicyInstanceMultiMatchRuleListUserTag                         string `json:"user-tag,omitempty"`
}

func PostSlbTemplateHttpPolicy(id string, inst SlbTemplateHttpPolicy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateHttpPolicy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/http-policy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttpPolicy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateHttpPolicy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateHttpPolicy(id string, name1 string, host string) (*SlbTemplateHttpPolicy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateHttpPolicy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/http-policy/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttpPolicy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateHttpPolicy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateHttpPolicy(id string, name1 string, inst SlbTemplateHttpPolicy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateHttpPolicy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/http-policy/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttpPolicy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateHttpPolicy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateHttpPolicy(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateHttpPolicy")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/http-policy/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttpPolicy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateHttpPolicy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
