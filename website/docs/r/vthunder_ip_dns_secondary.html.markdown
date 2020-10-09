---
layout: "thunder"
page_title: "thunder: thunder_ip_dns_secondary"
sidebar_current: "docs-thunder-resource-ip-dns-secondary"
description: |-
	Provides details about thunder ip dns secondary resource for A10
---

# thunder\_ip\_dns\_secondary

`thunder_ip_dns_secondary` Provides details about thunder ip dns secondary
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_dns_secondary" "dnssecondary" {
    ip_v4_addr = "10.10.10.3"
}
```

## Argument Reference

* `ip_v4_addr` - DNS server address
* `ip_v6_addr` - DNS server address
* `uuid` - uuid of the object

