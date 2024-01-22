

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySource struct {
	Inst struct {

    Destination SlbTemplatePolicyForwardPolicySourceDestination1450 `json:"destination"`
    MatchAny int `json:"match-any"`
    MatchAuthorizePolicy string `json:"match-authorize-policy"`
    MatchClassList string `json:"match-class-list"`
    Name string `json:"name"`
    Priority int `json:"priority"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"source"`
}


type SlbTemplatePolicyForwardPolicySourceDestination1450 struct {
    AdvMatchList []SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList `json:"adv-match-list"`
    ClassListList []SlbTemplatePolicyForwardPolicySourceDestinationClassListList `json:"class-list-list"`
    WebReputationScopeList []SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList `json:"web-reputation-scope-list"`
    WebCategoryListList []SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList `json:"web-category-list-list"`
    Any SlbTemplatePolicyForwardPolicySourceDestinationAny1451 `json:"any"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchList struct {
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
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationClassListList struct {
    DestClassList string `json:"dest-class-list"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationWebReputationScopeList struct {
    WebReputationScope string `json:"web-reputation-scope"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationWebCategoryListList struct {
    WebCategoryList string `json:"web-category-list"`
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Type string `json:"type"`
    Priority int `json:"priority"`
    Uuid string `json:"uuid"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAny1451 struct {
    Action string `json:"action"`
    DualStackAction string `json:"dual-stack-action"`
    Uuid string `json:"uuid"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452 `json:"sampling-enable"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAnySamplingEnable1452 struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplatePolicyForwardPolicySourceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicyForwardPolicySource) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePolicyForwardPolicySource) getPath() string{
    return "slb/template/policy/forward-policy/source"
}

func (p *SlbTemplatePolicyForwardPolicySource) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySource::Post")
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

func (p *SlbTemplatePolicyForwardPolicySource) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySource::Get")
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
func (p *SlbTemplatePolicyForwardPolicySource) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySource::Put")
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

func (p *SlbTemplatePolicyForwardPolicySource) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySource::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
