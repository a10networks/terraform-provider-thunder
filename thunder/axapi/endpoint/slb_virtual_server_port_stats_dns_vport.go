

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerPortStats49 struct {
	Inst struct {

    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbVirtualServerPortStats49Stats `json:"stats"`
    Name string 

	} `json:"stats"`
}


type SlbVirtualServerPortStats49Stats struct {
    Dns_vport SlbVirtualServerPortStats49StatsDns_vport `json:"dns_vport"`
}


type SlbVirtualServerPortStats49StatsDns_vport struct {
    Dns_total_request int `json:"dns_total_request"`
    Dns_total_response int `json:"dns_total_response"`
    Dns_total_drop int `json:"dns_total_drop"`
    Dns_request_response int `json:"dns_request_response"`
    Dns_request_send int `json:"dns_request_send"`
    Dns_request_drop int `json:"dns_request_drop"`
    Dns_response_drop int `json:"dns_response_drop"`
    Dns_response_send int `json:"dns_response_send"`
    Dns_request_timeout int `json:"dns_request_timeout"`
    Dns_request_rexmit int `json:"dns_request_rexmit"`
    Dns_cache_hit int `json:"dns_cache_hit"`
    Dnsrrl_total_dropped int `json:"dnsrrl_total_dropped"`
    Total_mf_dns_pkt int `json:"total_mf_dns_pkt"`
    Total_filter_drop int `json:"total_filter_drop"`
    Total_max_query_len_drop int `json:"total_max_query_len_drop"`
    Rcode_formerr_receive int `json:"rcode_formerr_receive"`
    Rcode_serverr_receive int `json:"rcode_serverr_receive"`
    Rcode_nxdomain_receive int `json:"rcode_nxdomain_receive"`
    Rcode_notimpl_receive int `json:"rcode_notimpl_receive"`
    Rcode_refuse_receive int `json:"rcode_refuse_receive"`
    Rcode_yxdomain_receive int `json:"rcode_yxdomain_receive"`
    Rcode_yxrrset_receive int `json:"rcode_yxrrset_receive"`
    Rcode_nxrrset_receive int `json:"rcode_nxrrset_receive"`
    Rcode_notauth_receive int `json:"rcode_notauth_receive"`
    Rcode_dsotypen_receive int `json:"rcode_dsotypen_receive"`
    Rcode_badver_receive int `json:"rcode_badver_receive"`
    Rcode_badkey_receive int `json:"rcode_badkey_receive"`
    Rcode_badtime_receive int `json:"rcode_badtime_receive"`
    Rcode_badmode_receive int `json:"rcode_badmode_receive"`
    Rcode_badname_receive int `json:"rcode_badname_receive"`
    Rcode_badalg_receive int `json:"rcode_badalg_receive"`
    Rcode_badtranc_receive int `json:"rcode_badtranc_receive"`
    Rcode_badcookie_receive int `json:"rcode_badcookie_receive"`
    Rcode_other_receive int `json:"rcode_other_receive"`
    Rcode_noerror_generate int `json:"rcode_noerror_generate"`
    Rcode_formerr_response int `json:"rcode_formerr_response"`
    Rcode_serverr_response int `json:"rcode_serverr_response"`
    Rcode_nxdomain_response int `json:"rcode_nxdomain_response"`
    Rcode_notimpl_response int `json:"rcode_notimpl_response"`
    Rcode_refuse_response int `json:"rcode_refuse_response"`
    Rcode_yxdomain_response int `json:"rcode_yxdomain_response"`
    Rcode_yxrrset_response int `json:"rcode_yxrrset_response"`
    Rcode_nxrrset_response int `json:"rcode_nxrrset_response"`
    Rcode_notauth_response int `json:"rcode_notauth_response"`
    Rcode_dsotypen_response int `json:"rcode_dsotypen_response"`
    Rcode_badver_response int `json:"rcode_badver_response"`
    Rcode_badkey_response int `json:"rcode_badkey_response"`
    Rcode_badtime_response int `json:"rcode_badtime_response"`
    Rcode_badmode_response int `json:"rcode_badmode_response"`
    Rcode_badname_response int `json:"rcode_badname_response"`
    Rcode_badalg_response int `json:"rcode_badalg_response"`
    Rcode_badtranc_response int `json:"rcode_badtranc_response"`
    Rcode_badcookie_response int `json:"rcode_badcookie_response"`
    Rcode_other_response int `json:"rcode_other_response"`
    Gslb_drop int `json:"gslb_drop"`
    Gslb_query_drop int `json:"gslb_query_drop"`
    Gslb_query_bad int `json:"gslb_query_bad"`
    Gslb_response_drop int `json:"gslb_response_drop"`
    Gslb_response_bad int `json:"gslb_response_bad"`
    Gslb_query_fwd int `json:"gslb_query_fwd"`
    Gslb_response_rvs int `json:"gslb_response_rvs"`
    Gslb_response_good int `json:"gslb_response_good"`
    Type_a_query int `json:"type_A_query"`
    Type_aaaa_query int `json:"type_AAAA_query"`
    Type_cname_query int `json:"type_CNAME_query"`
    Type_mx_query int `json:"type_MX_query"`
    Type_ns_query int `json:"type_NS_query"`
    Type_srv_query int `json:"type_SRV_query"`
    Type_ptr_query int `json:"type_PTR_query"`
    Type_soa_query int `json:"type_SOA_query"`
    Type_txt_query int `json:"type_TXT_query"`
    Type_any_query int `json:"type_ANY_query"`
    Type_other_query int `json:"type_other_query"`
    Type_nsid_query int `json:"type_NSID_query"`
    Type_dau_query int `json:"type_DAU_query"`
    Type_n3u_query int `json:"type_N3U_query"`
    Type_expire_query int `json:"type_EXPIRE_query"`
    Type_cookie_query int `json:"type_COOKIE_query"`
    Type_keepalive_query int `json:"type_KEEPALIVE_query"`
    Type_padding_query int `json:"type_PADDING_query"`
    Type_chain_query int `json:"type_CHAIN_query"`
    Total_dns_filter_type_drop int `json:"total_dns_filter_type_drop"`
    Total_dns_filter_class_drop int `json:"total_dns_filter_class_drop"`
    Dns_filter_type_a_drop int `json:"dns_filter_type_a_drop"`
    Dns_filter_type_aaaa_drop int `json:"dns_filter_type_aaaa_drop"`
    Dns_filter_type_cname_drop int `json:"dns_filter_type_cname_drop"`
    Dns_filter_type_mx_drop int `json:"dns_filter_type_mx_drop"`
    Dns_filter_type_ns_drop int `json:"dns_filter_type_ns_drop"`
    Dns_filter_type_srv_drop int `json:"dns_filter_type_srv_drop"`
    Dns_filter_type_ptr_drop int `json:"dns_filter_type_ptr_drop"`
    Dns_filter_type_soa_drop int `json:"dns_filter_type_soa_drop"`
    Dns_filter_type_txt_drop int `json:"dns_filter_type_txt_drop"`
    Dns_filter_type_any_drop int `json:"dns_filter_type_any_drop"`
    Dns_filter_type_others_drop int `json:"dns_filter_type_others_drop"`
    Dns_filter_class_internet_drop int `json:"dns_filter_class_internet_drop"`
    Dns_filter_class_chaos_drop int `json:"dns_filter_class_chaos_drop"`
    Dns_filter_class_hesiod_drop int `json:"dns_filter_class_hesiod_drop"`
    Dns_filter_class_none_drop int `json:"dns_filter_class_none_drop"`
    Dns_filter_class_any_drop int `json:"dns_filter_class_any_drop"`
    Dns_filter_class_others_drop int `json:"dns_filter_class_others_drop"`
    Dns_recursive_resolution_started int `json:"dns_recursive_resolution_started"`
    Dns_recursive_resolution_succeeded int `json:"dns_recursive_resolution_succeeded"`
    Dns_recursive_resolution_send_failed int `json:"dns_recursive_resolution_send_failed"`
    Dns_recursive_resolution_timeout int `json:"dns_recursive_resolution_timeout"`
    Dns_recursive_resolution_retransmit_sent int `json:"dns_recursive_resolution_retransmit_sent"`
    Dns_recursive_resolution_retransmit_exceeded int `json:"dns_recursive_resolution_retransmit_exceeded"`
    Dns_recursive_resolution_buff_alloc_failed int `json:"dns_recursive_resolution_buff_alloc_failed"`
    Dns_recursive_resolution_ongoing_client_retransmit int `json:"dns_recursive_resolution_ongoing_client_retransmit"`
    Slb_dns_client_ssl_succ int `json:"slb_dns_client_ssl_succ"`
    Slb_dns_server_ssl_succ int `json:"slb_dns_server_ssl_succ"`
    Slb_dns_udp_conn int `json:"slb_dns_udp_conn"`
    Slb_dns_udp_conn_succ int `json:"slb_dns_udp_conn_succ"`
    Slb_dns_padding_to_server_removed int `json:"slb_dns_padding_to_server_removed"`
    Slb_dns_padding_to_client_added int `json:"slb_dns_padding_to_client_added"`
    Slb_dns_edns_subnet_to_server_removed int `json:"slb_dns_edns_subnet_to_server_removed"`
    Slb_dns_udp_retransmit int `json:"slb_dns_udp_retransmit"`
    Slb_dns_udp_retransmit_fail int `json:"slb_dns_udp_retransmit_fail"`
    Dns_rpz_action_drop int `json:"dns_rpz_action_drop"`
    Dns_rpz_action_pass_thru int `json:"dns_rpz_action_pass_thru"`
    Dns_rpz_action_tcp_only int `json:"dns_rpz_action_tcp_only"`
    Dns_rpz_action_nxdomain int `json:"dns_rpz_action_nxdomain"`
    Dns_rpz_action_nodata int `json:"dns_rpz_action_nodata"`
    Dns_rpz_action_local_data int `json:"dns_rpz_action_local_data"`
    Dns_rpz_trigger_client_ip int `json:"dns_rpz_trigger_client_ip"`
    Dns_rpz_trigger_resp_ip int `json:"dns_rpz_trigger_resp_ip"`
    Dns_rpz_trigger_ns_ip int `json:"dns_rpz_trigger_ns_ip"`
    Dns_rpz_trigger_qname int `json:"dns_rpz_trigger_qname"`
    Dns_rpz_trigger_ns_name int `json:"dns_rpz_trigger_ns_name"`
    Dnsrrl_total_allowed int `json:"dnsrrl_total_allowed"`
    Dnsrrl_total_slipped int `json:"dnsrrl_total_slipped"`
    Dnsrrl_bad_fqdn int `json:"dnsrrl_bad_fqdn"`
    Total_mf_dns_pkt_detect int `json:"total_mf_dns_pkt_detect"`
    Type_rrsig_query int `json:"type_RRSIG_query"`
    Type_tsig_query int `json:"type_TSIG_query"`
    Type_dnskey_query int `json:"type_DNSKEY_query"`
    Type_axfr_query int `json:"type_AXFR_query"`
    Type_ixfr_query int `json:"type_IXFR_query"`
    Type_caa_query int `json:"type_CAA_query"`
    Type_naptr_query int `json:"type_NAPTR_query"`
    Type_ds_query int `json:"type_DS_query"`
    Type_cert_query int `json:"type_CERT_query"`
    Dns_recursive_resolution_not_dplane int `json:"dns_recursive_resolution_not_dplane"`
    Dns_recursive_resolution_no_resolver int `json:"dns_recursive_resolution_no_resolver"`
    Dns_recursive_resolution_max_trials_exceeded int `json:"dns_recursive_resolution_max_trials_exceeded"`
    Dns_recursive_resolution_no_hints int `json:"dns_recursive_resolution_no_hints"`
    Dns_recursive_resolution_res_submit_err int `json:"dns_recursive_resolution_res_submit_err"`
    Dns_recursive_resolution_res_check_err int `json:"dns_recursive_resolution_res_check_err"`
    Dns_recursive_resolution_udp_conn_err int `json:"dns_recursive_resolution_udp_conn_err"`
    Dns_recursive_resolution_tcp_conn_err int `json:"dns_recursive_resolution_tcp_conn_err"`
    Dns_recursive_resolution_udp_send_failed int `json:"dns_recursive_resolution_udp_send_failed"`
    Dns_recursive_resolution_icmp_err int `json:"dns_recursive_resolution_icmp_err"`
    Dns_recursive_resolution_query_not_sent int `json:"dns_recursive_resolution_query_not_sent"`
    Dns_tcp_pipeline_request_drop int `json:"dns_tcp_pipeline_request_drop"`
    Dns_recursive_resolution_resp_truncated int `json:"dns_recursive_resolution_resp_truncated"`
    Dns_recursive_resolution_full_response int `json:"dns_recursive_resolution_full_response"`
    Dns_full_response_from_cache int `json:"dns_full_response_from_cache"`
    Dns_recursive_resolution_missing_glue int `json:"dns_recursive_resolution_missing_glue"`
    Dns_recursive_resolution_ns_cache_hit int `json:"dns_recursive_resolution_ns_cache_hit"`
    Dns_recursive_resolution_ns_cache_miss int `json:"dns_recursive_resolution_ns_cache_miss"`
    Dns_recursive_resolution_lookup_ip_proto_switch_46 int `json:"dns_recursive_resolution_lookup_ip_proto_switch_46"`
    Dns_recursive_resolution_lookup_ip_proto_switch_64 int `json:"dns_recursive_resolution_lookup_ip_proto_switch_64"`
    Slb_dns_edns_ecs_received int `json:"slb_dns_edns_ecs_received"`
    Slb_dns_edns_ecs_inserted int `json:"slb_dns_edns_ecs_inserted"`
    Slb_dns_edns_ecs_insertion_fail int `json:"slb_dns_edns_ecs_insertion_fail"`
    Dns_recursive_resolution_invalid_hints int `json:"dns_recursive_resolution_invalid_hints"`
    Dns_recursive_resolution_pending_resolution int `json:"dns_recursive_resolution_pending_resolution"`
    Dns_recursive_resolution_query_dropped int `json:"dns_recursive_resolution_query_dropped"`
    Dns_recursive_resolution_respond_with_servfail int `json:"dns_recursive_resolution_respond_with_servfail"`
    Dns_recursive_resolution_total_trials_1 int `json:"dns_recursive_resolution_total_trials_1"`
    Dns_recursive_resolution_total_trials_3 int `json:"dns_recursive_resolution_total_trials_3"`
    Dns_recursive_resolution_total_trials_6 int `json:"dns_recursive_resolution_total_trials_6"`
    Dns_recursive_resolution_total_trials_12 int `json:"dns_recursive_resolution_total_trials_12"`
    Dns_recursive_resolution_total_trials_24 int `json:"dns_recursive_resolution_total_trials_24"`
    Dns_recursive_resolution_total_trials_48 int `json:"dns_recursive_resolution_total_trials_48"`
    Dns_recursive_resolution_total_trials_64 int `json:"dns_recursive_resolution_total_trials_64"`
    Dns_recursive_resolution_total_trials_128 int `json:"dns_recursive_resolution_total_trials_128"`
    Dns_recursive_resolution_total_trials_max int `json:"dns_recursive_resolution_total_trials_max"`
    Type_https_query int `json:"type_HTTPS_query"`
    Empty_response int `json:"empty_response"`
    Dnsrrl_total_tc int `json:"dnsrrl_total_tc"`
    Dns_negative_served int `json:"dns_negative_served"`
    Dns_recursive_resolution_reach_max_depth int `json:"dns_recursive_resolution_reach_max_depth"`
    Dns_rr_dnssec_req_received int `json:"dns_rr_dnssec_req_received"`
    Dns_rr_dnssec_resp_served int `json:"dns_rr_dnssec_resp_served"`
    Dns_rr_dnssec_validation_failed int `json:"dns_rr_dnssec_validation_failed"`
    Dns_rr_dnssec_val_alg_not_supported int `json:"dns_rr_dnssec_val_alg_not_supported"`
    Dns_rr_dnssec_val_dgst_not_supported int `json:"dns_rr_dnssec_val_dgst_not_supported"`
    Dns_rr_dnssec_val_rrsig_signer_err int `json:"dns_rr_dnssec_val_rrsig_signer_err"`
    Dns_rr_dnssec_val_rrsig_labels_err int `json:"dns_rr_dnssec_val_rrsig_labels_err"`
    Dns_rr_dnssec_val_rrsig_non_validity_period int `json:"dns_rr_dnssec_val_rrsig_non_validity_period"`
    Dns_rr_dnssec_val_dnskey_proto_err int `json:"dns_rr_dnssec_val_dnskey_proto_err"`
    Dns_rr_dnssec_val_incorrect_sig int `json:"dns_rr_dnssec_val_incorrect_sig"`
    Dns_rr_dnssec_val_incorrect_key_dgst int `json:"dns_rr_dnssec_val_incorrect_key_dgst"`
    Dns_rr_dnssec_val_with_trust_anchor_failed int `json:"dns_rr_dnssec_val_with_trust_anchor_failed"`
    Dns_rr_dnssec_val_rrset_size_exceed_limit int `json:"dns_rr_dnssec_val_rrset_size_exceed_limit"`
}

func (p *SlbVirtualServerPortStats49) GetId() string{
    return "1"
}

func (p *SlbVirtualServerPortStats49) getPath() string{
    return "slb/virtual-server/" +p.Inst.Name + "/port/" +strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol+"/stats?dns_vport=true"
}

func (p *SlbVirtualServerPortStats49) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats49::Post")
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

func (p *SlbVirtualServerPortStats49) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats49::Get")
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
func (p *SlbVirtualServerPortStats49) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats49::Put")
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

func (p *SlbVirtualServerPortStats49) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbVirtualServerPortStats49::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
