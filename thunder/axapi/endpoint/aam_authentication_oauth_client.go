

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationOauthClient struct {
	Inst struct {

    ClientId string `json:"client-id"`
    ClientSecret string `json:"client-secret"`
    Encrypted string `json:"encrypted"`
    GrantType string `json:"grant-type"`
    Infinity int `json:"infinity"`
    Name string `json:"name"`
    NoReply int `json:"no-reply"`
    ParameterNonceEnable int `json:"parameter-nonce-enable"`
    RedirectionEndpoint string `json:"redirection-endpoint"`
    Scope string `json:"scope"`
    SessionInitTtl int `json:"session-init-ttl"`
    TokenLifetime int `json:"token-lifetime"`
    Type string `json:"type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"client"`
}

func (p *AamAuthenticationOauthClient) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationOauthClient) getPath() string{
    return "aam/authentication/oauth/client"
}

func (p *AamAuthenticationOauthClient) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthClient::Post")
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

func (p *AamAuthenticationOauthClient) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthClient::Get")
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
func (p *AamAuthenticationOauthClient) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthClient::Put")
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

func (p *AamAuthenticationOauthClient) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationOauthClient::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
