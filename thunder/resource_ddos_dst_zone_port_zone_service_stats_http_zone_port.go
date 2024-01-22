package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceStats34() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_stats_http_zone_port`: Statistics for the object zone-service\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceStats34Create,
		UpdateContext: resourceDdosDstZonePortZoneServiceStats34Update,
		ReadContext:   resourceDdosDstZonePortZoneServiceStats34Read,
		DeleteContext: resourceDdosDstZonePortZoneServiceStats34Delete,

		Schema: map[string]*schema.Schema{
			"port_num": {
				Type: schema.TypeInt, Required: true, Description: "Port Number",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'dns-tcp': DNS-TCP Port; 'dns-udp': DNS-UDP Port; 'http': HTTP Port; 'tcp': TCP Port; 'udp': UDP Port; 'ssl-l4': SSL-L4 Port; 'sip-udp': SIP-UDP Port; 'sip-tcp': SIP-TCP Port; 'quic': QUIC Port;",
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_zone_port": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ofo_timer_expired": {
										Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Timeout",
									},
									"partial_header": {
										Type: schema.TypeInt, Optional: true, Description: "Partial Header",
									},
									"slow_post": {
										Type: schema.TypeInt, Optional: true, Description: "Slow Post",
									},
									"idle_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Idle Timeout",
									},
									"pkts_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Packets",
									},
									"pkts_retrans_fin": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit FIN",
									},
									"pkts_retrans_rst": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit RST",
									},
									"pkts_retrans_push": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit PSH",
									},
									"pkts_retrans": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Retransmit",
									},
									"chunk_bad": {
										Type: schema.TypeInt, Optional: true, Description: "Bad Payload Chunk",
									},
									"chunk_sz_512": {
										Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 512",
									},
									"chunk_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 1K",
									},
									"chunk_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 2K",
									},
									"chunk_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Less Than or Equal to 4K",
									},
									"chunk_sz_gt_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Payload Chunk Size Larger Than 4K",
									},
									"negative_resp_remain": {
										Type: schema.TypeInt, Optional: true, Description: "Negative Response Remain",
									},
									"too_many_headers": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Header Too Many",
									},
									"header_name_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "Header Name Too Long",
									},
									"req_http11": {
										Type: schema.TypeInt, Optional: true, Description: "Request HTTP 1.1",
									},
									"req_http10": {
										Type: schema.TypeInt, Optional: true, Description: "Request HTTP 1.0",
									},
									"req_get": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method GET",
									},
									"req_head": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method HEAD",
									},
									"req_put": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method PUT",
									},
									"req_post": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method POST",
									},
									"req_trace": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method TRACE",
									},
									"req_options": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method OPTIONS",
									},
									"req_connect": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method CONNECT",
									},
									"req_delete": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method DELETE",
									},
									"req_unknown": {
										Type: schema.TypeInt, Optional: true, Description: "Request Method UNKNOWN",
									},
									"line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "Line Too Long",
									},
									"req_content_len": {
										Type: schema.TypeInt, Optional: true, Description: "Request Content Length Received",
									},
									"rsp_chunk": {
										Type: schema.TypeInt, Optional: true, Description: "Response Chunk",
									},
									"req": {
										Type: schema.TypeInt, Optional: true, Description: "Request Total",
									},
									"negative_req_remain": {
										Type: schema.TypeInt, Optional: true, Description: "Negative Request Remain",
									},
									"client_rst": {
										Type: schema.TypeInt, Optional: true, Description: "Client TCP RST Received",
									},
									"parsereq_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Parse Request Failed",
									},
									"req_retran": {
										Type: schema.TypeInt, Optional: true, Description: "Request Retransmit",
									},
									"req_ofo": {
										Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Request",
									},
									"invalid_header": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Header Invalid",
									},
									"policy_violation": {
										Type: schema.TypeInt, Optional: true, Description: "Policy Violation",
									},
									"less_than_mss_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Min Payload Size Fail Exceeded",
									},
									"dst_req_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate Exceeded",
									},
									"src_req_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Src Request Rate Exceeded",
									},
									"processed": {
										Type: schema.TypeInt, Optional: true, Description: "Packets Processed",
									},
									"new_syn": {
										Type: schema.TypeInt, Optional: true, Description: "TCP SYN",
									},
									"policy_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Policy Dropped",
									},
									"er_condition": {
										Type: schema.TypeInt, Optional: true, Description: "Error Condition",
									},
									"ofo_queue_sz_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Queue Exceeded",
									},
									"fail_alloc": {
										Type: schema.TypeInt, Optional: true, Description: "Alloc Failed",
									},
									"fail_get_line": {
										Type: schema.TypeInt, Optional: true, Description: "Get Line Failed",
									},
									"fail_alloc_hdr": {
										Type: schema.TypeInt, Optional: true, Description: "Alloc Header Failed",
									},
									"invalid_hdr_name": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Header Name Invalid",
									},
									"invalid_hdr_val": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Header Value Invalid",
									},
									"ud_challenge_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Challenge URL Redirect Failed",
									},
									"js_challenge_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Challenge Javascript Failed",
									},
									"challenge_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Challenge Failed",
									},
									"js_challenge_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Challenge Javascript Sent",
									},
									"ud_challenge_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Challenge URL Redirect Sent",
									},
									"parse_bad_chunk": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Bad Chunk",
									},
									"content_length_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Content Length Too Long",
									},
									"parse_too_many_headers": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Too Many Headers",
									},
									"parse_header_name_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Header Name Too Long",
									},
									"parse_line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Line Too Long",
									},
									"parse_req_line_too_long": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Request Line Too Long",
									},
									"window_small": {
										Type: schema.TypeInt, Optional: true, Description: "Window Size Small",
									},
									"window_small_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Window Size Small Dropped",
									},
									"use_hdr_src_ip": {
										Type: schema.TypeInt, Optional: true, Description: "Use IP In Header As Src",
									},
									"user_agent_filter_match": {
										Type: schema.TypeInt, Optional: true, Description: "Agent Filter Match",
									},
									"user_agent_filter_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Agent Filter Blacklisted",
									},
									"referer_filter_match": {
										Type: schema.TypeInt, Optional: true, Description: "Referer Filter Match",
									},
									"referer_filter_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Referer Filter Blacklisted",
									},
									"req_sz_1k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 1K",
									},
									"req_sz_2k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 2K",
									},
									"req_sz_4k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 4K",
									},
									"req_sz_8k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 8K",
									},
									"req_sz_16k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 16K",
									},
									"req_sz_32k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 32K",
									},
									"req_sz_64k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 64K",
									},
									"req_sz_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Less Than or Equal to 256K",
									},
									"req_sz_gt_256k": {
										Type: schema.TypeInt, Optional: true, Description: "Request Payload Size Larger Than 256K",
									},
									"filter_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Total Match",
									},
									"filter_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Not Matched on Pkt",
									},
									"filter_action_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Blacklist",
									},
									"filter_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Drop",
									},
									"filter_action_default_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Default Pass",
									},
									"dst_post_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Post Rate Exceeded",
									},
									"src_post_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Src Post Rate Exceeded",
									},
									"dst_resp_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Response Rate Exceeded",
									},
									"header_processing_time_0": {
										Type: schema.TypeInt, Optional: true, Description: "Header Process Time Less Than 1s",
									},
									"header_processing_time_1": {
										Type: schema.TypeInt, Optional: true, Description: "Header Process Time Less Than 10s",
									},
									"header_processing_time_2": {
										Type: schema.TypeInt, Optional: true, Description: "Header Process Time Less Than 30S",
									},
									"header_processing_time_3": {
										Type: schema.TypeInt, Optional: true, Description: "Header Process Time Larger or Equal to 30s",
									},
									"header_processing_incomplete": {
										Type: schema.TypeInt, Optional: true, Description: "Header Process Incomplete",
									},
									"resp_code_1xx": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Status Code 1XX",
									},
									"resp_code_2xx": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Status Code 2XX",
									},
									"resp_code_3xx": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Status Code 3XX",
									},
									"resp_code_4xx": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Status Code 4XX",
									},
									"resp_code_5xx": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Status Code 5XX",
									},
									"resp_code_other": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Status Code OTHER",
									},
									"filter_action_whitelist": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Whitelist",
									},
									"filter1_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter1 Match",
									},
									"filter2_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter2 Match",
									},
									"filter3_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter3 Match",
									},
									"filter4_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter4 Match",
									},
									"filter5_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter5 Match",
									},
									"filter_none_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter No Match",
									},
									"port_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Received",
									},
									"port_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Dropped",
									},
									"port_pkt_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Packets Forwarded",
									},
									"port_pkt_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Packet Rate Exceeded",
									},
									"port_kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded",
									},
									"port_conn_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn Rate Exceeded",
									},
									"port_conn_limm_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Conn Limit Exceeded",
									},
									"filter1_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter1 Match Dropped",
									},
									"filter2_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter2 Match Dropped",
									},
									"filter3_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter3 Match Dropped",
									},
									"filter4_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter4 Match Dropped",
									},
									"filter5_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter5 Match Dropped",
									},
									"filter6_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter6 Match",
									},
									"filter6_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter6 Match Dropped",
									},
									"filter7_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter7 Match",
									},
									"filter7_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter7 Match Dropped",
									},
									"filter8_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter8 Match",
									},
									"filter8_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter8 Match Dropped",
									},
									"filter9_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter9 Match",
									},
									"filter9_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter9 Match Dropped",
									},
									"filter10_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter10 Match",
									},
									"filter10_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter10 Match Dropped",
									},
									"port_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Received",
									},
									"outbound_port_bytes": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Received",
									},
									"outbound_port_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Received",
									},
									"outbound_port_pkt_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Forwarded",
									},
									"port_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Forwarded",
									},
									"port_bytes_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Bytes Dropped",
									},
									"port_src_bl": {
										Type: schema.TypeInt, Optional: true, Description: "Src Blacklisted",
									},
									"port_src_escalation": {
										Type: schema.TypeInt, Optional: true, Description: "Src Escalation",
									},
									"current_es_level": {
										Type: schema.TypeInt, Optional: true, Description: "Current Escalation Level",
									},
									"filter_auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Auth Failed",
									},
									"syn_auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Auth Failed",
									},
									"ack_auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Dropped",
									},
									"syn_cookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Cookie Failed",
									},
									"sess_create": {
										Type: schema.TypeInt, Optional: true, Description: "Session Create",
									},
									"exceed_drop_prate_src": {
										Type: schema.TypeInt, Optional: true, Description: "Src Pkt Rate Exceeded",
									},
									"exceed_drop_crate_src": {
										Type: schema.TypeInt, Optional: true, Description: "Src Conn Rate Exceeded",
									},
									"exceed_drop_climit_src": {
										Type: schema.TypeInt, Optional: true, Description: "Src Conn Limit Exceeded",
									},
									"exceed_drop_brate_src": {
										Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded",
									},
									"outbound_port_bytes_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Forwarded",
									},
									"outbound_port_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Packets Dropped",
									},
									"outbound_port_bytes_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Bytes Dropped",
									},
									"uri_filter_match": {
										Type: schema.TypeInt, Optional: true, Description: "URI Filter Match",
									},
									"exceed_drop_brate_src_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded Count",
									},
									"port_kbit_rate_exceed_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded Count",
									},
									"syn_cookie_sent": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Cookie Sent",
									},
									"ack_retry_init": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Init",
									},
									"ack_retry_gap_drop": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Retry-Gap Dropped",
									},
									"conn_prate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Conn Pkt Rate Exceeded",
									},
									"out_of_seq_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Out-Of-Seq Exceeded",
									},
									"retransmit_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Retransmit Exceeded",
									},
									"zero_window_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Zero-Window Exceeded",
									},
									"syn_retry_init": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry Init",
									},
									"syn_retry_gap_drop": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry-Gap Dropped",
									},
									"ack_retry_pass": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Passed",
									},
									"syn_retry_pass": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry Passed",
									},
									"tcp_filter_action_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Blacklist",
									},
									"tcp_filter_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Drop",
									},
									"src_policy_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Policy Dropped",
									},
									"src_dos_policy_violation": {
										Type: schema.TypeInt, Optional: true, Description: "Src Policy Violation",
									},
									"src_challenge_ud_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src Challenge URL Redirect Failed",
									},
									"src_challenge_js_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src Challenge Javascript Failed",
									},
									"src_js_challenge_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Src Challenge Javascript Sent",
									},
									"src_ud_challenge_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Src Challenge URL Redirect Sent",
									},
									"tcp_auth_drop": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Auth Dropped",
									},
									"tcp_auth_resp": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Auth Responded",
									},
									"http_auth_drop": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Auth Dropped",
									},
									"http_auth_resp": {
										Type: schema.TypeInt, Optional: true, Description: "HTTP Auth Responded",
									},
									"bl": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Blacklisted",
									},
									"src_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Packets Dropped",
									},
									"frag_rcvd": {
										Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Received",
									},
									"frag_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Dropped",
									},
									"sess_create_inbound": {
										Type: schema.TypeInt, Optional: true, Description: "Inbound Sessions Created",
									},
									"sess_create_outbound": {
										Type: schema.TypeInt, Optional: true, Description: "Outbound Sessions Created",
									},
									"conn_create_from_syn": {
										Type: schema.TypeInt, Optional: true, Description: "Connections Created From SYN",
									},
									"conn_create_from_ack": {
										Type: schema.TypeInt, Optional: true, Description: "Connections Created From ACK",
									},
									"conn_close": {
										Type: schema.TypeInt, Optional: true, Description: "Connections Closed",
									},
									"conn_close_w_rst": {
										Type: schema.TypeInt, Optional: true, Description: "RST Connections Closed",
									},
									"conn_close_w_fin": {
										Type: schema.TypeInt, Optional: true, Description: "FIN Connections Closed",
									},
									"conn_close_w_idle": {
										Type: schema.TypeInt, Optional: true, Description: "Idle Connections Closed",
									},
									"conn_close_half_open": {
										Type: schema.TypeInt, Optional: true, Description: "Half Open Connections Closed",
									},
									"sess_aged": {
										Type: schema.TypeInt, Optional: true, Description: "Sessions Aged Out",
									},
									"syn_drop": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Dropped",
									},
									"syn_auth_pass": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Auth Passed",
									},
									"unauth_drop": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Unauth Dropped",
									},
									"rst_cookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "RST Cookie Failed",
									},
									"syn_retry_failed": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry Dropped",
									},
									"tcp_filter1_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter1 Match",
									},
									"tcp_filter2_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter2 Match",
									},
									"tcp_filter3_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter3 Match",
									},
									"tcp_filter4_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter4 Match",
									},
									"tcp_filter5_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter5 Match",
									},
									"tcp_filter_none_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter No Match",
									},
									"tcp_filter_total_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Not Matched on Pkt",
									},
									"tcp_filter_action_default_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action Default Pass",
									},
									"tcp_filter_action_whitelist": {
										Type: schema.TypeInt, Optional: true, Description: "Filter Action WL",
									},
									"src_syn_auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Auth Failed",
									},
									"src_syn_cookie_sent": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Cookie Sent",
									},
									"src_syn_cookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Cookie Failed",
									},
									"src_unauth_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Unauth Dropped",
									},
									"src_rst_cookie_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src RST Cookie Failed",
									},
									"src_syn_retry_init": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry Init",
									},
									"src_syn_retry_gap_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry-Gap Dropped",
									},
									"src_syn_retry_failed": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry Dropped",
									},
									"src_ack_retry_init": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry Init",
									},
									"src_ack_retry_gap_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry Retry-Gap Dropped",
									},
									"src_ack_auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry Dropped",
									},
									"src_out_of_seq_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src Out-Of-Seq Exceeded",
									},
									"src_retransmit_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src Retransmit Exceeded",
									},
									"src_zero_window_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src Zero-Window Exceeded",
									},
									"src_conn_pkt_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Conn Pkt Rate Exceeded",
									},
									"src_filter_action_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
									},
									"src_filter_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
									},
									"src_filter_action_default_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
									},
									"src_filter_action_whitelist": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action WL",
									},
									"resp_mib_chunk_bad": {
										Type: schema.TypeInt, Optional: true, Description: "Bad HTTP Response Chunk",
									},
									"tcp_rexmit_syn_limit_drop": {
										Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retransmit Exceeded Drop",
									},
									"tcp_rexmit_syn_limit_bl": {
										Type: schema.TypeInt, Optional: true, Description: "TCP SYN Retransmit Exceeded Blacklist",
									},
									"conn_ofo_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Conn Out-Of-Seq Rate Exceeded",
									},
									"conn_rexmit_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Conn Retransmit Rate Exceeded",
									},
									"conn_zwindow_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Conn Zero-Window Rate Exceeded",
									},
									"src_conn_ofo_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Conn Out-Of-Seq Rate Exceeded",
									},
									"src_conn_rexmit_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Conn Retransmit Rate Exceeded",
									},
									"src_conn_zwindow_rate_excd": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Conn Zero-Window Rate Exceeded",
									},
									"ack_retry_rto_pass": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry RTO Passed",
									},
									"ack_retry_rto_fail": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry RTO Dropped",
									},
									"ack_retry_rto_progress": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry RTO Progress",
									},
									"syn_retry_rto_pass": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry RTO Passed",
									},
									"syn_retry_rto_fail": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry RTO Dropped",
									},
									"syn_retry_rto_progress": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry RTO Progress",
									},
									"src_syn_retry_rto_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry RTO Passed",
									},
									"src_syn_retry_rto_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry RTO Dropped",
									},
									"src_syn_retry_rto_progress": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry RTO Progress",
									},
									"src_ack_retry_rto_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry RTO Passed",
									},
									"src_ack_retry_rto_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry RTO Dropped",
									},
									"src_ack_retry_rto_progress": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry RTO Progress",
									},
									"wellknown_sport_drop": {
										Type: schema.TypeInt, Optional: true, Description: "TCP SrcPort Wellknown",
									},
									"src_well_known_port": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP SrcPort Wellknown",
									},
									"secondary_port_pkt_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port Packet Rate Exceeded",
									},
									"secondary_port_kbit_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port KiBit Rate Exceeded",
									},
									"secondary_port_kbit_rate_exceed_pkt": {
										Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port KiBit Rate Exceeded Count",
									},
									"secondary_port_conn_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port Conn Rate Exceeded",
									},
									"secondary_port_conn_limm_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port Conn Limit Exceeded",
									},
									"src_http_auth_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src HTTP Auth Dropped",
									},
									"src_auth_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth Dropped",
									},
									"src_frag_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
									},
									"no_policy_class_list_match": {
										Type: schema.TypeInt, Optional: true, Description: "No Policy Class-list Match",
									},
									"frag_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
									},
									"create_conn_non_syn_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "Create Conn with non-SYN Packets Dropped",
									},
									"src_create_conn_non_syn_dropped": {
										Type: schema.TypeInt, Optional: true, Description: "Src Create Conn with non-SYN Packets Dropped",
									},
									"pkts_connect_passthru": {
										Type: schema.TypeInt, Optional: true, Description: "Connect Passthru Packets",
									},
									"parse_req_line_too_small": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Request Line Too Small",
									},
									"parse_req_line_invalid_method": {
										Type: schema.TypeInt, Optional: true, Description: "Malform Request Line Invalid Method",
									},
									"port_syn_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Exceeded",
									},
									"src_filter1_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter1 Match",
									},
									"src_filter2_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter2 Match",
									},
									"src_filter3_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter3 Match",
									},
									"src_filter4_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter4 Match",
									},
									"src_filter5_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter5 Match",
									},
									"src_filter_none_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
									},
									"src_filter_total_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Not Matched on Pkt",
									},
									"src_filter_auth_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Auth Failed",
									},
									"syn_tfo_rcv": {
										Type: schema.TypeInt, Optional: true, Description: "SYN TFO Received",
									},
									"ack_retry_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Timeout",
									},
									"ack_retry_reset": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Timeout Reset",
									},
									"ack_retry_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "ACK Retry Timeout Blacklisted",
									},
									"src_ack_retry_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry Timeout",
									},
									"src_ack_retry_reset": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry Timeout Reset",
									},
									"src_ack_retry_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Src ACK Retry Timeout Blacklisted",
									},
									"syn_retry_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry Timeout",
									},
									"syn_retry_reset": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry Timeout Reset",
									},
									"syn_retry_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Retry Timeout Blacklisted",
									},
									"src_syn_retry_timeout": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry Timeout",
									},
									"src_syn_retry_reset": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry Timeout Reset",
									},
									"src_syn_retry_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Src SYN Retry Timeout Blacklisted",
									},
									"src_syn_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP SYN Rate Exceeded",
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
									"src_header_filter_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Total Match",
									},
									"src_header_filter_not_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Not Matched on Pkt",
									},
									"src_header_filter_action_blacklist": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Blacklist",
									},
									"src_header_filter_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Drop",
									},
									"src_header_filter_action_default_pass": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Default Pass",
									},
									"src_header_filter_action_whitelist": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Whitelist",
									},
									"src_header_filter_none_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter No Match",
									},
									"src_header_filter1_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter1 Match",
									},
									"src_header_filter2_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter2 Match",
									},
									"src_header_filter3_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter3 Match",
									},
									"src_header_filter4_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter4 Match",
									},
									"src_header_filter5_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter5 Match",
									},
									"src_header_filter6_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter6 Match",
									},
									"src_header_filter7_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter7 Match",
									},
									"src_header_filter8_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter8 Match",
									},
									"src_header_filter9_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter9 Match",
									},
									"src_header_filter10_match": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter10 Match",
									},
									"src_header_filter1_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter1 Match Dropped",
									},
									"src_header_filter2_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter2 Match Dropped",
									},
									"src_header_filter3_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter3 Match Dropped",
									},
									"src_header_filter4_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter4 Match Dropped",
									},
									"src_header_filter5_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter5 Match Dropped",
									},
									"src_header_filter6_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter6 Match Dropped",
									},
									"src_header_filter7_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter7 Match Dropped",
									},
									"src_header_filter8_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter8 Match Dropped",
									},
									"src_header_filter9_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter9 Match Dropped",
									},
									"src_header_filter10_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Filter10 Match Dropped",
									},
									"exceed_action_tunnel": {
										Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Tunnel",
									},
									"pattern_recognition_proceeded": {
										Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Engine Started",
									},
									"pattern_not_found": {
										Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Pattern Not Found",
									},
									"pattern_recognition_generic_error": {
										Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Exceptions",
									},
									"pattern_filter1_match": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted Filter1 Match",
									},
									"pattern_filter2_match": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted Filter2 Match",
									},
									"pattern_filter3_match": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted Filter3 Match",
									},
									"pattern_filter4_match": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted Filter4 Match",
									},
									"pattern_filter5_match": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted Filter5 Match",
									},
									"pattern_filter_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Extracted Filter Drop",
									},
									"pattern_recognition_sampling_started": {
										Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Sampling Started",
									},
									"pattern_recognition_pattern_changed": {
										Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Pattern Change Detected",
									},
									"dst_hw_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Packets Dropped",
									},
									"synack_reset_sent": {
										Type: schema.TypeInt, Optional: true, Description: "SYNACK Reset Sent",
									},
									"synack_multiple_attempts_per_ip_detected": {
										Type: schema.TypeInt, Optional: true, Description: "SYNACK Multiple Attempts Per IP Detected",
									},
									"secondary_port_hit": {
										Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port Hit",
									},
									"src_zone_service_entry_learned": {
										Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Learned",
									},
									"src_zone_service_entry_aged": {
										Type: schema.TypeInt, Optional: true, Description: "SrcZoneService Entry Aged",
									},
									"dst_hw_drop_inserted": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Drop Rules Inserted",
									},
									"dst_hw_drop_removed": {
										Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Drop Rules Removed",
									},
									"src_hw_drop_inserted": {
										Type: schema.TypeInt, Optional: true, Description: "Src Hardware Drop Rules Inserted",
									},
									"src_hw_drop_removed": {
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
									"snat_fail": {
										Type: schema.TypeInt, Optional: true, Description: "Source NAT Failure",
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
									"exceed_action_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Dropped",
									},
									"syn_auth_rst_ack_drop": {
										Type: schema.TypeInt, Optional: true, Description: "SYN Auth RST-ACK Dropped",
									},
									"prog_exceed_reset": {
										Type: schema.TypeInt, Optional: true, Description: "Req-Resp: Violation Exceed Reset",
									},
									"prog_conn_exceed_reset": {
										Type: schema.TypeInt, Optional: true, Description: "Connection: Violation Exceed Reset",
									},
									"prog_win_exceed_reset": {
										Type: schema.TypeInt, Optional: true, Description: "Time Window: Violation Exceed Reset",
									},
									"conn_create_from_synack": {
										Type: schema.TypeInt, Optional: true, Description: "Connections Created From SYNACK",
									},
									"port_synack_rate_exceed": {
										Type: schema.TypeInt, Optional: true, Description: "TCP SYNACK Rate Exceeded",
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
									"ew_inbound_port_rcv": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Inbound Packets Received",
									},
									"ew_inbound_port_drop": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Inbound Packets Dropped",
									},
									"ew_inbound_port_sent": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Inbound Packets Forwarded",
									},
									"ew_inbound_port_byte_rcv": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Inbound Bytes Recevied",
									},
									"ew_inbound_port_byte_drop": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Inbound Bytes Dropped",
									},
									"ew_inbound_port_byte_sent": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Inbound Bytes Forwarded",
									},
									"ew_outbound_port_rcv": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Outbound Packets Received",
									},
									"ew_outbound_port_drop": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Outbound Packets Dropped",
									},
									"ew_outbound_port_sent": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Outbound Packets Forwarded",
									},
									"ew_outbound_port_byte_rcv": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Outbound Bytes Recevied",
									},
									"ew_outbound_port_byte_sent": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Outbound Bytes Dropped",
									},
									"ew_outbound_port_byte_drop": {
										Type: schema.TypeInt, Optional: true, Description: "East-West Outbound Bytes Forwarded",
									},
									"no_route_drop": {
										Type: schema.TypeInt, Optional: true, Description: "No Route Dropped",
									},
									"unauth_src_session_reset": {
										Type: schema.TypeInt, Optional: true, Description: "Session Reset for Unauthenticated Src",
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
									"src_hw_drop": {
										Type: schema.TypeInt, Optional: true, Description: "Src Hardware Packets Dropped",
									},
									"tcp_auth_rst": {
										Type: schema.TypeInt, Optional: true, Description: "TCP Auth Reset",
									},
									"src_tcp_auth_rst": {
										Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth Reset",
									},
									"addr_filter_drop": {
										Type: schema.TypeInt, Optional: true, Description: "IP Filtering Policy: Dropped",
									},
									"addr_filter_bl": {
										Type: schema.TypeInt, Optional: true, Description: "IP Filtering Policy: Blacklisted",
									},
									"src_learn_overflow": {
										Type: schema.TypeInt, Optional: true, Description: "Source Dynamic Entry Overflow",
									},
								},
							},
						},
					},
				},
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceStats34Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceStats34Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceStats34(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceStats34Read(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceStats34Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceStats34Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceStats34(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceStats34Read(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceStats34Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceStats34Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceStats34(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceStats34Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceStats34Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceStats34(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosDstZonePortZoneServiceStats34Stats(d []interface{}) edpt.DdosDstZonePortZoneServiceStats34Stats {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceStats34Stats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpZonePort = getObjectDdosDstZonePortZoneServiceStats34StatsHttpZonePort(in["http_zone_port"].([]interface{}))
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceStats34StatsHttpZonePort(d []interface{}) edpt.DdosDstZonePortZoneServiceStats34StatsHttpZonePort {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceStats34StatsHttpZonePort
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ofo_timer_expired = in["ofo_timer_expired"].(int)
		ret.Partial_header = in["partial_header"].(int)
		ret.Slow_post = in["slow_post"].(int)
		ret.Idle_timeout = in["idle_timeout"].(int)
		ret.Pkts_ofo = in["pkts_ofo"].(int)
		ret.Pkts_retrans_fin = in["pkts_retrans_fin"].(int)
		ret.Pkts_retrans_rst = in["pkts_retrans_rst"].(int)
		ret.Pkts_retrans_push = in["pkts_retrans_push"].(int)
		ret.Pkts_retrans = in["pkts_retrans"].(int)
		ret.Chunk_bad = in["chunk_bad"].(int)
		ret.Chunk_sz_512 = in["chunk_sz_512"].(int)
		ret.Chunk_sz_1k = in["chunk_sz_1k"].(int)
		ret.Chunk_sz_2k = in["chunk_sz_2k"].(int)
		ret.Chunk_sz_4k = in["chunk_sz_4k"].(int)
		ret.Chunk_sz_gt_4k = in["chunk_sz_gt_4k"].(int)
		ret.Negative_resp_remain = in["negative_resp_remain"].(int)
		ret.Too_many_headers = in["too_many_headers"].(int)
		ret.Header_name_too_long = in["header_name_too_long"].(int)
		ret.Req_http11 = in["req_http11"].(int)
		ret.Req_http10 = in["req_http10"].(int)
		ret.Req_get = in["req_get"].(int)
		ret.Req_head = in["req_head"].(int)
		ret.Req_put = in["req_put"].(int)
		ret.Req_post = in["req_post"].(int)
		ret.Req_trace = in["req_trace"].(int)
		ret.Req_options = in["req_options"].(int)
		ret.Req_connect = in["req_connect"].(int)
		ret.Req_delete = in["req_delete"].(int)
		ret.Req_unknown = in["req_unknown"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Req_content_len = in["req_content_len"].(int)
		ret.Rsp_chunk = in["rsp_chunk"].(int)
		ret.Req = in["req"].(int)
		ret.Negative_req_remain = in["negative_req_remain"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Parsereq_fail = in["parsereq_fail"].(int)
		ret.Req_retran = in["req_retran"].(int)
		ret.Req_ofo = in["req_ofo"].(int)
		ret.Invalid_header = in["invalid_header"].(int)
		ret.Policy_violation = in["policy_violation"].(int)
		ret.Less_than_mss_exceed = in["less_than_mss_exceed"].(int)
		ret.Dst_req_rate_exceed = in["dst_req_rate_exceed"].(int)
		ret.Src_req_rate_exceed = in["src_req_rate_exceed"].(int)
		ret.Processed = in["processed"].(int)
		ret.New_syn = in["new_syn"].(int)
		ret.Policy_drop = in["policy_drop"].(int)
		ret.Er_condition = in["er_condition"].(int)
		ret.Ofo_queue_sz_exceed = in["ofo_queue_sz_exceed"].(int)
		ret.Fail_alloc = in["fail_alloc"].(int)
		ret.Fail_get_line = in["fail_get_line"].(int)
		ret.Fail_alloc_hdr = in["fail_alloc_hdr"].(int)
		ret.Invalid_hdr_name = in["invalid_hdr_name"].(int)
		ret.Invalid_hdr_val = in["invalid_hdr_val"].(int)
		ret.Ud_challenge_failed = in["ud_challenge_failed"].(int)
		ret.Js_challenge_failed = in["js_challenge_failed"].(int)
		ret.Challenge_failed = in["challenge_failed"].(int)
		ret.Js_challenge_sent = in["js_challenge_sent"].(int)
		ret.Ud_challenge_sent = in["ud_challenge_sent"].(int)
		ret.Parse_bad_chunk = in["parse_bad_chunk"].(int)
		ret.Content_length_too_long = in["content_length_too_long"].(int)
		ret.Parse_too_many_headers = in["parse_too_many_headers"].(int)
		ret.Parse_header_name_too_long = in["parse_header_name_too_long"].(int)
		ret.Parse_line_too_long = in["parse_line_too_long"].(int)
		ret.Parse_req_line_too_long = in["parse_req_line_too_long"].(int)
		ret.Window_small = in["window_small"].(int)
		ret.Window_small_drop = in["window_small_drop"].(int)
		ret.Use_hdr_src_ip = in["use_hdr_src_ip"].(int)
		ret.User_agent_filter_match = in["user_agent_filter_match"].(int)
		ret.User_agent_filter_blacklist = in["user_agent_filter_blacklist"].(int)
		ret.Referer_filter_match = in["referer_filter_match"].(int)
		ret.Referer_filter_blacklist = in["referer_filter_blacklist"].(int)
		ret.Req_sz_1k = in["req_sz_1k"].(int)
		ret.Req_sz_2k = in["req_sz_2k"].(int)
		ret.Req_sz_4k = in["req_sz_4k"].(int)
		ret.Req_sz_8k = in["req_sz_8k"].(int)
		ret.Req_sz_16k = in["req_sz_16k"].(int)
		ret.Req_sz_32k = in["req_sz_32k"].(int)
		ret.Req_sz_64k = in["req_sz_64k"].(int)
		ret.Req_sz_256k = in["req_sz_256k"].(int)
		ret.Req_sz_gt_256k = in["req_sz_gt_256k"].(int)
		ret.Filter_match = in["filter_match"].(int)
		ret.Filter_not_match = in["filter_not_match"].(int)
		ret.Filter_action_blacklist = in["filter_action_blacklist"].(int)
		ret.Filter_action_drop = in["filter_action_drop"].(int)
		ret.Filter_action_default_pass = in["filter_action_default_pass"].(int)
		ret.Dst_post_rate_exceed = in["dst_post_rate_exceed"].(int)
		ret.Src_post_rate_exceed = in["src_post_rate_exceed"].(int)
		ret.Dst_resp_rate_exceed = in["dst_resp_rate_exceed"].(int)
		ret.Header_processing_time_0 = in["header_processing_time_0"].(int)
		ret.Header_processing_time_1 = in["header_processing_time_1"].(int)
		ret.Header_processing_time_2 = in["header_processing_time_2"].(int)
		ret.Header_processing_time_3 = in["header_processing_time_3"].(int)
		ret.Header_processing_incomplete = in["header_processing_incomplete"].(int)
		ret.Resp_code_1xx = in["resp_code_1xx"].(int)
		ret.Resp_code_2xx = in["resp_code_2xx"].(int)
		ret.Resp_code_3xx = in["resp_code_3xx"].(int)
		ret.Resp_code_4xx = in["resp_code_4xx"].(int)
		ret.Resp_code_5xx = in["resp_code_5xx"].(int)
		ret.Resp_code_other = in["resp_code_other"].(int)
		ret.Filter_action_whitelist = in["filter_action_whitelist"].(int)
		ret.Filter1_match = in["filter1_match"].(int)
		ret.Filter2_match = in["filter2_match"].(int)
		ret.Filter3_match = in["filter3_match"].(int)
		ret.Filter4_match = in["filter4_match"].(int)
		ret.Filter5_match = in["filter5_match"].(int)
		ret.Filter_none_match = in["filter_none_match"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
		ret.Port_conn_rate_exceed = in["port_conn_rate_exceed"].(int)
		ret.Port_conn_limm_exceed = in["port_conn_limm_exceed"].(int)
		ret.Filter1_drop = in["filter1_drop"].(int)
		ret.Filter2_drop = in["filter2_drop"].(int)
		ret.Filter3_drop = in["filter3_drop"].(int)
		ret.Filter4_drop = in["filter4_drop"].(int)
		ret.Filter5_drop = in["filter5_drop"].(int)
		ret.Filter6_match = in["filter6_match"].(int)
		ret.Filter6_drop = in["filter6_drop"].(int)
		ret.Filter7_match = in["filter7_match"].(int)
		ret.Filter7_drop = in["filter7_drop"].(int)
		ret.Filter8_match = in["filter8_match"].(int)
		ret.Filter8_drop = in["filter8_drop"].(int)
		ret.Filter9_match = in["filter9_match"].(int)
		ret.Filter9_drop = in["filter9_drop"].(int)
		ret.Filter10_match = in["filter10_match"].(int)
		ret.Filter10_drop = in["filter10_drop"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Port_src_escalation = in["port_src_escalation"].(int)
		ret.Current_es_level = in["current_es_level"].(int)
		ret.Filter_auth_fail = in["filter_auth_fail"].(int)
		ret.Syn_auth_fail = in["syn_auth_fail"].(int)
		ret.Ack_auth_fail = in["ack_auth_fail"].(int)
		ret.Syn_cookie_fail = in["syn_cookie_fail"].(int)
		ret.Sess_create = in["sess_create"].(int)
		ret.Exceed_drop_prate_src = in["exceed_drop_prate_src"].(int)
		ret.Exceed_drop_crate_src = in["exceed_drop_crate_src"].(int)
		ret.Exceed_drop_climit_src = in["exceed_drop_climit_src"].(int)
		ret.Exceed_drop_brate_src = in["exceed_drop_brate_src"].(int)
		ret.Outbound_port_bytes_sent = in["outbound_port_bytes_sent"].(int)
		ret.Outbound_port_drop = in["outbound_port_drop"].(int)
		ret.Outbound_port_bytes_drop = in["outbound_port_bytes_drop"].(int)
		ret.Uri_filter_match = in["uri_filter_match"].(int)
		ret.Exceed_drop_brate_src_pkt = in["exceed_drop_brate_src_pkt"].(int)
		ret.Port_kbit_rate_exceed_pkt = in["port_kbit_rate_exceed_pkt"].(int)
		ret.Syn_cookie_sent = in["syn_cookie_sent"].(int)
		ret.Ack_retry_init = in["ack_retry_init"].(int)
		ret.Ack_retry_gap_drop = in["ack_retry_gap_drop"].(int)
		ret.Conn_prate_excd = in["conn_prate_excd"].(int)
		ret.Out_of_seq_excd = in["out_of_seq_excd"].(int)
		ret.Retransmit_excd = in["retransmit_excd"].(int)
		ret.Zero_window_excd = in["zero_window_excd"].(int)
		ret.Syn_retry_init = in["syn_retry_init"].(int)
		ret.Syn_retry_gap_drop = in["syn_retry_gap_drop"].(int)
		ret.Ack_retry_pass = in["ack_retry_pass"].(int)
		ret.Syn_retry_pass = in["syn_retry_pass"].(int)
		ret.Tcp_filter_action_blacklist = in["tcp_filter_action_blacklist"].(int)
		ret.Tcp_filter_action_drop = in["tcp_filter_action_drop"].(int)
		ret.Src_policy_drop = in["src_policy_drop"].(int)
		ret.Src_dos_policy_violation = in["src_dos_policy_violation"].(int)
		ret.Src_challenge_ud_fail = in["src_challenge_ud_fail"].(int)
		ret.Src_challenge_js_fail = in["src_challenge_js_fail"].(int)
		ret.Src_js_challenge_sent = in["src_js_challenge_sent"].(int)
		ret.Src_ud_challenge_sent = in["src_ud_challenge_sent"].(int)
		ret.Tcp_auth_drop = in["tcp_auth_drop"].(int)
		ret.Tcp_auth_resp = in["tcp_auth_resp"].(int)
		ret.Http_auth_drop = in["http_auth_drop"].(int)
		ret.Http_auth_resp = in["http_auth_resp"].(int)
		ret.Bl = in["bl"].(int)
		ret.Src_drop = in["src_drop"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.Frag_drop = in["frag_drop"].(int)
		ret.Sess_create_inbound = in["sess_create_inbound"].(int)
		ret.Sess_create_outbound = in["sess_create_outbound"].(int)
		ret.Conn_create_from_syn = in["conn_create_from_syn"].(int)
		ret.Conn_create_from_ack = in["conn_create_from_ack"].(int)
		ret.Conn_close = in["conn_close"].(int)
		ret.Conn_close_w_rst = in["conn_close_w_rst"].(int)
		ret.Conn_close_w_fin = in["conn_close_w_fin"].(int)
		ret.Conn_close_w_idle = in["conn_close_w_idle"].(int)
		ret.Conn_close_half_open = in["conn_close_half_open"].(int)
		ret.Sess_aged = in["sess_aged"].(int)
		ret.Syn_drop = in["syn_drop"].(int)
		ret.Syn_auth_pass = in["syn_auth_pass"].(int)
		ret.Unauth_drop = in["unauth_drop"].(int)
		ret.Rst_cookie_fail = in["rst_cookie_fail"].(int)
		ret.Syn_retry_failed = in["syn_retry_failed"].(int)
		ret.Tcp_filter1_match = in["tcp_filter1_match"].(int)
		ret.Tcp_filter2_match = in["tcp_filter2_match"].(int)
		ret.Tcp_filter3_match = in["tcp_filter3_match"].(int)
		ret.Tcp_filter4_match = in["tcp_filter4_match"].(int)
		ret.Tcp_filter5_match = in["tcp_filter5_match"].(int)
		ret.Tcp_filter_none_match = in["tcp_filter_none_match"].(int)
		ret.Tcp_filter_total_not_match = in["tcp_filter_total_not_match"].(int)
		ret.Tcp_filter_action_default_pass = in["tcp_filter_action_default_pass"].(int)
		ret.Tcp_filter_action_whitelist = in["tcp_filter_action_whitelist"].(int)
		ret.Src_syn_auth_fail = in["src_syn_auth_fail"].(int)
		ret.Src_syn_cookie_sent = in["src_syn_cookie_sent"].(int)
		ret.Src_syn_cookie_fail = in["src_syn_cookie_fail"].(int)
		ret.Src_unauth_drop = in["src_unauth_drop"].(int)
		ret.Src_rst_cookie_fail = in["src_rst_cookie_fail"].(int)
		ret.Src_syn_retry_init = in["src_syn_retry_init"].(int)
		ret.Src_syn_retry_gap_drop = in["src_syn_retry_gap_drop"].(int)
		ret.Src_syn_retry_failed = in["src_syn_retry_failed"].(int)
		ret.Src_ack_retry_init = in["src_ack_retry_init"].(int)
		ret.Src_ack_retry_gap_drop = in["src_ack_retry_gap_drop"].(int)
		ret.Src_ack_auth_fail = in["src_ack_auth_fail"].(int)
		ret.Src_out_of_seq_excd = in["src_out_of_seq_excd"].(int)
		ret.Src_retransmit_excd = in["src_retransmit_excd"].(int)
		ret.Src_zero_window_excd = in["src_zero_window_excd"].(int)
		ret.Src_conn_pkt_rate_excd = in["src_conn_pkt_rate_excd"].(int)
		ret.Src_filter_action_blacklist = in["src_filter_action_blacklist"].(int)
		ret.Src_filter_action_drop = in["src_filter_action_drop"].(int)
		ret.Src_filter_action_default_pass = in["src_filter_action_default_pass"].(int)
		ret.Src_filter_action_whitelist = in["src_filter_action_whitelist"].(int)
		ret.Resp_mib_chunk_bad = in["resp_mib_chunk_bad"].(int)
		ret.Tcp_rexmit_syn_limit_drop = in["tcp_rexmit_syn_limit_drop"].(int)
		ret.Tcp_rexmit_syn_limit_bl = in["tcp_rexmit_syn_limit_bl"].(int)
		ret.Conn_ofo_rate_excd = in["conn_ofo_rate_excd"].(int)
		ret.Conn_rexmit_rate_excd = in["conn_rexmit_rate_excd"].(int)
		ret.Conn_zwindow_rate_excd = in["conn_zwindow_rate_excd"].(int)
		ret.Src_conn_ofo_rate_excd = in["src_conn_ofo_rate_excd"].(int)
		ret.Src_conn_rexmit_rate_excd = in["src_conn_rexmit_rate_excd"].(int)
		ret.Src_conn_zwindow_rate_excd = in["src_conn_zwindow_rate_excd"].(int)
		ret.Ack_retry_rto_pass = in["ack_retry_rto_pass"].(int)
		ret.Ack_retry_rto_fail = in["ack_retry_rto_fail"].(int)
		ret.Ack_retry_rto_progress = in["ack_retry_rto_progress"].(int)
		ret.Syn_retry_rto_pass = in["syn_retry_rto_pass"].(int)
		ret.Syn_retry_rto_fail = in["syn_retry_rto_fail"].(int)
		ret.Syn_retry_rto_progress = in["syn_retry_rto_progress"].(int)
		ret.Src_syn_retry_rto_pass = in["src_syn_retry_rto_pass"].(int)
		ret.Src_syn_retry_rto_fail = in["src_syn_retry_rto_fail"].(int)
		ret.Src_syn_retry_rto_progress = in["src_syn_retry_rto_progress"].(int)
		ret.Src_ack_retry_rto_pass = in["src_ack_retry_rto_pass"].(int)
		ret.Src_ack_retry_rto_fail = in["src_ack_retry_rto_fail"].(int)
		ret.Src_ack_retry_rto_progress = in["src_ack_retry_rto_progress"].(int)
		ret.Wellknown_sport_drop = in["wellknown_sport_drop"].(int)
		ret.Src_well_known_port = in["src_well_known_port"].(int)
		ret.Secondary_port_pkt_rate_exceed = in["secondary_port_pkt_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed = in["secondary_port_kbit_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed_pkt = in["secondary_port_kbit_rate_exceed_pkt"].(int)
		ret.Secondary_port_conn_rate_exceed = in["secondary_port_conn_rate_exceed"].(int)
		ret.Secondary_port_conn_limm_exceed = in["secondary_port_conn_limm_exceed"].(int)
		ret.Src_http_auth_drop = in["src_http_auth_drop"].(int)
		ret.Src_auth_drop = in["src_auth_drop"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.No_policy_class_list_match = in["no_policy_class_list_match"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Create_conn_non_syn_dropped = in["create_conn_non_syn_dropped"].(int)
		ret.Src_create_conn_non_syn_dropped = in["src_create_conn_non_syn_dropped"].(int)
		ret.Pkts_connect_passthru = in["pkts_connect_passthru"].(int)
		ret.Parse_req_line_too_small = in["parse_req_line_too_small"].(int)
		ret.Parse_req_line_invalid_method = in["parse_req_line_invalid_method"].(int)
		ret.Port_syn_rate_exceed = in["port_syn_rate_exceed"].(int)
		ret.Src_filter1_match = in["src_filter1_match"].(int)
		ret.Src_filter2_match = in["src_filter2_match"].(int)
		ret.Src_filter3_match = in["src_filter3_match"].(int)
		ret.Src_filter4_match = in["src_filter4_match"].(int)
		ret.Src_filter5_match = in["src_filter5_match"].(int)
		ret.Src_filter_none_match = in["src_filter_none_match"].(int)
		ret.Src_filter_total_not_match = in["src_filter_total_not_match"].(int)
		ret.Src_filter_auth_fail = in["src_filter_auth_fail"].(int)
		ret.Syn_tfo_rcv = in["syn_tfo_rcv"].(int)
		ret.Ack_retry_timeout = in["ack_retry_timeout"].(int)
		ret.Ack_retry_reset = in["ack_retry_reset"].(int)
		ret.Ack_retry_blacklist = in["ack_retry_blacklist"].(int)
		ret.Src_ack_retry_timeout = in["src_ack_retry_timeout"].(int)
		ret.Src_ack_retry_reset = in["src_ack_retry_reset"].(int)
		ret.Src_ack_retry_blacklist = in["src_ack_retry_blacklist"].(int)
		ret.Syn_retry_timeout = in["syn_retry_timeout"].(int)
		ret.Syn_retry_reset = in["syn_retry_reset"].(int)
		ret.Syn_retry_blacklist = in["syn_retry_blacklist"].(int)
		ret.Src_syn_retry_timeout = in["src_syn_retry_timeout"].(int)
		ret.Src_syn_retry_reset = in["src_syn_retry_reset"].(int)
		ret.Src_syn_retry_blacklist = in["src_syn_retry_blacklist"].(int)
		ret.Src_syn_rate_exceed = in["src_syn_rate_exceed"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Src_header_filter_match = in["src_header_filter_match"].(int)
		ret.Src_header_filter_not_match = in["src_header_filter_not_match"].(int)
		ret.Src_header_filter_action_blacklist = in["src_header_filter_action_blacklist"].(int)
		ret.Src_header_filter_action_drop = in["src_header_filter_action_drop"].(int)
		ret.Src_header_filter_action_default_pass = in["src_header_filter_action_default_pass"].(int)
		ret.Src_header_filter_action_whitelist = in["src_header_filter_action_whitelist"].(int)
		ret.Src_header_filter_none_match = in["src_header_filter_none_match"].(int)
		ret.Src_header_filter1_match = in["src_header_filter1_match"].(int)
		ret.Src_header_filter2_match = in["src_header_filter2_match"].(int)
		ret.Src_header_filter3_match = in["src_header_filter3_match"].(int)
		ret.Src_header_filter4_match = in["src_header_filter4_match"].(int)
		ret.Src_header_filter5_match = in["src_header_filter5_match"].(int)
		ret.Src_header_filter6_match = in["src_header_filter6_match"].(int)
		ret.Src_header_filter7_match = in["src_header_filter7_match"].(int)
		ret.Src_header_filter8_match = in["src_header_filter8_match"].(int)
		ret.Src_header_filter9_match = in["src_header_filter9_match"].(int)
		ret.Src_header_filter10_match = in["src_header_filter10_match"].(int)
		ret.Src_header_filter1_drop = in["src_header_filter1_drop"].(int)
		ret.Src_header_filter2_drop = in["src_header_filter2_drop"].(int)
		ret.Src_header_filter3_drop = in["src_header_filter3_drop"].(int)
		ret.Src_header_filter4_drop = in["src_header_filter4_drop"].(int)
		ret.Src_header_filter5_drop = in["src_header_filter5_drop"].(int)
		ret.Src_header_filter6_drop = in["src_header_filter6_drop"].(int)
		ret.Src_header_filter7_drop = in["src_header_filter7_drop"].(int)
		ret.Src_header_filter8_drop = in["src_header_filter8_drop"].(int)
		ret.Src_header_filter9_drop = in["src_header_filter9_drop"].(int)
		ret.Src_header_filter10_drop = in["src_header_filter10_drop"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Pattern_recognition_proceeded = in["pattern_recognition_proceeded"].(int)
		ret.Pattern_not_found = in["pattern_not_found"].(int)
		ret.Pattern_recognition_generic_error = in["pattern_recognition_generic_error"].(int)
		ret.Pattern_filter1_match = in["pattern_filter1_match"].(int)
		ret.Pattern_filter2_match = in["pattern_filter2_match"].(int)
		ret.Pattern_filter3_match = in["pattern_filter3_match"].(int)
		ret.Pattern_filter4_match = in["pattern_filter4_match"].(int)
		ret.Pattern_filter5_match = in["pattern_filter5_match"].(int)
		ret.Pattern_filter_drop = in["pattern_filter_drop"].(int)
		ret.Pattern_recognition_sampling_started = in["pattern_recognition_sampling_started"].(int)
		ret.Pattern_recognition_pattern_changed = in["pattern_recognition_pattern_changed"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Synack_reset_sent = in["synack_reset_sent"].(int)
		ret.Synack_multiple_attempts_per_ip_detected = in["synack_multiple_attempts_per_ip_detected"].(int)
		ret.Secondary_port_hit = in["secondary_port_hit"].(int)
		ret.Src_zone_service_entry_learned = in["src_zone_service_entry_learned"].(int)
		ret.Src_zone_service_entry_aged = in["src_zone_service_entry_aged"].(int)
		ret.Dst_hw_drop_inserted = in["dst_hw_drop_inserted"].(int)
		ret.Dst_hw_drop_removed = in["dst_hw_drop_removed"].(int)
		ret.Src_hw_drop_inserted = in["src_hw_drop_inserted"].(int)
		ret.Src_hw_drop_removed = in["src_hw_drop_removed"].(int)
		ret.Prog_first_req_time_exceed = in["prog_first_req_time_exceed"].(int)
		ret.Prog_req_resp_time_exceed = in["prog_req_resp_time_exceed"].(int)
		ret.Prog_request_len_exceed = in["prog_request_len_exceed"].(int)
		ret.Prog_response_len_exceed = in["prog_response_len_exceed"].(int)
		ret.Prog_resp_req_ratio_exceed = in["prog_resp_req_ratio_exceed"].(int)
		ret.Prog_resp_req_time_exceed = in["prog_resp_req_time_exceed"].(int)
		ret.Prog_conn_sent_exceed = in["prog_conn_sent_exceed"].(int)
		ret.Prog_conn_rcvd_exceed = in["prog_conn_rcvd_exceed"].(int)
		ret.Prog_conn_time_exceed = in["prog_conn_time_exceed"].(int)
		ret.Prog_conn_rcvd_sent_ratio_exceed = in["prog_conn_rcvd_sent_ratio_exceed"].(int)
		ret.Prog_win_sent_exceed = in["prog_win_sent_exceed"].(int)
		ret.Prog_win_rcvd_exceed = in["prog_win_rcvd_exceed"].(int)
		ret.Prog_win_rcvd_sent_ratio_exceed = in["prog_win_rcvd_sent_ratio_exceed"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Prog_exceed_drop = in["prog_exceed_drop"].(int)
		ret.Prog_exceed_bl = in["prog_exceed_bl"].(int)
		ret.Prog_conn_exceed_drop = in["prog_conn_exceed_drop"].(int)
		ret.Prog_conn_exceed_bl = in["prog_conn_exceed_bl"].(int)
		ret.Prog_win_exceed_drop = in["prog_win_exceed_drop"].(int)
		ret.Prog_win_exceed_bl = in["prog_win_exceed_bl"].(int)
		ret.Exceed_action_drop = in["exceed_action_drop"].(int)
		ret.Syn_auth_rst_ack_drop = in["syn_auth_rst_ack_drop"].(int)
		ret.Prog_exceed_reset = in["prog_exceed_reset"].(int)
		ret.Prog_conn_exceed_reset = in["prog_conn_exceed_reset"].(int)
		ret.Prog_win_exceed_reset = in["prog_win_exceed_reset"].(int)
		ret.Conn_create_from_synack = in["conn_create_from_synack"].(int)
		ret.Port_synack_rate_exceed = in["port_synack_rate_exceed"].(int)
		ret.Prog_conn_samples = in["prog_conn_samples"].(int)
		ret.Prog_req_samples = in["prog_req_samples"].(int)
		ret.Prog_win_samples = in["prog_win_samples"].(int)
		ret.Ew_inbound_port_rcv = in["ew_inbound_port_rcv"].(int)
		ret.Ew_inbound_port_drop = in["ew_inbound_port_drop"].(int)
		ret.Ew_inbound_port_sent = in["ew_inbound_port_sent"].(int)
		ret.Ew_inbound_port_byte_rcv = in["ew_inbound_port_byte_rcv"].(int)
		ret.Ew_inbound_port_byte_drop = in["ew_inbound_port_byte_drop"].(int)
		ret.Ew_inbound_port_byte_sent = in["ew_inbound_port_byte_sent"].(int)
		ret.Ew_outbound_port_rcv = in["ew_outbound_port_rcv"].(int)
		ret.Ew_outbound_port_drop = in["ew_outbound_port_drop"].(int)
		ret.Ew_outbound_port_sent = in["ew_outbound_port_sent"].(int)
		ret.Ew_outbound_port_byte_rcv = in["ew_outbound_port_byte_rcv"].(int)
		ret.Ew_outbound_port_byte_sent = in["ew_outbound_port_byte_sent"].(int)
		ret.Ew_outbound_port_byte_drop = in["ew_outbound_port_byte_drop"].(int)
		ret.No_route_drop = in["no_route_drop"].(int)
		ret.Unauth_src_session_reset = in["unauth_src_session_reset"].(int)
		ret.Prog_conn_samples_processed = in["prog_conn_samples_processed"].(int)
		ret.Prog_req_samples_processed = in["prog_req_samples_processed"].(int)
		ret.Prog_win_samples_processed = in["prog_win_samples_processed"].(int)
		ret.Src_hw_drop = in["src_hw_drop"].(int)
		ret.Tcp_auth_rst = in["tcp_auth_rst"].(int)
		ret.Src_tcp_auth_rst = in["src_tcp_auth_rst"].(int)
		ret.Addr_filter_drop = in["addr_filter_drop"].(int)
		ret.Addr_filter_bl = in["addr_filter_bl"].(int)
		ret.Src_learn_overflow = in["src_learn_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceStats34(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceStats34 {
	var ret edpt.DdosDstZonePortZoneServiceStats34
	ret.Inst.PortNum = d.Get("port_num").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.Stats = getObjectDdosDstZonePortZoneServiceStats34Stats(d.Get("stats").([]interface{}))
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
