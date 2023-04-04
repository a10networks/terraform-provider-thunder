package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateHttp struct {
	SlbTemplateHTTPInstanceName SlbTemplateHTTPInstance `json:"http,omitempty"`
}

type SlbTemplateHTTPInstance struct {
	SlbTemplateHTTPInstanceBypassSg                                         string                                                 `json:"bypass-sg,omitempty"`
	SlbTemplateHTTPInstanceClientIPHdrReplace                               int                                                    `json:"client-ip-hdr-replace,omitempty"`
	SlbTemplateHTTPInstanceClientPortHdrReplace                             int                                                    `json:"client-port-hdr-replace,omitempty"`
	SlbTemplateHTTPInstanceCompressionAutoDisableOnHighCPU                  int                                                    `json:"compression-auto-disable-on-high-cpu,omitempty"`
	SlbTemplateHTTPInstanceCompressionContentTypeContentType                []SlbTemplateHTTPInstanceCompressionContentType        `json:"compression-content-type,omitempty"`
	SlbTemplateHTTPInstanceCompressionEnable                                int                                                    `json:"compression-enable,omitempty"`
	SlbTemplateHTTPInstanceCompressionExcludeContentTypeExcludeContentType  []SlbTemplateHTTPInstanceCompressionExcludeContentType `json:"compression-exclude-content-type,omitempty"`
	SlbTemplateHTTPInstanceCompressionExcludeURIExcludeURI                  []SlbTemplateHTTPInstanceCompressionExcludeURI         `json:"compression-exclude-uri,omitempty"`
	SlbTemplateHTTPInstanceCompressionKeepAcceptEncoding                    int                                                    `json:"compression-keep-accept-encoding,omitempty"`
	SlbTemplateHTTPInstanceCompressionKeepAcceptEncodingEnable              int                                                    `json:"compression-keep-accept-encoding-enable,omitempty"`
	SlbTemplateHTTPInstanceCompressionLevel                                 int                                                    `json:"compression-level,omitempty"`
	SlbTemplateHTTPInstanceCompressionMinimumContentLength                  int                                                    `json:"compression-minimum-content-length,omitempty"`
	SlbTemplateHTTPInstanceCookieFormat                                     string                                                 `json:"cookie-format,omitempty"`
	SlbTemplateHTTPInstanceCookieSamesite                                   string                                                 `json:"cookie-samesite,omitempty"`
	SlbTemplateHTTPInstanceDefaultCharset                                   string                                                 `json:"default-charset,omitempty"`
	SlbTemplateHTTPInstanceFailoverUrl                                      string                                                 `json:"failover-url,omitempty"`
	SlbTemplateHTTPInstanceFrameLimit                                       int                                                    `json:"frame-limit,omitempty"`
	SlbTemplateHTTPInstanceHostSwitchingHostSwitchingType                   []SlbTemplateHTTPInstanceHostSwitching                 `json:"host-switching,omitempty"`
	SlbTemplateHTTPInstanceInsertClientIP                                   int                                                    `json:"insert-client-ip,omitempty"`
	SlbTemplateHTTPInstanceInsertClientIPHeaderName                         string                                                 `json:"insert-client-ip-header-name,omitempty"`
	SlbTemplateHTTPInstanceInsertClientPort                                 int                                                    `json:"insert-client-port,omitempty"`
	SlbTemplateHTTPInstanceInsertClientPortHeaderName                       string                                                 `json:"insert-client-port-header-name,omitempty"`
	SlbTemplateHTTPInstanceKeepClientAlive                                  int                                                    `json:"keep-client-alive,omitempty"`
	SlbTemplateHTTPInstanceLogRetry                                         int                                                    `json:"log-retry,omitempty"`
	SlbTemplateHTTPInstanceMaxConcurrentStreams                             int                                                    `json:"max-concurrent-streams,omitempty"`
	SlbTemplateHTTPInstanceName                                             string                                                 `json:"name,omitempty"`
	SlbTemplateHTTPInstanceNonHTTPBypass                                    int                                                    `json:"non-http-bypass,omitempty"`
	SlbTemplateHTTPInstanceNum100ContWaitForReqComplete                     int                                                    `json:"100-cont-wait-for-req-complete,omitempty"`
	SlbTemplateHTTPInstancePersistOn401                                     int                                                    `json:"persist-on-401,omitempty"`
	SlbTemplateHTTPInstancePrefix                                           string                                                 `json:"prefix,omitempty"`
	SlbTemplateHTTPInstanceRdPort                                           int                                                    `json:"rd-port,omitempty"`
	SlbTemplateHTTPInstanceRdRespCode                                       string                                                 `json:"rd-resp-code,omitempty"`
	SlbTemplateHTTPInstanceRdSecure                                         int                                                    `json:"rd-secure,omitempty"`
	SlbTemplateHTTPInstanceRdSimpleLoc                                      string                                                 `json:"rd-simple-loc,omitempty"`
	SlbTemplateHTTPInstanceRedirect                                         int                                                    `json:"redirect,omitempty"`
	SlbTemplateHTTPInstanceRedirectRewriteMatchList                         SlbTemplateHTTPInstanceRedirectRewrite                 `json:"redirect-rewrite,omitempty"`
	SlbTemplateHTTPInstanceReqHdrWaitTime                                   int                                                    `json:"req-hdr-wait-time,omitempty"`
	SlbTemplateHTTPInstanceReqHdrWaitTimeVal                                int                                                    `json:"req-hdr-wait-time-val,omitempty"`
	SlbTemplateHTTPInstanceRequestHeaderEraseListRequestHeaderErase         []SlbTemplateHTTPInstanceRequestHeaderEraseList        `json:"request-header-erase-list,omitempty"`
	SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsert       []SlbTemplateHTTPInstanceRequestHeaderInsertList       `json:"request-header-insert-list,omitempty"`
	SlbTemplateHTTPInstanceRequestLineCaseInsensitive                       int                                                    `json:"request-line-case-insensitive,omitempty"`
	SlbTemplateHTTPInstanceRequestTimeout                                   int                                                    `json:"request-timeout,omitempty"`
	SlbTemplateHTTPInstanceResponseContentReplaceListResponseContentReplace []SlbTemplateHTTPInstanceResponseContentReplaceList    `json:"response-content-replace-list,omitempty"`
	SlbTemplateHTTPInstanceResponseHeaderEraseListResponseHeaderErase       []SlbTemplateHTTPInstanceResponseHeaderEraseList       `json:"response-header-erase-list,omitempty"`
	SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsert     []SlbTemplateHTTPInstanceResponseHeaderInsertList      `json:"response-header-insert-list,omitempty"`
	SlbTemplateHTTPInstanceRetryOn5Xx                                       int                                                    `json:"retry-on-5xx,omitempty"`
	SlbTemplateHTTPInstanceRetryOn5XxPerReq                                 int                                                    `json:"retry-on-5xx-per-req,omitempty"`
	SlbTemplateHTTPInstanceRetryOn5XxPerReqVal                              int                                                    `json:"retry-on-5xx-per-req-val,omitempty"`
	SlbTemplateHTTPInstanceRetryOn5XxVal                                    int                                                    `json:"retry-on-5xx-val,omitempty"`
	SlbTemplateHTTPInstanceStrictTransactionSwitch                          int                                                    `json:"strict-transaction-switch,omitempty"`
	SlbTemplateHTTPInstanceTemplateLogging                                  SlbTemplateHTTPInstanceTemplate                        `json:"template,omitempty"`
	SlbTemplateHTTPInstanceTerm11ClientHdrConnClose                         int                                                    `json:"term-11client-hdr-conn-close,omitempty"`
	SlbTemplateHTTPInstanceUUID                                             string                                                 `json:"uuid,omitempty"`
	SlbTemplateHTTPInstanceUrlHashFirst                                     int                                                    `json:"url-hash-first,omitempty"`
	SlbTemplateHTTPInstanceUrlHashLast                                      int                                                    `json:"url-hash-last,omitempty"`
	SlbTemplateHTTPInstanceUrlHashOffset                                    int                                                    `json:"url-hash-offset,omitempty"`
	SlbTemplateHTTPInstanceUrlHashPersist                                   int                                                    `json:"url-hash-persist,omitempty"`
	SlbTemplateHTTPInstanceUrlSwitchingUrlSwitchingType                     []SlbTemplateHTTPInstanceUrlSwitching                  `json:"url-switching,omitempty"`
	SlbTemplateHTTPInstanceUseServerStatus                                  int                                                    `json:"use-server-status,omitempty"`
	SlbTemplateHTTPInstanceUserTag                                          string                                                 `json:"user-tag,omitempty"`
}

type SlbTemplateHTTPInstanceCompressionContentType struct {
	SlbTemplateHTTPInstanceCompressionContentTypeContentType string `json:"content-type,omitempty"`
}

type SlbTemplateHTTPInstanceCompressionExcludeContentType struct {
	SlbTemplateHTTPInstanceCompressionExcludeContentTypeExcludeContentType string `json:"exclude-content-type,omitempty"`
}

type SlbTemplateHTTPInstanceCompressionExcludeURI struct {
	SlbTemplateHTTPInstanceCompressionExcludeURIExcludeURI string `json:"exclude-uri,omitempty"`
}

type SlbTemplateHTTPInstanceHostSwitching struct {
	SlbTemplateHTTPInstanceHostSwitchingHostMatchString   string `json:"host-match-string,omitempty"`
	SlbTemplateHTTPInstanceHostSwitchingHostServiceGroup  string `json:"host-service-group,omitempty"`
	SlbTemplateHTTPInstanceHostSwitchingHostSwitchingType string `json:"host-switching-type,omitempty"`
}

type SlbTemplateHTTPInstanceRedirectRewrite struct {
	SlbTemplateHTTPInstanceRedirectRewriteMatchListRedirectMatch []SlbTemplateHTTPInstanceRedirectRewriteMatchList `json:"match-list,omitempty"`
	SlbTemplateHTTPInstanceRedirectRewriteRedirectSecure         int                                               `json:"redirect-secure,omitempty"`
	SlbTemplateHTTPInstanceRedirectRewriteRedirectSecurePort     int                                               `json:"redirect-secure-port,omitempty"`
}

type SlbTemplateHTTPInstanceRequestHeaderEraseList struct {
	SlbTemplateHTTPInstanceRequestHeaderEraseListRequestHeaderErase string `json:"request-header-erase,omitempty"`
}

type SlbTemplateHTTPInstanceRequestHeaderInsertList struct {
	SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsert     string `json:"request-header-insert,omitempty"`
	SlbTemplateHTTPInstanceRequestHeaderInsertListRequestHeaderInsertType string `json:"request-header-insert-type,omitempty"`
}

type SlbTemplateHTTPInstanceResponseContentReplaceList struct {
	SlbTemplateHTTPInstanceResponseContentReplaceListResponseContentReplace string `json:"response-content-replace,omitempty"`
	SlbTemplateHTTPInstanceResponseContentReplaceListResponseNewString      string `json:"response-new-string,omitempty"`
}

type SlbTemplateHTTPInstanceResponseHeaderEraseList struct {
	SlbTemplateHTTPInstanceResponseHeaderEraseListResponseHeaderErase string `json:"response-header-erase,omitempty"`
}

type SlbTemplateHTTPInstanceResponseHeaderInsertList struct {
	SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsert     string `json:"response-header-insert,omitempty"`
	SlbTemplateHTTPInstanceResponseHeaderInsertListResponseHeaderInsertType string `json:"response-header-insert-type,omitempty"`
}

type SlbTemplateHTTPInstanceTemplate struct {
	SlbTemplateHTTPInstanceTemplateLogging string `json:"logging,omitempty"`
}

type SlbTemplateHTTPInstanceUrlSwitching struct {
	SlbTemplateHTTPInstanceUrlSwitchingUrlMatchString   string `json:"url-match-string,omitempty"`
	SlbTemplateHTTPInstanceUrlSwitchingUrlServiceGroup  string `json:"url-service-group,omitempty"`
	SlbTemplateHTTPInstanceUrlSwitchingUrlSwitchingType string `json:"url-switching-type,omitempty"`
}

type SlbTemplateHTTPInstanceRedirectRewriteMatchList struct {
	SlbTemplateHTTPInstanceRedirectRewriteMatchListRedirectMatch string `json:"redirect-match,omitempty"`
	SlbTemplateHTTPInstanceRedirectRewriteMatchListRewriteTo     string `json:"rewrite-to,omitempty"`
}

func PostSlbTemplateHttp(id string, inst SlbTemplateHttp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateHttp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/http", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateHttp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateHttp(id string, name1 string, host string) (*SlbTemplateHttp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateHttp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/http/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateHttp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateHttp(id string, name1 string, inst SlbTemplateHttp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateHttp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/http/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateHttp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateHttp(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateHttp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/http/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateHttp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateHttp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
