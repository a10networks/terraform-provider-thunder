---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_network_ip_cidr"
sidebar_current: "docs-thunder-resource-router-bgp-network-ip-cidr"
description: |-
    Provides details about thunder router bgp network ip cidr resource for A10
---

# thunder\_router\_bgp\_network\_ip\_cidr

`thunder_router_bgp_network_ip_cidr` Provides details about thunder router bgp network ip cidr
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_network_ip_cidr" "resourceRouterBgpNetworkIpCidrTest" {
	network_ipv4_cidr = "string"
route_map = "string"
backdoor = 0
description = "string"
comm_value = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ip-cidr` - Specify a ip network to announce via BGP
* `network-ipv4-cidr` - Specify network mask
* `route-map` - Route-map to modify the attributes (Name of the route map)
* `backdoor` - Specify a BGP backdoor route
* `description` - Network specific description (Up to 80 characters describing this network)
* `comm-value` - community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export
* `uuid` - uuid of the object

