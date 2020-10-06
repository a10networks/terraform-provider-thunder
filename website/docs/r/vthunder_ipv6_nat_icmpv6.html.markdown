---
layout: "vthunder"
page_title: "vthunder: vthunder_ipv6_nat_icmpv6"
sidebar_current: "docs-vthunder-resource-ipv6-nat-icmpv6"
description: |-
  Provides details about vthunder ipv6 nat icmpv6 resource for A10
---

# vthunder\_ipv6\_nat\_icmpv6

`vthunder_ipv6_nat_icmpv6` Provides details about vthunder ipv6 nat icmpv6
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ipv6_nat_icmpv6" "natIcmpV6" {
  respond_to_ping = 0
}
```

## Argument Reference

* `respond_to_ping` - Respond to ICMPv6 echo requests to NAT pool IPs (default: disabled)
* `uuid` - uuid of the object

