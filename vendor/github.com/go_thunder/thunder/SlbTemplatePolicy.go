package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type Policy struct {
	UUID PolicyInstance `json:"policy,omitempty"`
}

type Filtering struct {
	SsliURLFiltering string `json:"ssli-url-filtering,omitempty"`
}
type SanFiltering struct {
	SsliURLFilteringSan string `json:"ssli-url-filtering-san,omitempty"`
}
type SamplingEnable3 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type ActionList struct {
	Log              int               `json:"log,omitempty"`
	HTTPStatusCode   string            `json:"http-status-code,omitempty"`
	ForwardSnat      string            `json:"forward-snat,omitempty"`
	UUID             string            `json:"uuid,omitempty"`
	DropResponseCode int               `json:"drop-response-code,omitempty"`
	Action1          string            `json:"action1,omitempty"`
	FakeSg           string            `json:"fake-sg,omitempty"`
	UserTag          string            `json:"user-tag,omitempty"`
	RealSg           string            `json:"real-sg,omitempty"`
	DropMessage      string            `json:"drop-message,omitempty"`
	Counters1        []SamplingEnable3 `json:"sampling-enable,omitempty"`
	FallBack         string            `json:"fall-back,omitempty"`
	FallBackSnat     string            `json:"fall-back-snat,omitempty"`
	DropRedirectURL  string            `json:"drop-redirect-url,omitempty"`
	Name             string            `json:"name,omitempty"`
}
type ClassListList struct {
	UUID          string `json:"uuid,omitempty"`
	DestClassList string `json:"dest-class-list,omitempty"`
	Priority      int    `json:"priority,omitempty"`

	Counters1 []SamplingEnable3 `json:"sampling-enable,omitempty"`
	Action    string            `json:"action,omitempty"`
	Type      string            `json:"type,omitempty"`
}
type WebCategoryListList struct {
	UUID            string            `json:"uuid,omitempty"`
	WebCategoryList string            `json:"web-category-list,omitempty"`
	Priority        int               `json:"priority,omitempty"`
	Counters1       []SamplingEnable3 `json:"sampling-enable,omitempty"`
	Action          string            `json:"action,omitempty"`
	Type            string            `json:"type,omitempty"`
}
type Any struct {
	Action    string            `json:"action,omitempty"`
	Counters1 []SamplingEnable3 `json:"sampling-enable,omitempty"`
	UUID      string            `json:"uuid,omitempty"`
}
type Destination struct {
	UUID            []ClassListList       `json:"class-list-list,omitempty"`
	WebCategoryList []WebCategoryListList `json:"web-category-list-list,omitempty"`
	Action          Any                   `json:"any,omitempty"`
}
type SourceList struct {
	MatchAny             int               `json:"match-any,omitempty"`
	Name                 string            `json:"name,omitempty"`
	MatchAuthorizePolicy string            `json:"match-authorize-policy,omitempty"`
	WebCategoryList      Destination       `json:"destination,omitempty"`
	UserTag              string            `json:"user-tag,omitempty"`
	Priority             int               `json:"priority,omitempty"`
	Counters1            []SamplingEnable3 `json:"sampling-enable,omitempty"`
	MatchClassList       string            `json:"match-class-list,omitempty"`
	UUID                 string            `json:"uuid,omitempty"`
}
type ForwardPolicy struct {
	SsliURLFiltering    []Filtering    `json:"filtering,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	LocalLogging        int            `json:"local-logging,omitempty"`
	SsliURLFilteringSan []SanFiltering `json:"san-filtering,omitempty"`
	Log                 []ActionList   `json:"action-list,omitempty"`
	NoClientConnReuse   int            `json:"no-client-conn-reuse,omitempty"`
	RequireWebCategory  int            `json:"require-web-category,omitempty"`
	AcosEventLog        int            `json:"acos-event-log,omitempty"`
	MatchAny            []SourceList   `json:"source-list,omitempty"`
}
type DNS64 struct {
	Prefix          string `json:"prefix,omitempty"`
	ExclusiveAnswer int    `json:"exclusive-answer,omitempty"`
	Disable         int    `json:"disable,omitempty"`
}
type ResponseCodeRateLimit struct {
	Threshold      int `json:"threshold,omitempty"`
	CodeRangeEnd   int `json:"code-range-end,omitempty"`
	CodeRangeStart int `json:"code-range-start,omitempty"`
	Period         int `json:"period,omitempty"`
}
type LidList2 struct {
	RequestRateLimit     int                     `json:"request-rate-limit,omitempty"`
	ActionValue          string                  `json:"action-value,omitempty"`
	RequestPer           int                     `json:"request-per,omitempty"`
	BwRateLimit          int                     `json:"bw-rate-limit,omitempty"`
	ConnLimit            int                     `json:"conn-limit,omitempty"`
	Log                  int                     `json:"log,omitempty"`
	DirectActionValue    string                  `json:"direct-action-value,omitempty"`
	ConnPer              int                     `json:"conn-per,omitempty"`
	DirectFail           int                     `json:"direct-fail,omitempty"`
	ConnRateLimit        int                     `json:"conn-rate-limit,omitempty"`
	DirectPbslbLogging   int                     `json:"direct-pbslb-logging,omitempty"`
	Prefix               DNS64                   `json:"dns64,omitempty"`
	Lidnum               int                     `json:"lidnum,omitempty"`
	OverLimitAction      int                     `json:"over-limit-action,omitempty"`
	Threshold            []ResponseCodeRateLimit `json:"response-code-rate-limit,omitempty"`
	DirectServiceGroup   string                  `json:"direct-service-group,omitempty"`
	UUID                 string                  `json:"uuid,omitempty"`
	RequestLimit         int                     `json:"request-limit,omitempty"`
	DirectActionInterval int                     `json:"direct-action-interval,omitempty"`
	BwPer                int                     `json:"bw-per,omitempty"`
	Interval             int                     `json:"interval,omitempty"`
	UserTag              string                  `json:"user-tag,omitempty"`
	DirectAction         int                     `json:"direct-action,omitempty"`
	Lockout              int                     `json:"lockout,omitempty"`
	DirectLoggingDrpRst  int                     `json:"direct-logging-drp-rst,omitempty"`
	DirectPbslbInterval  int                     `json:"direct-pbslb-interval,omitempty"`
}
type ClassList2 struct {
	HeaderName       string     `json:"header-name,omitempty"`
	RequestRateLimit []LidList2 `json:"lid-list,omitempty"`
	Name             string     `json:"name,omitempty"`
	ClientIPL3Dest   int        `json:"client-ip-l3-dest,omitempty"`
	ClientIPL7Header int        `json:"client-ip-l7-header,omitempty"`
	UUID             string     `json:"uuid,omitempty"`
}
type BwListID struct {
	PbslbInterval  int    `json:"pbslb-interval,omitempty"`
	ActionInterval int    `json:"action-interval,omitempty"`
	ServiceGroup   string `json:"service-group,omitempty"`
	LoggingDrpRst  int    `json:"logging-drp-rst,omitempty"`
	Fail           int    `json:"fail,omitempty"`
	PbslbLogging   int    `json:"pbslb-logging,omitempty"`
	ID             int    `json:"id,omitempty"`
	BwListAction   string `json:"bw-list-action,omitempty"`
}
type PolicyInstance struct {
	SsliURLFiltering ForwardPolicy     `json:"forward-policy,omitempty"`
	UseDestinationIP int               `json:"use-destination-ip,omitempty"`
	Name             string            `json:"name,omitempty"`
	OverLimit        int               `json:"over-limit,omitempty"`
	HeaderName       ClassList2        `json:"class-list,omitempty"`
	Interval         int               `json:"interval,omitempty"`
	Share            int               `json:"share,omitempty"`
	FullDomainTree   int               `json:"full-domain-tree,omitempty"`
	OverLimitLogging int               `json:"over-limit-logging,omitempty"`
	BwListName       string            `json:"bw-list-name,omitempty"`
	Timeout          int               `json:"timeout,omitempty"`
	Counters1        []SamplingEnable3 `json:"sampling-enable,omitempty"`
	UserTag          string            `json:"user-tag,omitempty"`
	PbslbInterval    []BwListID        `json:"bw-list-id,omitempty"`
	OverLimitLockup  int               `json:"over-limit-lockup,omitempty"`
	UUID             string            `json:"uuid,omitempty"`
	OverLimitReset   int               `json:"over-limit-reset,omitempty"`
	Overlap          int               `json:"overlap,omitempty"`
}

func PostSlbTemplatePolicy(id string, inst Policy, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplatePolicy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/policy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Policy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetSlbTemplatePolicy(id string, name string, host string) (*Policy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplatePolicy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/policy/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Policy
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

func PutSlbTemplatePolicy(id string, name string, inst Policy, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplatePolicy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/policy/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Policy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteSlbTemplatePolicy(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplatePolicy")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/policy/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Policy
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
