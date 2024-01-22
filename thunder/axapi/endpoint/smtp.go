

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Smtp struct {
	Inst struct {

    Mailfrom string `json:"mailfrom"`
    Needauthentication int `json:"needauthentication"`
    Port int `json:"port"`
    SmtpServer string `json:"smtp-server"`
    SmtpServerV6 string `json:"smtp-server-v6"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UsernameCfg SmtpUsernameCfg `json:"username-cfg"`
    Uuid string `json:"uuid"`

	} `json:"smtp"`
}


type SmtpUsernameCfg struct {
    Username string `json:"username"`
    Password SmtpUsernameCfgPassword `json:"password"`
}


type SmtpUsernameCfgPassword struct {
    SmtpPassword string `json:"smtp-password"`
    Encrypted string `json:"encrypted"`
}

func (p *Smtp) GetId() string{
    return "1"
}

func (p *Smtp) getPath() string{
    return "smtp"
}

func (p *Smtp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Smtp::Post")
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

func (p *Smtp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Smtp::Get")
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
func (p *Smtp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Smtp::Put")
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

func (p *Smtp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Smtp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
