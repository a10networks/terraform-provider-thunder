---
layout: "thunder"
page_title: "thunder: thunder_overlay_tunnel_vtep"
sidebar_current: "docs-thunder-resource-overlay-tunnel-vtep"
description: |-
    Provides details about thunder overlay tunnel vtep resource for A10
---

# thunder\_overlay\_tunnel\_vtep

`thunder_overlay_tunnel_vtep` Provides details about thunder overlay tunnel vtep
## Example Usage


```hcl
provider "thunder" {
    address  = var.address
    username = var.username  
    password = var.password
}


resource "thunder_overlay_tunnel_vtep" "resourceOverlayTunnelVtepTest" {
id1 = 0
encap = "string"
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
local_ip_address {  
 	ip_address =  "string" 
	uuid =  "string" 
vni_list {   
	segment =  0 
	uuid =  "string" 
	}
	}
local_ipv6_address {  
 	ipv6_address =  "string" 
	uuid =  "string" 
	}
remote_ip_address-list {   
	ip_address =  "string" 
	encap =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
use_lif {  
 	partition =  "string" 
	lif =  "string" 
	uuid =  "string" 
	}
gre_keepalive {  
 	retry_time =  0 
	retry_count =  0 
	uuid =  "string" 
	}
use_gre_key {  
 	gre_key =  0 
	uuid =  "string" 
	}
vni-list {   
	segment =  0 
	uuid =  "string" 
	}
	}
remote-ipv6-address-list {   
	ipv6_address =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
use_lif {  
 	partition =  "string" 
	lif =  "string" 
	uuid =  "string" 
	}
	}
host-list {   
	ip_addr =  "string" 
	overlay_mac_addr =  "string" 
	vni =  0 
	remote_vtep =  "string" 
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `vtep` - Virtual Tunnel end point Configuration
* `id1` - VTEP Identifier
* `encap` - 'nvgre': Tunnel Encapsulation Type is NVGRE; 'vxlan': Tunnel Encapsulation Type is VXLAN;
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'cfg_err_count': Config errors; 'flooded_pkt_count': Flooded packet count; 'encap_unresolved_count': Encap unresolved failures; 'unknown_encap_rx_pkt': Encap miss rx pkts; 'unknown_encap_tx_pkt': Encap miss tx pkts; 'arp_req_sent': Arp request sent; 'vtep_host_learned': Hosts learned; 'vtep_host_learn_error': Host learn error; 'invalid_lif_rx': Invalid Lif pkts in; 'invalid_lif_tx': Invalid Lif pkts out; 'unknown_vtep_tx': Vtep unknown tx; 'unknown_vtep_rx': Vtep Unkown rx; 'unhandled_pkt_rx': Unhandled packets in; 'unhandled_pkt_tx': Unhandled packets out; 'total_pkts_rx': Total packets out; 'total_bytes_rx': Total packet bytes in; 'unicast_pkt_rx': Total unicast packets in; 'bcast_pkt_rx': Total broadcast packets in; 'mcast_pkt_rx': Total multicast packets in; 'dropped_pkt_rx': Dropped received packets; 'encap_miss_pkts_rx': Encap missed in received packets; 'bad_chksum_pks_rx': Bad checksum in received packets; 'requeue_pkts_in': Requeued packets in; 'pkts_out': Packets out; 'total_bytes_tx': Packet bytes out; 'unicast_pkt_tx': Unicast packets out; 'bcast_pkt_tx': Broadcast packets out; 'mcast_pkt_tx': Multicast packets out; 'dropped_pkts_tx': Dropped packets out; 'large_pkts_rx': Too large packets in; 'dot1q_pkts_rx': Dot1q packets in; 'frag_pkts_tx': Frag packets out; 'reassembled_pkts_rx': Reassembled packets in; 'bad_inner_ipv4_len_rx': bad inner ipv4 packet len; 'bad_inner_ipv6_len_rx': Bad inner ipv6 packet len; 'frag_drop_pkts_tx': Frag dropped packets out; 'lif_un_init_rx': Lif uninitialized packets in;
* `ip-address` - IP Address of the remote VTEP
* `segment` - VNI configured for the remote VTEP
* `ipv6-address` - IPv6 Address of the remote VTEP
* `partition` - Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)
* `lif` - Logical interface binding the Provider Partition to the User Partition (logical interface name)
* `retry-time` - Keepalive retry interval in seconds
* `retry-count` - Keepalive multiplier
* `gre-key` - key
* `ip-addr` - IPv4 address of the overlay host
* `overlay-mac-addr` - MAC Address of the overlay host
* `vni` - Configure the segment id ( VNI of the remote host)
* `remote-vtep` - Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)

