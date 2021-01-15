---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_network_synchronization"
sidebar_current: "docs-thunder-resource-router-bgp-network-synchronization"
description: |-
    Provides details about thunder router bgp network synchronization resource for A10
---

# thunder\_router\_bgp\_network\_synchronization

`thunder_router_bgp_network_synchronization` Provides details about thunder router bgp network synchronization
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_network_synchronization" "resourceRouterBgpNetworkSynchronizationTest" {
	network_synchronization = 0
uuid = "string"
 
}

```

## Argument Reference

* `synchronization` - help Perform IGP synchronization
* `network-synchronization` - Perform IGP synchronization
* `uuid` - uuid of the object

