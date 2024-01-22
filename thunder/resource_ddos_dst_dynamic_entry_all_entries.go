package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstDynamicEntryAllEntries() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_dynamic_entry_all_entries`: All Entries\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstDynamicEntryAllEntriesCreate,
		UpdateContext: resourceDdosDstDynamicEntryAllEntriesUpdate,
		ReadContext:   resourceDdosDstDynamicEntryAllEntriesRead,
		DeleteContext: resourceDdosDstDynamicEntryAllEntriesDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'dst_tcp_any_exceed': TCP Dst L4-Type Rate: Total Exceeded; 'dst_tcp_pkt_rate_exceed': TCP Dst L4-Type Rate: Packet Exceeded; 'dst_tcp_conn_rate_exceed': TCP Dst L4-Type Rate: Conn Exceeded; 'dst_udp_any_exceed': UDP Dst L4-Type Rate: Total Exceeded; 'dst_udp_pkt_rate_exceed': UDP Dst L4-Type Rate: Packet Exceeded; 'dst_udp_conn_limit_exceed': UDP Dst L4-Type Limit: Conn Exceeded; 'dst_udp_conn_rate_exceed': UDP Dst L4-Type Rate: Conn Exceeded; 'dst_icmp_pkt_rate_exceed': ICMP Dst Rate: Packet Exceeded; 'dst_other_pkt_rate_exceed': OTHER Dst L4-Type Rate: Packet Exceeded; 'dst_other_frag_pkt_rate_exceed': OTHER Dst L4-Type Rate: Frag Exceeded; 'dst_port_pkt_rate_exceed': Port Rate: Packet Exceeded; 'dst_port_conn_limit_exceed': Port Limit: Conn Exceeded; 'dst_port_conn_rate_exceed': Port Rate: Conn Exceeded; 'dst_pkt_sent': Inbound: Packets Forwarded; 'dst_udp_pkt_sent': UDP Total Packets Forwarded; 'dst_tcp_pkt_sent': TCP Total Packets Forwarded; 'dst_icmp_pkt_sent': ICMP Total Packets Forwarded; 'dst_other_pkt_sent': OTHER Total Packets Forwarded; 'dst_tcp_conn_limit_exceed': TCP Dst L4-Type Limit: Conn Exceeded; 'dst_tcp_pkt_rcvd': TCP Total Packets Received; 'dst_udp_pkt_rcvd': UDP Total Packets Received; 'dst_icmp_pkt_rcvd': ICMP Total Packets Received; 'dst_other_pkt_rcvd': OTHER Total Packets Received; 'dst_udp_filter_match': UDP Filter Match; 'dst_udp_filter_not_match': UDP Filter Not Matched on Pkt; 'dst_udp_filter_action_blacklist': UDP Filter Action Blacklist; 'dst_udp_filter_action_drop': UDP Filter Action Drop; 'dst_tcp_syn': TCP Total SYN Received; 'dst_tcp_syn_drop': TCP SYN Packets Dropped; 'dst_tcp_src_rate_drop': TCP Src Rate: Total Exceeded; 'dst_udp_src_rate_drop': UDP Src Rate: Total Exceeded; 'dst_icmp_src_rate_drop': ICMP Src Rate: Total Exceeded; 'dst_other_frag_src_rate_drop': OTHER Src Rate: Frag Exceeded; 'dst_other_src_rate_drop': OTHER Src Rate: Total Exceeded; 'dst_tcp_drop': TCP Total Packets Dropped; 'dst_udp_drop': UDP Total Packets Dropped; 'dst_icmp_drop': ICMP Total Packets Dropped; 'dst_frag_drop': Fragmented Packets Dropped; 'dst_other_drop': OTHER Total Packets Dropped; 'dst_tcp_auth': TCP Auth: SYN Cookie Sent; 'dst_udp_filter_action_default_pass': UDP Filter Action Default Pass; 'dst_tcp_filter_match': TCP Filter Match; 'dst_tcp_filter_not_match': TCP Filter Not Matched on Pkt; 'dst_tcp_filter_action_blacklist': TCP Filter Action Blacklist; 'dst_tcp_filter_action_drop': TCP Filter Action Drop; 'dst_tcp_filter_action_default_pass': TCP Filter Action Default Pass; 'dst_udp_filter_action_whitelist': UDP Filter Action WL; 'dst_over_limit_on': DST overlimit Trigger ON; 'dst_over_limit_off': DST overlimit Trigger OFF; 'dst_port_over_limit_on': DST port overlimit Trigger ON; 'dst_port_over_limit_off': DST port overlimit Trigger OFF; 'dst_over_limit_action': DST overlimit action; 'dst_port_over_limit_action': DST port overlimit action; 'scanning_detected_drop': Scanning Detected drop (deprecated); 'scanning_detected_blacklist': Scanning Detected blacklist (deprecated); 'dst_udp_kibit_rate_drop': UDP Dst L4-Type Rate: KiBit Exceeded; 'dst_tcp_kibit_rate_drop': TCP Dst L4-Type Rate: KiBit Exceeded; 'dst_icmp_kibit_rate_drop': ICMP Dst Rate: KiBit Exceeded; 'dst_other_kibit_rate_drop': OTHER Dst L4-Type Rate: KiBit Exceeded; 'dst_port_undef_drop': Dst Port Undefined Dropped; 'dst_port_bl': Dst Port Blacklist Packets Dropped; 'dst_src_port_bl': Dst SrcPort Blacklist Packets Dropped; 'dst_port_kbit_rate_exceed': Port Rate: KiBit Exceeded; 'dst_tcp_src_drop': TCP Src Packets Dropped; 'dst_udp_src_drop': UDP Src Packets Dropped; 'dst_icmp_src_drop': ICMP Src Packets Dropped; 'dst_other_src_drop': OTHER Src Packets Dropped; 'tcp_syn_rcvd': TCP Inbound SYN Received; 'tcp_syn_ack_rcvd': TCP SYN ACK Received; 'tcp_ack_rcvd': TCP ACK Received; 'tcp_fin_rcvd': TCP FIN Received; 'tcp_rst_rcvd': TCP RST Received; 'ingress_bytes': Inbound: Bytes Received; 'egress_bytes': Outbound: Bytes Received; 'ingress_packets': Inbound: Packets Received; 'egress_packets': Outbound: Packets Received; 'tcp_fwd_recv': TCP Inbound Packets Received; 'udp_fwd_recv': UDP Inbound Packets Received; 'icmp_fwd_recv': ICMP Inbound Packets Received; 'tcp_syn_cookie_fail': TCP Auth: SYN Cookie Failed; 'dst_tcp_session_created': TCP Sessions Created; 'dst_udp_session_created': UDP Sessions Created; 'dst_tcp_filter_action_whitelist': TCP Filter Action WL; 'dst_other_filter_match': OTHER Filter Match; 'dst_other_filter_not_match': OTHER Filter Not Matched on Pkt; 'dst_other_filter_action_blacklist': OTHER Filter Action Blacklist; 'dst_other_filter_action_drop': OTHER Filter Action Drop; 'dst_other_filter_action_whitelist': OTHER Filter Action WL; 'dst_other_filter_action_default_pass': OTHER Filter Action Default Pass; 'dst_blackhole_inject': Dst Blackhole Inject; 'dst_blackhole_withdraw': Dst Blackhole Withdraw; 'dst_tcp_out_of_seq_excd': TCP Out-Of-Seq Exceeded; 'dst_tcp_retransmit_excd': TCP Retransmit Exceeded; 'dst_tcp_zero_window_excd': TCP Zero-Window Exceeded; 'dst_tcp_conn_prate_excd': TCP Rate: Conn Pkt Exceeded; 'dst_tcp_action_on_ack_init': TCP Auth: ACK Retry Init; 'dst_tcp_action_on_ack_gap_drop': TCP Auth: ACK Retry Retry-Gap Dropped; 'dst_tcp_action_on_ack_fail': TCP Auth: ACK Retry Dropped; 'dst_tcp_action_on_ack_pass': TCP Auth: ACK Retry Passed; 'dst_tcp_action_on_syn_init': TCP Auth: SYN Retry Init; 'dst_tcp_action_on_syn_gap_drop': TCP Auth: SYN Retry-Gap Dropped; 'dst_tcp_action_on_syn_fail': TCP Auth: SYN Retry Dropped; 'dst_tcp_action_on_syn_pass': TCP Auth: SYN Retry Passed; 'udp_payload_too_small': UDP Payload Too Small; 'udp_payload_too_big': UDP Payload Too Large; 'dst_udp_conn_prate_excd': UDP Rate: Conn Pkt Exceeded; 'dst_udp_ntp_monlist_req': UDP NTP Monlist Request; 'dst_udp_ntp_monlist_resp': UDP NTP Monlist Response; 'dst_udp_wellknown_sport_drop': UDP SrcPort Wellknown; 'dst_udp_retry_init': UDP Auth: Retry Init; 'dst_udp_retry_pass': UDP Auth: Retry Passed; 'dst_tcp_bytes_drop': TCP Total Bytes Dropped; 'dst_udp_bytes_drop': UDP Total Bytes Dropped; 'dst_icmp_bytes_drop': ICMP Total Bytes Dropped; 'dst_other_bytes_drop': OTHER Total Bytes Dropped; 'dst_out_no_route': Dst IPv4/v6 Out No Route; 'outbound_bytes_sent': Outbound: Bytes Forwarded; 'outbound_pkt_drop': Outbound: Packets Dropped; 'outbound_bytes_drop': Outbound: Bytes Dropped; 'outbound_pkt_sent': Outbound: Packets Forwarded; 'inbound_bytes_sent': Inbound: Bytes Forwarded; 'inbound_bytes_drop': Inbound: Bytes Dropped; 'dst_src_port_pkt_rate_exceed': SrcPort Rate: Packet Exceeded; 'dst_src_port_kbit_rate_exceed': SrcPort Rate: KiBit Exceeded; 'dst_src_port_conn_limit_exceed': SrcPort Limit: Conn Exceeded; 'dst_src_port_conn_rate_exceed': SrcPort Rate: Conn Exceeded; 'dst_ip_proto_pkt_rate_exceed': IP-Proto Rate: Packet Exceeded; 'dst_ip_proto_kbit_rate_exceed': IP-Proto Rate: KiBit Exceeded; 'dst_tcp_port_any_exceed': TCP Port Rate: Total Exceed; 'dst_udp_port_any_exceed': UDP Port Rate: Total Exceed; 'dst_tcp_auth_pass': TCP Auth: SYN Auth Passed; 'dst_tcp_rst_cookie_fail': TCP Auth: RST Cookie Failed; 'dst_tcp_unauth_drop': TCP Auth: Unauth Dropped; 'src_tcp_syn_auth_fail': Src TCP Auth: SYN Auth Failed; 'src_tcp_syn_cookie_sent': Src TCP Auth: SYN Cookie Sent; 'src_tcp_syn_cookie_fail': Src TCP Auth: SYN Cookie Failed; 'src_tcp_rst_cookie_fail': Src TCP Auth: RST Cookie Failed; 'src_tcp_unauth_drop': Src TCP Auth: Unauth Dropped; 'src_tcp_action_on_syn_init': Src TCP Auth: SYN Retry Init;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'src_tcp_action_on_syn_gap_drop': Src TCP Auth: SYN Retry-Gap Dropped; 'src_tcp_action_on_syn_fail': Src TCP Auth: SYN Retry Dropped; 'src_tcp_action_on_ack_init': Src TCP Auth: ACK Retry Init; 'src_tcp_action_on_ack_gap_drop': Src TCP Auth: ACK Retry Retry-Gap Dropped; 'src_tcp_action_on_ack_fail': Src TCP Auth: ACK Retry Dropped; 'src_tcp_out_of_seq_excd': Src TCP Out-Of-Seq Exceeded; 'src_tcp_retransmit_excd': Src TCP Retransmit Exceeded; 'src_tcp_zero_window_excd': Src TCP Zero-Window Exceeded; 'src_tcp_conn_prate_excd': Src TCP Rate: Conn Pkt Exceeded; 'src_udp_min_payload': Src UDP Payload Too Small; 'src_udp_max_payload': Src UDP Payload Too Large; 'src_udp_conn_prate_excd': Src UDP Rate: Conn Pkt Exceeded; 'src_udp_ntp_monlist_req': Src UDP NTP Monlist Request; 'src_udp_ntp_monlist_resp': Src UDP NTP Monlist Response; 'src_udp_wellknown_sport_drop': Src UDP SrcPort Wellknown; 'src_udp_retry_init': Src UDP Auth: Retry Init; 'dst_udp_retry_gap_drop': UDP Auth: Retry-Gap Dropped; 'dst_udp_retry_fail': UDP P Sessions Aged; 'dst_tcp_session_aged': TCP Sessions Aged; 'dst_udp_session_aged': UDP Sessions Aged; 'dst_tcp_conn_close': TCP Connections Closed; 'dst_tcp_conn_close_half_open': TCP Half Open Connections Closed; 'dst_l4_tcp_auth': TCP Dst L4-Type Auth: SYN Cookie Sent; 'tcp_l4_syn_cookie_fail': TCP Dst L4-Type Auth: SYN Cookie Failed; 'tcp_l4_rst_cookie_fail': TCP Dst L4-Type Auth: RST Cookie Failed; 'tcp_l4_unauth_drop': TCP Dst L4-Type Auth: Unauth Dropped; 'dst_drop_frag_pkt': Dst Fragmented Packets Dropped; 'src_tcp_filter_action_blacklist': Src TCP Filter Action Blacklist; 'src_tcp_filter_action_whitelist': Src TCP Filter Action WL; 'src_tcp_filter_action_drop': Src TCP Filter Action Drop; 'src_tcp_filter_action_default_pass': Src TCP Filter Action Default Pass; 'src_udp_filter_action_blacklist': Src UDP Filter Action Blacklist; 'src_udp_filter_action_whitelist': Src UDP Filter Action WL; 'src_udp_filter_action_drop': Src UDP Filter Action Drop; 'src_udp_filter_action_default_pass': Src UDP Filter Action Default Pass; 'src_other_filter_action_blacklist': Src OTHER Filter Action Blacklist; 'src_other_filter_action_whitelist': Src OTHER Filter Action WL; 'src_other_filter_action_drop': Src OTHER Filter Action Drop; 'src_other_filter_action_default_pass': Src OTHER Filter Action Default Pass; 'tcp_invalid_syn': TCP Invalid SYN Received; 'dst_tcp_conn_close_w_rst': TCP RST Connections Closed; 'dst_tcp_conn_close_w_fin': TCP FIN Connections Closed; 'dst_tcp_conn_close_w_idle': TCP Idle Connections Closed; 'dst_tcp_conn_create_from_syn': TCP Connections Created From SYN; 'dst_tcp_conn_create_from_ack': TCP Connections Created From ACK; 'src_frag_drop': Src Fragmented Packets Dropped; 'dst_l4_tcp_blacklist_drop': Dst L4-type TCP Blacklist Dropped; 'dst_l4_udp_blacklist_drop': Dst L4-type UDP Blacklist Dropped; 'dst_l4_icmp_blacklist_drop': No Policy Class-list Match; 'dst_l4_other_blacklist_drop': Dst L4-type OTHER Blacklist Dropped; 'src_l4_tcp_blacklist_drop': Src L4-type TCP Blacklist Dropped; 'src_l4_udp_blacklist_drop': Src L4-type UDP Blacklist Dropped; 'src_l4_icmp_blacklist_drop': Src L4-type ICMP Blacklist Dropped; 'src_l4_other_blacklist_drop': Src L4-type OTHER Blacklist Dropped; 'drop_frag_timeout_drop': Fragment Reassemble Timeout Drop; 'dst_port_kbit_rate_exceed_pkt': Port Rate: KiBit Pkt Exceeded; 'dst_tcp_bytes_rcv': TCP Total Bytes Received; 'dst_udp_bytes_rcv': UDP Total Bytes Received; 'dst_icmp_bytes_rcv': ICMP Total Bytes Received; 'dst_other_bytes_rcv': OTHER Total Bytes Received; 'dst_tcp_bytes_sent': TCP Total Bytes Forwarded; 'dst_udp_bytes_sent': UDP Total Bytes Forwarded; 'dst_icmp_bytes_sent': ICMP Total Bytes Forwarded; 'dst_other_bytes_sent': OTHER Total Bytes Forwarded; 'dst_udp_auth_drop': UDP Auth: Dropped; 'dst_tcp_auth_drop': TCP Auth: Dropped; 'dst_tcp_auth_resp': TCP Auth: Responded; 'inbound_pkt_drop': Inbound: Packets Dropped; 'dst_entry_pkt_rate_exceed': Entry Rate: Packet Exceeded; 'dst_entry_kbit_rate_exceed': Entry Rate: KiBit Exceeded; 'dst_entry_conn_limit_exceed': Entry Limit: Conn Exceeded; 'dst_entry_conn_rate_exceed': Entry Rate: Conn Exceeded; 'dst_entry_frag_pkt_rate_exceed': Entry Rate: Frag Packet Exceeded; 'dst_icmp_any_exceed': ICMP Rate: Total Exceed; 'dst_other_any_exceed': OTHER Rate: Total Exceed; 'src_dst_pair_entry_total': Src-Dst Pair Entry Total Count; 'src_dst_pair_entry_udp': Src-Dst Pair Entry UDP Count; 'src_dst_pair_entry_tcp': Src-Dst Pair Entry TCP Count; 'src_dst_pair_entry_icmp': Src-Dst Pair Entry ICMP Count; 'src_dst_pair_entry_other': Src-Dst Pair Entry OTHER Count; 'dst_clist_overflow_policy_at_learning': Dst Src-Based Overflow Policy Hit; 'tcp_rexmit_syn_limit_drop': TCP SYN Retransmit Exceeded Drop; 'tcp_rexmit_syn_limit_bl': TCP SYN Retransmit Exceeded Blacklist; 'dst_tcp_wellknown_sport_drop': TCP SrcPort Wellknown; 'src_tcp_wellknown_sport_drop': Src TCP SrcPort Wellknown; 'dst_frag_rcvd': Fragmented Packets Received; 'no_policy_class_list_match': No Policy Class-list Match; 'src_udp_retry_gap_drop': Src UDP Auth: Retry-Gap Dropped; 'dst_entry_kbit_rate_exceed_count': Entry Rate: KiBit Exceeded Count; 'dst_port_undef_hit': Dst Port Undefined Hit; 'dst_tcp_action_on_ack_timeout': TCP Auth: ACK Retry Timeout; 'dst_tcp_action_on_ack_reset': TCP Auth: ACK Retry Timeout Reset; 'dst_tcp_action_on_ack_blacklist': TCP Auth: ACK Retry Timeout Blacklisted; 'src_tcp_action_on_ack_timeout': Src TCP Auth: ACK Retry Timeout; 'src_tcp_action_on_ack_reset': Src TCP Auth: ACK Retry Timeout Reset; 'src_tcp_action_on_ack_blacklist': Src TCP Auth: ACK Retry Timeout Blacklisted; 'dst_tcp_action_on_syn_timeout': TCP Auth: SYN Retry Timeout; 'dst_tcp_action_on_syn_reset': TCP Auth: SYN Retry Timeout Reset; 'dst_tcp_action_on_syn_blacklist': TCP Auth: SYN Retry Timeout Blacklisted; 'src_tcp_action_on_syn_timeout': Src TCP Auth: SYN Retry Timeout; 'src_tcp_action_on_syn_reset': Src TCP Auth: SYN Retry Timeout Reset; 'src_tcp_action_on_syn_blacklist': Src TCP Auth: SYN Retry Timeout Blacklisted; 'dst_udp_frag_pkt_rate_exceed': UDP Dst L4-Type Rate: Frag Exceeded; 'dst_udp_frag_src_rate_drop': UDP Src Rate: Frag Exceeded; 'dst_tcp_frag_pkt_rate_exceed': TCP Dst L4-Type Rate: Frag Exceeded; 'dst_tcp_frag_src_rate_drop': TCP Src Rate: Frag Exceeded; 'dst_icmp_frag_pkt_rate_exceed': ICMP Dst L4-Type Rate: Frag Exceeded; 'dst_icmp_frag_src_rate_drop': ICMP Src Rate: Frag Exceeded; 'sflow_internal_samples_packed': Sflow Internal Samples Packed; 'sflow_external_samples_packed': Sflow External Samples Packed; 'sflow_internal_packets_sent': Sflow Internal Packets Sent; 'sflow_external_packets_sent': Sflow External Packets Sent; 'dns_outbound_total_query': DNS Outbound Total Query; 'dns_outbound_query_malformed': DNS Outbound Query Malformed; 'dns_outbound_query_resp_chk_failed': DNS Outbound Query Resp Check Failed; 'dns_outbound_query_resp_chk_blacklisted': DNS Outbound Query Resp Check Blacklisted; 'dns_outbound_query_resp_chk_refused_sent': DNS Outbound Query Resp Check REFUSED Sent; 'dns_outbound_query_resp_chk_reset_sent': DNS Outbound Query Resp Check RESET Sent; 'dns_outbound_query_resp_chk_no_resp_sent': DNS Outbound Query Resp Check No Response Sent; 'dns_outbound_query_resp_size_exceed': DNS Outbound Query Response Size Exceed; 'dns_outbound_query_sess_timed_out': DNS Outbound Query Session Timed Out; 'dst_exceed_action_tunnel': Entry Exceed Action: Tunnel; 'src_udp_auth_timeout': Src UDP Auth: Retry Timeout; 'src_udp_retry_pass': Src UDP Retry Passed;",
						},
						"counters3": {
							Type: schema.TypeString, Optional: true, Description: "'dst_hw_drop_rule_insert': Dst Hardware Drop Rules Inserted; 'dst_hw_drop_rule_remove': Dst Hardware Drop Rules Removed; 'src_hw_drop_rule_insert': Src Hardware Drop Rules Inserted; 'src_hw_drop_rule_remove': Src Hardware Drop Rules Removed; 'prog_first_req_time_exceed': Req-Resp: First Request Time Exceed; 'prog_req_resp_time_exceed': Req-Resp: Request to Response Time Exceed; 'prog_request_len_exceed': Req-Resp: Request Length Exceed; 'prog_response_len_exceed': Req-Resp: Response Length Exceed; 'prog_resp_req_ratio_exceed': Req-Resp: Response to Request Ratio Exceed; 'prog_resp_req_time_exceed': Req-Resp: Response to Request Time Exceed; 'entry_sync_message_received': Entry Sync Message Received; 'entry_sync_message_sent': Entry Sync Message Sent; 'prog_conn_sent_exceed': Connection: Sent Exceed; 'prog_conn_rcvd_exceed': Connection: Received Exceed; 'prog_conn_time_exceed': Connection: Time Exceed; 'prog_conn_rcvd_sent_ratio_exceed': Connection: Reveived to Sent Ratio Exceed; 'prog_win_sent_exceed': Time Window: Sent Exceed; 'prog_win_rcvd_exceed': Time Window: Received Exceed; 'prog_win_rcvd_sent_ratio_exceed': Time Window: Received to Sent Exceed; 'prog_exceed_drop': Req-Resp: Violation Exceed Dropped; 'prog_exceed_bl': Req-Resp: Violation Exceed Blacklisted; 'prog_conn_exceed_drop': Connection: Violation Exceed Dropped; 'prog_conn_exceed_bl': Connection: Violation Exceed Blacklisted; 'prog_win_exceed_drop': Time Window: Violation Exceed Dropped; 'prog_win_exceed_bl': Time Window: Violation Exceed Blacklisted; 'dst_exceed_action_drop': Entry Exceed Action: Dropped; 'src_hw_drop': Src Hardware Packets Dropped; 'dst_tcp_auth_rst': TCP Auth: Reset; 'dst_src_learn_overflow': Src Dynamic Entry Count Overflow; 'tcp_fwd_sent': TCP Inbound Packets Forwarded; 'udp_fwd_sent': UDP Inbound Packets Forwarded;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDstDynamicEntryAllEntriesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryAllEntriesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryAllEntries(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDynamicEntryAllEntriesRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstDynamicEntryAllEntriesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryAllEntriesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryAllEntries(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstDynamicEntryAllEntriesRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstDynamicEntryAllEntriesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryAllEntriesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryAllEntries(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstDynamicEntryAllEntriesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstDynamicEntryAllEntriesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstDynamicEntryAllEntries(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstDynamicEntryAllEntriesSamplingEnable(d []interface{}) []edpt.DdosDstDynamicEntryAllEntriesSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.DdosDstDynamicEntryAllEntriesSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstDynamicEntryAllEntriesSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		oi.Counters3 = in["counters3"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstDynamicEntryAllEntries(d *schema.ResourceData) edpt.DdosDstDynamicEntryAllEntries {
	var ret edpt.DdosDstDynamicEntryAllEntries
	ret.Inst.SamplingEnable = getSliceDdosDstDynamicEntryAllEntriesSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
