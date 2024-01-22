

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationCaptchaInstance struct {
	Inst struct {

    ClientIpParamName string `json:"client-ip-param-name"`
    Encrypted string `json:"encrypted"`
    Method string `json:"method" dval:"POST"`
    Name string `json:"name"`
    RespErrorCodeFieldName string `json:"resp-error-code-field-name"`
    RespResultFieldName string `json:"resp-result-field-name"`
    SecretKey int `json:"secret-key"`
    SecretKeyParamName string `json:"secret-key-param-name"`
    SecretKeyString string `json:"secret-key-string"`
    SendClientIp int `json:"send-client-ip"`
    Timeout int `json:"timeout" dval:"10"`
    TokenParamName string `json:"token-param-name"`
    Url string `json:"url"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}

func (p *AamAuthenticationCaptchaInstance) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationCaptchaInstance) getPath() string{
    return "aam/authentication/captcha/instance"
}

func (p *AamAuthenticationCaptchaInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptchaInstance::Post")
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

func (p *AamAuthenticationCaptchaInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptchaInstance::Get")
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
func (p *AamAuthenticationCaptchaInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptchaInstance::Put")
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

func (p *AamAuthenticationCaptchaInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationCaptchaInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
