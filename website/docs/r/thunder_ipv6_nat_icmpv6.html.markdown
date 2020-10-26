---
layout: "thunder"
page_title: "thunder: thunder_ipv6_nat_icmpv6"
sidebar_current: "docs-thunder-resource-ipv6-nat-icmpv6"
description: |-
  Provides details about thunder ipv6 nat icmpv6 resource for A10
---

# thunder\_ipv6\_nat\_icmpv6

`thunder_ipv6_nat_icmpv6` Provides details about thunder ipv6 nat icmpv6
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ipv6_nat_icmpv6" "natIcmpV6" {
  respond_to_ping = 0
}
```

## Argument Reference

* `respond_to_ping` - Respond to ICMPv6 echo requests to NAT pool IPs (default: disabled)
* `uuid` - uuid of the object

