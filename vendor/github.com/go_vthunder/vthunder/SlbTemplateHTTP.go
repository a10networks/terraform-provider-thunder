package go_vthunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type HTTP struct {
	UUID HTTPInstance `json:"http,omitempty"`
}

type CompressionExcludeURI struct {
	ExcludeURI string `json:"exclude-uri,omitempty"`
}
type MatchList struct {
	RewriteTo     string `json:"rewrite-to,omitempty"`
	RedirectMatch string `json:"redirect-match,omitempty"`
}
type RedirectRewrite struct {
	RedirectSecurePort int         `json:"redirect-secure-port,omitempty"`
	RedirectSecure     int         `json:"redirect-secure,omitempty"`
	RewriteTo          []MatchList `json:"match-list,omitempty"`
}
type RequestHeaderEraseList struct {
	RequestHeaderErase string `json:"request-header-erase,omitempty"`
}
type HostSwitching struct {
	HostSwitchingType string `json:"host-switching-type,omitempty"`
	HostServiceGroup  string `json:"host-service-group,omitempty"`
	HostMatchString   string `json:"host-match-string,omitempty"`
}
type ResponseHeaderInsertList struct {
	ResponseHeaderInsertType string `json:"response-header-insert-type,omitempty"`
	ResponseHeaderInsert     string `json:"response-header-insert,omitempty"`
}
type Template2 struct {
	Logging string `json:"logging,omitempty"`
}
type URLSwitching struct {
	URLServiceGroup  string `json:"url-service-group,omitempty"`
	URLMatchString   string `json:"url-match-string,omitempty"`
	URLSwitchingType string `json:"url-switching-type,omitempty"`
}
type ResponseContentReplaceList struct {
	ResponseNewString      string `json:"response-new-string,omitempty"`
	ResponseContentReplace string `json:"response-content-replace,omitempty"`
}
type RequestHeaderInsertList struct {
	RequestHeaderInsert     string `json:"request-header-insert,omitempty"`
	RequestHeaderInsertType string `json:"request-header-insert-type,omitempty"`
}
type ResponseHeaderEraseList struct {
	ResponseHeaderErase string `json:"response-header-erase,omitempty"`
}
type CompressionContentType struct {
	ContentType string `json:"content-type,omitempty"`
}
type CompressionExcludeContentType struct {
	ExcludeContentType string `json:"exclude-content-type,omitempty"`
}
type HTTPInstance struct {
	KeepClientAlive                     int                             `json:"keep-client-alive,omitempty"`
	CompressionAutoDisableOnHighCPU     int                             `json:"compression-auto-disable-on-high-cpu,omitempty"`
	ReqHdrWaitTime                      int                             `json:"req-hdr-wait-time,omitempty"`
	ExcludeURI                          []CompressionExcludeURI         `json:"compression-exclude-uri,omitempty"`
	CompressionEnable                   int                             `json:"compression-enable,omitempty"`
	CompressionKeepAcceptEncoding       int                             `json:"compression-keep-accept-encoding,omitempty"`
	FailoverURL                         string                          `json:"failover-url,omitempty"`
	RedirectSecurePort                  RedirectRewrite                 `json:"redirect-rewrite,omitempty"`
	RequestHeaderErase                  []RequestHeaderEraseList        `json:"request-header-erase-list,omitempty"`
	RdPort                              int                             `json:"rd-port,omitempty"`
	HostSwitchingType                   []HostSwitching                 `json:"host-switching,omitempty"`
	URLHashLast                         int                             `json:"url-hash-last,omitempty"`
	ClientIPHdrReplace                  int                             `json:"client-ip-hdr-replace,omitempty"`
	UseServerStatus                     int                             `json:"use-server-status,omitempty"`
	ReqHdrWaitTimeVal                   int                             `json:"req-hdr-wait-time-val,omitempty"`
	ResponseHeaderInsertType            []ResponseHeaderInsertList      `json:"response-header-insert-list,omitempty"`
	PersistOn401                        int                             `json:"persist-on-401,omitempty"`
	Redirect                            int                             `json:"redirect,omitempty"`
	InsertClientPort                    int                             `json:"insert-client-port,omitempty"`
	RetryOn5XxPerReqVal                 int                             `json:"retry-on-5xx-per-req-val,omitempty"`
	URLHashOffset                       int                             `json:"url-hash-offset,omitempty"`
	RdSimpleLoc                         string                          `json:"rd-simple-loc,omitempty"`
	LogRetry                            int                             `json:"log-retry,omitempty"`
	NonHTTPBypass                       int                             `json:"non-http-bypass,omitempty"`
	RetryOn5XxPerReq                    int                             `json:"retry-on-5xx-per-req,omitempty"`
	InsertClientIP                      int                             `json:"insert-client-ip,omitempty"`
	Logging                             Template2                       `json:"template,omitempty"`
	RequestTimeout                      int                             `json:"request-timeout,omitempty"`
	URLServiceGroup                     []URLSwitching                  `json:"url-switching,omitempty"`
	InsertClientPortHeaderName          string                          `json:"insert-client-port-header-name,omitempty"`
	StrictTransactionSwitch             int                             `json:"strict-transaction-switch,omitempty"`
	ResponseNewString                   []ResponseContentReplaceList    `json:"response-content-replace-list,omitempty"`
	One00ContWaitForReqComplete         int                             `json:"100-cont-wait-for-req-complete,omitempty"`
	MaxConcurrentStreams                int                             `json:"max-concurrent-streams,omitempty"`
	RequestHeaderInsert                 []RequestHeaderInsertList       `json:"request-header-insert-list,omitempty"`
	CompressionMinimumContentLength     int                             `json:"compression-minimum-content-length,omitempty"`
	CompressionLevel                    int                             `json:"compression-level,omitempty"`
	RequestLineCaseInsensitive          int                             `json:"request-line-case-insensitive,omitempty"`
	URLHashPersist                      int                             `json:"url-hash-persist,omitempty"`
	ResponseHeaderErase                 []ResponseHeaderEraseList       `json:"response-header-erase-list,omitempty"`
	UUID                                string                          `json:"uuid,omitempty"`
	BypassSg                            string                          `json:"bypass-sg,omitempty"`
	Name                                string                          `json:"name,omitempty"`
	RetryOn5XxVal                       int                             `json:"retry-on-5xx-val,omitempty"`
	URLHashFirst                        int                             `json:"url-hash-first,omitempty"`
	CompressionKeepAcceptEncodingEnable int                             `json:"compression-keep-accept-encoding-enable,omitempty"`
	UserTag                             string                          `json:"user-tag,omitempty"`
	ContentType                         []CompressionContentType        `json:"compression-content-type,omitempty"`
	ClientPortHdrReplace                int                             `json:"client-port-hdr-replace,omitempty"`
	InsertClientIPHeaderName            string                          `json:"insert-client-ip-header-name,omitempty"`
	RdSecure                            int                             `json:"rd-secure,omitempty"`
	RetryOn5Xx                          int                             `json:"retry-on-5xx,omitempty"`
	CookieFormat                        string                          `json:"cookie-format,omitempty"`
	Term11ClientHdrConnClose            int                             `json:"term-11client-hdr-conn-close,omitempty"`
	ExcludeContentType                  []CompressionExcludeContentType `json:"compression-exclude-content-type,omitempty"`
	RdRespCode                          string                          `json:"rd-resp-code,omitempty"`
}

func PostSlbTemplateHTTP(id string, inst HTTP, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateHTTP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/http", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HTTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func GetSlbTemplateHTTP(id string, name string, host string) (*HTTP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateHTTP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/http/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HTTP
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

func PutSlbTemplateHTTP(id string, name string, inst HTTP, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateHTTP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/http/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HTTP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}

}

func DeleteSlbTemplateHTTP(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateHTTP")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/http/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HTTP
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
