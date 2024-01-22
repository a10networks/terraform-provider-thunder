package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosIpprotoIcmpStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ddos_ipproto_icmp_stats`: Statistics for the object ipproto-icmp\n\n__PLACEHOLDER__",
		ReadContext: resourceDdosIpprotoIcmpStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rate_type0_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Exceeded",
						},
						"rate_type1_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Exceeded",
						},
						"rate_type2_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Exceeded",
						},
						"type_deny_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dropped",
						},
						"icmpv4_rfc_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv4 RFC Undef Type Dropped",
						},
						"icmpv6_rfc_undef_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 RFC Undef Type Dropped",
						},
						"wildcard_deny_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Dropped",
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
						"type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type",
						},
						"type_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Blacklisted",
						},
						"wildcard": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard",
						},
						"wildcard_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Wildcard Blacklisted",
						},
						"rate_type0_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Dropped",
						},
						"rate_type0_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 1 Blacklisted",
						},
						"rate_type1_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Dropped",
						},
						"rate_type1_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 2 Blacklisted",
						},
						"rate_type2_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Dropped",
						},
						"rate_type2_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Dst Rate 3 Blacklisted",
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
						"port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Packet Rate Exceeded",
						},
						"port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded",
						},
						"exceed_drop_prate_src": {
							Type: schema.TypeInt, Optional: true, Description: "Src Pkt Rate Exceeded",
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
						"src_rate_type0_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 1 Exceeded",
						},
						"src_rate_type0_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 1 Dropped",
						},
						"src_rate_type0_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 1 Blacklisted",
						},
						"src_rate_type1_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 2 Exceeded",
						},
						"src_rate_type1_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 2 Dropped",
						},
						"src_rate_type1_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 2 Blacklisted",
						},
						"src_rate_type2_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 3 Exceeded",
						},
						"src_rate_type2_exceed_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 3 Dropped",
						},
						"src_rate_type2_exceed_bl": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Type Src Rate 3 Blacklisted",
						},
						"exceed_drop_brate_src_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Src KiBit Rate Exceeded Count",
						},
						"port_kbit_rate_exceed_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "KiBit Rate Exceeded Count",
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
						"secondary_port_pkt_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port Packet Rate Exceeded",
						},
						"secondary_port_kbit_rate_exceed": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port KiBit Rate Exceeded",
						},
						"secondary_port_kbit_rate_exceed_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Per Addr-Port KiBit Rate Exceeded Count",
						},
						"no_policy_class_list_match": {
							Type: schema.TypeInt, Optional: true, Description: "No Policy Class-list Match",
						},
						"frag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmented Packets Timeout",
						},
						"src_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Src Fragmented Packets Dropped",
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
						"filter_action_blacklist": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Blacklist",
						},
						"filter_action_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Drop",
						},
						"filter_action_default_pass": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Action Default Pass",
						},
						"filter_total_not_match": {
							Type: schema.TypeInt, Optional: true, Description: "Filter Not Matched on Pkt",
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

func resourceDdosIpprotoIcmpStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosIpprotoIcmpStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosIpprotoIcmpStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		DdosIpprotoIcmpStatsStats := setObjectDdosIpprotoIcmpStatsStats(res)
		d.Set("stats", DdosIpprotoIcmpStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectDdosIpprotoIcmpStatsStats(ret edpt.DataDdosIpprotoIcmpStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"rate_type0_exceed":                   ret.DtDdosIpprotoIcmpStats.Stats.Rate_type0_exceed,
			"rate_type1_exceed":                   ret.DtDdosIpprotoIcmpStats.Stats.Rate_type1_exceed,
			"rate_type2_exceed":                   ret.DtDdosIpprotoIcmpStats.Stats.Rate_type2_exceed,
			"type_deny_drop":                      ret.DtDdosIpprotoIcmpStats.Stats.Type_deny_drop,
			"icmpv4_rfc_undef_drop":               ret.DtDdosIpprotoIcmpStats.Stats.Icmpv4_rfc_undef_drop,
			"icmpv6_rfc_undef_drop":               ret.DtDdosIpprotoIcmpStats.Stats.Icmpv6_rfc_undef_drop,
			"wildcard_deny_drop":                  ret.DtDdosIpprotoIcmpStats.Stats.Wildcard_deny_drop,
			"port_rcvd":                           ret.DtDdosIpprotoIcmpStats.Stats.Port_rcvd,
			"port_drop":                           ret.DtDdosIpprotoIcmpStats.Stats.Port_drop,
			"port_pkt_sent":                       ret.DtDdosIpprotoIcmpStats.Stats.Port_pkt_sent,
			"type":                                ret.DtDdosIpprotoIcmpStats.Stats.Type,
			"type_bl":                             ret.DtDdosIpprotoIcmpStats.Stats.Type_bl,
			"wildcard":                            ret.DtDdosIpprotoIcmpStats.Stats.Wildcard,
			"wildcard_bl":                         ret.DtDdosIpprotoIcmpStats.Stats.Wildcard_bl,
			"rate_type0_exceed_drop":              ret.DtDdosIpprotoIcmpStats.Stats.Rate_type0_exceed_drop,
			"rate_type0_exceed_bl":                ret.DtDdosIpprotoIcmpStats.Stats.Rate_type0_exceed_bl,
			"rate_type1_exceed_drop":              ret.DtDdosIpprotoIcmpStats.Stats.Rate_type1_exceed_drop,
			"rate_type1_exceed_bl":                ret.DtDdosIpprotoIcmpStats.Stats.Rate_type1_exceed_bl,
			"rate_type2_exceed_drop":              ret.DtDdosIpprotoIcmpStats.Stats.Rate_type2_exceed_drop,
			"rate_type2_exceed_bl":                ret.DtDdosIpprotoIcmpStats.Stats.Rate_type2_exceed_bl,
			"port_bytes":                          ret.DtDdosIpprotoIcmpStats.Stats.Port_bytes,
			"outbound_port_bytes":                 ret.DtDdosIpprotoIcmpStats.Stats.Outbound_port_bytes,
			"outbound_port_rcvd":                  ret.DtDdosIpprotoIcmpStats.Stats.Outbound_port_rcvd,
			"outbound_port_pkt_sent":              ret.DtDdosIpprotoIcmpStats.Stats.Outbound_port_pkt_sent,
			"port_bytes_sent":                     ret.DtDdosIpprotoIcmpStats.Stats.Port_bytes_sent,
			"port_bytes_drop":                     ret.DtDdosIpprotoIcmpStats.Stats.Port_bytes_drop,
			"port_src_bl":                         ret.DtDdosIpprotoIcmpStats.Stats.Port_src_bl,
			"port_src_escalation":                 ret.DtDdosIpprotoIcmpStats.Stats.Port_src_escalation,
			"current_es_level":                    ret.DtDdosIpprotoIcmpStats.Stats.Current_es_level,
			"port_pkt_rate_exceed":                ret.DtDdosIpprotoIcmpStats.Stats.Port_pkt_rate_exceed,
			"port_kbit_rate_exceed":               ret.DtDdosIpprotoIcmpStats.Stats.Port_kbit_rate_exceed,
			"exceed_drop_prate_src":               ret.DtDdosIpprotoIcmpStats.Stats.Exceed_drop_prate_src,
			"exceed_drop_brate_src":               ret.DtDdosIpprotoIcmpStats.Stats.Exceed_drop_brate_src,
			"outbound_port_bytes_sent":            ret.DtDdosIpprotoIcmpStats.Stats.Outbound_port_bytes_sent,
			"outbound_port_drop":                  ret.DtDdosIpprotoIcmpStats.Stats.Outbound_port_drop,
			"outbound_port_bytes_drop":            ret.DtDdosIpprotoIcmpStats.Stats.Outbound_port_bytes_drop,
			"src_rate_type0_exceed":               ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type0_exceed,
			"src_rate_type0_exceed_drop":          ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type0_exceed_drop,
			"src_rate_type0_exceed_bl":            ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type0_exceed_bl,
			"src_rate_type1_exceed":               ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type1_exceed,
			"src_rate_type1_exceed_drop":          ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type1_exceed_drop,
			"src_rate_type1_exceed_bl":            ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type1_exceed_bl,
			"src_rate_type2_exceed":               ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type2_exceed,
			"src_rate_type2_exceed_drop":          ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type2_exceed_drop,
			"src_rate_type2_exceed_bl":            ret.DtDdosIpprotoIcmpStats.Stats.Src_rate_type2_exceed_bl,
			"exceed_drop_brate_src_pkt":           ret.DtDdosIpprotoIcmpStats.Stats.Exceed_drop_brate_src_pkt,
			"port_kbit_rate_exceed_pkt":           ret.DtDdosIpprotoIcmpStats.Stats.Port_kbit_rate_exceed_pkt,
			"bl":                                  ret.DtDdosIpprotoIcmpStats.Stats.Bl,
			"src_drop":                            ret.DtDdosIpprotoIcmpStats.Stats.Src_drop,
			"frag_rcvd":                           ret.DtDdosIpprotoIcmpStats.Stats.Frag_rcvd,
			"frag_drop":                           ret.DtDdosIpprotoIcmpStats.Stats.Frag_drop,
			"secondary_port_pkt_rate_exceed":      ret.DtDdosIpprotoIcmpStats.Stats.Secondary_port_pkt_rate_exceed,
			"secondary_port_kbit_rate_exceed":     ret.DtDdosIpprotoIcmpStats.Stats.Secondary_port_kbit_rate_exceed,
			"secondary_port_kbit_rate_exceed_pkt": ret.DtDdosIpprotoIcmpStats.Stats.Secondary_port_kbit_rate_exceed_pkt,
			"no_policy_class_list_match":          ret.DtDdosIpprotoIcmpStats.Stats.No_policy_class_list_match,
			"frag_timeout":                        ret.DtDdosIpprotoIcmpStats.Stats.Frag_timeout,
			"src_frag_drop":                       ret.DtDdosIpprotoIcmpStats.Stats.Src_frag_drop,
			"sflow_internal_samples_packed":       ret.DtDdosIpprotoIcmpStats.Stats.Sflow_internal_samples_packed,
			"sflow_external_samples_packed":       ret.DtDdosIpprotoIcmpStats.Stats.Sflow_external_samples_packed,
			"sflow_internal_packets_sent":         ret.DtDdosIpprotoIcmpStats.Stats.Sflow_internal_packets_sent,
			"sflow_external_packets_sent":         ret.DtDdosIpprotoIcmpStats.Stats.Sflow_external_packets_sent,
			"exceed_action_tunnel":                ret.DtDdosIpprotoIcmpStats.Stats.Exceed_action_tunnel,
			"dst_hw_drop":                         ret.DtDdosIpprotoIcmpStats.Stats.Dst_hw_drop,
			"secondary_port_hit":                  ret.DtDdosIpprotoIcmpStats.Stats.Secondary_port_hit,
			"src_zone_service_entry_learned":      ret.DtDdosIpprotoIcmpStats.Stats.Src_zone_service_entry_learned,
			"src_zone_service_entry_aged":         ret.DtDdosIpprotoIcmpStats.Stats.Src_zone_service_entry_aged,
			"dst_hw_drop_inserted":                ret.DtDdosIpprotoIcmpStats.Stats.Dst_hw_drop_inserted,
			"dst_hw_drop_removed":                 ret.DtDdosIpprotoIcmpStats.Stats.Dst_hw_drop_removed,
			"src_hw_drop_inserted":                ret.DtDdosIpprotoIcmpStats.Stats.Src_hw_drop_inserted,
			"src_hw_drop_removed":                 ret.DtDdosIpprotoIcmpStats.Stats.Src_hw_drop_removed,
			"exceed_action_drop":                  ret.DtDdosIpprotoIcmpStats.Stats.Exceed_action_drop,
			"ew_inbound_port_rcv":                 ret.DtDdosIpprotoIcmpStats.Stats.Ew_inbound_port_rcv,
			"ew_inbound_port_drop":                ret.DtDdosIpprotoIcmpStats.Stats.Ew_inbound_port_drop,
			"ew_inbound_port_sent":                ret.DtDdosIpprotoIcmpStats.Stats.Ew_inbound_port_sent,
			"ew_inbound_port_byte_rcv":            ret.DtDdosIpprotoIcmpStats.Stats.Ew_inbound_port_byte_rcv,
			"ew_inbound_port_byte_drop":           ret.DtDdosIpprotoIcmpStats.Stats.Ew_inbound_port_byte_drop,
			"ew_inbound_port_byte_sent":           ret.DtDdosIpprotoIcmpStats.Stats.Ew_inbound_port_byte_sent,
			"ew_outbound_port_rcv":                ret.DtDdosIpprotoIcmpStats.Stats.Ew_outbound_port_rcv,
			"ew_outbound_port_drop":               ret.DtDdosIpprotoIcmpStats.Stats.Ew_outbound_port_drop,
			"ew_outbound_port_sent":               ret.DtDdosIpprotoIcmpStats.Stats.Ew_outbound_port_sent,
			"ew_outbound_port_byte_rcv":           ret.DtDdosIpprotoIcmpStats.Stats.Ew_outbound_port_byte_rcv,
			"ew_outbound_port_byte_sent":          ret.DtDdosIpprotoIcmpStats.Stats.Ew_outbound_port_byte_sent,
			"ew_outbound_port_byte_drop":          ret.DtDdosIpprotoIcmpStats.Stats.Ew_outbound_port_byte_drop,
			"no_route_drop":                       ret.DtDdosIpprotoIcmpStats.Stats.No_route_drop,
			"filter1_match":                       ret.DtDdosIpprotoIcmpStats.Stats.Filter1_match,
			"filter2_match":                       ret.DtDdosIpprotoIcmpStats.Stats.Filter2_match,
			"filter3_match":                       ret.DtDdosIpprotoIcmpStats.Stats.Filter3_match,
			"filter4_match":                       ret.DtDdosIpprotoIcmpStats.Stats.Filter4_match,
			"filter5_match":                       ret.DtDdosIpprotoIcmpStats.Stats.Filter5_match,
			"filter_none_match":                   ret.DtDdosIpprotoIcmpStats.Stats.Filter_none_match,
			"filter_action_blacklist":             ret.DtDdosIpprotoIcmpStats.Stats.Filter_action_blacklist,
			"filter_action_drop":                  ret.DtDdosIpprotoIcmpStats.Stats.Filter_action_drop,
			"filter_action_default_pass":          ret.DtDdosIpprotoIcmpStats.Stats.Filter_action_default_pass,
			"filter_total_not_match":              ret.DtDdosIpprotoIcmpStats.Stats.Filter_total_not_match,
			"src_hw_drop":                         ret.DtDdosIpprotoIcmpStats.Stats.Src_hw_drop,
			"addr_filter_drop":                    ret.DtDdosIpprotoIcmpStats.Stats.Addr_filter_drop,
			"addr_filter_bl":                      ret.DtDdosIpprotoIcmpStats.Stats.Addr_filter_bl,
			"src_learn_overflow":                  ret.DtDdosIpprotoIcmpStats.Stats.Src_learn_overflow,
		},
	}
}

func getObjectDdosIpprotoIcmpStatsStats(d []interface{}) edpt.DdosIpprotoIcmpStatsStats {

	count1 := len(d)
	var ret edpt.DdosIpprotoIcmpStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Rate_type0_exceed = in["rate_type0_exceed"].(int)
		ret.Rate_type1_exceed = in["rate_type1_exceed"].(int)
		ret.Rate_type2_exceed = in["rate_type2_exceed"].(int)
		ret.Type_deny_drop = in["type_deny_drop"].(int)
		ret.Icmpv4_rfc_undef_drop = in["icmpv4_rfc_undef_drop"].(int)
		ret.Icmpv6_rfc_undef_drop = in["icmpv6_rfc_undef_drop"].(int)
		ret.Wildcard_deny_drop = in["wildcard_deny_drop"].(int)
		ret.Port_rcvd = in["port_rcvd"].(int)
		ret.Port_drop = in["port_drop"].(int)
		ret.Port_pkt_sent = in["port_pkt_sent"].(int)
		ret.Type = in["type"].(int)
		ret.Type_bl = in["type_bl"].(int)
		ret.Wildcard = in["wildcard"].(int)
		ret.Wildcard_bl = in["wildcard_bl"].(int)
		ret.Rate_type0_exceed_drop = in["rate_type0_exceed_drop"].(int)
		ret.Rate_type0_exceed_bl = in["rate_type0_exceed_bl"].(int)
		ret.Rate_type1_exceed_drop = in["rate_type1_exceed_drop"].(int)
		ret.Rate_type1_exceed_bl = in["rate_type1_exceed_bl"].(int)
		ret.Rate_type2_exceed_drop = in["rate_type2_exceed_drop"].(int)
		ret.Rate_type2_exceed_bl = in["rate_type2_exceed_bl"].(int)
		ret.Port_bytes = in["port_bytes"].(int)
		ret.Outbound_port_bytes = in["outbound_port_bytes"].(int)
		ret.Outbound_port_rcvd = in["outbound_port_rcvd"].(int)
		ret.Outbound_port_pkt_sent = in["outbound_port_pkt_sent"].(int)
		ret.Port_bytes_sent = in["port_bytes_sent"].(int)
		ret.Port_bytes_drop = in["port_bytes_drop"].(int)
		ret.Port_src_bl = in["port_src_bl"].(int)
		ret.Port_src_escalation = in["port_src_escalation"].(int)
		ret.Current_es_level = in["current_es_level"].(int)
		ret.Port_pkt_rate_exceed = in["port_pkt_rate_exceed"].(int)
		ret.Port_kbit_rate_exceed = in["port_kbit_rate_exceed"].(int)
		ret.Exceed_drop_prate_src = in["exceed_drop_prate_src"].(int)
		ret.Exceed_drop_brate_src = in["exceed_drop_brate_src"].(int)
		ret.Outbound_port_bytes_sent = in["outbound_port_bytes_sent"].(int)
		ret.Outbound_port_drop = in["outbound_port_drop"].(int)
		ret.Outbound_port_bytes_drop = in["outbound_port_bytes_drop"].(int)
		ret.Src_rate_type0_exceed = in["src_rate_type0_exceed"].(int)
		ret.Src_rate_type0_exceed_drop = in["src_rate_type0_exceed_drop"].(int)
		ret.Src_rate_type0_exceed_bl = in["src_rate_type0_exceed_bl"].(int)
		ret.Src_rate_type1_exceed = in["src_rate_type1_exceed"].(int)
		ret.Src_rate_type1_exceed_drop = in["src_rate_type1_exceed_drop"].(int)
		ret.Src_rate_type1_exceed_bl = in["src_rate_type1_exceed_bl"].(int)
		ret.Src_rate_type2_exceed = in["src_rate_type2_exceed"].(int)
		ret.Src_rate_type2_exceed_drop = in["src_rate_type2_exceed_drop"].(int)
		ret.Src_rate_type2_exceed_bl = in["src_rate_type2_exceed_bl"].(int)
		ret.Exceed_drop_brate_src_pkt = in["exceed_drop_brate_src_pkt"].(int)
		ret.Port_kbit_rate_exceed_pkt = in["port_kbit_rate_exceed_pkt"].(int)
		ret.Bl = in["bl"].(int)
		ret.Src_drop = in["src_drop"].(int)
		ret.Frag_rcvd = in["frag_rcvd"].(int)
		ret.Frag_drop = in["frag_drop"].(int)
		ret.Secondary_port_pkt_rate_exceed = in["secondary_port_pkt_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed = in["secondary_port_kbit_rate_exceed"].(int)
		ret.Secondary_port_kbit_rate_exceed_pkt = in["secondary_port_kbit_rate_exceed_pkt"].(int)
		ret.No_policy_class_list_match = in["no_policy_class_list_match"].(int)
		ret.Frag_timeout = in["frag_timeout"].(int)
		ret.Src_frag_drop = in["src_frag_drop"].(int)
		ret.Sflow_internal_samples_packed = in["sflow_internal_samples_packed"].(int)
		ret.Sflow_external_samples_packed = in["sflow_external_samples_packed"].(int)
		ret.Sflow_internal_packets_sent = in["sflow_internal_packets_sent"].(int)
		ret.Sflow_external_packets_sent = in["sflow_external_packets_sent"].(int)
		ret.Exceed_action_tunnel = in["exceed_action_tunnel"].(int)
		ret.Dst_hw_drop = in["dst_hw_drop"].(int)
		ret.Secondary_port_hit = in["secondary_port_hit"].(int)
		ret.Src_zone_service_entry_learned = in["src_zone_service_entry_learned"].(int)
		ret.Src_zone_service_entry_aged = in["src_zone_service_entry_aged"].(int)
		ret.Dst_hw_drop_inserted = in["dst_hw_drop_inserted"].(int)
		ret.Dst_hw_drop_removed = in["dst_hw_drop_removed"].(int)
		ret.Src_hw_drop_inserted = in["src_hw_drop_inserted"].(int)
		ret.Src_hw_drop_removed = in["src_hw_drop_removed"].(int)
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
		ret.Filter1_match = in["filter1_match"].(int)
		ret.Filter2_match = in["filter2_match"].(int)
		ret.Filter3_match = in["filter3_match"].(int)
		ret.Filter4_match = in["filter4_match"].(int)
		ret.Filter5_match = in["filter5_match"].(int)
		ret.Filter_none_match = in["filter_none_match"].(int)
		ret.Filter_action_blacklist = in["filter_action_blacklist"].(int)
		ret.Filter_action_drop = in["filter_action_drop"].(int)
		ret.Filter_action_default_pass = in["filter_action_default_pass"].(int)
		ret.Filter_total_not_match = in["filter_total_not_match"].(int)
		ret.Src_hw_drop = in["src_hw_drop"].(int)
		ret.Addr_filter_drop = in["addr_filter_drop"].(int)
		ret.Addr_filter_bl = in["addr_filter_bl"].(int)
		ret.Src_learn_overflow = in["src_learn_overflow"].(int)
	}
	return ret
}

func dataToEndpointDdosIpprotoIcmpStats(d *schema.ResourceData) edpt.DdosIpprotoIcmpStats {
	var ret edpt.DdosIpprotoIcmpStats

	ret.Stats = getObjectDdosIpprotoIcmpStatsStats(d.Get("stats").([]interface{}))
	return ret
}
