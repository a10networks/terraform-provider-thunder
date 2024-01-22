

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateSmtp struct {
	Inst struct {

    ClientDomainSwitching []SlbTemplateSmtpClientDomainSwitching `json:"client-domain-switching"`
    ClientStarttlsType string `json:"client-starttls-type"`
    CommandDisable []SlbTemplateSmtpCommandDisable `json:"command-disable"`
    ErrorCodeToClient int `json:"error-code-to-client"`
    LfToCrlf int `json:"LF-to-CRLF"`
    Name string `json:"name"`
    ServerDomain string `json:"server-domain" dval:"mail-server-domain"`
    ServerStarttlsType string `json:"server-starttls-type"`
    ServiceReadyMsg string `json:"service-ready-msg" dval:"ESMTP mail service ready"`
    Template SlbTemplateSmtpTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"smtp"`
}


type SlbTemplateSmtpClientDomainSwitching struct {
    SwitchingType string `json:"switching-type"`
    MatchString string `json:"match-string"`
    ServiceGroup string `json:"service-group"`
}


type SlbTemplateSmtpCommandDisable struct {
    DisableType string `json:"disable-type"`
}


type SlbTemplateSmtpTemplate struct {
    Logging string `json:"logging"`
}

func (p *SlbTemplateSmtp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateSmtp) getPath() string{
    return "slb/template/smtp"
}

func (p *SlbTemplateSmtp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmtp::Post")
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

func (p *SlbTemplateSmtp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmtp::Get")
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
func (p *SlbTemplateSmtp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmtp::Put")
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

func (p *SlbTemplateSmtp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmtp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
