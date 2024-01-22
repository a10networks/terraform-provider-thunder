

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PkiScepCert struct {
	Inst struct {

    Days int `json:"days" dval:"1825"`
    Dn string `json:"dn"`
    Encrypted string `json:"encrypted"`
    EndDate string `json:"end-date"`
    Enroll int `json:"enroll"`
    Interval int `json:"interval" dval:"5"`
    KeyLength string `json:"key-length" dval:"2048"`
    LogLevel int `json:"log-level" dval:"1"`
    MaxPolltime int `json:"max-polltime" dval:"180"`
    Method string `json:"method" dval:"GET"`
    Minute int `json:"minute"`
    Name string `json:"name"`
    Password int `json:"password"`
    RenewBefore int `json:"renew-before"`
    RenewBeforeType string `json:"renew-before-type"`
    RenewBeforeValue int `json:"renew-before-value"`
    RenewEvery int `json:"renew-every"`
    RenewEveryType string `json:"renew-every-type"`
    RenewEveryValue int `json:"renew-every-value"`
    SecretString string `json:"secret-string"`
    StartDate string `json:"start-date"`
    SubjectAlternateName PkiScepCertSubjectAlternateName `json:"subject-alternate-name"`
    Url string `json:"url"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"scep-cert"`
}


type PkiScepCertSubjectAlternateName struct {
    SanType string `json:"san-type"`
    SanValue string `json:"san-value"`
}

func (p *PkiScepCert) GetId() string{
    return p.Inst.Name
}

func (p *PkiScepCert) getPath() string{
    return "pki/scep-cert"
}

func (p *PkiScepCert) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PkiScepCert::Post")
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

func (p *PkiScepCert) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PkiScepCert::Get")
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
func (p *PkiScepCert) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("PkiScepCert::Put")
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

func (p *PkiScepCert) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("PkiScepCert::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
