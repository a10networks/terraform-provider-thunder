

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationPortalLogon struct {
	Inst struct {

    ActionUrl string `json:"action-url"`
    Background AamAuthenticationPortalLogonBackground `json:"background"`
    CaptchaType string `json:"captcha-type"`
    EnableCaptcha int `json:"enable-CAPTCHA"`
    EnablePasscode int `json:"enable-passcode"`
    Encrypted string `json:"encrypted"`
    FailMsgCfg AamAuthenticationPortalLogonFailMsgCfg `json:"fail-msg-cfg"`
    PasscodeCfg AamAuthenticationPortalLogonPasscodeCfg `json:"passcode-cfg"`
    PasscodeVar string `json:"passcode-var"`
    PasswordCfg AamAuthenticationPortalLogonPasswordCfg `json:"password-cfg"`
    PasswordVar string `json:"password-var"`
    RecaptchaCfg AamAuthenticationPortalLogonRecaptchaCfg `json:"reCAPTCHA-cfg"`
    SiteKeyString string `json:"site-key-string"`
    SubmitText string `json:"submit-text"`
    UsernameCfg AamAuthenticationPortalLogonUsernameCfg `json:"username-cfg"`
    UsernameVar string `json:"username-var"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"logon"`
}


type AamAuthenticationPortalLogonBackground struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalLogonFailMsgCfg struct {
    FailMsg int `json:"fail-msg"`
    FailText string `json:"fail-text"`
    FailFont int `json:"fail-font" dval:"1"`
    FailFace string `json:"fail-face" dval:"Arial"`
    FailFontCustom string `json:"fail-font-custom"`
    FailSize int `json:"fail-size" dval:"5"`
    FailColor int `json:"fail-color" dval:"1"`
    FailColorName string `json:"fail-color-name" dval:"red"`
    FailColorValue string `json:"fail-color-value"`
    AuthzFailMsg string `json:"authz-fail-msg"`
}


type AamAuthenticationPortalLogonPasscodeCfg struct {
    Passcode int `json:"passcode"`
    PasscodeText string `json:"passcode-text"`
    PasscodeFont int `json:"passcode-font" dval:"1"`
    PasscodeFace string `json:"passcode-face" dval:"Arial"`
    PasscodeFontCustom string `json:"passcode-font-custom"`
    PasscodeSize int `json:"passcode-size" dval:"3"`
    PasscodeColor int `json:"passcode-color" dval:"1"`
    PasscodeColorName string `json:"passcode-color-name" dval:"black"`
    PasscodeColorValue string `json:"passcode-color-value"`
}


type AamAuthenticationPortalLogonPasswordCfg struct {
    Password int `json:"password"`
    PassText string `json:"pass-text"`
    PassFont int `json:"pass-font" dval:"1"`
    PassFace string `json:"pass-face" dval:"Arial"`
    PassFontCustom string `json:"pass-font-custom"`
    PassSize int `json:"pass-size" dval:"3"`
    PassColor int `json:"pass-color" dval:"1"`
    PassColorName string `json:"pass-color-name" dval:"black"`
    PassColorValue string `json:"pass-color-value"`
}


type AamAuthenticationPortalLogonRecaptchaCfg struct {
    RecaptchaTheme string `json:"reCAPTCHA-theme" dval:"light"`
    RecaptchaSize string `json:"reCAPTCHA-size" dval:"normal"`
    RecaptchaBadge string `json:"reCAPTCHA-badge" dval:"bottom-right"`
    RecaptchaAction string `json:"reCAPTCHA-action" dval:"A10_DEFAULT_LOGON"`
}


type AamAuthenticationPortalLogonUsernameCfg struct {
    Username int `json:"username"`
    UserText string `json:"user-text"`
    UserFont int `json:"user-font" dval:"1"`
    UserFace string `json:"user-face" dval:"Arial"`
    UserFontCustom string `json:"user-font-custom"`
    UserSize int `json:"user-size" dval:"3"`
    UserColor int `json:"user-color" dval:"1"`
    UserColorName string `json:"user-color-name" dval:"black"`
    UserColorValue string `json:"user-color-value"`
}

func (p *AamAuthenticationPortalLogon) GetId() string{
    return "1"
}

func (p *AamAuthenticationPortalLogon) getPath() string{
    return "aam/authentication/portal/" +p.Inst.Name + "/logon"
}

func (p *AamAuthenticationPortalLogon) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogon::Post")
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

func (p *AamAuthenticationPortalLogon) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogon::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *AamAuthenticationPortalLogon) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogon::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *AamAuthenticationPortalLogon) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogon::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
