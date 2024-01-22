

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateHttp struct {
	Inst struct {

    Action string `json:"action" dval:"drop"`
    AgentFilter DdosTemplateHttpAgentFilter `json:"agent-filter"`
    ChallengeCookieName string `json:"challenge-cookie-name" dval:"sto-idd"`
    ChallengeInterval int `json:"challenge-interval" dval:"8"`
    ChallengeKeepCookie int `json:"challenge-keep-cookie"`
    ChallengeMethod string `json:"challenge-method"`
    ChallengeRedirectCode string `json:"challenge-redirect-code" dval:"302"`
    ChallengeUriEncode int `json:"challenge-uri-encode"`
    Disable int `json:"disable"`
    DisallowConnectMethod int `json:"disallow-connect-method"`
    FilterHeaderList []DdosTemplateHttpFilterHeaderList `json:"filter-header-list"`
    HttpTmplName string `json:"http-tmpl-name"`
    IdleTimeout int `json:"idle-timeout"`
    IgnoreZeroPayload int `json:"ignore-zero-payload"`
    MalformedHttp DdosTemplateHttpMalformedHttp `json:"malformed-http"`
    MssCfg DdosTemplateHttpMssCfg `json:"mss-cfg"`
    MultiPuThresholdDistribution DdosTemplateHttpMultiPuThresholdDistribution `json:"multi-pu-threshold-distribution"`
    NonHttpBypass int `json:"non-http-bypass"`
    OutOfOrderQueueSize int `json:"out-of-order-queue-size" dval:"3"`
    OutOfOrderQueueTimeout int `json:"out-of-order-queue-timeout" dval:"3"`
    PostRateLimit int `json:"post-rate-limit"`
    RefererFilter DdosTemplateHttpRefererFilter `json:"referer-filter"`
    RequestHeader DdosTemplateHttpRequestHeader `json:"request-header"`
    RequestRateLimit DdosTemplateHttpRequestRateLimit `json:"request-rate-limit"`
    ResponseRateLimit DdosTemplateHttpResponseRateLimit `json:"response-rate-limit"`
    SlowReadDrop DdosTemplateHttpSlowReadDrop `json:"slow-read-drop"`
    UseHdrIpCfg DdosTemplateHttpUseHdrIpCfg `json:"use-hdr-ip-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"http"`
}


type DdosTemplateHttpAgentFilter struct {
    AgentFilterBlacklist int `json:"agent-filter-blacklist"`
    AgentEqualsCfg []DdosTemplateHttpAgentFilterAgentEqualsCfg `json:"agent-equals-cfg"`
    AgentContainsCfg []DdosTemplateHttpAgentFilterAgentContainsCfg `json:"agent-contains-cfg"`
    AgentStartsCfg []DdosTemplateHttpAgentFilterAgentStartsCfg `json:"agent-starts-cfg"`
    AgentEndsCfg []DdosTemplateHttpAgentFilterAgentEndsCfg `json:"agent-ends-cfg"`
}


type DdosTemplateHttpAgentFilterAgentEqualsCfg struct {
    AgentEquals string `json:"agent-equals"`
}


type DdosTemplateHttpAgentFilterAgentContainsCfg struct {
    AgentContains string `json:"agent-contains"`
}


type DdosTemplateHttpAgentFilterAgentStartsCfg struct {
    AgentStartsWith string `json:"agent-starts-with"`
}


type DdosTemplateHttpAgentFilterAgentEndsCfg struct {
    AgentEndsWith string `json:"agent-ends-with"`
}


type DdosTemplateHttpFilterHeaderList struct {
    HttpFilterHeaderSeq int `json:"http-filter-header-seq"`
    HttpFilterHeaderRegex string `json:"http-filter-header-regex"`
    HttpFilterHeaderUnmatched int `json:"http-filter-header-unmatched"`
    HttpFilterHeaderBlacklist int `json:"http-filter-header-blacklist"`
    HttpFilterHeaderWhitelist int `json:"http-filter-header-whitelist"`
    HttpFilterHeaderCountOnly int `json:"http-filter-header-count-only"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosTemplateHttpMalformedHttp struct {
    MalformedHttpEnabled int `json:"malformed-http-enabled"`
    MalformedHttpMaxLineSize int `json:"malformed-http-max-line-size" dval:"32512"`
    MalformedHttpMaxNumHeaders int `json:"malformed-http-max-num-headers" dval:"90"`
    MalformedHttpMaxReqLineSize int `json:"malformed-http-max-req-line-size" dval:"32512"`
    MalformedHttpMaxHeaderNameSize int `json:"malformed-http-max-header-name-size" dval:"64"`
    MalformedHttpMaxContentLength int `json:"malformed-http-max-content-length" dval:"4294967295"`
    MalformedHttpBadChunkMonEnabled int `json:"malformed-http-bad-chunk-mon-enabled"`
}


type DdosTemplateHttpMssCfg struct {
    MssTimeout int `json:"mss-timeout"`
    MssPercent int `json:"mss-percent"`
    NumberPackets int `json:"number-packets"`
}


type DdosTemplateHttpMultiPuThresholdDistribution struct {
    MultiPuThresholdDistributionValue int `json:"multi-pu-threshold-distribution-value"`
    MultiPuThresholdDistributionDisable string `json:"multi-pu-threshold-distribution-disable"`
}


type DdosTemplateHttpRefererFilter struct {
    RefFilterBlacklist int `json:"ref-filter-blacklist"`
    RefererEqualsCfg []DdosTemplateHttpRefererFilterRefererEqualsCfg `json:"referer-equals-cfg"`
    RefererContainsCfg []DdosTemplateHttpRefererFilterRefererContainsCfg `json:"referer-contains-cfg"`
    RefererStartsCfg []DdosTemplateHttpRefererFilterRefererStartsCfg `json:"referer-starts-cfg"`
    RefererEndsCfg []DdosTemplateHttpRefererFilterRefererEndsCfg `json:"referer-ends-cfg"`
}


type DdosTemplateHttpRefererFilterRefererEqualsCfg struct {
    RefererEquals string `json:"referer-equals"`
}


type DdosTemplateHttpRefererFilterRefererContainsCfg struct {
    RefererContains string `json:"referer-contains"`
}


type DdosTemplateHttpRefererFilterRefererStartsCfg struct {
    RefererStartsWith string `json:"referer-starts-with"`
}


type DdosTemplateHttpRefererFilterRefererEndsCfg struct {
    RefererEndsWith string `json:"referer-ends-with"`
}


type DdosTemplateHttpRequestHeader struct {
    Timeout int `json:"timeout"`
}


type DdosTemplateHttpRequestRateLimit struct {
    RequestRate int `json:"request-rate"`
    Uri []DdosTemplateHttpRequestRateLimitUri `json:"uri"`
}


type DdosTemplateHttpRequestRateLimitUri struct {
    EqualCfg DdosTemplateHttpRequestRateLimitUriEqualCfg `json:"equal-cfg"`
    ContainsCfg DdosTemplateHttpRequestRateLimitUriContainsCfg `json:"contains-cfg"`
    StartsCfg DdosTemplateHttpRequestRateLimitUriStartsCfg `json:"starts-cfg"`
    EndsCfg DdosTemplateHttpRequestRateLimitUriEndsCfg `json:"ends-cfg"`
}


type DdosTemplateHttpRequestRateLimitUriEqualCfg struct {
    UrlEquals string `json:"url-equals"`
    UrlEqualsRate int `json:"url-equals-rate"`
}


type DdosTemplateHttpRequestRateLimitUriContainsCfg struct {
    UrlContains string `json:"url-contains"`
    UrlContainsRate int `json:"url-contains-rate"`
}


type DdosTemplateHttpRequestRateLimitUriStartsCfg struct {
    UrlStartsWith string `json:"url-starts-with"`
    UrlStartsWithRate int `json:"url-starts-with-rate"`
}


type DdosTemplateHttpRequestRateLimitUriEndsCfg struct {
    UrlEndsWith string `json:"url-ends-with"`
    UrlEndsWithRate int `json:"url-ends-with-rate"`
}


type DdosTemplateHttpResponseRateLimit struct {
    ObjSize DdosTemplateHttpResponseRateLimitObjSize `json:"obj-size"`
}


type DdosTemplateHttpResponseRateLimitObjSize struct {
    LessCfg []DdosTemplateHttpResponseRateLimitObjSizeLessCfg `json:"less-cfg"`
    GreaterCfg []DdosTemplateHttpResponseRateLimitObjSizeGreaterCfg `json:"greater-cfg"`
    BetweenCfg []DdosTemplateHttpResponseRateLimitObjSizeBetweenCfg `json:"between-cfg"`
}


type DdosTemplateHttpResponseRateLimitObjSizeLessCfg struct {
    ObjLess int `json:"obj-less"`
    ObjLessRate int `json:"obj-less-rate"`
}


type DdosTemplateHttpResponseRateLimitObjSizeGreaterCfg struct {
    ObjGreater int `json:"obj-greater"`
    ObjGreaterRate int `json:"obj-greater-rate"`
}


type DdosTemplateHttpResponseRateLimitObjSizeBetweenCfg struct {
    ObjBetween1 int `json:"obj-between1"`
    ObjBetween2 int `json:"obj-between2"`
    ObjBetweenRate int `json:"obj-between-rate"`
}


type DdosTemplateHttpSlowReadDrop struct {
    MinWindowSize int `json:"min-window-size"`
    MinWindowCount int `json:"min-window-count"`
}


type DdosTemplateHttpUseHdrIpCfg struct {
    UseHdrIpAsSource int `json:"use-hdr-ip-as-source"`
    L7HdrName string `json:"l7-hdr-name" dval:"X-Forwarded-For"`
}

func (p *DdosTemplateHttp) GetId() string{
    return url.QueryEscape(p.Inst.HttpTmplName)
}

func (p *DdosTemplateHttp) getPath() string{
    return "ddos/template/http"
}

func (p *DdosTemplateHttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttp::Post")
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

func (p *DdosTemplateHttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttp::Get")
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
func (p *DdosTemplateHttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttp::Put")
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

func (p *DdosTemplateHttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateHttp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
