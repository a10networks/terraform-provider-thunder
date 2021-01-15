---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_address_family_ipv6_network_synchronization"
sidebar_current: "docs-thunder-resource-router-bgp-address-family-ipv6-network-synchronization"
description: |-
    Provides details about thunder router bgp address family ipv6 network synchronization resource for A10
---

# thunder\_router\_bgp\_address\_family\_ipv6\_network\_synchronization

`thunder_router_bgp_address_family_ipv6_network_synchronization` Provides details about thunder router bgp address family ipv6 network synchronization
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_address_family_ipv6_network_synchronization" "resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationTest" {
	network_synchronization = 0
uuid = "string"
 
}

```

## Argument Reference

* `synchronization` - help Perform IGP synchronization
* `network-synchronization` - Perform IGP synchronization
* `uuid` - uuid of the object

