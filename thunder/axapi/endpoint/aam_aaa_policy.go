

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAaaPolicy struct {
	Inst struct {

    AaaRuleList []AamAaaPolicyAaaRuleList `json:"aaa-rule-list"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []AamAaaPolicySamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"aaa-policy"`
}


type AamAaaPolicyAaaRuleList struct {
    Index int `json:"index"`
    Uri []AamAaaPolicyAaaRuleListUri `json:"uri"`
    Host []AamAaaPolicyAaaRuleListHost `json:"host"`
    DomainWhitelist string `json:"domain-whitelist"`
    Port int `json:"port"`
    MatchEncodedUri int `json:"match-encoded-uri"`
    AccessList AamAaaPolicyAaaRuleListAccessList `json:"access-list"`
    DomainName string `json:"domain-name"`
    UserAgent []AamAaaPolicyAaaRuleListUserAgent `json:"user-agent"`
    Action string `json:"action"`
    AuthenticationTemplate string `json:"authentication-template"`
    AuthorizePolicy string `json:"authorize-policy"`
    CaptchaAuthzPolicy string `json:"captcha-authz-policy"`
    AuthFailureBypass int `json:"auth-failure-bypass"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []AamAaaPolicyAaaRuleListSamplingEnable `json:"sampling-enable"`
}


type AamAaaPolicyAaaRuleListUri struct {
    MatchType string `json:"match-type"`
    UriStr string `json:"uri-str"`
}


type AamAaaPolicyAaaRuleListHost struct {
    HostMatchType string `json:"host-match-type"`
    HostStr string `json:"host-str"`
}


type AamAaaPolicyAaaRuleListAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
    Name string `json:"name"`
}


type AamAaaPolicyAaaRuleListUserAgent struct {
    UserAgentMatchType string `json:"user-agent-match-type"`
    UserAgentStr string `json:"user-agent-str"`
}


type AamAaaPolicyAaaRuleListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type AamAaaPolicySamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAaaPolicy) GetId() string{
    return p.Inst.Name
}

func (p *AamAaaPolicy) getPath() string{
    return "aam/aaa-policy"
}

func (p *AamAaaPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicy::Post")
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

func (p *AamAaaPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicy::Get")
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
func (p *AamAaaPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicy::Put")
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

func (p *AamAaaPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAaaPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
