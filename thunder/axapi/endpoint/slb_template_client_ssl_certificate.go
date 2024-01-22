

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateClientSslCertificate struct {
	Inst struct {

    Cert string `json:"cert"`
    ChainCert string `json:"chain-cert"`
    Key string `json:"key"`
    KeyEncrypted string `json:"key-encrypted"`
    Passphrase string `json:"passphrase"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"certificate"`
}

func (p *SlbTemplateClientSslCertificate) GetId() string{
    return url.QueryEscape(p.Inst.Cert)
}

func (p *SlbTemplateClientSslCertificate) getPath() string{
    return "slb/template/client-ssl/" +p.Inst.Name + "/certificate"
}

func (p *SlbTemplateClientSslCertificate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSslCertificate::Post")
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

func (p *SlbTemplateClientSslCertificate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSslCertificate::Get")
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
func (p *SlbTemplateClientSslCertificate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSslCertificate::Put")
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

func (p *SlbTemplateClientSslCertificate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateClientSslCertificate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
