

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationPortalLogonFail struct {
	Inst struct {

    Background AamAuthenticationPortalLogonFailBackground `json:"background"`
    FailMsgCfg AamAuthenticationPortalLogonFailFailMsgCfg `json:"fail-msg-cfg"`
    TitleCfg AamAuthenticationPortalLogonFailTitleCfg `json:"title-cfg"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"logon-fail"`
}


type AamAuthenticationPortalLogonFailBackground struct {
    Bgfile string `json:"bgfile"`
    Bgstyle string `json:"bgstyle" dval:"tile"`
    BgcolorName string `json:"bgcolor-name" dval:"white"`
    BgcolorValue string `json:"bgcolor-value"`
}


type AamAuthenticationPortalLogonFailFailMsgCfg struct {
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


type AamAuthenticationPortalLogonFailTitleCfg struct {
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

func (p *AamAuthenticationPortalLogonFail) GetId() string{
    return "1"
}

func (p *AamAuthenticationPortalLogonFail) getPath() string{
    return "aam/authentication/portal/" +p.Inst.Name + "/logon-fail"
}

func (p *AamAuthenticationPortalLogonFail) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogonFail::Post")
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

func (p *AamAuthenticationPortalLogonFail) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogonFail::Get")
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
func (p *AamAuthenticationPortalLogonFail) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogonFail::Put")
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

func (p *AamAuthenticationPortalLogonFail) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationPortalLogonFail::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
