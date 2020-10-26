---
layout: "thunder"
page_title: "thunder: thunder_dns_primary"
sidebar_current: "docs-thunder-resource-dns-primary"
description: |-
    Provides details about thunder dns primary resource for A10
---

# thunder\_dns\_primary

`thunder_dns_primary provides` details about Dns Primary
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_dns_primary" "dns_primary" {
		ip_v4_addr = "8.8.8.8"
}
```

## Argument Reference

* `ip_v4_addr` - IP v4 address
