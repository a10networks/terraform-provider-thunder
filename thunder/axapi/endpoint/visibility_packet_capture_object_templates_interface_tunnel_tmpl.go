

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplTriggerStatsInc2696 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplTriggerStatsRate2697 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplTriggerStatsSeverity2698 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"interface-tunnel-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplTriggerStatsInc2696 struct {
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplTriggerStatsRate2697 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    NumRxErrPkts int `json:"num-rx-err-pkts"`
    NumTxErrPkts int `json:"num-tx-err-pkts"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmplTriggerStatsSeverity2698 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/interface-tunnel-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesInterfaceTunnelTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
