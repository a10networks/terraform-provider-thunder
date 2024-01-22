

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateHttp struct {
	Inst struct {

    Challenge DdosZoneTemplateHttpChallenge `json:"challenge"`
    ClientSourceIp DdosZoneTemplateHttpClientSourceIp `json:"client-source-ip"`
    Disable int `json:"disable"`
    DisallowConnectMethod int `json:"disallow-connect-method"`
    Dst DdosZoneTemplateHttpDst `json:"dst"`
    FilterList []DdosZoneTemplateHttpFilterList `json:"filter-list"`
    HttpTmplName string `json:"http-tmpl-name"`
    IdleTimeout DdosZoneTemplateHttpIdleTimeout `json:"idle-timeout"`
    MalformedHttp DdosZoneTemplateHttpMalformedHttp308 `json:"malformed-http"`
    MssTimeout DdosZoneTemplateHttpMssTimeout `json:"mss-timeout"`
    MultiPuThresholdDistribution DdosZoneTemplateHttpMultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    NonHttpBypass int `json:"non-http-bypass"`
    OutOfOrderQueueSize int `json:"out-of-order-queue-size" dval:"3"`
    OutOfOrderQueueTimeout int `json:"out-of-order-queue-timeout" dval:"3"`
    RequestHeader DdosZoneTemplateHttpRequestHeader `json:"request-header"`
    SlowRead DdosZoneTemplateHttpSlowRead `json:"slow-read"`
    Src DdosZoneTemplateHttpSrc `json:"src"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"http"`
}


type DdosZoneTemplateHttpChallenge struct {
    ChallengeMethod string `json:"challenge-method"`
    ChallengeRedirectCode string `json:"challenge-redirect-code" dval:"302"`
    ChallengeUriEncode int `json:"challenge-uri-encode"`
    ChallengeCookieName string `json:"challenge-cookie-name" dval:"sto-idd"`
    ChallengeKeepCookie int `json:"challenge-keep-cookie"`
    ChallengeInterval int `json:"challenge-interval" dval:"8"`
    ChallengePassActionListName string `json:"challenge-pass-action-list-name"`
    ChallengePassAction string `json:"challenge-pass-action"`
    ChallengeFailActionListName string `json:"challenge-fail-action-list-name"`
    ChallengeFailAction string `json:"challenge-fail-action" dval:"reset"`
}


type DdosZoneTemplateHttpClientSourceIp struct {
    ClientSourceIp int `json:"client-source-ip"`
    HttpHeaderName string `json:"http-header-name" dval:"X-Forwarded-For"`
}


type DdosZoneTemplateHttpDst struct {
    RateLimit DdosZoneTemplateHttpDstRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateHttpDstRateLimit struct {
    HttpPost DdosZoneTemplateHttpDstRateLimitHttpPost `json:"http-post"`
    HttpRequest DdosZoneTemplateHttpDstRateLimitHttpRequest `json:"http-request"`
    ResponseSize DdosZoneTemplateHttpDstRateLimitResponseSize `json:"response-size"`
}


type DdosZoneTemplateHttpDstRateLimitHttpPost struct {
    DstPostRateLimit int `json:"dst-post-rate-limit"`
    DstPostRateLimitActionListName string `json:"dst-post-rate-limit-action-list-name"`
    DstPostRateLimitAction string `json:"dst-post-rate-limit-action" dval:"drop"`
}


type DdosZoneTemplateHttpDstRateLimitHttpRequest struct {
    DstRequestRate int `json:"dst-request-rate"`
    DstRequestRateLimitActionListName string `json:"dst-request-rate-limit-action-list-name"`
    DstRequestRateLimitAction string `json:"dst-request-rate-limit-action" dval:"drop"`
}


type DdosZoneTemplateHttpDstRateLimitResponseSize struct {
    LessCfg []DdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg `json:"less-cfg"`
    GreaterCfg []DdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg `json:"greater-cfg"`
    BetweenCfg []DdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg `json:"between-cfg"`
    ResponseSizeActionListName string `json:"response-size-action-list-name"`
    ResponseSizeAction string `json:"response-size-action"`
}


type DdosZoneTemplateHttpDstRateLimitResponseSizeLessCfg struct {
    ObjLess int `json:"obj-less"`
    ObjLessRate int `json:"obj-less-rate"`
}


type DdosZoneTemplateHttpDstRateLimitResponseSizeGreaterCfg struct {
    ObjGreater int `json:"obj-greater"`
    ObjGreaterRate int `json:"obj-greater-rate"`
}


type DdosZoneTemplateHttpDstRateLimitResponseSizeBetweenCfg struct {
    ObjBetween1 int `json:"obj-between1"`
    ObjBetween2 int `json:"obj-between2"`
    ObjBetweenRate int `json:"obj-between-rate"`
}


type DdosZoneTemplateHttpFilterList struct {
    HttpFilterName string `json:"http-filter-name"`
    HttpFilterSeq int `json:"http-filter-seq"`
    HttpHeaderCfg DdosZoneTemplateHttpFilterListHttpHeaderCfg `json:"http-header-cfg"`
    HttpRefererCfg DdosZoneTemplateHttpFilterListHttpRefererCfg `json:"http-referer-cfg"`
    HttpAgentCfg DdosZoneTemplateHttpFilterListHttpAgentCfg `json:"http-agent-cfg"`
    HttpUriCfg DdosZoneTemplateHttpFilterListHttpUriCfg `json:"http-uri-cfg"`
    Dst DdosZoneTemplateHttpFilterListDst `json:"dst"`
    HttpFilterActionListName string `json:"http-filter-action-list-name"`
    HttpFilterAction string `json:"http-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosZoneTemplateHttpFilterListHttpHeaderCfg struct {
    HttpFilterHeaderRegex string `json:"http-filter-header-regex"`
    HttpFilterHeaderInverseMatch int `json:"http-filter-header-inverse-match"`
}


type DdosZoneTemplateHttpFilterListHttpRefererCfg struct {
    RefererEqualsCfg []DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg `json:"referer-equals-cfg"`
    RefererContainsCfg []DdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg `json:"referer-contains-cfg"`
    RefererStartsCfg []DdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg `json:"referer-starts-cfg"`
    RefererEndsCfg []DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg `json:"referer-ends-cfg"`
}


type DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEqualsCfg struct {
    HttpFilterRefererEquals string `json:"http-filter-referer-equals"`
}


type DdosZoneTemplateHttpFilterListHttpRefererCfgRefererContainsCfg struct {
    HttpFilterRefererContains string `json:"http-filter-referer-contains"`
}


type DdosZoneTemplateHttpFilterListHttpRefererCfgRefererStartsCfg struct {
    HttpFilterRefererStartsWith string `json:"http-filter-referer-starts-with"`
}


type DdosZoneTemplateHttpFilterListHttpRefererCfgRefererEndsCfg struct {
    HttpFilterRefererEndsWith string `json:"http-filter-referer-ends-with"`
}


type DdosZoneTemplateHttpFilterListHttpAgentCfg struct {
    AgentEqualsCfg []DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg `json:"agent-equals-cfg"`
    AgentContainsCfg []DdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg `json:"agent-contains-cfg"`
    AgentStartsCfg []DdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg `json:"agent-starts-cfg"`
    AgentEndsCfg []DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg `json:"agent-ends-cfg"`
}


type DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEqualsCfg struct {
    HttpFilterAgentEquals string `json:"http-filter-agent-equals"`
}


type DdosZoneTemplateHttpFilterListHttpAgentCfgAgentContainsCfg struct {
    HttpFilterAgentContains string `json:"http-filter-agent-contains"`
}


type DdosZoneTemplateHttpFilterListHttpAgentCfgAgentStartsCfg struct {
    HttpFilterAgentStartsWith string `json:"http-filter-agent-starts-with"`
}


type DdosZoneTemplateHttpFilterListHttpAgentCfgAgentEndsCfg struct {
    HttpFilterAgentEndsWith string `json:"http-filter-agent-ends-with"`
}


type DdosZoneTemplateHttpFilterListHttpUriCfg struct {
    UriEqualCfg []DdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg `json:"uri-equal-cfg"`
    UriContainsCfg []DdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg `json:"uri-contains-cfg"`
    UriStartsCfg []DdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg `json:"uri-starts-cfg"`
    UriEndsCfg []DdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg `json:"uri-ends-cfg"`
}


type DdosZoneTemplateHttpFilterListHttpUriCfgUriEqualCfg struct {
    HttpFilterUriEquals string `json:"http-filter-uri-equals"`
}


type DdosZoneTemplateHttpFilterListHttpUriCfgUriContainsCfg struct {
    HttpFilterUriContains string `json:"http-filter-uri-contains"`
}


type DdosZoneTemplateHttpFilterListHttpUriCfgUriStartsCfg struct {
    HttpFilterUriStartsWith string `json:"http-filter-uri-starts-with"`
}


type DdosZoneTemplateHttpFilterListHttpUriCfgUriEndsCfg struct {
    HttpFilterUriEndsWith string `json:"http-filter-uri-ends-with"`
}


type DdosZoneTemplateHttpFilterListDst struct {
    HttpFilterRateLimit int `json:"http-filter-rate-limit"`
}


type DdosZoneTemplateHttpIdleTimeout struct {
    IdleTimeoutValue int `json:"idle-timeout-value"`
    IgnoreZeroPayload int `json:"ignore-zero-payload"`
    IdleTimeoutActionListName string `json:"idle-timeout-action-list-name"`
    IdleTimeoutAction string `json:"idle-timeout-action"`
}


type DdosZoneTemplateHttpMalformedHttp308 struct {
    MalformedHttp string `json:"malformed-http" dval:"check"`
    MalformedHttpMaxLineSize int `json:"malformed-http-max-line-size" dval:"32512"`
    MalformedHttpMaxNumHeaders int `json:"malformed-http-max-num-headers" dval:"90"`
    MalformedHttpMaxReqLineSize int `json:"malformed-http-max-req-line-size" dval:"32512"`
    MalformedHttpMaxHeaderNameSize int `json:"malformed-http-max-header-name-size" dval:"64"`
    MalformedHttpMaxContentLength int `json:"malformed-http-max-content-length" dval:"4294967295"`
    MalformedHttpBadChunkMonEnabled int `json:"malformed-http-bad-chunk-mon-enabled"`
    MalformedHttpActionListName string `json:"malformed-http-action-list-name"`
    MalformedHttpAction string `json:"malformed-http-action"`
    Uuid string `json:"uuid"`
}


type DdosZoneTemplateHttpMssTimeout struct {
    MssPercent int `json:"mss-percent"`
    NumberPackets int `json:"number-packets"`
    MssTimeoutActionListName string `json:"mss-timeout-action-list-name"`
    MssTimeoutAction string `json:"mss-timeout-action"`
}


type DdosZoneTemplateHttpMultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosZoneTemplateHttpRequestHeader struct {
    Timeout int `json:"timeout"`
    HeaderTimeoutActionListName string `json:"header-timeout-action-list-name"`
    HeaderTimeoutAction string `json:"header-timeout-action" dval:"drop"`
}


type DdosZoneTemplateHttpSlowRead struct {
    MinWindowSize int `json:"min-window-size"`
    MinWindowCount int `json:"min-window-count"`
    SlowReadActionListName string `json:"slow-read-action-list-name"`
    SlowReadAction string `json:"slow-read-action"`
}


type DdosZoneTemplateHttpSrc struct {
    RateLimit DdosZoneTemplateHttpSrcRateLimit `json:"rate-limit"`
}


type DdosZoneTemplateHttpSrcRateLimit struct {
    HttpPost DdosZoneTemplateHttpSrcRateLimitHttpPost `json:"http-post"`
    HttpRequest DdosZoneTemplateHttpSrcRateLimitHttpRequest `json:"http-request"`
}


type DdosZoneTemplateHttpSrcRateLimitHttpPost struct {
    SrcPostRateLimit int `json:"src-post-rate-limit"`
    SrcPostRateLimitActionListName string `json:"src-post-rate-limit-action-list-name"`
    SrcPostRateLimitAction string `json:"src-post-rate-limit-action" dval:"drop"`
}


type DdosZoneTemplateHttpSrcRateLimitHttpRequest struct {
    SrcRequestRate int `json:"src-request-rate"`
    SrcRequestRateLimitActionListName string `json:"src-request-rate-limit-action-list-name"`
    SrcRequestRateLimitAction string `json:"src-request-rate-limit-action" dval:"drop"`
}

func (p *DdosZoneTemplateHttp) GetId() string{
    return url.QueryEscape(p.Inst.HttpTmplName)
}

func (p *DdosZoneTemplateHttp) getPath() string{
    return "ddos/zone-template/http"
}

func (p *DdosZoneTemplateHttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttp::Post")
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

func (p *DdosZoneTemplateHttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttp::Get")
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
func (p *DdosZoneTemplateHttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttp::Put")
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

func (p *DdosZoneTemplateHttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
