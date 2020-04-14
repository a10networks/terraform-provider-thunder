---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_switch"
sidebar_current: "docs-vthunder-slb-switch"
description: |-
    Provides details about vthunder SLB switch resource for A10
---

# vthunder\_slb\_switch

`vthunder_slb_switch` Provides details about vthunder SLB switch
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_switch" "switch" {
	sampling_enable  {
	   counters1 = all
	} 
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'fwlb’: FWLB; 'licexpire_drop’: License Expire Drop; 'bwl_drop’: BW Limit Drop; 'rx_kernel’: Received kernel; 'rx_arp_req’: ARP REQ Rcvd; 'rx_arp_resp’: ARP RESP Rcvd; 'vlan_flood’: VLAN Flood; 'l2_def_vlan_drop’: L2 Default Vlan FWD Drop; 'ipv4_noroute_drop’: IPv4 No Route Drop; 'ipv6_noroute_drop’: IPv6 No Route Drop; 'prot_down_drop’: Prot Down Drop; 'l2_forward’: L2 Forward; 'l3_forward_ip’: L3 IP Forward; 'l3_forward_ipv6’: L3 IPv6 Forward; 'l4_process’: L4 Process; 'unknown_prot_drop’: Unknown Prot Drop; 'ttl_exceeded_drop’: TTL Exceeded Drop; 'linkdown_drop’: Link Down Drop; 'sport_drop’: SPORT Drop; 'incorrect_len_drop’: Incorrect Length Drop; 'ip_defrag’: IP Defrag; 'acl_deny’: ACL Denys; 'ipfrag_tcp’: IP(TCP) Fragment Rcvd; 'ipfrag_overlap’: IP Fragment Overlap; 'ipfrag_timeout’: IP Fragment Timeout; 'ipfrag_overload’: IP Frag Overload Drops; 'ipfrag_reasmoks’: IP Fragment Reasm OKs; 'ipfrag_reasmfails’: IP Fragment Reasm Fails; 'land_drop’: Anomaly Land Attack Drop; 'ipoptions_drop’: Anomaly IP OPT Drops; 'badpkt_drop’: Bad Pkt Drop; 'pingofdeath_drop’: Anomaly PingDeath Drop; 'allfrag_drop’: Anomaly All Frag Drop; 'tcpnoflag_drop’: Anomaly TCP noFlag Drop; 'tcpsynfrag_drop’: Anomaly SYN Frag Drop; 'tcpsynfin_drop’: Anomaly TCP SYNFIN Drop; 'ipsec_drop’: IPSec Drop; 'bpdu_rcvd’: BPDUs Received; 'bpdu_sent’: BPDUs Sent; 'ctrl_syn_rate_drop’: SYN rate exceeded Drop; 'ip_defrag_invalid_len’: IP Invalid Length Frag; 'ipv4_frag_6rd_ok’: IPv4 Frag 6RD OK; 'ipv4_frag_6rd_drop’: IPv4 Frag 6RD Dropped; 'no_ip_drop’: No IP Drop; 'ipv6frag_udp’: IPv6 Frag UDP; 'ipv6frag_udp_dropped’: IPv6 Frag UDP Dropped; 'ipv6frag_tcp_dropped’: IPv6 Frag TCP Dropped; 'ipv6frag_ipip_ok’: IPv6 Frag IPIP OKs; 'ipv6frag_ipip_dropped’: IPv6 Frag IPIP Drop; 'ip_frag_oversize’: IP Fragment oversize; 'ip_frag_too_many’: IP Fragment too many; 'ipv4_novlanfwd_drop’: IPv4 No L3 VLAN FWD Drop; 'ipv6_novlanfwd_drop’: IPv6 No L3 VLAN FWD Drop; 'fpga_error_pkt1’: FPGA Error PKT1; 'fpga_error_pkt2’: FPGA Error PKT2; 'max_arp_drop’: Max ARP Drop; 'ipv6frag_tcp’: IPv6 Frag TCP; 'ipv6frag_icmp’: IPv6 Frag ICMP; 'ipv6frag_ospf’: IPv6 Frag OSPF; 'ipv6frag_esp’: IPv6 Frag ESP; 'l4_in_ctrl_cpu’: L4 In Ctrl CPU; 'mgmt_svc_drop’: Management Service Drop; 'jumbo_frag_drop’: Jumbo Frag Drop; 'ipv6_jumbo_frag_drop’: IPv6 Jumbo Frag Drop; 'ipipv6_jumbo_frag_drop’: IPIPv6 Jumbo Frag Drop; 'ipv6_ndisc_dad_solicits’: IPv6 DAD on Solicits; 'ipv6_ndisc_dad_adverts’: IPv6 DAD on Adverts; 'ipv6_ndisc_mac_changes’: IPv6 DAD MAC Changed; 'ipv6_ndisc_out_of_memory’: IPv6 DAD Out-of-memory; 'sp_non_ctrl_pkt_drop’: Shared IP mode non ctrl packet to linux drop; 'urpf_pkt_drop’: URPF check packet drop; 'fw_smp_zone_mismatch’: FW SMP Zone Mismatch; 'ipfrag_udp’: IP(UDP) Fragment Rcvd; 'ipfrag_icmp’: IP(ICMP) Fragment Rcvd; 'ipfrag_ospf’: IP(OSPF) Fragment Rcvd; 'ipfrag_esp’: IP(ESP) Fragment Rcvd; 'ipfrag_tcp_dropped’: IP Frag TCP Dropped; 'ipfrag_udp_dropped’: IP Frag UDP Dropped; 'ipfrag_ipip_dropped’: IP Frag IPIP Drop; 'redirect_fwd_fail’: Redirect failed in the fwd direction; 'redirect_fwd_sent’: Redirect succeeded in the fwd direction; 'redirect_rev_fail’: Redirect failed in the rev direction; 'redirect_rev_sent’: Redirect succeeded in the rev direction; 'redirect_setup_fail’: Redirect connection setup failed; 'ip_frag_sent’: IP frag sent; 'invalid_rx_arp_pkt’: Invalid ARP PKT Rcvd;





