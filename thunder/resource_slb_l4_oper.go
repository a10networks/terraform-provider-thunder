package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbL4Oper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_l4_oper`: Operational Status for the object l4\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbL4OperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_cpu_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_outnoroute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst_nosyn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst_broker": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst_ack_attack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst_aflex": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst_stale_sess": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syn_stale_sess": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_outrst_tcpproxy": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_synreceived": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_synreceived_hw": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syn_rate": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syncookiessent": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syncookiessent_ts": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syncookiessentfailed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"intcp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inudp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"svr_sel_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_no_fwd_route": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_no_rev_route": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_icmp_error_process": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_icmp_no_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"smart_nat_id_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syncookiescheckfailed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"novport_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_vport_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nosyn_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nosyn_drop_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nosyn_drop_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nosyn_drop_ack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connlimit_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"connlimit_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_rate_limit_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_rate_limit_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxy_nosock_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aflex_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"sess_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_sess_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_sess_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"other_sess_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_no_slb": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"udp_no_slb": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"throttle_syn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inband_hm_retry": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"inband_hm_reassign": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"auto_reassign": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fast_aging_set": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fast_aging_reset": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_invalid_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"out_seq_ack_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"anomaly_out_seq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"anomaly_zero_win": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"anomaly_bad_content": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"anomaly_pbslb_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"no_resource_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reset_unknown_conn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reset_l7_on_failover": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syn_otherflags": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_syn_withdata": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ignore_msl": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l2_dsr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l3_dsr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port_preserve_attempt": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"port_preserve_succ": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bw_rate_limit_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"bw_watermark_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l4_cps_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nat_cps_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l7_cps_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_cps_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_tpt_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ssl_watermark_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"conn_limit_exceed_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"l4_svr_handshake_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"stateless_conn_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"ax_rexmit_syn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_ack_on_syn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rst_on_syn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_noest_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"noest_client_syn_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"noest_server_syn_xmit_aged_out": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rexmit_syn": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rexmit_syn_delq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rexmit_synack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rexmit_synack_delq": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_fwd_last_ack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rev_last_ack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_fwd_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_fwd_fin_dup": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_fwd_finack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rev_fin": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rev_fin_dup": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rev_finack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_fwd_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rev_rst": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_reqs_no_rsp": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_req_rsps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_req_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_req_morethan_rsps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_rsps_morethan_reqs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_udp_reqs": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"rcv_udp_rsps": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_est": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"synattack": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"skip_insert_client_ip": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dns_id_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"dnssec_switch": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syncookies_buff_queue": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"syncookies_buff_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"server_down_del": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"tcp_connections_closed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"gtp_c_invalid_port": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"gtp_c_invalid_header": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"gtp_c_invalid_message": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"reselect_svrselfail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_port_overload_fail": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_force_preserve_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"snat_force_preserve_free": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxy_header_insert": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxy_header_rexmit": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxy_prot_error": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"proxy_prot_drop": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_pkt_rcv_rr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_smp_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_smp_no_match": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_c_process_local_rr": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_smp_creation_failed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_smp_created": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_smp_free_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_smp_freed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"slb_gtp_proxy_retx_requests": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"pbslb_entry_limit_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fast_path_reroute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"fast_path_l2_reroute": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"cpu_count": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSlbL4OperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL4OperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL4Oper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbL4OperOper := setObjectSlbL4OperOper(res)
		d.Set("oper", SlbL4OperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbL4OperOper(ret edpt.DataSlbL4Oper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"l4_cpu_list": setSliceSlbL4OperOperL4CpuList(ret.DtSlbL4Oper.Oper.L4CpuList),
			"cpu_count":   ret.DtSlbL4Oper.Oper.CpuCount,
		},
	}
}

func setSliceSlbL4OperOperL4CpuList(d []edpt.SlbL4OperOperL4CpuList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ip_outnoroute"] = item.Ip_outnoroute
		in["tcp_outrst"] = item.Tcp_outrst
		in["tcp_outrst_nosyn"] = item.Tcp_outrst_nosyn
		in["tcp_outrst_broker"] = item.Tcp_outrst_broker
		in["tcp_outrst_ack_attack"] = item.Tcp_outrst_ack_attack
		in["tcp_outrst_aflex"] = item.Tcp_outrst_aflex
		in["tcp_outrst_stale_sess"] = item.Tcp_outrst_stale_sess
		in["tcp_syn_stale_sess"] = item.Tcp_syn_stale_sess
		in["tcp_outrst_tcpproxy"] = item.Tcp_outrst_tcpproxy
		in["tcp_synreceived"] = item.Tcp_synreceived
		in["tcp_synreceived_hw"] = item.Tcp_synreceived_hw
		in["tcp_syn_rate"] = item.Tcp_syn_rate
		in["tcp_syncookiessent"] = item.Tcp_syncookiessent
		in["tcp_syncookiessent_ts"] = item.Tcp_syncookiessent_ts
		in["tcp_syncookiessentfailed"] = item.Tcp_syncookiessentfailed
		in["intcp"] = item.Intcp
		in["inudp"] = item.Inudp
		in["svr_sel_failed"] = item.Svr_sel_failed
		in["snat_fail"] = item.Snat_fail
		in["snat_no_fwd_route"] = item.Snat_no_fwd_route
		in["snat_no_rev_route"] = item.Snat_no_rev_route
		in["snat_icmp_error_process"] = item.Snat_icmp_error_process
		in["snat_icmp_no_match"] = item.Snat_icmp_no_match
		in["smart_nat_id_mismatch"] = item.Smart_nat_id_mismatch
		in["tcp_syncookiescheckfailed"] = item.Tcp_syncookiescheckfailed
		in["novport_drop"] = item.Novport_drop
		in["no_vport_drop"] = item.No_vport_drop
		in["nosyn_drop"] = item.Nosyn_drop
		in["nosyn_drop_fin"] = item.Nosyn_drop_fin
		in["nosyn_drop_rst"] = item.Nosyn_drop_rst
		in["nosyn_drop_ack"] = item.Nosyn_drop_ack
		in["connlimit_drop"] = item.Connlimit_drop
		in["connlimit_reset"] = item.Connlimit_reset
		in["conn_rate_limit_drop"] = item.Conn_rate_limit_drop
		in["conn_rate_limit_reset"] = item.Conn_rate_limit_reset
		in["proxy_nosock_drop"] = item.Proxy_nosock_drop
		in["aflex_drop"] = item.Aflex_drop
		in["sess_aged_out"] = item.Sess_aged_out
		in["tcp_sess_aged_out"] = item.Tcp_sess_aged_out
		in["udp_sess_aged_out"] = item.Udp_sess_aged_out
		in["other_sess_aged_out"] = item.Other_sess_aged_out
		in["tcp_no_slb"] = item.Tcp_no_slb
		in["udp_no_slb"] = item.Udp_no_slb
		in["throttle_syn"] = item.Throttle_syn
		in["inband_hm_retry"] = item.Inband_hm_retry
		in["inband_hm_reassign"] = item.Inband_hm_reassign
		in["auto_reassign"] = item.Auto_reassign
		in["fast_aging_set"] = item.Fast_aging_set
		in["fast_aging_reset"] = item.Fast_aging_reset
		in["tcp_invalid_drop"] = item.Tcp_invalid_drop
		in["out_seq_ack_drop"] = item.Out_seq_ack_drop
		in["anomaly_out_seq"] = item.Anomaly_out_seq
		in["anomaly_zero_win"] = item.Anomaly_zero_win
		in["anomaly_bad_content"] = item.Anomaly_bad_content
		in["anomaly_pbslb_drop"] = item.Anomaly_pbslb_drop
		in["no_resource_drop"] = item.No_resource_drop
		in["reset_unknown_conn"] = item.Reset_unknown_conn
		in["reset_l7_on_failover"] = item.Reset_l7_on_failover
		in["tcp_syn_otherflags"] = item.Tcp_syn_otherflags
		in["tcp_syn_withdata"] = item.Tcp_syn_withdata
		in["ignore_msl"] = item.Ignore_msl
		in["l2_dsr"] = item.L2_dsr
		in["l3_dsr"] = item.L3_dsr
		in["port_preserve_attempt"] = item.Port_preserve_attempt
		in["port_preserve_succ"] = item.Port_preserve_succ
		in["bw_rate_limit_exceed_drop"] = item.Bw_rate_limit_exceed_drop
		in["bw_watermark_drop"] = item.Bw_watermark_drop
		in["l4_cps_exceed_drop"] = item.L4_cps_exceed_drop
		in["nat_cps_exceed_drop"] = item.Nat_cps_exceed_drop
		in["l7_cps_exceed_drop"] = item.L7_cps_exceed_drop
		in["ssl_cps_exceed_drop"] = item.Ssl_cps_exceed_drop
		in["ssl_tpt_exceed_drop"] = item.Ssl_tpt_exceed_drop
		in["ssl_watermark_drop"] = item.Ssl_watermark_drop
		in["conn_limit_exceed_drop"] = item.Conn_limit_exceed_drop
		in["l4_svr_handshake_fail"] = item.L4_svr_handshake_fail
		in["stateless_conn_timeout"] = item.Stateless_conn_timeout
		in["ax_rexmit_syn"] = item.Ax_rexmit_syn
		in["rcv_ack_on_syn"] = item.Rcv_ack_on_syn
		in["rcv_rst_on_syn"] = item.Rcv_rst_on_syn
		in["tcp_noest_aged_out"] = item.Tcp_noest_aged_out
		in["noest_client_syn_aged_out"] = item.Noest_client_syn_aged_out
		in["noest_server_syn_xmit_aged_out"] = item.Noest_server_syn_xmit_aged_out
		in["rcv_rexmit_syn"] = item.Rcv_rexmit_syn
		in["rcv_rexmit_syn_delq"] = item.Rcv_rexmit_syn_delq
		in["rcv_rexmit_synack"] = item.Rcv_rexmit_synack
		in["rcv_rexmit_synack_delq"] = item.Rcv_rexmit_synack_delq
		in["rcv_fwd_last_ack"] = item.Rcv_fwd_last_ack
		in["rcv_rev_last_ack"] = item.Rcv_rev_last_ack
		in["rcv_fwd_fin"] = item.Rcv_fwd_fin
		in["rcv_fwd_fin_dup"] = item.Rcv_fwd_fin_dup
		in["rcv_fwd_finack"] = item.Rcv_fwd_finack
		in["rcv_rev_fin"] = item.Rcv_rev_fin
		in["rcv_rev_fin_dup"] = item.Rcv_rev_fin_dup
		in["rcv_rev_finack"] = item.Rcv_rev_finack
		in["rcv_fwd_rst"] = item.Rcv_fwd_rst
		in["rcv_rev_rst"] = item.Rcv_rev_rst
		in["rcv_reqs_no_rsp"] = item.Rcv_reqs_no_rsp
		in["rcv_req_rsps"] = item.Rcv_req_rsps
		in["rcv_req_not_match"] = item.Rcv_req_not_match
		in["rcv_req_morethan_rsps"] = item.Rcv_req_morethan_rsps
		in["rcv_rsps_morethan_reqs"] = item.Rcv_rsps_morethan_reqs
		in["rcv_udp_reqs"] = item.Rcv_udp_reqs
		in["rcv_udp_rsps"] = item.Rcv_udp_rsps
		in["tcp_est"] = item.Tcp_est
		in["synattack"] = item.Synattack
		in["skip_insert_client_ip"] = item.Skip_insert_client_ip
		in["dns_id_switch"] = item.Dns_id_switch
		in["dnssec_switch"] = item.Dnssec_switch
		in["syncookies_buff_queue"] = item.Syncookies_buff_queue
		in["syncookies_buff_drop"] = item.Syncookies_buff_drop
		in["server_down_del"] = item.Server_down_del
		in["tcp_connections_closed"] = item.Tcp_connections_closed
		in["gtp_c_invalid_port"] = item.Gtp_c_invalid_port
		in["gtp_c_invalid_header"] = item.Gtp_c_invalid_header
		in["gtp_c_invalid_message"] = item.Gtp_c_invalid_message
		in["reselect_svrselfail"] = item.Reselect_svrselfail
		in["snat_port_overload_fail"] = item.Snat_port_overload_fail
		in["snat_force_preserve_alloc"] = item.Snat_force_preserve_alloc
		in["snat_force_preserve_free"] = item.Snat_force_preserve_free
		in["proxy_header_insert"] = item.Proxy_header_insert
		in["proxy_header_rexmit"] = item.Proxy_header_rexmit
		in["proxy_prot_error"] = item.Proxy_prot_error
		in["proxy_prot_drop"] = item.Proxy_prot_drop
		in["slb_gtp_proxy_pkt_rcv_rr"] = item.Slb_gtp_proxy_pkt_rcv_rr
		in["slb_gtp_proxy_smp_match"] = item.Slb_gtp_proxy_smp_match
		in["slb_gtp_proxy_smp_no_match"] = item.Slb_gtp_proxy_smp_no_match
		in["slb_gtp_proxy_c_process_local_rr"] = item.Slb_gtp_proxy_c_process_local_rr
		in["slb_gtp_proxy_smp_creation_failed"] = item.Slb_gtp_proxy_smp_creation_failed
		in["slb_gtp_proxy_smp_created"] = item.Slb_gtp_proxy_smp_created
		in["slb_gtp_proxy_smp_free_not_found"] = item.Slb_gtp_proxy_smp_free_not_found
		in["slb_gtp_proxy_smp_freed"] = item.Slb_gtp_proxy_smp_freed
		in["slb_gtp_proxy_retx_requests"] = item.Slb_gtp_proxy_retx_requests
		in["pbslb_entry_limit_exceed"] = item.Pbslb_entry_limit_exceed
		in["fast_path_reroute"] = item.Fast_path_reroute
		in["fast_path_l2_reroute"] = item.Fast_path_l2_reroute
		result = append(result, in)
	}
	return result
}

func getObjectSlbL4OperOper(d []interface{}) edpt.SlbL4OperOper {

	count1 := len(d)
	var ret edpt.SlbL4OperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4CpuList = getSliceSlbL4OperOperL4CpuList(in["l4_cpu_list"].([]interface{}))
		ret.CpuCount = in["cpu_count"].(int)
	}
	return ret
}

func getSliceSlbL4OperOperL4CpuList(d []interface{}) []edpt.SlbL4OperOperL4CpuList {

	count1 := len(d)
	ret := make([]edpt.SlbL4OperOperL4CpuList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbL4OperOperL4CpuList
		oi.Ip_outnoroute = in["ip_outnoroute"].(int)
		oi.Tcp_outrst = in["tcp_outrst"].(int)
		oi.Tcp_outrst_nosyn = in["tcp_outrst_nosyn"].(int)
		oi.Tcp_outrst_broker = in["tcp_outrst_broker"].(int)
		oi.Tcp_outrst_ack_attack = in["tcp_outrst_ack_attack"].(int)
		oi.Tcp_outrst_aflex = in["tcp_outrst_aflex"].(int)
		oi.Tcp_outrst_stale_sess = in["tcp_outrst_stale_sess"].(int)
		oi.Tcp_syn_stale_sess = in["tcp_syn_stale_sess"].(int)
		oi.Tcp_outrst_tcpproxy = in["tcp_outrst_tcpproxy"].(int)
		oi.Tcp_synreceived = in["tcp_synreceived"].(int)
		oi.Tcp_synreceived_hw = in["tcp_synreceived_hw"].(int)
		oi.Tcp_syn_rate = in["tcp_syn_rate"].(int)
		oi.Tcp_syncookiessent = in["tcp_syncookiessent"].(int)
		oi.Tcp_syncookiessent_ts = in["tcp_syncookiessent_ts"].(int)
		oi.Tcp_syncookiessentfailed = in["tcp_syncookiessentfailed"].(int)
		oi.Intcp = in["intcp"].(int)
		oi.Inudp = in["inudp"].(int)
		oi.Svr_sel_failed = in["svr_sel_failed"].(int)
		oi.Snat_fail = in["snat_fail"].(int)
		oi.Snat_no_fwd_route = in["snat_no_fwd_route"].(int)
		oi.Snat_no_rev_route = in["snat_no_rev_route"].(int)
		oi.Snat_icmp_error_process = in["snat_icmp_error_process"].(int)
		oi.Snat_icmp_no_match = in["snat_icmp_no_match"].(int)
		oi.Smart_nat_id_mismatch = in["smart_nat_id_mismatch"].(int)
		oi.Tcp_syncookiescheckfailed = in["tcp_syncookiescheckfailed"].(int)
		oi.Novport_drop = in["novport_drop"].(int)
		oi.No_vport_drop = in["no_vport_drop"].(int)
		oi.Nosyn_drop = in["nosyn_drop"].(int)
		oi.Nosyn_drop_fin = in["nosyn_drop_fin"].(int)
		oi.Nosyn_drop_rst = in["nosyn_drop_rst"].(int)
		oi.Nosyn_drop_ack = in["nosyn_drop_ack"].(int)
		oi.Connlimit_drop = in["connlimit_drop"].(int)
		oi.Connlimit_reset = in["connlimit_reset"].(int)
		oi.Conn_rate_limit_drop = in["conn_rate_limit_drop"].(int)
		oi.Conn_rate_limit_reset = in["conn_rate_limit_reset"].(int)
		oi.Proxy_nosock_drop = in["proxy_nosock_drop"].(int)
		oi.Aflex_drop = in["aflex_drop"].(int)
		oi.Sess_aged_out = in["sess_aged_out"].(int)
		oi.Tcp_sess_aged_out = in["tcp_sess_aged_out"].(int)
		oi.Udp_sess_aged_out = in["udp_sess_aged_out"].(int)
		oi.Other_sess_aged_out = in["other_sess_aged_out"].(int)
		oi.Tcp_no_slb = in["tcp_no_slb"].(int)
		oi.Udp_no_slb = in["udp_no_slb"].(int)
		oi.Throttle_syn = in["throttle_syn"].(int)
		oi.Inband_hm_retry = in["inband_hm_retry"].(int)
		oi.Inband_hm_reassign = in["inband_hm_reassign"].(int)
		oi.Auto_reassign = in["auto_reassign"].(int)
		oi.Fast_aging_set = in["fast_aging_set"].(int)
		oi.Fast_aging_reset = in["fast_aging_reset"].(int)
		oi.Tcp_invalid_drop = in["tcp_invalid_drop"].(int)
		oi.Out_seq_ack_drop = in["out_seq_ack_drop"].(int)
		oi.Anomaly_out_seq = in["anomaly_out_seq"].(int)
		oi.Anomaly_zero_win = in["anomaly_zero_win"].(int)
		oi.Anomaly_bad_content = in["anomaly_bad_content"].(int)
		oi.Anomaly_pbslb_drop = in["anomaly_pbslb_drop"].(int)
		oi.No_resource_drop = in["no_resource_drop"].(int)
		oi.Reset_unknown_conn = in["reset_unknown_conn"].(int)
		oi.Reset_l7_on_failover = in["reset_l7_on_failover"].(int)
		oi.Tcp_syn_otherflags = in["tcp_syn_otherflags"].(int)
		oi.Tcp_syn_withdata = in["tcp_syn_withdata"].(int)
		oi.Ignore_msl = in["ignore_msl"].(int)
		oi.L2_dsr = in["l2_dsr"].(int)
		oi.L3_dsr = in["l3_dsr"].(int)
		oi.Port_preserve_attempt = in["port_preserve_attempt"].(int)
		oi.Port_preserve_succ = in["port_preserve_succ"].(int)
		oi.Bw_rate_limit_exceed_drop = in["bw_rate_limit_exceed_drop"].(int)
		oi.Bw_watermark_drop = in["bw_watermark_drop"].(int)
		oi.L4_cps_exceed_drop = in["l4_cps_exceed_drop"].(int)
		oi.Nat_cps_exceed_drop = in["nat_cps_exceed_drop"].(int)
		oi.L7_cps_exceed_drop = in["l7_cps_exceed_drop"].(int)
		oi.Ssl_cps_exceed_drop = in["ssl_cps_exceed_drop"].(int)
		oi.Ssl_tpt_exceed_drop = in["ssl_tpt_exceed_drop"].(int)
		oi.Ssl_watermark_drop = in["ssl_watermark_drop"].(int)
		oi.Conn_limit_exceed_drop = in["conn_limit_exceed_drop"].(int)
		oi.L4_svr_handshake_fail = in["l4_svr_handshake_fail"].(int)
		oi.Stateless_conn_timeout = in["stateless_conn_timeout"].(int)
		oi.Ax_rexmit_syn = in["ax_rexmit_syn"].(int)
		oi.Rcv_ack_on_syn = in["rcv_ack_on_syn"].(int)
		oi.Rcv_rst_on_syn = in["rcv_rst_on_syn"].(int)
		oi.Tcp_noest_aged_out = in["tcp_noest_aged_out"].(int)
		oi.Noest_client_syn_aged_out = in["noest_client_syn_aged_out"].(int)
		oi.Noest_server_syn_xmit_aged_out = in["noest_server_syn_xmit_aged_out"].(int)
		oi.Rcv_rexmit_syn = in["rcv_rexmit_syn"].(int)
		oi.Rcv_rexmit_syn_delq = in["rcv_rexmit_syn_delq"].(int)
		oi.Rcv_rexmit_synack = in["rcv_rexmit_synack"].(int)
		oi.Rcv_rexmit_synack_delq = in["rcv_rexmit_synack_delq"].(int)
		oi.Rcv_fwd_last_ack = in["rcv_fwd_last_ack"].(int)
		oi.Rcv_rev_last_ack = in["rcv_rev_last_ack"].(int)
		oi.Rcv_fwd_fin = in["rcv_fwd_fin"].(int)
		oi.Rcv_fwd_fin_dup = in["rcv_fwd_fin_dup"].(int)
		oi.Rcv_fwd_finack = in["rcv_fwd_finack"].(int)
		oi.Rcv_rev_fin = in["rcv_rev_fin"].(int)
		oi.Rcv_rev_fin_dup = in["rcv_rev_fin_dup"].(int)
		oi.Rcv_rev_finack = in["rcv_rev_finack"].(int)
		oi.Rcv_fwd_rst = in["rcv_fwd_rst"].(int)
		oi.Rcv_rev_rst = in["rcv_rev_rst"].(int)
		oi.Rcv_reqs_no_rsp = in["rcv_reqs_no_rsp"].(int)
		oi.Rcv_req_rsps = in["rcv_req_rsps"].(int)
		oi.Rcv_req_not_match = in["rcv_req_not_match"].(int)
		oi.Rcv_req_morethan_rsps = in["rcv_req_morethan_rsps"].(int)
		oi.Rcv_rsps_morethan_reqs = in["rcv_rsps_morethan_reqs"].(int)
		oi.Rcv_udp_reqs = in["rcv_udp_reqs"].(int)
		oi.Rcv_udp_rsps = in["rcv_udp_rsps"].(int)
		oi.Tcp_est = in["tcp_est"].(int)
		oi.Synattack = in["synattack"].(int)
		oi.Skip_insert_client_ip = in["skip_insert_client_ip"].(int)
		oi.Dns_id_switch = in["dns_id_switch"].(int)
		oi.Dnssec_switch = in["dnssec_switch"].(int)
		oi.Syncookies_buff_queue = in["syncookies_buff_queue"].(int)
		oi.Syncookies_buff_drop = in["syncookies_buff_drop"].(int)
		oi.Server_down_del = in["server_down_del"].(int)
		oi.Tcp_connections_closed = in["tcp_connections_closed"].(int)
		oi.Gtp_c_invalid_port = in["gtp_c_invalid_port"].(int)
		oi.Gtp_c_invalid_header = in["gtp_c_invalid_header"].(int)
		oi.Gtp_c_invalid_message = in["gtp_c_invalid_message"].(int)
		oi.Reselect_svrselfail = in["reselect_svrselfail"].(int)
		oi.Snat_port_overload_fail = in["snat_port_overload_fail"].(int)
		oi.Snat_force_preserve_alloc = in["snat_force_preserve_alloc"].(int)
		oi.Snat_force_preserve_free = in["snat_force_preserve_free"].(int)
		oi.Proxy_header_insert = in["proxy_header_insert"].(int)
		oi.Proxy_header_rexmit = in["proxy_header_rexmit"].(int)
		oi.Proxy_prot_error = in["proxy_prot_error"].(int)
		oi.Proxy_prot_drop = in["proxy_prot_drop"].(int)
		oi.Slb_gtp_proxy_pkt_rcv_rr = in["slb_gtp_proxy_pkt_rcv_rr"].(int)
		oi.Slb_gtp_proxy_smp_match = in["slb_gtp_proxy_smp_match"].(int)
		oi.Slb_gtp_proxy_smp_no_match = in["slb_gtp_proxy_smp_no_match"].(int)
		oi.Slb_gtp_proxy_c_process_local_rr = in["slb_gtp_proxy_c_process_local_rr"].(int)
		oi.Slb_gtp_proxy_smp_creation_failed = in["slb_gtp_proxy_smp_creation_failed"].(int)
		oi.Slb_gtp_proxy_smp_created = in["slb_gtp_proxy_smp_created"].(int)
		oi.Slb_gtp_proxy_smp_free_not_found = in["slb_gtp_proxy_smp_free_not_found"].(int)
		oi.Slb_gtp_proxy_smp_freed = in["slb_gtp_proxy_smp_freed"].(int)
		oi.Slb_gtp_proxy_retx_requests = in["slb_gtp_proxy_retx_requests"].(int)
		oi.Pbslb_entry_limit_exceed = in["pbslb_entry_limit_exceed"].(int)
		oi.Fast_path_reroute = in["fast_path_reroute"].(int)
		oi.Fast_path_l2_reroute = in["fast_path_l2_reroute"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbL4Oper(d *schema.ResourceData) edpt.SlbL4Oper {
	var ret edpt.SlbL4Oper

	ret.Oper = getObjectSlbL4OperOper(d.Get("oper").([]interface{}))
	return ret
}
