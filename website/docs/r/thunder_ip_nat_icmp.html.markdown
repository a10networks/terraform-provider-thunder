---
layout: "thunder"
page_title: "thunder: thunder_ip_nat_icmp"
sidebar_current: "docs-thunder-resource-ip-nat-icmp"
description: |-
  Provides details about thunder ip nat icmp resource for A10
---

# thunder\_ip\_nat\_icmp

`thunder_ip_nat_icmp` Provides details about thunder ip nat icmp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_nat_icmp" "NatIcmp" {
  respond_to_ping = 0
  always_source_nat_errors = 0 
}
```

## Argument Reference

* `always_source_nat_errors` - Source NAT intermediate routers’ IPs for ICMP errors (default: disabled)
* `respond_to_ping` - Respond to ICMP echo requests to NAT pool IPs (default: disabled)
* `uuid` - uuid of the object

