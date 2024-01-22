

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateSctp struct {
	Inst struct {

    ChecksumCheck string `json:"checksum-check"`
    Log TemplateSctpLog `json:"log"`
    Name string `json:"name"`
    PermitPayloadProtocol TemplateSctpPermitPayloadProtocol `json:"permit-payload-protocol"`
    SctpHalfOpenIdleTimeout int `json:"sctp-half-open-idle-timeout" dval:"4"`
    SctpIdleTimeout int `json:"sctp-idle-timeout" dval:"5"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"sctp"`
}


type TemplateSctpLog struct {
    PayloadProtoFiltering int `json:"payload-proto-filtering"`
}


type TemplateSctpPermitPayloadProtocol struct {
    PermitConfigId []TemplateSctpPermitPayloadProtocolPermitConfigId `json:"permit-config-id"`
    PermitConfigName []TemplateSctpPermitPayloadProtocolPermitConfigName `json:"permit-config-name"`
}


type TemplateSctpPermitPayloadProtocolPermitConfigId struct {
    ProtocolId int `json:"protocol-id"`
}


type TemplateSctpPermitPayloadProtocolPermitConfigName struct {
    ProtocolName string `json:"protocol-name"`
}

func (p *TemplateSctp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateSctp) getPath() string{
    return "template/sctp"
}

func (p *TemplateSctp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateSctp::Post")
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

func (p *TemplateSctp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateSctp::Get")
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
func (p *TemplateSctp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateSctp::Put")
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

func (p *TemplateSctp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateSctp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
