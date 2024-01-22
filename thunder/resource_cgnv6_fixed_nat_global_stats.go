package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6FixedNatGlobalStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_cgnv6_fixed_nat_global_stats`: Statistics for the object global\n\n__PLACEHOLDER__",
		ReadContext: resourceCgnv6FixedNatGlobalStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total_nat_in_use": {
							Type: schema.TypeInt, Optional: true, Description: "Total NAT Addresses in-use",
						},
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
						"nat44_data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Data Sessions Created",
						},
						"nat44_data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Data Sessions Freed",
						},
						"nat64_data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Data Sessions Created",
						},
						"nat64_data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Data Sessions Freed",
						},
						"dslite_data_session_created": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Data Sessions Created",
						},
						"dslite_data_session_freed": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Data Sessions Freed",
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
						"session_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "Sessions User Quota Exceeded",
						},
						"nat44_tcp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 TCP Full-Cone Created",
						},
						"nat44_tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 TCP Full-Cone Freed",
						},
						"nat44_udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 UDP Full-Cone Created",
						},
						"nat44_udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 UDP Full-Cone Freed",
						},
						"nat44_udp_alg_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 UDP ALG Full-Cone Created",
						},
						"nat44_udp_alg_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 UDP ALG Full-Cone Freed",
						},
						"nat64_tcp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 TCP Full-Cone Created",
						},
						"nat64_tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 TCP Full-Cone Freed",
						},
						"nat64_udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 UDP Full-Cone Created",
						},
						"nat64_udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 UDP Full-Cone Freed",
						},
						"nat64_udp_alg_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 UDP ALG Full-Cone Created",
						},
						"nat64_udp_alg_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 UDP ALG Full-Cone Freed",
						},
						"dslite_tcp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite TCP Full-Cone Created",
						},
						"dslite_tcp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite TCP Full-Cone Freed",
						},
						"dslite_udp_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite UDP Full-Cone Created",
						},
						"dslite_udp_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite UDP Full-Cone Freed",
						},
						"dslite_udp_alg_fullcone_created": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite UDP ALG Full-Cone Created",
						},
						"dslite_udp_alg_fullcone_freed": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite UDP ALG Full-Cone Freed",
						},
						"fullcone_failure": {
							Type: schema.TypeInt, Optional: true, Description: "Full-Cone Session Creation Failed",
						},
						"nat44_eim_match": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Endpoint-Independent-Mapping Matched",
						},
						"nat64_eim_match": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Endpoint-Independent-Mapping Matched",
						},
						"dslite_eim_match": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Endpoint-Independent-Mapping Matched",
						},
						"nat44_eif_match": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Endpoint-Independent-Filtering Matched",
						},
						"nat64_eif_match": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Endpoint-Independent-Filtering Matched",
						},
						"dslite_eif_match": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Endpoint-Independent-Filtering Matched",
						},
						"nat44_inbound_filtered": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Endpoint-Dependent Filtering Drop",
						},
						"nat64_inbound_filtered": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Endpoint-Dependent Filtering Drop",
						},
						"dslite_inbound_filtered": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Endpoint-Dependent Filtering Drop",
						},
						"nat44_eif_limit_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Endpoint-Independent-Filtering Limit Exceeded",
						},
						"nat64_eif_limit_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Endpoint-Independent-Filtering Limit Exceeded",
						},
						"dslite_eif_limit_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Endpoint-Independent-Filtering Limit Exceeded",
						},
						"nat44_hairpin": {
							Type: schema.TypeInt, Optional: true, Description: "NAT44 Hairpin Session Created",
						},
						"nat64_hairpin": {
							Type: schema.TypeInt, Optional: true, Description: "NAT64 Hairpin Session Created",
						},
						"dslite_hairpin": {
							Type: schema.TypeInt, Optional: true, Description: "DS-Lite Hairpin Session Created",
						},
						"standby_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT LID Standby Drop",
						},
						"fixed_nat_fullcone_self_hairpinning_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Self-Hairpinning Drop",
						},
						"sixrd_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT IPv6 in IPv4 Packet Drop",
						},
						"dest_rlist_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Dest Rule List Drop",
						},
						"dest_rlist_pass_through": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Dest Rule List Pass-Through",
						},
						"dest_rlist_snat_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Dest Rules List Source NAT Drop",
						},
						"config_not_found": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT Config not Found",
						},
						"total_tcp_overload_acquired": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP ports acquired for port overloading",
						},
						"total_udp_overload_acquired": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP ports acquired for port overloading",
						},
						"total_tcp_overload_released": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP ports released from port overloading",
						},
						"total_udp_overload_released": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP ports released from port overloading",
						},
						"total_tcp_alloc_overload": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP ports allocated via overload",
						},
						"total_udp_alloc_overload": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP ports allocated via overload",
						},
						"total_tcp_free_overload": {
							Type: schema.TypeInt, Optional: true, Description: "Total TCP ports freed via overload",
						},
						"total_udp_free_overload": {
							Type: schema.TypeInt, Optional: true, Description: "Total UDP ports freed via overload",
						},
						"port_overload_smp_delete_scheduled": {
							Type: schema.TypeInt, Optional: true, Description: "Port overload SMP conn delete scheduled",
						},
						"port_overload_failed": {
							Type: schema.TypeInt, Optional: true, Description: "Port overload failed",
						},
						"ha_session_user_quota_exceeded": {
							Type: schema.TypeInt, Optional: true, Description: "HA Sessions User Quota Exceeded",
						},
						"fnat44_fwd_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packets TCP",
						},
						"fnat44_fwd_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packets TCP",
						},
						"fnat44_rev_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packets TCP",
						},
						"fnat44_rev_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packets TCP",
						},
						"fnat44_fwd_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Bytes TCP",
						},
						"fnat44_fwd_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Bytes TCP",
						},
						"fnat44_rev_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Bytes TCP",
						},
						"fnat44_rev_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Bytes TCP",
						},
						"fnat44_fwd_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packets UDP",
						},
						"fnat44_fwd_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packets UDP",
						},
						"fnat44_rev_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packets UDP",
						},
						"fnat44_rev_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packets UDP",
						},
						"fnat44_fwd_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Bytes UDP",
						},
						"fnat44_fwd_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Bytes UDP",
						},
						"fnat44_rev_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Bytes UDP",
						},
						"fnat44_rev_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Bytes UDP",
						},
						"fnat44_fwd_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packets ICMP",
						},
						"fnat44_fwd_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packets ICMP",
						},
						"fnat44_rev_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packets ICMP",
						},
						"fnat44_rev_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packets ICMP",
						},
						"fnat44_fwd_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Bytes ICMP",
						},
						"fnat44_fwd_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Bytes ICMP",
						},
						"fnat44_rev_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Bytes ICMP",
						},
						"fnat44_rev_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Bytes ICMP",
						},
						"fnat44_fwd_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packets OTHERS",
						},
						"fnat44_fwd_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packets OTHERS",
						},
						"fnat44_rev_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packets OTHERS",
						},
						"fnat44_rev_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packets OTHERS",
						},
						"fnat44_fwd_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Bytes OTHERS",
						},
						"fnat44_fwd_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Bytes OTHERS",
						},
						"fnat44_rev_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Bytes OTHERS",
						},
						"fnat44_rev_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Bytes OTHERS",
						},
						"fnat44_fwd_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packet size between 0 and 200",
						},
						"fnat44_fwd_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packet size between 201 and 800",
						},
						"fnat44_fwd_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packet size between 801 and 1550",
						},
						"fnat44_fwd_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Ingress Packet size between 1551 and 9000",
						},
						"fnat44_fwd_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packet size between 0 and 200",
						},
						"fnat44_fwd_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packet size between 201 and 800",
						},
						"fnat44_fwd_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packet size between 801 and 1550",
						},
						"fnat44_fwd_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Forward Egress Packet size between 1551 and 9000",
						},
						"fnat44_rev_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packet size between 0 and 200",
						},
						"fnat44_rev_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packet size between 201 and 800",
						},
						"fnat44_rev_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packet size between 801 and 1550",
						},
						"fnat44_rev_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Ingress Packet size between 1551 and 9000",
						},
						"fnat44_rev_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packet size between 0 and 200",
						},
						"fnat44_rev_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packet size between 201 and 800",
						},
						"fnat44_rev_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packet size between 801 and 1550",
						},
						"fnat44_rev_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT44 Reverse Egress Packet size between 1551 and 9000",
						},
						"fnat64_fwd_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packets TCP",
						},
						"fnat64_fwd_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packets TCP",
						},
						"fnat64_rev_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packets TCP",
						},
						"fnat64_rev_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packets TCP",
						},
						"fnat64_fwd_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Bytes TCP",
						},
						"fnat64_fwd_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Bytes TCP",
						},
						"fnat64_rev_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Bytes TCP",
						},
						"fnat64_rev_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Bytes TCP",
						},
						"fnat64_fwd_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packets UDP",
						},
						"fnat64_fwd_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packets UDP",
						},
						"fnat64_rev_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packets UDP",
						},
						"fnat64_rev_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packets UDP",
						},
						"fnat64_fwd_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Bytes UDP",
						},
						"fnat64_fwd_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Bytes UDP",
						},
						"fnat64_rev_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Bytes UDP",
						},
						"fnat64_rev_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Bytes UDP",
						},
						"fnat64_fwd_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packets ICMP",
						},
						"fnat64_fwd_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packets ICMP",
						},
						"fnat64_rev_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packets ICMP",
						},
						"fnat64_rev_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packets ICMP",
						},
						"fnat64_fwd_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Bytes ICMP",
						},
						"fnat64_fwd_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Bytes ICMP",
						},
						"fnat64_rev_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Bytes ICMP",
						},
						"fnat64_rev_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Bytes ICMP",
						},
						"fnat64_fwd_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packets OTHERS",
						},
						"fnat64_fwd_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packets OTHERS",
						},
						"fnat64_rev_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packets OTHERS",
						},
						"fnat64_rev_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packets OTHERS",
						},
						"fnat64_fwd_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Bytes OTHERS",
						},
						"fnat64_fwd_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Bytes OTHERS",
						},
						"fnat64_rev_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Bytes OTHERS",
						},
						"fnat64_rev_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Bytes OTHERS",
						},
						"fnat64_fwd_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packet size between 0 and 200",
						},
						"fnat64_fwd_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packet size between 201 and 800",
						},
						"fnat64_fwd_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packet size between 801 and 1550",
						},
						"fnat64_fwd_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Ingress Packet size between 1551 and 9000",
						},
						"fnat64_fwd_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packet size between 0 and 200",
						},
						"fnat64_fwd_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packet size between 201 and 800",
						},
						"fnat64_fwd_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packet size between 801 and 1550",
						},
						"fnat64_fwd_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Forward Egress Packet size between 1551 and 9000",
						},
						"fnat64_rev_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packet size between 0 and 200",
						},
						"fnat64_rev_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packet size between 201 and 800",
						},
						"fnat64_rev_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packet size between 801 and 1550",
						},
						"fnat64_rev_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Ingress Packet size between 1551 and 9000",
						},
						"fnat64_rev_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packet size between 0 and 200",
						},
						"fnat64_rev_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packet size between 201 and 800",
						},
						"fnat64_rev_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packet size between 801 and 1550",
						},
						"fnat64_rev_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed NAT64 Reverse Egress Packet size between 1551 and 9000",
						},
						"fnatdslite_fwd_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packets TCP",
						},
						"fnatdslite_fwd_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packets TCP",
						},
						"fnatdslite_rev_ingress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packets TCP",
						},
						"fnatdslite_rev_egress_packets_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packets TCP",
						},
						"fnatdslite_fwd_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Bytes TCP",
						},
						"fnatdslite_fwd_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Bytes TCP",
						},
						"fnatdslite_rev_ingress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Bytes TCP",
						},
						"fnatdslite_rev_egress_bytes_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Bytes TCP",
						},
						"fnatdslite_fwd_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packets UDP",
						},
						"fnatdslite_fwd_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packets UDP",
						},
						"fnatdslite_rev_ingress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packets UDP",
						},
						"fnatdslite_rev_egress_packets_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packets UDP",
						},
						"fnatdslite_fwd_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Bytes UDP",
						},
						"fnatdslite_fwd_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Bytes UDP",
						},
						"fnatdslite_rev_ingress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Bytes UDP",
						},
						"fnatdslite_rev_egress_bytes_udp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Bytes UDP",
						},
						"fnatdslite_fwd_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packets ICMP",
						},
						"fnatdslite_fwd_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packets ICMP",
						},
						"fnatdslite_rev_ingress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packets ICMP",
						},
						"fnatdslite_rev_egress_packets_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packets ICMP",
						},
						"fnatdslite_fwd_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Bytes ICMP",
						},
						"fnatdslite_fwd_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Bytes ICMP",
						},
						"fnatdslite_rev_ingress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Bytes ICMP",
						},
						"fnatdslite_rev_egress_bytes_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Bytes ICMP",
						},
						"fnatdslite_fwd_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packets OTHERS",
						},
						"fnatdslite_fwd_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packets OTHERS",
						},
						"fnatdslite_rev_ingress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packets OTHERS",
						},
						"fnatdslite_rev_egress_packets_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packets OTHERS",
						},
						"fnatdslite_fwd_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Bytes OTHERS",
						},
						"fnatdslite_fwd_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Bytes OTHERS",
						},
						"fnatdslite_rev_ingress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Bytes OTHERS",
						},
						"fnatdslite_rev_egress_bytes_others": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Bytes OTHERS",
						},
						"fnatdslite_fwd_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packet size between 0 and 200",
						},
						"fnatdslite_fwd_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packet size between 201 and 800",
						},
						"fnatdslite_fwd_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packet size between 801 and 1550",
						},
						"fnatdslite_fwd_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Ingress Packet size between 1551 and 9000",
						},
						"fnatdslite_fwd_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packet size between 0 and 200",
						},
						"fnatdslite_fwd_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packet size between 201 and 800",
						},
						"fnatdslite_fwd_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packet size between 801 and 1550",
						},
						"fnatdslite_fwd_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Forward Egress Packet size between 1551 and 9000",
						},
						"fnatdslite_rev_ingress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packet size between 0 and 200",
						},
						"fnatdslite_rev_ingress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packet size between 201 and 800",
						},
						"fnatdslite_rev_ingress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packet size between 801 and 1550",
						},
						"fnatdslite_rev_ingress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Ingress Packet size between 1551 and 9000",
						},
						"fnatdslite_rev_egress_pkt_size_range1": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packet size between 0 and 200",
						},
						"fnatdslite_rev_egress_pkt_size_range2": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packet size between 201 and 800",
						},
						"fnatdslite_rev_egress_pkt_size_range3": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packet size between 801 and 1550",
						},
						"fnatdslite_rev_egress_pkt_size_range4": {
							Type: schema.TypeInt, Optional: true, Description: "Fixed DS-Lite Reverse Egress Packet size between 1551 and 9000",
						},
						"active_subscriber_added": {
							Type: schema.TypeInt, Optional: true, Description: "Active Subscriber Added",
						},
						"active_subscriber_removed": {
							Type: schema.TypeInt, Optional: true, Description: "Active Subscriber Removed",
						},
					},
				},
			},
		},
	}
}

func resourceCgnv6FixedNatGlobalStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6FixedNatGlobalStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6FixedNatGlobalStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		Cgnv6FixedNatGlobalStatsStats := setObjectCgnv6FixedNatGlobalStatsStats(res)
		d.Set("stats", Cgnv6FixedNatGlobalStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectCgnv6FixedNatGlobalStatsStats(ret edpt.DataCgnv6FixedNatGlobalStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"total_nat_in_use":                         ret.DtCgnv6FixedNatGlobalStats.Stats.TotalNatInUse,
			"total_tcp_allocated":                      ret.DtCgnv6FixedNatGlobalStats.Stats.TotalTcpAllocated,
			"total_tcp_freed":                          ret.DtCgnv6FixedNatGlobalStats.Stats.TotalTcpFreed,
			"total_udp_allocated":                      ret.DtCgnv6FixedNatGlobalStats.Stats.TotalUdpAllocated,
			"total_udp_freed":                          ret.DtCgnv6FixedNatGlobalStats.Stats.TotalUdpFreed,
			"total_icmp_allocated":                     ret.DtCgnv6FixedNatGlobalStats.Stats.TotalIcmpAllocated,
			"total_icmp_freed":                         ret.DtCgnv6FixedNatGlobalStats.Stats.TotalIcmpFreed,
			"nat44_data_session_created":               ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44DataSessionCreated,
			"nat44_data_session_freed":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44DataSessionFreed,
			"nat64_data_session_created":               ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64DataSessionCreated,
			"nat64_data_session_freed":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64DataSessionFreed,
			"dslite_data_session_created":              ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteDataSessionCreated,
			"dslite_data_session_freed":                ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteDataSessionFreed,
			"nat_port_unavailable_tcp":                 ret.DtCgnv6FixedNatGlobalStats.Stats.NatPortUnavailableTcp,
			"nat_port_unavailable_udp":                 ret.DtCgnv6FixedNatGlobalStats.Stats.NatPortUnavailableUdp,
			"nat_port_unavailable_icmp":                ret.DtCgnv6FixedNatGlobalStats.Stats.NatPortUnavailableIcmp,
			"session_user_quota_exceeded":              ret.DtCgnv6FixedNatGlobalStats.Stats.SessionUserQuotaExceeded,
			"nat44_tcp_fullcone_created":               ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44TcpFullconeCreated,
			"nat44_tcp_fullcone_freed":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44TcpFullconeFreed,
			"nat44_udp_fullcone_created":               ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44UdpFullconeCreated,
			"nat44_udp_fullcone_freed":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44UdpFullconeFreed,
			"nat44_udp_alg_fullcone_created":           ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44UdpAlgFullconeCreated,
			"nat44_udp_alg_fullcone_freed":             ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44UdpAlgFullconeFreed,
			"nat64_tcp_fullcone_created":               ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64TcpFullconeCreated,
			"nat64_tcp_fullcone_freed":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64TcpFullconeFreed,
			"nat64_udp_fullcone_created":               ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64UdpFullconeCreated,
			"nat64_udp_fullcone_freed":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64UdpFullconeFreed,
			"nat64_udp_alg_fullcone_created":           ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64UdpAlgFullconeCreated,
			"nat64_udp_alg_fullcone_freed":             ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64UdpAlgFullconeFreed,
			"dslite_tcp_fullcone_created":              ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteTcpFullconeCreated,
			"dslite_tcp_fullcone_freed":                ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteTcpFullconeFreed,
			"dslite_udp_fullcone_created":              ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteUdpFullconeCreated,
			"dslite_udp_fullcone_freed":                ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteUdpFullconeFreed,
			"dslite_udp_alg_fullcone_created":          ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteUdpAlgFullconeCreated,
			"dslite_udp_alg_fullcone_freed":            ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteUdpAlgFullconeFreed,
			"fullcone_failure":                         ret.DtCgnv6FixedNatGlobalStats.Stats.FullconeFailure,
			"nat44_eim_match":                          ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44EimMatch,
			"nat64_eim_match":                          ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64EimMatch,
			"dslite_eim_match":                         ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteEimMatch,
			"nat44_eif_match":                          ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44EifMatch,
			"nat64_eif_match":                          ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64EifMatch,
			"dslite_eif_match":                         ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteEifMatch,
			"nat44_inbound_filtered":                   ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44InboundFiltered,
			"nat64_inbound_filtered":                   ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64InboundFiltered,
			"dslite_inbound_filtered":                  ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteInboundFiltered,
			"nat44_eif_limit_exceeded":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44EifLimitExceeded,
			"nat64_eif_limit_exceeded":                 ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64EifLimitExceeded,
			"dslite_eif_limit_exceeded":                ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteEifLimitExceeded,
			"nat44_hairpin":                            ret.DtCgnv6FixedNatGlobalStats.Stats.Nat44Hairpin,
			"nat64_hairpin":                            ret.DtCgnv6FixedNatGlobalStats.Stats.Nat64Hairpin,
			"dslite_hairpin":                           ret.DtCgnv6FixedNatGlobalStats.Stats.DsliteHairpin,
			"standby_drop":                             ret.DtCgnv6FixedNatGlobalStats.Stats.StandbyDrop,
			"fixed_nat_fullcone_self_hairpinning_drop": ret.DtCgnv6FixedNatGlobalStats.Stats.FixedNatFullconeSelfHairpinningDrop,
			"sixrd_drop":                               ret.DtCgnv6FixedNatGlobalStats.Stats.SixrdDrop,
			"dest_rlist_drop":                          ret.DtCgnv6FixedNatGlobalStats.Stats.DestRlistDrop,
			"dest_rlist_pass_through":                  ret.DtCgnv6FixedNatGlobalStats.Stats.DestRlistPassThrough,
			"dest_rlist_snat_drop":                     ret.DtCgnv6FixedNatGlobalStats.Stats.DestRlistSnatDrop,
			"config_not_found":                         ret.DtCgnv6FixedNatGlobalStats.Stats.ConfigNotFound,
			"total_tcp_overload_acquired":              ret.DtCgnv6FixedNatGlobalStats.Stats.TotalTcpOverloadAcquired,
			"total_udp_overload_acquired":              ret.DtCgnv6FixedNatGlobalStats.Stats.TotalUdpOverloadAcquired,
			"total_tcp_overload_released":              ret.DtCgnv6FixedNatGlobalStats.Stats.TotalTcpOverloadReleased,
			"total_udp_overload_released":              ret.DtCgnv6FixedNatGlobalStats.Stats.TotalUdpOverloadReleased,
			"total_tcp_alloc_overload":                 ret.DtCgnv6FixedNatGlobalStats.Stats.TotalTcpAllocOverload,
			"total_udp_alloc_overload":                 ret.DtCgnv6FixedNatGlobalStats.Stats.TotalUdpAllocOverload,
			"total_tcp_free_overload":                  ret.DtCgnv6FixedNatGlobalStats.Stats.TotalTcpFreeOverload,
			"total_udp_free_overload":                  ret.DtCgnv6FixedNatGlobalStats.Stats.TotalUdpFreeOverload,
			"port_overload_smp_delete_scheduled":       ret.DtCgnv6FixedNatGlobalStats.Stats.PortOverloadSmpDeleteScheduled,
			"port_overload_failed":                     ret.DtCgnv6FixedNatGlobalStats.Stats.PortOverloadFailed,
			"ha_session_user_quota_exceeded":           ret.DtCgnv6FixedNatGlobalStats.Stats.HaSessionUserQuotaExceeded,
			"fnat44_fwd_ingress_packets_tcp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_packets_tcp,
			"fnat44_fwd_egress_packets_tcp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_packets_tcp,
			"fnat44_rev_ingress_packets_tcp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_packets_tcp,
			"fnat44_rev_egress_packets_tcp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_packets_tcp,
			"fnat44_fwd_ingress_bytes_tcp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_bytes_tcp,
			"fnat44_fwd_egress_bytes_tcp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_bytes_tcp,
			"fnat44_rev_ingress_bytes_tcp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_bytes_tcp,
			"fnat44_rev_egress_bytes_tcp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_bytes_tcp,
			"fnat44_fwd_ingress_packets_udp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_packets_udp,
			"fnat44_fwd_egress_packets_udp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_packets_udp,
			"fnat44_rev_ingress_packets_udp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_packets_udp,
			"fnat44_rev_egress_packets_udp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_packets_udp,
			"fnat44_fwd_ingress_bytes_udp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_bytes_udp,
			"fnat44_fwd_egress_bytes_udp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_bytes_udp,
			"fnat44_rev_ingress_bytes_udp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_bytes_udp,
			"fnat44_rev_egress_bytes_udp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_bytes_udp,
			"fnat44_fwd_ingress_packets_icmp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_packets_icmp,
			"fnat44_fwd_egress_packets_icmp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_packets_icmp,
			"fnat44_rev_ingress_packets_icmp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_packets_icmp,
			"fnat44_rev_egress_packets_icmp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_packets_icmp,
			"fnat44_fwd_ingress_bytes_icmp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_bytes_icmp,
			"fnat44_fwd_egress_bytes_icmp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_bytes_icmp,
			"fnat44_rev_ingress_bytes_icmp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_bytes_icmp,
			"fnat44_rev_egress_bytes_icmp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_bytes_icmp,
			"fnat44_fwd_ingress_packets_others":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_packets_others,
			"fnat44_fwd_egress_packets_others":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_packets_others,
			"fnat44_rev_ingress_packets_others":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_packets_others,
			"fnat44_rev_egress_packets_others":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_packets_others,
			"fnat44_fwd_ingress_bytes_others":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_bytes_others,
			"fnat44_fwd_egress_bytes_others":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_bytes_others,
			"fnat44_rev_ingress_bytes_others":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_bytes_others,
			"fnat44_rev_egress_bytes_others":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_bytes_others,
			"fnat44_fwd_ingress_pkt_size_range1":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_pkt_size_range1,
			"fnat44_fwd_ingress_pkt_size_range2":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_pkt_size_range2,
			"fnat44_fwd_ingress_pkt_size_range3":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_pkt_size_range3,
			"fnat44_fwd_ingress_pkt_size_range4":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_ingress_pkt_size_range4,
			"fnat44_fwd_egress_pkt_size_range1":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_pkt_size_range1,
			"fnat44_fwd_egress_pkt_size_range2":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_pkt_size_range2,
			"fnat44_fwd_egress_pkt_size_range3":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_pkt_size_range3,
			"fnat44_fwd_egress_pkt_size_range4":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_fwd_egress_pkt_size_range4,
			"fnat44_rev_ingress_pkt_size_range1":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_pkt_size_range1,
			"fnat44_rev_ingress_pkt_size_range2":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_pkt_size_range2,
			"fnat44_rev_ingress_pkt_size_range3":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_pkt_size_range3,
			"fnat44_rev_ingress_pkt_size_range4":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_ingress_pkt_size_range4,
			"fnat44_rev_egress_pkt_size_range1":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_pkt_size_range1,
			"fnat44_rev_egress_pkt_size_range2":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_pkt_size_range2,
			"fnat44_rev_egress_pkt_size_range3":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_pkt_size_range3,
			"fnat44_rev_egress_pkt_size_range4":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat44_rev_egress_pkt_size_range4,
			"fnat64_fwd_ingress_packets_tcp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_packets_tcp,
			"fnat64_fwd_egress_packets_tcp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_packets_tcp,
			"fnat64_rev_ingress_packets_tcp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_packets_tcp,
			"fnat64_rev_egress_packets_tcp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_packets_tcp,
			"fnat64_fwd_ingress_bytes_tcp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_bytes_tcp,
			"fnat64_fwd_egress_bytes_tcp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_bytes_tcp,
			"fnat64_rev_ingress_bytes_tcp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_bytes_tcp,
			"fnat64_rev_egress_bytes_tcp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_bytes_tcp,
			"fnat64_fwd_ingress_packets_udp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_packets_udp,
			"fnat64_fwd_egress_packets_udp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_packets_udp,
			"fnat64_rev_ingress_packets_udp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_packets_udp,
			"fnat64_rev_egress_packets_udp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_packets_udp,
			"fnat64_fwd_ingress_bytes_udp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_bytes_udp,
			"fnat64_fwd_egress_bytes_udp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_bytes_udp,
			"fnat64_rev_ingress_bytes_udp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_bytes_udp,
			"fnat64_rev_egress_bytes_udp":              ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_bytes_udp,
			"fnat64_fwd_ingress_packets_icmp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_packets_icmp,
			"fnat64_fwd_egress_packets_icmp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_packets_icmp,
			"fnat64_rev_ingress_packets_icmp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_packets_icmp,
			"fnat64_rev_egress_packets_icmp":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_packets_icmp,
			"fnat64_fwd_ingress_bytes_icmp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_bytes_icmp,
			"fnat64_fwd_egress_bytes_icmp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_bytes_icmp,
			"fnat64_rev_ingress_bytes_icmp":            ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_bytes_icmp,
			"fnat64_rev_egress_bytes_icmp":             ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_bytes_icmp,
			"fnat64_fwd_ingress_packets_others":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_packets_others,
			"fnat64_fwd_egress_packets_others":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_packets_others,
			"fnat64_rev_ingress_packets_others":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_packets_others,
			"fnat64_rev_egress_packets_others":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_packets_others,
			"fnat64_fwd_ingress_bytes_others":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_bytes_others,
			"fnat64_fwd_egress_bytes_others":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_bytes_others,
			"fnat64_rev_ingress_bytes_others":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_bytes_others,
			"fnat64_rev_egress_bytes_others":           ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_bytes_others,
			"fnat64_fwd_ingress_pkt_size_range1":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_pkt_size_range1,
			"fnat64_fwd_ingress_pkt_size_range2":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_pkt_size_range2,
			"fnat64_fwd_ingress_pkt_size_range3":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_pkt_size_range3,
			"fnat64_fwd_ingress_pkt_size_range4":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_ingress_pkt_size_range4,
			"fnat64_fwd_egress_pkt_size_range1":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_pkt_size_range1,
			"fnat64_fwd_egress_pkt_size_range2":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_pkt_size_range2,
			"fnat64_fwd_egress_pkt_size_range3":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_pkt_size_range3,
			"fnat64_fwd_egress_pkt_size_range4":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_fwd_egress_pkt_size_range4,
			"fnat64_rev_ingress_pkt_size_range1":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_pkt_size_range1,
			"fnat64_rev_ingress_pkt_size_range2":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_pkt_size_range2,
			"fnat64_rev_ingress_pkt_size_range3":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_pkt_size_range3,
			"fnat64_rev_ingress_pkt_size_range4":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_ingress_pkt_size_range4,
			"fnat64_rev_egress_pkt_size_range1":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_pkt_size_range1,
			"fnat64_rev_egress_pkt_size_range2":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_pkt_size_range2,
			"fnat64_rev_egress_pkt_size_range3":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_pkt_size_range3,
			"fnat64_rev_egress_pkt_size_range4":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnat64_rev_egress_pkt_size_range4,
			"fnatdslite_fwd_ingress_packets_tcp":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_packets_tcp,
			"fnatdslite_fwd_egress_packets_tcp":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_packets_tcp,
			"fnatdslite_rev_ingress_packets_tcp":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_packets_tcp,
			"fnatdslite_rev_egress_packets_tcp":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_packets_tcp,
			"fnatdslite_fwd_ingress_bytes_tcp":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_bytes_tcp,
			"fnatdslite_fwd_egress_bytes_tcp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_bytes_tcp,
			"fnatdslite_rev_ingress_bytes_tcp":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_bytes_tcp,
			"fnatdslite_rev_egress_bytes_tcp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_bytes_tcp,
			"fnatdslite_fwd_ingress_packets_udp":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_packets_udp,
			"fnatdslite_fwd_egress_packets_udp":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_packets_udp,
			"fnatdslite_rev_ingress_packets_udp":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_packets_udp,
			"fnatdslite_rev_egress_packets_udp":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_packets_udp,
			"fnatdslite_fwd_ingress_bytes_udp":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_bytes_udp,
			"fnatdslite_fwd_egress_bytes_udp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_bytes_udp,
			"fnatdslite_rev_ingress_bytes_udp":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_bytes_udp,
			"fnatdslite_rev_egress_bytes_udp":          ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_bytes_udp,
			"fnatdslite_fwd_ingress_packets_icmp":      ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_packets_icmp,
			"fnatdslite_fwd_egress_packets_icmp":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_packets_icmp,
			"fnatdslite_rev_ingress_packets_icmp":      ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_packets_icmp,
			"fnatdslite_rev_egress_packets_icmp":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_packets_icmp,
			"fnatdslite_fwd_ingress_bytes_icmp":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_bytes_icmp,
			"fnatdslite_fwd_egress_bytes_icmp":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_bytes_icmp,
			"fnatdslite_rev_ingress_bytes_icmp":        ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_bytes_icmp,
			"fnatdslite_rev_egress_bytes_icmp":         ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_bytes_icmp,
			"fnatdslite_fwd_ingress_packets_others":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_packets_others,
			"fnatdslite_fwd_egress_packets_others":     ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_packets_others,
			"fnatdslite_rev_ingress_packets_others":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_packets_others,
			"fnatdslite_rev_egress_packets_others":     ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_packets_others,
			"fnatdslite_fwd_ingress_bytes_others":      ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_bytes_others,
			"fnatdslite_fwd_egress_bytes_others":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_bytes_others,
			"fnatdslite_rev_ingress_bytes_others":      ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_bytes_others,
			"fnatdslite_rev_egress_bytes_others":       ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_bytes_others,
			"fnatdslite_fwd_ingress_pkt_size_range1":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_pkt_size_range1,
			"fnatdslite_fwd_ingress_pkt_size_range2":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_pkt_size_range2,
			"fnatdslite_fwd_ingress_pkt_size_range3":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_pkt_size_range3,
			"fnatdslite_fwd_ingress_pkt_size_range4":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_ingress_pkt_size_range4,
			"fnatdslite_fwd_egress_pkt_size_range1":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_pkt_size_range1,
			"fnatdslite_fwd_egress_pkt_size_range2":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_pkt_size_range2,
			"fnatdslite_fwd_egress_pkt_size_range3":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_pkt_size_range3,
			"fnatdslite_fwd_egress_pkt_size_range4":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_fwd_egress_pkt_size_range4,
			"fnatdslite_rev_ingress_pkt_size_range1":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_pkt_size_range1,
			"fnatdslite_rev_ingress_pkt_size_range2":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_pkt_size_range2,
			"fnatdslite_rev_ingress_pkt_size_range3":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_pkt_size_range3,
			"fnatdslite_rev_ingress_pkt_size_range4":   ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_ingress_pkt_size_range4,
			"fnatdslite_rev_egress_pkt_size_range1":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_pkt_size_range1,
			"fnatdslite_rev_egress_pkt_size_range2":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_pkt_size_range2,
			"fnatdslite_rev_egress_pkt_size_range3":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_pkt_size_range3,
			"fnatdslite_rev_egress_pkt_size_range4":    ret.DtCgnv6FixedNatGlobalStats.Stats.Fnatdslite_rev_egress_pkt_size_range4,
			"active_subscriber_added":                  ret.DtCgnv6FixedNatGlobalStats.Stats.ActiveSubscriberAdded,
			"active_subscriber_removed":                ret.DtCgnv6FixedNatGlobalStats.Stats.ActiveSubscriberRemoved,
		},
	}
}

func getObjectCgnv6FixedNatGlobalStatsStats(d []interface{}) edpt.Cgnv6FixedNatGlobalStatsStats {

	count1 := len(d)
	var ret edpt.Cgnv6FixedNatGlobalStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TotalNatInUse = in["total_nat_in_use"].(int)
		ret.TotalTcpAllocated = in["total_tcp_allocated"].(int)
		ret.TotalTcpFreed = in["total_tcp_freed"].(int)
		ret.TotalUdpAllocated = in["total_udp_allocated"].(int)
		ret.TotalUdpFreed = in["total_udp_freed"].(int)
		ret.TotalIcmpAllocated = in["total_icmp_allocated"].(int)
		ret.TotalIcmpFreed = in["total_icmp_freed"].(int)
		ret.Nat44DataSessionCreated = in["nat44_data_session_created"].(int)
		ret.Nat44DataSessionFreed = in["nat44_data_session_freed"].(int)
		ret.Nat64DataSessionCreated = in["nat64_data_session_created"].(int)
		ret.Nat64DataSessionFreed = in["nat64_data_session_freed"].(int)
		ret.DsliteDataSessionCreated = in["dslite_data_session_created"].(int)
		ret.DsliteDataSessionFreed = in["dslite_data_session_freed"].(int)
		ret.NatPortUnavailableTcp = in["nat_port_unavailable_tcp"].(int)
		ret.NatPortUnavailableUdp = in["nat_port_unavailable_udp"].(int)
		ret.NatPortUnavailableIcmp = in["nat_port_unavailable_icmp"].(int)
		ret.SessionUserQuotaExceeded = in["session_user_quota_exceeded"].(int)
		ret.Nat44TcpFullconeCreated = in["nat44_tcp_fullcone_created"].(int)
		ret.Nat44TcpFullconeFreed = in["nat44_tcp_fullcone_freed"].(int)
		ret.Nat44UdpFullconeCreated = in["nat44_udp_fullcone_created"].(int)
		ret.Nat44UdpFullconeFreed = in["nat44_udp_fullcone_freed"].(int)
		ret.Nat44UdpAlgFullconeCreated = in["nat44_udp_alg_fullcone_created"].(int)
		ret.Nat44UdpAlgFullconeFreed = in["nat44_udp_alg_fullcone_freed"].(int)
		ret.Nat64TcpFullconeCreated = in["nat64_tcp_fullcone_created"].(int)
		ret.Nat64TcpFullconeFreed = in["nat64_tcp_fullcone_freed"].(int)
		ret.Nat64UdpFullconeCreated = in["nat64_udp_fullcone_created"].(int)
		ret.Nat64UdpFullconeFreed = in["nat64_udp_fullcone_freed"].(int)
		ret.Nat64UdpAlgFullconeCreated = in["nat64_udp_alg_fullcone_created"].(int)
		ret.Nat64UdpAlgFullconeFreed = in["nat64_udp_alg_fullcone_freed"].(int)
		ret.DsliteTcpFullconeCreated = in["dslite_tcp_fullcone_created"].(int)
		ret.DsliteTcpFullconeFreed = in["dslite_tcp_fullcone_freed"].(int)
		ret.DsliteUdpFullconeCreated = in["dslite_udp_fullcone_created"].(int)
		ret.DsliteUdpFullconeFreed = in["dslite_udp_fullcone_freed"].(int)
		ret.DsliteUdpAlgFullconeCreated = in["dslite_udp_alg_fullcone_created"].(int)
		ret.DsliteUdpAlgFullconeFreed = in["dslite_udp_alg_fullcone_freed"].(int)
		ret.FullconeFailure = in["fullcone_failure"].(int)
		ret.Nat44EimMatch = in["nat44_eim_match"].(int)
		ret.Nat64EimMatch = in["nat64_eim_match"].(int)
		ret.DsliteEimMatch = in["dslite_eim_match"].(int)
		ret.Nat44EifMatch = in["nat44_eif_match"].(int)
		ret.Nat64EifMatch = in["nat64_eif_match"].(int)
		ret.DsliteEifMatch = in["dslite_eif_match"].(int)
		ret.Nat44InboundFiltered = in["nat44_inbound_filtered"].(int)
		ret.Nat64InboundFiltered = in["nat64_inbound_filtered"].(int)
		ret.DsliteInboundFiltered = in["dslite_inbound_filtered"].(int)
		ret.Nat44EifLimitExceeded = in["nat44_eif_limit_exceeded"].(int)
		ret.Nat64EifLimitExceeded = in["nat64_eif_limit_exceeded"].(int)
		ret.DsliteEifLimitExceeded = in["dslite_eif_limit_exceeded"].(int)
		ret.Nat44Hairpin = in["nat44_hairpin"].(int)
		ret.Nat64Hairpin = in["nat64_hairpin"].(int)
		ret.DsliteHairpin = in["dslite_hairpin"].(int)
		ret.StandbyDrop = in["standby_drop"].(int)
		ret.FixedNatFullconeSelfHairpinningDrop = in["fixed_nat_fullcone_self_hairpinning_drop"].(int)
		ret.SixrdDrop = in["sixrd_drop"].(int)
		ret.DestRlistDrop = in["dest_rlist_drop"].(int)
		ret.DestRlistPassThrough = in["dest_rlist_pass_through"].(int)
		ret.DestRlistSnatDrop = in["dest_rlist_snat_drop"].(int)
		ret.ConfigNotFound = in["config_not_found"].(int)
		ret.TotalTcpOverloadAcquired = in["total_tcp_overload_acquired"].(int)
		ret.TotalUdpOverloadAcquired = in["total_udp_overload_acquired"].(int)
		ret.TotalTcpOverloadReleased = in["total_tcp_overload_released"].(int)
		ret.TotalUdpOverloadReleased = in["total_udp_overload_released"].(int)
		ret.TotalTcpAllocOverload = in["total_tcp_alloc_overload"].(int)
		ret.TotalUdpAllocOverload = in["total_udp_alloc_overload"].(int)
		ret.TotalTcpFreeOverload = in["total_tcp_free_overload"].(int)
		ret.TotalUdpFreeOverload = in["total_udp_free_overload"].(int)
		ret.PortOverloadSmpDeleteScheduled = in["port_overload_smp_delete_scheduled"].(int)
		ret.PortOverloadFailed = in["port_overload_failed"].(int)
		ret.HaSessionUserQuotaExceeded = in["ha_session_user_quota_exceeded"].(int)
		ret.Fnat44_fwd_ingress_packets_tcp = in["fnat44_fwd_ingress_packets_tcp"].(int)
		ret.Fnat44_fwd_egress_packets_tcp = in["fnat44_fwd_egress_packets_tcp"].(int)
		ret.Fnat44_rev_ingress_packets_tcp = in["fnat44_rev_ingress_packets_tcp"].(int)
		ret.Fnat44_rev_egress_packets_tcp = in["fnat44_rev_egress_packets_tcp"].(int)
		ret.Fnat44_fwd_ingress_bytes_tcp = in["fnat44_fwd_ingress_bytes_tcp"].(int)
		ret.Fnat44_fwd_egress_bytes_tcp = in["fnat44_fwd_egress_bytes_tcp"].(int)
		ret.Fnat44_rev_ingress_bytes_tcp = in["fnat44_rev_ingress_bytes_tcp"].(int)
		ret.Fnat44_rev_egress_bytes_tcp = in["fnat44_rev_egress_bytes_tcp"].(int)
		ret.Fnat44_fwd_ingress_packets_udp = in["fnat44_fwd_ingress_packets_udp"].(int)
		ret.Fnat44_fwd_egress_packets_udp = in["fnat44_fwd_egress_packets_udp"].(int)
		ret.Fnat44_rev_ingress_packets_udp = in["fnat44_rev_ingress_packets_udp"].(int)
		ret.Fnat44_rev_egress_packets_udp = in["fnat44_rev_egress_packets_udp"].(int)
		ret.Fnat44_fwd_ingress_bytes_udp = in["fnat44_fwd_ingress_bytes_udp"].(int)
		ret.Fnat44_fwd_egress_bytes_udp = in["fnat44_fwd_egress_bytes_udp"].(int)
		ret.Fnat44_rev_ingress_bytes_udp = in["fnat44_rev_ingress_bytes_udp"].(int)
		ret.Fnat44_rev_egress_bytes_udp = in["fnat44_rev_egress_bytes_udp"].(int)
		ret.Fnat44_fwd_ingress_packets_icmp = in["fnat44_fwd_ingress_packets_icmp"].(int)
		ret.Fnat44_fwd_egress_packets_icmp = in["fnat44_fwd_egress_packets_icmp"].(int)
		ret.Fnat44_rev_ingress_packets_icmp = in["fnat44_rev_ingress_packets_icmp"].(int)
		ret.Fnat44_rev_egress_packets_icmp = in["fnat44_rev_egress_packets_icmp"].(int)
		ret.Fnat44_fwd_ingress_bytes_icmp = in["fnat44_fwd_ingress_bytes_icmp"].(int)
		ret.Fnat44_fwd_egress_bytes_icmp = in["fnat44_fwd_egress_bytes_icmp"].(int)
		ret.Fnat44_rev_ingress_bytes_icmp = in["fnat44_rev_ingress_bytes_icmp"].(int)
		ret.Fnat44_rev_egress_bytes_icmp = in["fnat44_rev_egress_bytes_icmp"].(int)
		ret.Fnat44_fwd_ingress_packets_others = in["fnat44_fwd_ingress_packets_others"].(int)
		ret.Fnat44_fwd_egress_packets_others = in["fnat44_fwd_egress_packets_others"].(int)
		ret.Fnat44_rev_ingress_packets_others = in["fnat44_rev_ingress_packets_others"].(int)
		ret.Fnat44_rev_egress_packets_others = in["fnat44_rev_egress_packets_others"].(int)
		ret.Fnat44_fwd_ingress_bytes_others = in["fnat44_fwd_ingress_bytes_others"].(int)
		ret.Fnat44_fwd_egress_bytes_others = in["fnat44_fwd_egress_bytes_others"].(int)
		ret.Fnat44_rev_ingress_bytes_others = in["fnat44_rev_ingress_bytes_others"].(int)
		ret.Fnat44_rev_egress_bytes_others = in["fnat44_rev_egress_bytes_others"].(int)
		ret.Fnat44_fwd_ingress_pkt_size_range1 = in["fnat44_fwd_ingress_pkt_size_range1"].(int)
		ret.Fnat44_fwd_ingress_pkt_size_range2 = in["fnat44_fwd_ingress_pkt_size_range2"].(int)
		ret.Fnat44_fwd_ingress_pkt_size_range3 = in["fnat44_fwd_ingress_pkt_size_range3"].(int)
		ret.Fnat44_fwd_ingress_pkt_size_range4 = in["fnat44_fwd_ingress_pkt_size_range4"].(int)
		ret.Fnat44_fwd_egress_pkt_size_range1 = in["fnat44_fwd_egress_pkt_size_range1"].(int)
		ret.Fnat44_fwd_egress_pkt_size_range2 = in["fnat44_fwd_egress_pkt_size_range2"].(int)
		ret.Fnat44_fwd_egress_pkt_size_range3 = in["fnat44_fwd_egress_pkt_size_range3"].(int)
		ret.Fnat44_fwd_egress_pkt_size_range4 = in["fnat44_fwd_egress_pkt_size_range4"].(int)
		ret.Fnat44_rev_ingress_pkt_size_range1 = in["fnat44_rev_ingress_pkt_size_range1"].(int)
		ret.Fnat44_rev_ingress_pkt_size_range2 = in["fnat44_rev_ingress_pkt_size_range2"].(int)
		ret.Fnat44_rev_ingress_pkt_size_range3 = in["fnat44_rev_ingress_pkt_size_range3"].(int)
		ret.Fnat44_rev_ingress_pkt_size_range4 = in["fnat44_rev_ingress_pkt_size_range4"].(int)
		ret.Fnat44_rev_egress_pkt_size_range1 = in["fnat44_rev_egress_pkt_size_range1"].(int)
		ret.Fnat44_rev_egress_pkt_size_range2 = in["fnat44_rev_egress_pkt_size_range2"].(int)
		ret.Fnat44_rev_egress_pkt_size_range3 = in["fnat44_rev_egress_pkt_size_range3"].(int)
		ret.Fnat44_rev_egress_pkt_size_range4 = in["fnat44_rev_egress_pkt_size_range4"].(int)
		ret.Fnat64_fwd_ingress_packets_tcp = in["fnat64_fwd_ingress_packets_tcp"].(int)
		ret.Fnat64_fwd_egress_packets_tcp = in["fnat64_fwd_egress_packets_tcp"].(int)
		ret.Fnat64_rev_ingress_packets_tcp = in["fnat64_rev_ingress_packets_tcp"].(int)
		ret.Fnat64_rev_egress_packets_tcp = in["fnat64_rev_egress_packets_tcp"].(int)
		ret.Fnat64_fwd_ingress_bytes_tcp = in["fnat64_fwd_ingress_bytes_tcp"].(int)
		ret.Fnat64_fwd_egress_bytes_tcp = in["fnat64_fwd_egress_bytes_tcp"].(int)
		ret.Fnat64_rev_ingress_bytes_tcp = in["fnat64_rev_ingress_bytes_tcp"].(int)
		ret.Fnat64_rev_egress_bytes_tcp = in["fnat64_rev_egress_bytes_tcp"].(int)
		ret.Fnat64_fwd_ingress_packets_udp = in["fnat64_fwd_ingress_packets_udp"].(int)
		ret.Fnat64_fwd_egress_packets_udp = in["fnat64_fwd_egress_packets_udp"].(int)
		ret.Fnat64_rev_ingress_packets_udp = in["fnat64_rev_ingress_packets_udp"].(int)
		ret.Fnat64_rev_egress_packets_udp = in["fnat64_rev_egress_packets_udp"].(int)
		ret.Fnat64_fwd_ingress_bytes_udp = in["fnat64_fwd_ingress_bytes_udp"].(int)
		ret.Fnat64_fwd_egress_bytes_udp = in["fnat64_fwd_egress_bytes_udp"].(int)
		ret.Fnat64_rev_ingress_bytes_udp = in["fnat64_rev_ingress_bytes_udp"].(int)
		ret.Fnat64_rev_egress_bytes_udp = in["fnat64_rev_egress_bytes_udp"].(int)
		ret.Fnat64_fwd_ingress_packets_icmp = in["fnat64_fwd_ingress_packets_icmp"].(int)
		ret.Fnat64_fwd_egress_packets_icmp = in["fnat64_fwd_egress_packets_icmp"].(int)
		ret.Fnat64_rev_ingress_packets_icmp = in["fnat64_rev_ingress_packets_icmp"].(int)
		ret.Fnat64_rev_egress_packets_icmp = in["fnat64_rev_egress_packets_icmp"].(int)
		ret.Fnat64_fwd_ingress_bytes_icmp = in["fnat64_fwd_ingress_bytes_icmp"].(int)
		ret.Fnat64_fwd_egress_bytes_icmp = in["fnat64_fwd_egress_bytes_icmp"].(int)
		ret.Fnat64_rev_ingress_bytes_icmp = in["fnat64_rev_ingress_bytes_icmp"].(int)
		ret.Fnat64_rev_egress_bytes_icmp = in["fnat64_rev_egress_bytes_icmp"].(int)
		ret.Fnat64_fwd_ingress_packets_others = in["fnat64_fwd_ingress_packets_others"].(int)
		ret.Fnat64_fwd_egress_packets_others = in["fnat64_fwd_egress_packets_others"].(int)
		ret.Fnat64_rev_ingress_packets_others = in["fnat64_rev_ingress_packets_others"].(int)
		ret.Fnat64_rev_egress_packets_others = in["fnat64_rev_egress_packets_others"].(int)
		ret.Fnat64_fwd_ingress_bytes_others = in["fnat64_fwd_ingress_bytes_others"].(int)
		ret.Fnat64_fwd_egress_bytes_others = in["fnat64_fwd_egress_bytes_others"].(int)
		ret.Fnat64_rev_ingress_bytes_others = in["fnat64_rev_ingress_bytes_others"].(int)
		ret.Fnat64_rev_egress_bytes_others = in["fnat64_rev_egress_bytes_others"].(int)
		ret.Fnat64_fwd_ingress_pkt_size_range1 = in["fnat64_fwd_ingress_pkt_size_range1"].(int)
		ret.Fnat64_fwd_ingress_pkt_size_range2 = in["fnat64_fwd_ingress_pkt_size_range2"].(int)
		ret.Fnat64_fwd_ingress_pkt_size_range3 = in["fnat64_fwd_ingress_pkt_size_range3"].(int)
		ret.Fnat64_fwd_ingress_pkt_size_range4 = in["fnat64_fwd_ingress_pkt_size_range4"].(int)
		ret.Fnat64_fwd_egress_pkt_size_range1 = in["fnat64_fwd_egress_pkt_size_range1"].(int)
		ret.Fnat64_fwd_egress_pkt_size_range2 = in["fnat64_fwd_egress_pkt_size_range2"].(int)
		ret.Fnat64_fwd_egress_pkt_size_range3 = in["fnat64_fwd_egress_pkt_size_range3"].(int)
		ret.Fnat64_fwd_egress_pkt_size_range4 = in["fnat64_fwd_egress_pkt_size_range4"].(int)
		ret.Fnat64_rev_ingress_pkt_size_range1 = in["fnat64_rev_ingress_pkt_size_range1"].(int)
		ret.Fnat64_rev_ingress_pkt_size_range2 = in["fnat64_rev_ingress_pkt_size_range2"].(int)
		ret.Fnat64_rev_ingress_pkt_size_range3 = in["fnat64_rev_ingress_pkt_size_range3"].(int)
		ret.Fnat64_rev_ingress_pkt_size_range4 = in["fnat64_rev_ingress_pkt_size_range4"].(int)
		ret.Fnat64_rev_egress_pkt_size_range1 = in["fnat64_rev_egress_pkt_size_range1"].(int)
		ret.Fnat64_rev_egress_pkt_size_range2 = in["fnat64_rev_egress_pkt_size_range2"].(int)
		ret.Fnat64_rev_egress_pkt_size_range3 = in["fnat64_rev_egress_pkt_size_range3"].(int)
		ret.Fnat64_rev_egress_pkt_size_range4 = in["fnat64_rev_egress_pkt_size_range4"].(int)
		ret.Fnatdslite_fwd_ingress_packets_tcp = in["fnatdslite_fwd_ingress_packets_tcp"].(int)
		ret.Fnatdslite_fwd_egress_packets_tcp = in["fnatdslite_fwd_egress_packets_tcp"].(int)
		ret.Fnatdslite_rev_ingress_packets_tcp = in["fnatdslite_rev_ingress_packets_tcp"].(int)
		ret.Fnatdslite_rev_egress_packets_tcp = in["fnatdslite_rev_egress_packets_tcp"].(int)
		ret.Fnatdslite_fwd_ingress_bytes_tcp = in["fnatdslite_fwd_ingress_bytes_tcp"].(int)
		ret.Fnatdslite_fwd_egress_bytes_tcp = in["fnatdslite_fwd_egress_bytes_tcp"].(int)
		ret.Fnatdslite_rev_ingress_bytes_tcp = in["fnatdslite_rev_ingress_bytes_tcp"].(int)
		ret.Fnatdslite_rev_egress_bytes_tcp = in["fnatdslite_rev_egress_bytes_tcp"].(int)
		ret.Fnatdslite_fwd_ingress_packets_udp = in["fnatdslite_fwd_ingress_packets_udp"].(int)
		ret.Fnatdslite_fwd_egress_packets_udp = in["fnatdslite_fwd_egress_packets_udp"].(int)
		ret.Fnatdslite_rev_ingress_packets_udp = in["fnatdslite_rev_ingress_packets_udp"].(int)
		ret.Fnatdslite_rev_egress_packets_udp = in["fnatdslite_rev_egress_packets_udp"].(int)
		ret.Fnatdslite_fwd_ingress_bytes_udp = in["fnatdslite_fwd_ingress_bytes_udp"].(int)
		ret.Fnatdslite_fwd_egress_bytes_udp = in["fnatdslite_fwd_egress_bytes_udp"].(int)
		ret.Fnatdslite_rev_ingress_bytes_udp = in["fnatdslite_rev_ingress_bytes_udp"].(int)
		ret.Fnatdslite_rev_egress_bytes_udp = in["fnatdslite_rev_egress_bytes_udp"].(int)
		ret.Fnatdslite_fwd_ingress_packets_icmp = in["fnatdslite_fwd_ingress_packets_icmp"].(int)
		ret.Fnatdslite_fwd_egress_packets_icmp = in["fnatdslite_fwd_egress_packets_icmp"].(int)
		ret.Fnatdslite_rev_ingress_packets_icmp = in["fnatdslite_rev_ingress_packets_icmp"].(int)
		ret.Fnatdslite_rev_egress_packets_icmp = in["fnatdslite_rev_egress_packets_icmp"].(int)
		ret.Fnatdslite_fwd_ingress_bytes_icmp = in["fnatdslite_fwd_ingress_bytes_icmp"].(int)
		ret.Fnatdslite_fwd_egress_bytes_icmp = in["fnatdslite_fwd_egress_bytes_icmp"].(int)
		ret.Fnatdslite_rev_ingress_bytes_icmp = in["fnatdslite_rev_ingress_bytes_icmp"].(int)
		ret.Fnatdslite_rev_egress_bytes_icmp = in["fnatdslite_rev_egress_bytes_icmp"].(int)
		ret.Fnatdslite_fwd_ingress_packets_others = in["fnatdslite_fwd_ingress_packets_others"].(int)
		ret.Fnatdslite_fwd_egress_packets_others = in["fnatdslite_fwd_egress_packets_others"].(int)
		ret.Fnatdslite_rev_ingress_packets_others = in["fnatdslite_rev_ingress_packets_others"].(int)
		ret.Fnatdslite_rev_egress_packets_others = in["fnatdslite_rev_egress_packets_others"].(int)
		ret.Fnatdslite_fwd_ingress_bytes_others = in["fnatdslite_fwd_ingress_bytes_others"].(int)
		ret.Fnatdslite_fwd_egress_bytes_others = in["fnatdslite_fwd_egress_bytes_others"].(int)
		ret.Fnatdslite_rev_ingress_bytes_others = in["fnatdslite_rev_ingress_bytes_others"].(int)
		ret.Fnatdslite_rev_egress_bytes_others = in["fnatdslite_rev_egress_bytes_others"].(int)
		ret.Fnatdslite_fwd_ingress_pkt_size_range1 = in["fnatdslite_fwd_ingress_pkt_size_range1"].(int)
		ret.Fnatdslite_fwd_ingress_pkt_size_range2 = in["fnatdslite_fwd_ingress_pkt_size_range2"].(int)
		ret.Fnatdslite_fwd_ingress_pkt_size_range3 = in["fnatdslite_fwd_ingress_pkt_size_range3"].(int)
		ret.Fnatdslite_fwd_ingress_pkt_size_range4 = in["fnatdslite_fwd_ingress_pkt_size_range4"].(int)
		ret.Fnatdslite_fwd_egress_pkt_size_range1 = in["fnatdslite_fwd_egress_pkt_size_range1"].(int)
		ret.Fnatdslite_fwd_egress_pkt_size_range2 = in["fnatdslite_fwd_egress_pkt_size_range2"].(int)
		ret.Fnatdslite_fwd_egress_pkt_size_range3 = in["fnatdslite_fwd_egress_pkt_size_range3"].(int)
		ret.Fnatdslite_fwd_egress_pkt_size_range4 = in["fnatdslite_fwd_egress_pkt_size_range4"].(int)
		ret.Fnatdslite_rev_ingress_pkt_size_range1 = in["fnatdslite_rev_ingress_pkt_size_range1"].(int)
		ret.Fnatdslite_rev_ingress_pkt_size_range2 = in["fnatdslite_rev_ingress_pkt_size_range2"].(int)
		ret.Fnatdslite_rev_ingress_pkt_size_range3 = in["fnatdslite_rev_ingress_pkt_size_range3"].(int)
		ret.Fnatdslite_rev_ingress_pkt_size_range4 = in["fnatdslite_rev_ingress_pkt_size_range4"].(int)
		ret.Fnatdslite_rev_egress_pkt_size_range1 = in["fnatdslite_rev_egress_pkt_size_range1"].(int)
		ret.Fnatdslite_rev_egress_pkt_size_range2 = in["fnatdslite_rev_egress_pkt_size_range2"].(int)
		ret.Fnatdslite_rev_egress_pkt_size_range3 = in["fnatdslite_rev_egress_pkt_size_range3"].(int)
		ret.Fnatdslite_rev_egress_pkt_size_range4 = in["fnatdslite_rev_egress_pkt_size_range4"].(int)
		ret.ActiveSubscriberAdded = in["active_subscriber_added"].(int)
		ret.ActiveSubscriberRemoved = in["active_subscriber_removed"].(int)
	}
	return ret
}

func dataToEndpointCgnv6FixedNatGlobalStats(d *schema.ResourceData) edpt.Cgnv6FixedNatGlobalStats {
	var ret edpt.Cgnv6FixedNatGlobalStats

	ret.Stats = getObjectCgnv6FixedNatGlobalStatsStats(d.Get("stats").([]interface{}))
	return ret
}
