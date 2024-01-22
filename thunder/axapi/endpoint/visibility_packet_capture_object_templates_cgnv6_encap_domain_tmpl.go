

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl struct {
	Inst struct {

    CaptureConfig string `json:"capture-config"`
    Name string `json:"name"`
    TriggerStatsInc VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc2669 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsRate2670 `json:"trigger-stats-rate"`
    TriggerStatsSeverity VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsSeverity2671 `json:"trigger-stats-severity"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"cgnv6-encap-domain-tmpl"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsInc2669 struct {
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsRate2670 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Interface_not_configured int `json:"interface_not_configured"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmplTriggerStatsSeverity2671 struct {
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

func (p *VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl) getPath() string{
    return "visibility/packet-capture/object-templates/cgnv6-encap-domain-tmpl"
}

func (p *VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesCgnv6EncapDomainTmpl::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
