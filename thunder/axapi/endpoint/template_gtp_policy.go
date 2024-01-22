

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpPolicy struct {
	Inst struct {

    FilteringPolicyName string `json:"filtering-policy-name"`
    GeneralPolicyName string `json:"general-policy-name"`
    LoggingPolicyName string `json:"logging-policy-name"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    RateLimitPolicyName string `json:"rate-limit-policy-name"`
    SamplingEnable []TemplateGtpPolicySamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ValidationPolicyName string `json:"validation-policy-name"`

	} `json:"gtp-policy"`
}


type TemplateGtpPolicySamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
}

func (p *TemplateGtpPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpPolicy) getPath() string{
    return "template/gtp-policy"
}

func (p *TemplateGtpPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpPolicy::Post")
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

func (p *TemplateGtpPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpPolicy::Get")
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
func (p *TemplateGtpPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpPolicy::Put")
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

func (p *TemplateGtpPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
