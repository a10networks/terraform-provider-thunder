

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AamAaaPolicyAaaRule struct {
	Inst struct {

    AccessList AamAaaPolicyAaaRuleAccessList `json:"access-list"`
    Action string `json:"action"`
    AuthFailureBypass int `json:"auth-failure-bypass"`
    AuthenticationTemplate string `json:"authentication-template"`
    AuthorizePolicy string `json:"authorize-policy"`
    CaptchaAuthzPolicy string `json:"captcha-authz-policy"`
    DomainName string `json:"domain-name"`
    DomainWhitelist string `json:"domain-whitelist"`
    Host []AamAaaPolicyAaaRuleHost `json:"host"`
    Index int `json:"index"`
    MatchEncodedUri int `json:"match-encoded-uri"`
    Port int `json:"port"`
    SamplingEnable []AamAaaPolicyAaaRuleSamplingEnable `json:"sampling-enable"`
    Uri []AamAaaPolicyAaaRuleUri `json:"uri"`
    UserAgent []AamAaaPolicyAaaRuleUserAgent `json:"user-agent"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"aaa-rule"`
}


type AamAaaPolicyAaaRuleAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
    Name string `json:"name"`
}


type AamAaaPolicyAaaRuleHost struct {
    HostMatchType string `json:"host-match-type"`
    HostStr string `json:"host-str"`
}


type AamAaaPolicyAaaRuleSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAaaPolicyAaaRuleUri struct {
    MatchType string `json:"match-type"`
    UriStr string `json:"uri-str"`
}


type AamAaaPolicyAaaRuleUserAgent struct {
    UserAgentMatchType string `json:"user-agent-match-type"`
    UserAgentStr string `json:"user-agent-str"`
}

func (p *AamAaaPolicyAaaRule) GetId() string{
    return strconv.Itoa(p.Inst.Index)
}

func (p *AamAaaPolicyAaaRule) getPath() string{
    return "aam/aaa-policy/" +p.Inst.Name + "/aaa-rule"
}

func (p *AamAaaPolicyAaaRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicyAaaRule::Post")
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

func (p *AamAaaPolicyAaaRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicyAaaRule::Get")
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
func (p *AamAaaPolicyAaaRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicyAaaRule::Put")
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

func (p *AamAaaPolicyAaaRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicyAaaRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
