

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayFormBasedInstanceRequestUri struct {
	Inst struct {

    ActionUri string `json:"action-uri"`
    Cookie AamAuthenticationRelayFormBasedInstanceRequestUriCookie `json:"cookie"`
    DomainVariable string `json:"domain-variable"`
    MatchType string `json:"match-type"`
    MaxPacketCollectSize int `json:"max-packet-collect-size" dval:"1048576"`
    OtherVariables string `json:"other-variables"`
    PasswordVariable string `json:"password-variable"`
    Uri string `json:"uri"`
    UserTag string `json:"user-tag"`
    UserVariable string `json:"user-variable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"request-uri"`
}


type AamAuthenticationRelayFormBasedInstanceRequestUriCookie struct {
    CookieValue AamAuthenticationRelayFormBasedInstanceRequestUriCookieCookieValue `json:"cookie-value"`
}


type AamAuthenticationRelayFormBasedInstanceRequestUriCookieCookieValue struct {
    CookieValue string `json:"cookie-value"`
}

func (p *AamAuthenticationRelayFormBasedInstanceRequestUri) GetId() string{
    return p.Inst.MatchType+"+"+url.QueryEscape(p.Inst.Uri)
}

func (p *AamAuthenticationRelayFormBasedInstanceRequestUri) getPath() string{
    return "aam/authentication/relay/form-based/instance/" +p.Inst.Name + "/request-uri"
}

func (p *AamAuthenticationRelayFormBasedInstanceRequestUri) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstanceRequestUri::Post")
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

func (p *AamAuthenticationRelayFormBasedInstanceRequestUri) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstanceRequestUri::Get")
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
func (p *AamAuthenticationRelayFormBasedInstanceRequestUri) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstanceRequestUri::Put")
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

func (p *AamAuthenticationRelayFormBasedInstanceRequestUri) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstanceRequestUri::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
