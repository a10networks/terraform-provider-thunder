package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosQuicZonePortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_quic_zone_port_stats`: Statistics for the object quic-zone-port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosQuicZonePortStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"filter_auth_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Auth Failed",
						},
						"spoof_detect_fail": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Retry Timeout",
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
						"sess_create": {
							Type: schema.TypeInt, Optional: true, Description: "Session Create",
						},
						"payload_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Small",
						},
						"payload_too_big": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Payload Too Large",
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
						"src_udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry-Gap Dropped",
						},
						"src_udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth Dropped",
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
						"dst_hw_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Packets Dropped",
						},
						"secondary_port_hit": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port Hit",
						},
						"processed": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Headers Processed",
						},
						"initial": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Initial",
						},
						"version_negotiation": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Version Negotiation",
						},
						"rtt0": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC 0-RTT",
						},
						"handshake": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Handshake",
						},
						"retry": {
							Type: schema.TypeInt, Optional: true, Description: "QUIC Retry",
						},
						"invalid_packet_type": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid Packet Type",
						},
						"packet_parse_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Parse Failed",
						},
						"version_match_action_taken": {
							Type: schema.TypeInt, Optional: true, Description: "Version Match Action Taken",
						},
						"version_match_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Version Match Action Drop",
						},
						"version_match_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Version Match Action Blacklist",
						},
						"malformed_action_taken": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Action Taken",
						},
						"malformed_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Action Drop",
						},
						"malformed_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed Action Blacklist",
						},
						"malformed_dcid_len_max_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed DCID Length Maximum Exceed",
						},
						"malformed_scid_len_max_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Malformed SCID Length Maximum Exceed",
						},
						"initial_force_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Initial Force-TCP",
						},
						"fixed_bit_not_set": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed Bit Not Set",
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

func resourceDdosQuicZonePortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosQuicZonePortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosQuicZonePortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosQuicZonePortStatsStats := setObjectDdosQuicZonePortStatsStats(res)
		d.Set("stats", DdosQuicZonePortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosQuicZonePortStatsStats(ret edpt.DataDdosQuicZonePortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter1_match":                             ret.DtDdosQuicZonePortStats.Stats.Filter1_match,
			"filter2_match":                             ret.DtDdosQuicZonePortStats.Stats.Filter2_match,
			"filter3_match":                             ret.DtDdosQuicZonePortStats.Stats.Filter3_match,
			"filter4_match":                             ret.DtDdosQuicZonePortStats.Stats.Filter4_match,
			"filter5_match":                             ret.DtDdosQuicZonePortStats.Stats.Filter5_match,
			"filter_none_match":                         ret.DtDdosQuicZonePortStats.Stats.Filter_none_match,
			"port_rcvd":                                 ret.DtDdosQuicZonePortStats.Stats.Port_rcvd,
			"port_drop":                                 ret.DtDdosQuicZonePortStats.Stats.Port_drop,
			"port_pkt_sent":                             ret.DtDdosQuicZonePortStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":                      ret.DtDdosQuicZonePortStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":                     ret.DtDdosQuicZonePortStats.Stats.Port_kbit_rate_exceed,
			"port_conn_rate_exceed":                     ret.DtDdosQuicZonePortStats.Stats.Port_conn_rate_exceed,
			"port_conn_limm_exceed":                     ret.DtDdosQuicZonePortStats.Stats.Port_conn_limm_exceed,
			"filter_auth_fail":                          ret.DtDdosQuicZonePortStats.Stats.Filter_auth_fail,
			"spoof_detect_fail":                         ret.DtDdosQuicZonePortStats.Stats.Spoof_detect_fail,
			"port_bytes":                                ret.DtDdosQuicZonePortStats.Stats.Port_bytes,
			"outbound_port_bytes":                       ret.DtDdosQuicZonePortStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                        ret.DtDdosQuicZonePortStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":                    ret.DtDdosQuicZonePortStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                           ret.DtDdosQuicZonePortStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                           ret.DtDdosQuicZonePortStats.Stats.Port_bytes_drop,
			"port_src_bl":                               ret.DtDdosQuicZonePortStats.Stats.Port_src_bl,
			"port_src_escalation":                       ret.DtDdosQuicZonePortStats.Stats.Port_src_escalation,
			"current_es_level":                          ret.DtDdosQuicZonePortStats.Stats.Current_es_level,
			"sess_create":                               ret.DtDdosQuicZonePortStats.Stats.Sess_create,
			"payload_too_small":                         ret.DtDdosQuicZonePortStats.Stats.Payload_too_small,
			"payload_too_big":                           ret.DtDdosQuicZonePortStats.Stats.Payload_too_big,
			"filter_action_blacklist":                   ret.DtDdosQuicZonePortStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                        ret.DtDdosQuicZonePortStats.Stats.Filter_action_drop,
			"filter_action_default_pass":                ret.DtDdosQuicZonePortStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":                   ret.DtDdosQuicZonePortStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":                     ret.DtDdosQuicZonePortStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_crate_src":                     ret.DtDdosQuicZonePortStats.Stats.Exceed_drop_crate_src,
			"exceed_drop_climit_src":                    ret.DtDdosQuicZonePortStats.Stats.Exceed_drop_climit_src,
			"exceed_drop_brate_src":                     ret.DtDdosQuicZonePortStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":                  ret.DtDdosQuicZonePortStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                        ret.DtDdosQuicZonePortStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":                  ret.DtDdosQuicZonePortStats.Stats.Outbound_port_bytes_drop,
			"exceed_drop_brate_src_pkt":                 ret.DtDdosQuicZonePortStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":                 ret.DtDdosQuicZonePortStats.Stats.Port_kbit_rate_exceed_pkt,
			"udp_retry_init":                            ret.DtDdosQuicZonePortStats.Stats.Udp_retry_init,
			"conn_prate_excd":                           ret.DtDdosQuicZonePortStats.Stats.Conn_prate_excd,
			"ntp_monlist_req":                           ret.DtDdosQuicZonePortStats.Stats.Ntp_monlist_req,
			"ntp_monlist_resp":                          ret.DtDdosQuicZonePortStats.Stats.Ntp_monlist_resp,
			"wellknown_sport_drop":                      ret.DtDdosQuicZonePortStats.Stats.Wellknown_sport_drop,
			"udp_auth_drop":                             ret.DtDdosQuicZonePortStats.Stats.Udp_auth_drop,
			"bl":                                        ret.DtDdosQuicZonePortStats.Stats.Bl,
			"src_drop":                                  ret.DtDdosQuicZonePortStats.Stats.Src_drop,
			"frag_rcvd":                                 ret.DtDdosQuicZonePortStats.Stats.Frag_rcvd,
			"frag_drop":                                 ret.DtDdosQuicZonePortStats.Stats.Frag_drop,
			"sess_create_inbound":                       ret.DtDdosQuicZonePortStats.Stats.Sess_create_inbound,
			"sess_create_outbound":                      ret.DtDdosQuicZonePortStats.Stats.Sess_create_outbound,
			"sess_aged":                                 ret.DtDdosQuicZonePortStats.Stats.Sess_aged,
			"udp_retry_pass":                            ret.DtDdosQuicZonePortStats.Stats.Udp_retry_pass,
			"udp_retry_gap_drop":                        ret.DtDdosQuicZonePortStats.Stats.Udp_retry_gap_drop,
			"src_payload_too_small":                     ret.DtDdosQuicZonePortStats.Stats.Src_payload_too_small,
			"src_payload_too_big":                       ret.DtDdosQuicZonePortStats.Stats.Src_payload_too_big,
			"src_ntp_monlist_req":                       ret.DtDdosQuicZonePortStats.Stats.Src_ntp_monlist_req,
			"src_ntp_monlist_resp":                      ret.DtDdosQuicZonePortStats.Stats.Src_ntp_monlist_resp,
			"src_well_known_port":                       ret.DtDdosQuicZonePortStats.Stats.Src_well_known_port,
			"src_conn_pkt_rate_excd":                    ret.DtDdosQuicZonePortStats.Stats.Src_conn_pkt_rate_excd,
			"src_udp_retry_init":                        ret.DtDdosQuicZonePortStats.Stats.Src_udp_retry_init,
			"src_filter_action_blacklist":               ret.DtDdosQuicZonePortStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":                    ret.DtDdosQuicZonePortStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":            ret.DtDdosQuicZonePortStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":               ret.DtDdosQuicZonePortStats.Stats.Src_filter_action_whitelist,
			"secondary_port_pkt_rate_exceed":            ret.DtDdosQuicZonePortStats.Stats.Secondary_port_pkt_rate_exceed,
			"secondary_port_kbit_rate_exceed":           ret.DtDdosQuicZonePortStats.Stats.Secondary_port_kbit_rate_exceed,
			"secondary_port_kbit_rate_exceed_pkt":       ret.DtDdosQuicZonePortStats.Stats.Secondary_port_kbit_rate_exceed_pkt,
			"secondary_port_conn_rate_exceed":           ret.DtDdosQuicZonePortStats.Stats.Secondary_port_conn_rate_exceed,
			"secondary_port_conn_limm_exceed":           ret.DtDdosQuicZonePortStats.Stats.Secondary_port_conn_limm_exceed,
			"src_udp_retry_gap_drop":                    ret.DtDdosQuicZonePortStats.Stats.Src_udp_retry_gap_drop,
			"src_udp_auth_drop":                         ret.DtDdosQuicZonePortStats.Stats.Src_udp_auth_drop,
			"src_frag_drop":                             ret.DtDdosQuicZonePortStats.Stats.Src_frag_drop,
			"no_policy_class_list_match":                ret.DtDdosQuicZonePortStats.Stats.No_policy_class_list_match,
			"frag_timeout":                              ret.DtDdosQuicZonePortStats.Stats.Frag_timeout,
			"src_filter1_match":                         ret.DtDdosQuicZonePortStats.Stats.Src_filter1_match,
			"src_filter2_match":                         ret.DtDdosQuicZonePortStats.Stats.Src_filter2_match,
			"src_filter3_match":                         ret.DtDdosQuicZonePortStats.Stats.Src_filter3_match,
			"src_filter4_match":                         ret.DtDdosQuicZonePortStats.Stats.Src_filter4_match,
			"src_filter5_match":                         ret.DtDdosQuicZonePortStats.Stats.Src_filter5_match,
			"src_filter_none_match":                     ret.DtDdosQuicZonePortStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":                ret.DtDdosQuicZonePortStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":                      ret.DtDdosQuicZonePortStats.Stats.Src_filter_auth_fail,
			"filter_total_not_match":                    ret.DtDdosQuicZonePortStats.Stats.Filter_total_not_match,
			"sflow_internal_samples_packed":             ret.DtDdosQuicZonePortStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":             ret.DtDdosQuicZonePortStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":               ret.DtDdosQuicZonePortStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":               ret.DtDdosQuicZonePortStats.Stats.Sflow_external_packets_sent,
			"exceed_action_tunnel":                      ret.DtDdosQuicZonePortStats.Stats.Exceed_action_tunnel,
			"dst_udp_retry_timeout_blacklist":           ret.DtDdosQuicZonePortStats.Stats.Dst_udp_retry_timeout_blacklist,
			"src_udp_auth_timeout":                      ret.DtDdosQuicZonePortStats.Stats.Src_udp_auth_timeout,
			"zone_src_udp_retry_timeout_blacklist":      ret.DtDdosQuicZonePortStats.Stats.Zone_src_udp_retry_timeout_blacklist,
			"src_udp_retry_pass":                        ret.DtDdosQuicZonePortStats.Stats.Src_udp_retry_pass,
			"dst_hw_drop":                               ret.DtDdosQuicZonePortStats.Stats.Dst_hw_drop,
			"secondary_port_hit":                        ret.DtDdosQuicZonePortStats.Stats.Secondary_port_hit,
			"processed":                                 ret.DtDdosQuicZonePortStats.Stats.Processed,
			"initial":                                   ret.DtDdosQuicZonePortStats.Stats.Initial,
			"version_negotiation":                       ret.DtDdosQuicZonePortStats.Stats.Version_negotiation,
			"rtt0":                                      ret.DtDdosQuicZonePortStats.Stats.Rtt0,
			"handshake":                                 ret.DtDdosQuicZonePortStats.Stats.Handshake,
			"retry":                                     ret.DtDdosQuicZonePortStats.Stats.Retry,
			"invalid_packet_type":                       ret.DtDdosQuicZonePortStats.Stats.Invalid_packet_type,
			"packet_parse_failed":                       ret.DtDdosQuicZonePortStats.Stats.Packet_parse_failed,
			"version_match_action_taken":                ret.DtDdosQuicZonePortStats.Stats.Version_match_action_taken,
			"version_match_action_drop":                 ret.DtDdosQuicZonePortStats.Stats.Version_match_action_drop,
			"version_match_action_blacklist":            ret.DtDdosQuicZonePortStats.Stats.Version_match_action_blacklist,
			"malformed_action_taken":                    ret.DtDdosQuicZonePortStats.Stats.Malformed_action_taken,
			"malformed_action_drop":                     ret.DtDdosQuicZonePortStats.Stats.Malformed_action_drop,
			"malformed_action_blacklist":                ret.DtDdosQuicZonePortStats.Stats.Malformed_action_blacklist,
			"malformed_dcid_len_max_exceed":             ret.DtDdosQuicZonePortStats.Stats.Malformed_dcid_len_max_exceed,
			"malformed_scid_len_max_exceed":             ret.DtDdosQuicZonePortStats.Stats.Malformed_scid_len_max_exceed,
			"initial_force_tcp":                         ret.DtDdosQuicZonePortStats.Stats.Initial_force_tcp,
			"fixed_bit_not_set":                         ret.DtDdosQuicZonePortStats.Stats.Fixed_bit_not_set,
			"src_zone_service_entry_learned":            ret.DtDdosQuicZonePortStats.Stats.Src_zone_service_entry_learned,
			"src_zone_service_entry_aged":               ret.DtDdosQuicZonePortStats.Stats.Src_zone_service_entry_aged,
			"dst_hw_drop_inserted":                      ret.DtDdosQuicZonePortStats.Stats.Dst_hw_drop_inserted,
			"dst_hw_drop_removed":                       ret.DtDdosQuicZonePortStats.Stats.Dst_hw_drop_removed,
			"src_hw_drop_inserted":                      ret.DtDdosQuicZonePortStats.Stats.Src_hw_drop_inserted,
			"src_hw_drop_removed":                       ret.DtDdosQuicZonePortStats.Stats.Src_hw_drop_removed,
			"token_authentication_mismatched":           ret.DtDdosQuicZonePortStats.Stats.Token_authentication_mismatched,
			"token_authentication_invalid":              ret.DtDdosQuicZonePortStats.Stats.Token_authentication_invalid,
			"token_authentication_curr_salt_matched":    ret.DtDdosQuicZonePortStats.Stats.Token_authentication_curr_salt_matched,
			"token_authentication_prev_salt_matched":    ret.DtDdosQuicZonePortStats.Stats.Token_authentication_prev_salt_matched,
			"token_authentication_session_created":      ret.DtDdosQuicZonePortStats.Stats.Token_authentication_session_created,
			"token_authentication_session_created_fail": ret.DtDdosQuicZonePortStats.Stats.Token_authentication_session_created_fail,
			"snat_fail":                                 ret.DtDdosQuicZonePortStats.Stats.Snat_fail,
			"exceed_action_drop":                        ret.DtDdosQuicZonePortStats.Stats.Exceed_action_drop,
			"ew_inbound_port_rcv":                       ret.DtDdosQuicZonePortStats.Stats.Ew_inbound_port_rcv,
			"ew_inbound_port_drop":                      ret.DtDdosQuicZonePortStats.Stats.Ew_inbound_port_drop,
			"ew_inbound_port_sent":                      ret.DtDdosQuicZonePortStats.Stats.Ew_inbound_port_sent,
			"ew_inbound_port_byte_rcv":                  ret.DtDdosQuicZonePortStats.Stats.Ew_inbound_port_byte_rcv,
			"ew_inbound_port_byte_drop":                 ret.DtDdosQuicZonePortStats.Stats.Ew_inbound_port_byte_drop,
			"ew_inbound_port_byte_sent":                 ret.DtDdosQuicZonePortStats.Stats.Ew_inbound_port_byte_sent,
			"ew_outbound_port_rcv":                      ret.DtDdosQuicZonePortStats.Stats.Ew_outbound_port_rcv,
			"ew_outbound_port_drop":                     ret.DtDdosQuicZonePortStats.Stats.Ew_outbound_port_drop,
			"ew_outbound_port_sent":                     ret.DtDdosQuicZonePortStats.Stats.Ew_outbound_port_sent,
			"ew_outbound_port_byte_rcv":                 ret.DtDdosQuicZonePortStats.Stats.Ew_outbound_port_byte_rcv,
			"ew_outbound_port_byte_sent":                ret.DtDdosQuicZonePortStats.Stats.Ew_outbound_port_byte_sent,
			"ew_outbound_port_byte_drop":                ret.DtDdosQuicZonePortStats.Stats.Ew_outbound_port_byte_drop,
			"no_route_drop":                             ret.DtDdosQuicZonePortStats.Stats.No_route_drop,
			"unauth_src_session_reset":                  ret.DtDdosQuicZonePortStats.Stats.Unauth_src_session_reset,
			"src_hw_drop":                               ret.DtDdosQuicZonePortStats.Stats.Src_hw_drop,
			"addr_filter_drop":                          ret.DtDdosQuicZonePortStats.Stats.Addr_filter_drop,
			"addr_filter_bl":                            ret.DtDdosQuicZonePortStats.Stats.Addr_filter_bl,
			"src_learn_overflow":                        ret.DtDdosQuicZonePortStats.Stats.Src_learn_overflow,
		},
	}
}

func getObjectDdosQuicZonePortStatsStats(d []interface{}) edpt.DdosQuicZonePortStatsStats {

	count1 := len(d)
	var ret edpt.DdosQuicZonePortStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
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
		ret.Filter_auth_fail = in["filter_auth_fail"].(int)
		ret.Spoof_detect_fail = in["spoof_detect_fail"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Port_src_escalation = in["port_src_escalation"].(int)
		ret.Current_es_level = in["current_es_level"].(int)
		ret.Sess_create = in["sess_create"].(int)
		ret.Payload_too_small = in["payload_too_small"].(int)
		ret.Payload_too_big = in["payload_too_big"].(int)
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
		ret.Udp_retry_init = in["udp_retry_init"].(int)
		ret.Conn_prate_excd = in["conn_prate_excd"].(int)
		ret.Ntp_monlist_req = in["ntp_monlist_req"].(int)
		ret.Ntp_monlist_resp = in["ntp_monlist_resp"].(int)
		ret.Wellknown_sport_drop = in["wellknown_sport_drop"].(int)
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
		ret.Secondary_port_pkt_rate_exceed = in["secondary_port_pkt_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed = in["secondary_port_kbit_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed_pkt = in["secondary_port_kbit_rate_exceed_pkt"].(int)
		ret.Secondary_port_conn_rate_exceed = in["secondary_port_conn_rate_exceed"].(int)
		ret.Secondary_port_conn_limm_exceed = in["secondary_port_conn_limm_exceed"].(int)
		ret.Src_udp_retry_gap_drop = in["src_udp_retry_gap_drop"].(int)
		ret.Src_udp_auth_drop = in["src_udp_auth_drop"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.No_policy_class_list_match = in["no_policy_class_list_match"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Src_filter1_match = in["src_filter1_match"].(int)
		ret.Src_filter2_match = in["src_filter2_match"].(int)
		ret.Src_filter3_match = in["src_filter3_match"].(int)
		ret.Src_filter4_match = in["src_filter4_match"].(int)
		ret.Src_filter5_match = in["src_filter5_match"].(int)
		ret.Src_filter_none_match = in["src_filter_none_match"].(int)
		ret.Src_filter_total_not_match = in["src_filter_total_not_match"].(int)
		ret.Src_filter_auth_fail = in["src_filter_auth_fail"].(int)
		ret.Filter_total_not_match = in["filter_total_not_match"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Dst_udp_retry_timeout_blacklist = in["dst_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_auth_timeout = in["src_udp_auth_timeout"].(int)
		ret.Zone_src_udp_retry_timeout_blacklist = in["zone_src_udp_retry_timeout_blacklist"].(int)
		ret.Src_udp_retry_pass = in["src_udp_retry_pass"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Secondary_port_hit = in["secondary_port_hit"].(int)
		ret.Processed = in["processed"].(int)
		ret.Initial = in["initial"].(int)
		ret.Version_negotiation = in["version_negotiation"].(int)
		ret.Rtt0 = in["rtt0"].(int)
		ret.Handshake = in["handshake"].(int)
		ret.Retry = in["retry"].(int)
		ret.Invalid_packet_type = in["invalid_packet_type"].(int)
		ret.Packet_parse_failed = in["packet_parse_failed"].(int)
		ret.Version_match_action_taken = in["version_match_action_taken"].(int)
		ret.Version_match_action_drop = in["version_match_action_drop"].(int)
		ret.Version_match_action_blacklist = in["version_match_action_blacklist"].(int)
		ret.Malformed_action_taken = in["malformed_action_taken"].(int)
		ret.Malformed_action_drop = in["malformed_action_drop"].(int)
		ret.Malformed_action_blacklist = in["malformed_action_blacklist"].(int)
		ret.Malformed_dcid_len_max_exceed = in["malformed_dcid_len_max_exceed"].(int)
		ret.Malformed_scid_len_max_exceed = in["malformed_scid_len_max_exceed"].(int)
		ret.Initial_force_tcp = in["initial_force_tcp"].(int)
		ret.Fixed_bit_not_set = in["fixed_bit_not_set"].(int)
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

func dataToEndpointDdosQuicZonePortStats(d *schema.ResourceData) edpt.DdosQuicZonePortStats {
	var ret edpt.DdosQuicZonePortStats

	ret.Stats = getObjectDdosQuicZonePortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
