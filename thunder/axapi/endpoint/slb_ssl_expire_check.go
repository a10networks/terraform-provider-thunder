

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslExpireCheck struct {
	Inst struct {

    Before int `json:"before" dval:"5"`
    Exception SlbSslExpireCheckException1414 `json:"exception"`
    ExpireAddress1 string `json:"expire-address1"`
    IntervalDays int `json:"interval-days" dval:"2"`
    SslExpireEmailAddress string `json:"ssl-expire-email-address"`
    Uuid string `json:"uuid"`

	} `json:"ssl-expire-check"`
}


type SlbSslExpireCheckException1414 struct {
    Action string `json:"action"`
    CertificateName string `json:"certificate-name"`
}

func (p *SlbSslExpireCheck) GetId() string{
    return "1"
}

func (p *SlbSslExpireCheck) getPath() string{
    return "slb/ssl-expire-check"
}

func (p *SlbSslExpireCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheck::Post")
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

func (p *SlbSslExpireCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheck::Get")
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
func (p *SlbSslExpireCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheck::Put")
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

func (p *SlbSslExpireCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
