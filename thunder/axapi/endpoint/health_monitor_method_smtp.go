

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodSmtp struct {
	Inst struct {

    MailFrom string `json:"mail-from"`
    RcptTo string `json:"rcpt-to"`
    Smtp int `json:"smtp"`
    SmtpDomain string `json:"smtp-domain"`
    SmtpPort int `json:"smtp-port" dval:"25"`
    SmtpStarttls int `json:"smtp-starttls"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"smtp"`
}

func (p *HealthMonitorMethodSmtp) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodSmtp) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/smtp"
}

func (p *HealthMonitorMethodSmtp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSmtp::Post")
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

func (p *HealthMonitorMethodSmtp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSmtp::Get")
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
func (p *HealthMonitorMethodSmtp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSmtp::Put")
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

func (p *HealthMonitorMethodSmtp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSmtp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
