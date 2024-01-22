package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosLongStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_long_stats`: Statistics for the object long\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosLongStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_syncookie_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Sent",
						},
						"tcp_syncookie_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Passed",
						},
						"tcp_syncookie_sent_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Send Failed",
						},
						"tcp_syncookie_check_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Check Failed",
						},
						"tcp_syncookie_fail_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Cookie Blacklist Failed",
						},
						"tcp_outrst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Outbound RST",
						},
						"tcp_syn_received": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Received",
						},
						"tcp_syn_rate": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Per Sec",
						},
						"udp_exceed_drop_conn_prate": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Pkt Rate Exceeded",
						},
						"dns_malform_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Malform Drop",
						},
						"dns_qry_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Query Any Drop",
						},
						"tcp_reset_client": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Reset Client",
						},
						"tcp_reset_server": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Reset Server",
						},
						"dst_entry_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Learned",
						},
						"dst_entry_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Hit",
						},
						"src_entry_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Learned",
						},
						"src_entry_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Hit",
						},
						"sync_src_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Src Sent",
						},
						"sync_src_dst_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcDst Sent",
						},
						"sync_dst_wl_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Dst Sent",
						},
						"sync_src_wl_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Src Received",
						},
						"sync_src_dst_wl_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync SrcDst Received",
						},
						"sync_dst_wl_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Sync Dst Received",
						},
						"dst_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Pkt Rate Exceeded",
						},
						"dst_port_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Conn Limit Exceeded",
						},
						"dst_port_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Conn Rate Exceeded",
						},
						"dst_sport_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Pkt Rate Exceeded",
						},
						"dst_sport_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Conn Limit Exceeded",
						},
						"dst_sport_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Conn Rate Exceeded",
						},
						"dst_ipproto_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IP-Proto Pkt Rate Exceeded",
						},
						"tcp_ack_no_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP ACK No SYN",
						},
						"tcp_out_of_order": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Total",
						},
						"tcp_zero_window": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Total",
						},
						"tcp_retransmit": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Total",
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
						"src_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Src Entry Aged",
						},
						"dst_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Entry Aged",
						},
						"tcp_zero_wind_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Blacklisted",
						},
						"tcp_out_of_seq_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Blacklisted",
						},
						"tcp_retransmit_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Blacklisted",
						},
						"syn_auth_skip": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Auth Skipped",
						},
						"udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Passed",
						},
						"dns_auth_udp_pass": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth UDP Passed",
						},
						"udp_dst_wellknown_port_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Wellknown Port Drop",
						},
						"udp_ntp_monlist_req_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Request Dropped",
						},
						"udp_ntp_monlist_resp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Response Dropped",
						},
						"udp_payload_too_big_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large Dropped",
						},
						"udp_payload_too_small_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small Dropped",
						},
						"tcp_rexmit_syn_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit SYN Exceed Dropped",
						},
						"tcp_rexmit_syn_limit_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit SYN Exceed Blacklisted",
						},
						"tcp_exceed_drop_conn_prate": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Conn Pkt Rate Dropped",
						},
						"ip_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Received",
						},
						"ipv6_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Received",
						},
						"gre_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Received",
						},
						"gre_v6_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Received",
						},
						"dns_tcp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Force-TCP Passed",
						},
						"jumbo_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Frag Drop",
						},
						"entry_create_fail_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Create Fail Drop",
						},
						"dst_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port KiBit Rate Exceeded (KiBit)",
						},
						"dst_sport_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort KiBit Rate Exceeded (KiBit)",
						},
						"ip_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Encap",
						},
						"ip_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Encap Failed",
						},
						"ip_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Decap",
						},
						"ip_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Decap Failed",
						},
						"ip_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Tunnel Rate Limit Inner Pkts",
						},
						"ipv6_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Encap",
						},
						"ipv6_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Encap Failed",
						},
						"ipv6_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Decap",
						},
						"ipv6_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Decap Failed",
						},
						"ipv6_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Tunnel Rate Limit Inner Pkts",
						},
						"ip_gre_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Encap",
						},
						"ip_gre_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Encap Failed",
						},
						"ip_gre_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap",
						},
						"ip_gre_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap Failed",
						},
						"ip_gre_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Rate Limit Inner Pkts",
						},
						"ip_gre_tunnel_encap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Encap W/ Key",
						},
						"ip_gre_tunnel_decap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap W/ Key",
						},
						"ip_gre_tunnel_decap_key_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GRE Tunnel Decap Key Mismatch Dropped",
						},
						"ipv6_gre_tunnel_encap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Encap",
						},
						"ipv6_gre_tunnel_encap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Encap Failed",
						},
						"ipv6_gre_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap",
						},
						"ipv6_gre_tunnel_decap_fail": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap Failed",
						},
						"ipv6_gre_tunnel_rate_limit_inner": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Rate Limit Inner Pkts",
						},
						"ipv6_gre_tunnel_encap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Encap W/ Key",
						},
						"ipv6_gre_tunnel_decap_key": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap W/ Key",
						},
						"ipv6_gre_tunnel_decap_key_drop": {
							Type: schema.TypeInt, Optional: true, Description: "GRE V6 Tunnel Decap Key Mismatch Dropped",
						},
						"ip_vxlan_tunnel_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Tunnel Received",
						},
						"ip_vxlan_tunnel_invalid_vni": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Tunnel Invalid VNI",
						},
						"ip_vxlan_tunnel_decap": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Tunnel Decap",
						},
						"ip_vxlan_tunnel_decap_err": {
							Type: schema.TypeInt, Optional: true, Description: "IP VxLAN Tunnel Decap Error",
						},
						"jumbo_frag_drop_by_filter": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Fragment Filter Miss Drop",
						},
						"jumbo_frag_drop_before_slb": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Fragment Non Data Plane Drop",
						},
						"jumbo_outgoing_mtu_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Outgoing MTU Exceed Drop",
						},
						"jumbo_in_tunnel_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Packet in Tunnel Drop",
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
					},
				},
			},
		},
	}
}

func resourceDdosLongStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosLongStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosLongStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosLongStatsStats := setObjectDdosLongStatsStats(res)
		d.Set("stats", DdosLongStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosLongStatsStats(ret edpt.DataDdosLongStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"tcp_syncookie_sent":                     ret.DtDdosLongStats.Stats.Tcp_syncookie_sent,
			"tcp_syncookie_pass":                     ret.DtDdosLongStats.Stats.Tcp_syncookie_pass,
			"tcp_syncookie_sent_fail":                ret.DtDdosLongStats.Stats.Tcp_syncookie_sent_fail,
			"tcp_syncookie_check_fail":               ret.DtDdosLongStats.Stats.Tcp_syncookie_check_fail,
			"tcp_syncookie_fail_bl":                  ret.DtDdosLongStats.Stats.Tcp_syncookie_fail_bl,
			"tcp_outrst":                             ret.DtDdosLongStats.Stats.Tcp_outrst,
			"tcp_syn_received":                       ret.DtDdosLongStats.Stats.Tcp_syn_received,
			"tcp_syn_rate":                           ret.DtDdosLongStats.Stats.Tcp_syn_rate,
			"udp_exceed_drop_conn_prate":             ret.DtDdosLongStats.Stats.Udp_exceed_drop_conn_prate,
			"dns_malform_drop":                       ret.DtDdosLongStats.Stats.Dns_malform_drop,
			"dns_qry_any_drop":                       ret.DtDdosLongStats.Stats.Dns_qry_any_drop,
			"tcp_reset_client":                       ret.DtDdosLongStats.Stats.Tcp_reset_client,
			"tcp_reset_server":                       ret.DtDdosLongStats.Stats.Tcp_reset_server,
			"dst_entry_learn":                        ret.DtDdosLongStats.Stats.Dst_entry_learn,
			"dst_entry_hit":                          ret.DtDdosLongStats.Stats.Dst_entry_hit,
			"src_entry_learn":                        ret.DtDdosLongStats.Stats.Src_entry_learn,
			"src_entry_hit":                          ret.DtDdosLongStats.Stats.Src_entry_hit,
			"sync_src_wl_sent":                       ret.DtDdosLongStats.Stats.Sync_src_wl_sent,
			"sync_src_dst_wl_sent":                   ret.DtDdosLongStats.Stats.Sync_src_dst_wl_sent,
			"sync_dst_wl_sent":                       ret.DtDdosLongStats.Stats.Sync_dst_wl_sent,
			"sync_src_wl_rcv":                        ret.DtDdosLongStats.Stats.Sync_src_wl_rcv,
			"sync_src_dst_wl_rcv":                    ret.DtDdosLongStats.Stats.Sync_src_dst_wl_rcv,
			"sync_dst_wl_rcv":                        ret.DtDdosLongStats.Stats.Sync_dst_wl_rcv,
			"dst_port_pkt_rate_exceed":               ret.DtDdosLongStats.Stats.Dst_port_pkt_rate_exceed,
			"dst_port_conn_limit_exceed":             ret.DtDdosLongStats.Stats.Dst_port_conn_limit_exceed,
			"dst_port_conn_rate_exceed":              ret.DtDdosLongStats.Stats.Dst_port_conn_rate_exceed,
			"dst_sport_pkt_rate_exceed":              ret.DtDdosLongStats.Stats.Dst_sport_pkt_rate_exceed,
			"dst_sport_conn_limit_exceed":            ret.DtDdosLongStats.Stats.Dst_sport_conn_limit_exceed,
			"dst_sport_conn_rate_exceed":             ret.DtDdosLongStats.Stats.Dst_sport_conn_rate_exceed,
			"dst_ipproto_pkt_rate_exceed":            ret.DtDdosLongStats.Stats.Dst_ipproto_pkt_rate_exceed,
			"tcp_ack_no_syn":                         ret.DtDdosLongStats.Stats.Tcp_ack_no_syn,
			"tcp_out_of_order":                       ret.DtDdosLongStats.Stats.Tcp_out_of_order,
			"tcp_zero_window":                        ret.DtDdosLongStats.Stats.Tcp_zero_window,
			"tcp_retransmit":                         ret.DtDdosLongStats.Stats.Tcp_retransmit,
			"tcp_action_on_ack_start":                ret.DtDdosLongStats.Stats.Tcp_action_on_ack_start,
			"tcp_action_on_ack_matched":              ret.DtDdosLongStats.Stats.Tcp_action_on_ack_matched,
			"tcp_action_on_ack_passed":               ret.DtDdosLongStats.Stats.Tcp_action_on_ack_passed,
			"tcp_action_on_ack_failed":               ret.DtDdosLongStats.Stats.Tcp_action_on_ack_failed,
			"tcp_action_on_ack_timeout":              ret.DtDdosLongStats.Stats.Tcp_action_on_ack_timeout,
			"tcp_action_on_ack_reset":                ret.DtDdosLongStats.Stats.Tcp_action_on_ack_reset,
			"src_entry_aged":                         ret.DtDdosLongStats.Stats.Src_entry_aged,
			"dst_entry_aged":                         ret.DtDdosLongStats.Stats.Dst_entry_aged,
			"tcp_zero_wind_bl":                       ret.DtDdosLongStats.Stats.Tcp_zero_wind_bl,
			"tcp_out_of_seq_bl":                      ret.DtDdosLongStats.Stats.Tcp_out_of_seq_bl,
			"tcp_retransmit_bl":                      ret.DtDdosLongStats.Stats.Tcp_retransmit_bl,
			"syn_auth_skip":                          ret.DtDdosLongStats.Stats.Syn_auth_skip,
			"udp_retry_pass":                         ret.DtDdosLongStats.Stats.Udp_retry_pass,
			"dns_auth_udp_pass":                      ret.DtDdosLongStats.Stats.Dns_auth_udp_pass,
			"udp_dst_wellknown_port_drop":            ret.DtDdosLongStats.Stats.Udp_dst_wellknown_port_drop,
			"udp_ntp_monlist_req_drop":               ret.DtDdosLongStats.Stats.Udp_ntp_monlist_req_drop,
			"udp_ntp_monlist_resp_drop":              ret.DtDdosLongStats.Stats.Udp_ntp_monlist_resp_drop,
			"udp_payload_too_big_drop":               ret.DtDdosLongStats.Stats.Udp_payload_too_big_drop,
			"udp_payload_too_small_drop":             ret.DtDdosLongStats.Stats.Udp_payload_too_small_drop,
			"tcp_rexmit_syn_limit_drop":              ret.DtDdosLongStats.Stats.Tcp_rexmit_syn_limit_drop,
			"tcp_rexmit_syn_limit_bl":                ret.DtDdosLongStats.Stats.Tcp_rexmit_syn_limit_bl,
			"tcp_exceed_drop_conn_prate":             ret.DtDdosLongStats.Stats.Tcp_exceed_drop_conn_prate,
			"ip_tunnel_rcvd":                         ret.DtDdosLongStats.Stats.Ip_tunnel_rcvd,
			"ipv6_tunnel_rcvd":                       ret.DtDdosLongStats.Stats.Ipv6_tunnel_rcvd,
			"gre_tunnel_rcvd":                        ret.DtDdosLongStats.Stats.Gre_tunnel_rcvd,
			"gre_v6_tunnel_rcvd":                     ret.DtDdosLongStats.Stats.Gre_v6_tunnel_rcvd,
			"dns_tcp_auth_pass":                      ret.DtDdosLongStats.Stats.Dns_tcp_auth_pass,
			"jumbo_frag_drop":                        ret.DtDdosLongStats.Stats.Jumbo_frag_drop,
			"entry_create_fail_drop":                 ret.DtDdosLongStats.Stats.Entry_create_fail_drop,
			"dst_port_kbit_rate_exceed":              ret.DtDdosLongStats.Stats.Dst_port_kbit_rate_exceed,
			"dst_sport_kbit_rate_exceed":             ret.DtDdosLongStats.Stats.Dst_sport_kbit_rate_exceed,
			"ip_tunnel_encap":                        ret.DtDdosLongStats.Stats.Ip_tunnel_encap,
			"ip_tunnel_encap_fail":                   ret.DtDdosLongStats.Stats.Ip_tunnel_encap_fail,
			"ip_tunnel_decap":                        ret.DtDdosLongStats.Stats.Ip_tunnel_decap,
			"ip_tunnel_decap_fail":                   ret.DtDdosLongStats.Stats.Ip_tunnel_decap_fail,
			"ip_tunnel_rate_limit_inner":             ret.DtDdosLongStats.Stats.Ip_tunnel_rate_limit_inner,
			"ipv6_tunnel_encap":                      ret.DtDdosLongStats.Stats.Ipv6_tunnel_encap,
			"ipv6_tunnel_encap_fail":                 ret.DtDdosLongStats.Stats.Ipv6_tunnel_encap_fail,
			"ipv6_tunnel_decap":                      ret.DtDdosLongStats.Stats.Ipv6_tunnel_decap,
			"ipv6_tunnel_decap_fail":                 ret.DtDdosLongStats.Stats.Ipv6_tunnel_decap_fail,
			"ipv6_tunnel_rate_limit_inner":           ret.DtDdosLongStats.Stats.Ipv6_tunnel_rate_limit_inner,
			"ip_gre_tunnel_encap":                    ret.DtDdosLongStats.Stats.Ip_gre_tunnel_encap,
			"ip_gre_tunnel_encap_fail":               ret.DtDdosLongStats.Stats.Ip_gre_tunnel_encap_fail,
			"ip_gre_tunnel_decap":                    ret.DtDdosLongStats.Stats.Ip_gre_tunnel_decap,
			"ip_gre_tunnel_decap_fail":               ret.DtDdosLongStats.Stats.Ip_gre_tunnel_decap_fail,
			"ip_gre_tunnel_rate_limit_inner":         ret.DtDdosLongStats.Stats.Ip_gre_tunnel_rate_limit_inner,
			"ip_gre_tunnel_encap_key":                ret.DtDdosLongStats.Stats.Ip_gre_tunnel_encap_key,
			"ip_gre_tunnel_decap_key":                ret.DtDdosLongStats.Stats.Ip_gre_tunnel_decap_key,
			"ip_gre_tunnel_decap_key_drop":           ret.DtDdosLongStats.Stats.Ip_gre_tunnel_decap_key_drop,
			"ipv6_gre_tunnel_encap":                  ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_encap,
			"ipv6_gre_tunnel_encap_fail":             ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_encap_fail,
			"ipv6_gre_tunnel_decap":                  ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_decap,
			"ipv6_gre_tunnel_decap_fail":             ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_decap_fail,
			"ipv6_gre_tunnel_rate_limit_inner":       ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_rate_limit_inner,
			"ipv6_gre_tunnel_encap_key":              ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_encap_key,
			"ipv6_gre_tunnel_decap_key":              ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_decap_key,
			"ipv6_gre_tunnel_decap_key_drop":         ret.DtDdosLongStats.Stats.Ipv6_gre_tunnel_decap_key_drop,
			"ip_vxlan_tunnel_rcvd":                   ret.DtDdosLongStats.Stats.Ip_vxlan_tunnel_rcvd,
			"ip_vxlan_tunnel_invalid_vni":            ret.DtDdosLongStats.Stats.Ip_vxlan_tunnel_invalid_vni,
			"ip_vxlan_tunnel_decap":                  ret.DtDdosLongStats.Stats.Ip_vxlan_tunnel_decap,
			"ip_vxlan_tunnel_decap_err":              ret.DtDdosLongStats.Stats.Ip_vxlan_tunnel_decap_err,
			"jumbo_frag_drop_by_filter":              ret.DtDdosLongStats.Stats.Jumbo_frag_drop_by_filter,
			"jumbo_frag_drop_before_slb":             ret.DtDdosLongStats.Stats.Jumbo_frag_drop_before_slb,
			"jumbo_outgoing_mtu_exceed_drop":         ret.DtDdosLongStats.Stats.Jumbo_outgoing_mtu_exceed_drop,
			"jumbo_in_tunnel_drop":                   ret.DtDdosLongStats.Stats.Jumbo_in_tunnel_drop,
			"tcp_progression_violation_exceed":       ret.DtDdosLongStats.Stats.Tcp_progression_violation_exceed,
			"tcp_progression_violation_exceed_bl":    ret.DtDdosLongStats.Stats.Tcp_progression_violation_exceed_bl,
			"tcp_progression_violation_exceed_drop":  ret.DtDdosLongStats.Stats.Tcp_progression_violation_exceed_drop,
			"tcp_progression_violation_exceed_reset": ret.DtDdosLongStats.Stats.Tcp_progression_violation_exceed_reset,
		},
	}
}

func getObjectDdosLongStatsStats(d []interface{}) edpt.DdosLongStatsStats {

	count1 := len(d)
	var ret edpt.DdosLongStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Tcp_syncookie_sent = in["tcp_syncookie_sent"].(int)
		ret.Tcp_syncookie_pass = in["tcp_syncookie_pass"].(int)
		ret.Tcp_syncookie_sent_fail = in["tcp_syncookie_sent_fail"].(int)
		ret.Tcp_syncookie_check_fail = in["tcp_syncookie_check_fail"].(int)
		ret.Tcp_syncookie_fail_bl = in["tcp_syncookie_fail_bl"].(int)
		ret.Tcp_outrst = in["tcp_outrst"].(int)
		ret.Tcp_syn_received = in["tcp_syn_received"].(int)
		ret.Tcp_syn_rate = in["tcp_syn_rate"].(int)
		ret.Udp_exceed_drop_conn_prate = in["udp_exceed_drop_conn_prate"].(int)
		ret.Dns_malform_drop = in["dns_malform_drop"].(int)
		ret.Dns_qry_any_drop = in["dns_qry_any_drop"].(int)
		ret.Tcp_reset_client = in["tcp_reset_client"].(int)
		ret.Tcp_reset_server = in["tcp_reset_server"].(int)
		ret.Dst_entry_learn = in["dst_entry_learn"].(int)
		ret.Dst_entry_hit = in["dst_entry_hit"].(int)
		ret.Src_entry_learn = in["src_entry_learn"].(int)
		ret.Src_entry_hit = in["src_entry_hit"].(int)
		ret.Sync_src_wl_sent = in["sync_src_wl_sent"].(int)
		ret.Sync_src_dst_wl_sent = in["sync_src_dst_wl_sent"].(int)
		ret.Sync_dst_wl_sent = in["sync_dst_wl_sent"].(int)
		ret.Sync_src_wl_rcv = in["sync_src_wl_rcv"].(int)
		ret.Sync_src_dst_wl_rcv = in["sync_src_dst_wl_rcv"].(int)
		ret.Sync_dst_wl_rcv = in["sync_dst_wl_rcv"].(int)
		ret.Dst_port_pkt_rate_exceed = in["dst_port_pkt_rate_exceed"].(int)
		ret.Dst_port_conn_limit_exceed = in["dst_port_conn_limit_exceed"].(int)
		ret.Dst_port_conn_rate_exceed = in["dst_port_conn_rate_exceed"].(int)
		ret.Dst_sport_pkt_rate_exceed = in["dst_sport_pkt_rate_exceed"].(int)
		ret.Dst_sport_conn_limit_exceed = in["dst_sport_conn_limit_exceed"].(int)
		ret.Dst_sport_conn_rate_exceed = in["dst_sport_conn_rate_exceed"].(int)
		ret.Dst_ipproto_pkt_rate_exceed = in["dst_ipproto_pkt_rate_exceed"].(int)
		ret.Tcp_ack_no_syn = in["tcp_ack_no_syn"].(int)
		ret.Tcp_out_of_order = in["tcp_out_of_order"].(int)
		ret.Tcp_zero_window = in["tcp_zero_window"].(int)
		ret.Tcp_retransmit = in["tcp_retransmit"].(int)
		ret.Tcp_action_on_ack_start = in["tcp_action_on_ack_start"].(int)
		ret.Tcp_action_on_ack_matched = in["tcp_action_on_ack_matched"].(int)
		ret.Tcp_action_on_ack_passed = in["tcp_action_on_ack_passed"].(int)
		ret.Tcp_action_on_ack_failed = in["tcp_action_on_ack_failed"].(int)
		ret.Tcp_action_on_ack_timeout = in["tcp_action_on_ack_timeout"].(int)
		ret.Tcp_action_on_ack_reset = in["tcp_action_on_ack_reset"].(int)
		ret.Src_entry_aged = in["src_entry_aged"].(int)
		ret.Dst_entry_aged = in["dst_entry_aged"].(int)
		ret.Tcp_zero_wind_bl = in["tcp_zero_wind_bl"].(int)
		ret.Tcp_out_of_seq_bl = in["tcp_out_of_seq_bl"].(int)
		ret.Tcp_retransmit_bl = in["tcp_retransmit_bl"].(int)
		ret.Syn_auth_skip = in["syn_auth_skip"].(int)
		ret.Udp_retry_pass = in["udp_retry_pass"].(int)
		ret.Dns_auth_udp_pass = in["dns_auth_udp_pass"].(int)
		ret.Udp_dst_wellknown_port_drop = in["udp_dst_wellknown_port_drop"].(int)
		ret.Udp_ntp_monlist_req_drop = in["udp_ntp_monlist_req_drop"].(int)
		ret.Udp_ntp_monlist_resp_drop = in["udp_ntp_monlist_resp_drop"].(int)
		ret.Udp_payload_too_big_drop = in["udp_payload_too_big_drop"].(int)
		ret.Udp_payload_too_small_drop = in["udp_payload_too_small_drop"].(int)
		ret.Tcp_rexmit_syn_limit_drop = in["tcp_rexmit_syn_limit_drop"].(int)
		ret.Tcp_rexmit_syn_limit_bl = in["tcp_rexmit_syn_limit_bl"].(int)
		ret.Tcp_exceed_drop_conn_prate = in["tcp_exceed_drop_conn_prate"].(int)
		ret.Ip_tunnel_rcvd = in["ip_tunnel_rcvd"].(int)
		ret.Ipv6_tunnel_rcvd = in["ipv6_tunnel_rcvd"].(int)
		ret.Gre_tunnel_rcvd = in["gre_tunnel_rcvd"].(int)
		ret.Gre_v6_tunnel_rcvd = in["gre_v6_tunnel_rcvd"].(int)
		ret.Dns_tcp_auth_pass = in["dns_tcp_auth_pass"].(int)
		ret.Jumbo_frag_drop = in["jumbo_frag_drop"].(int)
		ret.Entry_create_fail_drop = in["entry_create_fail_drop"].(int)
		ret.Dst_port_kbit_rate_exceed = in["dst_port_kbit_rate_exceed"].(int)
		ret.Dst_sport_kbit_rate_exceed = in["dst_sport_kbit_rate_exceed"].(int)
		ret.Ip_tunnel_encap = in["ip_tunnel_encap"].(int)
		ret.Ip_tunnel_encap_fail = in["ip_tunnel_encap_fail"].(int)
		ret.Ip_tunnel_decap = in["ip_tunnel_decap"].(int)
		ret.Ip_tunnel_decap_fail = in["ip_tunnel_decap_fail"].(int)
		ret.Ip_tunnel_rate_limit_inner = in["ip_tunnel_rate_limit_inner"].(int)
		ret.Ipv6_tunnel_encap = in["ipv6_tunnel_encap"].(int)
		ret.Ipv6_tunnel_encap_fail = in["ipv6_tunnel_encap_fail"].(int)
		ret.Ipv6_tunnel_decap = in["ipv6_tunnel_decap"].(int)
		ret.Ipv6_tunnel_decap_fail = in["ipv6_tunnel_decap_fail"].(int)
		ret.Ipv6_tunnel_rate_limit_inner = in["ipv6_tunnel_rate_limit_inner"].(int)
		ret.Ip_gre_tunnel_encap = in["ip_gre_tunnel_encap"].(int)
		ret.Ip_gre_tunnel_encap_fail = in["ip_gre_tunnel_encap_fail"].(int)
		ret.Ip_gre_tunnel_decap = in["ip_gre_tunnel_decap"].(int)
		ret.Ip_gre_tunnel_decap_fail = in["ip_gre_tunnel_decap_fail"].(int)
		ret.Ip_gre_tunnel_rate_limit_inner = in["ip_gre_tunnel_rate_limit_inner"].(int)
		ret.Ip_gre_tunnel_encap_key = in["ip_gre_tunnel_encap_key"].(int)
		ret.Ip_gre_tunnel_decap_key = in["ip_gre_tunnel_decap_key"].(int)
		ret.Ip_gre_tunnel_decap_key_drop = in["ip_gre_tunnel_decap_key_drop"].(int)
		ret.Ipv6_gre_tunnel_encap = in["ipv6_gre_tunnel_encap"].(int)
		ret.Ipv6_gre_tunnel_encap_fail = in["ipv6_gre_tunnel_encap_fail"].(int)
		ret.Ipv6_gre_tunnel_decap = in["ipv6_gre_tunnel_decap"].(int)
		ret.Ipv6_gre_tunnel_decap_fail = in["ipv6_gre_tunnel_decap_fail"].(int)
		ret.Ipv6_gre_tunnel_rate_limit_inner = in["ipv6_gre_tunnel_rate_limit_inner"].(int)
		ret.Ipv6_gre_tunnel_encap_key = in["ipv6_gre_tunnel_encap_key"].(int)
		ret.Ipv6_gre_tunnel_decap_key = in["ipv6_gre_tunnel_decap_key"].(int)
		ret.Ipv6_gre_tunnel_decap_key_drop = in["ipv6_gre_tunnel_decap_key_drop"].(int)
		ret.Ip_vxlan_tunnel_rcvd = in["ip_vxlan_tunnel_rcvd"].(int)
		ret.Ip_vxlan_tunnel_invalid_vni = in["ip_vxlan_tunnel_invalid_vni"].(int)
		ret.Ip_vxlan_tunnel_decap = in["ip_vxlan_tunnel_decap"].(int)
		ret.Ip_vxlan_tunnel_decap_err = in["ip_vxlan_tunnel_decap_err"].(int)
		ret.Jumbo_frag_drop_by_filter = in["jumbo_frag_drop_by_filter"].(int)
		ret.Jumbo_frag_drop_before_slb = in["jumbo_frag_drop_before_slb"].(int)
		ret.Jumbo_outgoing_mtu_exceed_drop = in["jumbo_outgoing_mtu_exceed_drop"].(int)
		ret.Jumbo_in_tunnel_drop = in["jumbo_in_tunnel_drop"].(int)
		ret.Tcp_progression_violation_exceed = in["tcp_progression_violation_exceed"].(int)
		ret.Tcp_progression_violation_exceed_bl = in["tcp_progression_violation_exceed_bl"].(int)
		ret.Tcp_progression_violation_exceed_drop = in["tcp_progression_violation_exceed_drop"].(int)
		ret.Tcp_progression_violation_exceed_reset = in["tcp_progression_violation_exceed_reset"].(int)
	}
	return ret
}

func dataToEndpointDdosLongStats(d *schema.ResourceData) edpt.DdosLongStats {
	var ret edpt.DdosLongStats

	ret.Stats = getObjectDdosLongStatsStats(d.Get("stats").([]interface{}))
	return ret
}
