---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_neighbor_ve_neighbor"
sidebar_current: "docs-thunder-resource-router-bgp-neighbor-ve-neighbor"
description: |-
    Provides details about thunder router bgp neighbor ve neighbor resource for A10
---

# thunder\_router\_bgp\_neighbor\_ve\_neighbor

`thunder_router_bgp_neighbor_ve_neighbor` Provides details about thunder router bgp neighbor ve neighbor
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_neighbor_ve_neighbor" "resourceRouterBgpNeighborVeNeighborTest" {
	ve = 0
unnumbered = 0
peer_group_name = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ve-neighbor` - Specify a VE unnumbered neighbor
* `ve` - Virtual ethernet interface number
* `uuid` - uuid of the object

