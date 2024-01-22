package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSipUdpZonePortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_sip_udp_zone_port_stats`: Statistics for the object sip-udp-zone-port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosSipUdpZonePortStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"udp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Init",
						},
						"udp_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Failed",
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
						"spoof_detect_fail": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Timeout",
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
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Whitelist",
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
						"udp_auth_retry_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Retry Failed",
						},
						"exceed_drop_brate_src_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded Count",
						},
						"port_kbit_rate_exceed_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded Count",
						},
						"udp_retry_init": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Init",
						},
						"conn_prate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Conn Pkt Rate Exceeded",
						},
						"ntp_monlist_req": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Request",
						},
						"ntp_monlist_resp": {
							Type: schema.TypeInt, Optional: true, Description: "NTP Monlist Response",
						},
						"wellknown_sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP SrcPort Wellknown",
						},
						"payload_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small",
						},
						"payload_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large",
						},
						"udp_auth_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Retry-Gap Dropped",
						},
						"src_udp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "Src Auth UDP Init",
						},
						"udp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Passed",
						},
						"udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth Dropped",
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
						"sess_aged": {
							Type: schema.TypeInt, Optional: true, Description: "Sessions Aged Out",
						},
						"udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Passed",
						},
						"udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry-Gap Dropped",
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
						"src_payload_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Payload Too Small",
						},
						"src_payload_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Payload Too Large",
						},
						"src_ntp_monlist_req": {
							Type: schema.TypeInt, Optional: true, Description: "Src NTP Monlist Request",
						},
						"src_ntp_monlist_resp": {
							Type: schema.TypeInt, Optional: true, Description: "Src NTP Monlist Response",
						},
						"src_well_known_port": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP SrcPort Wellknown",
						},
						"src_conn_pkt_rate_excd": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Conn Pkt Rate Exceeded",
						},
						"src_udp_retry_init": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Init",
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
							Type: schema.TypeInt, Optional: true, Description: "Src Filter Action Whitelist",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
						},
						"frag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
						},
						"current_es_level": {
							Type: schema.TypeInt, Optional: true, Description: "Current Escalation Level",
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
						"no_policy_class_list_match": {
							Type: schema.TypeInt, Optional: true, Description: "No Policy Class-list Match",
						},
						"port_src_escalation": {
							Type: schema.TypeInt, Optional: true, Description: "Src Escalation",
						},
						"src_udp_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src Auth UDP Failed",
						},
						"policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Drop",
						},
						"policy_violation": {
							Type: schema.TypeInt, Optional: true, Description: "Policy Violation",
						},
						"src_udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth Dropped",
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
						"filter_total_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Not Matched on Pkt",
						},
						"src_udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry-Gap Dropped",
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
						"dst_udp_retry_timeout_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Timeout Blacklisted",
						},
						"src_udp_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Timeout",
						},
						"zone_src_udp_retry_timeout_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Timeout Blacklisted",
						},
						"src_udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Passed",
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
						"token_authentication_mismatched": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Mismatched Packets",
						},
						"token_authentication_invalid": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Invalid Packets",
						},
						"token_authentication_curr_salt_matched": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Current Salt Matched",
						},
						"token_authentication_prev_salt_matched": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Previous Salt Matched",
						},
						"token_authentication_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Session Created",
						},
						"token_authentication_session_created_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Token Authentication Session Created Fail",
						},
						"snat_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Source NAT Failure",
						},
						"exceed_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Exceed Action: Dropped",
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
						"src_hw_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Hardware Packets Dropped",
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
	}
}

func resourceDdosSipUdpZonePortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSipUdpZonePortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSipUdpZonePortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosSipUdpZonePortStatsStats := setObjectDdosSipUdpZonePortStatsStats(res)
		d.Set("stats", DdosSipUdpZonePortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosSipUdpZonePortStatsStats(ret edpt.DataDdosSipUdpZonePortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"request_method_ack":                        ret.DtDdosSipUdpZonePortStats.Stats.Request_method_ack,
			"request_method_bye":                        ret.DtDdosSipUdpZonePortStats.Stats.Request_method_bye,
			"request_method_cancel":                     ret.DtDdosSipUdpZonePortStats.Stats.Request_method_cancel,
			"request_method_invite":                     ret.DtDdosSipUdpZonePortStats.Stats.Request_method_invite,
			"request_method_info":                       ret.DtDdosSipUdpZonePortStats.Stats.Request_method_info,
			"request_method_message":                    ret.DtDdosSipUdpZonePortStats.Stats.Request_method_message,
			"request_method_notify":                     ret.DtDdosSipUdpZonePortStats.Stats.Request_method_notify,
			"request_method_options":                    ret.DtDdosSipUdpZonePortStats.Stats.Request_method_options,
			"request_method_prack":                      ret.DtDdosSipUdpZonePortStats.Stats.Request_method_prack,
			"request_method_publish":                    ret.DtDdosSipUdpZonePortStats.Stats.Request_method_publish,
			"request_method_register":                   ret.DtDdosSipUdpZonePortStats.Stats.Request_method_register,
			"request_method_refer":                      ret.DtDdosSipUdpZonePortStats.Stats.Request_method_refer,
			"request_method_subscribe":                  ret.DtDdosSipUdpZonePortStats.Stats.Request_method_subscribe,
			"request_method_update":                     ret.DtDdosSipUdpZonePortStats.Stats.Request_method_update,
			"request_method_unknown":                    ret.DtDdosSipUdpZonePortStats.Stats.Request_method_unknown,
			"request_unknown_version":                   ret.DtDdosSipUdpZonePortStats.Stats.Request_unknown_version,
			"keep_alive_msg":                            ret.DtDdosSipUdpZonePortStats.Stats.Keep_alive_msg,
			"rate1_limit_exceed":                        ret.DtDdosSipUdpZonePortStats.Stats.Rate1_limit_exceed,
			"rate2_limit_exceed":                        ret.DtDdosSipUdpZonePortStats.Stats.Rate2_limit_exceed,
			"rate3_limit_exceed":                        ret.DtDdosSipUdpZonePortStats.Stats.Rate3_limit_exceed,
			"src_rate1_limit_exceed":                    ret.DtDdosSipUdpZonePortStats.Stats.Src_rate1_limit_exceed,
			"src_rate2_limit_exceed":                    ret.DtDdosSipUdpZonePortStats.Stats.Src_rate2_limit_exceed,
			"src_rate3_limit_exceed":                    ret.DtDdosSipUdpZonePortStats.Stats.Src_rate3_limit_exceed,
			"response_1xx":                              ret.DtDdosSipUdpZonePortStats.Stats.Response_1xx,
			"response_2xx":                              ret.DtDdosSipUdpZonePortStats.Stats.Response_2xx,
			"response_3xx":                              ret.DtDdosSipUdpZonePortStats.Stats.Response_3xx,
			"response_4xx":                              ret.DtDdosSipUdpZonePortStats.Stats.Response_4xx,
			"response_5xx":                              ret.DtDdosSipUdpZonePortStats.Stats.Response_5xx,
			"response_6xx":                              ret.DtDdosSipUdpZonePortStats.Stats.Response_6xx,
			"response_unknown":                          ret.DtDdosSipUdpZonePortStats.Stats.Response_unknown,
			"response_unknown_version":                  ret.DtDdosSipUdpZonePortStats.Stats.Response_unknown_version,
			"read_start_line_error":                     ret.DtDdosSipUdpZonePortStats.Stats.Read_start_line_error,
			"invalid_start_line_error":                  ret.DtDdosSipUdpZonePortStats.Stats.Invalid_start_line_error,
			"parse_start_line_error":                    ret.DtDdosSipUdpZonePortStats.Stats.Parse_start_line_error,
			"line_too_long":                             ret.DtDdosSipUdpZonePortStats.Stats.Line_too_long,
			"line_mem_allocated":                        ret.DtDdosSipUdpZonePortStats.Stats.Line_mem_allocated,
			"line_mem_freed":                            ret.DtDdosSipUdpZonePortStats.Stats.Line_mem_freed,
			"max_uri_len_exceed":                        ret.DtDdosSipUdpZonePortStats.Stats.Max_uri_len_exceed,
			"too_many_header":                           ret.DtDdosSipUdpZonePortStats.Stats.Too_many_header,
			"invalid_header":                            ret.DtDdosSipUdpZonePortStats.Stats.Invalid_header,
			"header_name_too_long":                      ret.DtDdosSipUdpZonePortStats.Stats.Header_name_too_long,
			"parse_header_fail_error":                   ret.DtDdosSipUdpZonePortStats.Stats.Parse_header_fail_error,
			"max_header_value_len_exceed":               ret.DtDdosSipUdpZonePortStats.Stats.Max_header_value_len_exceed,
			"max_call_id_len_exceed":                    ret.DtDdosSipUdpZonePortStats.Stats.Max_call_id_len_exceed,
			"header_filter_match":                       ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_match,
			"header_filter_not_match":                   ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_not_match,
			"header_filter_none_match":                  ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_none_match,
			"header_filter_action_drop":                 ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_action_drop,
			"header_filter_action_blacklist":            ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_action_blacklist,
			"header_filter_action_whitelist":            ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_action_whitelist,
			"header_filter_action_default_pass":         ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_action_default_pass,
			"header_filter_filter1_match":               ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_filter1_match,
			"header_filter_filter2_match":               ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_filter2_match,
			"header_filter_filter3_match":               ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_filter3_match,
			"header_filter_filter4_match":               ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_filter4_match,
			"header_filter_filter5_match":               ret.DtDdosSipUdpZonePortStats.Stats.Header_filter_filter5_match,
			"max_sdp_len_exceed":                        ret.DtDdosSipUdpZonePortStats.Stats.Max_sdp_len_exceed,
			"body_too_big":                              ret.DtDdosSipUdpZonePortStats.Stats.Body_too_big,
			"get_content_fail_error":                    ret.DtDdosSipUdpZonePortStats.Stats.Get_content_fail_error,
			"concatenate_msg":                           ret.DtDdosSipUdpZonePortStats.Stats.Concatenate_msg,
			"mem_alloc_fail_error":                      ret.DtDdosSipUdpZonePortStats.Stats.Mem_alloc_fail_error,
			"malform_request":                           ret.DtDdosSipUdpZonePortStats.Stats.Malform_request,
			"udp_auth":                                  ret.DtDdosSipUdpZonePortStats.Stats.Udp_auth,
			"udp_auth_fail":                             ret.DtDdosSipUdpZonePortStats.Stats.Udp_auth_fail,
			"port_rcvd":                                 ret.DtDdosSipUdpZonePortStats.Stats.Port_rcvd,
			"port_drop":                                 ret.DtDdosSipUdpZonePortStats.Stats.Port_drop,
			"port_pkt_sent":                             ret.DtDdosSipUdpZonePortStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":                      ret.DtDdosSipUdpZonePortStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":                     ret.DtDdosSipUdpZonePortStats.Stats.Port_kbit_rate_exceed,
			"port_conn_rate_exceed":                     ret.DtDdosSipUdpZonePortStats.Stats.Port_conn_rate_exceed,
			"port_conn_limm_exceed":                     ret.DtDdosSipUdpZonePortStats.Stats.Port_conn_limm_exceed,
			"port_bytes":                                ret.DtDdosSipUdpZonePortStats.Stats.Port_bytes,
			"outbound_port_bytes":                       ret.DtDdosSipUdpZonePortStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                        ret.DtDdosSipUdpZonePortStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":                    ret.DtDdosSipUdpZonePortStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                           ret.DtDdosSipUdpZonePortStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                           ret.DtDdosSipUdpZonePortStats.Stats.Port_bytes_drop,
			"port_src_bl":                               ret.DtDdosSipUdpZonePortStats.Stats.Port_src_bl,
			"filter_auth_fail":                          ret.DtDdosSipUdpZonePortStats.Stats.Filter_auth_fail,
			"spoof_detect_fail":                         ret.DtDdosSipUdpZonePortStats.Stats.Spoof_detect_fail,
			"sess_create":                               ret.DtDdosSipUdpZonePortStats.Stats.Sess_create,
			"filter_action_blacklist":                   ret.DtDdosSipUdpZonePortStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                        ret.DtDdosSipUdpZonePortStats.Stats.Filter_action_drop,
			"filter_action_default_pass":                ret.DtDdosSipUdpZonePortStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":                   ret.DtDdosSipUdpZonePortStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":                     ret.DtDdosSipUdpZonePortStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_crate_src":                     ret.DtDdosSipUdpZonePortStats.Stats.Exceed_drop_crate_src,
			"exceed_drop_climit_src":                    ret.DtDdosSipUdpZonePortStats.Stats.Exceed_drop_climit_src,
			"exceed_drop_brate_src":                     ret.DtDdosSipUdpZonePortStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":                  ret.DtDdosSipUdpZonePortStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                        ret.DtDdosSipUdpZonePortStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":                  ret.DtDdosSipUdpZonePortStats.Stats.Outbound_port_bytes_drop,
			"udp_auth_retry_fail":                       ret.DtDdosSipUdpZonePortStats.Stats.Udp_auth_retry_fail,
			"exceed_drop_brate_src_pkt":                 ret.DtDdosSipUdpZonePortStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":                 ret.DtDdosSipUdpZonePortStats.Stats.Port_kbit_rate_exceed_pkt,
			"udp_retry_init":                            ret.DtDdosSipUdpZonePortStats.Stats.Udp_retry_init,
			"conn_prate_excd":                           ret.DtDdosSipUdpZonePortStats.Stats.Conn_prate_excd,
			"ntp_monlist_req":                           ret.DtDdosSipUdpZonePortStats.Stats.Ntp_monlist_req,
			"ntp_monlist_resp":                          ret.DtDdosSipUdpZonePortStats.Stats.Ntp_monlist_resp,
			"wellknown_sport_drop":                      ret.DtDdosSipUdpZonePortStats.Stats.Wellknown_sport_drop,
			"payload_too_small":                         ret.DtDdosSipUdpZonePortStats.Stats.Payload_too_small,
			"payload_too_big":                           ret.DtDdosSipUdpZonePortStats.Stats.Payload_too_big,
			"udp_auth_retry_gap_drop":                   ret.DtDdosSipUdpZonePortStats.Stats.Udp_auth_retry_gap_drop,
			"src_udp_auth":                              ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_auth,
			"udp_auth_pass":                             ret.DtDdosSipUdpZonePortStats.Stats.Udp_auth_pass,
			"udp_auth_drop":                             ret.DtDdosSipUdpZonePortStats.Stats.Udp_auth_drop,
			"bl":                                        ret.DtDdosSipUdpZonePortStats.Stats.Bl,
			"src_drop":                                  ret.DtDdosSipUdpZonePortStats.Stats.Src_drop,
			"frag_rcvd":                                 ret.DtDdosSipUdpZonePortStats.Stats.Frag_rcvd,
			"frag_drop":                                 ret.DtDdosSipUdpZonePortStats.Stats.Frag_drop,
			"sess_create_inbound":                       ret.DtDdosSipUdpZonePortStats.Stats.Sess_create_inbound,
			"sess_create_outbound":                      ret.DtDdosSipUdpZonePortStats.Stats.Sess_create_outbound,
			"sess_aged":                                 ret.DtDdosSipUdpZonePortStats.Stats.Sess_aged,
			"udp_retry_pass":                            ret.DtDdosSipUdpZonePortStats.Stats.Udp_retry_pass,
			"udp_retry_gap_drop":                        ret.DtDdosSipUdpZonePortStats.Stats.Udp_retry_gap_drop,
			"filter1_match":                             ret.DtDdosSipUdpZonePortStats.Stats.Filter1_match,
			"filter2_match":                             ret.DtDdosSipUdpZonePortStats.Stats.Filter2_match,
			"filter3_match":                             ret.DtDdosSipUdpZonePortStats.Stats.Filter3_match,
			"filter4_match":                             ret.DtDdosSipUdpZonePortStats.Stats.Filter4_match,
			"filter5_match":                             ret.DtDdosSipUdpZonePortStats.Stats.Filter5_match,
			"filter_none_match":                         ret.DtDdosSipUdpZonePortStats.Stats.Filter_none_match,
			"src_payload_too_small":                     ret.DtDdosSipUdpZonePortStats.Stats.Src_payload_too_small,
			"src_payload_too_big":                       ret.DtDdosSipUdpZonePortStats.Stats.Src_payload_too_big,
			"src_ntp_monlist_req":                       ret.DtDdosSipUdpZonePortStats.Stats.Src_ntp_monlist_req,
			"src_ntp_monlist_resp":                      ret.DtDdosSipUdpZonePortStats.Stats.Src_ntp_monlist_resp,
			"src_well_known_port":                       ret.DtDdosSipUdpZonePortStats.Stats.Src_well_known_port,
			"src_conn_pkt_rate_excd":                    ret.DtDdosSipUdpZonePortStats.Stats.Src_conn_pkt_rate_excd,
			"src_udp_retry_init":                        ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_retry_init,
			"src_filter_action_blacklist":               ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":                    ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":            ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":               ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_action_whitelist,
			"src_frag_drop":                             ret.DtDdosSipUdpZonePortStats.Stats.Src_frag_drop,
			"frag_timeout":                              ret.DtDdosSipUdpZonePortStats.Stats.Frag_timeout,
			"current_es_level":                          ret.DtDdosSipUdpZonePortStats.Stats.Current_es_level,
			"secondary_port_pkt_rate_exceed":            ret.DtDdosSipUdpZonePortStats.Stats.Secondary_port_pkt_rate_exceed,
			"secondary_port_kbit_rate_exceed":           ret.DtDdosSipUdpZonePortStats.Stats.Secondary_port_kbit_rate_exceed,
			"secondary_port_kbit_rate_exceed_pkt":       ret.DtDdosSipUdpZonePortStats.Stats.Secondary_port_kbit_rate_exceed_pkt,
			"secondary_port_conn_rate_exceed":           ret.DtDdosSipUdpZonePortStats.Stats.Secondary_port_conn_rate_exceed,
			"secondary_port_conn_limm_exceed":           ret.DtDdosSipUdpZonePortStats.Stats.Secondary_port_conn_limm_exceed,
			"no_policy_class_list_match":                ret.DtDdosSipUdpZonePortStats.Stats.No_policy_class_list_match,
			"port_src_escalation":                       ret.DtDdosSipUdpZonePortStats.Stats.Port_src_escalation,
			"src_udp_auth_fail":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_auth_fail,
			"policy_drop":                               ret.DtDdosSipUdpZonePortStats.Stats.Policy_drop,
			"policy_violation":                          ret.DtDdosSipUdpZonePortStats.Stats.Policy_violation,
			"src_udp_auth_drop":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_auth_drop,
			"src_filter1_match":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_filter1_match,
			"src_filter2_match":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_filter2_match,
			"src_filter3_match":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_filter3_match,
			"src_filter4_match":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_filter4_match,
			"src_filter5_match":                         ret.DtDdosSipUdpZonePortStats.Stats.Src_filter5_match,
			"src_filter_none_match":                     ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":                ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":                      ret.DtDdosSipUdpZonePortStats.Stats.Src_filter_auth_fail,
			"filter_total_not_match":                    ret.DtDdosSipUdpZonePortStats.Stats.Filter_total_not_match,
			"src_udp_retry_gap_drop":                    ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_retry_gap_drop,
			"sflow_internal_samples_packed":             ret.DtDdosSipUdpZonePortStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":             ret.DtDdosSipUdpZonePortStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":               ret.DtDdosSipUdpZonePortStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":               ret.DtDdosSipUdpZonePortStats.Stats.Sflow_external_packets_sent,
			"exceed_action_tunnel":                      ret.DtDdosSipUdpZonePortStats.Stats.Exceed_action_tunnel,
			"dst_udp_retry_timeout_blacklist":           ret.DtDdosSipUdpZonePortStats.Stats.Dst_udp_retry_timeout_blacklist,
			"src_udp_auth_timeout":                      ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_auth_timeout,
			"zone_src_udp_retry_timeout_blacklist":      ret.DtDdosSipUdpZonePortStats.Stats.Zone_src_udp_retry_timeout_blacklist,
			"src_udp_retry_pass":                        ret.DtDdosSipUdpZonePortStats.Stats.Src_udp_retry_pass,
			"pattern_recognition_proceeded":             ret.DtDdosSipUdpZonePortStats.Stats.Pattern_recognition_proceeded,
			"pattern_not_found":                         ret.DtDdosSipUdpZonePortStats.Stats.Pattern_not_found,
			"pattern_recognition_generic_error":         ret.DtDdosSipUdpZonePortStats.Stats.Pattern_recognition_generic_error,
			"pattern_filter1_match":                     ret.DtDdosSipUdpZonePortStats.Stats.Pattern_filter1_match,
			"pattern_filter2_match":                     ret.DtDdosSipUdpZonePortStats.Stats.Pattern_filter2_match,
			"pattern_filter3_match":                     ret.DtDdosSipUdpZonePortStats.Stats.Pattern_filter3_match,
			"pattern_filter4_match":                     ret.DtDdosSipUdpZonePortStats.Stats.Pattern_filter4_match,
			"pattern_filter5_match":                     ret.DtDdosSipUdpZonePortStats.Stats.Pattern_filter5_match,
			"pattern_filter_drop":                       ret.DtDdosSipUdpZonePortStats.Stats.Pattern_filter_drop,
			"pattern_recognition_sampling_started":      ret.DtDdosSipUdpZonePortStats.Stats.Pattern_recognition_sampling_started,
			"pattern_recognition_pattern_changed":       ret.DtDdosSipUdpZonePortStats.Stats.Pattern_recognition_pattern_changed,
			"dst_hw_drop":                               ret.DtDdosSipUdpZonePortStats.Stats.Dst_hw_drop,
			"secondary_port_hit":                        ret.DtDdosSipUdpZonePortStats.Stats.Secondary_port_hit,
			"src_zone_service_entry_learned":            ret.DtDdosSipUdpZonePortStats.Stats.Src_zone_service_entry_learned,
			"src_zone_service_entry_aged":               ret.DtDdosSipUdpZonePortStats.Stats.Src_zone_service_entry_aged,
			"dst_hw_drop_inserted":                      ret.DtDdosSipUdpZonePortStats.Stats.Dst_hw_drop_inserted,
			"dst_hw_drop_removed":                       ret.DtDdosSipUdpZonePortStats.Stats.Dst_hw_drop_removed,
			"src_hw_drop_inserted":                      ret.DtDdosSipUdpZonePortStats.Stats.Src_hw_drop_inserted,
			"src_hw_drop_removed":                       ret.DtDdosSipUdpZonePortStats.Stats.Src_hw_drop_removed,
			"token_authentication_mismatched":           ret.DtDdosSipUdpZonePortStats.Stats.Token_authentication_mismatched,
			"token_authentication_invalid":              ret.DtDdosSipUdpZonePortStats.Stats.Token_authentication_invalid,
			"token_authentication_curr_salt_matched":    ret.DtDdosSipUdpZonePortStats.Stats.Token_authentication_curr_salt_matched,
			"token_authentication_prev_salt_matched":    ret.DtDdosSipUdpZonePortStats.Stats.Token_authentication_prev_salt_matched,
			"token_authentication_session_created":      ret.DtDdosSipUdpZonePortStats.Stats.Token_authentication_session_created,
			"token_authentication_session_created_fail": ret.DtDdosSipUdpZonePortStats.Stats.Token_authentication_session_created_fail,
			"snat_fail":                                 ret.DtDdosSipUdpZonePortStats.Stats.Snat_fail,
			"exceed_action_drop":                        ret.DtDdosSipUdpZonePortStats.Stats.Exceed_action_drop,
			"ew_inbound_port_rcv":                       ret.DtDdosSipUdpZonePortStats.Stats.Ew_inbound_port_rcv,
			"ew_inbound_port_drop":                      ret.DtDdosSipUdpZonePortStats.Stats.Ew_inbound_port_drop,
			"ew_inbound_port_sent":                      ret.DtDdosSipUdpZonePortStats.Stats.Ew_inbound_port_sent,
			"ew_inbound_port_byte_rcv":                  ret.DtDdosSipUdpZonePortStats.Stats.Ew_inbound_port_byte_rcv,
			"ew_inbound_port_byte_drop":                 ret.DtDdosSipUdpZonePortStats.Stats.Ew_inbound_port_byte_drop,
			"ew_inbound_port_byte_sent":                 ret.DtDdosSipUdpZonePortStats.Stats.Ew_inbound_port_byte_sent,
			"ew_outbound_port_rcv":                      ret.DtDdosSipUdpZonePortStats.Stats.Ew_outbound_port_rcv,
			"ew_outbound_port_drop":                     ret.DtDdosSipUdpZonePortStats.Stats.Ew_outbound_port_drop,
			"ew_outbound_port_sent":                     ret.DtDdosSipUdpZonePortStats.Stats.Ew_outbound_port_sent,
			"ew_outbound_port_byte_rcv":                 ret.DtDdosSipUdpZonePortStats.Stats.Ew_outbound_port_byte_rcv,
			"ew_outbound_port_byte_sent":                ret.DtDdosSipUdpZonePortStats.Stats.Ew_outbound_port_byte_sent,
			"ew_outbound_port_byte_drop":                ret.DtDdosSipUdpZonePortStats.Stats.Ew_outbound_port_byte_drop,
			"no_route_drop":                             ret.DtDdosSipUdpZonePortStats.Stats.No_route_drop,
			"unauth_src_session_reset":                  ret.DtDdosSipUdpZonePortStats.Stats.Unauth_src_session_reset,
			"src_hw_drop":                               ret.DtDdosSipUdpZonePortStats.Stats.Src_hw_drop,
			"addr_filter_drop":                          ret.DtDdosSipUdpZonePortStats.Stats.Addr_filter_drop,
			"addr_filter_bl":                            ret.DtDdosSipUdpZonePortStats.Stats.Addr_filter_bl,
			"src_learn_overflow":                        ret.DtDdosSipUdpZonePortStats.Stats.Src_learn_overflow,
		},
	}
}

func getObjectDdosSipUdpZonePortStatsStats(d []interface{}) edpt.DdosSipUdpZonePortStatsStats {

	count1 := len(d)
	var ret edpt.DdosSipUdpZonePortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.Udp_auth = in["udp_auth"].(int)
		ret.Udp_auth_fail = in["udp_auth_fail"].(int)
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
		ret.Spoof_detect_fail = in["spoof_detect_fail"].(int)
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
		ret.Udp_auth_retry_fail = in["udp_auth_retry_fail"].(int)
		ret.Exceed_drop_brate_src_pkt = in["exceed_drop_brate_src_pkt"].(int)
		ret.Port_kbit_rate_exceed_pkt = in["port_kbit_rate_exceed_pkt"].(int)
		ret.Udp_retry_init = in["udp_retry_init"].(int)
		ret.Conn_prate_excd = in["conn_prate_excd"].(int)
		ret.Ntp_monlist_req = in["ntp_monlist_req"].(int)
		ret.Ntp_monlist_resp = in["ntp_monlist_resp"].(int)
		ret.Wellknown_sport_drop = in["wellknown_sport_drop"].(int)
		ret.Payload_too_small = in["payload_too_small"].(int)
		ret.Payload_too_big = in["payload_too_big"].(int)
		ret.Udp_auth_retry_gap_drop = in["udp_auth_retry_gap_drop"].(int)
		ret.Src_udp_auth = in["src_udp_auth"].(int)
		ret.Udp_auth_pass = in["udp_auth_pass"].(int)
		ret.Udp_auth_drop = in["udp_auth_drop"].(int)
		ret.Bl = in["bl"].(int)
		ret.Src_drop = in["src_drop"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.Frag_drop = in["frag_drop"].(int)
		ret.Sess_create_inbound = in["sess_create_inbound"].(int)
		ret.Sess_create_outbound = in["sess_create_outbound"].(int)
		ret.Sess_aged = in["sess_aged"].(int)
		ret.Udp_retry_pass = in["udp_retry_pass"].(int)
		ret.Udp_retry_gap_drop = in["udp_retry_gap_drop"].(int)
		ret.Filter1_match = in["filter1_match"].(int)
		ret.Filter2_match = in["filter2_match"].(int)
		ret.Filter3_match = in["filter3_match"].(int)
		ret.Filter4_match = in["filter4_match"].(int)
		ret.Filter5_match = in["filter5_match"].(int)
		ret.Filter_none_match = in["filter_none_match"].(int)
		ret.Src_payload_too_small = in["src_payload_too_small"].(int)
		ret.Src_payload_too_big = in["src_payload_too_big"].(int)
		ret.Src_ntp_monlist_req = in["src_ntp_monlist_req"].(int)
		ret.Src_ntp_monlist_resp = in["src_ntp_monlist_resp"].(int)
		ret.Src_well_known_port = in["src_well_known_port"].(int)
		ret.Src_conn_pkt_rate_excd = in["src_conn_pkt_rate_excd"].(int)
		ret.Src_udp_retry_init = in["src_udp_retry_init"].(int)
		ret.Src_filter_action_blacklist = in["src_filter_action_blacklist"].(int)
		ret.Src_filter_action_drop = in["src_filter_action_drop"].(int)
		ret.Src_filter_action_default_pass = in["src_filter_action_default_pass"].(int)
		ret.Src_filter_action_whitelist = in["src_filter_action_whitelist"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Current_es_level = in["current_es_level"].(int)
		ret.Secondary_port_pkt_rate_exceed = in["secondary_port_pkt_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed = in["secondary_port_kbit_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed_pkt = in["secondary_port_kbit_rate_exceed_pkt"].(int)
		ret.Secondary_port_conn_rate_exceed = in["secondary_port_conn_rate_exceed"].(int)
		ret.Secondary_port_conn_limm_exceed = in["secondary_port_conn_limm_exceed"].(int)
		ret.No_policy_class_list_match = in["no_policy_class_list_match"].(int)
		ret.Port_src_escalation = in["port_src_escalation"].(int)
		ret.Src_udp_auth_fail = in["src_udp_auth_fail"].(int)
		ret.Policy_drop = in["policy_drop"].(int)
		ret.Policy_violation = in["policy_violation"].(int)
		ret.Src_udp_auth_drop = in["src_udp_auth_drop"].(int)
		ret.Src_filter1_match = in["src_filter1_match"].(int)
		ret.Src_filter2_match = in["src_filter2_match"].(int)
		ret.Src_filter3_match = in["src_filter3_match"].(int)
		ret.Src_filter4_match = in["src_filter4_match"].(int)
		ret.Src_filter5_match = in["src_filter5_match"].(int)
		ret.Src_filter_none_match = in["src_filter_none_match"].(int)
		ret.Src_filter_total_not_match = in["src_filter_total_not_match"].(int)
		ret.Src_filter_auth_fail = in["src_filter_auth_fail"].(int)
		ret.Filter_total_not_match = in["filter_total_not_match"].(int)
		ret.Src_udp_retry_gap_drop = in["src_udp_retry_gap_drop"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Dst_udp_retry_timeout_blacklist = in["dst_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_auth_timeout = in["src_udp_auth_timeout"].(int)
		ret.Zone_src_udp_retry_timeout_blacklist = in["zone_src_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_retry_pass = in["src_udp_retry_pass"].(int)
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
		ret.Secondary_port_hit = in["secondary_port_hit"].(int)
		ret.Src_zone_service_entry_learned = in["src_zone_service_entry_learned"].(int)
		ret.Src_zone_service_entry_aged = in["src_zone_service_entry_aged"].(int)
		ret.Dst_hw_drop_inserted = in["dst_hw_drop_inserted"].(int)
		ret.Dst_hw_drop_removed = in["dst_hw_drop_removed"].(int)
		ret.Src_hw_drop_inserted = in["src_hw_drop_inserted"].(int)
		ret.Src_hw_drop_removed = in["src_hw_drop_removed"].(int)
		ret.Token_authentication_mismatched = in["token_authentication_mismatched"].(int)
		ret.Token_authentication_invalid = in["token_authentication_invalid"].(int)
		ret.Token_authentication_curr_salt_matched = in["token_authentication_curr_salt_matched"].(int)
		ret.Token_authentication_prev_salt_matched = in["token_authentication_prev_salt_matched"].(int)
		ret.Token_authentication_session_created = in["token_authentication_session_created"].(int)
		ret.Token_authentication_session_created_fail = in["token_authentication_session_created_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Exceed_action_drop = in["exceed_action_drop"].(int)
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
		ret.Src_hw_drop = in["src_hw_drop"].(int)
		ret.Addr_filter_drop = in["addr_filter_drop"].(int)
		ret.Addr_filter_bl = in["addr_filter_bl"].(int)
		ret.Src_learn_overflow = in["src_learn_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosSipUdpZonePortStats(d *schema.ResourceData) edpt.DdosSipUdpZonePortStats {
	var ret edpt.DdosSipUdpZonePortStats

	ret.Stats = getObjectDdosSipUdpZonePortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
