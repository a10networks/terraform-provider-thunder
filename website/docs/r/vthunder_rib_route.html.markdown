---
layout: "vthunder"
page_title: "vthunder: vthunder_rib_route"
sidebar_current: "docs-vthunder-resource-rib-route"
description: |-
    Provides details about vthunder rib route resource for A10
---

# vthunder\_rib\_route

`vthunder_rib_route` provides details about configuring rib route on a device
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_rib_route" "rib"{
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
