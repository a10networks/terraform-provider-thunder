

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamJwtAuthorization struct {
	Inst struct {

    Encrypted string `json:"encrypted"`
    ExpClaimRequried int `json:"exp-claim-requried"`
    JwtCacheEnable int `json:"jwt-cache-enable"`
    JwtExpDefault int `json:"jwt-exp-default"`
    JwtForwarding int `json:"jwt-forwarding"`
    LogLevel string `json:"log-level"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []AamJwtAuthorizationSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VerificationCert string `json:"verification-cert"`
    VerificationJwks string `json:"verification-jwks"`
    VerificationSecret string `json:"verification-secret"`

	} `json:"jwt-authorization"`
}


type AamJwtAuthorizationSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamJwtAuthorization) GetId() string{
    return p.Inst.Name
}

func (p *AamJwtAuthorization) getPath() string{
    return "aam/jwt-authorization"
}

func (p *AamJwtAuthorization) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamJwtAuthorization::Post")
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

func (p *AamJwtAuthorization) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamJwtAuthorization::Get")
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
func (p *AamJwtAuthorization) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamJwtAuthorization::Put")
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

func (p *AamJwtAuthorization) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamJwtAuthorization::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
