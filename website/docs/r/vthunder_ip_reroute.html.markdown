---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_reroute"
sidebar_current: "docs-vthunder-resource-ip-reroute"
description: |-
  Provides details about vthunder ip reroute resource for A10
---

# vthunder\_ip\_reroute

`vthunder_ip_reroute` Provides details about vthunder ip reroute
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_reroute" "reRoute" {
  suppress_protocols  {
        static = 0
        ospf = 0
        connected = 0
        ibgp = 0
        isis = 0
        rip = 0
        ebgp = 0
  }
}
```

## Argument Reference

* `uuid` - uuid of the object
* `ebgp` - EBGP
* `ibgp` - IBGP
* `isis` - ISIS
* `ospf` - OSPF
* `rip` - RIP

