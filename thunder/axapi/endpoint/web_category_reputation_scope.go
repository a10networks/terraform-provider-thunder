

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryReputationScope struct {
	Inst struct {

    GreaterThan WebCategoryReputationScopeGreaterThan `json:"greater-than"`
    LessThan WebCategoryReputationScopeLessThan `json:"less-than"`
    Name string `json:"name"`
    SamplingEnable []WebCategoryReputationScopeSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"reputation-scope"`
}


type WebCategoryReputationScopeGreaterThan struct {
    GreaterTrustworthy int `json:"greater-trustworthy"`
    GreaterLowRisk int `json:"greater-low-risk"`
    GreaterModerateRisk int `json:"greater-moderate-risk"`
    GreaterSuspicious int `json:"greater-suspicious"`
    GreaterMalicious int `json:"greater-malicious"`
    GreaterThreshold int `json:"greater-threshold"`
}


type WebCategoryReputationScopeLessThan struct {
    LessTrustworthy int `json:"less-trustworthy"`
    LessLowRisk int `json:"less-low-risk"`
    LessModerateRisk int `json:"less-moderate-risk"`
    LessSuspicious int `json:"less-suspicious"`
    LessMalicious int `json:"less-malicious"`
    LessThreshold int `json:"less-threshold"`
}


type WebCategoryReputationScopeSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *WebCategoryReputationScope) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *WebCategoryReputationScope) getPath() string{
    return "web-category/reputation-scope"
}

func (p *WebCategoryReputationScope) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryReputationScope::Post")
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

func (p *WebCategoryReputationScope) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryReputationScope::Get")
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
func (p *WebCategoryReputationScope) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryReputationScope::Put")
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

func (p *WebCategoryReputationScope) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebCategoryReputationScope::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
