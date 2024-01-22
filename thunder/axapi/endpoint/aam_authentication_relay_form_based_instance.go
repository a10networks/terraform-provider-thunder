

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayFormBasedInstance struct {
	Inst struct {

    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    RequestUriList []AamAuthenticationRelayFormBasedInstanceRequestUriList `json:"request-uri-list"`
    SamplingEnable []AamAuthenticationRelayFormBasedInstanceSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"instance"`
}


type AamAuthenticationRelayFormBasedInstanceRequestUriList struct {
    MatchType string `json:"match-type"`
    Uri string `json:"uri"`
    UserVariable string `json:"user-variable"`
    PasswordVariable string `json:"password-variable"`
    DomainVariable string `json:"domain-variable"`
    OtherVariables string `json:"other-variables"`
    MaxPacketCollectSize int `json:"max-packet-collect-size" dval:"1048576"`
    Cookie AamAuthenticationRelayFormBasedInstanceRequestUriListCookie `json:"cookie"`
    ActionUri string `json:"action-uri"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type AamAuthenticationRelayFormBasedInstanceRequestUriListCookie struct {
    CookieValue AamAuthenticationRelayFormBasedInstanceRequestUriListCookieCookieValue `json:"cookie-value"`
}


type AamAuthenticationRelayFormBasedInstanceRequestUriListCookieCookieValue struct {
    CookieValue string `json:"cookie-value"`
}


type AamAuthenticationRelayFormBasedInstanceSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelayFormBasedInstance) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationRelayFormBasedInstance) getPath() string{
    return "aam/authentication/relay/form-based/instance"
}

func (p *AamAuthenticationRelayFormBasedInstance) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstance::Post")
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

func (p *AamAuthenticationRelayFormBasedInstance) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstance::Get")
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
func (p *AamAuthenticationRelayFormBasedInstance) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstance::Put")
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

func (p *AamAuthenticationRelayFormBasedInstance) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayFormBasedInstance::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
