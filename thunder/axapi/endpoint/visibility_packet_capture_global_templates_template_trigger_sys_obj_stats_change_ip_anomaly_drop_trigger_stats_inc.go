

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc struct {
	Inst struct {

    Bad_ip_flg int `json:"bad_ip_flg"`
    Bad_ip_frg_offset int `json:"bad_ip_frg_offset"`
    Bad_ip_hdrlen int `json:"bad_ip_hdrlen"`
    Bad_ip_payload_len int `json:"bad_ip_payload_len"`
    Bad_ip_ttl int `json:"bad_ip_ttl"`
    Bad_tcp_urg_offset int `json:"bad_tcp_urg_offset"`
    Csum int `json:"csum"`
    Emp_frg int `json:"emp_frg"`
    Emp_mic_frg int `json:"emp_mic_frg"`
    Frg int `json:"frg"`
    Gre_pptp_err int `json:"gre_pptp_err"`
    Ipip_tnl_err int `json:"ipip_tnl_err"`
    Ipip_tnl_msmtch int `json:"ipip_tnl_msmtch"`
    Ipv6_eh_ah int `json:"ipv6_eh_ah"`
    Ipv6_eh_dest int `json:"ipv6_eh_dest"`
    Ipv6_eh_esp int `json:"ipv6_eh_esp"`
    Ipv6_eh_frag int `json:"ipv6_eh_frag"`
    Ipv6_eh_hbh int `json:"ipv6_eh_hbh"`
    Ipv6_eh_malformed int `json:"ipv6_eh_malformed"`
    Ipv6_eh_mobility int `json:"ipv6_eh_mobility"`
    Ipv6_eh_none int `json:"ipv6_eh_none"`
    Ipv6_eh_other int `json:"ipv6_eh_other"`
    Ipv6_eh_routing int `json:"ipv6_eh_routing"`
    Land int `json:"land"`
    No_ip_payload int `json:"no_ip_payload"`
    Nvgre_err int `json:"nvgre_err"`
    Opt int `json:"opt"`
    Over_ip_payload int `json:"over_ip_payload"`
    Pod int `json:"pod"`
    Runt_ip_hdr int `json:"runt_ip_hdr"`
    Runt_tcp_udp_hdr int `json:"runt_tcp_udp_hdr"`
    Tcp_bad_csum int `json:"tcp_bad_csum"`
    Tcp_bad_iplen int `json:"tcp_bad_iplen"`
    Tcp_frg_hdr int `json:"tcp_frg_hdr"`
    Tcp_null_frg int `json:"tcp_null_frg"`
    Tcp_null_scan int `json:"tcp_null_scan"`
    Tcp_opt_err int `json:"tcp_opt_err"`
    Tcp_sht_hdr int `json:"tcp_sht_hdr"`
    Tcp_syn_fin int `json:"tcp_syn_fin"`
    Tcp_syn_frg int `json:"tcp_syn_frg"`
    Tcp_xmas int `json:"tcp_xmas"`
    Tcp_xmas_scan int `json:"tcp_xmas_scan"`
    Udp_bad_csum int `json:"udp_bad_csum"`
    Udp_bad_len int `json:"udp_bad_len"`
    Udp_kerb_frg int `json:"udp_kerb_frg"`
    Udp_port_lb int `json:"udp_port_lb"`
    Udp_srt_hdr int `json:"udp_srt_hdr"`
    Uuid string `json:"uuid"`
    Vxlan_err int `json:"vxlan_err"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/ip-anomaly-drop/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeIpAnomalyDropTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
