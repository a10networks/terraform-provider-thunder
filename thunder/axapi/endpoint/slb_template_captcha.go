package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateCaptcha struct {
	Inst struct {
		InstanceList []SlbTemplateCaptchaInstanceList `json:"instance-list"`
	} `json:"captcha"`
}

type SlbTemplateCaptchaInstanceList struct {
	Name                   string `json:"name"`
	DefaultCaptcha         int    `json:"default-captcha"`
	CaptchaType            string `json:"captcha-type"`
	SiteKeyString          string `json:"site-key-string"`
	SiteKeyEncrypted       string `json:"site-key-encrypted"`
	SecretKey              int    `json:"secret-key"`
	SecretKeyString        string `json:"secret-key-string"`
	SecretKeyEncrypted     string `json:"secret-key-encrypted"`
	Url                    string `json:"url"`
	Method                 string `json:"method" dval:"POST"`
	Timeout                int    `json:"timeout" dval:"10"`
	SecretKeyParamName     string `json:"secret-key-param-name"`
	TokenParamName         string `json:"token-param-name"`
	RespResultFieldName    string `json:"resp-result-field-name"`
	RespErrorCodeFieldName string `json:"resp-error-code-field-name"`
	SendClientIp           int    `json:"send-client-ip"`
	ClientIpParamName      string `json:"client-ip-param-name"`
	Uuid                   string `json:"uuid"`
}

func (p *SlbTemplateCaptcha) GetId() string {
	return "1"
}

func (p *SlbTemplateCaptcha) getPath() string {
	return "slb/template/captcha"
}

func (p *SlbTemplateCaptcha) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCaptcha::Post")
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

func (p *SlbTemplateCaptcha) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCaptcha::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *SlbTemplateCaptcha) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCaptcha::Put")
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

func (p *SlbTemplateCaptcha) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplateCaptcha::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
