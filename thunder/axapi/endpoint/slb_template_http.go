

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateHttp struct {
	Inst struct {

    AllowedMethods string `json:"allowed-methods"`
    AllowedMethodsAction string `json:"allowed-methods-action" dval:"drop"`
    BypassSg string `json:"bypass-sg"`
    ClientIdleTimeout int `json:"client-idle-timeout"`
    ClientIpHdrReplace int `json:"client-ip-hdr-replace"`
    ClientPortHdrReplace int `json:"client-port-hdr-replace"`
    CompressionAutoDisableOnHighCpu int `json:"compression-auto-disable-on-high-cpu"`
    CompressionBrLevel int `json:"compression-br-level" dval:"1"`
    CompressionBrSlidingWindowSize int `json:"compression-br-sliding-window-size"`
    CompressionContentType []SlbTemplateHttpCompressionContentType `json:"compression-content-type"`
    CompressionEnable int `json:"compression-enable"`
    CompressionExcludeContentType []SlbTemplateHttpCompressionExcludeContentType `json:"compression-exclude-content-type"`
    CompressionExcludeUri []SlbTemplateHttpCompressionExcludeUri `json:"compression-exclude-uri"`
    CompressionKeepAcceptEncoding int `json:"compression-keep-accept-encoding"`
    CompressionKeepAcceptEncodingEnable int `json:"compression-keep-accept-encoding-enable"`
    CompressionLevel int `json:"compression-level" dval:"1"`
    CompressionMethodOrder string `json:"compression-method-order"`
    CompressionMinimumContentLength int `json:"compression-minimum-content-length" dval:"120"`
    ContWaitForReqComplete100 int `json:"cont-wait-for-req-complete100"`
    CookieFormat string `json:"cookie-format"`
    CookieSamesite string `json:"cookie-samesite"`
    DefaultCharset string `json:"default-charset" dval:"utf-8"`
    DisallowedMethods string `json:"disallowed-methods"`
    DisallowedMethodsAction string `json:"disallowed-methods-action" dval:"drop"`
    FailoverUrl string `json:"failover-url"`
    FrameLimit int `json:"frame-limit" dval:"10000"`
    HostSwitching []SlbTemplateHttpHostSwitching `json:"host-switching"`
    HttpProtocolCheck SlbTemplateHttpHttpProtocolCheck1447 `json:"http-protocol-check"`
    Http2ClientNoSnat int `json:"http2-client-no-snat"`
    InsertClientIp int `json:"insert-client-ip"`
    InsertClientIpHeaderName string `json:"insert-client-ip-header-name"`
    InsertClientPort int `json:"insert-client-port"`
    InsertClientPortHeaderName string `json:"insert-client-port-header-name"`
    KeepClientAlive int `json:"keep-client-alive"`
    LogRetry int `json:"log-retry"`
    MaxConcurrentStreams int `json:"max-concurrent-streams" dval:"50"`
    Name string `json:"name"`
    NonHttpBypass int `json:"non-http-bypass"`
    PersistOn401 int `json:"persist-on-401"`
    Prefix string `json:"prefix"`
    RdPort int `json:"rd-port"`
    RdRespCode string `json:"rd-resp-code"`
    RdSecure int `json:"rd-secure"`
    RdSimpleLoc string `json:"rd-simple-loc"`
    Redirect int `json:"redirect"`
    RedirectRewrite SlbTemplateHttpRedirectRewrite `json:"redirect-rewrite"`
    ReqHdrWaitTime int `json:"req-hdr-wait-time"`
    ReqHdrWaitTimeVal int `json:"req-hdr-wait-time-val" dval:"7"`
    RequestHeaderEraseList []SlbTemplateHttpRequestHeaderEraseList `json:"request-header-erase-list"`
    RequestHeaderInsertList []SlbTemplateHttpRequestHeaderInsertList `json:"request-header-insert-list"`
    RequestLineCaseInsensitive int `json:"request-line-case-insensitive"`
    RequestTimeout int `json:"request-timeout"`
    ResponseContentReplaceList []SlbTemplateHttpResponseContentReplaceList `json:"response-content-replace-list"`
    ResponseHeaderEraseList []SlbTemplateHttpResponseHeaderEraseList `json:"response-header-erase-list"`
    ResponseHeaderInsertList []SlbTemplateHttpResponseHeaderInsertList `json:"response-header-insert-list"`
    RetryOn5xx int `json:"retry-on-5xx"`
    RetryOn5xxPerReq int `json:"retry-on-5xx-per-req"`
    RetryOn5xxPerReqVal int `json:"retry-on-5xx-per-req-val" dval:"3"`
    RetryOn5xxVal int `json:"retry-on-5xx-val" dval:"3"`
    ServerSupportHttp2Only int `json:"server-support-http2-only"`
    ServerSupportHttp2OnlyValue string `json:"server-support-http2-only-value" dval:"auto-detect"`
    StreamCancellationLimit int `json:"stream-cancellation-limit"`
    StreamCancellationRate int `json:"stream-cancellation-rate" dval:"10"`
    StrictTransactionSwitch int `json:"strict-transaction-switch"`
    Template SlbTemplateHttpTemplate `json:"template"`
    Term11clientHdrConnClose int `json:"term-11client-hdr-conn-close"`
    UrlHashFirst int `json:"url-hash-first"`
    UrlHashLast int `json:"url-hash-last"`
    UrlHashOffset int `json:"url-hash-offset"`
    UrlHashPersist int `json:"url-hash-persist"`
    UrlSwitching []SlbTemplateHttpUrlSwitching `json:"url-switching"`
    UseServerStatus int `json:"use-server-status"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"http"`
}


type SlbTemplateHttpCompressionContentType struct {
    ContentType string `json:"content-type"`
}


type SlbTemplateHttpCompressionExcludeContentType struct {
    ExcludeContentType string `json:"exclude-content-type"`
}


type SlbTemplateHttpCompressionExcludeUri struct {
    ExcludeUri string `json:"exclude-uri"`
}


type SlbTemplateHttpHostSwitching struct {
    HostSwitchingType string `json:"host-switching-type"`
    HostMatchString string `json:"host-match-string"`
    HostServiceGroup string `json:"host-service-group"`
}


type SlbTemplateHttpHttpProtocolCheck1447 struct {
    H2upContentLengthAlias string `json:"h2up-content-length-alias"`
    MalformedH2upHeaderValue string `json:"malformed-h2up-header-value"`
    MalformedH2upSchemeValue string `json:"malformed-h2up-scheme-value"`
    H2upWithTransferEncoding string `json:"h2up-with-transfer-encoding"`
    MultipleContentLength string `json:"multiple-content-length"`
    MultipleTransferEncoding string `json:"multiple-transfer-encoding"`
    TransferEncodingAndContentLength string `json:"transfer-encoding-and-content-length"`
    GetAndPayload string `json:"get-and-payload"`
    H2upWithHostAndAuth string `json:"h2up-with-host-and-auth"`
    Uuid string `json:"uuid"`
    HeaderFilterRuleList []SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448 `json:"header-filter-rule-list"`
}


type SlbTemplateHttpHttpProtocolCheckHeaderFilterRuleList1448 struct {
    SeqNum int `json:"seq-num"`
    MatchTypeValue string `json:"match-type-value"`
    HeaderNameValue string `json:"header-name-value"`
    HeaderValueValue string `json:"header-value-value"`
    ActionValue string `json:"action-value"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SlbTemplateHttpRedirectRewrite struct {
    MatchList []SlbTemplateHttpRedirectRewriteMatchList `json:"match-list"`
    RedirectSecure int `json:"redirect-secure"`
    RedirectSecurePort int `json:"redirect-secure-port" dval:"443"`
}


type SlbTemplateHttpRedirectRewriteMatchList struct {
    RedirectMatch string `json:"redirect-match"`
    RewriteTo string `json:"rewrite-to"`
}


type SlbTemplateHttpRequestHeaderEraseList struct {
    RequestHeaderErase string `json:"request-header-erase"`
}


type SlbTemplateHttpRequestHeaderInsertList struct {
    RequestHeaderInsert string `json:"request-header-insert"`
    RequestHeaderInsertType string `json:"request-header-insert-type"`
}


type SlbTemplateHttpResponseContentReplaceList struct {
    ResponseContentReplace string `json:"response-content-replace"`
    ResponseNewString string `json:"response-new-string"`
}


type SlbTemplateHttpResponseHeaderEraseList struct {
    ResponseHeaderErase string `json:"response-header-erase"`
}


type SlbTemplateHttpResponseHeaderInsertList struct {
    ResponseHeaderInsert string `json:"response-header-insert"`
    ResponseHeaderInsertType string `json:"response-header-insert-type"`
}


type SlbTemplateHttpTemplate struct {
    Logging string `json:"logging"`
}


type SlbTemplateHttpUrlSwitching struct {
    UrlSwitchingType string `json:"url-switching-type"`
    UrlMatchString string `json:"url-match-string"`
    UrlServiceGroup string `json:"url-service-group"`
}

func (p *SlbTemplateHttp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateHttp) getPath() string{
    return "slb/template/http"
}

func (p *SlbTemplateHttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttp::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *SlbTemplateHttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttp::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SlbTemplateHttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttp::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *SlbTemplateHttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
