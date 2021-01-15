---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_address_family_ipv6_network_ipv6_network"
sidebar_current: "docs-thunder-resource-router-bgp-address-family-ipv6-network-ipv6-network"
description: |-
    Provides details about thunder router bgp address family ipv6 network ipv6 network resource for A10
---

# thunder\_router\_bgp\_address\_family\_ipv6\_network\_ipv6\_network

`thunder_router_bgp_address_family_ipv6_network_ipv6_network` Provides details about thunder router bgp address family ipv6 network ipv6 network
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_address_family_ipv6_network_ipv6_network" "resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkTest" {
	network_ipv6 = "string"
route_map = "string"
backdoor = 0
description = "string"
comm_value = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ipv6-network` - Specify a ip address mask network to announce via BGP
* `network-ipv6` - Specify a network to announce via BGP
* `route-map` - Route-map to modify the attributes (Name of the route map)
* `backdoor` - Specify a BGP backdoor route
* `description` - Network specific description (Up to 80 characters describing this network)
* `comm-value` - community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export
* `uuid` - uuid of the object

