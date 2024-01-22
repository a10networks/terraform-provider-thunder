

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpMessageFilteringPolicyVersionV0 struct {
	Inst struct {

    CreateAaPdp string `json:"create-aa-pdp" dval:"enable"`
    CreatePdp string `json:"create-pdp" dval:"enable"`
    DeleteAaPdp string `json:"delete-aa-pdp" dval:"enable"`
    DeletePdp string `json:"delete-pdp" dval:"enable"`
    EnableDisableAction string `json:"enable-disable-action"`
    GtpPdu string `json:"gtp-pdu" dval:"enable"`
    MessageType int `json:"message-type"`
    PduNotification string `json:"pdu-notification" dval:"enable"`
    ReservedMessages string `json:"reserved-messages" dval:"disable"`
    UpdatePdp string `json:"update-pdp" dval:"enable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"version-v0"`
}

func (p *TemplateGtpMessageFilteringPolicyVersionV0) GetId() string{
    return "1"
}

func (p *TemplateGtpMessageFilteringPolicyVersionV0) getPath() string{
    return "template/gtp/message-filtering-policy/" +p.Inst.Name + "/version-v0"
}

func (p *TemplateGtpMessageFilteringPolicyVersionV0) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV0::Post")
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

func (p *TemplateGtpMessageFilteringPolicyVersionV0) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV0::Get")
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
func (p *TemplateGtpMessageFilteringPolicyVersionV0) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV0::Put")
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

func (p *TemplateGtpMessageFilteringPolicyVersionV0) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV0::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
