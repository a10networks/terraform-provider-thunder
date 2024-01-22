

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicy struct {
	Inst struct {

    AcosEventLog int `json:"acos-event-log"`
    ActionList []SlbTemplatePolicyForwardPolicyActionList `json:"action-list"`
    DualStackActionList []SlbTemplatePolicyForwardPolicyDualStackActionList `json:"dual-stack-action-list"`
    EnableAdvMatch int `json:"enable-adv-match"`
    Filtering []SlbTemplatePolicyForwardPolicyFiltering `json:"filtering"`
    ForwardHttpConnectToIcap int `json:"forward-http-connect-to-icap"`
    LocalLogging int `json:"local-logging"`
    NoClientConnReuse int `json:"no-client-conn-reuse"`
    ReqmodIcap string `json:"reqmod-icap"`
    RequireWebCategory int `json:"require-web-category"`
    SanFiltering []SlbTemplatePolicyForwardPolicySanFiltering `json:"san-filtering"`
    SourceList []SlbTemplatePolicyForwardPolicySourceList `json:"source-list"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"forward-policy"`
}


type SlbTemplatePolicyForwardPolicyActionList struct {
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
    SamplingEnable []SlbTemplatePolicyForwardPolicyActionListSamplingEnable `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicyActionListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicyDualStackActionList struct {
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
    SamplingEnable []SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicyDualStackActionListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicyFiltering struct {
    SsliUrlFiltering string `json:"ssli-url-filtering"`
}


type SlbTemplatePolicyForwardPolicySanFiltering struct {
    SsliUrlFilteringSan string `json:"ssli-url-filtering-san"`
}


type SlbTemplatePolicyForwardPolicySourceList struct {
    Name string `json:"name"`
    MatchClassList string `json:"match-class-list"`
    MatchAny int `json:"match-any"`
    MatchAuthorizePolicy string `json:"match-authorize-policy"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceListSamplingEnable `json:"sampling-enable"`
    Destination SlbTemplatePolicyForwardPolicySourceListDestination `json:"destination"`
}


type SlbTemplatePolicyForwardPolicySourceListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceListDestination struct {
    AdvMatchList []SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList `json:"adv-match-list"`
    ClassListList []SlbTemplatePolicyForwardPolicySourceListDestinationClassListList `json:"class-list-list"`
    WebReputationScopeList []SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList `json:"web-reputation-scope-list"`
    WebCategoryListList []SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList `json:"web-category-list-list"`
    Any SlbTemplatePolicyForwardPolicySourceListDestinationAny `json:"any"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchList struct {
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
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAdvMatchListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationClassListList struct {
    DestClassList string `json:"dest-class-list"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList struct {
    WebReputationScope string `json:"web-reputation-scope"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList struct {
    WebCategoryList string `json:"web-category-list"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAny struct {
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Uuid string `json:"uuid"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicyForwardPolicy) GetId() string{
    return "1"
}

func (p *SlbTemplatePolicyForwardPolicy) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/forward-policy"
}

func (p *SlbTemplatePolicyForwardPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicy::Post")
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

func (p *SlbTemplatePolicyForwardPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicy::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *SlbTemplatePolicyForwardPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicy::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SlbTemplatePolicyForwardPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
