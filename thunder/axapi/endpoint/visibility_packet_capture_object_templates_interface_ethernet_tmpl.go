

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsInc2693 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsRate2694 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsSeverity2695 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"interface-ethernet-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsInc2693 struct {
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsRate2694 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Input_errors int `json:"input_errors"`
    Crc int `json:"crc"`
    Runts int `json:"runts"`
    Giants int `json:"giants"`
    Output_errors int `json:"output_errors"`
    Collisions int `json:"collisions"`
    Giants_output int `json:"giants_output"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmplTriggerStatsSeverity2695 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/interface-ethernet-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceEthernetTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
