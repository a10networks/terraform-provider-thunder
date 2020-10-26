---
layout: "thunder"
page_title: "thunder: thunder_ip_icmp"
sidebar_current: "docs-thunder-resource-ip-icmp"
description: |-
  Provides details about thunder ip icmp resource for A10
---

# thunder\_ip\_icmp

`thunder_ip_icmp` Provides details about thunder ip icmp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_icmp" "ICMP" {
  redirect = 0
  unreachable = 0
}
```

## Argument Reference

* `redirect` - Disable outbound ICMP redirect messages
* `unreachable` - Disable outbound ICMP unreachable messages
* `uuid` - uuid of the object

