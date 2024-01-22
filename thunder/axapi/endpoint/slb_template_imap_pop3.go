

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateImapPop3 struct {
	Inst struct {

    Logindisabled int `json:"logindisabled"`
    Name string `json:"name"`
    Starttls string `json:"starttls" dval:"disabled"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"imap-pop3"`
}

func (p *SlbTemplateImapPop3) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateImapPop3) getPath() string{
    return "slb/template/imap-pop3"
}

func (p *SlbTemplateImapPop3) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateImapPop3::Post")
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

func (p *SlbTemplateImapPop3) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateImapPop3::Get")
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
func (p *SlbTemplateImapPop3) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateImapPop3::Put")
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

func (p *SlbTemplateImapPop3) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateImapPop3::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
