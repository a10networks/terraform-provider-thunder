

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PkiCmpCert struct {
	Inst struct {

    AllowUnprotectedErrors int `json:"allow-unprotected-errors"`
    CertType int `json:"cert-type"`
    CmpTrustedCa string `json:"cmp-trusted-ca"`
    CmpTrustedCert string `json:"cmp-trusted-cert"`
    EcKeyLength string `json:"ec-key-length" dval:"384"`
    EcdsaType int `json:"ecdsa-type"`
    Encrypted string `json:"encrypted"`
    Enroll int `json:"enroll"`
    LogLevel int `json:"log-level" dval:"1"`
    MaxPolltime int `json:"max-polltime" dval:"120"`
    Minute int `json:"minute"`
    Name string `json:"name"`
    RecipientDn string `json:"recipient-dn"`
    RenewBefore int `json:"renew-before"`
    RenewBeforeType string `json:"renew-before-type"`
    RenewBeforeValue int `json:"renew-before-value"`
    RenewEvery int `json:"renew-every"`
    RenewEveryType string `json:"renew-every-type"`
    RenewEveryValue int `json:"renew-every-value"`
    RsaKeyLength string `json:"rsa-key-length" dval:"2048"`
    RsaType int `json:"rsa-type"`
    Secret int `json:"secret"`
    SecretString string `json:"secret-string"`
    SubjectAlternateName PkiCmpCertSubjectAlternateName `json:"subject-alternate-name"`
    SubjectDn string `json:"subject-dn"`
    Url string `json:"url"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"cmp-cert"`
}


type PkiCmpCertSubjectAlternateName struct {
    SanType string `json:"san-type"`
    SanValue string `json:"san-value"`
}

func (p *PkiCmpCert) GetId() string{
    return p.Inst.Name
}

func (p *PkiCmpCert) getPath() string{
    return "pki/cmp-cert"
}

func (p *PkiCmpCert) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PkiCmpCert::Post")
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

func (p *PkiCmpCert) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PkiCmpCert::Get")
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
func (p *PkiCmpCert) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PkiCmpCert::Put")
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

func (p *PkiCmpCert) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PkiCmpCert::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
