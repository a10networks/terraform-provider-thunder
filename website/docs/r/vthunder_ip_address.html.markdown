---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_address"
sidebar_current: "docs-vthunder-resource-ip-address"
description: |-
  Provides details about vthunder ip address resource for A10
---

# vthunder\_ip\_address

`vthunder_ip_address` Provides details about vthunder ip address
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_address" "testname" {
  ip_addr = "3.3.3.3"
  ip_mask = "255.255.0.0"
}
```

## Argument Reference

* `ip_addr` - IP address
* `ip_mask` - IP subnet mask
* `uuid` - uuid of the object

