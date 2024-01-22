package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dst_zone_stats`: Statistics for the object zone\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDstZoneStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_tcp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst IP-Proto Rate: Total Exceeded",
						},
						"zone_tcp_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst IP-Proto Rate: Packet Exceeded",
						},
						"zone_tcp_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst IP-Proto Rate: Conn Exceeded",
						},
						"zone_udp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst IP-Proto Rate: Total Exceeded",
						},
						"zone_udp_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst IP-Proto Rate: Packet Exceeded",
						},
						"zone_udp_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst IP-Proto Limit: Conn Exceeded",
						},
						"zone_udp_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst IP-Proto Rate: Conn Exceeded",
						},
						"zone_icmp_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst Rate: Packet Exceeded",
						},
						"zone_other_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst IP-Proto Rate: Packet Exceeded",
						},
						"zone_other_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst IP-Proto Rate: Frag Exceeded",
						},
						"zone_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: Packet Exceeded",
						},
						"zone_port_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Limit: Conn Exceeded",
						},
						"zone_port_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: Conn Exceeded",
						},
						"zone_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Inbound: Packets Forwarded",
						},
						"zone_udp_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Forwarded",
						},
						"zone_tcp_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Forwarded",
						},
						"zone_icmp_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Forwarded",
						},
						"zone_other_pkt_sent": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Forwarded",
						},
						"zone_tcp_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst IP-Proto Limit: Conn Exceeded",
						},
						"zone_tcp_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Received",
						},
						"zone_udp_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Received",
						},
						"zone_icmp_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Received",
						},
						"zone_other_pkt_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Received",
						},
						"zone_udp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Match",
						},
						"zone_udp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Not Matched on Pkt",
						},
						"zone_udp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action Blacklist",
						},
						"zone_udp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action Drop",
						},
						"zone_tcp_syn": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total SYN Received",
						},
						"zone_tcp_syn_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Packets Dropped",
						},
						"zone_tcp_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Rate: Total Exceeded",
						},
						"zone_udp_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Rate: Total Exceeded",
						},
						"zone_icmp_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Rate: Total Exceeded",
						},
						"zone_other_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Rate: Frag Exceeded",
						},
						"zone_other_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Src Rate: Total Exceeded",
						},
						"zone_tcp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Packets Dropped",
						},
						"zone_udp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Packets Dropped",
						},
						"zone_icmp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Packets Dropped",
						},
						"zone_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Dropped",
						},
						"zone_other_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Packets Dropped",
						},
						"zone_tcp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Cookie Sent",
						},
						"zone_udp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action Default Pass",
						},
						"zone_tcp_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Match",
						},
						"zone_tcp_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Not Matched on Pkt",
						},
						"zone_tcp_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action Blacklist",
						},
						"zone_tcp_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action Drop",
						},
						"zone_tcp_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action Default Pass",
						},
						"zone_udp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Filter Action WL",
						},
						"zone_udp_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst IP-Proto Rate: KiBit Exceeded",
						},
						"zone_tcp_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst IP-Proto Rate: KiBit Exceeded",
						},
						"zone_icmp_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst Rate: KiBit Exceeded",
						},
						"zone_other_kibit_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Dst IP-Proto Rate: KiBit Exceeded",
						},
						"zone_port_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Undefined Dropped",
						},
						"zone_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port Blacklist Packets Dropped",
						},
						"zone_src_port_bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst SrcPort Blacklist Packets Dropped",
						},
						"zone_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Port Rate: KiBit Exceeded",
						},
						"zone_tcp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Packets Dropped",
						},
						"zone_udp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Packets Dropped",
						},
						"zone_icmp_src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Src Packets Dropped",
						},
						"zone_other_src_drop": {
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
						"zone_tcp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Sessions Created",
						},
						"zone_udp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Sessions Created",
						},
						"zone_tcp_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Filter Action WL",
						},
						"zone_other_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Match",
						},
						"zone_other_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Not Matched on Pkt",
						},
						"zone_other_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action Blacklist",
						},
						"zone_other_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action Drop",
						},
						"zone_other_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action WL",
						},
						"zone_other_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Filter Action Default Pass",
						},
						"zone_blackhole_inject": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Inject",
						},
						"zone_blackhole_withdraw": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blackhole Withdraw",
						},
						"zone_tcp_out_of_seq_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Out-Of-Seq Exceeded",
						},
						"zone_tcp_retransmit_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit Exceeded",
						},
						"zone_tcp_zero_window_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Zero-Window Exceeded",
						},
						"zone_tcp_conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Rate: Conn Pkt Exceeded",
						},
						"zone_tcp_action_on_ack_init": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Init",
						},
						"zone_tcp_action_on_ack_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Retry-Gap Dropped",
						},
						"zone_tcp_action_on_ack_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Dropped",
						},
						"zone_tcp_action_on_ack_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Passed",
						},
						"zone_tcp_action_on_syn_init": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Init",
						},
						"zone_tcp_action_on_syn_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry-Gap Dropped",
						},
						"zone_tcp_action_on_syn_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Dropped",
						},
						"zone_tcp_action_on_syn_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Passed",
						},
						"zone_payload_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small",
						},
						"zone_payload_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large",
						},
						"zone_udp_conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Rate: Conn Pkt Exceeded",
						},
						"zone_udp_ntp_monlist_req": {
							Type: schema.TypeInt, Optional: true, Description: "UDP NTP Monlist Request",
						},
						"zone_udp_ntp_monlist_resp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP NTP Monlist Response",
						},
						"zone_udp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcPort Wellknown",
						},
						"zone_udp_retry_init": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry Init",
						},
						"zone_udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry Passed",
						},
						"zone_tcp_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Total Bytes Dropped",
						},
						"zone_udp_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Total Bytes Dropped",
						},
						"zone_icmp_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Total Bytes Dropped",
						},
						"zone_other_bytes_drop": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Total Bytes Dropped",
						},
						"zone_out_no_route": {
							Type: schema.TypeInt, Optional: true, Description: "Dst IPv4/v6 Out No Route",
						},
						"outbound_bytes_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: Bytes Forwarded",
						},
						"outbound_drop": {
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
						"zone_src_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Rate: Packet Exceeded",
						},
						"zone_src_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Rate: KiBit Exceeded",
						},
						"zone_src_port_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Limit: Conn Exceeded",
						},
						"zone_src_port_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "SrcPort Rate: Conn Exceeded",
						},
						"zone_ip_proto_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Proto Rate: Packet Exceeded",
						},
						"zone_ip_proto_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "IP-Proto Rate: KiBit Exceeded",
						},
						"zone_tcp_port_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Rate: Total Exceed",
						},
						"zone_udp_port_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Rate: Total Exceed",
						},
						"zone_tcp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Auth Passed",
						},
						"zone_tcp_rst_cookie_fail": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: RST Cookie Failed",
						},
						"zone_tcp_unauth_drop": {
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
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry Timeout",
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
						"zone_port_kbit_rate_exceed_pkt": {
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
						"dst_drop": {
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
						"dst_l4_tcp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst TCP IP-Proto Blacklist Dropped",
						},
						"dst_l4_udp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst UDP IP-Proto Blacklist Dropped",
						},
						"dst_l4_icmp_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst ICMP IP-Proto Blacklist Dropped",
						},
						"dst_l4_other_blacklist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst OTHER IP-Proto Blacklist Dropped",
						},
						"dst_icmp_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Rate: Total Exceed",
						},
						"dst_other_any_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "OTHER Rate: Total Exceed",
						},
						"tcp_rexmit_syn_limit_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retransmit Exceeded Drop",
						},
						"tcp_rexmit_syn_limit_bl": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retransmit Exceeded Blacklist",
						},
						"dst_clist_overflow_policy_at_learning": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Src-Based Overflow Policy Hit",
						},
						"zone_frag_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Received",
						},
						"zone_tcp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SrcPort Wellknown",
						},
						"src_tcp_wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP SrcPort Wellknown",
						},
						"secondary_dst_entry_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Rate: Packet Exceeded",
						},
						"secondary_dst_entry_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Rate: KiBit Exceeded",
						},
						"secondary_dst_entry_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Limit: Conn Exceeded",
						},
						"secondary_dst_entry_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Rate: Conn Exceeded",
						},
						"secondary_dst_entry_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Rate: Frag Packet Exceeded",
						},
						"src_udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth: Retry-Gap Dropped",
						},
						"dst_entry_kbit_rate_exceed_count": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Rate: KiBit Exceeded Count",
						},
						"secondary_entry_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Entry Learned",
						},
						"secondary_entry_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Entry Hit",
						},
						"secondary_entry_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Entry Missed",
						},
						"secondary_entry_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Entry Aged",
						},
						"secondary_entry_learning_thre_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Entry Count Overflow",
						},
						"zone_port_undef_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Port undefined Hit",
						},
						"zone_tcp_action_on_ack_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Timeout",
						},
						"zone_tcp_action_on_ack_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: ACK Retry Timeout Reset",
						},
						"zone_tcp_action_on_ack_blacklist": {
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
						"zone_tcp_action_on_syn_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Timeout",
						},
						"zone_tcp_action_on_syn_reset": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: SYN Retry Timeout Reset",
						},
						"zone_tcp_action_on_syn_blacklist": {
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
						"zone_udp_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dst IP-Proto Rate: Frag Exceeded",
						},
						"zone_udp_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Src Rate: Frag Exceeded",
						},
						"zone_tcp_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dst IP-Proto Rate: Frag Exceeded",
						},
						"zone_tcp_frag_src_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Src Rate: Frag Exceeded",
						},
						"zone_icmp_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dst IP-Proto Rate: Frag Exceeded",
						},
						"zone_icmp_frag_src_rate_drop": {
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
						"source_entry_total": {
							Type: schema.TypeInt, Optional: true, Description: "Source Entry Total Count",
						},
						"source_entry_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Source Entry UDP Count",
						},
						"source_entry_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Source Entry TCP Count",
						},
						"source_entry_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Source Entry ICMP Count",
						},
						"source_entry_other": {
							Type: schema.TypeInt, Optional: true, Description: "Source Entry OTHER Count",
						},
						"dst_exceed_action_tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Exceed Action: Tunnel",
						},
						"dst_udp_retry_timeout_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth: Retry Timeout Blacklisted",
						},
						"src_udp_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth: Retry Timeout",
						},
						"zone_src_udp_retry_timeout_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth: Retry Timeout Blacklisted",
						},
						"src_udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Passed",
						},
						"secondary_port_learn": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Port Learned",
						},
						"secondary_port_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr Port Aged",
						},
						"dst_entry_outbound_udp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: UDP Sessions Created",
						},
						"dst_entry_outbound_udp_session_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: UDP Sessions Aged",
						},
						"dst_entry_outbound_tcp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: TCP Sessions Created",
						},
						"dst_entry_outbound_tcp_session_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound: TCP Sessions Aged",
						},
						"dst_entry_outbound_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Rate: Packet Exceeded",
						},
						"dst_entry_outbound_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Rate: KiBit Exceeded",
						},
						"dst_entry_outbound_kbit_rate_exceed_count": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Rate: KiBit Exceeded Count",
						},
						"dst_entry_outbound_conn_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Limit: Conn Exceeded",
						},
						"dst_entry_outbound_conn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Rate: Conn Exceeded",
						},
						"dst_entry_outbound_frag_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Outbound Rate: Frag Packet Exceeded",
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
							Type: schema.TypeInt, Optional: true, Description: "Connection: Received to Sent Ratio Exceed",
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
						"east_west_inbound_rcv_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Inbound Packets Received",
						},
						"east_west_inbound_drop_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Inbound Packets Dropped",
						},
						"east_west_inbound_fwd_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Inbound Packets Forwarded",
						},
						"east_west_inbound_rcv_byte": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Inbound Bytes Received",
						},
						"east_west_inbound_drop_byte": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Inbound Bytes Dropped",
						},
						"east_west_inbound_fwd_byte": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Inbound Bytes Forwarded",
						},
						"east_west_outbound_rcv_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Outbound Packets Received",
						},
						"east_west_outbound_drop_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Outbound Packets Dropped",
						},
						"east_west_outbound_fwd_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Outbound Packets Forwarded",
						},
						"east_west_outbound_rcv_byte": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Outbound Bytes Received",
						},
						"east_west_outbound_drop_byte": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Outbound Bytes Dropped",
						},
						"east_west_outbound_fwd_byte": {
							Type: schema.TypeInt, Optional: true, Description: "East West: Outbound Bytes Forwarded",
						},
						"dst_exceed_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Entry Exceed Action: Dropped",
						},
						"prog_conn_samples": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Collected: Connection",
						},
						"prog_req_samples": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Collected: Req-Resp",
						},
						"prog_win_samples": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Collected: Time Window",
						},
						"victim_ip_learned": {
							Type: schema.TypeInt, Optional: true, Description: "Victim Identification: IP Entry Learned",
						},
						"victim_ip_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Victim Identification: IP Entry Aged",
						},
						"prog_conn_samples_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed: Connnection",
						},
						"prog_req_samples_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed: Req-Resp",
						},
						"prog_win_samples_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed: Time Window",
						},
						"dst_src_learn_overflow": {
							Type: schema.TypeInt, Optional: true, Description: "Src Dynamic Entry Count Overflow",
						},
						"dst_tcp_auth_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth: Reset",
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "",
			},
		},
	}
}

func resourceDdosDstZoneStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDstZoneStatsStats := setObjectDdosDstZoneStatsStats(res)
		d.Set("stats", DdosDstZoneStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDstZoneStatsStats(ret edpt.DataDdosDstZoneStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"zone_tcp_any_exceed":                       ret.DtDdosDstZoneStats.Stats.Zone_tcp_any_exceed,
			"zone_tcp_pkt_rate_exceed":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_pkt_rate_exceed,
			"zone_tcp_conn_rate_exceed":                 ret.DtDdosDstZoneStats.Stats.Zone_tcp_conn_rate_exceed,
			"zone_udp_any_exceed":                       ret.DtDdosDstZoneStats.Stats.Zone_udp_any_exceed,
			"zone_udp_pkt_rate_exceed":                  ret.DtDdosDstZoneStats.Stats.Zone_udp_pkt_rate_exceed,
			"zone_udp_conn_limit_exceed":                ret.DtDdosDstZoneStats.Stats.Zone_udp_conn_limit_exceed,
			"zone_udp_conn_rate_exceed":                 ret.DtDdosDstZoneStats.Stats.Zone_udp_conn_rate_exceed,
			"zone_icmp_pkt_rate_exceed":                 ret.DtDdosDstZoneStats.Stats.Zone_icmp_pkt_rate_exceed,
			"zone_other_pkt_rate_exceed":                ret.DtDdosDstZoneStats.Stats.Zone_other_pkt_rate_exceed,
			"zone_other_frag_pkt_rate_exceed":           ret.DtDdosDstZoneStats.Stats.Zone_other_frag_pkt_rate_exceed,
			"zone_port_pkt_rate_exceed":                 ret.DtDdosDstZoneStats.Stats.Zone_port_pkt_rate_exceed,
			"zone_port_conn_limit_exceed":               ret.DtDdosDstZoneStats.Stats.Zone_port_conn_limit_exceed,
			"zone_port_conn_rate_exceed":                ret.DtDdosDstZoneStats.Stats.Zone_port_conn_rate_exceed,
			"zone_pkt_sent":                             ret.DtDdosDstZoneStats.Stats.Zone_pkt_sent,
			"zone_udp_pkt_sent":                         ret.DtDdosDstZoneStats.Stats.Zone_udp_pkt_sent,
			"zone_tcp_pkt_sent":                         ret.DtDdosDstZoneStats.Stats.Zone_tcp_pkt_sent,
			"zone_icmp_pkt_sent":                        ret.DtDdosDstZoneStats.Stats.Zone_icmp_pkt_sent,
			"zone_other_pkt_sent":                       ret.DtDdosDstZoneStats.Stats.Zone_other_pkt_sent,
			"zone_tcp_conn_limit_exceed":                ret.DtDdosDstZoneStats.Stats.Zone_tcp_conn_limit_exceed,
			"zone_tcp_pkt_rcvd":                         ret.DtDdosDstZoneStats.Stats.Zone_tcp_pkt_rcvd,
			"zone_udp_pkt_rcvd":                         ret.DtDdosDstZoneStats.Stats.Zone_udp_pkt_rcvd,
			"zone_icmp_pkt_rcvd":                        ret.DtDdosDstZoneStats.Stats.Zone_icmp_pkt_rcvd,
			"zone_other_pkt_rcvd":                       ret.DtDdosDstZoneStats.Stats.Zone_other_pkt_rcvd,
			"zone_udp_filter_match":                     ret.DtDdosDstZoneStats.Stats.Zone_udp_filter_match,
			"zone_udp_filter_not_match":                 ret.DtDdosDstZoneStats.Stats.Zone_udp_filter_not_match,
			"zone_udp_filter_action_blacklist":          ret.DtDdosDstZoneStats.Stats.Zone_udp_filter_action_blacklist,
			"zone_udp_filter_action_drop":               ret.DtDdosDstZoneStats.Stats.Zone_udp_filter_action_drop,
			"zone_tcp_syn":                              ret.DtDdosDstZoneStats.Stats.Zone_tcp_syn,
			"zone_tcp_syn_drop":                         ret.DtDdosDstZoneStats.Stats.Zone_tcp_syn_drop,
			"zone_tcp_src_rate_drop":                    ret.DtDdosDstZoneStats.Stats.Zone_tcp_src_rate_drop,
			"zone_udp_src_rate_drop":                    ret.DtDdosDstZoneStats.Stats.Zone_udp_src_rate_drop,
			"zone_icmp_src_rate_drop":                   ret.DtDdosDstZoneStats.Stats.Zone_icmp_src_rate_drop,
			"zone_other_frag_src_rate_drop":             ret.DtDdosDstZoneStats.Stats.Zone_other_frag_src_rate_drop,
			"zone_other_src_rate_drop":                  ret.DtDdosDstZoneStats.Stats.Zone_other_src_rate_drop,
			"zone_tcp_drop":                             ret.DtDdosDstZoneStats.Stats.Zone_tcp_drop,
			"zone_udp_drop":                             ret.DtDdosDstZoneStats.Stats.Zone_udp_drop,
			"zone_icmp_drop":                            ret.DtDdosDstZoneStats.Stats.Zone_icmp_drop,
			"zone_frag_drop":                            ret.DtDdosDstZoneStats.Stats.Zone_frag_drop,
			"zone_other_drop":                           ret.DtDdosDstZoneStats.Stats.Zone_other_drop,
			"zone_tcp_auth":                             ret.DtDdosDstZoneStats.Stats.Zone_tcp_auth,
			"zone_udp_filter_action_default_pass":       ret.DtDdosDstZoneStats.Stats.Zone_udp_filter_action_default_pass,
			"zone_tcp_filter_match":                     ret.DtDdosDstZoneStats.Stats.Zone_tcp_filter_match,
			"zone_tcp_filter_not_match":                 ret.DtDdosDstZoneStats.Stats.Zone_tcp_filter_not_match,
			"zone_tcp_filter_action_blacklist":          ret.DtDdosDstZoneStats.Stats.Zone_tcp_filter_action_blacklist,
			"zone_tcp_filter_action_drop":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_filter_action_drop,
			"zone_tcp_filter_action_default_pass":       ret.DtDdosDstZoneStats.Stats.Zone_tcp_filter_action_default_pass,
			"zone_udp_filter_action_whitelist":          ret.DtDdosDstZoneStats.Stats.Zone_udp_filter_action_whitelist,
			"zone_udp_kibit_rate_drop":                  ret.DtDdosDstZoneStats.Stats.Zone_udp_kibit_rate_drop,
			"zone_tcp_kibit_rate_drop":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_kibit_rate_drop,
			"zone_icmp_kibit_rate_drop":                 ret.DtDdosDstZoneStats.Stats.Zone_icmp_kibit_rate_drop,
			"zone_other_kibit_rate_drop":                ret.DtDdosDstZoneStats.Stats.Zone_other_kibit_rate_drop,
			"zone_port_undef_drop":                      ret.DtDdosDstZoneStats.Stats.Zone_port_undef_drop,
			"zone_port_bl":                              ret.DtDdosDstZoneStats.Stats.Zone_port_bl,
			"zone_src_port_bl":                          ret.DtDdosDstZoneStats.Stats.Zone_src_port_bl,
			"zone_port_kbit_rate_exceed":                ret.DtDdosDstZoneStats.Stats.Zone_port_kbit_rate_exceed,
			"zone_tcp_src_drop":                         ret.DtDdosDstZoneStats.Stats.Zone_tcp_src_drop,
			"zone_udp_src_drop":                         ret.DtDdosDstZoneStats.Stats.Zone_udp_src_drop,
			"zone_icmp_src_drop":                        ret.DtDdosDstZoneStats.Stats.Zone_icmp_src_drop,
			"zone_other_src_drop":                       ret.DtDdosDstZoneStats.Stats.Zone_other_src_drop,
			"tcp_syn_rcvd":                              ret.DtDdosDstZoneStats.Stats.Tcp_syn_rcvd,
			"tcp_syn_ack_rcvd":                          ret.DtDdosDstZoneStats.Stats.Tcp_syn_ack_rcvd,
			"tcp_ack_rcvd":                              ret.DtDdosDstZoneStats.Stats.Tcp_ack_rcvd,
			"tcp_fin_rcvd":                              ret.DtDdosDstZoneStats.Stats.Tcp_fin_rcvd,
			"tcp_rst_rcvd":                              ret.DtDdosDstZoneStats.Stats.Tcp_rst_rcvd,
			"ingress_bytes":                             ret.DtDdosDstZoneStats.Stats.Ingress_bytes,
			"egress_bytes":                              ret.DtDdosDstZoneStats.Stats.Egress_bytes,
			"ingress_packets":                           ret.DtDdosDstZoneStats.Stats.Ingress_packets,
			"egress_packets":                            ret.DtDdosDstZoneStats.Stats.Egress_packets,
			"tcp_fwd_recv":                              ret.DtDdosDstZoneStats.Stats.Tcp_fwd_recv,
			"udp_fwd_recv":                              ret.DtDdosDstZoneStats.Stats.Udp_fwd_recv,
			"icmp_fwd_recv":                             ret.DtDdosDstZoneStats.Stats.Icmp_fwd_recv,
			"tcp_syn_cookie_fail":                       ret.DtDdosDstZoneStats.Stats.Tcp_syn_cookie_fail,
			"zone_tcp_session_created":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_session_created,
			"zone_udp_session_created":                  ret.DtDdosDstZoneStats.Stats.Zone_udp_session_created,
			"zone_tcp_filter_action_whitelist":          ret.DtDdosDstZoneStats.Stats.Zone_tcp_filter_action_whitelist,
			"zone_other_filter_match":                   ret.DtDdosDstZoneStats.Stats.Zone_other_filter_match,
			"zone_other_filter_not_match":               ret.DtDdosDstZoneStats.Stats.Zone_other_filter_not_match,
			"zone_other_filter_action_blacklist":        ret.DtDdosDstZoneStats.Stats.Zone_other_filter_action_blacklist,
			"zone_other_filter_action_drop":             ret.DtDdosDstZoneStats.Stats.Zone_other_filter_action_drop,
			"zone_other_filter_action_whitelist":        ret.DtDdosDstZoneStats.Stats.Zone_other_filter_action_whitelist,
			"zone_other_filter_action_default_pass":     ret.DtDdosDstZoneStats.Stats.Zone_other_filter_action_default_pass,
			"zone_blackhole_inject":                     ret.DtDdosDstZoneStats.Stats.Zone_blackhole_inject,
			"zone_blackhole_withdraw":                   ret.DtDdosDstZoneStats.Stats.Zone_blackhole_withdraw,
			"zone_tcp_out_of_seq_excd":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_out_of_seq_excd,
			"zone_tcp_retransmit_excd":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_retransmit_excd,
			"zone_tcp_zero_window_excd":                 ret.DtDdosDstZoneStats.Stats.Zone_tcp_zero_window_excd,
			"zone_tcp_conn_prate_excd":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_conn_prate_excd,
			"zone_tcp_action_on_ack_init":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_init,
			"zone_tcp_action_on_ack_gap_drop":           ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_gap_drop,
			"zone_tcp_action_on_ack_fail":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_fail,
			"zone_tcp_action_on_ack_pass":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_pass,
			"zone_tcp_action_on_syn_init":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_init,
			"zone_tcp_action_on_syn_gap_drop":           ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_gap_drop,
			"zone_tcp_action_on_syn_fail":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_fail,
			"zone_tcp_action_on_syn_pass":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_pass,
			"zone_payload_too_small":                    ret.DtDdosDstZoneStats.Stats.Zone_payload_too_small,
			"zone_payload_too_big":                      ret.DtDdosDstZoneStats.Stats.Zone_payload_too_big,
			"zone_udp_conn_prate_excd":                  ret.DtDdosDstZoneStats.Stats.Zone_udp_conn_prate_excd,
			"zone_udp_ntp_monlist_req":                  ret.DtDdosDstZoneStats.Stats.Zone_udp_ntp_monlist_req,
			"zone_udp_ntp_monlist_resp":                 ret.DtDdosDstZoneStats.Stats.Zone_udp_ntp_monlist_resp,
			"zone_udp_wellknown_sport_drop":             ret.DtDdosDstZoneStats.Stats.Zone_udp_wellknown_sport_drop,
			"zone_udp_retry_init":                       ret.DtDdosDstZoneStats.Stats.Zone_udp_retry_init,
			"zone_udp_retry_pass":                       ret.DtDdosDstZoneStats.Stats.Zone_udp_retry_pass,
			"zone_tcp_bytes_drop":                       ret.DtDdosDstZoneStats.Stats.Zone_tcp_bytes_drop,
			"zone_udp_bytes_drop":                       ret.DtDdosDstZoneStats.Stats.Zone_udp_bytes_drop,
			"zone_icmp_bytes_drop":                      ret.DtDdosDstZoneStats.Stats.Zone_icmp_bytes_drop,
			"zone_other_bytes_drop":                     ret.DtDdosDstZoneStats.Stats.Zone_other_bytes_drop,
			"zone_out_no_route":                         ret.DtDdosDstZoneStats.Stats.Zone_out_no_route,
			"outbound_bytes_sent":                       ret.DtDdosDstZoneStats.Stats.Outbound_bytes_sent,
			"outbound_drop":                             ret.DtDdosDstZoneStats.Stats.Outbound_drop,
			"outbound_bytes_drop":                       ret.DtDdosDstZoneStats.Stats.Outbound_bytes_drop,
			"outbound_pkt_sent":                         ret.DtDdosDstZoneStats.Stats.Outbound_pkt_sent,
			"inbound_bytes_sent":                        ret.DtDdosDstZoneStats.Stats.Inbound_bytes_sent,
			"inbound_bytes_drop":                        ret.DtDdosDstZoneStats.Stats.Inbound_bytes_drop,
			"zone_src_port_pkt_rate_exceed":             ret.DtDdosDstZoneStats.Stats.Zone_src_port_pkt_rate_exceed,
			"zone_src_port_kbit_rate_exceed":            ret.DtDdosDstZoneStats.Stats.Zone_src_port_kbit_rate_exceed,
			"zone_src_port_conn_limit_exceed":           ret.DtDdosDstZoneStats.Stats.Zone_src_port_conn_limit_exceed,
			"zone_src_port_conn_rate_exceed":            ret.DtDdosDstZoneStats.Stats.Zone_src_port_conn_rate_exceed,
			"zone_ip_proto_pkt_rate_exceed":             ret.DtDdosDstZoneStats.Stats.Zone_ip_proto_pkt_rate_exceed,
			"zone_ip_proto_kbit_rate_exceed":            ret.DtDdosDstZoneStats.Stats.Zone_ip_proto_kbit_rate_exceed,
			"zone_tcp_port_any_exceed":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_port_any_exceed,
			"zone_udp_port_any_exceed":                  ret.DtDdosDstZoneStats.Stats.Zone_udp_port_any_exceed,
			"zone_tcp_auth_pass":                        ret.DtDdosDstZoneStats.Stats.Zone_tcp_auth_pass,
			"zone_tcp_rst_cookie_fail":                  ret.DtDdosDstZoneStats.Stats.Zone_tcp_rst_cookie_fail,
			"zone_tcp_unauth_drop":                      ret.DtDdosDstZoneStats.Stats.Zone_tcp_unauth_drop,
			"src_tcp_syn_auth_fail":                     ret.DtDdosDstZoneStats.Stats.Src_tcp_syn_auth_fail,
			"src_tcp_syn_cookie_sent":                   ret.DtDdosDstZoneStats.Stats.Src_tcp_syn_cookie_sent,
			"src_tcp_syn_cookie_fail":                   ret.DtDdosDstZoneStats.Stats.Src_tcp_syn_cookie_fail,
			"src_tcp_rst_cookie_fail":                   ret.DtDdosDstZoneStats.Stats.Src_tcp_rst_cookie_fail,
			"src_tcp_unauth_drop":                       ret.DtDdosDstZoneStats.Stats.Src_tcp_unauth_drop,
			"src_tcp_action_on_syn_init":                ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_syn_init,
			"src_tcp_action_on_syn_gap_drop":            ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_syn_gap_drop,
			"src_tcp_action_on_syn_fail":                ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_syn_fail,
			"src_tcp_action_on_ack_init":                ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_ack_init,
			"src_tcp_action_on_ack_gap_drop":            ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_ack_gap_drop,
			"src_tcp_action_on_ack_fail":                ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_ack_fail,
			"src_tcp_out_of_seq_excd":                   ret.DtDdosDstZoneStats.Stats.Src_tcp_out_of_seq_excd,
			"src_tcp_retransmit_excd":                   ret.DtDdosDstZoneStats.Stats.Src_tcp_retransmit_excd,
			"src_tcp_zero_window_excd":                  ret.DtDdosDstZoneStats.Stats.Src_tcp_zero_window_excd,
			"src_tcp_conn_prate_excd":                   ret.DtDdosDstZoneStats.Stats.Src_tcp_conn_prate_excd,
			"src_udp_min_payload":                       ret.DtDdosDstZoneStats.Stats.Src_udp_min_payload,
			"src_udp_max_payload":                       ret.DtDdosDstZoneStats.Stats.Src_udp_max_payload,
			"src_udp_conn_prate_excd":                   ret.DtDdosDstZoneStats.Stats.Src_udp_conn_prate_excd,
			"src_udp_ntp_monlist_req":                   ret.DtDdosDstZoneStats.Stats.Src_udp_ntp_monlist_req,
			"src_udp_ntp_monlist_resp":                  ret.DtDdosDstZoneStats.Stats.Src_udp_ntp_monlist_resp,
			"src_udp_wellknown_sport_drop":              ret.DtDdosDstZoneStats.Stats.Src_udp_wellknown_sport_drop,
			"src_udp_retry_init":                        ret.DtDdosDstZoneStats.Stats.Src_udp_retry_init,
			"dst_udp_retry_gap_drop":                    ret.DtDdosDstZoneStats.Stats.Dst_udp_retry_gap_drop,
			"dst_udp_retry_fail":                        ret.DtDdosDstZoneStats.Stats.Dst_udp_retry_fail,
			"dst_tcp_session_aged":                      ret.DtDdosDstZoneStats.Stats.Dst_tcp_session_aged,
			"dst_udp_session_aged":                      ret.DtDdosDstZoneStats.Stats.Dst_udp_session_aged,
			"dst_tcp_conn_close":                        ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_close,
			"dst_tcp_conn_close_half_open":              ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_close_half_open,
			"src_tcp_filter_action_blacklist":           ret.DtDdosDstZoneStats.Stats.Src_tcp_filter_action_blacklist,
			"src_tcp_filter_action_whitelist":           ret.DtDdosDstZoneStats.Stats.Src_tcp_filter_action_whitelist,
			"src_tcp_filter_action_drop":                ret.DtDdosDstZoneStats.Stats.Src_tcp_filter_action_drop,
			"src_tcp_filter_action_default_pass":        ret.DtDdosDstZoneStats.Stats.Src_tcp_filter_action_default_pass,
			"src_udp_filter_action_blacklist":           ret.DtDdosDstZoneStats.Stats.Src_udp_filter_action_blacklist,
			"src_udp_filter_action_whitelist":           ret.DtDdosDstZoneStats.Stats.Src_udp_filter_action_whitelist,
			"src_udp_filter_action_drop":                ret.DtDdosDstZoneStats.Stats.Src_udp_filter_action_drop,
			"src_udp_filter_action_default_pass":        ret.DtDdosDstZoneStats.Stats.Src_udp_filter_action_default_pass,
			"src_other_filter_action_blacklist":         ret.DtDdosDstZoneStats.Stats.Src_other_filter_action_blacklist,
			"src_other_filter_action_whitelist":         ret.DtDdosDstZoneStats.Stats.Src_other_filter_action_whitelist,
			"src_other_filter_action_drop":              ret.DtDdosDstZoneStats.Stats.Src_other_filter_action_drop,
			"src_other_filter_action_default_pass":      ret.DtDdosDstZoneStats.Stats.Src_other_filter_action_default_pass,
			"tcp_invalid_syn":                           ret.DtDdosDstZoneStats.Stats.Tcp_invalid_syn,
			"dst_tcp_conn_close_w_rst":                  ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_close_w_rst,
			"dst_tcp_conn_close_w_fin":                  ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_close_w_fin,
			"dst_tcp_conn_close_w_idle":                 ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_close_w_idle,
			"dst_tcp_conn_create_from_syn":              ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_create_from_syn,
			"dst_tcp_conn_create_from_ack":              ret.DtDdosDstZoneStats.Stats.Dst_tcp_conn_create_from_ack,
			"src_frag_drop":                             ret.DtDdosDstZoneStats.Stats.Src_frag_drop,
			"zone_port_kbit_rate_exceed_pkt":            ret.DtDdosDstZoneStats.Stats.Zone_port_kbit_rate_exceed_pkt,
			"dst_tcp_bytes_rcv":                         ret.DtDdosDstZoneStats.Stats.Dst_tcp_bytes_rcv,
			"dst_udp_bytes_rcv":                         ret.DtDdosDstZoneStats.Stats.Dst_udp_bytes_rcv,
			"dst_icmp_bytes_rcv":                        ret.DtDdosDstZoneStats.Stats.Dst_icmp_bytes_rcv,
			"dst_other_bytes_rcv":                       ret.DtDdosDstZoneStats.Stats.Dst_other_bytes_rcv,
			"dst_tcp_bytes_sent":                        ret.DtDdosDstZoneStats.Stats.Dst_tcp_bytes_sent,
			"dst_udp_bytes_sent":                        ret.DtDdosDstZoneStats.Stats.Dst_udp_bytes_sent,
			"dst_icmp_bytes_sent":                       ret.DtDdosDstZoneStats.Stats.Dst_icmp_bytes_sent,
			"dst_other_bytes_sent":                      ret.DtDdosDstZoneStats.Stats.Dst_other_bytes_sent,
			"dst_udp_auth_drop":                         ret.DtDdosDstZoneStats.Stats.Dst_udp_auth_drop,
			"dst_tcp_auth_drop":                         ret.DtDdosDstZoneStats.Stats.Dst_tcp_auth_drop,
			"dst_tcp_auth_resp":                         ret.DtDdosDstZoneStats.Stats.Dst_tcp_auth_resp,
			"dst_drop":                                  ret.DtDdosDstZoneStats.Stats.Dst_drop,
			"dst_entry_pkt_rate_exceed":                 ret.DtDdosDstZoneStats.Stats.Dst_entry_pkt_rate_exceed,
			"dst_entry_kbit_rate_exceed":                ret.DtDdosDstZoneStats.Stats.Dst_entry_kbit_rate_exceed,
			"dst_entry_conn_limit_exceed":               ret.DtDdosDstZoneStats.Stats.Dst_entry_conn_limit_exceed,
			"dst_entry_conn_rate_exceed":                ret.DtDdosDstZoneStats.Stats.Dst_entry_conn_rate_exceed,
			"dst_entry_frag_pkt_rate_exceed":            ret.DtDdosDstZoneStats.Stats.Dst_entry_frag_pkt_rate_exceed,
			"dst_l4_tcp_blacklist_drop":                 ret.DtDdosDstZoneStats.Stats.Dst_l4_tcp_blacklist_drop,
			"dst_l4_udp_blacklist_drop":                 ret.DtDdosDstZoneStats.Stats.Dst_l4_udp_blacklist_drop,
			"dst_l4_icmp_blacklist_drop":                ret.DtDdosDstZoneStats.Stats.Dst_l4_icmp_blacklist_drop,
			"dst_l4_other_blacklist_drop":               ret.DtDdosDstZoneStats.Stats.Dst_l4_other_blacklist_drop,
			"dst_icmp_any_exceed":                       ret.DtDdosDstZoneStats.Stats.Dst_icmp_any_exceed,
			"dst_other_any_exceed":                      ret.DtDdosDstZoneStats.Stats.Dst_other_any_exceed,
			"tcp_rexmit_syn_limit_drop":                 ret.DtDdosDstZoneStats.Stats.Tcp_rexmit_syn_limit_drop,
			"tcp_rexmit_syn_limit_bl":                   ret.DtDdosDstZoneStats.Stats.Tcp_rexmit_syn_limit_bl,
			"dst_clist_overflow_policy_at_learning":     ret.DtDdosDstZoneStats.Stats.Dst_clist_overflow_policy_at_learning,
			"zone_frag_rcvd":                            ret.DtDdosDstZoneStats.Stats.Zone_frag_rcvd,
			"zone_tcp_wellknown_sport_drop":             ret.DtDdosDstZoneStats.Stats.Zone_tcp_wellknown_sport_drop,
			"src_tcp_wellknown_sport_drop":              ret.DtDdosDstZoneStats.Stats.Src_tcp_wellknown_sport_drop,
			"secondary_dst_entry_pkt_rate_exceed":       ret.DtDdosDstZoneStats.Stats.Secondary_dst_entry_pkt_rate_exceed,
			"secondary_dst_entry_kbit_rate_exceed":      ret.DtDdosDstZoneStats.Stats.Secondary_dst_entry_kbit_rate_exceed,
			"secondary_dst_entry_conn_limit_exceed":     ret.DtDdosDstZoneStats.Stats.Secondary_dst_entry_conn_limit_exceed,
			"secondary_dst_entry_conn_rate_exceed":      ret.DtDdosDstZoneStats.Stats.Secondary_dst_entry_conn_rate_exceed,
			"secondary_dst_entry_frag_pkt_rate_exceed":  ret.DtDdosDstZoneStats.Stats.Secondary_dst_entry_frag_pkt_rate_exceed,
			"src_udp_retry_gap_drop":                    ret.DtDdosDstZoneStats.Stats.Src_udp_retry_gap_drop,
			"dst_entry_kbit_rate_exceed_count":          ret.DtDdosDstZoneStats.Stats.Dst_entry_kbit_rate_exceed_count,
			"secondary_entry_learn":                     ret.DtDdosDstZoneStats.Stats.Secondary_entry_learn,
			"secondary_entry_hit":                       ret.DtDdosDstZoneStats.Stats.Secondary_entry_hit,
			"secondary_entry_miss":                      ret.DtDdosDstZoneStats.Stats.Secondary_entry_miss,
			"secondary_entry_aged":                      ret.DtDdosDstZoneStats.Stats.Secondary_entry_aged,
			"secondary_entry_learning_thre_exceed":      ret.DtDdosDstZoneStats.Stats.Secondary_entry_learning_thre_exceed,
			"zone_port_undef_hit":                       ret.DtDdosDstZoneStats.Stats.Zone_port_undef_hit,
			"zone_tcp_action_on_ack_timeout":            ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_timeout,
			"zone_tcp_action_on_ack_reset":              ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_reset,
			"zone_tcp_action_on_ack_blacklist":          ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_ack_blacklist,
			"src_tcp_action_on_ack_timeout":             ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_ack_timeout,
			"src_tcp_action_on_ack_reset":               ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_ack_reset,
			"src_tcp_action_on_ack_blacklist":           ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_ack_blacklist,
			"zone_tcp_action_on_syn_timeout":            ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_timeout,
			"zone_tcp_action_on_syn_reset":              ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_reset,
			"zone_tcp_action_on_syn_blacklist":          ret.DtDdosDstZoneStats.Stats.Zone_tcp_action_on_syn_blacklist,
			"src_tcp_action_on_syn_timeout":             ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_syn_timeout,
			"src_tcp_action_on_syn_reset":               ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_syn_reset,
			"src_tcp_action_on_syn_blacklist":           ret.DtDdosDstZoneStats.Stats.Src_tcp_action_on_syn_blacklist,
			"zone_udp_frag_pkt_rate_exceed":             ret.DtDdosDstZoneStats.Stats.Zone_udp_frag_pkt_rate_exceed,
			"zone_udp_frag_src_rate_drop":               ret.DtDdosDstZoneStats.Stats.Zone_udp_frag_src_rate_drop,
			"zone_tcp_frag_pkt_rate_exceed":             ret.DtDdosDstZoneStats.Stats.Zone_tcp_frag_pkt_rate_exceed,
			"zone_tcp_frag_src_rate_drop":               ret.DtDdosDstZoneStats.Stats.Zone_tcp_frag_src_rate_drop,
			"zone_icmp_frag_pkt_rate_exceed":            ret.DtDdosDstZoneStats.Stats.Zone_icmp_frag_pkt_rate_exceed,
			"zone_icmp_frag_src_rate_drop":              ret.DtDdosDstZoneStats.Stats.Zone_icmp_frag_src_rate_drop,
			"sflow_internal_samples_packed":             ret.DtDdosDstZoneStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":             ret.DtDdosDstZoneStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":               ret.DtDdosDstZoneStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":               ret.DtDdosDstZoneStats.Stats.Sflow_external_packets_sent,
			"dns_outbound_total_query":                  ret.DtDdosDstZoneStats.Stats.Dns_outbound_total_query,
			"dns_outbound_query_malformed":              ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_malformed,
			"dns_outbound_query_resp_chk_failed":        ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_resp_chk_failed,
			"dns_outbound_query_resp_chk_blacklisted":   ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_resp_chk_blacklisted,
			"dns_outbound_query_resp_chk_refused_sent":  ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_resp_chk_refused_sent,
			"dns_outbound_query_resp_chk_reset_sent":    ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_resp_chk_reset_sent,
			"dns_outbound_query_resp_chk_no_resp_sent":  ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_resp_chk_no_resp_sent,
			"dns_outbound_query_resp_size_exceed":       ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_resp_size_exceed,
			"dns_outbound_query_sess_timed_out":         ret.DtDdosDstZoneStats.Stats.Dns_outbound_query_sess_timed_out,
			"source_entry_total":                        ret.DtDdosDstZoneStats.Stats.Source_entry_total,
			"source_entry_udp":                          ret.DtDdosDstZoneStats.Stats.Source_entry_udp,
			"source_entry_tcp":                          ret.DtDdosDstZoneStats.Stats.Source_entry_tcp,
			"source_entry_icmp":                         ret.DtDdosDstZoneStats.Stats.Source_entry_icmp,
			"source_entry_other":                        ret.DtDdosDstZoneStats.Stats.Source_entry_other,
			"dst_exceed_action_tunnel":                  ret.DtDdosDstZoneStats.Stats.Dst_exceed_action_tunnel,
			"dst_udp_retry_timeout_blacklist":           ret.DtDdosDstZoneStats.Stats.Dst_udp_retry_timeout_blacklist,
			"src_udp_auth_timeout":                      ret.DtDdosDstZoneStats.Stats.Src_udp_auth_timeout,
			"zone_src_udp_retry_timeout_blacklist":      ret.DtDdosDstZoneStats.Stats.Zone_src_udp_retry_timeout_blacklist,
			"src_udp_retry_pass":                        ret.DtDdosDstZoneStats.Stats.Src_udp_retry_pass,
			"secondary_port_learn":                      ret.DtDdosDstZoneStats.Stats.Secondary_port_learn,
			"secondary_port_aged":                       ret.DtDdosDstZoneStats.Stats.Secondary_port_aged,
			"dst_entry_outbound_udp_session_created":    ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_udp_session_created,
			"dst_entry_outbound_udp_session_aged":       ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_udp_session_aged,
			"dst_entry_outbound_tcp_session_created":    ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_tcp_session_created,
			"dst_entry_outbound_tcp_session_aged":       ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_tcp_session_aged,
			"dst_entry_outbound_pkt_rate_exceed":        ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_pkt_rate_exceed,
			"dst_entry_outbound_kbit_rate_exceed":       ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_kbit_rate_exceed,
			"dst_entry_outbound_kbit_rate_exceed_count": ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_kbit_rate_exceed_count,
			"dst_entry_outbound_conn_limit_exceed":      ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_conn_limit_exceed,
			"dst_entry_outbound_conn_rate_exceed":       ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_conn_rate_exceed,
			"dst_entry_outbound_frag_pkt_rate_exceed":   ret.DtDdosDstZoneStats.Stats.Dst_entry_outbound_frag_pkt_rate_exceed,
			"prog_first_req_time_exceed":                ret.DtDdosDstZoneStats.Stats.Prog_first_req_time_exceed,
			"prog_req_resp_time_exceed":                 ret.DtDdosDstZoneStats.Stats.Prog_req_resp_time_exceed,
			"prog_request_len_exceed":                   ret.DtDdosDstZoneStats.Stats.Prog_request_len_exceed,
			"prog_response_len_exceed":                  ret.DtDdosDstZoneStats.Stats.Prog_response_len_exceed,
			"prog_resp_req_ratio_exceed":                ret.DtDdosDstZoneStats.Stats.Prog_resp_req_ratio_exceed,
			"prog_resp_req_time_exceed":                 ret.DtDdosDstZoneStats.Stats.Prog_resp_req_time_exceed,
			"entry_sync_message_received":               ret.DtDdosDstZoneStats.Stats.Entry_sync_message_received,
			"entry_sync_message_sent":                   ret.DtDdosDstZoneStats.Stats.Entry_sync_message_sent,
			"prog_conn_sent_exceed":                     ret.DtDdosDstZoneStats.Stats.Prog_conn_sent_exceed,
			"prog_conn_rcvd_exceed":                     ret.DtDdosDstZoneStats.Stats.Prog_conn_rcvd_exceed,
			"prog_conn_time_exceed":                     ret.DtDdosDstZoneStats.Stats.Prog_conn_time_exceed,
			"prog_conn_rcvd_sent_ratio_exceed":          ret.DtDdosDstZoneStats.Stats.Prog_conn_rcvd_sent_ratio_exceed,
			"prog_win_sent_exceed":                      ret.DtDdosDstZoneStats.Stats.Prog_win_sent_exceed,
			"prog_win_rcvd_exceed":                      ret.DtDdosDstZoneStats.Stats.Prog_win_rcvd_exceed,
			"prog_win_rcvd_sent_ratio_exceed":           ret.DtDdosDstZoneStats.Stats.Prog_win_rcvd_sent_ratio_exceed,
			"prog_exceed_drop":                          ret.DtDdosDstZoneStats.Stats.Prog_exceed_drop,
			"prog_exceed_bl":                            ret.DtDdosDstZoneStats.Stats.Prog_exceed_bl,
			"prog_conn_exceed_drop":                     ret.DtDdosDstZoneStats.Stats.Prog_conn_exceed_drop,
			"prog_conn_exceed_bl":                       ret.DtDdosDstZoneStats.Stats.Prog_conn_exceed_bl,
			"prog_win_exceed_drop":                      ret.DtDdosDstZoneStats.Stats.Prog_win_exceed_drop,
			"prog_win_exceed_bl":                        ret.DtDdosDstZoneStats.Stats.Prog_win_exceed_bl,
			"east_west_inbound_rcv_pkt":                 ret.DtDdosDstZoneStats.Stats.East_west_inbound_rcv_pkt,
			"east_west_inbound_drop_pkt":                ret.DtDdosDstZoneStats.Stats.East_west_inbound_drop_pkt,
			"east_west_inbound_fwd_pkt":                 ret.DtDdosDstZoneStats.Stats.East_west_inbound_fwd_pkt,
			"east_west_inbound_rcv_byte":                ret.DtDdosDstZoneStats.Stats.East_west_inbound_rcv_byte,
			"east_west_inbound_drop_byte":               ret.DtDdosDstZoneStats.Stats.East_west_inbound_drop_byte,
			"east_west_inbound_fwd_byte":                ret.DtDdosDstZoneStats.Stats.East_west_inbound_fwd_byte,
			"east_west_outbound_rcv_pkt":                ret.DtDdosDstZoneStats.Stats.East_west_outbound_rcv_pkt,
			"east_west_outbound_drop_pkt":               ret.DtDdosDstZoneStats.Stats.East_west_outbound_drop_pkt,
			"east_west_outbound_fwd_pkt":                ret.DtDdosDstZoneStats.Stats.East_west_outbound_fwd_pkt,
			"east_west_outbound_rcv_byte":               ret.DtDdosDstZoneStats.Stats.East_west_outbound_rcv_byte,
			"east_west_outbound_drop_byte":              ret.DtDdosDstZoneStats.Stats.East_west_outbound_drop_byte,
			"east_west_outbound_fwd_byte":               ret.DtDdosDstZoneStats.Stats.East_west_outbound_fwd_byte,
			"dst_exceed_action_drop":                    ret.DtDdosDstZoneStats.Stats.Dst_exceed_action_drop,
			"prog_conn_samples":                         ret.DtDdosDstZoneStats.Stats.Prog_conn_samples,
			"prog_req_samples":                          ret.DtDdosDstZoneStats.Stats.Prog_req_samples,
			"prog_win_samples":                          ret.DtDdosDstZoneStats.Stats.Prog_win_samples,
			"victim_ip_learned":                         ret.DtDdosDstZoneStats.Stats.Victim_ip_learned,
			"victim_ip_aged":                            ret.DtDdosDstZoneStats.Stats.Victim_ip_aged,
			"prog_conn_samples_processed":               ret.DtDdosDstZoneStats.Stats.Prog_conn_samples_processed,
			"prog_req_samples_processed":                ret.DtDdosDstZoneStats.Stats.Prog_req_samples_processed,
			"prog_win_samples_processed":                ret.DtDdosDstZoneStats.Stats.Prog_win_samples_processed,
			"dst_src_learn_overflow":                    ret.DtDdosDstZoneStats.Stats.Dst_src_learn_overflow,
			"dst_tcp_auth_rst":                          ret.DtDdosDstZoneStats.Stats.Dst_tcp_auth_rst,
		},
	}
}

func getObjectDdosDstZoneStatsStats(d []interface{}) edpt.DdosDstZoneStatsStats {

	count1 := len(d)
	var ret edpt.DdosDstZoneStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Zone_tcp_any_exceed = in["zone_tcp_any_exceed"].(int)
		ret.Zone_tcp_pkt_rate_exceed = in["zone_tcp_pkt_rate_exceed"].(int)
		ret.Zone_tcp_conn_rate_exceed = in["zone_tcp_conn_rate_exceed"].(int)
		ret.Zone_udp_any_exceed = in["zone_udp_any_exceed"].(int)
		ret.Zone_udp_pkt_rate_exceed = in["zone_udp_pkt_rate_exceed"].(int)
		ret.Zone_udp_conn_limit_exceed = in["zone_udp_conn_limit_exceed"].(int)
		ret.Zone_udp_conn_rate_exceed = in["zone_udp_conn_rate_exceed"].(int)
		ret.Zone_icmp_pkt_rate_exceed = in["zone_icmp_pkt_rate_exceed"].(int)
		ret.Zone_other_pkt_rate_exceed = in["zone_other_pkt_rate_exceed"].(int)
		ret.Zone_other_frag_pkt_rate_exceed = in["zone_other_frag_pkt_rate_exceed"].(int)
		ret.Zone_port_pkt_rate_exceed = in["zone_port_pkt_rate_exceed"].(int)
		ret.Zone_port_conn_limit_exceed = in["zone_port_conn_limit_exceed"].(int)
		ret.Zone_port_conn_rate_exceed = in["zone_port_conn_rate_exceed"].(int)
		ret.Zone_pkt_sent = in["zone_pkt_sent"].(int)
		ret.Zone_udp_pkt_sent = in["zone_udp_pkt_sent"].(int)
		ret.Zone_tcp_pkt_sent = in["zone_tcp_pkt_sent"].(int)
		ret.Zone_icmp_pkt_sent = in["zone_icmp_pkt_sent"].(int)
		ret.Zone_other_pkt_sent = in["zone_other_pkt_sent"].(int)
		ret.Zone_tcp_conn_limit_exceed = in["zone_tcp_conn_limit_exceed"].(int)
		ret.Zone_tcp_pkt_rcvd = in["zone_tcp_pkt_rcvd"].(int)
		ret.Zone_udp_pkt_rcvd = in["zone_udp_pkt_rcvd"].(int)
		ret.Zone_icmp_pkt_rcvd = in["zone_icmp_pkt_rcvd"].(int)
		ret.Zone_other_pkt_rcvd = in["zone_other_pkt_rcvd"].(int)
		ret.Zone_udp_filter_match = in["zone_udp_filter_match"].(int)
		ret.Zone_udp_filter_not_match = in["zone_udp_filter_not_match"].(int)
		ret.Zone_udp_filter_action_blacklist = in["zone_udp_filter_action_blacklist"].(int)
		ret.Zone_udp_filter_action_drop = in["zone_udp_filter_action_drop"].(int)
		ret.Zone_tcp_syn = in["zone_tcp_syn"].(int)
		ret.Zone_tcp_syn_drop = in["zone_tcp_syn_drop"].(int)
		ret.Zone_tcp_src_rate_drop = in["zone_tcp_src_rate_drop"].(int)
		ret.Zone_udp_src_rate_drop = in["zone_udp_src_rate_drop"].(int)
		ret.Zone_icmp_src_rate_drop = in["zone_icmp_src_rate_drop"].(int)
		ret.Zone_other_frag_src_rate_drop = in["zone_other_frag_src_rate_drop"].(int)
		ret.Zone_other_src_rate_drop = in["zone_other_src_rate_drop"].(int)
		ret.Zone_tcp_drop = in["zone_tcp_drop"].(int)
		ret.Zone_udp_drop = in["zone_udp_drop"].(int)
		ret.Zone_icmp_drop = in["zone_icmp_drop"].(int)
		ret.Zone_frag_drop = in["zone_frag_drop"].(int)
		ret.Zone_other_drop = in["zone_other_drop"].(int)
		ret.Zone_tcp_auth = in["zone_tcp_auth"].(int)
		ret.Zone_udp_filter_action_default_pass = in["zone_udp_filter_action_default_pass"].(int)
		ret.Zone_tcp_filter_match = in["zone_tcp_filter_match"].(int)
		ret.Zone_tcp_filter_not_match = in["zone_tcp_filter_not_match"].(int)
		ret.Zone_tcp_filter_action_blacklist = in["zone_tcp_filter_action_blacklist"].(int)
		ret.Zone_tcp_filter_action_drop = in["zone_tcp_filter_action_drop"].(int)
		ret.Zone_tcp_filter_action_default_pass = in["zone_tcp_filter_action_default_pass"].(int)
		ret.Zone_udp_filter_action_whitelist = in["zone_udp_filter_action_whitelist"].(int)
		ret.Zone_udp_kibit_rate_drop = in["zone_udp_kibit_rate_drop"].(int)
		ret.Zone_tcp_kibit_rate_drop = in["zone_tcp_kibit_rate_drop"].(int)
		ret.Zone_icmp_kibit_rate_drop = in["zone_icmp_kibit_rate_drop"].(int)
		ret.Zone_other_kibit_rate_drop = in["zone_other_kibit_rate_drop"].(int)
		ret.Zone_port_undef_drop = in["zone_port_undef_drop"].(int)
		ret.Zone_port_bl = in["zone_port_bl"].(int)
		ret.Zone_src_port_bl = in["zone_src_port_bl"].(int)
		ret.Zone_port_kbit_rate_exceed = in["zone_port_kbit_rate_exceed"].(int)
		ret.Zone_tcp_src_drop = in["zone_tcp_src_drop"].(int)
		ret.Zone_udp_src_drop = in["zone_udp_src_drop"].(int)
		ret.Zone_icmp_src_drop = in["zone_icmp_src_drop"].(int)
		ret.Zone_other_src_drop = in["zone_other_src_drop"].(int)
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
		ret.Zone_tcp_session_created = in["zone_tcp_session_created"].(int)
		ret.Zone_udp_session_created = in["zone_udp_session_created"].(int)
		ret.Zone_tcp_filter_action_whitelist = in["zone_tcp_filter_action_whitelist"].(int)
		ret.Zone_other_filter_match = in["zone_other_filter_match"].(int)
		ret.Zone_other_filter_not_match = in["zone_other_filter_not_match"].(int)
		ret.Zone_other_filter_action_blacklist = in["zone_other_filter_action_blacklist"].(int)
		ret.Zone_other_filter_action_drop = in["zone_other_filter_action_drop"].(int)
		ret.Zone_other_filter_action_whitelist = in["zone_other_filter_action_whitelist"].(int)
		ret.Zone_other_filter_action_default_pass = in["zone_other_filter_action_default_pass"].(int)
		ret.Zone_blackhole_inject = in["zone_blackhole_inject"].(int)
		ret.Zone_blackhole_withdraw = in["zone_blackhole_withdraw"].(int)
		ret.Zone_tcp_out_of_seq_excd = in["zone_tcp_out_of_seq_excd"].(int)
		ret.Zone_tcp_retransmit_excd = in["zone_tcp_retransmit_excd"].(int)
		ret.Zone_tcp_zero_window_excd = in["zone_tcp_zero_window_excd"].(int)
		ret.Zone_tcp_conn_prate_excd = in["zone_tcp_conn_prate_excd"].(int)
		ret.Zone_tcp_action_on_ack_init = in["zone_tcp_action_on_ack_init"].(int)
		ret.Zone_tcp_action_on_ack_gap_drop = in["zone_tcp_action_on_ack_gap_drop"].(int)
		ret.Zone_tcp_action_on_ack_fail = in["zone_tcp_action_on_ack_fail"].(int)
		ret.Zone_tcp_action_on_ack_pass = in["zone_tcp_action_on_ack_pass"].(int)
		ret.Zone_tcp_action_on_syn_init = in["zone_tcp_action_on_syn_init"].(int)
		ret.Zone_tcp_action_on_syn_gap_drop = in["zone_tcp_action_on_syn_gap_drop"].(int)
		ret.Zone_tcp_action_on_syn_fail = in["zone_tcp_action_on_syn_fail"].(int)
		ret.Zone_tcp_action_on_syn_pass = in["zone_tcp_action_on_syn_pass"].(int)
		ret.Zone_payload_too_small = in["zone_payload_too_small"].(int)
		ret.Zone_payload_too_big = in["zone_payload_too_big"].(int)
		ret.Zone_udp_conn_prate_excd = in["zone_udp_conn_prate_excd"].(int)
		ret.Zone_udp_ntp_monlist_req = in["zone_udp_ntp_monlist_req"].(int)
		ret.Zone_udp_ntp_monlist_resp = in["zone_udp_ntp_monlist_resp"].(int)
		ret.Zone_udp_wellknown_sport_drop = in["zone_udp_wellknown_sport_drop"].(int)
		ret.Zone_udp_retry_init = in["zone_udp_retry_init"].(int)
		ret.Zone_udp_retry_pass = in["zone_udp_retry_pass"].(int)
		ret.Zone_tcp_bytes_drop = in["zone_tcp_bytes_drop"].(int)
		ret.Zone_udp_bytes_drop = in["zone_udp_bytes_drop"].(int)
		ret.Zone_icmp_bytes_drop = in["zone_icmp_bytes_drop"].(int)
		ret.Zone_other_bytes_drop = in["zone_other_bytes_drop"].(int)
		ret.Zone_out_no_route = in["zone_out_no_route"].(int)
		ret.Outbound_bytes_sent = in["outbound_bytes_sent"].(int)
		ret.Outbound_drop = in["outbound_drop"].(int)
		ret.Outbound_bytes_drop = in["outbound_bytes_drop"].(int)
		ret.Outbound_pkt_sent = in["outbound_pkt_sent"].(int)
		ret.Inbound_bytes_sent = in["inbound_bytes_sent"].(int)
		ret.Inbound_bytes_drop = in["inbound_bytes_drop"].(int)
		ret.Zone_src_port_pkt_rate_exceed = in["zone_src_port_pkt_rate_exceed"].(int)
		ret.Zone_src_port_kbit_rate_exceed = in["zone_src_port_kbit_rate_exceed"].(int)
		ret.Zone_src_port_conn_limit_exceed = in["zone_src_port_conn_limit_exceed"].(int)
		ret.Zone_src_port_conn_rate_exceed = in["zone_src_port_conn_rate_exceed"].(int)
		ret.Zone_ip_proto_pkt_rate_exceed = in["zone_ip_proto_pkt_rate_exceed"].(int)
		ret.Zone_ip_proto_kbit_rate_exceed = in["zone_ip_proto_kbit_rate_exceed"].(int)
		ret.Zone_tcp_port_any_exceed = in["zone_tcp_port_any_exceed"].(int)
		ret.Zone_udp_port_any_exceed = in["zone_udp_port_any_exceed"].(int)
		ret.Zone_tcp_auth_pass = in["zone_tcp_auth_pass"].(int)
		ret.Zone_tcp_rst_cookie_fail = in["zone_tcp_rst_cookie_fail"].(int)
		ret.Zone_tcp_unauth_drop = in["zone_tcp_unauth_drop"].(int)
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
		ret.Zone_port_kbit_rate_exceed_pkt = in["zone_port_kbit_rate_exceed_pkt"].(int)
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
		ret.Dst_drop = in["dst_drop"].(int)
		ret.Dst_entry_pkt_rate_exceed = in["dst_entry_pkt_rate_exceed"].(int)
		ret.Dst_entry_kbit_rate_exceed = in["dst_entry_kbit_rate_exceed"].(int)
		ret.Dst_entry_conn_limit_exceed = in["dst_entry_conn_limit_exceed"].(int)
		ret.Dst_entry_conn_rate_exceed = in["dst_entry_conn_rate_exceed"].(int)
		ret.Dst_entry_frag_pkt_rate_exceed = in["dst_entry_frag_pkt_rate_exceed"].(int)
		ret.Dst_l4_tcp_blacklist_drop = in["dst_l4_tcp_blacklist_drop"].(int)
		ret.Dst_l4_udp_blacklist_drop = in["dst_l4_udp_blacklist_drop"].(int)
		ret.Dst_l4_icmp_blacklist_drop = in["dst_l4_icmp_blacklist_drop"].(int)
		ret.Dst_l4_other_blacklist_drop = in["dst_l4_other_blacklist_drop"].(int)
		ret.Dst_icmp_any_exceed = in["dst_icmp_any_exceed"].(int)
		ret.Dst_other_any_exceed = in["dst_other_any_exceed"].(int)
		ret.Tcp_rexmit_syn_limit_drop = in["tcp_rexmit_syn_limit_drop"].(int)
		ret.Tcp_rexmit_syn_limit_bl = in["tcp_rexmit_syn_limit_bl"].(int)
		ret.Dst_clist_overflow_policy_at_learning = in["dst_clist_overflow_policy_at_learning"].(int)
		ret.Zone_frag_rcvd = in["zone_frag_rcvd"].(int)
		ret.Zone_tcp_wellknown_sport_drop = in["zone_tcp_wellknown_sport_drop"].(int)
		ret.Src_tcp_wellknown_sport_drop = in["src_tcp_wellknown_sport_drop"].(int)
		ret.Secondary_dst_entry_pkt_rate_exceed = in["secondary_dst_entry_pkt_rate_exceed"].(int)
		ret.Secondary_dst_entry_kbit_rate_exceed = in["secondary_dst_entry_kbit_rate_exceed"].(int)
		ret.Secondary_dst_entry_conn_limit_exceed = in["secondary_dst_entry_conn_limit_exceed"].(int)
		ret.Secondary_dst_entry_conn_rate_exceed = in["secondary_dst_entry_conn_rate_exceed"].(int)
		ret.Secondary_dst_entry_frag_pkt_rate_exceed = in["secondary_dst_entry_frag_pkt_rate_exceed"].(int)
		ret.Src_udp_retry_gap_drop = in["src_udp_retry_gap_drop"].(int)
		ret.Dst_entry_kbit_rate_exceed_count = in["dst_entry_kbit_rate_exceed_count"].(int)
		ret.Secondary_entry_learn = in["secondary_entry_learn"].(int)
		ret.Secondary_entry_hit = in["secondary_entry_hit"].(int)
		ret.Secondary_entry_miss = in["secondary_entry_miss"].(int)
		ret.Secondary_entry_aged = in["secondary_entry_aged"].(int)
		ret.Secondary_entry_learning_thre_exceed = in["secondary_entry_learning_thre_exceed"].(int)
		ret.Zone_port_undef_hit = in["zone_port_undef_hit"].(int)
		ret.Zone_tcp_action_on_ack_timeout = in["zone_tcp_action_on_ack_timeout"].(int)
		ret.Zone_tcp_action_on_ack_reset = in["zone_tcp_action_on_ack_reset"].(int)
		ret.Zone_tcp_action_on_ack_blacklist = in["zone_tcp_action_on_ack_blacklist"].(int)
		ret.Src_tcp_action_on_ack_timeout = in["src_tcp_action_on_ack_timeout"].(int)
		ret.Src_tcp_action_on_ack_reset = in["src_tcp_action_on_ack_reset"].(int)
		ret.Src_tcp_action_on_ack_blacklist = in["src_tcp_action_on_ack_blacklist"].(int)
		ret.Zone_tcp_action_on_syn_timeout = in["zone_tcp_action_on_syn_timeout"].(int)
		ret.Zone_tcp_action_on_syn_reset = in["zone_tcp_action_on_syn_reset"].(int)
		ret.Zone_tcp_action_on_syn_blacklist = in["zone_tcp_action_on_syn_blacklist"].(int)
		ret.Src_tcp_action_on_syn_timeout = in["src_tcp_action_on_syn_timeout"].(int)
		ret.Src_tcp_action_on_syn_reset = in["src_tcp_action_on_syn_reset"].(int)
		ret.Src_tcp_action_on_syn_blacklist = in["src_tcp_action_on_syn_blacklist"].(int)
		ret.Zone_udp_frag_pkt_rate_exceed = in["zone_udp_frag_pkt_rate_exceed"].(int)
		ret.Zone_udp_frag_src_rate_drop = in["zone_udp_frag_src_rate_drop"].(int)
		ret.Zone_tcp_frag_pkt_rate_exceed = in["zone_tcp_frag_pkt_rate_exceed"].(int)
		ret.Zone_tcp_frag_src_rate_drop = in["zone_tcp_frag_src_rate_drop"].(int)
		ret.Zone_icmp_frag_pkt_rate_exceed = in["zone_icmp_frag_pkt_rate_exceed"].(int)
		ret.Zone_icmp_frag_src_rate_drop = in["zone_icmp_frag_src_rate_drop"].(int)
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
		ret.Source_entry_total = in["source_entry_total"].(int)
		ret.Source_entry_udp = in["source_entry_udp"].(int)
		ret.Source_entry_tcp = in["source_entry_tcp"].(int)
		ret.Source_entry_icmp = in["source_entry_icmp"].(int)
		ret.Source_entry_other = in["source_entry_other"].(int)
		ret.Dst_exceed_action_tunnel = in["dst_exceed_action_tunnel"].(int)
		ret.Dst_udp_retry_timeout_blacklist = in["dst_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_auth_timeout = in["src_udp_auth_timeout"].(int)
		ret.Zone_src_udp_retry_timeout_blacklist = in["zone_src_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_retry_pass = in["src_udp_retry_pass"].(int)
		ret.Secondary_port_learn = in["secondary_port_learn"].(int)
		ret.Secondary_port_aged = in["secondary_port_aged"].(int)
		ret.Dst_entry_outbound_udp_session_created = in["dst_entry_outbound_udp_session_created"].(int)
		ret.Dst_entry_outbound_udp_session_aged = in["dst_entry_outbound_udp_session_aged"].(int)
		ret.Dst_entry_outbound_tcp_session_created = in["dst_entry_outbound_tcp_session_created"].(int)
		ret.Dst_entry_outbound_tcp_session_aged = in["dst_entry_outbound_tcp_session_aged"].(int)
		ret.Dst_entry_outbound_pkt_rate_exceed = in["dst_entry_outbound_pkt_rate_exceed"].(int)
		ret.Dst_entry_outbound_kbit_rate_exceed = in["dst_entry_outbound_kbit_rate_exceed"].(int)
		ret.Dst_entry_outbound_kbit_rate_exceed_count = in["dst_entry_outbound_kbit_rate_exceed_count"].(int)
		ret.Dst_entry_outbound_conn_limit_exceed = in["dst_entry_outbound_conn_limit_exceed"].(int)
		ret.Dst_entry_outbound_conn_rate_exceed = in["dst_entry_outbound_conn_rate_exceed"].(int)
		ret.Dst_entry_outbound_frag_pkt_rate_exceed = in["dst_entry_outbound_frag_pkt_rate_exceed"].(int)
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
		ret.East_west_inbound_rcv_pkt = in["east_west_inbound_rcv_pkt"].(int)
		ret.East_west_inbound_drop_pkt = in["east_west_inbound_drop_pkt"].(int)
		ret.East_west_inbound_fwd_pkt = in["east_west_inbound_fwd_pkt"].(int)
		ret.East_west_inbound_rcv_byte = in["east_west_inbound_rcv_byte"].(int)
		ret.East_west_inbound_drop_byte = in["east_west_inbound_drop_byte"].(int)
		ret.East_west_inbound_fwd_byte = in["east_west_inbound_fwd_byte"].(int)
		ret.East_west_outbound_rcv_pkt = in["east_west_outbound_rcv_pkt"].(int)
		ret.East_west_outbound_drop_pkt = in["east_west_outbound_drop_pkt"].(int)
		ret.East_west_outbound_fwd_pkt = in["east_west_outbound_fwd_pkt"].(int)
		ret.East_west_outbound_rcv_byte = in["east_west_outbound_rcv_byte"].(int)
		ret.East_west_outbound_drop_byte = in["east_west_outbound_drop_byte"].(int)
		ret.East_west_outbound_fwd_byte = in["east_west_outbound_fwd_byte"].(int)
		ret.Dst_exceed_action_drop = in["dst_exceed_action_drop"].(int)
		ret.Prog_conn_samples = in["prog_conn_samples"].(int)
		ret.Prog_req_samples = in["prog_req_samples"].(int)
		ret.Prog_win_samples = in["prog_win_samples"].(int)
		ret.Victim_ip_learned = in["victim_ip_learned"].(int)
		ret.Victim_ip_aged = in["victim_ip_aged"].(int)
		ret.Prog_conn_samples_processed = in["prog_conn_samples_processed"].(int)
		ret.Prog_req_samples_processed = in["prog_req_samples_processed"].(int)
		ret.Prog_win_samples_processed = in["prog_win_samples_processed"].(int)
		ret.Dst_src_learn_overflow = in["dst_src_learn_overflow"].(int)
		ret.Dst_tcp_auth_rst = in["dst_tcp_auth_rst"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZoneStats(d *schema.ResourceData) edpt.DdosDstZoneStats {
	var ret edpt.DdosDstZoneStats

	ret.Stats = getObjectDdosDstZoneStatsStats(d.Get("stats").([]interface{}))

	ret.ZoneName = d.Get("zone_name").(string)
	return ret
}
