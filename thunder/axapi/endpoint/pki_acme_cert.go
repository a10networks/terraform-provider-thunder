

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PkiAcmeCert struct {
	Inst struct {

    CertType int `json:"cert-type"`
    Domain string `json:"domain"`
    EabHmacKey int `json:"eab-hmac-key"`
    EabKeyId string `json:"eab-key-id"`
    EcKeyLength string `json:"ec-key-length" dval:"384"`
    EcdsaType int `json:"ecdsa-type"`
    Email string `json:"email"`
    Encrypted string `json:"encrypted"`
    Enroll int `json:"enroll"`
    Force int `json:"force"`
    LogLevel int `json:"log-level" dval:"1"`
    Minute int `json:"minute"`
    Name string `json:"name"`
    RenewBefore int `json:"renew-before"`
    RenewBeforeType string `json:"renew-before-type"`
    RenewBeforeValue int `json:"renew-before-value"`
    RenewEvery int `json:"renew-every"`
    RenewEveryType string `json:"renew-every-type"`
    RenewEveryValue int `json:"renew-every-value"`
    RsaKeyLength string `json:"rsa-key-length" dval:"2048"`
    RsaType int `json:"rsa-type"`
    SanDomain string `json:"san-domain"`
    SecretString string `json:"secret-string"`
    Staging int `json:"staging"`
    StagingUrl string `json:"staging-url"`
    Url string `json:"url"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"acme-cert"`
}

func (p *PkiAcmeCert) GetId() string{
    return p.Inst.Name
}

func (p *PkiAcmeCert) getPath() string{
    return "pki/acme-cert"
}

func (p *PkiAcmeCert) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PkiAcmeCert::Post")
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

func (p *PkiAcmeCert) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PkiAcmeCert::Get")
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
func (p *PkiAcmeCert) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PkiAcmeCert::Put")
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

func (p *PkiAcmeCert) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PkiAcmeCert::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
