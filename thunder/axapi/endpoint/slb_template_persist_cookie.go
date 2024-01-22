

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePersistCookie struct {
	Inst struct {

    CookieName string `json:"cookie-name" dval:"sto-id"`
    Domain string `json:"domain"`
    DontHonorConnRules int `json:"dont-honor-conn-rules"`
    EncryptLevel int `json:"encrypt-level" dval:"1"`
    Encrypted string `json:"encrypted"`
    Expire int `json:"expire" dval:"31536000"`
    Httponly int `json:"httponly"`
    InsertAlways int `json:"insert-always"`
    MatchType int `json:"match-type"`
    Name string `json:"name"`
    PassPhrase string `json:"pass-phrase" dval:"ACOS4KEY"`
    PassThru int `json:"pass-thru"`
    Path string `json:"path" dval:"/"`
    Prefix string `json:"prefix"`
    Samesite string `json:"samesite"`
    ScanAllMembers int `json:"scan-all-members"`
    Secure int `json:"secure"`
    Server int `json:"server"`
    ServerServiceGroup int `json:"server-service-group"`
    ServiceGroup int `json:"service-group"`
    UseAttribute string `json:"use-attribute" dval:"expires"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"cookie"`
}

func (p *SlbTemplatePersistCookie) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePersistCookie) getPath() string{
    return "slb/template/persist/cookie"
}

func (p *SlbTemplatePersistCookie) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistCookie::Post")
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

func (p *SlbTemplatePersistCookie) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistCookie::Get")
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
func (p *SlbTemplatePersistCookie) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistCookie::Put")
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

func (p *SlbTemplatePersistCookie) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePersistCookie::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
