

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch struct {
	Inst struct {

    Action string `json:"action"`
    DisableReqmodIcap int `json:"disable-reqmod-icap"`
    DisableRespmodIcap int `json:"disable-respmod-icap"`
    DualStackAction string `json:"dual-stack-action"`
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
    MatchHttpUrl string `json:"match-http-url"`
    MatchHttpUrlRegex string `json:"match-http-url-regex"`
    MatchHttpUserAgent string `json:"match-http-user-agent"`
    MatchServerAddress string `json:"match-server-address"`
    MatchServerPort int `json:"match-server-port"`
    MatchServerPortRangeBegin int `json:"match-server-port-range-begin"`
    MatchServerPortRangeEnd int `json:"match-server-port-range-end"`
    MatchTimeRange string `json:"match-time-range"`
    MatchWebCategoryList string `json:"match-web-category-list"`
    MatchWebReputationScope string `json:"match-web-reputation-scope"`
    NotifyPage string `json:"notify-page"`
    Priority int `json:"priority"`
    SamplingEnable []SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"adv-match"`
}


type SlbTemplatePolicyForwardPolicySourceDestinationAdvMatchSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch) GetId() string{
    return strconv.Itoa(p.Inst.Priority)
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch) getPath() string{
    return "slb/template/policy/" +p.Inst.Name + "/forward-policy/source/" +p.Inst.Name + "/destination/adv-match"
}

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch::Post")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch::Get")
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
func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch::Put")
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

func (p *SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePolicyForwardPolicySourceDestinationAdvMatch::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
