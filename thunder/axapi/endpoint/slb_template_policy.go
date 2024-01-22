

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicy struct {
	Inst struct {

    BwListId []SlbTemplatePolicyBwListId `json:"bw-list-id"`
    BwListName string `json:"bw-list-name"`
    ClassList SlbTemplatePolicyClassList1453 `json:"class-list"`
    ForwardPolicy SlbTemplatePolicyForwardPolicy1457 `json:"forward-policy"`
    FullDomainTree int `json:"full-domain-tree"`
    Interval int `json:"interval"`
    Name string `json:"name"`
    OverLimit int `json:"over-limit"`
    OverLimitLockup int `json:"over-limit-lockup"`
    OverLimitLogging int `json:"over-limit-logging"`
    OverLimitReset int `json:"over-limit-reset"`
    Overlap int `json:"overlap"`
    SamplingEnable []SlbTemplatePolicySamplingEnable `json:"sampling-enable"`
    Share int `json:"share"`
    Timeout int `json:"timeout" dval:"5"`
    UseDestinationIp int `json:"use-destination-ip"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"policy"`
}


type SlbTemplatePolicyBwListId struct {
    Id1 int `json:"id1"`
    ServiceGroup string `json:"service-group"`
    PbslbLogging int `json:"pbslb-logging"`
    PbslbInterval int `json:"pbslb-interval" dval:"3"`
    Fail int `json:"fail"`
    BwListAction string `json:"bw-list-action"`
    LoggingDrpRst int `json:"logging-drp-rst"`
    ActionInterval int `json:"action-interval" dval:"3"`
}


type SlbTemplatePolicyClassList1453 struct {
    Name string `json:"name"`
    ClientIpL3Dest int `json:"client-ip-l3-dest"`
    ClientIpL7Header int `json:"client-ip-l7-header"`
    HeaderName string `json:"header-name"`
    Uuid string `json:"uuid"`
    LidList []SlbTemplatePolicyClassListLidList1454 `json:"lid-list"`
}


type SlbTemplatePolicyClassListLidList1454 struct {
    Lidnum int `json:"lidnum"`
    ConnLimit int `json:"conn-limit"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnPer int `json:"conn-per"`
    RequestLimit int `json:"request-limit"`
    RequestRateLimit int `json:"request-rate-limit"`
    RequestPer int `json:"request-per"`
    BwRateLimit int `json:"bw-rate-limit"`
    BwPer int `json:"bw-per"`
    OverLimitAction int `json:"over-limit-action"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    Interval int `json:"interval"`
    DirectAction int `json:"direct-action"`
    DirectServiceGroup string `json:"direct-service-group"`
    DirectPbslbLogging int `json:"direct-pbslb-logging"`
    DirectPbslbInterval int `json:"direct-pbslb-interval" dval:"3"`
    DirectFail int `json:"direct-fail"`
    DirectActionValue string `json:"direct-action-value"`
    DirectLoggingDrpRst int `json:"direct-logging-drp-rst"`
    DirectActionInterval int `json:"direct-action-interval" dval:"3"`
    ResponseCodeRateLimit []SlbTemplatePolicyClassListLidListResponseCodeRateLimit1455 `json:"response-code-rate-limit"`
    Dns64 SlbTemplatePolicyClassListLidListDns641456 `json:"dns64"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SlbTemplatePolicyClassListLidListResponseCodeRateLimit1455 struct {
    CodeRangeStart int `json:"code-range-start"`
    CodeRangeEnd int `json:"code-range-end"`
    Threshold int `json:"threshold"`
    Period int `json:"period"`
}


type SlbTemplatePolicyClassListLidListDns641456 struct {
    Disable int `json:"disable"`
    ExclusiveAnswer int `json:"exclusive-answer"`
    Prefix string `json:"prefix"`
}


type SlbTemplatePolicyForwardPolicy1457 struct {
    NoClientConnReuse int `json:"no-client-conn-reuse"`
    AcosEventLog int `json:"acos-event-log"`
    LocalLogging int `json:"local-logging"`
    RequireWebCategory int `json:"require-web-category"`
    ForwardHttpConnectToIcap int `json:"forward-http-connect-to-icap"`
    ReqmodIcap string `json:"reqmod-icap"`
    Filtering []SlbTemplatePolicyForwardPolicyFiltering1458 `json:"filtering"`
    SanFiltering []SlbTemplatePolicyForwardPolicySanFiltering1459 `json:"san-filtering"`
    EnableAdvMatch int `json:"enable-adv-match"`
    Uuid string `json:"uuid"`
    ActionList []SlbTemplatePolicyForwardPolicyActionList1460 `json:"action-list"`
    DualStackActionList []SlbTemplatePolicyForwardPolicyDualStackActionList1462 `json:"dual-stack-action-list"`
    SourceList []SlbTemplatePolicyForwardPolicySourceList1464 `json:"source-list"`
}


type SlbTemplatePolicyForwardPolicyFiltering1458 struct {
    SsliUrlFiltering string `json:"ssli-url-filtering"`
}


type SlbTemplatePolicyForwardPolicySanFiltering1459 struct {
    SsliUrlFilteringSan string `json:"ssli-url-filtering-san"`
}


type SlbTemplatePolicyForwardPolicyActionList1460 struct {
    Name string `json:"name"`
    Action1 string `json:"action1"`
    FakeSg string `json:"fake-sg"`
    RealSg string `json:"real-sg"`
    ForwardSnat string `json:"forward-snat"`
    ForwardSnatPtOnly int `json:"forward-snat-pt-only"`
    FallBack string `json:"fall-back"`
    FallBackSnat string `json:"fall-back-snat"`
    FallBackSnatPtOnly int `json:"fall-back-snat-pt-only"`
    ProxyChaining int `json:"proxy-chaining"`
    ProxyChainingBypass int `json:"proxy-chaining-bypass"`
    SupportCertFetch int `json:"support-cert-fetch"`
    Log int `json:"log"`
    DropResponseCode int `json:"drop-response-code"`
    DropMessage string `json:"drop-message"`
    DropRedirectUrl string `json:"drop-redirect-url"`
    HttpStatusCode string `json:"http-status-code" dval:"302"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbTemplatePolicyForwardPolicyActionListSamplingEnable1461 `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicyActionListSamplingEnable1461 struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicyDualStackActionList1462 struct {
    Name string `json:"name"`
    Ipv4 string `json:"ipv4"`
    Ipv4Snat string `json:"ipv4-snat"`
    Ipv6 string `json:"ipv6"`
    Ipv6Snat string `json:"ipv6-snat"`
    FallBack string `json:"fall-back"`
    FallBackSnat string `json:"fall-back-snat"`
    Log int `json:"log"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463 `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable1463 struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceList1464 struct {
    Name string `json:"name"`
    MatchClassList string `json:"match-class-list"`
    MatchAny int `json:"match-any"`
    MatchAuthorizePolicy string `json:"match-authorize-policy"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceListSamplingEnable1465 `json:"sampling-enable"`
    Destination SlbTemplatePolicyForwardPolicySourceListDestination1466 `json:"destination"`
}


type SlbTemplatePolicyForwardPolicySourceListSamplingEnable1465 struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceListDestination1466 struct {
    AdvMatchList []SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467 `json:"adv-match-list"`
    ClassListList []SlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469 `json:"class-list-list"`
    WebReputationScopeList []SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470 `json:"web-reputation-scope-list"`
    WebCategoryListList []SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471 `json:"web-category-list-list"`
    Any SlbTemplatePolicyForwardPolicySourceListDestinationAny1472 `json:"any"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList1467 struct {
    Priority int `json:"priority"`
    MatchHost string `json:"match-host"`
    MatchHttpContentEncoding string `json:"match-http-content-encoding"`
    MatchHttpContentLengthRangeBegin int `json:"match-http-content-length-range-begin"`
    MatchHttpContentLengthRangeEnd int `json:"match-http-content-length-range-end"`
    MatchHttpContentType string `json:"match-http-content-type"`
    MatchHttpHeader string `json:"match-http-header"`
    MatchHttpMethodConnect int `json:"match-http-method-connect"`
    MatchHttpMethodDelete int `json:"match-http-method-delete"`
    MatchHttpMethodGet int `json:"match-http-method-get"`
    MatchHttpMethodHead int `json:"match-http-method-head"`
    MatchHttpMethodOptions int `json:"match-http-method-options"`
    MatchHttpMethodPatch int `json:"match-http-method-patch"`
    MatchHttpMethodPost int `json:"match-http-method-post"`
    MatchHttpMethodPut int `json:"match-http-method-put"`
    MatchHttpMethodTrace int `json:"match-http-method-trace"`
    MatchHttpRequestFileExtension string `json:"match-http-request-file-extension"`
    MatchHttpUrlRegex string `json:"match-http-url-regex"`
    MatchHttpUrl string `json:"match-http-url"`
    MatchHttpUserAgent string `json:"match-http-user-agent"`
    MatchServerAddress string `json:"match-server-address"`
    MatchServerPort int `json:"match-server-port"`
    MatchServerPortRangeBegin int `json:"match-server-port-range-begin"`
    MatchServerPortRangeEnd int `json:"match-server-port-range-end"`
    MatchTimeRange string `json:"match-time-range"`
    MatchWebCategoryList string `json:"match-web-category-list"`
    MatchWebReputationScope string `json:"match-web-reputation-scope"`
    DisableReqmodIcap int `json:"disable-reqmod-icap"`
    DisableRespmodIcap int `json:"disable-respmod-icap"`
    NotifyPage string `json:"notify-page"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468 `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable1468 struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationClassListList1469 struct {
    DestClassList string `json:"dest-class-list"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList1470 struct {
    WebReputationScope string `json:"web-reputation-scope"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList1471 struct {
    WebCategoryList string `json:"web-category-list"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAny1472 struct {
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Uuid string `json:"uuid"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473 `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable1473 struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePolicy) getPath() string{
    return "slb/template/policy"
}

func (p *SlbTemplatePolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicy::Post")
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

func (p *SlbTemplatePolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicy::Get")
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
func (p *SlbTemplatePolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicy::Put")
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

func (p *SlbTemplatePolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
