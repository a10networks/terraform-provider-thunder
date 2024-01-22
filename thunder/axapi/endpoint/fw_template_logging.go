

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type FwTemplateLogging struct {
	Inst struct {

    Facility string `json:"facility" dval:"local0"`
    Format string `json:"format" dval:"cef"`
    IncludeDestFqdn int `json:"include-dest-fqdn"`
    IncludeHttp FwTemplateLoggingIncludeHttp `json:"include-http"`
    IncludeRadiusAttribute FwTemplateLoggingIncludeRadiusAttribute `json:"include-radius-attribute"`
    Log FwTemplateLoggingLog `json:"log"`
    MergedStyle int `json:"merged-style"`
    Name string `json:"name"`
    Resolution string `json:"resolution" dval:"seconds"`
    Rule FwTemplateLoggingRule `json:"rule"`
    ServiceGroup string `json:"service-group"`
    SessionPeriodicLog FwTemplateLoggingSessionPeriodicLog378 `json:"session-periodic-log"`
    Severity string `json:"severity" dval:"informational"`
    SourceAddress FwTemplateLoggingSourceAddress379 `json:"source-address"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}


type FwTemplateLoggingIncludeHttp struct {
    HeaderCfg []FwTemplateLoggingIncludeHttpHeaderCfg `json:"header-cfg"`
    L4SessionInfo int `json:"l4-session-info"`
    Method int `json:"method"`
    RequestNumber int `json:"request-number"`
    FileExtension int `json:"file-extension"`
}


type FwTemplateLoggingIncludeHttpHeaderCfg struct {
    HttpHeader string `json:"http-header"`
    MaxLength int `json:"max-length" dval:"100"`
    CustomHeaderName string `json:"custom-header-name"`
    CustomMaxLength int `json:"custom-max-length" dval:"100"`
}


type FwTemplateLoggingIncludeRadiusAttribute struct {
    AttrCfg []FwTemplateLoggingIncludeRadiusAttributeAttrCfg `json:"attr-cfg"`
    NoQuote int `json:"no-quote"`
    FramedIpv6Prefix int `json:"framed-ipv6-prefix"`
    PrefixLength string `json:"prefix-length"`
    InsertIfNotExisting int `json:"insert-if-not-existing"`
    ZeroInCustomAttr int `json:"zero-in-custom-attr"`
}


type FwTemplateLoggingIncludeRadiusAttributeAttrCfg struct {
    Attr string `json:"attr"`
    AttrEvent string `json:"attr-event"`
}


type FwTemplateLoggingLog struct {
    HttpRequests string `json:"http-requests"`
}


type FwTemplateLoggingRule struct {
    RuleHttpRequests FwTemplateLoggingRuleRuleHttpRequests `json:"rule-http-requests"`
}


type FwTemplateLoggingRuleRuleHttpRequests struct {
    DestPort []FwTemplateLoggingRuleRuleHttpRequestsDestPort `json:"dest-port"`
    LogEveryHttpRequest int `json:"log-every-http-request"`
    MaxUrlLen int `json:"max-url-len" dval:"128"`
    IncludeAllHeaders int `json:"include-all-headers"`
    DisableSequenceCheck int `json:"disable-sequence-check"`
}


type FwTemplateLoggingRuleRuleHttpRequestsDestPort struct {
    DestPortNumber int `json:"dest-port-number"`
    IncludeByteCount int `json:"include-byte-count"`
}


type FwTemplateLoggingSessionPeriodicLog378 struct {
    Interval int `json:"interval"`
    Uuid string `json:"uuid"`
}


type FwTemplateLoggingSourceAddress379 struct {
    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`
}

func (p *FwTemplateLogging) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *FwTemplateLogging) getPath() string{
    return "fw/template/logging"
}

func (p *FwTemplateLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLogging::Post")
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

func (p *FwTemplateLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLogging::Get")
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
func (p *FwTemplateLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLogging::Put")
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

func (p *FwTemplateLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
