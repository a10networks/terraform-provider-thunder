---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_neighbor_trunk_neighbor"
sidebar_current: "docs-thunder-resource-router-bgp-neighbor-trunk-neighbor"
description: |-
    Provides details about thunder router bgp neighbor trunk neighbor resource for A10
---

# thunder\_router\_bgp\_neighbor\_trunk\_neighbor

`thunder_router_bgp_neighbor_trunk_neighbor` Provides details about thunder router bgp neighbor trunk neighbor
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_neighbor_trunk_neighbor" "resourceRouterBgpNeighborTrunkNeighborTest" {
	trunk = 0
unnumbered = 0
peer_group_name = "string"
uuid = "string"
 
}

```

## Argument Reference

* `trunk-neighbor` - Specify a trunk unnumbered neighbor
* `trunk` - Trunk interface number
* `uuid` - uuid of the object

