---
layout: "thunder"
page_title: "thunder: thunder_ipv6_icmpv6"
sidebar_current: "docs-thunder-resource-ipv6-icmpv6"
description: |-
  Provides details about thunder ipv6 icmpv6 resource for A10
---

# thunder\_ipv6\_icmpv6

`thunder_ipv6_icmpv6` Provides details about thunder ipv6 icmpv6
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ipv6_icmpv6" "testname" {
    redirect = 0
  unreachable = 0
}
```

## Argument Reference

* `redirect` - Disable outbound ICMPv6 redirect messages
* `unreachable` - Disable outbound ICMPv6 unreachable messages
* `uuid` - uuid of the object
