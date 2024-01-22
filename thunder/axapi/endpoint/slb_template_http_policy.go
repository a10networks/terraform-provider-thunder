

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateHttpPolicy struct {
	Inst struct {

    CookieName string `json:"cookie-name"`
    GeoLocationMatch []SlbTemplateHttpPolicyGeoLocationMatch `json:"geo-location-match"`
    HttpPolicyMatch []SlbTemplateHttpPolicyHttpPolicyMatch `json:"http-policy-match"`
    MultiMatchRuleList []SlbTemplateHttpPolicyMultiMatchRuleList `json:"multi-match-rule-list"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"http-policy"`
}


type SlbTemplateHttpPolicyGeoLocationMatch struct {
    GeoLocation string `json:"geo-location"`
    GeoLocationServiceGroup string `json:"geo-location-service-group"`
}


type SlbTemplateHttpPolicyHttpPolicyMatch struct {
    Type string `json:"type"`
    MatchType string `json:"match-type"`
    MatchString string `json:"match-string"`
    ServiceGroup string `json:"service-group"`
}


type SlbTemplateHttpPolicyMultiMatchRuleList struct {
    MultiMatch string `json:"multi-match"`
    SeqNum int `json:"seq-num"`
    HostEqualsType string `json:"host-equals-type"`
    HostEqualsString string `json:"host-equals-string"`
    HostContainsType string `json:"host-contains-type"`
    HostContainsString string `json:"host-contains-string"`
    HostStartsWithType string `json:"host-starts-with-type"`
    HostStartsWithString string `json:"host-starts-with-string"`
    HostEndsWithType string `json:"host-ends-with-type"`
    HostEndsWithString string `json:"host-ends-with-string"`
    CookieNameEqualsType string `json:"cookie-name-equals-type"`
    CookieNameEqualsString string `json:"cookie-name-equals-string"`
    CookieNameContainsType string `json:"cookie-name-contains-type"`
    CookieNameContainsString string `json:"cookie-name-contains-string"`
    CookieNameStartsWithType string `json:"cookie-name-starts-with-type"`
    CookieNameStartsWithString string `json:"cookie-name-starts-with-string"`
    CookieNameEndsWithType string `json:"cookie-name-ends-with-type"`
    CookieNameEndsWithString string `json:"cookie-name-ends-with-string"`
    CookieValueEqualsType string `json:"cookie-value-equals-type"`
    CookieValueEqualsString string `json:"cookie-value-equals-string"`
    CookieValueContainsType string `json:"cookie-value-contains-type"`
    CookieValueContainsString string `json:"cookie-value-contains-string"`
    CookieValueStartsWithType string `json:"cookie-value-starts-with-type"`
    CookieValueStartsWithString string `json:"cookie-value-starts-with-string"`
    CookieValueEndsWithType string `json:"cookie-value-ends-with-type"`
    CookieValueEndsWithString string `json:"cookie-value-ends-with-string"`
    UrlEqualsType string `json:"url-equals-type"`
    UrlEqualsString string `json:"url-equals-string"`
    UrlContainsType string `json:"url-contains-type"`
    UrlContainsString string `json:"url-contains-string"`
    UrlStartsWithType string `json:"url-starts-with-type"`
    UrlStartsWithString string `json:"url-starts-with-string"`
    UrlEndsWithType string `json:"url-ends-with-type"`
    UrlEndsWithString string `json:"url-ends-with-string"`
    HeaderNameEqualsType string `json:"header-name-equals-type"`
    HeaderNameEqualsString string `json:"header-name-equals-string"`
    HeaderNameContainsType string `json:"header-name-contains-type"`
    HeaderNameContainsString string `json:"header-name-contains-string"`
    HeaderNameStartsWithType string `json:"header-name-starts-with-type"`
    HeaderNameStartsWithString string `json:"header-name-starts-with-string"`
    HeaderNameEndsWithType string `json:"header-name-ends-with-type"`
    HeaderNameEndsWithString string `json:"header-name-ends-with-string"`
    HeaderValueEqualsType string `json:"header-value-equals-type"`
    HeaderValueEqualsString string `json:"header-value-equals-string"`
    HeaderValueContainsType string `json:"header-value-contains-type"`
    HeaderValueContainsString string `json:"header-value-contains-string"`
    HeaderValueStartsWithType string `json:"header-value-starts-with-type"`
    HeaderValueStartsWithString string `json:"header-value-starts-with-string"`
    HeaderValueEndsWithType string `json:"header-value-ends-with-type"`
    HeaderValueEndsWithString string `json:"header-value-ends-with-string"`
    QueryParamNameEqualsType string `json:"query-param-name-equals-type"`
    QueryParamNameEqualsString string `json:"query-param-name-equals-string"`
    QueryParamNameContainsType string `json:"query-param-name-contains-type"`
    QueryParamNameContainsString string `json:"query-param-name-contains-string"`
    QueryParamNameStartsWithType string `json:"query-param-name-starts-with-type"`
    QueryParamNameStartsWithString string `json:"query-param-name-starts-with-string"`
    QueryParamNameEndsWithType string `json:"query-param-name-ends-with-type"`
    QueryParamNameEndsWithString string `json:"query-param-name-ends-with-string"`
    QueryParamValueEqualsType string `json:"query-param-value-equals-type"`
    QueryParamValueEqualsString string `json:"query-param-value-equals-string"`
    QueryParamValueContainsType string `json:"query-param-value-contains-type"`
    QueryParamValueContainsString string `json:"query-param-value-contains-string"`
    QueryParamValueStartsWithType string `json:"query-param-value-starts-with-type"`
    QueryParamValueStartsWithString string `json:"query-param-value-starts-with-string"`
    QueryParamValueEndsWithType string `json:"query-param-value-ends-with-type"`
    QueryParamValueEndsWithString string `json:"query-param-value-ends-with-string"`
    ServiceGroup string `json:"service-group"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *SlbTemplateHttpPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateHttpPolicy) getPath() string{
    return "slb/template/http-policy"
}

func (p *SlbTemplateHttpPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicy::Post")
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

func (p *SlbTemplateHttpPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicy::Get")
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
func (p *SlbTemplateHttpPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicy::Put")
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

func (p *SlbTemplateHttpPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
