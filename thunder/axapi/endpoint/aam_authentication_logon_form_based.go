

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationLogonFormBased struct {
	Inst struct {

    AccountLock int `json:"account-lock"`
    ChallengeVariable string `json:"challenge-variable"`
    CpPageCfg AamAuthenticationLogonFormBasedCpPageCfg `json:"cp-page-cfg"`
    CspSupport AamAuthenticationLogonFormBasedCspSupport `json:"csp-support"`
    Duration int `json:"duration" dval:"1800"`
    HstsTimeout int `json:"hsts-timeout"`
    LogonPageCfg AamAuthenticationLogonFormBasedLogonPageCfg `json:"logon-page-cfg"`
    Name string `json:"name"`
    NewPinVariable string `json:"new-pin-variable"`
    NextTokenVariable string `json:"next-token-variable"`
    NotifyCpPageCfg AamAuthenticationLogonFormBasedNotifyCpPageCfg `json:"notify-cp-page-cfg"`
    Portal AamAuthenticationLogonFormBasedPortal `json:"portal"`
    Retry int `json:"retry" dval:"3"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"form-based"`
}


type AamAuthenticationLogonFormBasedCpPageCfg struct {
    ChangepasswordUrl string `json:"changepassword-url"`
    CpUserEnum string `json:"cp-user-enum"`
    CpUserVar string `json:"cp-user-var"`
    CpOldPwdEnum string `json:"cp-old-pwd-enum"`
    CpOldPwdVar string `json:"cp-old-pwd-var"`
    CpNewPwdEnum string `json:"cp-new-pwd-enum"`
    CpNewPwdVar string `json:"cp-new-pwd-var"`
    CpCfmPwdEnum string `json:"cp-cfm-pwd-enum"`
    CpCfmPwdVar string `json:"cp-cfm-pwd-var"`
}


type AamAuthenticationLogonFormBasedCspSupport struct {
    None int `json:"none"`
    Self int `json:"self"`
    Specificuri string `json:"specificURI"`
    OptionalSecondUri string `json:"optional-second-URI"`
}


type AamAuthenticationLogonFormBasedLogonPageCfg struct {
    ActionUrl string `json:"action-url"`
    UsernameVariable string `json:"username-variable"`
    PasswordVariable string `json:"password-variable"`
    PasscodeVariable string `json:"passcode-variable"`
    CaptchaVariable string `json:"captcha-variable"`
    LoginFailureMessage string `json:"login-failure-message"`
    AuthzFailureMessage string `json:"authz-failure-message"`
    DisableChangePasswordLink int `json:"disable-change-password-link"`
}


type AamAuthenticationLogonFormBasedNotifyCpPageCfg struct {
    NotifychangepasswordChangeUrl string `json:"notifychangepassword-change-url"`
    NotifychangepasswordContinueUrl string `json:"notifychangepassword-continue-url"`
}


type AamAuthenticationLogonFormBasedPortal struct {
    DefaultPortal int `json:"default-portal"`
    PortalName string `json:"portal-name"`
    Logon string `json:"logon"`
    Failpage string `json:"failpage"`
    Changepasswordpage string `json:"changepasswordpage"`
    Notifychangepasswordpage string `json:"notifychangepasswordpage"`
    ChallengePage string `json:"challenge-page"`
    NewPinPage string `json:"new-pin-page"`
    NextTokenPage string `json:"next-token-page"`
}

func (p *AamAuthenticationLogonFormBased) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationLogonFormBased) getPath() string{
    return "aam/authentication/logon/form-based"
}

func (p *AamAuthenticationLogonFormBased) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonFormBased::Post")
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

func (p *AamAuthenticationLogonFormBased) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonFormBased::Get")
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
func (p *AamAuthenticationLogonFormBased) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonFormBased::Put")
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

func (p *AamAuthenticationLogonFormBased) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonFormBased::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
