---
layout: "vthunder"
page_title: "vthunder: vthunder_ipv6_icmpv6"
sidebar_current: "docs-vthunder-resource-ipv6-icmpv6"
description: |-
  Provides details about vthunder ipv6 icmpv6 resource for A10
---

# vthunder\_ipv6\_icmpv6

`vthunder_ipv6_icmpv6` Provides details about vthunder ipv6 icmpv6
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ipv6_icmpv6" "testname" {
    redirect = 0
  unreachable = 0
}
```

## Argument Reference

* `redirect` - Disable outbound ICMPv6 redirect messages
* `unreachable` - Disable outbound ICMPv6 unreachable messages
* `uuid` - uuid of the object
