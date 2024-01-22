

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayOauth struct {
	Inst struct {

    All int `json:"all"`
    MatchType string `json:"match-type"`
    MatchUri string `json:"match-uri"`
    Name string `json:"name"`
    RelayType string `json:"relay-type"`
    SamplingEnable []AamAuthenticationRelayOauthSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"oauth"`
}


type AamAuthenticationRelayOauthSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelayOauth) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationRelayOauth) getPath() string{
    return "aam/authentication/relay/oauth"
}

func (p *AamAuthenticationRelayOauth) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayOauth::Post")
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

func (p *AamAuthenticationRelayOauth) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayOauth::Get")
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
func (p *AamAuthenticationRelayOauth) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayOauth::Put")
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

func (p *AamAuthenticationRelayOauth) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayOauth::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
