---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_nat_global"
sidebar_current: "docs-vthunder-resource-ip-nat-global"
description: |-
  Provides details about vthunder ip nat global resource for A10
---

# vthunder\_ip\_nat\_global

`vthunder_ip_nat_global` Provides details about vthunder ip nat global
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_nat_global" "IpNatGlobal" {
  reset_idle_tcp_conn = 0
}
```

## Argument Reference

* `reset_idle_tcp_conn` - Reset Idle TCP Connections
* `uuid` - uuid of the object

