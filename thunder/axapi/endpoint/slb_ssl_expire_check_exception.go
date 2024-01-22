

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslExpireCheckException struct {
	Inst struct {

    Action string `json:"action"`
    CertificateName string `json:"certificate-name"`

	} `json:"exception"`
}

func (p *SlbSslExpireCheckException) GetId() string{
    return "1"
}

func (p *SlbSslExpireCheckException) getPath() string{
    return "slb/ssl-expire-check/exception"
}

func (p *SlbSslExpireCheckException) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheckException::Post")
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

func (p *SlbSslExpireCheckException) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheckException::Get")
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
func (p *SlbSslExpireCheckException) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheckException::Put")
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

func (p *SlbSslExpireCheckException) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbSslExpireCheckException::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
