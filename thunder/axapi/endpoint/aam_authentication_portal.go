

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationPortal struct {
	Inst struct {

    ChangePassword AamAuthenticationPortalChangePassword0 `json:"change-password"`
    LogoCfg AamAuthenticationPortalLogoCfg `json:"logo-cfg"`
    Logon AamAuthenticationPortalLogon7 `json:"logon"`
    LogonFail AamAuthenticationPortalLogonFail14 `json:"logon-fail"`
    Name string `json:"name"`
    NotifyChangePassword AamAuthenticationPortalNotifyChangePassword18 `json:"notify-change-password"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"portal"`
}


type AamAuthenticationPortalChangePassword0 struct {
    Background AamAuthenticationPortalChangePasswordBackground1 `json:"background"`
    TitleCfg AamAuthenticationPortalChangePasswordTitleCfg2 `json:"title-cfg"`
    ActionUrl string `json:"action-url"`
    UsernameCfg AamAuthenticationPortalChangePasswordUsernameCfg3 `json:"username-cfg"`
    UsernameVar string `json:"username-var"`
    OldPwdCfg AamAuthenticationPortalChangePasswordOldPwdCfg4 `json:"old-pwd-cfg"`
    OldPasswordVar string `json:"old-password-var"`
    NewPwdCfg AamAuthenticationPortalChangePasswordNewPwdCfg5 `json:"new-pwd-cfg"`
    NewPasswordVar string `json:"new-password-var"`
    CfmPwdCfg AamAuthenticationPortalChangePasswordCfmPwdCfg6 `json:"cfm-pwd-cfg"`
    ConfirmPasswordVar string `json:"confirm-password-var"`
    SubmitText string `json:"submit-text"`
    ResetText string `json:"reset-text"`
    Uuid string `json:"uuid"`
}


type AamAuthenticationPortalChangePasswordBackground1 struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalChangePasswordTitleCfg2 struct {
    Title int `json:"title"`
    TitleText string `json:"title-text"`
    TitleFont int `json:"title-font" dval:"1"`
    TitleFace string `json:"title-face" dval:"Arial"`
    TitleFontCustom string `json:"title-font-custom"`
    TitleSize int `json:"title-size" dval:"5"`
    TitleColor int `json:"title-color" dval:"1"`
    TitleColorName string `json:"title-color-name" dval:"black"`
    TitleColorValue string `json:"title-color-value"`
}


type AamAuthenticationPortalChangePasswordUsernameCfg3 struct {
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


type AamAuthenticationPortalChangePasswordOldPwdCfg4 struct {
    OldPassword int `json:"old-password"`
    OldText string `json:"old-text"`
    OldFont int `json:"old-font" dval:"1"`
    OldFace string `json:"old-face" dval:"Arial"`
    OldFontCustom string `json:"old-font-custom"`
    OldSize int `json:"old-size" dval:"3"`
    OldColor int `json:"old-color" dval:"1"`
    OldColorName string `json:"old-color-name" dval:"black"`
    OldColorValue string `json:"old-color-value"`
}


type AamAuthenticationPortalChangePasswordNewPwdCfg5 struct {
    NewPassword int `json:"new-password"`
    NewText string `json:"new-text"`
    NewFont int `json:"new-font" dval:"1"`
    NewFace string `json:"new-face" dval:"Arial"`
    NewFontCustom string `json:"new-font-custom"`
    NewSize int `json:"new-size" dval:"3"`
    NewColor int `json:"new-color" dval:"1"`
    NewColorName string `json:"new-color-name" dval:"black"`
    NewColorValue string `json:"new-color-value"`
}


type AamAuthenticationPortalChangePasswordCfmPwdCfg6 struct {
    ConfirmPassword int `json:"confirm-password"`
    CfmText string `json:"cfm-text"`
    CfmFont int `json:"cfm-font" dval:"1"`
    CfmFace string `json:"cfm-face" dval:"Arial"`
    CfmFontCustom string `json:"cfm-font-custom"`
    CfmSize int `json:"cfm-size" dval:"3"`
    CfmColor int `json:"cfm-color" dval:"1"`
    CfmColorName string `json:"cfm-color-name" dval:"black"`
    CfmColorValue string `json:"cfm-color-value"`
}


type AamAuthenticationPortalLogoCfg struct {
    Logo string `json:"logo"`
    Width int `json:"width" dval:"134"`
    Height int `json:"height" dval:"71"`
}


type AamAuthenticationPortalLogon7 struct {
    Background AamAuthenticationPortalLogonBackground8 `json:"background"`
    FailMsgCfg AamAuthenticationPortalLogonFailMsgCfg9 `json:"fail-msg-cfg"`
    ActionUrl string `json:"action-url"`
    UsernameCfg AamAuthenticationPortalLogonUsernameCfg10 `json:"username-cfg"`
    UsernameVar string `json:"username-var"`
    PasswordCfg AamAuthenticationPortalLogonPasswordCfg11 `json:"password-cfg"`
    PasswordVar string `json:"password-var"`
    EnablePasscode int `json:"enable-passcode"`
    PasscodeCfg AamAuthenticationPortalLogonPasscodeCfg12 `json:"passcode-cfg"`
    PasscodeVar string `json:"passcode-var"`
    EnableCaptcha int `json:"enable-CAPTCHA"`
    CaptchaType string `json:"captcha-type"`
    SiteKeyString string `json:"site-key-string"`
    Encrypted string `json:"encrypted"`
    RecaptchaCfg AamAuthenticationPortalLogonRecaptchaCfg13 `json:"reCAPTCHA-cfg"`
    SubmitText string `json:"submit-text"`
    Uuid string `json:"uuid"`
}


type AamAuthenticationPortalLogonBackground8 struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalLogonFailMsgCfg9 struct {
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


type AamAuthenticationPortalLogonUsernameCfg10 struct {
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


type AamAuthenticationPortalLogonPasswordCfg11 struct {
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


type AamAuthenticationPortalLogonPasscodeCfg12 struct {
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


type AamAuthenticationPortalLogonRecaptchaCfg13 struct {
    RecaptchaTheme string `json:"reCAPTCHA-theme" dval:"light"`
    RecaptchaSize string `json:"reCAPTCHA-size" dval:"normal"`
    RecaptchaBadge string `json:"reCAPTCHA-badge" dval:"bottom-right"`
    RecaptchaAction string `json:"reCAPTCHA-action" dval:"A10_DEFAULT_LOGON"`
}


type AamAuthenticationPortalLogonFail14 struct {
    Background AamAuthenticationPortalLogonFailBackground15 `json:"background"`
    TitleCfg AamAuthenticationPortalLogonFailTitleCfg16 `json:"title-cfg"`
    FailMsgCfg AamAuthenticationPortalLogonFailFailMsgCfg17 `json:"fail-msg-cfg"`
    Uuid string `json:"uuid"`
}


type AamAuthenticationPortalLogonFailBackground15 struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalLogonFailTitleCfg16 struct {
    Title int `json:"title"`
    TitleText string `json:"title-text"`
    TitleFont int `json:"title-font" dval:"1"`
    TitleFace string `json:"title-face" dval:"Arial"`
    TitleFontCustom string `json:"title-font-custom"`
    TitleSize int `json:"title-size" dval:"5"`
    TitleColor int `json:"title-color" dval:"1"`
    TitleColorName string `json:"title-color-name" dval:"black"`
    TitleColorValue string `json:"title-color-value"`
}


type AamAuthenticationPortalLogonFailFailMsgCfg17 struct {
    FailMsg int `json:"fail-msg"`
    FailText string `json:"fail-text"`
    FailFont int `json:"fail-font" dval:"1"`
    FailFace string `json:"fail-face" dval:"Arial"`
    FailFontCustom string `json:"fail-font-custom"`
    FailSize int `json:"fail-size" dval:"3"`
    FailColor int `json:"fail-color" dval:"1"`
    FailColorName string `json:"fail-color-name" dval:"black"`
    FailColorValue string `json:"fail-color-value"`
}


type AamAuthenticationPortalNotifyChangePassword18 struct {
    Background AamAuthenticationPortalNotifyChangePasswordBackground19 `json:"background"`
    ContinueUrl string `json:"continue-url"`
    ChangeUrl string `json:"change-url"`
    UsernameCfg AamAuthenticationPortalNotifyChangePasswordUsernameCfg20 `json:"username-cfg"`
    UsernameVar string `json:"username-var"`
    OldPwdCfg AamAuthenticationPortalNotifyChangePasswordOldPwdCfg21 `json:"old-pwd-cfg"`
    OldPasswordVar string `json:"old-password-var"`
    NewPwdCfg AamAuthenticationPortalNotifyChangePasswordNewPwdCfg22 `json:"new-pwd-cfg"`
    NewPasswordVar string `json:"new-password-var"`
    CfmPwdCfg AamAuthenticationPortalNotifyChangePasswordCfmPwdCfg23 `json:"cfm-pwd-cfg"`
    ConfirmPasswordVar string `json:"confirm-password-var"`
    ChangeText string `json:"change-text"`
    ContinueText string `json:"continue-text"`
    Uuid string `json:"uuid"`
}


type AamAuthenticationPortalNotifyChangePasswordBackground19 struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalNotifyChangePasswordUsernameCfg20 struct {
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


type AamAuthenticationPortalNotifyChangePasswordOldPwdCfg21 struct {
    OldPassword int `json:"old-password"`
    OldText string `json:"old-text"`
    OldFont int `json:"old-font" dval:"1"`
    OldFace string `json:"old-face" dval:"Arial"`
    OldFontCustom string `json:"old-font-custom"`
    OldSize int `json:"old-size" dval:"3"`
    OldColor int `json:"old-color" dval:"1"`
    OldColorName string `json:"old-color-name" dval:"black"`
    OldColorValue string `json:"old-color-value"`
}


type AamAuthenticationPortalNotifyChangePasswordNewPwdCfg22 struct {
    NewPassword int `json:"new-password"`
    NewText string `json:"new-text"`
    NewFont int `json:"new-font" dval:"1"`
    NewFace string `json:"new-face" dval:"Arial"`
    NewFontCustom string `json:"new-font-custom"`
    NewSize int `json:"new-size" dval:"3"`
    NewColor int `json:"new-color" dval:"1"`
    NewColorName string `json:"new-color-name" dval:"black"`
    NewColorValue string `json:"new-color-value"`
}


type AamAuthenticationPortalNotifyChangePasswordCfmPwdCfg23 struct {
    ConfirmPassword int `json:"confirm-password"`
    CfmText string `json:"cfm-text"`
    CfmFont int `json:"cfm-font" dval:"1"`
    CfmFace string `json:"cfm-face" dval:"Arial"`
    CfmFontCustom string `json:"cfm-font-custom"`
    CfmSize int `json:"cfm-size" dval:"3"`
    CfmColor int `json:"cfm-color" dval:"1"`
    CfmColorName string `json:"cfm-color-name" dval:"black"`
    CfmColorValue string `json:"cfm-color-value"`
}

func (p *AamAuthenticationPortal) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationPortal) getPath() string{
    return "aam/authentication/portal"
}

func (p *AamAuthenticationPortal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortal::Post")
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

func (p *AamAuthenticationPortal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortal::Get")
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
func (p *AamAuthenticationPortal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortal::Put")
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

func (p *AamAuthenticationPortal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
