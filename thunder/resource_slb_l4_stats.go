package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbL4Stats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_l4_stats`: Statistics for the object l4\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbL4StatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"intcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP received",
						},
						"synreceived": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN received",
						},
						"tcp_fwd_last_ack": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv fwd last ACK",
						},
						"tcp_rev_last_ack": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rev last ACK",
						},
						"tcp_rev_fin": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rev FIN",
						},
						"tcp_fwd_fin": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv fwd FIN",
						},
						"tcp_fwd_ackfin": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv fwd FIN|ACK",
						},
						"inudp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP received",
						},
						"syncookiessent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN cookie snt",
						},
						"syncookiessent_ts": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN cookie snt ts",
						},
						"syncookiessentfailed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN cookie snt fail",
						},
						"outrst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST",
						},
						"outrst_nosyn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST no SYN",
						},
						"outrst_broker": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST L4 proxy",
						},
						"outrst_ack_attack": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST ACK attack",
						},
						"outrst_aflex": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST aFleX",
						},
						"outrst_stale_sess": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST stale sess",
						},
						"syn_stale_sess": {
							Type: schema.TypeInt, Optional: true, Description: "SYN stale sess drop",
						},
						"outrst_tcpproxy": {
							Type: schema.TypeInt, Optional: true, Description: "TCP out RST TCP proxy",
						},
						"svrselfail": {
							Type: schema.TypeInt, Optional: true, Description: "Server sel failure",
						},
						"noroute": {
							Type: schema.TypeInt, Optional: true, Description: "IP out noroute",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT failure",
						},
						"snat_no_fwd_route": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT no fwd route",
						},
						"snat_no_rev_route": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT no rev route",
						},
						"snat_icmp_error_process": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT ICMP Process",
						},
						"snat_icmp_no_match": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT ICMP No Match",
						},
						"smart_nat_id_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Auto NAT id mismatch",
						},
						"syncookiescheckfailed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN cookie failed",
						},
						"novport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NAT no session drops",
						},
						"no_vport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "vport not matching drops",
						},
						"nosyn_drop": {
							Type: schema.TypeInt, Optional: true, Description: "No SYN pkt drops",
						},
						"nosyn_drop_fin": {
							Type: schema.TypeInt, Optional: true, Description: "No SYN pkt drops - FIN",
						},
						"nosyn_drop_rst": {
							Type: schema.TypeInt, Optional: true, Description: "No SYN pkt drops - RST",
						},
						"nosyn_drop_ack": {
							Type: schema.TypeInt, Optional: true, Description: "No SYN pkt drops - ACK",
						},
						"connlimit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Limit drops",
						},
						"connlimit_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Limit resets",
						},
						"conn_rate_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Conn rate limit drops",
						},
						"conn_rate_limit_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Conn rate limit resets",
						},
						"proxy_nosock_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy no sock drops",
						},
						"drop_aflex": {
							Type: schema.TypeInt, Optional: true, Description: "aFleX drops",
						},
						"sess_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "Session aged out",
						},
						"tcp_sess_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session aged out",
						},
						"udp_sess_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session aged out",
						},
						"other_sess_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "Other Session aged out",
						},
						"tcp_no_slb": {
							Type: schema.TypeInt, Optional: true, Description: "TCP no SLB",
						},
						"udp_no_slb": {
							Type: schema.TypeInt, Optional: true, Description: "UDP no SLB",
						},
						"throttle_syn": {
							Type: schema.TypeInt, Optional: true, Description: "SYN Throttle",
						},
						"drop_gslb": {
							Type: schema.TypeInt, Optional: true, Description: "Drop GSLB",
						},
						"inband_hm_retry": {
							Type: schema.TypeInt, Optional: true, Description: "Inband HM retry",
						},
						"inband_hm_reassign": {
							Type: schema.TypeInt, Optional: true, Description: "Inband HM reassign",
						},
						"auto_reassign": {
							Type: schema.TypeInt, Optional: true, Description: "Auto-reselect server",
						},
						"fast_aging_set": {
							Type: schema.TypeInt, Optional: true, Description: "Fast aging set",
						},
						"fast_aging_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Fast aging reset",
						},
						"dns_policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Policy Drop",
						},
						"tcp_invalid_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP invalid drop",
						},
						"anomaly_out_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Anomaly out of sequence",
						},
						"anomaly_zero_win": {
							Type: schema.TypeInt, Optional: true, Description: "Anomaly zero window",
						},
						"anomaly_bad_content": {
							Type: schema.TypeInt, Optional: true, Description: "Anomaly bad content",
						},
						"anomaly_pbslb_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Anomaly pbslb drop",
						},
						"no_resourse_drop": {
							Type: schema.TypeInt, Optional: true, Description: "No resource drop",
						},
						"reset_unknown_conn": {
							Type: schema.TypeInt, Optional: true, Description: "Reset unknown conn",
						},
						"reset_l7_on_failover": {
							Type: schema.TypeInt, Optional: true, Description: "RST L7 on failover",
						},
						"ignore_msl": {
							Type: schema.TypeInt, Optional: true, Description: "ignore msl",
						},
						"l2_dsr": {
							Type: schema.TypeInt, Optional: true, Description: "L2 DSR received",
						},
						"l3_dsr": {
							Type: schema.TypeInt, Optional: true, Description: "L3 DSR received",
						},
						"port_preserve_attempt": {
							Type: schema.TypeInt, Optional: true, Description: "NAT Port Preserve Try",
						},
						"port_preserve_succ": {
							Type: schema.TypeInt, Optional: true, Description: "NAT Port Preserve Succ",
						},
						"tcpsyndata_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN With Data Drop",
						},
						"tcpotherflags_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Other Flags Drop",
						},
						"bw_rate_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "BW-Limit Exceed drop",
						},
						"bw_watermark_drop": {
							Type: schema.TypeInt, Optional: true, Description: "BW-Watermark drop",
						},
						"l4_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "L4 CPS exceed drop",
						},
						"nat_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT CPS exceed drop",
						},
						"l7_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "L7 CPS exceed drop",
						},
						"ssl_cps_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SSL CPS exceed drop",
						},
						"ssl_tpt_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SSL TPT exceed drop",
						},
						"ssl_watermark_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SSL TPT-Watermark drop",
						},
						"concurrent_conn_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "L3V Conn Limit Drop",
						},
						"svr_syn_handshake_fail": {
							Type: schema.TypeInt, Optional: true, Description: "L4 server handshake fail",
						},
						"stateless_conn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "L4 stateless Conn TO",
						},
						"tcp_ax_rexmit_syn": {
							Type: schema.TypeInt, Optional: true, Description: "L4 AX re-xmit SYN",
						},
						"tcp_syn_rcv_ack": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv ACK on SYN",
						},
						"tcp_syn_rcv_rst": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv RST on SYN",
						},
						"tcp_sess_noest_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "TCP no-Est Sess aged out",
						},
						"tcp_sess_noest_csyn_rcv_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "no-Est CSYN rcv aged out",
						},
						"tcp_sess_noest_ssyn_xmit_aged_out": {
							Type: schema.TypeInt, Optional: true, Description: "no-Est SSYN snt aged out",
						},
						"tcp_rexmit_syn": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rexmit SYN",
						},
						"tcp_rexmit_syn_delq": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rexmit SYN (delq)",
						},
						"tcp_rexmit_synack": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rexmit SYN|ACK",
						},
						"tcp_rexmit_synack_delq": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rexmit SYN|ACK DQ",
						},
						"tcp_fwd_fin_dup": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv fwd FIN dup",
						},
						"tcp_rev_fin_dup": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rev FIN dup",
						},
						"tcp_rev_ackfin": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rev FIN|ACK",
						},
						"tcp_fwd_rst": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv fwd RST",
						},
						"tcp_rev_rst": {
							Type: schema.TypeInt, Optional: true, Description: "L4 rcv rev RST",
						},
						"udp_req_oneplus_no_resp": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP reqs no rsp",
						},
						"udp_req_one_oneplus_resp": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP req rsps",
						},
						"udp_req_resp_notmatch": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP req/rsp not match",
						},
						"udp_req_more_resp": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP req greater than rsps",
						},
						"udp_resp_more_req": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP rsps greater than reqs",
						},
						"udp_req_oneplus": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP reqs",
						},
						"udp_resp_oneplus": {
							Type: schema.TypeInt, Optional: true, Description: "L4 UDP rsps",
						},
						"out_seq_ack_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Out of sequence ACK drop",
						},
						"tcp_est": {
							Type: schema.TypeInt, Optional: true, Description: "L4 TCP Established",
						},
						"synattack": {
							Type: schema.TypeInt, Optional: true, Description: "L4 SYN attack",
						},
						"syn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN rate per sec",
						},
						"syncookie_buff_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN cookie buff drop",
						},
						"syncookie_buff_queue": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN cookie buff queue",
						},
						"skip_insert_client_ip": {
							Type: schema.TypeInt, Optional: true, Description: "Skip Insert-client-ip",
						},
						"synreceived_hw": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN (HW SYN cookie)",
						},
						"dns_id_switch": {
							Type: schema.TypeInt, Optional: true, Description: "DNS query id switch",
						},
						"server_down_del": {
							Type: schema.TypeInt, Optional: true, Description: "Server Down Del switch",
						},
						"dnssec_switch": {
							Type: schema.TypeInt, Optional: true, Description: "DNSSEC SG switch",
						},
						"rate_drop_reset_unkn": {
							Type: schema.TypeInt, Optional: true, Description: "Rate Drop reset",
						},
						"tcp_connections_closed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Connections Closed",
						},
						"gtp_c_invalid_port": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Packet Received on GTP VIP",
						},
						"gtp_c_invalid_header": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Header Received on GTP VIP",
						},
						"gtp_c_invalid_message": {
							Type: schema.TypeInt, Optional: true, Description: "Non Create Session/PDP Context Request/Response Received on GTP VIP",
						},
						"reselect_svrselfail": {
							Type: schema.TypeInt, Optional: true, Description: "Server reselect failure",
						},
						"snat_port_overload_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Snat port overload fail",
						},
						"snat_force_preserve_alloc": {
							Type: schema.TypeInt, Optional: true, Description: "Snat port preserve allocated",
						},
						"snat_force_preserve_free": {
							Type: schema.TypeInt, Optional: true, Description: "Snat port preserve freed",
						},
						"proxy_header_insert": {
							Type: schema.TypeInt, Optional: true, Description: "PROXY protocol header inserted",
						},
						"proxy_header_rexmit": {
							Type: schema.TypeInt, Optional: true, Description: "PROXY protocol header retransmitted",
						},
						"proxy_prot_error": {
							Type: schema.TypeInt, Optional: true, Description: "PROXY protocol error",
						},
						"proxy_prot_drop": {
							Type: schema.TypeInt, Optional: true, Description: "PROXY protocol drop",
						},
						"slb_gtp_proxy_smp_match": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy helper session found",
						},
						"slb_gtp_proxy_smp_no_match": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy helper session not found",
						},
						"slb_gtp_proxy_c_process_local_rr": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy messageprocessed locally on RR",
						},
						"slb_gtp_proxy_smp_creation_failed": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy helper session creation failed",
						},
						"slb_gtp_proxy_smp_created": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy helper session created",
						},
						"slb_gtp_proxy_smp_free_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy session helper not found during cleanup",
						},
						"slb_gtp_proxy_smp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy session helper freed",
						},
						"slb_gtp_proxy_retx_requests": {
							Type: schema.TypeInt, Optional: true, Description: "SLB GTP proxy retx requests",
						},
						"pbslb_entry_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "pbslb entry limit Exceed",
						},
						"fast_path_reroute": {
							Type: schema.TypeInt, Optional: true, Description: "Fast Path Reroute",
						},
						"fast_path_l2_reroute": {
							Type: schema.TypeInt, Optional: true, Description: "Fast Path L2 Reroute",
						},
					},
				},
			},
		},
	}
}

func resourceSlbL4StatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbL4StatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbL4Stats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbL4StatsStats := setObjectSlbL4StatsStats(res)
		d.Set("stats", SlbL4StatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbL4StatsStats(ret edpt.DataSlbL4Stats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"intcp":                             ret.DtSlbL4Stats.Stats.Intcp,
			"synreceived":                       ret.DtSlbL4Stats.Stats.Synreceived,
			"tcp_fwd_last_ack":                  ret.DtSlbL4Stats.Stats.Tcp_fwd_last_ack,
			"tcp_rev_last_ack":                  ret.DtSlbL4Stats.Stats.Tcp_rev_last_ack,
			"tcp_rev_fin":                       ret.DtSlbL4Stats.Stats.Tcp_rev_fin,
			"tcp_fwd_fin":                       ret.DtSlbL4Stats.Stats.Tcp_fwd_fin,
			"tcp_fwd_ackfin":                    ret.DtSlbL4Stats.Stats.Tcp_fwd_ackfin,
			"inudp":                             ret.DtSlbL4Stats.Stats.Inudp,
			"syncookiessent":                    ret.DtSlbL4Stats.Stats.Syncookiessent,
			"syncookiessent_ts":                 ret.DtSlbL4Stats.Stats.Syncookiessent_ts,
			"syncookiessentfailed":              ret.DtSlbL4Stats.Stats.Syncookiessentfailed,
			"outrst":                            ret.DtSlbL4Stats.Stats.Outrst,
			"outrst_nosyn":                      ret.DtSlbL4Stats.Stats.Outrst_nosyn,
			"outrst_broker":                     ret.DtSlbL4Stats.Stats.Outrst_broker,
			"outrst_ack_attack":                 ret.DtSlbL4Stats.Stats.Outrst_ack_attack,
			"outrst_aflex":                      ret.DtSlbL4Stats.Stats.Outrst_aflex,
			"outrst_stale_sess":                 ret.DtSlbL4Stats.Stats.Outrst_stale_sess,
			"syn_stale_sess":                    ret.DtSlbL4Stats.Stats.Syn_stale_sess,
			"outrst_tcpproxy":                   ret.DtSlbL4Stats.Stats.Outrst_tcpproxy,
			"svrselfail":                        ret.DtSlbL4Stats.Stats.Svrselfail,
			"noroute":                           ret.DtSlbL4Stats.Stats.Noroute,
			"snat_fail":                         ret.DtSlbL4Stats.Stats.Snat_fail,
			"snat_no_fwd_route":                 ret.DtSlbL4Stats.Stats.Snat_no_fwd_route,
			"snat_no_rev_route":                 ret.DtSlbL4Stats.Stats.Snat_no_rev_route,
			"snat_icmp_error_process":           ret.DtSlbL4Stats.Stats.Snat_icmp_error_process,
			"snat_icmp_no_match":                ret.DtSlbL4Stats.Stats.Snat_icmp_no_match,
			"smart_nat_id_mismatch":             ret.DtSlbL4Stats.Stats.Smart_nat_id_mismatch,
			"syncookiescheckfailed":             ret.DtSlbL4Stats.Stats.Syncookiescheckfailed,
			"novport_drop":                      ret.DtSlbL4Stats.Stats.Novport_drop,
			"no_vport_drop":                     ret.DtSlbL4Stats.Stats.No_vport_drop,
			"nosyn_drop":                        ret.DtSlbL4Stats.Stats.Nosyn_drop,
			"nosyn_drop_fin":                    ret.DtSlbL4Stats.Stats.Nosyn_drop_fin,
			"nosyn_drop_rst":                    ret.DtSlbL4Stats.Stats.Nosyn_drop_rst,
			"nosyn_drop_ack":                    ret.DtSlbL4Stats.Stats.Nosyn_drop_ack,
			"connlimit_drop":                    ret.DtSlbL4Stats.Stats.Connlimit_drop,
			"connlimit_reset":                   ret.DtSlbL4Stats.Stats.Connlimit_reset,
			"conn_rate_limit_drop":              ret.DtSlbL4Stats.Stats.Conn_rate_limit_drop,
			"conn_rate_limit_reset":             ret.DtSlbL4Stats.Stats.Conn_rate_limit_reset,
			"proxy_nosock_drop":                 ret.DtSlbL4Stats.Stats.Proxy_nosock_drop,
			"drop_aflex":                        ret.DtSlbL4Stats.Stats.Drop_aflex,
			"sess_aged_out":                     ret.DtSlbL4Stats.Stats.Sess_aged_out,
			"tcp_sess_aged_out":                 ret.DtSlbL4Stats.Stats.Tcp_sess_aged_out,
			"udp_sess_aged_out":                 ret.DtSlbL4Stats.Stats.Udp_sess_aged_out,
			"other_sess_aged_out":               ret.DtSlbL4Stats.Stats.Other_sess_aged_out,
			"tcp_no_slb":                        ret.DtSlbL4Stats.Stats.Tcp_no_slb,
			"udp_no_slb":                        ret.DtSlbL4Stats.Stats.Udp_no_slb,
			"throttle_syn":                      ret.DtSlbL4Stats.Stats.Throttle_syn,
			"drop_gslb":                         ret.DtSlbL4Stats.Stats.Drop_gslb,
			"inband_hm_retry":                   ret.DtSlbL4Stats.Stats.Inband_hm_retry,
			"inband_hm_reassign":                ret.DtSlbL4Stats.Stats.Inband_hm_reassign,
			"auto_reassign":                     ret.DtSlbL4Stats.Stats.Auto_reassign,
			"fast_aging_set":                    ret.DtSlbL4Stats.Stats.Fast_aging_set,
			"fast_aging_reset":                  ret.DtSlbL4Stats.Stats.Fast_aging_reset,
			"dns_policy_drop":                   ret.DtSlbL4Stats.Stats.Dns_policy_drop,
			"tcp_invalid_drop":                  ret.DtSlbL4Stats.Stats.Tcp_invalid_drop,
			"anomaly_out_seq":                   ret.DtSlbL4Stats.Stats.Anomaly_out_seq,
			"anomaly_zero_win":                  ret.DtSlbL4Stats.Stats.Anomaly_zero_win,
			"anomaly_bad_content":               ret.DtSlbL4Stats.Stats.Anomaly_bad_content,
			"anomaly_pbslb_drop":                ret.DtSlbL4Stats.Stats.Anomaly_pbslb_drop,
			"no_resourse_drop":                  ret.DtSlbL4Stats.Stats.No_resourse_drop,
			"reset_unknown_conn":                ret.DtSlbL4Stats.Stats.Reset_unknown_conn,
			"reset_l7_on_failover":              ret.DtSlbL4Stats.Stats.Reset_l7_on_failover,
			"ignore_msl":                        ret.DtSlbL4Stats.Stats.Ignore_msl,
			"l2_dsr":                            ret.DtSlbL4Stats.Stats.L2_dsr,
			"l3_dsr":                            ret.DtSlbL4Stats.Stats.L3_dsr,
			"port_preserve_attempt":             ret.DtSlbL4Stats.Stats.Port_preserve_attempt,
			"port_preserve_succ":                ret.DtSlbL4Stats.Stats.Port_preserve_succ,
			"tcpsyndata_drop":                   ret.DtSlbL4Stats.Stats.Tcpsyndata_drop,
			"tcpotherflags_drop":                ret.DtSlbL4Stats.Stats.Tcpotherflags_drop,
			"bw_rate_limit_exceed":              ret.DtSlbL4Stats.Stats.Bw_rate_limit_exceed,
			"bw_watermark_drop":                 ret.DtSlbL4Stats.Stats.Bw_watermark_drop,
			"l4_cps_exceed":                     ret.DtSlbL4Stats.Stats.L4_cps_exceed,
			"nat_cps_exceed":                    ret.DtSlbL4Stats.Stats.Nat_cps_exceed,
			"l7_cps_exceed":                     ret.DtSlbL4Stats.Stats.L7_cps_exceed,
			"ssl_cps_exceed":                    ret.DtSlbL4Stats.Stats.Ssl_cps_exceed,
			"ssl_tpt_exceed":                    ret.DtSlbL4Stats.Stats.Ssl_tpt_exceed,
			"ssl_watermark_drop":                ret.DtSlbL4Stats.Stats.Ssl_watermark_drop,
			"concurrent_conn_exceed":            ret.DtSlbL4Stats.Stats.Concurrent_conn_exceed,
			"svr_syn_handshake_fail":            ret.DtSlbL4Stats.Stats.Svr_syn_handshake_fail,
			"stateless_conn_timeout":            ret.DtSlbL4Stats.Stats.Stateless_conn_timeout,
			"tcp_ax_rexmit_syn":                 ret.DtSlbL4Stats.Stats.Tcp_ax_rexmit_syn,
			"tcp_syn_rcv_ack":                   ret.DtSlbL4Stats.Stats.Tcp_syn_rcv_ack,
			"tcp_syn_rcv_rst":                   ret.DtSlbL4Stats.Stats.Tcp_syn_rcv_rst,
			"tcp_sess_noest_aged_out":           ret.DtSlbL4Stats.Stats.Tcp_sess_noest_aged_out,
			"tcp_sess_noest_csyn_rcv_aged_out":  ret.DtSlbL4Stats.Stats.Tcp_sess_noest_csyn_rcv_aged_out,
			"tcp_sess_noest_ssyn_xmit_aged_out": ret.DtSlbL4Stats.Stats.Tcp_sess_noest_ssyn_xmit_aged_out,
			"tcp_rexmit_syn":                    ret.DtSlbL4Stats.Stats.Tcp_rexmit_syn,
			"tcp_rexmit_syn_delq":               ret.DtSlbL4Stats.Stats.Tcp_rexmit_syn_delq,
			"tcp_rexmit_synack":                 ret.DtSlbL4Stats.Stats.Tcp_rexmit_synack,
			"tcp_rexmit_synack_delq":            ret.DtSlbL4Stats.Stats.Tcp_rexmit_synack_delq,
			"tcp_fwd_fin_dup":                   ret.DtSlbL4Stats.Stats.Tcp_fwd_fin_dup,
			"tcp_rev_fin_dup":                   ret.DtSlbL4Stats.Stats.Tcp_rev_fin_dup,
			"tcp_rev_ackfin":                    ret.DtSlbL4Stats.Stats.Tcp_rev_ackfin,
			"tcp_fwd_rst":                       ret.DtSlbL4Stats.Stats.Tcp_fwd_rst,
			"tcp_rev_rst":                       ret.DtSlbL4Stats.Stats.Tcp_rev_rst,
			"udp_req_oneplus_no_resp":           ret.DtSlbL4Stats.Stats.Udp_req_oneplus_no_resp,
			"udp_req_one_oneplus_resp":          ret.DtSlbL4Stats.Stats.Udp_req_one_oneplus_resp,
			"udp_req_resp_notmatch":             ret.DtSlbL4Stats.Stats.Udp_req_resp_notmatch,
			"udp_req_more_resp":                 ret.DtSlbL4Stats.Stats.Udp_req_more_resp,
			"udp_resp_more_req":                 ret.DtSlbL4Stats.Stats.Udp_resp_more_req,
			"udp_req_oneplus":                   ret.DtSlbL4Stats.Stats.Udp_req_oneplus,
			"udp_resp_oneplus":                  ret.DtSlbL4Stats.Stats.Udp_resp_oneplus,
			"out_seq_ack_drop":                  ret.DtSlbL4Stats.Stats.Out_seq_ack_drop,
			"tcp_est":                           ret.DtSlbL4Stats.Stats.Tcp_est,
			"synattack":                         ret.DtSlbL4Stats.Stats.Synattack,
			"syn_rate":                          ret.DtSlbL4Stats.Stats.Syn_rate,
			"syncookie_buff_drop":               ret.DtSlbL4Stats.Stats.Syncookie_buff_drop,
			"syncookie_buff_queue":              ret.DtSlbL4Stats.Stats.Syncookie_buff_queue,
			"skip_insert_client_ip":             ret.DtSlbL4Stats.Stats.Skip_insert_client_ip,
			"synreceived_hw":                    ret.DtSlbL4Stats.Stats.Synreceived_hw,
			"dns_id_switch":                     ret.DtSlbL4Stats.Stats.Dns_id_switch,
			"server_down_del":                   ret.DtSlbL4Stats.Stats.Server_down_del,
			"dnssec_switch":                     ret.DtSlbL4Stats.Stats.Dnssec_switch,
			"rate_drop_reset_unkn":              ret.DtSlbL4Stats.Stats.Rate_drop_reset_unkn,
			"tcp_connections_closed":            ret.DtSlbL4Stats.Stats.Tcp_connections_closed,
			"gtp_c_invalid_port":                ret.DtSlbL4Stats.Stats.Gtp_c_invalid_port,
			"gtp_c_invalid_header":              ret.DtSlbL4Stats.Stats.Gtp_c_invalid_header,
			"gtp_c_invalid_message":             ret.DtSlbL4Stats.Stats.Gtp_c_invalid_message,
			"reselect_svrselfail":               ret.DtSlbL4Stats.Stats.Reselect_svrselfail,
			"snat_port_overload_fail":           ret.DtSlbL4Stats.Stats.Snat_port_overload_fail,
			"snat_force_preserve_alloc":         ret.DtSlbL4Stats.Stats.Snat_force_preserve_alloc,
			"snat_force_preserve_free":          ret.DtSlbL4Stats.Stats.Snat_force_preserve_free,
			"proxy_header_insert":               ret.DtSlbL4Stats.Stats.Proxy_header_insert,
			"proxy_header_rexmit":               ret.DtSlbL4Stats.Stats.Proxy_header_rexmit,
			"proxy_prot_error":                  ret.DtSlbL4Stats.Stats.Proxy_prot_error,
			"proxy_prot_drop":                   ret.DtSlbL4Stats.Stats.Proxy_prot_drop,
			"slb_gtp_proxy_smp_match":           ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_smp_match,
			"slb_gtp_proxy_smp_no_match":        ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_smp_no_match,
			"slb_gtp_proxy_c_process_local_rr":  ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_c_process_local_rr,
			"slb_gtp_proxy_smp_creation_failed": ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_smp_creation_failed,
			"slb_gtp_proxy_smp_created":         ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_smp_created,
			"slb_gtp_proxy_smp_free_not_found":  ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_smp_free_not_found,
			"slb_gtp_proxy_smp_freed":           ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_smp_freed,
			"slb_gtp_proxy_retx_requests":       ret.DtSlbL4Stats.Stats.Slb_gtp_proxy_retx_requests,
			"pbslb_entry_limit_exceed":          ret.DtSlbL4Stats.Stats.Pbslb_entry_limit_exceed,
			"fast_path_reroute":                 ret.DtSlbL4Stats.Stats.Fast_path_reroute,
			"fast_path_l2_reroute":              ret.DtSlbL4Stats.Stats.Fast_path_l2_reroute,
		},
	}
}

func getObjectSlbL4StatsStats(d []interface{}) edpt.SlbL4StatsStats {

	count1 := len(d)
	var ret edpt.SlbL4StatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Intcp = in["intcp"].(int)
		ret.Synreceived = in["synreceived"].(int)
		ret.Tcp_fwd_last_ack = in["tcp_fwd_last_ack"].(int)
		ret.Tcp_rev_last_ack = in["tcp_rev_last_ack"].(int)
		ret.Tcp_rev_fin = in["tcp_rev_fin"].(int)
		ret.Tcp_fwd_fin = in["tcp_fwd_fin"].(int)
		ret.Tcp_fwd_ackfin = in["tcp_fwd_ackfin"].(int)
		ret.Inudp = in["inudp"].(int)
		ret.Syncookiessent = in["syncookiessent"].(int)
		ret.Syncookiessent_ts = in["syncookiessent_ts"].(int)
		ret.Syncookiessentfailed = in["syncookiessentfailed"].(int)
		ret.Outrst = in["outrst"].(int)
		ret.Outrst_nosyn = in["outrst_nosyn"].(int)
		ret.Outrst_broker = in["outrst_broker"].(int)
		ret.Outrst_ack_attack = in["outrst_ack_attack"].(int)
		ret.Outrst_aflex = in["outrst_aflex"].(int)
		ret.Outrst_stale_sess = in["outrst_stale_sess"].(int)
		ret.Syn_stale_sess = in["syn_stale_sess"].(int)
		ret.Outrst_tcpproxy = in["outrst_tcpproxy"].(int)
		ret.Svrselfail = in["svrselfail"].(int)
		ret.Noroute = in["noroute"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Snat_no_fwd_route = in["snat_no_fwd_route"].(int)
		ret.Snat_no_rev_route = in["snat_no_rev_route"].(int)
		ret.Snat_icmp_error_process = in["snat_icmp_error_process"].(int)
		ret.Snat_icmp_no_match = in["snat_icmp_no_match"].(int)
		ret.Smart_nat_id_mismatch = in["smart_nat_id_mismatch"].(int)
		ret.Syncookiescheckfailed = in["syncookiescheckfailed"].(int)
		ret.Novport_drop = in["novport_drop"].(int)
		ret.No_vport_drop = in["no_vport_drop"].(int)
		ret.Nosyn_drop = in["nosyn_drop"].(int)
		ret.Nosyn_drop_fin = in["nosyn_drop_fin"].(int)
		ret.Nosyn_drop_rst = in["nosyn_drop_rst"].(int)
		ret.Nosyn_drop_ack = in["nosyn_drop_ack"].(int)
		ret.Connlimit_drop = in["connlimit_drop"].(int)
		ret.Connlimit_reset = in["connlimit_reset"].(int)
		ret.Conn_rate_limit_drop = in["conn_rate_limit_drop"].(int)
		ret.Conn_rate_limit_reset = in["conn_rate_limit_reset"].(int)
		ret.Proxy_nosock_drop = in["proxy_nosock_drop"].(int)
		ret.Drop_aflex = in["drop_aflex"].(int)
		ret.Sess_aged_out = in["sess_aged_out"].(int)
		ret.Tcp_sess_aged_out = in["tcp_sess_aged_out"].(int)
		ret.Udp_sess_aged_out = in["udp_sess_aged_out"].(int)
		ret.Other_sess_aged_out = in["other_sess_aged_out"].(int)
		ret.Tcp_no_slb = in["tcp_no_slb"].(int)
		ret.Udp_no_slb = in["udp_no_slb"].(int)
		ret.Throttle_syn = in["throttle_syn"].(int)
		ret.Drop_gslb = in["drop_gslb"].(int)
		ret.Inband_hm_retry = in["inband_hm_retry"].(int)
		ret.Inband_hm_reassign = in["inband_hm_reassign"].(int)
		ret.Auto_reassign = in["auto_reassign"].(int)
		ret.Fast_aging_set = in["fast_aging_set"].(int)
		ret.Fast_aging_reset = in["fast_aging_reset"].(int)
		ret.Dns_policy_drop = in["dns_policy_drop"].(int)
		ret.Tcp_invalid_drop = in["tcp_invalid_drop"].(int)
		ret.Anomaly_out_seq = in["anomaly_out_seq"].(int)
		ret.Anomaly_zero_win = in["anomaly_zero_win"].(int)
		ret.Anomaly_bad_content = in["anomaly_bad_content"].(int)
		ret.Anomaly_pbslb_drop = in["anomaly_pbslb_drop"].(int)
		ret.No_resourse_drop = in["no_resourse_drop"].(int)
		ret.Reset_unknown_conn = in["reset_unknown_conn"].(int)
		ret.Reset_l7_on_failover = in["reset_l7_on_failover"].(int)
		ret.Ignore_msl = in["ignore_msl"].(int)
		ret.L2_dsr = in["l2_dsr"].(int)
		ret.L3_dsr = in["l3_dsr"].(int)
		ret.Port_preserve_attempt = in["port_preserve_attempt"].(int)
		ret.Port_preserve_succ = in["port_preserve_succ"].(int)
		ret.Tcpsyndata_drop = in["tcpsyndata_drop"].(int)
		ret.Tcpotherflags_drop = in["tcpotherflags_drop"].(int)
		ret.Bw_rate_limit_exceed = in["bw_rate_limit_exceed"].(int)
		ret.Bw_watermark_drop = in["bw_watermark_drop"].(int)
		ret.L4_cps_exceed = in["l4_cps_exceed"].(int)
		ret.Nat_cps_exceed = in["nat_cps_exceed"].(int)
		ret.L7_cps_exceed = in["l7_cps_exceed"].(int)
		ret.Ssl_cps_exceed = in["ssl_cps_exceed"].(int)
		ret.Ssl_tpt_exceed = in["ssl_tpt_exceed"].(int)
		ret.Ssl_watermark_drop = in["ssl_watermark_drop"].(int)
		ret.Concurrent_conn_exceed = in["concurrent_conn_exceed"].(int)
		ret.Svr_syn_handshake_fail = in["svr_syn_handshake_fail"].(int)
		ret.Stateless_conn_timeout = in["stateless_conn_timeout"].(int)
		ret.Tcp_ax_rexmit_syn = in["tcp_ax_rexmit_syn"].(int)
		ret.Tcp_syn_rcv_ack = in["tcp_syn_rcv_ack"].(int)
		ret.Tcp_syn_rcv_rst = in["tcp_syn_rcv_rst"].(int)
		ret.Tcp_sess_noest_aged_out = in["tcp_sess_noest_aged_out"].(int)
		ret.Tcp_sess_noest_csyn_rcv_aged_out = in["tcp_sess_noest_csyn_rcv_aged_out"].(int)
		ret.Tcp_sess_noest_ssyn_xmit_aged_out = in["tcp_sess_noest_ssyn_xmit_aged_out"].(int)
		ret.Tcp_rexmit_syn = in["tcp_rexmit_syn"].(int)
		ret.Tcp_rexmit_syn_delq = in["tcp_rexmit_syn_delq"].(int)
		ret.Tcp_rexmit_synack = in["tcp_rexmit_synack"].(int)
		ret.Tcp_rexmit_synack_delq = in["tcp_rexmit_synack_delq"].(int)
		ret.Tcp_fwd_fin_dup = in["tcp_fwd_fin_dup"].(int)
		ret.Tcp_rev_fin_dup = in["tcp_rev_fin_dup"].(int)
		ret.Tcp_rev_ackfin = in["tcp_rev_ackfin"].(int)
		ret.Tcp_fwd_rst = in["tcp_fwd_rst"].(int)
		ret.Tcp_rev_rst = in["tcp_rev_rst"].(int)
		ret.Udp_req_oneplus_no_resp = in["udp_req_oneplus_no_resp"].(int)
		ret.Udp_req_one_oneplus_resp = in["udp_req_one_oneplus_resp"].(int)
		ret.Udp_req_resp_notmatch = in["udp_req_resp_notmatch"].(int)
		ret.Udp_req_more_resp = in["udp_req_more_resp"].(int)
		ret.Udp_resp_more_req = in["udp_resp_more_req"].(int)
		ret.Udp_req_oneplus = in["udp_req_oneplus"].(int)
		ret.Udp_resp_oneplus = in["udp_resp_oneplus"].(int)
		ret.Out_seq_ack_drop = in["out_seq_ack_drop"].(int)
		ret.Tcp_est = in["tcp_est"].(int)
		ret.Synattack = in["synattack"].(int)
		ret.Syn_rate = in["syn_rate"].(int)
		ret.Syncookie_buff_drop = in["syncookie_buff_drop"].(int)
		ret.Syncookie_buff_queue = in["syncookie_buff_queue"].(int)
		ret.Skip_insert_client_ip = in["skip_insert_client_ip"].(int)
		ret.Synreceived_hw = in["synreceived_hw"].(int)
		ret.Dns_id_switch = in["dns_id_switch"].(int)
		ret.Server_down_del = in["server_down_del"].(int)
		ret.Dnssec_switch = in["dnssec_switch"].(int)
		ret.Rate_drop_reset_unkn = in["rate_drop_reset_unkn"].(int)
		ret.Tcp_connections_closed = in["tcp_connections_closed"].(int)
		ret.Gtp_c_invalid_port = in["gtp_c_invalid_port"].(int)
		ret.Gtp_c_invalid_header = in["gtp_c_invalid_header"].(int)
		ret.Gtp_c_invalid_message = in["gtp_c_invalid_message"].(int)
		ret.Reselect_svrselfail = in["reselect_svrselfail"].(int)
		ret.Snat_port_overload_fail = in["snat_port_overload_fail"].(int)
		ret.Snat_force_preserve_alloc = in["snat_force_preserve_alloc"].(int)
		ret.Snat_force_preserve_free = in["snat_force_preserve_free"].(int)
		ret.Proxy_header_insert = in["proxy_header_insert"].(int)
		ret.Proxy_header_rexmit = in["proxy_header_rexmit"].(int)
		ret.Proxy_prot_error = in["proxy_prot_error"].(int)
		ret.Proxy_prot_drop = in["proxy_prot_drop"].(int)
		ret.Slb_gtp_proxy_smp_match = in["slb_gtp_proxy_smp_match"].(int)
		ret.Slb_gtp_proxy_smp_no_match = in["slb_gtp_proxy_smp_no_match"].(int)
		ret.Slb_gtp_proxy_c_process_local_rr = in["slb_gtp_proxy_c_process_local_rr"].(int)
		ret.Slb_gtp_proxy_smp_creation_failed = in["slb_gtp_proxy_smp_creation_failed"].(int)
		ret.Slb_gtp_proxy_smp_created = in["slb_gtp_proxy_smp_created"].(int)
		ret.Slb_gtp_proxy_smp_free_not_found = in["slb_gtp_proxy_smp_free_not_found"].(int)
		ret.Slb_gtp_proxy_smp_freed = in["slb_gtp_proxy_smp_freed"].(int)
		ret.Slb_gtp_proxy_retx_requests = in["slb_gtp_proxy_retx_requests"].(int)
		ret.Pbslb_entry_limit_exceed = in["pbslb_entry_limit_exceed"].(int)
		ret.Fast_path_reroute = in["fast_path_reroute"].(int)
		ret.Fast_path_l2_reroute = in["fast_path_l2_reroute"].(int)
	}
	return ret
}

func dataToEndpointSlbL4Stats(d *schema.ResourceData) edpt.SlbL4Stats {
	var ret edpt.SlbL4Stats

	ret.Stats = getObjectSlbL4StatsStats(d.Get("stats").([]interface{}))
	return ret
}
