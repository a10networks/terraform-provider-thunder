---
layout: "thunder"
page_title: "thunder: thunder_ip_nat_global"
sidebar_current: "docs-thunder-resource-ip-nat-global"
description: |-
  Provides details about thunder ip nat global resource for A10
---

# thunder\_ip\_nat\_global

`thunder_ip_nat_global` Provides details about thunder ip nat global
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_nat_global" "IpNatGlobal" {
  reset_idle_tcp_conn = 0
}
```

## Argument Reference

* `reset_idle_tcp_conn` - Reset Idle TCP Connections
* `uuid` - uuid of the object

