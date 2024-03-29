

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationJwt struct {
	Inst struct {

    Action string `json:"action"`
    Encrypted string `json:"encrypted"`
    Issuer string `json:"issuer"`
    JwtRelayUri string `json:"jwt-relay-uri"`
    Name string `json:"name"`
    SecretString string `json:"secret-string"`
    SignatureSecret int `json:"signature-secret"`
    TokenLifetime int `json:"token-lifetime" dval:"300"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"jwt"`
}

func (p *AamAuthenticationJwt) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationJwt) getPath() string{
    return "aam/authentication/jwt"
}

func (p *AamAuthenticationJwt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationJwt::Post")
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

func (p *AamAuthenticationJwt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationJwt::Get")
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
func (p *AamAuthenticationJwt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationJwt::Put")
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

func (p *AamAuthenticationJwt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationJwt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
