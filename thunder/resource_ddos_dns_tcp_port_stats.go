package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsTcpPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dns_tcp_port_stats`: Statistics for the object dns-tcp-port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDnsTcpPortStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_limit_type0": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 1 Exceeded",
						},
						"rate_limit_type1": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 2 Exceeded",
						},
						"rate_limit_type2": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 3 Exceeded",
						},
						"rate_limit_type3": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 4 Exceeded",
						},
						"rate_limit_type4": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Request Rate 5 Exceeded",
						},
						"nxdomain_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Rate Exceeded",
						},
						"is_nxdomain": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response",
						},
						"nxdomain_bl_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Blacklisted",
						},
						"dns_fqdn_label_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Label Length Exceeded",
						},
						"dns_malform_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Query Dropped",
						},
						"dns_fqdn_stage_2_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate Exceeded",
						},
						"req_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Requests Forwarded",
						},
						"req_incomplete": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Request Incomplete",
						},
						"req_size_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Request Size Exceeded",
						},
						"req_retrans": {
							Type: schema.TypeInt, Optional: true, Description: "Request Retransmit",
						},
						"query_type_a": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type A",
						},
						"query_type_aaaa": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type AAAA",
						},
						"query_type_ns": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type NS",
						},
						"query_type_cname": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type CNAME",
						},
						"query_type_any": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type ANY",
						},
						"query_type_srv": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type SRV",
						},
						"query_type_mx": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type MX",
						},
						"query_type_soa": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type SOA",
						},
						"query_type_opt": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type OPT",
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
						"dg_action_permit": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Action Permit",
						},
						"dg_action_deny": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Action Deny",
						},
						"dg_action_default": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Action Default",
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
						"src_rate_limit_type0": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 1 Exceeded",
						},
						"src_rate_limit_type1": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 2 Exceeded",
						},
						"src_rate_limit_type2": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 3 Exceeded",
						},
						"src_rate_limit_type3": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 4 Exceeded",
						},
						"src_rate_limit_type4": {
							Type: schema.TypeInt, Optional: true, Description: "Src Request Rate 5 Exceeded",
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
						"fqdn_rate_by_label_count_check": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate by Label Count Checked",
						},
						"fqdn_rate_by_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate by Label Count Exceeded",
						},
						"policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Dropped",
						},
						"fqdn_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Label Count Exceeded",
						},
						"rrtype_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Record Type Dropped",
						},
						"src_policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Policy Dropped",
						},
						"src_dns_fqdn_label_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src FQDN Label Length Exceeded",
						},
						"src_fqdn_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src FQDN Label Count Exceeded",
						},
						"src_dns_malform_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Malform Query Dropped",
						},
						"tcp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Dropped",
						},
						"tcp_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Auth Responded",
						},
						"dns_query_class_in": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class INTERNET",
						},
						"dns_query_class_csnet": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class CSNET",
						},
						"dns_query_class_chaos": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class CHAOS",
						},
						"dns_query_class_hs": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class HESIOD",
						},
						"dns_query_class_none": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class NONE",
						},
						"dns_query_class_any": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class ANY",
						},
						"dns_query_class_whitelist_miss": {
							Type: schema.TypeInt, Optional: true, Description: "Query Class Whitelist Miss",
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
						"force_tcp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Force-TCP Passed",
						},
						"nxdomain_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Dropped",
						},
						"dns_malform": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Query",
						},
						"src_dns_malform": {
							Type: schema.TypeInt, Optional: true, Description: "Src Malform Query",
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
						"dg_request_check_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Request Check Fail",
						},
						"query_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type ANY Dropped",
						},
						"src_query_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Query Type ANY Dropped",
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
						"create_conn_non_syn_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Create Conn with non-SYN Packets Dropped",
						},
						"src_create_conn_non_syn_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Src Create Conn with non-SYN Packets Dropped",
						},
						"dg_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Domain Query Rate Exceeded",
						},
						"port_syn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP SYN Rate Exceeded",
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
						"src_syn_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src TCP SYN Rate Exceeded",
						},
						"non_query_opcode_pass_through": {
							Type: schema.TypeInt, Optional: true, Description: "Non Query Opcode Pass Through",
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
						"pattern_recognition_sampling_started": {
							Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Sampling Started",
						},
						"pattern_recognition_pattern_changed": {
							Type: schema.TypeInt, Optional: true, Description: "Pattern Recognition: Pattern Change Detected",
						},
						"exceed_action_tunnel": {
							Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Tunnel",
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

func resourceDdosDnsTcpPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsTcpPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsTcpPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDnsTcpPortStatsStats := setObjectDdosDnsTcpPortStatsStats(res)
		d.Set("stats", DdosDnsTcpPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDnsTcpPortStatsStats(ret edpt.DataDdosDnsTcpPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rate_limit_type0":                         ret.DtDdosDnsTcpPortStats.Stats.Rate_limit_type0,
			"rate_limit_type1":                         ret.DtDdosDnsTcpPortStats.Stats.Rate_limit_type1,
			"rate_limit_type2":                         ret.DtDdosDnsTcpPortStats.Stats.Rate_limit_type2,
			"rate_limit_type3":                         ret.DtDdosDnsTcpPortStats.Stats.Rate_limit_type3,
			"rate_limit_type4":                         ret.DtDdosDnsTcpPortStats.Stats.Rate_limit_type4,
			"nxdomain_rate_exceed":                     ret.DtDdosDnsTcpPortStats.Stats.Nxdomain_rate_exceed,
			"is_nxdomain":                              ret.DtDdosDnsTcpPortStats.Stats.Is_nxdomain,
			"nxdomain_bl_drop":                         ret.DtDdosDnsTcpPortStats.Stats.Nxdomain_bl_drop,
			"dns_fqdn_label_len_exceed":                ret.DtDdosDnsTcpPortStats.Stats.Dns_fqdn_label_len_exceed,
			"dns_malform_drop":                         ret.DtDdosDnsTcpPortStats.Stats.Dns_malform_drop,
			"dns_fqdn_stage_2_rate_exceed":             ret.DtDdosDnsTcpPortStats.Stats.Dns_fqdn_stage_2_rate_exceed,
			"req_sent":                                 ret.DtDdosDnsTcpPortStats.Stats.Req_sent,
			"req_incomplete":                           ret.DtDdosDnsTcpPortStats.Stats.Req_incomplete,
			"req_size_exceed":                          ret.DtDdosDnsTcpPortStats.Stats.Req_size_exceed,
			"req_retrans":                              ret.DtDdosDnsTcpPortStats.Stats.Req_retrans,
			"query_type_a":                             ret.DtDdosDnsTcpPortStats.Stats.Query_type_a,
			"query_type_aaaa":                          ret.DtDdosDnsTcpPortStats.Stats.Query_type_aaaa,
			"query_type_ns":                            ret.DtDdosDnsTcpPortStats.Stats.Query_type_ns,
			"query_type_cname":                         ret.DtDdosDnsTcpPortStats.Stats.Query_type_cname,
			"query_type_any":                           ret.DtDdosDnsTcpPortStats.Stats.Query_type_any,
			"query_type_srv":                           ret.DtDdosDnsTcpPortStats.Stats.Query_type_srv,
			"query_type_mx":                            ret.DtDdosDnsTcpPortStats.Stats.Query_type_mx,
			"query_type_soa":                           ret.DtDdosDnsTcpPortStats.Stats.Query_type_soa,
			"query_type_opt":                           ret.DtDdosDnsTcpPortStats.Stats.Query_type_opt,
			"port_rcvd":                                ret.DtDdosDnsTcpPortStats.Stats.Port_rcvd,
			"port_drop":                                ret.DtDdosDnsTcpPortStats.Stats.Port_drop,
			"port_pkt_sent":                            ret.DtDdosDnsTcpPortStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":                     ret.DtDdosDnsTcpPortStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":                    ret.DtDdosDnsTcpPortStats.Stats.Port_kbit_rate_exceed,
			"port_conn_rate_exceed":                    ret.DtDdosDnsTcpPortStats.Stats.Port_conn_rate_exceed,
			"port_conn_limm_exceed":                    ret.DtDdosDnsTcpPortStats.Stats.Port_conn_limm_exceed,
			"dg_action_permit":                         ret.DtDdosDnsTcpPortStats.Stats.Dg_action_permit,
			"dg_action_deny":                           ret.DtDdosDnsTcpPortStats.Stats.Dg_action_deny,
			"dg_action_default":                        ret.DtDdosDnsTcpPortStats.Stats.Dg_action_default,
			"port_bytes":                               ret.DtDdosDnsTcpPortStats.Stats.Port_bytes,
			"outbound_port_bytes":                      ret.DtDdosDnsTcpPortStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                       ret.DtDdosDnsTcpPortStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":                   ret.DtDdosDnsTcpPortStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                          ret.DtDdosDnsTcpPortStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                          ret.DtDdosDnsTcpPortStats.Stats.Port_bytes_drop,
			"port_src_bl":                              ret.DtDdosDnsTcpPortStats.Stats.Port_src_bl,
			"filter_auth_fail":                         ret.DtDdosDnsTcpPortStats.Stats.Filter_auth_fail,
			"syn_auth_fail":                            ret.DtDdosDnsTcpPortStats.Stats.Syn_auth_fail,
			"ack_auth_fail":                            ret.DtDdosDnsTcpPortStats.Stats.Ack_auth_fail,
			"syn_cookie_fail":                          ret.DtDdosDnsTcpPortStats.Stats.Syn_cookie_fail,
			"sess_create":                              ret.DtDdosDnsTcpPortStats.Stats.Sess_create,
			"filter_action_blacklist":                  ret.DtDdosDnsTcpPortStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                       ret.DtDdosDnsTcpPortStats.Stats.Filter_action_drop,
			"filter_action_default_pass":               ret.DtDdosDnsTcpPortStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":                  ret.DtDdosDnsTcpPortStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":                    ret.DtDdosDnsTcpPortStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_crate_src":                    ret.DtDdosDnsTcpPortStats.Stats.Exceed_drop_crate_src,
			"exceed_drop_climit_src":                   ret.DtDdosDnsTcpPortStats.Stats.Exceed_drop_climit_src,
			"exceed_drop_brate_src":                    ret.DtDdosDnsTcpPortStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":                 ret.DtDdosDnsTcpPortStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                       ret.DtDdosDnsTcpPortStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":                 ret.DtDdosDnsTcpPortStats.Stats.Outbound_port_bytes_drop,
			"src_rate_limit_type0":                     ret.DtDdosDnsTcpPortStats.Stats.Src_rate_limit_type0,
			"src_rate_limit_type1":                     ret.DtDdosDnsTcpPortStats.Stats.Src_rate_limit_type1,
			"src_rate_limit_type2":                     ret.DtDdosDnsTcpPortStats.Stats.Src_rate_limit_type2,
			"src_rate_limit_type3":                     ret.DtDdosDnsTcpPortStats.Stats.Src_rate_limit_type3,
			"src_rate_limit_type4":                     ret.DtDdosDnsTcpPortStats.Stats.Src_rate_limit_type4,
			"exceed_drop_brate_src_pkt":                ret.DtDdosDnsTcpPortStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":                ret.DtDdosDnsTcpPortStats.Stats.Port_kbit_rate_exceed_pkt,
			"syn_cookie_sent":                          ret.DtDdosDnsTcpPortStats.Stats.Syn_cookie_sent,
			"ack_retry_init":                           ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_init,
			"ack_retry_gap_drop":                       ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_gap_drop,
			"conn_prate_excd":                          ret.DtDdosDnsTcpPortStats.Stats.Conn_prate_excd,
			"out_of_seq_excd":                          ret.DtDdosDnsTcpPortStats.Stats.Out_of_seq_excd,
			"retransmit_excd":                          ret.DtDdosDnsTcpPortStats.Stats.Retransmit_excd,
			"zero_window_excd":                         ret.DtDdosDnsTcpPortStats.Stats.Zero_window_excd,
			"syn_retry_init":                           ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_init,
			"syn_retry_gap_drop":                       ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_gap_drop,
			"ack_retry_pass":                           ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_pass,
			"syn_retry_pass":                           ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_pass,
			"fqdn_rate_by_label_count_check":           ret.DtDdosDnsTcpPortStats.Stats.Fqdn_rate_by_label_count_check,
			"fqdn_rate_by_label_count_exceed":          ret.DtDdosDnsTcpPortStats.Stats.Fqdn_rate_by_label_count_exceed,
			"policy_drop":                              ret.DtDdosDnsTcpPortStats.Stats.Policy_drop,
			"fqdn_label_count_exceed":                  ret.DtDdosDnsTcpPortStats.Stats.Fqdn_label_count_exceed,
			"rrtype_drop":                              ret.DtDdosDnsTcpPortStats.Stats.Rrtype_drop,
			"src_policy_drop":                          ret.DtDdosDnsTcpPortStats.Stats.Src_policy_drop,
			"src_dns_fqdn_label_len_exceed":            ret.DtDdosDnsTcpPortStats.Stats.Src_dns_fqdn_label_len_exceed,
			"src_fqdn_label_count_exceed":              ret.DtDdosDnsTcpPortStats.Stats.Src_fqdn_label_count_exceed,
			"src_dns_malform_drop":                     ret.DtDdosDnsTcpPortStats.Stats.Src_dns_malform_drop,
			"tcp_auth_drop":                            ret.DtDdosDnsTcpPortStats.Stats.Tcp_auth_drop,
			"tcp_auth_resp":                            ret.DtDdosDnsTcpPortStats.Stats.Tcp_auth_resp,
			"dns_query_class_in":                       ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_in,
			"dns_query_class_csnet":                    ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_csnet,
			"dns_query_class_chaos":                    ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_chaos,
			"dns_query_class_hs":                       ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_hs,
			"dns_query_class_none":                     ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_none,
			"dns_query_class_any":                      ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_any,
			"dns_query_class_whitelist_miss":           ret.DtDdosDnsTcpPortStats.Stats.Dns_query_class_whitelist_miss,
			"bl":                                       ret.DtDdosDnsTcpPortStats.Stats.Bl,
			"src_drop":                                 ret.DtDdosDnsTcpPortStats.Stats.Src_drop,
			"frag_rcvd":                                ret.DtDdosDnsTcpPortStats.Stats.Frag_rcvd,
			"frag_drop":                                ret.DtDdosDnsTcpPortStats.Stats.Frag_drop,
			"sess_create_inbound":                      ret.DtDdosDnsTcpPortStats.Stats.Sess_create_inbound,
			"sess_create_outbound":                     ret.DtDdosDnsTcpPortStats.Stats.Sess_create_outbound,
			"conn_create_from_syn":                     ret.DtDdosDnsTcpPortStats.Stats.Conn_create_from_syn,
			"conn_create_from_ack":                     ret.DtDdosDnsTcpPortStats.Stats.Conn_create_from_ack,
			"conn_close":                               ret.DtDdosDnsTcpPortStats.Stats.Conn_close,
			"conn_close_w_rst":                         ret.DtDdosDnsTcpPortStats.Stats.Conn_close_w_rst,
			"conn_close_w_fin":                         ret.DtDdosDnsTcpPortStats.Stats.Conn_close_w_fin,
			"conn_close_w_idle":                        ret.DtDdosDnsTcpPortStats.Stats.Conn_close_w_idle,
			"conn_close_half_open":                     ret.DtDdosDnsTcpPortStats.Stats.Conn_close_half_open,
			"sess_aged":                                ret.DtDdosDnsTcpPortStats.Stats.Sess_aged,
			"syn_drop":                                 ret.DtDdosDnsTcpPortStats.Stats.Syn_drop,
			"syn_auth_pass":                            ret.DtDdosDnsTcpPortStats.Stats.Syn_auth_pass,
			"unauth_drop":                              ret.DtDdosDnsTcpPortStats.Stats.Unauth_drop,
			"rst_cookie_fail":                          ret.DtDdosDnsTcpPortStats.Stats.Rst_cookie_fail,
			"syn_retry_failed":                         ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_failed,
			"filter1_match":                            ret.DtDdosDnsTcpPortStats.Stats.Filter1_match,
			"filter2_match":                            ret.DtDdosDnsTcpPortStats.Stats.Filter2_match,
			"filter3_match":                            ret.DtDdosDnsTcpPortStats.Stats.Filter3_match,
			"filter4_match":                            ret.DtDdosDnsTcpPortStats.Stats.Filter4_match,
			"filter5_match":                            ret.DtDdosDnsTcpPortStats.Stats.Filter5_match,
			"filter_none_match":                        ret.DtDdosDnsTcpPortStats.Stats.Filter_none_match,
			"filter_total_not_match":                   ret.DtDdosDnsTcpPortStats.Stats.Filter_total_not_match,
			"src_syn_auth_fail":                        ret.DtDdosDnsTcpPortStats.Stats.Src_syn_auth_fail,
			"src_syn_cookie_sent":                      ret.DtDdosDnsTcpPortStats.Stats.Src_syn_cookie_sent,
			"src_syn_cookie_fail":                      ret.DtDdosDnsTcpPortStats.Stats.Src_syn_cookie_fail,
			"src_unauth_drop":                          ret.DtDdosDnsTcpPortStats.Stats.Src_unauth_drop,
			"src_rst_cookie_fail":                      ret.DtDdosDnsTcpPortStats.Stats.Src_rst_cookie_fail,
			"src_syn_retry_init":                       ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_init,
			"src_syn_retry_gap_drop":                   ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_gap_drop,
			"src_syn_retry_failed":                     ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_failed,
			"src_ack_retry_init":                       ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_init,
			"src_ack_retry_gap_drop":                   ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_gap_drop,
			"src_ack_auth_fail":                        ret.DtDdosDnsTcpPortStats.Stats.Src_ack_auth_fail,
			"src_out_of_seq_excd":                      ret.DtDdosDnsTcpPortStats.Stats.Src_out_of_seq_excd,
			"src_retransmit_excd":                      ret.DtDdosDnsTcpPortStats.Stats.Src_retransmit_excd,
			"src_zero_window_excd":                     ret.DtDdosDnsTcpPortStats.Stats.Src_zero_window_excd,
			"src_conn_pkt_rate_excd":                   ret.DtDdosDnsTcpPortStats.Stats.Src_conn_pkt_rate_excd,
			"src_filter_action_blacklist":              ret.DtDdosDnsTcpPortStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":                   ret.DtDdosDnsTcpPortStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":           ret.DtDdosDnsTcpPortStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":              ret.DtDdosDnsTcpPortStats.Stats.Src_filter_action_whitelist,
			"force_tcp_auth_pass":                      ret.DtDdosDnsTcpPortStats.Stats.Force_tcp_auth_pass,
			"nxdomain_drop":                            ret.DtDdosDnsTcpPortStats.Stats.Nxdomain_drop,
			"dns_malform":                              ret.DtDdosDnsTcpPortStats.Stats.Dns_malform,
			"src_dns_malform":                          ret.DtDdosDnsTcpPortStats.Stats.Src_dns_malform,
			"tcp_rexmit_syn_limit_drop":                ret.DtDdosDnsTcpPortStats.Stats.Tcp_rexmit_syn_limit_drop,
			"tcp_rexmit_syn_limit_bl":                  ret.DtDdosDnsTcpPortStats.Stats.Tcp_rexmit_syn_limit_bl,
			"conn_ofo_rate_excd":                       ret.DtDdosDnsTcpPortStats.Stats.Conn_ofo_rate_excd,
			"conn_rexmit_rate_excd":                    ret.DtDdosDnsTcpPortStats.Stats.Conn_rexmit_rate_excd,
			"conn_zwindow_rate_excd":                   ret.DtDdosDnsTcpPortStats.Stats.Conn_zwindow_rate_excd,
			"src_conn_ofo_rate_excd":                   ret.DtDdosDnsTcpPortStats.Stats.Src_conn_ofo_rate_excd,
			"src_conn_rexmit_rate_excd":                ret.DtDdosDnsTcpPortStats.Stats.Src_conn_rexmit_rate_excd,
			"src_conn_zwindow_rate_excd":               ret.DtDdosDnsTcpPortStats.Stats.Src_conn_zwindow_rate_excd,
			"ack_retry_rto_pass":                       ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_rto_pass,
			"ack_retry_rto_fail":                       ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_rto_fail,
			"ack_retry_rto_progress":                   ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_rto_progress,
			"syn_retry_rto_pass":                       ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_rto_pass,
			"syn_retry_rto_fail":                       ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_rto_fail,
			"syn_retry_rto_progress":                   ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_rto_progress,
			"src_syn_retry_rto_pass":                   ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_rto_pass,
			"src_syn_retry_rto_fail":                   ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_rto_fail,
			"src_syn_retry_rto_progress":               ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_rto_progress,
			"src_ack_retry_rto_pass":                   ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_rto_pass,
			"src_ack_retry_rto_fail":                   ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_rto_fail,
			"src_ack_retry_rto_progress":               ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_rto_progress,
			"wellknown_sport_drop":                     ret.DtDdosDnsTcpPortStats.Stats.Wellknown_sport_drop,
			"src_well_known_port":                      ret.DtDdosDnsTcpPortStats.Stats.Src_well_known_port,
			"dg_request_check_fail":                    ret.DtDdosDnsTcpPortStats.Stats.Dg_request_check_fail,
			"query_type_any_drop":                      ret.DtDdosDnsTcpPortStats.Stats.Query_type_any_drop,
			"src_query_type_any_drop":                  ret.DtDdosDnsTcpPortStats.Stats.Src_query_type_any_drop,
			"src_auth_drop":                            ret.DtDdosDnsTcpPortStats.Stats.Src_auth_drop,
			"src_frag_drop":                            ret.DtDdosDnsTcpPortStats.Stats.Src_frag_drop,
			"frag_timeout":                             ret.DtDdosDnsTcpPortStats.Stats.Frag_timeout,
			"create_conn_non_syn_dropped":              ret.DtDdosDnsTcpPortStats.Stats.Create_conn_non_syn_dropped,
			"src_create_conn_non_syn_dropped":          ret.DtDdosDnsTcpPortStats.Stats.Src_create_conn_non_syn_dropped,
			"dg_rate_exceed":                           ret.DtDdosDnsTcpPortStats.Stats.Dg_rate_exceed,
			"port_syn_rate_exceed":                     ret.DtDdosDnsTcpPortStats.Stats.Port_syn_rate_exceed,
			"pattern_recognition_proceeded":            ret.DtDdosDnsTcpPortStats.Stats.Pattern_recognition_proceeded,
			"pattern_not_found":                        ret.DtDdosDnsTcpPortStats.Stats.Pattern_not_found,
			"pattern_recognition_generic_error":        ret.DtDdosDnsTcpPortStats.Stats.Pattern_recognition_generic_error,
			"pattern_filter1_match":                    ret.DtDdosDnsTcpPortStats.Stats.Pattern_filter1_match,
			"pattern_filter2_match":                    ret.DtDdosDnsTcpPortStats.Stats.Pattern_filter2_match,
			"pattern_filter3_match":                    ret.DtDdosDnsTcpPortStats.Stats.Pattern_filter3_match,
			"pattern_filter4_match":                    ret.DtDdosDnsTcpPortStats.Stats.Pattern_filter4_match,
			"pattern_filter5_match":                    ret.DtDdosDnsTcpPortStats.Stats.Pattern_filter5_match,
			"pattern_filter_drop":                      ret.DtDdosDnsTcpPortStats.Stats.Pattern_filter_drop,
			"src_syn_rate_exceed":                      ret.DtDdosDnsTcpPortStats.Stats.Src_syn_rate_exceed,
			"non_query_opcode_pass_through":            ret.DtDdosDnsTcpPortStats.Stats.Non_query_opcode_pass_through,
			"src_filter1_match":                        ret.DtDdosDnsTcpPortStats.Stats.Src_filter1_match,
			"src_filter2_match":                        ret.DtDdosDnsTcpPortStats.Stats.Src_filter2_match,
			"src_filter3_match":                        ret.DtDdosDnsTcpPortStats.Stats.Src_filter3_match,
			"src_filter4_match":                        ret.DtDdosDnsTcpPortStats.Stats.Src_filter4_match,
			"src_filter5_match":                        ret.DtDdosDnsTcpPortStats.Stats.Src_filter5_match,
			"src_filter_none_match":                    ret.DtDdosDnsTcpPortStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":               ret.DtDdosDnsTcpPortStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":                     ret.DtDdosDnsTcpPortStats.Stats.Src_filter_auth_fail,
			"syn_tfo_rcv":                              ret.DtDdosDnsTcpPortStats.Stats.Syn_tfo_rcv,
			"ack_retry_timeout":                        ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_timeout,
			"ack_retry_reset":                          ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_reset,
			"ack_retry_blacklist":                      ret.DtDdosDnsTcpPortStats.Stats.Ack_retry_blacklist,
			"src_ack_retry_timeout":                    ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_timeout,
			"src_ack_retry_reset":                      ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_reset,
			"src_ack_retry_blacklist":                  ret.DtDdosDnsTcpPortStats.Stats.Src_ack_retry_blacklist,
			"syn_retry_timeout":                        ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_timeout,
			"syn_retry_reset":                          ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_reset,
			"syn_retry_blacklist":                      ret.DtDdosDnsTcpPortStats.Stats.Syn_retry_blacklist,
			"src_syn_retry_timeout":                    ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_timeout,
			"src_syn_retry_reset":                      ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_reset,
			"src_syn_retry_blacklist":                  ret.DtDdosDnsTcpPortStats.Stats.Src_syn_retry_blacklist,
			"sflow_internal_samples_packed":            ret.DtDdosDnsTcpPortStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":            ret.DtDdosDnsTcpPortStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":              ret.DtDdosDnsTcpPortStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":              ret.DtDdosDnsTcpPortStats.Stats.Sflow_external_packets_sent,
			"pattern_recognition_sampling_started":     ret.DtDdosDnsTcpPortStats.Stats.Pattern_recognition_sampling_started,
			"pattern_recognition_pattern_changed":      ret.DtDdosDnsTcpPortStats.Stats.Pattern_recognition_pattern_changed,
			"exceed_action_tunnel":                     ret.DtDdosDnsTcpPortStats.Stats.Exceed_action_tunnel,
			"dst_hw_drop":                              ret.DtDdosDnsTcpPortStats.Stats.Dst_hw_drop,
			"synack_reset_sent":                        ret.DtDdosDnsTcpPortStats.Stats.Synack_reset_sent,
			"synack_multiple_attempts_per_ip_detected": ret.DtDdosDnsTcpPortStats.Stats.Synack_multiple_attempts_per_ip_detected,
			"prog_first_req_time_exceed":               ret.DtDdosDnsTcpPortStats.Stats.Prog_first_req_time_exceed,
			"prog_req_resp_time_exceed":                ret.DtDdosDnsTcpPortStats.Stats.Prog_req_resp_time_exceed,
			"prog_request_len_exceed":                  ret.DtDdosDnsTcpPortStats.Stats.Prog_request_len_exceed,
			"prog_response_len_exceed":                 ret.DtDdosDnsTcpPortStats.Stats.Prog_response_len_exceed,
			"prog_resp_req_ratio_exceed":               ret.DtDdosDnsTcpPortStats.Stats.Prog_resp_req_ratio_exceed,
			"prog_resp_req_time_exceed":                ret.DtDdosDnsTcpPortStats.Stats.Prog_resp_req_time_exceed,
			"prog_conn_sent_exceed":                    ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_sent_exceed,
			"prog_conn_rcvd_exceed":                    ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_rcvd_exceed,
			"prog_conn_time_exceed":                    ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_time_exceed,
			"prog_conn_rcvd_sent_ratio_exceed":         ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_rcvd_sent_ratio_exceed,
			"prog_win_sent_exceed":                     ret.DtDdosDnsTcpPortStats.Stats.Prog_win_sent_exceed,
			"prog_win_rcvd_exceed":                     ret.DtDdosDnsTcpPortStats.Stats.Prog_win_rcvd_exceed,
			"prog_win_rcvd_sent_ratio_exceed":          ret.DtDdosDnsTcpPortStats.Stats.Prog_win_rcvd_sent_ratio_exceed,
			"snat_fail":                                ret.DtDdosDnsTcpPortStats.Stats.Snat_fail,
			"prog_exceed_drop":                         ret.DtDdosDnsTcpPortStats.Stats.Prog_exceed_drop,
			"prog_exceed_bl":                           ret.DtDdosDnsTcpPortStats.Stats.Prog_exceed_bl,
			"prog_conn_exceed_drop":                    ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_exceed_drop,
			"prog_conn_exceed_bl":                      ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_exceed_bl,
			"prog_win_exceed_drop":                     ret.DtDdosDnsTcpPortStats.Stats.Prog_win_exceed_drop,
			"prog_win_exceed_bl":                       ret.DtDdosDnsTcpPortStats.Stats.Prog_win_exceed_bl,
			"exceed_action_drop":                       ret.DtDdosDnsTcpPortStats.Stats.Exceed_action_drop,
			"syn_auth_rst_ack_drop":                    ret.DtDdosDnsTcpPortStats.Stats.Syn_auth_rst_ack_drop,
			"prog_exceed_reset":                        ret.DtDdosDnsTcpPortStats.Stats.Prog_exceed_reset,
			"prog_conn_exceed_reset":                   ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_exceed_reset,
			"prog_win_exceed_reset":                    ret.DtDdosDnsTcpPortStats.Stats.Prog_win_exceed_reset,
			"conn_create_from_synack":                  ret.DtDdosDnsTcpPortStats.Stats.Conn_create_from_synack,
			"port_synack_rate_exceed":                  ret.DtDdosDnsTcpPortStats.Stats.Port_synack_rate_exceed,
			"prog_conn_samples":                        ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_samples,
			"prog_req_samples":                         ret.DtDdosDnsTcpPortStats.Stats.Prog_req_samples,
			"prog_win_samples":                         ret.DtDdosDnsTcpPortStats.Stats.Prog_win_samples,
			"prog_conn_samples_processed":              ret.DtDdosDnsTcpPortStats.Stats.Prog_conn_samples_processed,
			"prog_req_samples_processed":               ret.DtDdosDnsTcpPortStats.Stats.Prog_req_samples_processed,
			"prog_win_samples_processed":               ret.DtDdosDnsTcpPortStats.Stats.Prog_win_samples_processed,
			"tcp_auth_rst":                             ret.DtDdosDnsTcpPortStats.Stats.Tcp_auth_rst,
			"src_tcp_auth_rst":                         ret.DtDdosDnsTcpPortStats.Stats.Src_tcp_auth_rst,
		},
	}
}

func getObjectDdosDnsTcpPortStatsStats(d []interface{}) edpt.DdosDnsTcpPortStatsStats {

	count1 := len(d)
	var ret edpt.DdosDnsTcpPortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rate_limit_type0 = in["rate_limit_type0"].(int)
		ret.Rate_limit_type1 = in["rate_limit_type1"].(int)
		ret.Rate_limit_type2 = in["rate_limit_type2"].(int)
		ret.Rate_limit_type3 = in["rate_limit_type3"].(int)
		ret.Rate_limit_type4 = in["rate_limit_type4"].(int)
		ret.Nxdomain_rate_exceed = in["nxdomain_rate_exceed"].(int)
		ret.Is_nxdomain = in["is_nxdomain"].(int)
		ret.Nxdomain_bl_drop = in["nxdomain_bl_drop"].(int)
		ret.Dns_fqdn_label_len_exceed = in["dns_fqdn_label_len_exceed"].(int)
		ret.Dns_malform_drop = in["dns_malform_drop"].(int)
		ret.Dns_fqdn_stage_2_rate_exceed = in["dns_fqdn_stage_2_rate_exceed"].(int)
		ret.Req_sent = in["req_sent"].(int)
		ret.Req_incomplete = in["req_incomplete"].(int)
		ret.Req_size_exceed = in["req_size_exceed"].(int)
		ret.Req_retrans = in["req_retrans"].(int)
		ret.Query_type_a = in["query_type_a"].(int)
		ret.Query_type_aaaa = in["query_type_aaaa"].(int)
		ret.Query_type_ns = in["query_type_ns"].(int)
		ret.Query_type_cname = in["query_type_cname"].(int)
		ret.Query_type_any = in["query_type_any"].(int)
		ret.Query_type_srv = in["query_type_srv"].(int)
		ret.Query_type_mx = in["query_type_mx"].(int)
		ret.Query_type_soa = in["query_type_soa"].(int)
		ret.Query_type_opt = in["query_type_opt"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
		ret.Port_conn_rate_exceed = in["port_conn_rate_exceed"].(int)
		ret.Port_conn_limm_exceed = in["port_conn_limm_exceed"].(int)
		ret.Dg_action_permit = in["dg_action_permit"].(int)
		ret.Dg_action_deny = in["dg_action_deny"].(int)
		ret.Dg_action_default = in["dg_action_default"].(int)
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
		ret.Src_rate_limit_type0 = in["src_rate_limit_type0"].(int)
		ret.Src_rate_limit_type1 = in["src_rate_limit_type1"].(int)
		ret.Src_rate_limit_type2 = in["src_rate_limit_type2"].(int)
		ret.Src_rate_limit_type3 = in["src_rate_limit_type3"].(int)
		ret.Src_rate_limit_type4 = in["src_rate_limit_type4"].(int)
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
		ret.Fqdn_rate_by_label_count_check = in["fqdn_rate_by_label_count_check"].(int)
		ret.Fqdn_rate_by_label_count_exceed = in["fqdn_rate_by_label_count_exceed"].(int)
		ret.Policy_drop = in["policy_drop"].(int)
		ret.Fqdn_label_count_exceed = in["fqdn_label_count_exceed"].(int)
		ret.Rrtype_drop = in["rrtype_drop"].(int)
		ret.Src_policy_drop = in["src_policy_drop"].(int)
		ret.Src_dns_fqdn_label_len_exceed = in["src_dns_fqdn_label_len_exceed"].(int)
		ret.Src_fqdn_label_count_exceed = in["src_fqdn_label_count_exceed"].(int)
		ret.Src_dns_malform_drop = in["src_dns_malform_drop"].(int)
		ret.Tcp_auth_drop = in["tcp_auth_drop"].(int)
		ret.Tcp_auth_resp = in["tcp_auth_resp"].(int)
		ret.Dns_query_class_in = in["dns_query_class_in"].(int)
		ret.Dns_query_class_csnet = in["dns_query_class_csnet"].(int)
		ret.Dns_query_class_chaos = in["dns_query_class_chaos"].(int)
		ret.Dns_query_class_hs = in["dns_query_class_hs"].(int)
		ret.Dns_query_class_none = in["dns_query_class_none"].(int)
		ret.Dns_query_class_any = in["dns_query_class_any"].(int)
		ret.Dns_query_class_whitelist_miss = in["dns_query_class_whitelist_miss"].(int)
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
		ret.Filter1_match = in["filter1_match"].(int)
		ret.Filter2_match = in["filter2_match"].(int)
		ret.Filter3_match = in["filter3_match"].(int)
		ret.Filter4_match = in["filter4_match"].(int)
		ret.Filter5_match = in["filter5_match"].(int)
		ret.Filter_none_match = in["filter_none_match"].(int)
		ret.Filter_total_not_match = in["filter_total_not_match"].(int)
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
		ret.Force_tcp_auth_pass = in["force_tcp_auth_pass"].(int)
		ret.Nxdomain_drop = in["nxdomain_drop"].(int)
		ret.Dns_malform = in["dns_malform"].(int)
		ret.Src_dns_malform = in["src_dns_malform"].(int)
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
		ret.Dg_request_check_fail = in["dg_request_check_fail"].(int)
		ret.Query_type_any_drop = in["query_type_any_drop"].(int)
		ret.Src_query_type_any_drop = in["src_query_type_any_drop"].(int)
		ret.Src_auth_drop = in["src_auth_drop"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Create_conn_non_syn_dropped = in["create_conn_non_syn_dropped"].(int)
		ret.Src_create_conn_non_syn_dropped = in["src_create_conn_non_syn_dropped"].(int)
		ret.Dg_rate_exceed = in["dg_rate_exceed"].(int)
		ret.Port_syn_rate_exceed = in["port_syn_rate_exceed"].(int)
		ret.Pattern_recognition_proceeded = in["pattern_recognition_proceeded"].(int)
		ret.Pattern_not_found = in["pattern_not_found"].(int)
		ret.Pattern_recognition_generic_error = in["pattern_recognition_generic_error"].(int)
		ret.Pattern_filter1_match = in["pattern_filter1_match"].(int)
		ret.Pattern_filter2_match = in["pattern_filter2_match"].(int)
		ret.Pattern_filter3_match = in["pattern_filter3_match"].(int)
		ret.Pattern_filter4_match = in["pattern_filter4_match"].(int)
		ret.Pattern_filter5_match = in["pattern_filter5_match"].(int)
		ret.Pattern_filter_drop = in["pattern_filter_drop"].(int)
		ret.Src_syn_rate_exceed = in["src_syn_rate_exceed"].(int)
		ret.Non_query_opcode_pass_through = in["non_query_opcode_pass_through"].(int)
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
		ret.Pattern_recognition_sampling_started = in["pattern_recognition_sampling_started"].(int)
		ret.Pattern_recognition_pattern_changed = in["pattern_recognition_pattern_changed"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
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

func dataToEndpointDdosDnsTcpPortStats(d *schema.ResourceData) edpt.DdosDnsTcpPortStats {
	var ret edpt.DdosDnsTcpPortStats

	ret.Stats = getObjectDdosDnsTcpPortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
