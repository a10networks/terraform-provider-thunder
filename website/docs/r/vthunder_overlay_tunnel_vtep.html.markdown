---
layout: "vthunder"
page_title: "vthunder: vthunder_overlay_tunnel_vtep"
sidebar_current: "docs-vthunder-resource-overlay_tunnel_vtep"
description: |-
    Provides details about vthunder overlay tunnel vtep resource for A10
---

# vthunder\_overlay\_tunnel\_vtep

`vthunder_overlay_tunnel_options` Provides details about vthunder overlay tunnel vtep
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_overlay_tunnel_vtep" "vtep"{
    source_ip_address{
        ip_address="1.2.3.4"
        vni_list{
            segment=1
            partition="part7"
            lif=1
    }
    }
    encap="nvgre"
    host_list{
    destination_vtep="1.4.3.2"
    ip_addr="1.4.3.5"
    overlay_mac_addr="00:1B:44:11:3A:B7"
    vni=1
    }
    id2=3
    destination_ip_address_list{
    ip_address="2.3.4.5"
    vni_list{
    segment=1
    }
    encap="nvgre"
}
```

## Argument Reference

* `source_ip_address`
  * `ip_address` - Source Tunnel End Point IPv4 address
  * `vni_list`
    * `segment` - VNI configured for the remote VTEP
    * `partition` - Logical interface binding the Provider Partition to the User Partition (logical interface number)
    * `lif` - Name of the Partition with the L2 segment being extended (Name of the User Partition with the L2 segment being extended)
* `encap` - 'nvgre’: Tunnel Encapsulation Type is NVGRE; 'vxlan’: Tunnel Encapsulation Type is VXLAN;
* `host_list`
    * `ip_addr` - uuid of the object
    * `destination_vtep` - Configure the VTEP IP address (IPv4 address of the VTEP for the remote host)
    * `overlay_mac_addr` - MAC Address of the overlay host
    * `vni` - Configure the segment id ( VNI of the remote host)
* `id2` - Disable TCP MSS adjustment in SYN packet for tunnels
* `destination_ip_address_list`
    * `ip_address` - Disable Flow-ID computation for NVGRE
    * `destination_ip_address_list`
        * `ip_address` - IPv4 address of the overlay host
        * `vni_list`
            * `segment` - Configure the segment id ( VNI of the remote host)


