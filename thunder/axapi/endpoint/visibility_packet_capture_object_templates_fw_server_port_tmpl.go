

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesFwServerPortTmplTriggerStatsInc2681 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesFwServerPortTmplTriggerStatsRate2682 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesFwServerPortTmplTriggerStatsSeverity2683 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"fw-server-port-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplTriggerStatsInc2681 struct {
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplTriggerStatsRate2682 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Es_resp_400 int `json:"es_resp_400"`
    Es_resp_500 int `json:"es_resp_500"`
    Es_resp_invalid_http int `json:"es_resp_invalid_http"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesFwServerPortTmplTriggerStatsSeverity2683 struct {
    Error int `json:"error"`
    ErrorAlert int `json:"error-alert"`
    ErrorWarning int `json:"error-warning"`
    ErrorCritical int `json:"error-critical"`
    Drop int `json:"drop"`
    DropAlert int `json:"drop-alert"`
    DropWarning int `json:"drop-warning"`
    DropCritical int `json:"drop-critical"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/fw-server-port-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesFwServerPortTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
