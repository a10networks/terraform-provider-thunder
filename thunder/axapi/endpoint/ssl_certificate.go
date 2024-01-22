

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SslCertificate struct {
	Inst struct {

    CertificateType string `json:"certificate-type"`
    Name string `json:"name"`
    PfxPassword string `json:"pfx-password"`
    PublicKey string `json:"public-key"`

	} `json:"certificate"`
}

func (p *SslCertificate) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SslCertificate) getPath() string{
    return "ssl/certificate"
}

func (p *SslCertificate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SslCertificate::Post")
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

func (p *SslCertificate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SslCertificate::Get")
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
func (p *SslCertificate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SslCertificate::Put")
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

func (p *SslCertificate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SslCertificate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
