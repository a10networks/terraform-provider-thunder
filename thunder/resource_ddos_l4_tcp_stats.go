package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosL4TcpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_l4_tcp_stats`: Statistics for the object l4-tcp\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosL4TcpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Sessions Created",
						},
						"intcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Received",
						},
						"tcp_syn_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Received",
						},
						"tcp_invalid_syn_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Invalid SYN Received",
						},
						"tcp_syn_ack_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN ACK Received",
						},
						"tcp_ack_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Received",
						},
						"tcp_fin_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Received",
						},
						"tcp_rst_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Received",
						},
						"tcp_outrst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Outbound RST",
						},
						"tcp_reset_client": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Reset Client",
						},
						"tcp_reset_server": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Reset Server",
						},
						"tcp_syn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Per Sec",
						},
						"tcp_total_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Dropped",
						},
						"tcp_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst Packets Dropped",
						},
						"tcp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Packets Dropped",
						},
						"tcp_drop_black_user_cfg_src": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Blacklist User Packets Dropped",
						},
						"tcp_src_dst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SrcDst Packets Dropped",
						},
						"tcp_drop_black_user_cfg_src_dst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SrcDst Blacklist User Packets Dropped",
						},
						"tcp_port_zero_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port 0 Packets Dropped",
						},
						"tcp_syncookie_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Sent",
						},
						"tcp_syncookie_sent_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Send Failed",
						},
						"tcp_syncookie_check_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Check Failed",
						},
						"tcp_syncookie_hw_missing": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie HW Missed",
						},
						"tcp_syncookie_fail_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Blacklist Failed",
						},
						"tcp_syncookie_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Passed",
						},
						"syn_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Auth Passed",
						},
						"syn_auth_skip": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Auth Skipped",
						},
						"tcp_action_on_ack_start": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Init",
						},
						"tcp_action_on_ack_matched": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Matched",
						},
						"tcp_action_on_ack_passed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Passed",
						},
						"tcp_action_on_ack_failed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Dropped",
						},
						"tcp_action_on_ack_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Timeout",
						},
						"tcp_action_on_ack_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Timeout Reset",
						},
						"tcp_ack_no_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK No SYN",
						},
						"tcp_out_of_seq": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Total",
						},
						"tcp_zero_window": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Total",
						},
						"tcp_retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Total",
						},
						"tcp_rexmit_syn_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit SYN Exceed Dropped",
						},
						"tcp_zero_window_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Blacklisted",
						},
						"tcp_out_of_seq_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Blacklisted",
						},
						"tcp_retransmit_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Blacklisted",
						},
						"tcp_rexmit_syn_limit_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit SYN Exceed Blacklisted",
						},
						"tcp_per_conn_prate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Pkt Rate Dropped",
						},
						"tcp_action_on_ack_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Retry-Gap Dropped",
						},
						"tcp_action_on_ack_gap_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Retry-Gap Passed",
						},
						"tcp_action_on_syn_start": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Init",
						},
						"tcp_action_on_syn_passed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Passed",
						},
						"tcp_action_on_syn_failed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Dropped",
						},
						"tcp_action_on_syn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Timeout",
						},
						"tcp_action_on_syn_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Timeout Reset",
						},
						"tcp_action_on_syn_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Retry-Gap Dropped",
						},
						"tcp_action_on_syn_gap_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Retry-Gap Passed",
						},
						"tcp_unauth_rst_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Unauth RST Dropped",
						},
						"dst_tcp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Match",
						},
						"dst_tcp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter No Match",
						},
						"dst_tcp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Blacklist",
						},
						"dst_tcp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Drop",
						},
						"dst_tcp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action Default Pass",
						},
						"tcp_concurrent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Concurrent Port Access",
						},
						"dst_tcp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Filter Action WL",
						},
						"src_tcp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Match",
						},
						"src_tcp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
						},
						"src_tcp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
						},
						"src_tcp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
						},
						"src_tcp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
						},
						"src_tcp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action WL",
						},
						"src_dst_tcp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Match",
						},
						"src_dst_tcp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter No Match",
						},
						"src_dst_tcp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Blacklist",
						},
						"src_dst_tcp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Drop",
						},
						"src_dst_tcp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action Default Pass",
						},
						"src_dst_tcp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "SrcDst Filter Action WL",
						},
						"syn_auth_pass_wl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Auth Pass WL",
						},
						"tcp_out_of_seq_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Dropped",
						},
						"tcp_zero_window_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Dropped",
						},
						"tcp_retransmit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Dropped",
						},
						"tcp_per_conn_prate_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Pkt Rate Blacklisted",
						},
						"tcp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Exceeded",
						},
						"tcp_drop_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Blacklist Packets Dropped",
						},
						"tcp_frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Frag Received",
						},
						"tcp_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Frag Dropped",
						},
						"tcp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Dropped",
						},
						"tcp_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Responded",
						},
						"tcp_total_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Received",
						},
						"tcp_total_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Dropped",
						},
						"tcp_action_on_ack_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK Retry Timeout Blacklisted",
						},
						"tcp_action_on_syn_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retry Timeout Blacklisted",
						},
						"tcp_per_conn_ofo_rate_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Out-Of-Seq Rate Dropped",
						},
						"tcp_per_conn_ofo_rate_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Out-Of-Seq Rate Blacklisted",
						},
						"tcp_per_conn_rexmit_rate_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Retransmit Rate Dropped",
						},
						"tcp_per_conn_rexmit_rate_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Retransmit Rate Blacklisted",
						},
						"tcp_per_conn_zwindow_rate_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Zero-Window Rate Dropped",
						},
						"tcp_per_conn_zwindow_rate_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Zero-Window Rate Blacklisted",
						},
						"tcp_syn_tfo_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN TFO Received",
						},
						"tcp_progression_violation_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Progression: Violation Exceeded",
						},
						"tcp_progression_violation_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Progression: Violation Exceeded Blacklisted",
						},
						"tcp_progression_violation_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Progression: Violation Exceeded Dropped",
						},
						"tcp_progression_violation_exceed_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Progression: Violation Exceeded Reset",
						},
						"tcp_auth_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Reset",
						},
					},
				},
			},
		},
	}
}

func resourceDdosL4TcpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosL4TcpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosL4TcpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosL4TcpStatsStats := setObjectDdosL4TcpStatsStats(res)
		d.Set("stats", DdosL4TcpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosL4TcpStatsStats(ret edpt.DataDdosL4TcpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_sess_create":                        ret.DtDdosL4TcpStats.Stats.Tcp_sess_create,
			"intcp":                                  ret.DtDdosL4TcpStats.Stats.Intcp,
			"tcp_syn_rcvd":                           ret.DtDdosL4TcpStats.Stats.Tcp_syn_rcvd,
			"tcp_invalid_syn_rcvd":                   ret.DtDdosL4TcpStats.Stats.Tcp_invalid_syn_rcvd,
			"tcp_syn_ack_rcvd":                       ret.DtDdosL4TcpStats.Stats.Tcp_syn_ack_rcvd,
			"tcp_ack_rcvd":                           ret.DtDdosL4TcpStats.Stats.Tcp_ack_rcvd,
			"tcp_fin_rcvd":                           ret.DtDdosL4TcpStats.Stats.Tcp_fin_rcvd,
			"tcp_rst_rcvd":                           ret.DtDdosL4TcpStats.Stats.Tcp_rst_rcvd,
			"tcp_outrst":                             ret.DtDdosL4TcpStats.Stats.Tcp_outrst,
			"tcp_reset_client":                       ret.DtDdosL4TcpStats.Stats.Tcp_reset_client,
			"tcp_reset_server":                       ret.DtDdosL4TcpStats.Stats.Tcp_reset_server,
			"tcp_syn_rate":                           ret.DtDdosL4TcpStats.Stats.Tcp_syn_rate,
			"tcp_total_drop":                         ret.DtDdosL4TcpStats.Stats.Tcp_total_drop,
			"tcp_dst_drop":                           ret.DtDdosL4TcpStats.Stats.Tcp_dst_drop,
			"tcp_src_drop":                           ret.DtDdosL4TcpStats.Stats.Tcp_src_drop,
			"tcp_drop_black_user_cfg_src":            ret.DtDdosL4TcpStats.Stats.Tcp_drop_black_user_cfg_src,
			"tcp_src_dst_drop":                       ret.DtDdosL4TcpStats.Stats.Tcp_src_dst_drop,
			"tcp_drop_black_user_cfg_src_dst":        ret.DtDdosL4TcpStats.Stats.Tcp_drop_black_user_cfg_src_dst,
			"tcp_port_zero_drop":                     ret.DtDdosL4TcpStats.Stats.Tcp_port_zero_drop,
			"tcp_syncookie_sent":                     ret.DtDdosL4TcpStats.Stats.Tcp_syncookie_sent,
			"tcp_syncookie_sent_fail":                ret.DtDdosL4TcpStats.Stats.Tcp_syncookie_sent_fail,
			"tcp_syncookie_check_fail":               ret.DtDdosL4TcpStats.Stats.Tcp_syncookie_check_fail,
			"tcp_syncookie_hw_missing":               ret.DtDdosL4TcpStats.Stats.Tcp_syncookie_hw_missing,
			"tcp_syncookie_fail_bl":                  ret.DtDdosL4TcpStats.Stats.Tcp_syncookie_fail_bl,
			"tcp_syncookie_pass":                     ret.DtDdosL4TcpStats.Stats.Tcp_syncookie_pass,
			"syn_auth_pass":                          ret.DtDdosL4TcpStats.Stats.Syn_auth_pass,
			"syn_auth_skip":                          ret.DtDdosL4TcpStats.Stats.Syn_auth_skip,
			"tcp_action_on_ack_start":                ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_start,
			"tcp_action_on_ack_matched":              ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_matched,
			"tcp_action_on_ack_passed":               ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_passed,
			"tcp_action_on_ack_failed":               ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_failed,
			"tcp_action_on_ack_timeout":              ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_timeout,
			"tcp_action_on_ack_reset":                ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_reset,
			"tcp_ack_no_syn":                         ret.DtDdosL4TcpStats.Stats.Tcp_ack_no_syn,
			"tcp_out_of_seq":                         ret.DtDdosL4TcpStats.Stats.Tcp_out_of_seq,
			"tcp_zero_window":                        ret.DtDdosL4TcpStats.Stats.Tcp_zero_window,
			"tcp_retransmit":                         ret.DtDdosL4TcpStats.Stats.Tcp_retransmit,
			"tcp_rexmit_syn_limit_drop":              ret.DtDdosL4TcpStats.Stats.Tcp_rexmit_syn_limit_drop,
			"tcp_zero_window_bl":                     ret.DtDdosL4TcpStats.Stats.Tcp_zero_window_bl,
			"tcp_out_of_seq_bl":                      ret.DtDdosL4TcpStats.Stats.Tcp_out_of_seq_bl,
			"tcp_retransmit_bl":                      ret.DtDdosL4TcpStats.Stats.Tcp_retransmit_bl,
			"tcp_rexmit_syn_limit_bl":                ret.DtDdosL4TcpStats.Stats.Tcp_rexmit_syn_limit_bl,
			"tcp_per_conn_prate_exceed":              ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_prate_exceed,
			"tcp_action_on_ack_gap_drop":             ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_gap_drop,
			"tcp_action_on_ack_gap_pass":             ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_gap_pass,
			"tcp_action_on_syn_start":                ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_start,
			"tcp_action_on_syn_passed":               ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_passed,
			"tcp_action_on_syn_failed":               ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_failed,
			"tcp_action_on_syn_timeout":              ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_timeout,
			"tcp_action_on_syn_reset":                ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_reset,
			"tcp_action_on_syn_gap_drop":             ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_gap_drop,
			"tcp_action_on_syn_gap_pass":             ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_gap_pass,
			"tcp_unauth_rst_drop":                    ret.DtDdosL4TcpStats.Stats.Tcp_unauth_rst_drop,
			"dst_tcp_filter_match":                   ret.DtDdosL4TcpStats.Stats.Dst_tcp_filter_match,
			"dst_tcp_filter_not_match":               ret.DtDdosL4TcpStats.Stats.Dst_tcp_filter_not_match,
			"dst_tcp_filter_action_blacklist":        ret.DtDdosL4TcpStats.Stats.Dst_tcp_filter_action_blacklist,
			"dst_tcp_filter_action_drop":             ret.DtDdosL4TcpStats.Stats.Dst_tcp_filter_action_drop,
			"dst_tcp_filter_action_default_pass":     ret.DtDdosL4TcpStats.Stats.Dst_tcp_filter_action_default_pass,
			"tcp_concurrent":                         ret.DtDdosL4TcpStats.Stats.Tcp_concurrent,
			"dst_tcp_filter_action_whitelist":        ret.DtDdosL4TcpStats.Stats.Dst_tcp_filter_action_whitelist,
			"src_tcp_filter_match":                   ret.DtDdosL4TcpStats.Stats.Src_tcp_filter_match,
			"src_tcp_filter_not_match":               ret.DtDdosL4TcpStats.Stats.Src_tcp_filter_not_match,
			"src_tcp_filter_action_blacklist":        ret.DtDdosL4TcpStats.Stats.Src_tcp_filter_action_blacklist,
			"src_tcp_filter_action_drop":             ret.DtDdosL4TcpStats.Stats.Src_tcp_filter_action_drop,
			"src_tcp_filter_action_default_pass":     ret.DtDdosL4TcpStats.Stats.Src_tcp_filter_action_default_pass,
			"src_tcp_filter_action_whitelist":        ret.DtDdosL4TcpStats.Stats.Src_tcp_filter_action_whitelist,
			"src_dst_tcp_filter_match":               ret.DtDdosL4TcpStats.Stats.Src_dst_tcp_filter_match,
			"src_dst_tcp_filter_not_match":           ret.DtDdosL4TcpStats.Stats.Src_dst_tcp_filter_not_match,
			"src_dst_tcp_filter_action_blacklist":    ret.DtDdosL4TcpStats.Stats.Src_dst_tcp_filter_action_blacklist,
			"src_dst_tcp_filter_action_drop":         ret.DtDdosL4TcpStats.Stats.Src_dst_tcp_filter_action_drop,
			"src_dst_tcp_filter_action_default_pass": ret.DtDdosL4TcpStats.Stats.Src_dst_tcp_filter_action_default_pass,
			"src_dst_tcp_filter_action_whitelist":    ret.DtDdosL4TcpStats.Stats.Src_dst_tcp_filter_action_whitelist,
			"syn_auth_pass_wl":                       ret.DtDdosL4TcpStats.Stats.Syn_auth_pass_wl,
			"tcp_out_of_seq_drop":                    ret.DtDdosL4TcpStats.Stats.Tcp_out_of_seq_drop,
			"tcp_zero_window_drop":                   ret.DtDdosL4TcpStats.Stats.Tcp_zero_window_drop,
			"tcp_retransmit_drop":                    ret.DtDdosL4TcpStats.Stats.Tcp_retransmit_drop,
			"tcp_per_conn_prate_exceed_bl":           ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_prate_exceed_bl,
			"tcp_any_exceed":                         ret.DtDdosL4TcpStats.Stats.Tcp_any_exceed,
			"tcp_drop_bl":                            ret.DtDdosL4TcpStats.Stats.Tcp_drop_bl,
			"tcp_frag_rcvd":                          ret.DtDdosL4TcpStats.Stats.Tcp_frag_rcvd,
			"tcp_frag_drop":                          ret.DtDdosL4TcpStats.Stats.Tcp_frag_drop,
			"tcp_auth_drop":                          ret.DtDdosL4TcpStats.Stats.Tcp_auth_drop,
			"tcp_auth_resp":                          ret.DtDdosL4TcpStats.Stats.Tcp_auth_resp,
			"tcp_total_bytes_rcv":                    ret.DtDdosL4TcpStats.Stats.Tcp_total_bytes_rcv,
			"tcp_total_bytes_drop":                   ret.DtDdosL4TcpStats.Stats.Tcp_total_bytes_drop,
			"tcp_action_on_ack_bl":                   ret.DtDdosL4TcpStats.Stats.Tcp_action_on_ack_bl,
			"tcp_action_on_syn_bl":                   ret.DtDdosL4TcpStats.Stats.Tcp_action_on_syn_bl,
			"tcp_per_conn_ofo_rate_exceed_drop":      ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_ofo_rate_exceed_drop,
			"tcp_per_conn_ofo_rate_exceed_bl":        ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_ofo_rate_exceed_bl,
			"tcp_per_conn_rexmit_rate_exceed_drop":   ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_rexmit_rate_exceed_drop,
			"tcp_per_conn_rexmit_rate_exceed_bl":     ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_rexmit_rate_exceed_bl,
			"tcp_per_conn_zwindow_rate_exceed_drop":  ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_zwindow_rate_exceed_drop,
			"tcp_per_conn_zwindow_rate_exceed_bl":    ret.DtDdosL4TcpStats.Stats.Tcp_per_conn_zwindow_rate_exceed_bl,
			"tcp_syn_tfo_rcvd":                       ret.DtDdosL4TcpStats.Stats.Tcp_syn_tfo_rcvd,
			"tcp_progression_violation_exceed":       ret.DtDdosL4TcpStats.Stats.Tcp_progression_violation_exceed,
			"tcp_progression_violation_exceed_bl":    ret.DtDdosL4TcpStats.Stats.Tcp_progression_violation_exceed_bl,
			"tcp_progression_violation_exceed_drop":  ret.DtDdosL4TcpStats.Stats.Tcp_progression_violation_exceed_drop,
			"tcp_progression_violation_exceed_reset": ret.DtDdosL4TcpStats.Stats.Tcp_progression_violation_exceed_reset,
			"tcp_auth_rst":                           ret.DtDdosL4TcpStats.Stats.Tcp_auth_rst,
		},
	}
}

func getObjectDdosL4TcpStatsStats(d []interface{}) edpt.DdosL4TcpStatsStats {

	count1 := len(d)
	var ret edpt.DdosL4TcpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp_sess_create = in["tcp_sess_create"].(int)
		ret.Intcp = in["intcp"].(int)
		ret.Tcp_syn_rcvd = in["tcp_syn_rcvd"].(int)
		ret.Tcp_invalid_syn_rcvd = in["tcp_invalid_syn_rcvd"].(int)
		ret.Tcp_syn_ack_rcvd = in["tcp_syn_ack_rcvd"].(int)
		ret.Tcp_ack_rcvd = in["tcp_ack_rcvd"].(int)
		ret.Tcp_fin_rcvd = in["tcp_fin_rcvd"].(int)
		ret.Tcp_rst_rcvd = in["tcp_rst_rcvd"].(int)
		ret.Tcp_outrst = in["tcp_outrst"].(int)
		ret.Tcp_reset_client = in["tcp_reset_client"].(int)
		ret.Tcp_reset_server = in["tcp_reset_server"].(int)
		ret.Tcp_syn_rate = in["tcp_syn_rate"].(int)
		ret.Tcp_total_drop = in["tcp_total_drop"].(int)
		ret.Tcp_dst_drop = in["tcp_dst_drop"].(int)
		ret.Tcp_src_drop = in["tcp_src_drop"].(int)
		ret.Tcp_drop_black_user_cfg_src = in["tcp_drop_black_user_cfg_src"].(int)
		ret.Tcp_src_dst_drop = in["tcp_src_dst_drop"].(int)
		ret.Tcp_drop_black_user_cfg_src_dst = in["tcp_drop_black_user_cfg_src_dst"].(int)
		ret.Tcp_port_zero_drop = in["tcp_port_zero_drop"].(int)
		ret.Tcp_syncookie_sent = in["tcp_syncookie_sent"].(int)
		ret.Tcp_syncookie_sent_fail = in["tcp_syncookie_sent_fail"].(int)
		ret.Tcp_syncookie_check_fail = in["tcp_syncookie_check_fail"].(int)
		ret.Tcp_syncookie_hw_missing = in["tcp_syncookie_hw_missing"].(int)
		ret.Tcp_syncookie_fail_bl = in["tcp_syncookie_fail_bl"].(int)
		ret.Tcp_syncookie_pass = in["tcp_syncookie_pass"].(int)
		ret.Syn_auth_pass = in["syn_auth_pass"].(int)
		ret.Syn_auth_skip = in["syn_auth_skip"].(int)
		ret.Tcp_action_on_ack_start = in["tcp_action_on_ack_start"].(int)
		ret.Tcp_action_on_ack_matched = in["tcp_action_on_ack_matched"].(int)
		ret.Tcp_action_on_ack_passed = in["tcp_action_on_ack_passed"].(int)
		ret.Tcp_action_on_ack_failed = in["tcp_action_on_ack_failed"].(int)
		ret.Tcp_action_on_ack_timeout = in["tcp_action_on_ack_timeout"].(int)
		ret.Tcp_action_on_ack_reset = in["tcp_action_on_ack_reset"].(int)
		ret.Tcp_ack_no_syn = in["tcp_ack_no_syn"].(int)
		ret.Tcp_out_of_seq = in["tcp_out_of_seq"].(int)
		ret.Tcp_zero_window = in["tcp_zero_window"].(int)
		ret.Tcp_retransmit = in["tcp_retransmit"].(int)
		ret.Tcp_rexmit_syn_limit_drop = in["tcp_rexmit_syn_limit_drop"].(int)
		ret.Tcp_zero_window_bl = in["tcp_zero_window_bl"].(int)
		ret.Tcp_out_of_seq_bl = in["tcp_out_of_seq_bl"].(int)
		ret.Tcp_retransmit_bl = in["tcp_retransmit_bl"].(int)
		ret.Tcp_rexmit_syn_limit_bl = in["tcp_rexmit_syn_limit_bl"].(int)
		ret.Tcp_per_conn_prate_exceed = in["tcp_per_conn_prate_exceed"].(int)
		ret.Tcp_action_on_ack_gap_drop = in["tcp_action_on_ack_gap_drop"].(int)
		ret.Tcp_action_on_ack_gap_pass = in["tcp_action_on_ack_gap_pass"].(int)
		ret.Tcp_action_on_syn_start = in["tcp_action_on_syn_start"].(int)
		ret.Tcp_action_on_syn_passed = in["tcp_action_on_syn_passed"].(int)
		ret.Tcp_action_on_syn_failed = in["tcp_action_on_syn_failed"].(int)
		ret.Tcp_action_on_syn_timeout = in["tcp_action_on_syn_timeout"].(int)
		ret.Tcp_action_on_syn_reset = in["tcp_action_on_syn_reset"].(int)
		ret.Tcp_action_on_syn_gap_drop = in["tcp_action_on_syn_gap_drop"].(int)
		ret.Tcp_action_on_syn_gap_pass = in["tcp_action_on_syn_gap_pass"].(int)
		ret.Tcp_unauth_rst_drop = in["tcp_unauth_rst_drop"].(int)
		ret.Dst_tcp_filter_match = in["dst_tcp_filter_match"].(int)
		ret.Dst_tcp_filter_not_match = in["dst_tcp_filter_not_match"].(int)
		ret.Dst_tcp_filter_action_blacklist = in["dst_tcp_filter_action_blacklist"].(int)
		ret.Dst_tcp_filter_action_drop = in["dst_tcp_filter_action_drop"].(int)
		ret.Dst_tcp_filter_action_default_pass = in["dst_tcp_filter_action_default_pass"].(int)
		ret.Tcp_concurrent = in["tcp_concurrent"].(int)
		ret.Dst_tcp_filter_action_whitelist = in["dst_tcp_filter_action_whitelist"].(int)
		ret.Src_tcp_filter_match = in["src_tcp_filter_match"].(int)
		ret.Src_tcp_filter_not_match = in["src_tcp_filter_not_match"].(int)
		ret.Src_tcp_filter_action_blacklist = in["src_tcp_filter_action_blacklist"].(int)
		ret.Src_tcp_filter_action_drop = in["src_tcp_filter_action_drop"].(int)
		ret.Src_tcp_filter_action_default_pass = in["src_tcp_filter_action_default_pass"].(int)
		ret.Src_tcp_filter_action_whitelist = in["src_tcp_filter_action_whitelist"].(int)
		ret.Src_dst_tcp_filter_match = in["src_dst_tcp_filter_match"].(int)
		ret.Src_dst_tcp_filter_not_match = in["src_dst_tcp_filter_not_match"].(int)
		ret.Src_dst_tcp_filter_action_blacklist = in["src_dst_tcp_filter_action_blacklist"].(int)
		ret.Src_dst_tcp_filter_action_drop = in["src_dst_tcp_filter_action_drop"].(int)
		ret.Src_dst_tcp_filter_action_default_pass = in["src_dst_tcp_filter_action_default_pass"].(int)
		ret.Src_dst_tcp_filter_action_whitelist = in["src_dst_tcp_filter_action_whitelist"].(int)
		ret.Syn_auth_pass_wl = in["syn_auth_pass_wl"].(int)
		ret.Tcp_out_of_seq_drop = in["tcp_out_of_seq_drop"].(int)
		ret.Tcp_zero_window_drop = in["tcp_zero_window_drop"].(int)
		ret.Tcp_retransmit_drop = in["tcp_retransmit_drop"].(int)
		ret.Tcp_per_conn_prate_exceed_bl = in["tcp_per_conn_prate_exceed_bl"].(int)
		ret.Tcp_any_exceed = in["tcp_any_exceed"].(int)
		ret.Tcp_drop_bl = in["tcp_drop_bl"].(int)
		ret.Tcp_frag_rcvd = in["tcp_frag_rcvd"].(int)
		ret.Tcp_frag_drop = in["tcp_frag_drop"].(int)
		ret.Tcp_auth_drop = in["tcp_auth_drop"].(int)
		ret.Tcp_auth_resp = in["tcp_auth_resp"].(int)
		ret.Tcp_total_bytes_rcv = in["tcp_total_bytes_rcv"].(int)
		ret.Tcp_total_bytes_drop = in["tcp_total_bytes_drop"].(int)
		ret.Tcp_action_on_ack_bl = in["tcp_action_on_ack_bl"].(int)
		ret.Tcp_action_on_syn_bl = in["tcp_action_on_syn_bl"].(int)
		ret.Tcp_per_conn_ofo_rate_exceed_drop = in["tcp_per_conn_ofo_rate_exceed_drop"].(int)
		ret.Tcp_per_conn_ofo_rate_exceed_bl = in["tcp_per_conn_ofo_rate_exceed_bl"].(int)
		ret.Tcp_per_conn_rexmit_rate_exceed_drop = in["tcp_per_conn_rexmit_rate_exceed_drop"].(int)
		ret.Tcp_per_conn_rexmit_rate_exceed_bl = in["tcp_per_conn_rexmit_rate_exceed_bl"].(int)
		ret.Tcp_per_conn_zwindow_rate_exceed_drop = in["tcp_per_conn_zwindow_rate_exceed_drop"].(int)
		ret.Tcp_per_conn_zwindow_rate_exceed_bl = in["tcp_per_conn_zwindow_rate_exceed_bl"].(int)
		ret.Tcp_syn_tfo_rcvd = in["tcp_syn_tfo_rcvd"].(int)
		ret.Tcp_progression_violation_exceed = in["tcp_progression_violation_exceed"].(int)
		ret.Tcp_progression_violation_exceed_bl = in["tcp_progression_violation_exceed_bl"].(int)
		ret.Tcp_progression_violation_exceed_drop = in["tcp_progression_violation_exceed_drop"].(int)
		ret.Tcp_progression_violation_exceed_reset = in["tcp_progression_violation_exceed_reset"].(int)
		ret.Tcp_auth_rst = in["tcp_auth_rst"].(int)
	}
	return ret
}

func dataToEndpointDdosL4TcpStats(d *schema.ResourceData) edpt.DdosL4TcpStats {
	var ret edpt.DdosL4TcpStats

	ret.Stats = getObjectDdosL4TcpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
