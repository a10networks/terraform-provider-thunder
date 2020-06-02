---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_nat_icmp"
sidebar_current: "docs-vthunder-resource-ip-nat-icmp"
description: |-
  Provides details about vthunder ip nat icmp resource for A10
---

# vthunder\_ip\_nat\_icmp

`vthunder_ip_nat_icmp` Provides details about vthunder ip nat icmp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_nat_icmp" "NatIcmp" {
  respond_to_ping = 0
  always_source_nat_errors = 0 
}
```

## Argument Reference

* `always_source_nat_errors` - Source NAT intermediate routers’ IPs for ICMP errors (default: disabled)
* `respond_to_ping` - Respond to ICMP echo requests to NAT pool IPs (default: disabled)
* `uuid` - uuid of the object

