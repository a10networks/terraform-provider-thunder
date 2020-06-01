---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_icmp"
sidebar_current: "docs-vthunder-resource-ip-icmp"
description: |-
  Provides details about vthunder ip icmp resource for A10
---

# vthunder\_ip\_icmp

`vthunder_ip_icmp` Provides details about vthunder ip icmp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_icmp" "ICMP" {
  redirect = 0
  unreachable = 0
}
```

## Argument Reference

* `redirect` - Disable outbound ICMP redirect messages
* `unreachable` - Disable outbound ICMP unreachable messages
* `uuid` - uuid of the object

