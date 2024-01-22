

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationPortalChangePassword struct {
	Inst struct {

    ActionUrl string `json:"action-url"`
    Background AamAuthenticationPortalChangePasswordBackground `json:"background"`
    CfmPwdCfg AamAuthenticationPortalChangePasswordCfmPwdCfg `json:"cfm-pwd-cfg"`
    ConfirmPasswordVar string `json:"confirm-password-var"`
    NewPasswordVar string `json:"new-password-var"`
    NewPwdCfg AamAuthenticationPortalChangePasswordNewPwdCfg `json:"new-pwd-cfg"`
    OldPasswordVar string `json:"old-password-var"`
    OldPwdCfg AamAuthenticationPortalChangePasswordOldPwdCfg `json:"old-pwd-cfg"`
    ResetText string `json:"reset-text"`
    SubmitText string `json:"submit-text"`
    TitleCfg AamAuthenticationPortalChangePasswordTitleCfg `json:"title-cfg"`
    UsernameCfg AamAuthenticationPortalChangePasswordUsernameCfg `json:"username-cfg"`
    UsernameVar string `json:"username-var"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"change-password"`
}


type AamAuthenticationPortalChangePasswordBackground struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalChangePasswordCfmPwdCfg struct {
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


type AamAuthenticationPortalChangePasswordNewPwdCfg struct {
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


type AamAuthenticationPortalChangePasswordOldPwdCfg struct {
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


type AamAuthenticationPortalChangePasswordTitleCfg struct {
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


type AamAuthenticationPortalChangePasswordUsernameCfg struct {
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

func (p *AamAuthenticationPortalChangePassword) GetId() string{
    return "1"
}

func (p *AamAuthenticationPortalChangePassword) getPath() string{
    return "aam/authentication/portal/" +p.Inst.Name + "/change-password"
}

func (p *AamAuthenticationPortalChangePassword) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalChangePassword::Post")
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

func (p *AamAuthenticationPortalChangePassword) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalChangePassword::Get")
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
func (p *AamAuthenticationPortalChangePassword) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalChangePassword::Put")
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

func (p *AamAuthenticationPortalChangePassword) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalChangePassword::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
