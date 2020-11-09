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
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_overlay_tunnel_vtep" "Overlay_Tunnel_Vtep_Test" {

uuid = "string"
user_tag = "string"
sampling-enable {   
        counters1 =  "string" 
        }
source_ip_address {  
        ip_address =  "string" 
        uuid =  "string" 
vni-list {   
        segment =  0 
        uuid =  "string" 
        }
        }
encap = "string"
host-list {   
        destination_vtep =  "string" 
        ip_addr =  "string" 
        overlay_mac_addr =  "string" 
        vni =  0 
        uuid =  "string" 
        }
id = 0
destination-ip-address-list {   
        uuid =  "string" 
        ip_address =  "string" 
vni-list {   
        segment =  0 
        uuid =  "string" 
        }
        user_tag =  "string" 
        encap =  "string" 
        }
 
}

```

## Argument Reference

* `encap` - ‘nvgre’: Tunnel Encapsulation Type is NVGRE; ‘vxlan’: Tunnel Encapsulation Type is VXLAN;
* `id` - VTEP Identifier
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘cfg_err_count’: Config errors; ‘flooded_pkt_count’: Flooded packet count; ‘encap_unresolved_count’: Encap unresolved failures; ‘unknown_encap_rx_pkt’: Encap miss rx pkts; ‘unknown_encap_tx_pkt’: Encap miss tx pkts; ‘arp_req_sent’: Arp request sent; ‘vtep_host_learned’: Hosts learned; ‘vtep_host_learn_error’: Host learn error; ‘invalid_lif_rx’: Invalid Lif pkts in; ‘invalid_lif_tx’: Invalid Lif pkts out; ‘unknown_vtep_tx’: Vtep unknown tx; ‘unknown_vtep_rx’: Vtep Unkown rx; ‘unhandled_pkt_rx’: Unhandled packets in; ‘unhandled_pkt_tx’: Unhandled packets out; ‘total_pkts_rx’: Total packets out; ‘total_bytes_rx’: Total packet bytes in; ‘unicast_pkt_rx’: Total unicast packets in; ‘bcast_pkt_rx’: Total broadcast packets in; ‘mcast_pkt_rx’: Total multicast packets in; ‘dropped_pkt_rx’: Dropped received packets; ‘encap_miss_pkts_rx’: Encap missed in received packets; ‘bad_chksum_pks_rx’: Bad checksum in received packets; ‘requeue_pkts_in’: Requeued packets in; ‘pkts_out’: Packets out; ‘total_bytes_tx’: Packet bytes out; ‘unicast_pkt_tx’: Unicast packets out; ‘bcast_pkt_tx’: Broadcast packets out; ‘mcast_pkt_tx’: Multicast packets out; ‘dropped_pkts_tx’: Dropped packets out; ‘large_pkts_rx’: Too large packets in; ‘dot1q_pkts_rx’: Dot1q packets in; ‘frag_pkts_tx’: Frag packets out; ‘reassembled_pkts_rx’: Reassembled packets in; ‘bad_inner_ipv4_len_rx’: bad inner ipv4 packet len; ‘bad_inner_ipv6_len_rx’: Bad inner ipv6 packet len; ‘lif_un_init_rx’: Lif uninitialized packets in;
* `ip_address` - IP Address of the remote VTEP
* `gateway` - This is a Gateway segment id
* `lif` - Logical interface binding the Provider Partition to the User Partition (logical interface number)
* `partition` - Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)
* `segment` - VNI configured for the remote VTEP
* `destination_vtep` - Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)
* `ip_addr` - IPv4 address of the overlay host
* `overlay_mac_addr` - MAC Address of the overlay host
* `vni` -  Configure the segment id ( VNI of the remote host)
