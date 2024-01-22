

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbL4Oper struct {
    
    Oper SlbL4OperOper `json:"oper"`

}
type DataSlbL4Oper struct {
    DtSlbL4Oper SlbL4Oper `json:"l4"`
}


type SlbL4OperOper struct {
    L4CpuList []SlbL4OperOperL4CpuList `json:"l4-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbL4OperOperL4CpuList struct {
    Ip_outnoroute int `json:"ip_outnoroute"`
    Tcp_outrst int `json:"tcp_outrst"`
    Tcp_outrst_nosyn int `json:"tcp_outrst_nosyn"`
    Tcp_outrst_broker int `json:"tcp_outrst_broker"`
    Tcp_outrst_ack_attack int `json:"tcp_outrst_ack_attack"`
    Tcp_outrst_aflex int `json:"tcp_outrst_aflex"`
    Tcp_outrst_stale_sess int `json:"tcp_outrst_stale_sess"`
    Tcp_syn_stale_sess int `json:"tcp_syn_stale_sess"`
    Tcp_outrst_tcpproxy int `json:"tcp_outrst_tcpproxy"`
    Tcp_synreceived int `json:"tcp_synreceived"`
    Tcp_synreceived_hw int `json:"tcp_synreceived_hw"`
    Tcp_syn_rate int `json:"tcp_syn_rate"`
    Tcp_syncookiessent int `json:"tcp_syncookiessent"`
    Tcp_syncookiessent_ts int `json:"tcp_syncookiessent_ts"`
    Tcp_syncookiessentfailed int `json:"tcp_syncookiessentfailed"`
    Intcp int `json:"intcp"`
    Inudp int `json:"inudp"`
    Svr_sel_failed int `json:"svr_sel_failed"`
    Snat_fail int `json:"snat_fail"`
    Snat_no_fwd_route int `json:"snat_no_fwd_route"`
    Snat_no_rev_route int `json:"snat_no_rev_route"`
    Snat_icmp_error_process int `json:"snat_icmp_error_process"`
    Snat_icmp_no_match int `json:"snat_icmp_no_match"`
    Smart_nat_id_mismatch int `json:"smart_nat_id_mismatch"`
    Tcp_syncookiescheckfailed int `json:"tcp_syncookiescheckfailed"`
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
    Aflex_drop int `json:"aflex_drop"`
    Sess_aged_out int `json:"sess_aged_out"`
    Tcp_sess_aged_out int `json:"tcp_sess_aged_out"`
    Udp_sess_aged_out int `json:"udp_sess_aged_out"`
    Other_sess_aged_out int `json:"other_sess_aged_out"`
    Tcp_no_slb int `json:"tcp_no_slb"`
    Udp_no_slb int `json:"udp_no_slb"`
    Throttle_syn int `json:"throttle_syn"`
    Inband_hm_retry int `json:"inband_hm_retry"`
    Inband_hm_reassign int `json:"inband_hm_reassign"`
    Auto_reassign int `json:"auto_reassign"`
    Fast_aging_set int `json:"fast_aging_set"`
    Fast_aging_reset int `json:"fast_aging_reset"`
    Tcp_invalid_drop int `json:"tcp_invalid_drop"`
    Out_seq_ack_drop int `json:"out_seq_ack_drop"`
    Anomaly_out_seq int `json:"anomaly_out_seq"`
    Anomaly_zero_win int `json:"anomaly_zero_win"`
    Anomaly_bad_content int `json:"anomaly_bad_content"`
    Anomaly_pbslb_drop int `json:"anomaly_pbslb_drop"`
    No_resource_drop int `json:"no_resource_drop"`
    Reset_unknown_conn int `json:"reset_unknown_conn"`
    Reset_l7_on_failover int `json:"reset_l7_on_failover"`
    Tcp_syn_otherflags int `json:"tcp_syn_otherflags"`
    Tcp_syn_withdata int `json:"tcp_syn_withdata"`
    Ignore_msl int `json:"ignore_msl"`
    L2_dsr int `json:"l2_dsr"`
    L3_dsr int `json:"l3_dsr"`
    Port_preserve_attempt int `json:"port_preserve_attempt"`
    Port_preserve_succ int `json:"port_preserve_succ"`
    Bw_rate_limit_exceed_drop int `json:"bw_rate_limit_exceed_drop"`
    Bw_watermark_drop int `json:"bw_watermark_drop"`
    L4_cps_exceed_drop int `json:"l4_cps_exceed_drop"`
    Nat_cps_exceed_drop int `json:"nat_cps_exceed_drop"`
    L7_cps_exceed_drop int `json:"l7_cps_exceed_drop"`
    Ssl_cps_exceed_drop int `json:"ssl_cps_exceed_drop"`
    Ssl_tpt_exceed_drop int `json:"ssl_tpt_exceed_drop"`
    Ssl_watermark_drop int `json:"ssl_watermark_drop"`
    Conn_limit_exceed_drop int `json:"conn_limit_exceed_drop"`
    L4_svr_handshake_fail int `json:"l4_svr_handshake_fail"`
    Stateless_conn_timeout int `json:"stateless_conn_timeout"`
    Ax_rexmit_syn int `json:"ax_rexmit_syn"`
    Rcv_ack_on_syn int `json:"rcv_ack_on_syn"`
    Rcv_rst_on_syn int `json:"rcv_rst_on_syn"`
    Tcp_noest_aged_out int `json:"tcp_noest_aged_out"`
    Noest_client_syn_aged_out int `json:"noest_client_syn_aged_out"`
    Noest_server_syn_xmit_aged_out int `json:"noest_server_syn_xmit_aged_out"`
    Rcv_rexmit_syn int `json:"rcv_rexmit_syn"`
    Rcv_rexmit_syn_delq int `json:"rcv_rexmit_syn_delq"`
    Rcv_rexmit_synack int `json:"rcv_rexmit_synack"`
    Rcv_rexmit_synack_delq int `json:"rcv_rexmit_synack_delq"`
    Rcv_fwd_last_ack int `json:"rcv_fwd_last_ack"`
    Rcv_rev_last_ack int `json:"rcv_rev_last_ack"`
    Rcv_fwd_fin int `json:"rcv_fwd_fin"`
    Rcv_fwd_fin_dup int `json:"rcv_fwd_fin_dup"`
    Rcv_fwd_finack int `json:"rcv_fwd_finack"`
    Rcv_rev_fin int `json:"rcv_rev_fin"`
    Rcv_rev_fin_dup int `json:"rcv_rev_fin_dup"`
    Rcv_rev_finack int `json:"rcv_rev_finack"`
    Rcv_fwd_rst int `json:"rcv_fwd_rst"`
    Rcv_rev_rst int `json:"rcv_rev_rst"`
    Rcv_reqs_no_rsp int `json:"rcv_reqs_no_rsp"`
    Rcv_req_rsps int `json:"rcv_req_rsps"`
    Rcv_req_not_match int `json:"rcv_req_not_match"`
    Rcv_req_morethan_rsps int `json:"rcv_req_morethan_rsps"`
    Rcv_rsps_morethan_reqs int `json:"rcv_rsps_morethan_reqs"`
    Rcv_udp_reqs int `json:"rcv_udp_reqs"`
    Rcv_udp_rsps int `json:"rcv_udp_rsps"`
    Tcp_est int `json:"tcp_est"`
    Synattack int `json:"synattack"`
    Skip_insert_client_ip int `json:"skip_insert_client_ip"`
    Dns_id_switch int `json:"dns_id_switch"`
    Dnssec_switch int `json:"dnssec_switch"`
    Syncookies_buff_queue int `json:"syncookies_buff_queue"`
    Syncookies_buff_drop int `json:"syncookies_buff_drop"`
    Server_down_del int `json:"server_down_del"`
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
    Slb_gtp_proxy_pkt_rcv_rr int `json:"slb_gtp_proxy_pkt_rcv_rr"`
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

func (p *SlbL4Oper) GetId() string{
    return "1"
}

func (p *SlbL4Oper) getPath() string{
    return "slb/l4/oper"
}

func (p *SlbL4Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbL4Oper,error) {
logger.Println("SlbL4Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbL4Oper
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
