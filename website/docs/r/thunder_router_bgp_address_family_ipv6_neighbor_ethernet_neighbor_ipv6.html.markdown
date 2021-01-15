---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6"
sidebar_current: "docs-thunder-resource-router-bgp-address-family-ipv6-neighbor-ethernet-neighbor-ipv6"
description: |-
    Provides details about thunder router bgp address family ipv6 neighbor ethernet neighbor ipv6 resource for A10
---

# thunder\_router\_bgp\_address\_family\_ipv6\_neighbor\_ethernet\_neighbor\_ipv6

`thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6` Provides details about thunder router bgp address family ipv6 neighbor ethernet neighbor ipv6
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_address_family_ipv6_neighbor_ethernet_neighbor_ipv6" "resourceRouterBgpAddressFamilyIpv6NeighborEthernetNeighborIpv6Test" {
	ethernet = 0
peer_group_name = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ethernet-neighbor-ipv6` - Specify an ethernet unnumbered neighbor
* `ethernet` - Ethernet interface number
* `uuid` - uuid of the object

