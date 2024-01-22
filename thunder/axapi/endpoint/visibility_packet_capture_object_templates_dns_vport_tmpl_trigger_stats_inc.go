

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc struct {
	Inst struct {

    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Gslb_query_bad int `json:"gslb_query_bad"`
    Gslb_response_bad int `json:"gslb_response_bad"`
    Rcode_notimpl_receive int `json:"rcode_notimpl_receive"`
    Rcode_notimpl_response int `json:"rcode_notimpl_response"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_filter_drop int `json:"total_filter_drop"`
    Total_max_query_len_drop int `json:"total_max_query_len_drop"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/object-templates/dns_vport-tmpl/" +p.Inst.Name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesDnsVportTmplTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
