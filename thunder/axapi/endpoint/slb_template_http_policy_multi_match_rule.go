

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateHttpPolicyMultiMatchRule struct {
	Inst struct {

    CookieNameContainsString string `json:"cookie-name-contains-string"`
    CookieNameContainsType string `json:"cookie-name-contains-type"`
    CookieNameEndsWithString string `json:"cookie-name-ends-with-string"`
    CookieNameEndsWithType string `json:"cookie-name-ends-with-type"`
    CookieNameEqualsString string `json:"cookie-name-equals-string"`
    CookieNameEqualsType string `json:"cookie-name-equals-type"`
    CookieNameStartsWithString string `json:"cookie-name-starts-with-string"`
    CookieNameStartsWithType string `json:"cookie-name-starts-with-type"`
    CookieValueContainsString string `json:"cookie-value-contains-string"`
    CookieValueContainsType string `json:"cookie-value-contains-type"`
    CookieValueEndsWithString string `json:"cookie-value-ends-with-string"`
    CookieValueEndsWithType string `json:"cookie-value-ends-with-type"`
    CookieValueEqualsString string `json:"cookie-value-equals-string"`
    CookieValueEqualsType string `json:"cookie-value-equals-type"`
    CookieValueStartsWithString string `json:"cookie-value-starts-with-string"`
    CookieValueStartsWithType string `json:"cookie-value-starts-with-type"`
    HeaderNameContainsString string `json:"header-name-contains-string"`
    HeaderNameContainsType string `json:"header-name-contains-type"`
    HeaderNameEndsWithString string `json:"header-name-ends-with-string"`
    HeaderNameEndsWithType string `json:"header-name-ends-with-type"`
    HeaderNameEqualsString string `json:"header-name-equals-string"`
    HeaderNameEqualsType string `json:"header-name-equals-type"`
    HeaderNameStartsWithString string `json:"header-name-starts-with-string"`
    HeaderNameStartsWithType string `json:"header-name-starts-with-type"`
    HeaderValueContainsString string `json:"header-value-contains-string"`
    HeaderValueContainsType string `json:"header-value-contains-type"`
    HeaderValueEndsWithString string `json:"header-value-ends-with-string"`
    HeaderValueEndsWithType string `json:"header-value-ends-with-type"`
    HeaderValueEqualsString string `json:"header-value-equals-string"`
    HeaderValueEqualsType string `json:"header-value-equals-type"`
    HeaderValueStartsWithString string `json:"header-value-starts-with-string"`
    HeaderValueStartsWithType string `json:"header-value-starts-with-type"`
    HostContainsString string `json:"host-contains-string"`
    HostContainsType string `json:"host-contains-type"`
    HostEndsWithString string `json:"host-ends-with-string"`
    HostEndsWithType string `json:"host-ends-with-type"`
    HostEqualsString string `json:"host-equals-string"`
    HostEqualsType string `json:"host-equals-type"`
    HostStartsWithString string `json:"host-starts-with-string"`
    HostStartsWithType string `json:"host-starts-with-type"`
    MultiMatch string `json:"multi-match"`
    QueryParamNameContainsString string `json:"query-param-name-contains-string"`
    QueryParamNameContainsType string `json:"query-param-name-contains-type"`
    QueryParamNameEndsWithString string `json:"query-param-name-ends-with-string"`
    QueryParamNameEndsWithType string `json:"query-param-name-ends-with-type"`
    QueryParamNameEqualsString string `json:"query-param-name-equals-string"`
    QueryParamNameEqualsType string `json:"query-param-name-equals-type"`
    QueryParamNameStartsWithString string `json:"query-param-name-starts-with-string"`
    QueryParamNameStartsWithType string `json:"query-param-name-starts-with-type"`
    QueryParamValueContainsString string `json:"query-param-value-contains-string"`
    QueryParamValueContainsType string `json:"query-param-value-contains-type"`
    QueryParamValueEndsWithString string `json:"query-param-value-ends-with-string"`
    QueryParamValueEndsWithType string `json:"query-param-value-ends-with-type"`
    QueryParamValueEqualsString string `json:"query-param-value-equals-string"`
    QueryParamValueEqualsType string `json:"query-param-value-equals-type"`
    QueryParamValueStartsWithString string `json:"query-param-value-starts-with-string"`
    QueryParamValueStartsWithType string `json:"query-param-value-starts-with-type"`
    SeqNum int `json:"seq-num"`
    ServiceGroup string `json:"service-group"`
    UrlContainsString string `json:"url-contains-string"`
    UrlContainsType string `json:"url-contains-type"`
    UrlEndsWithString string `json:"url-ends-with-string"`
    UrlEndsWithType string `json:"url-ends-with-type"`
    UrlEqualsString string `json:"url-equals-string"`
    UrlEqualsType string `json:"url-equals-type"`
    UrlStartsWithString string `json:"url-starts-with-string"`
    UrlStartsWithType string `json:"url-starts-with-type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"multi-match-rule"`
}

func (p *SlbTemplateHttpPolicyMultiMatchRule) GetId() string{
    return url.QueryEscape(p.Inst.MultiMatch)
}

func (p *SlbTemplateHttpPolicyMultiMatchRule) getPath() string{
    return "slb/template/http-policy/" +p.Inst.Name + "/multi-match-rule"
}

func (p *SlbTemplateHttpPolicyMultiMatchRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicyMultiMatchRule::Post")
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

func (p *SlbTemplateHttpPolicyMultiMatchRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicyMultiMatchRule::Get")
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
func (p *SlbTemplateHttpPolicyMultiMatchRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicyMultiMatchRule::Put")
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

func (p *SlbTemplateHttpPolicyMultiMatchRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateHttpPolicyMultiMatchRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
