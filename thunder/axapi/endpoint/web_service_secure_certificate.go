

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebServiceSecureCertificate struct {
	Inst struct {

    FileUrl string `json:"file-url"`
    Load int `json:"load"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"certificate"`
}

func (p *WebServiceSecureCertificate) GetId() string{
    return "1"
}

func (p *WebServiceSecureCertificate) getPath() string{
    return "web-service/secure/certificate"
}

func (p *WebServiceSecureCertificate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecureCertificate::Post")
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

func (p *WebServiceSecureCertificate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecureCertificate::Get")
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
func (p *WebServiceSecureCertificate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecureCertificate::Put")
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

func (p *WebServiceSecureCertificate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WebServiceSecureCertificate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
