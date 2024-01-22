

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbL4Stats struct {
    
    Stats SlbL4StatsStats `json:"stats"`

}
type DataSlbL4Stats struct {
    DtSlbL4Stats SlbL4Stats `json:"l4"`
}


type SlbL4StatsStats struct {
    Intcp int `json:"intcp"`
    Synreceived int `json:"synreceived"`
    Tcp_fwd_last_ack int `json:"tcp_fwd_last_ack"`
    Tcp_rev_last_ack int `json:"tcp_rev_last_ack"`
    Tcp_rev_fin int `json:"tcp_rev_fin"`
    Tcp_fwd_fin int `json:"tcp_fwd_fin"`
    Tcp_fwd_ackfin int `json:"tcp_fwd_ackfin"`
    Inudp int `json:"inudp"`
    Syncookiessent int `json:"syncookiessent"`
    Syncookiessent_ts int `json:"syncookiessent_ts"`
    Syncookiessentfailed int `json:"syncookiessentfailed"`
    Outrst int `json:"outrst"`
    Outrst_nosyn int `json:"outrst_nosyn"`
    Outrst_broker int `json:"outrst_broker"`
    Outrst_ack_attack int `json:"outrst_ack_attack"`
    Outrst_aflex int `json:"outrst_aflex"`
    Outrst_stale_sess int `json:"outrst_stale_sess"`
    Syn_stale_sess int `json:"syn_stale_sess"`
    Outrst_tcpproxy int `json:"outrst_tcpproxy"`
    Svrselfail int `json:"svrselfail"`
    Noroute int `json:"noroute"`
    Snat_fail int `json:"snat_fail"`
    Snat_no_fwd_route int `json:"snat_no_fwd_route"`
    Snat_no_rev_route int `json:"snat_no_rev_route"`
    Snat_icmp_error_process int `json:"snat_icmp_error_process"`
    Snat_icmp_no_match int `json:"snat_icmp_no_match"`
    Smart_nat_id_mismatch int `json:"smart_nat_id_mismatch"`
    Syncookiescheckfailed int `json:"syncookiescheckfailed"`
    Novport_drop int `json:"novport_drop"`
    No_vport_drop int `json:"no_vport_drop"`
    Nosyn_drop int `json:"nosyn_drop"`
    Nosyn_drop_fin int `json:"nosyn_drop_fin"`
    Nosyn_drop_rst int `json:"nosyn_drop_rst"`
    Nosyn_drop_ack int `json:"nosyn_drop_ack"`
    Connlimit_drop int `json:"connlimit_drop"`
    Connlimit_reset int `json:"connlimit_reset"`
    Conn_rate_limit_drop int `json:"conn_rate_limit_drop"`
    Conn_rate_limit_reset int `json:"conn_rate_limit_reset"`
    Proxy_nosock_drop int `json:"proxy_nosock_drop"`
    Drop_aflex int `json:"drop_aflex"`
    Sess_aged_out int `json:"sess_aged_out"`
    Tcp_sess_aged_out int `json:"tcp_sess_aged_out"`
    Udp_sess_aged_out int `json:"udp_sess_aged_out"`
    Other_sess_aged_out int `json:"other_sess_aged_out"`
    Tcp_no_slb int `json:"tcp_no_slb"`
    Udp_no_slb int `json:"udp_no_slb"`
    Throttle_syn int `json:"throttle_syn"`
    Drop_gslb int `json:"drop_gslb"`
    Inband_hm_retry int `json:"inband_hm_retry"`
    Inband_hm_reassign int `json:"inband_hm_reassign"`
    Auto_reassign int `json:"auto_reassign"`
    Fast_aging_set int `json:"fast_aging_set"`
    Fast_aging_reset int `json:"fast_aging_reset"`
    Dns_policy_drop int `json:"dns_policy_drop"`
    Tcp_invalid_drop int `json:"tcp_invalid_drop"`
    Anomaly_out_seq int `json:"anomaly_out_seq"`
    Anomaly_zero_win int `json:"anomaly_zero_win"`
    Anomaly_bad_content int `json:"anomaly_bad_content"`
    Anomaly_pbslb_drop int `json:"anomaly_pbslb_drop"`
    No_resourse_drop int `json:"no_resourse_drop"`
    Reset_unknown_conn int `json:"reset_unknown_conn"`
    Reset_l7_on_failover int `json:"reset_l7_on_failover"`
    Ignore_msl int `json:"ignore_msl"`
    L2_dsr int `json:"l2_dsr"`
    L3_dsr int `json:"l3_dsr"`
    Port_preserve_attempt int `json:"port_preserve_attempt"`
    Port_preserve_succ int `json:"port_preserve_succ"`
    Tcpsyndata_drop int `json:"tcpsyndata_drop"`
    Tcpotherflags_drop int `json:"tcpotherflags_drop"`
    Bw_rate_limit_exceed int `json:"bw_rate_limit_exceed"`
    Bw_watermark_drop int `json:"bw_watermark_drop"`
    L4_cps_exceed int `json:"l4_cps_exceed"`
    Nat_cps_exceed int `json:"nat_cps_exceed"`
    L7_cps_exceed int `json:"l7_cps_exceed"`
    Ssl_cps_exceed int `json:"ssl_cps_exceed"`
    Ssl_tpt_exceed int `json:"ssl_tpt_exceed"`
    Ssl_watermark_drop int `json:"ssl_watermark_drop"`
    Concurrent_conn_exceed int `json:"concurrent_conn_exceed"`
    Svr_syn_handshake_fail int `json:"svr_syn_handshake_fail"`
    Stateless_conn_timeout int `json:"stateless_conn_timeout"`
    Tcp_ax_rexmit_syn int `json:"tcp_ax_rexmit_syn"`
    Tcp_syn_rcv_ack int `json:"tcp_syn_rcv_ack"`
    Tcp_syn_rcv_rst int `json:"tcp_syn_rcv_rst"`
    Tcp_sess_noest_aged_out int `json:"tcp_sess_noest_aged_out"`
    Tcp_sess_noest_csyn_rcv_aged_out int `json:"tcp_sess_noest_csyn_rcv_aged_out"`
    Tcp_sess_noest_ssyn_xmit_aged_out int `json:"tcp_sess_noest_ssyn_xmit_aged_out"`
    Tcp_rexmit_syn int `json:"tcp_rexmit_syn"`
    Tcp_rexmit_syn_delq int `json:"tcp_rexmit_syn_delq"`
    Tcp_rexmit_synack int `json:"tcp_rexmit_synack"`
    Tcp_rexmit_synack_delq int `json:"tcp_rexmit_synack_delq"`
    Tcp_fwd_fin_dup int `json:"tcp_fwd_fin_dup"`
    Tcp_rev_fin_dup int `json:"tcp_rev_fin_dup"`
    Tcp_rev_ackfin int `json:"tcp_rev_ackfin"`
    Tcp_fwd_rst int `json:"tcp_fwd_rst"`
    Tcp_rev_rst int `json:"tcp_rev_rst"`
    Udp_req_oneplus_no_resp int `json:"udp_req_oneplus_no_resp"`
    Udp_req_one_oneplus_resp int `json:"udp_req_one_oneplus_resp"`
    Udp_req_resp_notmatch int `json:"udp_req_resp_notmatch"`
    Udp_req_more_resp int `json:"udp_req_more_resp"`
    Udp_resp_more_req int `json:"udp_resp_more_req"`
    Udp_req_oneplus int `json:"udp_req_oneplus"`
    Udp_resp_oneplus int `json:"udp_resp_oneplus"`
    Out_seq_ack_drop int `json:"out_seq_ack_drop"`
    Tcp_est int `json:"tcp_est"`
    Synattack int `json:"synattack"`
    Syn_rate int `json:"syn_rate"`
    Syncookie_buff_drop int `json:"syncookie_buff_drop"`
    Syncookie_buff_queue int `json:"syncookie_buff_queue"`
    Skip_insert_client_ip int `json:"skip_insert_client_ip"`
    Synreceived_hw int `json:"synreceived_hw"`
    Dns_id_switch int `json:"dns_id_switch"`
    Server_down_del int `json:"server_down_del"`
    Dnssec_switch int `json:"dnssec_switch"`
    Rate_drop_reset_unkn int `json:"rate_drop_reset_unkn"`
    Tcp_connections_closed int `json:"tcp_connections_closed"`
    Gtp_c_invalid_port int `json:"gtp_c_invalid_port"`
    Gtp_c_invalid_header int `json:"gtp_c_invalid_header"`
    Gtp_c_invalid_message int `json:"gtp_c_invalid_message"`
    Reselect_svrselfail int `json:"reselect_svrselfail"`
    Snat_port_overload_fail int `json:"snat_port_overload_fail"`
    Snat_force_preserve_alloc int `json:"snat_force_preserve_alloc"`
    Snat_force_preserve_free int `json:"snat_force_preserve_free"`
    Proxy_header_insert int `json:"proxy_header_insert"`
    Proxy_header_rexmit int `json:"proxy_header_rexmit"`
    Proxy_prot_error int `json:"proxy_prot_error"`
    Proxy_prot_drop int `json:"proxy_prot_drop"`
    Slb_gtp_proxy_smp_match int `json:"slb_gtp_proxy_smp_match"`
    Slb_gtp_proxy_smp_no_match int `json:"slb_gtp_proxy_smp_no_match"`
    Slb_gtp_proxy_c_process_local_rr int `json:"slb_gtp_proxy_c_process_local_rr"`
    Slb_gtp_proxy_smp_creation_failed int `json:"slb_gtp_proxy_smp_creation_failed"`
    Slb_gtp_proxy_smp_created int `json:"slb_gtp_proxy_smp_created"`
    Slb_gtp_proxy_smp_free_not_found int `json:"slb_gtp_proxy_smp_free_not_found"`
    Slb_gtp_proxy_smp_freed int `json:"slb_gtp_proxy_smp_freed"`
    Slb_gtp_proxy_retx_requests int `json:"slb_gtp_proxy_retx_requests"`
    Pbslb_entry_limit_exceed int `json:"pbslb_entry_limit_exceed"`
    Fast_path_reroute int `json:"fast_path_reroute"`
    Fast_path_l2_reroute int `json:"fast_path_l2_reroute"`
}

func (p *SlbL4Stats) GetId() string{
    return "1"
}

func (p *SlbL4Stats) getPath() string{
    return "slb/l4/stats"
}

func (p *SlbL4Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbL4Stats,error) {
logger.Println("SlbL4Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbL4Stats
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
