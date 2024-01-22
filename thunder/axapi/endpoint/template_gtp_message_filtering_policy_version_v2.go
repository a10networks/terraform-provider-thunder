

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpMessageFilteringPolicyVersionV2 struct {
	Inst struct {

    BearerResource string `json:"bearer-resource" dval:"enable"`
    ChangeNotification string `json:"change-notification" dval:"enable"`
    CreateBearer string `json:"create-bearer" dval:"enable"`
    CreateSession string `json:"create-session" dval:"enable"`
    DeleteBearer string `json:"delete-bearer" dval:"enable"`
    DeleteCommand string `json:"delete-command" dval:"enable"`
    DeletePdn string `json:"delete-pdn" dval:"enable"`
    DeleteSession string `json:"delete-session" dval:"enable"`
    EnableDisableAction string `json:"enable-disable-action"`
    MessageType int `json:"message-type"`
    ModifyBearer string `json:"modify-bearer" dval:"enable"`
    ModifyCommand string `json:"modify-command" dval:"enable"`
    PgwDownlinkTrigger string `json:"pgw-downlink-trigger" dval:"disable"`
    RemoteUeReport string `json:"remote-ue-report" dval:"enable"`
    ReservedMessages string `json:"reserved-messages" dval:"disable"`
    Resume string `json:"resume" dval:"enable"`
    Suspend string `json:"suspend" dval:"enable"`
    TraceSession string `json:"trace-session" dval:"disable"`
    UpdateBearer string `json:"update-bearer" dval:"enable"`
    UpdatePdn string `json:"update-pdn" dval:"enable"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"version-v2"`
}

func (p *TemplateGtpMessageFilteringPolicyVersionV2) GetId() string{
    return "1"
}

func (p *TemplateGtpMessageFilteringPolicyVersionV2) getPath() string{
    return "template/gtp/message-filtering-policy/" +p.Inst.Name + "/version-v2"
}

func (p *TemplateGtpMessageFilteringPolicyVersionV2) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV2::Post")
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

func (p *TemplateGtpMessageFilteringPolicyVersionV2) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV2::Get")
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
func (p *TemplateGtpMessageFilteringPolicyVersionV2) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV2::Put")
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

func (p *TemplateGtpMessageFilteringPolicyVersionV2) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicyVersionV2::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
