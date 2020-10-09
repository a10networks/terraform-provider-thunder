---
layout: "thunder"
page_title: "thunder: thunder_ip_address"
sidebar_current: "docs-thunder-resource-ip-address"
description: |-
  Provides details about thunder ip address resource for A10
---

# thunder\_ip\_address

`thunder_ip_address` Provides details about thunder ip address
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_address" "testname" {
  ip_addr = "3.3.3.3"
  ip_mask = "255.255.0.0"
}
```

## Argument Reference

* `ip_addr` - IP address
* `ip_mask` - IP subnet mask
* `uuid` - uuid of the object

