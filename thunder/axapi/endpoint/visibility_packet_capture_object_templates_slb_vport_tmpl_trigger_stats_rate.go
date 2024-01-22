

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate struct {
	Inst struct {

    Compression_miss_no_client int `json:"compression_miss_no_client"`
    Compression_miss_template_exclusion int `json:"compression_miss_template_exclusion"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Duration int `json:"duration" dval:"60"`
    Es_total_failure_actions int `json:"es_total_failure_actions"`
    Loc_deny int `json:"loc_deny"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Total_mf_dns_pkts int `json:"total_mf_dns_pkts"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/object-templates/slb-vport-tmpl/" +p.Inst.Name + "/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesSlbVportTmplTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
