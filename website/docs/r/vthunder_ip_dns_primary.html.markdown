---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_dns_primary"
sidebar_current: "docs-vthunder-resource-ip-dns-primary"
description: |-
	Provides details about vthunder ip dns primary resource for A10
---

# vthunder\_ip\_dns\_primary

`vthunder_ip_dns_primary` Provides details about vthunder ip dns primary
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_dns_primary" "dnsPrimary" {
    ip_v4_addr = "10.10.10.3"
}
```

## Argument Reference

* `ip_v4_addr` - DNS server address
* `ip_v6_addr` - DNS server address
* `uuid` - uuid of the object

