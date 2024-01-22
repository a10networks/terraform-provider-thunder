

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity struct {
	Inst struct {

    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropCritical int `json:"drop-critical"`
    DropWarning int `json:"drop-warning"`
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorCritical int `json:"error-critical"`
    ErrorWarning int `json:"error-warning"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-severity"`
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity) getPath() string{
    return "visibility/packet-capture/object-templates/aam-auth-service-group-tmpl/" +p.Inst.Name + "/trigger-stats-severity"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServiceGroupTmplTriggerStatsSeverity::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
