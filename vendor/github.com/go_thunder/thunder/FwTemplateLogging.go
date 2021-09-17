package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTemplateLogging struct {
	UUID FwTemplateLoggingInstance `json:"logging-instance,omitempty"`
}

type FwTemplateLoggingInstance struct {
	Facility         string                                  `json:"facility,omitempty"`
	Format           string                                  `json:"format,omitempty"`
	IncludeDestFqdn  int                                     `json:"include-dest-fqdn,omitempty"`
	FileExtension    FwTemplateLoggingIncludeHTTP            `json:"include-http,omitempty"`
	FramedIpv6Prefix FwTemplateLoggingIncludeRadiusAttribute `json:"include-radius-attribute,omitempty"`
	HTTPRequests     FwTemplateLoggingLog                    `json:"log,omitempty"`
	MergedStyle      int                                     `json:"merged-style,omitempty"`
	Name             string                                  `json:"name,omitempty"`
	Resolution       string                                  `json:"resolution,omitempty"`
	IncludeByteCount FwTemplateLoggingRule                   `json:"rule,omitempty"`
	ServiceGroup     string                                  `json:"service-group,omitempty"`
	Severity         string                                  `json:"severity,omitempty"`
	IP               FwTemplateLoggingSourceAddress          `json:"source-address,omitempty"`
	UUID             string                                  `json:"uuid,omitempty"`
	UserTag          string                                  `json:"user-tag,omitempty"`
}

type FwTemplateLoggingIncludeHTTP struct {
	FileExtension   int                          `json:"file-extension,omitempty"`
	CustomMaxLength []FwTemplateLoggingHeaderCfg `json:"header-cfg,omitempty"`
	L4SessionInfo   int                          `json:"l4-session-info,omitempty"`
	Method          int                          `json:"method,omitempty"`
	RequestNumber   int                          `json:"request-number,omitempty"`
}

type FwTemplateLoggingIncludeRadiusAttribute struct {
	AttrEvent           []FwTemplateLoggingAttrCfg `json:"attr-cfg,omitempty"`
	FramedIpv6Prefix    int                        `json:"framed-ipv6-prefix,omitempty"`
	InsertIfNotExisting int                        `json:"insert-if-not-existing,omitempty"`
	NoQuote             int                        `json:"no-quote,omitempty"`
	PrefixLength        string                     `json:"prefix-length,omitempty"`
	ZeroInCustomAttr    int                        `json:"zero-in-custom-attr,omitempty"`
}

type FwTemplateLoggingLog struct {
	HTTPRequests string `json:"http-requests,omitempty"`
}

type FwTemplateLoggingRule struct {
	IncludeByteCount FwTemplateLoggingRuleHTTPRequests `json:"rule-http-requests,omitempty"`
}

type FwTemplateLoggingSourceAddress struct {
	IP   string `json:"ip,omitempty"`
	Ipv6 string `json:"ipv6,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

type FwTemplateLoggingHeaderCfg struct {
	CustomHeaderName string `json:"custom-header-name,omitempty"`
	CustomMaxLength  int    `json:"custom-max-length,omitempty"`
	HTTPHeader       string `json:"http-header,omitempty"`
	MaxLength        int    `json:"max-length,omitempty"`
}

type FwTemplateLoggingAttrCfg struct {
	Attr      string `json:"attr,omitempty"`
	AttrEvent string `json:"attr-event,omitempty"`
}

type FwTemplateLoggingRuleHTTPRequests struct {
	IncludeByteCount     []FwTemplateLoggingDestPort `json:"dest-port,omitempty"`
	DisableSequenceCheck int                         `json:"disable-sequence-check,omitempty"`
	IncludeAllHeaders    int                         `json:"include-all-headers,omitempty"`
	LogEveryHTTPRequest  int                         `json:"log-every-http-request,omitempty"`
	MaxURLLen            int                         `json:"max-url-len,omitempty"`
}

type FwTemplateLoggingDestPort struct {
	DestPortNumber   int `json:"dest-port-number,omitempty"`
	IncludeByteCount int `json:"include-byte-count,omitempty"`
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
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/template/logging", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTemplateLogging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTemplateLogging REQ RES..........................", m)
			err := check_api_status("PostFwTemplateLogging", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwTemplateLogging(id string, name string, host string) (*FwTemplateLogging, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTemplateLogging")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/template/logging/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTemplateLogging
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwTemplateLogging REQ RES..........................", m)
			err := check_api_status("GetFwTemplateLogging", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
