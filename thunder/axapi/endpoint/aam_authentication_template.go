

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationTemplate struct {
	Inst struct {

    Account string `json:"account"`
    AccountingServer string `json:"accounting-server"`
    AccountingServiceGroup string `json:"accounting-service-group"`
    AuthSessMode string `json:"auth-sess-mode"`
    Captcha string `json:"captcha"`
    Chain []AamAuthenticationTemplateChain `json:"chain"`
    CookieDomain []AamAuthenticationTemplateCookieDomain `json:"cookie-domain"`
    CookieDomainGroup []AamAuthenticationTemplateCookieDomainGroup `json:"cookie-domain-group"`
    CookieHttponlyEnable int `json:"cookie-httponly-enable"`
    CookieMaxAge int `json:"cookie-max-age" dval:"604800"`
    CookieSamesite string `json:"cookie-samesite"`
    CookieSecureEnable int `json:"cookie-secure-enable"`
    ForwardLogoutDisable int `json:"forward-logout-disable"`
    Jwt string `json:"jwt"`
    LocalLogging int `json:"local-logging"`
    Log string `json:"log" dval:"use-partition-level-config"`
    Logon string `json:"logon"`
    LogoutIdleTimeout int `json:"logout-idle-timeout" dval:"300"`
    LogoutUrl string `json:"logout-url"`
    MaxSessionTime int `json:"max-session-time"`
    ModifyContentSecurityPolicy int `json:"modify-content-security-policy"`
    Name string `json:"name"`
    OauthAuthorizationServer string `json:"oauth-authorization-server"`
    OauthClient string `json:"oauth-client"`
    RedirectHostname string `json:"redirect-hostname"`
    Relay string `json:"relay"`
    SamlIdp string `json:"saml-idp"`
    SamlSp string `json:"saml-sp"`
    Server string `json:"server"`
    ServiceGroup string `json:"service-group"`
    Type string `json:"type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"template"`
}


type AamAuthenticationTemplateChain struct {
    ChainServer string `json:"chain-server"`
    ChainServerPriority int `json:"chain-server-priority" dval:"3"`
    ChainSg string `json:"chain-sg"`
    ChainSgPriority int `json:"chain-sg-priority" dval:"3"`
}


type AamAuthenticationTemplateCookieDomain struct {
    CookieDmn string `json:"cookie-dmn"`
}


type AamAuthenticationTemplateCookieDomainGroup struct {
    CookieDmngrp int `json:"cookie-dmngrp"`
}

func (p *AamAuthenticationTemplate) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationTemplate) getPath() string{
    return "aam/authentication/template"
}

func (p *AamAuthenticationTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationTemplate::Post")
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

func (p *AamAuthenticationTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationTemplate::Get")
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
func (p *AamAuthenticationTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationTemplate::Put")
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

func (p *AamAuthenticationTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
