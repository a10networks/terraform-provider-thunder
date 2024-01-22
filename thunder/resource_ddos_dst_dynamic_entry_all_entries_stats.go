package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDynamicEntryAllEntriesStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_dynamic_entry_all_entries_stats`: Statistics for the object all-entries\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstDynamicEntryAllEntriesStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_tcp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Rate: Total Exceeded",
						},
						"dst_tcp_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Rate: Packet Exceeded",
						},
						"dst_tcp_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Rate: Conn Exceeded",
						},
						"dst_udp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst L4-Type Rate: Total Exceeded",
						},
						"dst_udp_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst L4-Type Rate: Packet Exceeded",
						},
						"dst_udp_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst L4-Type Limit: Conn Exceeded",
						},
						"dst_udp_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst L4-Type Rate: Conn Exceeded",
						},
						"dst_icmp_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst Rate: Packet Exceeded",
						},
						"dst_other_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst L4-Type Rate: Packet Exceeded",
						},
						"dst_other_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst L4-Type Rate: Frag Exceeded",
						},
						"dst_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: Packet Exceeded",
						},
						"dst_port_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Limit: Conn Exceeded",
						},
						"dst_port_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: Conn Exceeded",
						},
						"dst_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Packets Forwarded",
						},
						"dst_udp_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Forwarded",
						},
						"dst_tcp_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Forwarded",
						},
						"dst_icmp_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Forwarded",
						},
						"dst_other_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Forwarded",
						},
						"dst_tcp_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Limit: Conn Exceeded",
						},
						"dst_tcp_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Received",
						},
						"dst_udp_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Received",
						},
						"dst_icmp_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Received",
						},
						"dst_other_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Received",
						},
						"dst_udp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Match",
						},
						"dst_udp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Not Matched on Pkt",
						},
						"dst_udp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action Blacklist",
						},
						"dst_udp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action Drop",
						},
						"dst_tcp_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total SYN Received",
						},
						"dst_tcp_syn_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Packets Dropped",
						},
						"dst_tcp_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Rate: Total Exceeded",
						},
						"dst_udp_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Rate: Total Exceeded",
						},
						"dst_icmp_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Rate: Total Exceeded",
						},
						"dst_other_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Rate: Frag Exceeded",
						},
						"dst_other_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Rate: Total Exceeded",
						},
						"dst_tcp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Dropped",
						},
						"dst_udp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Dropped",
						},
						"dst_icmp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Dropped",
						},
						"dst_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Dropped",
						},
						"dst_other_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Dropped",
						},
						"dst_tcp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Cookie Sent",
						},
						"dst_udp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action Default Pass",
						},
						"dst_tcp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Match",
						},
						"dst_tcp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Not Matched on Pkt",
						},
						"dst_tcp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action Blacklist",
						},
						"dst_tcp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action Drop",
						},
						"dst_tcp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action Default Pass",
						},
						"dst_udp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action WL",
						},
						"dst_udp_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst L4-Type Rate: KiBit Exceeded",
						},
						"dst_tcp_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Rate: KiBit Exceeded",
						},
						"dst_icmp_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst Rate: KiBit Exceeded",
						},
						"dst_other_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst L4-Type Rate: KiBit Exceeded",
						},
						"dst_port_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Undefined Dropped",
						},
						"dst_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Blacklist Packets Dropped",
						},
						"dst_src_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Blacklist Packets Dropped",
						},
						"dst_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: KiBit Exceeded",
						},
						"dst_tcp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Packets Dropped",
						},
						"dst_udp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Packets Dropped",
						},
						"dst_icmp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Packets Dropped",
						},
						"dst_other_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Packets Dropped",
						},
						"tcp_syn_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Inbound SYN Received",
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
						"ingress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Bytes Received",
						},
						"egress_bytes": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Bytes Received",
						},
						"ingress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Packets Received",
						},
						"egress_packets": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Packets Received",
						},
						"tcp_fwd_recv": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Inbound Packets Received",
						},
						"udp_fwd_recv": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Inbound Packets Received",
						},
						"icmp_fwd_recv": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Inbound Packets Received",
						},
						"tcp_syn_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Cookie Failed",
						},
						"dst_tcp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Sessions Created",
						},
						"dst_udp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Sessions Created",
						},
						"dst_tcp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action WL",
						},
						"dst_other_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Match",
						},
						"dst_other_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Not Matched on Pkt",
						},
						"dst_other_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action Blacklist",
						},
						"dst_other_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action Drop",
						},
						"dst_other_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action WL",
						},
						"dst_other_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action Default Pass",
						},
						"dst_blackhole_inject": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Inject",
						},
						"dst_blackhole_withdraw": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Withdraw",
						},
						"dst_tcp_out_of_seq_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Exceeded",
						},
						"dst_tcp_retransmit_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Exceeded",
						},
						"dst_tcp_zero_window_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Exceeded",
						},
						"dst_tcp_conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Rate: Conn Pkt Exceeded",
						},
						"dst_tcp_action_on_ack_init": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Init",
						},
						"dst_tcp_action_on_ack_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Retry-Gap Dropped",
						},
						"dst_tcp_action_on_ack_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Dropped",
						},
						"dst_tcp_action_on_ack_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Passed",
						},
						"dst_tcp_action_on_syn_init": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Init",
						},
						"dst_tcp_action_on_syn_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry-Gap Dropped",
						},
						"dst_tcp_action_on_syn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Dropped",
						},
						"dst_tcp_action_on_syn_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Passed",
						},
						"udp_payload_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small",
						},
						"udp_payload_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large",
						},
						"dst_udp_conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Rate: Conn Pkt Exceeded",
						},
						"dst_udp_ntp_monlist_req": {
							Type: schema.TypeInt, Optional: true, Description: "UDP NTP Monlist Request",
						},
						"dst_udp_ntp_monlist_resp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP NTP Monlist Response",
						},
						"dst_udp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcPort Wellknown",
						},
						"dst_udp_retry_init": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry Init",
						},
						"dst_udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry Passed",
						},
						"dst_tcp_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Dropped",
						},
						"dst_udp_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Dropped",
						},
						"dst_icmp_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Dropped",
						},
						"dst_other_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Dropped",
						},
						"dst_out_no_route": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IPv4/v6 Out No Route",
						},
						"outbound_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Bytes Forwarded",
						},
						"outbound_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Packets Dropped",
						},
						"outbound_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Bytes Dropped",
						},
						"outbound_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Packets Forwarded",
						},
						"inbound_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Bytes Forwarded",
						},
						"inbound_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Bytes Dropped",
						},
						"dst_src_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Rate: Packet Exceeded",
						},
						"dst_src_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Rate: KiBit Exceeded",
						},
						"dst_src_port_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Limit: Conn Exceeded",
						},
						"dst_src_port_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Rate: Conn Exceeded",
						},
						"dst_ip_proto_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Proto Rate: Packet Exceeded",
						},
						"dst_ip_proto_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Proto Rate: KiBit Exceeded",
						},
						"dst_tcp_port_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Rate: Total Exceed",
						},
						"dst_udp_port_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Rate: Total Exceed",
						},
						"dst_tcp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Auth Passed",
						},
						"dst_tcp_rst_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: RST Cookie Failed",
						},
						"dst_tcp_unauth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: Unauth Dropped",
						},
						"src_tcp_syn_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Auth Failed",
						},
						"src_tcp_syn_cookie_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Cookie Sent",
						},
						"src_tcp_syn_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Cookie Failed",
						},
						"src_tcp_rst_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: RST Cookie Failed",
						},
						"src_tcp_unauth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: Unauth Dropped",
						},
						"src_tcp_action_on_syn_init": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Retry Init",
						},
						"src_tcp_action_on_syn_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Retry-Gap Dropped",
						},
						"src_tcp_action_on_syn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Retry Dropped",
						},
						"src_tcp_action_on_ack_init": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: ACK Retry Init",
						},
						"src_tcp_action_on_ack_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: ACK Retry Retry-Gap Dropped",
						},
						"src_tcp_action_on_ack_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: ACK Retry Dropped",
						},
						"src_tcp_out_of_seq_excd": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Out-Of-Seq Exceeded",
						},
						"src_tcp_retransmit_excd": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Retransmit Exceeded",
						},
						"src_tcp_zero_window_excd": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Zero-Window Exceeded",
						},
						"src_tcp_conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Rate: Conn Pkt Exceeded",
						},
						"src_udp_min_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Payload Too Small",
						},
						"src_udp_max_payload": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Payload Too Large",
						},
						"src_udp_conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Rate: Conn Pkt Exceeded",
						},
						"src_udp_ntp_monlist_req": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP NTP Monlist Request",
						},
						"src_udp_ntp_monlist_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP NTP Monlist Response",
						},
						"src_udp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP SrcPort Wellknown",
						},
						"src_udp_retry_init": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth: Retry Init",
						},
						"dst_udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry-Gap Dropped",
						},
						"dst_udp_retry_fail": {
							Type: schema.TypeInt, Optional: true, Description: "UDP P Sessions Aged",
						},
						"dst_tcp_session_aged": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Sessions Aged",
						},
						"dst_udp_session_aged": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Sessions Aged",
						},
						"dst_tcp_conn_close": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Connections Closed",
						},
						"dst_tcp_conn_close_half_open": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Half Open Connections Closed",
						},
						"dst_l4_tcp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Auth: SYN Cookie Sent",
						},
						"tcp_l4_syn_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Auth: SYN Cookie Failed",
						},
						"tcp_l4_rst_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Auth: RST Cookie Failed",
						},
						"tcp_l4_unauth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Auth: Unauth Dropped",
						},
						"src_tcp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Filter Action Blacklist",
						},
						"src_tcp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Filter Action WL",
						},
						"src_tcp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Filter Action Drop",
						},
						"src_tcp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Filter Action Default Pass",
						},
						"src_udp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Filter Action Blacklist",
						},
						"src_udp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Filter Action WL",
						},
						"src_udp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Filter Action Drop",
						},
						"src_udp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Filter Action Default Pass",
						},
						"src_other_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src OTHER Filter Action Blacklist",
						},
						"src_other_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Src OTHER Filter Action WL",
						},
						"src_other_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src OTHER Filter Action Drop",
						},
						"src_other_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src OTHER Filter Action Default Pass",
						},
						"tcp_invalid_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Invalid SYN Received",
						},
						"dst_tcp_conn_close_w_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP RST Connections Closed",
						},
						"dst_tcp_conn_close_w_fin": {
							Type: schema.TypeInt, Optional: true, Description: "TCP FIN Connections Closed",
						},
						"dst_tcp_conn_close_w_idle": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Idle Connections Closed",
						},
						"dst_tcp_conn_create_from_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Connections Created From SYN",
						},
						"dst_tcp_conn_create_from_ack": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Connections Created From ACK",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
						},
						"dst_l4_tcp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst L4-type TCP Blacklist Dropped",
						},
						"dst_l4_udp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst L4-type UDP Blacklist Dropped",
						},
						"dst_l4_icmp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "No Policy Class-list Match",
						},
						"dst_l4_other_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst L4-type OTHER Blacklist Dropped",
						},
						"src_l4_tcp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src L4-type TCP Blacklist Dropped",
						},
						"src_l4_udp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src L4-type UDP Blacklist Dropped",
						},
						"src_l4_icmp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src L4-type ICMP Blacklist Dropped",
						},
						"src_l4_other_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src L4-type OTHER Blacklist Dropped",
						},
						"dst_port_kbit_rate_exceed_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: KiBit Pkt Exceeded",
						},
						"dst_tcp_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Received",
						},
						"dst_udp_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Received",
						},
						"dst_icmp_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Received",
						},
						"dst_other_bytes_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Received",
						},
						"dst_tcp_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Forwarded",
						},
						"dst_udp_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Forwarded",
						},
						"dst_icmp_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Forwarded",
						},
						"dst_other_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Forwarded",
						},
						"dst_udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Dropped",
						},
						"dst_tcp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: Dropped",
						},
						"dst_tcp_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: Responded",
						},
						"inbound_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Packets Dropped",
						},
						"dst_entry_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Rate: Packet Exceeded",
						},
						"dst_entry_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Rate: KiBit Exceeded",
						},
						"dst_entry_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Limit: Conn Exceeded",
						},
						"dst_entry_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Rate: Conn Exceeded",
						},
						"dst_entry_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Rate: Frag Packet Exceeded",
						},
						"dst_icmp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Rate: Total Exceed",
						},
						"dst_other_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Rate: Total Exceed",
						},
						"src_dst_pair_entry_total": {
							Type: schema.TypeInt, Optional: true, Description: "Src-Dst Pair Entry Total Count",
						},
						"src_dst_pair_entry_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Src-Dst Pair Entry UDP Count",
						},
						"src_dst_pair_entry_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Src-Dst Pair Entry TCP Count",
						},
						"src_dst_pair_entry_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Src-Dst Pair Entry ICMP Count",
						},
						"src_dst_pair_entry_other": {
							Type: schema.TypeInt, Optional: true, Description: "Src-Dst Pair Entry OTHER Count",
						},
						"dst_clist_overflow_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Src-Based Overflow Policy Hit",
						},
						"tcp_rexmit_syn_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retransmit Exceeded Drop",
						},
						"tcp_rexmit_syn_limit_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retransmit Exceeded Blacklist",
						},
						"dst_tcp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SrcPort Wellknown",
						},
						"src_tcp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP SrcPort Wellknown",
						},
						"dst_frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Received",
						},
						"no_policy_class_list_match": {
							Type: schema.TypeInt, Optional: true, Description: "No Policy Class-list Match",
						},
						"src_udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth: Retry-Gap Dropped",
						},
						"dst_entry_kbit_rate_exceed_count": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Rate: KiBit Exceeded Count",
						},
						"dst_port_undef_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Undefined Hit",
						},
						"dst_tcp_action_on_ack_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Timeout",
						},
						"dst_tcp_action_on_ack_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Timeout Reset",
						},
						"dst_tcp_action_on_ack_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Timeout Blacklisted",
						},
						"src_tcp_action_on_ack_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: ACK Retry Timeout",
						},
						"src_tcp_action_on_ack_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: ACK Retry Timeout Reset",
						},
						"src_tcp_action_on_ack_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: ACK Retry Timeout Blacklisted",
						},
						"dst_tcp_action_on_syn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Timeout",
						},
						"dst_tcp_action_on_syn_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Timeout Reset",
						},
						"dst_tcp_action_on_syn_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Timeout Blacklisted",
						},
						"src_tcp_action_on_syn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Retry Timeout",
						},
						"src_tcp_action_on_syn_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Retry Timeout Reset",
						},
						"src_tcp_action_on_syn_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth: SYN Retry Timeout Blacklisted",
						},
						"dst_udp_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst L4-Type Rate: Frag Exceeded",
						},
						"dst_udp_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Rate: Frag Exceeded",
						},
						"dst_tcp_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst L4-Type Rate: Frag Exceeded",
						},
						"dst_tcp_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Rate: Frag Exceeded",
						},
						"dst_icmp_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst L4-Type Rate: Frag Exceeded",
						},
						"dst_icmp_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Rate: Frag Exceeded",
						},
						"sflow_internal_samples_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow Internal Samples Packed",
						},
						"sflow_external_samples_packed": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow External Samples Packed",
						},
						"sflow_internal_packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow Internal Packets Sent",
						},
						"sflow_external_packets_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Sflow External Packets Sent",
						},
						"dns_outbound_total_query": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Total Query",
						},
						"dns_outbound_query_malformed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Malformed",
						},
						"dns_outbound_query_resp_chk_failed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Resp Check Failed",
						},
						"dns_outbound_query_resp_chk_blacklisted": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Resp Check Blacklisted",
						},
						"dns_outbound_query_resp_chk_refused_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Resp Check REFUSED Sent",
						},
						"dns_outbound_query_resp_chk_reset_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Resp Check RESET Sent",
						},
						"dns_outbound_query_resp_chk_no_resp_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Resp Check No Response Sent",
						},
						"dns_outbound_query_resp_size_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Response Size Exceed",
						},
						"dns_outbound_query_sess_timed_out": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Outbound Query Session Timed Out",
						},
						"dst_exceed_action_tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Exceed Action: Tunnel",
						},
						"src_udp_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth: Retry Timeout",
						},
						"src_udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Passed",
						},
						"dst_hw_drop_rule_insert": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Drop Rules Inserted",
						},
						"dst_hw_drop_rule_remove": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Drop Rules Removed",
						},
						"src_hw_drop_rule_insert": {
							Type: schema.TypeInt, Optional: true, Description: "Src Hardware Drop Rules Inserted",
						},
						"src_hw_drop_rule_remove": {
							Type: schema.TypeInt, Optional: true, Description: "Src Hardware Drop Rules Removed",
						},
						"prog_first_req_time_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: First Request Time Exceed",
						},
						"prog_req_resp_time_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Request to Response Time Exceed",
						},
						"prog_request_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Request Length Exceed",
						},
						"prog_response_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Response Length Exceed",
						},
						"prog_resp_req_ratio_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Response to Request Ratio Exceed",
						},
						"prog_resp_req_time_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Response to Request Time Exceed",
						},
						"entry_sync_message_received": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Sync Message Received",
						},
						"entry_sync_message_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Sync Message Sent",
						},
						"prog_conn_sent_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Connection: Sent Exceed",
						},
						"prog_conn_rcvd_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Connection: Received Exceed",
						},
						"prog_conn_time_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Connection: Time Exceed",
						},
						"prog_conn_rcvd_sent_ratio_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Connection: Reveived to Sent Ratio Exceed",
						},
						"prog_win_sent_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Time Window: Sent Exceed",
						},
						"prog_win_rcvd_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Time Window: Received Exceed",
						},
						"prog_win_rcvd_sent_ratio_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Time Window: Received to Sent Exceed",
						},
						"prog_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Violation Exceed Dropped",
						},
						"prog_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Violation Exceed Blacklisted",
						},
						"prog_conn_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Connection: Violation Exceed Dropped",
						},
						"prog_conn_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Connection: Violation Exceed Blacklisted",
						},
						"prog_win_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Time Window: Violation Exceed Dropped",
						},
						"prog_win_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Time Window: Violation Exceed Blacklisted",
						},
						"dst_exceed_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Exceed Action: Dropped",
						},
						"src_hw_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Hardware Packets Dropped",
						},
						"dst_tcp_auth_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: Reset",
						},
						"dst_src_learn_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "Src Dynamic Entry Count Overflow",
						},
						"tcp_fwd_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Inbound Packets Forwarded",
						},
						"udp_fwd_sent": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Inbound Packets Forwarded",
						},
					},
				},
			},
		},
	}
}

func resourceDdosDstDynamicEntryAllEntriesStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryAllEntriesStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryAllEntriesStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstDynamicEntryAllEntriesStatsStats := setObjectDdosDstDynamicEntryAllEntriesStatsStats(res)
		d.Set("stats", DdosDstDynamicEntryAllEntriesStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstDynamicEntryAllEntriesStatsStats(ret edpt.DataDdosDstDynamicEntryAllEntriesStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"dst_tcp_any_exceed":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_any_exceed,
			"dst_tcp_pkt_rate_exceed":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_pkt_rate_exceed,
			"dst_tcp_conn_rate_exceed":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_rate_exceed,
			"dst_udp_any_exceed":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_any_exceed,
			"dst_udp_pkt_rate_exceed":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_pkt_rate_exceed,
			"dst_udp_conn_limit_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_conn_limit_exceed,
			"dst_udp_conn_rate_exceed":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_conn_rate_exceed,
			"dst_icmp_pkt_rate_exceed":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_pkt_rate_exceed,
			"dst_other_pkt_rate_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_pkt_rate_exceed,
			"dst_other_frag_pkt_rate_exceed":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_frag_pkt_rate_exceed,
			"dst_port_pkt_rate_exceed":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_pkt_rate_exceed,
			"dst_port_conn_limit_exceed":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_conn_limit_exceed,
			"dst_port_conn_rate_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_conn_rate_exceed,
			"dst_pkt_sent":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_pkt_sent,
			"dst_udp_pkt_sent":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_pkt_sent,
			"dst_tcp_pkt_sent":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_pkt_sent,
			"dst_icmp_pkt_sent":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_pkt_sent,
			"dst_other_pkt_sent":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_pkt_sent,
			"dst_tcp_conn_limit_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_limit_exceed,
			"dst_tcp_pkt_rcvd":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_pkt_rcvd,
			"dst_udp_pkt_rcvd":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_pkt_rcvd,
			"dst_icmp_pkt_rcvd":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_pkt_rcvd,
			"dst_other_pkt_rcvd":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_pkt_rcvd,
			"dst_udp_filter_match":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_filter_match,
			"dst_udp_filter_not_match":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_filter_not_match,
			"dst_udp_filter_action_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_filter_action_blacklist,
			"dst_udp_filter_action_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_filter_action_drop,
			"dst_tcp_syn":                              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_syn,
			"dst_tcp_syn_drop":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_syn_drop,
			"dst_tcp_src_rate_drop":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_src_rate_drop,
			"dst_udp_src_rate_drop":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_src_rate_drop,
			"dst_icmp_src_rate_drop":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_src_rate_drop,
			"dst_other_frag_src_rate_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_frag_src_rate_drop,
			"dst_other_src_rate_drop":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_src_rate_drop,
			"dst_tcp_drop":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_drop,
			"dst_udp_drop":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_drop,
			"dst_icmp_drop":                            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_drop,
			"dst_frag_drop":                            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_frag_drop,
			"dst_other_drop":                           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_drop,
			"dst_tcp_auth":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_auth,
			"dst_udp_filter_action_default_pass":       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_filter_action_default_pass,
			"dst_tcp_filter_match":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_filter_match,
			"dst_tcp_filter_not_match":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_filter_not_match,
			"dst_tcp_filter_action_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_filter_action_blacklist,
			"dst_tcp_filter_action_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_filter_action_drop,
			"dst_tcp_filter_action_default_pass":       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_filter_action_default_pass,
			"dst_udp_filter_action_whitelist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_filter_action_whitelist,
			"dst_udp_kibit_rate_drop":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_kibit_rate_drop,
			"dst_tcp_kibit_rate_drop":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_kibit_rate_drop,
			"dst_icmp_kibit_rate_drop":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_kibit_rate_drop,
			"dst_other_kibit_rate_drop":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_kibit_rate_drop,
			"dst_port_undef_drop":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_undef_drop,
			"dst_port_bl":                              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_bl,
			"dst_src_port_bl":                          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_src_port_bl,
			"dst_port_kbit_rate_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_kbit_rate_exceed,
			"dst_tcp_src_drop":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_src_drop,
			"dst_udp_src_drop":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_src_drop,
			"dst_icmp_src_drop":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_src_drop,
			"dst_other_src_drop":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_src_drop,
			"tcp_syn_rcvd":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_syn_rcvd,
			"tcp_syn_ack_rcvd":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_syn_ack_rcvd,
			"tcp_ack_rcvd":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_ack_rcvd,
			"tcp_fin_rcvd":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_fin_rcvd,
			"tcp_rst_rcvd":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_rst_rcvd,
			"ingress_bytes":                            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Ingress_bytes,
			"egress_bytes":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Egress_bytes,
			"ingress_packets":                          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Ingress_packets,
			"egress_packets":                           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Egress_packets,
			"tcp_fwd_recv":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_fwd_recv,
			"udp_fwd_recv":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Udp_fwd_recv,
			"icmp_fwd_recv":                            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Icmp_fwd_recv,
			"tcp_syn_cookie_fail":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_syn_cookie_fail,
			"dst_tcp_session_created":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_session_created,
			"dst_udp_session_created":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_session_created,
			"dst_tcp_filter_action_whitelist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_filter_action_whitelist,
			"dst_other_filter_match":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_filter_match,
			"dst_other_filter_not_match":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_filter_not_match,
			"dst_other_filter_action_blacklist":        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_filter_action_blacklist,
			"dst_other_filter_action_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_filter_action_drop,
			"dst_other_filter_action_whitelist":        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_filter_action_whitelist,
			"dst_other_filter_action_default_pass":     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_filter_action_default_pass,
			"dst_blackhole_inject":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_blackhole_inject,
			"dst_blackhole_withdraw":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_blackhole_withdraw,
			"dst_tcp_out_of_seq_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_out_of_seq_excd,
			"dst_tcp_retransmit_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_retransmit_excd,
			"dst_tcp_zero_window_excd":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_zero_window_excd,
			"dst_tcp_conn_prate_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_prate_excd,
			"dst_tcp_action_on_ack_init":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_init,
			"dst_tcp_action_on_ack_gap_drop":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_gap_drop,
			"dst_tcp_action_on_ack_fail":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_fail,
			"dst_tcp_action_on_ack_pass":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_pass,
			"dst_tcp_action_on_syn_init":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_init,
			"dst_tcp_action_on_syn_gap_drop":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_gap_drop,
			"dst_tcp_action_on_syn_fail":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_fail,
			"dst_tcp_action_on_syn_pass":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_pass,
			"udp_payload_too_small":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Udp_payload_too_small,
			"udp_payload_too_big":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Udp_payload_too_big,
			"dst_udp_conn_prate_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_conn_prate_excd,
			"dst_udp_ntp_monlist_req":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_ntp_monlist_req,
			"dst_udp_ntp_monlist_resp":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_ntp_monlist_resp,
			"dst_udp_wellknown_sport_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_wellknown_sport_drop,
			"dst_udp_retry_init":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_retry_init,
			"dst_udp_retry_pass":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_retry_pass,
			"dst_tcp_bytes_drop":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_bytes_drop,
			"dst_udp_bytes_drop":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_bytes_drop,
			"dst_icmp_bytes_drop":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_bytes_drop,
			"dst_other_bytes_drop":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_bytes_drop,
			"dst_out_no_route":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_out_no_route,
			"outbound_bytes_sent":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Outbound_bytes_sent,
			"outbound_pkt_drop":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Outbound_pkt_drop,
			"outbound_bytes_drop":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Outbound_bytes_drop,
			"outbound_pkt_sent":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Outbound_pkt_sent,
			"inbound_bytes_sent":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Inbound_bytes_sent,
			"inbound_bytes_drop":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Inbound_bytes_drop,
			"dst_src_port_pkt_rate_exceed":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_src_port_pkt_rate_exceed,
			"dst_src_port_kbit_rate_exceed":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_src_port_kbit_rate_exceed,
			"dst_src_port_conn_limit_exceed":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_src_port_conn_limit_exceed,
			"dst_src_port_conn_rate_exceed":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_src_port_conn_rate_exceed,
			"dst_ip_proto_pkt_rate_exceed":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_ip_proto_pkt_rate_exceed,
			"dst_ip_proto_kbit_rate_exceed":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_ip_proto_kbit_rate_exceed,
			"dst_tcp_port_any_exceed":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_port_any_exceed,
			"dst_udp_port_any_exceed":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_port_any_exceed,
			"dst_tcp_auth_pass":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_auth_pass,
			"dst_tcp_rst_cookie_fail":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_rst_cookie_fail,
			"dst_tcp_unauth_drop":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_unauth_drop,
			"src_tcp_syn_auth_fail":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_syn_auth_fail,
			"src_tcp_syn_cookie_sent":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_syn_cookie_sent,
			"src_tcp_syn_cookie_fail":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_syn_cookie_fail,
			"src_tcp_rst_cookie_fail":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_rst_cookie_fail,
			"src_tcp_unauth_drop":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_unauth_drop,
			"src_tcp_action_on_syn_init":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_syn_init,
			"src_tcp_action_on_syn_gap_drop":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_syn_gap_drop,
			"src_tcp_action_on_syn_fail":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_syn_fail,
			"src_tcp_action_on_ack_init":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_ack_init,
			"src_tcp_action_on_ack_gap_drop":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_ack_gap_drop,
			"src_tcp_action_on_ack_fail":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_ack_fail,
			"src_tcp_out_of_seq_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_out_of_seq_excd,
			"src_tcp_retransmit_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_retransmit_excd,
			"src_tcp_zero_window_excd":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_zero_window_excd,
			"src_tcp_conn_prate_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_conn_prate_excd,
			"src_udp_min_payload":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_min_payload,
			"src_udp_max_payload":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_max_payload,
			"src_udp_conn_prate_excd":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_conn_prate_excd,
			"src_udp_ntp_monlist_req":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_ntp_monlist_req,
			"src_udp_ntp_monlist_resp":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_ntp_monlist_resp,
			"src_udp_wellknown_sport_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_wellknown_sport_drop,
			"src_udp_retry_init":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_retry_init,
			"dst_udp_retry_gap_drop":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_retry_gap_drop,
			"dst_udp_retry_fail":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_retry_fail,
			"dst_tcp_session_aged":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_session_aged,
			"dst_udp_session_aged":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_session_aged,
			"dst_tcp_conn_close":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_close,
			"dst_tcp_conn_close_half_open":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_close_half_open,
			"dst_l4_tcp_auth":                          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_l4_tcp_auth,
			"tcp_l4_syn_cookie_fail":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_l4_syn_cookie_fail,
			"tcp_l4_rst_cookie_fail":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_l4_rst_cookie_fail,
			"tcp_l4_unauth_drop":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_l4_unauth_drop,
			"src_tcp_filter_action_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_filter_action_blacklist,
			"src_tcp_filter_action_whitelist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_filter_action_whitelist,
			"src_tcp_filter_action_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_filter_action_drop,
			"src_tcp_filter_action_default_pass":       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_filter_action_default_pass,
			"src_udp_filter_action_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_filter_action_blacklist,
			"src_udp_filter_action_whitelist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_filter_action_whitelist,
			"src_udp_filter_action_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_filter_action_drop,
			"src_udp_filter_action_default_pass":       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_filter_action_default_pass,
			"src_other_filter_action_blacklist":        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_other_filter_action_blacklist,
			"src_other_filter_action_whitelist":        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_other_filter_action_whitelist,
			"src_other_filter_action_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_other_filter_action_drop,
			"src_other_filter_action_default_pass":     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_other_filter_action_default_pass,
			"tcp_invalid_syn":                          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_invalid_syn,
			"dst_tcp_conn_close_w_rst":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_close_w_rst,
			"dst_tcp_conn_close_w_fin":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_close_w_fin,
			"dst_tcp_conn_close_w_idle":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_close_w_idle,
			"dst_tcp_conn_create_from_syn":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_create_from_syn,
			"dst_tcp_conn_create_from_ack":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_conn_create_from_ack,
			"src_frag_drop":                            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_frag_drop,
			"dst_l4_tcp_blacklist_drop":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_l4_tcp_blacklist_drop,
			"dst_l4_udp_blacklist_drop":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_l4_udp_blacklist_drop,
			"dst_l4_icmp_blacklist_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_l4_icmp_blacklist_drop,
			"dst_l4_other_blacklist_drop":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_l4_other_blacklist_drop,
			"src_l4_tcp_blacklist_drop":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_l4_tcp_blacklist_drop,
			"src_l4_udp_blacklist_drop":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_l4_udp_blacklist_drop,
			"src_l4_icmp_blacklist_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_l4_icmp_blacklist_drop,
			"src_l4_other_blacklist_drop":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_l4_other_blacklist_drop,
			"dst_port_kbit_rate_exceed_pkt":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_kbit_rate_exceed_pkt,
			"dst_tcp_bytes_rcv":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_bytes_rcv,
			"dst_udp_bytes_rcv":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_bytes_rcv,
			"dst_icmp_bytes_rcv":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_bytes_rcv,
			"dst_other_bytes_rcv":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_bytes_rcv,
			"dst_tcp_bytes_sent":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_bytes_sent,
			"dst_udp_bytes_sent":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_bytes_sent,
			"dst_icmp_bytes_sent":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_bytes_sent,
			"dst_other_bytes_sent":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_bytes_sent,
			"dst_udp_auth_drop":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_auth_drop,
			"dst_tcp_auth_drop":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_auth_drop,
			"dst_tcp_auth_resp":                        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_auth_resp,
			"inbound_pkt_drop":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Inbound_pkt_drop,
			"dst_entry_pkt_rate_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_entry_pkt_rate_exceed,
			"dst_entry_kbit_rate_exceed":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_entry_kbit_rate_exceed,
			"dst_entry_conn_limit_exceed":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_entry_conn_limit_exceed,
			"dst_entry_conn_rate_exceed":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_entry_conn_rate_exceed,
			"dst_entry_frag_pkt_rate_exceed":           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_entry_frag_pkt_rate_exceed,
			"dst_icmp_any_exceed":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_any_exceed,
			"dst_other_any_exceed":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_other_any_exceed,
			"src_dst_pair_entry_total":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_dst_pair_entry_total,
			"src_dst_pair_entry_udp":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_dst_pair_entry_udp,
			"src_dst_pair_entry_tcp":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_dst_pair_entry_tcp,
			"src_dst_pair_entry_icmp":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_dst_pair_entry_icmp,
			"src_dst_pair_entry_other":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_dst_pair_entry_other,
			"dst_clist_overflow_policy_at_learning":    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_clist_overflow_policy_at_learning,
			"tcp_rexmit_syn_limit_drop":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_rexmit_syn_limit_drop,
			"tcp_rexmit_syn_limit_bl":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_rexmit_syn_limit_bl,
			"dst_tcp_wellknown_sport_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_wellknown_sport_drop,
			"src_tcp_wellknown_sport_drop":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_wellknown_sport_drop,
			"dst_frag_rcvd":                            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_frag_rcvd,
			"no_policy_class_list_match":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.No_policy_class_list_match,
			"src_udp_retry_gap_drop":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_retry_gap_drop,
			"dst_entry_kbit_rate_exceed_count":         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_entry_kbit_rate_exceed_count,
			"dst_port_undef_hit":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_port_undef_hit,
			"dst_tcp_action_on_ack_timeout":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_timeout,
			"dst_tcp_action_on_ack_reset":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_reset,
			"dst_tcp_action_on_ack_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_ack_blacklist,
			"src_tcp_action_on_ack_timeout":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_ack_timeout,
			"src_tcp_action_on_ack_reset":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_ack_reset,
			"src_tcp_action_on_ack_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_ack_blacklist,
			"dst_tcp_action_on_syn_timeout":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_timeout,
			"dst_tcp_action_on_syn_reset":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_reset,
			"dst_tcp_action_on_syn_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_action_on_syn_blacklist,
			"src_tcp_action_on_syn_timeout":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_syn_timeout,
			"src_tcp_action_on_syn_reset":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_syn_reset,
			"src_tcp_action_on_syn_blacklist":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_tcp_action_on_syn_blacklist,
			"dst_udp_frag_pkt_rate_exceed":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_frag_pkt_rate_exceed,
			"dst_udp_frag_src_rate_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_udp_frag_src_rate_drop,
			"dst_tcp_frag_pkt_rate_exceed":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_frag_pkt_rate_exceed,
			"dst_tcp_frag_src_rate_drop":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_frag_src_rate_drop,
			"dst_icmp_frag_pkt_rate_exceed":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_frag_pkt_rate_exceed,
			"dst_icmp_frag_src_rate_drop":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_icmp_frag_src_rate_drop,
			"sflow_internal_samples_packed":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":            ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Sflow_external_packets_sent,
			"dns_outbound_total_query":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_total_query,
			"dns_outbound_query_malformed":             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_malformed,
			"dns_outbound_query_resp_chk_failed":       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_resp_chk_failed,
			"dns_outbound_query_resp_chk_blacklisted":  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_resp_chk_blacklisted,
			"dns_outbound_query_resp_chk_refused_sent": ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_resp_chk_refused_sent,
			"dns_outbound_query_resp_chk_reset_sent":   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_resp_chk_reset_sent,
			"dns_outbound_query_resp_chk_no_resp_sent": ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_resp_chk_no_resp_sent,
			"dns_outbound_query_resp_size_exceed":      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_resp_size_exceed,
			"dns_outbound_query_sess_timed_out":        ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dns_outbound_query_sess_timed_out,
			"dst_exceed_action_tunnel":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_exceed_action_tunnel,
			"src_udp_auth_timeout":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_auth_timeout,
			"src_udp_retry_pass":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_udp_retry_pass,
			"dst_hw_drop_rule_insert":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_hw_drop_rule_insert,
			"dst_hw_drop_rule_remove":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_hw_drop_rule_remove,
			"src_hw_drop_rule_insert":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_hw_drop_rule_insert,
			"src_hw_drop_rule_remove":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_hw_drop_rule_remove,
			"prog_first_req_time_exceed":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_first_req_time_exceed,
			"prog_req_resp_time_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_req_resp_time_exceed,
			"prog_request_len_exceed":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_request_len_exceed,
			"prog_response_len_exceed":                 ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_response_len_exceed,
			"prog_resp_req_ratio_exceed":               ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_resp_req_ratio_exceed,
			"prog_resp_req_time_exceed":                ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_resp_req_time_exceed,
			"entry_sync_message_received":              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Entry_sync_message_received,
			"entry_sync_message_sent":                  ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Entry_sync_message_sent,
			"prog_conn_sent_exceed":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_conn_sent_exceed,
			"prog_conn_rcvd_exceed":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_conn_rcvd_exceed,
			"prog_conn_time_exceed":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_conn_time_exceed,
			"prog_conn_rcvd_sent_ratio_exceed":         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_conn_rcvd_sent_ratio_exceed,
			"prog_win_sent_exceed":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_win_sent_exceed,
			"prog_win_rcvd_exceed":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_win_rcvd_exceed,
			"prog_win_rcvd_sent_ratio_exceed":          ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_win_rcvd_sent_ratio_exceed,
			"prog_exceed_drop":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_exceed_drop,
			"prog_exceed_bl":                           ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_exceed_bl,
			"prog_conn_exceed_drop":                    ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_conn_exceed_drop,
			"prog_conn_exceed_bl":                      ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_conn_exceed_bl,
			"prog_win_exceed_drop":                     ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_win_exceed_drop,
			"prog_win_exceed_bl":                       ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Prog_win_exceed_bl,
			"dst_exceed_action_drop":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_exceed_action_drop,
			"src_hw_drop":                              ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Src_hw_drop,
			"dst_tcp_auth_rst":                         ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_tcp_auth_rst,
			"dst_src_learn_overflow":                   ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Dst_src_learn_overflow,
			"tcp_fwd_sent":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Tcp_fwd_sent,
			"udp_fwd_sent":                             ret.DtDdosDstDynamicEntryAllEntriesStats.Stats.Udp_fwd_sent,
		},
	}
}

func getObjectDdosDstDynamicEntryAllEntriesStatsStats(d []interface{}) edpt.DdosDstDynamicEntryAllEntriesStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstDynamicEntryAllEntriesStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dst_tcp_any_exceed = in["dst_tcp_any_exceed"].(int)
		ret.Dst_tcp_pkt_rate_exceed = in["dst_tcp_pkt_rate_exceed"].(int)
		ret.Dst_tcp_conn_rate_exceed = in["dst_tcp_conn_rate_exceed"].(int)
		ret.Dst_udp_any_exceed = in["dst_udp_any_exceed"].(int)
		ret.Dst_udp_pkt_rate_exceed = in["dst_udp_pkt_rate_exceed"].(int)
		ret.Dst_udp_conn_limit_exceed = in["dst_udp_conn_limit_exceed"].(int)
		ret.Dst_udp_conn_rate_exceed = in["dst_udp_conn_rate_exceed"].(int)
		ret.Dst_icmp_pkt_rate_exceed = in["dst_icmp_pkt_rate_exceed"].(int)
		ret.Dst_other_pkt_rate_exceed = in["dst_other_pkt_rate_exceed"].(int)
		ret.Dst_other_frag_pkt_rate_exceed = in["dst_other_frag_pkt_rate_exceed"].(int)
		ret.Dst_port_pkt_rate_exceed = in["dst_port_pkt_rate_exceed"].(int)
		ret.Dst_port_conn_limit_exceed = in["dst_port_conn_limit_exceed"].(int)
		ret.Dst_port_conn_rate_exceed = in["dst_port_conn_rate_exceed"].(int)
		ret.Dst_pkt_sent = in["dst_pkt_sent"].(int)
		ret.Dst_udp_pkt_sent = in["dst_udp_pkt_sent"].(int)
		ret.Dst_tcp_pkt_sent = in["dst_tcp_pkt_sent"].(int)
		ret.Dst_icmp_pkt_sent = in["dst_icmp_pkt_sent"].(int)
		ret.Dst_other_pkt_sent = in["dst_other_pkt_sent"].(int)
		ret.Dst_tcp_conn_limit_exceed = in["dst_tcp_conn_limit_exceed"].(int)
		ret.Dst_tcp_pkt_rcvd = in["dst_tcp_pkt_rcvd"].(int)
		ret.Dst_udp_pkt_rcvd = in["dst_udp_pkt_rcvd"].(int)
		ret.Dst_icmp_pkt_rcvd = in["dst_icmp_pkt_rcvd"].(int)
		ret.Dst_other_pkt_rcvd = in["dst_other_pkt_rcvd"].(int)
		ret.Dst_udp_filter_match = in["dst_udp_filter_match"].(int)
		ret.Dst_udp_filter_not_match = in["dst_udp_filter_not_match"].(int)
		ret.Dst_udp_filter_action_blacklist = in["dst_udp_filter_action_blacklist"].(int)
		ret.Dst_udp_filter_action_drop = in["dst_udp_filter_action_drop"].(int)
		ret.Dst_tcp_syn = in["dst_tcp_syn"].(int)
		ret.Dst_tcp_syn_drop = in["dst_tcp_syn_drop"].(int)
		ret.Dst_tcp_src_rate_drop = in["dst_tcp_src_rate_drop"].(int)
		ret.Dst_udp_src_rate_drop = in["dst_udp_src_rate_drop"].(int)
		ret.Dst_icmp_src_rate_drop = in["dst_icmp_src_rate_drop"].(int)
		ret.Dst_other_frag_src_rate_drop = in["dst_other_frag_src_rate_drop"].(int)
		ret.Dst_other_src_rate_drop = in["dst_other_src_rate_drop"].(int)
		ret.Dst_tcp_drop = in["dst_tcp_drop"].(int)
		ret.Dst_udp_drop = in["dst_udp_drop"].(int)
		ret.Dst_icmp_drop = in["dst_icmp_drop"].(int)
		ret.Dst_frag_drop = in["dst_frag_drop"].(int)
		ret.Dst_other_drop = in["dst_other_drop"].(int)
		ret.Dst_tcp_auth = in["dst_tcp_auth"].(int)
		ret.Dst_udp_filter_action_default_pass = in["dst_udp_filter_action_default_pass"].(int)
		ret.Dst_tcp_filter_match = in["dst_tcp_filter_match"].(int)
		ret.Dst_tcp_filter_not_match = in["dst_tcp_filter_not_match"].(int)
		ret.Dst_tcp_filter_action_blacklist = in["dst_tcp_filter_action_blacklist"].(int)
		ret.Dst_tcp_filter_action_drop = in["dst_tcp_filter_action_drop"].(int)
		ret.Dst_tcp_filter_action_default_pass = in["dst_tcp_filter_action_default_pass"].(int)
		ret.Dst_udp_filter_action_whitelist = in["dst_udp_filter_action_whitelist"].(int)
		ret.Dst_udp_kibit_rate_drop = in["dst_udp_kibit_rate_drop"].(int)
		ret.Dst_tcp_kibit_rate_drop = in["dst_tcp_kibit_rate_drop"].(int)
		ret.Dst_icmp_kibit_rate_drop = in["dst_icmp_kibit_rate_drop"].(int)
		ret.Dst_other_kibit_rate_drop = in["dst_other_kibit_rate_drop"].(int)
		ret.Dst_port_undef_drop = in["dst_port_undef_drop"].(int)
		ret.Dst_port_bl = in["dst_port_bl"].(int)
		ret.Dst_src_port_bl = in["dst_src_port_bl"].(int)
		ret.Dst_port_kbit_rate_exceed = in["dst_port_kbit_rate_exceed"].(int)
		ret.Dst_tcp_src_drop = in["dst_tcp_src_drop"].(int)
		ret.Dst_udp_src_drop = in["dst_udp_src_drop"].(int)
		ret.Dst_icmp_src_drop = in["dst_icmp_src_drop"].(int)
		ret.Dst_other_src_drop = in["dst_other_src_drop"].(int)
		ret.Tcp_syn_rcvd = in["tcp_syn_rcvd"].(int)
		ret.Tcp_syn_ack_rcvd = in["tcp_syn_ack_rcvd"].(int)
		ret.Tcp_ack_rcvd = in["tcp_ack_rcvd"].(int)
		ret.Tcp_fin_rcvd = in["tcp_fin_rcvd"].(int)
		ret.Tcp_rst_rcvd = in["tcp_rst_rcvd"].(int)
		ret.Ingress_bytes = in["ingress_bytes"].(int)
		ret.Egress_bytes = in["egress_bytes"].(int)
		ret.Ingress_packets = in["ingress_packets"].(int)
		ret.Egress_packets = in["egress_packets"].(int)
		ret.Tcp_fwd_recv = in["tcp_fwd_recv"].(int)
		ret.Udp_fwd_recv = in["udp_fwd_recv"].(int)
		ret.Icmp_fwd_recv = in["icmp_fwd_recv"].(int)
		ret.Tcp_syn_cookie_fail = in["tcp_syn_cookie_fail"].(int)
		ret.Dst_tcp_session_created = in["dst_tcp_session_created"].(int)
		ret.Dst_udp_session_created = in["dst_udp_session_created"].(int)
		ret.Dst_tcp_filter_action_whitelist = in["dst_tcp_filter_action_whitelist"].(int)
		ret.Dst_other_filter_match = in["dst_other_filter_match"].(int)
		ret.Dst_other_filter_not_match = in["dst_other_filter_not_match"].(int)
		ret.Dst_other_filter_action_blacklist = in["dst_other_filter_action_blacklist"].(int)
		ret.Dst_other_filter_action_drop = in["dst_other_filter_action_drop"].(int)
		ret.Dst_other_filter_action_whitelist = in["dst_other_filter_action_whitelist"].(int)
		ret.Dst_other_filter_action_default_pass = in["dst_other_filter_action_default_pass"].(int)
		ret.Dst_blackhole_inject = in["dst_blackhole_inject"].(int)
		ret.Dst_blackhole_withdraw = in["dst_blackhole_withdraw"].(int)
		ret.Dst_tcp_out_of_seq_excd = in["dst_tcp_out_of_seq_excd"].(int)
		ret.Dst_tcp_retransmit_excd = in["dst_tcp_retransmit_excd"].(int)
		ret.Dst_tcp_zero_window_excd = in["dst_tcp_zero_window_excd"].(int)
		ret.Dst_tcp_conn_prate_excd = in["dst_tcp_conn_prate_excd"].(int)
		ret.Dst_tcp_action_on_ack_init = in["dst_tcp_action_on_ack_init"].(int)
		ret.Dst_tcp_action_on_ack_gap_drop = in["dst_tcp_action_on_ack_gap_drop"].(int)
		ret.Dst_tcp_action_on_ack_fail = in["dst_tcp_action_on_ack_fail"].(int)
		ret.Dst_tcp_action_on_ack_pass = in["dst_tcp_action_on_ack_pass"].(int)
		ret.Dst_tcp_action_on_syn_init = in["dst_tcp_action_on_syn_init"].(int)
		ret.Dst_tcp_action_on_syn_gap_drop = in["dst_tcp_action_on_syn_gap_drop"].(int)
		ret.Dst_tcp_action_on_syn_fail = in["dst_tcp_action_on_syn_fail"].(int)
		ret.Dst_tcp_action_on_syn_pass = in["dst_tcp_action_on_syn_pass"].(int)
		ret.Udp_payload_too_small = in["udp_payload_too_small"].(int)
		ret.Udp_payload_too_big = in["udp_payload_too_big"].(int)
		ret.Dst_udp_conn_prate_excd = in["dst_udp_conn_prate_excd"].(int)
		ret.Dst_udp_ntp_monlist_req = in["dst_udp_ntp_monlist_req"].(int)
		ret.Dst_udp_ntp_monlist_resp = in["dst_udp_ntp_monlist_resp"].(int)
		ret.Dst_udp_wellknown_sport_drop = in["dst_udp_wellknown_sport_drop"].(int)
		ret.Dst_udp_retry_init = in["dst_udp_retry_init"].(int)
		ret.Dst_udp_retry_pass = in["dst_udp_retry_pass"].(int)
		ret.Dst_tcp_bytes_drop = in["dst_tcp_bytes_drop"].(int)
		ret.Dst_udp_bytes_drop = in["dst_udp_bytes_drop"].(int)
		ret.Dst_icmp_bytes_drop = in["dst_icmp_bytes_drop"].(int)
		ret.Dst_other_bytes_drop = in["dst_other_bytes_drop"].(int)
		ret.Dst_out_no_route = in["dst_out_no_route"].(int)
		ret.Outbound_bytes_sent = in["outbound_bytes_sent"].(int)
		ret.Outbound_pkt_drop = in["outbound_pkt_drop"].(int)
		ret.Outbound_bytes_drop = in["outbound_bytes_drop"].(int)
		ret.Outbound_pkt_sent = in["outbound_pkt_sent"].(int)
		ret.Inbound_bytes_sent = in["inbound_bytes_sent"].(int)
		ret.Inbound_bytes_drop = in["inbound_bytes_drop"].(int)
		ret.Dst_src_port_pkt_rate_exceed = in["dst_src_port_pkt_rate_exceed"].(int)
		ret.Dst_src_port_kbit_rate_exceed = in["dst_src_port_kbit_rate_exceed"].(int)
		ret.Dst_src_port_conn_limit_exceed = in["dst_src_port_conn_limit_exceed"].(int)
		ret.Dst_src_port_conn_rate_exceed = in["dst_src_port_conn_rate_exceed"].(int)
		ret.Dst_ip_proto_pkt_rate_exceed = in["dst_ip_proto_pkt_rate_exceed"].(int)
		ret.Dst_ip_proto_kbit_rate_exceed = in["dst_ip_proto_kbit_rate_exceed"].(int)
		ret.Dst_tcp_port_any_exceed = in["dst_tcp_port_any_exceed"].(int)
		ret.Dst_udp_port_any_exceed = in["dst_udp_port_any_exceed"].(int)
		ret.Dst_tcp_auth_pass = in["dst_tcp_auth_pass"].(int)
		ret.Dst_tcp_rst_cookie_fail = in["dst_tcp_rst_cookie_fail"].(int)
		ret.Dst_tcp_unauth_drop = in["dst_tcp_unauth_drop"].(int)
		ret.Src_tcp_syn_auth_fail = in["src_tcp_syn_auth_fail"].(int)
		ret.Src_tcp_syn_cookie_sent = in["src_tcp_syn_cookie_sent"].(int)
		ret.Src_tcp_syn_cookie_fail = in["src_tcp_syn_cookie_fail"].(int)
		ret.Src_tcp_rst_cookie_fail = in["src_tcp_rst_cookie_fail"].(int)
		ret.Src_tcp_unauth_drop = in["src_tcp_unauth_drop"].(int)
		ret.Src_tcp_action_on_syn_init = in["src_tcp_action_on_syn_init"].(int)
		ret.Src_tcp_action_on_syn_gap_drop = in["src_tcp_action_on_syn_gap_drop"].(int)
		ret.Src_tcp_action_on_syn_fail = in["src_tcp_action_on_syn_fail"].(int)
		ret.Src_tcp_action_on_ack_init = in["src_tcp_action_on_ack_init"].(int)
		ret.Src_tcp_action_on_ack_gap_drop = in["src_tcp_action_on_ack_gap_drop"].(int)
		ret.Src_tcp_action_on_ack_fail = in["src_tcp_action_on_ack_fail"].(int)
		ret.Src_tcp_out_of_seq_excd = in["src_tcp_out_of_seq_excd"].(int)
		ret.Src_tcp_retransmit_excd = in["src_tcp_retransmit_excd"].(int)
		ret.Src_tcp_zero_window_excd = in["src_tcp_zero_window_excd"].(int)
		ret.Src_tcp_conn_prate_excd = in["src_tcp_conn_prate_excd"].(int)
		ret.Src_udp_min_payload = in["src_udp_min_payload"].(int)
		ret.Src_udp_max_payload = in["src_udp_max_payload"].(int)
		ret.Src_udp_conn_prate_excd = in["src_udp_conn_prate_excd"].(int)
		ret.Src_udp_ntp_monlist_req = in["src_udp_ntp_monlist_req"].(int)
		ret.Src_udp_ntp_monlist_resp = in["src_udp_ntp_monlist_resp"].(int)
		ret.Src_udp_wellknown_sport_drop = in["src_udp_wellknown_sport_drop"].(int)
		ret.Src_udp_retry_init = in["src_udp_retry_init"].(int)
		ret.Dst_udp_retry_gap_drop = in["dst_udp_retry_gap_drop"].(int)
		ret.Dst_udp_retry_fail = in["dst_udp_retry_fail"].(int)
		ret.Dst_tcp_session_aged = in["dst_tcp_session_aged"].(int)
		ret.Dst_udp_session_aged = in["dst_udp_session_aged"].(int)
		ret.Dst_tcp_conn_close = in["dst_tcp_conn_close"].(int)
		ret.Dst_tcp_conn_close_half_open = in["dst_tcp_conn_close_half_open"].(int)
		ret.Dst_l4_tcp_auth = in["dst_l4_tcp_auth"].(int)
		ret.Tcp_l4_syn_cookie_fail = in["tcp_l4_syn_cookie_fail"].(int)
		ret.Tcp_l4_rst_cookie_fail = in["tcp_l4_rst_cookie_fail"].(int)
		ret.Tcp_l4_unauth_drop = in["tcp_l4_unauth_drop"].(int)
		ret.Src_tcp_filter_action_blacklist = in["src_tcp_filter_action_blacklist"].(int)
		ret.Src_tcp_filter_action_whitelist = in["src_tcp_filter_action_whitelist"].(int)
		ret.Src_tcp_filter_action_drop = in["src_tcp_filter_action_drop"].(int)
		ret.Src_tcp_filter_action_default_pass = in["src_tcp_filter_action_default_pass"].(int)
		ret.Src_udp_filter_action_blacklist = in["src_udp_filter_action_blacklist"].(int)
		ret.Src_udp_filter_action_whitelist = in["src_udp_filter_action_whitelist"].(int)
		ret.Src_udp_filter_action_drop = in["src_udp_filter_action_drop"].(int)
		ret.Src_udp_filter_action_default_pass = in["src_udp_filter_action_default_pass"].(int)
		ret.Src_other_filter_action_blacklist = in["src_other_filter_action_blacklist"].(int)
		ret.Src_other_filter_action_whitelist = in["src_other_filter_action_whitelist"].(int)
		ret.Src_other_filter_action_drop = in["src_other_filter_action_drop"].(int)
		ret.Src_other_filter_action_default_pass = in["src_other_filter_action_default_pass"].(int)
		ret.Tcp_invalid_syn = in["tcp_invalid_syn"].(int)
		ret.Dst_tcp_conn_close_w_rst = in["dst_tcp_conn_close_w_rst"].(int)
		ret.Dst_tcp_conn_close_w_fin = in["dst_tcp_conn_close_w_fin"].(int)
		ret.Dst_tcp_conn_close_w_idle = in["dst_tcp_conn_close_w_idle"].(int)
		ret.Dst_tcp_conn_create_from_syn = in["dst_tcp_conn_create_from_syn"].(int)
		ret.Dst_tcp_conn_create_from_ack = in["dst_tcp_conn_create_from_ack"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Dst_l4_tcp_blacklist_drop = in["dst_l4_tcp_blacklist_drop"].(int)
		ret.Dst_l4_udp_blacklist_drop = in["dst_l4_udp_blacklist_drop"].(int)
		ret.Dst_l4_icmp_blacklist_drop = in["dst_l4_icmp_blacklist_drop"].(int)
		ret.Dst_l4_other_blacklist_drop = in["dst_l4_other_blacklist_drop"].(int)
		ret.Src_l4_tcp_blacklist_drop = in["src_l4_tcp_blacklist_drop"].(int)
		ret.Src_l4_udp_blacklist_drop = in["src_l4_udp_blacklist_drop"].(int)
		ret.Src_l4_icmp_blacklist_drop = in["src_l4_icmp_blacklist_drop"].(int)
		ret.Src_l4_other_blacklist_drop = in["src_l4_other_blacklist_drop"].(int)
		ret.Dst_port_kbit_rate_exceed_pkt = in["dst_port_kbit_rate_exceed_pkt"].(int)
		ret.Dst_tcp_bytes_rcv = in["dst_tcp_bytes_rcv"].(int)
		ret.Dst_udp_bytes_rcv = in["dst_udp_bytes_rcv"].(int)
		ret.Dst_icmp_bytes_rcv = in["dst_icmp_bytes_rcv"].(int)
		ret.Dst_other_bytes_rcv = in["dst_other_bytes_rcv"].(int)
		ret.Dst_tcp_bytes_sent = in["dst_tcp_bytes_sent"].(int)
		ret.Dst_udp_bytes_sent = in["dst_udp_bytes_sent"].(int)
		ret.Dst_icmp_bytes_sent = in["dst_icmp_bytes_sent"].(int)
		ret.Dst_other_bytes_sent = in["dst_other_bytes_sent"].(int)
		ret.Dst_udp_auth_drop = in["dst_udp_auth_drop"].(int)
		ret.Dst_tcp_auth_drop = in["dst_tcp_auth_drop"].(int)
		ret.Dst_tcp_auth_resp = in["dst_tcp_auth_resp"].(int)
		ret.Inbound_pkt_drop = in["inbound_pkt_drop"].(int)
		ret.Dst_entry_pkt_rate_exceed = in["dst_entry_pkt_rate_exceed"].(int)
		ret.Dst_entry_kbit_rate_exceed = in["dst_entry_kbit_rate_exceed"].(int)
		ret.Dst_entry_conn_limit_exceed = in["dst_entry_conn_limit_exceed"].(int)
		ret.Dst_entry_conn_rate_exceed = in["dst_entry_conn_rate_exceed"].(int)
		ret.Dst_entry_frag_pkt_rate_exceed = in["dst_entry_frag_pkt_rate_exceed"].(int)
		ret.Dst_icmp_any_exceed = in["dst_icmp_any_exceed"].(int)
		ret.Dst_other_any_exceed = in["dst_other_any_exceed"].(int)
		ret.Src_dst_pair_entry_total = in["src_dst_pair_entry_total"].(int)
		ret.Src_dst_pair_entry_udp = in["src_dst_pair_entry_udp"].(int)
		ret.Src_dst_pair_entry_tcp = in["src_dst_pair_entry_tcp"].(int)
		ret.Src_dst_pair_entry_icmp = in["src_dst_pair_entry_icmp"].(int)
		ret.Src_dst_pair_entry_other = in["src_dst_pair_entry_other"].(int)
		ret.Dst_clist_overflow_policy_at_learning = in["dst_clist_overflow_policy_at_learning"].(int)
		ret.Tcp_rexmit_syn_limit_drop = in["tcp_rexmit_syn_limit_drop"].(int)
		ret.Tcp_rexmit_syn_limit_bl = in["tcp_rexmit_syn_limit_bl"].(int)
		ret.Dst_tcp_wellknown_sport_drop = in["dst_tcp_wellknown_sport_drop"].(int)
		ret.Src_tcp_wellknown_sport_drop = in["src_tcp_wellknown_sport_drop"].(int)
		ret.Dst_frag_rcvd = in["dst_frag_rcvd"].(int)
		ret.No_policy_class_list_match = in["no_policy_class_list_match"].(int)
		ret.Src_udp_retry_gap_drop = in["src_udp_retry_gap_drop"].(int)
		ret.Dst_entry_kbit_rate_exceed_count = in["dst_entry_kbit_rate_exceed_count"].(int)
		ret.Dst_port_undef_hit = in["dst_port_undef_hit"].(int)
		ret.Dst_tcp_action_on_ack_timeout = in["dst_tcp_action_on_ack_timeout"].(int)
		ret.Dst_tcp_action_on_ack_reset = in["dst_tcp_action_on_ack_reset"].(int)
		ret.Dst_tcp_action_on_ack_blacklist = in["dst_tcp_action_on_ack_blacklist"].(int)
		ret.Src_tcp_action_on_ack_timeout = in["src_tcp_action_on_ack_timeout"].(int)
		ret.Src_tcp_action_on_ack_reset = in["src_tcp_action_on_ack_reset"].(int)
		ret.Src_tcp_action_on_ack_blacklist = in["src_tcp_action_on_ack_blacklist"].(int)
		ret.Dst_tcp_action_on_syn_timeout = in["dst_tcp_action_on_syn_timeout"].(int)
		ret.Dst_tcp_action_on_syn_reset = in["dst_tcp_action_on_syn_reset"].(int)
		ret.Dst_tcp_action_on_syn_blacklist = in["dst_tcp_action_on_syn_blacklist"].(int)
		ret.Src_tcp_action_on_syn_timeout = in["src_tcp_action_on_syn_timeout"].(int)
		ret.Src_tcp_action_on_syn_reset = in["src_tcp_action_on_syn_reset"].(int)
		ret.Src_tcp_action_on_syn_blacklist = in["src_tcp_action_on_syn_blacklist"].(int)
		ret.Dst_udp_frag_pkt_rate_exceed = in["dst_udp_frag_pkt_rate_exceed"].(int)
		ret.Dst_udp_frag_src_rate_drop = in["dst_udp_frag_src_rate_drop"].(int)
		ret.Dst_tcp_frag_pkt_rate_exceed = in["dst_tcp_frag_pkt_rate_exceed"].(int)
		ret.Dst_tcp_frag_src_rate_drop = in["dst_tcp_frag_src_rate_drop"].(int)
		ret.Dst_icmp_frag_pkt_rate_exceed = in["dst_icmp_frag_pkt_rate_exceed"].(int)
		ret.Dst_icmp_frag_src_rate_drop = in["dst_icmp_frag_src_rate_drop"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Dns_outbound_total_query = in["dns_outbound_total_query"].(int)
		ret.Dns_outbound_query_malformed = in["dns_outbound_query_malformed"].(int)
		ret.Dns_outbound_query_resp_chk_failed = in["dns_outbound_query_resp_chk_failed"].(int)
		ret.Dns_outbound_query_resp_chk_blacklisted = in["dns_outbound_query_resp_chk_blacklisted"].(int)
		ret.Dns_outbound_query_resp_chk_refused_sent = in["dns_outbound_query_resp_chk_refused_sent"].(int)
		ret.Dns_outbound_query_resp_chk_reset_sent = in["dns_outbound_query_resp_chk_reset_sent"].(int)
		ret.Dns_outbound_query_resp_chk_no_resp_sent = in["dns_outbound_query_resp_chk_no_resp_sent"].(int)
		ret.Dns_outbound_query_resp_size_exceed = in["dns_outbound_query_resp_size_exceed"].(int)
		ret.Dns_outbound_query_sess_timed_out = in["dns_outbound_query_sess_timed_out"].(int)
		ret.Dst_exceed_action_tunnel = in["dst_exceed_action_tunnel"].(int)
		ret.Src_udp_auth_timeout = in["src_udp_auth_timeout"].(int)
		ret.Src_udp_retry_pass = in["src_udp_retry_pass"].(int)
		ret.Dst_hw_drop_rule_insert = in["dst_hw_drop_rule_insert"].(int)
		ret.Dst_hw_drop_rule_remove = in["dst_hw_drop_rule_remove"].(int)
		ret.Src_hw_drop_rule_insert = in["src_hw_drop_rule_insert"].(int)
		ret.Src_hw_drop_rule_remove = in["src_hw_drop_rule_remove"].(int)
		ret.Prog_first_req_time_exceed = in["prog_first_req_time_exceed"].(int)
		ret.Prog_req_resp_time_exceed = in["prog_req_resp_time_exceed"].(int)
		ret.Prog_request_len_exceed = in["prog_request_len_exceed"].(int)
		ret.Prog_response_len_exceed = in["prog_response_len_exceed"].(int)
		ret.Prog_resp_req_ratio_exceed = in["prog_resp_req_ratio_exceed"].(int)
		ret.Prog_resp_req_time_exceed = in["prog_resp_req_time_exceed"].(int)
		ret.Entry_sync_message_received = in["entry_sync_message_received"].(int)
		ret.Entry_sync_message_sent = in["entry_sync_message_sent"].(int)
		ret.Prog_conn_sent_exceed = in["prog_conn_sent_exceed"].(int)
		ret.Prog_conn_rcvd_exceed = in["prog_conn_rcvd_exceed"].(int)
		ret.Prog_conn_time_exceed = in["prog_conn_time_exceed"].(int)
		ret.Prog_conn_rcvd_sent_ratio_exceed = in["prog_conn_rcvd_sent_ratio_exceed"].(int)
		ret.Prog_win_sent_exceed = in["prog_win_sent_exceed"].(int)
		ret.Prog_win_rcvd_exceed = in["prog_win_rcvd_exceed"].(int)
		ret.Prog_win_rcvd_sent_ratio_exceed = in["prog_win_rcvd_sent_ratio_exceed"].(int)
		ret.Prog_exceed_drop = in["prog_exceed_drop"].(int)
		ret.Prog_exceed_bl = in["prog_exceed_bl"].(int)
		ret.Prog_conn_exceed_drop = in["prog_conn_exceed_drop"].(int)
		ret.Prog_conn_exceed_bl = in["prog_conn_exceed_bl"].(int)
		ret.Prog_win_exceed_drop = in["prog_win_exceed_drop"].(int)
		ret.Prog_win_exceed_bl = in["prog_win_exceed_bl"].(int)
		ret.Dst_exceed_action_drop = in["dst_exceed_action_drop"].(int)
		ret.Src_hw_drop = in["src_hw_drop"].(int)
		ret.Dst_tcp_auth_rst = in["dst_tcp_auth_rst"].(int)
		ret.Dst_src_learn_overflow = in["dst_src_learn_overflow"].(int)
		ret.Tcp_fwd_sent = in["tcp_fwd_sent"].(int)
		ret.Udp_fwd_sent = in["udp_fwd_sent"].(int)
	}
	return ret
}

func dataToEndpointDdosDstDynamicEntryAllEntriesStats(d *schema.ResourceData) edpt.DdosDstDynamicEntryAllEntriesStats {
	var ret edpt.DdosDstDynamicEntryAllEntriesStats

	ret.Stats = getObjectDdosDstDynamicEntryAllEntriesStatsStats(d.Get("stats").([]interface{}))
	return ret
}
