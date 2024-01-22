

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationRelayNtlm struct {
	Inst struct {

    Domain string `json:"domain"`
    LargeRequestDisable int `json:"large-request-disable"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    SamplingEnable []AamAuthenticationRelayNtlmSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Version int `json:"version" dval:"2"`

	} `json:"ntlm"`
}


type AamAuthenticationRelayNtlmSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *AamAuthenticationRelayNtlm) GetId() string{
    return p.Inst.Name
}

func (p *AamAuthenticationRelayNtlm) getPath() string{
    return "aam/authentication/relay/ntlm"
}

func (p *AamAuthenticationRelayNtlm) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayNtlm::Post")
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

func (p *AamAuthenticationRelayNtlm) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayNtlm::Get")
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
func (p *AamAuthenticationRelayNtlm) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayNtlm::Put")
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

func (p *AamAuthenticationRelayNtlm) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationRelayNtlm::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
