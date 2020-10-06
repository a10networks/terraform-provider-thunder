---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_dns_secondary"
sidebar_current: "docs-vthunder-resource-ip-dns-secondary"
description: |-
	Provides details about vthunder ip dns secondary resource for A10
---

# vthunder\_ip\_dns\_secondary

`vthunder_ip_dns_secondary` Provides details about vthunder ip dns secondary
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_dns_secondary" "dnssecondary" {
    ip_v4_addr = "10.10.10.3"
}
```

## Argument Reference

* `ip_v4_addr` - DNS server address
* `ip_v6_addr` - DNS server address
* `uuid` - uuid of the object

