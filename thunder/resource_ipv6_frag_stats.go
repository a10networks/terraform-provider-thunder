package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6FragStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_ipv6_frag_stats`: Statistics for the object frag\n\n__PLACEHOLDER__",
		ReadContext: resourceIpv6FragStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Session Inserted",
						},
						"session_expired": {
							Type: schema.TypeInt, Optional: true, Description: "Session Expired",
						},
						"icmp_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Received",
						},
						"icmpv6_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 Received",
						},
						"udp_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Received",
						},
						"tcp_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Received",
						},
						"ipip_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "IP-in-IP Received",
						},
						"ipv6ip_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6-in-IP Received",
						},
						"other_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "Other Received",
						},
						"icmp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP Dropped",
						},
						"icmpv6_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 Dropped",
						},
						"udp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Dropped",
						},
						"tcp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Dropped",
						},
						"ipip_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IP-in-IP Dropped",
						},
						"ipv6ip_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6-in-IP Dropped",
						},
						"other_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "Other Dropped",
						},
						"overlap_error": {
							Type: schema.TypeInt, Optional: true, Description: "Overlapping Fragment Dropped",
						},
						"bad_ip_len": {
							Type: schema.TypeInt, Optional: true, Description: "Bad IP Length",
						},
						"too_small": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Too Small Drop",
						},
						"first_tcp_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "First TCP Fragment Too Small Drop",
						},
						"first_l4_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "First L4 Fragment Too Small Drop",
						},
						"total_sessions_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Total Sessions Exceeded Drop",
						},
						"no_session_memory": {
							Type: schema.TypeInt, Optional: true, Description: "Out of Session Memory",
						},
						"fast_aging_set": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmentation Fast Aging Set",
						},
						"fast_aging_unset": {
							Type: schema.TypeInt, Optional: true, Description: "Fragmentation Fast Aging Unset",
						},
						"fragment_queue_success": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Queue Success",
						},
						"unaligned_len": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Length Unaligned",
						},
						"exceeded_len": {
							Type: schema.TypeInt, Optional: true, Description: "Payload Length Out of Bounds",
						},
						"duplicate_first_frag": {
							Type: schema.TypeInt, Optional: true, Description: "Duplicate First Fragment",
						},
						"duplicate_last_frag": {
							Type: schema.TypeInt, Optional: true, Description: "Duplicate Last Fragment",
						},
						"total_fragments_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Total Queued Fragments Exceeded",
						},
						"fragment_queue_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Queue Failure",
						},
						"reassembly_success": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Reassembly Success",
						},
						"max_len_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Max Data Length Exceeded",
						},
						"reassembly_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Reassembly Failure",
						},
						"policy_drop": {
							Type: schema.TypeInt, Optional: true, Description: "MTU Exceeded Policy Drop",
						},
						"error_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fragment Processing Drop",
						},
						"high_cpu_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "High CPU Threshold Reached",
						},
						"low_cpu_threshold": {
							Type: schema.TypeInt, Optional: true, Description: "Low CPU Threshold Reached",
						},
						"cpu_threshold_drop": {
							Type: schema.TypeInt, Optional: true, Description: "High CPU Drop",
						},
						"ipd_entry_drop": {
							Type: schema.TypeInt, Optional: true, Description: "DDoS Protection Drop",
						},
						"max_packets_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Too Many Packets Per Reassembly Drop",
						},
						"session_packets_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Session Max Packets Exceeded",
						},
						"sctp_rcv": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Received",
						},
						"sctp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "SCTP Dropped",
						},
						"first_gtp_packet_too_small": {
							Type: schema.TypeInt, Optional: true, Description: "First GTP Fragment Too Small Drop",
						},
					},
				},
			},
		},
	}
}

func resourceIpv6FragStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6FragStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6FragStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Ipv6FragStatsStats := setObjectIpv6FragStatsStats(res)
		d.Set("stats", Ipv6FragStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectIpv6FragStatsStats(ret edpt.DataIpv6FragStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"session_inserted":           ret.DtIpv6FragStats.Stats.SessionInserted,
			"session_expired":            ret.DtIpv6FragStats.Stats.SessionExpired,
			"icmp_rcv":                   ret.DtIpv6FragStats.Stats.IcmpRcv,
			"icmpv6_rcv":                 ret.DtIpv6FragStats.Stats.Icmpv6Rcv,
			"udp_rcv":                    ret.DtIpv6FragStats.Stats.UdpRcv,
			"tcp_rcv":                    ret.DtIpv6FragStats.Stats.TcpRcv,
			"ipip_rcv":                   ret.DtIpv6FragStats.Stats.IpipRcv,
			"ipv6ip_rcv":                 ret.DtIpv6FragStats.Stats.Ipv6ipRcv,
			"other_rcv":                  ret.DtIpv6FragStats.Stats.OtherRcv,
			"icmp_dropped":               ret.DtIpv6FragStats.Stats.IcmpDropped,
			"icmpv6_dropped":             ret.DtIpv6FragStats.Stats.Icmpv6Dropped,
			"udp_dropped":                ret.DtIpv6FragStats.Stats.UdpDropped,
			"tcp_dropped":                ret.DtIpv6FragStats.Stats.TcpDropped,
			"ipip_dropped":               ret.DtIpv6FragStats.Stats.IpipDropped,
			"ipv6ip_dropped":             ret.DtIpv6FragStats.Stats.Ipv6ipDropped,
			"other_dropped":              ret.DtIpv6FragStats.Stats.OtherDropped,
			"overlap_error":              ret.DtIpv6FragStats.Stats.OverlapError,
			"bad_ip_len":                 ret.DtIpv6FragStats.Stats.BadIpLen,
			"too_small":                  ret.DtIpv6FragStats.Stats.TooSmall,
			"first_tcp_too_small":        ret.DtIpv6FragStats.Stats.FirstTcpTooSmall,
			"first_l4_too_small":         ret.DtIpv6FragStats.Stats.FirstL4TooSmall,
			"total_sessions_exceeded":    ret.DtIpv6FragStats.Stats.TotalSessionsExceeded,
			"no_session_memory":          ret.DtIpv6FragStats.Stats.NoSessionMemory,
			"fast_aging_set":             ret.DtIpv6FragStats.Stats.FastAgingSet,
			"fast_aging_unset":           ret.DtIpv6FragStats.Stats.FastAgingUnset,
			"fragment_queue_success":     ret.DtIpv6FragStats.Stats.FragmentQueueSuccess,
			"unaligned_len":              ret.DtIpv6FragStats.Stats.UnalignedLen,
			"exceeded_len":               ret.DtIpv6FragStats.Stats.ExceededLen,
			"duplicate_first_frag":       ret.DtIpv6FragStats.Stats.DuplicateFirstFrag,
			"duplicate_last_frag":        ret.DtIpv6FragStats.Stats.DuplicateLastFrag,
			"total_fragments_exceeded":   ret.DtIpv6FragStats.Stats.TotalFragmentsExceeded,
			"fragment_queue_failure":     ret.DtIpv6FragStats.Stats.FragmentQueueFailure,
			"reassembly_success":         ret.DtIpv6FragStats.Stats.ReassemblySuccess,
			"max_len_exceeded":           ret.DtIpv6FragStats.Stats.MaxLenExceeded,
			"reassembly_failure":         ret.DtIpv6FragStats.Stats.ReassemblyFailure,
			"policy_drop":                ret.DtIpv6FragStats.Stats.PolicyDrop,
			"error_drop":                 ret.DtIpv6FragStats.Stats.ErrorDrop,
			"high_cpu_threshold":         ret.DtIpv6FragStats.Stats.HighCpuThreshold,
			"low_cpu_threshold":          ret.DtIpv6FragStats.Stats.LowCpuThreshold,
			"cpu_threshold_drop":         ret.DtIpv6FragStats.Stats.CpuThresholdDrop,
			"ipd_entry_drop":             ret.DtIpv6FragStats.Stats.IpdEntryDrop,
			"max_packets_exceeded":       ret.DtIpv6FragStats.Stats.MaxPacketsExceeded,
			"session_packets_exceeded":   ret.DtIpv6FragStats.Stats.SessionPacketsExceeded,
			"sctp_rcv":                   ret.DtIpv6FragStats.Stats.SctpRcv,
			"sctp_dropped":               ret.DtIpv6FragStats.Stats.SctpDropped,
			"first_gtp_packet_too_small": ret.DtIpv6FragStats.Stats.FirstGtpPacketTooSmall,
		},
	}
}

func getObjectIpv6FragStatsStats(d []interface{}) edpt.Ipv6FragStatsStats {

	count1 := len(d)
	var ret edpt.Ipv6FragStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionInserted = in["session_inserted"].(int)
		ret.SessionExpired = in["session_expired"].(int)
		ret.IcmpRcv = in["icmp_rcv"].(int)
		ret.Icmpv6Rcv = in["icmpv6_rcv"].(int)
		ret.UdpRcv = in["udp_rcv"].(int)
		ret.TcpRcv = in["tcp_rcv"].(int)
		ret.IpipRcv = in["ipip_rcv"].(int)
		ret.Ipv6ipRcv = in["ipv6ip_rcv"].(int)
		ret.OtherRcv = in["other_rcv"].(int)
		ret.IcmpDropped = in["icmp_dropped"].(int)
		ret.Icmpv6Dropped = in["icmpv6_dropped"].(int)
		ret.UdpDropped = in["udp_dropped"].(int)
		ret.TcpDropped = in["tcp_dropped"].(int)
		ret.IpipDropped = in["ipip_dropped"].(int)
		ret.Ipv6ipDropped = in["ipv6ip_dropped"].(int)
		ret.OtherDropped = in["other_dropped"].(int)
		ret.OverlapError = in["overlap_error"].(int)
		ret.BadIpLen = in["bad_ip_len"].(int)
		ret.TooSmall = in["too_small"].(int)
		ret.FirstTcpTooSmall = in["first_tcp_too_small"].(int)
		ret.FirstL4TooSmall = in["first_l4_too_small"].(int)
		ret.TotalSessionsExceeded = in["total_sessions_exceeded"].(int)
		ret.NoSessionMemory = in["no_session_memory"].(int)
		ret.FastAgingSet = in["fast_aging_set"].(int)
		ret.FastAgingUnset = in["fast_aging_unset"].(int)
		ret.FragmentQueueSuccess = in["fragment_queue_success"].(int)
		ret.UnalignedLen = in["unaligned_len"].(int)
		ret.ExceededLen = in["exceeded_len"].(int)
		ret.DuplicateFirstFrag = in["duplicate_first_frag"].(int)
		ret.DuplicateLastFrag = in["duplicate_last_frag"].(int)
		ret.TotalFragmentsExceeded = in["total_fragments_exceeded"].(int)
		ret.FragmentQueueFailure = in["fragment_queue_failure"].(int)
		ret.ReassemblySuccess = in["reassembly_success"].(int)
		ret.MaxLenExceeded = in["max_len_exceeded"].(int)
		ret.ReassemblyFailure = in["reassembly_failure"].(int)
		ret.PolicyDrop = in["policy_drop"].(int)
		ret.ErrorDrop = in["error_drop"].(int)
		ret.HighCpuThreshold = in["high_cpu_threshold"].(int)
		ret.LowCpuThreshold = in["low_cpu_threshold"].(int)
		ret.CpuThresholdDrop = in["cpu_threshold_drop"].(int)
		ret.IpdEntryDrop = in["ipd_entry_drop"].(int)
		ret.MaxPacketsExceeded = in["max_packets_exceeded"].(int)
		ret.SessionPacketsExceeded = in["session_packets_exceeded"].(int)
		ret.SctpRcv = in["sctp_rcv"].(int)
		ret.SctpDropped = in["sctp_dropped"].(int)
		ret.FirstGtpPacketTooSmall = in["first_gtp_packet_too_small"].(int)
	}
	return ret
}

func dataToEndpointIpv6FragStats(d *schema.ResourceData) edpt.Ipv6FragStats {
	var ret edpt.Ipv6FragStats

	ret.Stats = getObjectIpv6FragStatsStats(d.Get("stats").([]interface{}))
	return ret
}
