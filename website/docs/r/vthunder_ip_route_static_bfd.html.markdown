---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_route_static_bfd"
sidebar_current: "docs-vthunder-resource-ip-route-static-bfd"
description: |-
	Provides details about vthunder ip route static bfd resource for A10
---

# vthunder\_ip\_route\_static\_bfd

`vthunder_ip_route_static_bfd` Provides details about vthunder ip route static bfd
## Example Usage

```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_route_static_bfd" "ipStaticBFD" {
  local_ip="3.3.3.3"
  nexthop_ip="3.3.3.3"
}
```

## Argument Reference

* `local_ip` - Local IP address
* `nexthop_ip` - Nexthop IP address
* `uuid` - uuid of the object
