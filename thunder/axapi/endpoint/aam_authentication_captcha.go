

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationCaptcha struct {
	Inst struct {

    InstanceList []AamAuthenticationCaptchaInstanceList `json:"instance-list"`
    Uuid string `json:"uuid"`

	} `json:"captcha"`
}


type AamAuthenticationCaptchaInstanceList struct {
    Name string `json:"name"`
    SecretKey int `json:"secret-key"`
    SecretKeyString string `json:"secret-key-string"`
    Encrypted string `json:"encrypted"`
    Url string `json:"url"`
    Method string `json:"method" dval:"POST"`
    Timeout int `json:"timeout" dval:"10"`
    SecretKeyParamName string `json:"secret-key-param-name"`
    TokenParamName string `json:"token-param-name"`
    RespResultFieldName string `json:"resp-result-field-name"`
    RespErrorCodeFieldName string `json:"resp-error-code-field-name"`
    SendClientIp int `json:"send-client-ip"`
    ClientIpParamName string `json:"client-ip-param-name"`
    Uuid string `json:"uuid"`
}

func (p *AamAuthenticationCaptcha) GetId() string{
    return "1"
}

func (p *AamAuthenticationCaptcha) getPath() string{
    return "aam/authentication/captcha"
}

func (p *AamAuthenticationCaptcha) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptcha::Post")
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

func (p *AamAuthenticationCaptcha) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptcha::Get")
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
func (p *AamAuthenticationCaptcha) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptcha::Put")
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

func (p *AamAuthenticationCaptcha) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptcha::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
