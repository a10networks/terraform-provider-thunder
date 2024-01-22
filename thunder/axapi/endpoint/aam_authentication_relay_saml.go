

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelaySaml struct {
	Inst struct {

    IdpAuthUri string `json:"idp-auth-uri"`
    MatchType string `json:"match-type"`
    MatchUri string `json:"match-uri"`
    Method string `json:"method"`
    Name string `json:"name"`
    RelayAcsUri string `json:"relay-acs-uri"`
    RetryNumber int `json:"retry-number"`
    SamplingEnable []AamAuthenticationRelaySamlSamplingEnable `json:"sampling-enable"`
    ServerCookieName string `json:"server-cookie-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Value string `json:"value"`

	} `json:"saml"`
}


type AamAuthenticationRelaySamlSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelaySaml) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationRelaySaml) getPath() string{
    return "aam/authentication/relay/saml"
}

func (p *AamAuthenticationRelaySaml) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelaySaml::Post")
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

func (p *AamAuthenticationRelaySaml) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelaySaml::Get")
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
func (p *AamAuthenticationRelaySaml) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelaySaml::Put")
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

func (p *AamAuthenticationRelaySaml) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelaySaml::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
