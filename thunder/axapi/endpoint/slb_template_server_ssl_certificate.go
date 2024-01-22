

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateServerSslCertificate struct {
	Inst struct {

    Cert string `json:"cert"`
    Encrypted string `json:"encrypted"`
    Key string `json:"key"`
    Passphrase string `json:"passphrase"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"certificate"`
}

func (p *SlbTemplateServerSslCertificate) GetId() string{
    return "1"
}

func (p *SlbTemplateServerSslCertificate) getPath() string{
    return "slb/template/server-ssl/" +p.Inst.Name + "/certificate"
}

func (p *SlbTemplateServerSslCertificate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSslCertificate::Post")
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

func (p *SlbTemplateServerSslCertificate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSslCertificate::Get")
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
func (p *SlbTemplateServerSslCertificate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSslCertificate::Put")
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

func (p *SlbTemplateServerSslCertificate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServerSslCertificate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
