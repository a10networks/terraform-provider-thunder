

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpFilteringPolicy struct {
	Inst struct {

    ApnImsiFiltAction string `json:"apn-imsi-filt-action" dval:"drop"`
    ApnImsiFiltering string `json:"apn-imsi-filtering"`
    GtpInGtpFiltAction string `json:"gtp-in-gtp-filt-action" dval:"drop"`
    GtpInGtpFiltering string `json:"gtp-in-gtp-filtering" dval:"enable"`
    MessageFiltAction string `json:"message-filt-action" dval:"drop"`
    MessageFilteringPolicyName string `json:"message-filtering-policy-name"`
    MsisdnFiltAction string `json:"msisdn-filt-action" dval:"drop"`
    MsisdnFiltering string `json:"msisdn-filtering"`
    Name string `json:"name"`
    RatTypeFiltering []TemplateGtpFilteringPolicyRatTypeFiltering `json:"rat-type-filtering"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"filtering-policy"`
}


type TemplateGtpFilteringPolicyRatTypeFiltering struct {
    RatTypeValue string `json:"rat-type-value"`
    RatTypeFiltAction string `json:"rat-type-filt-action" dval:"drop"`
}

func (p *TemplateGtpFilteringPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpFilteringPolicy) getPath() string{
    return "template/gtp/filtering-policy"
}

func (p *TemplateGtpFilteringPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpFilteringPolicy::Post")
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

func (p *TemplateGtpFilteringPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpFilteringPolicy::Get")
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
func (p *TemplateGtpFilteringPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpFilteringPolicy::Put")
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

func (p *TemplateGtpFilteringPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpFilteringPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
