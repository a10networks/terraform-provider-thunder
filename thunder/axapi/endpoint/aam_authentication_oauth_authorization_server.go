

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationOauthAuthorizationServer struct {
	Inst struct {

    AuthorizationEndpoint string `json:"authorization-endpoint"`
    ClientMethod string `json:"client-method"`
    Issuer string `json:"issuer"`
    Name string `json:"name"`
    SamplingEnable []AamAuthenticationOauthAuthorizationServerSamplingEnable `json:"sampling-enable"`
    ServerMethod string `json:"server-method"`
    TokenEndpoint string `json:"token-endpoint"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VerificationCert string `json:"verification-cert"`
    VerificationJwks string `json:"verification-jwks"`

	} `json:"authorization-server"`
}


type AamAuthenticationOauthAuthorizationServerSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationOauthAuthorizationServer) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationOauthAuthorizationServer) getPath() string{
    return "aam/authentication/oauth/authorization-server"
}

func (p *AamAuthenticationOauthAuthorizationServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthAuthorizationServer::Post")
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

func (p *AamAuthenticationOauthAuthorizationServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthAuthorizationServer::Get")
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
func (p *AamAuthenticationOauthAuthorizationServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthAuthorizationServer::Put")
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

func (p *AamAuthenticationOauthAuthorizationServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthAuthorizationServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
