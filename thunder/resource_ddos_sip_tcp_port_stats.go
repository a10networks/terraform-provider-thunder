package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSipTcpPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_sip_tcp_port_stats`: Statistics for the object sip-tcp-port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSipTcpPortStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Drop",
						},
						"policy_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Violation",
						},
						"idle_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Idle Timeout",
						},
						"ofo_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Out-of-Order Timeout",
						},
						"seq_check_ofo": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Out-Of-Order",
						},
						"pkts_ofo_total": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Out-Of-Order Total",
						},
						"ofo_queue_size_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Out-Of-Order Queue Size Exceed",
						},
						"seq_check_retrans_fin": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Fin",
						},
						"seq_check_retrans_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Rst",
						},
						"seq_check_retrans_push": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Push",
						},
						"seq_check_retrans_other": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence Check Retransmit Other",
						},
						"pkts_retrans_total": {
							Type: schema.TypeInt, Optional: true, Description: "Packets Retransmit Total",
						},
						"client_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Client Rst",
						},
						"error_condition": {
							Type: schema.TypeInt, Optional: true, Description: "Error Condition",
						},
						"request_method_ack": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method ACK",
						},
						"request_method_bye": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method BYE",
						},
						"request_method_cancel": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method CANCEL",
						},
						"request_method_invite": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method INVITE",
						},
						"request_method_info": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method INFO",
						},
						"request_method_message": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method MESSAGE",
						},
						"request_method_notify": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method NOTIFY",
						},
						"request_method_options": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method OPTIONS",
						},
						"request_method_prack": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method PRACK",
						},
						"request_method_publish": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method PUBLISH",
						},
						"request_method_register": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method REGISTER",
						},
						"request_method_refer": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method REFER",
						},
						"request_method_subscribe": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method SUBSCRIBE",
						},
						"request_method_update": {
							Type: schema.TypeInt, Optional: true, Description: "Request Method UPDATE",
						},
						"request_method_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Request Method",
						},
						"request_unknown_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Request Version",
						},
						"keep_alive_msg": {
							Type: schema.TypeInt, Optional: true, Description: "KeepAlive Message",
						},
						"rate1_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 1 Limit Exceed",
						},
						"rate2_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 2 Limit Exceed",
						},
						"rate3_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 3 Limit Exceed",
						},
						"src_rate1_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 1 Limit Exceed",
						},
						"src_rate2_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 2 Limit Exceed",
						},
						"src_rate3_limit_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 3 Limit Exceed",
						},
						"response_1xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 1xx",
						},
						"response_2xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 2xx",
						},
						"response_3xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 3xx",
						},
						"response_4xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 4xx",
						},
						"response_5xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 5xx",
						},
						"response_6xx": {
							Type: schema.TypeInt, Optional: true, Description: "Response Status Code 6xx",
						},
						"response_unknown": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Response Status Code",
						},
						"response_unknown_version": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Response Version",
						},
						"read_start_line_error": {
							Type: schema.TypeInt, Optional: true, Description: "Start Line Read Error",
						},
						"invalid_start_line_error": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Start Line",
						},
						"parse_start_line_error": {
							Type: schema.TypeInt, Optional: true, Description: "Start Line Parse Error",
						},
						"line_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Line Too Long",
						},
						"line_mem_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Line Memory Allocated",
						},
						"line_mem_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Line Memory Freed",
						},
						"max_uri_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max URI Length Exceed",
						},
						"too_many_header": {
							Type: schema.TypeInt, Optional: true, Description: "Max Header Count Exceed",
						},
						"invalid_header": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Header",
						},
						"header_name_too_long": {
							Type: schema.TypeInt, Optional: true, Description: "Max Header Name Length Exceed",
						},
						"parse_header_fail_error": {
							Type: schema.TypeInt, Optional: true, Description: "Header Parse Fail",
						},
						"max_header_value_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max Header Value Length Exceed",
						},
						"max_call_id_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max Call ID Length Exceed",
						},
						"header_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Match",
						},
						"header_filter_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Not Match",
						},
						"header_filter_none_match": {
							Type: schema.TypeInt, Optional: true, Description: "None Header Filter Match",
						},
						"header_filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Drop",
						},
						"header_filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Blacklist",
						},
						"header_filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Whitelist",
						},
						"header_filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter Action Default Pass",
						},
						"header_filter_filter1_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter 1 Match",
						},
						"header_filter_filter2_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter 2 Match",
						},
						"header_filter_filter3_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter 3 Match",
						},
						"header_filter_filter4_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter 4 Match",
						},
						"header_filter_filter5_match": {
							Type: schema.TypeInt, Optional: true, Description: "Header Filter 5 Match",
						},
						"max_sdp_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Max SDP Length Exceed",
						},
						"body_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Body Too Big",
						},
						"get_content_fail_error": {
							Type: schema.TypeInt, Optional: true, Description: "Get Content Fail",
						},
						"concatenate_msg": {
							Type: schema.TypeInt, Optional: true, Description: "Concatenate Message",
						},
						"mem_alloc_fail_error": {
							Type: schema.TypeInt, Optional: true, Description: "Memory Allocate Fail",
						},
						"malform_request": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Request",
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
						"filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Blacklist",
						},
						"filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Drop",
						},
						"filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Default Pass",
						},
						"filter_action_whitelist": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action WL",
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
						"tcp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Dropped",
						},
						"tcp_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Responded",
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
						"src_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth Dropped",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
						},
						"frag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
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
						"filter_total_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Not Matched on Pkt",
						},
						"bl": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Blacklisted",
						},
						"src_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Packets Dropped",
						},
						"create_conn_non_syn_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Create Conn with non-SYN Packets Dropped",
						},
						"src_create_conn_non_syn_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Src Create Conn with non-SYN Packets Dropped",
						},
						"port_syn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Exceeded",
						},
						"src_syn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP SYN Rate Exceeded",
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
						"prog_conn_samples_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed: Connnection",
						},
						"prog_req_samples_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed: Req-Resp",
						},
						"prog_win_samples_processed": {
							Type: schema.TypeInt, Optional: true, Description: "Sample Processed: Time Window",
						},
						"tcp_auth_rst": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Reset",
						},
						"src_tcp_auth_rst": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP Auth Reset",
						},
					},
				},
			},
		},
	}
}

func resourceDdosSipTcpPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSipTcpPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSipTcpPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSipTcpPortStatsStats := setObjectDdosSipTcpPortStatsStats(res)
		d.Set("stats", DdosSipTcpPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSipTcpPortStatsStats(ret edpt.DataDdosSipTcpPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"policy_drop":                              ret.DtDdosSipTcpPortStats.Stats.Policy_drop,
			"policy_violation":                         ret.DtDdosSipTcpPortStats.Stats.Policy_violation,
			"idle_timeout":                             ret.DtDdosSipTcpPortStats.Stats.Idle_timeout,
			"ofo_timeout":                              ret.DtDdosSipTcpPortStats.Stats.Ofo_timeout,
			"seq_check_ofo":                            ret.DtDdosSipTcpPortStats.Stats.Seq_check_ofo,
			"pkts_ofo_total":                           ret.DtDdosSipTcpPortStats.Stats.Pkts_ofo_total,
			"ofo_queue_size_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Ofo_queue_size_exceed,
			"seq_check_retrans_fin":                    ret.DtDdosSipTcpPortStats.Stats.Seq_check_retrans_fin,
			"seq_check_retrans_rst":                    ret.DtDdosSipTcpPortStats.Stats.Seq_check_retrans_rst,
			"seq_check_retrans_push":                   ret.DtDdosSipTcpPortStats.Stats.Seq_check_retrans_push,
			"seq_check_retrans_other":                  ret.DtDdosSipTcpPortStats.Stats.Seq_check_retrans_other,
			"pkts_retrans_total":                       ret.DtDdosSipTcpPortStats.Stats.Pkts_retrans_total,
			"client_rst":                               ret.DtDdosSipTcpPortStats.Stats.Client_rst,
			"error_condition":                          ret.DtDdosSipTcpPortStats.Stats.Error_condition,
			"request_method_ack":                       ret.DtDdosSipTcpPortStats.Stats.Request_method_ack,
			"request_method_bye":                       ret.DtDdosSipTcpPortStats.Stats.Request_method_bye,
			"request_method_cancel":                    ret.DtDdosSipTcpPortStats.Stats.Request_method_cancel,
			"request_method_invite":                    ret.DtDdosSipTcpPortStats.Stats.Request_method_invite,
			"request_method_info":                      ret.DtDdosSipTcpPortStats.Stats.Request_method_info,
			"request_method_message":                   ret.DtDdosSipTcpPortStats.Stats.Request_method_message,
			"request_method_notify":                    ret.DtDdosSipTcpPortStats.Stats.Request_method_notify,
			"request_method_options":                   ret.DtDdosSipTcpPortStats.Stats.Request_method_options,
			"request_method_prack":                     ret.DtDdosSipTcpPortStats.Stats.Request_method_prack,
			"request_method_publish":                   ret.DtDdosSipTcpPortStats.Stats.Request_method_publish,
			"request_method_register":                  ret.DtDdosSipTcpPortStats.Stats.Request_method_register,
			"request_method_refer":                     ret.DtDdosSipTcpPortStats.Stats.Request_method_refer,
			"request_method_subscribe":                 ret.DtDdosSipTcpPortStats.Stats.Request_method_subscribe,
			"request_method_update":                    ret.DtDdosSipTcpPortStats.Stats.Request_method_update,
			"request_method_unknown":                   ret.DtDdosSipTcpPortStats.Stats.Request_method_unknown,
			"request_unknown_version":                  ret.DtDdosSipTcpPortStats.Stats.Request_unknown_version,
			"keep_alive_msg":                           ret.DtDdosSipTcpPortStats.Stats.Keep_alive_msg,
			"rate1_limit_exceed":                       ret.DtDdosSipTcpPortStats.Stats.Rate1_limit_exceed,
			"rate2_limit_exceed":                       ret.DtDdosSipTcpPortStats.Stats.Rate2_limit_exceed,
			"rate3_limit_exceed":                       ret.DtDdosSipTcpPortStats.Stats.Rate3_limit_exceed,
			"src_rate1_limit_exceed":                   ret.DtDdosSipTcpPortStats.Stats.Src_rate1_limit_exceed,
			"src_rate2_limit_exceed":                   ret.DtDdosSipTcpPortStats.Stats.Src_rate2_limit_exceed,
			"src_rate3_limit_exceed":                   ret.DtDdosSipTcpPortStats.Stats.Src_rate3_limit_exceed,
			"response_1xx":                             ret.DtDdosSipTcpPortStats.Stats.Response_1xx,
			"response_2xx":                             ret.DtDdosSipTcpPortStats.Stats.Response_2xx,
			"response_3xx":                             ret.DtDdosSipTcpPortStats.Stats.Response_3xx,
			"response_4xx":                             ret.DtDdosSipTcpPortStats.Stats.Response_4xx,
			"response_5xx":                             ret.DtDdosSipTcpPortStats.Stats.Response_5xx,
			"response_6xx":                             ret.DtDdosSipTcpPortStats.Stats.Response_6xx,
			"response_unknown":                         ret.DtDdosSipTcpPortStats.Stats.Response_unknown,
			"response_unknown_version":                 ret.DtDdosSipTcpPortStats.Stats.Response_unknown_version,
			"read_start_line_error":                    ret.DtDdosSipTcpPortStats.Stats.Read_start_line_error,
			"invalid_start_line_error":                 ret.DtDdosSipTcpPortStats.Stats.Invalid_start_line_error,
			"parse_start_line_error":                   ret.DtDdosSipTcpPortStats.Stats.Parse_start_line_error,
			"line_too_long":                            ret.DtDdosSipTcpPortStats.Stats.Line_too_long,
			"line_mem_allocated":                       ret.DtDdosSipTcpPortStats.Stats.Line_mem_allocated,
			"line_mem_freed":                           ret.DtDdosSipTcpPortStats.Stats.Line_mem_freed,
			"max_uri_len_exceed":                       ret.DtDdosSipTcpPortStats.Stats.Max_uri_len_exceed,
			"too_many_header":                          ret.DtDdosSipTcpPortStats.Stats.Too_many_header,
			"invalid_header":                           ret.DtDdosSipTcpPortStats.Stats.Invalid_header,
			"header_name_too_long":                     ret.DtDdosSipTcpPortStats.Stats.Header_name_too_long,
			"parse_header_fail_error":                  ret.DtDdosSipTcpPortStats.Stats.Parse_header_fail_error,
			"max_header_value_len_exceed":              ret.DtDdosSipTcpPortStats.Stats.Max_header_value_len_exceed,
			"max_call_id_len_exceed":                   ret.DtDdosSipTcpPortStats.Stats.Max_call_id_len_exceed,
			"header_filter_match":                      ret.DtDdosSipTcpPortStats.Stats.Header_filter_match,
			"header_filter_not_match":                  ret.DtDdosSipTcpPortStats.Stats.Header_filter_not_match,
			"header_filter_none_match":                 ret.DtDdosSipTcpPortStats.Stats.Header_filter_none_match,
			"header_filter_action_drop":                ret.DtDdosSipTcpPortStats.Stats.Header_filter_action_drop,
			"header_filter_action_blacklist":           ret.DtDdosSipTcpPortStats.Stats.Header_filter_action_blacklist,
			"header_filter_action_whitelist":           ret.DtDdosSipTcpPortStats.Stats.Header_filter_action_whitelist,
			"header_filter_action_default_pass":        ret.DtDdosSipTcpPortStats.Stats.Header_filter_action_default_pass,
			"header_filter_filter1_match":              ret.DtDdosSipTcpPortStats.Stats.Header_filter_filter1_match,
			"header_filter_filter2_match":              ret.DtDdosSipTcpPortStats.Stats.Header_filter_filter2_match,
			"header_filter_filter3_match":              ret.DtDdosSipTcpPortStats.Stats.Header_filter_filter3_match,
			"header_filter_filter4_match":              ret.DtDdosSipTcpPortStats.Stats.Header_filter_filter4_match,
			"header_filter_filter5_match":              ret.DtDdosSipTcpPortStats.Stats.Header_filter_filter5_match,
			"max_sdp_len_exceed":                       ret.DtDdosSipTcpPortStats.Stats.Max_sdp_len_exceed,
			"body_too_big":                             ret.DtDdosSipTcpPortStats.Stats.Body_too_big,
			"get_content_fail_error":                   ret.DtDdosSipTcpPortStats.Stats.Get_content_fail_error,
			"concatenate_msg":                          ret.DtDdosSipTcpPortStats.Stats.Concatenate_msg,
			"mem_alloc_fail_error":                     ret.DtDdosSipTcpPortStats.Stats.Mem_alloc_fail_error,
			"malform_request":                          ret.DtDdosSipTcpPortStats.Stats.Malform_request,
			"port_rcvd":                                ret.DtDdosSipTcpPortStats.Stats.Port_rcvd,
			"port_drop":                                ret.DtDdosSipTcpPortStats.Stats.Port_drop,
			"port_pkt_sent":                            ret.DtDdosSipTcpPortStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":                     ret.DtDdosSipTcpPortStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Port_kbit_rate_exceed,
			"port_conn_rate_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Port_conn_rate_exceed,
			"port_conn_limm_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Port_conn_limm_exceed,
			"port_bytes":                               ret.DtDdosSipTcpPortStats.Stats.Port_bytes,
			"outbound_port_bytes":                      ret.DtDdosSipTcpPortStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                       ret.DtDdosSipTcpPortStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":                   ret.DtDdosSipTcpPortStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                          ret.DtDdosSipTcpPortStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                          ret.DtDdosSipTcpPortStats.Stats.Port_bytes_drop,
			"port_src_bl":                              ret.DtDdosSipTcpPortStats.Stats.Port_src_bl,
			"filter_auth_fail":                         ret.DtDdosSipTcpPortStats.Stats.Filter_auth_fail,
			"syn_auth_fail":                            ret.DtDdosSipTcpPortStats.Stats.Syn_auth_fail,
			"ack_auth_fail":                            ret.DtDdosSipTcpPortStats.Stats.Ack_auth_fail,
			"syn_cookie_fail":                          ret.DtDdosSipTcpPortStats.Stats.Syn_cookie_fail,
			"sess_create":                              ret.DtDdosSipTcpPortStats.Stats.Sess_create,
			"filter_action_blacklist":                  ret.DtDdosSipTcpPortStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                       ret.DtDdosSipTcpPortStats.Stats.Filter_action_drop,
			"filter_action_default_pass":               ret.DtDdosSipTcpPortStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":                  ret.DtDdosSipTcpPortStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":                    ret.DtDdosSipTcpPortStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_crate_src":                    ret.DtDdosSipTcpPortStats.Stats.Exceed_drop_crate_src,
			"exceed_drop_climit_src":                   ret.DtDdosSipTcpPortStats.Stats.Exceed_drop_climit_src,
			"exceed_drop_brate_src":                    ret.DtDdosSipTcpPortStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":                 ret.DtDdosSipTcpPortStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                       ret.DtDdosSipTcpPortStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":                 ret.DtDdosSipTcpPortStats.Stats.Outbound_port_bytes_drop,
			"exceed_drop_brate_src_pkt":                ret.DtDdosSipTcpPortStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":                ret.DtDdosSipTcpPortStats.Stats.Port_kbit_rate_exceed_pkt,
			"syn_cookie_sent":                          ret.DtDdosSipTcpPortStats.Stats.Syn_cookie_sent,
			"ack_retry_init":                           ret.DtDdosSipTcpPortStats.Stats.Ack_retry_init,
			"ack_retry_gap_drop":                       ret.DtDdosSipTcpPortStats.Stats.Ack_retry_gap_drop,
			"conn_prate_excd":                          ret.DtDdosSipTcpPortStats.Stats.Conn_prate_excd,
			"out_of_seq_excd":                          ret.DtDdosSipTcpPortStats.Stats.Out_of_seq_excd,
			"retransmit_excd":                          ret.DtDdosSipTcpPortStats.Stats.Retransmit_excd,
			"zero_window_excd":                         ret.DtDdosSipTcpPortStats.Stats.Zero_window_excd,
			"syn_retry_init":                           ret.DtDdosSipTcpPortStats.Stats.Syn_retry_init,
			"syn_retry_gap_drop":                       ret.DtDdosSipTcpPortStats.Stats.Syn_retry_gap_drop,
			"ack_retry_pass":                           ret.DtDdosSipTcpPortStats.Stats.Ack_retry_pass,
			"syn_retry_pass":                           ret.DtDdosSipTcpPortStats.Stats.Syn_retry_pass,
			"tcp_auth_drop":                            ret.DtDdosSipTcpPortStats.Stats.Tcp_auth_drop,
			"tcp_auth_resp":                            ret.DtDdosSipTcpPortStats.Stats.Tcp_auth_resp,
			"frag_rcvd":                                ret.DtDdosSipTcpPortStats.Stats.Frag_rcvd,
			"frag_drop":                                ret.DtDdosSipTcpPortStats.Stats.Frag_drop,
			"sess_create_inbound":                      ret.DtDdosSipTcpPortStats.Stats.Sess_create_inbound,
			"sess_create_outbound":                     ret.DtDdosSipTcpPortStats.Stats.Sess_create_outbound,
			"conn_create_from_syn":                     ret.DtDdosSipTcpPortStats.Stats.Conn_create_from_syn,
			"conn_create_from_ack":                     ret.DtDdosSipTcpPortStats.Stats.Conn_create_from_ack,
			"conn_close":                               ret.DtDdosSipTcpPortStats.Stats.Conn_close,
			"conn_close_w_rst":                         ret.DtDdosSipTcpPortStats.Stats.Conn_close_w_rst,
			"conn_close_w_fin":                         ret.DtDdosSipTcpPortStats.Stats.Conn_close_w_fin,
			"conn_close_w_idle":                        ret.DtDdosSipTcpPortStats.Stats.Conn_close_w_idle,
			"conn_close_half_open":                     ret.DtDdosSipTcpPortStats.Stats.Conn_close_half_open,
			"sess_aged":                                ret.DtDdosSipTcpPortStats.Stats.Sess_aged,
			"syn_drop":                                 ret.DtDdosSipTcpPortStats.Stats.Syn_drop,
			"syn_auth_pass":                            ret.DtDdosSipTcpPortStats.Stats.Syn_auth_pass,
			"unauth_drop":                              ret.DtDdosSipTcpPortStats.Stats.Unauth_drop,
			"rst_cookie_fail":                          ret.DtDdosSipTcpPortStats.Stats.Rst_cookie_fail,
			"syn_retry_failed":                         ret.DtDdosSipTcpPortStats.Stats.Syn_retry_failed,
			"src_syn_auth_fail":                        ret.DtDdosSipTcpPortStats.Stats.Src_syn_auth_fail,
			"src_syn_cookie_sent":                      ret.DtDdosSipTcpPortStats.Stats.Src_syn_cookie_sent,
			"src_syn_cookie_fail":                      ret.DtDdosSipTcpPortStats.Stats.Src_syn_cookie_fail,
			"src_unauth_drop":                          ret.DtDdosSipTcpPortStats.Stats.Src_unauth_drop,
			"src_rst_cookie_fail":                      ret.DtDdosSipTcpPortStats.Stats.Src_rst_cookie_fail,
			"src_syn_retry_init":                       ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_init,
			"src_syn_retry_gap_drop":                   ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_gap_drop,
			"src_syn_retry_failed":                     ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_failed,
			"src_ack_retry_init":                       ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_init,
			"src_ack_retry_gap_drop":                   ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_gap_drop,
			"src_ack_auth_fail":                        ret.DtDdosSipTcpPortStats.Stats.Src_ack_auth_fail,
			"src_out_of_seq_excd":                      ret.DtDdosSipTcpPortStats.Stats.Src_out_of_seq_excd,
			"src_retransmit_excd":                      ret.DtDdosSipTcpPortStats.Stats.Src_retransmit_excd,
			"src_zero_window_excd":                     ret.DtDdosSipTcpPortStats.Stats.Src_zero_window_excd,
			"src_conn_pkt_rate_excd":                   ret.DtDdosSipTcpPortStats.Stats.Src_conn_pkt_rate_excd,
			"src_filter_action_blacklist":              ret.DtDdosSipTcpPortStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":                   ret.DtDdosSipTcpPortStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":           ret.DtDdosSipTcpPortStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":              ret.DtDdosSipTcpPortStats.Stats.Src_filter_action_whitelist,
			"tcp_rexmit_syn_limit_drop":                ret.DtDdosSipTcpPortStats.Stats.Tcp_rexmit_syn_limit_drop,
			"tcp_rexmit_syn_limit_bl":                  ret.DtDdosSipTcpPortStats.Stats.Tcp_rexmit_syn_limit_bl,
			"conn_ofo_rate_excd":                       ret.DtDdosSipTcpPortStats.Stats.Conn_ofo_rate_excd,
			"conn_rexmit_rate_excd":                    ret.DtDdosSipTcpPortStats.Stats.Conn_rexmit_rate_excd,
			"conn_zwindow_rate_excd":                   ret.DtDdosSipTcpPortStats.Stats.Conn_zwindow_rate_excd,
			"src_conn_ofo_rate_excd":                   ret.DtDdosSipTcpPortStats.Stats.Src_conn_ofo_rate_excd,
			"src_conn_rexmit_rate_excd":                ret.DtDdosSipTcpPortStats.Stats.Src_conn_rexmit_rate_excd,
			"src_conn_zwindow_rate_excd":               ret.DtDdosSipTcpPortStats.Stats.Src_conn_zwindow_rate_excd,
			"ack_retry_rto_pass":                       ret.DtDdosSipTcpPortStats.Stats.Ack_retry_rto_pass,
			"ack_retry_rto_fail":                       ret.DtDdosSipTcpPortStats.Stats.Ack_retry_rto_fail,
			"ack_retry_rto_progress":                   ret.DtDdosSipTcpPortStats.Stats.Ack_retry_rto_progress,
			"syn_retry_rto_pass":                       ret.DtDdosSipTcpPortStats.Stats.Syn_retry_rto_pass,
			"syn_retry_rto_fail":                       ret.DtDdosSipTcpPortStats.Stats.Syn_retry_rto_fail,
			"syn_retry_rto_progress":                   ret.DtDdosSipTcpPortStats.Stats.Syn_retry_rto_progress,
			"src_syn_retry_rto_pass":                   ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_rto_pass,
			"src_syn_retry_rto_fail":                   ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_rto_fail,
			"src_syn_retry_rto_progress":               ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_rto_progress,
			"src_ack_retry_rto_pass":                   ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_rto_pass,
			"src_ack_retry_rto_fail":                   ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_rto_fail,
			"src_ack_retry_rto_progress":               ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_rto_progress,
			"wellknown_sport_drop":                     ret.DtDdosSipTcpPortStats.Stats.Wellknown_sport_drop,
			"src_well_known_port":                      ret.DtDdosSipTcpPortStats.Stats.Src_well_known_port,
			"src_auth_drop":                            ret.DtDdosSipTcpPortStats.Stats.Src_auth_drop,
			"src_frag_drop":                            ret.DtDdosSipTcpPortStats.Stats.Src_frag_drop,
			"frag_timeout":                             ret.DtDdosSipTcpPortStats.Stats.Frag_timeout,
			"filter1_match":                            ret.DtDdosSipTcpPortStats.Stats.Filter1_match,
			"filter2_match":                            ret.DtDdosSipTcpPortStats.Stats.Filter2_match,
			"filter3_match":                            ret.DtDdosSipTcpPortStats.Stats.Filter3_match,
			"filter4_match":                            ret.DtDdosSipTcpPortStats.Stats.Filter4_match,
			"filter5_match":                            ret.DtDdosSipTcpPortStats.Stats.Filter5_match,
			"filter_none_match":                        ret.DtDdosSipTcpPortStats.Stats.Filter_none_match,
			"filter_total_not_match":                   ret.DtDdosSipTcpPortStats.Stats.Filter_total_not_match,
			"bl":                                       ret.DtDdosSipTcpPortStats.Stats.Bl,
			"src_drop":                                 ret.DtDdosSipTcpPortStats.Stats.Src_drop,
			"create_conn_non_syn_dropped":              ret.DtDdosSipTcpPortStats.Stats.Create_conn_non_syn_dropped,
			"src_create_conn_non_syn_dropped":          ret.DtDdosSipTcpPortStats.Stats.Src_create_conn_non_syn_dropped,
			"port_syn_rate_exceed":                     ret.DtDdosSipTcpPortStats.Stats.Port_syn_rate_exceed,
			"src_syn_rate_exceed":                      ret.DtDdosSipTcpPortStats.Stats.Src_syn_rate_exceed,
			"src_filter1_match":                        ret.DtDdosSipTcpPortStats.Stats.Src_filter1_match,
			"src_filter2_match":                        ret.DtDdosSipTcpPortStats.Stats.Src_filter2_match,
			"src_filter3_match":                        ret.DtDdosSipTcpPortStats.Stats.Src_filter3_match,
			"src_filter4_match":                        ret.DtDdosSipTcpPortStats.Stats.Src_filter4_match,
			"src_filter5_match":                        ret.DtDdosSipTcpPortStats.Stats.Src_filter5_match,
			"src_filter_none_match":                    ret.DtDdosSipTcpPortStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":               ret.DtDdosSipTcpPortStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":                     ret.DtDdosSipTcpPortStats.Stats.Src_filter_auth_fail,
			"syn_tfo_rcv":                              ret.DtDdosSipTcpPortStats.Stats.Syn_tfo_rcv,
			"ack_retry_timeout":                        ret.DtDdosSipTcpPortStats.Stats.Ack_retry_timeout,
			"ack_retry_reset":                          ret.DtDdosSipTcpPortStats.Stats.Ack_retry_reset,
			"ack_retry_blacklist":                      ret.DtDdosSipTcpPortStats.Stats.Ack_retry_blacklist,
			"src_ack_retry_timeout":                    ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_timeout,
			"src_ack_retry_reset":                      ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_reset,
			"src_ack_retry_blacklist":                  ret.DtDdosSipTcpPortStats.Stats.Src_ack_retry_blacklist,
			"syn_retry_timeout":                        ret.DtDdosSipTcpPortStats.Stats.Syn_retry_timeout,
			"syn_retry_reset":                          ret.DtDdosSipTcpPortStats.Stats.Syn_retry_reset,
			"syn_retry_blacklist":                      ret.DtDdosSipTcpPortStats.Stats.Syn_retry_blacklist,
			"src_syn_retry_timeout":                    ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_timeout,
			"src_syn_retry_reset":                      ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_reset,
			"src_syn_retry_blacklist":                  ret.DtDdosSipTcpPortStats.Stats.Src_syn_retry_blacklist,
			"sflow_internal_samples_packed":            ret.DtDdosSipTcpPortStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":            ret.DtDdosSipTcpPortStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":              ret.DtDdosSipTcpPortStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":              ret.DtDdosSipTcpPortStats.Stats.Sflow_external_packets_sent,
			"exceed_action_tunnel":                     ret.DtDdosSipTcpPortStats.Stats.Exceed_action_tunnel,
			"pattern_recognition_proceeded":            ret.DtDdosSipTcpPortStats.Stats.Pattern_recognition_proceeded,
			"pattern_not_found":                        ret.DtDdosSipTcpPortStats.Stats.Pattern_not_found,
			"pattern_recognition_generic_error":        ret.DtDdosSipTcpPortStats.Stats.Pattern_recognition_generic_error,
			"pattern_filter1_match":                    ret.DtDdosSipTcpPortStats.Stats.Pattern_filter1_match,
			"pattern_filter2_match":                    ret.DtDdosSipTcpPortStats.Stats.Pattern_filter2_match,
			"pattern_filter3_match":                    ret.DtDdosSipTcpPortStats.Stats.Pattern_filter3_match,
			"pattern_filter4_match":                    ret.DtDdosSipTcpPortStats.Stats.Pattern_filter4_match,
			"pattern_filter5_match":                    ret.DtDdosSipTcpPortStats.Stats.Pattern_filter5_match,
			"pattern_filter_drop":                      ret.DtDdosSipTcpPortStats.Stats.Pattern_filter_drop,
			"pattern_recognition_sampling_started":     ret.DtDdosSipTcpPortStats.Stats.Pattern_recognition_sampling_started,
			"pattern_recognition_pattern_changed":      ret.DtDdosSipTcpPortStats.Stats.Pattern_recognition_pattern_changed,
			"dst_hw_drop":                              ret.DtDdosSipTcpPortStats.Stats.Dst_hw_drop,
			"synack_reset_sent":                        ret.DtDdosSipTcpPortStats.Stats.Synack_reset_sent,
			"synack_multiple_attempts_per_ip_detected": ret.DtDdosSipTcpPortStats.Stats.Synack_multiple_attempts_per_ip_detected,
			"prog_first_req_time_exceed":               ret.DtDdosSipTcpPortStats.Stats.Prog_first_req_time_exceed,
			"prog_req_resp_time_exceed":                ret.DtDdosSipTcpPortStats.Stats.Prog_req_resp_time_exceed,
			"prog_request_len_exceed":                  ret.DtDdosSipTcpPortStats.Stats.Prog_request_len_exceed,
			"prog_response_len_exceed":                 ret.DtDdosSipTcpPortStats.Stats.Prog_response_len_exceed,
			"prog_resp_req_ratio_exceed":               ret.DtDdosSipTcpPortStats.Stats.Prog_resp_req_ratio_exceed,
			"prog_resp_req_time_exceed":                ret.DtDdosSipTcpPortStats.Stats.Prog_resp_req_time_exceed,
			"prog_conn_sent_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Prog_conn_sent_exceed,
			"prog_conn_rcvd_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Prog_conn_rcvd_exceed,
			"prog_conn_time_exceed":                    ret.DtDdosSipTcpPortStats.Stats.Prog_conn_time_exceed,
			"prog_conn_rcvd_sent_ratio_exceed":         ret.DtDdosSipTcpPortStats.Stats.Prog_conn_rcvd_sent_ratio_exceed,
			"prog_win_sent_exceed":                     ret.DtDdosSipTcpPortStats.Stats.Prog_win_sent_exceed,
			"prog_win_rcvd_exceed":                     ret.DtDdosSipTcpPortStats.Stats.Prog_win_rcvd_exceed,
			"prog_win_rcvd_sent_ratio_exceed":          ret.DtDdosSipTcpPortStats.Stats.Prog_win_rcvd_sent_ratio_exceed,
			"snat_fail":                                ret.DtDdosSipTcpPortStats.Stats.Snat_fail,
			"prog_exceed_drop":                         ret.DtDdosSipTcpPortStats.Stats.Prog_exceed_drop,
			"prog_exceed_bl":                           ret.DtDdosSipTcpPortStats.Stats.Prog_exceed_bl,
			"prog_conn_exceed_drop":                    ret.DtDdosSipTcpPortStats.Stats.Prog_conn_exceed_drop,
			"prog_conn_exceed_bl":                      ret.DtDdosSipTcpPortStats.Stats.Prog_conn_exceed_bl,
			"prog_win_exceed_drop":                     ret.DtDdosSipTcpPortStats.Stats.Prog_win_exceed_drop,
			"prog_win_exceed_bl":                       ret.DtDdosSipTcpPortStats.Stats.Prog_win_exceed_bl,
			"exceed_action_drop":                       ret.DtDdosSipTcpPortStats.Stats.Exceed_action_drop,
			"syn_auth_rst_ack_drop":                    ret.DtDdosSipTcpPortStats.Stats.Syn_auth_rst_ack_drop,
			"prog_exceed_reset":                        ret.DtDdosSipTcpPortStats.Stats.Prog_exceed_reset,
			"prog_conn_exceed_reset":                   ret.DtDdosSipTcpPortStats.Stats.Prog_conn_exceed_reset,
			"prog_win_exceed_reset":                    ret.DtDdosSipTcpPortStats.Stats.Prog_win_exceed_reset,
			"conn_create_from_synack":                  ret.DtDdosSipTcpPortStats.Stats.Conn_create_from_synack,
			"port_synack_rate_exceed":                  ret.DtDdosSipTcpPortStats.Stats.Port_synack_rate_exceed,
			"prog_conn_samples":                        ret.DtDdosSipTcpPortStats.Stats.Prog_conn_samples,
			"prog_req_samples":                         ret.DtDdosSipTcpPortStats.Stats.Prog_req_samples,
			"prog_win_samples":                         ret.DtDdosSipTcpPortStats.Stats.Prog_win_samples,
			"prog_conn_samples_processed":              ret.DtDdosSipTcpPortStats.Stats.Prog_conn_samples_processed,
			"prog_req_samples_processed":               ret.DtDdosSipTcpPortStats.Stats.Prog_req_samples_processed,
			"prog_win_samples_processed":               ret.DtDdosSipTcpPortStats.Stats.Prog_win_samples_processed,
			"tcp_auth_rst":                             ret.DtDdosSipTcpPortStats.Stats.Tcp_auth_rst,
			"src_tcp_auth_rst":                         ret.DtDdosSipTcpPortStats.Stats.Src_tcp_auth_rst,
		},
	}
}

func getObjectDdosSipTcpPortStatsStats(d []interface{}) edpt.DdosSipTcpPortStatsStats {

	count1 := len(d)
	var ret edpt.DdosSipTcpPortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Policy_drop = in["policy_drop"].(int)
		ret.Policy_violation = in["policy_violation"].(int)
		ret.Idle_timeout = in["idle_timeout"].(int)
		ret.Ofo_timeout = in["ofo_timeout"].(int)
		ret.Seq_check_ofo = in["seq_check_ofo"].(int)
		ret.Pkts_ofo_total = in["pkts_ofo_total"].(int)
		ret.Ofo_queue_size_exceed = in["ofo_queue_size_exceed"].(int)
		ret.Seq_check_retrans_fin = in["seq_check_retrans_fin"].(int)
		ret.Seq_check_retrans_rst = in["seq_check_retrans_rst"].(int)
		ret.Seq_check_retrans_push = in["seq_check_retrans_push"].(int)
		ret.Seq_check_retrans_other = in["seq_check_retrans_other"].(int)
		ret.Pkts_retrans_total = in["pkts_retrans_total"].(int)
		ret.Client_rst = in["client_rst"].(int)
		ret.Error_condition = in["error_condition"].(int)
		ret.Request_method_ack = in["request_method_ack"].(int)
		ret.Request_method_bye = in["request_method_bye"].(int)
		ret.Request_method_cancel = in["request_method_cancel"].(int)
		ret.Request_method_invite = in["request_method_invite"].(int)
		ret.Request_method_info = in["request_method_info"].(int)
		ret.Request_method_message = in["request_method_message"].(int)
		ret.Request_method_notify = in["request_method_notify"].(int)
		ret.Request_method_options = in["request_method_options"].(int)
		ret.Request_method_prack = in["request_method_prack"].(int)
		ret.Request_method_publish = in["request_method_publish"].(int)
		ret.Request_method_register = in["request_method_register"].(int)
		ret.Request_method_refer = in["request_method_refer"].(int)
		ret.Request_method_subscribe = in["request_method_subscribe"].(int)
		ret.Request_method_update = in["request_method_update"].(int)
		ret.Request_method_unknown = in["request_method_unknown"].(int)
		ret.Request_unknown_version = in["request_unknown_version"].(int)
		ret.Keep_alive_msg = in["keep_alive_msg"].(int)
		ret.Rate1_limit_exceed = in["rate1_limit_exceed"].(int)
		ret.Rate2_limit_exceed = in["rate2_limit_exceed"].(int)
		ret.Rate3_limit_exceed = in["rate3_limit_exceed"].(int)
		ret.Src_rate1_limit_exceed = in["src_rate1_limit_exceed"].(int)
		ret.Src_rate2_limit_exceed = in["src_rate2_limit_exceed"].(int)
		ret.Src_rate3_limit_exceed = in["src_rate3_limit_exceed"].(int)
		ret.Response_1xx = in["response_1xx"].(int)
		ret.Response_2xx = in["response_2xx"].(int)
		ret.Response_3xx = in["response_3xx"].(int)
		ret.Response_4xx = in["response_4xx"].(int)
		ret.Response_5xx = in["response_5xx"].(int)
		ret.Response_6xx = in["response_6xx"].(int)
		ret.Response_unknown = in["response_unknown"].(int)
		ret.Response_unknown_version = in["response_unknown_version"].(int)
		ret.Read_start_line_error = in["read_start_line_error"].(int)
		ret.Invalid_start_line_error = in["invalid_start_line_error"].(int)
		ret.Parse_start_line_error = in["parse_start_line_error"].(int)
		ret.Line_too_long = in["line_too_long"].(int)
		ret.Line_mem_allocated = in["line_mem_allocated"].(int)
		ret.Line_mem_freed = in["line_mem_freed"].(int)
		ret.Max_uri_len_exceed = in["max_uri_len_exceed"].(int)
		ret.Too_many_header = in["too_many_header"].(int)
		ret.Invalid_header = in["invalid_header"].(int)
		ret.Header_name_too_long = in["header_name_too_long"].(int)
		ret.Parse_header_fail_error = in["parse_header_fail_error"].(int)
		ret.Max_header_value_len_exceed = in["max_header_value_len_exceed"].(int)
		ret.Max_call_id_len_exceed = in["max_call_id_len_exceed"].(int)
		ret.Header_filter_match = in["header_filter_match"].(int)
		ret.Header_filter_not_match = in["header_filter_not_match"].(int)
		ret.Header_filter_none_match = in["header_filter_none_match"].(int)
		ret.Header_filter_action_drop = in["header_filter_action_drop"].(int)
		ret.Header_filter_action_blacklist = in["header_filter_action_blacklist"].(int)
		ret.Header_filter_action_whitelist = in["header_filter_action_whitelist"].(int)
		ret.Header_filter_action_default_pass = in["header_filter_action_default_pass"].(int)
		ret.Header_filter_filter1_match = in["header_filter_filter1_match"].(int)
		ret.Header_filter_filter2_match = in["header_filter_filter2_match"].(int)
		ret.Header_filter_filter3_match = in["header_filter_filter3_match"].(int)
		ret.Header_filter_filter4_match = in["header_filter_filter4_match"].(int)
		ret.Header_filter_filter5_match = in["header_filter_filter5_match"].(int)
		ret.Max_sdp_len_exceed = in["max_sdp_len_exceed"].(int)
		ret.Body_too_big = in["body_too_big"].(int)
		ret.Get_content_fail_error = in["get_content_fail_error"].(int)
		ret.Concatenate_msg = in["concatenate_msg"].(int)
		ret.Mem_alloc_fail_error = in["mem_alloc_fail_error"].(int)
		ret.Malform_request = in["malform_request"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
		ret.Port_conn_rate_exceed = in["port_conn_rate_exceed"].(int)
		ret.Port_conn_limm_exceed = in["port_conn_limm_exceed"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Filter_auth_fail = in["filter_auth_fail"].(int)
		ret.Syn_auth_fail = in["syn_auth_fail"].(int)
		ret.Ack_auth_fail = in["ack_auth_fail"].(int)
		ret.Syn_cookie_fail = in["syn_cookie_fail"].(int)
		ret.Sess_create = in["sess_create"].(int)
		ret.Filter_action_blacklist = in["filter_action_blacklist"].(int)
		ret.Filter_action_drop = in["filter_action_drop"].(int)
		ret.Filter_action_default_pass = in["filter_action_default_pass"].(int)
		ret.Filter_action_whitelist = in["filter_action_whitelist"].(int)
		ret.Exceed_drop_prate_src = in["exceed_drop_prate_src"].(int)
		ret.Exceed_drop_crate_src = in["exceed_drop_crate_src"].(int)
		ret.Exceed_drop_climit_src = in["exceed_drop_climit_src"].(int)
		ret.Exceed_drop_brate_src = in["exceed_drop_brate_src"].(int)
		ret.Outbound_port_bytes_sent = in["outbound_port_bytes_sent"].(int)
		ret.Outbound_port_drop = in["outbound_port_drop"].(int)
		ret.Outbound_port_bytes_drop = in["outbound_port_bytes_drop"].(int)
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
		ret.Tcp_auth_drop = in["tcp_auth_drop"].(int)
		ret.Tcp_auth_resp = in["tcp_auth_resp"].(int)
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
		ret.Src_auth_drop = in["src_auth_drop"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Filter1_match = in["filter1_match"].(int)
		ret.Filter2_match = in["filter2_match"].(int)
		ret.Filter3_match = in["filter3_match"].(int)
		ret.Filter4_match = in["filter4_match"].(int)
		ret.Filter5_match = in["filter5_match"].(int)
		ret.Filter_none_match = in["filter_none_match"].(int)
		ret.Filter_total_not_match = in["filter_total_not_match"].(int)
		ret.Bl = in["bl"].(int)
		ret.Src_drop = in["src_drop"].(int)
		ret.Create_conn_non_syn_dropped = in["create_conn_non_syn_dropped"].(int)
		ret.Src_create_conn_non_syn_dropped = in["src_create_conn_non_syn_dropped"].(int)
		ret.Port_syn_rate_exceed = in["port_syn_rate_exceed"].(int)
		ret.Src_syn_rate_exceed = in["src_syn_rate_exceed"].(int)
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
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
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
		ret.Prog_conn_samples_processed = in["prog_conn_samples_processed"].(int)
		ret.Prog_req_samples_processed = in["prog_req_samples_processed"].(int)
		ret.Prog_win_samples_processed = in["prog_win_samples_processed"].(int)
		ret.Tcp_auth_rst = in["tcp_auth_rst"].(int)
		ret.Src_tcp_auth_rst = in["src_tcp_auth_rst"].(int)
	}
	return ret
}

func dataToEndpointDdosSipTcpPortStats(d *schema.ResourceData) edpt.DdosSipTcpPortStats {
	var ret edpt.DdosSipTcpPortStats

	ret.Stats = getObjectDdosSipTcpPortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
