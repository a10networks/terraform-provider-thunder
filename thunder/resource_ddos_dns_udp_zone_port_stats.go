package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDnsUdpZonePortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_dns_udp_zone_port_stats`: Statistics for the object dns-udp-zone-port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosDnsUdpZonePortStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"force_tcp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Force-TCP",
						},
						"udp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Passed",
						},
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
						"nxdomain_bl_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Blacklisted",
						},
						"is_nxdomain": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response",
						},
						"nxdomain_drop": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Dropped",
						},
						"udp_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Failed",
						},
						"malform_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Query Dropped",
						},
						"fqdn_stage_2_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate Exceeded",
						},
						"req_sent": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Requests Forwarded",
						},
						"req_size_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Request Size Exceeded",
						},
						"req_retrans": {
							Type: schema.TypeInt, Optional: true, Description: "Request Retransmit",
						},
						"fqdn_label_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Label Length Exceeded",
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
						"port_src_escalation": {
							Type: schema.TypeInt, Optional: true, Description: "Src Escalation",
						},
						"current_es_level": {
							Type: schema.TypeInt, Optional: true, Description: "Current Escalation Level",
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
						"fqdn_rate_by_label_count_check": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate by Label Count Checked",
						},
						"fqdn_rate_by_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Rate by Label Count Exceeded",
						},
						"udp_auth_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Retry-Gap Dropped",
						},
						"fqdn_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "FQDN Label Count Exceeded",
						},
						"rrtype_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Record Type Dropped",
						},
						"src_dns_fqdn_label_len_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src FQDN Label Length Exceeded",
						},
						"src_fqdn_label_count_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Src FQDN Label Count Exceeded",
						},
						"src_udp_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Src Auth UDP Failed",
						},
						"src_force_tcp_auth": {
							Type: schema.TypeInt, Optional: true, Description: "Src Auth Force-TCP",
						},
						"src_dns_malform_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Malform Query Dropped",
						},
						"force_tcp_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Auth Force-TCP With UDP Timeout",
						},
						"udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Auth Dropped",
						},
						"dns_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Dropped",
						},
						"dns_auth_resp": {
							Type: schema.TypeInt, Optional: true, Description: "DNS Auth Responded",
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
						"udp_auth_init": {
							Type: schema.TypeInt, Optional: true, Description: "Auth UDP Init",
						},
						"nxdomain_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "NXDOMAIN Response Rate Exceeded",
						},
						"src_udp_auth_init": {
							Type: schema.TypeInt, Optional: true, Description: "Src Auth UDP Init",
						},
						"dns_malform": {
							Type: schema.TypeInt, Optional: true, Description: "Malform Query",
						},
						"src_dns_malform": {
							Type: schema.TypeInt, Optional: true, Description: "Src Malform Query",
						},
						"dg_request_check_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Request Check Fail",
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
						"query_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Query Type ANY Dropped",
						},
						"src_query_type_any_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Query Type ANY Dropped",
						},
						"src_udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth Dropped",
						},
						"src_dns_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src DNS Auth Dropped",
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
						"dg_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Domain Group Domain Query Rate Exceeded",
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
						"non_query_opcode_pass_through": {
							Type: schema.TypeInt, Optional: true, Description: "Non Query Opcode Pass Through",
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

func resourceDdosDnsUdpZonePortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDnsUdpZonePortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDnsUdpZonePortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosDnsUdpZonePortStatsStats := setObjectDdosDnsUdpZonePortStatsStats(res)
		d.Set("stats", DdosDnsUdpZonePortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosDnsUdpZonePortStatsStats(ret edpt.DataDdosDnsUdpZonePortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"force_tcp_auth":                            ret.DtDdosDnsUdpZonePortStats.Stats.Force_tcp_auth,
			"udp_auth_pass":                             ret.DtDdosDnsUdpZonePortStats.Stats.Udp_auth_pass,
			"rate_limit_type0":                          ret.DtDdosDnsUdpZonePortStats.Stats.Rate_limit_type0,
			"rate_limit_type1":                          ret.DtDdosDnsUdpZonePortStats.Stats.Rate_limit_type1,
			"rate_limit_type2":                          ret.DtDdosDnsUdpZonePortStats.Stats.Rate_limit_type2,
			"rate_limit_type3":                          ret.DtDdosDnsUdpZonePortStats.Stats.Rate_limit_type3,
			"rate_limit_type4":                          ret.DtDdosDnsUdpZonePortStats.Stats.Rate_limit_type4,
			"nxdomain_bl_drop":                          ret.DtDdosDnsUdpZonePortStats.Stats.Nxdomain_bl_drop,
			"is_nxdomain":                               ret.DtDdosDnsUdpZonePortStats.Stats.Is_nxdomain,
			"nxdomain_drop":                             ret.DtDdosDnsUdpZonePortStats.Stats.Nxdomain_drop,
			"udp_auth_fail":                             ret.DtDdosDnsUdpZonePortStats.Stats.Udp_auth_fail,
			"malform_drop":                              ret.DtDdosDnsUdpZonePortStats.Stats.Malform_drop,
			"fqdn_stage_2_rate_exceed":                  ret.DtDdosDnsUdpZonePortStats.Stats.Fqdn_stage_2_rate_exceed,
			"req_sent":                                  ret.DtDdosDnsUdpZonePortStats.Stats.Req_sent,
			"req_size_exceed":                           ret.DtDdosDnsUdpZonePortStats.Stats.Req_size_exceed,
			"req_retrans":                               ret.DtDdosDnsUdpZonePortStats.Stats.Req_retrans,
			"fqdn_label_len_exceed":                     ret.DtDdosDnsUdpZonePortStats.Stats.Fqdn_label_len_exceed,
			"query_type_a":                              ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_a,
			"query_type_aaaa":                           ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_aaaa,
			"query_type_ns":                             ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_ns,
			"query_type_cname":                          ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_cname,
			"query_type_any":                            ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_any,
			"query_type_srv":                            ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_srv,
			"query_type_mx":                             ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_mx,
			"query_type_soa":                            ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_soa,
			"query_type_opt":                            ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_opt,
			"port_rcvd":                                 ret.DtDdosDnsUdpZonePortStats.Stats.Port_rcvd,
			"port_drop":                                 ret.DtDdosDnsUdpZonePortStats.Stats.Port_drop,
			"port_pkt_sent":                             ret.DtDdosDnsUdpZonePortStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":                      ret.DtDdosDnsUdpZonePortStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":                     ret.DtDdosDnsUdpZonePortStats.Stats.Port_kbit_rate_exceed,
			"port_conn_rate_exceed":                     ret.DtDdosDnsUdpZonePortStats.Stats.Port_conn_rate_exceed,
			"port_conn_limm_exceed":                     ret.DtDdosDnsUdpZonePortStats.Stats.Port_conn_limm_exceed,
			"dg_action_permit":                          ret.DtDdosDnsUdpZonePortStats.Stats.Dg_action_permit,
			"dg_action_deny":                            ret.DtDdosDnsUdpZonePortStats.Stats.Dg_action_deny,
			"dg_action_default":                         ret.DtDdosDnsUdpZonePortStats.Stats.Dg_action_default,
			"port_bytes":                                ret.DtDdosDnsUdpZonePortStats.Stats.Port_bytes,
			"outbound_port_bytes":                       ret.DtDdosDnsUdpZonePortStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                        ret.DtDdosDnsUdpZonePortStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":                    ret.DtDdosDnsUdpZonePortStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                           ret.DtDdosDnsUdpZonePortStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                           ret.DtDdosDnsUdpZonePortStats.Stats.Port_bytes_drop,
			"port_src_bl":                               ret.DtDdosDnsUdpZonePortStats.Stats.Port_src_bl,
			"port_src_escalation":                       ret.DtDdosDnsUdpZonePortStats.Stats.Port_src_escalation,
			"current_es_level":                          ret.DtDdosDnsUdpZonePortStats.Stats.Current_es_level,
			"filter_auth_fail":                          ret.DtDdosDnsUdpZonePortStats.Stats.Filter_auth_fail,
			"spoof_detect_fail":                         ret.DtDdosDnsUdpZonePortStats.Stats.Spoof_detect_fail,
			"sess_create":                               ret.DtDdosDnsUdpZonePortStats.Stats.Sess_create,
			"filter_action_blacklist":                   ret.DtDdosDnsUdpZonePortStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                        ret.DtDdosDnsUdpZonePortStats.Stats.Filter_action_drop,
			"filter_action_default_pass":                ret.DtDdosDnsUdpZonePortStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":                   ret.DtDdosDnsUdpZonePortStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":                     ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_crate_src":                     ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_drop_crate_src,
			"exceed_drop_climit_src":                    ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_drop_climit_src,
			"exceed_drop_brate_src":                     ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":                  ret.DtDdosDnsUdpZonePortStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                        ret.DtDdosDnsUdpZonePortStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":                  ret.DtDdosDnsUdpZonePortStats.Stats.Outbound_port_bytes_drop,
			"udp_auth_retry_fail":                       ret.DtDdosDnsUdpZonePortStats.Stats.Udp_auth_retry_fail,
			"src_rate_limit_type0":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_rate_limit_type0,
			"src_rate_limit_type1":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_rate_limit_type1,
			"src_rate_limit_type2":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_rate_limit_type2,
			"src_rate_limit_type3":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_rate_limit_type3,
			"src_rate_limit_type4":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_rate_limit_type4,
			"exceed_drop_brate_src_pkt":                 ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":                 ret.DtDdosDnsUdpZonePortStats.Stats.Port_kbit_rate_exceed_pkt,
			"udp_retry_init":                            ret.DtDdosDnsUdpZonePortStats.Stats.Udp_retry_init,
			"conn_prate_excd":                           ret.DtDdosDnsUdpZonePortStats.Stats.Conn_prate_excd,
			"ntp_monlist_req":                           ret.DtDdosDnsUdpZonePortStats.Stats.Ntp_monlist_req,
			"ntp_monlist_resp":                          ret.DtDdosDnsUdpZonePortStats.Stats.Ntp_monlist_resp,
			"wellknown_sport_drop":                      ret.DtDdosDnsUdpZonePortStats.Stats.Wellknown_sport_drop,
			"payload_too_small":                         ret.DtDdosDnsUdpZonePortStats.Stats.Payload_too_small,
			"payload_too_big":                           ret.DtDdosDnsUdpZonePortStats.Stats.Payload_too_big,
			"fqdn_rate_by_label_count_check":            ret.DtDdosDnsUdpZonePortStats.Stats.Fqdn_rate_by_label_count_check,
			"fqdn_rate_by_label_count_exceed":           ret.DtDdosDnsUdpZonePortStats.Stats.Fqdn_rate_by_label_count_exceed,
			"udp_auth_retry_gap_drop":                   ret.DtDdosDnsUdpZonePortStats.Stats.Udp_auth_retry_gap_drop,
			"fqdn_label_count_exceed":                   ret.DtDdosDnsUdpZonePortStats.Stats.Fqdn_label_count_exceed,
			"rrtype_drop":                               ret.DtDdosDnsUdpZonePortStats.Stats.Rrtype_drop,
			"src_dns_fqdn_label_len_exceed":             ret.DtDdosDnsUdpZonePortStats.Stats.Src_dns_fqdn_label_len_exceed,
			"src_fqdn_label_count_exceed":               ret.DtDdosDnsUdpZonePortStats.Stats.Src_fqdn_label_count_exceed,
			"src_udp_auth_fail":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_auth_fail,
			"src_force_tcp_auth":                        ret.DtDdosDnsUdpZonePortStats.Stats.Src_force_tcp_auth,
			"src_dns_malform_drop":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_dns_malform_drop,
			"force_tcp_auth_timeout":                    ret.DtDdosDnsUdpZonePortStats.Stats.Force_tcp_auth_timeout,
			"udp_auth_drop":                             ret.DtDdosDnsUdpZonePortStats.Stats.Udp_auth_drop,
			"dns_auth_drop":                             ret.DtDdosDnsUdpZonePortStats.Stats.Dns_auth_drop,
			"dns_auth_resp":                             ret.DtDdosDnsUdpZonePortStats.Stats.Dns_auth_resp,
			"dns_query_class_in":                        ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_in,
			"dns_query_class_csnet":                     ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_csnet,
			"dns_query_class_chaos":                     ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_chaos,
			"dns_query_class_hs":                        ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_hs,
			"dns_query_class_none":                      ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_none,
			"dns_query_class_any":                       ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_any,
			"dns_query_class_whitelist_miss":            ret.DtDdosDnsUdpZonePortStats.Stats.Dns_query_class_whitelist_miss,
			"bl":                                        ret.DtDdosDnsUdpZonePortStats.Stats.Bl,
			"src_drop":                                  ret.DtDdosDnsUdpZonePortStats.Stats.Src_drop,
			"frag_rcvd":                                 ret.DtDdosDnsUdpZonePortStats.Stats.Frag_rcvd,
			"frag_drop":                                 ret.DtDdosDnsUdpZonePortStats.Stats.Frag_drop,
			"sess_create_inbound":                       ret.DtDdosDnsUdpZonePortStats.Stats.Sess_create_inbound,
			"sess_create_outbound":                      ret.DtDdosDnsUdpZonePortStats.Stats.Sess_create_outbound,
			"sess_aged":                                 ret.DtDdosDnsUdpZonePortStats.Stats.Sess_aged,
			"udp_retry_pass":                            ret.DtDdosDnsUdpZonePortStats.Stats.Udp_retry_pass,
			"udp_retry_gap_drop":                        ret.DtDdosDnsUdpZonePortStats.Stats.Udp_retry_gap_drop,
			"filter1_match":                             ret.DtDdosDnsUdpZonePortStats.Stats.Filter1_match,
			"filter2_match":                             ret.DtDdosDnsUdpZonePortStats.Stats.Filter2_match,
			"filter3_match":                             ret.DtDdosDnsUdpZonePortStats.Stats.Filter3_match,
			"filter4_match":                             ret.DtDdosDnsUdpZonePortStats.Stats.Filter4_match,
			"filter5_match":                             ret.DtDdosDnsUdpZonePortStats.Stats.Filter5_match,
			"filter_none_match":                         ret.DtDdosDnsUdpZonePortStats.Stats.Filter_none_match,
			"src_payload_too_small":                     ret.DtDdosDnsUdpZonePortStats.Stats.Src_payload_too_small,
			"src_payload_too_big":                       ret.DtDdosDnsUdpZonePortStats.Stats.Src_payload_too_big,
			"src_ntp_monlist_req":                       ret.DtDdosDnsUdpZonePortStats.Stats.Src_ntp_monlist_req,
			"src_ntp_monlist_resp":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_ntp_monlist_resp,
			"src_well_known_port":                       ret.DtDdosDnsUdpZonePortStats.Stats.Src_well_known_port,
			"src_conn_pkt_rate_excd":                    ret.DtDdosDnsUdpZonePortStats.Stats.Src_conn_pkt_rate_excd,
			"src_udp_retry_init":                        ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_retry_init,
			"src_filter_action_blacklist":               ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":                    ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":            ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":               ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_action_whitelist,
			"udp_auth_init":                             ret.DtDdosDnsUdpZonePortStats.Stats.Udp_auth_init,
			"nxdomain_rate_exceed":                      ret.DtDdosDnsUdpZonePortStats.Stats.Nxdomain_rate_exceed,
			"src_udp_auth_init":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_auth_init,
			"dns_malform":                               ret.DtDdosDnsUdpZonePortStats.Stats.Dns_malform,
			"src_dns_malform":                           ret.DtDdosDnsUdpZonePortStats.Stats.Src_dns_malform,
			"dg_request_check_fail":                     ret.DtDdosDnsUdpZonePortStats.Stats.Dg_request_check_fail,
			"secondary_port_pkt_rate_exceed":            ret.DtDdosDnsUdpZonePortStats.Stats.Secondary_port_pkt_rate_exceed,
			"secondary_port_kbit_rate_exceed":           ret.DtDdosDnsUdpZonePortStats.Stats.Secondary_port_kbit_rate_exceed,
			"secondary_port_kbit_rate_exceed_pkt":       ret.DtDdosDnsUdpZonePortStats.Stats.Secondary_port_kbit_rate_exceed_pkt,
			"secondary_port_conn_rate_exceed":           ret.DtDdosDnsUdpZonePortStats.Stats.Secondary_port_conn_rate_exceed,
			"secondary_port_conn_limm_exceed":           ret.DtDdosDnsUdpZonePortStats.Stats.Secondary_port_conn_limm_exceed,
			"query_type_any_drop":                       ret.DtDdosDnsUdpZonePortStats.Stats.Query_type_any_drop,
			"src_query_type_any_drop":                   ret.DtDdosDnsUdpZonePortStats.Stats.Src_query_type_any_drop,
			"src_udp_auth_drop":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_auth_drop,
			"src_dns_auth_drop":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_dns_auth_drop,
			"src_frag_drop":                             ret.DtDdosDnsUdpZonePortStats.Stats.Src_frag_drop,
			"no_policy_class_list_match":                ret.DtDdosDnsUdpZonePortStats.Stats.No_policy_class_list_match,
			"frag_timeout":                              ret.DtDdosDnsUdpZonePortStats.Stats.Frag_timeout,
			"dg_rate_exceed":                            ret.DtDdosDnsUdpZonePortStats.Stats.Dg_rate_exceed,
			"pattern_recognition_proceeded":             ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_recognition_proceeded,
			"pattern_not_found":                         ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_not_found,
			"pattern_recognition_generic_error":         ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_recognition_generic_error,
			"pattern_filter1_match":                     ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_filter1_match,
			"pattern_filter2_match":                     ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_filter2_match,
			"pattern_filter3_match":                     ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_filter3_match,
			"pattern_filter4_match":                     ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_filter4_match,
			"pattern_filter5_match":                     ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_filter5_match,
			"pattern_filter_drop":                       ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_filter_drop,
			"src_filter1_match":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter1_match,
			"src_filter2_match":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter2_match,
			"src_filter3_match":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter3_match,
			"src_filter4_match":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter4_match,
			"src_filter5_match":                         ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter5_match,
			"src_filter_none_match":                     ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":                ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_filter_auth_fail,
			"filter_total_not_match":                    ret.DtDdosDnsUdpZonePortStats.Stats.Filter_total_not_match,
			"src_udp_retry_gap_drop":                    ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_retry_gap_drop,
			"non_query_opcode_pass_through":             ret.DtDdosDnsUdpZonePortStats.Stats.Non_query_opcode_pass_through,
			"sflow_internal_samples_packed":             ret.DtDdosDnsUdpZonePortStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":             ret.DtDdosDnsUdpZonePortStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":               ret.DtDdosDnsUdpZonePortStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":               ret.DtDdosDnsUdpZonePortStats.Stats.Sflow_external_packets_sent,
			"pattern_recognition_sampling_started":      ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_recognition_sampling_started,
			"pattern_recognition_pattern_changed":       ret.DtDdosDnsUdpZonePortStats.Stats.Pattern_recognition_pattern_changed,
			"exceed_action_tunnel":                      ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_action_tunnel,
			"dst_udp_retry_timeout_blacklist":           ret.DtDdosDnsUdpZonePortStats.Stats.Dst_udp_retry_timeout_blacklist,
			"src_udp_auth_timeout":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_auth_timeout,
			"zone_src_udp_retry_timeout_blacklist":      ret.DtDdosDnsUdpZonePortStats.Stats.Zone_src_udp_retry_timeout_blacklist,
			"src_udp_retry_pass":                        ret.DtDdosDnsUdpZonePortStats.Stats.Src_udp_retry_pass,
			"dst_hw_drop":                               ret.DtDdosDnsUdpZonePortStats.Stats.Dst_hw_drop,
			"secondary_port_hit":                        ret.DtDdosDnsUdpZonePortStats.Stats.Secondary_port_hit,
			"src_zone_service_entry_learned":            ret.DtDdosDnsUdpZonePortStats.Stats.Src_zone_service_entry_learned,
			"src_zone_service_entry_aged":               ret.DtDdosDnsUdpZonePortStats.Stats.Src_zone_service_entry_aged,
			"dst_hw_drop_inserted":                      ret.DtDdosDnsUdpZonePortStats.Stats.Dst_hw_drop_inserted,
			"dst_hw_drop_removed":                       ret.DtDdosDnsUdpZonePortStats.Stats.Dst_hw_drop_removed,
			"src_hw_drop_inserted":                      ret.DtDdosDnsUdpZonePortStats.Stats.Src_hw_drop_inserted,
			"src_hw_drop_removed":                       ret.DtDdosDnsUdpZonePortStats.Stats.Src_hw_drop_removed,
			"token_authentication_mismatched":           ret.DtDdosDnsUdpZonePortStats.Stats.Token_authentication_mismatched,
			"token_authentication_invalid":              ret.DtDdosDnsUdpZonePortStats.Stats.Token_authentication_invalid,
			"token_authentication_curr_salt_matched":    ret.DtDdosDnsUdpZonePortStats.Stats.Token_authentication_curr_salt_matched,
			"token_authentication_prev_salt_matched":    ret.DtDdosDnsUdpZonePortStats.Stats.Token_authentication_prev_salt_matched,
			"token_authentication_session_created":      ret.DtDdosDnsUdpZonePortStats.Stats.Token_authentication_session_created,
			"token_authentication_session_created_fail": ret.DtDdosDnsUdpZonePortStats.Stats.Token_authentication_session_created_fail,
			"snat_fail":                                 ret.DtDdosDnsUdpZonePortStats.Stats.Snat_fail,
			"exceed_action_drop":                        ret.DtDdosDnsUdpZonePortStats.Stats.Exceed_action_drop,
			"ew_inbound_port_rcv":                       ret.DtDdosDnsUdpZonePortStats.Stats.Ew_inbound_port_rcv,
			"ew_inbound_port_drop":                      ret.DtDdosDnsUdpZonePortStats.Stats.Ew_inbound_port_drop,
			"ew_inbound_port_sent":                      ret.DtDdosDnsUdpZonePortStats.Stats.Ew_inbound_port_sent,
			"ew_inbound_port_byte_rcv":                  ret.DtDdosDnsUdpZonePortStats.Stats.Ew_inbound_port_byte_rcv,
			"ew_inbound_port_byte_drop":                 ret.DtDdosDnsUdpZonePortStats.Stats.Ew_inbound_port_byte_drop,
			"ew_inbound_port_byte_sent":                 ret.DtDdosDnsUdpZonePortStats.Stats.Ew_inbound_port_byte_sent,
			"ew_outbound_port_rcv":                      ret.DtDdosDnsUdpZonePortStats.Stats.Ew_outbound_port_rcv,
			"ew_outbound_port_drop":                     ret.DtDdosDnsUdpZonePortStats.Stats.Ew_outbound_port_drop,
			"ew_outbound_port_sent":                     ret.DtDdosDnsUdpZonePortStats.Stats.Ew_outbound_port_sent,
			"ew_outbound_port_byte_rcv":                 ret.DtDdosDnsUdpZonePortStats.Stats.Ew_outbound_port_byte_rcv,
			"ew_outbound_port_byte_sent":                ret.DtDdosDnsUdpZonePortStats.Stats.Ew_outbound_port_byte_sent,
			"ew_outbound_port_byte_drop":                ret.DtDdosDnsUdpZonePortStats.Stats.Ew_outbound_port_byte_drop,
			"no_route_drop":                             ret.DtDdosDnsUdpZonePortStats.Stats.No_route_drop,
			"unauth_src_session_reset":                  ret.DtDdosDnsUdpZonePortStats.Stats.Unauth_src_session_reset,
			"src_hw_drop":                               ret.DtDdosDnsUdpZonePortStats.Stats.Src_hw_drop,
			"addr_filter_drop":                          ret.DtDdosDnsUdpZonePortStats.Stats.Addr_filter_drop,
			"addr_filter_bl":                            ret.DtDdosDnsUdpZonePortStats.Stats.Addr_filter_bl,
			"src_learn_overflow":                        ret.DtDdosDnsUdpZonePortStats.Stats.Src_learn_overflow,
		},
	}
}

func getObjectDdosDnsUdpZonePortStatsStats(d []interface{}) edpt.DdosDnsUdpZonePortStatsStats {

	count1 := len(d)
	var ret edpt.DdosDnsUdpZonePortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Force_tcp_auth = in["force_tcp_auth"].(int)
		ret.Udp_auth_pass = in["udp_auth_pass"].(int)
		ret.Rate_limit_type0 = in["rate_limit_type0"].(int)
		ret.Rate_limit_type1 = in["rate_limit_type1"].(int)
		ret.Rate_limit_type2 = in["rate_limit_type2"].(int)
		ret.Rate_limit_type3 = in["rate_limit_type3"].(int)
		ret.Rate_limit_type4 = in["rate_limit_type4"].(int)
		ret.Nxdomain_bl_drop = in["nxdomain_bl_drop"].(int)
		ret.Is_nxdomain = in["is_nxdomain"].(int)
		ret.Nxdomain_drop = in["nxdomain_drop"].(int)
		ret.Udp_auth_fail = in["udp_auth_fail"].(int)
		ret.Malform_drop = in["malform_drop"].(int)
		ret.Fqdn_stage_2_rate_exceed = in["fqdn_stage_2_rate_exceed"].(int)
		ret.Req_sent = in["req_sent"].(int)
		ret.Req_size_exceed = in["req_size_exceed"].(int)
		ret.Req_retrans = in["req_retrans"].(int)
		ret.Fqdn_label_len_exceed = in["fqdn_label_len_exceed"].(int)
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
		ret.Port_src_escalation = in["port_src_escalation"].(int)
		ret.Current_es_level = in["current_es_level"].(int)
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
		ret.Src_rate_limit_type0 = in["src_rate_limit_type0"].(int)
		ret.Src_rate_limit_type1 = in["src_rate_limit_type1"].(int)
		ret.Src_rate_limit_type2 = in["src_rate_limit_type2"].(int)
		ret.Src_rate_limit_type3 = in["src_rate_limit_type3"].(int)
		ret.Src_rate_limit_type4 = in["src_rate_limit_type4"].(int)
		ret.Exceed_drop_brate_src_pkt = in["exceed_drop_brate_src_pkt"].(int)
		ret.Port_kbit_rate_exceed_pkt = in["port_kbit_rate_exceed_pkt"].(int)
		ret.Udp_retry_init = in["udp_retry_init"].(int)
		ret.Conn_prate_excd = in["conn_prate_excd"].(int)
		ret.Ntp_monlist_req = in["ntp_monlist_req"].(int)
		ret.Ntp_monlist_resp = in["ntp_monlist_resp"].(int)
		ret.Wellknown_sport_drop = in["wellknown_sport_drop"].(int)
		ret.Payload_too_small = in["payload_too_small"].(int)
		ret.Payload_too_big = in["payload_too_big"].(int)
		ret.Fqdn_rate_by_label_count_check = in["fqdn_rate_by_label_count_check"].(int)
		ret.Fqdn_rate_by_label_count_exceed = in["fqdn_rate_by_label_count_exceed"].(int)
		ret.Udp_auth_retry_gap_drop = in["udp_auth_retry_gap_drop"].(int)
		ret.Fqdn_label_count_exceed = in["fqdn_label_count_exceed"].(int)
		ret.Rrtype_drop = in["rrtype_drop"].(int)
		ret.Src_dns_fqdn_label_len_exceed = in["src_dns_fqdn_label_len_exceed"].(int)
		ret.Src_fqdn_label_count_exceed = in["src_fqdn_label_count_exceed"].(int)
		ret.Src_udp_auth_fail = in["src_udp_auth_fail"].(int)
		ret.Src_force_tcp_auth = in["src_force_tcp_auth"].(int)
		ret.Src_dns_malform_drop = in["src_dns_malform_drop"].(int)
		ret.Force_tcp_auth_timeout = in["force_tcp_auth_timeout"].(int)
		ret.Udp_auth_drop = in["udp_auth_drop"].(int)
		ret.Dns_auth_drop = in["dns_auth_drop"].(int)
		ret.Dns_auth_resp = in["dns_auth_resp"].(int)
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
		ret.Udp_auth_init = in["udp_auth_init"].(int)
		ret.Nxdomain_rate_exceed = in["nxdomain_rate_exceed"].(int)
		ret.Src_udp_auth_init = in["src_udp_auth_init"].(int)
		ret.Dns_malform = in["dns_malform"].(int)
		ret.Src_dns_malform = in["src_dns_malform"].(int)
		ret.Dg_request_check_fail = in["dg_request_check_fail"].(int)
		ret.Secondary_port_pkt_rate_exceed = in["secondary_port_pkt_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed = in["secondary_port_kbit_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed_pkt = in["secondary_port_kbit_rate_exceed_pkt"].(int)
		ret.Secondary_port_conn_rate_exceed = in["secondary_port_conn_rate_exceed"].(int)
		ret.Secondary_port_conn_limm_exceed = in["secondary_port_conn_limm_exceed"].(int)
		ret.Query_type_any_drop = in["query_type_any_drop"].(int)
		ret.Src_query_type_any_drop = in["src_query_type_any_drop"].(int)
		ret.Src_udp_auth_drop = in["src_udp_auth_drop"].(int)
		ret.Src_dns_auth_drop = in["src_dns_auth_drop"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.No_policy_class_list_match = in["no_policy_class_list_match"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Dg_rate_exceed = in["dg_rate_exceed"].(int)
		ret.Pattern_recognition_proceeded = in["pattern_recognition_proceeded"].(int)
		ret.Pattern_not_found = in["pattern_not_found"].(int)
		ret.Pattern_recognition_generic_error = in["pattern_recognition_generic_error"].(int)
		ret.Pattern_filter1_match = in["pattern_filter1_match"].(int)
		ret.Pattern_filter2_match = in["pattern_filter2_match"].(int)
		ret.Pattern_filter3_match = in["pattern_filter3_match"].(int)
		ret.Pattern_filter4_match = in["pattern_filter4_match"].(int)
		ret.Pattern_filter5_match = in["pattern_filter5_match"].(int)
		ret.Pattern_filter_drop = in["pattern_filter_drop"].(int)
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
		ret.Non_query_opcode_pass_through = in["non_query_opcode_pass_through"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Pattern_recognition_sampling_started = in["pattern_recognition_sampling_started"].(int)
		ret.Pattern_recognition_pattern_changed = in["pattern_recognition_pattern_changed"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Dst_udp_retry_timeout_blacklist = in["dst_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_auth_timeout = in["src_udp_auth_timeout"].(int)
		ret.Zone_src_udp_retry_timeout_blacklist = in["zone_src_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_retry_pass = in["src_udp_retry_pass"].(int)
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

func dataToEndpointDdosDnsUdpZonePortStats(d *schema.ResourceData) edpt.DdosDnsUdpZonePortStats {
	var ret edpt.DdosDnsUdpZonePortStats

	ret.Stats = getObjectDdosDnsUdpZonePortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
