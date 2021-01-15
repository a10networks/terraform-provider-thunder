---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_neighbor_ethernet_neighbor"
sidebar_current: "docs-thunder-resource-router-bgp-neighbor-ethernet-neighbor"
description: |-
    Provides details about thunder router bgp neighbor ethernet neighbor resource for A10
---

# thunder\_router\_bgp\_neighbor\_ethernet\_neighbor

`thunder_router_bgp_neighbor_ethernet_neighbor` Provides details about thunder router bgp neighbor ethernet neighbor
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_neighbor_ethernet_neighbor" "resourceRouterBgpNeighborEthernetNeighborTest" {
	ethernet = 0
unnumbered = 0
peer_group_name = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ethernet-neighbor` - Specify an ethernet unnumbered neighbor
* `ethernet` - Ethernet interface number
* `uuid` - uuid of the object

