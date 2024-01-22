

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplTriggerStatsInc2648 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplTriggerStatsRate2649 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplTriggerStatsSeverity2650 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"aam-auth-server-ocsp-inst-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplTriggerStatsInc2648 struct {
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplTriggerStatsRate2649 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Timeout int `json:"timeout"`
    Fail int `json:"fail"`
    StaplingTimeout int `json:"stapling-timeout"`
    StaplingFail int `json:"stapling-fail"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmplTriggerStatsSeverity2650 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/aam-auth-server-ocsp-inst-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamAuthServerOcspInstTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
