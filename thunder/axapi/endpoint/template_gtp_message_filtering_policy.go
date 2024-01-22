

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpMessageFilteringPolicy struct {
	Inst struct {

    InterfaceType string `json:"interface-type"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VersionV0 TemplateGtpMessageFilteringPolicyVersionV01900 `json:"version-v0"`
    VersionV1 TemplateGtpMessageFilteringPolicyVersionV11901 `json:"version-v1"`
    VersionV2 TemplateGtpMessageFilteringPolicyVersionV21902 `json:"version-v2"`

	} `json:"message-filtering-policy"`
}


type TemplateGtpMessageFilteringPolicyVersionV01900 struct {
    EnableDisableAction string `json:"enable-disable-action" dval:"enable"`
    MessageType int `json:"message-type"`
    CreatePdp string `json:"create-pdp" dval:"enable"`
    UpdatePdp string `json:"update-pdp" dval:"enable"`
    DeletePdp string `json:"delete-pdp" dval:"enable"`
    PduNotification string `json:"pdu-notification" dval:"enable"`
    GtpPdu string `json:"gtp-pdu" dval:"enable"`
    CreateAaPdp string `json:"create-aa-pdp" dval:"enable"`
    DeleteAaPdp string `json:"delete-aa-pdp" dval:"enable"`
    ReservedMessages string `json:"reserved-messages" dval:"disable"`
    Uuid string `json:"uuid"`
}


type TemplateGtpMessageFilteringPolicyVersionV11901 struct {
    EnableDisableAction string `json:"enable-disable-action" dval:"enable"`
    MessageType int `json:"message-type"`
    CreatePdp string `json:"create-pdp" dval:"enable"`
    UpdatePdp string `json:"update-pdp" dval:"enable"`
    DeletePdp string `json:"delete-pdp" dval:"enable"`
    InitiatePdp string `json:"initiate-pdp" dval:"enable"`
    PduNotification string `json:"pdu-notification" dval:"enable"`
    MsInfoChange string `json:"ms-info-change" dval:"enable"`
    GtpPdu string `json:"gtp-pdu" dval:"enable"`
    MbmsSession string `json:"mbms-session" dval:"disable"`
    MbmsNotification string `json:"mbms-notification" dval:"disable"`
    MbmsRegistration string `json:"mbms-registration" dval:"disable"`
    MbmsDeregistration string `json:"mbms-deregistration" dval:"disable"`
    CreateMbms string `json:"create-mbms" dval:"disable"`
    DeleteMbms string `json:"delete-mbms" dval:"disable"`
    UpdateMbms string `json:"update-mbms" dval:"disable"`
    ReservedMessages string `json:"reserved-messages" dval:"disable"`
    Uuid string `json:"uuid"`
}


type TemplateGtpMessageFilteringPolicyVersionV21902 struct {
    EnableDisableAction string `json:"enable-disable-action" dval:"enable"`
    MessageType int `json:"message-type"`
    ChangeNotification string `json:"change-notification" dval:"enable"`
    CreateSession string `json:"create-session" dval:"enable"`
    DeleteSession string `json:"delete-session" dval:"enable"`
    ModifyBearer string `json:"modify-bearer" dval:"enable"`
    RemoteUeReport string `json:"remote-ue-report" dval:"enable"`
    ModifyCommand string `json:"modify-command" dval:"enable"`
    DeleteCommand string `json:"delete-command" dval:"enable"`
    BearerResource string `json:"bearer-resource" dval:"enable"`
    CreateBearer string `json:"create-bearer" dval:"enable"`
    UpdateBearer string `json:"update-bearer" dval:"enable"`
    DeleteBearer string `json:"delete-bearer" dval:"enable"`
    DeletePdn string `json:"delete-pdn" dval:"enable"`
    UpdatePdn string `json:"update-pdn" dval:"enable"`
    Suspend string `json:"suspend" dval:"enable"`
    Resume string `json:"resume" dval:"enable"`
    PgwDownlinkTrigger string `json:"pgw-downlink-trigger" dval:"disable"`
    TraceSession string `json:"trace-session" dval:"disable"`
    ReservedMessages string `json:"reserved-messages" dval:"disable"`
    Uuid string `json:"uuid"`
}

func (p *TemplateGtpMessageFilteringPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpMessageFilteringPolicy) getPath() string{
    return "template/gtp/message-filtering-policy"
}

func (p *TemplateGtpMessageFilteringPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicy::Post")
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

func (p *TemplateGtpMessageFilteringPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicy::Get")
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
func (p *TemplateGtpMessageFilteringPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicy::Put")
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

func (p *TemplateGtpMessageFilteringPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpMessageFilteringPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
