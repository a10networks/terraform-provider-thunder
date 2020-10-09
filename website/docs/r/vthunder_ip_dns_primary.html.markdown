---
layout: "thunder"
page_title: "thunder: thunder_ip_dns_primary"
sidebar_current: "docs-thunder-resource-ip-dns-primary"
description: |-
	Provides details about thunder ip dns primary resource for A10
---

# thunder\_ip\_dns\_primary

`thunder_ip_dns_primary` Provides details about thunder ip dns primary
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_dns_primary" "dnsPrimary" {
    ip_v4_addr = "10.10.10.3"
}
```

## Argument Reference

* `ip_v4_addr` - DNS server address
* `ip_v6_addr` - DNS server address
* `uuid` - uuid of the object

