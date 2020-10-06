---
layout: "vthunder"
page_title: "vthunder: vthunder_dns_primary"
sidebar_current: "docs-vthunder-resource-dns-primary"
description: |-
    Provides details about vthunder dns primary resource for A10
---

# vthunder\_dns\_primary

`vthunder_dns_primary provides` details about Dns Primary
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_dns_primary" "dns_primary" {
		ip_v4_addr = "8.8.8.8"
}
```

## Argument Reference

* `ip_v4_addr` - IP v4 address
