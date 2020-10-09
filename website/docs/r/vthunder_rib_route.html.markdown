---
layout: "thunder"
page_title: "thunder: thunder_rib_route"
sidebar_current: "docs-thunder-resource-rib-route"
description: |-
    Provides details about thunder rib route resource for A10
---

# thunder\_rib\_route

`thunder_rib_route` provides details about configuring rib route on a device
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_rib_route" "rib"{
  ip_dest_addr="0.0.0.0"
  ip_mask="/0"
  ip_nexthop_ipv4{
    ip_next_hop="10.0.2.1"
    distance_nexthop_ip=1
  }
}
```

## Argument Reference

* `ip_dest_addr` - Destination ip address
* `ip_mask` - Ip address subnet mask
* `ip_next_hop` - Ip next hop
* `distance_nexthop_ip` - Distance of next hop ip address
