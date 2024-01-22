package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwLoggingStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_fw_logging_stats`: Statistics for the object logging\n\n__PLACEHOLDER__",
		ReadContext: resourceFwLoggingStatsRead,

		Schema: map[string]*schema.Schema{
			"gtp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stats": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_type_gtp_invalid_teid": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Invalid TEID",
									},
									"log_gtp_type_reserved_ie_present": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Reserved IE Present",
									},
									"log_type_gtp_mandatory_ie_missing": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Mandatory IE Missing",
									},
									"log_type_gtp_mandatory_ie_inside_grouped_ie_missing": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Mandatory IE Missing Inside Grouped IE",
									},
									"log_type_gtp_msisdn_filtering": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP MSISDN Filtering",
									},
									"log_type_gtp_out_of_order_ie": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Out of Order IE V1",
									},
									"log_type_gtp_out_of_state_ie": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Out of State IE",
									},
									"log_type_enduser_ip_spoofed": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Enduser IP Spoofed",
									},
									"log_type_crosslayer_correlation": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Crosslayer Correlation",
									},
									"log_type_message_not_supported": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Reserved Message Found",
									},
									"log_type_out_of_state": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Out of State Message",
									},
									"log_type_max_msg_length": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Message Length Exceeded Max",
									},
									"log_type_gtp_message_filtering": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Message Filtering",
									},
									"log_type_gtp_apn_filtering": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event Type GTP Apn Filtering",
									},
									"log_type_gtp_rat_type_filtering": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP RAT Type Filtering",
									},
									"log_type_country_code_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Country Code Mismatch",
									},
									"log_type_gtp_in_gtp_filtering": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP in GTP Filtering",
									},
									"log_type_gtp_node_restart": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP SGW/PGW restarted",
									},
									"log_type_gtp_seq_num_mismatch": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Response Sequence number Mismatch",
									},
									"log_type_gtp_rate_limit_periodic": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Rate Limit Periodic",
									},
									"log_type_gtp_invalid_message_length": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Invalid message length across layers",
									},
									"log_type_gtp_hdr_invalid_protocol_flag": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP Protocol flag in header",
									},
									"log_type_gtp_hdr_invalid_spare_bits": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP invalid spare bits in header",
									},
									"log_type_gtp_hdr_invalid_piggy_flag": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event GTP invalid piggyback flag in header",
									},
									"log_type_gtp_invalid_version": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event invalid GTP version",
									},
									"log_type_gtp_invalid_ports": {
										Type: schema.TypeInt, Optional: true, Description: "Log Event mismatch of GTP message and ports",
									},
								},
							},
						},
					},
				},
			},
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"log_message_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Log Packet Sent",
						},
						"log_type_reset": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type Reset",
						},
						"log_type_deny": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type Deny",
						},
						"log_type_session_closed": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type Session Close",
						},
						"log_type_session_opened": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type Session Open",
						},
						"rule_not_logged": {
							Type: schema.TypeInt, Optional: true, Description: "Firewall Rule Not Logged",
						},
						"log_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Log Packets Dropped",
						},
						"tcp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Created",
						},
						"tcp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Session Deleted",
						},
						"udp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session Created",
						},
						"udp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Session Deleted",
						},
						"icmp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Session Deleted",
						},
						"icmp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Session Created",
						},
						"icmpv6_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPV6 Session Deleted",
						},
						"icmpv6_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPV6 Session Created",
						},
						"other_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "Other Session Deleted",
						},
						"other_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Other Session Created",
						},
						"http_request_logged": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Request Logged",
						},
						"http_logging_invalid_format": {
							Type: schema.TypeInt, Optional: true, Description: "HTTP Logging Invalid Format Error",
						},
						"sctp_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Session Created",
						},
						"sctp_session_deleted": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Session Deleted",
						},
						"log_type_sctp_inner_proto_filter": {
							Type: schema.TypeInt, Optional: true, Description: "Log Event Type SCTP Inner Proto Filter",
						},
						"iddos_blackhole_entry_create": {
							Type: schema.TypeInt, Optional: true, Description: "iDDoS IP Entry Created",
						},
						"iddos_blackhole_entry_delete": {
							Type: schema.TypeInt, Optional: true, Description: "iDDoS IP Entry Deleted",
						},
						"session_limit_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Session Limit Exceeded",
						},
					},
				},
			},
		},
	}
}

func resourceFwLoggingStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwLoggingStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwLoggingStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		FwLoggingStatsGtp := setObjectFwLoggingStatsGtp(res)
		d.Set("gtp", FwLoggingStatsGtp)
		FwLoggingStatsStats := setObjectFwLoggingStatsStats(res)
		d.Set("stats", FwLoggingStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectFwLoggingStatsGtp(ret edpt.DataFwLoggingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"stats": setObjectFwLoggingStatsGtpStats(ret.DtFwLoggingStats.Gtp.Stats),
		},
	}
}

func setObjectFwLoggingStatsGtpStats(d edpt.FwLoggingStatsGtpStats) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["log_type_gtp_invalid_teid"] = d.Log_type_gtp_invalid_teid

	in["log_gtp_type_reserved_ie_present"] = d.Log_gtp_type_reserved_ie_present

	in["log_type_gtp_mandatory_ie_missing"] = d.Log_type_gtp_mandatory_ie_missing

	in["log_type_gtp_mandatory_ie_inside_grouped_ie_missing"] = d.Log_type_gtp_mandatory_ie_inside_grouped_ie_missing

	in["log_type_gtp_msisdn_filtering"] = d.Log_type_gtp_msisdn_filtering

	in["log_type_gtp_out_of_order_ie"] = d.Log_type_gtp_out_of_order_ie

	in["log_type_gtp_out_of_state_ie"] = d.Log_type_gtp_out_of_state_ie

	in["log_type_enduser_ip_spoofed"] = d.Log_type_enduser_ip_spoofed

	in["log_type_crosslayer_correlation"] = d.Log_type_crosslayer_correlation

	in["log_type_message_not_supported"] = d.Log_type_message_not_supported

	in["log_type_out_of_state"] = d.Log_type_out_of_state

	in["log_type_max_msg_length"] = d.Log_type_max_msg_length

	in["log_type_gtp_message_filtering"] = d.Log_type_gtp_message_filtering

	in["log_type_gtp_apn_filtering"] = d.Log_type_gtp_apn_filtering

	in["log_type_gtp_rat_type_filtering"] = d.Log_type_gtp_rat_type_filtering

	in["log_type_country_code_mismatch"] = d.Log_type_country_code_mismatch

	in["log_type_gtp_in_gtp_filtering"] = d.Log_type_gtp_in_gtp_filtering

	in["log_type_gtp_node_restart"] = d.Log_type_gtp_node_restart

	in["log_type_gtp_seq_num_mismatch"] = d.Log_type_gtp_seq_num_mismatch

	in["log_type_gtp_rate_limit_periodic"] = d.Log_type_gtp_rate_limit_periodic

	in["log_type_gtp_invalid_message_length"] = d.Log_type_gtp_invalid_message_length

	in["log_type_gtp_hdr_invalid_protocol_flag"] = d.Log_type_gtp_hdr_invalid_protocol_flag

	in["log_type_gtp_hdr_invalid_spare_bits"] = d.Log_type_gtp_hdr_invalid_spare_bits

	in["log_type_gtp_hdr_invalid_piggy_flag"] = d.Log_type_gtp_hdr_invalid_piggy_flag

	in["log_type_gtp_invalid_version"] = d.Log_type_gtp_invalid_version

	in["log_type_gtp_invalid_ports"] = d.Log_type_gtp_invalid_ports
	result = append(result, in)
	return result
}

func setObjectFwLoggingStatsStats(ret edpt.DataFwLoggingStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"log_message_sent":                 ret.DtFwLoggingStats.Stats.Log_message_sent,
			"log_type_reset":                   ret.DtFwLoggingStats.Stats.Log_type_reset,
			"log_type_deny":                    ret.DtFwLoggingStats.Stats.Log_type_deny,
			"log_type_session_closed":          ret.DtFwLoggingStats.Stats.Log_type_session_closed,
			"log_type_session_opened":          ret.DtFwLoggingStats.Stats.Log_type_session_opened,
			"rule_not_logged":                  ret.DtFwLoggingStats.Stats.Rule_not_logged,
			"log_dropped":                      ret.DtFwLoggingStats.Stats.LogDropped,
			"tcp_session_created":              ret.DtFwLoggingStats.Stats.TcpSessionCreated,
			"tcp_session_deleted":              ret.DtFwLoggingStats.Stats.TcpSessionDeleted,
			"udp_session_created":              ret.DtFwLoggingStats.Stats.UdpSessionCreated,
			"udp_session_deleted":              ret.DtFwLoggingStats.Stats.UdpSessionDeleted,
			"icmp_session_deleted":             ret.DtFwLoggingStats.Stats.IcmpSessionDeleted,
			"icmp_session_created":             ret.DtFwLoggingStats.Stats.IcmpSessionCreated,
			"icmpv6_session_deleted":           ret.DtFwLoggingStats.Stats.Icmpv6SessionDeleted,
			"icmpv6_session_created":           ret.DtFwLoggingStats.Stats.Icmpv6SessionCreated,
			"other_session_deleted":            ret.DtFwLoggingStats.Stats.OtherSessionDeleted,
			"other_session_created":            ret.DtFwLoggingStats.Stats.OtherSessionCreated,
			"http_request_logged":              ret.DtFwLoggingStats.Stats.HttpRequestLogged,
			"http_logging_invalid_format":      ret.DtFwLoggingStats.Stats.HttpLoggingInvalidFormat,
			"sctp_session_created":             ret.DtFwLoggingStats.Stats.SctpSessionCreated,
			"sctp_session_deleted":             ret.DtFwLoggingStats.Stats.SctpSessionDeleted,
			"log_type_sctp_inner_proto_filter": ret.DtFwLoggingStats.Stats.Log_type_sctp_inner_proto_filter,
			"iddos_blackhole_entry_create":     ret.DtFwLoggingStats.Stats.IddosBlackholeEntryCreate,
			"iddos_blackhole_entry_delete":     ret.DtFwLoggingStats.Stats.IddosBlackholeEntryDelete,
			"session_limit_exceeded":           ret.DtFwLoggingStats.Stats.SessionLimitExceeded,
		},
	}
}

func getObjectFwLoggingStatsGtp(d []interface{}) edpt.FwLoggingStatsGtp {

	count1 := len(d)
	var ret edpt.FwLoggingStatsGtp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Stats = getObjectFwLoggingStatsGtpStats(in["stats"].([]interface{}))
	}
	return ret
}

func getObjectFwLoggingStatsGtpStats(d []interface{}) edpt.FwLoggingStatsGtpStats {

	count1 := len(d)
	var ret edpt.FwLoggingStatsGtpStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Log_type_gtp_invalid_teid = in["log_type_gtp_invalid_teid"].(int)
		ret.Log_gtp_type_reserved_ie_present = in["log_gtp_type_reserved_ie_present"].(int)
		ret.Log_type_gtp_mandatory_ie_missing = in["log_type_gtp_mandatory_ie_missing"].(int)
		ret.Log_type_gtp_mandatory_ie_inside_grouped_ie_missing = in["log_type_gtp_mandatory_ie_inside_grouped_ie_missing"].(int)
		ret.Log_type_gtp_msisdn_filtering = in["log_type_gtp_msisdn_filtering"].(int)
		ret.Log_type_gtp_out_of_order_ie = in["log_type_gtp_out_of_order_ie"].(int)
		ret.Log_type_gtp_out_of_state_ie = in["log_type_gtp_out_of_state_ie"].(int)
		ret.Log_type_enduser_ip_spoofed = in["log_type_enduser_ip_spoofed"].(int)
		ret.Log_type_crosslayer_correlation = in["log_type_crosslayer_correlation"].(int)
		ret.Log_type_message_not_supported = in["log_type_message_not_supported"].(int)
		ret.Log_type_out_of_state = in["log_type_out_of_state"].(int)
		ret.Log_type_max_msg_length = in["log_type_max_msg_length"].(int)
		ret.Log_type_gtp_message_filtering = in["log_type_gtp_message_filtering"].(int)
		ret.Log_type_gtp_apn_filtering = in["log_type_gtp_apn_filtering"].(int)
		ret.Log_type_gtp_rat_type_filtering = in["log_type_gtp_rat_type_filtering"].(int)
		ret.Log_type_country_code_mismatch = in["log_type_country_code_mismatch"].(int)
		ret.Log_type_gtp_in_gtp_filtering = in["log_type_gtp_in_gtp_filtering"].(int)
		ret.Log_type_gtp_node_restart = in["log_type_gtp_node_restart"].(int)
		ret.Log_type_gtp_seq_num_mismatch = in["log_type_gtp_seq_num_mismatch"].(int)
		ret.Log_type_gtp_rate_limit_periodic = in["log_type_gtp_rate_limit_periodic"].(int)
		ret.Log_type_gtp_invalid_message_length = in["log_type_gtp_invalid_message_length"].(int)
		ret.Log_type_gtp_hdr_invalid_protocol_flag = in["log_type_gtp_hdr_invalid_protocol_flag"].(int)
		ret.Log_type_gtp_hdr_invalid_spare_bits = in["log_type_gtp_hdr_invalid_spare_bits"].(int)
		ret.Log_type_gtp_hdr_invalid_piggy_flag = in["log_type_gtp_hdr_invalid_piggy_flag"].(int)
		ret.Log_type_gtp_invalid_version = in["log_type_gtp_invalid_version"].(int)
		ret.Log_type_gtp_invalid_ports = in["log_type_gtp_invalid_ports"].(int)
	}
	return ret
}

func getObjectFwLoggingStatsStats(d []interface{}) edpt.FwLoggingStatsStats {

	count1 := len(d)
	var ret edpt.FwLoggingStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Log_message_sent = in["log_message_sent"].(int)
		ret.Log_type_reset = in["log_type_reset"].(int)
		ret.Log_type_deny = in["log_type_deny"].(int)
		ret.Log_type_session_closed = in["log_type_session_closed"].(int)
		ret.Log_type_session_opened = in["log_type_session_opened"].(int)
		ret.Rule_not_logged = in["rule_not_logged"].(int)
		ret.LogDropped = in["log_dropped"].(int)
		ret.TcpSessionCreated = in["tcp_session_created"].(int)
		ret.TcpSessionDeleted = in["tcp_session_deleted"].(int)
		ret.UdpSessionCreated = in["udp_session_created"].(int)
		ret.UdpSessionDeleted = in["udp_session_deleted"].(int)
		ret.IcmpSessionDeleted = in["icmp_session_deleted"].(int)
		ret.IcmpSessionCreated = in["icmp_session_created"].(int)
		ret.Icmpv6SessionDeleted = in["icmpv6_session_deleted"].(int)
		ret.Icmpv6SessionCreated = in["icmpv6_session_created"].(int)
		ret.OtherSessionDeleted = in["other_session_deleted"].(int)
		ret.OtherSessionCreated = in["other_session_created"].(int)
		ret.HttpRequestLogged = in["http_request_logged"].(int)
		ret.HttpLoggingInvalidFormat = in["http_logging_invalid_format"].(int)
		ret.SctpSessionCreated = in["sctp_session_created"].(int)
		ret.SctpSessionDeleted = in["sctp_session_deleted"].(int)
		ret.Log_type_sctp_inner_proto_filter = in["log_type_sctp_inner_proto_filter"].(int)
		ret.IddosBlackholeEntryCreate = in["iddos_blackhole_entry_create"].(int)
		ret.IddosBlackholeEntryDelete = in["iddos_blackhole_entry_delete"].(int)
		ret.SessionLimitExceeded = in["session_limit_exceeded"].(int)
	}
	return ret
}

func dataToEndpointFwLoggingStats(d *schema.ResourceData) edpt.FwLoggingStats {
	var ret edpt.FwLoggingStats

	ret.Gtp = getObjectFwLoggingStatsGtp(d.Get("gtp").([]interface{}))

	ret.Stats = getObjectFwLoggingStatsStats(d.Get("stats").([]interface{}))
	return ret
}
