package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosUdpPortStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_udp_port_stats`: Statistics for the object udp-port\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosUdpPortStatsRead,

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
						"src_udp_retry_gap_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry-Gap Dropped",
						},
						"src_udp_auth_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Auth Dropped",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
						},
						"frag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
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
						"src_udp_auth_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Timeout",
						},
						"src_udp_retry_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Src UDP Retry Passed",
						},
						"dst_hw_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Dst Hardware Packets Dropped",
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
					},
				},
			},
		},
	}
}

func resourceDdosUdpPortStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosUdpPortStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosUdpPortStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosUdpPortStatsStats := setObjectDdosUdpPortStatsStats(res)
		d.Set("stats", DdosUdpPortStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosUdpPortStatsStats(ret edpt.DataDdosUdpPortStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"filter1_match":                             ret.DtDdosUdpPortStats.Stats.Filter1_match,
			"filter2_match":                             ret.DtDdosUdpPortStats.Stats.Filter2_match,
			"filter3_match":                             ret.DtDdosUdpPortStats.Stats.Filter3_match,
			"filter4_match":                             ret.DtDdosUdpPortStats.Stats.Filter4_match,
			"filter5_match":                             ret.DtDdosUdpPortStats.Stats.Filter5_match,
			"filter_none_match":                         ret.DtDdosUdpPortStats.Stats.Filter_none_match,
			"port_rcvd":                                 ret.DtDdosUdpPortStats.Stats.Port_rcvd,
			"port_drop":                                 ret.DtDdosUdpPortStats.Stats.Port_drop,
			"port_pkt_sent":                             ret.DtDdosUdpPortStats.Stats.Port_pkt_sent,
			"port_pkt_rate_exceed":                      ret.DtDdosUdpPortStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":                     ret.DtDdosUdpPortStats.Stats.Port_kbit_rate_exceed,
			"port_conn_rate_exceed":                     ret.DtDdosUdpPortStats.Stats.Port_conn_rate_exceed,
			"port_conn_limm_exceed":                     ret.DtDdosUdpPortStats.Stats.Port_conn_limm_exceed,
			"filter_auth_fail":                          ret.DtDdosUdpPortStats.Stats.Filter_auth_fail,
			"spoof_detect_fail":                         ret.DtDdosUdpPortStats.Stats.Spoof_detect_fail,
			"port_bytes":                                ret.DtDdosUdpPortStats.Stats.Port_bytes,
			"outbound_port_bytes":                       ret.DtDdosUdpPortStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                        ret.DtDdosUdpPortStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":                    ret.DtDdosUdpPortStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                           ret.DtDdosUdpPortStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                           ret.DtDdosUdpPortStats.Stats.Port_bytes_drop,
			"port_src_bl":                               ret.DtDdosUdpPortStats.Stats.Port_src_bl,
			"sess_create":                               ret.DtDdosUdpPortStats.Stats.Sess_create,
			"payload_too_small":                         ret.DtDdosUdpPortStats.Stats.Payload_too_small,
			"payload_too_big":                           ret.DtDdosUdpPortStats.Stats.Payload_too_big,
			"filter_action_blacklist":                   ret.DtDdosUdpPortStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                        ret.DtDdosUdpPortStats.Stats.Filter_action_drop,
			"filter_action_default_pass":                ret.DtDdosUdpPortStats.Stats.Filter_action_default_pass,
			"filter_action_whitelist":                   ret.DtDdosUdpPortStats.Stats.Filter_action_whitelist,
			"exceed_drop_prate_src":                     ret.DtDdosUdpPortStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_crate_src":                     ret.DtDdosUdpPortStats.Stats.Exceed_drop_crate_src,
			"exceed_drop_climit_src":                    ret.DtDdosUdpPortStats.Stats.Exceed_drop_climit_src,
			"exceed_drop_brate_src":                     ret.DtDdosUdpPortStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":                  ret.DtDdosUdpPortStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                        ret.DtDdosUdpPortStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":                  ret.DtDdosUdpPortStats.Stats.Outbound_port_bytes_drop,
			"exceed_drop_brate_src_pkt":                 ret.DtDdosUdpPortStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":                 ret.DtDdosUdpPortStats.Stats.Port_kbit_rate_exceed_pkt,
			"udp_retry_init":                            ret.DtDdosUdpPortStats.Stats.Udp_retry_init,
			"conn_prate_excd":                           ret.DtDdosUdpPortStats.Stats.Conn_prate_excd,
			"ntp_monlist_req":                           ret.DtDdosUdpPortStats.Stats.Ntp_monlist_req,
			"ntp_monlist_resp":                          ret.DtDdosUdpPortStats.Stats.Ntp_monlist_resp,
			"wellknown_sport_drop":                      ret.DtDdosUdpPortStats.Stats.Wellknown_sport_drop,
			"udp_auth_drop":                             ret.DtDdosUdpPortStats.Stats.Udp_auth_drop,
			"bl":                                        ret.DtDdosUdpPortStats.Stats.Bl,
			"src_drop":                                  ret.DtDdosUdpPortStats.Stats.Src_drop,
			"frag_rcvd":                                 ret.DtDdosUdpPortStats.Stats.Frag_rcvd,
			"frag_drop":                                 ret.DtDdosUdpPortStats.Stats.Frag_drop,
			"sess_create_inbound":                       ret.DtDdosUdpPortStats.Stats.Sess_create_inbound,
			"sess_create_outbound":                      ret.DtDdosUdpPortStats.Stats.Sess_create_outbound,
			"sess_aged":                                 ret.DtDdosUdpPortStats.Stats.Sess_aged,
			"udp_retry_pass":                            ret.DtDdosUdpPortStats.Stats.Udp_retry_pass,
			"udp_retry_gap_drop":                        ret.DtDdosUdpPortStats.Stats.Udp_retry_gap_drop,
			"src_payload_too_small":                     ret.DtDdosUdpPortStats.Stats.Src_payload_too_small,
			"src_payload_too_big":                       ret.DtDdosUdpPortStats.Stats.Src_payload_too_big,
			"src_ntp_monlist_req":                       ret.DtDdosUdpPortStats.Stats.Src_ntp_monlist_req,
			"src_ntp_monlist_resp":                      ret.DtDdosUdpPortStats.Stats.Src_ntp_monlist_resp,
			"src_well_known_port":                       ret.DtDdosUdpPortStats.Stats.Src_well_known_port,
			"src_conn_pkt_rate_excd":                    ret.DtDdosUdpPortStats.Stats.Src_conn_pkt_rate_excd,
			"src_udp_retry_init":                        ret.DtDdosUdpPortStats.Stats.Src_udp_retry_init,
			"src_filter_action_blacklist":               ret.DtDdosUdpPortStats.Stats.Src_filter_action_blacklist,
			"src_filter_action_drop":                    ret.DtDdosUdpPortStats.Stats.Src_filter_action_drop,
			"src_filter_action_default_pass":            ret.DtDdosUdpPortStats.Stats.Src_filter_action_default_pass,
			"src_filter_action_whitelist":               ret.DtDdosUdpPortStats.Stats.Src_filter_action_whitelist,
			"src_udp_retry_gap_drop":                    ret.DtDdosUdpPortStats.Stats.Src_udp_retry_gap_drop,
			"src_udp_auth_drop":                         ret.DtDdosUdpPortStats.Stats.Src_udp_auth_drop,
			"src_frag_drop":                             ret.DtDdosUdpPortStats.Stats.Src_frag_drop,
			"frag_timeout":                              ret.DtDdosUdpPortStats.Stats.Frag_timeout,
			"pattern_recognition_proceeded":             ret.DtDdosUdpPortStats.Stats.Pattern_recognition_proceeded,
			"pattern_not_found":                         ret.DtDdosUdpPortStats.Stats.Pattern_not_found,
			"pattern_recognition_generic_error":         ret.DtDdosUdpPortStats.Stats.Pattern_recognition_generic_error,
			"pattern_filter1_match":                     ret.DtDdosUdpPortStats.Stats.Pattern_filter1_match,
			"pattern_filter2_match":                     ret.DtDdosUdpPortStats.Stats.Pattern_filter2_match,
			"pattern_filter3_match":                     ret.DtDdosUdpPortStats.Stats.Pattern_filter3_match,
			"pattern_filter4_match":                     ret.DtDdosUdpPortStats.Stats.Pattern_filter4_match,
			"pattern_filter5_match":                     ret.DtDdosUdpPortStats.Stats.Pattern_filter5_match,
			"pattern_filter_drop":                       ret.DtDdosUdpPortStats.Stats.Pattern_filter_drop,
			"src_filter1_match":                         ret.DtDdosUdpPortStats.Stats.Src_filter1_match,
			"src_filter2_match":                         ret.DtDdosUdpPortStats.Stats.Src_filter2_match,
			"src_filter3_match":                         ret.DtDdosUdpPortStats.Stats.Src_filter3_match,
			"src_filter4_match":                         ret.DtDdosUdpPortStats.Stats.Src_filter4_match,
			"src_filter5_match":                         ret.DtDdosUdpPortStats.Stats.Src_filter5_match,
			"src_filter_none_match":                     ret.DtDdosUdpPortStats.Stats.Src_filter_none_match,
			"src_filter_total_not_match":                ret.DtDdosUdpPortStats.Stats.Src_filter_total_not_match,
			"src_filter_auth_fail":                      ret.DtDdosUdpPortStats.Stats.Src_filter_auth_fail,
			"filter_total_not_match":                    ret.DtDdosUdpPortStats.Stats.Filter_total_not_match,
			"sflow_internal_samples_packed":             ret.DtDdosUdpPortStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":             ret.DtDdosUdpPortStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":               ret.DtDdosUdpPortStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":               ret.DtDdosUdpPortStats.Stats.Sflow_external_packets_sent,
			"pattern_recognition_sampling_started":      ret.DtDdosUdpPortStats.Stats.Pattern_recognition_sampling_started,
			"pattern_recognition_pattern_changed":       ret.DtDdosUdpPortStats.Stats.Pattern_recognition_pattern_changed,
			"exceed_action_tunnel":                      ret.DtDdosUdpPortStats.Stats.Exceed_action_tunnel,
			"src_udp_auth_timeout":                      ret.DtDdosUdpPortStats.Stats.Src_udp_auth_timeout,
			"src_udp_retry_pass":                        ret.DtDdosUdpPortStats.Stats.Src_udp_retry_pass,
			"dst_hw_drop":                               ret.DtDdosUdpPortStats.Stats.Dst_hw_drop,
			"token_authentication_mismatched":           ret.DtDdosUdpPortStats.Stats.Token_authentication_mismatched,
			"token_authentication_invalid":              ret.DtDdosUdpPortStats.Stats.Token_authentication_invalid,
			"token_authentication_curr_salt_matched":    ret.DtDdosUdpPortStats.Stats.Token_authentication_curr_salt_matched,
			"token_authentication_prev_salt_matched":    ret.DtDdosUdpPortStats.Stats.Token_authentication_prev_salt_matched,
			"token_authentication_session_created":      ret.DtDdosUdpPortStats.Stats.Token_authentication_session_created,
			"token_authentication_session_created_fail": ret.DtDdosUdpPortStats.Stats.Token_authentication_session_created_fail,
			"snat_fail":                                 ret.DtDdosUdpPortStats.Stats.Snat_fail,
			"exceed_action_drop":                        ret.DtDdosUdpPortStats.Stats.Exceed_action_drop,
		},
	}
}

func getObjectDdosUdpPortStatsStats(d []interface{}) edpt.DdosUdpPortStatsStats {

	count1 := len(d)
	var ret edpt.DdosUdpPortStatsStats
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
		ret.Src_udp_retry_gap_drop = in["src_udp_retry_gap_drop"].(int)
		ret.Src_udp_auth_drop = in["src_udp_auth_drop"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
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
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Pattern_recognition_sampling_started = in["pattern_recognition_sampling_started"].(int)
		ret.Pattern_recognition_pattern_changed = in["pattern_recognition_pattern_changed"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Src_udp_auth_timeout = in["src_udp_auth_timeout"].(int)
		ret.Src_udp_retry_pass = in["src_udp_retry_pass"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Token_authentication_mismatched = in["token_authentication_mismatched"].(int)
		ret.Token_authentication_invalid = in["token_authentication_invalid"].(int)
		ret.Token_authentication_curr_salt_matched = in["token_authentication_curr_salt_matched"].(int)
		ret.Token_authentication_prev_salt_matched = in["token_authentication_prev_salt_matched"].(int)
		ret.Token_authentication_session_created = in["token_authentication_session_created"].(int)
		ret.Token_authentication_session_created_fail = in["token_authentication_session_created_fail"].(int)
		ret.Snat_fail = in["snat_fail"].(int)
		ret.Exceed_action_drop = in["exceed_action_drop"].(int)
	}
	return ret
}

func dataToEndpointDdosUdpPortStats(d *schema.ResourceData) edpt.DdosUdpPortStats {
	var ret edpt.DdosUdpPortStats

	ret.Stats = getObjectDdosUdpPortStatsStats(d.Get("stats").([]interface{}))
	return ret
}
