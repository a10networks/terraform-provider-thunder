package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbSwitchStats() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_switch_stats`: Statistics for the object switch\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbSwitchStatsRead,

		Schema: map[string]*schema.Schema{
			"stats": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fwlb": {
							Type: schema.TypeInt, Optional: true, Description: "FWLB",
						},
						"licexpire_drop": {
							Type: schema.TypeInt, Optional: true, Description: "License Expire Drop",
						},
						"bwl_drop": {
							Type: schema.TypeInt, Optional: true, Description: "BW Limit Drop",
						},
						"rx_kernel": {
							Type: schema.TypeInt, Optional: true, Description: "Received kernel",
						},
						"rx_arp_req": {
							Type: schema.TypeInt, Optional: true, Description: "ARP REQ Rcvd",
						},
						"rx_arp_resp": {
							Type: schema.TypeInt, Optional: true, Description: "ARP RESP Rcvd",
						},
						"vlan_flood": {
							Type: schema.TypeInt, Optional: true, Description: "VLAN Flood",
						},
						"l2_def_vlan_drop": {
							Type: schema.TypeInt, Optional: true, Description: "L2 Default Vlan FWD Drop",
						},
						"ipv4_noroute_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 No Route Drop",
						},
						"ipv6_noroute_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 No Route Drop",
						},
						"prot_down_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Prot Down Drop",
						},
						"l2_forward": {
							Type: schema.TypeInt, Optional: true, Description: "L2 Forward",
						},
						"l3_forward_ip": {
							Type: schema.TypeInt, Optional: true, Description: "L3 IP Forward",
						},
						"l3_forward_ipv6": {
							Type: schema.TypeInt, Optional: true, Description: "L3 IPv6 Forward",
						},
						"l4_process": {
							Type: schema.TypeInt, Optional: true, Description: "L4 Process",
						},
						"unknown_prot_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Unknown Prot Drop",
						},
						"ttl_exceeded_drop": {
							Type: schema.TypeInt, Optional: true, Description: "TTL Exceeded Drop",
						},
						"linkdown_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Link Down Drop",
						},
						"sport_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SPORT Drop",
						},
						"incorrect_len_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Incorrect Length Drop",
						},
						"ip_defrag": {
							Type: schema.TypeInt, Optional: true, Description: "IP Defrag",
						},
						"acl_deny": {
							Type: schema.TypeInt, Optional: true, Description: "ACL Denys",
						},
						"ipfrag_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "IP(TCP) Fragment Rcvd",
						},
						"ipfrag_overlap": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment Overlap",
						},
						"ipfrag_timeout": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment Timeout",
						},
						"ipfrag_overload": {
							Type: schema.TypeInt, Optional: true, Description: "IP Frag Overload Drops",
						},
						"ipfrag_reasmoks": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment Reasm OKs",
						},
						"ipfrag_reasmfails": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment Reasm Fails",
						},
						"badpkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Bad Pkt Drop",
						},
						"ipsec_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPSec Drop",
						},
						"bpdu_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "BPDUs Received",
						},
						"bpdu_sent": {
							Type: schema.TypeInt, Optional: true, Description: "BPDUs Sent",
						},
						"ctrl_syn_rate_drop": {
							Type: schema.TypeInt, Optional: true, Description: "SYN rate exceeded Drop",
						},
						"ip_defrag_invalid_len": {
							Type: schema.TypeInt, Optional: true, Description: "IP Invalid Length Frag",
						},
						"ipv4_frag_6rd_ok": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Frag 6RD OK",
						},
						"ipv4_frag_6rd_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 Frag 6RD Dropped",
						},
						"no_ip_drop": {
							Type: schema.TypeInt, Optional: true, Description: "No IP Drop",
						},
						"ipv6frag_udp": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag UDP",
						},
						"ipv6frag_udp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag UDP Dropped",
						},
						"ipv6frag_tcp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag TCP Dropped",
						},
						"ipv6frag_ipip_ok": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag IPIP OKs",
						},
						"ipv6frag_ipip_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag IPIP Drop",
						},
						"ip_frag_oversize": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment oversize",
						},
						"ip_frag_too_many": {
							Type: schema.TypeInt, Optional: true, Description: "IP Fragment too many",
						},
						"ipv4_novlanfwd_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv4 No L3 VLAN FWD Drop",
						},
						"ipv6_novlanfwd_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 No L3 VLAN FWD Drop",
						},
						"fpga_error_pkt1": {
							Type: schema.TypeInt, Optional: true, Description: "FPGA Error PKT1",
						},
						"fpga_error_pkt2": {
							Type: schema.TypeInt, Optional: true, Description: "FPGA Error PKT2",
						},
						"max_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Max ARP Drop",
						},
						"ipv6frag_tcp": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag TCP",
						},
						"ipv6frag_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag ICMP",
						},
						"ipv6frag_ospf": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag OSPF",
						},
						"ipv6frag_esp": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag ESP",
						},
						"l4_in_ctrl_cpu": {
							Type: schema.TypeInt, Optional: true, Description: "L4 In Ctrl CPU",
						},
						"mgmt_svc_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Management Service Drop",
						},
						"jumbo_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Jumbo Frag Drop",
						},
						"ipv6_jumbo_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Jumbo Frag Drop",
						},
						"ipipv6_jumbo_frag_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPIPv6 Jumbo Frag Drop",
						},
						"ipv6_ndisc_dad_solicits": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 DAD on Solicits",
						},
						"ipv6_ndisc_dad_adverts": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 DAD on Adverts",
						},
						"ipv6_ndisc_mac_changes": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 DAD MAC Changed",
						},
						"ipv6_ndisc_out_of_memory": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 DAD Out-of-memory",
						},
						"sp_non_ctrl_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Shared IP mode non ctrl packet to linux drop",
						},
						"urpf_pkt_drop": {
							Type: schema.TypeInt, Optional: true, Description: "URPF check packet drop",
						},
						"fw_smp_zone_mismatch": {
							Type: schema.TypeInt, Optional: true, Description: "FW SMP Zone Mismatch",
						},
						"ipfrag_udp": {
							Type: schema.TypeInt, Optional: true, Description: "IP(UDP) Fragment Rcvd",
						},
						"ipfrag_icmp": {
							Type: schema.TypeInt, Optional: true, Description: "IP(ICMP) Fragment Rcvd",
						},
						"ipfrag_ospf": {
							Type: schema.TypeInt, Optional: true, Description: "IP(OSPF) Fragment Rcvd",
						},
						"ipfrag_esp": {
							Type: schema.TypeInt, Optional: true, Description: "IP(ESP) Fragment Rcvd",
						},
						"ipfrag_tcp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IP Frag TCP Dropped",
						},
						"ipfrag_udp_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IP Frag UDP Dropped",
						},
						"ipfrag_ipip_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IP Frag IPIP Drop",
						},
						"redirect_fwd_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Redirect failed in the fwd direction",
						},
						"redirect_fwd_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Redirect succeeded in the fwd direction",
						},
						"redirect_rev_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Redirect failed in the rev direction",
						},
						"redirect_rev_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Redirect succeeded in the rev direction",
						},
						"redirect_setup_fail": {
							Type: schema.TypeInt, Optional: true, Description: "Redirect connection setup failed",
						},
						"ip_frag_sent": {
							Type: schema.TypeInt, Optional: true, Description: "IP frag sent",
						},
						"invalid_rx_arp_pkt": {
							Type: schema.TypeInt, Optional: true, Description: "Invalid ARP PKT Rcvd",
						},
						"invalid_sender_mac_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to invalid sender MAC",
						},
						"dev_based_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to interface state checks",
						},
						"scaleout_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to scaleout checks",
						},
						"virtual_ip_not_found_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to virtual IP not found",
						},
						"inactive_static_nat_pool_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to inactive static nat pool",
						},
						"inactive_nat_pool_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to inactive nat pool",
						},
						"scaleout_hairpin_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to scaleout hairpin checks",
						},
						"self_grat_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Self generated grat ARP PKT dropped",
						},
						"self_grat_nat_ip_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Self generated grat ARP PKT dropped for NAT IP",
						},
						"ip_not_found_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to IP not found",
						},
						"dev_link_down_arp_drop": {
							Type: schema.TypeInt, Optional: true, Description: "ARP PKT dropped due to interface is down",
						},
						"lacp_tx_intf_err_drop": {
							Type: schema.TypeInt, Optional: true, Description: "LACP interface error corrected",
						},
						"service_chain_sent": {
							Type: schema.TypeInt, Optional: true, Description: "Service Chain Packets Sent",
						},
						"service_chain_rcvd": {
							Type: schema.TypeInt, Optional: true, Description: "Service Chain Packets Rcvd",
						},
						"unnumbered_nat_error": {
							Type: schema.TypeInt, Optional: true, Description: "Unnumbered NAT error",
						},
						"unnumbered_unsupported_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Unsupported protocol for unnumbered",
						},
						"ipv6frag_gre_dropped": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 Frag gre Drop",
						},
						"ipv6_ndisc_dad_prefix_mismatch_drop": {
							Type: schema.TypeInt, Optional: true, Description: "IPv6 DAD on Advertise drop for prefix mismatch",
						},
						"bw_ignore_limit": {
							Type: schema.TypeInt, Optional: true, Description: "BW Limit ignored packets count",
						},
						"ppsl_drop_egr": {
							Type: schema.TypeInt, Optional: true, Description: "Packet-Per-Sec Limit Drop at egress",
						},
						"ppsl_drop_ing": {
							Type: schema.TypeInt, Optional: true, Description: "Packet-Per-Sec Limit Drop at ingress",
						},
						"ppsl_ignore_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Packet-Per-Sec Limit ignored packets count",
						},
						"closed_port_syn_drop": {
							Type: schema.TypeInt, Optional: true, Description: "Linux Closed Port SYN Drop",
						},
						"tls13_drop_req": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13-Request-Per-Sec Limit Drop at ingress",
						},
						"tls13_ignore_req": {
							Type: schema.TypeInt, Optional: true, Description: "TLS13-Request-Per-Sec Limit ignored packets count",
						},
						"tls12_drop_req": {
							Type: schema.TypeInt, Optional: true, Description: "TLS12-Request-Per-Sec Limit Drop at ingress",
						},
						"tls12_ignore_req": {
							Type: schema.TypeInt, Optional: true, Description: "TLS12-Request-Per-Sec Limit ignored packets count",
						},
						"tls12_tls13_drop_req": {
							Type: schema.TypeInt, Optional: true, Description: "TLS12--TLS13-Request-Per-Sec Limit Drop at ingress",
						},
						"tls12_tls13_ignore_req": {
							Type: schema.TypeInt, Optional: true, Description: "TLS12-TLS13-Request-Per-Sec Limit ignored packets count",
						},
					},
				},
			},
		},
	}
}

func resourceSlbSwitchStatsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbSwitchStatsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbSwitchStats(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbSwitchStatsStats := setObjectSlbSwitchStatsStats(res)
		d.Set("stats", SlbSwitchStatsStats)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbSwitchStatsStats(ret edpt.DataSlbSwitchStats) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"fwlb":                                ret.DtSlbSwitchStats.Stats.Fwlb,
			"licexpire_drop":                      ret.DtSlbSwitchStats.Stats.Licexpire_drop,
			"bwl_drop":                            ret.DtSlbSwitchStats.Stats.Bwl_drop,
			"rx_kernel":                           ret.DtSlbSwitchStats.Stats.Rx_kernel,
			"rx_arp_req":                          ret.DtSlbSwitchStats.Stats.Rx_arp_req,
			"rx_arp_resp":                         ret.DtSlbSwitchStats.Stats.Rx_arp_resp,
			"vlan_flood":                          ret.DtSlbSwitchStats.Stats.Vlan_flood,
			"l2_def_vlan_drop":                    ret.DtSlbSwitchStats.Stats.L2_def_vlan_drop,
			"ipv4_noroute_drop":                   ret.DtSlbSwitchStats.Stats.Ipv4_noroute_drop,
			"ipv6_noroute_drop":                   ret.DtSlbSwitchStats.Stats.Ipv6_noroute_drop,
			"prot_down_drop":                      ret.DtSlbSwitchStats.Stats.Prot_down_drop,
			"l2_forward":                          ret.DtSlbSwitchStats.Stats.L2_forward,
			"l3_forward_ip":                       ret.DtSlbSwitchStats.Stats.L3_forward_ip,
			"l3_forward_ipv6":                     ret.DtSlbSwitchStats.Stats.L3_forward_ipv6,
			"l4_process":                          ret.DtSlbSwitchStats.Stats.L4_process,
			"unknown_prot_drop":                   ret.DtSlbSwitchStats.Stats.Unknown_prot_drop,
			"ttl_exceeded_drop":                   ret.DtSlbSwitchStats.Stats.Ttl_exceeded_drop,
			"linkdown_drop":                       ret.DtSlbSwitchStats.Stats.Linkdown_drop,
			"sport_drop":                          ret.DtSlbSwitchStats.Stats.Sport_drop,
			"incorrect_len_drop":                  ret.DtSlbSwitchStats.Stats.Incorrect_len_drop,
			"ip_defrag":                           ret.DtSlbSwitchStats.Stats.Ip_defrag,
			"acl_deny":                            ret.DtSlbSwitchStats.Stats.Acl_deny,
			"ipfrag_tcp":                          ret.DtSlbSwitchStats.Stats.Ipfrag_tcp,
			"ipfrag_overlap":                      ret.DtSlbSwitchStats.Stats.Ipfrag_overlap,
			"ipfrag_timeout":                      ret.DtSlbSwitchStats.Stats.Ipfrag_timeout,
			"ipfrag_overload":                     ret.DtSlbSwitchStats.Stats.Ipfrag_overload,
			"ipfrag_reasmoks":                     ret.DtSlbSwitchStats.Stats.Ipfrag_reasmoks,
			"ipfrag_reasmfails":                   ret.DtSlbSwitchStats.Stats.Ipfrag_reasmfails,
			"badpkt_drop":                         ret.DtSlbSwitchStats.Stats.Badpkt_drop,
			"ipsec_drop":                          ret.DtSlbSwitchStats.Stats.Ipsec_drop,
			"bpdu_rcvd":                           ret.DtSlbSwitchStats.Stats.Bpdu_rcvd,
			"bpdu_sent":                           ret.DtSlbSwitchStats.Stats.Bpdu_sent,
			"ctrl_syn_rate_drop":                  ret.DtSlbSwitchStats.Stats.Ctrl_syn_rate_drop,
			"ip_defrag_invalid_len":               ret.DtSlbSwitchStats.Stats.Ip_defrag_invalid_len,
			"ipv4_frag_6rd_ok":                    ret.DtSlbSwitchStats.Stats.Ipv4_frag_6rd_ok,
			"ipv4_frag_6rd_drop":                  ret.DtSlbSwitchStats.Stats.Ipv4_frag_6rd_drop,
			"no_ip_drop":                          ret.DtSlbSwitchStats.Stats.No_ip_drop,
			"ipv6frag_udp":                        ret.DtSlbSwitchStats.Stats.Ipv6frag_udp,
			"ipv6frag_udp_dropped":                ret.DtSlbSwitchStats.Stats.Ipv6frag_udp_dropped,
			"ipv6frag_tcp_dropped":                ret.DtSlbSwitchStats.Stats.Ipv6frag_tcp_dropped,
			"ipv6frag_ipip_ok":                    ret.DtSlbSwitchStats.Stats.Ipv6frag_ipip_ok,
			"ipv6frag_ipip_dropped":               ret.DtSlbSwitchStats.Stats.Ipv6frag_ipip_dropped,
			"ip_frag_oversize":                    ret.DtSlbSwitchStats.Stats.Ip_frag_oversize,
			"ip_frag_too_many":                    ret.DtSlbSwitchStats.Stats.Ip_frag_too_many,
			"ipv4_novlanfwd_drop":                 ret.DtSlbSwitchStats.Stats.Ipv4_novlanfwd_drop,
			"ipv6_novlanfwd_drop":                 ret.DtSlbSwitchStats.Stats.Ipv6_novlanfwd_drop,
			"fpga_error_pkt1":                     ret.DtSlbSwitchStats.Stats.Fpga_error_pkt1,
			"fpga_error_pkt2":                     ret.DtSlbSwitchStats.Stats.Fpga_error_pkt2,
			"max_arp_drop":                        ret.DtSlbSwitchStats.Stats.Max_arp_drop,
			"ipv6frag_tcp":                        ret.DtSlbSwitchStats.Stats.Ipv6frag_tcp,
			"ipv6frag_icmp":                       ret.DtSlbSwitchStats.Stats.Ipv6frag_icmp,
			"ipv6frag_ospf":                       ret.DtSlbSwitchStats.Stats.Ipv6frag_ospf,
			"ipv6frag_esp":                        ret.DtSlbSwitchStats.Stats.Ipv6frag_esp,
			"l4_in_ctrl_cpu":                      ret.DtSlbSwitchStats.Stats.L4_in_ctrl_cpu,
			"mgmt_svc_drop":                       ret.DtSlbSwitchStats.Stats.Mgmt_svc_drop,
			"jumbo_frag_drop":                     ret.DtSlbSwitchStats.Stats.Jumbo_frag_drop,
			"ipv6_jumbo_frag_drop":                ret.DtSlbSwitchStats.Stats.Ipv6_jumbo_frag_drop,
			"ipipv6_jumbo_frag_drop":              ret.DtSlbSwitchStats.Stats.Ipipv6_jumbo_frag_drop,
			"ipv6_ndisc_dad_solicits":             ret.DtSlbSwitchStats.Stats.Ipv6_ndisc_dad_solicits,
			"ipv6_ndisc_dad_adverts":              ret.DtSlbSwitchStats.Stats.Ipv6_ndisc_dad_adverts,
			"ipv6_ndisc_mac_changes":              ret.DtSlbSwitchStats.Stats.Ipv6_ndisc_mac_changes,
			"ipv6_ndisc_out_of_memory":            ret.DtSlbSwitchStats.Stats.Ipv6_ndisc_out_of_memory,
			"sp_non_ctrl_pkt_drop":                ret.DtSlbSwitchStats.Stats.Sp_non_ctrl_pkt_drop,
			"urpf_pkt_drop":                       ret.DtSlbSwitchStats.Stats.Urpf_pkt_drop,
			"fw_smp_zone_mismatch":                ret.DtSlbSwitchStats.Stats.Fw_smp_zone_mismatch,
			"ipfrag_udp":                          ret.DtSlbSwitchStats.Stats.Ipfrag_udp,
			"ipfrag_icmp":                         ret.DtSlbSwitchStats.Stats.Ipfrag_icmp,
			"ipfrag_ospf":                         ret.DtSlbSwitchStats.Stats.Ipfrag_ospf,
			"ipfrag_esp":                          ret.DtSlbSwitchStats.Stats.Ipfrag_esp,
			"ipfrag_tcp_dropped":                  ret.DtSlbSwitchStats.Stats.Ipfrag_tcp_dropped,
			"ipfrag_udp_dropped":                  ret.DtSlbSwitchStats.Stats.Ipfrag_udp_dropped,
			"ipfrag_ipip_dropped":                 ret.DtSlbSwitchStats.Stats.Ipfrag_ipip_dropped,
			"redirect_fwd_fail":                   ret.DtSlbSwitchStats.Stats.Redirect_fwd_fail,
			"redirect_fwd_sent":                   ret.DtSlbSwitchStats.Stats.Redirect_fwd_sent,
			"redirect_rev_fail":                   ret.DtSlbSwitchStats.Stats.Redirect_rev_fail,
			"redirect_rev_sent":                   ret.DtSlbSwitchStats.Stats.Redirect_rev_sent,
			"redirect_setup_fail":                 ret.DtSlbSwitchStats.Stats.Redirect_setup_fail,
			"ip_frag_sent":                        ret.DtSlbSwitchStats.Stats.Ip_frag_sent,
			"invalid_rx_arp_pkt":                  ret.DtSlbSwitchStats.Stats.Invalid_rx_arp_pkt,
			"invalid_sender_mac_arp_drop":         ret.DtSlbSwitchStats.Stats.Invalid_sender_mac_arp_drop,
			"dev_based_arp_drop":                  ret.DtSlbSwitchStats.Stats.Dev_based_arp_drop,
			"scaleout_arp_drop":                   ret.DtSlbSwitchStats.Stats.Scaleout_arp_drop,
			"virtual_ip_not_found_arp_drop":       ret.DtSlbSwitchStats.Stats.Virtual_ip_not_found_arp_drop,
			"inactive_static_nat_pool_arp_drop":   ret.DtSlbSwitchStats.Stats.Inactive_static_nat_pool_arp_drop,
			"inactive_nat_pool_arp_drop":          ret.DtSlbSwitchStats.Stats.Inactive_nat_pool_arp_drop,
			"scaleout_hairpin_arp_drop":           ret.DtSlbSwitchStats.Stats.Scaleout_hairpin_arp_drop,
			"self_grat_arp_drop":                  ret.DtSlbSwitchStats.Stats.Self_grat_arp_drop,
			"self_grat_nat_ip_arp_drop":           ret.DtSlbSwitchStats.Stats.Self_grat_nat_ip_arp_drop,
			"ip_not_found_arp_drop":               ret.DtSlbSwitchStats.Stats.Ip_not_found_arp_drop,
			"dev_link_down_arp_drop":              ret.DtSlbSwitchStats.Stats.Dev_link_down_arp_drop,
			"lacp_tx_intf_err_drop":               ret.DtSlbSwitchStats.Stats.Lacp_tx_intf_err_drop,
			"service_chain_sent":                  ret.DtSlbSwitchStats.Stats.Service_chain_sent,
			"service_chain_rcvd":                  ret.DtSlbSwitchStats.Stats.Service_chain_rcvd,
			"unnumbered_nat_error":                ret.DtSlbSwitchStats.Stats.Unnumbered_nat_error,
			"unnumbered_unsupported_drop":         ret.DtSlbSwitchStats.Stats.Unnumbered_unsupported_drop,
			"ipv6frag_gre_dropped":                ret.DtSlbSwitchStats.Stats.Ipv6frag_gre_dropped,
			"ipv6_ndisc_dad_prefix_mismatch_drop": ret.DtSlbSwitchStats.Stats.Ipv6_ndisc_dad_prefix_mismatch_drop,
			"bw_ignore_limit":                     ret.DtSlbSwitchStats.Stats.Bw_ignore_limit,
			"ppsl_drop_egr":                       ret.DtSlbSwitchStats.Stats.Ppsl_drop_egr,
			"ppsl_drop_ing":                       ret.DtSlbSwitchStats.Stats.Ppsl_drop_ing,
			"ppsl_ignore_limit":                   ret.DtSlbSwitchStats.Stats.Ppsl_ignore_limit,
			"closed_port_syn_drop":                ret.DtSlbSwitchStats.Stats.Closed_port_syn_drop,
			"tls13_drop_req":                      ret.DtSlbSwitchStats.Stats.Tls13_drop_req,
			"tls13_ignore_req":                    ret.DtSlbSwitchStats.Stats.Tls13_ignore_req,
			"tls12_drop_req":                      ret.DtSlbSwitchStats.Stats.Tls12_drop_req,
			"tls12_ignore_req":                    ret.DtSlbSwitchStats.Stats.Tls12_ignore_req,
			"tls12_tls13_drop_req":                ret.DtSlbSwitchStats.Stats.Tls12_tls13_drop_req,
			"tls12_tls13_ignore_req":              ret.DtSlbSwitchStats.Stats.Tls12_tls13_ignore_req,
		},
	}
}

func getObjectSlbSwitchStatsStats(d []interface{}) edpt.SlbSwitchStatsStats {

	count1 := len(d)
	var ret edpt.SlbSwitchStatsStats
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Fwlb = in["fwlb"].(int)
		ret.Licexpire_drop = in["licexpire_drop"].(int)
		ret.Bwl_drop = in["bwl_drop"].(int)
		ret.Rx_kernel = in["rx_kernel"].(int)
		ret.Rx_arp_req = in["rx_arp_req"].(int)
		ret.Rx_arp_resp = in["rx_arp_resp"].(int)
		ret.Vlan_flood = in["vlan_flood"].(int)
		ret.L2_def_vlan_drop = in["l2_def_vlan_drop"].(int)
		ret.Ipv4_noroute_drop = in["ipv4_noroute_drop"].(int)
		ret.Ipv6_noroute_drop = in["ipv6_noroute_drop"].(int)
		ret.Prot_down_drop = in["prot_down_drop"].(int)
		ret.L2_forward = in["l2_forward"].(int)
		ret.L3_forward_ip = in["l3_forward_ip"].(int)
		ret.L3_forward_ipv6 = in["l3_forward_ipv6"].(int)
		ret.L4_process = in["l4_process"].(int)
		ret.Unknown_prot_drop = in["unknown_prot_drop"].(int)
		ret.Ttl_exceeded_drop = in["ttl_exceeded_drop"].(int)
		ret.Linkdown_drop = in["linkdown_drop"].(int)
		ret.Sport_drop = in["sport_drop"].(int)
		ret.Incorrect_len_drop = in["incorrect_len_drop"].(int)
		ret.Ip_defrag = in["ip_defrag"].(int)
		ret.Acl_deny = in["acl_deny"].(int)
		ret.Ipfrag_tcp = in["ipfrag_tcp"].(int)
		ret.Ipfrag_overlap = in["ipfrag_overlap"].(int)
		ret.Ipfrag_timeout = in["ipfrag_timeout"].(int)
		ret.Ipfrag_overload = in["ipfrag_overload"].(int)
		ret.Ipfrag_reasmoks = in["ipfrag_reasmoks"].(int)
		ret.Ipfrag_reasmfails = in["ipfrag_reasmfails"].(int)
		ret.Badpkt_drop = in["badpkt_drop"].(int)
		ret.Ipsec_drop = in["ipsec_drop"].(int)
		ret.Bpdu_rcvd = in["bpdu_rcvd"].(int)
		ret.Bpdu_sent = in["bpdu_sent"].(int)
		ret.Ctrl_syn_rate_drop = in["ctrl_syn_rate_drop"].(int)
		ret.Ip_defrag_invalid_len = in["ip_defrag_invalid_len"].(int)
		ret.Ipv4_frag_6rd_ok = in["ipv4_frag_6rd_ok"].(int)
		ret.Ipv4_frag_6rd_drop = in["ipv4_frag_6rd_drop"].(int)
		ret.No_ip_drop = in["no_ip_drop"].(int)
		ret.Ipv6frag_udp = in["ipv6frag_udp"].(int)
		ret.Ipv6frag_udp_dropped = in["ipv6frag_udp_dropped"].(int)
		ret.Ipv6frag_tcp_dropped = in["ipv6frag_tcp_dropped"].(int)
		ret.Ipv6frag_ipip_ok = in["ipv6frag_ipip_ok"].(int)
		ret.Ipv6frag_ipip_dropped = in["ipv6frag_ipip_dropped"].(int)
		ret.Ip_frag_oversize = in["ip_frag_oversize"].(int)
		ret.Ip_frag_too_many = in["ip_frag_too_many"].(int)
		ret.Ipv4_novlanfwd_drop = in["ipv4_novlanfwd_drop"].(int)
		ret.Ipv6_novlanfwd_drop = in["ipv6_novlanfwd_drop"].(int)
		ret.Fpga_error_pkt1 = in["fpga_error_pkt1"].(int)
		ret.Fpga_error_pkt2 = in["fpga_error_pkt2"].(int)
		ret.Max_arp_drop = in["max_arp_drop"].(int)
		ret.Ipv6frag_tcp = in["ipv6frag_tcp"].(int)
		ret.Ipv6frag_icmp = in["ipv6frag_icmp"].(int)
		ret.Ipv6frag_ospf = in["ipv6frag_ospf"].(int)
		ret.Ipv6frag_esp = in["ipv6frag_esp"].(int)
		ret.L4_in_ctrl_cpu = in["l4_in_ctrl_cpu"].(int)
		ret.Mgmt_svc_drop = in["mgmt_svc_drop"].(int)
		ret.Jumbo_frag_drop = in["jumbo_frag_drop"].(int)
		ret.Ipv6_jumbo_frag_drop = in["ipv6_jumbo_frag_drop"].(int)
		ret.Ipipv6_jumbo_frag_drop = in["ipipv6_jumbo_frag_drop"].(int)
		ret.Ipv6_ndisc_dad_solicits = in["ipv6_ndisc_dad_solicits"].(int)
		ret.Ipv6_ndisc_dad_adverts = in["ipv6_ndisc_dad_adverts"].(int)
		ret.Ipv6_ndisc_mac_changes = in["ipv6_ndisc_mac_changes"].(int)
		ret.Ipv6_ndisc_out_of_memory = in["ipv6_ndisc_out_of_memory"].(int)
		ret.Sp_non_ctrl_pkt_drop = in["sp_non_ctrl_pkt_drop"].(int)
		ret.Urpf_pkt_drop = in["urpf_pkt_drop"].(int)
		ret.Fw_smp_zone_mismatch = in["fw_smp_zone_mismatch"].(int)
		ret.Ipfrag_udp = in["ipfrag_udp"].(int)
		ret.Ipfrag_icmp = in["ipfrag_icmp"].(int)
		ret.Ipfrag_ospf = in["ipfrag_ospf"].(int)
		ret.Ipfrag_esp = in["ipfrag_esp"].(int)
		ret.Ipfrag_tcp_dropped = in["ipfrag_tcp_dropped"].(int)
		ret.Ipfrag_udp_dropped = in["ipfrag_udp_dropped"].(int)
		ret.Ipfrag_ipip_dropped = in["ipfrag_ipip_dropped"].(int)
		ret.Redirect_fwd_fail = in["redirect_fwd_fail"].(int)
		ret.Redirect_fwd_sent = in["redirect_fwd_sent"].(int)
		ret.Redirect_rev_fail = in["redirect_rev_fail"].(int)
		ret.Redirect_rev_sent = in["redirect_rev_sent"].(int)
		ret.Redirect_setup_fail = in["redirect_setup_fail"].(int)
		ret.Ip_frag_sent = in["ip_frag_sent"].(int)
		ret.Invalid_rx_arp_pkt = in["invalid_rx_arp_pkt"].(int)
		ret.Invalid_sender_mac_arp_drop = in["invalid_sender_mac_arp_drop"].(int)
		ret.Dev_based_arp_drop = in["dev_based_arp_drop"].(int)
		ret.Scaleout_arp_drop = in["scaleout_arp_drop"].(int)
		ret.Virtual_ip_not_found_arp_drop = in["virtual_ip_not_found_arp_drop"].(int)
		ret.Inactive_static_nat_pool_arp_drop = in["inactive_static_nat_pool_arp_drop"].(int)
		ret.Inactive_nat_pool_arp_drop = in["inactive_nat_pool_arp_drop"].(int)
		ret.Scaleout_hairpin_arp_drop = in["scaleout_hairpin_arp_drop"].(int)
		ret.Self_grat_arp_drop = in["self_grat_arp_drop"].(int)
		ret.Self_grat_nat_ip_arp_drop = in["self_grat_nat_ip_arp_drop"].(int)
		ret.Ip_not_found_arp_drop = in["ip_not_found_arp_drop"].(int)
		ret.Dev_link_down_arp_drop = in["dev_link_down_arp_drop"].(int)
		ret.Lacp_tx_intf_err_drop = in["lacp_tx_intf_err_drop"].(int)
		ret.Service_chain_sent = in["service_chain_sent"].(int)
		ret.Service_chain_rcvd = in["service_chain_rcvd"].(int)
		ret.Unnumbered_nat_error = in["unnumbered_nat_error"].(int)
		ret.Unnumbered_unsupported_drop = in["unnumbered_unsupported_drop"].(int)
		ret.Ipv6frag_gre_dropped = in["ipv6frag_gre_dropped"].(int)
		ret.Ipv6_ndisc_dad_prefix_mismatch_drop = in["ipv6_ndisc_dad_prefix_mismatch_drop"].(int)
		ret.Bw_ignore_limit = in["bw_ignore_limit"].(int)
		ret.Ppsl_drop_egr = in["ppsl_drop_egr"].(int)
		ret.Ppsl_drop_ing = in["ppsl_drop_ing"].(int)
		ret.Ppsl_ignore_limit = in["ppsl_ignore_limit"].(int)
		ret.Closed_port_syn_drop = in["closed_port_syn_drop"].(int)
		ret.Tls13_drop_req = in["tls13_drop_req"].(int)
		ret.Tls13_ignore_req = in["tls13_ignore_req"].(int)
		ret.Tls12_drop_req = in["tls12_drop_req"].(int)
		ret.Tls12_ignore_req = in["tls12_ignore_req"].(int)
		ret.Tls12_tls13_drop_req = in["tls12_tls13_drop_req"].(int)
		ret.Tls12_tls13_ignore_req = in["tls12_tls13_ignore_req"].(int)
	}
	return ret
}

func dataToEndpointSlbSwitchStats(d *schema.ResourceData) edpt.SlbSwitchStats {
	var ret edpt.SlbSwitchStats

	ret.Stats = getObjectSlbSwitchStatsStats(d.Get("stats").([]interface{}))
	return ret
}
