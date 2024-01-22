

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpValidationPolicy struct {
	Inst struct {

    AnomalyAction string `json:"anomaly-action" dval:"drop"`
    AnomalyChecks string `json:"anomaly-checks" dval:"enable"`
    AntiSpoofingAction string `json:"anti-spoofing-action" dval:"drop"`
    AntiSpoofingCheck string `json:"anti-spoofing-check" dval:"disable"`
    CrosslayerCorrAction string `json:"crosslayer-corr-action" dval:"drop"`
    CrosslayerCorrelation string `json:"crosslayer-correlation" dval:"disable"`
    MandatoryIeAction string `json:"mandatory-ie-action" dval:"drop"`
    MandatoryIeCheck string `json:"mandatory-ie-check" dval:"enable"`
    MsisdnImsiCorrAction string `json:"msisdn-imsi-corr-action" dval:"drop"`
    MsisdnImsiCorrelation string `json:"msisdn-imsi-correlation" dval:"disable"`
    Name string `json:"name"`
    OutOfOrderIeAction string `json:"out-of-order-ie-action" dval:"drop"`
    OutOfOrderIeCheck string `json:"out-of-order-ie-check" dval:"disable"`
    OutOfStateIeAction string `json:"out-of-state-ie-action" dval:"drop"`
    OutOfStateIeCheck string `json:"out-of-state-ie-check" dval:"disable"`
    ReservedIeAction string `json:"reserved-ie-action" dval:"drop"`
    ReservedIeCheck string `json:"reserved-ie-check" dval:"enable"`
    SequenceNumCorrAction string `json:"sequence-num-corr-action" dval:"drop"`
    SequenceNumCorrelation string `json:"sequence-num-correlation" dval:"disable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"validation-policy"`
}

func (p *TemplateGtpValidationPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpValidationPolicy) getPath() string{
    return "template/gtp/validation-policy"
}

func (p *TemplateGtpValidationPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpValidationPolicy::Post")
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

func (p *TemplateGtpValidationPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpValidationPolicy::Get")
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
func (p *TemplateGtpValidationPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpValidationPolicy::Put")
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

func (p *TemplateGtpValidationPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpValidationPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
