package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6Nat64GlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_nat64_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6Nat64GlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_tcp_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP Ports Allocated",
						},
						"total_tcp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP Ports Freed",
						},
						"total_udp_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP Ports Allocated",
						},
						"total_udp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP Ports Freed",
						},
						"total_icmp_allocated": {
							Type: schema.TypeInt, Optional: true, Description: "Total ICMP Ports Allocated",
						},
						"total_icmp_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Total ICMP Ports Freed",
						},
						"data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Created",
						},
						"data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session Freed",
						},
						"user_quota_created": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Created",
						},
						"user_quota_put_in_del_q": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Freed",
						},
						"user_quota_failure": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Creation Failed",
						},
						"nat_port_unavailable_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP NAT Port Unavailable",
						},
						"nat_port_unavailable_udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP NAT Port Unavailable",
						},
						"nat_port_unavailable_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP NAT Port Unavailable",
						},
						"new_user_resource_unavailable": {
							Type: schema.TypeInt, Optional: true, Description: "New User NAT Resource Unavailable",
						},
						"tcp_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "TCP User-Quota Exceeded",
						},
						"udp_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "UDP User-Quota Exceeded",
						},
						"icmp_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP User-Quota Exceeded",
						},
						"extended_quota_matched": {
							Type: schema.TypeInt, Optional: true, Description: "Extended User-Quota Matched",
						},
						"extended_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Extended User-Quota Exceeded",
						},
						"data_sesn_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Data Session User-Quota Exceeded",
						},
						"data_sesn_rate_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Conn Rate User-Quota Exceeded",
						},
						"tcp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Session Created",
						},
						"tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Full-cone Session Freed",
						},
						"udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Session Created",
						},
						"udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Full-cone Session Freed",
						},
						"fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Full-cone Session Creation Failed",
						},
						"hairpin": {
							Type: schema.TypeInt, Optional: true, Description: "Hairpin Session Created",
						},
						"fullcone_self_hairpinning_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Self-Hairpinning Drop",
						},
						"endpoint_indep_map_match": {
							Type: schema.TypeInt, Optional: true, Description: "Endpoint-Independent Mapping Matched",
						},
						"endpoint_indep_filter_match": {
							Type: schema.TypeInt, Optional: true, Description: "Endpoint-Independent Filtering Matched",
						},
						"inbound_filtered": {
							Type: schema.TypeInt, Optional: true, Description: "Endpoint-Dependent Filtering Drop",
						},
						"eif_limit_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Endpoint-Independent Filtering Inbound Limit Exceeded",
						},
						"total_tcp_overloaded": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloaded",
						},
						"total_udp_overloaded": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloaded",
						},
						"port_overloading_smp_inserted_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading Session Created",
						},
						"port_overloading_smp_inserted_udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading Session Created",
						},
						"port_overloading_smp_free_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading Session Freed",
						},
						"port_overloading_smp_free_udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading Session Freed",
						},
						"nat_pool_unusable": {
							Type: schema.TypeInt, Optional: true, Description: "NAT Pool Unusable",
						},
						"ha_nat_pool_unusable": {
							Type: schema.TypeInt, Optional: true, Description: "HA NAT Pool Unusable",
						},
						"ha_nat_pool_batch_type_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "HA NAT Pool Batch Type Mismatch",
						},
						"no_radius_profile_match": {
							Type: schema.TypeInt, Optional: true, Description: "No RADIUS Profile Match",
						},
						"l3": {
							Type: schema.TypeInt, Optional: true, Description: "Layer 3 Forwarded Packets",
						},
						"source_matches_prefix": {
							Type: schema.TypeInt, Optional: true, Description: "Source Address Prefix Match Drop",
						},
						"lid_drop": {
							Type: schema.TypeInt, Optional: true, Description: "LSN LID Drop",
						},
						"lid_pass_through": {
							Type: schema.TypeInt, Optional: true, Description: "LSN LID Pass-through",
						},
						"no_class_list_match": {
							Type: schema.TypeInt, Optional: true, Description: "No Class-List Match",
						},
						"user_quota_unusable_drop": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Unusable Drop",
						},
						"user_quota_unusable": {
							Type: schema.TypeInt, Optional: true, Description: "User-Quota Marked Unusable",
						},
						"fwd_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets TCP",
						},
						"fwd_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets TCP",
						},
						"rev_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets TCP",
						},
						"rev_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets TCP",
						},
						"fwd_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes TCP",
						},
						"fwd_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes TCP",
						},
						"rev_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes TCP",
						},
						"rev_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes TCP",
						},
						"fwd_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets UDP",
						},
						"fwd_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets UDP",
						},
						"rev_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets UDP",
						},
						"rev_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets UDP",
						},
						"fwd_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes UDP",
						},
						"fwd_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes UDP",
						},
						"rev_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes UDP",
						},
						"rev_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes UDP",
						},
						"fwd_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets ICMP",
						},
						"fwd_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets ICMP",
						},
						"rev_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets ICMP",
						},
						"rev_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets ICMP",
						},
						"fwd_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes ICMP",
						},
						"fwd_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes ICMP",
						},
						"rev_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes ICMP",
						},
						"rev_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes ICMP",
						},
						"fwd_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packets OTHERS",
						},
						"fwd_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packets OTHERS",
						},
						"rev_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packets OTHERS",
						},
						"rev_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packets OTHERS",
						},
						"fwd_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Bytes OTHERS",
						},
						"fwd_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Bytes OTHERS",
						},
						"rev_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Bytes OTHERS",
						},
						"rev_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Bytes OTHERS",
						},
						"fwd_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 0 and 200",
						},
						"fwd_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 201 and 800",
						},
						"fwd_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 801 and 1550",
						},
						"fwd_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Ingress Packet size between 1551 and 9000",
						},
						"fwd_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 0 and 200",
						},
						"fwd_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 201 and 800",
						},
						"fwd_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 801 and 1550",
						},
						"fwd_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Egress Packet size between 1551 and 9000",
						},
						"rev_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 0 and 200",
						},
						"rev_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 201 and 800",
						},
						"rev_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 801 and 1550",
						},
						"rev_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Ingress Packet size between 1551 and 9000",
						},
						"rev_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 0 and 200",
						},
						"rev_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 201 and 800",
						},
						"rev_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 801 and 1550",
						},
						"rev_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Reverse Egress Packet size between 1551 and 9000",
						},
						"prefix_quota_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "Prefix Quota NAT IP Mismatch",
						},
						"port_overloading_port_tcp_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Port Overloading NAT Port TCP Inserted",
						},
						"port_overloading_port_udp_inserted": {
							Type: schema.TypeInt, Optional: true, Description: "Port Overloading NAT Port UDP Inserted",
						},
						"port_overloading_port_free_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "TCP Port Overloading NAT Port Freed",
						},
						"port_overloading_port_free_udp": {
							Type: schema.TypeInt, Optional: true, Description: "UDP Port Overloading NAT Port Freed",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6Nat64GlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6Nat64GlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6Nat64GlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6Nat64GlobalStatsStats := setObjectCgnv6Nat64GlobalStatsStats(res)
		d.Set("stats", Cgnv6Nat64GlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6Nat64GlobalStatsStats(ret edpt.DataCgnv6Nat64GlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_tcp_allocated":                ret.DtCgnv6Nat64GlobalStats.Stats.Total_tcp_allocated,
			"total_tcp_freed":                    ret.DtCgnv6Nat64GlobalStats.Stats.Total_tcp_freed,
			"total_udp_allocated":                ret.DtCgnv6Nat64GlobalStats.Stats.Total_udp_allocated,
			"total_udp_freed":                    ret.DtCgnv6Nat64GlobalStats.Stats.Total_udp_freed,
			"total_icmp_allocated":               ret.DtCgnv6Nat64GlobalStats.Stats.Total_icmp_allocated,
			"total_icmp_freed":                   ret.DtCgnv6Nat64GlobalStats.Stats.Total_icmp_freed,
			"data_session_created":               ret.DtCgnv6Nat64GlobalStats.Stats.Data_session_created,
			"data_session_freed":                 ret.DtCgnv6Nat64GlobalStats.Stats.Data_session_freed,
			"user_quota_created":                 ret.DtCgnv6Nat64GlobalStats.Stats.User_quota_created,
			"user_quota_put_in_del_q":            ret.DtCgnv6Nat64GlobalStats.Stats.User_quota_put_in_del_q,
			"user_quota_failure":                 ret.DtCgnv6Nat64GlobalStats.Stats.User_quota_failure,
			"nat_port_unavailable_tcp":           ret.DtCgnv6Nat64GlobalStats.Stats.Nat_port_unavailable_tcp,
			"nat_port_unavailable_udp":           ret.DtCgnv6Nat64GlobalStats.Stats.Nat_port_unavailable_udp,
			"nat_port_unavailable_icmp":          ret.DtCgnv6Nat64GlobalStats.Stats.Nat_port_unavailable_icmp,
			"new_user_resource_unavailable":      ret.DtCgnv6Nat64GlobalStats.Stats.New_user_resource_unavailable,
			"tcp_user_quota_exceeded":            ret.DtCgnv6Nat64GlobalStats.Stats.Tcp_user_quota_exceeded,
			"udp_user_quota_exceeded":            ret.DtCgnv6Nat64GlobalStats.Stats.Udp_user_quota_exceeded,
			"icmp_user_quota_exceeded":           ret.DtCgnv6Nat64GlobalStats.Stats.Icmp_user_quota_exceeded,
			"extended_quota_matched":             ret.DtCgnv6Nat64GlobalStats.Stats.Extended_quota_matched,
			"extended_quota_exceeded":            ret.DtCgnv6Nat64GlobalStats.Stats.Extended_quota_exceeded,
			"data_sesn_user_quota_exceeded":      ret.DtCgnv6Nat64GlobalStats.Stats.Data_sesn_user_quota_exceeded,
			"data_sesn_rate_user_quota_exceeded": ret.DtCgnv6Nat64GlobalStats.Stats.Data_sesn_rate_user_quota_exceeded,
			"tcp_fullcone_created":               ret.DtCgnv6Nat64GlobalStats.Stats.Tcp_fullcone_created,
			"tcp_fullcone_freed":                 ret.DtCgnv6Nat64GlobalStats.Stats.Tcp_fullcone_freed,
			"udp_fullcone_created":               ret.DtCgnv6Nat64GlobalStats.Stats.Udp_fullcone_created,
			"udp_fullcone_freed":                 ret.DtCgnv6Nat64GlobalStats.Stats.Udp_fullcone_freed,
			"fullcone_failure":                   ret.DtCgnv6Nat64GlobalStats.Stats.Fullcone_failure,
			"hairpin":                            ret.DtCgnv6Nat64GlobalStats.Stats.Hairpin,
			"fullcone_self_hairpinning_drop":     ret.DtCgnv6Nat64GlobalStats.Stats.Fullcone_self_hairpinning_drop,
			"endpoint_indep_map_match":           ret.DtCgnv6Nat64GlobalStats.Stats.Endpoint_indep_map_match,
			"endpoint_indep_filter_match":        ret.DtCgnv6Nat64GlobalStats.Stats.Endpoint_indep_filter_match,
			"inbound_filtered":                   ret.DtCgnv6Nat64GlobalStats.Stats.Inbound_filtered,
			"eif_limit_exceeded":                 ret.DtCgnv6Nat64GlobalStats.Stats.Eif_limit_exceeded,
			"total_tcp_overloaded":               ret.DtCgnv6Nat64GlobalStats.Stats.Total_tcp_overloaded,
			"total_udp_overloaded":               ret.DtCgnv6Nat64GlobalStats.Stats.Total_udp_overloaded,
			"port_overloading_smp_inserted_tcp":  ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_smp_inserted_tcp,
			"port_overloading_smp_inserted_udp":  ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_smp_inserted_udp,
			"port_overloading_smp_free_tcp":      ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_smp_free_tcp,
			"port_overloading_smp_free_udp":      ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_smp_free_udp,
			"nat_pool_unusable":                  ret.DtCgnv6Nat64GlobalStats.Stats.Nat_pool_unusable,
			"ha_nat_pool_unusable":               ret.DtCgnv6Nat64GlobalStats.Stats.Ha_nat_pool_unusable,
			"ha_nat_pool_batch_type_mismatch":    ret.DtCgnv6Nat64GlobalStats.Stats.Ha_nat_pool_batch_type_mismatch,
			"no_radius_profile_match":            ret.DtCgnv6Nat64GlobalStats.Stats.No_radius_profile_match,
			"l3":                                 ret.DtCgnv6Nat64GlobalStats.Stats.L3,
			"source_matches_prefix":              ret.DtCgnv6Nat64GlobalStats.Stats.Source_matches_prefix,
			"lid_drop":                           ret.DtCgnv6Nat64GlobalStats.Stats.Lid_drop,
			"lid_pass_through":                   ret.DtCgnv6Nat64GlobalStats.Stats.Lid_pass_through,
			"no_class_list_match":                ret.DtCgnv6Nat64GlobalStats.Stats.No_class_list_match,
			"user_quota_unusable_drop":           ret.DtCgnv6Nat64GlobalStats.Stats.User_quota_unusable_drop,
			"user_quota_unusable":                ret.DtCgnv6Nat64GlobalStats.Stats.User_quota_unusable,
			"fwd_ingress_packets_tcp":            ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_packets_tcp,
			"fwd_egress_packets_tcp":             ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_packets_tcp,
			"rev_ingress_packets_tcp":            ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_packets_tcp,
			"rev_egress_packets_tcp":             ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_packets_tcp,
			"fwd_ingress_bytes_tcp":              ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_bytes_tcp,
			"fwd_egress_bytes_tcp":               ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_bytes_tcp,
			"rev_ingress_bytes_tcp":              ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_bytes_tcp,
			"rev_egress_bytes_tcp":               ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_bytes_tcp,
			"fwd_ingress_packets_udp":            ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_packets_udp,
			"fwd_egress_packets_udp":             ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_packets_udp,
			"rev_ingress_packets_udp":            ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_packets_udp,
			"rev_egress_packets_udp":             ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_packets_udp,
			"fwd_ingress_bytes_udp":              ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_bytes_udp,
			"fwd_egress_bytes_udp":               ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_bytes_udp,
			"rev_ingress_bytes_udp":              ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_bytes_udp,
			"rev_egress_bytes_udp":               ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_bytes_udp,
			"fwd_ingress_packets_icmp":           ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_packets_icmp,
			"fwd_egress_packets_icmp":            ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_packets_icmp,
			"rev_ingress_packets_icmp":           ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_packets_icmp,
			"rev_egress_packets_icmp":            ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_packets_icmp,
			"fwd_ingress_bytes_icmp":             ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_bytes_icmp,
			"fwd_egress_bytes_icmp":              ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_bytes_icmp,
			"rev_ingress_bytes_icmp":             ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_bytes_icmp,
			"rev_egress_bytes_icmp":              ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_bytes_icmp,
			"fwd_ingress_packets_others":         ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_packets_others,
			"fwd_egress_packets_others":          ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_packets_others,
			"rev_ingress_packets_others":         ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_packets_others,
			"rev_egress_packets_others":          ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_packets_others,
			"fwd_ingress_bytes_others":           ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_bytes_others,
			"fwd_egress_bytes_others":            ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_bytes_others,
			"rev_ingress_bytes_others":           ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_bytes_others,
			"rev_egress_bytes_others":            ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_bytes_others,
			"fwd_ingress_pkt_size_range1":        ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_pkt_size_range1,
			"fwd_ingress_pkt_size_range2":        ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_pkt_size_range2,
			"fwd_ingress_pkt_size_range3":        ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_pkt_size_range3,
			"fwd_ingress_pkt_size_range4":        ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_ingress_pkt_size_range4,
			"fwd_egress_pkt_size_range1":         ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_pkt_size_range1,
			"fwd_egress_pkt_size_range2":         ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_pkt_size_range2,
			"fwd_egress_pkt_size_range3":         ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_pkt_size_range3,
			"fwd_egress_pkt_size_range4":         ret.DtCgnv6Nat64GlobalStats.Stats.Fwd_egress_pkt_size_range4,
			"rev_ingress_pkt_size_range1":        ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_pkt_size_range1,
			"rev_ingress_pkt_size_range2":        ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_pkt_size_range2,
			"rev_ingress_pkt_size_range3":        ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_pkt_size_range3,
			"rev_ingress_pkt_size_range4":        ret.DtCgnv6Nat64GlobalStats.Stats.Rev_ingress_pkt_size_range4,
			"rev_egress_pkt_size_range1":         ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_pkt_size_range1,
			"rev_egress_pkt_size_range2":         ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_pkt_size_range2,
			"rev_egress_pkt_size_range3":         ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_pkt_size_range3,
			"rev_egress_pkt_size_range4":         ret.DtCgnv6Nat64GlobalStats.Stats.Rev_egress_pkt_size_range4,
			"prefix_quota_mismatch":              ret.DtCgnv6Nat64GlobalStats.Stats.Prefix_quota_mismatch,
			"port_overloading_port_tcp_inserted": ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_port_tcp_inserted,
			"port_overloading_port_udp_inserted": ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_port_udp_inserted,
			"port_overloading_port_free_tcp":     ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_port_free_tcp,
			"port_overloading_port_free_udp":     ret.DtCgnv6Nat64GlobalStats.Stats.Port_overloading_port_free_udp,
		},
	}
}

func getObjectCgnv6Nat64GlobalStatsStats(d []interface{}) edpt.Cgnv6Nat64GlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6Nat64GlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Total_tcp_allocated = in["total_tcp_allocated"].(int)
		ret.Total_tcp_freed = in["total_tcp_freed"].(int)
		ret.Total_udp_allocated = in["total_udp_allocated"].(int)
		ret.Total_udp_freed = in["total_udp_freed"].(int)
		ret.Total_icmp_allocated = in["total_icmp_allocated"].(int)
		ret.Total_icmp_freed = in["total_icmp_freed"].(int)
		ret.Data_session_created = in["data_session_created"].(int)
		ret.Data_session_freed = in["data_session_freed"].(int)
		ret.User_quota_created = in["user_quota_created"].(int)
		ret.User_quota_put_in_del_q = in["user_quota_put_in_del_q"].(int)
		ret.User_quota_failure = in["user_quota_failure"].(int)
		ret.Nat_port_unavailable_tcp = in["nat_port_unavailable_tcp"].(int)
		ret.Nat_port_unavailable_udp = in["nat_port_unavailable_udp"].(int)
		ret.Nat_port_unavailable_icmp = in["nat_port_unavailable_icmp"].(int)
		ret.New_user_resource_unavailable = in["new_user_resource_unavailable"].(int)
		ret.Tcp_user_quota_exceeded = in["tcp_user_quota_exceeded"].(int)
		ret.Udp_user_quota_exceeded = in["udp_user_quota_exceeded"].(int)
		ret.Icmp_user_quota_exceeded = in["icmp_user_quota_exceeded"].(int)
		ret.Extended_quota_matched = in["extended_quota_matched"].(int)
		ret.Extended_quota_exceeded = in["extended_quota_exceeded"].(int)
		ret.Data_sesn_user_quota_exceeded = in["data_sesn_user_quota_exceeded"].(int)
		ret.Data_sesn_rate_user_quota_exceeded = in["data_sesn_rate_user_quota_exceeded"].(int)
		ret.Tcp_fullcone_created = in["tcp_fullcone_created"].(int)
		ret.Tcp_fullcone_freed = in["tcp_fullcone_freed"].(int)
		ret.Udp_fullcone_created = in["udp_fullcone_created"].(int)
		ret.Udp_fullcone_freed = in["udp_fullcone_freed"].(int)
		ret.Fullcone_failure = in["fullcone_failure"].(int)
		ret.Hairpin = in["hairpin"].(int)
		ret.Fullcone_self_hairpinning_drop = in["fullcone_self_hairpinning_drop"].(int)
		ret.Endpoint_indep_map_match = in["endpoint_indep_map_match"].(int)
		ret.Endpoint_indep_filter_match = in["endpoint_indep_filter_match"].(int)
		ret.Inbound_filtered = in["inbound_filtered"].(int)
		ret.Eif_limit_exceeded = in["eif_limit_exceeded"].(int)
		ret.Total_tcp_overloaded = in["total_tcp_overloaded"].(int)
		ret.Total_udp_overloaded = in["total_udp_overloaded"].(int)
		ret.Port_overloading_smp_inserted_tcp = in["port_overloading_smp_inserted_tcp"].(int)
		ret.Port_overloading_smp_inserted_udp = in["port_overloading_smp_inserted_udp"].(int)
		ret.Port_overloading_smp_free_tcp = in["port_overloading_smp_free_tcp"].(int)
		ret.Port_overloading_smp_free_udp = in["port_overloading_smp_free_udp"].(int)
		ret.Nat_pool_unusable = in["nat_pool_unusable"].(int)
		ret.Ha_nat_pool_unusable = in["ha_nat_pool_unusable"].(int)
		ret.Ha_nat_pool_batch_type_mismatch = in["ha_nat_pool_batch_type_mismatch"].(int)
		ret.No_radius_profile_match = in["no_radius_profile_match"].(int)
		ret.L3 = in["l3"].(int)
		ret.Source_matches_prefix = in["source_matches_prefix"].(int)
		ret.Lid_drop = in["lid_drop"].(int)
		ret.Lid_pass_through = in["lid_pass_through"].(int)
		ret.No_class_list_match = in["no_class_list_match"].(int)
		ret.User_quota_unusable_drop = in["user_quota_unusable_drop"].(int)
		ret.User_quota_unusable = in["user_quota_unusable"].(int)
		ret.Fwd_ingress_packets_tcp = in["fwd_ingress_packets_tcp"].(int)
		ret.Fwd_egress_packets_tcp = in["fwd_egress_packets_tcp"].(int)
		ret.Rev_ingress_packets_tcp = in["rev_ingress_packets_tcp"].(int)
		ret.Rev_egress_packets_tcp = in["rev_egress_packets_tcp"].(int)
		ret.Fwd_ingress_bytes_tcp = in["fwd_ingress_bytes_tcp"].(int)
		ret.Fwd_egress_bytes_tcp = in["fwd_egress_bytes_tcp"].(int)
		ret.Rev_ingress_bytes_tcp = in["rev_ingress_bytes_tcp"].(int)
		ret.Rev_egress_bytes_tcp = in["rev_egress_bytes_tcp"].(int)
		ret.Fwd_ingress_packets_udp = in["fwd_ingress_packets_udp"].(int)
		ret.Fwd_egress_packets_udp = in["fwd_egress_packets_udp"].(int)
		ret.Rev_ingress_packets_udp = in["rev_ingress_packets_udp"].(int)
		ret.Rev_egress_packets_udp = in["rev_egress_packets_udp"].(int)
		ret.Fwd_ingress_bytes_udp = in["fwd_ingress_bytes_udp"].(int)
		ret.Fwd_egress_bytes_udp = in["fwd_egress_bytes_udp"].(int)
		ret.Rev_ingress_bytes_udp = in["rev_ingress_bytes_udp"].(int)
		ret.Rev_egress_bytes_udp = in["rev_egress_bytes_udp"].(int)
		ret.Fwd_ingress_packets_icmp = in["fwd_ingress_packets_icmp"].(int)
		ret.Fwd_egress_packets_icmp = in["fwd_egress_packets_icmp"].(int)
		ret.Rev_ingress_packets_icmp = in["rev_ingress_packets_icmp"].(int)
		ret.Rev_egress_packets_icmp = in["rev_egress_packets_icmp"].(int)
		ret.Fwd_ingress_bytes_icmp = in["fwd_ingress_bytes_icmp"].(int)
		ret.Fwd_egress_bytes_icmp = in["fwd_egress_bytes_icmp"].(int)
		ret.Rev_ingress_bytes_icmp = in["rev_ingress_bytes_icmp"].(int)
		ret.Rev_egress_bytes_icmp = in["rev_egress_bytes_icmp"].(int)
		ret.Fwd_ingress_packets_others = in["fwd_ingress_packets_others"].(int)
		ret.Fwd_egress_packets_others = in["fwd_egress_packets_others"].(int)
		ret.Rev_ingress_packets_others = in["rev_ingress_packets_others"].(int)
		ret.Rev_egress_packets_others = in["rev_egress_packets_others"].(int)
		ret.Fwd_ingress_bytes_others = in["fwd_ingress_bytes_others"].(int)
		ret.Fwd_egress_bytes_others = in["fwd_egress_bytes_others"].(int)
		ret.Rev_ingress_bytes_others = in["rev_ingress_bytes_others"].(int)
		ret.Rev_egress_bytes_others = in["rev_egress_bytes_others"].(int)
		ret.Fwd_ingress_pkt_size_range1 = in["fwd_ingress_pkt_size_range1"].(int)
		ret.Fwd_ingress_pkt_size_range2 = in["fwd_ingress_pkt_size_range2"].(int)
		ret.Fwd_ingress_pkt_size_range3 = in["fwd_ingress_pkt_size_range3"].(int)
		ret.Fwd_ingress_pkt_size_range4 = in["fwd_ingress_pkt_size_range4"].(int)
		ret.Fwd_egress_pkt_size_range1 = in["fwd_egress_pkt_size_range1"].(int)
		ret.Fwd_egress_pkt_size_range2 = in["fwd_egress_pkt_size_range2"].(int)
		ret.Fwd_egress_pkt_size_range3 = in["fwd_egress_pkt_size_range3"].(int)
		ret.Fwd_egress_pkt_size_range4 = in["fwd_egress_pkt_size_range4"].(int)
		ret.Rev_ingress_pkt_size_range1 = in["rev_ingress_pkt_size_range1"].(int)
		ret.Rev_ingress_pkt_size_range2 = in["rev_ingress_pkt_size_range2"].(int)
		ret.Rev_ingress_pkt_size_range3 = in["rev_ingress_pkt_size_range3"].(int)
		ret.Rev_ingress_pkt_size_range4 = in["rev_ingress_pkt_size_range4"].(int)
		ret.Rev_egress_pkt_size_range1 = in["rev_egress_pkt_size_range1"].(int)
		ret.Rev_egress_pkt_size_range2 = in["rev_egress_pkt_size_range2"].(int)
		ret.Rev_egress_pkt_size_range3 = in["rev_egress_pkt_size_range3"].(int)
		ret.Rev_egress_pkt_size_range4 = in["rev_egress_pkt_size_range4"].(int)
		ret.Prefix_quota_mismatch = in["prefix_quota_mismatch"].(int)
		ret.Port_overloading_port_tcp_inserted = in["port_overloading_port_tcp_inserted"].(int)
		ret.Port_overloading_port_udp_inserted = in["port_overloading_port_udp_inserted"].(int)
		ret.Port_overloading_port_free_tcp = in["port_overloading_port_free_tcp"].(int)
		ret.Port_overloading_port_free_udp = in["port_overloading_port_free_udp"].(int)
	}
	return ret
}

func dataToEndpointCgnv6Nat64GlobalStats(d *schema.ResourceData) edpt.Cgnv6Nat64GlobalStats {
	var ret edpt.Cgnv6Nat64GlobalStats

	ret.Stats = getObjectCgnv6Nat64GlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
