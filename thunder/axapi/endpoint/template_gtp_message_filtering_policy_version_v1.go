

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpMessageFilteringPolicyVersionV1 struct {
	Inst struct {

    CreateMbms string `json:"create-mbms" dval:"disable"`
    CreatePdp string `json:"create-pdp" dval:"enable"`
    DeleteMbms string `json:"delete-mbms" dval:"disable"`
    DeletePdp string `json:"delete-pdp" dval:"enable"`
    EnableDisableAction string `json:"enable-disable-action"`
    GtpPdu string `json:"gtp-pdu" dval:"enable"`
    InitiatePdp string `json:"initiate-pdp" dval:"enable"`
    MbmsDeregistration string `json:"mbms-deregistration" dval:"disable"`
    MbmsNotification string `json:"mbms-notification" dval:"disable"`
    MbmsRegistration string `json:"mbms-registration" dval:"disable"`
    MbmsSession string `json:"mbms-session" dval:"disable"`
    MessageType int `json:"message-type"`
    MsInfoChange string `json:"ms-info-change" dval:"enable"`
    PduNotification string `json:"pdu-notification" dval:"enable"`
    ReservedMessages string `json:"reserved-messages" dval:"disable"`
    UpdateMbms string `json:"update-mbms" dval:"disable"`
    UpdatePdp string `json:"update-pdp" dval:"enable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"version-v1"`
}

func (p *TemplateGtpMessageFilteringPolicyVersionV1) GetId() string{
    return "1"
}

func (p *TemplateGtpMessageFilteringPolicyVersionV1) getPath() string{
    return "template/gtp/message-filtering-policy/" +p.Inst.Name + "/version-v1"
}

func (p *TemplateGtpMessageFilteringPolicyVersionV1) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV1::Post")
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

func (p *TemplateGtpMessageFilteringPolicyVersionV1) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV1::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *TemplateGtpMessageFilteringPolicyVersionV1) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV1::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *TemplateGtpMessageFilteringPolicyVersionV1) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV1::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
