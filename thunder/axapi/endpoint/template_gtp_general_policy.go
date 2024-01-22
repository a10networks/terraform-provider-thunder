

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpGeneralPolicy struct {
	Inst struct {

    HandoverTimeout int `json:"handover-timeout" dval:"2"`
    MaxMesgLengthAction string `json:"max-mesg-length-action" dval:"drop"`
    MaximumMessageLength int `json:"maximum-message-length" dval:"1500"`
    Name string `json:"name"`
    TunnelTimeout int `json:"tunnel-timeout" dval:"1440"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    V0Action string `json:"v0-action" dval:"drop"`

	} `json:"general-policy"`
}

func (p *TemplateGtpGeneralPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpGeneralPolicy) getPath() string{
    return "template/gtp/general-policy"
}

func (p *TemplateGtpGeneralPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpGeneralPolicy::Post")
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

func (p *TemplateGtpGeneralPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpGeneralPolicy::Get")
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
func (p *TemplateGtpGeneralPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpGeneralPolicy::Put")
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

func (p *TemplateGtpGeneralPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpGeneralPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
