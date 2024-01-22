

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpLoggingPolicy struct {
	Inst struct {

    Log TemplateGtpLoggingPolicyLog `json:"log"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"logging-policy"`
}


type TemplateGtpLoggingPolicyLog struct {
    MandatoryIeCheck int `json:"mandatory-ie-check"`
    OutOfStateIeCheck int `json:"out-of-state-ie-check"`
    OutOfOrderIeCheck int `json:"out-of-order-ie-check"`
    InvalidTeidCheck int `json:"invalid-teid-check"`
    ReservedIeCheck int `json:"reserved-ie-check"`
    MessageFiltering int `json:"message-filtering"`
    ApnImsiFiltering int `json:"apn-imsi-filtering"`
    RatTypeFiltering int `json:"rat-type-filtering"`
    MsisdnFiltering int `json:"msisdn-filtering"`
    CrosslayerCorrelation int `json:"crosslayer-correlation"`
    AntiSpoofingCheck int `json:"anti-spoofing-check"`
    MsisdnImsiCorrelation int `json:"msisdn-imsi-correlation"`
    MaxMessageLengthCheck int `json:"max-message-length-check"`
    GtpInGtpFiltering int `json:"gtp-in-gtp-filtering"`
    SequenceNumCorrelation int `json:"sequence-num-correlation"`
    InvalidHeaderCheck int `json:"invalid-header-check"`
}

func (p *TemplateGtpLoggingPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpLoggingPolicy) getPath() string{
    return "template/gtp/logging-policy"
}

func (p *TemplateGtpLoggingPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpLoggingPolicy::Post")
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

func (p *TemplateGtpLoggingPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpLoggingPolicy::Get")
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
func (p *TemplateGtpLoggingPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpLoggingPolicy::Put")
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

func (p *TemplateGtpLoggingPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpLoggingPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
