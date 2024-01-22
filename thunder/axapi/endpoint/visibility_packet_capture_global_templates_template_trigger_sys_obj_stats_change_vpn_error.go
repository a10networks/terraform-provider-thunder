

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2103 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2104 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"vpn-error"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsInc2103 struct {
    Bad_opcode int `json:"bad_opcode"`
    Bad_sg_write_len int `json:"bad_sg_write_len"`
    Bad_len int `json:"bad_len"`
    Bad_ipsec_protocol int `json:"bad_ipsec_protocol"`
    Bad_ipsec_auth int `json:"bad_ipsec_auth"`
    Bad_ipsec_padding int `json:"bad_ipsec_padding"`
    Bad_ip_version int `json:"bad_ip_version"`
    Bad_auth_type int `json:"bad_auth_type"`
    Bad_encrypt_type int `json:"bad_encrypt_type"`
    Bad_ipsec_spi int `json:"bad_ipsec_spi"`
    Bad_checksum int `json:"bad_checksum"`
    Bad_ipsec_context int `json:"bad_ipsec_context"`
    Bad_ipsec_context_direction int `json:"bad_ipsec_context_direction"`
    Bad_ipsec_context_flag_mismatch int `json:"bad_ipsec_context_flag_mismatch"`
    Ipcomp_payload int `json:"ipcomp_payload"`
    Bad_selector_match int `json:"bad_selector_match"`
    Bad_fragment_size int `json:"bad_fragment_size"`
    Bad_inline_data int `json:"bad_inline_data"`
    Bad_frag_size_configuration int `json:"bad_frag_size_configuration"`
    Dummy_payload int `json:"dummy_payload"`
    Bad_ip_payload_type int `json:"bad_ip_payload_type"`
    Bad_min_frag_size_auth_sha384_512 int `json:"bad_min_frag_size_auth_sha384_512"`
    Bad_esp_next_header int `json:"bad_esp_next_header"`
    Bad_gre_header int `json:"bad_gre_header"`
    Bad_gre_protocol int `json:"bad_gre_protocol"`
    Ipv6_extension_headers_too_big int `json:"ipv6_extension_headers_too_big"`
    Ipv6_hop_by_hop_error int `json:"ipv6_hop_by_hop_error"`
    Error_ipv6_decrypt_rh_segs_left_error int `json:"error_ipv6_decrypt_rh_segs_left_error"`
    Ipv6_rh_length_error int `json:"ipv6_rh_length_error"`
    Ipv6_outbound_rh_copy_addr_error int `json:"ipv6_outbound_rh_copy_addr_error"`
    Error_ipv6_extension_header_bad int `json:"error_IPv6_extension_header_bad"`
    Bad_encrypt_type_ctr_gcm int `json:"bad_encrypt_type_ctr_gcm"`
    Ah_not_supported_with_gcm_gmac_sha2 int `json:"ah_not_supported_with_gcm_gmac_sha2"`
    Tfc_padding_with_prefrag_not_supported int `json:"tfc_padding_with_prefrag_not_supported"`
    Bad_srtp_auth_tag int `json:"bad_srtp_auth_tag"`
    Bad_ipcomp_configuration int `json:"bad_ipcomp_configuration"`
    Dsiv_incorrect_param int `json:"dsiv_incorrect_param"`
    Bad_ipsec_unknown int `json:"bad_ipsec_unknown"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnErrorTriggerStatsRate2104 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Bad_opcode int `json:"bad_opcode"`
    Bad_sg_write_len int `json:"bad_sg_write_len"`
    Bad_len int `json:"bad_len"`
    Bad_ipsec_protocol int `json:"bad_ipsec_protocol"`
    Bad_ipsec_auth int `json:"bad_ipsec_auth"`
    Bad_ipsec_padding int `json:"bad_ipsec_padding"`
    Bad_ip_version int `json:"bad_ip_version"`
    Bad_auth_type int `json:"bad_auth_type"`
    Bad_encrypt_type int `json:"bad_encrypt_type"`
    Bad_ipsec_spi int `json:"bad_ipsec_spi"`
    Bad_checksum int `json:"bad_checksum"`
    Bad_ipsec_context int `json:"bad_ipsec_context"`
    Bad_ipsec_context_direction int `json:"bad_ipsec_context_direction"`
    Bad_ipsec_context_flag_mismatch int `json:"bad_ipsec_context_flag_mismatch"`
    Ipcomp_payload int `json:"ipcomp_payload"`
    Bad_selector_match int `json:"bad_selector_match"`
    Bad_fragment_size int `json:"bad_fragment_size"`
    Bad_inline_data int `json:"bad_inline_data"`
    Bad_frag_size_configuration int `json:"bad_frag_size_configuration"`
    Dummy_payload int `json:"dummy_payload"`
    Bad_ip_payload_type int `json:"bad_ip_payload_type"`
    Bad_min_frag_size_auth_sha384_512 int `json:"bad_min_frag_size_auth_sha384_512"`
    Bad_esp_next_header int `json:"bad_esp_next_header"`
    Bad_gre_header int `json:"bad_gre_header"`
    Bad_gre_protocol int `json:"bad_gre_protocol"`
    Ipv6_extension_headers_too_big int `json:"ipv6_extension_headers_too_big"`
    Ipv6_hop_by_hop_error int `json:"ipv6_hop_by_hop_error"`
    Error_ipv6_decrypt_rh_segs_left_error int `json:"error_ipv6_decrypt_rh_segs_left_error"`
    Ipv6_rh_length_error int `json:"ipv6_rh_length_error"`
    Ipv6_outbound_rh_copy_addr_error int `json:"ipv6_outbound_rh_copy_addr_error"`
    Error_ipv6_extension_header_bad int `json:"error_IPv6_extension_header_bad"`
    Bad_encrypt_type_ctr_gcm int `json:"bad_encrypt_type_ctr_gcm"`
    Ah_not_supported_with_gcm_gmac_sha2 int `json:"ah_not_supported_with_gcm_gmac_sha2"`
    Tfc_padding_with_prefrag_not_supported int `json:"tfc_padding_with_prefrag_not_supported"`
    Bad_srtp_auth_tag int `json:"bad_srtp_auth_tag"`
    Bad_ipcomp_configuration int `json:"bad_ipcomp_configuration"`
    Dsiv_incorrect_param int `json:"dsiv_incorrect_param"`
    Bad_ipsec_unknown int `json:"bad_ipsec_unknown"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/vpn-error"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeVpnError::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
