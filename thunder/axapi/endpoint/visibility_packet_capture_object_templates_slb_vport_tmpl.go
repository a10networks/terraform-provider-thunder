

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesSlbVportTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsInc2714 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate2715 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsSeverity2716 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"slb-vport-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsInc2714 struct {
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate2715 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Loc_deny int `json:"loc_deny"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsSeverity2716 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/slb-vport-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
