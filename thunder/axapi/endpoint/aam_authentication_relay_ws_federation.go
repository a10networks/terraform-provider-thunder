

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayWsFederation struct {
	Inst struct {

    ApplicationServer string `json:"application-server"`
    AuthenticationUri string `json:"authentication-uri"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []AamAuthenticationRelayWsFederationSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ws-federation"`
}


type AamAuthenticationRelayWsFederationSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelayWsFederation) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationRelayWsFederation) getPath() string{
    return "aam/authentication/relay/ws-federation"
}

func (p *AamAuthenticationRelayWsFederation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayWsFederation::Post")
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

func (p *AamAuthenticationRelayWsFederation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayWsFederation::Get")
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
func (p *AamAuthenticationRelayWsFederation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayWsFederation::Put")
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

func (p *AamAuthenticationRelayWsFederation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayWsFederation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
