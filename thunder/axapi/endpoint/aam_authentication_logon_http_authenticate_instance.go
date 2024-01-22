

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationLogonHttpAuthenticateInstance struct {
	Inst struct {

    AccountLock int `json:"account-lock"`
    AuthMethod AamAuthenticationLogonHttpAuthenticateInstanceAuthMethod `json:"auth-method"`
    Duration int `json:"duration" dval:"1800"`
    HstsTimeout int `json:"hsts-timeout"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Retry int `json:"retry" dval:"3"`
    SamplingEnable []AamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}


type AamAuthenticationLogonHttpAuthenticateInstanceAuthMethod struct {
    Basic AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodBasic `json:"basic"`
    Ntlm AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNtlm `json:"ntlm"`
    Negotiate AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNegotiate `json:"negotiate"`
}


type AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodBasic struct {
    BasicRealm string `json:"basic-realm"`
    ChallengeResponseForm string `json:"challenge-response-form"`
    ChallengePage string `json:"challenge-page"`
    ChallengeVariable string `json:"challenge-variable"`
    NewPinPage string `json:"new-pin-page"`
    NextTokenPage string `json:"next-token-page"`
    NewPinVariable string `json:"new-pin-variable"`
    NextTokenVariable string `json:"next-token-variable"`
    BasicEnable int `json:"basic-enable"`
}


type AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNtlm struct {
    NtlmEnable int `json:"ntlm-enable"`
}


type AamAuthenticationLogonHttpAuthenticateInstanceAuthMethodNegotiate struct {
    NegotiateEnable int `json:"negotiate-enable"`
}


type AamAuthenticationLogonHttpAuthenticateInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationLogonHttpAuthenticateInstance) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *AamAuthenticationLogonHttpAuthenticateInstance) getPath() string{
    return "aam/authentication/logon/http-authenticate/instance"
}

func (p *AamAuthenticationLogonHttpAuthenticateInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonHttpAuthenticateInstance::Post")
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

func (p *AamAuthenticationLogonHttpAuthenticateInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonHttpAuthenticateInstance::Get")
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
func (p *AamAuthenticationLogonHttpAuthenticateInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonHttpAuthenticateInstance::Put")
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

func (p *AamAuthenticationLogonHttpAuthenticateInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationLogonHttpAuthenticateInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
