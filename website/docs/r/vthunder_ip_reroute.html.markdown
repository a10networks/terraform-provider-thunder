---
layout: "thunder"
page_title: "thunder: thunder_ip_reroute"
sidebar_current: "docs-thunder-resource-ip-reroute"
description: |-
  Provides details about thunder ip reroute resource for A10
---

# thunder\_ip\_reroute

`thunder_ip_reroute` Provides details about thunder ip reroute
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_reroute" "reRoute" {
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

